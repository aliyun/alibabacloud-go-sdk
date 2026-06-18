// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartSupabaseProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProjectId(v string) *RestartSupabaseProjectRequest
	GetProjectId() *string
	SetRegionId(v string) *RestartSupabaseProjectRequest
	GetRegionId() *string
}

type RestartSupabaseProjectRequest struct {
	// The Supabase project ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// spb-xxxx
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The region ID. Specifies the region in which to query or perform the operation.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s RestartSupabaseProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s RestartSupabaseProjectRequest) GoString() string {
	return s.String()
}

func (s *RestartSupabaseProjectRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *RestartSupabaseProjectRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RestartSupabaseProjectRequest) SetProjectId(v string) *RestartSupabaseProjectRequest {
	s.ProjectId = &v
	return s
}

func (s *RestartSupabaseProjectRequest) SetRegionId(v string) *RestartSupabaseProjectRequest {
	s.RegionId = &v
	return s
}

func (s *RestartSupabaseProjectRequest) Validate() error {
	return dara.Validate(s)
}
