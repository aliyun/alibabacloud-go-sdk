// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSecurityInfoPreferenceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetSecurityInfoPreferenceResponseBody
	GetRequestId() *string
	SetSecurityInfoPreference(v *GetSecurityInfoPreferenceResponseBodySecurityInfoPreference) *GetSecurityInfoPreferenceResponseBody
	GetSecurityInfoPreference() *GetSecurityInfoPreferenceResponseBodySecurityInfoPreference
}

type GetSecurityInfoPreferenceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 8CE8B990-193D-50CE-A604-69F3E7DCE740
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The security information management preferences.
	SecurityInfoPreference *GetSecurityInfoPreferenceResponseBodySecurityInfoPreference `json:"SecurityInfoPreference,omitempty" xml:"SecurityInfoPreference,omitempty" type:"Struct"`
}

func (s GetSecurityInfoPreferenceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSecurityInfoPreferenceResponseBody) GoString() string {
	return s.String()
}

func (s *GetSecurityInfoPreferenceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSecurityInfoPreferenceResponseBody) GetSecurityInfoPreference() *GetSecurityInfoPreferenceResponseBodySecurityInfoPreference {
	return s.SecurityInfoPreference
}

func (s *GetSecurityInfoPreferenceResponseBody) SetRequestId(v string) *GetSecurityInfoPreferenceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSecurityInfoPreferenceResponseBody) SetSecurityInfoPreference(v *GetSecurityInfoPreferenceResponseBodySecurityInfoPreference) *GetSecurityInfoPreferenceResponseBody {
	s.SecurityInfoPreference = v
	return s
}

func (s *GetSecurityInfoPreferenceResponseBody) Validate() error {
	if s.SecurityInfoPreference != nil {
		if err := s.SecurityInfoPreference.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSecurityInfoPreferenceResponseBodySecurityInfoPreference struct {
	// Indicates whether CloudSSO users in the directory are allowed to manage their security email addresses in the CloudSSO user portal. Valid values:
	//
	// - true: Allowed.
	//
	// - false: Not allowed. This is the default value.
	//
	// example:
	//
	// false
	AllowUserManageSecurityEmail *bool `json:"AllowUserManageSecurityEmail,omitempty" xml:"AllowUserManageSecurityEmail,omitempty"`
}

func (s GetSecurityInfoPreferenceResponseBodySecurityInfoPreference) String() string {
	return dara.Prettify(s)
}

func (s GetSecurityInfoPreferenceResponseBodySecurityInfoPreference) GoString() string {
	return s.String()
}

func (s *GetSecurityInfoPreferenceResponseBodySecurityInfoPreference) GetAllowUserManageSecurityEmail() *bool {
	return s.AllowUserManageSecurityEmail
}

func (s *GetSecurityInfoPreferenceResponseBodySecurityInfoPreference) SetAllowUserManageSecurityEmail(v bool) *GetSecurityInfoPreferenceResponseBodySecurityInfoPreference {
	s.AllowUserManageSecurityEmail = &v
	return s
}

func (s *GetSecurityInfoPreferenceResponseBodySecurityInfoPreference) Validate() error {
	return dara.Validate(s)
}
