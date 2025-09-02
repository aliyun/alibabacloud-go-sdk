// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFileTypeStatisticRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProjectEnv(v string) *GetFileTypeStatisticRequest
	GetProjectEnv() *string
	SetProjectId(v int64) *GetFileTypeStatisticRequest
	GetProjectId() *int64
}

type GetFileTypeStatisticRequest struct {
	// The environment of the workspace. Valid values: PROD and DEV. The value PROD indicates the production environment. The value DEV indicates the development environment.
	//
	// This parameter is required.
	//
	// example:
	//
	// PROD
	ProjectEnv *string `json:"ProjectEnv,omitempty" xml:"ProjectEnv,omitempty"`
	// The ID of the DataWorks workspace. You can log on to the DataWorks console and go to the Workspace Management page to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123465
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GetFileTypeStatisticRequest) String() string {
	return dara.Prettify(s)
}

func (s GetFileTypeStatisticRequest) GoString() string {
	return s.String()
}

func (s *GetFileTypeStatisticRequest) GetProjectEnv() *string {
	return s.ProjectEnv
}

func (s *GetFileTypeStatisticRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetFileTypeStatisticRequest) SetProjectEnv(v string) *GetFileTypeStatisticRequest {
	s.ProjectEnv = &v
	return s
}

func (s *GetFileTypeStatisticRequest) SetProjectId(v int64) *GetFileTypeStatisticRequest {
	s.ProjectId = &v
	return s
}

func (s *GetFileTypeStatisticRequest) Validate() error {
	return dara.Validate(s)
}
