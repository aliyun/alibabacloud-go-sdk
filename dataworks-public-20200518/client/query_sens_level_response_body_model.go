// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySensLevelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v interface{}) *QuerySensLevelResponseBody
	GetData() interface{}
	SetErrorCode(v string) *QuerySensLevelResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *QuerySensLevelResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *QuerySensLevelResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *QuerySensLevelResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QuerySensLevelResponseBody
	GetSuccess() *bool
}

type QuerySensLevelResponseBody struct {
	// The returned data about data sensitivity levels. The data is in the JSON array format.
	//
	// example:
	//
	// [{"level":1,"isSensitive":false,"levelName":"1level"},{"level":2,"isSensitive":false,"levelName":"2level"}]
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

func (s QuerySensLevelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QuerySensLevelResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySensLevelResponseBody) GetData() interface{} {
	return s.Data
}

func (s *QuerySensLevelResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *QuerySensLevelResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *QuerySensLevelResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *QuerySensLevelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QuerySensLevelResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QuerySensLevelResponseBody) SetData(v interface{}) *QuerySensLevelResponseBody {
	s.Data = v
	return s
}

func (s *QuerySensLevelResponseBody) SetErrorCode(v string) *QuerySensLevelResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *QuerySensLevelResponseBody) SetErrorMessage(v string) *QuerySensLevelResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QuerySensLevelResponseBody) SetHttpStatusCode(v int32) *QuerySensLevelResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QuerySensLevelResponseBody) SetRequestId(v string) *QuerySensLevelResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySensLevelResponseBody) SetSuccess(v bool) *QuerySensLevelResponseBody {
	s.Success = &v
	return s
}

func (s *QuerySensLevelResponseBody) Validate() error {
	return dara.Validate(s)
}
