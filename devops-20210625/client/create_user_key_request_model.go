// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUserKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *CreateUserKeyRequest
	GetAccessToken() *string
	SetExpireTime(v string) *CreateUserKeyRequest
	GetExpireTime() *string
	SetKeyScope(v string) *CreateUserKeyRequest
	GetKeyScope() *string
	SetPublicKey(v string) *CreateUserKeyRequest
	GetPublicKey() *string
	SetTitle(v string) *CreateUserKeyRequest
	GetTitle() *string
	SetOrganizationId(v string) *CreateUserKeyRequest
	GetOrganizationId() *string
}

type CreateUserKeyRequest struct {
	// example:
	//
	// f0b1e61db5961df5975a93f9129d2513
	AccessToken *string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	// example:
	//
	// 2022-03-12 12:00:00
	ExpireTime *string `json:"expireTime,omitempty" xml:"expireTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ALL
	KeyScope *string `json:"keyScope,omitempty" xml:"keyScope,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxx
	PublicKey *string `json:"publicKey,omitempty" xml:"publicKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// SSH Title
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 5ebbc0228123212b59xxxxx
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
}

func (s CreateUserKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateUserKeyRequest) GoString() string {
	return s.String()
}

func (s *CreateUserKeyRequest) GetAccessToken() *string {
	return s.AccessToken
}

func (s *CreateUserKeyRequest) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *CreateUserKeyRequest) GetKeyScope() *string {
	return s.KeyScope
}

func (s *CreateUserKeyRequest) GetPublicKey() *string {
	return s.PublicKey
}

func (s *CreateUserKeyRequest) GetTitle() *string {
	return s.Title
}

func (s *CreateUserKeyRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *CreateUserKeyRequest) SetAccessToken(v string) *CreateUserKeyRequest {
	s.AccessToken = &v
	return s
}

func (s *CreateUserKeyRequest) SetExpireTime(v string) *CreateUserKeyRequest {
	s.ExpireTime = &v
	return s
}

func (s *CreateUserKeyRequest) SetKeyScope(v string) *CreateUserKeyRequest {
	s.KeyScope = &v
	return s
}

func (s *CreateUserKeyRequest) SetPublicKey(v string) *CreateUserKeyRequest {
	s.PublicKey = &v
	return s
}

func (s *CreateUserKeyRequest) SetTitle(v string) *CreateUserKeyRequest {
	s.Title = &v
	return s
}

func (s *CreateUserKeyRequest) SetOrganizationId(v string) *CreateUserKeyRequest {
	s.OrganizationId = &v
	return s
}

func (s *CreateUserKeyRequest) Validate() error {
	return dara.Validate(s)
}
