// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *StartJobResponseBody
	GetAccessDeniedDetail() *string
	SetData(v *Job) *StartJobResponseBody
	GetData() *Job
	SetErrorCode(v string) *StartJobResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *StartJobResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *StartJobResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *StartJobResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *StartJobResponseBody
	GetSuccess() *bool
}

type StartJobResponseBody struct {
	// The detailed information about the access denial.
	AccessDeniedDetail *string `json:"accessDeniedDetail,omitempty" xml:"accessDeniedDetail,omitempty"`
	// - When \\`success\\` is \\`true\\`, the created instance information is returned.
	//
	// - Returns an empty value if \\`success\\` is false.
	Data *Job `json:"data,omitempty" xml:"data,omitempty"`
	// - If success is false, the call returns a business error code.If success is true, the call returns business data. The following table describes the data format.If success is true, the call returns business data. The following table describes the data format.
	//
	// - When success is true, an empty value is returned.
	//
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// - If success is false, a business error message is returned.
	//
	// - When \\`success\\` is true, an empty value is returned.
	//
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// Fixed value: 200.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// Request ID.
	//
	// example:
	//
	// CBC799F0-AS7S-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the service request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s StartJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartJobResponseBody) GoString() string {
	return s.String()
}

func (s *StartJobResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *StartJobResponseBody) GetData() *Job {
	return s.Data
}

func (s *StartJobResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *StartJobResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *StartJobResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *StartJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *StartJobResponseBody) SetAccessDeniedDetail(v string) *StartJobResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *StartJobResponseBody) SetData(v *Job) *StartJobResponseBody {
	s.Data = v
	return s
}

func (s *StartJobResponseBody) SetErrorCode(v string) *StartJobResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *StartJobResponseBody) SetErrorMessage(v string) *StartJobResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *StartJobResponseBody) SetHttpCode(v int32) *StartJobResponseBody {
	s.HttpCode = &v
	return s
}

func (s *StartJobResponseBody) SetRequestId(v string) *StartJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartJobResponseBody) SetSuccess(v bool) *StartJobResponseBody {
	s.Success = &v
	return s
}

func (s *StartJobResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
