// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySupabaseProjectDescriptionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProjectDescription(v string) *ModifySupabaseProjectDescriptionRequest
	GetProjectDescription() *string
	SetProjectId(v string) *ModifySupabaseProjectDescriptionRequest
	GetProjectId() *string
	SetRegionId(v string) *ModifySupabaseProjectDescriptionRequest
	GetRegionId() *string
}

type ModifySupabaseProjectDescriptionRequest struct {
	// The detailed description of the Supabase project.
	//
	// example:
	//
	// for-test-project
	ProjectDescription *string `json:"ProjectDescription,omitempty" xml:"ProjectDescription,omitempty"`
	// The Supabase instance ID.
	//
	// example:
	//
	// sbp-twmoe9bakow
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifySupabaseProjectDescriptionRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifySupabaseProjectDescriptionRequest) GoString() string {
	return s.String()
}

func (s *ModifySupabaseProjectDescriptionRequest) GetProjectDescription() *string {
	return s.ProjectDescription
}

func (s *ModifySupabaseProjectDescriptionRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *ModifySupabaseProjectDescriptionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifySupabaseProjectDescriptionRequest) SetProjectDescription(v string) *ModifySupabaseProjectDescriptionRequest {
	s.ProjectDescription = &v
	return s
}

func (s *ModifySupabaseProjectDescriptionRequest) SetProjectId(v string) *ModifySupabaseProjectDescriptionRequest {
	s.ProjectId = &v
	return s
}

func (s *ModifySupabaseProjectDescriptionRequest) SetRegionId(v string) *ModifySupabaseProjectDescriptionRequest {
	s.RegionId = &v
	return s
}

func (s *ModifySupabaseProjectDescriptionRequest) Validate() error {
	return dara.Validate(s)
}
