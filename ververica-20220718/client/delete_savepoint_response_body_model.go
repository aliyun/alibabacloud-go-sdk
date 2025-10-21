// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSavepointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *DeleteSavepointResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DeleteSavepointResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *DeleteSavepointResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *DeleteSavepointResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteSavepointResponseBody
	GetSuccess() *bool
}

type DeleteSavepointResponseBody struct {
	// 	- If the value of success was false, an error code was returned.
	//
	// 	- If the value of success was true, a null value was returned.
	//
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 	- If the value of success was false, an error message was returned.
	//
	// 	- If the value of success was true, a null value was returned.
	//
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The value was fixed to 200.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CBC799F0-AS7S-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteSavepointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSavepointResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSavepointResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteSavepointResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteSavepointResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *DeleteSavepointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSavepointResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteSavepointResponseBody) SetErrorCode(v string) *DeleteSavepointResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteSavepointResponseBody) SetErrorMessage(v string) *DeleteSavepointResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteSavepointResponseBody) SetHttpCode(v int32) *DeleteSavepointResponseBody {
	s.HttpCode = &v
	return s
}

func (s *DeleteSavepointResponseBody) SetRequestId(v string) *DeleteSavepointResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSavepointResponseBody) SetSuccess(v bool) *DeleteSavepointResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteSavepointResponseBody) Validate() error {
	return dara.Validate(s)
}
