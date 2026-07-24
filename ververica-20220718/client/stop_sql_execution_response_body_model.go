// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopSqlExecutionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *StopSqlExecutionResponseBodyData) *StopSqlExecutionResponseBody
	GetData() *StopSqlExecutionResponseBodyData
	SetErrorCode(v string) *StopSqlExecutionResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *StopSqlExecutionResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *StopSqlExecutionResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *StopSqlExecutionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *StopSqlExecutionResponseBody
	GetSuccess() *bool
}

type StopSqlExecutionResponseBody struct {
	// SqlExecutionStopResult
	Data *StopSqlExecutionResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The business error code. This parameter is not empty when success is false. This parameter is empty when success is true.
	//
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// The business error message. This parameter is not empty when success is false. This parameter is empty when success is true.
	//
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The business status code, which is always 200. Use success to determine whether the business request was successful.
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
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s StopSqlExecutionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopSqlExecutionResponseBody) GoString() string {
	return s.String()
}

func (s *StopSqlExecutionResponseBody) GetData() *StopSqlExecutionResponseBodyData {
	return s.Data
}

func (s *StopSqlExecutionResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *StopSqlExecutionResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *StopSqlExecutionResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *StopSqlExecutionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopSqlExecutionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *StopSqlExecutionResponseBody) SetData(v *StopSqlExecutionResponseBodyData) *StopSqlExecutionResponseBody {
	s.Data = v
	return s
}

func (s *StopSqlExecutionResponseBody) SetErrorCode(v string) *StopSqlExecutionResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *StopSqlExecutionResponseBody) SetErrorMessage(v string) *StopSqlExecutionResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *StopSqlExecutionResponseBody) SetHttpCode(v int32) *StopSqlExecutionResponseBody {
	s.HttpCode = &v
	return s
}

func (s *StopSqlExecutionResponseBody) SetRequestId(v string) *StopSqlExecutionResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopSqlExecutionResponseBody) SetSuccess(v bool) *StopSqlExecutionResponseBody {
	s.Success = &v
	return s
}

func (s *StopSqlExecutionResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type StopSqlExecutionResponseBodyData struct {
	// The supplementary information about stopping the query script.
	//
	// example:
	//
	// “”
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Indicates whether the stop instruction was successfully initiated.
	//
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s StopSqlExecutionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s StopSqlExecutionResponseBodyData) GoString() string {
	return s.String()
}

func (s *StopSqlExecutionResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *StopSqlExecutionResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *StopSqlExecutionResponseBodyData) SetMessage(v string) *StopSqlExecutionResponseBodyData {
	s.Message = &v
	return s
}

func (s *StopSqlExecutionResponseBodyData) SetSuccess(v bool) *StopSqlExecutionResponseBodyData {
	s.Success = &v
	return s
}

func (s *StopSqlExecutionResponseBodyData) Validate() error {
	return dara.Validate(s)
}
