// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRepositoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *DeleteRepositoryRequest
	GetAccessToken() *string
	SetReason(v string) *DeleteRepositoryRequest
	GetReason() *string
	SetOrganizationId(v string) *DeleteRepositoryRequest
	GetOrganizationId() *string
}

type DeleteRepositoryRequest struct {
	// example:
	//
	// f0b1e61db5961df5975a93f9129d2513
	AccessToken *string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	// This parameter is required.
	Reason *string `json:"reason,omitempty" xml:"reason,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 6082a9b0c7972588ac363793
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
}

func (s DeleteRepositoryRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRepositoryRequest) GoString() string {
	return s.String()
}

func (s *DeleteRepositoryRequest) GetAccessToken() *string {
	return s.AccessToken
}

func (s *DeleteRepositoryRequest) GetReason() *string {
	return s.Reason
}

func (s *DeleteRepositoryRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *DeleteRepositoryRequest) SetAccessToken(v string) *DeleteRepositoryRequest {
	s.AccessToken = &v
	return s
}

func (s *DeleteRepositoryRequest) SetReason(v string) *DeleteRepositoryRequest {
	s.Reason = &v
	return s
}

func (s *DeleteRepositoryRequest) SetOrganizationId(v string) *DeleteRepositoryRequest {
	s.OrganizationId = &v
	return s
}

func (s *DeleteRepositoryRequest) Validate() error {
	return dara.Validate(s)
}
