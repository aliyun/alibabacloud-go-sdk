// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVideoClipTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateVideoClipTaskResponseBody
	GetRequestId() *string
	SetTaskId(v string) *CreateVideoClipTaskResponseBody
	GetTaskId() *string
}

type CreateVideoClipTaskResponseBody struct {
	// example:
	//
	// 86A90C40-D1AB-50DA-A4B1-0D545F80F2FE
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 837091359375048704
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s CreateVideoClipTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateVideoClipTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVideoClipTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateVideoClipTaskResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateVideoClipTaskResponseBody) SetRequestId(v string) *CreateVideoClipTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVideoClipTaskResponseBody) SetTaskId(v string) *CreateVideoClipTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateVideoClipTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
