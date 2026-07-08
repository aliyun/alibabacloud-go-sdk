// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunDeepWritingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCursor(v int32) *RunDeepWritingRequest
	GetCursor() *int32
	SetTaskId(v string) *RunDeepWritingRequest
	GetTaskId() *string
	SetWorkspaceId(v string) *RunDeepWritingRequest
	GetWorkspaceId() *string
}

type RunDeepWritingRequest struct {
	// The cursor, which is the ordinal number of the streaming event.
	//
	// example:
	//
	// 10
	Cursor *int32 `json:"Cursor,omitempty" xml:"Cursor,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 95c2fbe6-5a20-4fc2-8a93-376ed05fbe13
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The unique identifier of the Alibaba Cloud Model Studio workspace. For more information, see [Workspace ID](https://help.aliyun.com/document_detail/2782167.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// llm-ir8zkqry2fncaxwq
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s RunDeepWritingRequest) String() string {
	return dara.Prettify(s)
}

func (s RunDeepWritingRequest) GoString() string {
	return s.String()
}

func (s *RunDeepWritingRequest) GetCursor() *int32 {
	return s.Cursor
}

func (s *RunDeepWritingRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *RunDeepWritingRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *RunDeepWritingRequest) SetCursor(v int32) *RunDeepWritingRequest {
	s.Cursor = &v
	return s
}

func (s *RunDeepWritingRequest) SetTaskId(v string) *RunDeepWritingRequest {
	s.TaskId = &v
	return s
}

func (s *RunDeepWritingRequest) SetWorkspaceId(v string) *RunDeepWritingRequest {
	s.WorkspaceId = &v
	return s
}

func (s *RunDeepWritingRequest) Validate() error {
	return dara.Validate(s)
}
