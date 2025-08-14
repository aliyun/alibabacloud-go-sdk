// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSceneRulePageListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeSceneRulePageListResponseBody
	GetCurrentPage() *int32
	SetPageSize(v int32) *DescribeSceneRulePageListResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeSceneRulePageListResponseBody
	GetRequestId() *string
	SetResultObject(v []*DescribeSceneRulePageListResponseBodyResultObject) *DescribeSceneRulePageListResponseBody
	GetResultObject() []*DescribeSceneRulePageListResponseBodyResultObject
	SetTotalItem(v int32) *DescribeSceneRulePageListResponseBody
	GetTotalItem() *int32
	SetTotalPage(v int32) *DescribeSceneRulePageListResponseBody
	GetTotalPage() *int32
}

type DescribeSceneRulePageListResponseBody struct {
	// Current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// Number of items per page in the returned results. Default value: 20, minimum value: 1, maximum value: 50.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Request ID, which is unique for each request, facilitating subsequent troubleshooting
	//
	// example:
	//
	// AE7E6105-7DEB-5125-9B24-DCBC139F6CD2
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Returned object
	ResultObject []*DescribeSceneRulePageListResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Repeated"`
	// Total number of items
	//
	// example:
	//
	// 3
	TotalItem *int32 `json:"totalItem,omitempty" xml:"totalItem,omitempty"`
	// Total number of pages
	//
	// example:
	//
	// 9
	TotalPage *int32 `json:"totalPage,omitempty" xml:"totalPage,omitempty"`
}

func (s DescribeSceneRulePageListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSceneRulePageListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSceneRulePageListResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeSceneRulePageListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSceneRulePageListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSceneRulePageListResponseBody) GetResultObject() []*DescribeSceneRulePageListResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeSceneRulePageListResponseBody) GetTotalItem() *int32 {
	return s.TotalItem
}

func (s *DescribeSceneRulePageListResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *DescribeSceneRulePageListResponseBody) SetCurrentPage(v int32) *DescribeSceneRulePageListResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeSceneRulePageListResponseBody) SetPageSize(v int32) *DescribeSceneRulePageListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeSceneRulePageListResponseBody) SetRequestId(v string) *DescribeSceneRulePageListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSceneRulePageListResponseBody) SetResultObject(v []*DescribeSceneRulePageListResponseBodyResultObject) *DescribeSceneRulePageListResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeSceneRulePageListResponseBody) SetTotalItem(v int32) *DescribeSceneRulePageListResponseBody {
	s.TotalItem = &v
	return s
}

func (s *DescribeSceneRulePageListResponseBody) SetTotalPage(v int32) *DescribeSceneRulePageListResponseBody {
	s.TotalPage = &v
	return s
}

