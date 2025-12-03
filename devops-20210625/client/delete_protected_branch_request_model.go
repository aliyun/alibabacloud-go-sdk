// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteProtectedBranchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *DeleteProtectedBranchRequest
	GetAccessToken() *string
	SetOrganizationId(v string) *DeleteProtectedBranchRequest
	GetOrganizationId() *string
}

type DeleteProtectedBranchRequest struct {
	// example:
	//
	// f0b1e61db5961df5975a93f9129d2513
	AccessToken *string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 611b75680fc7bf0dbe1dce55
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
}

func (s DeleteProtectedBranchRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteProtectedBranchRequest) GoString() string {
	return s.String()
}

func (s *DeleteProtectedBranchRequest) GetAccessToken() *string {
	return s.AccessToken
}

func (s *DeleteProtectedBranchRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *DeleteProtectedBranchRequest) SetAccessToken(v string) *DeleteProtectedBranchRequest {
	s.AccessToken = &v
	return s
}

func (s *DeleteProtectedBranchRequest) SetOrganizationId(v string) *DeleteProtectedBranchRequest {
	s.OrganizationId = &v
	return s
}

func (s *DeleteProtectedBranchRequest) Validate() error {
	return dara.Validate(s)
}
