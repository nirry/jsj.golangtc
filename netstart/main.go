package main

import (
	"log"
	"fmt"
	"os"
	"jsj.golangtc/netstart/netstar"
	"jsj.golangtc/netstart/utils"
	"jsj.golangtc/netstart/model"
)

var (
	GoodsNumber string
	Code        string
)

func main() {

	tmp := `*********************************************************
	1.抢购商品
	2.查看订单
	exit.退出
*********************************************************`

	fmt.Println()
	fmt.Println()
	fmt.Println(tmp)
	fmt.Println()
	fmt.Println()

	ngo := netstar.NewNgo()

	for {

		fmt.Scanln(&Code)
		exit(Code)
		if Code != "1" && Code != "2" {
			log.Printf("序号错误：%s", Code)
			continue
		}

		switch Code {
		case "1":
			initProducts(ngo)
			break
		case "2":
			initOrders(ngo)
			break
		}

	}

}

func exit(Code string) {
	if Code == "exit" {
		os.Exit(0)
	}
}

func initProducts(ngo *netstar.Ngo) {

	dto, err := ngo.Init()

	if err != nil {
		log.Fatal(err)
		return
	}

	for index, v := range dto.Data.GoodsDetails {
		format := utils.TimeFormat(v.SaleTime / 1000)
		log.Printf("[%-2d] %s %-10f %d %-20s %-30s\n", index, v.GoodsNumber, v.GoodsNbdPrice, v.GoodsStatus, format, v.GoodsName)
	}

	var goodDto model.GoodDto

	for {

		log.Println("输入商品编号: ")
		fmt.Scanln(&GoodsNumber)
		exit(GoodsNumber)

		goodDto, err = ngo.First(GoodsNumber, dto.Data.GoodsDetails)

		if err == nil {
			break
		} else {
			log.Println(err)
		}
	}

	log.Printf("找到商品：%s %-10f %d %-30s\n", goodDto.GoodsNumber, goodDto.GoodsNbdPrice, goodDto.GoodsStatus, goodDto.GoodsName)
	ngo.Start(goodDto)

}

func initOrders(ngo *netstar.Ngo) {

	dto, err := ngo.NOrders()
	if err != nil {
		log.Fatal(err)
	}

	if dto.Data == nil || dto.Data.RecordList == nil || len(dto.Data.RecordList) == 0 {
		return
	}

	for index, o := range dto.Data.RecordList {
		format := utils.TimeFormat(o.OrderCreateTime / 1000)
		log.Printf("[%-2d] %s %-10f %d %-20s %-30s\n", index, o.GoodsNumber, o.OrderAmount, o.OrderStatus, format, o.GoodsName)
	}

}

//出兑操作

