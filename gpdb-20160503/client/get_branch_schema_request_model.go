// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBranchSchemaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBranchId(v string) *GetBranchSchemaRequest
	GetBranchId() *string
	SetDBName(v string) *GetBranchSchemaRequest
	GetDBName() *string
	SetProjectId(v string) *GetBranchSchemaRequest
	GetProjectId() *string
	SetRegionId(v string) *GetBranchSchemaRequest
	GetRegionId() *string
}

type GetBranchSchemaRequest struct {
	// The branch ID that uniquely identifies a Supabase branch.
	//
	// This parameter is required.
	//
	// example:
	//
	// br-xxxx
	BranchId *string `json:"BranchId,omitempty" xml:"BranchId,omitempty"`
	// The database name. The system databases postgres, template0, and template1 do not support schema queries.
	//
	// This parameter is required.
	//
	// example:
	//
	// neondb
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// The Supabase project ID that corresponds to the primary branch.
	//
	// This parameter is required.
	//
	// example:
	//
	// spb-xxxx
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The region ID. This parameter is required when you create a primary branch. When you create a sub-branch, the region of the primary branch is inherited by default.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetBranchSchemaRequest) String() string {
	return dara.Prettify(s)
}

func (s GetBranchSchemaRequest) GoString() string {
	return s.String()
}

func (s *GetBranchSchemaRequest) GetBranchId() *string {
	return s.BranchId
}

func (s *GetBranchSchemaRequest) GetDBName() *string {
	return s.DBName
}

func (s *GetBranchSchemaRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *GetBranchSchemaRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetBranchSchemaRequest) SetBranchId(v string) *GetBranchSchemaRequest {
	s.BranchId = &v
	return s
}

func (s *GetBranchSchemaRequest) SetDBName(v string) *GetBranchSchemaRequest {
	s.DBName = &v
	return s
}

func (s *GetBranchSchemaRequest) SetProjectId(v string) *GetBranchSchemaRequest {
	s.ProjectId = &v
	return s
}

func (s *GetBranchSchemaRequest) SetRegionId(v string) *GetBranchSchemaRequest {
	s.RegionId = &v
	return s
}

func (s *GetBranchSchemaRequest) Validate() error {
	return dara.Validate(s)
}
