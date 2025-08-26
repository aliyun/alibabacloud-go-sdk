// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPodcastTaskResultQueryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *PodcastTaskResultQueryRequest
	GetTaskId() *string
	SetWorkspaceId(v string) *PodcastTaskResultQueryRequest
	GetWorkspaceId() *string
}

type PodcastTaskResultQueryRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 63c4e0eaab3b4c0db208ecafa990e8d1
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-ep8ba0dr6seiddri
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s PodcastTaskResultQueryRequest) String() string {
	return dara.Prettify(s)
}

func (s PodcastTaskResultQueryRequest) GoString() string {
	return s.String()
}

func (s *PodcastTaskResultQueryRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *PodcastTaskResultQueryRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *PodcastTaskResultQueryRequest) SetTaskId(v string) *PodcastTaskResultQueryRequest {
	s.TaskId = &v
	return s
}

func (s *PodcastTaskResultQueryRequest) SetWorkspaceId(v string) *PodcastTaskResultQueryRequest {
	s.WorkspaceId = &v
	return s
}

func (s *PodcastTaskResultQueryRequest) Validate() error {
	return dara.Validate(s)
}
