// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUdfResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateUdfResponseBody
	GetCode() *string
	SetData(v int64) *CreateUdfResponseBody
	GetData() *int64
	SetHttpStatusCode(v int32) *CreateUdfResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateUdfResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateUdfResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateUdfResponseBody
	GetSuccess() *bool
}

type CreateUdfResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 10212101
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
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

func (s CreateUdfResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateUdfResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUdfResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateUdfResponseBody) GetData() *int64 {
	return s.Data
}

func (s *CreateUdfResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateUdfResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateUdfResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateUdfResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateUdfResponseBody) SetCode(v string) *CreateUdfResponseBody {
	s.Code = &v
	return s
}

func (s *CreateUdfResponseBody) SetData(v int64) *CreateUdfResponseBody {
	s.Data = &v
	return s
}

func (s *CreateUdfResponseBody) SetHttpStatusCode(v int32) *CreateUdfResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateUdfResponseBody) SetMessage(v string) *CreateUdfResponseBody {
	s.Message = &v
	return s
}

func (s *CreateUdfResponseBody) SetRequestId(v string) *CreateUdfResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateUdfResponseBody) SetSuccess(v bool) *CreateUdfResponseBody {
	s.Success = &v
	return s
}

func (s *CreateUdfResponseBody) Validate() error {
	return dara.Validate(s)
}
