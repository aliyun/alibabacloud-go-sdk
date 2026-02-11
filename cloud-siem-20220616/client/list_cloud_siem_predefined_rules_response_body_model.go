// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCloudSiemPredefinedRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListCloudSiemPredefinedRulesResponseBody
	GetCode() *int32
	SetData(v *ListCloudSiemPredefinedRulesResponseBodyData) *ListCloudSiemPredefinedRulesResponseBody
	GetData() *ListCloudSiemPredefinedRulesResponseBodyData
	SetMessage(v string) *ListCloudSiemPredefinedRulesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListCloudSiemPredefinedRulesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListCloudSiemPredefinedRulesResponseBody
	GetSuccess() *bool
}

type ListCloudSiemPredefinedRulesResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	//
	// example:
	//
	// 123456
	Data *ListCloudSiemPredefinedRulesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9AAA9ED9-78F4-5021-86DC-D51C7511****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListCloudSiemPredefinedRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCloudSiemPredefinedRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListCloudSiemPredefinedRulesResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListCloudSiemPredefinedRulesResponseBody) GetData() *ListCloudSiemPredefinedRulesResponseBodyData {
	return s.Data
}

func (s *ListCloudSiemPredefinedRulesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListCloudSiemPredefinedRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCloudSiemPredefinedRulesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListCloudSiemPredefinedRulesResponseBody) SetCode(v int32) *ListCloudSiemPredefinedRulesResponseBody {
	s.Code = &v
	return s
}

func (s *ListCloudSiemPredefinedRulesResponseBody) SetData(v *ListCloudSiemPredefinedRulesResponseBodyData) *ListCloudSiemPredefinedRulesResponseBody {
	s.Data = v
	return s
}

func (s *ListCloudSiemPredefinedRulesResponseBody) SetMessage(v string) *ListCloudSiemPredefinedRulesResponseBody {
	s.Message = &v
	return s
}

func (s *ListCloudSiemPredefinedRulesResponseBody) SetRequestId(v string) *ListCloudSiemPredefinedRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCloudSiemPredefinedRulesResponseBody) SetSuccess(v bool) *ListCloudSiemPredefinedRulesResponseBody {
	s.Success = &v
	return s
}

func (s *ListCloudSiemPredefinedRulesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListCloudSiemPredefinedRulesResponseBodyData struct {
	// The pagination information.
	PageInfo *ListCloudSiemPredefinedRulesResponseBodyDataPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The detailed data.
	ResponseData []*ListCloudSiemPredefinedRulesResponseBodyDataResponseData `json:"ResponseData,omitempty" xml:"ResponseData,omitempty" type:"Repeated"`
}

func (s ListCloudSiemPredefinedRulesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListCloudSiemPredefinedRulesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListCloudSiemPredefinedRulesResponseBodyData) GetPageInfo() *ListCloudSiemPredefinedRulesResponseBodyDataPageInfo {
	return s.PageInfo
}

func (s *ListCloudSiemPredefinedRulesResponseBodyData) GetResponseData() []*ListCloudSiemPredefinedRulesResponseBodyDataResponseData {
	return s.ResponseData
}

func (s *ListCloudSiemPredefinedRulesResponseBodyData) SetPageInfo(v *ListCloudSiemPredefinedRulesResponseBodyDataPageInfo) *ListCloudSiemPredefinedRulesResponseBodyData {
	s.PageInfo = v
	return s
}

func (s *ListCloudSiemPredefinedRulesResponseBodyData) SetResponseData(v []*ListCloudSiemPredefinedRulesResponseBodyDataResponseData) *ListCloudSiemPredefinedRulesResponseBodyData {
	s.ResponseData = v
	return s
}

func (s *ListCloudSiemPredefinedRulesResponseBodyData) Validate() error {
	if s.PageInfo != nil {
		if err := s.PageInfo.Validate(); err != nil {
			return err
		}
	}
	if s.ResponseData != nil {
		for _, item := range s.ResponseData {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCloudSiemPredefinedRulesResponseBodyDataPageInfo struct {
	// The current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCloudSiemPredefinedRulesResponseBodyDataPageInfo) String() string {
	return dara.Prettify(s)
}

func (s ListCloudSiemPredefinedRulesResponseBodyDataPageInfo) GoString() string {
	return s.String()
}

func (s *ListCloudSiemPredefinedRulesResponseBodyDataPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListCloudSiemPredefinedRulesResponseBodyDataPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCloudSiemPredefinedRulesResponseBodyDataPageInfo) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListCloudSiemPredefinedRulesResponseBodyDataPageInfo) SetCurrentPage(v int32) *ListCloudSiemPredefinedRulesResponseBodyDataPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *ListCloudSiemPredefinedRulesResponseBodyDataPageInfo) SetPageSize(v int32) *ListCloudSiemPredefinedRulesResponseBodyDataPageInfo {
	s.PageSize = &v
	return s
}

func (s *ListCloudSiemPredefinedRulesResponseBodyDataPageInfo) SetTotalCount(v int64) *ListCloudSiemPredefinedRulesResponseBodyDataPageInfo {
	s.TotalCount = &v
	return s
}

func (s *ListCloudSiemPredefinedRulesResponseBodyDataPageInfo) Validate() error {
	return dara.Validate(s)
}

type ListCloudSiemPredefinedRulesResponseBodyDataResponseData struct {
	// The type of the risk.
	//
	// example:
	//
	// WEBSHELL
	AlertType *string `json:"AlertType,omitempty" xml:"AlertType,omitempty"`
	// The alert additional field for ATT\\&CK.
	//
	// example:
	//
	// T1595.002 Vulnerability Scanning
	AttCk *string `json:"AttCk,omitempty" xml:"AttCk,omitempty"`
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
	// The time when the rule was created.
	//
	// example:
	//
	// 2021-01-06 16:37:29
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the rule was modified.
	//
	// example:
	//
	// 2021-01-06 16:37:29
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the predefined rule.
	//
	// example:
	//
	// 123456789
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The internal code of the rule description.
	//
	// example:
	//
	// ${siem_rule_description_siem_cfw-attack-count-level-up_cfw-attack}
	RuleDescMds *string `json:"RuleDescMds,omitempty" xml:"RuleDescMds,omitempty"`
	// The name of the rule.
	//
	// example:
	//
	// siem_base64-command-exec_aegis-proc
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The rule name in Chinese.
	//
	// example:
	//
	// siem_base64-command-exec_aegis-proc
	RuleNameCn *string `json:"RuleNameCn,omitempty" xml:"RuleNameCn,omitempty"`
	// The rule name in English.
	//
	// example:
	//
	// siem_base64-command-exec_aegis-proc
	RuleNameEn *string `json:"RuleNameEn,omitempty" xml:"RuleNameEn,omitempty"`
	// The internal code of the rule name.
	//
	// example:
	//
	// ${siem_rule_name_siem_cfw-attack-count-level-up_cfw-attack}
	RuleNameMds *string `json:"RuleNameMds,omitempty" xml:"RuleNameMds,omitempty"`
	// The log source of the rule.
	//
	// example:
	//
	// cloud_siem_aegis_proc
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The status of the predefined rule. Valid values:
	//
	// 	- 0: The rule is in the initial state.
	//
	// 	- 100: The rule takes effect.
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The risk level. Valid values:
	//
	// 	- serious: high.
	//
	// 	- suspicious: medium.
	//
	// 	- remind: low.
	//
	// example:
	//
	// remind
	ThreatLevel *string `json:"ThreatLevel,omitempty" xml:"ThreatLevel,omitempty"`
}

func (s ListCloudSiemPredefinedRulesResponseBodyDataResponseData) String() string {
	return dara.Prettify(s)
}

func (s ListCloudSiemPredefinedRulesResponseBodyDataResponseData) GoString() string {
	return s.String()
}

func (s *ListCloudSiemPredefinedRulesResponseBodyDataResponseData) GetAlertType() *string {
	return s.AlertType
}

func (s *ListCloudSiemPredefinedRulesResponseBodyDataResponseData) GetAttCk() *string {
	return s.AttCk
}

func (s *ListCloudSiemPredefinedRulesResponseBodyDataResponseData) GetEventTransferType() *string {
	return s.EventTransferType
}

func (s *ListCloudSiemPredefinedRulesResponseBodyDataResponseData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListCloudSiemPredefinedRulesResponseBodyDataResponseData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListCloudSiemPredefinedRulesResponseBodyDataResponseData) GetId() *int64 {
	return s.Id
}

func (s *ListCloudSiemPredefinedRulesResponseBodyDataResponseData) GetRuleDescMds() *string {
	return s.RuleDescMds
}

func (s *ListCloudSiemPredefinedRulesResponseBodyDataResponseData) GetRuleName() *string {
	return s.RuleName
}

func (s *ListCloudSiemPredefinedRulesResponseBodyDataResponseData) GetRuleNameCn() *string {
	return s.RuleNameCn
}

func (s *ListCloudSiemPredefinedRulesResponseBodyDataResponseData) GetRuleNameEn() *string {
	return s.RuleNameEn
}

func (s *ListCloudSiemPredefinedRulesResponseBodyDataResponseData) GetRuleNameMds() *string {
	return s.RuleNameMds
}

func (s *ListCloudSiemPredefinedRulesResponseBodyDataResponseData) GetSource() *string {
	return s.Source
}

func (s *ListCloudSiemPredefinedRulesResponseBodyDataResponseData) GetStatus() *int32 {
	return s.Status
}

func (s *ListCloudSiemPredefinedRulesResponseBodyDataResponseData) GetThreatLevel() *string {
	return s.ThreatLevel
}

func (s *ListCloudSiemPredefinedRulesResponseBodyDataResponseData) SetAlertType(v string) *ListCloudSiemPredefinedRulesResponseBodyDataResponseData {
	s.AlertType = &v
	return s
}

func (s *ListCloudSiemPredefinedRulesResponseBodyDataResponseData) SetAttCk(v string) *ListCloudSiemPredefinedRulesResponseBodyDataResponseData {
	s.AttCk = &v
	return s
}

func (s *ListCloudSiemPredefinedRulesResponseBodyDataResponseData) SetEventTransferType(v string) *ListCloudSiemPredefinedRulesResponseBodyDataResponseData {
	s.EventTransferType = &v
	return s
}

func (s *ListCloudSiemPredefinedRulesResponseBodyDataResponseData) SetGmtCreate(v string) *ListCloudSiemPredefinedRulesResponseBodyDataResponseData {
	s.GmtCreate = &v
	return s
}

func (s *ListCloudSiemPredefinedRulesResponseBodyDataResponseData) SetGmtModified(v string) *ListCloudSiemPredefinedRulesResponseBodyDataResponseData {
	s.GmtModified = &v
	return s
}

func (s *ListCloudSiemPredefinedRulesResponseBodyDataResponseData) SetId(v int64) *ListCloudSiemPredefinedRulesResponseBodyDataResponseData {
	s.Id = &v
	return s
}

func (s *ListCloudSiemPredefinedRulesResponseBodyDataResponseData) SetRuleDescMds(v string) *ListCloudSiemPredefinedRulesResponseBodyDataResponseData {
	s.RuleDescMds = &v
	return s
}

func (s *ListCloudSiemPredefinedRulesResponseBodyDataResponseData) SetRuleName(v string) *ListCloudSiemPredefinedRulesResponseBodyDataResponseData {
	s.RuleName = &v
	return s
}

func (s *ListCloudSiemPredefinedRulesResponseBodyDataResponseData) SetRuleNameCn(v string) *ListCloudSiemPredefinedRulesResponseBodyDataResponseData {
	s.RuleNameCn = &v
	return s
}

func (s *ListCloudSiemPredefinedRulesResponseBodyDataResponseData) SetRuleNameEn(v string) *ListCloudSiemPredefinedRulesResponseBodyDataResponseData {
	s.RuleNameEn = &v
	return s
}

func (s *ListCloudSiemPredefinedRulesResponseBodyDataResponseData) SetRuleNameMds(v string) *ListCloudSiemPredefinedRulesResponseBodyDataResponseData {
	s.RuleNameMds = &v
	return s
}

func (s *ListCloudSiemPredefinedRulesResponseBodyDataResponseData) SetSource(v string) *ListCloudSiemPredefinedRulesResponseBodyDataResponseData {
	s.Source = &v
	return s
}

func (s *ListCloudSiemPredefinedRulesResponseBodyDataResponseData) SetStatus(v int32) *ListCloudSiemPredefinedRulesResponseBodyDataResponseData {
	s.Status = &v
	return s
}

func (s *ListCloudSiemPredefinedRulesResponseBodyDataResponseData) SetThreatLevel(v string) *ListCloudSiemPredefinedRulesResponseBodyDataResponseData {
	s.ThreatLevel = &v
	return s
}

func (s *ListCloudSiemPredefinedRulesResponseBodyDataResponseData) Validate() error {
	return dara.Validate(s)
}
