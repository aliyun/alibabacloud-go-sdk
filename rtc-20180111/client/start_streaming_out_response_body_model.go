// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartStreamingOutResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StartStreamingOutResponseBody
	GetRequestId() *string
	SetTaskId(v string) *StartStreamingOutResponseBody
	GetTaskId() *string
}

type StartStreamingOutResponseBody struct {
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CF8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 123
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s StartStreamingOutResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartStreamingOutResponseBody) GoString() string {
	return s.String()
}

func (s *StartStreamingOutResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartStreamingOutResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *StartStreamingOutResponseBody) SetRequestId(v string) *StartStreamingOutResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartStreamingOutResponseBody) SetTaskId(v string) *StartStreamingOutResponseBody {
	s.TaskId = &v
	return s
}

func (s *StartStreamingOutResponseBody) Validate() error {
	return dara.Validate(s)
}