//
//Host: star.8.163.com
//Connection: keep-alive
//Content-Length: 116
//Accept: application/json, text/plain, */*
//Origin: https://star.8.163.com
//X-Requested-With: XMLHttpRequest
//User-Agent: Mozilla/5.0 (Linux; Android 9; EML-AL00 Build/HUAWEIEML-AL00; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/68.0.3440.91 Mobile Safari/537.36 hybrid/1.0.0 star_client_info_begin {hybridVersion: "1.0.0",clientVersion: "1.8.1",accountId: "ab09afc1788e2737ac8d7c464dcb7f2cf2f69a8e4b1c0dd49c11759348539bb9",channel: "e01170001"}star_client_info_end
//Content-Type: application/json;charset=UTF-8
//Referer: https://star.8.163.com/m
//Accept-Encoding: gzip, deflate
//Accept-Language: zh-CN,en-US;q=0.9
//Cookie: NTES_YD_SESS=OGT5Yf74YX5NkjGKBe.CailyKe.3SDQ8Mq0MxFZOlM9Ia7s5atXR4y.mEK3avpCXpu6Lz81NSXmo2SwuGq1UOGwxfmuWSpa2vf3RwM6TAx53mB4IjQXoZB8Tlt.qGsE7_5iKk2mm6xX_bufliSGikdDzv9QBrxqm5cbCV7huQi55o_mfwYlE.wzWrdo8ueFi.dwRU9yue0EpNOno3v9LbAZq90QG.8InuDCUdVeP7p90n; STAR_YD_SESS=OGT5Yf74YX5NkjGKBe.CailyKe.3SDQ8Mq0MxFZOlM9Ia7s5atXR4y.mEK3avpCXpu6Lz81NSXmo2SwuGq1UOGwxfmuWSpa2vf3RwM6TAx53mB4IjQXoZB8Tlt.qGsE7_5iKk2mm6xX_bufliSGikdDzv9QBrxqm5cbCV7huQi55o_mfwYlE.wzWrdo8ueFi.dwRU9yue0EpNOno3v9LbAZq90QG.8InuDCUdVeP7p90n; mail_psc_fingerprint=c75dfde6e42f7a25fb99b09e1d3a4720; _ga=GA1.2.459119499.1538527418; UM_distinctid=166376187b31e0-06f766dd4ffa55-70624200-41be0-166376187b4148; STAREIG=050e9c6dee6f7f798536aa357967f77b78b3bce1; mp_MA-AF7C-B137646D6BBB_hubble=%7B%22sessionReferrer%22%3A%20%22%22%2C%22updatedTime%22%3A%201544531994505%2C%22sessionStartTime%22%3A%201544531984883%2C%22deviceUdid%22%3A%20%22976ee96f-5b6a-45ab-8563-eabd2c4cea07%22%2C%22persistedTime%22%3A%201538527417478%2C%22LASTEVENT%22%3A%20%7B%22eventId%22%3A%20%22da_u_login%22%2C%22time%22%3A%201544531994506%7D%2C%22sessionUuid%22%3A%20%220188a839-2afa-4c23-918a-fa9accbaca77%22%2C%22user_id%22%3A%20%2268AF18D314C3B8A2E71023C41D8ACBA6E8CF62411BBA5C23372D06D39D6F1396%22%2C%22superProperties%22%3A%20%7B%22env%22%3A%20%22browser%22%2C%22channel%22%3A%20%22star%22%2C%22product%22%3A%20%22product_mucfc2.0%22%7D%2C%22currentReferrer%22%3A%20%22https%3A%2F%2Flq.163.com%2Fplatform%2Fopen-activate.html%23%2FIdCardVerify%3FshowProtocal%3Dtrue%22%7D

//{"goodsNumber":"120001541578607296840660","addressId":32332,"buyToken":"f80f3c34-7723-4ab7-b55a-ea161a4320ef363377"}

