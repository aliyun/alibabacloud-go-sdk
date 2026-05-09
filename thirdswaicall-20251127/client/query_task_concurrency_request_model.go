// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTaskConcurrencyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v int64) *QueryTaskConcurrencyRequest
	GetTaskId() *int64
}

type QueryTaskConcurrencyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 12345
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s QueryTaskConcurrencyRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryTaskConcurrencyRequest) GoString() string {
	return s.String()
}

func (s *QueryTaskConcurrencyRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *QueryTaskConcurrencyRequest) SetTaskId(v int64) *QueryTaskConcurrencyRequest {
	s.TaskId = &v
	return s
}

func (s *QueryTaskConcurrencyRequest) Validate() error {
	return dara.Validate(s)
}
