// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProgramTypeCountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProjectEnv(v string) *ListProgramTypeCountRequest
	GetProjectEnv() *string
	SetProjectId(v int64) *ListProgramTypeCountRequest
	GetProjectId() *int64
}

type ListProgramTypeCountRequest struct {
	// The environment of the workspace. Valid values: PROD and DEV.
	//
	// This parameter is required.
	//
	// example:
	//
	// PROD
	ProjectEnv *string `json:"ProjectEnv,omitempty" xml:"ProjectEnv,omitempty"`
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ListProgramTypeCountRequest) String() string {
	return dara.Prettify(s)
}

func (s ListProgramTypeCountRequest) GoString() string {
	return s.String()
}

func (s *ListProgramTypeCountRequest) GetProjectEnv() *string {
	return s.ProjectEnv
}

func (s *ListProgramTypeCountRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListProgramTypeCountRequest) SetProjectEnv(v string) *ListProgramTypeCountRequest {
	s.ProjectEnv = &v
	return s
}

func (s *ListProgramTypeCountRequest) SetProjectId(v int64) *ListProgramTypeCountRequest {
	s.ProjectId = &v
	return s
}

func (s *ListProgramTypeCountRequest) Validate() error {
	return dara.Validate(s)
}
