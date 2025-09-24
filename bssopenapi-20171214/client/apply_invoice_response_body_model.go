// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyInvoiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ApplyInvoiceResponseBody
	GetCode() *string
	SetData(v *ApplyInvoiceResponseBodyData) *ApplyInvoiceResponseBody
	GetData() *ApplyInvoiceResponseBodyData
	SetMessage(v string) *ApplyInvoiceResponseBody
	GetMessage() *string
	SetRequestId(v string) *ApplyInvoiceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ApplyInvoiceResponseBody
	GetSuccess() *bool
}

type ApplyInvoiceResponseBody struct {
	// The status code.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *ApplyInvoiceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message returned.
	//
	// example:
	//
	// Successful!
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 6000EE23-274B-4E07-A697-FF2E999520A4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ApplyInvoiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ApplyInvoiceResponseBody) GoString() string {
	return s.String()
}

func (s *ApplyInvoiceResponseBody) GetCode() *string {
	return s.Code
}

func (s *ApplyInvoiceResponseBody) GetData() *ApplyInvoiceResponseBodyData {
	return s.Data
}

func (s *ApplyInvoiceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ApplyInvoiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ApplyInvoiceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ApplyInvoiceResponseBody) SetCode(v string) *ApplyInvoiceResponseBody {
	s.Code = &v
	return s
}

func (s *ApplyInvoiceResponseBody) SetData(v *ApplyInvoiceResponseBodyData) *ApplyInvoiceResponseBody {
	s.Data = v
	return s
}

func (s *ApplyInvoiceResponseBody) SetMessage(v string) *ApplyInvoiceResponseBody {
	s.Message = &v
	return s
}

func (s *ApplyInvoiceResponseBody) SetRequestId(v string) *ApplyInvoiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ApplyInvoiceResponseBody) SetSuccess(v bool) *ApplyInvoiceResponseBody {
	s.Success = &v
	return s
}

func (s *ApplyInvoiceResponseBody) Validate() error {
	return dara.Validate(s)
}

type ApplyInvoiceResponseBodyData struct {
	// The ID of the application.
	//
	// example:
	//
	// 1323125534
	InvoiceApplyId *int64 `json:"InvoiceApplyId,omitempty" xml:"InvoiceApplyId,omitempty"`
}

func (s ApplyInvoiceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ApplyInvoiceResponseBodyData) GoString() string {
	return s.String()
}

func (s *ApplyInvoiceResponseBodyData) GetInvoiceApplyId() *int64 {
	return s.InvoiceApplyId
}

func (s *ApplyInvoiceResponseBodyData) SetInvoiceApplyId(v int64) *ApplyInvoiceResponseBodyData {
	s.InvoiceApplyId = &v
	return s
}

func (s *ApplyInvoiceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
