// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumeC2paBucketSignTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ResumeC2paBucketSignTaskResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ResumeC2paBucketSignTaskResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ResumeC2paBucketSignTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *ResumeC2paBucketSignTaskResponseBody
	GetRequestId() *string
	SetStatus(v string) *ResumeC2paBucketSignTaskResponseBody
	GetStatus() *string
	SetSuccess(v bool) *ResumeC2paBucketSignTaskResponseBody
	GetSuccess() *bool
	SetTaskId(v string) *ResumeC2paBucketSignTaskResponseBody
	GetTaskId() *string
}

type ResumeC2paBucketSignTaskResponseBody struct {
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
	// The task status after resumption. The value PROCESSING is returned.
	//
	// example:
	//
	// PROCESSING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The Bucket batch signing task ID.
	//
	// example:
	//
	// 047b8305-9572-4eed-b7fb-8588b989b6e5
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ResumeC2paBucketSignTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResumeC2paBucketSignTaskResponseBody) GoString() string {
	return s.String()
}

func (s *ResumeC2paBucketSignTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *ResumeC2paBucketSignTaskResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ResumeC2paBucketSignTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ResumeC2paBucketSignTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResumeC2paBucketSignTaskResponseBody) GetStatus() *string {
	return s.Status
}

func (s *ResumeC2paBucketSignTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ResumeC2paBucketSignTaskResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *ResumeC2paBucketSignTaskResponseBody) SetCode(v string) *ResumeC2paBucketSignTaskResponseBody {
	s.Code = &v
	return s
}

func (s *ResumeC2paBucketSignTaskResponseBody) SetHttpStatusCode(v int32) *ResumeC2paBucketSignTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ResumeC2paBucketSignTaskResponseBody) SetMessage(v string) *ResumeC2paBucketSignTaskResponseBody {
	s.Message = &v
	return s
}

func (s *ResumeC2paBucketSignTaskResponseBody) SetRequestId(v string) *ResumeC2paBucketSignTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResumeC2paBucketSignTaskResponseBody) SetStatus(v string) *ResumeC2paBucketSignTaskResponseBody {
	s.Status = &v
	return s
}

func (s *ResumeC2paBucketSignTaskResponseBody) SetSuccess(v bool) *ResumeC2paBucketSignTaskResponseBody {
	s.Success = &v
	return s
}

func (s *ResumeC2paBucketSignTaskResponseBody) SetTaskId(v string) *ResumeC2paBucketSignTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *ResumeC2paBucketSignTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
