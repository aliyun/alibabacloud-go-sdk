// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserPoolClientsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClients(v []*ListUserPoolClientsResponseBodyClients) *ListUserPoolClientsResponseBody
	GetClients() []*ListUserPoolClientsResponseBodyClients
	SetMaxResults(v int32) *ListUserPoolClientsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListUserPoolClientsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListUserPoolClientsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListUserPoolClientsResponseBody
	GetTotalCount() *int32
}

type ListUserPoolClientsResponseBody struct {
	Clients []*ListUserPoolClientsResponseBodyClients `json:"Clients,omitempty" xml:"Clients,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// TGlzdENsaWVudHM6OjIw
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// AABD6E03-4B3A-5687-88FF-72232670ED0C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 5
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListUserPoolClientsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUserPoolClientsResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserPoolClientsResponseBody) GetClients() []*ListUserPoolClientsResponseBodyClients {
	return s.Clients
}

func (s *ListUserPoolClientsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListUserPoolClientsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListUserPoolClientsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListUserPoolClientsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListUserPoolClientsResponseBody) SetClients(v []*ListUserPoolClientsResponseBodyClients) *ListUserPoolClientsResponseBody {
	s.Clients = v
	return s
}

func (s *ListUserPoolClientsResponseBody) SetMaxResults(v int32) *ListUserPoolClientsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListUserPoolClientsResponseBody) SetNextToken(v string) *ListUserPoolClientsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListUserPoolClientsResponseBody) SetRequestId(v string) *ListUserPoolClientsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserPoolClientsResponseBody) SetTotalCount(v int32) *ListUserPoolClientsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListUserPoolClientsResponseBody) Validate() error {
	if s.Clients != nil {
		for _, item := range s.Clients {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListUserPoolClientsResponseBodyClients struct {
	// example:
	//
	// 3600
	AccessTokenValidity *string `json:"AccessTokenValidity,omitempty" xml:"AccessTokenValidity,omitempty"`
	// example:
	//
	// client-xxxxxxxxxxxxxxxxxxxx
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// example:
	//
	// my-web-app
	ClientName   *string                                               `json:"ClientName,omitempty" xml:"ClientName,omitempty"`
	ClientScopes []*ListUserPoolClientsResponseBodyClientsClientScopes `json:"ClientScopes,omitempty" xml:"ClientScopes,omitempty" type:"Repeated"`
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

func (s ListUserPoolClientsResponseBodyClients) String() string {
	return dara.Prettify(s)
}

func (s ListUserPoolClientsResponseBodyClients) GoString() string {
	return s.String()
}

func (s *ListUserPoolClientsResponseBodyClients) GetAccessTokenValidity() *string {
	return s.AccessTokenValidity
}

func (s *ListUserPoolClientsResponseBodyClients) GetClientId() *string {
	return s.ClientId
}

func (s *ListUserPoolClientsResponseBodyClients) GetClientName() *string {
	return s.ClientName
}

func (s *ListUserPoolClientsResponseBodyClients) GetClientScopes() []*ListUserPoolClientsResponseBodyClientsClientScopes {
	return s.ClientScopes
}

func (s *ListUserPoolClientsResponseBodyClients) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListUserPoolClientsResponseBodyClients) GetEnforcePKCE() *bool {
	return s.EnforcePKCE
}

func (s *ListUserPoolClientsResponseBodyClients) GetRedirectURIs() []*string {
	return s.RedirectURIs
}

func (s *ListUserPoolClientsResponseBodyClients) GetRefreshTokenValidity() *string {
	return s.RefreshTokenValidity
}

func (s *ListUserPoolClientsResponseBodyClients) GetSecretRequired() *bool {
	return s.SecretRequired
}

func (s *ListUserPoolClientsResponseBodyClients) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListUserPoolClientsResponseBodyClients) GetUserPoolName() *string {
	return s.UserPoolName
}

func (s *ListUserPoolClientsResponseBodyClients) SetAccessTokenValidity(v string) *ListUserPoolClientsResponseBodyClients {
	s.AccessTokenValidity = &v
	return s
}

func (s *ListUserPoolClientsResponseBodyClients) SetClientId(v string) *ListUserPoolClientsResponseBodyClients {
	s.ClientId = &v
	return s
}

func (s *ListUserPoolClientsResponseBodyClients) SetClientName(v string) *ListUserPoolClientsResponseBodyClients {
	s.ClientName = &v
	return s
}

func (s *ListUserPoolClientsResponseBodyClients) SetClientScopes(v []*ListUserPoolClientsResponseBodyClientsClientScopes) *ListUserPoolClientsResponseBodyClients {
	s.ClientScopes = v
	return s
}

func (s *ListUserPoolClientsResponseBodyClients) SetCreateTime(v string) *ListUserPoolClientsResponseBodyClients {
	s.CreateTime = &v
	return s
}

func (s *ListUserPoolClientsResponseBodyClients) SetEnforcePKCE(v bool) *ListUserPoolClientsResponseBodyClients {
	s.EnforcePKCE = &v
	return s
}

func (s *ListUserPoolClientsResponseBodyClients) SetRedirectURIs(v []*string) *ListUserPoolClientsResponseBodyClients {
	s.RedirectURIs = v
	return s
}

func (s *ListUserPoolClientsResponseBodyClients) SetRefreshTokenValidity(v string) *ListUserPoolClientsResponseBodyClients {
	s.RefreshTokenValidity = &v
	return s
}

func (s *ListUserPoolClientsResponseBodyClients) SetSecretRequired(v bool) *ListUserPoolClientsResponseBodyClients {
	s.SecretRequired = &v
	return s
}

func (s *ListUserPoolClientsResponseBodyClients) SetUpdateTime(v string) *ListUserPoolClientsResponseBodyClients {
	s.UpdateTime = &v
	return s
}

func (s *ListUserPoolClientsResponseBodyClients) SetUserPoolName(v string) *ListUserPoolClientsResponseBodyClients {
	s.UserPoolName = &v
	return s
}

func (s *ListUserPoolClientsResponseBodyClients) Validate() error {
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

type ListUserPoolClientsResponseBodyClientsClientScopes struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// openid
	ScopeName *string `json:"ScopeName,omitempty" xml:"ScopeName,omitempty"`
}

func (s ListUserPoolClientsResponseBodyClientsClientScopes) String() string {
	return dara.Prettify(s)
}

func (s ListUserPoolClientsResponseBodyClientsClientScopes) GoString() string {
	return s.String()
}

func (s *ListUserPoolClientsResponseBodyClientsClientScopes) GetDescription() *string {
	return s.Description
}

func (s *ListUserPoolClientsResponseBodyClientsClientScopes) GetScopeName() *string {
	return s.ScopeName
}

func (s *ListUserPoolClientsResponseBodyClientsClientScopes) SetDescription(v string) *ListUserPoolClientsResponseBodyClientsClientScopes {
	s.Description = &v
	return s
}

func (s *ListUserPoolClientsResponseBodyClientsClientScopes) SetScopeName(v string) *ListUserPoolClientsResponseBodyClientsClientScopes {
	s.ScopeName = &v
	return s
}

func (s *ListUserPoolClientsResponseBodyClientsClientScopes) Validate() error {
	return dara.Validate(s)
}
