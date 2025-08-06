// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCustomTemplateAndGroupConsoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigs(v string) *AddCustomTemplateAndGroupConsoleRequest
	GetConfigs() *string
	SetOwnerId(v int64) *AddCustomTemplateAndGroupConsoleRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *AddCustomTemplateAndGroupConsoleRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AddCustomTemplateAndGroupConsoleRequest
	GetResourceOwnerId() *int64
	SetResourceRealOwnerId(v int64) *AddCustomTemplateAndGroupConsoleRequest
	GetResourceRealOwnerId() *int64
}

type AddCustomTemplateAndGroupConsoleRequest struct {
	Configs              *string `json:"Configs,omitempty" xml:"Configs,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ResourceRealOwnerId  *int64  `json:"ResourceRealOwnerId,omitempty" xml:"ResourceRealOwnerId,omitempty"`
}

func (s AddCustomTemplateAndGroupConsoleRequest) String() string {
	return dara.Prettify(s)
}

func (s AddCustomTemplateAndGroupConsoleRequest) GoString() string {
	return s.String()
}

func (s *AddCustomTemplateAndGroupConsoleRequest) GetConfigs() *string {
	return s.Configs
}

func (s *AddCustomTemplateAndGroupConsoleRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddCustomTemplateAndGroupConsoleRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AddCustomTemplateAndGroupConsoleRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AddCustomTemplateAndGroupConsoleRequest) GetResourceRealOwnerId() *int64 {
	return s.ResourceRealOwnerId
}

func (s *AddCustomTemplateAndGroupConsoleRequest) SetConfigs(v string) *AddCustomTemplateAndGroupConsoleRequest {
	s.Configs = &v
	return s
}

func (s *AddCustomTemplateAndGroupConsoleRequest) SetOwnerId(v int64) *AddCustomTemplateAndGroupConsoleRequest {
	s.OwnerId = &v
	return s
}

func (s *AddCustomTemplateAndGroupConsoleRequest) SetResourceOwnerAccount(v string) *AddCustomTemplateAndGroupConsoleRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AddCustomTemplateAndGroupConsoleRequest) SetResourceOwnerId(v int64) *AddCustomTemplateAndGroupConsoleRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AddCustomTemplateAndGroupConsoleRequest) SetResourceRealOwnerId(v int64) *AddCustomTemplateAndGroupConsoleRequest {
	s.ResourceRealOwnerId = &v
	return s
}

func (s *AddCustomTemplateAndGroupConsoleRequest) Validate() error {
	return dara.Validate(s)
}
