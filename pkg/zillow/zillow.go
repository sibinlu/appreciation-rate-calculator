package zillow

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type ZillowData struct {
	Data struct {
		Property struct {
			HomeValueChartData []struct {
				Points []struct {
					X int64   `json:"x"`
					Y float32 `json:"y"`
				} `json:"points"`
				Name string `json:"name"`
			} `json:"homeValueChartData"`
		} `json:"property"`
	} `json:"data"`
}

func GetAppreciationRate(zpid string) {
	// Test Data
	//text := `{"data":{"property":{"homeValueChartData":[{"points":[{"x":1375254000000,"y":661568},{"x":1377932400000,"y":659984},{"x":1380524400000,"y":665582},{"x":1383202800000,"y":676746},{"x":1385798400000,"y":699866},{"x":1388476800000,"y":701472},{"x":1391155200000,"y":709900},{"x":1393574400000,"y":725111},{"x":1396249200000,"y":729830},{"x":1398841200000,"y":740391},{"x":1401519600000,"y":719772},{"x":1404111600000,"y":722177},{"x":1406790000000,"y":677043},{"x":1409468400000,"y":685700},{"x":1412060400000,"y":682027},{"x":1414738800000,"y":686055},{"x":1417334400000,"y":704797},{"x":1420012800000,"y":710781},{"x":1422691200000,"y":710355},{"x":1425110400000,"y":724441},{"x":1427785200000,"y":738931},{"x":1430377200000,"y":748571},{"x":1433055600000,"y":746333},{"x":1435647600000,"y":748985},{"x":1438326000000,"y":720327},{"x":1441004400000,"y":758816},{"x":1443596400000,"y":790999},{"x":1446274800000,"y":795292},{"x":1448870400000,"y":794413},{"x":1451548800000,"y":826400},{"x":1454227200000,"y":842120},{"x":1456732800000,"y":861622},{"x":1459407600000,"y":852230},{"x":1461999600000,"y":857141},{"x":1464678000000,"y":872328},{"x":1467270000000,"y":858867},{"x":1469948400000,"y":854423},{"x":1472626800000,"y":866029},{"x":1475218800000,"y":848212},{"x":1477897200000,"y":873905},{"x":1480492800000,"y":883021},{"x":1483171200000,"y":884110},{"x":1485849600000,"y":886754},{"x":1488268800000,"y":899227},{"x":1490943600000,"y":891027},{"x":1493535600000,"y":896211},{"x":1496214000000,"y":895635},{"x":1498806000000,"y":936221},{"x":1501484400000,"y":968874},{"x":1504162800000,"y":968542},{"x":1506754800000,"y":971003},{"x":1509433200000,"y":985961},{"x":1512028800000,"y":982534},{"x":1514707200000,"y":1013296},{"x":1517385600000,"y":1036891},{"x":1519804800000,"y":1058805},{"x":1522479600000,"y":1090812},{"x":1525071600000,"y":1175599},{"x":1527750000000,"y":1212625},{"x":1530342000000,"y":1196497},{"x":1533020400000,"y":1253803},{"x":1535698800000,"y":1216801},{"x":1538290800000,"y":1125369},{"x":1540969200000,"y":1126343},{"x":1543564800000,"y":1125584},{"x":1546243200000,"y":1105134},{"x":1548921600000,"y":1101325},{"x":1551340800000,"y":1116460},{"x":1554015600000,"y":1108127},{"x":1556607600000,"y":1104903},{"x":1559286000000,"y":1120608},{"x":1561878000000,"y":1113521},{"x":1564556400000,"y":1095840},{"x":1567234800000,"y":1105584},{"x":1569826800000,"y":1098714},{"x":1572505200000,"y":1097887},{"x":1575100800000,"y":1101900},{"x":1577779200000,"y":1123108},{"x":1580457600000,"y":1144196},{"x":1582963200000,"y":1138622},{"x":1585638000000,"y":1169462},{"x":1588230000000,"y":1143042},{"x":1590908400000,"y":1107600},{"x":1593500400000,"y":1122635},{"x":1596178800000,"y":1143408},{"x":1598857200000,"y":1157860},{"x":1601449200000,"y":1155412},{"x":1604127600000,"y":1184390},{"x":1606723200000,"y":1214733},{"x":1609401600000,"y":1203559},{"x":1612080000000,"y":1214392},{"x":1614499200000,"y":1263568},{"x":1617174000000,"y":1293432},{"x":1619766000000,"y":1340959},{"x":1622444400000,"y":1388697},{"x":1625036400000,"y":1430000},{"x":1627714800000,"y":1449000},{"x":1630393200000,"y":1429700},{"x":1632985200000,"y":1422200},{"x":1635663600000,"y":1434300},{"x":1638259200000,"y":1452400},{"x":1640937600000,"y":1477800},{"x":1643616000000,"y":1509500},{"x":1646035200000,"y":1563700},{"x":1648710000000,"y":1691200},{"x":1651302000000,"y":1738300},{"x":1653980400000,"y":1711300},{"x":1656572400000,"y":1657900},{"x":1659250800000,"y":1559600},{"x":1661929200000,"y":1454900},{"x":1664521200000,"y":1440100},{"x":1667199600000,"y":1461200},{"x":1669795200000,"y":1426900},{"x":1672473600000,"y":1384000},{"x":1675152000000,"y":1378100},{"x":1677571200000,"y":1396800},{"x":1680246000000,"y":1436900},{"x":1682838000000,"y":1463000},{"x":1685516400000,"y":1477400},{"x":1688108400000,"y":1496600},{"x":1689663600000,"y":1490800}],"name":"This home"},{"points":[{"x":1535698800000,"y":1120000}],"name":"Sale"}]}},"extensions":{}}`
	//textBytes := []byte(text)

	//Get Data from Zillow
	data := ZillowData{}

	zpid = normalize(zpid)
	if err := json.Unmarshal(getData(zpid), &data); err != nil {
		fmt.Println("Can not unmarshal JSON")
		return
	}

	points := data.Data.Property.HomeValueChartData[0].Points
	var rates float32 = 0
	for i := 0; i < len(points)-12; i++ {
		rates += points[i+12].Y/points[i].Y - 1
	}

	result := rates / float32(len(points)-12)

	fmt.Printf("You apprecition rate for zpid %v is\n", zpid)
	fmt.Printf("%.2f%%\n", result*100)
}

