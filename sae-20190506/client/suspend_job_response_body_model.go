// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSuspendJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SuspendJobResponseBody
	GetCode() *string
	SetData(v string) *SuspendJobResponseBody
	GetData() *string
	SetErrorCode(v string) *SuspendJobResponseBody
	GetErrorCode() *string
	SetMessage(v string) *SuspendJobResponseBody
	GetMessage() *string
	SetRequestId(v string) *SuspendJobResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SuspendJobResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *SuspendJobResponseBody
	GetTraceId() *string
}

type SuspendJobResponseBody struct {
	// The HTTP status code. Valid values:
	//
	// 	- **2xx**: The call was successful.
	//
	// 	- **3xx**: The call was redirected.
	//
	// 	- **4xx**: The call failed.
	//
	// 	- **5xx**: A server error occurred.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Whether the execution is successful.
	//
	// example:
	//
	// {success: true}
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error code returned. Valid values:
	//
	// 	- The **ErrorCode*	- parameter is not returned if the request succeeds.
	//
	// 	- If the call fails, the **ErrorCode*	- parameter is returned. For more information, see **Error codes*	- in this topic.
	//
	// example:
	//
	// Null
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The additional information that is returned. Valid values:
	//
	// 	- success: If the call is successful, **success*	- is returned.
	//
	// 	- An error code: If the call fails, an error code is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 67DD9A98-9CCC-5BE8-8C9E-B45E72F4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the job was executed. Valid values:
	//
	// 	- **true**: The job was executed.
	//
	// 	- **false**: The job failed to be executed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The trace ID that is used to query the details of the request.
	//
	// example:
	//
	// 0b87b7e716575071334387401e****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s SuspendJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SuspendJobResponseBody) GoString() string {
	return s.String()
}

func (s *SuspendJobResponseBody) GetCode() *string {
	return s.Code
}

func (s *SuspendJobResponseBody) GetData() *string {
	return s.Data
}

func (s *SuspendJobResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *SuspendJobResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SuspendJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SuspendJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SuspendJobResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *SuspendJobResponseBody) SetCode(v string) *SuspendJobResponseBody {
	s.Code = &v
	return s
}

func (s *SuspendJobResponseBody) SetData(v string) *SuspendJobResponseBody {
	s.Data = &v
	return s
}

func (s *SuspendJobResponseBody) SetErrorCode(v string) *SuspendJobResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SuspendJobResponseBody) SetMessage(v string) *SuspendJobResponseBody {
	s.Message = &v
	return s
}

func (s *SuspendJobResponseBody) SetRequestId(v string) *SuspendJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SuspendJobResponseBody) SetSuccess(v bool) *SuspendJobResponseBody {
	s.Success = &v
	return s
}

func (s *SuspendJobResponseBody) SetTraceId(v string) *SuspendJobResponseBody {
	s.TraceId = &v
	return s
}

func (s *SuspendJobResponseBody) Validate() error {
	return dara.Validate(s)
}
