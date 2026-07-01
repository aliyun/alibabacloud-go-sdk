// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetValidateDeploymentDraftResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *AsyncDraftValidateResult) *GetValidateDeploymentDraftResultResponseBody
	GetData() *AsyncDraftValidateResult
	SetErrorCode(v string) *GetValidateDeploymentDraftResultResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetValidateDeploymentDraftResultResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *GetValidateDeploymentDraftResultResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *GetValidateDeploymentDraftResultResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetValidateDeploymentDraftResultResponseBody
	GetSuccess() *bool
}

type GetValidateDeploymentDraftResultResponseBody struct {
	// The result object of the in-depth check for the job draft.
	Data *AsyncDraftValidateResult `json:"data,omitempty" xml:"data,omitempty"`
	// If success is false, this parameter is not empty and indicates the business error code. If success is true, this parameter is empty.
	//
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// If success is false, this parameter is not empty and indicates the business error message. If success is true, this parameter is empty.
	//
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The business status code. This is always 200. The success parameter indicates whether the business request was successful.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CBC799F0-AS7S-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the business request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetValidateDeploymentDraftResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetValidateDeploymentDraftResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetValidateDeploymentDraftResultResponseBody) GetData() *AsyncDraftValidateResult {
	return s.Data
}

func (s *GetValidateDeploymentDraftResultResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetValidateDeploymentDraftResultResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetValidateDeploymentDraftResultResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *GetValidateDeploymentDraftResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetValidateDeploymentDraftResultResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetValidateDeploymentDraftResultResponseBody) SetData(v *AsyncDraftValidateResult) *GetValidateDeploymentDraftResultResponseBody {
	s.Data = v
	return s
}

func (s *GetValidateDeploymentDraftResultResponseBody) SetErrorCode(v string) *GetValidateDeploymentDraftResultResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetValidateDeploymentDraftResultResponseBody) SetErrorMessage(v string) *GetValidateDeploymentDraftResultResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetValidateDeploymentDraftResultResponseBody) SetHttpCode(v int32) *GetValidateDeploymentDraftResultResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetValidateDeploymentDraftResultResponseBody) SetRequestId(v string) *GetValidateDeploymentDraftResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetValidateDeploymentDraftResultResponseBody) SetSuccess(v bool) *GetValidateDeploymentDraftResultResponseBody {
	s.Success = &v
	return s
}

func (s *GetValidateDeploymentDraftResultResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
