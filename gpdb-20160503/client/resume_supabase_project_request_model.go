// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumeSupabaseProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProjectId(v string) *ResumeSupabaseProjectRequest
	GetProjectId() *string
	SetRegionId(v string) *ResumeSupabaseProjectRequest
	GetRegionId() *string
}

type ResumeSupabaseProjectRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// sbp-1****
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// Region ID
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ResumeSupabaseProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s ResumeSupabaseProjectRequest) GoString() string {
	return s.String()
}

func (s *ResumeSupabaseProjectRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *ResumeSupabaseProjectRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ResumeSupabaseProjectRequest) SetProjectId(v string) *ResumeSupabaseProjectRequest {
	s.ProjectId = &v
	return s
}

func (s *ResumeSupabaseProjectRequest) SetRegionId(v string) *ResumeSupabaseProjectRequest {
	s.RegionId = &v
	return s
}

func (s *ResumeSupabaseProjectRequest) Validate() error {
	return dara.Validate(s)
}
