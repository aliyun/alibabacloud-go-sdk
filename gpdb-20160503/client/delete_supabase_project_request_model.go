// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSupabaseProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProjectId(v string) *DeleteSupabaseProjectRequest
	GetProjectId() *string
	SetRegionId(v string) *DeleteSupabaseProjectRequest
	GetRegionId() *string
}

type DeleteSupabaseProjectRequest struct {
	// The Supabase project ID. You can go to the Supabase page in the AnalyticDB for PostgreSQL console to obtain the workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// sbp-1****
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteSupabaseProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSupabaseProjectRequest) GoString() string {
	return s.String()
}

func (s *DeleteSupabaseProjectRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *DeleteSupabaseProjectRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteSupabaseProjectRequest) SetProjectId(v string) *DeleteSupabaseProjectRequest {
	s.ProjectId = &v
	return s
}

func (s *DeleteSupabaseProjectRequest) SetRegionId(v string) *DeleteSupabaseProjectRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteSupabaseProjectRequest) Validate() error {
	return dara.Validate(s)
}
