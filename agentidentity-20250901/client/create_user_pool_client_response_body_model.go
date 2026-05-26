// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUserPoolClientResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClient(v *CreateUserPoolClientResponseBodyClient) *CreateUserPoolClientResponseBody
	GetClient() *CreateUserPoolClientResponseBodyClient
	SetRequestId(v string) *CreateUserPoolClientResponseBody
	GetRequestId() *string
}

type CreateUserPoolClientResponseBody struct {
	Client *CreateUserPoolClientResponseBodyClient `json:"Client,omitempty" xml:"Client,omitempty" type:"Struct"`
	// example:
	//
	// AABD6E03-4B3A-5687-88FF-72232670ED0C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateUserPoolClientResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateUserPoolClientResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUserPoolClientResponseBody) GetClient() *CreateUserPoolClientResponseBodyClient {
	return s.Client
}

func (s *CreateUserPoolClientResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateUserPoolClientResponseBody) SetClient(v *CreateUserPoolClientResponseBodyClient) *CreateUserPoolClientResponseBody {
	s.Client = v
	return s
}

func (s *CreateUserPoolClientResponseBody) SetRequestId(v string) *CreateUserPoolClientResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateUserPoolClientResponseBody) Validate() error {
	if s.Client != nil {
		if err := s.Client.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateUserPoolClientResponseBodyClient struct {
	// example:
	//
	// 3600
	AccessTokenValidity *string `json:"AccessTokenValidity,omitempty" xml:"AccessTokenValidity,omitempty"`
	// example:
	//
	// client_xxxxxxxxxxxxxxxxxxxx
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// example:
	//
	// my-web-app
	ClientName   *string                                               `json:"ClientName,omitempty" xml:"ClientName,omitempty"`
	ClientScopes []*CreateUserPoolClientResponseBodyClientClientScopes `json:"ClientScopes,omitempty" xml:"ClientScopes,omitempty" type:"Repeated"`
	// example:
	//
	// 2026-05-07T06:19:17Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// true
	EnforcePKCE  *bool     `json:"EnforcePKCE,omitempty" xml:"EnforcePKCE,omitempty"`
	RedirectURIs []*string `json:"RedirectURIs,omitempty" xml:"RedirectURIs,omitempty" type:"Repeated"`
	// example:
	//
	// 86400
	RefreshTokenValidity *string `json:"RefreshTokenValidity,omitempty" xml:"RefreshTokenValidity,omitempty"`
	// example:
	//
	// true
	SecretRequired *bool `json:"SecretRequired,omitempty" xml:"SecretRequired,omitempty"`
	// example:
	//
	// 2026-05-07T06:19:17Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// example:
	//
	// my-agent-userpool
	UserPoolName *string `json:"UserPoolName,omitempty" xml:"UserPoolName,omitempty"`
}

func (s CreateUserPoolClientResponseBodyClient) String() string {
	return dara.Prettify(s)
}

func (s CreateUserPoolClientResponseBodyClient) GoString() string {
	return s.String()
}

func (s *CreateUserPoolClientResponseBodyClient) GetAccessTokenValidity() *string {
	return s.AccessTokenValidity
}

func (s *CreateUserPoolClientResponseBodyClient) GetClientId() *string {
	return s.ClientId
}

func (s *CreateUserPoolClientResponseBodyClient) GetClientName() *string {
	return s.ClientName
}

func (s *CreateUserPoolClientResponseBodyClient) GetClientScopes() []*CreateUserPoolClientResponseBodyClientClientScopes {
	return s.ClientScopes
}

func (s *CreateUserPoolClientResponseBodyClient) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CreateUserPoolClientResponseBodyClient) GetEnforcePKCE() *bool {
	return s.EnforcePKCE
}

func (s *CreateUserPoolClientResponseBodyClient) GetRedirectURIs() []*string {
	return s.RedirectURIs
}

func (s *CreateUserPoolClientResponseBodyClient) GetRefreshTokenValidity() *string {
	return s.RefreshTokenValidity
}

func (s *CreateUserPoolClientResponseBodyClient) GetSecretRequired() *bool {
	return s.SecretRequired
}

func (s *CreateUserPoolClientResponseBodyClient) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *CreateUserPoolClientResponseBodyClient) GetUserPoolName() *string {
	return s.UserPoolName
}

