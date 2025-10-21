// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopSessionClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *StopSessionClusterResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *StopSessionClusterResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *StopSessionClusterResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *StopSessionClusterResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *StopSessionClusterResponseBody
	GetSuccess() *bool
}

type StopSessionClusterResponseBody struct {
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// example:
	//
	// CBC799F0-AS7S-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s StopSessionClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopSessionClusterResponseBody) GoString() string {
	return s.String()
}

func (s *StopSessionClusterResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *StopSessionClusterResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *StopSessionClusterResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *StopSessionClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopSessionClusterResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *StopSessionClusterResponseBody) SetErrorCode(v string) *StopSessionClusterResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *StopSessionClusterResponseBody) SetErrorMessage(v string) *StopSessionClusterResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *StopSessionClusterResponseBody) SetHttpCode(v int32) *StopSessionClusterResponseBody {
	s.HttpCode = &v
	return s
}

func (s *StopSessionClusterResponseBody) SetRequestId(v string) *StopSessionClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopSessionClusterResponseBody) SetSuccess(v bool) *StopSessionClusterResponseBody {
	s.Success = &v
	return s
}

func (s *StopSessionClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
