package jtpay

import (
	"fmt"
	"testing"
)

func TestPay(t *testing.T) {
	cfg := &JTpayConfig{
		Usercode:      "XXXXXXXX",
		CompKey:       "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
		NotifyUrl:     "http://xx.cn/game/jtpay",
		ReturnUrl:     "http://xx.cn/game/jtpay",
		PlaceOrderUrl: "http://pay.jtpay.com/form/pay",
	}
	trans, err := NewJTpayTrans(cfg)
	if err != nil {
		t.Logf("err %v", err)
	}
	order := &JTpayOrder{
		P3_money:       "1",
		P9_paymethod:   "3",
		P20_productnum: "1",
		P25_terminal:   "2",
		P26_iswappay:   "3",
		//P4_returnurl:   trans.Config.ReturnUrl,
		//P5_notifyurl:   trans.Config.NotifyUrl,
		//P14_customname: "10226740",
	}
	order.P1_usercode = trans.Config.Usercode
	order.P6_ordertime = "20180202053128"   //OrderTimeStr()
	order.P2_order = "20180202173128191676" //GenOrderid(order.P1_usercode, order.P6_ordertime)
	//E5080E76A63C6803A3D8A4F8D1AE8CA2
	//E5080E76A63C6803A3D8A4F8D1AE8CA2
	trans.InitOrder(order)
	_, err = trans.Submit(order)
	//t.Logf("resp %s, err %v", resp, err)
	t.Logf("err %v", err)
	fmt.Printf("%s\n", order.P7_sign)
	t.Log(order)
}

func TestVerify(t *testing.T) {
	cfg := &JTpayConfig{
		Usercode:      "XXXXXXXX",
		CompKey:       "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
		NotifyUrl:     "http://xx.cn/game/jtpay",
		ReturnUrl:     "http://xx.cn/game/jtpay",
		PlaceOrderUrl: "http://pay.jtpay.com/form/pay",
	}
	trans, err := NewJTpayTrans(cfg)
	if err != nil {
		panic(err)
	}
	result := &NotifyResult{
		P1_usercode:      "10231993",
		P2_order:         "10231993-20180327000845-48314",
		P3_money:         "1",
		P4_status:        "1",
		P5_jtpayorder:    "20180326160844234406",
		P6_paymethod:     "3",
		P7_paychannelnum: "",
		P8_charset:       "UTF-8",
		P9_signtype:      "MD5",
		P10_sign:         "EC03C50E252D2A3CC550130243A540E9",
		P11_remark:       "",
	}
	//
	t.Log(trans.NotifyVerify(result))
}

//下单页面
func ioswapHtml(order *JTpayOrder) (str string) {
	str = `
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html>
<head>
	<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
	<title>竣付通</title>
</head>
<!--支付宝IOSwap支付请求提交页-->
<body onLoad="document.yeepay.submit();">
	<form name='yeepay' action='http://pay.jtpay.com/form/pay' method='post'  >
	`
	str += fmt.Sprintf("<input type='hidden' name='p1_usercode'				value='%s'>", order.P1_usercode)
	str += fmt.Sprintf("<input type='hidden' name='p2_order'				value='%s'>      ", order.P2_order)
	str += fmt.Sprintf("<input type='hidden' name='p3_money'				value='%s'>      ", order.P3_money)
	str += fmt.Sprintf("<input type='hidden' name='p4_returnurl'			value='%s'>  ", order.P4_returnurl)
	str += fmt.Sprintf("<input type='hidden' name='p5_notifyurl'			value='%s'>  ", order.P5_notifyurl)
	str += fmt.Sprintf("<input type='hidden' name='p6_ordertime'			value='%s'>  ", order.P6_ordertime)
	str += fmt.Sprintf("<input type='hidden' name='p7_sign'					value='%s'>       ", order.P7_sign)
	str += fmt.Sprintf("<input type='hidden' name='p9_paymethod'			value='%s'>  ", order.P9_paymethod)
	str += fmt.Sprintf("<input type='hidden' name='p14_customname'			value='%s'>", order.P14_customname)
	str += fmt.Sprintf("<input type='hidden' name='p17_customip'			value='%s'>  ", order.P17_customip)
	str += fmt.Sprintf("<input type='hidden' name='p25_terminal'			value='%s'>  ", order.P25_terminal)
	str += fmt.Sprintf("<input type='hidden' name='p26_iswappay'			value='%s'>  ", order.P26_iswappay)
	str += `
	</form>
</body>
</html>
`
	return
}