func (s *CreateUserPoolClientResponseBodyClient) SetAccessTokenValidity(v string) *CreateUserPoolClientResponseBodyClient {
	s.AccessTokenValidity = &v
	return s
}

func (s *CreateUserPoolClientResponseBodyClient) SetClientId(v string) *CreateUserPoolClientResponseBodyClient {
	s.ClientId = &v
	return s
}

func (s *CreateUserPoolClientResponseBodyClient) SetClientName(v string) *CreateUserPoolClientResponseBodyClient {
	s.ClientName = &v
	return s
}

func (s *CreateUserPoolClientResponseBodyClient) SetClientScopes(v []*CreateUserPoolClientResponseBodyClientClientScopes) *CreateUserPoolClientResponseBodyClient {
	s.ClientScopes = v
	return s
}

func (s *CreateUserPoolClientResponseBodyClient) SetCreateTime(v string) *CreateUserPoolClientResponseBodyClient {
	s.CreateTime = &v
	return s
}

func (s *CreateUserPoolClientResponseBodyClient) SetEnforcePKCE(v bool) *CreateUserPoolClientResponseBodyClient {
	s.EnforcePKCE = &v
	return s
}

func (s *CreateUserPoolClientResponseBodyClient) SetRedirectURIs(v []*string) *CreateUserPoolClientResponseBodyClient {
	s.RedirectURIs = v
	return s
}

func (s *CreateUserPoolClientResponseBodyClient) SetRefreshTokenValidity(v string) *CreateUserPoolClientResponseBodyClient {
	s.RefreshTokenValidity = &v
	return s
}

func (s *CreateUserPoolClientResponseBodyClient) SetSecretRequired(v bool) *CreateUserPoolClientResponseBodyClient {
	s.SecretRequired = &v
	return s
}

func (s *CreateUserPoolClientResponseBodyClient) SetUpdateTime(v string) *CreateUserPoolClientResponseBodyClient {
	s.UpdateTime = &v
	return s
}

func (s *CreateUserPoolClientResponseBodyClient) SetUserPoolName(v string) *CreateUserPoolClientResponseBodyClient {
	s.UserPoolName = &v
	return s
}

func (s *CreateUserPoolClientResponseBodyClient) Validate() error {
	if s.ClientScopes != nil {
		for _, item := range s.ClientScopes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateUserPoolClientResponseBodyClientClientScopes struct {
	// example:
	//
	// OpenID Connect authentication
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// openid
	ScopeName *string `json:"ScopeName,omitempty" xml:"ScopeName,omitempty"`
}

func (s CreateUserPoolClientResponseBodyClientClientScopes) String() string {
	return dara.Prettify(s)
}

func (s CreateUserPoolClientResponseBodyClientClientScopes) GoString() string {
	return s.String()
}

func (s *CreateUserPoolClientResponseBodyClientClientScopes) GetDescription() *string {
	return s.Description
}

func (s *CreateUserPoolClientResponseBodyClientClientScopes) GetScopeName() *string {
	return s.ScopeName
}

func (s *CreateUserPoolClientResponseBodyClientClientScopes) SetDescription(v string) *CreateUserPoolClientResponseBodyClientClientScopes {
	s.Description = &v
	return s
}

func (s *CreateUserPoolClientResponseBodyClientClientScopes) SetScopeName(v string) *CreateUserPoolClientResponseBodyClientClientScopes {
	s.ScopeName = &v
	return s
}

func (s *CreateUserPoolClientResponseBodyClientClientScopes) Validate() error {
	return dara.Validate(s)
}
