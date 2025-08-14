// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRuleVersionListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeRuleVersionListResponseBody
	GetCurrentPage() *int32
	SetPageSize(v int32) *DescribeRuleVersionListResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeRuleVersionListResponseBody
	GetRequestId() *string
	SetResultObject(v []*DescribeRuleVersionListResponseBodyResultObject) *DescribeRuleVersionListResponseBody
	GetResultObject() []*DescribeRuleVersionListResponseBodyResultObject
	SetTotalItem(v int32) *DescribeRuleVersionListResponseBody
	GetTotalItem() *int32
	SetTotalPage(v int32) *DescribeRuleVersionListResponseBody
	GetTotalPage() *int32
}

type DescribeRuleVersionListResponseBody struct {
	// Current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// Page size, default value is 10
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
	// Returned object
	ResultObject []*DescribeRuleVersionListResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Repeated"`
	// Total number of items
	//
	// example:
	//
	// 3
	TotalItem *int32 `json:"totalItem,omitempty" xml:"totalItem,omitempty"`
	// Total number of pages.
	//
	// example:
	//
	// 7
	TotalPage *int32 `json:"totalPage,omitempty" xml:"totalPage,omitempty"`
}

func (s DescribeRuleVersionListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRuleVersionListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRuleVersionListResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeRuleVersionListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeRuleVersionListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRuleVersionListResponseBody) GetResultObject() []*DescribeRuleVersionListResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeRuleVersionListResponseBody) GetTotalItem() *int32 {
	return s.TotalItem
}

func (s *DescribeRuleVersionListResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *DescribeRuleVersionListResponseBody) SetCurrentPage(v int32) *DescribeRuleVersionListResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeRuleVersionListResponseBody) SetPageSize(v int32) *DescribeRuleVersionListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeRuleVersionListResponseBody) SetRequestId(v string) *DescribeRuleVersionListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRuleVersionListResponseBody) SetResultObject(v []*DescribeRuleVersionListResponseBodyResultObject) *DescribeRuleVersionListResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeRuleVersionListResponseBody) SetTotalItem(v int32) *DescribeRuleVersionListResponseBody {
	s.TotalItem = &v
	return s
}

func (s *DescribeRuleVersionListResponseBody) SetTotalPage(v int32) *DescribeRuleVersionListResponseBody {
	s.TotalPage = &v
	return s
}

func (s *DescribeRuleVersionListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeRuleVersionListResponseBodyResultObject struct {
	// Audit object
	ConsoleAudit *DescribeRuleVersionListResponseBodyResultObjectConsoleAudit `json:"consoleAudit,omitempty" xml:"consoleAudit,omitempty" type:"Struct"`
	// Console rule ID.
	//
	// example:
	//
	// 6715
	ConsoleRuleId *int64 `json:"consoleRuleId,omitempty" xml:"consoleRuleId,omitempty"`
	// Associated policy remarks
	//
	// example:
	//
	// 备注
	ConsoleRuleMemo *string `json:"consoleRuleMemo,omitempty" xml:"consoleRuleMemo,omitempty"`
	// Associated policy name
	//
	// example:
	//
	// 营销风险识别
	ConsoleRuleName *string `json:"consoleRuleName,omitempty" xml:"consoleRuleName,omitempty"`
	// Event code
	//
	// example:
	//
	// de_aszbjb7236
	EventCode *string `json:"eventCode,omitempty" xml:"eventCode,omitempty"`
	// Event type
	//
	// example:
	//
	// MAIN
	EventType *string `json:"eventType,omitempty" xml:"eventType,omitempty"`
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
	// Primary key ID of the policy
	//
	// example:
	//
	// 376773
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// The user who last operated.
	//
	// example:
	//
	// 1519714049632764
	LastOperator *string `json:"lastOperator,omitempty" xml:"lastOperator,omitempty"`
	// Policy priority, the higher the number, the higher the priority.
	//
	// example:
	//
	// 10
	Priority *int64 `json:"priority,omitempty" xml:"priority,omitempty"`
	// Policy ID
	//
	// example:
	//
	// 102224
	RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty"`
	// Policy status
	//
	// example:
	//
	// RUNNING
	RuleStatus *string `json:"ruleStatus,omitempty" xml:"ruleStatus,omitempty"`
	// Rule Type
	//
	// example:
	//
	// DEFAULT
	RuleType *string `json:"ruleType,omitempty" xml:"ruleType,omitempty"`
	// Version number
	//
	// example:
	//
	// 2
	Version *int64 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s DescribeRuleVersionListResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeRuleVersionListResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeRuleVersionListResponseBodyResultObject) GetConsoleAudit() *DescribeRuleVersionListResponseBodyResultObjectConsoleAudit {
	return s.ConsoleAudit
}

func (s *DescribeRuleVersionListResponseBodyResultObject) GetConsoleRuleId() *int64 {
	return s.ConsoleRuleId
}

func (s *DescribeRuleVersionListResponseBodyResultObject) GetConsoleRuleMemo() *string {
	return s.ConsoleRuleMemo
}

func (s *DescribeRuleVersionListResponseBodyResultObject) GetConsoleRuleName() *string {
	return s.ConsoleRuleName
}

func (s *DescribeRuleVersionListResponseBodyResultObject) GetEventCode() *string {
	return s.EventCode
}

func (s *DescribeRuleVersionListResponseBodyResultObject) GetEventType() *string {
	return s.EventType
}

func (s *DescribeRuleVersionListResponseBodyResultObject) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *DescribeRuleVersionListResponseBodyResultObject) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *DescribeRuleVersionListResponseBodyResultObject) GetId() *int64 {
	return s.Id
}