//https://star.8.163.com/api/goods/detail
//
//Host: star.8.163.com
//Connection: keep-alive
//Content-Length: 42
//Accept: application/json, text/plain, */*
//Origin: https://star.8.163.com
//X-Requested-With: XMLHttpRequest
//User-Agent: Mozilla/5.0 (Linux; Android 9; EML-AL00 Build/HUAWEIEML-AL00; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/68.0.3440.91 Mobile Safari/537.36 hybrid/1.0.0 star_client_info_begin {hybridVersion: "1.0.0",clientVersion: "1.8.1",accountId: "ab09afc1788e2737ac8d7c464dcb7f2cf2f69a8e4b1c0dd49c11759348539bb9",channel: "e01170001"}star_client_info_end
//Content-Type: application/json;charset=UTF-8
//Referer: https://star.8.163.com/m
//Accept-Encoding: gzip, deflate
//Accept-Language: zh-CN,en-US;q=0.9
//Cookie: NTES_YD_SESS=OGT5Yf74YX5NkjGKBe.CailyKe.3SDQ8Mq0MxFZOlM9Ia7s5atXR4y.mEK3avpCXpu6Lz81NSXmo2SwuGq1UOGwxfmuWSpa2vf3RwM6TAx53mB4IjQXoZB8Tlt.qGsE7_5iKk2mm6xX_bufliSGikdDzv9QBrxqm5cbCV7huQi55o_mfwYlE.wzWrdo8ueFi.dwRU9yue0EpNOno3v9LbAZq90QG.8InuDCUdVeP7p90n; STAR_YD_SESS=OGT5Yf74YX5NkjGKBe.CailyKe.3SDQ8Mq0MxFZOlM9Ia7s5atXR4y.mEK3avpCXpu6Lz81NSXmo2SwuGq1UOGwxfmuWSpa2vf3RwM6TAx53mB4IjQXoZB8Tlt.qGsE7_5iKk2mm6xX_bufliSGikdDzv9QBrxqm5cbCV7huQi55o_mfwYlE.wzWrdo8ueFi.dwRU9yue0EpNOno3v9LbAZq90QG.8InuDCUdVeP7p90n; mail_psc_fingerprint=c75dfde6e42f7a25fb99b09e1d3a4720; _ga=GA1.2.459119499.1538527418; UM_distinctid=166376187b31e0-06f766dd4ffa55-70624200-41be0-166376187b4148; STAREIG=050e9c6dee6f7f798536aa357967f77b78b3bce1; mp_MA-AF7C-B137646D6BBB_hubble=%7B%22sessionReferrer%22%3A%20%22%22%2C%22updatedTime%22%3A%201544531994505%2C%22sessionStartTime%22%3A%201544531984883%2C%22deviceUdid%22%3A%20%22976ee96f-5b6a-45ab-8563-eabd2c4cea07%22%2C%22persistedTime%22%3A%201538527417478%2C%22LASTEVENT%22%3A%20%7B%22eventId%22%3A%20%22da_u_login%22%2C%22time%22%3A%201544531994506%7D%2C%22sessionUuid%22%3A%20%220188a839-2afa-4c23-918a-fa9accbaca77%22%2C%22user_id%22%3A%20%2268AF18D314C3B8A2E71023C41D8ACBA6E8CF62411BBA5C23372D06D39D6F1396%22%2C%22superProperties%22%3A%20%7B%22env%22%3A%20%22browser%22%2C%22channel%22%3A%20%22star%22%2C%22product%22%3A%20%22product_mucfc2.0%22%7D%2C%22currentReferrer%22%3A%20%22https%3A%2F%2Flq.163.com%2Fplatform%2Fopen-activate.html%23%2FIdCardVerify%3FshowProtocal%3Dtrue%22%7D

//{"goodsNumber":"120001541578607296840660"}

//goodsNumber=600101544660434750309349

