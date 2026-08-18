// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAppendCasesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AppendCasesResponseBody
	GetCode() *string
	SetData(v []*AppendCasesResponseBodyData) *AppendCasesResponseBody
	GetData() []*AppendCasesResponseBodyData
	SetHttpStatusCode(v int32) *AppendCasesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *AppendCasesResponseBody
	GetMessage() *string
	SetParams(v []*string) *AppendCasesResponseBody
	GetParams() []*string
	SetRequestId(v string) *AppendCasesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AppendCasesResponseBody
	GetSuccess() *bool
}

type AppendCasesResponseBody struct {
	// The response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The list of contacts that failed to be added.
	Data []*AppendCasesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
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

func (s AppendCasesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AppendCasesResponseBody) GoString() string {
	return s.String()
}

func (s *AppendCasesResponseBody) GetCode() *string {
	return s.Code
}

func (s *AppendCasesResponseBody) GetData() []*AppendCasesResponseBodyData {
	return s.Data
}

func (s *AppendCasesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *AppendCasesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AppendCasesResponseBody) GetParams() []*string {
	return s.Params
}

func (s *AppendCasesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AppendCasesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AppendCasesResponseBody) SetCode(v string) *AppendCasesResponseBody {
	s.Code = &v
	return s
}

func (s *AppendCasesResponseBody) SetData(v []*AppendCasesResponseBodyData) *AppendCasesResponseBody {
	s.Data = v
	return s
}

func (s *AppendCasesResponseBody) SetHttpStatusCode(v int32) *AppendCasesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AppendCasesResponseBody) SetMessage(v string) *AppendCasesResponseBody {
	s.Message = &v
	return s
}

func (s *AppendCasesResponseBody) SetParams(v []*string) *AppendCasesResponseBody {
	s.Params = v
	return s
}

func (s *AppendCasesResponseBody) SetRequestId(v string) *AppendCasesResponseBody {
	s.RequestId = &v
	return s
}

func (s *AppendCasesResponseBody) SetSuccess(v bool) *AppendCasesResponseBody {
	s.Success = &v
	return s
}

func (s *AppendCasesResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AppendCasesResponseBodyData struct {
	// The phone number.
	//
	// example:
	//
	// 133xxxxxxxx
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// The business ID.
	//
	// example:
	//
	// bizId-xxxxxx
	ReferenceId *string `json:"ReferenceId,omitempty" xml:"ReferenceId,omitempty"`
}

func (s AppendCasesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AppendCasesResponseBodyData) GoString() string {
	return s.String()
}

func (s *AppendCasesResponseBodyData) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *AppendCasesResponseBodyData) GetReferenceId() *string {
	return s.ReferenceId
}

func (s *AppendCasesResponseBodyData) SetPhoneNumber(v string) *AppendCasesResponseBodyData {
	s.PhoneNumber = &v
	return s
}

func (s *AppendCasesResponseBodyData) SetReferenceId(v string) *AppendCasesResponseBodyData {
	s.ReferenceId = &v
	return s
}

func (s *AppendCasesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
