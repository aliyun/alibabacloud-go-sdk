// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDeploymentDraftResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DeploymentDraft) *UpdateDeploymentDraftResponseBody
	GetData() *DeploymentDraft
	SetErrorCode(v string) *UpdateDeploymentDraftResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *UpdateDeploymentDraftResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *UpdateDeploymentDraftResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *UpdateDeploymentDraftResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateDeploymentDraftResponseBody
	GetSuccess() *bool
}

type UpdateDeploymentDraftResponseBody struct {
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
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateDeploymentDraftResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDeploymentDraftResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDeploymentDraftResponseBody) GetData() *DeploymentDraft {
	return s.Data
}

func (s *UpdateDeploymentDraftResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateDeploymentDraftResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateDeploymentDraftResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *UpdateDeploymentDraftResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDeploymentDraftResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateDeploymentDraftResponseBody) SetData(v *DeploymentDraft) *UpdateDeploymentDraftResponseBody {
	s.Data = v
	return s
}

func (s *UpdateDeploymentDraftResponseBody) SetErrorCode(v string) *UpdateDeploymentDraftResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateDeploymentDraftResponseBody) SetErrorMessage(v string) *UpdateDeploymentDraftResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateDeploymentDraftResponseBody) SetHttpCode(v int32) *UpdateDeploymentDraftResponseBody {
	s.HttpCode = &v
	return s
}

func (s *UpdateDeploymentDraftResponseBody) SetRequestId(v string) *UpdateDeploymentDraftResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDeploymentDraftResponseBody) SetSuccess(v bool) *UpdateDeploymentDraftResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateDeploymentDraftResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
