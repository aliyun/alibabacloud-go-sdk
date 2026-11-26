// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetSecurityInfoPreferenceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetSecurityInfoPreferenceResponseBody
	GetRequestId() *string
	SetSecurityInfoPreference(v *SetSecurityInfoPreferenceResponseBodySecurityInfoPreference) *SetSecurityInfoPreferenceResponseBody
	GetSecurityInfoPreference() *SetSecurityInfoPreferenceResponseBodySecurityInfoPreference
}

type SetSecurityInfoPreferenceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 9B13E4EE-3853-5852-9165-597C32AD8FB7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The security information management preferences.
	SecurityInfoPreference *SetSecurityInfoPreferenceResponseBodySecurityInfoPreference `json:"SecurityInfoPreference,omitempty" xml:"SecurityInfoPreference,omitempty" type:"Struct"`
}

func (s SetSecurityInfoPreferenceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetSecurityInfoPreferenceResponseBody) GoString() string {
	return s.String()
}

func (s *SetSecurityInfoPreferenceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetSecurityInfoPreferenceResponseBody) GetSecurityInfoPreference() *SetSecurityInfoPreferenceResponseBodySecurityInfoPreference {
	return s.SecurityInfoPreference
}

func (s *SetSecurityInfoPreferenceResponseBody) SetRequestId(v string) *SetSecurityInfoPreferenceResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetSecurityInfoPreferenceResponseBody) SetSecurityInfoPreference(v *SetSecurityInfoPreferenceResponseBodySecurityInfoPreference) *SetSecurityInfoPreferenceResponseBody {
	s.SecurityInfoPreference = v
	return s
}

func (s *SetSecurityInfoPreferenceResponseBody) Validate() error {
	if s.SecurityInfoPreference != nil {
		if err := s.SecurityInfoPreference.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SetSecurityInfoPreferenceResponseBodySecurityInfoPreference struct {
	// Specifies whether CloudSSO users in the directory are allowed to manage their security email addresses in the CloudSSO user portal. Valid values:
	//
	// - true: allowed.
	//
	// - false: not allowed.
	//
	// example:
	//
	// false
	AllowUserManageSecurityEmail *bool `json:"AllowUserManageSecurityEmail,omitempty" xml:"AllowUserManageSecurityEmail,omitempty"`
}

func (s SetSecurityInfoPreferenceResponseBodySecurityInfoPreference) String() string {
	return dara.Prettify(s)
}

func (s SetSecurityInfoPreferenceResponseBodySecurityInfoPreference) GoString() string {
	return s.String()
}

func (s *SetSecurityInfoPreferenceResponseBodySecurityInfoPreference) GetAllowUserManageSecurityEmail() *bool {
	return s.AllowUserManageSecurityEmail
}

func (s *SetSecurityInfoPreferenceResponseBodySecurityInfoPreference) SetAllowUserManageSecurityEmail(v bool) *SetSecurityInfoPreferenceResponseBodySecurityInfoPreference {
	s.AllowUserManageSecurityEmail = &v
	return s
}

func (s *SetSecurityInfoPreferenceResponseBodySecurityInfoPreference) Validate() error {
	return dara.Validate(s)
}
