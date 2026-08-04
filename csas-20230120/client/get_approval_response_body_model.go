// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApprovalResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApproval(v []*GetApprovalResponseBodyApproval) *GetApprovalResponseBody
	GetApproval() []*GetApprovalResponseBodyApproval
	SetRequestId(v string) *GetApprovalResponseBody
	GetRequestId() *string
}

type GetApprovalResponseBody struct {
	// The approval instance.
	Approval []*GetApprovalResponseBodyApproval `json:"Approval,omitempty" xml:"Approval,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 7E9D7ACD-53D5-56EF-A913-79D148D06299
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetApprovalResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetApprovalResponseBody) GoString() string {
	return s.String()
}

func (s *GetApprovalResponseBody) GetApproval() []*GetApprovalResponseBodyApproval {
	return s.Approval
}

func (s *GetApprovalResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetApprovalResponseBody) SetApproval(v []*GetApprovalResponseBodyApproval) *GetApprovalResponseBody {
	s.Approval = v
	return s
}

func (s *GetApprovalResponseBody) SetRequestId(v string) *GetApprovalResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetApprovalResponseBody) Validate() error {
	if s.Approval != nil {
		for _, item := range s.Approval {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetApprovalResponseBodyApproval struct {
	// The details of the approval instance.
	//
	// example:
	//
	// {"initiatorName":"王先生","initiatorDept":"测试部","devType":"windows","deviceType":"usbStorage","deviceId":"FC216E9E3****","approvalEndTimestamp":1736524799,"approvalReason":"这是一个测试"}
	ApprovalDetail *string `json:"ApprovalDetail,omitempty" xml:"ApprovalDetail,omitempty"`
	// The ID of the approval instance.
	//
	// example:
	//
	// approval-3564b140642f****
	ApprovalId *string `json:"ApprovalId,omitempty" xml:"ApprovalId,omitempty"`
	// A list of progress nodes for the approval instance.
	ApprovalProgresses []*GetApprovalResponseBodyApprovalApprovalProgresses `json:"ApprovalProgresses,omitempty" xml:"ApprovalProgresses,omitempty" type:"Repeated"`
	// The time when the approval instance was created.
	//
	// example:
	//
	// 2022-11-15 22:11:55
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The department of the user who created the approval instance.
	//
	// example:
	//
	// 测试部
	CreatorDepartment *string `json:"CreatorDepartment,omitempty" xml:"CreatorDepartment,omitempty"`
	// The ID of the client device from which the approval instance was created.
	//
	// example:
	//
	// 36efa42d-2c32-c4dc-e3fc-8541e33a****
	CreatorDevTag *string `json:"CreatorDevTag,omitempty" xml:"CreatorDevTag,omitempty"`
	// The ID of the user who created the approval instance.
	//
	// example:
	//
	// su_e8f218fb171edd167c2ad917d21f53148bdefc510ca1f3c3cc0249d3643d****
	CreatorUserId *string `json:"CreatorUserId,omitempty" xml:"CreatorUserId,omitempty"`
	// The username of the user who created the approval instance.
	//
	// example:
	//
	// 王先生
	CreatorUsername *string `json:"CreatorUsername,omitempty" xml:"CreatorUsername,omitempty"`
	// The expiration time of the approval instance. This is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1757952000
	EndTimestamp *int64 `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	// The type of the associated policy. Valid values:
	//
	// - **DomainBlacklist**: Domain blacklist
	//
	// - **DomainWhitelist**: Domain whitelist
	//
	// - **SoftwareBlock**: Software block
	//
	// - **AppUninstall**: Client uninstall
	//
	// - **DlpSend**: Outbound file transfer
	//
	// - **PeripheralBlock**: Peripheral control
	//
	// example:
	//
	// DlpSend
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	// The ID of the associated approval flow.
	//
	// example:
	//
	// approval-process-fcc351b8a95b****
	ProcessId *string `json:"ProcessId,omitempty" xml:"ProcessId,omitempty"`
	// The name of the associated approval flow.
	//
	// example:
	//
	// 测试
	ProcessName *string `json:"ProcessName,omitempty" xml:"ProcessName,omitempty"`
	// The reason for creating the approval instance.
	//
	// example:
	//
	// 这是一个测试
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The content of the associated template.
	//
	// example:
	//
	// {"form": {"labelCol": 6,"wrapperCol": 12}}
	SchemaContent *string `json:"SchemaContent,omitempty" xml:"SchemaContent,omitempty"`
	// The ID of the associated template.
	//
	// example:
	//
	// approval-schema-090134f1ebff****
	SchemaId *string `json:"SchemaId,omitempty" xml:"SchemaId,omitempty"`
	// The name of the associated template.
	//
	// example:
	//
	// 测试
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// The status of the approval instance. Valid values:
	//
	// - **Pending**: Pending
	//
	// - **Approved**: Approved
	//
	// - **Rejected**: Rejected
	//
	// - **Revoked**: Revoked
	//
	// - **Expired**: Expired
	//
	// example:
	//
	// Pending
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetApprovalResponseBodyApproval) String() string {
	return dara.Prettify(s)
}

