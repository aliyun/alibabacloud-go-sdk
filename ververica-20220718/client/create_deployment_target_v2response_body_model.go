// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDeploymentTargetV2ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DeploymentTarget) *CreateDeploymentTargetV2ResponseBody
	GetData() *DeploymentTarget
	SetErrorCode(v string) *CreateDeploymentTargetV2ResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateDeploymentTargetV2ResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *CreateDeploymentTargetV2ResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *CreateDeploymentTargetV2ResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateDeploymentTargetV2ResponseBody
	GetSuccess() *bool
}

type CreateDeploymentTargetV2ResponseBody struct {
	// example:
	//
	// { "jobs": [ { "jid": "4df35f8e54554b23bf7dcd38a151****", "name": "69d001d5-419a-4bfc-9c2e-849cacd3****", "state": "RUNNING", "start-time": 1659154942068, "end-time": -1, "duration": 188161756, "last-modification": 1659154968305, "tasks": { "total": 2, "created": 0, "scheduled": 0, "deploying": 0, "running": 2, "finished": 0, "canceling": 0, "canceled": 0, "failed": 0, "reconciling": 0, "initializing": 0 } } ] }
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

func (s CreateDeploymentTargetV2ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDeploymentTargetV2ResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDeploymentTargetV2ResponseBody) GetData() *DeploymentTarget {
	return s.Data
}

func (s *CreateDeploymentTargetV2ResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateDeploymentTargetV2ResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateDeploymentTargetV2ResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *CreateDeploymentTargetV2ResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDeploymentTargetV2ResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateDeploymentTargetV2ResponseBody) SetData(v *DeploymentTarget) *CreateDeploymentTargetV2ResponseBody {
	s.Data = v
	return s
}

func (s *CreateDeploymentTargetV2ResponseBody) SetErrorCode(v string) *CreateDeploymentTargetV2ResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateDeploymentTargetV2ResponseBody) SetErrorMessage(v string) *CreateDeploymentTargetV2ResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateDeploymentTargetV2ResponseBody) SetHttpCode(v int32) *CreateDeploymentTargetV2ResponseBody {
	s.HttpCode = &v
	return s
}

func (s *CreateDeploymentTargetV2ResponseBody) SetRequestId(v string) *CreateDeploymentTargetV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDeploymentTargetV2ResponseBody) SetSuccess(v bool) *CreateDeploymentTargetV2ResponseBody {
	s.Success = &v
	return s
}

func (s *CreateDeploymentTargetV2ResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
