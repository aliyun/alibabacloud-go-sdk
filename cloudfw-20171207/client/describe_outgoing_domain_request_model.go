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
	// The product category. Default value: empty. Valid values:
	//
	// - **All**: All categories.
	//
	// - **RiskDomain**: Risky domain category.
	//
	// - **RiskIP**: Risky IP category.
	//
	// - **AliYun**: Alibaba Cloud product category.
	//
	// - **NotAliYun**: Non-Alibaba Cloud product category.
	//
	// example:
	//
	// All
	CategoryId *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// The page number of the results to return in a paged query.
	//
	// Default value: 1, which indicates the first page.
	//
	// example:
	//
	// 1
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The source of the traffic statistics. Default value: Internet firewall. Valid values:
	//
	// - **internet**: Internet firewall.
	//
	// - **nat**: NAT firewall.
	//
	// example:
	//
	// nat
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// The domain name of outbound connections.
	//
	// example:
	//
	// www.aliyundoc.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The end time of the query. The value is a UNIX timestamp in seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1656750960
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Specifies whether to collect statistics only on traffic that accesses AI services. Default value: false. Valid values:
	//
	// - **true**: Yes.
	//
	// - **false**: No.
	//
	// example:
	//
	// true
	IsAITraffic *string `json:"IsAITraffic,omitempty" xml:"IsAITraffic,omitempty"`
	// The language type of the request message. Valid values:
	//
	// - **zh*	- (default): Chinese
	//
	// - **en**: English
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
	// The number of entries per page in a paged query.
	//
	// Default value: 6. Maximum value: 100.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The public IP address of the ECS instance that initiates the outbound connection.
	//
	// example:
	//
	// 192.0.XX.XX
	PublicIP *string `json:"PublicIP,omitempty" xml:"PublicIP,omitempty"`
	// The field by which to sort the results. Valid values:
	//
	// - **SessionCount*	- (default): the number of requests.
	//
	// - **TotalBytes**: the total traffic volume.
	//
	// example:
	//
	// SessionCount
	Sort *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// The start time of the query. The value is a UNIX timestamp in seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1656664560
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The intelligence tags label ID. Valid values:
	//
	// - **AliYun**: Alibaba Cloud product.
	//
	// - **RiskDomain**: Risky domain.
	//
	// - **RiskIP**: Risky IP.
	//
	// - **TrustedDomain**: Trusted website.
	//
	// - **AliPay**: Alipay.
	//
	// - **DingDing**: DingTalk.
	//
	// - **WeChat**: WeChat.
	//
	// - **QQ**: Tencent QQ.
	//
	// - **SecurityService**: Security service.
	//
	// - **Microsoft**: Microsoft.
	//
	// - **Amazon**: Amazon.
	//
	// - **Pan**: Cloud drive.
	//
	// - **Map**: Map.
	//
	// - **Code**: Code hosting.
	//
	// - **SystemService**: System service.
	//
	// - **Taobao**: Taobao.
	//
	// - **Google**: Google.
	//
	// - **ThirdPartyService**: Third-party platform service.
	//
	// - **FirstFlow**: First Visit.
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
	// - **Suspicious**: Suspicious website.
	//
	// - **C&C**: Remote control.
	//
	// - **Gang**: Gang.
	//
	// - **CVE**: CVE vulnerability.
	//
	// - **Backdoor**: Backdoor Trojan.
	//
	// - **Phishing**: Phishing website.
	//
	// - **APT**: APT attack.
	//
	// - **Supply Chain Attack**: Supply chain attack.
	//
	// - **Malicious software**: Malware.
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
