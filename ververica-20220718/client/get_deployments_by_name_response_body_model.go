// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeploymentsByNameResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*Deployment) *GetDeploymentsByNameResponseBody
	GetData() []*Deployment
	SetErrorCode(v string) *GetDeploymentsByNameResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetDeploymentsByNameResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *GetDeploymentsByNameResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *GetDeploymentsByNameResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDeploymentsByNameResponseBody
	GetSuccess() *bool
}

type GetDeploymentsByNameResponseBody struct {
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

func (s GetDeploymentsByNameResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDeploymentsByNameResponseBody) GoString() string {
	return s.String()
}

func (s *GetDeploymentsByNameResponseBody) GetData() []*Deployment {
	return s.Data
}

func (s *GetDeploymentsByNameResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetDeploymentsByNameResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetDeploymentsByNameResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *GetDeploymentsByNameResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDeploymentsByNameResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDeploymentsByNameResponseBody) SetData(v []*Deployment) *GetDeploymentsByNameResponseBody {
	s.Data = v
	return s
}

func (s *GetDeploymentsByNameResponseBody) SetErrorCode(v string) *GetDeploymentsByNameResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetDeploymentsByNameResponseBody) SetErrorMessage(v string) *GetDeploymentsByNameResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetDeploymentsByNameResponseBody) SetHttpCode(v int32) *GetDeploymentsByNameResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetDeploymentsByNameResponseBody) SetRequestId(v string) *GetDeploymentsByNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDeploymentsByNameResponseBody) SetSuccess(v bool) *GetDeploymentsByNameResponseBody {
	s.Success = &v
	return s
}

func (s *GetDeploymentsByNameResponseBody) Validate() error {
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
