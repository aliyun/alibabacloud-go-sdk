// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCloudSiemCustomizeRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertType(v string) *ListCloudSiemCustomizeRulesRequest
	GetAlertType() *string
	SetCurrentPage(v int32) *ListCloudSiemCustomizeRulesRequest
	GetCurrentPage() *int32
	SetEndTime(v int64) *ListCloudSiemCustomizeRulesRequest
	GetEndTime() *int64
	SetId(v string) *ListCloudSiemCustomizeRulesRequest
	GetId() *string
	SetOrder(v string) *ListCloudSiemCustomizeRulesRequest
	GetOrder() *string
	SetOrderField(v string) *ListCloudSiemCustomizeRulesRequest
	GetOrderField() *string
	SetPageSize(v int32) *ListCloudSiemCustomizeRulesRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListCloudSiemCustomizeRulesRequest
	GetRegionId() *string
	SetRoleFor(v int64) *ListCloudSiemCustomizeRulesRequest
	GetRoleFor() *int64
	SetRoleType(v int32) *ListCloudSiemCustomizeRulesRequest
	GetRoleType() *int32
	SetRuleName(v string) *ListCloudSiemCustomizeRulesRequest
	GetRuleName() *string
	SetRuleType(v string) *ListCloudSiemCustomizeRulesRequest
	GetRuleType() *string
	SetStartTime(v int64) *ListCloudSiemCustomizeRulesRequest
	GetStartTime() *int64
	SetStatus(v int32) *ListCloudSiemCustomizeRulesRequest
	GetStatus() *int32
	SetThreatLevel(v []*string) *ListCloudSiemCustomizeRulesRequest
	GetThreatLevel() []*string
}

type ListCloudSiemCustomizeRulesRequest struct {
	// The alert type.
	//
	// example:
	//
	// scan
	AlertType *string `json:"AlertType,omitempty" xml:"AlertType,omitempty"`
	// The page number. Pages start from page 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The end time of the query. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1577808000000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the custom rule.
	//
	// example:
	//
	// 10223
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The order in which you want to sort the custom rules. Valid values:
	//
	// - desc: descending order
	//
	// - asc: ascending order
	//
	// example:
	//
	// desc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The field that you use to sort the custom rules. Valid values:
	//
	// - GmtModified: sorts the rules by modification time.
	//
	// - Id: sorts the rules by rule ID. This is the default value.
	//
	// example:
	//
	// Id
	OrderField *string `json:"OrderField,omitempty" xml:"OrderField,omitempty"`
	// The number of entries per page. Maximum value: 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region where the data management center of Threat Analysis is deployed. You must select the region where your assets reside. Valid values:
	//
	// - **cn-hangzhou**: your assets are in the Chinese mainland or Hong Kong (China).
	//
	// - **ap-southeast-1**: your assets are outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The user ID that is used to switch the administrator\\"s view to the view of a member.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The view type.
	//
	// - 0: the view of the current Alibaba Cloud account.
	//
	// - 1: the view of all members in the enterprise.
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	// The name of the rule. The name can contain letters, digits, underscores (_), and periods (.).
	//
	// example:
	//
	// waf_scan
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The type of the rule. Valid values:
	//
	// - **predefine**: predefined
	//
	// - **customize**: custom
	//
	// example:
	//
	// customize
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	// The start time of the query. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1577808000000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the rule. Valid values:
	//
	// - **0**: initial
	//
	// - **10**: testing with simulated data
	//
	// - **15**: testing with business data
	//
	// - **20**: testing with business data is complete
	//
	// - **100**: published
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The threat level. The value is a JSON array. Valid values:
	//
	// - **serious**: high
	//
	// - **suspicious**: medium
	//
	// - **remind**: low
	//
	// example:
	//
	// ["serious","suspicious","remind"]
	ThreatLevel []*string `json:"ThreatLevel,omitempty" xml:"ThreatLevel,omitempty" type:"Repeated"`
}

func (s ListCloudSiemCustomizeRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCloudSiemCustomizeRulesRequest) GoString() string {
	return s.String()
}

func (s *ListCloudSiemCustomizeRulesRequest) GetAlertType() *string {
	return s.AlertType
}

func (s *ListCloudSiemCustomizeRulesRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListCloudSiemCustomizeRulesRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListCloudSiemCustomizeRulesRequest) GetId() *string {
	return s.Id
}

func (s *ListCloudSiemCustomizeRulesRequest) GetOrder() *string {
	return s.Order
}

func (s *ListCloudSiemCustomizeRulesRequest) GetOrderField() *string {
	return s.OrderField
}

func (s *ListCloudSiemCustomizeRulesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCloudSiemCustomizeRulesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListCloudSiemCustomizeRulesRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *ListCloudSiemCustomizeRulesRequest) GetRoleType() *int32 {
	return s.RoleType
}

func (s *ListCloudSiemCustomizeRulesRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *ListCloudSiemCustomizeRulesRequest) GetRuleType() *string {
	return s.RuleType
}

func (s *ListCloudSiemCustomizeRulesRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListCloudSiemCustomizeRulesRequest) GetStatus() *int32 {
	return s.Status
}

func (s *ListCloudSiemCustomizeRulesRequest) GetThreatLevel() []*string {
	return s.ThreatLevel
}

func (s *ListCloudSiemCustomizeRulesRequest) SetAlertType(v string) *ListCloudSiemCustomizeRulesRequest {
	s.AlertType = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesRequest) SetCurrentPage(v int32) *ListCloudSiemCustomizeRulesRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesRequest) SetEndTime(v int64) *ListCloudSiemCustomizeRulesRequest {
	s.EndTime = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesRequest) SetId(v string) *ListCloudSiemCustomizeRulesRequest {
	s.Id = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesRequest) SetOrder(v string) *ListCloudSiemCustomizeRulesRequest {
	s.Order = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesRequest) SetOrderField(v string) *ListCloudSiemCustomizeRulesRequest {
	s.OrderField = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesRequest) SetPageSize(v int32) *ListCloudSiemCustomizeRulesRequest {
	s.PageSize = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesRequest) SetRegionId(v string) *ListCloudSiemCustomizeRulesRequest {
	s.RegionId = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesRequest) SetRoleFor(v int64) *ListCloudSiemCustomizeRulesRequest {
	s.RoleFor = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesRequest) SetRoleType(v int32) *ListCloudSiemCustomizeRulesRequest {
	s.RoleType = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesRequest) SetRuleName(v string) *ListCloudSiemCustomizeRulesRequest {
	s.RuleName = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesRequest) SetRuleType(v string) *ListCloudSiemCustomizeRulesRequest {
	s.RuleType = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesRequest) SetStartTime(v int64) *ListCloudSiemCustomizeRulesRequest {
	s.StartTime = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesRequest) SetStatus(v int32) *ListCloudSiemCustomizeRulesRequest {
	s.Status = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesRequest) SetThreatLevel(v []*string) *ListCloudSiemCustomizeRulesRequest {
	s.ThreatLevel = v
	return s
}

func (s *ListCloudSiemCustomizeRulesRequest) Validate() error {
	return dara.Validate(s)
}