func (s *DescribeSceneRulePageListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeSceneRulePageListResponseBodyResultObject struct {
	// Service authorization type
	//
	// example:
	//
	// admin
	AuthType *string `json:"authType,omitempty" xml:"authType,omitempty"`
	// Audit object
	ConsoleAudit *DescribeSceneRulePageListResponseBodyResultObjectConsoleAudit `json:"consoleAudit,omitempty" xml:"consoleAudit,omitempty" type:"Struct"`
	// Event code
	//
	// example:
	//
	// de_aszbjb7236
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
	// External rule name
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
	// Modification time
	//
	// example:
	//
	// 1565701886000
	GmtModified *int64 `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// Primary key ID of the rule
	//
	// example:
	//
	// 497
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// Main rule ID
	//
	// example:
	//
	// 4399
	MainRuleId *string `json:"mainRuleId,omitempty" xml:"mainRuleId,omitempty"`
	// Strategy priority, where a higher number indicates a higher priority.
	//
	// example:
	//
	// 10
	Priority *int64 `json:"priority,omitempty" xml:"priority,omitempty"`
	// Rule Auth type
	//
	// example:
	//
	// CUSTMER
	RuleAuthType *string `json:"ruleAuthType,omitempty" xml:"ruleAuthType,omitempty"`
	// Strategy ID
	//
	// example:
	//
	// 4730
	RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty"`
	// Strategy description
	//
	// example:
	//
	// 描述信息
	RuleMemo *string `json:"ruleMemo,omitempty" xml:"ruleMemo,omitempty"`
	// Strategy name
	//
	// example:
	//
	// 营销风险识别
	RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty"`
	// Strategy status
	//
	// example:
	//
	// RUNNING
	RuleStatus *string `json:"ruleStatus,omitempty" xml:"ruleStatus,omitempty"`
	// Rule type
	//
	// example:
	//
	// DEFAULT
	RuleType *string `json:"ruleType,omitempty" xml:"ruleType,omitempty"`
	// Primary key ID of the rule version
	//
	// example:
	//
	// 3823
	RuleVersionId *int64 `json:"ruleVersionId,omitempty" xml:"ruleVersionId,omitempty"`
	// Template ID
	//
	// example:
	//
	// 6
	TemplateId *int64 `json:"templateId,omitempty" xml:"templateId,omitempty"`
	// Version number
	//
	// example:
	//
	// 1.0
	Version *int32 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s DescribeSceneRulePageListResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeSceneRulePageListResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeSceneRulePageListResponseBodyResultObject) GetAuthType() *string {
	return s.AuthType
}

func (s *DescribeSceneRulePageListResponseBodyResultObject) GetConsoleAudit() *DescribeSceneRulePageListResponseBodyResultObjectConsoleAudit {
	return s.ConsoleAudit
}

func (s *DescribeSceneRulePageListResponseBodyResultObject) GetEventCode() *string {
	return s.EventCode
}

func (s *DescribeSceneRulePageListResponseBodyResultObject) GetEventName() *string {
	return s.EventName
}

func (s *DescribeSceneRulePageListResponseBodyResultObject) GetEventType() *string {
	return s.EventType
}

func (s *DescribeSceneRulePageListResponseBodyResultObject) GetExternalRuleName() *string {
	return s.ExternalRuleName
}

func (s *DescribeSceneRulePageListResponseBodyResultObject) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *DescribeSceneRulePageListResponseBodyResultObject) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *DescribeSceneRulePageListResponseBodyResultObject) GetId() *int64 {
	return s.Id
}

func (s *DescribeSceneRulePageListResponseBodyResultObject) GetMainRuleId() *string {
	return s.MainRuleId
}

func (s *DescribeSceneRulePageListResponseBodyResultObject) GetPriority() *int64 {
	return s.Priority
}

func (s *DescribeSceneRulePageListResponseBodyResultObject) GetRuleAuthType() *string {
	return s.RuleAuthType
}

func (s *DescribeSceneRulePageListResponseBodyResultObject) GetRuleId() *string {
	return s.RuleId
}

func (s *DescribeSceneRulePageListResponseBodyResultObject) GetRuleMemo() *string {
	return s.RuleMemo
}

func (s *DescribeSceneRulePageListResponseBodyResultObject) GetRuleName() *string {
	return s.RuleName
}

func (s *DescribeSceneRulePageListResponseBodyResultObject) GetRuleStatus() *string {
	return s.RuleStatus
}

func (s *DescribeSceneRulePageListResponseBodyResultObject) GetRuleType() *string {
	return s.RuleType
}

func (s *DescribeSceneRulePageListResponseBodyResultObject) GetRuleVersionId() *int64 {
	return s.RuleVersionId
}

func (s *DescribeSceneRulePageListResponseBodyResultObject) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *DescribeSceneRulePageListResponseBodyResultObject) GetVersion() *int32 {
	return s.Version
}

func (s *DescribeSceneRulePageListResponseBodyResultObject) SetAuthType(v string) *DescribeSceneRulePageListResponseBodyResultObject {
	s.AuthType = &v
	return s
}

func (s *DescribeSceneRulePageListResponseBodyResultObject) SetConsoleAudit(v *DescribeSceneRulePageListResponseBodyResultObjectConsoleAudit) *DescribeSceneRulePageListResponseBodyResultObject {
	s.ConsoleAudit = v
	return s
}

func (s *DescribeSceneRulePageListResponseBodyResultObject) SetEventCode(v string) *DescribeSceneRulePageListResponseBodyResultObject {
	s.EventCode = &v
	return s
}

