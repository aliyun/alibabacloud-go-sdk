// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDeploymentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *Deployment) *CreateDeploymentResponseBody
	GetData() *Deployment
	SetErrorCode(v string) *CreateDeploymentResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateDeploymentResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *CreateDeploymentResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *CreateDeploymentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateDeploymentResponseBody
	GetSuccess() *bool
}

type CreateDeploymentResponseBody struct {
	// 	- If the value of success was true, the deployment that you created was returned.
	//
	// 	- If the value of success was false, a null value was returned.
	Data *Deployment `json:"data,omitempty" xml:"data,omitempty"`
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

func (s CreateDeploymentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDeploymentResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDeploymentResponseBody) GetData() *Deployment {
	return s.Data
}

func (s *CreateDeploymentResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateDeploymentResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateDeploymentResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *CreateDeploymentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDeploymentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateDeploymentResponseBody) SetData(v *Deployment) *CreateDeploymentResponseBody {
	s.Data = v
	return s
}

func (s *CreateDeploymentResponseBody) SetErrorCode(v string) *CreateDeploymentResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateDeploymentResponseBody) SetErrorMessage(v string) *CreateDeploymentResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateDeploymentResponseBody) SetHttpCode(v int32) *CreateDeploymentResponseBody {
	s.HttpCode = &v
	return s
}

func (s *CreateDeploymentResponseBody) SetRequestId(v string) *CreateDeploymentResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDeploymentResponseBody) SetSuccess(v bool) *CreateDeploymentResponseBody {
	s.Success = &v
	return s
}

func (s *CreateDeploymentResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
