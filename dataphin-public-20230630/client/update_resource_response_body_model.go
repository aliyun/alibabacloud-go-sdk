// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateResourceResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *UpdateResourceResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateResourceResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateResourceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateResourceResponseBody
	GetSuccess() *bool
}

type UpdateResourceResponseBody struct {
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

func (s UpdateResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateResourceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateResourceResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateResourceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateResourceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateResourceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateResourceResponseBody) SetCode(v string) *UpdateResourceResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateResourceResponseBody) SetHttpStatusCode(v int32) *UpdateResourceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateResourceResponseBody) SetMessage(v string) *UpdateResourceResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateResourceResponseBody) SetRequestId(v string) *UpdateResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateResourceResponseBody) SetSuccess(v bool) *UpdateResourceResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateResourceResponseBody) Validate() error {
	return dara.Validate(s)
}
