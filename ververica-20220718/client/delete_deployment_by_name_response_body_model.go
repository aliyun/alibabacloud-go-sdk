// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDeploymentByNameResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *DeleteDeploymentByNameResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DeleteDeploymentByNameResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *DeleteDeploymentByNameResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *DeleteDeploymentByNameResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteDeploymentByNameResponseBody
	GetSuccess() *bool
}

type DeleteDeploymentByNameResponseBody struct {
	// The error code. This field is not empty when success is false, indicating a business error code. This field is empty when success is true.
	//
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// The error message. This field is not empty when success is false, indicating a business error message. This field is empty when success is true.
	//
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The HTTP status code, which is always 200. Use the success field to determine whether the request was successful.
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
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteDeploymentByNameResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDeploymentByNameResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDeploymentByNameResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteDeploymentByNameResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteDeploymentByNameResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *DeleteDeploymentByNameResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDeploymentByNameResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteDeploymentByNameResponseBody) SetErrorCode(v string) *DeleteDeploymentByNameResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteDeploymentByNameResponseBody) SetErrorMessage(v string) *DeleteDeploymentByNameResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteDeploymentByNameResponseBody) SetHttpCode(v int32) *DeleteDeploymentByNameResponseBody {
	s.HttpCode = &v
	return s
}

func (s *DeleteDeploymentByNameResponseBody) SetRequestId(v string) *DeleteDeploymentByNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDeploymentByNameResponseBody) SetSuccess(v bool) *DeleteDeploymentByNameResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteDeploymentByNameResponseBody) Validate() error {
	return dara.Validate(s)
}
