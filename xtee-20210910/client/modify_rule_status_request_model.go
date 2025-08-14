// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRuleStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *ModifyRuleStatusRequest
	GetLang() *string
	SetApplyUserId(v string) *ModifyRuleStatusRequest
	GetApplyUserId() *string
	SetApplyUserName(v string) *ModifyRuleStatusRequest
	GetApplyUserName() *string
	SetAuditRemark(v string) *ModifyRuleStatusRequest
	GetAuditRemark() *string
	SetAuditUserId(v string) *ModifyRuleStatusRequest
	GetAuditUserId() *string
	SetAuditUserName(v string) *ModifyRuleStatusRequest
	GetAuditUserName() *string
	SetConsoleRuleId(v int64) *ModifyRuleStatusRequest
	GetConsoleRuleId() *int64
	SetEventType(v string) *ModifyRuleStatusRequest
	GetEventType() *string
	SetRegId(v string) *ModifyRuleStatusRequest
	GetRegId() *string
	SetRuleAuditType(v string) *ModifyRuleStatusRequest
	GetRuleAuditType() *string
	SetRuleId(v string) *ModifyRuleStatusRequest
	GetRuleId() *string
	SetRuleVersionId(v int64) *ModifyRuleStatusRequest
	GetRuleVersionId() *int64
}

type ModifyRuleStatusRequest struct {
	// Set the language type for requests and received messages, default value is **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// UID of the applicant.
	//
	// example:
	//
	// 1519714049632764
	ApplyUserId *string `json:"applyUserId,omitempty" xml:"applyUserId,omitempty"`
	// Name of the applicant.
	//
	// example:
	//
	// root
	ApplyUserName *string `json:"applyUserName,omitempty" xml:"applyUserName,omitempty"`
	// Approval remarks.
	//
	// example:
	//
	// 申请发布上线
	AuditRemark *string `json:"auditRemark,omitempty" xml:"auditRemark,omitempty"`
	// UID of the auditor.
	//
	// example:
	//
	// 1519714049632764
	AuditUserId *string `json:"auditUserId,omitempty" xml:"auditUserId,omitempty"`
	// Name of the auditor.
	//
	// example:
	//
	// root
	AuditUserName *string `json:"auditUserName,omitempty" xml:"auditUserName,omitempty"`
	// Primary key ID of the policy.
	//
	// example:
	//
	// 6843
	ConsoleRuleId *int64 `json:"consoleRuleId,omitempty" xml:"consoleRuleId,omitempty"`
	// Event type.
	//
	// example:
	//
	// MAIN
	EventType *string `json:"eventType,omitempty" xml:"eventType,omitempty"`
	// Region code.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
	// Audit status.
	//
	// example:
	//
	// DRAFT_TO_RUNNING
	RuleAuditType *string `json:"ruleAuditType,omitempty" xml:"ruleAuditType,omitempty"`
	// Policy ID.
	//
	// example:
	//
	// 101544
	RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty"`
	// Primary key ID of the policy version.
	//
	// example:
	//
	// 11519
	RuleVersionId *int64 `json:"ruleVersionId,omitempty" xml:"ruleVersionId,omitempty"`
}

func (s ModifyRuleStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyRuleStatusRequest) GoString() string {
	return s.String()
}

func (s *ModifyRuleStatusRequest) GetLang() *string {
	return s.Lang
}

func (s *ModifyRuleStatusRequest) GetApplyUserId() *string {
	return s.ApplyUserId
}

func (s *ModifyRuleStatusRequest) GetApplyUserName() *string {
	return s.ApplyUserName
}

func (s *ModifyRuleStatusRequest) GetAuditRemark() *string {
	return s.AuditRemark
}

func (s *ModifyRuleStatusRequest) GetAuditUserId() *string {
	return s.AuditUserId
}

func (s *ModifyRuleStatusRequest) GetAuditUserName() *string {
	return s.AuditUserName
}

func (s *ModifyRuleStatusRequest) GetConsoleRuleId() *int64 {
	return s.ConsoleRuleId
}

func (s *ModifyRuleStatusRequest) GetEventType() *string {
	return s.EventType
}

func (s *ModifyRuleStatusRequest) GetRegId() *string {
	return s.RegId
}

func (s *ModifyRuleStatusRequest) GetRuleAuditType() *string {
	return s.RuleAuditType
}

func (s *ModifyRuleStatusRequest) GetRuleId() *string {
	return s.RuleId
}

func (s *ModifyRuleStatusRequest) GetRuleVersionId() *int64 {
	return s.RuleVersionId
}

func (s *ModifyRuleStatusRequest) SetLang(v string) *ModifyRuleStatusRequest {
	s.Lang = &v
	return s
}

func (s *ModifyRuleStatusRequest) SetApplyUserId(v string) *ModifyRuleStatusRequest {
	s.ApplyUserId = &v
	return s
}

func (s *ModifyRuleStatusRequest) SetApplyUserName(v string) *ModifyRuleStatusRequest {
	s.ApplyUserName = &v
	return s
}

func (s *ModifyRuleStatusRequest) SetAuditRemark(v string) *ModifyRuleStatusRequest {
	s.AuditRemark = &v
	return s
}

func (s *ModifyRuleStatusRequest) SetAuditUserId(v string) *ModifyRuleStatusRequest {
	s.AuditUserId = &v
	return s
}

func (s *ModifyRuleStatusRequest) SetAuditUserName(v string) *ModifyRuleStatusRequest {
	s.AuditUserName = &v
	return s
}

func (s *ModifyRuleStatusRequest) SetConsoleRuleId(v int64) *ModifyRuleStatusRequest {
	s.ConsoleRuleId = &v
	return s
}

func (s *ModifyRuleStatusRequest) SetEventType(v string) *ModifyRuleStatusRequest {
	s.EventType = &v
	return s
}

func (s *ModifyRuleStatusRequest) SetRegId(v string) *ModifyRuleStatusRequest {
	s.RegId = &v
	return s
}

func (s *ModifyRuleStatusRequest) SetRuleAuditType(v string) *ModifyRuleStatusRequest {
	s.RuleAuditType = &v
	return s
}

func (s *ModifyRuleStatusRequest) SetRuleId(v string) *ModifyRuleStatusRequest {
	s.RuleId = &v
	return s
}

func (s *ModifyRuleStatusRequest) SetRuleVersionId(v int64) *ModifyRuleStatusRequest {
	s.RuleVersionId = &v
	return s
}

func (s *ModifyRuleStatusRequest) Validate() error {
	return dara.Validate(s)
}
