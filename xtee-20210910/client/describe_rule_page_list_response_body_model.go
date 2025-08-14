// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRulePageListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeRulePageListResponseBody
	GetCurrentPage() *int32
	SetPageSize(v int32) *DescribeRulePageListResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeRulePageListResponseBody
	GetRequestId() *string
	SetResultObject(v []*DescribeRulePageListResponseBodyResultObject) *DescribeRulePageListResponseBody
	GetResultObject() []*DescribeRulePageListResponseBodyResultObject
	SetTotalItem(v int32) *DescribeRulePageListResponseBody
	GetTotalItem() *int32
	SetTotalPage(v int32) *DescribeRulePageListResponseBody
	GetTotalPage() *int32
}

type DescribeRulePageListResponseBody struct {
	// Current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// Page size, default value is 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Request ID.
	//
	// example:
	//
	// AE7E6105-7DEB-5125-9B24-DCBC139F6CD2
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Returned object.
	ResultObject []*DescribeRulePageListResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Repeated"`
	// Total number of items.
	//
	// example:
	//
	// 28
	TotalItem *int32 `json:"totalItem,omitempty" xml:"totalItem,omitempty"`
	// Total number of pages.
	//
	// example:
	//
	// 3
	TotalPage *int32 `json:"totalPage,omitempty" xml:"totalPage,omitempty"`
}

func (s DescribeRulePageListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRulePageListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRulePageListResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeRulePageListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeRulePageListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRulePageListResponseBody) GetResultObject() []*DescribeRulePageListResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeRulePageListResponseBody) GetTotalItem() *int32 {
	return s.TotalItem
}

func (s *DescribeRulePageListResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *DescribeRulePageListResponseBody) SetCurrentPage(v int32) *DescribeRulePageListResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeRulePageListResponseBody) SetPageSize(v int32) *DescribeRulePageListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeRulePageListResponseBody) SetRequestId(v string) *DescribeRulePageListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRulePageListResponseBody) SetResultObject(v []*DescribeRulePageListResponseBodyResultObject) *DescribeRulePageListResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeRulePageListResponseBody) SetTotalItem(v int32) *DescribeRulePageListResponseBody {
	s.TotalItem = &v
	return s
}

func (s *DescribeRulePageListResponseBody) SetTotalPage(v int32) *DescribeRulePageListResponseBody {
	s.TotalPage = &v
	return s
}

