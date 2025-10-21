// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDeploymentTargetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DeploymentTarget) *CreateDeploymentTargetResponseBody
	GetData() *DeploymentTarget
	SetErrorCode(v string) *CreateDeploymentTargetResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateDeploymentTargetResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *CreateDeploymentTargetResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *CreateDeploymentTargetResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateDeploymentTargetResponseBody
	GetSuccess() *bool
}

type CreateDeploymentTargetResponseBody struct {
	Data *DeploymentTarget `json:"data,omitempty" xml:"data,omitempty"`
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

func (s CreateDeploymentTargetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDeploymentTargetResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDeploymentTargetResponseBody) GetData() *DeploymentTarget {
	return s.Data
}

func (s *CreateDeploymentTargetResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateDeploymentTargetResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateDeploymentTargetResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *CreateDeploymentTargetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDeploymentTargetResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateDeploymentTargetResponseBody) SetData(v *DeploymentTarget) *CreateDeploymentTargetResponseBody {
	s.Data = v
	return s
}

func (s *CreateDeploymentTargetResponseBody) SetErrorCode(v string) *CreateDeploymentTargetResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateDeploymentTargetResponseBody) SetErrorMessage(v string) *CreateDeploymentTargetResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateDeploymentTargetResponseBody) SetHttpCode(v int32) *CreateDeploymentTargetResponseBody {
	s.HttpCode = &v
	return s
}

func (s *CreateDeploymentTargetResponseBody) SetRequestId(v string) *CreateDeploymentTargetResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDeploymentTargetResponseBody) SetSuccess(v bool) *CreateDeploymentTargetResponseBody {
	s.Success = &v
	return s
}

func (s *CreateDeploymentTargetResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
