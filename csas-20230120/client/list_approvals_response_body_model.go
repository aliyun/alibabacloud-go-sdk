// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApprovalsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApprovals(v []*ListApprovalsResponseBodyApprovals) *ListApprovalsResponseBody
	GetApprovals() []*ListApprovalsResponseBodyApprovals
	SetRequestId(v string) *ListApprovalsResponseBody
	GetRequestId() *string
	SetTotalNum(v string) *ListApprovalsResponseBody
	GetTotalNum() *string
}

type ListApprovalsResponseBody struct {
	Approvals []*ListApprovalsResponseBodyApprovals `json:"Approvals,omitempty" xml:"Approvals,omitempty" type:"Repeated"`
	// example:
	//
	// 6965F5BA-53B6-5650-A708-51F090F843BB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalNum *string `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s ListApprovalsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListApprovalsResponseBody) GoString() string {
	return s.String()
}

func (s *ListApprovalsResponseBody) GetApprovals() []*ListApprovalsResponseBodyApprovals {
	return s.Approvals
}

func (s *ListApprovalsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListApprovalsResponseBody) GetTotalNum() *string {
	return s.TotalNum
}

func (s *ListApprovalsResponseBody) SetApprovals(v []*ListApprovalsResponseBodyApprovals) *ListApprovalsResponseBody {
	s.Approvals = v
	return s
}

func (s *ListApprovalsResponseBody) SetRequestId(v string) *ListApprovalsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListApprovalsResponseBody) SetTotalNum(v string) *ListApprovalsResponseBody {
	s.TotalNum = &v
	return s
}

