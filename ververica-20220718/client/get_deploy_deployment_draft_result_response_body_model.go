// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeployDeploymentDraftResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *AsyncDraftDeployResult) *GetDeployDeploymentDraftResultResponseBody
	GetData() *AsyncDraftDeployResult
	SetErrorCode(v string) *GetDeployDeploymentDraftResultResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetDeployDeploymentDraftResultResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *GetDeployDeploymentDraftResultResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *GetDeployDeploymentDraftResultResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDeployDeploymentDraftResultResponseBody
	GetSuccess() *bool
}

type GetDeployDeploymentDraftResultResponseBody struct {
	Data *AsyncDraftDeployResult `json:"data,omitempty" xml:"data,omitempty"`
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

func (s GetDeployDeploymentDraftResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDeployDeploymentDraftResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetDeployDeploymentDraftResultResponseBody) GetData() *AsyncDraftDeployResult {
	return s.Data
}

func (s *GetDeployDeploymentDraftResultResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetDeployDeploymentDraftResultResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetDeployDeploymentDraftResultResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *GetDeployDeploymentDraftResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDeployDeploymentDraftResultResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDeployDeploymentDraftResultResponseBody) SetData(v *AsyncDraftDeployResult) *GetDeployDeploymentDraftResultResponseBody {
	s.Data = v
	return s
}

func (s *GetDeployDeploymentDraftResultResponseBody) SetErrorCode(v string) *GetDeployDeploymentDraftResultResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetDeployDeploymentDraftResultResponseBody) SetErrorMessage(v string) *GetDeployDeploymentDraftResultResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetDeployDeploymentDraftResultResponseBody) SetHttpCode(v int32) *GetDeployDeploymentDraftResultResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetDeployDeploymentDraftResultResponseBody) SetRequestId(v string) *GetDeployDeploymentDraftResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDeployDeploymentDraftResultResponseBody) SetSuccess(v bool) *GetDeployDeploymentDraftResultResponseBody {
	s.Success = &v
	return s
}

func (s *GetDeployDeploymentDraftResultResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
