// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSupabaseProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProjectId(v string) *GetSupabaseProjectRequest
	GetProjectId() *string
	SetRegionId(v string) *GetSupabaseProjectRequest
	GetRegionId() *string
}

type GetSupabaseProjectRequest struct {
	// The ID of the Supabase instance. You can obtain the ID on the Supabase page in the console.
	//
	// This parameter is required.
	//
	// example:
	//
	// sbp-263****
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The region ID of the cluster.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/86912.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetSupabaseProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSupabaseProjectRequest) GoString() string {
	return s.String()
}

func (s *GetSupabaseProjectRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *GetSupabaseProjectRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetSupabaseProjectRequest) SetProjectId(v string) *GetSupabaseProjectRequest {
	s.ProjectId = &v
	return s
}

func (s *GetSupabaseProjectRequest) SetRegionId(v string) *GetSupabaseProjectRequest {
	s.RegionId = &v
	return s
}

func (s *GetSupabaseProjectRequest) Validate() error {
	return dara.Validate(s)
}
