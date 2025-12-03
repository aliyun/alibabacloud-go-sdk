// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBranchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *CreateBranchRequest
	GetAccessToken() *string
	SetBranchName(v string) *CreateBranchRequest
	GetBranchName() *string
	SetRef(v string) *CreateBranchRequest
	GetRef() *string
	SetOrganizationId(v string) *CreateBranchRequest
	GetOrganizationId() *string
}

type CreateBranchRequest struct {
	// example:
	//
	// 0cf2c8458ac44d9481aab2dd6ec10596v3
	AccessToken *string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// createBranch
	BranchName *string `json:"branchName,omitempty" xml:"branchName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// master
	Ref *string `json:"ref,omitempty" xml:"ref,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 60de7a6852743a5162b5f957
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
}

func (s CreateBranchRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateBranchRequest) GoString() string {
	return s.String()
}

func (s *CreateBranchRequest) GetAccessToken() *string {
	return s.AccessToken
}

func (s *CreateBranchRequest) GetBranchName() *string {
	return s.BranchName
}

func (s *CreateBranchRequest) GetRef() *string {
	return s.Ref
}

func (s *CreateBranchRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *CreateBranchRequest) SetAccessToken(v string) *CreateBranchRequest {
	s.AccessToken = &v
	return s
}

func (s *CreateBranchRequest) SetBranchName(v string) *CreateBranchRequest {
	s.BranchName = &v
	return s
}

func (s *CreateBranchRequest) SetRef(v string) *CreateBranchRequest {
	s.Ref = &v
	return s
}

func (s *CreateBranchRequest) SetOrganizationId(v string) *CreateBranchRequest {
	s.OrganizationId = &v
	return s
}

func (s *CreateBranchRequest) Validate() error {
	return dara.Validate(s)
}
