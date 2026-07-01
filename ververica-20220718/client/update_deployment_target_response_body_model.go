// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDeploymentTargetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DeploymentTarget) *UpdateDeploymentTargetResponseBody
	GetData() *DeploymentTarget
	SetErrorCode(v string) *UpdateDeploymentTargetResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *UpdateDeploymentTargetResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *UpdateDeploymentTargetResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *UpdateDeploymentTargetResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateDeploymentTargetResponseBody
	GetSuccess() *bool
}

type UpdateDeploymentTargetResponseBody struct {
	// The updated deployment target.
	Data *DeploymentTarget `json:"data,omitempty" xml:"data,omitempty"`
	// The error code returned if the request fails. This parameter is empty if the request is successful.
	//
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// The error message returned if the request fails. This parameter is empty if the request is successful.
	//
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The HTTP status code. A value of 200 is always returned. Use the \\`success\\` parameter to determine whether the request was successful.
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

func (s UpdateDeploymentTargetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDeploymentTargetResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDeploymentTargetResponseBody) GetData() *DeploymentTarget {
	return s.Data
}

func (s *UpdateDeploymentTargetResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateDeploymentTargetResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateDeploymentTargetResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *UpdateDeploymentTargetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDeploymentTargetResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateDeploymentTargetResponseBody) SetData(v *DeploymentTarget) *UpdateDeploymentTargetResponseBody {
	s.Data = v
	return s
}

func (s *UpdateDeploymentTargetResponseBody) SetErrorCode(v string) *UpdateDeploymentTargetResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateDeploymentTargetResponseBody) SetErrorMessage(v string) *UpdateDeploymentTargetResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateDeploymentTargetResponseBody) SetHttpCode(v int32) *UpdateDeploymentTargetResponseBody {
	s.HttpCode = &v
	return s
}

func (s *UpdateDeploymentTargetResponseBody) SetRequestId(v string) *UpdateDeploymentTargetResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDeploymentTargetResponseBody) SetSuccess(v bool) *UpdateDeploymentTargetResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateDeploymentTargetResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
