// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopCloudNoteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StopCloudNoteResponseBody
	GetRequestId() *string
	SetTaskId(v string) *StopCloudNoteResponseBody
	GetTaskId() *string
}

type StopCloudNoteResponseBody struct {
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CF8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 123
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s StopCloudNoteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopCloudNoteResponseBody) GoString() string {
	return s.String()
}

func (s *StopCloudNoteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopCloudNoteResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *StopCloudNoteResponseBody) SetRequestId(v string) *StopCloudNoteResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopCloudNoteResponseBody) SetTaskId(v string) *StopCloudNoteResponseBody {
	s.TaskId = &v
	return s
}

func (s *StopCloudNoteResponseBody) Validate() error {
	return dara.Validate(s)
}