func (s *DescribeRuleVersionListResponseBodyResultObject) GetLastOperator() *string {
	return s.LastOperator
}

func (s *DescribeRuleVersionListResponseBodyResultObject) GetPriority() *int64 {
	return s.Priority
}

func (s *DescribeRuleVersionListResponseBodyResultObject) GetRuleId() *string {
	return s.RuleId
}

func (s *DescribeRuleVersionListResponseBodyResultObject) GetRuleStatus() *string {
	return s.RuleStatus
}

func (s *DescribeRuleVersionListResponseBodyResultObject) GetRuleType() *string {
	return s.RuleType
}

func (s *DescribeRuleVersionListResponseBodyResultObject) GetVersion() *int64 {
	return s.Version
}

func (s *DescribeRuleVersionListResponseBodyResultObject) SetConsoleAudit(v *DescribeRuleVersionListResponseBodyResultObjectConsoleAudit) *DescribeRuleVersionListResponseBodyResultObject {
	s.ConsoleAudit = v
	return s
}

func (s *DescribeRuleVersionListResponseBodyResultObject) SetConsoleRuleId(v int64) *DescribeRuleVersionListResponseBodyResultObject {
	s.ConsoleRuleId = &v
	return s
}

func (s *DescribeRuleVersionListResponseBodyResultObject) SetConsoleRuleMemo(v string) *DescribeRuleVersionListResponseBodyResultObject {
	s.ConsoleRuleMemo = &v
	return s
}

func (s *DescribeRuleVersionListResponseBodyResultObject) SetConsoleRuleName(v string) *DescribeRuleVersionListResponseBodyResultObject {
	s.ConsoleRuleName = &v
	return s
}

func (s *DescribeRuleVersionListResponseBodyResultObject) SetEventCode(v string) *DescribeRuleVersionListResponseBodyResultObject {
	s.EventCode = &v
	return s
}

func (s *DescribeRuleVersionListResponseBodyResultObject) SetEventType(v string) *DescribeRuleVersionListResponseBodyResultObject {
	s.EventType = &v
	return s
}

func (s *DescribeRuleVersionListResponseBodyResultObject) SetGmtCreate(v int64) *DescribeRuleVersionListResponseBodyResultObject {
	s.GmtCreate = &v
	return s
}

func (s *DescribeRuleVersionListResponseBodyResultObject) SetGmtModified(v int64) *DescribeRuleVersionListResponseBodyResultObject {
	s.GmtModified = &v
	return s
}

func (s *DescribeRuleVersionListResponseBodyResultObject) SetId(v int64) *DescribeRuleVersionListResponseBodyResultObject {
	s.Id = &v
	return s
}

func (s *DescribeRuleVersionListResponseBodyResultObject) SetLastOperator(v string) *DescribeRuleVersionListResponseBodyResultObject {
	s.LastOperator = &v
	return s
}

func (s *DescribeRuleVersionListResponseBodyResultObject) SetPriority(v int64) *DescribeRuleVersionListResponseBodyResultObject {
	s.Priority = &v
	return s
}

func (s *DescribeRuleVersionListResponseBodyResultObject) SetRuleId(v string) *DescribeRuleVersionListResponseBodyResultObject {
	s.RuleId = &v
	return s
}

func (s *DescribeRuleVersionListResponseBodyResultObject) SetRuleStatus(v string) *DescribeRuleVersionListResponseBodyResultObject {
	s.RuleStatus = &v
	return s
}

func (s *DescribeRuleVersionListResponseBodyResultObject) SetRuleType(v string) *DescribeRuleVersionListResponseBodyResultObject {
	s.RuleType = &v
	return s
}

