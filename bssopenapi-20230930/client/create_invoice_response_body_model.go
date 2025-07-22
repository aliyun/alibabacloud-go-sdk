// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInvoiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*CreateInvoiceResponseBodyData) *CreateInvoiceResponseBody
	GetData() []*CreateInvoiceResponseBodyData
	SetMetadata(v interface{}) *CreateInvoiceResponseBody
	GetMetadata() interface{}
	SetRequestId(v string) *CreateInvoiceResponseBody
	GetRequestId() *string
}

type CreateInvoiceResponseBody struct {
	Data []*CreateInvoiceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// {}
	Metadata interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// example:
	//
	// 79EE7556-0CFD-44EB-9CD6-B3B526E3A85F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateInvoiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateInvoiceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInvoiceResponseBody) GetData() []*CreateInvoiceResponseBodyData {
	return s.Data
}

func (s *CreateInvoiceResponseBody) GetMetadata() interface{} {
	return s.Metadata
}

func (s *CreateInvoiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateInvoiceResponseBody) SetData(v []*CreateInvoiceResponseBodyData) *CreateInvoiceResponseBody {
	s.Data = v
	return s
}

func (s *CreateInvoiceResponseBody) SetMetadata(v interface{}) *CreateInvoiceResponseBody {
	s.Metadata = v
	return s
}

func (s *CreateInvoiceResponseBody) SetRequestId(v string) *CreateInvoiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateInvoiceResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateInvoiceResponseBodyData struct {
	// example:
	//
	// 1990699401005016
	AccountId *int64 `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// example:
	//
	// 0.01
	Amount *string `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// example:
	//
	// 1001
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// ALIYUN_SERVICE
	InvoiceIssuer *string `json:"InvoiceIssuer,omitempty" xml:"InvoiceIssuer,omitempty"`
	Message       *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s CreateInvoiceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateInvoiceResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateInvoiceResponseBodyData) GetAccountId() *int64 {
	return s.AccountId
}

func (s *CreateInvoiceResponseBodyData) GetAmount() *string {
	return s.Amount
}

func (s *CreateInvoiceResponseBodyData) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateInvoiceResponseBodyData) GetInvoiceIssuer() *string {
	return s.InvoiceIssuer
}

func (s *CreateInvoiceResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *CreateInvoiceResponseBodyData) SetAccountId(v int64) *CreateInvoiceResponseBodyData {
	s.AccountId = &v
	return s
}

func (s *CreateInvoiceResponseBodyData) SetAmount(v string) *CreateInvoiceResponseBodyData {
	s.Amount = &v
	return s
}

func (s *CreateInvoiceResponseBodyData) SetErrorCode(v string) *CreateInvoiceResponseBodyData {
	s.ErrorCode = &v
	return s
}

func (s *CreateInvoiceResponseBodyData) SetInvoiceIssuer(v string) *CreateInvoiceResponseBodyData {
	s.InvoiceIssuer = &v
	return s
}

func (s *CreateInvoiceResponseBodyData) SetMessage(v string) *CreateInvoiceResponseBodyData {
	s.Message = &v
	return s
}

func (s *CreateInvoiceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
