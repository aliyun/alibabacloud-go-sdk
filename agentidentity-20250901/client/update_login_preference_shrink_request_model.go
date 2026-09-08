// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLoginPreferenceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllowedPostLogoutRedirectUrisShrink(v string) *UpdateLoginPreferenceShrinkRequest
	GetAllowedPostLogoutRedirectUrisShrink() *string
	SetLoginPreferenceShrink(v string) *UpdateLoginPreferenceShrinkRequest
	GetLoginPreferenceShrink() *string
	SetUserPoolName(v string) *UpdateLoginPreferenceShrinkRequest
	GetUserPoolName() *string
}

type UpdateLoginPreferenceShrinkRequest struct {
	AllowedPostLogoutRedirectUrisShrink *string `json:"AllowedPostLogoutRedirectUris,omitempty" xml:"AllowedPostLogoutRedirectUris,omitempty"`
	LoginPreferenceShrink               *string `json:"LoginPreference,omitempty" xml:"LoginPreference,omitempty"`
	UserPoolName                        *string `json:"UserPoolName,omitempty" xml:"UserPoolName,omitempty"`
}

func (s UpdateLoginPreferenceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateLoginPreferenceShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateLoginPreferenceShrinkRequest) GetAllowedPostLogoutRedirectUrisShrink() *string {
	return s.AllowedPostLogoutRedirectUrisShrink
}

func (s *UpdateLoginPreferenceShrinkRequest) GetLoginPreferenceShrink() *string {
	return s.LoginPreferenceShrink
}

func (s *UpdateLoginPreferenceShrinkRequest) GetUserPoolName() *string {
	return s.UserPoolName
}

func (s *UpdateLoginPreferenceShrinkRequest) SetAllowedPostLogoutRedirectUrisShrink(v string) *UpdateLoginPreferenceShrinkRequest {
	s.AllowedPostLogoutRedirectUrisShrink = &v
	return s
}

func (s *UpdateLoginPreferenceShrinkRequest) SetLoginPreferenceShrink(v string) *UpdateLoginPreferenceShrinkRequest {
	s.LoginPreferenceShrink = &v
	return s
}

func (s *UpdateLoginPreferenceShrinkRequest) SetUserPoolName(v string) *UpdateLoginPreferenceShrinkRequest {
	s.UserPoolName = &v
	return s
}

func (s *UpdateLoginPreferenceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