func (s *DescribeRuleVersionListResponseBodyResultObject) SetVersion(v int64) *DescribeRuleVersionListResponseBodyResultObject {
	s.Version = &v
	return s
}

func (s *DescribeRuleVersionListResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}

type DescribeRuleVersionListResponseBodyResultObjectConsoleAudit struct {
	// Initiator UID
	//
	// example:
	//
	// 1519714049632764
	ApplyUserId *string `json:"applyUserId,omitempty" xml:"applyUserId,omitempty"`
	// Initiator name
	//
	// example:
	//
	// root
	ApplyUserName *string `json:"applyUserName,omitempty" xml:"applyUserName,omitempty"`
	// Approval comment
	//
	// example:
	//
	// 同意
	AuditMsg *string `json:"auditMsg,omitempty" xml:"auditMsg,omitempty"`
	// Final approver UID
	//
	// example:
	//
	// 1519714049632764
	AuditRealUserId *string `json:"auditRealUserId,omitempty" xml:"auditRealUserId,omitempty"`
	// Final approver name
	//
	// example:
	//
	// root
	AuditRealUserName *string `json:"auditRealUserName,omitempty" xml:"auditRealUserName,omitempty"`
	// Approver\\"s remarks.
	//
	// example:
	//
	// 备注
	AuditRemark *string `json:"auditRemark,omitempty" xml:"auditRemark,omitempty"`
	// Approval status
	//
	// example:
	//
	// AGREE
	AuditStatus *string `json:"auditStatus,omitempty" xml:"auditStatus,omitempty"`
	// Approval time.
	//
	// example:
	//
	// 1545726028000
	AuditTime *int64 `json:"auditTime,omitempty" xml:"auditTime,omitempty"`
	// Designated approver UID
	//
	// example:
	//
	// 1519714049632764
	AuditUserId *string `json:"auditUserId,omitempty" xml:"auditUserId,omitempty"`
	// Designated auditor\\"s name
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
	// Information of related others (in JSON format)
	//
	// example:
	//
	// {}
	RelationExt *string `json:"relationExt,omitempty" xml:"relationExt,omitempty"`
	// ID of the associated transaction
	//
	// example:
	//
	// 123
	RelationId *int64 `json:"relationId,omitempty" xml:"relationId,omitempty"`
	// Name of the associated item
	//
	// example:
	//
	// 营销风险识别
	RelationName *string `json:"relationName,omitempty" xml:"relationName,omitempty"`
	// Type of approval (e.g., `rule` for policy approval)
	//
	// example:
	//
	// RULE
	RelationType *string `json:"relationType,omitempty" xml:"relationType,omitempty"`
}

func (s DescribeRuleVersionListResponseBodyResultObjectConsoleAudit) String() string {
	return dara.Prettify(s)
}

func (s DescribeRuleVersionListResponseBodyResultObjectConsoleAudit) GoString() string {
	return s.String()
}

func (s *DescribeRuleVersionListResponseBodyResultObjectConsoleAudit) GetApplyUserId() *string {
	return s.ApplyUserId
}

func (s *DescribeRuleVersionListResponseBodyResultObjectConsoleAudit) GetApplyUserName() *string {
	return s.ApplyUserName
}

func (s *DescribeRuleVersionListResponseBodyResultObjectConsoleAudit) GetAuditMsg() *string {
	return s.AuditMsg
}

func (s *DescribeRuleVersionListResponseBodyResultObjectConsoleAudit) GetAuditRealUserId() *string {
	return s.AuditRealUserId
}

func (s *DescribeRuleVersionListResponseBodyResultObjectConsoleAudit) GetAuditRealUserName() *string {
	return s.AuditRealUserName
}

func (s *DescribeRuleVersionListResponseBodyResultObjectConsoleAudit) GetAuditRemark() *string {
	return s.AuditRemark
}

func (s *DescribeRuleVersionListResponseBodyResultObjectConsoleAudit) GetAuditStatus() *string {
	return s.AuditStatus
}

func (s *DescribeRuleVersionListResponseBodyResultObjectConsoleAudit) GetAuditTime() *int64 {
	return s.AuditTime
}

func (s *DescribeRuleVersionListResponseBodyResultObjectConsoleAudit) GetAuditUserId() *string {
	return s.AuditUserId
}

func (s *DescribeRuleVersionListResponseBodyResultObjectConsoleAudit) GetAuditUserName() *string {
	return s.AuditUserName
}

func (s *DescribeRuleVersionListResponseBodyResultObjectConsoleAudit) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *DescribeRuleVersionListResponseBodyResultObjectConsoleAudit) GetId() *int64 {
	return s.Id
}

