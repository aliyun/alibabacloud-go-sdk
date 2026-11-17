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
	// The Protocol Type supported by the access control policy. Valid values:
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
	// > The valid values of this parameter depend on the value of the Proto parameter. If Proto is set to TCP, ApplicationNameList supports all the preceding Protocol Types. If both ApplicationNameList and ApplicationName are specified, ApplicationNameList takes precedence.
	//
	// example:
	//
	// FTP
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// The category ID. Default value: null. Valid values:
	//
	// - **All**: All categories.
	//
	// - **RiskDomain**: Risky domain names.
	//
	// - **RiskIP**: Risky IP addresses.
	//
	// - **AliYun**: Alibaba Cloud services.
	//
	// - **NotAliYun**: Non-Alibaba Cloud services.
	//
	// example:
	//
	// All
	CategoryId *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// Sets the paging settings to specify the page number to return for a paged query.
	//
	// Default value: 1, which means the first page is returned.
	//
	// example:
	//
	// 1
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The destination IP address of the domain name for outbound connections.
	//
	// example:
	//
	// 10.0.XX.XX
	DstIP *string `json:"DstIP,omitempty" xml:"DstIP,omitempty"`
	// The end time of the query. The value is a UNIX timestamp accurate to seconds.
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
	// - **asc**: Ascending order.
	//
	// - **desc*	- (default): Descending order.
	//
	// example:
	//
	// desc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// Sets the paging settings to specify the number of entries to return on each page for a paged query.
	//
	// Default value: 6, which means 6 entries are returned on each page. Maximum value: 10.
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
	// - **SessionCount*	- (default): The number of requests.
	//
	// - **TotalBytes**: The total volume of traffic.
	//
	// example:
	//
	// SessionCount
	Sort *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// The start time of the query. The value is a UNIX timestamp accurate to seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1656837360
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The intelligence tag ID. Valid values:
	//
	// - **AliYun**: Alibaba Cloud services.
	//
	// - **RiskDomain**: Risky domain names.
	//
	// - **RiskIP**: Risky IP addresses.
	//
	// - **TrustedDomain**: Trusted websites.
	//
	// - **AliPay**: Alipay.
	//
	// - **DingDing**: DingTalk.
	//
	// - **WeChat**: WeChat.
	//
	// - **QQ**: Tencent QQ.
	//
	// - **SecurityService**: Security services.
	//
	// - **Microsoft**: Microsoft.
	//
	// - **Amazon**: Amazon.
	//
	// - **Pan**: Cloud disks.
	//
	// - **Map**: Maps.
	//
	// - **Code**: Code hosting.
	//
	// - **SystemService**: System services.
	//
	// - **Taobao**: Taobao.
	//
	// - **Google**: Google.
	//
	// - **ThirdPartyService**: Third-party platform services.
	//
	// - **FirstFlow**: First visit.
	//
	// - **Downloader**: Malicious download.
	//
	// - **Alexa Top1M**: Popular website.
	//
	// - **Miner**: Miner Pool.
	//
	// - **Intelligence**: Threat intelligence.
	//
	// - **DDoS**: DDoS Trojan.
	//
	// - **Ransomware**: Ransomware.
	//
	// - **Spyware**: Spyware.
	//
	// - **Rogue**: Rogue software.
	//
	// - **Botnet**: Botnets.
	//
	// - **Suspicious**: Suspicious websites.
	//
	// - **C&C**: Remote control.
	//
	// - **Gang**: Gangs.
	//
	// - **CVE**: CVE vulnerabilities.
	//
	// - **Backdoor**: Backdoor Trojans.
	//
	// - **Phishing**: Phishing websites.
	//
	// - **APT**: APT attacks.
	//
	// - **Supply Chain Attack**: Supply chain attacks.
	//
	// - **Malicious software**: Malware.
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
