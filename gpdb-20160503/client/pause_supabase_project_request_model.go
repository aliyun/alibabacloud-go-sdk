// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPauseSupabaseProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProjectId(v string) *PauseSupabaseProjectRequest
	GetProjectId() *string
	SetRegionId(v string) *PauseSupabaseProjectRequest
	GetRegionId() *string
}

type PauseSupabaseProjectRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// sbp-180****
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s PauseSupabaseProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s PauseSupabaseProjectRequest) GoString() string {
	return s.String()
}

func (s *PauseSupabaseProjectRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *PauseSupabaseProjectRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *PauseSupabaseProjectRequest) SetProjectId(v string) *PauseSupabaseProjectRequest {
	s.ProjectId = &v
	return s
}

func (s *PauseSupabaseProjectRequest) SetRegionId(v string) *PauseSupabaseProjectRequest {
	s.RegionId = &v
	return s
}

func (s *PauseSupabaseProjectRequest) Validate() error {
	return dara.Validate(s)
}