func (s GetApprovalResponseBodyApproval) GoString() string {
	return s.String()
}

func (s *GetApprovalResponseBodyApproval) GetApprovalDetail() *string {
	return s.ApprovalDetail
}

func (s *GetApprovalResponseBodyApproval) GetApprovalId() *string {
	return s.ApprovalId
}

func (s *GetApprovalResponseBodyApproval) GetApprovalProgresses() []*GetApprovalResponseBodyApprovalApprovalProgresses {
	return s.ApprovalProgresses
}

func (s *GetApprovalResponseBodyApproval) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetApprovalResponseBodyApproval) GetCreatorDepartment() *string {
	return s.CreatorDepartment
}

func (s *GetApprovalResponseBodyApproval) GetCreatorDevTag() *string {
	return s.CreatorDevTag
}

func (s *GetApprovalResponseBodyApproval) GetCreatorUserId() *string {
	return s.CreatorUserId
}

func (s *GetApprovalResponseBodyApproval) GetCreatorUsername() *string {
	return s.CreatorUsername
}

func (s *GetApprovalResponseBodyApproval) GetEndTimestamp() *int64 {
	return s.EndTimestamp
}

func (s *GetApprovalResponseBodyApproval) GetPolicyType() *string {
	return s.PolicyType
}

func (s *GetApprovalResponseBodyApproval) GetProcessId() *string {
	return s.ProcessId
}

func (s *GetApprovalResponseBodyApproval) GetProcessName() *string {
	return s.ProcessName
}

func (s *GetApprovalResponseBodyApproval) GetReason() *string {
	return s.Reason
}

func (s *GetApprovalResponseBodyApproval) GetSchemaContent() *string {
	return s.SchemaContent
}

func (s *GetApprovalResponseBodyApproval) GetSchemaId() *string {
	return s.SchemaId
}

func (s *GetApprovalResponseBodyApproval) GetSchemaName() *string {
	return s.SchemaName
}

func (s *GetApprovalResponseBodyApproval) GetStatus() *string {
	return s.Status
}

func (s *GetApprovalResponseBodyApproval) SetApprovalDetail(v string) *GetApprovalResponseBodyApproval {
	s.ApprovalDetail = &v
	return s
}

func (s *GetApprovalResponseBodyApproval) SetApprovalId(v string) *GetApprovalResponseBodyApproval {
	s.ApprovalId = &v
	return s
}

func (s *GetApprovalResponseBodyApproval) SetApprovalProgresses(v []*GetApprovalResponseBodyApprovalApprovalProgresses) *GetApprovalResponseBodyApproval {
	s.ApprovalProgresses = v
	return s
}