func (s *DescribeRulePageListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeRulePageListResponseBodyResultObject struct {
	// Service authorization type
	//
	// example:
	//
	// all
	AuthType *string `json:"authType,omitempty" xml:"authType,omitempty"`
	// Audit object
	ConsoleAudit *DescribeRulePageListResponseBodyResultObjectConsoleAudit `json:"consoleAudit,omitempty" xml:"consoleAudit,omitempty" type:"Struct"`
	// Event code.
	//
	// example:
	//
	// de_arcehq4370
	EventCode *string `json:"eventCode,omitempty" xml:"eventCode,omitempty"`
	// Event name.
	//
	// example:
	//
	// 注册风险
	EventName *string `json:"eventName,omitempty" xml:"eventName,omitempty"`
	// Event type
	//
	// example:
	//
	// MAIN
	EventType *string `json:"eventType,omitempty" xml:"eventType,omitempty"`
	// External name of the rule
	//
	// example:
	//
	// 策略1
	ExternalRuleName *string `json:"externalRuleName,omitempty" xml:"externalRuleName,omitempty"`
	// Creation time.
	//
	// example:
	//
	// 1621578648000
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// Modification time.
	//
	// example:
	//
	// 1565701886000
	GmtModified *int64 `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// Whether there is a new version
	//
	// example:
	//
	// true
	HasNewVersion *bool `json:"hasNewVersion,omitempty" xml:"hasNewVersion,omitempty"`
	// Primary key ID of the policy.
	//
	// example:
	//
	// 2793
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// Main rule ID
	//
	// example:
	//
	// 4399
	MainRuleId *string `json:"mainRuleId,omitempty" xml:"mainRuleId,omitempty"`
	// Policy priority, the higher the number, the higher the priority.
	//
	// example:
	//
	// 10
	Priority *int64 `json:"priority,omitempty" xml:"priority,omitempty"`
	// Rule type
	//
	// example:
	//
	// NORMAL
	RuleAuthType *string `json:"ruleAuthType,omitempty" xml:"ruleAuthType,omitempty"`
	// Policy ID.
	//
	// example:
	//
	// 102059
	RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty"`
	// Policy description.
	//
	// example:
	//
	// 描述信息
	RuleMemo *string `json:"ruleMemo,omitempty" xml:"ruleMemo,omitempty"`
	// Policy name.
	//
	// example:
	//
	// 营销风险识别
	RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty"`
	// Policy status.
	//
	// example:
	//
	// RUNNING
	RuleStatus *string `json:"ruleStatus,omitempty" xml:"ruleStatus,omitempty"`
	// Rule type
	//
	// example:
	//
	// NORMAL
	RuleType *string `json:"ruleType,omitempty" xml:"ruleType,omitempty"`
	// Primary key ID of the rule version.
	//
	// example:
	//
	// 11300
	RuleVersionId *int64 `json:"ruleVersionId,omitempty" xml:"ruleVersionId,omitempty"`
	// Template ID.
	//
	// example:
	//
	// register
	TemplateId *int64 `json:"templateId,omitempty" xml:"templateId,omitempty"`
	// Version number.
	//
	// example:
	//
	// 1
	Version *int32 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s DescribeRulePageListResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeRulePageListResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeRulePageListResponseBodyResultObject) GetAuthType() *string {
	return s.AuthType
}

func (s *DescribeRulePageListResponseBodyResultObject) GetConsoleAudit() *DescribeRulePageListResponseBodyResultObjectConsoleAudit {
	return s.ConsoleAudit
}

func (s *DescribeRulePageListResponseBodyResultObject) GetEventCode() *string {
	return s.EventCode
}

func (s *DescribeRulePageListResponseBodyResultObject) GetEventName() *string {
	return s.EventName
}

func (s *DescribeRulePageListResponseBodyResultObject) GetEventType() *string {
	return s.EventType
}

func (s *DescribeRulePageListResponseBodyResultObject) GetExternalRuleName() *string {
	return s.ExternalRuleName
}

func (s *DescribeRulePageListResponseBodyResultObject) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *DescribeRulePageListResponseBodyResultObject) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *DescribeRulePageListResponseBodyResultObject) GetHasNewVersion() *bool {
	return s.HasNewVersion
}

func (s *DescribeRulePageListResponseBodyResultObject) GetId() *int64 {
	return s.Id
}

func (s *DescribeRulePageListResponseBodyResultObject) GetMainRuleId() *string {
	return s.MainRuleId
}

func (s *DescribeRulePageListResponseBodyResultObject) GetPriority() *int64 {
	return s.Priority
}

func (s *DescribeRulePageListResponseBodyResultObject) GetRuleAuthType() *string {
	return s.RuleAuthType
}

func (s *DescribeRulePageListResponseBodyResultObject) GetRuleId() *string {
	return s.RuleId
}

func (s *DescribeRulePageListResponseBodyResultObject) GetRuleMemo() *string {
	return s.RuleMemo
}

func (s *DescribeRulePageListResponseBodyResultObject) GetRuleName() *string {
	return s.RuleName
}

func (s *DescribeRulePageListResponseBodyResultObject) GetRuleStatus() *string {
	return s.RuleStatus
}

func (s *DescribeRulePageListResponseBodyResultObject) GetRuleType() *string {
	return s.RuleType
}

func (s *DescribeRulePageListResponseBodyResultObject) GetRuleVersionId() *int64 {
	return s.RuleVersionId
}

func (s *DescribeRulePageListResponseBodyResultObject) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *DescribeRulePageListResponseBodyResultObject) GetVersion() *int32 {
	return s.Version
}

func (s *DescribeRulePageListResponseBodyResultObject) SetAuthType(v string) *DescribeRulePageListResponseBodyResultObject {
	s.AuthType = &v
	return s
}

func (s *DescribeRulePageListResponseBodyResultObject) SetConsoleAudit(v *DescribeRulePageListResponseBodyResultObjectConsoleAudit) *DescribeRulePageListResponseBodyResultObject {
	s.ConsoleAudit = v
	return s
}

