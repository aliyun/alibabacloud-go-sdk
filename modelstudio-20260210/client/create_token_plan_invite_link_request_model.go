// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTokenPlanInviteLinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExpireType(v string) *CreateTokenPlanInviteLinkRequest
	GetExpireType() *string
	SetSsoSource(v string) *CreateTokenPlanInviteLinkRequest
	GetSsoSource() *string
}

type CreateTokenPlanInviteLinkRequest struct {
	// The expiration period. Default value: DAYS_7. Valid values:
	//
	// - DAYS_7
	//
	// - DAYS_30
	//
	// - MONTHS_6
	//
	// - YEAR_1
	//
	// example:
	//
	// DAYS_7
	ExpireType *string `json:"ExpireType,omitempty" xml:"ExpireType,omitempty"`
	// The SSO logon method bound to the invitation link. Valid values:
	//
	// - SAML
	//
	// - DINGTALK
	//
	// This parameter is required.
	//
	// example:
	//
	// SAML
	SsoSource *string `json:"SsoSource,omitempty" xml:"SsoSource,omitempty"`
}

func (s CreateTokenPlanInviteLinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTokenPlanInviteLinkRequest) GoString() string {
	return s.String()
}

func (s *CreateTokenPlanInviteLinkRequest) GetExpireType() *string {
	return s.ExpireType
}

func (s *CreateTokenPlanInviteLinkRequest) GetSsoSource() *string {
	return s.SsoSource
}

func (s *CreateTokenPlanInviteLinkRequest) SetExpireType(v string) *CreateTokenPlanInviteLinkRequest {
	s.ExpireType = &v
	return s
}

func (s *CreateTokenPlanInviteLinkRequest) SetSsoSource(v string) *CreateTokenPlanInviteLinkRequest {
	s.SsoSource = &v
	return s
}

func (s *CreateTokenPlanInviteLinkRequest) Validate() error {
	return dara.Validate(s)
}
