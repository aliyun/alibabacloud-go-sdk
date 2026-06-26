// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLoginPreferenceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLoginPreference(v *UpdateLoginPreferenceRequestLoginPreference) *UpdateLoginPreferenceRequest
	GetLoginPreference() *UpdateLoginPreferenceRequestLoginPreference
	SetUserPoolName(v string) *UpdateLoginPreferenceRequest
	GetUserPoolName() *string
}

type UpdateLoginPreferenceRequest struct {
	LoginPreference *UpdateLoginPreferenceRequestLoginPreference `json:"LoginPreference,omitempty" xml:"LoginPreference,omitempty" type:"Struct"`
	UserPoolName    *string                                      `json:"UserPoolName,omitempty" xml:"UserPoolName,omitempty"`
}

func (s UpdateLoginPreferenceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateLoginPreferenceRequest) GoString() string {
	return s.String()
}

func (s *UpdateLoginPreferenceRequest) GetLoginPreference() *UpdateLoginPreferenceRequestLoginPreference {
	return s.LoginPreference
}

func (s *UpdateLoginPreferenceRequest) GetUserPoolName() *string {
	return s.UserPoolName
}

func (s *UpdateLoginPreferenceRequest) SetLoginPreference(v *UpdateLoginPreferenceRequestLoginPreference) *UpdateLoginPreferenceRequest {
	s.LoginPreference = v
	return s
}

func (s *UpdateLoginPreferenceRequest) SetUserPoolName(v string) *UpdateLoginPreferenceRequest {
	s.UserPoolName = &v
	return s
}

func (s *UpdateLoginPreferenceRequest) Validate() error {
	if s.LoginPreference != nil {
		if err := s.LoginPreference.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateLoginPreferenceRequestLoginPreference struct {
	EnablePasswordLogin *bool `json:"EnablePasswordLogin,omitempty" xml:"EnablePasswordLogin,omitempty"`
}

func (s UpdateLoginPreferenceRequestLoginPreference) String() string {
	return dara.Prettify(s)
}

func (s UpdateLoginPreferenceRequestLoginPreference) GoString() string {
	return s.String()
}

func (s *UpdateLoginPreferenceRequestLoginPreference) GetEnablePasswordLogin() *bool {
	return s.EnablePasswordLogin
}

func (s *UpdateLoginPreferenceRequestLoginPreference) SetEnablePasswordLogin(v bool) *UpdateLoginPreferenceRequestLoginPreference {
	s.EnablePasswordLogin = &v
	return s
}

func (s *UpdateLoginPreferenceRequestLoginPreference) Validate() error {
	return dara.Validate(s)
}
