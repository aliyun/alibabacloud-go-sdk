// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateChangeRequestRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppCodeRepoSn(v string) *CreateChangeRequestRequest
	GetAppCodeRepoSn() *string
	SetAutoDeleteBranchWhenEnd(v bool) *CreateChangeRequestRequest
	GetAutoDeleteBranchWhenEnd() *bool
	SetBranchName(v string) *CreateChangeRequestRequest
	GetBranchName() *string
	SetCreateBranch(v bool) *CreateChangeRequestRequest
	GetCreateBranch() *bool
	SetOwnerAccountId(v string) *CreateChangeRequestRequest
	GetOwnerAccountId() *string
	SetOwnerId(v string) *CreateChangeRequestRequest
	GetOwnerId() *string
	SetTitle(v string) *CreateChangeRequestRequest
	GetTitle() *string
	SetOrganizationId(v string) *CreateChangeRequestRequest
	GetOrganizationId() *string
}

type CreateChangeRequestRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// sn123
	AppCodeRepoSn *string `json:"appCodeRepoSn,omitempty" xml:"appCodeRepoSn,omitempty"`
	// example:
	//
	// false
	AutoDeleteBranchWhenEnd *bool `json:"autoDeleteBranchWhenEnd,omitempty" xml:"autoDeleteBranchWhenEnd,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// hotfix/20240524
	BranchName *string `json:"branchName,omitempty" xml:"branchName,omitempty"`
	// example:
	//
	// false
	CreateBranch *bool `json:"createBranch,omitempty" xml:"createBranch,omitempty"`
	// example:
	//
	// 1332695887xxxxxx
	OwnerAccountId *string `json:"ownerAccountId,omitempty" xml:"ownerAccountId,omitempty"`
	// This parameter is required.
	OwnerId *string `json:"ownerId,omitempty" xml:"ownerId,omitempty"`
	// This parameter is required.
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// example:
	//
	// 66c0c9fffeb86b450c199fcd
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
}

func (s CreateChangeRequestRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateChangeRequestRequest) GoString() string {
	return s.String()
}

func (s *CreateChangeRequestRequest) GetAppCodeRepoSn() *string {
	return s.AppCodeRepoSn
}

func (s *CreateChangeRequestRequest) GetAutoDeleteBranchWhenEnd() *bool {
	return s.AutoDeleteBranchWhenEnd
}

func (s *CreateChangeRequestRequest) GetBranchName() *string {
	return s.BranchName
}

func (s *CreateChangeRequestRequest) GetCreateBranch() *bool {
	return s.CreateBranch
}

func (s *CreateChangeRequestRequest) GetOwnerAccountId() *string {
	return s.OwnerAccountId
}

func (s *CreateChangeRequestRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *CreateChangeRequestRequest) GetTitle() *string {
	return s.Title
}

func (s *CreateChangeRequestRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *CreateChangeRequestRequest) SetAppCodeRepoSn(v string) *CreateChangeRequestRequest {
	s.AppCodeRepoSn = &v
	return s
}

func (s *CreateChangeRequestRequest) SetAutoDeleteBranchWhenEnd(v bool) *CreateChangeRequestRequest {
	s.AutoDeleteBranchWhenEnd = &v
	return s
}

func (s *CreateChangeRequestRequest) SetBranchName(v string) *CreateChangeRequestRequest {
	s.BranchName = &v
	return s
}

func (s *CreateChangeRequestRequest) SetCreateBranch(v bool) *CreateChangeRequestRequest {
	s.CreateBranch = &v
	return s
}

func (s *CreateChangeRequestRequest) SetOwnerAccountId(v string) *CreateChangeRequestRequest {
	s.OwnerAccountId = &v
	return s
}

func (s *CreateChangeRequestRequest) SetOwnerId(v string) *CreateChangeRequestRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateChangeRequestRequest) SetTitle(v string) *CreateChangeRequestRequest {
	s.Title = &v
	return s
}

func (s *CreateChangeRequestRequest) SetOrganizationId(v string) *CreateChangeRequestRequest {
	s.OrganizationId = &v
	return s
}

func (s *CreateChangeRequestRequest) Validate() error {
	return dara.Validate(s)
}
