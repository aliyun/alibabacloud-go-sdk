// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPauseC2paBucketSignTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *PauseC2paBucketSignTaskResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *PauseC2paBucketSignTaskResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *PauseC2paBucketSignTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *PauseC2paBucketSignTaskResponseBody
	GetRequestId() *string
	SetStatus(v string) *PauseC2paBucketSignTaskResponseBody
	GetStatus() *string
	SetSuccess(v bool) *PauseC2paBucketSignTaskResponseBody
	GetSuccess() *bool
	SetTaskId(v string) *PauseC2paBucketSignTaskResponseBody
	GetTaskId() *string
}

type PauseC2paBucketSignTaskResponseBody struct {
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
	// The pause processing status. PAUSING is returned if submitted files exist. Otherwise, PAUSED is returned.
	//
	// example:
	//
	// PAUSING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The bucket batch signing task ID.
	//
	// example:
	//
	// 047b8305-9572-4eed-b7fb-8588b989b6e5
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s PauseC2paBucketSignTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PauseC2paBucketSignTaskResponseBody) GoString() string {
	return s.String()
}

func (s *PauseC2paBucketSignTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *PauseC2paBucketSignTaskResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *PauseC2paBucketSignTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PauseC2paBucketSignTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PauseC2paBucketSignTaskResponseBody) GetStatus() *string {
	return s.Status
}

func (s *PauseC2paBucketSignTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *PauseC2paBucketSignTaskResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *PauseC2paBucketSignTaskResponseBody) SetCode(v string) *PauseC2paBucketSignTaskResponseBody {
	s.Code = &v
	return s
}

func (s *PauseC2paBucketSignTaskResponseBody) SetHttpStatusCode(v int32) *PauseC2paBucketSignTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *PauseC2paBucketSignTaskResponseBody) SetMessage(v string) *PauseC2paBucketSignTaskResponseBody {
	s.Message = &v
	return s
}

func (s *PauseC2paBucketSignTaskResponseBody) SetRequestId(v string) *PauseC2paBucketSignTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *PauseC2paBucketSignTaskResponseBody) SetStatus(v string) *PauseC2paBucketSignTaskResponseBody {
	s.Status = &v
	return s
}

func (s *PauseC2paBucketSignTaskResponseBody) SetSuccess(v bool) *PauseC2paBucketSignTaskResponseBody {
	s.Success = &v
	return s
}

func (s *PauseC2paBucketSignTaskResponseBody) SetTaskId(v string) *PauseC2paBucketSignTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *PauseC2paBucketSignTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
