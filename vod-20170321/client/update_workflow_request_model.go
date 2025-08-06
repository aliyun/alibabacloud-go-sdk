// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkflowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActionList(v string) *UpdateWorkflowRequest
	GetActionList() *string
	SetCallbackConfig(v string) *UpdateWorkflowRequest
	GetCallbackConfig() *string
	SetDescription(v string) *UpdateWorkflowRequest
	GetDescription() *string
	SetName(v string) *UpdateWorkflowRequest
	GetName() *string
	SetOwnerId(v int64) *UpdateWorkflowRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *UpdateWorkflowRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdateWorkflowRequest
	GetResourceOwnerId() *int64
	SetState(v string) *UpdateWorkflowRequest
	GetState() *string
	SetWorkflowId(v string) *UpdateWorkflowRequest
	GetWorkflowId() *string
}

type UpdateWorkflowRequest struct {
	ActionList           *string `json:"ActionList,omitempty" xml:"ActionList,omitempty"`
	CallbackConfig       *string `json:"CallbackConfig,omitempty" xml:"CallbackConfig,omitempty"`
	Description          *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name                 *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	State                *string `json:"State,omitempty" xml:"State,omitempty"`
	// This parameter is required.
	WorkflowId *string `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
}

func (s UpdateWorkflowRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkflowRequest) GoString() string {
	return s.String()
}

func (s *UpdateWorkflowRequest) GetActionList() *string {
	return s.ActionList
}

func (s *UpdateWorkflowRequest) GetCallbackConfig() *string {
	return s.CallbackConfig
}

func (s *UpdateWorkflowRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateWorkflowRequest) GetName() *string {
	return s.Name
}

func (s *UpdateWorkflowRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateWorkflowRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateWorkflowRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdateWorkflowRequest) GetState() *string {
	return s.State
}

func (s *UpdateWorkflowRequest) GetWorkflowId() *string {
	return s.WorkflowId
}

func (s *UpdateWorkflowRequest) SetActionList(v string) *UpdateWorkflowRequest {
	s.ActionList = &v
	return s
}

func (s *UpdateWorkflowRequest) SetCallbackConfig(v string) *UpdateWorkflowRequest {
	s.CallbackConfig = &v
	return s
}

func (s *UpdateWorkflowRequest) SetDescription(v string) *UpdateWorkflowRequest {
	s.Description = &v
	return s
}

func (s *UpdateWorkflowRequest) SetName(v string) *UpdateWorkflowRequest {
	s.Name = &v
	return s
}

func (s *UpdateWorkflowRequest) SetOwnerId(v int64) *UpdateWorkflowRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateWorkflowRequest) SetResourceOwnerAccount(v string) *UpdateWorkflowRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateWorkflowRequest) SetResourceOwnerId(v int64) *UpdateWorkflowRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateWorkflowRequest) SetState(v string) *UpdateWorkflowRequest {
	s.State = &v
	return s
}

func (s *UpdateWorkflowRequest) SetWorkflowId(v string) *UpdateWorkflowRequest {
	s.WorkflowId = &v
	return s
}

func (s *UpdateWorkflowRequest) Validate() error {
	return dara.Validate(s)
}
