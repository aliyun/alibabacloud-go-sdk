// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTxtRecordForVerifyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *GetTxtRecordForVerifyResponseBody
	GetDomainName() *string
	SetParentDomainName(v string) *GetTxtRecordForVerifyResponseBody
	GetParentDomainName() *string
	SetRR(v string) *GetTxtRecordForVerifyResponseBody
	GetRR() *string
	SetRequestId(v string) *GetTxtRecordForVerifyResponseBody
	GetRequestId() *string
	SetValue(v string) *GetTxtRecordForVerifyResponseBody
	GetValue() *string
}

type GetTxtRecordForVerifyResponseBody struct {
	// The domain name.
	//
	// >  If you do not specify this parameter, it is not returned.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The top-level domain name.
	//
	// example:
	//
	// com
	ParentDomainName *string `json:"ParentDomainName,omitempty" xml:"ParentDomainName,omitempty"`
	// The hostname.
	//
	// example:
	//
	// aliyunRetrieval
	RR *string `json:"RR,omitempty" xml:"RR,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9CC0D642-49D4-48DE-A1A5-9F218652E4A7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The record value.
	//
	// >  The validity period is three days.
	//
	// example:
	//
	// c99419e6997f41daaa3e*****
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetTxtRecordForVerifyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTxtRecordForVerifyResponseBody) GoString() string {
	return s.String()
}

func (s *GetTxtRecordForVerifyResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *GetTxtRecordForVerifyResponseBody) GetParentDomainName() *string {
	return s.ParentDomainName
}

func (s *GetTxtRecordForVerifyResponseBody) GetRR() *string {
	return s.RR
}

func (s *GetTxtRecordForVerifyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTxtRecordForVerifyResponseBody) GetValue() *string {
	return s.Value
}

func (s *GetTxtRecordForVerifyResponseBody) SetDomainName(v string) *GetTxtRecordForVerifyResponseBody {
	s.DomainName = &v
	return s
}

func (s *GetTxtRecordForVerifyResponseBody) SetParentDomainName(v string) *GetTxtRecordForVerifyResponseBody {
	s.ParentDomainName = &v
	return s
}

func (s *GetTxtRecordForVerifyResponseBody) SetRR(v string) *GetTxtRecordForVerifyResponseBody {
	s.RR = &v
	return s
}

func (s *GetTxtRecordForVerifyResponseBody) SetRequestId(v string) *GetTxtRecordForVerifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTxtRecordForVerifyResponseBody) SetValue(v string) *GetTxtRecordForVerifyResponseBody {
	s.Value = &v
	return s
}

func (s *GetTxtRecordForVerifyResponseBody) Validate() error {
	return dara.Validate(s)
}
