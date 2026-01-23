// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteStandardLookupTableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteStandardLookupTableResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DeleteStandardLookupTableResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteStandardLookupTableResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteStandardLookupTableResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteStandardLookupTableResponseBody
	GetSuccess() *bool
}

type DeleteStandardLookupTableResponseBody struct {
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

func (s DeleteStandardLookupTableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteStandardLookupTableResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteStandardLookupTableResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteStandardLookupTableResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteStandardLookupTableResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteStandardLookupTableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteStandardLookupTableResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteStandardLookupTableResponseBody) SetCode(v string) *DeleteStandardLookupTableResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteStandardLookupTableResponseBody) SetHttpStatusCode(v int32) *DeleteStandardLookupTableResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteStandardLookupTableResponseBody) SetMessage(v string) *DeleteStandardLookupTableResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteStandardLookupTableResponseBody) SetRequestId(v string) *DeleteStandardLookupTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteStandardLookupTableResponseBody) SetSuccess(v bool) *DeleteStandardLookupTableResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteStandardLookupTableResponseBody) Validate() error {
	return dara.Validate(s)
}
