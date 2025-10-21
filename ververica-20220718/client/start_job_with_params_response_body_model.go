// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartJobWithParamsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *StartJobWithParamsResponseBody
	GetAccessDeniedDetail() *string
	SetData(v *Job) *StartJobWithParamsResponseBody
	GetData() *Job
	SetErrorCode(v string) *StartJobWithParamsResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *StartJobWithParamsResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *StartJobWithParamsResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *StartJobWithParamsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *StartJobWithParamsResponseBody
	GetSuccess() *bool
}

type StartJobWithParamsResponseBody struct {
	AccessDeniedDetail *string `json:"accessDeniedDetail,omitempty" xml:"accessDeniedDetail,omitempty"`
	// The details of the job of the deployment returned.
	Data *Job `json:"data,omitempty" xml:"data,omitempty"`
	// If the value of success was false, an error code was returned. If the value of success was true, a null value was returned.
	//
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// If the value of success was false, an error message was returned. If the value of success was true, a null value was returned.
	//
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The status code returned. The value was fixed to 200. The status code 200 indicates that the request was successful.
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
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s StartJobWithParamsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartJobWithParamsResponseBody) GoString() string {
	return s.String()
}

func (s *StartJobWithParamsResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *StartJobWithParamsResponseBody) GetData() *Job {
	return s.Data
}

func (s *StartJobWithParamsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *StartJobWithParamsResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *StartJobWithParamsResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *StartJobWithParamsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartJobWithParamsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *StartJobWithParamsResponseBody) SetAccessDeniedDetail(v string) *StartJobWithParamsResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *StartJobWithParamsResponseBody) SetData(v *Job) *StartJobWithParamsResponseBody {
	s.Data = v
	return s
}

func (s *StartJobWithParamsResponseBody) SetErrorCode(v string) *StartJobWithParamsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *StartJobWithParamsResponseBody) SetErrorMessage(v string) *StartJobWithParamsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *StartJobWithParamsResponseBody) SetHttpCode(v int32) *StartJobWithParamsResponseBody {
	s.HttpCode = &v
	return s
}

func (s *StartJobWithParamsResponseBody) SetRequestId(v string) *StartJobWithParamsResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartJobWithParamsResponseBody) SetSuccess(v bool) *StartJobWithParamsResponseBody {
	s.Success = &v
	return s
}

func (s *StartJobWithParamsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
