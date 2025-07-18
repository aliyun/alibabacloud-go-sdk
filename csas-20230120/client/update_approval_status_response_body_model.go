// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApprovalStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApproval(v []*UpdateApprovalStatusResponseBodyApproval) *UpdateApprovalStatusResponseBody
	GetApproval() []*UpdateApprovalStatusResponseBodyApproval
	SetRequestId(v string) *UpdateApprovalStatusResponseBody
	GetRequestId() *string
}

type UpdateApprovalStatusResponseBody struct {
	Approval []*UpdateApprovalStatusResponseBodyApproval `json:"Approval,omitempty" xml:"Approval,omitempty" type:"Repeated"`
	// example:
	//
	// 58D6B23E-E5DA-5418-8F61-51A3B5A30049
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateApprovalStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateApprovalStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateApprovalStatusResponseBody) GetApproval() []*UpdateApprovalStatusResponseBodyApproval {
	return s.Approval
}

func (s *UpdateApprovalStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateApprovalStatusResponseBody) SetApproval(v []*UpdateApprovalStatusResponseBodyApproval) *UpdateApprovalStatusResponseBody {
	s.Approval = v
	return s
}

func (s *UpdateApprovalStatusResponseBody) SetRequestId(v string) *UpdateApprovalStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateApprovalStatusResponseBody) Validate() error {
	return dara.Validate(s)
}

