// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOutgoingDestinationIPRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationName(v string) *DescribeOutgoingDestinationIPRequest
	GetApplicationName() *string
	SetCategoryId(v string) *DescribeOutgoingDestinationIPRequest
	GetCategoryId() *string
	SetCurrentPage(v string) *DescribeOutgoingDestinationIPRequest
	GetCurrentPage() *string
	SetDstIP(v string) *DescribeOutgoingDestinationIPRequest
	GetDstIP() *string
	SetEndTime(v string) *DescribeOutgoingDestinationIPRequest
	GetEndTime() *string
	SetLang(v string) *DescribeOutgoingDestinationIPRequest
	GetLang() *string
	SetOrder(v string) *DescribeOutgoingDestinationIPRequest
	GetOrder() *string
	SetPageSize(v string) *DescribeOutgoingDestinationIPRequest
	GetPageSize() *string
	SetPort(v string) *DescribeOutgoingDestinationIPRequest
	GetPort() *string
	SetPrivateIP(v string) *DescribeOutgoingDestinationIPRequest
	GetPrivateIP() *string
	SetPublicIP(v string) *DescribeOutgoingDestinationIPRequest
	GetPublicIP() *string
	SetSort(v string) *DescribeOutgoingDestinationIPRequest
	GetSort() *string
	SetStartTime(v string) *DescribeOutgoingDestinationIPRequest
	GetStartTime() *string
	SetTagIdNew(v string) *DescribeOutgoingDestinationIPRequest
	GetTagIdNew() *string
}

type DescribeOutgoingDestinationIPRequest struct {
	// The application type supported by the access control policy.
	//
	// - **FTP**
	//
	// - **HTTP**
	//
	// - **HTTPS**
	//
	// - **Memcache**
	//
	// - **MongoDB**
	//
	// - **MQTT**
	//
	// - **MySQL**
	//
	// - **RDP**
	//
	// - **Redis**
	//
	// - **SMTP**
	//
	// - **SMTPS**
	//
	// - **SSH**
	//
	// - **SSL_No_Cert**
	//
	// - **SSL**
	//
	// - **VNC**
	//
	// > The supported application types depend on the protocol type specified in the Proto parameter. If Proto is set to TCP, all application types listed above are supported. If both ApplicationName and ApplicationNameList are specified, the value of ApplicationNameList takes precedence.
	//
	// example:
	//
	// FTP
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// The ID of the service category. Valid values:
	//
	// - **All**: all categories
	//
	// - **RiskDomain**: risk domains
	//
	// - **RiskIP**: risk IPs
	//
	// - **AliYun**: Alibaba Cloud services
	//
	// - **NotAliYun**: services other than Alibaba Cloud services
	//
	// example:
	//
	// All
	CategoryId *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// The page number to return.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The destination IP address of the outbound connection.
	//
	// example:
	//
	// 10.0.XX.XX
	DstIP *string `json:"DstIP,omitempty" xml:"DstIP,omitempty"`
	// The end of the time range to query. The value is a timestamp in seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1656923760
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The language of the response. Valid values:
	//
	// - **zh*	- (default): Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The sort order. Valid values:
	//
	// - **asc**: ascending order.
	//
	// - **desc*	- (default): descending order.
	//
	// example:
	//
	// desc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The number of entries to return on each page.
	//
	// Default value: 6. Maximum value: 10.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The port number.
	//
	// example:
	//
	// 80
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The private IP address of the ECS instance that initiates the outbound connection.
	//
	// example:
	//
	// 192.168.XX.XX
	PrivateIP *string `json:"PrivateIP,omitempty" xml:"PrivateIP,omitempty"`
	// The public IP address of the ECS instance that initiates the outbound connection.
	//
	// example:
	//
	// 192.0.XX.XX
	PublicIP *string `json:"PublicIP,omitempty" xml:"PublicIP,omitempty"`
	// The field by which to sort the results. Valid values:
	//
	// - **SessionCount*	- (default): request count.
	//
	// - **TotalBytes**: total traffic.
	//
	// example:
	//
	// SessionCount
	Sort *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// The start of the time range to query. The value is a timestamp in seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1656837360
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The ID of the threat intelligence tag. Valid values:
	//
	// - **AliYun**: Alibaba Cloud service
	//
	// - **RiskDomain**: risk domain
	//
	// - **RiskIP**: risk IP
	//
	// - **TrustedDomain**: trusted website
	//
	// - **AliPay**: Alipay
	//
	// - **DingDing**: DingTalk
	//
	// - **WeChat**: WeChat
	//
	// - **QQ**: Tencent QQ
	//
	// - **SecurityService**: security service
	//
	// - **Microsoft**: Microsoft
	//
	// - **Amazon**: Amazon
	//
	// - **Pan**: cloud drive
	//
	// - **Map**: map
	//
	// - **Code**: code hosting
	//
	// - **SystemService**: system service
	//
	// - **Taobao**: Taobao
	//
	// - **Google**: Google
	//
	// - **ThirdPartyService**: third-party service
	//
	// - **FirstFlow**: first access
	//
	// - **Downloader**: malicious downloader
	//
	// - **Alexa Top1M**: popular website
	//
	// - **Miner**: mining pool
	//
	// - **Intelligence**: threat intelligence
	//
	// - **DDoS**: DDoS trojan
	//
	// - **Ransomware**: ransomware
	//
	// - **Spyware**: spyware
	//
	// - **Rogue**: rogue software
	//
	// - **Botnet**: botnet
	//
	// - **Suspicious**: suspicious website
	//
	// - **C\\&C**: command and control (C\\&C)
	//
	// - **Gang**: threat actor group
	//
	// - **CVE**: CVE
	//
	// - **Backdoor**: backdoor
	//
	// - **Phishing**: phishing website
	//
	// - **APT**: APT attack
	//
	// - **Supply Chain Attack**: supply chain attack
	//
	// - **Malicious software**: malware
	//
	// example:
	//
	// AliYun
	TagIdNew *string `json:"TagIdNew,omitempty" xml:"TagIdNew,omitempty"`
}

func (s DescribeOutgoingDestinationIPRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeOutgoingDestinationIPRequest) GoString() string {
	return s.String()
}

func (s *DescribeOutgoingDestinationIPRequest) GetApplicationName() *string {
	return s.ApplicationName
}

func (s *DescribeOutgoingDestinationIPRequest) GetCategoryId() *string {
	return s.CategoryId
}

func (s *DescribeOutgoingDestinationIPRequest) GetCurrentPage() *string {
	return s.CurrentPage
}

func (s *DescribeOutgoingDestinationIPRequest) GetDstIP() *string {
	return s.DstIP
}

func (s *DescribeOutgoingDestinationIPRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeOutgoingDestinationIPRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeOutgoingDestinationIPRequest) GetOrder() *string {
	return s.Order
}

func (s *DescribeOutgoingDestinationIPRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeOutgoingDestinationIPRequest) GetPort() *string {
	return s.Port
}

func (s *DescribeOutgoingDestinationIPRequest) GetPrivateIP() *string {
	return s.PrivateIP
}

func (s *DescribeOutgoingDestinationIPRequest) GetPublicIP() *string {
	return s.PublicIP
}

func (s *DescribeOutgoingDestinationIPRequest) GetSort() *string {
	return s.Sort
}

func (s *DescribeOutgoingDestinationIPRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeOutgoingDestinationIPRequest) GetTagIdNew() *string {
	return s.TagIdNew
}

func (s *DescribeOutgoingDestinationIPRequest) SetApplicationName(v string) *DescribeOutgoingDestinationIPRequest {
	s.ApplicationName = &v
	return s
}

func (s *DescribeOutgoingDestinationIPRequest) SetCategoryId(v string) *DescribeOutgoingDestinationIPRequest {
	s.CategoryId = &v
	return s
}

func (s *DescribeOutgoingDestinationIPRequest) SetCurrentPage(v string) *DescribeOutgoingDestinationIPRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeOutgoingDestinationIPRequest) SetDstIP(v string) *DescribeOutgoingDestinationIPRequest {
	s.DstIP = &v
	return s
}

func (s *DescribeOutgoingDestinationIPRequest) SetEndTime(v string) *DescribeOutgoingDestinationIPRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeOutgoingDestinationIPRequest) SetLang(v string) *DescribeOutgoingDestinationIPRequest {
	s.Lang = &v
	return s
}

func (s *DescribeOutgoingDestinationIPRequest) SetOrder(v string) *DescribeOutgoingDestinationIPRequest {
	s.Order = &v
	return s
}

func (s *DescribeOutgoingDestinationIPRequest) SetPageSize(v string) *DescribeOutgoingDestinationIPRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeOutgoingDestinationIPRequest) SetPort(v string) *DescribeOutgoingDestinationIPRequest {
	s.Port = &v
	return s
}

func (s *DescribeOutgoingDestinationIPRequest) SetPrivateIP(v string) *DescribeOutgoingDestinationIPRequest {
	s.PrivateIP = &v
	return s
}

func (s *DescribeOutgoingDestinationIPRequest) SetPublicIP(v string) *DescribeOutgoingDestinationIPRequest {
	s.PublicIP = &v
	return s
}

func (s *DescribeOutgoingDestinationIPRequest) SetSort(v string) *DescribeOutgoingDestinationIPRequest {
	s.Sort = &v
	return s
}

func (s *DescribeOutgoingDestinationIPRequest) SetStartTime(v string) *DescribeOutgoingDestinationIPRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeOutgoingDestinationIPRequest) SetTagIdNew(v string) *DescribeOutgoingDestinationIPRequest {
	s.TagIdNew = &v
	return s
}

func (s *DescribeOutgoingDestinationIPRequest) Validate() error {
	return dara.Validate(s)
}
