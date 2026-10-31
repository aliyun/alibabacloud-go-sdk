// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApprovalTaskDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFlowId(v int64) *GetApprovalTaskDetailRequest
	GetFlowId() *int64
	SetOpTenantId(v int64) *GetApprovalTaskDetailRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *GetApprovalTaskDetailRequest
	GetOpUserId() *string
}

type GetApprovalTaskDetailRequest struct {
	// The approval flow ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	FlowId *int64 `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The ID of the operator.
	//
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
}

func (s GetApprovalTaskDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetApprovalTaskDetailRequest) GoString() string {
	return s.String()
}

func (s *GetApprovalTaskDetailRequest) GetFlowId() *int64 {
	return s.FlowId
}

func (s *GetApprovalTaskDetailRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetApprovalTaskDetailRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *GetApprovalTaskDetailRequest) SetFlowId(v int64) *GetApprovalTaskDetailRequest {
	s.FlowId = &v
	return s
}

func (s *GetApprovalTaskDetailRequest) SetOpTenantId(v int64) *GetApprovalTaskDetailRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetApprovalTaskDetailRequest) SetOpUserId(v string) *GetApprovalTaskDetailRequest {
	s.OpUserId = &v
	return s
}

func (s *GetApprovalTaskDetailRequest) Validate() error {
	return dara.Validate(s)
}
