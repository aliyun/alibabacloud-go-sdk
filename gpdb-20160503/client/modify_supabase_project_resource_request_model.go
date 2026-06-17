// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySupabaseProjectResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetModifyType(v string) *ModifySupabaseProjectResourceRequest
	GetModifyType() *string
	SetProjectId(v string) *ModifySupabaseProjectResourceRequest
	GetProjectId() *string
	SetProjectSpec(v string) *ModifySupabaseProjectResourceRequest
	GetProjectSpec() *string
	SetRegionId(v string) *ModifySupabaseProjectResourceRequest
	GetRegionId() *string
	SetStorageSize(v int64) *ModifySupabaseProjectResourceRequest
	GetStorageSize() *int64
}

type ModifySupabaseProjectResourceRequest struct {
	// The modification type.
	//
	// This parameter is required.
	//
	// example:
	//
	// UPGRADE,DOWNGRADE
	ModifyType *string `json:"ModifyType,omitempty" xml:"ModifyType,omitempty"`
	// The Supabase project ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// sbp-tyarplz****
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The new project specifications.
	//
	// example:
	//
	// 2C4G
	ProjectSpec *string `json:"ProjectSpec,omitempty" xml:"ProjectSpec,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/86912.html) operation to query available region IDs.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The storage size, in GB.
	//
	// example:
	//
	// 100
	StorageSize *int64 `json:"StorageSize,omitempty" xml:"StorageSize,omitempty"`
}

func (s ModifySupabaseProjectResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifySupabaseProjectResourceRequest) GoString() string {
	return s.String()
}

func (s *ModifySupabaseProjectResourceRequest) GetModifyType() *string {
	return s.ModifyType
}

func (s *ModifySupabaseProjectResourceRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *ModifySupabaseProjectResourceRequest) GetProjectSpec() *string {
	return s.ProjectSpec
}

func (s *ModifySupabaseProjectResourceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifySupabaseProjectResourceRequest) GetStorageSize() *int64 {
	return s.StorageSize
}

func (s *ModifySupabaseProjectResourceRequest) SetModifyType(v string) *ModifySupabaseProjectResourceRequest {
	s.ModifyType = &v
	return s
}

func (s *ModifySupabaseProjectResourceRequest) SetProjectId(v string) *ModifySupabaseProjectResourceRequest {
	s.ProjectId = &v
	return s
}

func (s *ModifySupabaseProjectResourceRequest) SetProjectSpec(v string) *ModifySupabaseProjectResourceRequest {
	s.ProjectSpec = &v
	return s
}

func (s *ModifySupabaseProjectResourceRequest) SetRegionId(v string) *ModifySupabaseProjectResourceRequest {
	s.RegionId = &v
	return s
}

func (s *ModifySupabaseProjectResourceRequest) SetStorageSize(v int64) *ModifySupabaseProjectResourceRequest {
	s.StorageSize = &v
	return s
}

func (s *ModifySupabaseProjectResourceRequest) Validate() error {
	return dara.Validate(s)
}