func (s *ListApprovalsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListApprovalsResponseBodyApprovals struct {
	ApprovalDetail *string `json:"ApprovalDetail,omitempty" xml:"ApprovalDetail,omitempty"`
	// example:
	//
	// approval-872b5e911b35****
	ApprovalId         *string                                                 `json:"ApprovalId,omitempty" xml:"ApprovalId,omitempty"`
	ApprovalProgresses []*ListApprovalsResponseBodyApprovalsApprovalProgresses `json:"ApprovalProgresses,omitempty" xml:"ApprovalProgresses,omitempty" type:"Repeated"`
	// example:
	//
	// 2022-11-15 22:11:55
	CreateTime        *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreatorDepartment *string `json:"CreatorDepartment,omitempty" xml:"CreatorDepartment,omitempty"`
	// example:
	//
	// 36efa42d-2c32-c4dc-e3fc-8541e33a****
	CreatorDevTag *string `json:"CreatorDevTag,omitempty" xml:"CreatorDevTag,omitempty"`
	// example:
	//
	// su_e8f218fb171edd167c2ad917d21f53148bdefc510ca1f3c3cc0249d3643d****
	CreatorUserId   *string `json:"CreatorUserId,omitempty" xml:"CreatorUserId,omitempty"`
	CreatorUsername *string `json:"CreatorUsername,omitempty" xml:"CreatorUsername,omitempty"`
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

func (s ListApprovalsResponseBodyApprovals) String() string {
	return dara.Prettify(s)
}

func (s ListApprovalsResponseBodyApprovals) GoString() string {
	return s.String()
}

func (s *ListApprovalsResponseBodyApprovals) GetApprovalDetail() *string {
	return s.ApprovalDetail
}

func (s *ListApprovalsResponseBodyApprovals) GetApprovalId() *string {
	return s.ApprovalId
}

func (s *ListApprovalsResponseBodyApprovals) GetApprovalProgresses() []*ListApprovalsResponseBodyApprovalsApprovalProgresses {
	return s.ApprovalProgresses
}

func (s *ListApprovalsResponseBodyApprovals) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListApprovalsResponseBodyApprovals) GetCreatorDepartment() *string {
	return s.CreatorDepartment
}

func (s *ListApprovalsResponseBodyApprovals) GetCreatorDevTag() *string {
	return s.CreatorDevTag
}

func (s *ListApprovalsResponseBodyApprovals) GetCreatorUserId() *string {
	return s.CreatorUserId
}

func (s *ListApprovalsResponseBodyApprovals) GetCreatorUsername() *string {
	return s.CreatorUsername
}

func (s *ListApprovalsResponseBodyApprovals) GetEndTimestamp() *int64 {
	return s.EndTimestamp
}

func (s *ListApprovalsResponseBodyApprovals) GetPolicyType() *string {
	return s.PolicyType
}

func (s *ListApprovalsResponseBodyApprovals) GetProcessId() *string {
	return s.ProcessId
}

func (s *ListApprovalsResponseBodyApprovals) GetProcessName() *string {
	return s.ProcessName
}

func (s *ListApprovalsResponseBodyApprovals) GetReason() *string {
	return s.Reason
}

func (s *ListApprovalsResponseBodyApprovals) GetSchemaContent() *string {
	return s.SchemaContent
}

func (s *ListApprovalsResponseBodyApprovals) GetSchemaId() *string {
	return s.SchemaId
}

func (s *ListApprovalsResponseBodyApprovals) GetSchemaName() *string {
	return s.SchemaName
}

func (s *ListApprovalsResponseBodyApprovals) GetStatus() *string {
	return s.Status
}

func (s *ListApprovalsResponseBodyApprovals) SetApprovalDetail(v string) *ListApprovalsResponseBodyApprovals {
	s.ApprovalDetail = &v
	return s
}

func (s *ListApprovalsResponseBodyApprovals) SetApprovalId(v string) *ListApprovalsResponseBodyApprovals {
	s.ApprovalId = &v
	return s
}

func (s *ListApprovalsResponseBodyApprovals) SetApprovalProgresses(v []*ListApprovalsResponseBodyApprovalsApprovalProgresses) *ListApprovalsResponseBodyApprovals {
	s.ApprovalProgresses = v
	return s
}

func (s *ListApprovalsResponseBodyApprovals) SetCreateTime(v string) *ListApprovalsResponseBodyApprovals {
	s.CreateTime = &v
	return s
}

func (s *ListApprovalsResponseBodyApprovals) SetCreatorDepartment(v string) *ListApprovalsResponseBodyApprovals {
	s.CreatorDepartment = &v
	return s
}

func (s *ListApprovalsResponseBodyApprovals) SetCreatorDevTag(v string) *ListApprovalsResponseBodyApprovals {
	s.CreatorDevTag = &v
	return s
}

func (s *ListApprovalsResponseBodyApprovals) SetCreatorUserId(v string) *ListApprovalsResponseBodyApprovals {
	s.CreatorUserId = &v
	return s
}

func (s *ListApprovalsResponseBodyApprovals) SetCreatorUsername(v string) *ListApprovalsResponseBodyApprovals {
	s.CreatorUsername = &v
	return s
}

func (s *ListApprovalsResponseBodyApprovals) SetEndTimestamp(v int64) *ListApprovalsResponseBodyApprovals {
	s.EndTimestamp = &v
	return s
}

func (s *ListApprovalsResponseBodyApprovals) SetPolicyType(v string) *ListApprovalsResponseBodyApprovals {
	s.PolicyType = &v
	return s
}

func (s *ListApprovalsResponseBodyApprovals) SetProcessId(v string) *ListApprovalsResponseBodyApprovals {
	s.ProcessId = &v
	return s
}

func (s *ListApprovalsResponseBodyApprovals) SetProcessName(v string) *ListApprovalsResponseBodyApprovals {
	s.ProcessName = &v
	return s
}

func (s *ListApprovalsResponseBodyApprovals) SetReason(v string) *ListApprovalsResponseBodyApprovals {
	s.Reason = &v
	return s
}

func (s *ListApprovalsResponseBodyApprovals) SetSchemaContent(v string) *ListApprovalsResponseBodyApprovals {
	s.SchemaContent = &v
	return s
}

func (s *ListApprovalsResponseBodyApprovals) SetSchemaId(v string) *ListApprovalsResponseBodyApprovals {
	s.SchemaId = &v
	return s
}

func (s *ListApprovalsResponseBodyApprovals) SetSchemaName(v string) *ListApprovalsResponseBodyApprovals {
	s.SchemaName = &v
	return s
}

func (s *ListApprovalsResponseBodyApprovals) SetStatus(v string) *ListApprovalsResponseBodyApprovals {
	s.Status = &v
	return s
}

func (s *ListApprovalsResponseBodyApprovals) Validate() error {
	return dara.Validate(s)
}

type ListApprovalsResponseBodyApprovalsApprovalProgresses struct {
	// example:
	//
	// Approve
	Action  *string `json:"Action,omitempty" xml:"Action,omitempty"`
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// example:
	//
	// su_e8f218fb171edd167c2ad917d21f53148bdefc510ca1f3c3cc0249d3643d****
	Executor  *string                                                          `json:"Executor,omitempty" xml:"Executor,omitempty"`
	Operators []*ListApprovalsResponseBodyApprovalsApprovalProgressesOperators `json:"Operators,omitempty" xml:"Operators,omitempty" type:"Repeated"`
	// example:
	//
	// Approved
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 1736752000
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s ListApprovalsResponseBodyApprovalsApprovalProgresses) String() string {
	return dara.Prettify(s)
}

func (s ListApprovalsResponseBodyApprovalsApprovalProgresses) GoString() string {
	return s.String()
}

func (s *ListApprovalsResponseBodyApprovalsApprovalProgresses) GetAction() *string {
	return s.Action
}

func (s *ListApprovalsResponseBodyApprovalsApprovalProgresses) GetComment() *string {
	return s.Comment
}

func (s *ListApprovalsResponseBodyApprovalsApprovalProgresses) GetExecutor() *string {
	return s.Executor
}

func (s *ListApprovalsResponseBodyApprovalsApprovalProgresses) GetOperators() []*ListApprovalsResponseBodyApprovalsApprovalProgressesOperators {
	return s.Operators
}

func (s *ListApprovalsResponseBodyApprovalsApprovalProgresses) GetStatus() *string {
	return s.Status
}

func (s *ListApprovalsResponseBodyApprovalsApprovalProgresses) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *ListApprovalsResponseBodyApprovalsApprovalProgresses) SetAction(v string) *ListApprovalsResponseBodyApprovalsApprovalProgresses {
	s.Action = &v
	return s
}

