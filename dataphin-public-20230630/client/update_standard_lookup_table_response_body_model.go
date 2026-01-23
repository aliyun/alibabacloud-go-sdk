// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateStandardLookupTableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateStandardLookupTableResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *UpdateStandardLookupTableResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateStandardLookupTableResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateStandardLookupTableResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateStandardLookupTableResponseBody
	GetSuccess() *bool
}

type UpdateStandardLookupTableResponseBody struct {
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

func (s UpdateStandardLookupTableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateStandardLookupTableResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateStandardLookupTableResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateStandardLookupTableResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateStandardLookupTableResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateStandardLookupTableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateStandardLookupTableResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateStandardLookupTableResponseBody) SetCode(v string) *UpdateStandardLookupTableResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateStandardLookupTableResponseBody) SetHttpStatusCode(v int32) *UpdateStandardLookupTableResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateStandardLookupTableResponseBody) SetMessage(v string) *UpdateStandardLookupTableResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateStandardLookupTableResponseBody) SetRequestId(v string) *UpdateStandardLookupTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateStandardLookupTableResponseBody) SetSuccess(v bool) *UpdateStandardLookupTableResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateStandardLookupTableResponseBody) Validate() error {
	return dara.Validate(s)
}
