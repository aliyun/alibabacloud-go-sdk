// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFindFinanceTaxDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *FindFinanceTaxDetailResponseBody
	GetCode() *string
	SetFinance(v *FindFinanceTaxDetailResponseBodyFinance) *FindFinanceTaxDetailResponseBody
	GetFinance() *FindFinanceTaxDetailResponseBodyFinance
	SetMessage(v string) *FindFinanceTaxDetailResponseBody
	GetMessage() *string
	SetRequestId(v string) *FindFinanceTaxDetailResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *FindFinanceTaxDetailResponseBody
	GetSuccess() *bool
}

type FindFinanceTaxDetailResponseBody struct {
	Code      *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Finance   *FindFinanceTaxDetailResponseBodyFinance `json:"Finance,omitempty" xml:"Finance,omitempty" type:"Struct"`
	Message   *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s FindFinanceTaxDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FindFinanceTaxDetailResponseBody) GoString() string {
	return s.String()
}

func (s *FindFinanceTaxDetailResponseBody) GetCode() *string {
	return s.Code
}

func (s *FindFinanceTaxDetailResponseBody) GetFinance() *FindFinanceTaxDetailResponseBodyFinance {
	return s.Finance
}

func (s *FindFinanceTaxDetailResponseBody) GetMessage() *string {
	return s.Message
}

func (s *FindFinanceTaxDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FindFinanceTaxDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *FindFinanceTaxDetailResponseBody) SetCode(v string) *FindFinanceTaxDetailResponseBody {
	s.Code = &v
	return s
}

func (s *FindFinanceTaxDetailResponseBody) SetFinance(v *FindFinanceTaxDetailResponseBodyFinance) *FindFinanceTaxDetailResponseBody {
	s.Finance = v
	return s
}

func (s *FindFinanceTaxDetailResponseBody) SetMessage(v string) *FindFinanceTaxDetailResponseBody {
	s.Message = &v
	return s
}

func (s *FindFinanceTaxDetailResponseBody) SetRequestId(v string) *FindFinanceTaxDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *FindFinanceTaxDetailResponseBody) SetSuccess(v bool) *FindFinanceTaxDetailResponseBody {
	s.Success = &v
	return s
}

func (s *FindFinanceTaxDetailResponseBody) Validate() error {
	if s.Finance != nil {
		if err := s.Finance.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type FindFinanceTaxDetailResponseBodyFinance struct {
	FinanceTaxCertificateImgName       *string `json:"FinanceTaxCertificateImgName,omitempty" xml:"FinanceTaxCertificateImgName,omitempty"`
	Tax                                *string `json:"Tax,omitempty" xml:"Tax,omitempty"`
	FinanceTaxCertificateImgUrl        *string `json:"financeTaxCertificateImgUrl,omitempty" xml:"financeTaxCertificateImgUrl,omitempty"`
	SecondFinanceTax                   *string `json:"secondFinanceTax,omitempty" xml:"secondFinanceTax,omitempty"`
	SecondFinanceTaxCertificateImgName *string `json:"secondFinanceTaxCertificateImgName,omitempty" xml:"secondFinanceTaxCertificateImgName,omitempty"`
	SecondFinanceTaxCertificateImgUrl  *string `json:"secondFinanceTaxCertificateImgUrl,omitempty" xml:"secondFinanceTaxCertificateImgUrl,omitempty"`
}

func (s FindFinanceTaxDetailResponseBodyFinance) String() string {
	return dara.Prettify(s)
}

func (s FindFinanceTaxDetailResponseBodyFinance) GoString() string {
	return s.String()
}

func (s *FindFinanceTaxDetailResponseBodyFinance) GetFinanceTaxCertificateImgName() *string {
	return s.FinanceTaxCertificateImgName
}

func (s *FindFinanceTaxDetailResponseBodyFinance) GetTax() *string {
	return s.Tax
}

func (s *FindFinanceTaxDetailResponseBodyFinance) GetFinanceTaxCertificateImgUrl() *string {
	return s.FinanceTaxCertificateImgUrl
}

func (s *FindFinanceTaxDetailResponseBodyFinance) GetSecondFinanceTax() *string {
	return s.SecondFinanceTax
}

func (s *FindFinanceTaxDetailResponseBodyFinance) GetSecondFinanceTaxCertificateImgName() *string {
	return s.SecondFinanceTaxCertificateImgName
}

func (s *FindFinanceTaxDetailResponseBodyFinance) GetSecondFinanceTaxCertificateImgUrl() *string {
	return s.SecondFinanceTaxCertificateImgUrl
}

func (s *FindFinanceTaxDetailResponseBodyFinance) SetFinanceTaxCertificateImgName(v string) *FindFinanceTaxDetailResponseBodyFinance {
	s.FinanceTaxCertificateImgName = &v
	return s
}

func (s *FindFinanceTaxDetailResponseBodyFinance) SetTax(v string) *FindFinanceTaxDetailResponseBodyFinance {
	s.Tax = &v
	return s
}

func (s *FindFinanceTaxDetailResponseBodyFinance) SetFinanceTaxCertificateImgUrl(v string) *FindFinanceTaxDetailResponseBodyFinance {
	s.FinanceTaxCertificateImgUrl = &v
	return s
}

func (s *FindFinanceTaxDetailResponseBodyFinance) SetSecondFinanceTax(v string) *FindFinanceTaxDetailResponseBodyFinance {
	s.SecondFinanceTax = &v
	return s
}

func (s *FindFinanceTaxDetailResponseBodyFinance) SetSecondFinanceTaxCertificateImgName(v string) *FindFinanceTaxDetailResponseBodyFinance {
	s.SecondFinanceTaxCertificateImgName = &v
	return s
}

func (s *FindFinanceTaxDetailResponseBodyFinance) SetSecondFinanceTaxCertificateImgUrl(v string) *FindFinanceTaxDetailResponseBodyFinance {
	s.SecondFinanceTaxCertificateImgUrl = &v
	return s
}

func (s *FindFinanceTaxDetailResponseBodyFinance) Validate() error {
	return dara.Validate(s)
}
