// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindIdentityProviderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIdentityProviderType(v string) *UnbindIdentityProviderRequest
	GetIdentityProviderType() *string
	SetInstanceId(v string) *UnbindIdentityProviderRequest
	GetInstanceId() *string
}

type UnbindIdentityProviderRequest struct {
	IdentityProviderType *string `json:"IdentityProviderType,omitempty" xml:"IdentityProviderType,omitempty"`
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s UnbindIdentityProviderRequest) String() string {
	return dara.Prettify(s)
}

func (s UnbindIdentityProviderRequest) GoString() string {
	return s.String()
}

func (s *UnbindIdentityProviderRequest) GetIdentityProviderType() *string {
	return s.IdentityProviderType
}

func (s *UnbindIdentityProviderRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UnbindIdentityProviderRequest) SetIdentityProviderType(v string) *UnbindIdentityProviderRequest {
	s.IdentityProviderType = &v
	return s
}

func (s *UnbindIdentityProviderRequest) SetInstanceId(v string) *UnbindIdentityProviderRequest {
	s.InstanceId = &v
	return s
}

func (s *UnbindIdentityProviderRequest) Validate() error {
	return dara.Validate(s)
}
