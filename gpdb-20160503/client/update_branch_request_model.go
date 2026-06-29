// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBranchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBranchId(v string) *UpdateBranchRequest
	GetBranchId() *string
	SetBranchName(v string) *UpdateBranchRequest
	GetBranchName() *string
	SetClearExpiresAt(v bool) *UpdateBranchRequest
	GetClearExpiresAt() *bool
	SetDescription(v string) *UpdateBranchRequest
	GetDescription() *string
	SetExpiresAt(v string) *UpdateBranchRequest
	GetExpiresAt() *string
	SetProjectId(v string) *UpdateBranchRequest
	GetProjectId() *string
	SetProtected(v bool) *UpdateBranchRequest
	GetProtected() *bool
	SetRegionId(v string) *UpdateBranchRequest
	GetRegionId() *string
	SetTag(v []*UpdateBranchRequestTag) *UpdateBranchRequest
	GetTag() []*UpdateBranchRequestTag
}

type UpdateBranchRequest struct {
	// The branch ID that uniquely identifies a Supabase branch.
	//
	// This parameter is required.
	//
	// example:
	//
	// br-xxxx
	BranchId *string `json:"BranchId,omitempty" xml:"BranchId,omitempty"`
	// The branch name.
	//
	// This parameter is required.
	//
	// example:
	//
	// dev
	BranchName *string `json:"BranchName,omitempty" xml:"BranchName,omitempty"`
	// Specifies whether to clear the branch expiration time.
	//
	// Valid values:
	//
	// - true: Clears ExpiresAt.
	//
	// - false: Does not clear ExpiresAt.
	//
	// If this parameter is not specified, the existing expiration time remains unchanged.
	//
	// example:
	//
	// false
	ClearExpiresAt *bool `json:"ClearExpiresAt,omitempty" xml:"ClearExpiresAt,omitempty"`
	// The branch description.
	//
	// example:
	//
	// test branch
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time when the branch automatically expires and is deleted. The value is in ISO 8601 UTC format.
	//
	// example:
	//
	// 2026-10-08T09:11:12Z
	ExpiresAt *string `json:"ExpiresAt,omitempty" xml:"ExpiresAt,omitempty"`
	// The Supabase project ID that corresponds to the primary branch.
	//
	// This parameter is required.
	//
	// example:
	//
	// spb-xxxx
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// Specifies whether to enable branch protection.
	//
	// Valid values:
	//
	// - true: Branch protection is enabled.
	//
	// - false: Branch protection is disabled. This is the default value.
	//
	// example:
	//
	// false
	Protected *bool `json:"Protected,omitempty" xml:"Protected,omitempty"`
	// The region ID. This parameter is required when you create a primary branch. When you create a sub-branch, the region of the primary branch is inherited by default.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The list of branch tags.
	Tag []*UpdateBranchRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s UpdateBranchRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateBranchRequest) GoString() string {
	return s.String()
}

func (s *UpdateBranchRequest) GetBranchId() *string {
	return s.BranchId
}

func (s *UpdateBranchRequest) GetBranchName() *string {
	return s.BranchName
}

func (s *UpdateBranchRequest) GetClearExpiresAt() *bool {
	return s.ClearExpiresAt
}

func (s *UpdateBranchRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateBranchRequest) GetExpiresAt() *string {
	return s.ExpiresAt
}

func (s *UpdateBranchRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *UpdateBranchRequest) GetProtected() *bool {
	return s.Protected
}

func (s *UpdateBranchRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateBranchRequest) GetTag() []*UpdateBranchRequestTag {
	return s.Tag
}

func (s *UpdateBranchRequest) SetBranchId(v string) *UpdateBranchRequest {
	s.BranchId = &v
	return s
}

func (s *UpdateBranchRequest) SetBranchName(v string) *UpdateBranchRequest {
	s.BranchName = &v
	return s
}

func (s *UpdateBranchRequest) SetClearExpiresAt(v bool) *UpdateBranchRequest {
	s.ClearExpiresAt = &v
	return s
}

func (s *UpdateBranchRequest) SetDescription(v string) *UpdateBranchRequest {
	s.Description = &v
	return s
}

func (s *UpdateBranchRequest) SetExpiresAt(v string) *UpdateBranchRequest {
	s.ExpiresAt = &v
	return s
}

func (s *UpdateBranchRequest) SetProjectId(v string) *UpdateBranchRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateBranchRequest) SetProtected(v bool) *UpdateBranchRequest {
	s.Protected = &v
	return s
}

func (s *UpdateBranchRequest) SetRegionId(v string) *UpdateBranchRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateBranchRequest) SetTag(v []*UpdateBranchRequestTag) *UpdateBranchRequest {
	s.Tag = v
	return s
}

func (s *UpdateBranchRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateBranchRequestTag struct {
	// The tag key.
	//
	// example:
	//
	// env
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// dev
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateBranchRequestTag) String() string {
	return dara.Prettify(s)
}

func (s UpdateBranchRequestTag) GoString() string {
	return s.String()
}

func (s *UpdateBranchRequestTag) GetKey() *string {
	return s.Key
}

func (s *UpdateBranchRequestTag) GetValue() *string {
	return s.Value
}

func (s *UpdateBranchRequestTag) SetKey(v string) *UpdateBranchRequestTag {
	s.Key = &v
	return s
}

func (s *UpdateBranchRequestTag) SetValue(v string) *UpdateBranchRequestTag {
	s.Value = &v
	return s
}

func (s *UpdateBranchRequestTag) Validate() error {
	return dara.Validate(s)
}
