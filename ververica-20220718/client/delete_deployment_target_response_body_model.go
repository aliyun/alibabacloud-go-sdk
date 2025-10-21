// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDeploymentTargetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DeploymentTarget) *DeleteDeploymentTargetResponseBody
	GetData() *DeploymentTarget
	SetErrorCode(v string) *DeleteDeploymentTargetResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DeleteDeploymentTargetResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *DeleteDeploymentTargetResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *DeleteDeploymentTargetResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteDeploymentTargetResponseBody
	GetSuccess() *bool
}

type DeleteDeploymentTargetResponseBody struct {
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

func (s DeleteDeploymentTargetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDeploymentTargetResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDeploymentTargetResponseBody) GetData() *DeploymentTarget {
	return s.Data
}

func (s *DeleteDeploymentTargetResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteDeploymentTargetResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteDeploymentTargetResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *DeleteDeploymentTargetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDeploymentTargetResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteDeploymentTargetResponseBody) SetData(v *DeploymentTarget) *DeleteDeploymentTargetResponseBody {
	s.Data = v
	return s
}

func (s *DeleteDeploymentTargetResponseBody) SetErrorCode(v string) *DeleteDeploymentTargetResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteDeploymentTargetResponseBody) SetErrorMessage(v string) *DeleteDeploymentTargetResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteDeploymentTargetResponseBody) SetHttpCode(v int32) *DeleteDeploymentTargetResponseBody {
	s.HttpCode = &v
	return s
}

func (s *DeleteDeploymentTargetResponseBody) SetRequestId(v string) *DeleteDeploymentTargetResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDeploymentTargetResponseBody) SetSuccess(v bool) *DeleteDeploymentTargetResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteDeploymentTargetResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
