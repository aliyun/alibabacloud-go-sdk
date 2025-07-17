// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *GetTaskRequest
	GetId() *int64
	SetProjectEnv(v string) *GetTaskRequest
	GetProjectEnv() *string
}

type GetTaskRequest struct {
	// The task ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The environment of the workspace. Valid values:
	//
	// 	- Prod: production environment
	//
	// 	- Dev: development environment
	//
	// example:
	//
	// Prod
	ProjectEnv *string `json:"ProjectEnv,omitempty" xml:"ProjectEnv,omitempty"`
}

func (s GetTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTaskRequest) GoString() string {
	return s.String()
}

func (s *GetTaskRequest) GetId() *int64 {
	return s.Id
}

func (s *GetTaskRequest) GetProjectEnv() *string {
	return s.ProjectEnv
}

func (s *GetTaskRequest) SetId(v int64) *GetTaskRequest {
	s.Id = &v
	return s
}

func (s *GetTaskRequest) SetProjectEnv(v string) *GetTaskRequest {
	s.ProjectEnv = &v
	return s
}

func (s *GetTaskRequest) Validate() error {
	return dara.Validate(s)
}
