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
	// The current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// The number of entries per page. Default value: 20. Minimum value: 1. Maximum value: 50.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The request ID. Each request has a unique ID for troubleshooting purposes.
	//
	// example:
	//
	// AE7E6105-7DEB-5125-9B24-DCBC139F6CD2
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The response object.
	ResultObject []*DescribeSceneRulePageListResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Repeated"`
	// The total number of entries.
	//
	// example:
	//
	// 3
	TotalItem *int32 `json:"totalItem,omitempty" xml:"totalItem,omitempty"`
	// The total number of pages.
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
	if s.ResultObject != nil {
		for _, item := range s.ResultObject {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSceneRulePageListResponseBodyResultObject struct {
	// The service authorization type.
	//
	// example:
	//
	// admin
	AuthType *string `json:"authType,omitempty" xml:"authType,omitempty"`
	// The audit object.
	ConsoleAudit *DescribeSceneRulePageListResponseBodyResultObjectConsoleAudit `json:"consoleAudit,omitempty" xml:"consoleAudit,omitempty" type:"Struct"`
	// The event code.
	//
	// example:
	//
	// de_aszbjb7236
	EventCode *string `json:"eventCode,omitempty" xml:"eventCode,omitempty"`
	// The event name.
	//
	// example:
	//
	// 注册风险
	EventName *string `json:"eventName,omitempty" xml:"eventName,omitempty"`
	// The event type. Valid values:
	//
	// - BYPASS: bypass event.
	//
	// - SHUNT: shunt event.
	//
	// - MAIN: main event.
	//
	// example:
	//
	// MAIN
	EventType *string `json:"eventType,omitempty" xml:"eventType,omitempty"`
	// The customer-facing policy name.
	//
	// example:
	//
	// 策略1
	ExternalRuleName *string `json:"externalRuleName,omitempty" xml:"externalRuleName,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 1621578648000
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// The modification time.
	//
	// example:
	//
	// 1565701886000
	GmtModified *int64 `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// The primary key ID of the policy.
	//
	// example:
	//
	// 497
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// The main policy ID.
	//
	// example:
	//
	// 4399
	MainRuleId *string `json:"mainRuleId,omitempty" xml:"mainRuleId,omitempty"`
	// The policy priority. A larger value indicates a higher priority.
	//
	// example:
	//
	// 10
	Priority *int64 `json:"priority,omitempty" xml:"priority,omitempty"`
	// The policy type.
	//
	// example:
	//
	// CUSTMER
	RuleAuthType *string `json:"ruleAuthType,omitempty" xml:"ruleAuthType,omitempty"`
	// The policy ID.
	//
	// example:
	//
	// 4730
	RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty"`
	// The policy description.
	//
	// example:
	//
	// 描述信息
	RuleMemo *string `json:"ruleMemo,omitempty" xml:"ruleMemo,omitempty"`
	// The policy name.
	//
	// example:
	//
	// 营销风险识别
	RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty"`
	// The policy status.
	//
	// example:
	//
	// RUNNING
	RuleStatus *string `json:"ruleStatus,omitempty" xml:"ruleStatus,omitempty"`
	// The rule type.
	//
	// example:
	//
	// DSL
	RuleType *string `json:"ruleType,omitempty" xml:"ruleType,omitempty"`
	// The primary key ID of the policy version.
	//
	// example:
	//
	// 3823
	RuleVersionId *int64 `json:"ruleVersionId,omitempty" xml:"ruleVersionId,omitempty"`
	// The template ID.
	//
	// example:
	//
	// 6
	TemplateId *int64 `json:"templateId,omitempty" xml:"templateId,omitempty"`
	// The version number.
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
	if s.ConsoleAudit != nil {
		if err := s.ConsoleAudit.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSceneRulePageListResponseBodyResultObjectConsoleAudit struct {
	// The account ID of the applicant.
	//
	// example:
	//
	// 1234567890999
	ApplyUserId *string `json:"applyUserId,omitempty" xml:"applyUserId,omitempty"`
	// The account name of the applicant.
	//
	// example:
	//
	// 张三
	ApplyUserName *string `json:"applyUserName,omitempty" xml:"applyUserName,omitempty"`
	// The approval comment.
	//
	// example:
	//
	// 同意
	AuditMsg *string `json:"auditMsg,omitempty" xml:"auditMsg,omitempty"`
	// The ID of the final approver.
	//
	// example:
	//
	// 1234567890
	AuditRealUserId *string `json:"auditRealUserId,omitempty" xml:"auditRealUserId,omitempty"`
	// The account name of the approver.
	//
	// example:
	//
	// 王五
	AuditRealUserName *string `json:"auditRealUserName,omitempty" xml:"auditRealUserName,omitempty"`
	// The remark for the approval request.
	//
	// example:
	//
	// 备注
	AuditRemark *string `json:"auditRemark,omitempty" xml:"auditRemark,omitempty"`
	// The approval status.
	//
	// example:
	//
	// AGREE
	AuditStatus *string `json:"auditStatus,omitempty" xml:"auditStatus,omitempty"`
	// The approval time.
	//
	// example:
	//
	// 1545726028000
	AuditTime *int64 `json:"auditTime,omitempty" xml:"auditTime,omitempty"`
	// The account IDs of the designated reviewers. Multiple IDs are separated by commas (,).
	//
	// example:
	//
	// 123
	AuditUserId *string `json:"auditUserId,omitempty" xml:"auditUserId,omitempty"`
	// The account names of the designated reviewers. Multiple names are separated by commas (,).
	//
	// example:
	//
	// 李四
	AuditUserName *string `json:"auditUserName,omitempty" xml:"auditUserName,omitempty"`
	// The creation time in UTC.
	//
	// example:
	//
	// 1545726028000
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// The primary key ID.
	//
	// example:
	//
	// 1728
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// The information about other associated persons in JSON format.
	//
	// example:
	//
	// {}
	RelationExt *string `json:"relationExt,omitempty" xml:"relationExt,omitempty"`
	// The transaction ID associated with the approval.
	//
	// example:
	//
	// 123
	RelationId *int64 `json:"relationId,omitempty" xml:"relationId,omitempty"`
	// The transaction name associated with the approval. This parameter can be empty.
	//
	// example:
	//
	// t
	RelationName *string `json:"relationName,omitempty" xml:"relationName,omitempty"`
	// The approval type. For example, rule indicates a policy approval.
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