func (s *ListApprovalsResponseBodyApprovalsApprovalProgresses) SetComment(v string) *ListApprovalsResponseBodyApprovalsApprovalProgresses {
	s.Comment = &v
	return s
}

func (s *ListApprovalsResponseBodyApprovalsApprovalProgresses) SetExecutor(v string) *ListApprovalsResponseBodyApprovalsApprovalProgresses {
	s.Executor = &v
	return s
}

func (s *ListApprovalsResponseBodyApprovalsApprovalProgresses) SetOperators(v []*ListApprovalsResponseBodyApprovalsApprovalProgressesOperators) *ListApprovalsResponseBodyApprovalsApprovalProgresses {
	s.Operators = v
	return s
}

func (s *ListApprovalsResponseBodyApprovalsApprovalProgresses) SetStatus(v string) *ListApprovalsResponseBodyApprovalsApprovalProgresses {
	s.Status = &v
	return s
}

func (s *ListApprovalsResponseBodyApprovalsApprovalProgresses) SetTimestamp(v int64) *ListApprovalsResponseBodyApprovalsApprovalProgresses {
	s.Timestamp = &v
	return s
}

func (s *ListApprovalsResponseBodyApprovalsApprovalProgresses) Validate() error {
	return dara.Validate(s)
}

type ListApprovalsResponseBodyApprovalsApprovalProgressesOperators struct {
	// example:
	//
	// su_e8f218fb171edd167c2ad917d21f53148bdefc510ca1f3c3cc0249d3643d****
	SaseUserId *string `json:"SaseUserId,omitempty" xml:"SaseUserId,omitempty"`
	Username   *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s ListApprovalsResponseBodyApprovalsApprovalProgressesOperators) String() string {
	return dara.Prettify(s)
}

func (s ListApprovalsResponseBodyApprovalsApprovalProgressesOperators) GoString() string {
	return s.String()
}

func (s *ListApprovalsResponseBodyApprovalsApprovalProgressesOperators) GetSaseUserId() *string {
	return s.SaseUserId
}

func (s *ListApprovalsResponseBodyApprovalsApprovalProgressesOperators) GetUsername() *string {
	return s.Username
}

func (s *ListApprovalsResponseBodyApprovalsApprovalProgressesOperators) SetSaseUserId(v string) *ListApprovalsResponseBodyApprovalsApprovalProgressesOperators {
	s.SaseUserId = &v
	return s
}

func (s *ListApprovalsResponseBodyApprovalsApprovalProgressesOperators) SetUsername(v string) *ListApprovalsResponseBodyApprovalsApprovalProgressesOperators {
	s.Username = &v
	return s
}

func (s *ListApprovalsResponseBodyApprovalsApprovalProgressesOperators) Validate() error {
	return dara.Validate(s)
}
