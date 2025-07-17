// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *DeleteTaskRequest
	GetId() *int64
	SetProjectEnv(v string) *DeleteTaskRequest
	GetProjectEnv() *string
}

type DeleteTaskRequest struct {
	// The task ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The environment of the workspace.
	//
	// Valid values:
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

func (s DeleteTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteTaskRequest) GoString() string {
	return s.String()
}

func (s *DeleteTaskRequest) GetId() *int64 {
	return s.Id
}

func (s *DeleteTaskRequest) GetProjectEnv() *string {
	return s.ProjectEnv
}

func (s *DeleteTaskRequest) SetId(v int64) *DeleteTaskRequest {
	s.Id = &v
	return s
}

func (s *DeleteTaskRequest) SetProjectEnv(v string) *DeleteTaskRequest {
	s.ProjectEnv = &v
	return s
}

func (s *DeleteTaskRequest) Validate() error {
	return dara.Validate(s)
}
