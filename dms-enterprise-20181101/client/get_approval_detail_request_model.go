// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApprovalDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTid(v int64) *GetApprovalDetailRequest
	GetTid() *int64
	SetWorkflowInstanceId(v int64) *GetApprovalDetailRequest
	GetWorkflowInstanceId() *int64
}

type GetApprovalDetailRequest struct {
	// The ID of the tenant. You can call the [GetUserActiveTenant](https://help.aliyun.com/document_detail/198073.html) operation to obtain the tenant ID.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
	// The ID of the approval process. You can call the [GetOrderBaseInfo](https://help.aliyun.com/document_detail/144642.html) operation to obtain the ID of the approval process.
	//
	// This parameter is required.
	//
	// example:
	//
	// 184****
	WorkflowInstanceId *int64 `json:"WorkflowInstanceId,omitempty" xml:"WorkflowInstanceId,omitempty"`
}

func (s GetApprovalDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetApprovalDetailRequest) GoString() string {
	return s.String()
}

func (s *GetApprovalDetailRequest) GetTid() *int64 {
	return s.Tid
}

func (s *GetApprovalDetailRequest) GetWorkflowInstanceId() *int64 {
	return s.WorkflowInstanceId
}

func (s *GetApprovalDetailRequest) SetTid(v int64) *GetApprovalDetailRequest {
	s.Tid = &v
	return s
}

func (s *GetApprovalDetailRequest) SetWorkflowInstanceId(v int64) *GetApprovalDetailRequest {
	s.WorkflowInstanceId = &v
	return s
}

func (s *GetApprovalDetailRequest) Validate() error {
	return dara.Validate(s)
}