//POST https://star.8.163.com/api/banner/getList HTTP/1.1
//Host: star.8.163.com
//Connection: keep-alive
//Content-Length: 0
//Accept: application/json, text/plain, */*
//Origin: https://star.8.163.com
//X-Requested-With: XMLHttpRequest
//User-Agent: Mozilla/5.0 (Linux; Android 9; EML-AL00 Build/HUAWEIEML-AL00; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/68.0.3440.91 Mobile Safari/537.36 hybrid/1.0.0 star_client_info_begin {hybridVersion: "1.0.0",clientVersion: "1.8.1",accountId: "ab09afc1788e2737ac8d7c464dcb7f2cf2f69a8e4b1c0dd49c11759348539bb9",channel: "e01170001"}star_client_info_end
//Referer: https://star.8.163.com/m
//Accept-Encoding: gzip, deflate
//Accept-Language: zh-CN,en-US;q=0.9
//Cookie: NTES_YD_SESS=OGT5Yf74YX5NkjGKBe.CailyKe.3SDQ8Mq0MxFZOlM9Ia7s5atXR4y.mEK3avpCXpu6Lz81NSXmo2SwuGq1UOGwxfmuWSpa2vf3RwM6TAx53mB4IjQXoZB8Tlt.qGsE7_5iKk2mm6xX_bufliSGikdDzv9QBrxqm5cbCV7huQi55o_mfwYlE.wzWrdo8ueFi.dwRU9yue0EpNOno3v9LbAZq90QG.8InuDCUdVeP7p90n; STAR_YD_SESS=OGT5Yf74YX5NkjGKBe.CailyKe.3SDQ8Mq0MxFZOlM9Ia7s5atXR4y.mEK3avpCXpu6Lz81NSXmo2SwuGq1UOGwxfmuWSpa2vf3RwM6TAx53mB4IjQXoZB8Tlt.qGsE7_5iKk2mm6xX_bufliSGikdDzv9QBrxqm5cbCV7huQi55o_mfwYlE.wzWrdo8ueFi.dwRU9yue0EpNOno3v9LbAZq90QG.8InuDCUdVeP7p90n; mail_psc_fingerprint=c75dfde6e42f7a25fb99b09e1d3a4720; _ga=GA1.2.459119499.1538527418; UM_distinctid=166376187b31e0-06f766dd4ffa55-70624200-41be0-166376187b4148; STAREIG=050e9c6dee6f7f798536aa357967f77b78b3bce1; mp_MA-AF7C-B137646D6BBB_hubble=%7B%22sessionReferrer%22%3A%20%22%22%2C%22updatedTime%22%3A%201544531994505%2C%22sessionStartTime%22%3A%201544531984883%2C%22deviceUdid%22%3A%20%22976ee96f-5b6a-45ab-8563-eabd2c4cea07%22%2C%22persistedTime%22%3A%201538527417478%2C%22LASTEVENT%22%3A%20%7B%22eventId%22%3A%20%22da_u_login%22%2C%22time%22%3A%201544531994506%7D%2C%22sessionUuid%22%3A%20%220188a839-2afa-4c23-918a-fa9accbaca77%22%2C%22user_id%22%3A%20%2268AF18D314C3B8A2E71023C41D8ACBA6E8CF62411BBA5C23372D06D39D6F1396%22%2C%22superProperties%22%3A%20%7B%22env%22%3A%20%22browser%22%2C%22channel%22%3A%20%22star%22%2C%22product%22%3A%20%22product_mucfc2.0%22%7D%2C%22currentReferrer%22%3A%20%22https%3A%2F%2Flq.163.com%2Fplatform%2Fopen-activate.html%23%2FIdCardVerify%3FshowProtocal%3Dtrue%22%7D
//
//
//
////所有商品
//{"code":200,"msg":"请求成功","data":{"list":["120001543482888051275234","120101544605530401789046","120001544436059877923708","120001541578607296840660","900101543544606481860051","300101544085767924292090","700101544518262243443720","130101544580129696267130","700101544587219758209409","700101544587887281610208","600101544660286442062312","600101544660434750309349","600101544660591556430172","130101544667910263598661","900101544597178570415219","300101544752663947662267","600101544496624396233074","130101544667653418868807","700101544668085492593626","700101544672581402383601","700101544701209932157040","200101544706361590865098","200101544706428959335189","400101544711739573489911","700101544754123177500443","500101544757344052734405","300101544771637381492268","130101544794548107715448","700101545012759855824913","200101545049340910766011","300101544083653037657195","200101544427206815005649","130101544585879657780367","900101543544940432958573","900101543546260839191696","900101543546363840563142","900101543926354423407490","130101543983185362695159","130101543983414672848775","130101543987576435176590","300101544007407675845209","300101544084595065548945","900101544095812979185916","130101544102266336679985","130101544103264917565810","130101544104045179202036","130101544156973930243665","500101544170522279909298","500101544172546621978424","500101544172672981898186","130101544249277795553080","200101544324217241681847","700101544347292124434275","130101544353099702479030","900101544355678258116183","700101544356885978594686","500101544364611630150830","600101544414787573963476","600101544415360457916252","500101544421692310998137","700101544427011464902644","600101544427096224406497","700101544427269942063364","700101544427326012344501","300101544427689393992731","300101544427876446722120","600101544428438066317702","200101544434547108199874","500101544446283703196265","300101544449330275031410","300101544498053452680517","300101544498148626188404","300101544498265847819554","300101544498337468085624","130101544512891380608849","130101544512955991449964","900101544521263338491443","700101544525087163172071","600101544534584010454671","200101544536759128112740","400101544539982118378787","700101544542331209585818","700101544576693974302904","500101544578087598592399","130101544586068620156790","700101544586976665393094","200101544591652077825454","600101544593500428908786","130101544612879318116789","900101544620926464355753","600101544669925761169146","800101544693792300943684"],"datetime":1545266700502}}
