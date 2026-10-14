// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateC2paBucketSignTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateC2paBucketSignTaskResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *CreateC2paBucketSignTaskResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateC2paBucketSignTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateC2paBucketSignTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateC2paBucketSignTaskResponseBody
	GetSuccess() *bool
	SetTaskId(v string) *CreateC2paBucketSignTaskResponseBody
	GetTaskId() *string
}

type CreateC2paBucketSignTaskResponseBody struct {
	// The return code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 01A03DCE-4091-1A91-9810-692C3D783877
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The ID of the batch signing task for the bucket. You can use this ID to query, pause, resume, or stop the task.
	//
	// example:
	//
	// 047b8305-9572-4eed-b7fb-8588b989b6e5
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateC2paBucketSignTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateC2paBucketSignTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateC2paBucketSignTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateC2paBucketSignTaskResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateC2paBucketSignTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateC2paBucketSignTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateC2paBucketSignTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateC2paBucketSignTaskResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateC2paBucketSignTaskResponseBody) SetCode(v string) *CreateC2paBucketSignTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CreateC2paBucketSignTaskResponseBody) SetHttpStatusCode(v int32) *CreateC2paBucketSignTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateC2paBucketSignTaskResponseBody) SetMessage(v string) *CreateC2paBucketSignTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreateC2paBucketSignTaskResponseBody) SetRequestId(v string) *CreateC2paBucketSignTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateC2paBucketSignTaskResponseBody) SetSuccess(v bool) *CreateC2paBucketSignTaskResponseBody {
	s.Success = &v
	return s
}

func (s *CreateC2paBucketSignTaskResponseBody) SetTaskId(v string) *CreateC2paBucketSignTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateC2paBucketSignTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
