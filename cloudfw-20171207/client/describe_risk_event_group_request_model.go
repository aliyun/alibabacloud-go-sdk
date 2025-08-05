// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRiskEventGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttackApp(v []*string) *DescribeRiskEventGroupRequest
	GetAttackApp() []*string
	SetAttackAppCategory(v []*string) *DescribeRiskEventGroupRequest
	GetAttackAppCategory() []*string
	SetAttackType(v string) *DescribeRiskEventGroupRequest
	GetAttackType() *string
	SetBuyVersion(v int64) *DescribeRiskEventGroupRequest
	GetBuyVersion() *int64
	SetCurrentPage(v string) *DescribeRiskEventGroupRequest
	GetCurrentPage() *string
	SetDataType(v string) *DescribeRiskEventGroupRequest
	GetDataType() *string
	SetDirection(v string) *DescribeRiskEventGroupRequest
	GetDirection() *string
	SetDstIP(v string) *DescribeRiskEventGroupRequest
	GetDstIP() *string
	SetDstNetworkInstanceId(v string) *DescribeRiskEventGroupRequest
	GetDstNetworkInstanceId() *string
	SetEndTime(v string) *DescribeRiskEventGroupRequest
	GetEndTime() *string
	SetEventName(v string) *DescribeRiskEventGroupRequest
	GetEventName() *string
	SetFirewallType(v string) *DescribeRiskEventGroupRequest
	GetFirewallType() *string
	SetIsOnlyPrivateAssoc(v string) *DescribeRiskEventGroupRequest
	GetIsOnlyPrivateAssoc() *string
	SetLang(v string) *DescribeRiskEventGroupRequest
	GetLang() *string
	SetNoLocation(v string) *DescribeRiskEventGroupRequest
	GetNoLocation() *string
	SetOrder(v string) *DescribeRiskEventGroupRequest
	GetOrder() *string
	SetPageSize(v string) *DescribeRiskEventGroupRequest
	GetPageSize() *string
	SetRuleResult(v string) *DescribeRiskEventGroupRequest
	GetRuleResult() *string
	SetRuleSource(v string) *DescribeRiskEventGroupRequest
	GetRuleSource() *string
	SetSort(v string) *DescribeRiskEventGroupRequest
	GetSort() *string
	SetSrcIP(v string) *DescribeRiskEventGroupRequest
	GetSrcIP() *string
	SetSrcNetworkInstanceId(v string) *DescribeRiskEventGroupRequest
	GetSrcNetworkInstanceId() *string
	SetStartTime(v string) *DescribeRiskEventGroupRequest
	GetStartTime() *string
	SetVulLevel(v string) *DescribeRiskEventGroupRequest
	GetVulLevel() *string
}

