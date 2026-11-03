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
	// The list of names of the attacked applications. Specify the value in the `["AttackApp1","AttackApp2"]` format.
	//
	// example:
	//
	// ["MySql","DNS"]
	AttackApp []*string `json:"AttackApp,omitempty" xml:"AttackApp,omitempty" type:"Repeated"`
	// The list of categories of the attacked applications. Specify the value in the `["AttackAppCategory1","AttackAppCategory2"]` format.
	AttackAppCategory []*string `json:"AttackAppCategory,omitempty" xml:"AttackAppCategory,omitempty" type:"Repeated"`
	// The attack type of the intrusion prevention events. Valid values:
	//
	// - **1**: abnormal connection.
	//
	// - **2**: command execute.
	//
	// - **3**: brute-force attacks.
	//
	// - **4**: scan.
	//
	// - **5**: other.
	//
	// - **6**: information leak.
	//
	// - **7**: DoS attack.
	//
	// - **8**: overflow attack.
	//
	// - **9**: web attack.
	//
	// - **10**: trojan backdoor.
	//
	// - **11**: virus worm.
	//
	// - **12**: mining behavior.
	//
	// - **13**: reverse shell.
	//
	// > If the Settings for this parameter are not specified, events of all attack types are queried.
	//
	// example:
	//
	// 1
	AttackType *string `json:"AttackType,omitempty" xml:"AttackType,omitempty"`
	// The purchased edition. Valid values:
	//
	// - **2**: Premium Edition.
	//
	// - **3**: Enterprise Edition.
	//
	// - **4**: Ultimate Edition.
	//
	// - **10**: Pay-as-you-go Edition.
	//
	// example:
	//
	// 10
	BuyVersion *int64 `json:"BuyVersion,omitempty" xml:"BuyVersion,omitempty"`
	// The paging Settings for the page number to return in a paged query.
	//
	// Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The risk event type. Valid value:
	//
	// - **session**: intrusion prevention event.
	//
	// This parameter is required.
	//
	// example:
	//
	// session
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// The traffic direction of the intrusion prevention events. Valid values:
	//
	// - **in**: inbound.
	//
	// - **out**: outbound.
	//
	// > If this parameter is not specified, events in all traffic directions are queried.
	//
	// example:
	//
	// in
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The destination IP address to query. If you specify this parameter, intrusion prevention events that involve the specified destination IP address are queried.
	//
	// example:
	//
	// 192.0.XX.XX
	DstIP *string `json:"DstIP,omitempty" xml:"DstIP,omitempty"`
	// The instance ID of the destination VPC-connected instance.
	//
	// > This parameter is required only when the FirewallType Settings is VpcFirewall.
	//
	// example:
	//
	// vpc-uf6e9a9zyokj2ywuo****
	DstNetworkInstanceId *string `json:"DstNetworkInstanceId,omitempty" xml:"DstNetworkInstanceId,omitempty"`
	// The end time of the query. The value is a UNIX timestamp in seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1534408267
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The name of the intrusion prevention event.
	//
	// example:
	//
	// wooden horse rear door communication
	EventName *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	// The type of the cloud firewall. Valid values:
	//
	// - **VpcFirewall**: virtual private cloud (VPC) firewalls.
	//
	// - **InternetFirewall*	- (default): the Internet Border firewall.
	//
	// example:
	//
	// InternetFirewall
	FirewallType *string `json:"FirewallType,omitempty" xml:"FirewallType,omitempty"`
	// Specifies whether to query only the data for which private IP tracing is completed.
	//
	// example:
	//
	// true
	IsOnlyPrivateAssoc *string `json:"IsOnlyPrivateAssoc,omitempty" xml:"IsOnlyPrivateAssoc,omitempty"`
	// The language of the request and response. Valid values:
	//
	// - **zh*	- (default): Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Specifies whether to query IP geolocation information. Valid values:
	//
	// - **true**: IP geolocation information is not queried.
	//
	// - **false*	- (default): IP geolocation information is queried.
	//
	// example:
	//
	// false
	NoLocation *string `json:"NoLocation,omitempty" xml:"NoLocation,omitempty"`
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
	// The paging Settings for the number of entries to return on each page in a paged query.
	//
	// Default value: **6**. Maximum value: **10**.
	//
	// example:
	//
	// 6
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The defense status of Cloud Firewall. Valid values:
	//
	// - **1**: alert.
	//
	// - **2**: block.
	//
	// > If this parameter is not specified, events in all defense statuses are queried.
	//
	// example:
	//
	// 1
	RuleResult *string `json:"RuleResult,omitempty" xml:"RuleResult,omitempty"`
	// The source of the detection rule for the intrusion prevention events. Valid values:
	//
	// - **1**: basic protection.
	//
	// - **2**: virtual patches.
	//
	// - **4**: threat intelligence.
	//
	// > If this parameter is not specified, events detected by rules from all sources are queried.
	//
	// example:
	//
	// 1
	RuleSource *string `json:"RuleSource,omitempty" xml:"RuleSource,omitempty"`
	// The field by which to sort the results. Valid values:
	//
	// - **VulLevel*	- (default): sort by severity level.
	//
	// - **LastTime**: sort by the most recent occurrence time.
	//
	// example:
	//
	// LastTime
	Sort *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// The source IP address to query. If you specify this parameter, intrusion prevention events that involve the specified source IP address are queried.
	//
	// example:
	//
	// 192.0.XX.XX
	SrcIP *string `json:"SrcIP,omitempty" xml:"SrcIP,omitempty"`
	// The instance ID of the source VPC-connected instance.
	//
	// > This parameter is required only when the FirewallType Settings is VpcFirewall.
	//
	// example:
	//
	// vpc-uf6e9a9zyokj2ywuo****
	SrcNetworkInstanceId *string `json:"SrcNetworkInstanceId,omitempty" xml:"SrcNetworkInstanceId,omitempty"`
	// The start time of the query. The value is a UNIX timestamp in seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1534408189
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The severity level of the intrusion prevention events. Valid values:
	//
	// - **1**: low.
	//
	// - **2**: medium.
	//
	// - **3**: high.
	//
	// > If this parameter is not specified, events at all severity levels are queried.
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
