package v0

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

//用户api_id和key
var api_id = "test"
var api_key = "JDJhJDEwJDRYSy4wdXN0Q3FsMi9VZ3Y5NGMvMnVYMEgzeExoc21CRXZuZ3ZnOUpEVThTWUtTZnlLUjJP"

//api地址
var base_url = "http://127.0.0.1:8000"
var API_NAME_ORDER = "order"
var API_NAME_LABEL = "label"
var API_POST = "POST"
var API_PUT = "PUT"
var API_GET = "GET"
var API_DELETE = "DELETE"

//var create_request = "{\"customer_number\":\"CSUM15122100003\",\"channel_code\":\"eee\",\"type\":\"10\",\"packs\":1,\"weight\":1.000,\"length\":10.00,\"width\":10.00,\"height\":6.00,\"cbm\":0.00,\"receiver\":{ \"name\":\"Suzzanne Khudruj\",\"address\":\"59 Abbington dr\",\"country_code\":\"US\",\"state\":\"VIC\",\"city\":\"BUNDOORA\",\"suburb\":\"BUNDOORA\",\"postcode\":\"3083\",\"phone\":\"0399311598\"},\"sender\":{ \"name\":\"速玛（深圳）物流有限公司\",\"address\":\"广东省深圳市南山区侨城东路\",\"state\":\"广东省\",\"city\":\"深圳市\",\"suburb\":\"深圳市\",\"postcode\":\"519000\",\"phone\":\"18681472186\",\"email\":\"daniel@sumxpress.com\"},\"items\":[{\"type\":\"O\",\"name\":\"Mobile Phone\",\"cname\":\"Mobile Phone\",\"quantity\":1,\"unit_value\":554.0}]}"