func (s *DescribeRulePageListResponseBodyResultObject) SetEventCode(v string) *DescribeRulePageListResponseBodyResultObject {
	s.EventCode = &v
	return s
}

func (s *DescribeRulePageListResponseBodyResultObject) SetEventName(v string) *DescribeRulePageListResponseBodyResultObject {
	s.EventName = &v
	return s
}

func (s *DescribeRulePageListResponseBodyResultObject) SetEventType(v string) *DescribeRulePageListResponseBodyResultObject {
	s.EventType = &v
	return s
}

func (s *DescribeRulePageListResponseBodyResultObject) SetExternalRuleName(v string) *DescribeRulePageListResponseBodyResultObject {
	s.ExternalRuleName = &v
	return s
}

func (s *DescribeRulePageListResponseBodyResultObject) SetGmtCreate(v int64) *DescribeRulePageListResponseBodyResultObject {
	s.GmtCreate = &v
	return s
}

func (s *DescribeRulePageListResponseBodyResultObject) SetGmtModified(v int64) *DescribeRulePageListResponseBodyResultObject {
	s.GmtModified = &v
	return s
}

func (s *DescribeRulePageListResponseBodyResultObject) SetHasNewVersion(v bool) *DescribeRulePageListResponseBodyResultObject {
	s.HasNewVersion = &v
	return s
}

func (s *DescribeRulePageListResponseBodyResultObject) SetId(v int64) *DescribeRulePageListResponseBodyResultObject {
	s.Id = &v
	return s
}

func (s *DescribeRulePageListResponseBodyResultObject) SetMainRuleId(v string) *DescribeRulePageListResponseBodyResultObject {
	s.MainRuleId = &v
	return s
}

func (s *DescribeRulePageListResponseBodyResultObject) SetPriority(v int64) *DescribeRulePageListResponseBodyResultObject {
	s.Priority = &v
	return s
}

func (s *DescribeRulePageListResponseBodyResultObject) SetRuleAuthType(v string) *DescribeRulePageListResponseBodyResultObject {
	s.RuleAuthType = &v
	return s
}

func (s *DescribeRulePageListResponseBodyResultObject) SetRuleId(v string) *DescribeRulePageListResponseBodyResultObject {
	s.RuleId = &v
	return s
}

func (s *DescribeRulePageListResponseBodyResultObject) SetRuleMemo(v string) *DescribeRulePageListResponseBodyResultObject {
	s.RuleMemo = &v
	return s
}

func (s *DescribeRulePageListResponseBodyResultObject) SetRuleName(v string) *DescribeRulePageListResponseBodyResultObject {
	s.RuleName = &v
	return s
}

func (s *DescribeRulePageListResponseBodyResultObject) SetRuleStatus(v string) *DescribeRulePageListResponseBodyResultObject {
	s.RuleStatus = &v
	return s
}

func (s *DescribeRulePageListResponseBodyResultObject) SetRuleType(v string) *DescribeRulePageListResponseBodyResultObject {
	s.RuleType = &v
	return s
}

func (s *DescribeRulePageListResponseBodyResultObject) SetRuleVersionId(v int64) *DescribeRulePageListResponseBodyResultObject {
	s.RuleVersionId = &v
	return s
}

func (s *DescribeRulePageListResponseBodyResultObject) SetTemplateId(v int64) *DescribeRulePageListResponseBodyResultObject {
	s.TemplateId = &v
	return s
}

func (s *DescribeRulePageListResponseBodyResultObject) SetVersion(v int32) *DescribeRulePageListResponseBodyResultObject {
	s.Version = &v
	return s
}

func (s *DescribeRulePageListResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}