type DescribeRiskEventGroupRequest struct {
	// The names of the attacked applications. Set the value in the `["AttackApp1","AttackApp2"]` format.
	//
	// example:
	//
	// ["MySql","DNS"]
	AttackApp []*string `json:"AttackApp,omitempty" xml:"AttackApp,omitempty" type:"Repeated"`
	// A list of categories of attacked applications, expressed in the format ["AttackAppCategory1","AttackAppCategory2"].
	AttackAppCategory []*string `json:"AttackAppCategory,omitempty" xml:"AttackAppCategory,omitempty" type:"Repeated"`
	// The attack type of the intrusion events. Valid values:
	//
	// 	- **1**: suspicious connection
	//
	// 	- **2**: command execution
	//
	// 	- **3**: brute-force attack
	//
	// 	- **4**: scanning
	//
	// 	- **5**: others
	//
	// 	- **6**: information leak
	//
	// 	- **7**: DoS attack
	//
	// 	- **8**: buffer overflow attack
	//
	// 	- **9**: web attack
	//
	// 	- **10**: trojan backdoor
	//
	// 	- **11**: computer worm
	//
	// 	- **12**: mining
	//
	// 	- **13**: reverse shell
	//
	// > If you do not specify this parameter, the intrusion events of all attack types are queried.
	//
	// example:
	//
	// 1
	AttackType *string `json:"AttackType,omitempty" xml:"AttackType,omitempty"`
	// The edition of Cloud Firewall that you purchase. Valid values:
	//
	// 	- **2**: Premium Edition
	//
	// 	- **3**: Enterprise Edition
	//
	// 	- **4**: Ultimate Edition
	//
	// 	- **10**: Cloud Firewall that uses the pay-as-you-go billing method
	//
	// example:
	//
	// 10
	BuyVersion *int64 `json:"BuyVersion,omitempty" xml:"BuyVersion,omitempty"`
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The type of the risk events.\\
	//
	// Set the value to **session**, which indicates intrusion events.
	//
	// This parameter is required.
	//
	// example:
	//
	// session
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// The direction of the traffic for the intrusion events. Valid values:
	//
	// 	- **in**: inbound
	//
	// 	- **out**: outbound
	//
	// > If you do not specify this parameter, the intrusion events that are recorded for both inbound and outbound traffic are queried.
	//
	// example:
	//
	// in
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The destination IP address to query. If you specify this parameter, all intrusion events with the specified destination IP address are queried.
	//
	// example:
	//
	// 192.0.XX.XX
	DstIP *string `json:"DstIP,omitempty" xml:"DstIP,omitempty"`
	// The ID of the destination VPC.
	//
	// > If the FirewallType parameter is set to VpcFirewall, you must specify this parameter.
	//
	// example:
	//
	// vpc-uf6e9a9zyokj2ywuo****
	DstNetworkInstanceId *string `json:"DstNetworkInstanceId,omitempty" xml:"DstNetworkInstanceId,omitempty"`
	// The end of the time range to query. The value is a UNIX timestamp. Unit: seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1534408267
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The name of the intrusion event.
	//
	// example:
	//
	// Webshell communication
	EventName *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	// The type of the firewall. Valid values:
	//
	// 	- **VpcFirewall**: virtual private cloud (VPC) firewall
	//
	// 	- **InternetFirewall**: Internet firewall (default)
	//
	// example:
	//
	// InternetFirewall
	FirewallType *string `json:"FirewallType,omitempty" xml:"FirewallType,omitempty"`
	// Whether to query only the data that has completed private network tracing.
	//
	// example:
	//
	// true
	IsOnlyPrivateAssoc *string `json:"IsOnlyPrivateAssoc,omitempty" xml:"IsOnlyPrivateAssoc,omitempty"`
	// The language of the content within the request and response. Valid values:
	//
	// 	- **zh**: Chinese (default)
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Specifies whether to query the information about the geographical locations of IP addresses.
	//
	// 	- **true**: does not query the information about the geographical locations of IP addresses.
	//
	// 	- **false**: queries the information about the geographical locations of IP addresses. This is the default value.
	//
	// example:
	//
	// false
	NoLocation *string `json:"NoLocation,omitempty" xml:"NoLocation,omitempty"`
	// The order in which you want to sort the results. Valid values:
	//
	// 	- **asc**: the ascending order.
	//
	// 	- **desc**: the descending order. This is the default value.
	//
	// example:
	//
	// desc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The number of entries to return on each page.
	//
	// Default value: **6**. Maximum value: **10**.
	//
	// example:
	//
	// 6
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The status of the firewall. Valid values:
	//
	// 	- **1**: alerting
	//
	// 	- **2**: blocking
	//
	// > If you do not specify this parameter, all intrusion events that are detected by the firewall are queried, regardless of the firewall status.
	//
	// example:
	//
	// 1
	RuleResult *string `json:"RuleResult,omitempty" xml:"RuleResult,omitempty"`
	// The module of the rule that is used to detect the intrusion events. Valid values:
	//
	// 	- **1**: basic protection
	//
	// 	- **2**: virtual patching
	//
	// 	- **4**: threat intelligence
	//
	// > If you do not specify this parameter, the intrusion events that are detected by all rules are queried.
	//
	// example:
	//
	// 1
	RuleSource *string `json:"RuleSource,omitempty" xml:"RuleSource,omitempty"`
	// The field based on which you want to sort the results. Valid values:
	//
	// 	- **VulLevel**: The results are sorted based on the risk level field. This is the default value.
	//
	// 	- **LastTime**: The results are sorted based on the most recent occurrence time.
	//
	// example:
	//
	// LastTime
	Sort *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// The source IP address to query. If you specify this parameter, all intrusion events with the specified source IP address are queried.
	//
	// example:
	//
	// 192.0.XX.XX
	SrcIP *string `json:"SrcIP,omitempty" xml:"SrcIP,omitempty"`
	// The ID of the source VPC.
	//
	// > If the FirewallType parameter is set to VpcFirewall, you must specify this parameter.
	//
	// example:
	//
	// vpc-uf6e9a9zyokj2ywuo****
	SrcNetworkInstanceId *string `json:"SrcNetworkInstanceId,omitempty" xml:"SrcNetworkInstanceId,omitempty"`
	// The beginning of the time range to query. The value is a UNIX timestamp. Unit: seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1534408189
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The risk level of the intrusion events. Valid values:
	//
	// 	- **1**: low
	//
	// 	- **2**: medium
	//
	// 	- **3**: high
	//
	// > If you do not specify this parameter, the intrusion events that are at all risk levels are queried.
	//
	// example:
	//
	// 1
	VulLevel *string `json:"VulLevel,omitempty" xml:"VulLevel,omitempty"`
}

