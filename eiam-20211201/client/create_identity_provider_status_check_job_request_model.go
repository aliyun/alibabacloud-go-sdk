// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIdentityProviderStatusCheckJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIdentityProviderId(v string) *CreateIdentityProviderStatusCheckJobRequest
	GetIdentityProviderId() *string
	SetInstanceId(v string) *CreateIdentityProviderStatusCheckJobRequest
	GetInstanceId() *string
}

type CreateIdentityProviderStatusCheckJobRequest struct {
	// IDaaS的身份提供方主键id
	//
	// This parameter is required.
	//
	// example:
	//
	// idp_11111
	IdentityProviderId *string `json:"IdentityProviderId,omitempty" xml:"IdentityProviderId,omitempty"`
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s CreateIdentityProviderStatusCheckJobRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateIdentityProviderStatusCheckJobRequest) GoString() string {
	return s.String()
}

func (s *CreateIdentityProviderStatusCheckJobRequest) GetIdentityProviderId() *string {
	return s.IdentityProviderId
}

func (s *CreateIdentityProviderStatusCheckJobRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateIdentityProviderStatusCheckJobRequest) SetIdentityProviderId(v string) *CreateIdentityProviderStatusCheckJobRequest {
	s.IdentityProviderId = &v
	return s
}

func (s *CreateIdentityProviderStatusCheckJobRequest) SetInstanceId(v string) *CreateIdentityProviderStatusCheckJobRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateIdentityProviderStatusCheckJobRequest) Validate() error {
	return dara.Validate(s)
}
