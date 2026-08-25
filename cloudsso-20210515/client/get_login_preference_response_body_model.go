// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLoginPreferenceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLoginPreference(v *GetLoginPreferenceResponseBodyLoginPreference) *GetLoginPreferenceResponseBody
	GetLoginPreference() *GetLoginPreferenceResponseBodyLoginPreference
	SetRequestId(v string) *GetLoginPreferenceResponseBody
	GetRequestId() *string
}

type GetLoginPreferenceResponseBody struct {
	// The logon preference.
	LoginPreference *GetLoginPreferenceResponseBodyLoginPreference `json:"LoginPreference,omitempty" xml:"LoginPreference,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 8CE8B990-193D-50CE-A604-69F3E7DCE740
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetLoginPreferenceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetLoginPreferenceResponseBody) GoString() string {
	return s.String()
}

func (s *GetLoginPreferenceResponseBody) GetLoginPreference() *GetLoginPreferenceResponseBodyLoginPreference {
	return s.LoginPreference
}

func (s *GetLoginPreferenceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetLoginPreferenceResponseBody) SetLoginPreference(v *GetLoginPreferenceResponseBodyLoginPreference) *GetLoginPreferenceResponseBody {
	s.LoginPreference = v
	return s
}

func (s *GetLoginPreferenceResponseBody) SetRequestId(v string) *GetLoginPreferenceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLoginPreferenceResponseBody) Validate() error {
	if s.LoginPreference != nil {
		if err := s.LoginPreference.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetLoginPreferenceResponseBodyLoginPreference struct {
	// Indicates whether a user can obtain the application access credential after logon to the portal. Valid values:
	//
	// 	- True
	//
	// 	- False (default)
	//
	// example:
	//
	// True
	AllowUserToGetCredentials *bool `json:"AllowUserToGetCredentials,omitempty" xml:"AllowUserToGetCredentials,omitempty"`
	// The IP address whitelist. CloudSSO users can log on to the CloudSSO user portal only by using the IP addresses in the whitelist.
	//
	// The IP address whitelist takes effect only on CloudSSO users who want to log on to the CloudSSO user portal by using the username-password logon or single sign-on (SSO) method. The IP address whitelist does not take effect on CloudSSO users who access accounts in a resource directory from the CloudSSO user portal.
	//
	// If the return value of this parameter is empty, no IP address whitelists are configured.
	//
	// example:
	//
	// 192.168.0.0/16;10.0.0.0/8
	LoginNetworkMasks *string `json:"LoginNetworkMasks,omitempty" xml:"LoginNetworkMasks,omitempty"`
}

func (s GetLoginPreferenceResponseBodyLoginPreference) String() string {
	return dara.Prettify(s)
}

func (s GetLoginPreferenceResponseBodyLoginPreference) GoString() string {
	return s.String()
}

func (s *GetLoginPreferenceResponseBodyLoginPreference) GetAllowUserToGetCredentials() *bool {
	return s.AllowUserToGetCredentials
}

func (s *GetLoginPreferenceResponseBodyLoginPreference) GetLoginNetworkMasks() *string {
	return s.LoginNetworkMasks
}

func (s *GetLoginPreferenceResponseBodyLoginPreference) SetAllowUserToGetCredentials(v bool) *GetLoginPreferenceResponseBodyLoginPreference {
	s.AllowUserToGetCredentials = &v
	return s
}

func (s *GetLoginPreferenceResponseBodyLoginPreference) SetLoginNetworkMasks(v string) *GetLoginPreferenceResponseBodyLoginPreference {
	s.LoginNetworkMasks = &v
	return s
}

func (s *GetLoginPreferenceResponseBodyLoginPreference) Validate() error {
	return dara.Validate(s)
}
