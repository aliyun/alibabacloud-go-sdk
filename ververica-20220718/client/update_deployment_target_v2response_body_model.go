// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDeploymentTargetV2ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DeploymentTarget) *UpdateDeploymentTargetV2ResponseBody
	GetData() *DeploymentTarget
	SetErrorCode(v string) *UpdateDeploymentTargetV2ResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *UpdateDeploymentTargetV2ResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *UpdateDeploymentTargetV2ResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *UpdateDeploymentTargetV2ResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateDeploymentTargetV2ResponseBody
	GetSuccess() *bool
}

type UpdateDeploymentTargetV2ResponseBody struct {
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

func (s UpdateDeploymentTargetV2ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDeploymentTargetV2ResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDeploymentTargetV2ResponseBody) GetData() *DeploymentTarget {
	return s.Data
}

func (s *UpdateDeploymentTargetV2ResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateDeploymentTargetV2ResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateDeploymentTargetV2ResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *UpdateDeploymentTargetV2ResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDeploymentTargetV2ResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateDeploymentTargetV2ResponseBody) SetData(v *DeploymentTarget) *UpdateDeploymentTargetV2ResponseBody {
	s.Data = v
	return s
}

func (s *UpdateDeploymentTargetV2ResponseBody) SetErrorCode(v string) *UpdateDeploymentTargetV2ResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateDeploymentTargetV2ResponseBody) SetErrorMessage(v string) *UpdateDeploymentTargetV2ResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateDeploymentTargetV2ResponseBody) SetHttpCode(v int32) *UpdateDeploymentTargetV2ResponseBody {
	s.HttpCode = &v
	return s
}

func (s *UpdateDeploymentTargetV2ResponseBody) SetRequestId(v string) *UpdateDeploymentTargetV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDeploymentTargetV2ResponseBody) SetSuccess(v bool) *UpdateDeploymentTargetV2ResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateDeploymentTargetV2ResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