func (s DescribeRiskEventGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRiskEventGroupRequest) GoString() string {
	return s.String()
}

func (s *DescribeRiskEventGroupRequest) GetAttackApp() []*string {
	return s.AttackApp
}

func (s *DescribeRiskEventGroupRequest) GetAttackAppCategory() []*string {
	return s.AttackAppCategory
}

func (s *DescribeRiskEventGroupRequest) GetAttackType() *string {
	return s.AttackType
}

func (s *DescribeRiskEventGroupRequest) GetBuyVersion() *int64 {
	return s.BuyVersion
}

func (s *DescribeRiskEventGroupRequest) GetCurrentPage() *string {
	return s.CurrentPage
}

func (s *DescribeRiskEventGroupRequest) GetDataType() *string {
	return s.DataType
}

func (s *DescribeRiskEventGroupRequest) GetDirection() *string {
	return s.Direction
}

func (s *DescribeRiskEventGroupRequest) GetDstIP() *string {
	return s.DstIP
}

func (s *DescribeRiskEventGroupRequest) GetDstNetworkInstanceId() *string {
	return s.DstNetworkInstanceId
}

func (s *DescribeRiskEventGroupRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeRiskEventGroupRequest) GetEventName() *string {
	return s.EventName
}

func (s *DescribeRiskEventGroupRequest) GetFirewallType() *string {
	return s.FirewallType
}

func (s *DescribeRiskEventGroupRequest) GetIsOnlyPrivateAssoc() *string {
	return s.IsOnlyPrivateAssoc
}

func (s *DescribeRiskEventGroupRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeRiskEventGroupRequest) GetNoLocation() *string {
	return s.NoLocation
}

func (s *DescribeRiskEventGroupRequest) GetOrder() *string {
	return s.Order
}

func (s *DescribeRiskEventGroupRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeRiskEventGroupRequest) GetRuleResult() *string {
	return s.RuleResult
}

func (s *DescribeRiskEventGroupRequest) GetRuleSource() *string {
	return s.RuleSource
}

func (s *DescribeRiskEventGroupRequest) GetSort() *string {
	return s.Sort
}

func (s *DescribeRiskEventGroupRequest) GetSrcIP() *string {
	return s.SrcIP
}

func (s *DescribeRiskEventGroupRequest) GetSrcNetworkInstanceId() *string {
	return s.SrcNetworkInstanceId
}

func (s *DescribeRiskEventGroupRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeRiskEventGroupRequest) GetVulLevel() *string {
	return s.VulLevel
}

