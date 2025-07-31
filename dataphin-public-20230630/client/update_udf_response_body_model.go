// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUdfResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateUdfResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *UpdateUdfResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateUdfResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateUdfResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateUdfResponseBody
	GetSuccess() *bool
}

type UpdateUdfResponseBody struct {
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
	// internal error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateUdfResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateUdfResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateUdfResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateUdfResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateUdfResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateUdfResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateUdfResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateUdfResponseBody) SetCode(v string) *UpdateUdfResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateUdfResponseBody) SetHttpStatusCode(v int32) *UpdateUdfResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateUdfResponseBody) SetMessage(v string) *UpdateUdfResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateUdfResponseBody) SetRequestId(v string) *UpdateUdfResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateUdfResponseBody) SetSuccess(v bool) *UpdateUdfResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateUdfResponseBody) Validate() error {
	return dara.Validate(s)
}
