// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeploymentsByLabelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*Deployment) *GetDeploymentsByLabelResponseBody
	GetData() []*Deployment
	SetErrorCode(v string) *GetDeploymentsByLabelResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetDeploymentsByLabelResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *GetDeploymentsByLabelResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *GetDeploymentsByLabelResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDeploymentsByLabelResponseBody
	GetSuccess() *bool
}

type GetDeploymentsByLabelResponseBody struct {
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

func (s GetDeploymentsByLabelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDeploymentsByLabelResponseBody) GoString() string {
	return s.String()
}

func (s *GetDeploymentsByLabelResponseBody) GetData() []*Deployment {
	return s.Data
}

func (s *GetDeploymentsByLabelResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetDeploymentsByLabelResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetDeploymentsByLabelResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *GetDeploymentsByLabelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDeploymentsByLabelResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDeploymentsByLabelResponseBody) SetData(v []*Deployment) *GetDeploymentsByLabelResponseBody {
	s.Data = v
	return s
}

func (s *GetDeploymentsByLabelResponseBody) SetErrorCode(v string) *GetDeploymentsByLabelResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetDeploymentsByLabelResponseBody) SetErrorMessage(v string) *GetDeploymentsByLabelResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetDeploymentsByLabelResponseBody) SetHttpCode(v int32) *GetDeploymentsByLabelResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetDeploymentsByLabelResponseBody) SetRequestId(v string) *GetDeploymentsByLabelResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDeploymentsByLabelResponseBody) SetSuccess(v bool) *GetDeploymentsByLabelResponseBody {
	s.Success = &v
	return s
}

func (s *GetDeploymentsByLabelResponseBody) Validate() error {
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
