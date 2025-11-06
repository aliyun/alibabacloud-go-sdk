// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopViewResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StopViewResponseBody
	GetRequestId() *string
	SetTaskId(v string) *StopViewResponseBody
	GetTaskId() *string
}

type StopViewResponseBody struct {
	// example:
	//
	// E8236D21-B690-5251-A361-5971FBF552BA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 2bd80921b81e4d4289f696606885606b
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s StopViewResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopViewResponseBody) GoString() string {
	return s.String()
}

func (s *StopViewResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopViewResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *StopViewResponseBody) SetRequestId(v string) *StopViewResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopViewResponseBody) SetTaskId(v string) *StopViewResponseBody {
	s.TaskId = &v
	return s
}

func (s *StopViewResponseBody) Validate() error {
	return dara.Validate(s)
}
