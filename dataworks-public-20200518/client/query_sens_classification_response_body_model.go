// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySensClassificationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v interface{}) *QuerySensClassificationResponseBody
	GetData() interface{}
	SetErrorCode(v string) *QuerySensClassificationResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *QuerySensClassificationResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *QuerySensClassificationResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *QuerySensClassificationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QuerySensClassificationResponseBody
	GetSuccess() *bool
}

type QuerySensClassificationResponseBody struct {
	// The returned data about data categories. The data is in the JSON format.
	//
	// example:
	//
	// [{         "nodeName": "teset1",         "sensitiveNotNull": true,         "nodeId": "1aac2e35-b437-486b-95c7-a5ae48371024",         "nodeOldId": "1aac2e35-b437-486b-95c7-a5ae48371024",         "parentId": "0"     }]
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
	// 400
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

func (s QuerySensClassificationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QuerySensClassificationResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySensClassificationResponseBody) GetData() interface{} {
	return s.Data
}

func (s *QuerySensClassificationResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *QuerySensClassificationResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *QuerySensClassificationResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *QuerySensClassificationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QuerySensClassificationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QuerySensClassificationResponseBody) SetData(v interface{}) *QuerySensClassificationResponseBody {
	s.Data = v
	return s
}

func (s *QuerySensClassificationResponseBody) SetErrorCode(v string) *QuerySensClassificationResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *QuerySensClassificationResponseBody) SetErrorMessage(v string) *QuerySensClassificationResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QuerySensClassificationResponseBody) SetHttpStatusCode(v int32) *QuerySensClassificationResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QuerySensClassificationResponseBody) SetRequestId(v string) *QuerySensClassificationResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySensClassificationResponseBody) SetSuccess(v bool) *QuerySensClassificationResponseBody {
	s.Success = &v
	return s
}

func (s *QuerySensClassificationResponseBody) Validate() error {
	return dara.Validate(s)
}