func (s *GetApprovalResponseBodyApproval) SetCreateTime(v string) *GetApprovalResponseBodyApproval {
	s.CreateTime = &v
	return s
}

func (s *GetApprovalResponseBodyApproval) SetCreatorDepartment(v string) *GetApprovalResponseBodyApproval {
	s.CreatorDepartment = &v
	return s
}

func (s *GetApprovalResponseBodyApproval) SetCreatorDevTag(v string) *GetApprovalResponseBodyApproval {
	s.CreatorDevTag = &v
	return s
}

func (s *GetApprovalResponseBodyApproval) SetCreatorUserId(v string) *GetApprovalResponseBodyApproval {
	s.CreatorUserId = &v
	return s
}

func (s *GetApprovalResponseBodyApproval) SetCreatorUsername(v string) *GetApprovalResponseBodyApproval {
	s.CreatorUsername = &v
	return s
}

func (s *GetApprovalResponseBodyApproval) SetEndTimestamp(v int64) *GetApprovalResponseBodyApproval {
	s.EndTimestamp = &v
	return s
}

func (s *GetApprovalResponseBodyApproval) SetPolicyType(v string) *GetApprovalResponseBodyApproval {
	s.PolicyType = &v
	return s
}

func (s *GetApprovalResponseBodyApproval) SetProcessId(v string) *GetApprovalResponseBodyApproval {
	s.ProcessId = &v
	return s
}

func (s *GetApprovalResponseBodyApproval) SetProcessName(v string) *GetApprovalResponseBodyApproval {
	s.ProcessName = &v
	return s
}

func (s *GetApprovalResponseBodyApproval) SetReason(v string) *GetApprovalResponseBodyApproval {
	s.Reason = &v
	return s
}

func (s *GetApprovalResponseBodyApproval) SetSchemaContent(v string) *GetApprovalResponseBodyApproval {
	s.SchemaContent = &v
	return s
}

func (s *GetApprovalResponseBodyApproval) SetSchemaId(v string) *GetApprovalResponseBodyApproval {
	s.SchemaId = &v
	return s
}

func (s *GetApprovalResponseBodyApproval) SetSchemaName(v string) *GetApprovalResponseBodyApproval {
	s.SchemaName = &v
	return s
}

func (s *GetApprovalResponseBodyApproval) SetStatus(v string) *GetApprovalResponseBodyApproval {
	s.Status = &v
	return s
}