type DescribeRulePageListResponseBodyResultObjectConsoleAudit struct {
	// UID of the user who passed the audit
	//
	// example:
	//
	// 用户uid
	ApplyUserId *string `json:"applyUserId,omitempty" xml:"applyUserId,omitempty"`
	// Name of the user who passed the audit
	//
	// example:
	//
	// root
	ApplyUserName *string `json:"applyUserName,omitempty" xml:"applyUserName,omitempty"`
	// Approval comments
	//
	// example:
	//
	// 同意
	AuditMsg *string `json:"auditMsg,omitempty" xml:"auditMsg,omitempty"`
	// UID of the final auditor
	//
	// example:
	//
	// 1728
	AuditRealUserId *string `json:"auditRealUserId,omitempty" xml:"auditRealUserId,omitempty"`
	// Name of the final auditor
	//
	// example:
	//
	// root
	AuditRealUserName *string `json:"auditRealUserName,omitempty" xml:"auditRealUserName,omitempty"`
	// Remarks by the approver.
	//
	// example:
	//
	// 备注
	AuditRemark *string `json:"auditRemark,omitempty" xml:"auditRemark,omitempty"`
	// Application audit status
	//
	// example:
	//
	// AGREE
	AuditStatus *string `json:"auditStatus,omitempty" xml:"auditStatus,omitempty"`
	// Approval time
	//
	// example:
	//
	// 1545726028000
	AuditTime *int64 `json:"auditTime,omitempty" xml:"auditTime,omitempty"`
	// UID of the auditor
	//
	// example:
	//
	// 1234xxxx
	AuditUserId *string `json:"auditUserId,omitempty" xml:"auditUserId,omitempty"`
	// Name of the auditor
	//
	// example:
	//
	// root
	AuditUserName *string `json:"auditUserName,omitempty" xml:"auditUserName,omitempty"`
	// Creation time.
	//
	// example:
	//
	// 1545726028000
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// Primary key ID
	//
	// example:
	//
	// 1728
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// Information of other related parties (in JSON format)
	//
	// example:
	//
	// {}
	RelationExt *string `json:"relationExt,omitempty" xml:"relationExt,omitempty"`
	// ID of the related transaction for the approval
	//
	// example:
	//
	// 123
	RelationId *int64 `json:"relationId,omitempty" xml:"relationId,omitempty"`
	// Name of the related transaction for the approval (can be null)
	//
	// example:
	//
	// 策略1
	RelationName *string `json:"relationName,omitempty" xml:"relationName,omitempty"`
	// Type of the approval (e.g., `rule` represents the approval of a rule)
	//
	// example:
	//
	// RULE
	RelationType *string `json:"relationType,omitempty" xml:"relationType,omitempty"`
}

func (s DescribeRulePageListResponseBodyResultObjectConsoleAudit) String() string {
	return dara.Prettify(s)
}

func (s DescribeRulePageListResponseBodyResultObjectConsoleAudit) GoString() string {
	return s.String()
}

func (s *DescribeRulePageListResponseBodyResultObjectConsoleAudit) GetApplyUserId() *string {
	return s.ApplyUserId
}

func (s *DescribeRulePageListResponseBodyResultObjectConsoleAudit) GetApplyUserName() *string {
	return s.ApplyUserName
}

func (s *DescribeRulePageListResponseBodyResultObjectConsoleAudit) GetAuditMsg() *string {
	return s.AuditMsg
}

func (s *DescribeRulePageListResponseBodyResultObjectConsoleAudit) GetAuditRealUserId() *string {
	return s.AuditRealUserId
}

func (s *DescribeRulePageListResponseBodyResultObjectConsoleAudit) GetAuditRealUserName() *string {
	return s.AuditRealUserName
}

func (s *DescribeRulePageListResponseBodyResultObjectConsoleAudit) GetAuditRemark() *string {
	return s.AuditRemark
}

func (s *DescribeRulePageListResponseBodyResultObjectConsoleAudit) GetAuditStatus() *string {
	return s.AuditStatus
}

func (s *DescribeRulePageListResponseBodyResultObjectConsoleAudit) GetAuditTime() *int64 {
	return s.AuditTime
}

func (s *DescribeRulePageListResponseBodyResultObjectConsoleAudit) GetAuditUserId() *string {
	return s.AuditUserId
}

func (s *DescribeRulePageListResponseBodyResultObjectConsoleAudit) GetAuditUserName() *string {
	return s.AuditUserName
}

func (s *DescribeRulePageListResponseBodyResultObjectConsoleAudit) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *DescribeRulePageListResponseBodyResultObjectConsoleAudit) GetId() *int64 {
	return s.Id
}

func (s *DescribeRulePageListResponseBodyResultObjectConsoleAudit) GetRelationExt() *string {
	return s.RelationExt
}

