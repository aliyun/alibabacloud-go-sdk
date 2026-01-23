// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStandardLookupTableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateStandardLookupTableResponseBody
	GetCode() *string
	SetData(v int64) *CreateStandardLookupTableResponseBody
	GetData() *int64
	SetHttpStatusCode(v int32) *CreateStandardLookupTableResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateStandardLookupTableResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateStandardLookupTableResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateStandardLookupTableResponseBody
	GetSuccess() *bool
}

type CreateStandardLookupTableResponseBody struct {
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

func (s CreateStandardLookupTableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateStandardLookupTableResponseBody) GoString() string {
	return s.String()
}

func (s *CreateStandardLookupTableResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateStandardLookupTableResponseBody) GetData() *int64 {
	return s.Data
}

func (s *CreateStandardLookupTableResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateStandardLookupTableResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateStandardLookupTableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateStandardLookupTableResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateStandardLookupTableResponseBody) SetCode(v string) *CreateStandardLookupTableResponseBody {
	s.Code = &v
	return s
}

func (s *CreateStandardLookupTableResponseBody) SetData(v int64) *CreateStandardLookupTableResponseBody {
	s.Data = &v
	return s
}

func (s *CreateStandardLookupTableResponseBody) SetHttpStatusCode(v int32) *CreateStandardLookupTableResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateStandardLookupTableResponseBody) SetMessage(v string) *CreateStandardLookupTableResponseBody {
	s.Message = &v
	return s
}

func (s *CreateStandardLookupTableResponseBody) SetRequestId(v string) *CreateStandardLookupTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateStandardLookupTableResponseBody) SetSuccess(v bool) *CreateStandardLookupTableResponseBody {
	s.Success = &v
	return s
}

func (s *CreateStandardLookupTableResponseBody) Validate() error {
	return dara.Validate(s)
}
