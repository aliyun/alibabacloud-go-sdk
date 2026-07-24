// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDeploymentByNameResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *Deployment) *UpdateDeploymentByNameResponseBody
	GetData() *Deployment
	SetErrorCode(v string) *UpdateDeploymentByNameResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *UpdateDeploymentByNameResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *UpdateDeploymentByNameResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *UpdateDeploymentByNameResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateDeploymentByNameResponseBody
	GetSuccess() *bool
}

type UpdateDeploymentByNameResponseBody struct {
	// The updated job content returned when success is true. This parameter is empty when success is false.
	//
	// example:
	//
	// 123
	Data *Deployment `json:"data,omitempty" xml:"data,omitempty"`
	// The error code. This parameter is not empty when success is false. This parameter is empty when success is true.
	//
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// The error message. This parameter is not empty when success is false. This parameter is empty when success is true.
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
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateDeploymentByNameResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDeploymentByNameResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDeploymentByNameResponseBody) GetData() *Deployment {
	return s.Data
}

func (s *UpdateDeploymentByNameResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateDeploymentByNameResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateDeploymentByNameResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *UpdateDeploymentByNameResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDeploymentByNameResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateDeploymentByNameResponseBody) SetData(v *Deployment) *UpdateDeploymentByNameResponseBody {
	s.Data = v
	return s
}

func (s *UpdateDeploymentByNameResponseBody) SetErrorCode(v string) *UpdateDeploymentByNameResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateDeploymentByNameResponseBody) SetErrorMessage(v string) *UpdateDeploymentByNameResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateDeploymentByNameResponseBody) SetHttpCode(v int32) *UpdateDeploymentByNameResponseBody {
	s.HttpCode = &v
	return s
}

func (s *UpdateDeploymentByNameResponseBody) SetRequestId(v string) *UpdateDeploymentByNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDeploymentByNameResponseBody) SetSuccess(v bool) *UpdateDeploymentByNameResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateDeploymentByNameResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