func (s *DescribeSceneRulePageListResponseBodyResultObject) SetEventName(v string) *DescribeSceneRulePageListResponseBodyResultObject {
	s.EventName = &v
	return s
}

func (s *DescribeSceneRulePageListResponseBodyResultObject) SetEventType(v string) *DescribeSceneRulePageListResponseBodyResultObject {
	s.EventType = &v
	return s
}

func (s *DescribeSceneRulePageListResponseBodyResultObject) SetExternalRuleName(v string) *DescribeSceneRulePageListResponseBodyResultObject {
	s.ExternalRuleName = &v
	return s
}

func (s *DescribeSceneRulePageListResponseBodyResultObject) SetGmtCreate(v int64) *DescribeSceneRulePageListResponseBodyResultObject {
	s.GmtCreate = &v
	return s
}

func (s *DescribeSceneRulePageListResponseBodyResultObject) SetGmtModified(v int64) *DescribeSceneRulePageListResponseBodyResultObject {
	s.GmtModified = &v
	return s
}

func (s *DescribeSceneRulePageListResponseBodyResultObject) SetId(v int64) *DescribeSceneRulePageListResponseBodyResultObject {
	s.Id = &v
	return s
}

func (s *DescribeSceneRulePageListResponseBodyResultObject) SetMainRuleId(v string) *DescribeSceneRulePageListResponseBodyResultObject {
	s.MainRuleId = &v
	return s
}

func (s *DescribeSceneRulePageListResponseBodyResultObject) SetPriority(v int64) *DescribeSceneRulePageListResponseBodyResultObject {
	s.Priority = &v
	return s
}

func (s *DescribeSceneRulePageListResponseBodyResultObject) SetRuleAuthType(v string) *DescribeSceneRulePageListResponseBodyResultObject {
	s.RuleAuthType = &v
	return s
}

func (s *DescribeSceneRulePageListResponseBodyResultObject) SetRuleId(v string) *DescribeSceneRulePageListResponseBodyResultObject {
	s.RuleId = &v
	return s
}

func (s *DescribeSceneRulePageListResponseBodyResultObject) SetRuleMemo(v string) *DescribeSceneRulePageListResponseBodyResultObject {
	s.RuleMemo = &v
	return s
}

func (s *DescribeSceneRulePageListResponseBodyResultObject) SetRuleName(v string) *DescribeSceneRulePageListResponseBodyResultObject {
	s.RuleName = &v
	return s
}

func (s *DescribeSceneRulePageListResponseBodyResultObject) SetRuleStatus(v string) *DescribeSceneRulePageListResponseBodyResultObject {
	s.RuleStatus = &v
	return s
}

func (s *DescribeSceneRulePageListResponseBodyResultObject) SetRuleType(v string) *DescribeSceneRulePageListResponseBodyResultObject {
	s.RuleType = &v
	return s
}

func (s *DescribeSceneRulePageListResponseBodyResultObject) SetRuleVersionId(v int64) *DescribeSceneRulePageListResponseBodyResultObject {
	s.RuleVersionId = &v
	return s
}

func (s *DescribeSceneRulePageListResponseBodyResultObject) SetTemplateId(v int64) *DescribeSceneRulePageListResponseBodyResultObject {
	s.TemplateId = &v
	return s
}

func (s *DescribeSceneRulePageListResponseBodyResultObject) SetVersion(v int32) *DescribeSceneRulePageListResponseBodyResultObject {
	s.Version = &v
	return s
}

func (s *DescribeSceneRulePageListResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}