func (s *GetApprovalResponseBodyApproval) Validate() error {
	if s.ApprovalProgresses != nil {
		for _, item := range s.ApprovalProgresses {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetApprovalResponseBodyApprovalApprovalProgresses struct {
	// The action performed on the progress node. Valid values:
	//
	// - **Approve**: Approve
	//
	// - **Reject**: Reject
	//
	// - **Revoke**: Revoke
	//
	// - **Comment**: Comment
	//
	// example:
	//
	// Approve
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The comment for the action on the progress node.
	//
	// example:
	//
	// 审核通过
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The ID of the executor for the progress node.
	//
	// example:
	//
	// su_e8f218fb171edd167c2ad917d21f53148bdefc510ca1f3c3cc0249d3643d****
	Executor *string `json:"Executor,omitempty" xml:"Executor,omitempty"`
	// A list of operators for the progress node.
	Operators []*GetApprovalResponseBodyApprovalApprovalProgressesOperators `json:"Operators,omitempty" xml:"Operators,omitempty" type:"Repeated"`
	// The status of the progress node. Valid values:
	//
	// - **Pending**: Pending
	//
	// - **Approved**: Approved
	//
	// - **Rejected**: Rejected
	//
	// - **Revoked**: Revoked
	//
	// example:
	//
	// Approved
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when the action was performed on the progress node. This is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1736752000
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s GetApprovalResponseBodyApprovalApprovalProgresses) String() string {
	return dara.Prettify(s)
}

func (s GetApprovalResponseBodyApprovalApprovalProgresses) GoString() string {
	return s.String()
}

func (s *GetApprovalResponseBodyApprovalApprovalProgresses) GetAction() *string {
	return s.Action
}

func (s *GetApprovalResponseBodyApprovalApprovalProgresses) GetComment() *string {
	return s.Comment
}

func (s *GetApprovalResponseBodyApprovalApprovalProgresses) GetExecutor() *string {
	return s.Executor
}

func (s *GetApprovalResponseBodyApprovalApprovalProgresses) GetOperators() []*GetApprovalResponseBodyApprovalApprovalProgressesOperators {
	return s.Operators
}

func (s *GetApprovalResponseBodyApprovalApprovalProgresses) GetStatus() *string {
	return s.Status
}

func (s *GetApprovalResponseBodyApprovalApprovalProgresses) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *GetApprovalResponseBodyApprovalApprovalProgresses) SetAction(v string) *GetApprovalResponseBodyApprovalApprovalProgresses {
	s.Action = &v
	return s
}

func (s *GetApprovalResponseBodyApprovalApprovalProgresses) SetComment(v string) *GetApprovalResponseBodyApprovalApprovalProgresses {
	s.Comment = &v
	return s
}

func (s *GetApprovalResponseBodyApprovalApprovalProgresses) SetExecutor(v string) *GetApprovalResponseBodyApprovalApprovalProgresses {
	s.Executor = &v
	return s
}

func (s *GetApprovalResponseBodyApprovalApprovalProgresses) SetOperators(v []*GetApprovalResponseBodyApprovalApprovalProgressesOperators) *GetApprovalResponseBodyApprovalApprovalProgresses {
	s.Operators = v
	return s
}

func (s *GetApprovalResponseBodyApprovalApprovalProgresses) SetStatus(v string) *GetApprovalResponseBodyApprovalApprovalProgresses {
	s.Status = &v
	return s
}

func (s *GetApprovalResponseBodyApprovalApprovalProgresses) SetTimestamp(v int64) *GetApprovalResponseBodyApprovalApprovalProgresses {
	s.Timestamp = &v
	return s
}

func (s *GetApprovalResponseBodyApprovalApprovalProgresses) Validate() error {
	if s.Operators != nil {
		for _, item := range s.Operators {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetApprovalResponseBodyApprovalApprovalProgressesOperators struct {
	// The ID of the operator for the progress node.
	//
	// example:
	//
	// su_e8f218fb171edd167c2ad917d21f53148bdefc510ca1f3c3cc0249d3643d****
	SaseUserId *string `json:"SaseUserId,omitempty" xml:"SaseUserId,omitempty"`
	// The username of the operator for the progress node.
	//
	// example:
	//
	// 王先生
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s GetApprovalResponseBodyApprovalApprovalProgressesOperators) String() string {
	return dara.Prettify(s)
}

func (s GetApprovalResponseBodyApprovalApprovalProgressesOperators) GoString() string {
	return s.String()
}

func (s *GetApprovalResponseBodyApprovalApprovalProgressesOperators) GetSaseUserId() *string {
	return s.SaseUserId
}

func (s *GetApprovalResponseBodyApprovalApprovalProgressesOperators) GetUsername() *string {
	return s.Username
}

func (s *GetApprovalResponseBodyApprovalApprovalProgressesOperators) SetSaseUserId(v string) *GetApprovalResponseBodyApprovalApprovalProgressesOperators {
	s.SaseUserId = &v
	return s
}

func (s *GetApprovalResponseBodyApprovalApprovalProgressesOperators) SetUsername(v string) *GetApprovalResponseBodyApprovalApprovalProgressesOperators {
	s.Username = &v
	return s
}

func (s *GetApprovalResponseBodyApprovalApprovalProgressesOperators) Validate() error {
	return dara.Validate(s)
}
