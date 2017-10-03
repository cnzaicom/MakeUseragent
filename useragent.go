package MakeUseragent

import (
	"bytes"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type useragent struct {
	nettype            []string
	android_phone_list []string
}

func New() useragent {
	nettype_ := []string{"3G", "4G", "WIFI"}
	android_phone_list_ := []string{
		"Nexus 4 Build/KOT49H", "Nexus 5 Build/MRA58N", "Nexus 6 Build/LYZ28E",
		"Nexus 7 Build/JSS15Q", "Nexus 5 Build/MRA58N", "Nexus 6 Build/LYZ28E",
		//三星
		"GT-I9152P Build/JLS36C", "SM-E7000 Build/KTU84P", "SM-G9200 Build/LMY47X",
		"GT-I9128I Build/JDQ39", "GT-I9500 Build/JDQ39", "SM-N9008V Build/LRX21V",
		"SM-N7506V Build/JLS36C", "SM-G3609 Build/KTU84P", "SCH-W2013 Build/IMM76D",
		//LG
		"LGMS323 Build/KOT49I.MS32310c",
		//OPPO/VIVO
		"OPPO R7 Build/KTU84P", "OPPO R7t Build/KTU84P", "R7007 Build/JLS36C", "R2017 Build/JLS36C", "R6007 Build/JLS36C",
		"1105 Build/KTU84P", "N5117 Build/JLS36C", "M571C Build/LMY47D", "R7Plus Build/LRX21M", "X909T Build/JDQ39",
		"A31t Build/KTU84P", "A31 Build/KTU84P", "R8207 Build/KTU84P", "R833T Build/JDQ39", "OPPO R7sm Build/LMY47V; wv",

		"vivo Y13iL Build/KTU84P", "vivo X5Pro D Build/LRX21M", "vivo Y22L Build/JLS36C", "vivo Y13T Build/JDQ39",
		"vivo X5Max Build/KTU84P", "ONE A2001 Build/LMY48W",
		//华为
		"VIE-AL10 Build/HUAWEIVIE-AL10; wv", "HUAWEI NXT-AL10 Build/HUAWEINXT-AL10", "HUAWEI NXT-CL00 Build/HUAWEINXT-CL00",
		"Che2-TL00M Build/HonorChe2-TL00M; wv", "FRD-AL10 Build/HUAWEIFRD-AL10", "HUAWEI RIO-AL00 Build/HuaweiRIO-AL00",
		"HUAWEI C199 Build/HuaweiC199", "HUAWEI RIO-TL00 Build/HUAWEIRIO-TL00; wv", "HUAWEI TAG-TL00 Build/HUAWEITAG-TL00",
		"HUAWEI MT7-CL00 Build/HuaweiMT7-CL00; wv", "PLE-703L Build/HuaweiMediaPad; wv", "PLK-TL01H Build/HONORPLK-TL01H",
		"EVA-AL10 Build/HUAWEIEVA-AL10",
		//小米
		"MI MAX Build/MMB29M", "MI 5 Build/NRD90M", "MI NOTE LTE Build/KTU84P", "MI 3C Build/MMB29M", "MI 5s Build/MXB48T",
		"MI NOTE LTE Build/MMB29M", "MI 2S Build/JRO03L", "MI 5 Build/MXB48T", "MI NOTE Pro Build/LRX22G",
		//联想ZUK
		"Z2 Plus Build/N2G47O; wv"}
	return useragent{nettype: nettype_, android_phone_list: android_phone_list_}
}

/*
	取单个随机数x  0 <= x <= (num-1)
*/
func (self useragent) getrandomNumber(num int) (int, error) {
	if num >= 0 {
		rand.Seed(time.Now().UnixNano())
		return rand.Intn(num), nil
	} else {
		return -1, errors.New("less than zero")
	}
}

/*
	取随机区间x      min <= x <= max
*/
func (self useragent) getrandomNumbers(min, max int) (int, error) {
	if max > min && min >= 0 {
		rand.Seed(time.Now().UnixNano())
		return rand.Intn(max-min+1) + min, nil
	} else {
		return -1, errors.New("max less than min or min less than zero")
	}
}

/*
	随机网络类型

*/
func (self useragent) rand_nettype() string {
	num, err := self.getrandomNumber(len(self.nettype))

	if err != nil {
		panic(err)
	}
	return self.nettype[num]
}

/*
*	随机android机型
 */
func (self useragent) rand_android_phone() string {
	num, err := self.getrandomNumber(len(self.android_phone_list))

	if err != nil {
		panic(err)
	}
	return self.android_phone_list[num]
}

/*
	随机android版本
*/
func (self useragent) rand_android_version() string {
	va, _ := self.getrandomNumbers(3, 9)
	vb, _ := self.getrandomNumber(10)
	vc, _ := self.getrandomNumber(10)
	return fmt.Sprintf("%d.%d.%d", va, vb, vc)
}
func (s useragent) rand_chrome() string {
	va, _ := s.getrandomNumbers(20, 61)
	vb, _ := s.getrandomNumbers(0, 9)
	vc, _ := s.getrandomNumbers(1000, 9999)
	vd, _ := s.getrandomNumbers(10, 100)
	return fmt.Sprintf("%d.%d.%d.%d", va, vb, vc, vd)
}
func (s useragent) rand_safari() string {
	va, _ := s.getrandomNumbers(100, 999)
	vb, _ := s.getrandomNumber(100)
	return fmt.Sprintf("%d.%d", va, vb)
}
func (s useragent) rand_mac_version() string {
	va, _ := s.getrandomNumbers(6, 12)
	vb, _ := s.getrandomNumbers(1, 9)
	vc, _ := s.getrandomNumbers(1, 9)
	return fmt.Sprintf("%d_%d_%d", va, vb, vc)
}
func (s useragent) rand_windows_version() string {
	va, _ := s.getrandomNumbers(6, 10)
	vb, _ := s.getrandomNumbers(0, 9)
	return fmt.Sprintf("%d.%d", va, vb)
}
func (s useragent) rand_key(lens int) string {
	str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var buffer bytes.Buffer
	for i := 0; i < lens; i++ {
		num, _ := s.getrandomNumber(36)
		buffer.WriteByte(str[num])
	}
	return buffer.String()
}
func (s useragent) rand_firefox() string {
	va, _ := s.getrandomNumbers(20, 60)
	return fmt.Sprintf("%d.0", va)
}
func (s useragent) internet_explorer() string {
	var str = [4]string{
		"Mozilla/5.0 (Windows NT 6.3; Trident/7.0; rv 11.0) like Gecko",
		"Mozilla/5.0 (compatible; WOW64; MSIE 10.0; Windows NT 6.2)",
		"Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Trident/5.0)",
		"Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 6.0; Trident/4.0)",
	}
	num, _ := s.getrandomNumber(4)
	return str[num]
}
func (s useragent) rand_linux_vertion() string {
	var str = [8]string{
		"FreeBSD amd64",
		"NetBSD amd64",
		"Ubuntu; Linux armv7l",
		"Linux i586;",
		"OpenBSD amd64",
		"Linux x86_64",
		"Ubuntu; Linux x86_64",
		"Linux i686",
	}
	num, _ := s.getrandomNumber(8)
	return str[num]
}
func (s useragent) Chrome_pc() string {
	num, _ := s.getrandomNumber(3)
	switch num {
	case 2:
		return s.chrome_pc_linux()
	case 1:
		return s.chrome_pc_mac()
	default:
		return s.chrome_pc_windows()
	}
}

func (s useragent) chrome_pc_linux() string {
	Safari := s.rand_safari()
	Chrome := s.rand_chrome()
	Linux := s.rand_linux_vertion()
	return fmt.Sprintf("Mozilla/5.0 (X11; %s) AppleWebKit/%s (KHTML, like Gecko) Chrome/%s Safari/%s", Linux, Safari, Chrome, Safari)
}

func (s useragent) chrome_pc_mac() string {
	Safari := s.rand_safari()
	Chrome := s.rand_chrome()
	mac_version := s.rand_mac_version()
	return fmt.Sprintf("Mozilla/5.0 (Macintosh; Intel Mac OS X %s) AppleWebKit/%s (KHTML, like Gecko) Chrome/%s Safari/%s", mac_version, Safari, Chrome, Safari)
}
func (s useragent) chrome_pc_windows() string {
	Safari := s.rand_safari()
	Chrome := s.rand_chrome()
	WindowsNT := s.rand_windows_version()
	return fmt.Sprintf("Mozilla/5.0 (Windows NT %s; WOW64) AppleWebKit/%s (KHTML, like Gecko) Chrome/%s Safari/%s", WindowsNT, Safari, Chrome, Safari)
}
func (s useragent) Firefox_pc() string {
	num, _ := s.getrandomNumber(3)
	switch num {
	case 2:
		return s.firefox_pc_linux()
	case 1:
		return s.firefox_pc_mac()
	default:
		return s.firefox_pc_windows()
	}
}
func (s useragent) firefox_pc_linux() string {
	Firefox := s.rand_firefox()
	Linux := s.rand_linux_vertion()
	return fmt.Sprintf("Mozilla/5.0 (X11; %s; rv:%s) Gecko/20100101 Firefox/%s", Linux, Firefox, Firefox)
}
func (s useragent) firefox_pc_mac() string {
	Firefox := s.rand_firefox()
	mac_version := s.rand_mac_version()
	return fmt.Sprintf("Mozilla/5.0 (Macintosh; Intel Mac OS X %s; rv:%s) Gecko/20100101 Firefox/%s", mac_version, Firefox, Firefox)
}
func (s useragent) firefox_pc_windows() string {
	Firefox := s.rand_firefox()
	WindowsNT := s.rand_windows_version()
	return fmt.Sprintf("Mozilla/5.0 (Windows NT %s; WOW64; rv:%s) Gecko/20100101 Firefox/%s", WindowsNT, Firefox, Firefox)
}

func (s useragent) Pc() string {
	num, _ := s.getrandomNumber(3)
	switch num {
	case 2:
		return s.Pc_mac()
	case 1:
		return s.Pc_linux()
	default:
		return s.Pc_windows()
	}
}

func (s useragent) Pc_linux() string {
	num, _ := s.getrandomNumber(2)
	switch num {
	case 1:
		return s.chrome_pc_linux()
	default:
		return s.firefox_pc_linux()
	}
}

func (s useragent) Pc_windows() string {
	num, _ := s.getrandomNumber(3)
	switch num {
	case 2:
		return s.firefox_pc_windows()
	case 1:
		return s.chrome_pc_windows()
	default:
		return s.firefox_pc_windows()
	}
}

func (s useragent) Pc_mac() string {
	num, _ := s.getrandomNumber(2)
	switch num {
	case 1:
		return s.firefox_pc_mac()
	default:
		return s.chrome_pc_mac()
	}
}

/*
	无线端
*/

func (s useragent) chrome_wap_android() string {
	androidVersion := s.rand_android_version()
	androidPhone := s.rand_android_phone()
	Safari := s.rand_safari()
	Chrome := s.rand_chrome()
	return fmt.Sprintf("Mozilla/5.0 (Linux; Android %s; %s) AppleWebKit/%s (KHTML, like Gecko) Chrome/%s Mobile Safari/%s", androidVersion, androidPhone, Safari, Chrome, Safari)
}

func (s useragent) chrome_wap_iphone() string {
	mac_version := s.rand_mac_version()
	Safari := s.rand_safari()
	Mobile := s.rand_key(6)
	return fmt.Sprintf("Mozilla/5.0 (iPhone; CPU iPhone OS %s like Mac OS X) AppleWebKit/%s (KHTML, like Gecko) Version/9.0 Mobile/%s Safari/%s", mac_version, Safari, Mobile, Safari)
}
func (s useragent) wechat_android() string {
	androidVersion := s.rand_android_version()
	androidPhone := s.rand_android_phone()
	Safari := s.rand_safari()
	Chrome := s.rand_chrome()
	num, _ := s.getrandomNumbers(1, 999999)
	TBS := fmt.Sprintf("%06d", num)
	va, _ := s.getrandomNumber(10)
	vb, _ := s.getrandomNumber(10)
	vc, _ := s.getrandomNumbers(1, 9999)
	MicroMessenger := fmt.Sprintf("6.%d.%d.%d", va, vb, vc)
	NetType := s.rand_nettype()
	return fmt.Sprintf("Mozilla/5.0 (Linux; Android %s; %s) AppleWebKit/%s (KHTML, like Gecko) Version/4.0 Chrome/%s Mobile MQQBrowser/6.2 TBS/%s Safari/%s MicroMessenger/%s NetType/%s Language/zh_CN", androidVersion, androidPhone, Safari, Chrome, TBS, Safari, MicroMessenger, NetType)
}
func (s useragent) wechat_iphone() string {
	mac_version := s.rand_mac_version()
	Safari := s.rand_safari()
	Mobile := s.rand_key(6)
	va, _ := s.getrandomNumber(10)
	vb, _ := s.getrandomNumber(10)
	vc, _ := s.getrandomNumbers(1, 9999)
	MicroMessenger := fmt.Sprintf("6.%d.%d.%d", va, vb, vc)
	NetType := s.rand_nettype()
	return fmt.Sprintf("Mozilla/5.0 (iPhone; CPU iPhone OS %s like Mac OS X) AppleWebKit/%s (KHTML, like Gecko) Mobile/%s MicroMessenger/%s NetType/%s Language/zh_CN", mac_version, Safari, Mobile, MicroMessenger, NetType)
}

func (s useragent) uc_browser_android() string {
	androidVersion := s.rand_android_version()
	androidPhone := s.rand_android_phone()
	Safari := s.rand_safari()
	return fmt.Sprintf("Mozilla/5.0 (Linux; U; Android %s; zh-cn; %s) AppleWebKit/%s (KHTML, like Gecko) Version/4.0 UCBrowser/1.0.0.100 U3/0.8.0 Mobile Safari/%s AliApp(TB/6.6.4) WindVane/8.0.0 1080X1920 GCanvas/1.4.2.21", androidVersion, androidPhone, Safari, Safari)
}

func (s useragent) baidu_box_app_android() string {
	androidVersion := s.rand_android_version()
	androidPhone := s.rand_android_phone()
	Safari := s.rand_safari()
	Chrome := s.rand_chrome()
	return fmt.Sprintf("Mozilla/5.0 (Linux; Android %s; %s) AppleWebKit/%s (KHTML, like Gecko) Version/4.0 Chrome/%s Mobile Safari/%s T7/7.4 baiduboxapp/8.4 (Baidu; P1 %s)", androidVersion, androidPhone, Safari, Chrome, Safari, androidVersion)
}
func (s useragent) baidu_box_app_iphone() string {
	mac_version := s.rand_mac_version()
	Safari := s.rand_safari()
	Mobile := s.rand_key(6)
	va, _ := s.getrandomNumbers(1, 20)
	vb, _ := s.getrandomNumber(10)
	vc, _ := s.getrandomNumber(10)
	vd, _ := s.getrandomNumber(10)
	ve, _ := s.getrandomNumbers(999, 9999)
	vf, _ := s.getrandomNumbers(1, 999)
	baiduboxapp := fmt.Sprintf("0_%d.%d.%d.%d_enohpi_%04d_%03d", va, vb, vc, vd, ve, vf)
	keystr := s.rand_key(51)
	return fmt.Sprintf("Mozilla/5.0 (iPhone; CPU iPhone OS %s like Mac OS X) AppleWebKit/%s (KHTML, like Gecko) Mobile/%s baiduboxapp/%s/2.01_4C2%%258enohPi/1099a/%s/1", mac_version, Safari, Mobile, baiduboxapp, keystr)
}

/*
	手机百度
*/
func (s useragent) Baidu_box_app() string {
	num, _ := s.getrandomNumber(2)
	switch num {
	case 1:
		return s.baidu_box_app_android()
	default:
		return s.baidu_box_app_iphone()
	}
}

/*
	android手机无线端
*/
func (s useragent) Android() string {
	num, _ := s.getrandomNumber(4)
	switch num {
	case 4:
		return s.baidu_box_app_android()
	case 2:
		return s.chrome_wap_android()
	case 1:
		return s.wechat_android()
	default:
		return s.uc_browser_android()
	}
}

/*
	iphone手机无线端
*/
func (s useragent) Iphone() string {
	num, _ := s.getrandomNumber(3)
	switch num {
	case 2:
		return s.baidu_box_app_iphone()
	case 1:
		return s.wechat_iphone()
	default:
		return s.chrome_wap_iphone()
	}
}

/*
	Chrome 无线端
*/
func (s useragent) Chrome_wap() string {
	num, _ := s.getrandomNumber(2)
	switch num {
	case 1:
		return s.chrome_wap_android()
	default:
		return s.chrome_wap_iphone()
	}
}

/*
	随机取
*/
func (s useragent) Random() string {
	num, _ := s.getrandomNumber(2)
	switch num {
	case 1:
		return s.Pc()
	default:
		return s.Wap()
	}
}

/*
	无线端
*/
func (s useragent) Wap() string {
	num, _ := s.getrandomNumber(2)
	switch num {
	case 1:
		return s.Android()
	default:
		return s.Iphone()
	}
}

/*
	wechat 微信
*/
func (s useragent) Wechat() string {
	num, _ := s.getrandomNumber(2)
	switch num {
	case 1:
		return s.wechat_android()
	default:
		return s.wechat_iphone()
	}
}
