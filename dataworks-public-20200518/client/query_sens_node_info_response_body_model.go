// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySensNodeInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v interface{}) *QuerySensNodeInfoResponseBody
	GetData() interface{}
	SetErrorCode(v string) *QuerySensNodeInfoResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *QuerySensNodeInfoResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *QuerySensNodeInfoResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *QuerySensNodeInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QuerySensNodeInfoResponseBody
	GetSuccess() *bool
}

type QuerySensNodeInfoResponseBody struct {
	// The returned business data in the JSON format.
	//
	// example:
	//
	// {"success": true, "httpStatusCode": 200, "data": { "result": [ { "sensitiveName": "certificate expiration date", "sensitiveId": "fd4ff5a2-9537-43d1-8e4f-1d0b5ffaac12", "status": 0, "templateName": "built-in classification and grading template", "keyRuleId": "228248921215042 certificate expiration date"} ], "totalCount": 1, "currentPage": 1, "pageSize": 10 }, "requestId": 10000001}
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

func (s QuerySensNodeInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QuerySensNodeInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySensNodeInfoResponseBody) GetData() interface{} {
	return s.Data
}

func (s *QuerySensNodeInfoResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *QuerySensNodeInfoResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *QuerySensNodeInfoResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *QuerySensNodeInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QuerySensNodeInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QuerySensNodeInfoResponseBody) SetData(v interface{}) *QuerySensNodeInfoResponseBody {
	s.Data = v
	return s
}

func (s *QuerySensNodeInfoResponseBody) SetErrorCode(v string) *QuerySensNodeInfoResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *QuerySensNodeInfoResponseBody) SetErrorMessage(v string) *QuerySensNodeInfoResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QuerySensNodeInfoResponseBody) SetHttpStatusCode(v int32) *QuerySensNodeInfoResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QuerySensNodeInfoResponseBody) SetRequestId(v string) *QuerySensNodeInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySensNodeInfoResponseBody) SetSuccess(v bool) *QuerySensNodeInfoResponseBody {
	s.Success = &v
	return s
}

func (s *QuerySensNodeInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
