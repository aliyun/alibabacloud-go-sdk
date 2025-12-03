// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRepositoryGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *DeleteRepositoryGroupRequest
	GetAccessToken() *string
	SetReason(v string) *DeleteRepositoryGroupRequest
	GetReason() *string
	SetOrganizationId(v string) *DeleteRepositoryGroupRequest
	GetOrganizationId() *string
}

type DeleteRepositoryGroupRequest struct {
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
	// 624666bd54d036291ae13a36
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
}

func (s DeleteRepositoryGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRepositoryGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteRepositoryGroupRequest) GetAccessToken() *string {
	return s.AccessToken
}

func (s *DeleteRepositoryGroupRequest) GetReason() *string {
	return s.Reason
}

func (s *DeleteRepositoryGroupRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *DeleteRepositoryGroupRequest) SetAccessToken(v string) *DeleteRepositoryGroupRequest {
	s.AccessToken = &v
	return s
}

func (s *DeleteRepositoryGroupRequest) SetReason(v string) *DeleteRepositoryGroupRequest {
	s.Reason = &v
	return s
}

func (s *DeleteRepositoryGroupRequest) SetOrganizationId(v string) *DeleteRepositoryGroupRequest {
	s.OrganizationId = &v
	return s
}

func (s *DeleteRepositoryGroupRequest) Validate() error {
	return dara.Validate(s)
}
