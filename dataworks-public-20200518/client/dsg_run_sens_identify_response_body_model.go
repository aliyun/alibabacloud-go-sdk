// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDsgRunSensIdentifyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v interface{}) *DsgRunSensIdentifyResponseBody
	GetData() interface{}
	SetErrorCode(v string) *DsgRunSensIdentifyResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DsgRunSensIdentifyResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *DsgRunSensIdentifyResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *DsgRunSensIdentifyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DsgRunSensIdentifyResponseBody
	GetSuccess() *bool
}

type DsgRunSensIdentifyResponseBody struct {
	// The ID of the generated sensitive data identification task. The value is of the INT 64 type.
	//
	// example:
	//
	// 1000001
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

func (s DsgRunSensIdentifyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DsgRunSensIdentifyResponseBody) GoString() string {
	return s.String()
}

func (s *DsgRunSensIdentifyResponseBody) GetData() interface{} {
	return s.Data
}

func (s *DsgRunSensIdentifyResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DsgRunSensIdentifyResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DsgRunSensIdentifyResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DsgRunSensIdentifyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DsgRunSensIdentifyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DsgRunSensIdentifyResponseBody) SetData(v interface{}) *DsgRunSensIdentifyResponseBody {
	s.Data = v
	return s
}

func (s *DsgRunSensIdentifyResponseBody) SetErrorCode(v string) *DsgRunSensIdentifyResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DsgRunSensIdentifyResponseBody) SetErrorMessage(v string) *DsgRunSensIdentifyResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DsgRunSensIdentifyResponseBody) SetHttpStatusCode(v int32) *DsgRunSensIdentifyResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DsgRunSensIdentifyResponseBody) SetRequestId(v string) *DsgRunSensIdentifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DsgRunSensIdentifyResponseBody) SetSuccess(v bool) *DsgRunSensIdentifyResponseBody {
	s.Success = &v
	return s
}

func (s *DsgRunSensIdentifyResponseBody) Validate() error {
	return dara.Validate(s)
}
