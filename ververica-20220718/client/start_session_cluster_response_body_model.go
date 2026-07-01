// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartSessionClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *StartSessionClusterResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *StartSessionClusterResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *StartSessionClusterResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *StartSessionClusterResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *StartSessionClusterResponseBody
	GetSuccess() *bool
}

type StartSessionClusterResponseBody struct {
	// The error code. This parameter is empty if the request is successful. If the request fails, this parameter is not empty.
	//
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// The error message. This parameter is empty if the request is successful. If the request fails, this parameter is not empty.
	//
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The HTTP status code. A value of 200 is always returned. Use the success parameter to determine whether the request was successful.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CBC799F0-ABCF-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s StartSessionClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartSessionClusterResponseBody) GoString() string {
	return s.String()
}

func (s *StartSessionClusterResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *StartSessionClusterResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *StartSessionClusterResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *StartSessionClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartSessionClusterResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *StartSessionClusterResponseBody) SetErrorCode(v string) *StartSessionClusterResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *StartSessionClusterResponseBody) SetErrorMessage(v string) *StartSessionClusterResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *StartSessionClusterResponseBody) SetHttpCode(v int32) *StartSessionClusterResponseBody {
	s.HttpCode = &v
	return s
}

func (s *StartSessionClusterResponseBody) SetRequestId(v string) *StartSessionClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartSessionClusterResponseBody) SetSuccess(v bool) *StartSessionClusterResponseBody {
	s.Success = &v
	return s
}

func (s *StartSessionClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
