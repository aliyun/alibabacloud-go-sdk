// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateStreamingOutResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateStreamingOutResponseBody
	GetRequestId() *string
	SetTaskId(v string) *UpdateStreamingOutResponseBody
	GetTaskId() *string
}

type UpdateStreamingOutResponseBody struct {
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CF8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 123
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s UpdateStreamingOutResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateStreamingOutResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateStreamingOutResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateStreamingOutResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *UpdateStreamingOutResponseBody) SetRequestId(v string) *UpdateStreamingOutResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateStreamingOutResponseBody) SetTaskId(v string) *UpdateStreamingOutResponseBody {
	s.TaskId = &v
	return s
}

func (s *UpdateStreamingOutResponseBody) Validate() error {
	return dara.Validate(s)
}
