// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotUpdateJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *HotUpdateJobResult) *HotUpdateJobResponseBody
	GetData() *HotUpdateJobResult
	SetErrorCode(v string) *HotUpdateJobResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *HotUpdateJobResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *HotUpdateJobResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *HotUpdateJobResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *HotUpdateJobResponseBody
	GetSuccess() *bool
}

type HotUpdateJobResponseBody struct {
	// The dynamic update result.
	Data *HotUpdateJobResult `json:"data,omitempty" xml:"data,omitempty"`
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

func (s HotUpdateJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s HotUpdateJobResponseBody) GoString() string {
	return s.String()
}

func (s *HotUpdateJobResponseBody) GetData() *HotUpdateJobResult {
	return s.Data
}

func (s *HotUpdateJobResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *HotUpdateJobResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *HotUpdateJobResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *HotUpdateJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *HotUpdateJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *HotUpdateJobResponseBody) SetData(v *HotUpdateJobResult) *HotUpdateJobResponseBody {
	s.Data = v
	return s
}

func (s *HotUpdateJobResponseBody) SetErrorCode(v string) *HotUpdateJobResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *HotUpdateJobResponseBody) SetErrorMessage(v string) *HotUpdateJobResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *HotUpdateJobResponseBody) SetHttpCode(v int32) *HotUpdateJobResponseBody {
	s.HttpCode = &v
	return s
}

func (s *HotUpdateJobResponseBody) SetRequestId(v string) *HotUpdateJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *HotUpdateJobResponseBody) SetSuccess(v bool) *HotUpdateJobResponseBody {
	s.Success = &v
	return s
}

func (s *HotUpdateJobResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
