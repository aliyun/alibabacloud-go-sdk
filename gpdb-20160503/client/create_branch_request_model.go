// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBranchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBranchName(v string) *CreateBranchRequest
	GetBranchName() *string
	SetClientToken(v string) *CreateBranchRequest
	GetClientToken() *string
	SetDescription(v string) *CreateBranchRequest
	GetDescription() *string
	SetExpiresAt(v string) *CreateBranchRequest
	GetExpiresAt() *string
	SetInitSource(v string) *CreateBranchRequest
	GetInitSource() *string
	SetParentBranchId(v string) *CreateBranchRequest
	GetParentBranchId() *string
	SetParentLsn(v string) *CreateBranchRequest
	GetParentLsn() *string
	SetParentTimestamp(v string) *CreateBranchRequest
	GetParentTimestamp() *string
	SetProjectId(v string) *CreateBranchRequest
	GetProjectId() *string
	SetProtected(v bool) *CreateBranchRequest
	GetProtected() *bool
	SetRegionId(v string) *CreateBranchRequest
	GetRegionId() *string
	SetTag(v []*CreateBranchRequestTag) *CreateBranchRequest
	GetTag() []*CreateBranchRequestTag
}

type CreateBranchRequest struct {
	// The branch name.
	//
	// This parameter is required.
	//
	// example:
	//
	// dev
	BranchName *string `json:"BranchName,omitempty" xml:"BranchName,omitempty"`
	// The client idempotency token. This token ensures the idempotence of retry requests.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426614174000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The branch description.
	//
	// example:
	//
	// test branch
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time at which the branch automatically expires and is deleted. The value is in ISO 8601 UTC format.
	//
	// example:
	//
	// 2026-10-08T09:11:12Z
	ExpiresAt *string `json:"ExpiresAt,omitempty" xml:"ExpiresAt,omitempty"`
	// The initialization source of the branch.
	//
	// Valid values:
	//
	// - ParentData: copies the schema and data from the parent branch. This is the default value.
	//
	// - SchemaOnly: copies only the schema structure.
	//
	// example:
	//
	// ParentData
	InitSource *string `json:"InitSource,omitempty" xml:"InitSource,omitempty"`
	// The parent branch ID. This parameter specifies the parent branch for the new branch or query condition.
	//
	// This parameter is required.
	//
	// example:
	//
	// br-main
	ParentBranchId *string `json:"ParentBranchId,omitempty" xml:"ParentBranchId,omitempty"`
	// The log sequence number (LSN) from the parent branch at which the branch is created.
	//
	// example:
	//
	// 0/3522648
	ParentLsn *string `json:"ParentLsn,omitempty" xml:"ParentLsn,omitempty"`
	// The point in time for data synchronization from the parent branch when creating the branch. The value is in ISO 8601 UTC format.
	//
	// Default value: the current time.
	//
	// example:
	//
	// 2026-04-08T09:11:12Z
	ParentTimestamp *string `json:"ParentTimestamp,omitempty" xml:"ParentTimestamp,omitempty"`
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
	// - true: Enables branch protection.
	//
	// - false: Disables branch protection. This is the default value.
	//
	// example:
	//
	// false
	Protected *bool `json:"Protected,omitempty" xml:"Protected,omitempty"`
	// The region ID. This parameter is required when you create a primary branch. When you create a child branch, the region is inherited from the primary branch by default.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The list of branch tags.
	Tag []*CreateBranchRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreateBranchRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateBranchRequest) GoString() string {
	return s.String()
}

func (s *CreateBranchRequest) GetBranchName() *string {
	return s.BranchName
}

func (s *CreateBranchRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateBranchRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateBranchRequest) GetExpiresAt() *string {
	return s.ExpiresAt
}

func (s *CreateBranchRequest) GetInitSource() *string {
	return s.InitSource
}

func (s *CreateBranchRequest) GetParentBranchId() *string {
	return s.ParentBranchId
}

func (s *CreateBranchRequest) GetParentLsn() *string {
	return s.ParentLsn
}

func (s *CreateBranchRequest) GetParentTimestamp() *string {
	return s.ParentTimestamp
}

func (s *CreateBranchRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *CreateBranchRequest) GetProtected() *bool {
	return s.Protected
}

func (s *CreateBranchRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateBranchRequest) GetTag() []*CreateBranchRequestTag {
	return s.Tag
}

func (s *CreateBranchRequest) SetBranchName(v string) *CreateBranchRequest {
	s.BranchName = &v
	return s
}

func (s *CreateBranchRequest) SetClientToken(v string) *CreateBranchRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateBranchRequest) SetDescription(v string) *CreateBranchRequest {
	s.Description = &v
	return s
}

func (s *CreateBranchRequest) SetExpiresAt(v string) *CreateBranchRequest {
	s.ExpiresAt = &v
	return s
}

func (s *CreateBranchRequest) SetInitSource(v string) *CreateBranchRequest {
	s.InitSource = &v
	return s
}

func (s *CreateBranchRequest) SetParentBranchId(v string) *CreateBranchRequest {
	s.ParentBranchId = &v
	return s
}

func (s *CreateBranchRequest) SetParentLsn(v string) *CreateBranchRequest {
	s.ParentLsn = &v
	return s
}

func (s *CreateBranchRequest) SetParentTimestamp(v string) *CreateBranchRequest {
	s.ParentTimestamp = &v
	return s
}

func (s *CreateBranchRequest) SetProjectId(v string) *CreateBranchRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateBranchRequest) SetProtected(v bool) *CreateBranchRequest {
	s.Protected = &v
	return s
}

func (s *CreateBranchRequest) SetRegionId(v string) *CreateBranchRequest {
	s.RegionId = &v
	return s
}

func (s *CreateBranchRequest) SetTag(v []*CreateBranchRequestTag) *CreateBranchRequest {
	s.Tag = v
	return s
}

func (s *CreateBranchRequest) Validate() error {
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

type CreateBranchRequestTag struct {
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

func (s CreateBranchRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateBranchRequestTag) GoString() string {
	return s.String()
}

func (s *CreateBranchRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateBranchRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateBranchRequestTag) SetKey(v string) *CreateBranchRequestTag {
	s.Key = &v
	return s
}

func (s *CreateBranchRequestTag) SetValue(v string) *CreateBranchRequestTag {
	s.Value = &v
	return s
}

func (s *CreateBranchRequestTag) Validate() error {
	return dara.Validate(s)
}
