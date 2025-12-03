// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBranchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *DeleteBranchRequest
	GetAccessToken() *string
	SetBranchName(v string) *DeleteBranchRequest
	GetBranchName() *string
	SetOrganizationId(v string) *DeleteBranchRequest
	GetOrganizationId() *string
}

type DeleteBranchRequest struct {
	// example:
	//
	// f0b1e61db5961df5975a93f9129d2513
	AccessToken *string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// deleteBranch
	BranchName *string `json:"branchName,omitempty" xml:"branchName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 619b80042f595dbd1b9b0de2
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
}

func (s DeleteBranchRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteBranchRequest) GoString() string {
	return s.String()
}

func (s *DeleteBranchRequest) GetAccessToken() *string {
	return s.AccessToken
}

func (s *DeleteBranchRequest) GetBranchName() *string {
	return s.BranchName
}

func (s *DeleteBranchRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *DeleteBranchRequest) SetAccessToken(v string) *DeleteBranchRequest {
	s.AccessToken = &v
	return s
}

func (s *DeleteBranchRequest) SetBranchName(v string) *DeleteBranchRequest {
	s.BranchName = &v
	return s
}

func (s *DeleteBranchRequest) SetOrganizationId(v string) *DeleteBranchRequest {
	s.OrganizationId = &v
	return s
}

func (s *DeleteBranchRequest) Validate() error {
	return dara.Validate(s)
}
