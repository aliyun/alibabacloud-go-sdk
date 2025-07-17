// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApproveOrderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApprovalNodeId(v int64) *ApproveOrderRequest
	GetApprovalNodeId() *int64
	SetApprovalNodePos(v string) *ApproveOrderRequest
	GetApprovalNodePos() *string
	SetApprovalType(v string) *ApproveOrderRequest
	GetApprovalType() *string
	SetComment(v string) *ApproveOrderRequest
	GetComment() *string
	SetNewApprover(v int64) *ApproveOrderRequest
	GetNewApprover() *int64
	SetNewApproverList(v string) *ApproveOrderRequest
	GetNewApproverList() *string
	SetOldApprover(v int64) *ApproveOrderRequest
	GetOldApprover() *int64
	SetRealLoginUserUid(v string) *ApproveOrderRequest
	GetRealLoginUserUid() *string
	SetTid(v int64) *ApproveOrderRequest
	GetTid() *int64
	SetWorkflowInstanceId(v int64) *ApproveOrderRequest
	GetWorkflowInstanceId() *int64
}

type ApproveOrderRequest struct {
	// If ApprovalType is set to ADD_APPROVAL_NODE, you need to specify this parameter. The ID of the user that is added as the new approval node. This node must be a user-defined approval node. You can call the ListUserDefineWorkFlowNodes operation to obtain the value of this parameter.
	//
	// example:
	//
	// 1
	ApprovalNodeId *int64 `json:"ApprovalNodeId,omitempty" xml:"ApprovalNodeId,omitempty"`
	// The position of the new approval node. You must specify this parameter if ApprovalType is set to ADD_APPROVAL_NODE. Valid values:
	//
	// 	- **PRE_ADD_APPROVAL_NODE**: before the current approval node.
	//
	// 	- **POST_ADD_APPROVAL_NODE**: after the current approval node.
	//
	// example:
	//
	// POST_ADD_APPROVAL_NODE
	ApprovalNodePos *string `json:"ApprovalNodePos,omitempty" xml:"ApprovalNodePos,omitempty"`
	// The action that you want to perform on the ticket. Valid values:
	//
	// 	- **AGREE**
	//
	// 	- **CANCEL**
	//
	// 	- **REJECT**
	//
	// 	- **TRANSFER**
	//
	// 	- **ADD_APPROVAL_NODE**
	//
	// This parameter is required.
	//
	// example:
	//
	// agree
	ApprovalType *string `json:"ApprovalType,omitempty" xml:"ApprovalType,omitempty"`
	// The description of the ticket.
	//
	// example:
	//
	// test
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The ID of the user to which the ticket is transferred. If ApprovalType is set to TRANSFER, you need to specify this parameter.
	//
	// example:
	//
	// 12***
	NewApprover *int64 `json:"NewApprover,omitempty" xml:"NewApprover,omitempty"`
	// >  You can specify this parameter if ApprovalType is set to TRANSFER. You need to only specify one of NewApproverList and NewApprover.
	//
	// The IDs of the users to whom the ticket is transferred. Separate multiple IDs with commas (,).
	//
	// example:
	//
	// 154***,155***,156***
	NewApproverList *string `json:"NewApproverList,omitempty" xml:"NewApproverList,omitempty"`
	// The ID of the user that transfers the ticket to another user. The default value is the ID of the current user. If the current user is an administrator or a database administrator (DBA), the user can change the value of this parameter to the ID of another user.
	//
	// example:
	//
	// 23***
	OldApprover *int64 `json:"OldApprover,omitempty" xml:"OldApprover,omitempty"`
	// The UID of the Alibaba Cloud account that actually calls the API.
	//
	// example:
	//
	// 21400447956867****
	RealLoginUserUid *string `json:"RealLoginUserUid,omitempty" xml:"RealLoginUserUid,omitempty"`
	// The ID of the tenant. You can call the [GetUserActiveTenant](https://help.aliyun.com/document_detail/198073.html) operation to obtain the tenant ID.
	//
	// example:
	//
	// -1
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
	// The ID of the approval process. You can call the [GetOrderBaseInfo](https://help.aliyun.com/document_detail/144642.html) operation to obtain the ID of the approval process.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234
	WorkflowInstanceId *int64 `json:"WorkflowInstanceId,omitempty" xml:"WorkflowInstanceId,omitempty"`
}

func (s ApproveOrderRequest) String() string {
	return dara.Prettify(s)
}

func (s ApproveOrderRequest) GoString() string {
	return s.String()
}

func (s *ApproveOrderRequest) GetApprovalNodeId() *int64 {
	return s.ApprovalNodeId
}

func (s *ApproveOrderRequest) GetApprovalNodePos() *string {
	return s.ApprovalNodePos
}

func (s *ApproveOrderRequest) GetApprovalType() *string {
	return s.ApprovalType
}

func (s *ApproveOrderRequest) GetComment() *string {
	return s.Comment
}

func (s *ApproveOrderRequest) GetNewApprover() *int64 {
	return s.NewApprover
}

func (s *ApproveOrderRequest) GetNewApproverList() *string {
	return s.NewApproverList
}

func (s *ApproveOrderRequest) GetOldApprover() *int64 {
	return s.OldApprover
}

func (s *ApproveOrderRequest) GetRealLoginUserUid() *string {
	return s.RealLoginUserUid
}

func (s *ApproveOrderRequest) GetTid() *int64 {
	return s.Tid
}

func (s *ApproveOrderRequest) GetWorkflowInstanceId() *int64 {
	return s.WorkflowInstanceId
}

func (s *ApproveOrderRequest) SetApprovalNodeId(v int64) *ApproveOrderRequest {
	s.ApprovalNodeId = &v
	return s
}

func (s *ApproveOrderRequest) SetApprovalNodePos(v string) *ApproveOrderRequest {
	s.ApprovalNodePos = &v
	return s
}

func (s *ApproveOrderRequest) SetApprovalType(v string) *ApproveOrderRequest {
	s.ApprovalType = &v
	return s
}

func (s *ApproveOrderRequest) SetComment(v string) *ApproveOrderRequest {
	s.Comment = &v
	return s
}

func (s *ApproveOrderRequest) SetNewApprover(v int64) *ApproveOrderRequest {
	s.NewApprover = &v
	return s
}

func (s *ApproveOrderRequest) SetNewApproverList(v string) *ApproveOrderRequest {
	s.NewApproverList = &v
	return s
}

func (s *ApproveOrderRequest) SetOldApprover(v int64) *ApproveOrderRequest {
	s.OldApprover = &v
	return s
}

func (s *ApproveOrderRequest) SetRealLoginUserUid(v string) *ApproveOrderRequest {
	s.RealLoginUserUid = &v
	return s
}

func (s *ApproveOrderRequest) SetTid(v int64) *ApproveOrderRequest {
	s.Tid = &v
	return s
}

func (s *ApproveOrderRequest) SetWorkflowInstanceId(v int64) *ApproveOrderRequest {
	s.WorkflowInstanceId = &v
	return s
}

func (s *ApproveOrderRequest) Validate() error {
	return dara.Validate(s)
}
