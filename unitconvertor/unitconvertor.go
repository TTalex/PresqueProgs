package unitconvertor

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Factor struct {
	before float64
	factor float64
	after  float64
	how    string
}

type FactorRandomizer struct {
	Factors []Factor
}

func initFactors(fr FactorRandomizer) {
	fr.Factors = []Factor{}
}
func getFactor(fr FactorRandomizer) Factor {
	if len(fr.Factors) == 0 {
		return Factor{}
	}
	return fr.Factors[rand.Intn(len(fr.Factors))]
}
func addFactor(fr FactorRandomizer, f Factor) FactorRandomizer {
	fr.Factors = append(fr.Factors, f)
	return fr
}

func addConvertor(maindata map[string]map[string]FactorRandomizer, fromunit string, tounit string, factor Factor) {
	if maindata[fromunit] == nil {
		maindata[fromunit] = make(map[string]FactorRandomizer)
	}
	if maindata[fromunit][tounit].Factors == nil {
		initFactors(maindata[fromunit][tounit])
	}
	maindata[fromunit][tounit] = addFactor(maindata[fromunit][tounit], factor)
}
func addConvertors(maindata map[string]map[string]FactorRandomizer, fromunit string, tounit string, factor Factor) {
	addConvertor(maindata, fromunit, tounit, factor)
	addConvertor(maindata, tounit, fromunit, Factor{factor.before, 1 / factor.factor, factor.after, factor.how})
}
func Convert(maindata map[string]map[string]FactorRandomizer, fromunit string, tounit string, value float64) (float64, string) {
	factor := getFactor(maindata[fromunit][tounit])
	return (factor.before+value)*factor.factor + factor.after, factor.how
}
func ConvertString(maindata map[string]map[string]FactorRandomizer, fromunit string, tounit string, value float64) string {
	res, how := Convert(maindata, fromunit, tounit, value)
	if how == "" {
		return fmt.Sprintf("Je ne sais pas convertir les %s en %s", fromunit, tounit)
	}
	return fmt.Sprintf("%s%s = %s%s %s", strconv.FormatFloat(value, 'f', -1, 64), fromunit, strconv.FormatFloat(res, 'f', -1, 64), tounit, how)
}
func Filldata() map[string]map[string]FactorRandomizer {
	data := map[string]map[string]FactorRandomizer{}
	addConvertors(data, "kg", "m", Factor{0, 21.0 / 9 * 1000000, 0, "par masse linétique moyenne de la soie"})
	addConvertors(data, "kg", "m", Factor{0, 0.2 / 0.4, 0, "de saucisse de morteau"})
	addConvertors(data, "kg", "m2", Factor{0, 39, 0, "par masse surfacique moyenne du béton"})
	addConvertors(data, "kg", "m3", Factor{0, 2000, 0, "par masse volumique moyenne du sable humide"})
	addConvertors(data, "m", "s", Factor{0, 2, 0, "par vitesse de marche"})
	addConvertors(data, "cd", "m2", Factor{0, 1.5 * 1000000000, 0, "par luminance solaire"})
	addConvertor(data, "m", "K", Factor{0, -7.0 / 1000, 15 - 273.15, "par temperature moyenne par altitude"})
	addConvertor(data, "K", "m", Factor{-273.15, 1 / -7.0 * 1000, -15, "par temperature moyenne par altitude"})
	addConvertor(data, "m", "C", Factor{0, -7.0 / 1000, 15, "par temperature moyenne par altitude"})
	addConvertor(data, "C", "m", Factor{0, 1 / -7.0 * 1000, -15, "par temperature moyenne par altitude"})
	return data
}

func mainhandler(w http.ResponseWriter, r *http.Request) {
	data := Filldata()
	params := mux.Vars(r)
	value, _ := strconv.ParseFloat(params["value"], 64)
	str := ConvertString(data, params["from"], params["to"], value)
	w.Header().Set("content-type", "application/json")
	w.Write([]byte(fmt.Sprintf(`{"str": "%s"}`, str)))
}
func Main() {
	// data := Filldata()
	// fmt.Println("data", data)
	// fmt.Println(getFactor(data["m"]["m"]))
	// fmt.Println(ConvertString(data, "kg", "m", 58))
	// fmt.Println(ConvertString(data, "m", "kg", 1))
	// fmt.Println(ConvertString(data, "m", "K", 1000))
	// fmt.Println(ConvertString(data, "K", "m", 280))
	// fmt.Println(ConvertString(data, "m", "K", 4850))
	router := mux.NewRouter()
	router.HandleFunc("/api/{from}/{to}/{value}", mainhandler).Methods("GET")
	http.ListenAndServe(":8001", router)
	// reader := bufio.NewReader(os.Stdin)
	// for {
	// 	fmt.Print("Enter quantity: ")
	// 	value, _ := reader.ReadString('\n')
	// 	valuef, _ := strconv.ParseFloat(value, 64)
	// 	fmt.Print("Enter fromunit: ")
	// 	fromunit, _ := reader.ReadString('\n')
	// 	fmt.Print("Enter tounit: ")
	// 	tounit, _ := reader.ReadString('\n')
	// 	fmt.Println(ConvertString(data, fromunit, tounit, valuef))
	// }
}