func (s *DescribeRulePageListResponseBodyResultObjectConsoleAudit) GetRelationId() *int64 {
	return s.RelationId
}

func (s *DescribeRulePageListResponseBodyResultObjectConsoleAudit) GetRelationName() *string {
	return s.RelationName
}

func (s *DescribeRulePageListResponseBodyResultObjectConsoleAudit) GetRelationType() *string {
	return s.RelationType
}

func (s *DescribeRulePageListResponseBodyResultObjectConsoleAudit) SetApplyUserId(v string) *DescribeRulePageListResponseBodyResultObjectConsoleAudit {
	s.ApplyUserId = &v
	return s
}

func (s *DescribeRulePageListResponseBodyResultObjectConsoleAudit) SetApplyUserName(v string) *DescribeRulePageListResponseBodyResultObjectConsoleAudit {
	s.ApplyUserName = &v
	return s
}

func (s *DescribeRulePageListResponseBodyResultObjectConsoleAudit) SetAuditMsg(v string) *DescribeRulePageListResponseBodyResultObjectConsoleAudit {
	s.AuditMsg = &v
	return s
}

func (s *DescribeRulePageListResponseBodyResultObjectConsoleAudit) SetAuditRealUserId(v string) *DescribeRulePageListResponseBodyResultObjectConsoleAudit {
	s.AuditRealUserId = &v
	return s
}

func (s *DescribeRulePageListResponseBodyResultObjectConsoleAudit) SetAuditRealUserName(v string) *DescribeRulePageListResponseBodyResultObjectConsoleAudit {
	s.AuditRealUserName = &v
	return s
}

func (s *DescribeRulePageListResponseBodyResultObjectConsoleAudit) SetAuditRemark(v string) *DescribeRulePageListResponseBodyResultObjectConsoleAudit {
	s.AuditRemark = &v
	return s
}

func (s *DescribeRulePageListResponseBodyResultObjectConsoleAudit) SetAuditStatus(v string) *DescribeRulePageListResponseBodyResultObjectConsoleAudit {
	s.AuditStatus = &v
	return s
}

func (s *DescribeRulePageListResponseBodyResultObjectConsoleAudit) SetAuditTime(v int64) *DescribeRulePageListResponseBodyResultObjectConsoleAudit {
	s.AuditTime = &v
	return s
}

func (s *DescribeRulePageListResponseBodyResultObjectConsoleAudit) SetAuditUserId(v string) *DescribeRulePageListResponseBodyResultObjectConsoleAudit {
	s.AuditUserId = &v
	return s
}

func (s *DescribeRulePageListResponseBodyResultObjectConsoleAudit) SetAuditUserName(v string) *DescribeRulePageListResponseBodyResultObjectConsoleAudit {
	s.AuditUserName = &v
	return s
}

func (s *DescribeRulePageListResponseBodyResultObjectConsoleAudit) SetGmtCreate(v int64) *DescribeRulePageListResponseBodyResultObjectConsoleAudit {
	s.GmtCreate = &v
	return s
}

func (s *DescribeRulePageListResponseBodyResultObjectConsoleAudit) SetId(v int64) *DescribeRulePageListResponseBodyResultObjectConsoleAudit {
	s.Id = &v
	return s
}

func (s *DescribeRulePageListResponseBodyResultObjectConsoleAudit) SetRelationExt(v string) *DescribeRulePageListResponseBodyResultObjectConsoleAudit {
	s.RelationExt = &v
	return s
}

func (s *DescribeRulePageListResponseBodyResultObjectConsoleAudit) SetRelationId(v int64) *DescribeRulePageListResponseBodyResultObjectConsoleAudit {
	s.RelationId = &v
	return s
}

func (s *DescribeRulePageListResponseBodyResultObjectConsoleAudit) SetRelationName(v string) *DescribeRulePageListResponseBodyResultObjectConsoleAudit {
	s.RelationName = &v
	return s
}

func (s *DescribeRulePageListResponseBodyResultObjectConsoleAudit) SetRelationType(v string) *DescribeRulePageListResponseBodyResultObjectConsoleAudit {
	s.RelationType = &v
	return s
}

func (s *DescribeRulePageListResponseBodyResultObjectConsoleAudit) Validate() error {
	return dara.Validate(s)
}
