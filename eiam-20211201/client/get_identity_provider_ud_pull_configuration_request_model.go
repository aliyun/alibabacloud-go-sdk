// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIdentityProviderUdPullConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIdentityProviderId(v string) *GetIdentityProviderUdPullConfigurationRequest
	GetIdentityProviderId() *string
	SetInstanceId(v string) *GetIdentityProviderUdPullConfigurationRequest
	GetInstanceId() *string
}

type GetIdentityProviderUdPullConfigurationRequest struct {
	// Identity provider ID
	//
	// This parameter is required.
	//
	// example:
	//
	// idp_my664lwkhpicbyzirog3xxxxx
	IdentityProviderId *string `json:"IdentityProviderId,omitempty" xml:"IdentityProviderId,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetIdentityProviderUdPullConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s GetIdentityProviderUdPullConfigurationRequest) GoString() string {
	return s.String()
}

func (s *GetIdentityProviderUdPullConfigurationRequest) GetIdentityProviderId() *string {
	return s.IdentityProviderId
}

func (s *GetIdentityProviderUdPullConfigurationRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetIdentityProviderUdPullConfigurationRequest) SetIdentityProviderId(v string) *GetIdentityProviderUdPullConfigurationRequest {
	s.IdentityProviderId = &v
	return s
}

func (s *GetIdentityProviderUdPullConfigurationRequest) SetInstanceId(v string) *GetIdentityProviderUdPullConfigurationRequest {
	s.InstanceId = &v
	return s
}

func (s *GetIdentityProviderUdPullConfigurationRequest) Validate() error {
	return dara.Validate(s)
}
