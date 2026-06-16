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
	// The current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// AE7E6105-7DEB-5125-9B24-DCBC139F6CD2
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The returned object.
	ResultObject []*DescribeRulePageListResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Repeated"`
	// The total number of entries.
	//
	// example:
	//
	// 28
	TotalItem *int32 `json:"totalItem,omitempty" xml:"totalItem,omitempty"`
	// The total number of pages.
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

type DescribeRulePageListResponseBodyResultObject struct {
	// The service authorization type.
	//
	// example:
	//
	// all
	AuthType *string `json:"authType,omitempty" xml:"authType,omitempty"`
	// The approval object.
	ConsoleAudit *DescribeRulePageListResponseBodyResultObjectConsoleAudit `json:"consoleAudit,omitempty" xml:"consoleAudit,omitempty" type:"Struct"`
	// The event code.
	//
	// example:
	//
	// de_arcehq4370
	EventCode *string `json:"eventCode,omitempty" xml:"eventCode,omitempty"`
	// The event name.
	//
	// example:
	//
	// 注册风险
	EventName *string `json:"eventName,omitempty" xml:"eventName,omitempty"`
	// The event type.
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
	// Indicates whether a new version is available.
	//
	// example:
	//
	// true
	HasNewVersion *bool `json:"hasNewVersion,omitempty" xml:"hasNewVersion,omitempty"`
	// The primary key ID of the policy.
	//
	// example:
	//
	// 2793
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// The ID of the main policy.
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
	// NORMAL
	RuleAuthType *string `json:"ruleAuthType,omitempty" xml:"ruleAuthType,omitempty"`
	// The policy ID.
	//
	// example:
	//
	// 102059
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
	// The policy type.
	//
	// example:
	//
	// NORMAL
	RuleType *string `json:"ruleType,omitempty" xml:"ruleType,omitempty"`
	// The primary key ID of the policy version.
	//
	// example:
	//
	// 11300
	RuleVersionId *int64 `json:"ruleVersionId,omitempty" xml:"ruleVersionId,omitempty"`
	// The template ID.
	//
	// example:
	//
	// register
	TemplateId *int64 `json:"templateId,omitempty" xml:"templateId,omitempty"`
	// The version number.
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
	if s.ConsoleAudit != nil {
		if err := s.ConsoleAudit.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRulePageListResponseBodyResultObjectConsoleAudit struct {
	// The UID of the user who approved the request.
	//
	// example:
	//
	// 用户uid
	ApplyUserId *string `json:"applyUserId,omitempty" xml:"applyUserId,omitempty"`
	// The name of the user who approved the request.
	//
	// example:
	//
	// root
	ApplyUserName *string `json:"applyUserName,omitempty" xml:"applyUserName,omitempty"`
	// The approval comment.
	//
	// example:
	//
	// 同意
	AuditMsg *string `json:"auditMsg,omitempty" xml:"auditMsg,omitempty"`
	// The UID of the final reviewer.
	//
	// example:
	//
	// 1728
	AuditRealUserId *string `json:"auditRealUserId,omitempty" xml:"auditRealUserId,omitempty"`
	// The name of the final reviewer.
	//
	// example:
	//
	// root
	AuditRealUserName *string `json:"auditRealUserName,omitempty" xml:"auditRealUserName,omitempty"`
	// The remark from the reviewer.
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
	// The UID of the reviewer.
	//
	// example:
	//
	// 1234xxxx
	AuditUserId *string `json:"auditUserId,omitempty" xml:"auditUserId,omitempty"`
	// The name of the reviewer.
	//
	// example:
	//
	// root
	AuditUserName *string `json:"auditUserName,omitempty" xml:"auditUserName,omitempty"`
	// The creation time.
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
	// The information about other associated users in JSON format.
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
	// 策略1
	RelationName *string `json:"relationName,omitempty" xml:"relationName,omitempty"`
	// The type of the approval. For example, rule indicates a policy approval.
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
