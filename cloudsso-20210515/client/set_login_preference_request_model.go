// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetLoginPreferenceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllowUserToGetCredentials(v bool) *SetLoginPreferenceRequest
	GetAllowUserToGetCredentials() *bool
	SetDirectoryId(v string) *SetLoginPreferenceRequest
	GetDirectoryId() *string
	SetLoginNetworkMasks(v string) *SetLoginPreferenceRequest
	GetLoginNetworkMasks() *string
}

type SetLoginPreferenceRequest struct {
	// Specifies whether to allow a user to obtain the application access credential after logon to the portal. Valid values:
	//
	// 	- True
	//
	// 	- False (default)
	//
	// example:
	//
	// True
	AllowUserToGetCredentials *bool `json:"AllowUserToGetCredentials,omitempty" xml:"AllowUserToGetCredentials,omitempty"`
	// The ID of the directory.
	//
	// example:
	//
	// d-00fc2p61****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The IP address whitelist. CloudSSO users can log on to the CloudSSO user portal only by using the IP addresses in the whitelist. Limits:
	//
	// 	- You can enter IP addresses or CIDR blocks. IPv4 addresses are supported.
	//
	// 	- You can enter up to 100 IP addresses or CIDR blocks. Separate multiple IP addresses or CIDR blocks with semicolons `(;)`.
	//
	// 	- If you do not specify this parameter, the original settings are retained.
	//
	// 	- If you set this parameter to a semicolon (`;`), the value of this parameter is cleared.
	//
	// 	- The IP address whitelist takes effect only on CloudSSO users who want to log on to the CloudSSO user portal by using the username-password logon or single sign-on (SSO) method. The IP address whitelist does not take effect on CloudSSO users who access accounts in a resource directory from the CloudSSO user portal.
	//
	// example:
	//
	// 192.168.0.0/16;10.0.0.0/8
	LoginNetworkMasks *string `json:"LoginNetworkMasks,omitempty" xml:"LoginNetworkMasks,omitempty"`
}

func (s SetLoginPreferenceRequest) String() string {
	return dara.Prettify(s)
}

func (s SetLoginPreferenceRequest) GoString() string {
	return s.String()
}

func (s *SetLoginPreferenceRequest) GetAllowUserToGetCredentials() *bool {
	return s.AllowUserToGetCredentials
}

func (s *SetLoginPreferenceRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *SetLoginPreferenceRequest) GetLoginNetworkMasks() *string {
	return s.LoginNetworkMasks
}

func (s *SetLoginPreferenceRequest) SetAllowUserToGetCredentials(v bool) *SetLoginPreferenceRequest {
	s.AllowUserToGetCredentials = &v
	return s
}

func (s *SetLoginPreferenceRequest) SetDirectoryId(v string) *SetLoginPreferenceRequest {
	s.DirectoryId = &v
	return s
}

func (s *SetLoginPreferenceRequest) SetLoginNetworkMasks(v string) *SetLoginPreferenceRequest {
	s.LoginNetworkMasks = &v
	return s
}

func (s *SetLoginPreferenceRequest) Validate() error {
	return dara.Validate(s)
}
