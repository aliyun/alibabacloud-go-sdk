// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStandardSetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateStandardSetResponseBody
	GetCode() *string
	SetData(v int64) *CreateStandardSetResponseBody
	GetData() *int64
	SetHttpStatusCode(v int32) *CreateStandardSetResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateStandardSetResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateStandardSetResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateStandardSetResponseBody
	GetSuccess() *bool
}

type CreateStandardSetResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 1201
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

func (s CreateStandardSetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateStandardSetResponseBody) GoString() string {
	return s.String()
}

func (s *CreateStandardSetResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateStandardSetResponseBody) GetData() *int64 {
	return s.Data
}

func (s *CreateStandardSetResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateStandardSetResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateStandardSetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateStandardSetResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateStandardSetResponseBody) SetCode(v string) *CreateStandardSetResponseBody {
	s.Code = &v
	return s
}

func (s *CreateStandardSetResponseBody) SetData(v int64) *CreateStandardSetResponseBody {
	s.Data = &v
	return s
}

func (s *CreateStandardSetResponseBody) SetHttpStatusCode(v int32) *CreateStandardSetResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateStandardSetResponseBody) SetMessage(v string) *CreateStandardSetResponseBody {
	s.Message = &v
	return s
}

func (s *CreateStandardSetResponseBody) SetRequestId(v string) *CreateStandardSetResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateStandardSetResponseBody) SetSuccess(v bool) *CreateStandardSetResponseBody {
	s.Success = &v
	return s
}

func (s *CreateStandardSetResponseBody) Validate() error {
	return dara.Validate(s)
}
