// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitiatePptCreationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOutline(v string) *InitiatePptCreationRequest
	GetOutline() *string
	SetTaskId(v string) *InitiatePptCreationRequest
	GetTaskId() *string
	SetWorkspaceId(v string) *InitiatePptCreationRequest
	GetWorkspaceId() *string
}

type InitiatePptCreationRequest struct {
	Outline *string `json:"Outline,omitempty" xml:"Outline,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 95c2fbe6-5a20-4fc2-8a93-376ed05fbe13
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// llm-3fy94b2rtadt01qa
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s InitiatePptCreationRequest) String() string {
	return dara.Prettify(s)
}

func (s InitiatePptCreationRequest) GoString() string {
	return s.String()
}

func (s *InitiatePptCreationRequest) GetOutline() *string {
	return s.Outline
}

func (s *InitiatePptCreationRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *InitiatePptCreationRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *InitiatePptCreationRequest) SetOutline(v string) *InitiatePptCreationRequest {
	s.Outline = &v
	return s
}

func (s *InitiatePptCreationRequest) SetTaskId(v string) *InitiatePptCreationRequest {
	s.TaskId = &v
	return s
}

func (s *InitiatePptCreationRequest) SetWorkspaceId(v string) *InitiatePptCreationRequest {
	s.WorkspaceId = &v
	return s
}

func (s *InitiatePptCreationRequest) Validate() error {
	return dara.Validate(s)
}
