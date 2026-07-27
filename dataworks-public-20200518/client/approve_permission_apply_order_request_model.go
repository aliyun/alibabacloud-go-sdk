// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApprovePermissionApplyOrderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApproveAction(v int32) *ApprovePermissionApplyOrderRequest
	GetApproveAction() *int32
	SetApproveComment(v string) *ApprovePermissionApplyOrderRequest
	GetApproveComment() *string
	SetFlowId(v string) *ApprovePermissionApplyOrderRequest
	GetFlowId() *string
}

type ApprovePermissionApplyOrderRequest struct {
	// The approval action to perform. Valid values:
	//
	// - 1: Approve.
	//
	// - 2: Reject.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ApproveAction *int32 `json:"ApproveAction,omitempty" xml:"ApproveAction,omitempty"`
	// The remarks for the approval.
	//
	// This parameter is required.
	//
	// example:
	//
	// agree
	ApproveComment *string `json:"ApproveComment,omitempty" xml:"ApproveComment,omitempty"`
	// The ID of the permission request order to approve. You can call the [ListPermissionApplyOrders](https://help.aliyun.com/document_detail/211008.html) operation to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 48f36729-05f9-4a40-9286-933fd940f30a
	FlowId *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
}

func (s ApprovePermissionApplyOrderRequest) String() string {
	return dara.Prettify(s)
}

func (s ApprovePermissionApplyOrderRequest) GoString() string {
	return s.String()
}

func (s *ApprovePermissionApplyOrderRequest) GetApproveAction() *int32 {
	return s.ApproveAction
}

func (s *ApprovePermissionApplyOrderRequest) GetApproveComment() *string {
	return s.ApproveComment
}

func (s *ApprovePermissionApplyOrderRequest) GetFlowId() *string {
	return s.FlowId
}

func (s *ApprovePermissionApplyOrderRequest) SetApproveAction(v int32) *ApprovePermissionApplyOrderRequest {
	s.ApproveAction = &v
	return s
}

func (s *ApprovePermissionApplyOrderRequest) SetApproveComment(v string) *ApprovePermissionApplyOrderRequest {
	s.ApproveComment = &v
	return s
}

func (s *ApprovePermissionApplyOrderRequest) SetFlowId(v string) *ApprovePermissionApplyOrderRequest {
	s.FlowId = &v
	return s
}

func (s *ApprovePermissionApplyOrderRequest) Validate() error {
	return dara.Validate(s)
}