type Contact struct {
	Name        string `json:"name"`
	Address     string `json:"address"`
	Country     string `json:"country"`
	CountryCode string `json:"country_code"`
	State       string `json:"state"`
	City        string `json:"city"`
	Suburb      string `json:"suburb"`
	Postcode    string `json:"postcode"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	TaxID       string `json:"tax_id"`
}
type Items struct {
	Type      string  `json:"type"`
	Name      string  `json:"name"`
	Cname     string  `json:"cname"`
	Brand     string  `json:"brand"`
	Model     string  `json:"model"`
	Quantity  int     `json:"quantity"`
	UnitValue float64 `json:"unit_value"`
	HsCode    string  `json:"hs_code"`
	Material  string  `json:"material"`
	Purpose   string  `json:"purpose"`
	Sku       string  `json:"sku"`
	URL       string  `json:"url"`
	Ean       string  `json:"ean"`
	SerialNo  string  `json:"serial_no"`
}
type Order struct {
	Number          string  `json:"number"`
	Name            string  `json:"name"`
	CustomerNumber  string  `json:"customer_number"`
	FirstMileNumber string  `json:"first_mile_number"`
	CustomerCode    string  `json:"customer_code"`
	CustomerName    string  `json:"customer_name"`
	Ttk             string  `json:"ttk"`
	ChannelCode     string  `json:"channel_code"`
	ChannelName     string  `json:"channel_name"`
	Type            string  `json:"type"`
	State           int     `json:"state"`
	Packs           int     `json:"packs"`
	Weight          float64 `json:"weight"`
	Length          float64 `json:"length"`
	Width           float64 `json:"width"`
	Height          float64 `json:"height"`
	Cbm             float64 `json:"cbm"`
	Clear           string  `json:"clear"`
	Insurance       float64 `json:"insurance"`
	Price           float64 `json:"price"`
	LabelState      string  `json:"label_state"`
	Currency        string  `json:"currency"`
	Receiver        Contact `json:"receiver"`
	Sender          Contact `json:"sender"`
	ReturnOfAddress Contact `json:"return_of_address"`
	Notes           string  `json:"notes"`
	Items           []Items `json:"items"`
}
type packageResponse struct {
	Code        int     `json:"code"`
	Msg         string  `json:"msg"`
	TotalCount  int     `json:"total_count"`
	PageSize    int     `json:"page_size"`
	TotalPage   int     `json:"total_page"`
	CurrentPage int     `json:"current_page"`
	Orders      []Order `json:"orders"`
}

func fillPacketGlobavend(customer_number string, number string) *Order {

	//	items := []*Items{
	//		&Items{
	items := make([]Items, 0)
	item := Items{
		Type:      "O",
		Name:      "Mobile",
		Cname:     "手机",
		Brand:     "iPhone",
		Model:     "11",
		Quantity:  1,
		UnitValue: 60,
		HsCode:    "560811",
		Material:  "",
		Purpose:   "",
		Sku:       "{2}ND1001-1(1), ML25400-1(1) {B}",
		URL:       "",
		//Ean:
		//SerialNo:
		//	},
	}
	items = append(items, item)

	order := Order{
		Items: items,

		//Number: "OWLAU220814000101",
		Number: number,
		//Name            string  `json:"name"`
		CustomerNumber: customer_number, //"CSUM15122100003",
		//FirstMileNumber string  `json:"first_mile_number"`
		CustomerCode: "test",
		//CustomerName    string  `json:"customer_name"`
		//Ttk             string  `json:"ttk"`
		ChannelCode: "test-Grobavend", //"test-sumx", // "test-brazil",
		//ChannelName     string  `json:"channel_name"`
		Type: "10",
		//State           int     `json:"state"`
		//Packs           int     `json:"packs"`
		Weight:    1,
		Length:    15.00,
		Width:     15.00,
		Height:    15.00,
		Cbm:       0.00,
		Clear:     "ddu",
		Insurance: 1,
		//Price           float64 `json:"price"`
		//LabelState      string  `json:"label_state"`
		Currency: "USD",
		Receiver: Contact{
			Name:        "James Birnie",
			Address:     "70 Iris Street",
			Country:     "AU",
			CountryCode: "AU",
			State:       "NSW",
			City:        "BEACON HILL",
			Suburb:      "Beacon Hill",
			Postcode:    "2100",
			Phone:       "0407297946",
			//Email: ,
			TaxID: "",
		},
		Sender: Contact{
			Name:        "SumXpress(SZ) Co.",
			Address:     "Nanshan District Guangdong State",
			Country:     "CN",
			CountryCode: "CN",
			State:       "GD",
			City:        "深圳",
			Suburb:      "HuaQiaoCheng Street",
			Postcode:    "519000",
			Phone:       "18681472186",
			Email:       "daniel@sumxpress.com",
			//TaxID       string `json:"tax_id"`
		},
		//ReturnOfAddress Contact `json:"return_of_address"`
		//Notes           string  `json:"notes"`

	}

	return &order
}

func fillPacketSumx(customer_number string, number string) *Order {

	//	items := []*Items{
	//		&Items{
	items := make([]Items, 0)
	item := Items{
		Type:      "O",
		Name:      "Mobile",
		Cname:     "手机",
		Brand:     "iPhone",
		Model:     "11",
		Quantity:  1,
		UnitValue: 60,
		HsCode:    "560811",
		Material:  "",
		Purpose:   "",
		Sku:       "{2}ND1001-1(1), ML25400-1(1) {B}",
		URL:       "",
		//Ean:
		//SerialNo:
		//	},
	}
	items = append(items, item)

	order := Order{
		Items: items,

		//Number: "OWLAU220814000101",
		Number: number,
		//Name            string  `json:"name"`
		CustomerNumber: customer_number, //"CSUM15122100003",
		//FirstMileNumber string  `json:"first_mile_number"`
		CustomerCode: "testa",
		//CustomerName    string  `json:"customer_name"`
		//Ttk             string  `json:"ttk"`
		ChannelCode: "test-sumx", //"test-sumx", // "test-brazil",
		//ChannelName     string  `json:"channel_name"`
		Type: "10",
		//State           int     `json:"state"`
		//Packs           int     `json:"packs"`
		Weight:    1,
		Length:    15.00,
		Width:     15.00,
		Height:    15.00,
		Cbm:       0.00,
		Clear:     "ddu",
		Insurance: 1,
		//Price           float64 `json:"price"`
		//LabelState      string  `json:"label_state"`
		Currency: "USD",
		Receiver: Contact{
			Name:        "James Birnie",
			Address:     "70 Iris Street",
			Country:     "AU",
			CountryCode: "AU",
			State:       "NSW",
			City:        "BEACON HILL",
			Suburb:      "Beacon Hill",
			Postcode:    "2100",
			Phone:       "0407297946",
			//Email: ,
			TaxID: "",
		},
		Sender: Contact{
			Name:        "SumXpress(SZ) Co.",
			Address:     "Nanshan District Guangdong State",
			Country:     "CN",
			CountryCode: "CN",
			State:       "GD",
			City:        "深圳",
			Suburb:      "HuaQiaoCheng Street",
			Postcode:    "519000",
			Phone:       "18681472186",
			Email:       "daniel@sumxpress.com",
			//TaxID       string `json:"tax_id"`
		},
		//ReturnOfAddress Contact `json:"return_of_address"`
		//Notes           string  `json:"notes"`

	}

	return &order
}

func fillPacketFeichi(customer_number string, number string) *Order {

	//	items := []*Items{
	//		&Items{
	items := make([]Items, 0)
	item := Items{
		Type:      "O",
		Name:      "Mobile",
		Cname:     "手机",
		Brand:     "iPhone",
		Model:     "11",
		Quantity:  1,
		UnitValue: 60,
		HsCode:    "392690",
		Material:  "",
		Purpose:   "",
		Sku:       "sku-1234567",
		URL:       "",
		//Ean:
		//SerialNo:
		//	},
	}
	items = append(items, item)

	order := Order{
		Items: items,

		//Number: "OWLAU220814000101",
		Number: number,
		//Name            string  `json:"name"`
		CustomerNumber: customer_number, //"CSUM15122100003",
		//FirstMileNumber string  `json:"first_mile_number"`
		CustomerCode: "testa",
		//CustomerName    string  `json:"customer_name"`
		//Ttk             string  `json:"ttk"`
		ChannelCode: "test-brazil", //"test-sumx", // "test-brazil",
		//ChannelName     string  `json:"channel_name"`
		Type: "10",
		//State           int     `json:"state"`
		//Packs           int     `json:"packs"`
		Weight:    1,
		Length:    15.00,
		Width:     15.00,
		Height:    15.00,
		Cbm:       0.00,
		Clear:     "ddu",
		Insurance: 1,
		//Price           float64 `json:"price"`
		//LabelState      string  `json:"label_state"`
		Currency: "USD",
		Receiver: Contact{
			Name:        "Marcia Amoriim",
			Address:     "rua artur de azevedo 1681 apt 61A",
			Country:     "BR",
			CountryCode: "BR",
			State:       "SP",
			City:        "PAULO",
			Suburb:      "a",
			Postcode:    "05404014",
			Phone:       "55011978456814",
			//Email: ,
			TaxID: "00390128830",
		},
		Sender: Contact{
			Name:        "SumXpress(SZ) Co.",
			Address:     "Nanshan District Guangdong State",
			Country:     "CN",
			CountryCode: "CN",
			State:       "GD",
			City:        "深圳",
			Suburb:      "HuaQiaoCheng Street",
			Postcode:    "519000",
			Phone:       "18681472186",
			Email:       "daniel@sumxpress.com",
			//TaxID       string `json:"tax_id"`
		},
		//ReturnOfAddress Contact `json:"return_of_address"`
		//Notes           string  `json:"notes"`

	}

	return &order
}

func fillPacket(customer_number string, number string) *Order {
	return fillPacketSumx(customer_number, number)
	//return fillPacketFeichi(customer_number, number)
	//return fillPacketGlobavend(customer_number, number)
}

func GetHttpResponse(apiMethod string, httpMethod string, orderNo string, body string) (res *string) {
	timestamp := time.Now().Unix()
	query_path_data := fmt.Sprintf("/api/v0/%s/%s?_id=%s&_t=%d", apiMethod, orderNo, api_id, timestamp)
	sign_data := api_key + api_id + query_path_data + body + strconv.FormatInt(timestamp, 10) + api_key

	sign := md5.Sum([]byte(sign_data))
	url := base_url + query_path_data + "&_s=" + strings.ToUpper(fmt.Sprintf("%x", sign))

	fmt.Printf("sign data: %s\n", sign_data)
	fmt.Printf("sign: %s\n", strings.ToUpper(fmt.Sprintf("%x", sign)))

	fmt.Println("Trying to get response from :" + url)
	fmt.Println("Request body :" + body)
	fmt.Println("API method :" + httpMethod)

	myClient := http.Client{Timeout: time.Second * 10000}

	reqBody := strings.NewReader(body)

	req, err := http.NewRequest(httpMethod, url, reqBody)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "JDJhJDEwJDRYSy4wdXN0Q3FsMi9VZ3Y5NGMvMnVYMEgzeExoc21CRXZuZ3ZnOUpEVThTWUtTZnlLUjJP")

	resp, err := myClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Println("======================Response Message=======================")
		fmt.Println("Http Response code :" + resp.Status)
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	response := string(respBody)
	return &response
}

func GetHttpResponse1(url string, httpMethod string, body string) (res *string) {

	fmt.Println("Trying to get response from :" + url)
	fmt.Println("Request body :" + body)
	fmt.Println("API method :" + httpMethod)

	myClient := http.Client{Timeout: time.Second * 10000}

	reqBody := strings.NewReader(body)

	req, err := http.NewRequest(httpMethod, url, reqBody)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "JDJhJDEwJDRYSy4wdXN0Q3FsMi9VZ3Y5NGMvMnVYMEgzeExoc21CRXZuZ3ZnOUpEVThTWUtTZnlLUjJP")

	resp, err := myClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Println("======================Response Message=======================")
		fmt.Println("Http Response code :" + resp.Status)
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	response := string(respBody)
	return &response
}

//number: 运单号，创建的时候传空即可，其他时候需要有准确的number
//request: json格式的请求报文
//api_name: 方法名称
func CreateOrderProcess(numbers string) {
	fmt.Println("Create order starting...")
	fmt.Println("===============Request Message(CreateOrder)==================")

	order := fillPacket(numbers, "")
	req, err := json.Marshal(&order)
	if err != nil {
		return
	}

	response := GetHttpResponse("order", API_POST, "", string(req))
	if response == nil {
		fmt.Println("failure")
		return
	}

	fmt.Println("======================Response Message=======================")
	fmt.Println(*response)

	//解析json body
	rsp := &packageResponse{}
	err = json.Unmarshal([]byte(*response), rsp)
	if err != nil {
		return
	}

	//code := rsp["code"]
	//fmt.Println("Code: ", int64(code.(float64)))
	//if code == 0 {
	//	fmt.Println("Create order successfully.")
	//} else if code == 8 {
	//	fmt.Println("Create order is not fully completed, it was saved as draft order.")
	//} else {
	//	msg := rsp["msg"]
	//	fmt.Printf("Create Order failed. err msg=%s\n", msg)
	//}
}

func UpdateOrderProcess(numbers string) {
	fmt.Println("Update order starting...")
	fmt.Println("===============Request Message(UpdateOrder)==================")

	num := strings.Split(numbers, ",")
	customer_number := num[0]

	number := ""
	if len(num) > 1 {
		number = num[1]
	}

	order := fillPacket(customer_number, number)
	req, err := json.Marshal(&order)
	if err != nil {
		return
	}

	response := GetHttpResponse("order", API_PUT, number, string(req))

	if response == nil {
		fmt.Println("failure")
		return
	}

	fmt.Println("======================Response Message=======================")
	fmt.Println(*response)

	result := &packageResponse{}
	err = json.Unmarshal([]byte(*response), result)
	if err != nil {
		return
	}
}

func GetOrderProcess(numbers string) {
	fmt.Println("Get order starting...")
	fmt.Println("===============Request Message==================")

	response := GetHttpResponse("order", API_GET, numbers, "")
	if response == nil {
		fmt.Println("failure")
		return
	}

	fmt.Println("======================Response Message=======================")
	fmt.Println(*response)

	//解析json body
	//var rsp map[string]interface{}
	//err := json.Unmarshal([]byte(*response), &rsp)
	result := &packageResponse{}
	err := json.Unmarshal([]byte(*response), result)
	if err != nil {
		return
	}
}

func DeleteOrderProcess(numbers string) {
	fmt.Println("Delete order starting...")
	fmt.Println("=============== Request Message ==================")

	response := GetHttpResponse("order", API_DELETE, numbers, "")

	if response == nil {
		fmt.Println("failure")
		return
	}

	fmt.Println("======================Response Message=======================")
	fmt.Println(*response)

	//解析json body
	//var rsp map[string]interface{}
	//err := json.Unmarshal([]byte(*response), &rsp)
	result := &packageResponse{}
	err := json.Unmarshal([]byte(*response), result)
	if err != nil {
		return
	}
}

func GetLabelProcess(numbers string) {
	fmt.Println("Get label starting...")
	fmt.Println("===============Request Message==================")

	response := GetHttpResponse("label", API_GET, numbers, "")
	if response == nil {
		fmt.Println("failure")
		return
	}

	fmt.Println("======================Response Message=======================")
	fmt.Println(*response)

	//解析json body
	//var rsp map[string]interface{}
	//err := json.Unmarshal([]byte(*response), &rsp)
	result := &packageResponse{}
	err := json.Unmarshal([]byte(*response), result)
	if err != nil {
		return
	}
}
func GetTrackingProcess(numbers string) {
	fmt.Println("Get tracking starting...")
	fmt.Println("===============Request Message==================")

	timestamp := time.Now().Unix()
	query_path_data := fmt.Sprintf("/api/v0/%s/%s?_id=%s&_t=%d&type=%d", "tracking", numbers, api_id, timestamp, 0)
	sign_data := api_key + api_id + query_path_data + "" + strconv.FormatInt(timestamp, 10) + api_key

	sign := md5.Sum([]byte(sign_data))
	url := base_url + query_path_data + "&_s=" + strings.ToUpper(fmt.Sprintf("%x", sign))

	response := GetHttpResponse1(url, API_GET, "")
	if response == nil {
		fmt.Println("failure")
		return
	}

	fmt.Println("======================Response Message=======================")
	fmt.Println(*response)

	//解析json body
	//var rsp map[string]interface{}
	//err := json.Unmarshal([]byte(*response), &rsp)
	result := &packageResponse{}
	err := json.Unmarshal([]byte(*response), result)
	if err != nil {
		return
	}
}
