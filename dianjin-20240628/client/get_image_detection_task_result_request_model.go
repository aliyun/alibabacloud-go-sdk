// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetImageDetectionTaskResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *GetImageDetectionTaskResultRequest
	GetTaskId() *string
	SetUserId(v string) *GetImageDetectionTaskResultRequest
	GetUserId() *string
}

type GetImageDetectionTaskResultRequest struct {
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

func (s GetImageDetectionTaskResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetImageDetectionTaskResultRequest) GoString() string {
	return s.String()
}

func (s *GetImageDetectionTaskResultRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetImageDetectionTaskResultRequest) GetUserId() *string {
	return s.UserId
}

func (s *GetImageDetectionTaskResultRequest) SetTaskId(v string) *GetImageDetectionTaskResultRequest {
	s.TaskId = &v
	return s
}

func (s *GetImageDetectionTaskResultRequest) SetUserId(v string) *GetImageDetectionTaskResultRequest {
	s.UserId = &v
	return s
}

func (s *GetImageDetectionTaskResultRequest) Validate() error {
	return dara.Validate(s)
}
