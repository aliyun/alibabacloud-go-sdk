// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetJobResponseBody
	GetAccessDeniedDetail() *string
	SetData(v *Job) *GetJobResponseBody
	GetData() *Job
	SetErrorCode(v string) *GetJobResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetJobResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *GetJobResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *GetJobResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetJobResponseBody
	GetSuccess() *bool
}

type GetJobResponseBody struct {
	AccessDeniedDetail *string `json:"accessDeniedDetail,omitempty" xml:"accessDeniedDetail,omitempty"`
	// 	- If the value of success was true, the details of the job was returned.
	//
	// 	- If the value of success was false, a null value was returned.
	Data *Job `json:"data,omitempty" xml:"data,omitempty"`
	// 	- If the value of success was false, an error code was returned.
	//
	// 	- If the value of success was true, a null value was returned.
	//
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 	- If the value of success was false, an error message was returned.
	//
	// 	- If the value of success was true, a null value was returned.
	//
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The value was fixed to 200.
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

func (s GetJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetJobResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetJobResponseBody) GetData() *Job {
	return s.Data
}

func (s *GetJobResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetJobResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetJobResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *GetJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetJobResponseBody) SetAccessDeniedDetail(v string) *GetJobResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetJobResponseBody) SetData(v *Job) *GetJobResponseBody {
	s.Data = v
	return s
}

func (s *GetJobResponseBody) SetErrorCode(v string) *GetJobResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetJobResponseBody) SetErrorMessage(v string) *GetJobResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetJobResponseBody) SetHttpCode(v int32) *GetJobResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetJobResponseBody) SetRequestId(v string) *GetJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetJobResponseBody) SetSuccess(v bool) *GetJobResponseBody {
	s.Success = &v
	return s
}

func (s *GetJobResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
