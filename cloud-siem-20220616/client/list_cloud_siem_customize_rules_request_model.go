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
	// The end of the time range to query. Unit: milliseconds.
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
	// The sort method. Valid values:
	//
	// 	- desc: descending order.
	//
	// 	- asc: ascending order.
	//
	// example:
	//
	// desc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The field that is used to sort the rules. Valid values:
	//
	// 	- GmtModified: The rules are sorted based on the modification time.
	//
	// 	- Id (default): The rules are sorted based on the rule ID.
	//
	// example:
	//
	// Id
	OrderField *string `json:"OrderField,omitempty" xml:"OrderField,omitempty"`
	// The number of entries per page. The value can be up to 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The data management center of the threat analysis feature. Specify this parameter based on the regions in which your assets reside. Valid values:
	//
	// 	- **cn-hangzhou**: Your assets reside in regions in China.
	//
	// 	- **ap-southeast-1**: Your assets reside in regions outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the destination account to which you switch the view from the management account.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of the view. Valid values:
	//
	// 	- 0: view of the current Alibaba Cloud account.
	//
	// 	- 1: view of all accounts for the enterprise.
	//
	// example:
	//
	// 0
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	// The name of the rule. The name can contain letters, digits, underscores (_), and periods (.).
	//
	// example:
	//
	// waf_scan
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The type of the rule. Valid values:
	//
	// 	- **predefine**
	//
	// 	- **customize**
	//
	// example:
	//
	// customize
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	// The beginning of the time range to query. Unit: milliseconds.
	//
	// example:
	//
	// 1577808000000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the rule. Valid values:
	//
	// 	- **0**: The rule is in the initial state.
	//
	// 	- **10**: The simulation data is tested.
	//
	// 	- **15**: The business data is being tested.
	//
	// 	- **20**: The business data test is complete.
	//
	// 	- **100**: The rule is in effect.
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The threat level. The value must be a JSON array. Valid values:
	//
	// 	- **serious**: high-risk.
	//
	// 	- **suspicious**: medium-risk.
	//
	// 	- **remind**: low-risk.
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
