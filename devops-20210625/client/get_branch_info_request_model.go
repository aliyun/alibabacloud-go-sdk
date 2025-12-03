// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBranchInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *GetBranchInfoRequest
	GetAccessToken() *string
	SetBranchName(v string) *GetBranchInfoRequest
	GetBranchName() *string
	SetOrganizationId(v string) *GetBranchInfoRequest
	GetOrganizationId() *string
}

type GetBranchInfoRequest struct {
	// example:
	//
	// f0b1e61db5961df5975a93f9129d2513
	AccessToken *string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// master
	BranchName *string `json:"branchName,omitempty" xml:"branchName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 5fbe3118672533690be72b12
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
}

func (s GetBranchInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetBranchInfoRequest) GoString() string {
	return s.String()
}

func (s *GetBranchInfoRequest) GetAccessToken() *string {
	return s.AccessToken
}

func (s *GetBranchInfoRequest) GetBranchName() *string {
	return s.BranchName
}

func (s *GetBranchInfoRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *GetBranchInfoRequest) SetAccessToken(v string) *GetBranchInfoRequest {
	s.AccessToken = &v
	return s
}

func (s *GetBranchInfoRequest) SetBranchName(v string) *GetBranchInfoRequest {
	s.BranchName = &v
	return s
}

func (s *GetBranchInfoRequest) SetOrganizationId(v string) *GetBranchInfoRequest {
	s.OrganizationId = &v
	return s
}

func (s *GetBranchInfoRequest) Validate() error {
	return dara.Validate(s)
}
