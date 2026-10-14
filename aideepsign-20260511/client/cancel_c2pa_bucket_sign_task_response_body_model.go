// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelC2paBucketSignTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CancelC2paBucketSignTaskResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *CancelC2paBucketSignTaskResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CancelC2paBucketSignTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *CancelC2paBucketSignTaskResponseBody
	GetRequestId() *string
	SetStatus(v string) *CancelC2paBucketSignTaskResponseBody
	GetStatus() *string
	SetSuccess(v bool) *CancelC2paBucketSignTaskResponseBody
	GetSuccess() *bool
	SetTaskId(v string) *CancelC2paBucketSignTaskResponseBody
	GetTaskId() *string
}

type CancelC2paBucketSignTaskResponseBody struct {
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
	// The response message.
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
	// The stop processing status. CANCELING indicates that the stop process is being finalized. Repeated requests may return the completed CANCELED status.
	//
	// example:
	//
	// CANCELING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The ID of the bucket batch signing task.
	//
	// example:
	//
	// 047b8305-9572-4eed-b7fb-8588b989b6e5
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CancelC2paBucketSignTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelC2paBucketSignTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CancelC2paBucketSignTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *CancelC2paBucketSignTaskResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CancelC2paBucketSignTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CancelC2paBucketSignTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelC2paBucketSignTaskResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CancelC2paBucketSignTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CancelC2paBucketSignTaskResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *CancelC2paBucketSignTaskResponseBody) SetCode(v string) *CancelC2paBucketSignTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CancelC2paBucketSignTaskResponseBody) SetHttpStatusCode(v int32) *CancelC2paBucketSignTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CancelC2paBucketSignTaskResponseBody) SetMessage(v string) *CancelC2paBucketSignTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CancelC2paBucketSignTaskResponseBody) SetRequestId(v string) *CancelC2paBucketSignTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelC2paBucketSignTaskResponseBody) SetStatus(v string) *CancelC2paBucketSignTaskResponseBody {
	s.Status = &v
	return s
}

func (s *CancelC2paBucketSignTaskResponseBody) SetSuccess(v bool) *CancelC2paBucketSignTaskResponseBody {
	s.Success = &v
	return s
}

func (s *CancelC2paBucketSignTaskResponseBody) SetTaskId(v string) *CancelC2paBucketSignTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *CancelC2paBucketSignTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