func (s *DescribeRuleVersionListResponseBodyResultObjectConsoleAudit) GetRelationExt() *string {
	return s.RelationExt
}

func (s *DescribeRuleVersionListResponseBodyResultObjectConsoleAudit) GetRelationId() *int64 {
	return s.RelationId
}

func (s *DescribeRuleVersionListResponseBodyResultObjectConsoleAudit) GetRelationName() *string {
	return s.RelationName
}

func (s *DescribeRuleVersionListResponseBodyResultObjectConsoleAudit) GetRelationType() *string {
	return s.RelationType
}

func (s *DescribeRuleVersionListResponseBodyResultObjectConsoleAudit) SetApplyUserId(v string) *DescribeRuleVersionListResponseBodyResultObjectConsoleAudit {
	s.ApplyUserId = &v
	return s
}

func (s *DescribeRuleVersionListResponseBodyResultObjectConsoleAudit) SetApplyUserName(v string) *DescribeRuleVersionListResponseBodyResultObjectConsoleAudit {
	s.ApplyUserName = &v
	return s
}

func (s *DescribeRuleVersionListResponseBodyResultObjectConsoleAudit) SetAuditMsg(v string) *DescribeRuleVersionListResponseBodyResultObjectConsoleAudit {
	s.AuditMsg = &v
	return s
}

func (s *DescribeRuleVersionListResponseBodyResultObjectConsoleAudit) SetAuditRealUserId(v string) *DescribeRuleVersionListResponseBodyResultObjectConsoleAudit {
	s.AuditRealUserId = &v
	return s
}

func (s *DescribeRuleVersionListResponseBodyResultObjectConsoleAudit) SetAuditRealUserName(v string) *DescribeRuleVersionListResponseBodyResultObjectConsoleAudit {
	s.AuditRealUserName = &v
	return s
}

func (s *DescribeRuleVersionListResponseBodyResultObjectConsoleAudit) SetAuditRemark(v string) *DescribeRuleVersionListResponseBodyResultObjectConsoleAudit {
	s.AuditRemark = &v
	return s
}

func (s *DescribeRuleVersionListResponseBodyResultObjectConsoleAudit) SetAuditStatus(v string) *DescribeRuleVersionListResponseBodyResultObjectConsoleAudit {
	s.AuditStatus = &v
	return s
}

func (s *DescribeRuleVersionListResponseBodyResultObjectConsoleAudit) SetAuditTime(v int64) *DescribeRuleVersionListResponseBodyResultObjectConsoleAudit {
	s.AuditTime = &v
	return s
}

func (s *DescribeRuleVersionListResponseBodyResultObjectConsoleAudit) SetAuditUserId(v string) *DescribeRuleVersionListResponseBodyResultObjectConsoleAudit {
	s.AuditUserId = &v
	return s
}

func (s *DescribeRuleVersionListResponseBodyResultObjectConsoleAudit) SetAuditUserName(v string) *DescribeRuleVersionListResponseBodyResultObjectConsoleAudit {
	s.AuditUserName = &v
	return s
}

func (s *DescribeRuleVersionListResponseBodyResultObjectConsoleAudit) SetGmtCreate(v int64) *DescribeRuleVersionListResponseBodyResultObjectConsoleAudit {
	s.GmtCreate = &v
	return s
}

func (s *DescribeRuleVersionListResponseBodyResultObjectConsoleAudit) SetId(v int64) *DescribeRuleVersionListResponseBodyResultObjectConsoleAudit {
	s.Id = &v
	return s
}

func (s *DescribeRuleVersionListResponseBodyResultObjectConsoleAudit) SetRelationExt(v string) *DescribeRuleVersionListResponseBodyResultObjectConsoleAudit {
	s.RelationExt = &v
	return s
}

func (s *DescribeRuleVersionListResponseBodyResultObjectConsoleAudit) SetRelationId(v int64) *DescribeRuleVersionListResponseBodyResultObjectConsoleAudit {
	s.RelationId = &v
	return s
}

func (s *DescribeRuleVersionListResponseBodyResultObjectConsoleAudit) SetRelationName(v string) *DescribeRuleVersionListResponseBodyResultObjectConsoleAudit {
	s.RelationName = &v
	return s
}

func (s *DescribeRuleVersionListResponseBodyResultObjectConsoleAudit) SetRelationType(v string) *DescribeRuleVersionListResponseBodyResultObjectConsoleAudit {
	s.RelationType = &v
	return s
}

func (s *DescribeRuleVersionListResponseBodyResultObjectConsoleAudit) Validate() error {
	return dara.Validate(s)
}
