// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *DeleteUserKeyRequest
	GetAccessToken() *string
	SetOrganizationId(v string) *DeleteUserKeyRequest
	GetOrganizationId() *string
}

type DeleteUserKeyRequest struct {
	// example:
	//
	// f0b1e61db5961df5975a93f9129d2513
	AccessToken *string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 5ebbc0228123212b59xxxxx
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
}

func (s DeleteUserKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserKeyRequest) GoString() string {
	return s.String()
}

func (s *DeleteUserKeyRequest) GetAccessToken() *string {
	return s.AccessToken
}

func (s *DeleteUserKeyRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *DeleteUserKeyRequest) SetAccessToken(v string) *DeleteUserKeyRequest {
	s.AccessToken = &v
	return s
}

func (s *DeleteUserKeyRequest) SetOrganizationId(v string) *DeleteUserKeyRequest {
	s.OrganizationId = &v
	return s
}

func (s *DeleteUserKeyRequest) Validate() error {
	return dara.Validate(s)
}
