// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTestAgenticDataSourceConnectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *TestAgenticDataSourceConnectionResponseBodyData) *TestAgenticDataSourceConnectionResponseBody
	GetData() *TestAgenticDataSourceConnectionResponseBodyData
	SetErrorCode(v string) *TestAgenticDataSourceConnectionResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *TestAgenticDataSourceConnectionResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *TestAgenticDataSourceConnectionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *TestAgenticDataSourceConnectionResponseBody
	GetSuccess() *bool
}

type TestAgenticDataSourceConnectionResponseBody struct {
	// The connection test result.
	Data *TestAgenticDataSourceConnectionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code returned when the call fails. This parameter is empty when the call is successful.
	//
	// example:
	//
	// InvalidParameter
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned when the call fails. This parameter is empty when the call is successful.
	//
	// example:
	//
	// The parameter is invalid
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The request ID, which is used to locate the current call.
	//
	// example:
	//
	// request-1234567890
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the current call is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s TestAgenticDataSourceConnectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TestAgenticDataSourceConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *TestAgenticDataSourceConnectionResponseBody) GetData() *TestAgenticDataSourceConnectionResponseBodyData {
	return s.Data
}

func (s *TestAgenticDataSourceConnectionResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *TestAgenticDataSourceConnectionResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *TestAgenticDataSourceConnectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TestAgenticDataSourceConnectionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *TestAgenticDataSourceConnectionResponseBody) SetData(v *TestAgenticDataSourceConnectionResponseBodyData) *TestAgenticDataSourceConnectionResponseBody {
	s.Data = v
	return s
}

func (s *TestAgenticDataSourceConnectionResponseBody) SetErrorCode(v string) *TestAgenticDataSourceConnectionResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *TestAgenticDataSourceConnectionResponseBody) SetErrorMessage(v string) *TestAgenticDataSourceConnectionResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *TestAgenticDataSourceConnectionResponseBody) SetRequestId(v string) *TestAgenticDataSourceConnectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *TestAgenticDataSourceConnectionResponseBody) SetSuccess(v bool) *TestAgenticDataSourceConnectionResponseBody {
	s.Success = &v
	return s
}

func (s *TestAgenticDataSourceConnectionResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type TestAgenticDataSourceConnectionResponseBodyData struct {
	// The description of the connection test result. This parameter is usually empty when the connection is successful, and returns the cause when the connection fails.
	//
	// example:
	//
	// Connection timed out. Check the connection endpoint
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Indicates whether the data source is connected.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s TestAgenticDataSourceConnectionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s TestAgenticDataSourceConnectionResponseBodyData) GoString() string {
	return s.String()
}

func (s *TestAgenticDataSourceConnectionResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *TestAgenticDataSourceConnectionResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *TestAgenticDataSourceConnectionResponseBodyData) SetMessage(v string) *TestAgenticDataSourceConnectionResponseBodyData {
	s.Message = &v
	return s
}

func (s *TestAgenticDataSourceConnectionResponseBodyData) SetSuccess(v bool) *TestAgenticDataSourceConnectionResponseBodyData {
	s.Success = &v
	return s
}

func (s *TestAgenticDataSourceConnectionResponseBodyData) Validate() error {
	return dara.Validate(s)
}
