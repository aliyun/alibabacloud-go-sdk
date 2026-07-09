// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAuthCodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountType(v string) *GetAuthCodeRequest
	GetAccountType() *string
	SetAdDomain(v string) *GetAuthCodeRequest
	GetAdDomain() *string
	SetAdPassword(v string) *GetAuthCodeRequest
	GetAdPassword() *string
	SetAutoCreateUser(v bool) *GetAuthCodeRequest
	GetAutoCreateUser() *bool
	SetEndUserId(v string) *GetAuthCodeRequest
	GetEndUserId() *string
	SetExternalUserId(v string) *GetAuthCodeRequest
	GetExternalUserId() *string
	SetPolicy(v string) *GetAuthCodeRequest
	GetPolicy() *string
	SetTokenType(v string) *GetAuthCodeRequest
	GetTokenType() *string
}

type GetAuthCodeRequest struct {
	AccountType *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	AdDomain    *string `json:"AdDomain,omitempty" xml:"AdDomain,omitempty"`
	AdPassword  *string `json:"AdPassword,omitempty" xml:"AdPassword,omitempty"`
	// Specifies whether to synchronously create an EndUserId based on ExternalUserId. This parameter takes effect only when EndUserId is empty.
	//
	// example:
	//
	// false
	AutoCreateUser *bool `json:"AutoCreateUser,omitempty" xml:"AutoCreateUser,omitempty"`
	// The username of the China Desktop Service (China Desktop Service) convenience account. The username must be unique within an Alibaba Cloud account. This parameter and ExternalUserId cannot both be empty.
	//
	// example:
	//
	// alice
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The external user ID. This ID is defined by the caller and must be unique within an Alibaba Cloud account. This parameter and EndUserId cannot both be empty.
	//
	// example:
	//
	// alice
	ExternalUserId *string `json:"ExternalUserId,omitempty" xml:"ExternalUserId,omitempty"`
	// The access policy that restricts the access permissions of the authorization code. If this parameter is left empty, no restrictions are applied.
	//
	// Syntax:
	//
	// ```json
	//
	// {
	//
	//       "Version": "1",
	//
	//       "Resource": {
	//
	//             "Type": "<Resource type>",
	//
	//             "Id": "<Resource ID>"
	//
	//       }
	//
	// }
	//
	// ```
	//
	// Valid values of <Resource type>:
	//
	// - AppInstanceGroup: delivery group. You can call the ListAppInstanceGroup operation to obtain the ID.
	//
	// - AppInstance: application instance (dedicated field).
	//
	// - App: application. You can call the ListAppInstanceGroup operation to obtain the ID.
	//
	// example:
	//
	// {
	//
	//       "Version": "1",
	//
	//       "Resource": {
	//
	//             "Type": "AppInstanceGroup",
	//
	//             "Id": "aig-9ciijz60n4xsv****"
	//
	//       }
	//
	// }
	Policy    *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	TokenType *string `json:"TokenType,omitempty" xml:"TokenType,omitempty"`
}

func (s GetAuthCodeRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAuthCodeRequest) GoString() string {
	return s.String()
}

func (s *GetAuthCodeRequest) GetAccountType() *string {
	return s.AccountType
}

func (s *GetAuthCodeRequest) GetAdDomain() *string {
	return s.AdDomain
}

func (s *GetAuthCodeRequest) GetAdPassword() *string {
	return s.AdPassword
}

func (s *GetAuthCodeRequest) GetAutoCreateUser() *bool {
	return s.AutoCreateUser
}

func (s *GetAuthCodeRequest) GetEndUserId() *string {
	return s.EndUserId
}

func (s *GetAuthCodeRequest) GetExternalUserId() *string {
	return s.ExternalUserId
}

func (s *GetAuthCodeRequest) GetPolicy() *string {
	return s.Policy
}

func (s *GetAuthCodeRequest) GetTokenType() *string {
	return s.TokenType
}

func (s *GetAuthCodeRequest) SetAccountType(v string) *GetAuthCodeRequest {
	s.AccountType = &v
	return s
}

func (s *GetAuthCodeRequest) SetAdDomain(v string) *GetAuthCodeRequest {
	s.AdDomain = &v
	return s
}

func (s *GetAuthCodeRequest) SetAdPassword(v string) *GetAuthCodeRequest {
	s.AdPassword = &v
	return s
}

func (s *GetAuthCodeRequest) SetAutoCreateUser(v bool) *GetAuthCodeRequest {
	s.AutoCreateUser = &v
	return s
}

func (s *GetAuthCodeRequest) SetEndUserId(v string) *GetAuthCodeRequest {
	s.EndUserId = &v
	return s
}

func (s *GetAuthCodeRequest) SetExternalUserId(v string) *GetAuthCodeRequest {
	s.ExternalUserId = &v
	return s
}

func (s *GetAuthCodeRequest) SetPolicy(v string) *GetAuthCodeRequest {
	s.Policy = &v
	return s
}

func (s *GetAuthCodeRequest) SetTokenType(v string) *GetAuthCodeRequest {
	s.TokenType = &v
	return s
}

func (s *GetAuthCodeRequest) Validate() error {
	return dara.Validate(s)
}
