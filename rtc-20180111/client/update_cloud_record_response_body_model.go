// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCloudRecordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateCloudRecordResponseBody
	GetRequestId() *string
	SetTaskId(v string) *UpdateCloudRecordResponseBody
	GetTaskId() *string
}

type UpdateCloudRecordResponseBody struct {
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CF8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 123
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s UpdateCloudRecordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudRecordResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCloudRecordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCloudRecordResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *UpdateCloudRecordResponseBody) SetRequestId(v string) *UpdateCloudRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCloudRecordResponseBody) SetTaskId(v string) *UpdateCloudRecordResponseBody {
	s.TaskId = &v
	return s
}

func (s *UpdateCloudRecordResponseBody) Validate() error {
	return dara.Validate(s)
}
