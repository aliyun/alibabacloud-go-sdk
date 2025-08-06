// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWorkflowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *DeleteWorkflowRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteWorkflowRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteWorkflowRequest
	GetResourceOwnerId() *int64
	SetWorkflowId(v string) *DeleteWorkflowRequest
	GetWorkflowId() *string
}

type DeleteWorkflowRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	WorkflowId *string `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
}

func (s DeleteWorkflowRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteWorkflowRequest) GoString() string {
	return s.String()
}

func (s *DeleteWorkflowRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteWorkflowRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteWorkflowRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteWorkflowRequest) GetWorkflowId() *string {
	return s.WorkflowId
}

func (s *DeleteWorkflowRequest) SetOwnerId(v int64) *DeleteWorkflowRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteWorkflowRequest) SetResourceOwnerAccount(v string) *DeleteWorkflowRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteWorkflowRequest) SetResourceOwnerId(v int64) *DeleteWorkflowRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteWorkflowRequest) SetWorkflowId(v string) *DeleteWorkflowRequest {
	s.WorkflowId = &v
	return s
}

func (s *DeleteWorkflowRequest) Validate() error {
	return dara.Validate(s)
}
