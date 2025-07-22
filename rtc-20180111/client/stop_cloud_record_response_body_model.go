// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopCloudRecordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StopCloudRecordResponseBody
	GetRequestId() *string
	SetTaskId(v string) *StopCloudRecordResponseBody
	GetTaskId() *string
}

type StopCloudRecordResponseBody struct {
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CF8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 123
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s StopCloudRecordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopCloudRecordResponseBody) GoString() string {
	return s.String()
}

func (s *StopCloudRecordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopCloudRecordResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *StopCloudRecordResponseBody) SetRequestId(v string) *StopCloudRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopCloudRecordResponseBody) SetTaskId(v string) *StopCloudRecordResponseBody {
	s.TaskId = &v
	return s
}

func (s *StopCloudRecordResponseBody) Validate() error {
	return dara.Validate(s)
}
