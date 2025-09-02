// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDsgStopSensIdentifyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v interface{}) *DsgStopSensIdentifyResponseBody
	GetData() interface{}
	SetErrorCode(v string) *DsgStopSensIdentifyResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DsgStopSensIdentifyResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *DsgStopSensIdentifyResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *DsgStopSensIdentifyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DsgStopSensIdentifyResponseBody
	GetSuccess() *bool
}

type DsgStopSensIdentifyResponseBody struct {
	// The returned data, which is of the Boolean type.
	//
	// example:
	//
	// true
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

func (s DsgStopSensIdentifyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DsgStopSensIdentifyResponseBody) GoString() string {
	return s.String()
}

func (s *DsgStopSensIdentifyResponseBody) GetData() interface{} {
	return s.Data
}

func (s *DsgStopSensIdentifyResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DsgStopSensIdentifyResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DsgStopSensIdentifyResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DsgStopSensIdentifyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DsgStopSensIdentifyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DsgStopSensIdentifyResponseBody) SetData(v interface{}) *DsgStopSensIdentifyResponseBody {
	s.Data = v
	return s
}

func (s *DsgStopSensIdentifyResponseBody) SetErrorCode(v string) *DsgStopSensIdentifyResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DsgStopSensIdentifyResponseBody) SetErrorMessage(v string) *DsgStopSensIdentifyResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DsgStopSensIdentifyResponseBody) SetHttpStatusCode(v int32) *DsgStopSensIdentifyResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DsgStopSensIdentifyResponseBody) SetRequestId(v string) *DsgStopSensIdentifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DsgStopSensIdentifyResponseBody) SetSuccess(v bool) *DsgStopSensIdentifyResponseBody {
	s.Success = &v
	return s
}

func (s *DsgStopSensIdentifyResponseBody) Validate() error {
	return dara.Validate(s)
}
