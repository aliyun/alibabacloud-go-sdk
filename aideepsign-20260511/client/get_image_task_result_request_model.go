// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetImageTaskResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *GetImageTaskResultRequest
	GetTaskId() *string
}

type GetImageTaskResultRequest struct {
	// The task ID returned by `CreateImageTask`.
	//
	// This parameter is required.
	//
	// example:
	//
	// f47ac10b-58cc-4372-a567-0e02b2c3d479
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetImageTaskResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetImageTaskResultRequest) GoString() string {
	return s.String()
}

func (s *GetImageTaskResultRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetImageTaskResultRequest) SetTaskId(v string) *GetImageTaskResultRequest {
	s.TaskId = &v
	return s
}

func (s *GetImageTaskResultRequest) Validate() error {
	return dara.Validate(s)
}
