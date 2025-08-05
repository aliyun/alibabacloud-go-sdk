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
	// The application type in the access control policy. Valid values:
	//
	// 	- **FTP**
	//
	// 	- **HTTP**
	//
	// 	- **HTTPS**
	//
	// 	- **Memcache**
	//
	// 	- **MongoDB**
	//
	// 	- **MQTT**
	//
	// 	- **MySQL**
	//
	// 	- **RDP**
	//
	// 	- **Redis**
	//
	// 	- **SMTP**
	//
	// 	- **SMTPS**
	//
	// 	- **SSH**
	//
	// 	- **SSL_No_Cert**
	//
	// 	- **SSL**
	//
	// 	- **VNC**
	//
	// >  The value of this parameter depends on the value of Proto. If you set Proto to TCP, you can set ApplicationNameList to any valid value. If you specify both ApplicationNameList and ApplicationName, only the value of ApplicationNameList is used.
	//
	// example:
	//
	// FTP
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// The ID of the service to which the destination IP address belongs. This parameter is left empty by default. Valid values:
	//
	// 	- **All**: all services
	//
	// 	- **RiskDomain**: risky domain names
	//
	// 	- **RiskIP**: risky IP addresses
	//
	// 	- **AliYun**: Alibaba Cloud services
	//
	// 	- **NotAliYun**: third-party services
	//
	// example:
	//
	// All
	CategoryId *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// The number of the page to return.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The destination IP address in the outbound connection that is initiated to access a domain name.
	//
	// example:
	//
	// 10.0.XX.XX
	DstIP *string `json:"DstIP,omitempty" xml:"DstIP,omitempty"`
	// The end of the time range to query. The value is a UNIX timestamp. Unit: seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1656923760
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The language of the content within the response. Valid values:
	//
	// 	- **zh*	- (default)
	//
	// 	- **en**
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The method that you want to use to sort the query results. Valid values:
	//
	// 	- **asc**
	//
	// 	- **desc*	- (default)
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
	// The public IP address of the Elastic Compute Service (ECS) instance that initiates the outbound connection.
	//
	// example:
	//
	// 192.0.XX.XX
	PublicIP *string `json:"PublicIP,omitempty" xml:"PublicIP,omitempty"`
	// The field based on which you want to sort the query results. Valid values:
	//
	// 	- **SessionCount*	- (default): the number of requests.
	//
	// 	- **TotalBytes**: the total volume of traffic.
	//
	// example:
	//
	// SessionCount
	Sort *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// The beginning of the time range to query. The value is a UNIX timestamp. Unit: seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1656837360
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The ID of the tag. Valid values:
	//
	// 	- **AliYun**: Alibaba Cloud service
	//
	// 	- **RiskDomain**: risky domain name
	//
	// 	- **RiskIP**: risky IP address
	//
	// 	- **TrustedDomain**: trusted website
	//
	// 	- **AliPay**: Alipay
	//
	// 	- **DingDing**: DingTalk
	//
	// 	- **WeChat**: WeChat
	//
	// 	- **QQ**: Tencent QQ
	//
	// 	- **SecurityService**: security service
	//
	// 	- **Microsoft**: Microsoft
	//
	// 	- **Amazon**: Amazon Web Services (AWS)
	//
	// 	- **Pan**: cloud disk
	//
	// 	- **Map**: map
	//
	// 	- **Code**: code hosting
	//
	// 	- **SystemService**: system service
	//
	// 	- **Taobao**: Taobao
	//
	// 	- **Google**: Google
	//
	// 	- **ThirdPartyService**: third-party service
	//
	// 	- **FirstFlow**: the first time
	//
	// 	- **Downloader**: malicious download
	//
	// 	- **Alexa Top1M**: popular website
	//
	// 	- **Miner**: mining pool
	//
	// 	- **Intelligence**: threat intelligence
	//
	// 	- **DDoS**: DDoS trojan
	//
	// 	- **Ransomware**: ransomware
	//
	// 	- **Spyware**: spyware
	//
	// 	- **Rogue**: rogue software
	//
	// 	- **Botnet**: botnet
	//
	// 	- **Suspicious**: suspicious website
	//
	// 	- **C\\&C**: command and control (C\\&C)
	//
	// 	- **Gang**: gang
	//
	// 	- **CVE**: Common Vulnerabilities and Exposures (CVE)
	//
	// 	- **Backdoor**: webshell
	//
	// 	- **Phishing**: phishing website
	//
	// 	- **APT**: advanced persistent threat (APT) attack
	//
	// 	- **Supply Chain Attack**: supply chain attack
	//
	// 	- **Malicious software**: malware
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
