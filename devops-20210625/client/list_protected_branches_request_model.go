// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProtectedBranchesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *ListProtectedBranchesRequest
	GetAccessToken() *string
	SetOrganizationId(v string) *ListProtectedBranchesRequest
	GetOrganizationId() *string
}

type ListProtectedBranchesRequest struct {
	// example:
	//
	// f0b1e61db5961df5975a93f9129d2513
	AccessToken *string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 60de7a6852743a5162b5f957
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
}

func (s ListProtectedBranchesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListProtectedBranchesRequest) GoString() string {
	return s.String()
}

func (s *ListProtectedBranchesRequest) GetAccessToken() *string {
	return s.AccessToken
}

func (s *ListProtectedBranchesRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *ListProtectedBranchesRequest) SetAccessToken(v string) *ListProtectedBranchesRequest {
	s.AccessToken = &v
	return s
}

func (s *ListProtectedBranchesRequest) SetOrganizationId(v string) *ListProtectedBranchesRequest {
	s.OrganizationId = &v
	return s
}

func (s *ListProtectedBranchesRequest) Validate() error {
	return dara.Validate(s)
}
