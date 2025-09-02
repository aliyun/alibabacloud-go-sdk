// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRecognizeRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v interface{}) *DeleteRecognizeRuleResponseBody
	GetData() interface{}
	SetErrorCode(v string) *DeleteRecognizeRuleResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DeleteRecognizeRuleResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *DeleteRecognizeRuleResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *DeleteRecognizeRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteRecognizeRuleResponseBody
	GetSuccess() *bool
}

type DeleteRecognizeRuleResponseBody struct {
	// The returned data about whether the deletion is successful.
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

func (s DeleteRecognizeRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteRecognizeRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRecognizeRuleResponseBody) GetData() interface{} {
	return s.Data
}

func (s *DeleteRecognizeRuleResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteRecognizeRuleResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteRecognizeRuleResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteRecognizeRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteRecognizeRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteRecognizeRuleResponseBody) SetData(v interface{}) *DeleteRecognizeRuleResponseBody {
	s.Data = v
	return s
}

func (s *DeleteRecognizeRuleResponseBody) SetErrorCode(v string) *DeleteRecognizeRuleResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteRecognizeRuleResponseBody) SetErrorMessage(v string) *DeleteRecognizeRuleResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteRecognizeRuleResponseBody) SetHttpStatusCode(v int32) *DeleteRecognizeRuleResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteRecognizeRuleResponseBody) SetRequestId(v string) *DeleteRecognizeRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRecognizeRuleResponseBody) SetSuccess(v bool) *DeleteRecognizeRuleResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteRecognizeRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
