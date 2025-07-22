// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartCloudRecordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StartCloudRecordResponseBody
	GetRequestId() *string
	SetTaskId(v string) *StartCloudRecordResponseBody
	GetTaskId() *string
}

type StartCloudRecordResponseBody struct {
	// requestId
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CF8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// taskId
	//
	// example:
	//
	// 123
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s StartCloudRecordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartCloudRecordResponseBody) GoString() string {
	return s.String()
}

func (s *StartCloudRecordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartCloudRecordResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *StartCloudRecordResponseBody) SetRequestId(v string) *StartCloudRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartCloudRecordResponseBody) SetTaskId(v string) *StartCloudRecordResponseBody {
	s.TaskId = &v
	return s
}

func (s *StartCloudRecordResponseBody) Validate() error {
	return dara.Validate(s)
}
