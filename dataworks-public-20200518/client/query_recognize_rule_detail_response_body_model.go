// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRecognizeRuleDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v interface{}) *QueryRecognizeRuleDetailResponseBody
	GetData() interface{}
	SetErrorCode(v string) *QueryRecognizeRuleDetailResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *QueryRecognizeRuleDetailResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *QueryRecognizeRuleDetailResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *QueryRecognizeRuleDetailResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryRecognizeRuleDetailResponseBody
	GetSuccess() *bool
}

type QueryRecognizeRuleDetailResponseBody struct {
	// The details of the sensitive field in the JSON format.
	//
	// example:
	//
	// {"nodeName":"basic unit overview","gmtModified":1709017449000,"hitThreshold":30,"sensitiveName":"mobile-yinni","templateId":"8222abeb-9784-4417-b420-0322926d5cbf","recognizeRulesType":2,"delete":false,"bydAccuracy":1,"colScan":"," defineType ":1,": ydAccuracy ":{" contentRule ":))," operationType ":0}," nodeParent ":" unit/unit basic information/unit basic overview "," level ":6," keyRuleId ":" 228248921215042mobile-yinni "," isDelete ":false," levelName ":" 6level "," sensitive ":false," operationType ":0," sourceName ": dsg-test-zuoyue","nodeId":"bea2fc81-90c9-45f3-b7a9-26d534208a0d","status":1}
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
	// The request ID. You can locate logs and troubleshoot issues based on the ID.
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

func (s QueryRecognizeRuleDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryRecognizeRuleDetailResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRecognizeRuleDetailResponseBody) GetData() interface{} {
	return s.Data
}

func (s *QueryRecognizeRuleDetailResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *QueryRecognizeRuleDetailResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *QueryRecognizeRuleDetailResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *QueryRecognizeRuleDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryRecognizeRuleDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryRecognizeRuleDetailResponseBody) SetData(v interface{}) *QueryRecognizeRuleDetailResponseBody {
	s.Data = v
	return s
}

func (s *QueryRecognizeRuleDetailResponseBody) SetErrorCode(v string) *QueryRecognizeRuleDetailResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *QueryRecognizeRuleDetailResponseBody) SetErrorMessage(v string) *QueryRecognizeRuleDetailResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryRecognizeRuleDetailResponseBody) SetHttpStatusCode(v int32) *QueryRecognizeRuleDetailResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QueryRecognizeRuleDetailResponseBody) SetRequestId(v string) *QueryRecognizeRuleDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryRecognizeRuleDetailResponseBody) SetSuccess(v bool) *QueryRecognizeRuleDetailResponseBody {
	s.Success = &v
	return s
}

func (s *QueryRecognizeRuleDetailResponseBody) Validate() error {
	return dara.Validate(s)
}
