// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPptInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *GetPptInfoRequest
	GetTaskId() *string
	SetWorkspaceId(v string) *GetPptInfoRequest
	GetWorkspaceId() *string
}

type GetPptInfoRequest struct {
	// example:
	//
	// 1f178f22-ec52-467d-8489-eef4468x0240
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// llm-2setzb9xb8mx6vss
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetPptInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPptInfoRequest) GoString() string {
	return s.String()
}

func (s *GetPptInfoRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetPptInfoRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetPptInfoRequest) SetTaskId(v string) *GetPptInfoRequest {
	s.TaskId = &v
	return s
}

func (s *GetPptInfoRequest) SetWorkspaceId(v string) *GetPptInfoRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetPptInfoRequest) Validate() error {
	return dara.Validate(s)
}
