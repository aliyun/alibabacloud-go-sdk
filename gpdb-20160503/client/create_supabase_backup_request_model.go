// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSupabaseBackupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProjectId(v string) *CreateSupabaseBackupRequest
	GetProjectId() *string
	SetRegionId(v string) *CreateSupabaseBackupRequest
	GetRegionId() *string
}

type CreateSupabaseBackupRequest struct {
	// Instance ID of the Supabase instance. You can obtain instance ID on the Supabase page in the console.
	//
	// This parameter is required.
	//
	// example:
	//
	// sbp-263****
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The region ID.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/86912.html) operation to query available region IDs.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateSupabaseBackupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSupabaseBackupRequest) GoString() string {
	return s.String()
}

func (s *CreateSupabaseBackupRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *CreateSupabaseBackupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateSupabaseBackupRequest) SetProjectId(v string) *CreateSupabaseBackupRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateSupabaseBackupRequest) SetRegionId(v string) *CreateSupabaseBackupRequest {
	s.RegionId = &v
	return s
}

func (s *CreateSupabaseBackupRequest) Validate() error {
	return dara.Validate(s)
}
