// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSupabaseVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMinorVersion(v string) *UpdateSupabaseVersionRequest
	GetMinorVersion() *string
	SetProjectId(v string) *UpdateSupabaseVersionRequest
	GetProjectId() *string
	SetRegionId(v string) *UpdateSupabaseVersionRequest
	GetRegionId() *string
}

type UpdateSupabaseVersionRequest struct {
	// The target minor version. You can query the supported upgrade versions for the current project by calling GetSupabaseUpdateVersion.
	//
	// example:
	//
	// 20240731
	MinorVersion *string `json:"MinorVersion,omitempty" xml:"MinorVersion,omitempty"`
	// The ID of the Supabase project.
	//
	// This parameter is required.
	//
	// example:
	//
	// spb-xxxx
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateSupabaseVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSupabaseVersionRequest) GoString() string {
	return s.String()
}

func (s *UpdateSupabaseVersionRequest) GetMinorVersion() *string {
	return s.MinorVersion
}

func (s *UpdateSupabaseVersionRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *UpdateSupabaseVersionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateSupabaseVersionRequest) SetMinorVersion(v string) *UpdateSupabaseVersionRequest {
	s.MinorVersion = &v
	return s
}

func (s *UpdateSupabaseVersionRequest) SetProjectId(v string) *UpdateSupabaseVersionRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateSupabaseVersionRequest) SetRegionId(v string) *UpdateSupabaseVersionRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateSupabaseVersionRequest) Validate() error {
	return dara.Validate(s)
}
