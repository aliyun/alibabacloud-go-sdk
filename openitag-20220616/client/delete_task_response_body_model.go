// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DeleteTaskResponseBody
	GetCode() *int32
	SetDetails(v string) *DeleteTaskResponseBody
	GetDetails() *string
	SetErrorCode(v string) *DeleteTaskResponseBody
	GetErrorCode() *string
	SetMessage(v string) *DeleteTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteTaskResponseBody
	GetSuccess() *bool
}

type DeleteTaskResponseBody struct {
	// Total number of data entries under the conditions of this request. This parameter is optional and is not returned by default.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Details
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
	// Response message of the request
	//
	// This parameter is required.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID
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

func (s DeleteTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTaskResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteTaskResponseBody) GetDetails() *string {
	return s.Details
}

func (s *DeleteTaskResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteTaskResponseBody) SetCode(v int32) *DeleteTaskResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteTaskResponseBody) SetDetails(v string) *DeleteTaskResponseBody {
	s.Details = &v
	return s
}

func (s *DeleteTaskResponseBody) SetErrorCode(v string) *DeleteTaskResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteTaskResponseBody) SetMessage(v string) *DeleteTaskResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteTaskResponseBody) SetRequestId(v string) *DeleteTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTaskResponseBody) SetSuccess(v bool) *DeleteTaskResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