// Private methods.
func normalize(zpid string) (normalized string) {
	parts := strings.Split(zpid, "_")

	if len(parts) > 0 {
		normalized = parts[0]
	} else {
		fmt.Println("zpid is missing!")
	}
	return
}

func getData(zpid string) (body []byte) {
	url := "https://www.zillow.com/graphql/?zpid=" + zpid + "&timePeriod=TEN_YEARS&metricType=LOCAL_HOME_VALUES&forecast=true&useNewChartAPI=false&operationName=HomeValueChartDataQuery"
	payload := []byte(`{"operationName":"HomeValueChartDataQuery","variables":{"zpid":"` + zpid + `","timePeriod":"TEN_YEARS","metricType":"LOCAL_HOME_VALUES","forecast":true,"useNewChartAPI":false},"query":"query HomeValueChartDataQuery($zpid: ID\u0021, $metricType: HomeValueChartMetricType, $timePeriod: HomeValueChartTimePeriod, $useNewChartAPI: Boolean) {\n  property(zpid: $zpid) {\n    homeValueChartData(\n      metricType: $metricType\n      timePeriod: $timePeriod\n      useNewChartAPI: $useNewChartAPI\n    ) {\n      points {\n        x\n        y\n      }\n      name\n    }\n  }\n}\n"}`)

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("authority", "www.zillow.com")
	req.Header.Set("accept", "*/*")
	req.Header.Set("accept-language", "en-US,en;q=0.9,zh-CN;q=0.8,zh-TW;q=0.7,zh;q=0.6")
	req.Header.Set("client-id", "for-sale-sub-app-browser-client_home-value-v2-zestimate-summary_home-value-v2-zestimate-summary")
	req.Header.Set("content-type", "application/json")
	req.Header.Set("cookie", "x-amz-continuous-deployment-state=AYABeFMLxNDfKRAAQRRAf3YDPRsAPgACAAFEAB1kM2Jsa2Q0azB3azlvai5jbG91ZGZyb250Lm5ldAABRwAVRzA3MjU1NjcyMVRZRFY4RDcyVlpWAAEAAkNEABpDb29raWUAAACAAAAADAqj20gqJ6MFQRGBFQAwujRyW6z1Ws8whMgjGpJ+F1NZBllHejrixDgs2Ju77Mc3r0CJ8kYU5ARsLYdLkMVMAgAAAAAMAAQAAAAAAAAAAAAAAAAAAK+WpXdtk0kLibZRaHa1nJD%2F%2F%2F%2FAAAAAQAAAAAAAAAAAAAAAQAAAAwmI%2FQkVFcTzIMLqn3bBEMWv%2FYWUVkZvBcrwSem; zguid=24|%2492d9aab1-3d2e-49da-be90-c1b003045606; zgsession=1|b1478bc7-6537-46ae-9a7e-80c2f06f252b; _ga=GA1.2.793706849.1686767059; zjs_anonymous_id=%2292d9aab1-3d2e-49da-be90-c1b003045606%22; zjs_user_id=null; zg_anonymous_id=%2299f4da3f-4cc1-4202-ac96-23f2e1d79246%22; pxcts=ad4cf1eb-0ae0-11ee-a40c-7556415a4864; _pxvid=ad4ce09c-0ae0-11ee-a40c-73617e2f402f; _gcl_au=1.1.1894376152.1686767063; DoubleClickSession=true; __pdst=919a00a740e5416291afeed80fa43e8c; _pin_unauth=dWlkPU1qSTNPRE5sTURJdFlXWXdOeTAwWldJd0xUazBOek10TWpFeU1XUmtORFJtTUdVNA; JSESSIONID=0652C2854BEDEBA547F94120169BC37A; _gid=GA1.2.1089876490.1689635038; x-amz-continuous-deployment-state=AYABeF0anl8CwrYSvkakf0mA+GAAPgACAAFEAB1kM2Jsa2Q0azB3azlvai5jbG91ZGZyb250Lm5ldAABRwAVRzA3MjU1NjcyMVRZRFY4RDcyVlpWAAEAAkNEABpDb29raWUAAACAAAAADE7thHNoB0AvTTkpHAAw6p+EjM8MizF1oHaPmahsjnrkAsvs2mDY66ovJkfOjmAqrM%2Fox8VZw7ce6w9N6V2MAgAAAAAMAAQAAAAAAAAAAAAAAAAAAATNzUB6y8+yc8Eghot0ZcT%2F%2F%2F%2F%2FAAAAAQAAAAAAAAAAAAAAAQAAAAym3RTXdljed5IsYowPNMrHcC01WPROUfjJRmcP; _clck=vl18i9|2|fdd|0|1260; search=6|1692227507256%7Crect%3D37.59328762229768%252C-121.9866943359375%252C37.47771894357508%252C-122.12625503540039%26rid%3D49611%26disp%3Dmap%26mdm%3Dauto%26p%3D1%26sort%3Ddays%26z%3D1%26listPriceActive%3D1%26fs%3D1%26fr%3D0%26mmm%3D0%26rs%3D0%26ah%3D0%26singlestory%3D0%26housing-connector%3D0%26abo%3D0%26garage%3D0%26pool%3D0%26ac%3D0%26waterfront%3D0%26finished%3D0%26unfinished%3D0%26cityview%3D0%26mountainview%3D0%26parkview%3D0%26waterview%3D0%26hoadata%3D1%26zillow-owned%3D0%263dhome%3D0%26featuredMultiFamilyBuilding%3D0%26commuteMode%3Ddriving%26commuteTimeOfDay%3Dnow%09%0949611%09%7B%22isList%22%3Atrue%2C%22isMap%22%3Atrue%7D%09%09%09%09%09; _uetsid=38ae31e024f611eeaf7e77b30b6ff0ee; _uetvid=afd590600ae011eea1827fa4720a5def; _clsk=1bxl456|1689635674189|4|1|n.clarity.ms/collect; _px3=1232733f8f48ffac062396a4dee401607f441ca31a9676a2332909213a364f8c:Qs90b71pzHGx1AMFnznJdlYQoabTvTy0ik0QH9djSvTzsQn0eVFAWy8OjLHaaPSwCG/VnV2rRX5LDqTWjGl4/A==:1000:PmcfS0F9Ffump6YATFqSENV2IfA3kQ1Cbk21SrhZoZQKUdV28s7NGA9H4j2NMQ5J53Kxj2aBxrXysZfrV8kvHW9CUqpMPMDKnbjkV7t4FvBf2Wc4ryouh2oFIn6Bas48DXdRof8ikjrHy8fGhKssIf7uWUkSGL5pl9z4svugdjZ/dTYV51z2Qk3OEK4aCgYokRfa/viCCYyW5bsX/7YwCQ==; _gat=1; AWSALB=qv5GHL16vrO6QMDy2t5orrPpk87b3Pt3j1Y5VTsxYHWptKWUDFrnUX0QY9pm26BGZM1Any04tnCh9YKaiwAoFjY6PzI5YFsL3543hSzLmeJTUvieOH2KP7wJz7WI; AWSALBCORS=qv5GHL16vrO6QMDy2t5orrPpk87b3Pt3j1Y5VTsxYHWptKWUDFrnUX0QY9pm26BGZM1Any04tnCh9YKaiwAoFjY6PzI5YFsL3543hSzLmeJTUvieOH2KP7wJz7WI")
	req.Header.Set("origin", "https://www.zillow.com")
	req.Header.Set("referer", "https://www.zillow.com/homedetails/3411-Park-Blvd-Palo-Alto-CA-94306/19503546_zpid/")
	req.Header.Set("sec-ch-ua", `"Not.A/Brand";v="8", "Chromium";v="114", "Google Chrome";v="114"`)
	req.Header.Set("sec-ch-ua-mobile", "?1")
	req.Header.Set("sec-ch-ua-platform", `"Android"`)
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("user-agent", "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Mobile Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	// Debugging the response.
	//fmt.Println(string(body))
	return
}
