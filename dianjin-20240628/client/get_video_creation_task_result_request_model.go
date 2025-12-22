// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVideoCreationTaskResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *GetVideoCreationTaskResultRequest
	GetTaskId() *string
	SetUserId(v string) *GetVideoCreationTaskResultRequest
	GetUserId() *string
}

type GetVideoCreationTaskResultRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxx
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// example:
	//
	// 1
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetVideoCreationTaskResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetVideoCreationTaskResultRequest) GoString() string {
	return s.String()
}

func (s *GetVideoCreationTaskResultRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetVideoCreationTaskResultRequest) GetUserId() *string {
	return s.UserId
}

func (s *GetVideoCreationTaskResultRequest) SetTaskId(v string) *GetVideoCreationTaskResultRequest {
	s.TaskId = &v
	return s
}

func (s *GetVideoCreationTaskResultRequest) SetUserId(v string) *GetVideoCreationTaskResultRequest {
	s.UserId = &v
	return s
}

func (s *GetVideoCreationTaskResultRequest) Validate() error {
	return dara.Validate(s)
}
