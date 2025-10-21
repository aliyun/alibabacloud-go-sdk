// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeploymentDraftResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DeploymentDraft) *GetDeploymentDraftResponseBody
	GetData() *DeploymentDraft
	SetErrorCode(v string) *GetDeploymentDraftResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetDeploymentDraftResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *GetDeploymentDraftResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *GetDeploymentDraftResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDeploymentDraftResponseBody
	GetSuccess() *bool
}

type GetDeploymentDraftResponseBody struct {
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

func (s GetDeploymentDraftResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDeploymentDraftResponseBody) GoString() string {
	return s.String()
}

func (s *GetDeploymentDraftResponseBody) GetData() *DeploymentDraft {
	return s.Data
}

func (s *GetDeploymentDraftResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetDeploymentDraftResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetDeploymentDraftResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *GetDeploymentDraftResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDeploymentDraftResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDeploymentDraftResponseBody) SetData(v *DeploymentDraft) *GetDeploymentDraftResponseBody {
	s.Data = v
	return s
}

func (s *GetDeploymentDraftResponseBody) SetErrorCode(v string) *GetDeploymentDraftResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetDeploymentDraftResponseBody) SetErrorMessage(v string) *GetDeploymentDraftResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetDeploymentDraftResponseBody) SetHttpCode(v int32) *GetDeploymentDraftResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetDeploymentDraftResponseBody) SetRequestId(v string) *GetDeploymentDraftResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDeploymentDraftResponseBody) SetSuccess(v bool) *GetDeploymentDraftResponseBody {
	s.Success = &v
	return s
}

func (s *GetDeploymentDraftResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
