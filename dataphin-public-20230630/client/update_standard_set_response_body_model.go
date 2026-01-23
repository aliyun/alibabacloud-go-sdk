// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateStandardSetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateStandardSetResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *UpdateStandardSetResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateStandardSetResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateStandardSetResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateStandardSetResponseBody
	GetSuccess() *bool
}

type UpdateStandardSetResponseBody struct {
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

func (s UpdateStandardSetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateStandardSetResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateStandardSetResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateStandardSetResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateStandardSetResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateStandardSetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateStandardSetResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateStandardSetResponseBody) SetCode(v string) *UpdateStandardSetResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateStandardSetResponseBody) SetHttpStatusCode(v int32) *UpdateStandardSetResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateStandardSetResponseBody) SetMessage(v string) *UpdateStandardSetResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateStandardSetResponseBody) SetRequestId(v string) *UpdateStandardSetResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateStandardSetResponseBody) SetSuccess(v bool) *UpdateStandardSetResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateStandardSetResponseBody) Validate() error {
	return dara.Validate(s)
}
