// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeploymentsByIpResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*Deployment) *GetDeploymentsByIpResponseBody
	GetData() []*Deployment
	SetErrorCode(v string) *GetDeploymentsByIpResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetDeploymentsByIpResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *GetDeploymentsByIpResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *GetDeploymentsByIpResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDeploymentsByIpResponseBody
	GetSuccess() *bool
}

type GetDeploymentsByIpResponseBody struct {
	Data []*Deployment `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
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

func (s GetDeploymentsByIpResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDeploymentsByIpResponseBody) GoString() string {
	return s.String()
}

func (s *GetDeploymentsByIpResponseBody) GetData() []*Deployment {
	return s.Data
}

func (s *GetDeploymentsByIpResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetDeploymentsByIpResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetDeploymentsByIpResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *GetDeploymentsByIpResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDeploymentsByIpResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDeploymentsByIpResponseBody) SetData(v []*Deployment) *GetDeploymentsByIpResponseBody {
	s.Data = v
	return s
}

func (s *GetDeploymentsByIpResponseBody) SetErrorCode(v string) *GetDeploymentsByIpResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetDeploymentsByIpResponseBody) SetErrorMessage(v string) *GetDeploymentsByIpResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetDeploymentsByIpResponseBody) SetHttpCode(v int32) *GetDeploymentsByIpResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetDeploymentsByIpResponseBody) SetRequestId(v string) *GetDeploymentsByIpResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDeploymentsByIpResponseBody) SetSuccess(v bool) *GetDeploymentsByIpResponseBody {
	s.Success = &v
	return s
}

func (s *GetDeploymentsByIpResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
