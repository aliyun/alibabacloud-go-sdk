// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCustomTemplateAndGroupConsoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigs(v string) *UpdateCustomTemplateAndGroupConsoleRequest
	GetConfigs() *string
	SetOwnerId(v int64) *UpdateCustomTemplateAndGroupConsoleRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *UpdateCustomTemplateAndGroupConsoleRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdateCustomTemplateAndGroupConsoleRequest
	GetResourceOwnerId() *int64
	SetResourceRealOwnerId(v int64) *UpdateCustomTemplateAndGroupConsoleRequest
	GetResourceRealOwnerId() *int64
}

type UpdateCustomTemplateAndGroupConsoleRequest struct {
	Configs              *string `json:"Configs,omitempty" xml:"Configs,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ResourceRealOwnerId  *int64  `json:"ResourceRealOwnerId,omitempty" xml:"ResourceRealOwnerId,omitempty"`
}

func (s UpdateCustomTemplateAndGroupConsoleRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomTemplateAndGroupConsoleRequest) GoString() string {
	return s.String()
}

func (s *UpdateCustomTemplateAndGroupConsoleRequest) GetConfigs() *string {
	return s.Configs
}

func (s *UpdateCustomTemplateAndGroupConsoleRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateCustomTemplateAndGroupConsoleRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateCustomTemplateAndGroupConsoleRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdateCustomTemplateAndGroupConsoleRequest) GetResourceRealOwnerId() *int64 {
	return s.ResourceRealOwnerId
}

func (s *UpdateCustomTemplateAndGroupConsoleRequest) SetConfigs(v string) *UpdateCustomTemplateAndGroupConsoleRequest {
	s.Configs = &v
	return s
}

func (s *UpdateCustomTemplateAndGroupConsoleRequest) SetOwnerId(v int64) *UpdateCustomTemplateAndGroupConsoleRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateCustomTemplateAndGroupConsoleRequest) SetResourceOwnerAccount(v string) *UpdateCustomTemplateAndGroupConsoleRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateCustomTemplateAndGroupConsoleRequest) SetResourceOwnerId(v int64) *UpdateCustomTemplateAndGroupConsoleRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateCustomTemplateAndGroupConsoleRequest) SetResourceRealOwnerId(v int64) *UpdateCustomTemplateAndGroupConsoleRequest {
	s.ResourceRealOwnerId = &v
	return s
}

func (s *UpdateCustomTemplateAndGroupConsoleRequest) Validate() error {
	return dara.Validate(s)
}
