// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddRecognizeRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v interface{}) *AddRecognizeRuleResponseBody
	GetData() interface{}
	SetErrorCode(v string) *AddRecognizeRuleResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *AddRecognizeRuleResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *AddRecognizeRuleResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *AddRecognizeRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddRecognizeRuleResponseBody
	GetSuccess() *bool
}

type AddRecognizeRuleResponseBody struct {
	// The returned result in the JSON format.
	//
	// example:
	//
	// {   "HttpStatusCode": 200,   "Success": true }
	Data interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error code.
	//
	// example:
	//
	// 9990030003
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
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 10000001
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

func (s AddRecognizeRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddRecognizeRuleResponseBody) GoString() string {
	return s.String()
}

func (s *AddRecognizeRuleResponseBody) GetData() interface{} {
	return s.Data
}

func (s *AddRecognizeRuleResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *AddRecognizeRuleResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *AddRecognizeRuleResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *AddRecognizeRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddRecognizeRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddRecognizeRuleResponseBody) SetData(v interface{}) *AddRecognizeRuleResponseBody {
	s.Data = v
	return s
}

func (s *AddRecognizeRuleResponseBody) SetErrorCode(v string) *AddRecognizeRuleResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *AddRecognizeRuleResponseBody) SetErrorMessage(v string) *AddRecognizeRuleResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *AddRecognizeRuleResponseBody) SetHttpStatusCode(v int32) *AddRecognizeRuleResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AddRecognizeRuleResponseBody) SetRequestId(v string) *AddRecognizeRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddRecognizeRuleResponseBody) SetSuccess(v bool) *AddRecognizeRuleResponseBody {
	s.Success = &v
	return s
}

func (s *AddRecognizeRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
