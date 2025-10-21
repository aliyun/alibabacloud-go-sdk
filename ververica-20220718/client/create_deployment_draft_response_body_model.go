// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDeploymentDraftResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DeploymentDraft) *CreateDeploymentDraftResponseBody
	GetData() *DeploymentDraft
	SetErrorCode(v string) *CreateDeploymentDraftResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateDeploymentDraftResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *CreateDeploymentDraftResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *CreateDeploymentDraftResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateDeploymentDraftResponseBody
	GetSuccess() *bool
}

type CreateDeploymentDraftResponseBody struct {
	Data *DeploymentDraft `json:"data,omitempty" xml:"data,omitempty"`
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

func (s CreateDeploymentDraftResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDeploymentDraftResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDeploymentDraftResponseBody) GetData() *DeploymentDraft {
	return s.Data
}

func (s *CreateDeploymentDraftResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateDeploymentDraftResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateDeploymentDraftResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *CreateDeploymentDraftResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDeploymentDraftResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateDeploymentDraftResponseBody) SetData(v *DeploymentDraft) *CreateDeploymentDraftResponseBody {
	s.Data = v
	return s
}

func (s *CreateDeploymentDraftResponseBody) SetErrorCode(v string) *CreateDeploymentDraftResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateDeploymentDraftResponseBody) SetErrorMessage(v string) *CreateDeploymentDraftResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateDeploymentDraftResponseBody) SetHttpCode(v int32) *CreateDeploymentDraftResponseBody {
	s.HttpCode = &v
	return s
}

func (s *CreateDeploymentDraftResponseBody) SetRequestId(v string) *CreateDeploymentDraftResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDeploymentDraftResponseBody) SetSuccess(v bool) *CreateDeploymentDraftResponseBody {
	s.Success = &v
	return s
}

func (s *CreateDeploymentDraftResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
