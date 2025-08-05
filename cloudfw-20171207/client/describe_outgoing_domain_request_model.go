// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOutgoingDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategoryId(v string) *DescribeOutgoingDomainRequest
	GetCategoryId() *string
	SetCurrentPage(v string) *DescribeOutgoingDomainRequest
	GetCurrentPage() *string
	SetDataType(v string) *DescribeOutgoingDomainRequest
	GetDataType() *string
	SetDomain(v string) *DescribeOutgoingDomainRequest
	GetDomain() *string
	SetEndTime(v string) *DescribeOutgoingDomainRequest
	GetEndTime() *string
	SetIsAITraffic(v string) *DescribeOutgoingDomainRequest
	GetIsAITraffic() *string
	SetLang(v string) *DescribeOutgoingDomainRequest
	GetLang() *string
	SetOrder(v string) *DescribeOutgoingDomainRequest
	GetOrder() *string
	SetPageSize(v string) *DescribeOutgoingDomainRequest
	GetPageSize() *string
	SetPublicIP(v string) *DescribeOutgoingDomainRequest
	GetPublicIP() *string
	SetSort(v string) *DescribeOutgoingDomainRequest
	GetSort() *string
	SetStartTime(v string) *DescribeOutgoingDomainRequest
	GetStartTime() *string
	SetTagIdNew(v string) *DescribeOutgoingDomainRequest
	GetTagIdNew() *string
}

type DescribeOutgoingDomainRequest struct {
	// The type of the service. This parameter is empty by default. Valid values:
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
	// The source of traffic for statistics. Valid values:
	//
	// 	- **internet*	- (default): the Internet firewall.
	//
	// 	- **nat**: NAT firewalls.
	//
	// example:
	//
	// nat
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// The domain name in outbound connections.
	//
	// example:
	//
	// www.aliyundoc.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The end of the time range to query. The value is a UNIX timestamp. Unit: seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1656750960
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Specifies whether to collect statistics only on AI service access traffic. Valid values:
	//
	// 	- **true**
	//
	// 	- **false*	- (default)
	//
	// example:
	//
	// true
	IsAITraffic *string `json:"IsAITraffic,omitempty" xml:"IsAITraffic,omitempty"`
	// The language of the content within the request. Valid values:
	//
	// 	- **zh**: Chinese (default)
	//
	// 	- **en**: English
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
	// Default value: 6. Maximum value: 100.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The public IP address of the Elastic Compute Service (ECS) instance that initiates outbound connections.
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
	// 1656664560
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
	// 	- **FirstFlow**: the first time when an outbound connection is initiated
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

func (s DescribeOutgoingDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeOutgoingDomainRequest) GoString() string {
	return s.String()
}

func (s *DescribeOutgoingDomainRequest) GetCategoryId() *string {
	return s.CategoryId
}

func (s *DescribeOutgoingDomainRequest) GetCurrentPage() *string {
	return s.CurrentPage
}

func (s *DescribeOutgoingDomainRequest) GetDataType() *string {
	return s.DataType
}

func (s *DescribeOutgoingDomainRequest) GetDomain() *string {
	return s.Domain
}

func (s *DescribeOutgoingDomainRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeOutgoingDomainRequest) GetIsAITraffic() *string {
	return s.IsAITraffic
}

func (s *DescribeOutgoingDomainRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeOutgoingDomainRequest) GetOrder() *string {
	return s.Order
}

func (s *DescribeOutgoingDomainRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeOutgoingDomainRequest) GetPublicIP() *string {
	return s.PublicIP
}

func (s *DescribeOutgoingDomainRequest) GetSort() *string {
	return s.Sort
}

func (s *DescribeOutgoingDomainRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeOutgoingDomainRequest) GetTagIdNew() *string {
	return s.TagIdNew
}

func (s *DescribeOutgoingDomainRequest) SetCategoryId(v string) *DescribeOutgoingDomainRequest {
	s.CategoryId = &v
	return s
}

func (s *DescribeOutgoingDomainRequest) SetCurrentPage(v string) *DescribeOutgoingDomainRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeOutgoingDomainRequest) SetDataType(v string) *DescribeOutgoingDomainRequest {
	s.DataType = &v
	return s
}

func (s *DescribeOutgoingDomainRequest) SetDomain(v string) *DescribeOutgoingDomainRequest {
	s.Domain = &v
	return s
}

func (s *DescribeOutgoingDomainRequest) SetEndTime(v string) *DescribeOutgoingDomainRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeOutgoingDomainRequest) SetIsAITraffic(v string) *DescribeOutgoingDomainRequest {
	s.IsAITraffic = &v
	return s
}

func (s *DescribeOutgoingDomainRequest) SetLang(v string) *DescribeOutgoingDomainRequest {
	s.Lang = &v
	return s
}

func (s *DescribeOutgoingDomainRequest) SetOrder(v string) *DescribeOutgoingDomainRequest {
	s.Order = &v
	return s
}

func (s *DescribeOutgoingDomainRequest) SetPageSize(v string) *DescribeOutgoingDomainRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeOutgoingDomainRequest) SetPublicIP(v string) *DescribeOutgoingDomainRequest {
	s.PublicIP = &v
	return s
}

func (s *DescribeOutgoingDomainRequest) SetSort(v string) *DescribeOutgoingDomainRequest {
	s.Sort = &v
	return s
}

func (s *DescribeOutgoingDomainRequest) SetStartTime(v string) *DescribeOutgoingDomainRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeOutgoingDomainRequest) SetTagIdNew(v string) *DescribeOutgoingDomainRequest {
	s.TagIdNew = &v
	return s
}

func (s *DescribeOutgoingDomainRequest) Validate() error {
	return dara.Validate(s)
}
