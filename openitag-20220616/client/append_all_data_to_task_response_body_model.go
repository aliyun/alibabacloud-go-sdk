// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAppendAllDataToTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *AppendAllDataToTaskResponseBody
	GetCode() *int32
	SetDetails(v string) *AppendAllDataToTaskResponseBody
	GetDetails() *string
	SetErrorCode(v string) *AppendAllDataToTaskResponseBody
	GetErrorCode() *string
	SetMessage(v string) *AppendAllDataToTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *AppendAllDataToTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AppendAllDataToTaskResponseBody
	GetSuccess() *bool
}

type AppendAllDataToTaskResponseBody struct {
	// Return code. The default value is 0, indicating normal execution.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Details.
	//
	// example:
	//
	// null
	Details *string `json:"Details,omitempty" xml:"Details,omitempty"`
	// Error code
	//
	// example:
	//
	// ""
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// Response message of the request.
	//
	// This parameter is required.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 90ABA848-AD74-1F6E-84BC-4182A7F1F29E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation succeeded.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AppendAllDataToTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AppendAllDataToTaskResponseBody) GoString() string {
	return s.String()
}

func (s *AppendAllDataToTaskResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *AppendAllDataToTaskResponseBody) GetDetails() *string {
	return s.Details
}

func (s *AppendAllDataToTaskResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *AppendAllDataToTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AppendAllDataToTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AppendAllDataToTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AppendAllDataToTaskResponseBody) SetCode(v int32) *AppendAllDataToTaskResponseBody {
	s.Code = &v
	return s
}

func (s *AppendAllDataToTaskResponseBody) SetDetails(v string) *AppendAllDataToTaskResponseBody {
	s.Details = &v
	return s
}

func (s *AppendAllDataToTaskResponseBody) SetErrorCode(v string) *AppendAllDataToTaskResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *AppendAllDataToTaskResponseBody) SetMessage(v string) *AppendAllDataToTaskResponseBody {
	s.Message = &v
	return s
}

func (s *AppendAllDataToTaskResponseBody) SetRequestId(v string) *AppendAllDataToTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *AppendAllDataToTaskResponseBody) SetSuccess(v bool) *AppendAllDataToTaskResponseBody {
	s.Success = &v
	return s
}

func (s *AppendAllDataToTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