type UpdateApprovalStatusResponseBodyApproval struct {
	ApprovalDetail *string `json:"ApprovalDetail,omitempty" xml:"ApprovalDetail,omitempty"`
	// example:
	//
	// approval-165e6738ad9d****
	ApprovalId         *string                                                       `json:"ApprovalId,omitempty" xml:"ApprovalId,omitempty"`
	ApprovalProgresses []*UpdateApprovalStatusResponseBodyApprovalApprovalProgresses `json:"ApprovalProgresses,omitempty" xml:"ApprovalProgresses,omitempty" type:"Repeated"`
	// example:
	//
	// 2022-11-15 22:11:55
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 36efa42d-2c32-c4dc-e3fc-8541e33a****
	CreatorDevTag *string `json:"CreatorDevTag,omitempty" xml:"CreatorDevTag,omitempty"`
	// example:
	//
	// su_e8f218fb171edd167c2ad917d21f53148bdefc510ca1f3c3cc0249d3643d****
	CreatorUserId *string `json:"CreatorUserId,omitempty" xml:"CreatorUserId,omitempty"`
	// example:
	//
	// 1757952000
	EndTimestamp *int64 `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	// example:
	//
	// DlpSend
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	// example:
	//
	// approval-process-fcc351b8a95b****
	ProcessId   *string `json:"ProcessId,omitempty" xml:"ProcessId,omitempty"`
	ProcessName *string `json:"ProcessName,omitempty" xml:"ProcessName,omitempty"`
	Reason      *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// example:
	//
	// {"form": {"labelCol": 6,"wrapperCol": 12}}
	SchemaContent *string `json:"SchemaContent,omitempty" xml:"SchemaContent,omitempty"`
	// example:
	//
	// approval-schema-090134f1ebff****
	SchemaId   *string `json:"SchemaId,omitempty" xml:"SchemaId,omitempty"`
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// example:
	//
	// Pending
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateApprovalStatusResponseBodyApproval) String() string {
	return dara.Prettify(s)
}

func (s UpdateApprovalStatusResponseBodyApproval) GoString() string {
	return s.String()
}

func (s *UpdateApprovalStatusResponseBodyApproval) GetApprovalDetail() *string {
	return s.ApprovalDetail
}

func (s *UpdateApprovalStatusResponseBodyApproval) GetApprovalId() *string {
	return s.ApprovalId
}

func (s *UpdateApprovalStatusResponseBodyApproval) GetApprovalProgresses() []*UpdateApprovalStatusResponseBodyApprovalApprovalProgresses {
	return s.ApprovalProgresses
}

func (s *UpdateApprovalStatusResponseBodyApproval) GetCreateTime() *string {
	return s.CreateTime
}

func (s *UpdateApprovalStatusResponseBodyApproval) GetCreatorDevTag() *string {
	return s.CreatorDevTag
}

func (s *UpdateApprovalStatusResponseBodyApproval) GetCreatorUserId() *string {
	return s.CreatorUserId
}

func (s *UpdateApprovalStatusResponseBodyApproval) GetEndTimestamp() *int64 {
	return s.EndTimestamp
}

func (s *UpdateApprovalStatusResponseBodyApproval) GetPolicyType() *string {
	return s.PolicyType
}

func (s *UpdateApprovalStatusResponseBodyApproval) GetProcessId() *string {
	return s.ProcessId
}

func (s *UpdateApprovalStatusResponseBodyApproval) GetProcessName() *string {
	return s.ProcessName
}

func (s *UpdateApprovalStatusResponseBodyApproval) GetReason() *string {
	return s.Reason
}

func (s *UpdateApprovalStatusResponseBodyApproval) GetSchemaContent() *string {
	return s.SchemaContent
}

func (s *UpdateApprovalStatusResponseBodyApproval) GetSchemaId() *string {
	return s.SchemaId
}

func (s *UpdateApprovalStatusResponseBodyApproval) GetSchemaName() *string {
	return s.SchemaName
}

func (s *UpdateApprovalStatusResponseBodyApproval) GetStatus() *string {
	return s.Status
}

func (s *UpdateApprovalStatusResponseBodyApproval) SetApprovalDetail(v string) *UpdateApprovalStatusResponseBodyApproval {
	s.ApprovalDetail = &v
	return s
}

func (s *UpdateApprovalStatusResponseBodyApproval) SetApprovalId(v string) *UpdateApprovalStatusResponseBodyApproval {
	s.ApprovalId = &v
	return s
}

func (s *UpdateApprovalStatusResponseBodyApproval) SetApprovalProgresses(v []*UpdateApprovalStatusResponseBodyApprovalApprovalProgresses) *UpdateApprovalStatusResponseBodyApproval {
	s.ApprovalProgresses = v
	return s
}

func (s *UpdateApprovalStatusResponseBodyApproval) SetCreateTime(v string) *UpdateApprovalStatusResponseBodyApproval {
	s.CreateTime = &v
	return s
}

func (s *UpdateApprovalStatusResponseBodyApproval) SetCreatorDevTag(v string) *UpdateApprovalStatusResponseBodyApproval {
	s.CreatorDevTag = &v
	return s
}

func (s *UpdateApprovalStatusResponseBodyApproval) SetCreatorUserId(v string) *UpdateApprovalStatusResponseBodyApproval {
	s.CreatorUserId = &v
	return s
}

func (s *UpdateApprovalStatusResponseBodyApproval) SetEndTimestamp(v int64) *UpdateApprovalStatusResponseBodyApproval {
	s.EndTimestamp = &v
	return s
}

func (s *UpdateApprovalStatusResponseBodyApproval) SetPolicyType(v string) *UpdateApprovalStatusResponseBodyApproval {
	s.PolicyType = &v
	return s
}

func (s *UpdateApprovalStatusResponseBodyApproval) SetProcessId(v string) *UpdateApprovalStatusResponseBodyApproval {
	s.ProcessId = &v
	return s
}

func (s *UpdateApprovalStatusResponseBodyApproval) SetProcessName(v string) *UpdateApprovalStatusResponseBodyApproval {
	s.ProcessName = &v
	return s
}

func (s *UpdateApprovalStatusResponseBodyApproval) SetReason(v string) *UpdateApprovalStatusResponseBodyApproval {
	s.Reason = &v
	return s
}

func (s *UpdateApprovalStatusResponseBodyApproval) SetSchemaContent(v string) *UpdateApprovalStatusResponseBodyApproval {
	s.SchemaContent = &v
	return s
}

func (s *UpdateApprovalStatusResponseBodyApproval) SetSchemaId(v string) *UpdateApprovalStatusResponseBodyApproval {
	s.SchemaId = &v
	return s
}

func (s *UpdateApprovalStatusResponseBodyApproval) SetSchemaName(v string) *UpdateApprovalStatusResponseBodyApproval {
	s.SchemaName = &v
	return s
}

func (s *UpdateApprovalStatusResponseBodyApproval) SetStatus(v string) *UpdateApprovalStatusResponseBodyApproval {
	s.Status = &v
	return s
}

func (s *UpdateApprovalStatusResponseBodyApproval) Validate() error {
	return dara.Validate(s)
}

type UpdateApprovalStatusResponseBodyApprovalApprovalProgresses struct {
	// example:
	//
	// Approve
	Action  *string `json:"Action,omitempty" xml:"Action,omitempty"`
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// example:
	//
	// su_e8f218fb171edd167c2ad917d21f53148bdefc510ca1f3c3cc0249d3643d****
	Executor  *string                                                                `json:"Executor,omitempty" xml:"Executor,omitempty"`
	Operators []*UpdateApprovalStatusResponseBodyApprovalApprovalProgressesOperators `json:"Operators,omitempty" xml:"Operators,omitempty" type:"Repeated"`
	// example:
	//
	// Approved
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 1736752000
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s UpdateApprovalStatusResponseBodyApprovalApprovalProgresses) String() string {
	return dara.Prettify(s)
}

func (s UpdateApprovalStatusResponseBodyApprovalApprovalProgresses) GoString() string {
	return s.String()
}

func (s *UpdateApprovalStatusResponseBodyApprovalApprovalProgresses) GetAction() *string {
	return s.Action
}

func (s *UpdateApprovalStatusResponseBodyApprovalApprovalProgresses) GetComment() *string {
	return s.Comment
}

func (s *UpdateApprovalStatusResponseBodyApprovalApprovalProgresses) GetExecutor() *string {
	return s.Executor
}

func (s *UpdateApprovalStatusResponseBodyApprovalApprovalProgresses) GetOperators() []*UpdateApprovalStatusResponseBodyApprovalApprovalProgressesOperators {
	return s.Operators
}

func (s *UpdateApprovalStatusResponseBodyApprovalApprovalProgresses) GetStatus() *string {
	return s.Status
}

func (s *UpdateApprovalStatusResponseBodyApprovalApprovalProgresses) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *UpdateApprovalStatusResponseBodyApprovalApprovalProgresses) SetAction(v string) *UpdateApprovalStatusResponseBodyApprovalApprovalProgresses {
	s.Action = &v
	return s
}

func (s *UpdateApprovalStatusResponseBodyApprovalApprovalProgresses) SetComment(v string) *UpdateApprovalStatusResponseBodyApprovalApprovalProgresses {
	s.Comment = &v
	return s
}

func (s *UpdateApprovalStatusResponseBodyApprovalApprovalProgresses) SetExecutor(v string) *UpdateApprovalStatusResponseBodyApprovalApprovalProgresses {
	s.Executor = &v
	return s
}

func (s *UpdateApprovalStatusResponseBodyApprovalApprovalProgresses) SetOperators(v []*UpdateApprovalStatusResponseBodyApprovalApprovalProgressesOperators) *UpdateApprovalStatusResponseBodyApprovalApprovalProgresses {
	s.Operators = v
	return s
}

func (s *UpdateApprovalStatusResponseBodyApprovalApprovalProgresses) SetStatus(v string) *UpdateApprovalStatusResponseBodyApprovalApprovalProgresses {
	s.Status = &v
	return s
}

func (s *UpdateApprovalStatusResponseBodyApprovalApprovalProgresses) SetTimestamp(v int64) *UpdateApprovalStatusResponseBodyApprovalApprovalProgresses {
	s.Timestamp = &v
	return s
}

func (s *UpdateApprovalStatusResponseBodyApprovalApprovalProgresses) Validate() error {
	return dara.Validate(s)
}

type UpdateApprovalStatusResponseBodyApprovalApprovalProgressesOperators struct {
	// example:
	//
	// su_e8f218fb171edd167c2ad917d21f53148bdefc510ca1f3c3cc0249d3643d****
	SaseUserId *string `json:"SaseUserId,omitempty" xml:"SaseUserId,omitempty"`
	Username   *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s UpdateApprovalStatusResponseBodyApprovalApprovalProgressesOperators) String() string {
	return dara.Prettify(s)
}

func (s UpdateApprovalStatusResponseBodyApprovalApprovalProgressesOperators) GoString() string {
	return s.String()
}

func (s *UpdateApprovalStatusResponseBodyApprovalApprovalProgressesOperators) GetSaseUserId() *string {
	return s.SaseUserId
}

func (s *UpdateApprovalStatusResponseBodyApprovalApprovalProgressesOperators) GetUsername() *string {
	return s.Username
}

func (s *UpdateApprovalStatusResponseBodyApprovalApprovalProgressesOperators) SetSaseUserId(v string) *UpdateApprovalStatusResponseBodyApprovalApprovalProgressesOperators {
	s.SaseUserId = &v
	return s
}

func (s *UpdateApprovalStatusResponseBodyApprovalApprovalProgressesOperators) SetUsername(v string) *UpdateApprovalStatusResponseBodyApprovalApprovalProgressesOperators {
	s.Username = &v
	return s
}

func (s *UpdateApprovalStatusResponseBodyApprovalApprovalProgressesOperators) Validate() error {
	return dara.Validate(s)
}