type DescribeSceneRulePageListResponseBodyResultObjectConsoleAudit struct {
	// Initiator account ID
	//
	// example:
	//
	// 1234567890999
	ApplyUserId *string `json:"applyUserId,omitempty" xml:"applyUserId,omitempty"`
	// Initiator account name
	//
	// example:
	//
	// 张三
	ApplyUserName *string `json:"applyUserName,omitempty" xml:"applyUserName,omitempty"`
	// Approval comments
	//
	// example:
	//
	// 同意
	AuditMsg *string `json:"auditMsg,omitempty" xml:"auditMsg,omitempty"`
	// Final approver ID
	//
	// example:
	//
	// 1234567890
	AuditRealUserId *string `json:"auditRealUserId,omitempty" xml:"auditRealUserId,omitempty"`
	// Approver account name
	//
	// example:
	//
	// 王五
	AuditRealUserName *string `json:"auditRealUserName,omitempty" xml:"auditRealUserName,omitempty"`
	// Approval application remarks
	//
	// example:
	//
	// 备注
	AuditRemark *string `json:"auditRemark,omitempty" xml:"auditRemark,omitempty"`
	// Status
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
	// Designated auditor account IDs (comma-separated for multiple)
	//
	// example:
	//
	// 123
	AuditUserId *string `json:"auditUserId,omitempty" xml:"auditUserId,omitempty"`
	// Designated auditor account names (comma-separated for multiple)
	//
	// example:
	//
	// 李四
	AuditUserName *string `json:"auditUserName,omitempty" xml:"auditUserName,omitempty"`
	// Creation time in UTC
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
	// Information about other related parties (in JSON format)
	//
	// example:
	//
	// {}
	RelationExt *string `json:"relationExt,omitempty" xml:"relationExt,omitempty"`
	// 审批关联的事务ID
	//
	// example:
	//
	// 123
	RelationId *int64 `json:"relationId,omitempty" xml:"relationId,omitempty"`
	// The name of the associated transaction (can be null)
	//
	// example:
	//
	// t
	RelationName *string `json:"relationName,omitempty" xml:"relationName,omitempty"`
	// 审批的类型（如rule代表策略的审批）
	//
	// example:
	//
	// RULE
	RelationType *string `json:"relationType,omitempty" xml:"relationType,omitempty"`
}

func (s DescribeSceneRulePageListResponseBodyResultObjectConsoleAudit) String() string {
	return dara.Prettify(s)
}

func (s DescribeSceneRulePageListResponseBodyResultObjectConsoleAudit) GoString() string {
	return s.String()
}

func (s *DescribeSceneRulePageListResponseBodyResultObjectConsoleAudit) GetApplyUserId() *string {
	return s.ApplyUserId
}

func (s *DescribeSceneRulePageListResponseBodyResultObjectConsoleAudit) GetApplyUserName() *string {
	return s.ApplyUserName
}

func (s *DescribeSceneRulePageListResponseBodyResultObjectConsoleAudit) GetAuditMsg() *string {
	return s.AuditMsg
}

func (s *DescribeSceneRulePageListResponseBodyResultObjectConsoleAudit) GetAuditRealUserId() *string {
	return s.AuditRealUserId
}

func (s *DescribeSceneRulePageListResponseBodyResultObjectConsoleAudit) GetAuditRealUserName() *string {
	return s.AuditRealUserName
}

func (s *DescribeSceneRulePageListResponseBodyResultObjectConsoleAudit) GetAuditRemark() *string {
	return s.AuditRemark
}

func (s *DescribeSceneRulePageListResponseBodyResultObjectConsoleAudit) GetAuditStatus() *string {
	return s.AuditStatus
}

func (s *DescribeSceneRulePageListResponseBodyResultObjectConsoleAudit) GetAuditTime() *int64 {
	return s.AuditTime
}

func (s *DescribeSceneRulePageListResponseBodyResultObjectConsoleAudit) GetAuditUserId() *string {
	return s.AuditUserId
}

func (s *DescribeSceneRulePageListResponseBodyResultObjectConsoleAudit) GetAuditUserName() *string {
	return s.AuditUserName
}

func (s *DescribeSceneRulePageListResponseBodyResultObjectConsoleAudit) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *DescribeSceneRulePageListResponseBodyResultObjectConsoleAudit) GetId() *int64 {
	return s.Id
}

func (s *DescribeSceneRulePageListResponseBodyResultObjectConsoleAudit) GetRelationExt() *string {
	return s.RelationExt
}

func (s *DescribeSceneRulePageListResponseBodyResultObjectConsoleAudit) GetRelationId() *int64 {
	return s.RelationId
}

func (s *DescribeSceneRulePageListResponseBodyResultObjectConsoleAudit) GetRelationName() *string {
	return s.RelationName
}

