// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStandardResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateStandardResponseBody
	GetCode() *string
	SetData(v int64) *CreateStandardResponseBody
	GetData() *int64
	SetHttpStatusCode(v int32) *CreateStandardResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateStandardResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateStandardResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateStandardResponseBody
	GetSuccess() *bool
}

type CreateStandardResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 1234
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

func (s CreateStandardResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateStandardResponseBody) GoString() string {
	return s.String()
}

func (s *CreateStandardResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateStandardResponseBody) GetData() *int64 {
	return s.Data
}

func (s *CreateStandardResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateStandardResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateStandardResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateStandardResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateStandardResponseBody) SetCode(v string) *CreateStandardResponseBody {
	s.Code = &v
	return s
}

func (s *CreateStandardResponseBody) SetData(v int64) *CreateStandardResponseBody {
	s.Data = &v
	return s
}

func (s *CreateStandardResponseBody) SetHttpStatusCode(v int32) *CreateStandardResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateStandardResponseBody) SetMessage(v string) *CreateStandardResponseBody {
	s.Message = &v
	return s
}

func (s *CreateStandardResponseBody) SetRequestId(v string) *CreateStandardResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateStandardResponseBody) SetSuccess(v bool) *CreateStandardResponseBody {
	s.Success = &v
	return s
}

func (s *CreateStandardResponseBody) Validate() error {
	return dara.Validate(s)
}
