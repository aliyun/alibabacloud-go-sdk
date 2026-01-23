// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateStandardWordRootResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateStandardWordRootResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *UpdateStandardWordRootResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateStandardWordRootResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateStandardWordRootResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateStandardWordRootResponseBody
	GetSuccess() *bool
}

type UpdateStandardWordRootResponseBody struct {
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

func (s UpdateStandardWordRootResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateStandardWordRootResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateStandardWordRootResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateStandardWordRootResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateStandardWordRootResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateStandardWordRootResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateStandardWordRootResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateStandardWordRootResponseBody) SetCode(v string) *UpdateStandardWordRootResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateStandardWordRootResponseBody) SetHttpStatusCode(v int32) *UpdateStandardWordRootResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateStandardWordRootResponseBody) SetMessage(v string) *UpdateStandardWordRootResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateStandardWordRootResponseBody) SetRequestId(v string) *UpdateStandardWordRootResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateStandardWordRootResponseBody) SetSuccess(v bool) *UpdateStandardWordRootResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateStandardWordRootResponseBody) Validate() error {
	return dara.Validate(s)
}
