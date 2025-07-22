// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartCloudNoteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StartCloudNoteResponseBody
	GetRequestId() *string
	SetTaskId(v string) *StartCloudNoteResponseBody
	GetTaskId() *string
}

type StartCloudNoteResponseBody struct {
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CF8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 123
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s StartCloudNoteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartCloudNoteResponseBody) GoString() string {
	return s.String()
}

func (s *StartCloudNoteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartCloudNoteResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *StartCloudNoteResponseBody) SetRequestId(v string) *StartCloudNoteResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartCloudNoteResponseBody) SetTaskId(v string) *StartCloudNoteResponseBody {
	s.TaskId = &v
	return s
}

func (s *StartCloudNoteResponseBody) Validate() error {
	return dara.Validate(s)
}
