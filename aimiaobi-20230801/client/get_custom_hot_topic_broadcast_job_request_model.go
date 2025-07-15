// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCustomHotTopicBroadcastJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *GetCustomHotTopicBroadcastJobRequest
	GetTaskId() *string
	SetWorkspaceId(v string) *GetCustomHotTopicBroadcastJobRequest
	GetWorkspaceId() *string
}

type GetCustomHotTopicBroadcastJobRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 2e27abb32cb64f80a0c6e829b6c87a09
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetCustomHotTopicBroadcastJobRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCustomHotTopicBroadcastJobRequest) GoString() string {
	return s.String()
}

func (s *GetCustomHotTopicBroadcastJobRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetCustomHotTopicBroadcastJobRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetCustomHotTopicBroadcastJobRequest) SetTaskId(v string) *GetCustomHotTopicBroadcastJobRequest {
	s.TaskId = &v
	return s
}

func (s *GetCustomHotTopicBroadcastJobRequest) SetWorkspaceId(v string) *GetCustomHotTopicBroadcastJobRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetCustomHotTopicBroadcastJobRequest) Validate() error {
	return dara.Validate(s)
}
