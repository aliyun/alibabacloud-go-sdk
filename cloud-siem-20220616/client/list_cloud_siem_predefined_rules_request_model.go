// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCloudSiemPredefinedRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertType(v string) *ListCloudSiemPredefinedRulesRequest
	GetAlertType() *string
	SetAttCk(v string) *ListCloudSiemPredefinedRulesRequest
	GetAttCk() *string
	SetCurrentPage(v int32) *ListCloudSiemPredefinedRulesRequest
	GetCurrentPage() *int32
	SetEndTime(v int64) *ListCloudSiemPredefinedRulesRequest
	GetEndTime() *int64
	SetEventTransferType(v string) *ListCloudSiemPredefinedRulesRequest
	GetEventTransferType() *string
	SetId(v string) *ListCloudSiemPredefinedRulesRequest
	GetId() *string
	SetLogSource(v string) *ListCloudSiemPredefinedRulesRequest
	GetLogSource() *string
	SetOrder(v string) *ListCloudSiemPredefinedRulesRequest
	GetOrder() *string
	SetOrderField(v string) *ListCloudSiemPredefinedRulesRequest
	GetOrderField() *string
	SetPageSize(v int32) *ListCloudSiemPredefinedRulesRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListCloudSiemPredefinedRulesRequest
	GetRegionId() *string
	SetRoleFor(v int64) *ListCloudSiemPredefinedRulesRequest
	GetRoleFor() *int64
	SetRoleType(v int32) *ListCloudSiemPredefinedRulesRequest
	GetRoleType() *int32
	SetRuleName(v string) *ListCloudSiemPredefinedRulesRequest
	GetRuleName() *string
	SetRuleType(v string) *ListCloudSiemPredefinedRulesRequest
	GetRuleType() *string
	SetStartTime(v int64) *ListCloudSiemPredefinedRulesRequest
	GetStartTime() *int64
	SetStatus(v int32) *ListCloudSiemPredefinedRulesRequest
	GetStatus() *int32
	SetThreatLevel(v []*string) *ListCloudSiemPredefinedRulesRequest
	GetThreatLevel() []*string
}

type ListCloudSiemPredefinedRulesRequest struct {
	// The alert type.
	//
	// example:
	//
	// scan
	AlertType *string `json:"AlertType,omitempty" xml:"AlertType,omitempty"`
	// The ATT\\&CK information.
	//
	// example:
	//
	// T1595.002 Vulnerability Scanning
	AttCk *string `json:"AttCk,omitempty" xml:"AttCk,omitempty"`
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
	// The method that is used to generate an event. Valid values:
	//
	// 	- default: built-in method.
	//
	// 	- singleToSingle: The system generates an event for each alert.
	//
	// 	- allToSingle: The system generates an event for alerts within a period of time.
	//
	// example:
	//
	// allToSingle
	EventTransferType *string `json:"EventTransferType,omitempty" xml:"EventTransferType,omitempty"`
	// The ID of the rule.
	//
	// example:
	//
	// 10223
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The log source.
	//
	// example:
	//
	// cloud_siem_aegis_sas_alert
	LogSource *string `json:"LogSource,omitempty" xml:"LogSource,omitempty"`
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
	// The number of entries per page. Maximum value: 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region in which the data management center of the threat analysis feature resides. Specify this parameter based on the regions in which your assets reside. Valid values:
	//
	// 	- cn-hangzhou: Your assets reside in regions in China.
	//
	// 	- ap-southeast-1: Your assets reside in regions outside China.
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
	// The type of the view.
	//
	// 	- 0: view of the current Alibaba Cloud account.
	//
	// 	- 1: view of all accounts for the enterprise.
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
	// 	- predefine
	//
	// 	- customize
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
	// 	- 0: The rule is in the initial state.
	//
	// 	- 10: The simulation data is tested.
	//
	// 	- 15: The business data is being tested.
	//
	// 	- 20: The business data test ends.
	//
	// 	- 100: The rule takes effect.
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The risk level. The value is a JSON array. Valid values:
	//
	// 	- serious: high
	//
	// 	- suspicious: medium
	//
	// 	- remind: low
	//
	// example:
	//
	// ["serious","suspicious","remind"]
	ThreatLevel []*string `json:"ThreatLevel,omitempty" xml:"ThreatLevel,omitempty" type:"Repeated"`
}

func (s ListCloudSiemPredefinedRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCloudSiemPredefinedRulesRequest) GoString() string {
	return s.String()
}

func (s *ListCloudSiemPredefinedRulesRequest) GetAlertType() *string {
	return s.AlertType
}

func (s *ListCloudSiemPredefinedRulesRequest) GetAttCk() *string {
	return s.AttCk
}

func (s *ListCloudSiemPredefinedRulesRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListCloudSiemPredefinedRulesRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListCloudSiemPredefinedRulesRequest) GetEventTransferType() *string {
	return s.EventTransferType
}

func (s *ListCloudSiemPredefinedRulesRequest) GetId() *string {
	return s.Id
}

func (s *ListCloudSiemPredefinedRulesRequest) GetLogSource() *string {
	return s.LogSource
}

func (s *ListCloudSiemPredefinedRulesRequest) GetOrder() *string {
	return s.Order
}

func (s *ListCloudSiemPredefinedRulesRequest) GetOrderField() *string {
	return s.OrderField
}

func (s *ListCloudSiemPredefinedRulesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCloudSiemPredefinedRulesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListCloudSiemPredefinedRulesRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *ListCloudSiemPredefinedRulesRequest) GetRoleType() *int32 {
	return s.RoleType
}

func (s *ListCloudSiemPredefinedRulesRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *ListCloudSiemPredefinedRulesRequest) GetRuleType() *string {
	return s.RuleType
}

func (s *ListCloudSiemPredefinedRulesRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListCloudSiemPredefinedRulesRequest) GetStatus() *int32 {
	return s.Status
}

func (s *ListCloudSiemPredefinedRulesRequest) GetThreatLevel() []*string {
	return s.ThreatLevel
}

func (s *ListCloudSiemPredefinedRulesRequest) SetAlertType(v string) *ListCloudSiemPredefinedRulesRequest {
	s.AlertType = &v
	return s
}

func (s *ListCloudSiemPredefinedRulesRequest) SetAttCk(v string) *ListCloudSiemPredefinedRulesRequest {
	s.AttCk = &v
	return s
}

func (s *ListCloudSiemPredefinedRulesRequest) SetCurrentPage(v int32) *ListCloudSiemPredefinedRulesRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListCloudSiemPredefinedRulesRequest) SetEndTime(v int64) *ListCloudSiemPredefinedRulesRequest {
	s.EndTime = &v
	return s
}

func (s *ListCloudSiemPredefinedRulesRequest) SetEventTransferType(v string) *ListCloudSiemPredefinedRulesRequest {
	s.EventTransferType = &v
	return s
}

func (s *ListCloudSiemPredefinedRulesRequest) SetId(v string) *ListCloudSiemPredefinedRulesRequest {
	s.Id = &v
	return s
}

func (s *ListCloudSiemPredefinedRulesRequest) SetLogSource(v string) *ListCloudSiemPredefinedRulesRequest {
	s.LogSource = &v
	return s
}

func (s *ListCloudSiemPredefinedRulesRequest) SetOrder(v string) *ListCloudSiemPredefinedRulesRequest {
	s.Order = &v
	return s
}

func (s *ListCloudSiemPredefinedRulesRequest) SetOrderField(v string) *ListCloudSiemPredefinedRulesRequest {
	s.OrderField = &v
	return s
}

func (s *ListCloudSiemPredefinedRulesRequest) SetPageSize(v int32) *ListCloudSiemPredefinedRulesRequest {
	s.PageSize = &v
	return s
}

func (s *ListCloudSiemPredefinedRulesRequest) SetRegionId(v string) *ListCloudSiemPredefinedRulesRequest {
	s.RegionId = &v
	return s
}

func (s *ListCloudSiemPredefinedRulesRequest) SetRoleFor(v int64) *ListCloudSiemPredefinedRulesRequest {
	s.RoleFor = &v
	return s
}

func (s *ListCloudSiemPredefinedRulesRequest) SetRoleType(v int32) *ListCloudSiemPredefinedRulesRequest {
	s.RoleType = &v
	return s
}

func (s *ListCloudSiemPredefinedRulesRequest) SetRuleName(v string) *ListCloudSiemPredefinedRulesRequest {
	s.RuleName = &v
	return s
}

func (s *ListCloudSiemPredefinedRulesRequest) SetRuleType(v string) *ListCloudSiemPredefinedRulesRequest {
	s.RuleType = &v
	return s
}

func (s *ListCloudSiemPredefinedRulesRequest) SetStartTime(v int64) *ListCloudSiemPredefinedRulesRequest {
	s.StartTime = &v
	return s
}

func (s *ListCloudSiemPredefinedRulesRequest) SetStatus(v int32) *ListCloudSiemPredefinedRulesRequest {
	s.Status = &v
	return s
}

func (s *ListCloudSiemPredefinedRulesRequest) SetThreatLevel(v []*string) *ListCloudSiemPredefinedRulesRequest {
	s.ThreatLevel = v
	return s
}

func (s *ListCloudSiemPredefinedRulesRequest) Validate() error {
	return dara.Validate(s)
}
