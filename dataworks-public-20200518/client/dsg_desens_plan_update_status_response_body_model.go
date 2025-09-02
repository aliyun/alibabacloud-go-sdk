// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDsgDesensPlanUpdateStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *DsgDesensPlanUpdateStatusResponseBody
	GetData() *bool
	SetErrorCode(v string) *DsgDesensPlanUpdateStatusResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DsgDesensPlanUpdateStatusResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *DsgDesensPlanUpdateStatusResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *DsgDesensPlanUpdateStatusResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DsgDesensPlanUpdateStatusResponseBody
	GetSuccess() *bool
}

type DsgDesensPlanUpdateStatusResponseBody struct {
	// The operation result. Valid values:
	//
	// 	- true: The operation is successful.
	//
	// 	- false: The operation fails.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error code.
	//
	// example:
	//
	// 1029030003
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// param error
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 400
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID. You can locate logs and troubleshoot issues based on the ID.
	//
	// example:
	//
	// 102400001
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

func (s DsgDesensPlanUpdateStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DsgDesensPlanUpdateStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DsgDesensPlanUpdateStatusResponseBody) GetData() *bool {
	return s.Data
}

func (s *DsgDesensPlanUpdateStatusResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DsgDesensPlanUpdateStatusResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DsgDesensPlanUpdateStatusResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DsgDesensPlanUpdateStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DsgDesensPlanUpdateStatusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DsgDesensPlanUpdateStatusResponseBody) SetData(v bool) *DsgDesensPlanUpdateStatusResponseBody {
	s.Data = &v
	return s
}

func (s *DsgDesensPlanUpdateStatusResponseBody) SetErrorCode(v string) *DsgDesensPlanUpdateStatusResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DsgDesensPlanUpdateStatusResponseBody) SetErrorMessage(v string) *DsgDesensPlanUpdateStatusResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DsgDesensPlanUpdateStatusResponseBody) SetHttpStatusCode(v int32) *DsgDesensPlanUpdateStatusResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DsgDesensPlanUpdateStatusResponseBody) SetRequestId(v string) *DsgDesensPlanUpdateStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DsgDesensPlanUpdateStatusResponseBody) SetSuccess(v bool) *DsgDesensPlanUpdateStatusResponseBody {
	s.Success = &v
	return s
}

func (s *DsgDesensPlanUpdateStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
