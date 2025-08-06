// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkflowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *GetWorkflowRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *GetWorkflowRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetWorkflowRequest
	GetResourceOwnerId() *int64
	SetWorkflowId(v string) *GetWorkflowRequest
	GetWorkflowId() *string
}

type GetWorkflowRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	WorkflowId *string `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
}

func (s GetWorkflowRequest) String() string {
	return dara.Prettify(s)
}

func (s GetWorkflowRequest) GoString() string {
	return s.String()
}

func (s *GetWorkflowRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetWorkflowRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetWorkflowRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetWorkflowRequest) GetWorkflowId() *string {
	return s.WorkflowId
}

func (s *GetWorkflowRequest) SetOwnerId(v int64) *GetWorkflowRequest {
	s.OwnerId = &v
	return s
}

func (s *GetWorkflowRequest) SetResourceOwnerAccount(v string) *GetWorkflowRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetWorkflowRequest) SetResourceOwnerId(v int64) *GetWorkflowRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetWorkflowRequest) SetWorkflowId(v string) *GetWorkflowRequest {
	s.WorkflowId = &v
	return s
}

func (s *GetWorkflowRequest) Validate() error {
	return dara.Validate(s)
}
