// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIdentityProviderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIdentityProviderId(v string) *DeleteIdentityProviderRequest
	GetIdentityProviderId() *string
	SetInstanceId(v string) *DeleteIdentityProviderRequest
	GetInstanceId() *string
}

type DeleteIdentityProviderRequest struct {
	// Identity provider ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idp_my664lwkhpicbyzirog3xxxxx
	IdentityProviderId *string `json:"IdentityProviderId,omitempty" xml:"IdentityProviderId,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteIdentityProviderRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteIdentityProviderRequest) GoString() string {
	return s.String()
}

func (s *DeleteIdentityProviderRequest) GetIdentityProviderId() *string {
	return s.IdentityProviderId
}

func (s *DeleteIdentityProviderRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteIdentityProviderRequest) SetIdentityProviderId(v string) *DeleteIdentityProviderRequest {
	s.IdentityProviderId = &v
	return s
}

func (s *DeleteIdentityProviderRequest) SetInstanceId(v string) *DeleteIdentityProviderRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteIdentityProviderRequest) Validate() error {
	return dara.Validate(s)
}