func (s *DescribeSceneRulePageListResponseBodyResultObjectConsoleAudit) GetRelationType() *string {
	return s.RelationType
}

func (s *DescribeSceneRulePageListResponseBodyResultObjectConsoleAudit) SetApplyUserId(v string) *DescribeSceneRulePageListResponseBodyResultObjectConsoleAudit {
	s.ApplyUserId = &v
	return s
}

func (s *DescribeSceneRulePageListResponseBodyResultObjectConsoleAudit) SetApplyUserName(v string) *DescribeSceneRulePageListResponseBodyResultObjectConsoleAudit {
	s.ApplyUserName = &v
	return s
}

func (s *DescribeSceneRulePageListResponseBodyResultObjectConsoleAudit) SetAuditMsg(v string) *DescribeSceneRulePageListResponseBodyResultObjectConsoleAudit {
	s.AuditMsg = &v
	return s
}

func (s *DescribeSceneRulePageListResponseBodyResultObjectConsoleAudit) SetAuditRealUserId(v string) *DescribeSceneRulePageListResponseBodyResultObjectConsoleAudit {
	s.AuditRealUserId = &v
	return s
}

func (s *DescribeSceneRulePageListResponseBodyResultObjectConsoleAudit) SetAuditRealUserName(v string) *DescribeSceneRulePageListResponseBodyResultObjectConsoleAudit {
	s.AuditRealUserName = &v
	return s
}

func (s *DescribeSceneRulePageListResponseBodyResultObjectConsoleAudit) SetAuditRemark(v string) *DescribeSceneRulePageListResponseBodyResultObjectConsoleAudit {
	s.AuditRemark = &v
	return s
}

func (s *DescribeSceneRulePageListResponseBodyResultObjectConsoleAudit) SetAuditStatus(v string) *DescribeSceneRulePageListResponseBodyResultObjectConsoleAudit {
	s.AuditStatus = &v
	return s
}

func (s *DescribeSceneRulePageListResponseBodyResultObjectConsoleAudit) SetAuditTime(v int64) *DescribeSceneRulePageListResponseBodyResultObjectConsoleAudit {
	s.AuditTime = &v
	return s
}

func (s *DescribeSceneRulePageListResponseBodyResultObjectConsoleAudit) SetAuditUserId(v string) *DescribeSceneRulePageListResponseBodyResultObjectConsoleAudit {
	s.AuditUserId = &v
	return s
}

func (s *DescribeSceneRulePageListResponseBodyResultObjectConsoleAudit) SetAuditUserName(v string) *DescribeSceneRulePageListResponseBodyResultObjectConsoleAudit {
	s.AuditUserName = &v
	return s
}

func (s *DescribeSceneRulePageListResponseBodyResultObjectConsoleAudit) SetGmtCreate(v int64) *DescribeSceneRulePageListResponseBodyResultObjectConsoleAudit {
	s.GmtCreate = &v
	return s
}

func (s *DescribeSceneRulePageListResponseBodyResultObjectConsoleAudit) SetId(v int64) *DescribeSceneRulePageListResponseBodyResultObjectConsoleAudit {
	s.Id = &v
	return s
}

func (s *DescribeSceneRulePageListResponseBodyResultObjectConsoleAudit) SetRelationExt(v string) *DescribeSceneRulePageListResponseBodyResultObjectConsoleAudit {
	s.RelationExt = &v
	return s
}

func (s *DescribeSceneRulePageListResponseBodyResultObjectConsoleAudit) SetRelationId(v int64) *DescribeSceneRulePageListResponseBodyResultObjectConsoleAudit {
	s.RelationId = &v
	return s
}

func (s *DescribeSceneRulePageListResponseBodyResultObjectConsoleAudit) SetRelationName(v string) *DescribeSceneRulePageListResponseBodyResultObjectConsoleAudit {
	s.RelationName = &v
	return s
}

func (s *DescribeSceneRulePageListResponseBodyResultObjectConsoleAudit) SetRelationType(v string) *DescribeSceneRulePageListResponseBodyResultObjectConsoleAudit {
	s.RelationType = &v
	return s
}

func (s *DescribeSceneRulePageListResponseBodyResultObjectConsoleAudit) Validate() error {
	return dara.Validate(s)
}
