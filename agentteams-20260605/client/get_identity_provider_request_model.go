// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIdentityProviderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIdentityProviderType(v string) *GetIdentityProviderRequest
	GetIdentityProviderType() *string
	SetInstanceId(v string) *GetIdentityProviderRequest
	GetInstanceId() *string
}

type GetIdentityProviderRequest struct {
	// This parameter is required.
	IdentityProviderType *string `json:"IdentityProviderType,omitempty" xml:"IdentityProviderType,omitempty"`
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetIdentityProviderRequest) String() string {
	return dara.Prettify(s)
}

func (s GetIdentityProviderRequest) GoString() string {
	return s.String()
}

func (s *GetIdentityProviderRequest) GetIdentityProviderType() *string {
	return s.IdentityProviderType
}

func (s *GetIdentityProviderRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetIdentityProviderRequest) SetIdentityProviderType(v string) *GetIdentityProviderRequest {
	s.IdentityProviderType = &v
	return s
}

func (s *GetIdentityProviderRequest) SetInstanceId(v string) *GetIdentityProviderRequest {
	s.InstanceId = &v
	return s
}

func (s *GetIdentityProviderRequest) Validate() error {
	return dara.Validate(s)
}
