// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRecognizeDataByRuleTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v interface{}) *QueryRecognizeDataByRuleTypeResponseBody
	GetData() interface{}
	SetErrorCode(v string) *QueryRecognizeDataByRuleTypeResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *QueryRecognizeDataByRuleTypeResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *QueryRecognizeDataByRuleTypeResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *QueryRecognizeDataByRuleTypeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryRecognizeDataByRuleTypeResponseBody
	GetSuccess() *bool
}

type QueryRecognizeDataByRuleTypeResponseBody struct {
	// The returned result in the JSON format.
	//
	// example:
	//
	// [{"custom": false,       "name": "EducationDegree",       "localeName": "EducationDegree",       "templateJsonStr": "{&quot;_clazz&quot;:&quot;com.alipay.dsgclient.sdk.dsg.fastscan.engine.cond.EducationDegreeCond&quot;}",       "desc": ""}]
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

func (s QueryRecognizeDataByRuleTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryRecognizeDataByRuleTypeResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRecognizeDataByRuleTypeResponseBody) GetData() interface{} {
	return s.Data
}

func (s *QueryRecognizeDataByRuleTypeResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *QueryRecognizeDataByRuleTypeResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *QueryRecognizeDataByRuleTypeResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *QueryRecognizeDataByRuleTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryRecognizeDataByRuleTypeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryRecognizeDataByRuleTypeResponseBody) SetData(v interface{}) *QueryRecognizeDataByRuleTypeResponseBody {
	s.Data = v
	return s
}

func (s *QueryRecognizeDataByRuleTypeResponseBody) SetErrorCode(v string) *QueryRecognizeDataByRuleTypeResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *QueryRecognizeDataByRuleTypeResponseBody) SetErrorMessage(v string) *QueryRecognizeDataByRuleTypeResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryRecognizeDataByRuleTypeResponseBody) SetHttpStatusCode(v int32) *QueryRecognizeDataByRuleTypeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QueryRecognizeDataByRuleTypeResponseBody) SetRequestId(v string) *QueryRecognizeDataByRuleTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryRecognizeDataByRuleTypeResponseBody) SetSuccess(v bool) *QueryRecognizeDataByRuleTypeResponseBody {
	s.Success = &v
	return s
}

func (s *QueryRecognizeDataByRuleTypeResponseBody) Validate() error {
	return dara.Validate(s)
}
