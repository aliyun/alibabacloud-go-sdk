// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAsyncCreateClipsTimeLineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProcessPrompt(v string) *AsyncCreateClipsTimeLineRequest
	GetProcessPrompt() *string
	SetTaskId(v string) *AsyncCreateClipsTimeLineRequest
	GetTaskId() *string
	SetWorkspaceId(v string) *AsyncCreateClipsTimeLineRequest
	GetWorkspaceId() *string
}

type AsyncCreateClipsTimeLineRequest struct {
	ProcessPrompt *string `json:"ProcessPrompt,omitempty" xml:"ProcessPrompt,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 7AA2AE16-D873-5C5F-9708-15396C382EB1
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-2setzb9x4ewsd
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s AsyncCreateClipsTimeLineRequest) String() string {
	return dara.Prettify(s)
}

func (s AsyncCreateClipsTimeLineRequest) GoString() string {
	return s.String()
}

func (s *AsyncCreateClipsTimeLineRequest) GetProcessPrompt() *string {
	return s.ProcessPrompt
}

func (s *AsyncCreateClipsTimeLineRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *AsyncCreateClipsTimeLineRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *AsyncCreateClipsTimeLineRequest) SetProcessPrompt(v string) *AsyncCreateClipsTimeLineRequest {
	s.ProcessPrompt = &v
	return s
}

func (s *AsyncCreateClipsTimeLineRequest) SetTaskId(v string) *AsyncCreateClipsTimeLineRequest {
	s.TaskId = &v
	return s
}

func (s *AsyncCreateClipsTimeLineRequest) SetWorkspaceId(v string) *AsyncCreateClipsTimeLineRequest {
	s.WorkspaceId = &v
	return s
}

func (s *AsyncCreateClipsTimeLineRequest) Validate() error {
	return dara.Validate(s)
}
