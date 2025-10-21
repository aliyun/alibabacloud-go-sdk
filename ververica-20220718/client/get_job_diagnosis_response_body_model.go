// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJobDiagnosisResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetJobDiagnosisResponseBody
	GetAccessDeniedDetail() *string
	SetData(v *JobDiagnosis) *GetJobDiagnosisResponseBody
	GetData() *JobDiagnosis
	SetErrorCode(v string) *GetJobDiagnosisResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetJobDiagnosisResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *GetJobDiagnosisResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *GetJobDiagnosisResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetJobDiagnosisResponseBody
	GetSuccess() *bool
}

type GetJobDiagnosisResponseBody struct {
	// example:
	//
	// “”
	AccessDeniedDetail *string       `json:"accessDeniedDetail,omitempty" xml:"accessDeniedDetail,omitempty"`
	Data               *JobDiagnosis `json:"data,omitempty" xml:"data,omitempty"`
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

func (s GetJobDiagnosisResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetJobDiagnosisResponseBody) GoString() string {
	return s.String()
}

func (s *GetJobDiagnosisResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetJobDiagnosisResponseBody) GetData() *JobDiagnosis {
	return s.Data
}

func (s *GetJobDiagnosisResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetJobDiagnosisResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetJobDiagnosisResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *GetJobDiagnosisResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetJobDiagnosisResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetJobDiagnosisResponseBody) SetAccessDeniedDetail(v string) *GetJobDiagnosisResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetJobDiagnosisResponseBody) SetData(v *JobDiagnosis) *GetJobDiagnosisResponseBody {
	s.Data = v
	return s
}

func (s *GetJobDiagnosisResponseBody) SetErrorCode(v string) *GetJobDiagnosisResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetJobDiagnosisResponseBody) SetErrorMessage(v string) *GetJobDiagnosisResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetJobDiagnosisResponseBody) SetHttpCode(v int32) *GetJobDiagnosisResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetJobDiagnosisResponseBody) SetRequestId(v string) *GetJobDiagnosisResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetJobDiagnosisResponseBody) SetSuccess(v bool) *GetJobDiagnosisResponseBody {
	s.Success = &v
	return s
}

func (s *GetJobDiagnosisResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
