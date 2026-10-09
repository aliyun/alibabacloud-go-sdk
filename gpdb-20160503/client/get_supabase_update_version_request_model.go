// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSupabaseUpdateVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProjectId(v string) *GetSupabaseUpdateVersionRequest
	GetProjectId() *string
	SetRegionId(v string) *GetSupabaseUpdateVersionRequest
	GetRegionId() *string
}

type GetSupabaseUpdateVersionRequest struct {
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

func (s GetSupabaseUpdateVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSupabaseUpdateVersionRequest) GoString() string {
	return s.String()
}

func (s *GetSupabaseUpdateVersionRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *GetSupabaseUpdateVersionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetSupabaseUpdateVersionRequest) SetProjectId(v string) *GetSupabaseUpdateVersionRequest {
	s.ProjectId = &v
	return s
}

func (s *GetSupabaseUpdateVersionRequest) SetRegionId(v string) *GetSupabaseUpdateVersionRequest {
	s.RegionId = &v
	return s
}

func (s *GetSupabaseUpdateVersionRequest) Validate() error {
	return dara.Validate(s)
}
