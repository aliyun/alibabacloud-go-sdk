// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableIdentityProviderUdPullRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIdentityProviderId(v string) *DisableIdentityProviderUdPullRequest
	GetIdentityProviderId() *string
	SetInstanceId(v string) *DisableIdentityProviderUdPullRequest
	GetInstanceId() *string
}

type DisableIdentityProviderUdPullRequest struct {
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

func (s DisableIdentityProviderUdPullRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableIdentityProviderUdPullRequest) GoString() string {
	return s.String()
}

func (s *DisableIdentityProviderUdPullRequest) GetIdentityProviderId() *string {
	return s.IdentityProviderId
}

func (s *DisableIdentityProviderUdPullRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DisableIdentityProviderUdPullRequest) SetIdentityProviderId(v string) *DisableIdentityProviderUdPullRequest {
	s.IdentityProviderId = &v
	return s
}

func (s *DisableIdentityProviderUdPullRequest) SetInstanceId(v string) *DisableIdentityProviderUdPullRequest {
	s.InstanceId = &v
	return s
}

func (s *DisableIdentityProviderUdPullRequest) Validate() error {
	return dara.Validate(s)
}
