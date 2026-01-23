// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateStandardResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateStandardResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *UpdateStandardResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateStandardResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateStandardResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateStandardResponseBody
	GetSuccess() *bool
}

type UpdateStandardResponseBody struct {
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

func (s UpdateStandardResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateStandardResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateStandardResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateStandardResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateStandardResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateStandardResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateStandardResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateStandardResponseBody) SetCode(v string) *UpdateStandardResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateStandardResponseBody) SetHttpStatusCode(v int32) *UpdateStandardResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateStandardResponseBody) SetMessage(v string) *UpdateStandardResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateStandardResponseBody) SetRequestId(v string) *UpdateStandardResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateStandardResponseBody) SetSuccess(v bool) *UpdateStandardResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateStandardResponseBody) Validate() error {
	return dara.Validate(s)
}
