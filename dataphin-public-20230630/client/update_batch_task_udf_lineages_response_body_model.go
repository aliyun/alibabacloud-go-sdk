// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBatchTaskUdfLineagesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateBatchTaskUdfLineagesResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *UpdateBatchTaskUdfLineagesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateBatchTaskUdfLineagesResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateBatchTaskUdfLineagesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateBatchTaskUdfLineagesResponseBody
	GetSuccess() *bool
}

type UpdateBatchTaskUdfLineagesResponseBody struct {
	// The error code. A value of OK indicates that the request was successful.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code returned by the backend.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateBatchTaskUdfLineagesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateBatchTaskUdfLineagesResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateBatchTaskUdfLineagesResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateBatchTaskUdfLineagesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateBatchTaskUdfLineagesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateBatchTaskUdfLineagesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateBatchTaskUdfLineagesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateBatchTaskUdfLineagesResponseBody) SetCode(v string) *UpdateBatchTaskUdfLineagesResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateBatchTaskUdfLineagesResponseBody) SetHttpStatusCode(v int32) *UpdateBatchTaskUdfLineagesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateBatchTaskUdfLineagesResponseBody) SetMessage(v string) *UpdateBatchTaskUdfLineagesResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateBatchTaskUdfLineagesResponseBody) SetRequestId(v string) *UpdateBatchTaskUdfLineagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateBatchTaskUdfLineagesResponseBody) SetSuccess(v bool) *UpdateBatchTaskUdfLineagesResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateBatchTaskUdfLineagesResponseBody) Validate() error {
	return dara.Validate(s)
}
