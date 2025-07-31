// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateComputeSourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateComputeSourceResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *UpdateComputeSourceResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateComputeSourceResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateComputeSourceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateComputeSourceResponseBody
	GetSuccess() *bool
}

type UpdateComputeSourceResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateComputeSourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateComputeSourceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateComputeSourceResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateComputeSourceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateComputeSourceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateComputeSourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateComputeSourceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateComputeSourceResponseBody) SetCode(v string) *UpdateComputeSourceResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateComputeSourceResponseBody) SetHttpStatusCode(v int32) *UpdateComputeSourceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateComputeSourceResponseBody) SetMessage(v string) *UpdateComputeSourceResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateComputeSourceResponseBody) SetRequestId(v string) *UpdateComputeSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateComputeSourceResponseBody) SetSuccess(v bool) *UpdateComputeSourceResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateComputeSourceResponseBody) Validate() error {
	return dara.Validate(s)
}
