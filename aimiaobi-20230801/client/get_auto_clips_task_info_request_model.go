// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAutoClipsTaskInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *GetAutoClipsTaskInfoRequest
	GetTaskId() *string
	SetWorkspaceId(v string) *GetAutoClipsTaskInfoRequest
	GetWorkspaceId() *string
}

type GetAutoClipsTaskInfoRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 0dbf1055f8a2475d99904c3b76a0ffba
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// This parameter is required.
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetAutoClipsTaskInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAutoClipsTaskInfoRequest) GoString() string {
	return s.String()
}

func (s *GetAutoClipsTaskInfoRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetAutoClipsTaskInfoRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetAutoClipsTaskInfoRequest) SetTaskId(v string) *GetAutoClipsTaskInfoRequest {
	s.TaskId = &v
	return s
}

func (s *GetAutoClipsTaskInfoRequest) SetWorkspaceId(v string) *GetAutoClipsTaskInfoRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetAutoClipsTaskInfoRequest) Validate() error {
	return dara.Validate(s)
}
