// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAsyncTaskResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *QueryAsyncTaskResultRequest
	GetTaskId() *string
}

type QueryAsyncTaskResultRequest struct {
	// The ID of the asynchronous task. This parameter is required.
	//
	// example:
	//
	// b67f6089-085a-9402-93c6-bac0561b3a06
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s QueryAsyncTaskResultRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryAsyncTaskResultRequest) GoString() string {
	return s.String()
}

func (s *QueryAsyncTaskResultRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *QueryAsyncTaskResultRequest) SetTaskId(v string) *QueryAsyncTaskResultRequest {
	s.TaskId = &v
	return s
}

func (s *QueryAsyncTaskResultRequest) Validate() error {
	return dara.Validate(s)
}
