// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAbortCasesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AbortCasesResponseBody
	GetCode() *string
	SetData(v bool) *AbortCasesResponseBody
	GetData() *bool
	SetHttpStatusCode(v int32) *AbortCasesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *AbortCasesResponseBody
	GetMessage() *string
	SetParams(v []*string) *AbortCasesResponseBody
	GetParams() []*string
	SetRequestId(v string) *AbortCasesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AbortCasesResponseBody
	GetSuccess() *bool
}

type AbortCasesResponseBody struct {
	// The return code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	//
	// example:
	//
	// Flash message configuration ID
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The message returned by the operation.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The variable values in the error message.
	Params []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AbortCasesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AbortCasesResponseBody) GoString() string {
	return s.String()
}

func (s *AbortCasesResponseBody) GetCode() *string {
	return s.Code
}

func (s *AbortCasesResponseBody) GetData() *bool {
	return s.Data
}

func (s *AbortCasesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *AbortCasesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AbortCasesResponseBody) GetParams() []*string {
	return s.Params
}

func (s *AbortCasesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AbortCasesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AbortCasesResponseBody) SetCode(v string) *AbortCasesResponseBody {
	s.Code = &v
	return s
}

func (s *AbortCasesResponseBody) SetData(v bool) *AbortCasesResponseBody {
	s.Data = &v
	return s
}

func (s *AbortCasesResponseBody) SetHttpStatusCode(v int32) *AbortCasesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AbortCasesResponseBody) SetMessage(v string) *AbortCasesResponseBody {
	s.Message = &v
	return s
}

func (s *AbortCasesResponseBody) SetParams(v []*string) *AbortCasesResponseBody {
	s.Params = v
	return s
}

func (s *AbortCasesResponseBody) SetRequestId(v string) *AbortCasesResponseBody {
	s.RequestId = &v
	return s
}

func (s *AbortCasesResponseBody) SetSuccess(v bool) *AbortCasesResponseBody {
	s.Success = &v
	return s
}

func (s *AbortCasesResponseBody) Validate() error {
	return dara.Validate(s)
}
