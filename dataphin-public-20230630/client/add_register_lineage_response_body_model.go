// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddRegisterLineageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AddRegisterLineageResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *AddRegisterLineageResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *AddRegisterLineageResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddRegisterLineageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddRegisterLineageResponseBody
	GetSuccess() *bool
}

type AddRegisterLineageResponseBody struct {
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

func (s AddRegisterLineageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddRegisterLineageResponseBody) GoString() string {
	return s.String()
}

func (s *AddRegisterLineageResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddRegisterLineageResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *AddRegisterLineageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddRegisterLineageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddRegisterLineageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddRegisterLineageResponseBody) SetCode(v string) *AddRegisterLineageResponseBody {
	s.Code = &v
	return s
}

func (s *AddRegisterLineageResponseBody) SetHttpStatusCode(v int32) *AddRegisterLineageResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AddRegisterLineageResponseBody) SetMessage(v string) *AddRegisterLineageResponseBody {
	s.Message = &v
	return s
}

func (s *AddRegisterLineageResponseBody) SetRequestId(v string) *AddRegisterLineageResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddRegisterLineageResponseBody) SetSuccess(v bool) *AddRegisterLineageResponseBody {
	s.Success = &v
	return s
}

func (s *AddRegisterLineageResponseBody) Validate() error {
	return dara.Validate(s)
}
