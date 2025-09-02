// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDsgQuerySensResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v interface{}) *DsgQuerySensResultResponseBody
	GetData() interface{}
	SetDynamicErrorCode(v string) *DsgQuerySensResultResponseBody
	GetDynamicErrorCode() *string
	SetDynamicErrorMessage(v string) *DsgQuerySensResultResponseBody
	GetDynamicErrorMessage() *string
	SetErrorCode(v string) *DsgQuerySensResultResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DsgQuerySensResultResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *DsgQuerySensResultResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *DsgQuerySensResultResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DsgQuerySensResultResponseBody
	GetSuccess() *bool
}

type DsgQuerySensResultResponseBody struct {
	// The data returned.
	//
	// example:
	//
	// 1234
	Data interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error code. The value is the same as that of ErrorCode.
	//
	// example:
	//
	// 400
	DynamicErrorCode *string `json:"DynamicErrorCode,omitempty" xml:"DynamicErrorCode,omitempty"`
	// The error message. The value is the same as that of ErrorMessage.
	//
	// example:
	//
	// Missing parameter
	DynamicErrorMessage *string `json:"DynamicErrorMessage,omitempty" xml:"DynamicErrorMessage,omitempty"`
	// The error code.
	//
	// example:
	//
	// 400
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Missing parameter
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 400
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 900000001
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DsgQuerySensResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DsgQuerySensResultResponseBody) GoString() string {
	return s.String()
}

func (s *DsgQuerySensResultResponseBody) GetData() interface{} {
	return s.Data
}

func (s *DsgQuerySensResultResponseBody) GetDynamicErrorCode() *string {
	return s.DynamicErrorCode
}

func (s *DsgQuerySensResultResponseBody) GetDynamicErrorMessage() *string {
	return s.DynamicErrorMessage
}

func (s *DsgQuerySensResultResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DsgQuerySensResultResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DsgQuerySensResultResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DsgQuerySensResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DsgQuerySensResultResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DsgQuerySensResultResponseBody) SetData(v interface{}) *DsgQuerySensResultResponseBody {
	s.Data = v
	return s
}

func (s *DsgQuerySensResultResponseBody) SetDynamicErrorCode(v string) *DsgQuerySensResultResponseBody {
	s.DynamicErrorCode = &v
	return s
}

func (s *DsgQuerySensResultResponseBody) SetDynamicErrorMessage(v string) *DsgQuerySensResultResponseBody {
	s.DynamicErrorMessage = &v
	return s
}

func (s *DsgQuerySensResultResponseBody) SetErrorCode(v string) *DsgQuerySensResultResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DsgQuerySensResultResponseBody) SetErrorMessage(v string) *DsgQuerySensResultResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DsgQuerySensResultResponseBody) SetHttpStatusCode(v int32) *DsgQuerySensResultResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DsgQuerySensResultResponseBody) SetRequestId(v string) *DsgQuerySensResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *DsgQuerySensResultResponseBody) SetSuccess(v bool) *DsgQuerySensResultResponseBody {
	s.Success = &v
	return s
}

func (s *DsgQuerySensResultResponseBody) Validate() error {
	return dara.Validate(s)
}
