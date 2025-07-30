// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIdentityProviderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIdentityProviderId(v string) *GetIdentityProviderRequest
	GetIdentityProviderId() *string
	SetInstanceId(v string) *GetIdentityProviderRequest
	GetInstanceId() *string
}

type GetIdentityProviderRequest struct {
	// Identity provider ID.
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

func (s GetIdentityProviderRequest) String() string {
	return dara.Prettify(s)
}

func (s GetIdentityProviderRequest) GoString() string {
	return s.String()
}

func (s *GetIdentityProviderRequest) GetIdentityProviderId() *string {
	return s.IdentityProviderId
}

func (s *GetIdentityProviderRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetIdentityProviderRequest) SetIdentityProviderId(v string) *GetIdentityProviderRequest {
	s.IdentityProviderId = &v
	return s
}

func (s *GetIdentityProviderRequest) SetInstanceId(v string) *GetIdentityProviderRequest {
	s.InstanceId = &v
	return s
}

func (s *GetIdentityProviderRequest) Validate() error {
	return dara.Validate(s)
}
