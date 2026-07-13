// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTaskConcurrencyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationCode(v string) *QueryTaskConcurrencyRequest
	GetApplicationCode() *string
	SetTaskId(v int64) *QueryTaskConcurrencyRequest
	GetTaskId() *int64
}

type QueryTaskConcurrencyRequest struct {
	// example:
	//
	// B9191F0E57
	ApplicationCode *string `json:"ApplicationCode,omitempty" xml:"ApplicationCode,omitempty"`
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

func (s *QueryTaskConcurrencyRequest) GetApplicationCode() *string {
	return s.ApplicationCode
}

func (s *QueryTaskConcurrencyRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *QueryTaskConcurrencyRequest) SetApplicationCode(v string) *QueryTaskConcurrencyRequest {
	s.ApplicationCode = &v
	return s
}

func (s *QueryTaskConcurrencyRequest) SetTaskId(v int64) *QueryTaskConcurrencyRequest {
	s.TaskId = &v
	return s
}

func (s *QueryTaskConcurrencyRequest) Validate() error {
	return dara.Validate(s)
}
