// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopStreamingOutResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StopStreamingOutResponseBody
	GetRequestId() *string
	SetTaskId(v string) *StopStreamingOutResponseBody
	GetTaskId() *string
}

type StopStreamingOutResponseBody struct {
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CF8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 123
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s StopStreamingOutResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopStreamingOutResponseBody) GoString() string {
	return s.String()
}

func (s *StopStreamingOutResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopStreamingOutResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *StopStreamingOutResponseBody) SetRequestId(v string) *StopStreamingOutResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopStreamingOutResponseBody) SetTaskId(v string) *StopStreamingOutResponseBody {
	s.TaskId = &v
	return s
}

func (s *StopStreamingOutResponseBody) Validate() error {
	return dara.Validate(s)
}
