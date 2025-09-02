// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRecognizeRulesTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v interface{}) *QueryRecognizeRulesTypeResponseBody
	GetData() interface{}
	SetErrorCode(v string) *QueryRecognizeRulesTypeResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *QueryRecognizeRulesTypeResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *QueryRecognizeRulesTypeResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *QueryRecognizeRulesTypeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryRecognizeRulesTypeResponseBody
	GetSuccess() *bool
}

type QueryRecognizeRulesTypeResponseBody struct {
	// The returned data about the built-in sensitive data identification rule that is used to configure a sensitive field. The data is in the JSON format.
	//
	// example:
	//
	// {   "HttpStatusCode": 200,   "Data": [     {       "Regular Expression": "1"     },     {       "Recognize Rule": "2"     },     {       "Sample Library": "3"     },     {       "Model": "4"     }   ],   "Success": true }
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
	// Indicates whether the request was successful. Valid values:
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

func (s QueryRecognizeRulesTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryRecognizeRulesTypeResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRecognizeRulesTypeResponseBody) GetData() interface{} {
	return s.Data
}

func (s *QueryRecognizeRulesTypeResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *QueryRecognizeRulesTypeResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *QueryRecognizeRulesTypeResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *QueryRecognizeRulesTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryRecognizeRulesTypeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryRecognizeRulesTypeResponseBody) SetData(v interface{}) *QueryRecognizeRulesTypeResponseBody {
	s.Data = v
	return s
}

func (s *QueryRecognizeRulesTypeResponseBody) SetErrorCode(v string) *QueryRecognizeRulesTypeResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *QueryRecognizeRulesTypeResponseBody) SetErrorMessage(v string) *QueryRecognizeRulesTypeResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryRecognizeRulesTypeResponseBody) SetHttpStatusCode(v int32) *QueryRecognizeRulesTypeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QueryRecognizeRulesTypeResponseBody) SetRequestId(v string) *QueryRecognizeRulesTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryRecognizeRulesTypeResponseBody) SetSuccess(v bool) *QueryRecognizeRulesTypeResponseBody {
	s.Success = &v
	return s
}

func (s *QueryRecognizeRulesTypeResponseBody) Validate() error {
	return dara.Validate(s)
}
