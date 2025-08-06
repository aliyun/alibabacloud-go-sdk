// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddWorkflowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActionList(v string) *AddWorkflowRequest
	GetActionList() *string
	SetAppId(v string) *AddWorkflowRequest
	GetAppId() *string
	SetBizVersion(v string) *AddWorkflowRequest
	GetBizVersion() *string
	SetCallbackConfig(v string) *AddWorkflowRequest
	GetCallbackConfig() *string
	SetDescription(v string) *AddWorkflowRequest
	GetDescription() *string
	SetName(v string) *AddWorkflowRequest
	GetName() *string
	SetOwnerId(v int64) *AddWorkflowRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *AddWorkflowRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AddWorkflowRequest
	GetResourceOwnerId() *int64
}

type AddWorkflowRequest struct {
	// This parameter is required.
	ActionList     *string `json:"ActionList,omitempty" xml:"ActionList,omitempty"`
	AppId          *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	BizVersion     *string `json:"BizVersion,omitempty" xml:"BizVersion,omitempty"`
	CallbackConfig *string `json:"CallbackConfig,omitempty" xml:"CallbackConfig,omitempty"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	Name                 *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s AddWorkflowRequest) String() string {
	return dara.Prettify(s)
}

func (s AddWorkflowRequest) GoString() string {
	return s.String()
}

func (s *AddWorkflowRequest) GetActionList() *string {
	return s.ActionList
}

func (s *AddWorkflowRequest) GetAppId() *string {
	return s.AppId
}

func (s *AddWorkflowRequest) GetBizVersion() *string {
	return s.BizVersion
}

func (s *AddWorkflowRequest) GetCallbackConfig() *string {
	return s.CallbackConfig
}

func (s *AddWorkflowRequest) GetDescription() *string {
	return s.Description
}

func (s *AddWorkflowRequest) GetName() *string {
	return s.Name
}

func (s *AddWorkflowRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddWorkflowRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AddWorkflowRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AddWorkflowRequest) SetActionList(v string) *AddWorkflowRequest {
	s.ActionList = &v
	return s
}

func (s *AddWorkflowRequest) SetAppId(v string) *AddWorkflowRequest {
	s.AppId = &v
	return s
}

func (s *AddWorkflowRequest) SetBizVersion(v string) *AddWorkflowRequest {
	s.BizVersion = &v
	return s
}

func (s *AddWorkflowRequest) SetCallbackConfig(v string) *AddWorkflowRequest {
	s.CallbackConfig = &v
	return s
}

func (s *AddWorkflowRequest) SetDescription(v string) *AddWorkflowRequest {
	s.Description = &v
	return s
}

func (s *AddWorkflowRequest) SetName(v string) *AddWorkflowRequest {
	s.Name = &v
	return s
}

func (s *AddWorkflowRequest) SetOwnerId(v int64) *AddWorkflowRequest {
	s.OwnerId = &v
	return s
}

func (s *AddWorkflowRequest) SetResourceOwnerAccount(v string) *AddWorkflowRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AddWorkflowRequest) SetResourceOwnerId(v int64) *AddWorkflowRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AddWorkflowRequest) Validate() error {
	return dara.Validate(s)
}
