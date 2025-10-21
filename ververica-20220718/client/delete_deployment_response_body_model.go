// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDeploymentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *DeleteDeploymentResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DeleteDeploymentResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *DeleteDeploymentResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *DeleteDeploymentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteDeploymentResponseBody
	GetSuccess() *bool
}

type DeleteDeploymentResponseBody struct {
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

func (s DeleteDeploymentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDeploymentResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDeploymentResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteDeploymentResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteDeploymentResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *DeleteDeploymentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDeploymentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteDeploymentResponseBody) SetErrorCode(v string) *DeleteDeploymentResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteDeploymentResponseBody) SetErrorMessage(v string) *DeleteDeploymentResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteDeploymentResponseBody) SetHttpCode(v int32) *DeleteDeploymentResponseBody {
	s.HttpCode = &v
	return s
}

func (s *DeleteDeploymentResponseBody) SetRequestId(v string) *DeleteDeploymentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDeploymentResponseBody) SetSuccess(v bool) *DeleteDeploymentResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteDeploymentResponseBody) Validate() error {
	return dara.Validate(s)
}
