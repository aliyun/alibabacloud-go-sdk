// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *StopJobResponseBody
	GetAccessDeniedDetail() *string
	SetData(v *Job) *StopJobResponseBody
	GetData() *Job
	SetErrorCode(v string) *StopJobResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *StopJobResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *StopJobResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *StopJobResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *StopJobResponseBody
	GetSuccess() *bool
}

type StopJobResponseBody struct {
	AccessDeniedDetail *string `json:"accessDeniedDetail,omitempty" xml:"accessDeniedDetail,omitempty"`
	// 	- If the value of success was true, the job that you stopped was returned.
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

func (s StopJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopJobResponseBody) GoString() string {
	return s.String()
}

func (s *StopJobResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *StopJobResponseBody) GetData() *Job {
	return s.Data
}

func (s *StopJobResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *StopJobResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *StopJobResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *StopJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *StopJobResponseBody) SetAccessDeniedDetail(v string) *StopJobResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *StopJobResponseBody) SetData(v *Job) *StopJobResponseBody {
	s.Data = v
	return s
}

func (s *StopJobResponseBody) SetErrorCode(v string) *StopJobResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *StopJobResponseBody) SetErrorMessage(v string) *StopJobResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *StopJobResponseBody) SetHttpCode(v int32) *StopJobResponseBody {
	s.HttpCode = &v
	return s
}

func (s *StopJobResponseBody) SetRequestId(v string) *StopJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopJobResponseBody) SetSuccess(v bool) *StopJobResponseBody {
	s.Success = &v
	return s
}

func (s *StopJobResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
