// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *StopInstanceResponseBody
	GetData() *bool
	SetErrorCode(v string) *StopInstanceResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *StopInstanceResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *StopInstanceResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *StopInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *StopInstanceResponseBody
	GetSuccess() *bool
}

type StopInstanceResponseBody struct {
	// The result returned.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error code.
	//
	// example:
	//
	// Invalid.Tenant.ConnectionNotExists
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// The connection does not exist.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E6F0DBDD-5AD****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StopInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *StopInstanceResponseBody) GetData() *bool {
	return s.Data
}

func (s *StopInstanceResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *StopInstanceResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *StopInstanceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *StopInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *StopInstanceResponseBody) SetData(v bool) *StopInstanceResponseBody {
	s.Data = &v
	return s
}

func (s *StopInstanceResponseBody) SetErrorCode(v string) *StopInstanceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *StopInstanceResponseBody) SetErrorMessage(v string) *StopInstanceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *StopInstanceResponseBody) SetHttpStatusCode(v int32) *StopInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *StopInstanceResponseBody) SetRequestId(v string) *StopInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopInstanceResponseBody) SetSuccess(v bool) *StopInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *StopInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