func (s *DescribeRiskEventGroupRequest) SetAttackApp(v []*string) *DescribeRiskEventGroupRequest {
	s.AttackApp = v
	return s
}

func (s *DescribeRiskEventGroupRequest) SetAttackAppCategory(v []*string) *DescribeRiskEventGroupRequest {
	s.AttackAppCategory = v
	return s
}

func (s *DescribeRiskEventGroupRequest) SetAttackType(v string) *DescribeRiskEventGroupRequest {
	s.AttackType = &v
	return s
}

func (s *DescribeRiskEventGroupRequest) SetBuyVersion(v int64) *DescribeRiskEventGroupRequest {
	s.BuyVersion = &v
	return s
}

func (s *DescribeRiskEventGroupRequest) SetCurrentPage(v string) *DescribeRiskEventGroupRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeRiskEventGroupRequest) SetDataType(v string) *DescribeRiskEventGroupRequest {
	s.DataType = &v
	return s
}

func (s *DescribeRiskEventGroupRequest) SetDirection(v string) *DescribeRiskEventGroupRequest {
	s.Direction = &v
	return s
}

func (s *DescribeRiskEventGroupRequest) SetDstIP(v string) *DescribeRiskEventGroupRequest {
	s.DstIP = &v
	return s
}

func (s *DescribeRiskEventGroupRequest) SetDstNetworkInstanceId(v string) *DescribeRiskEventGroupRequest {
	s.DstNetworkInstanceId = &v
	return s
}

func (s *DescribeRiskEventGroupRequest) SetEndTime(v string) *DescribeRiskEventGroupRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeRiskEventGroupRequest) SetEventName(v string) *DescribeRiskEventGroupRequest {
	s.EventName = &v
	return s
}

func (s *DescribeRiskEventGroupRequest) SetFirewallType(v string) *DescribeRiskEventGroupRequest {
	s.FirewallType = &v
	return s
}

func (s *DescribeRiskEventGroupRequest) SetIsOnlyPrivateAssoc(v string) *DescribeRiskEventGroupRequest {
	s.IsOnlyPrivateAssoc = &v
	return s
}

func (s *DescribeRiskEventGroupRequest) SetLang(v string) *DescribeRiskEventGroupRequest {
	s.Lang = &v
	return s
}

func (s *DescribeRiskEventGroupRequest) SetNoLocation(v string) *DescribeRiskEventGroupRequest {
	s.NoLocation = &v
	return s
}

func (s *DescribeRiskEventGroupRequest) SetOrder(v string) *DescribeRiskEventGroupRequest {
	s.Order = &v
	return s
}

func (s *DescribeRiskEventGroupRequest) SetPageSize(v string) *DescribeRiskEventGroupRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeRiskEventGroupRequest) SetRuleResult(v string) *DescribeRiskEventGroupRequest {
	s.RuleResult = &v
	return s
}

func (s *DescribeRiskEventGroupRequest) SetRuleSource(v string) *DescribeRiskEventGroupRequest {
	s.RuleSource = &v
	return s
}

func (s *DescribeRiskEventGroupRequest) SetSort(v string) *DescribeRiskEventGroupRequest {
	s.Sort = &v
	return s
}

func (s *DescribeRiskEventGroupRequest) SetSrcIP(v string) *DescribeRiskEventGroupRequest {
	s.SrcIP = &v
	return s
}

func (s *DescribeRiskEventGroupRequest) SetSrcNetworkInstanceId(v string) *DescribeRiskEventGroupRequest {
	s.SrcNetworkInstanceId = &v
	return s
}

func (s *DescribeRiskEventGroupRequest) SetStartTime(v string) *DescribeRiskEventGroupRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeRiskEventGroupRequest) SetVulLevel(v string) *DescribeRiskEventGroupRequest {
	s.VulLevel = &v
	return s
}

func (s *DescribeRiskEventGroupRequest) Validate() error {
	return dara.Validate(s)
}
