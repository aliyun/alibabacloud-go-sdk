// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserPoolClientResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClient(v *GetUserPoolClientResponseBodyClient) *GetUserPoolClientResponseBody
	GetClient() *GetUserPoolClientResponseBodyClient
	SetRequestId(v string) *GetUserPoolClientResponseBody
	GetRequestId() *string
}

type GetUserPoolClientResponseBody struct {
	Client *GetUserPoolClientResponseBodyClient `json:"Client,omitempty" xml:"Client,omitempty" type:"Struct"`
	// example:
	//
	// AABD6E03-4B3A-5687-88FF-72232670ED0C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetUserPoolClientResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUserPoolClientResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserPoolClientResponseBody) GetClient() *GetUserPoolClientResponseBodyClient {
	return s.Client
}

func (s *GetUserPoolClientResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUserPoolClientResponseBody) SetClient(v *GetUserPoolClientResponseBodyClient) *GetUserPoolClientResponseBody {
	s.Client = v
	return s
}

func (s *GetUserPoolClientResponseBody) SetRequestId(v string) *GetUserPoolClientResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserPoolClientResponseBody) Validate() error {
	if s.Client != nil {
		if err := s.Client.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetUserPoolClientResponseBodyClient struct {
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
	ClientName   *string                                            `json:"ClientName,omitempty" xml:"ClientName,omitempty"`
	ClientScopes []*GetUserPoolClientResponseBodyClientClientScopes `json:"ClientScopes,omitempty" xml:"ClientScopes,omitempty" type:"Repeated"`
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

func (s GetUserPoolClientResponseBodyClient) String() string {
	return dara.Prettify(s)
}

func (s GetUserPoolClientResponseBodyClient) GoString() string {
	return s.String()
}

func (s *GetUserPoolClientResponseBodyClient) GetAccessTokenValidity() *string {
	return s.AccessTokenValidity
}

func (s *GetUserPoolClientResponseBodyClient) GetClientId() *string {
	return s.ClientId
}

func (s *GetUserPoolClientResponseBodyClient) GetClientName() *string {
	return s.ClientName
}

func (s *GetUserPoolClientResponseBodyClient) GetClientScopes() []*GetUserPoolClientResponseBodyClientClientScopes {
	return s.ClientScopes
}

func (s *GetUserPoolClientResponseBodyClient) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetUserPoolClientResponseBodyClient) GetEnforcePKCE() *bool {
	return s.EnforcePKCE
}

func (s *GetUserPoolClientResponseBodyClient) GetRedirectURIs() []*string {
	return s.RedirectURIs
}

func (s *GetUserPoolClientResponseBodyClient) GetRefreshTokenValidity() *string {
	return s.RefreshTokenValidity
}

func (s *GetUserPoolClientResponseBodyClient) GetSecretRequired() *bool {
	return s.SecretRequired
}

func (s *GetUserPoolClientResponseBodyClient) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetUserPoolClientResponseBodyClient) GetUserPoolName() *string {
	return s.UserPoolName
}

func (s *GetUserPoolClientResponseBodyClient) SetAccessTokenValidity(v string) *GetUserPoolClientResponseBodyClient {
	s.AccessTokenValidity = &v
	return s
}

func (s *GetUserPoolClientResponseBodyClient) SetClientId(v string) *GetUserPoolClientResponseBodyClient {
	s.ClientId = &v
	return s
}

func (s *GetUserPoolClientResponseBodyClient) SetClientName(v string) *GetUserPoolClientResponseBodyClient {
	s.ClientName = &v
	return s
}

func (s *GetUserPoolClientResponseBodyClient) SetClientScopes(v []*GetUserPoolClientResponseBodyClientClientScopes) *GetUserPoolClientResponseBodyClient {
	s.ClientScopes = v
	return s
}

func (s *GetUserPoolClientResponseBodyClient) SetCreateTime(v string) *GetUserPoolClientResponseBodyClient {
	s.CreateTime = &v
	return s
}

func (s *GetUserPoolClientResponseBodyClient) SetEnforcePKCE(v bool) *GetUserPoolClientResponseBodyClient {
	s.EnforcePKCE = &v
	return s
}

func (s *GetUserPoolClientResponseBodyClient) SetRedirectURIs(v []*string) *GetUserPoolClientResponseBodyClient {
	s.RedirectURIs = v
	return s
}

func (s *GetUserPoolClientResponseBodyClient) SetRefreshTokenValidity(v string) *GetUserPoolClientResponseBodyClient {
	s.RefreshTokenValidity = &v
	return s
}

func (s *GetUserPoolClientResponseBodyClient) SetSecretRequired(v bool) *GetUserPoolClientResponseBodyClient {
	s.SecretRequired = &v
	return s
}

func (s *GetUserPoolClientResponseBodyClient) SetUpdateTime(v string) *GetUserPoolClientResponseBodyClient {
	s.UpdateTime = &v
	return s
}

func (s *GetUserPoolClientResponseBodyClient) SetUserPoolName(v string) *GetUserPoolClientResponseBodyClient {
	s.UserPoolName = &v
	return s
}

func (s *GetUserPoolClientResponseBodyClient) Validate() error {
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

type GetUserPoolClientResponseBodyClientClientScopes struct {
	// example:
	//
	// OpenID Connect authentication
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// openid
	ScopeName *string `json:"ScopeName,omitempty" xml:"ScopeName,omitempty"`
}

func (s GetUserPoolClientResponseBodyClientClientScopes) String() string {
	return dara.Prettify(s)
}

func (s GetUserPoolClientResponseBodyClientClientScopes) GoString() string {
	return s.String()
}

func (s *GetUserPoolClientResponseBodyClientClientScopes) GetDescription() *string {
	return s.Description
}

func (s *GetUserPoolClientResponseBodyClientClientScopes) GetScopeName() *string {
	return s.ScopeName
}

func (s *GetUserPoolClientResponseBodyClientClientScopes) SetDescription(v string) *GetUserPoolClientResponseBodyClientClientScopes {
	s.Description = &v
	return s
}

func (s *GetUserPoolClientResponseBodyClientClientScopes) SetScopeName(v string) *GetUserPoolClientResponseBodyClientClientScopes {
	s.ScopeName = &v
	return s
}

func (s *GetUserPoolClientResponseBodyClientClientScopes) Validate() error {
	return dara.Validate(s)
}
