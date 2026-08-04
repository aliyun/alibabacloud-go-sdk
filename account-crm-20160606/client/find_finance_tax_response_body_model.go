// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFindFinanceTaxResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *FindFinanceTaxResponseBody
	GetCode() *string
	SetFinanceVersion(v *FindFinanceTaxResponseBodyFinanceVersion) *FindFinanceTaxResponseBody
	GetFinanceVersion() *FindFinanceTaxResponseBodyFinanceVersion
	SetMessage(v string) *FindFinanceTaxResponseBody
	GetMessage() *string
	SetRequestId(v string) *FindFinanceTaxResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *FindFinanceTaxResponseBody
	GetSuccess() *bool
}

type FindFinanceTaxResponseBody struct {
	Code           *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	FinanceVersion *FindFinanceTaxResponseBodyFinanceVersion `json:"FinanceVersion,omitempty" xml:"FinanceVersion,omitempty" type:"Struct"`
	Message        *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s FindFinanceTaxResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FindFinanceTaxResponseBody) GoString() string {
	return s.String()
}

func (s *FindFinanceTaxResponseBody) GetCode() *string {
	return s.Code
}

func (s *FindFinanceTaxResponseBody) GetFinanceVersion() *FindFinanceTaxResponseBodyFinanceVersion {
	return s.FinanceVersion
}

func (s *FindFinanceTaxResponseBody) GetMessage() *string {
	return s.Message
}

func (s *FindFinanceTaxResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FindFinanceTaxResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *FindFinanceTaxResponseBody) SetCode(v string) *FindFinanceTaxResponseBody {
	s.Code = &v
	return s
}

func (s *FindFinanceTaxResponseBody) SetFinanceVersion(v *FindFinanceTaxResponseBodyFinanceVersion) *FindFinanceTaxResponseBody {
	s.FinanceVersion = v
	return s
}

func (s *FindFinanceTaxResponseBody) SetMessage(v string) *FindFinanceTaxResponseBody {
	s.Message = &v
	return s
}

func (s *FindFinanceTaxResponseBody) SetRequestId(v string) *FindFinanceTaxResponseBody {
	s.RequestId = &v
	return s
}

func (s *FindFinanceTaxResponseBody) SetSuccess(v bool) *FindFinanceTaxResponseBody {
	s.Success = &v
	return s
}

func (s *FindFinanceTaxResponseBody) Validate() error {
	if s.FinanceVersion != nil {
		if err := s.FinanceVersion.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type FindFinanceTaxResponseBodyFinanceVersion struct {
	FinanceTaxCertificateImgName       *string `json:"FinanceTaxCertificateImgName,omitempty" xml:"FinanceTaxCertificateImgName,omitempty"`
	FinanceTaxCertificateImgUrl        *string `json:"FinanceTaxCertificateImgUrl,omitempty" xml:"FinanceTaxCertificateImgUrl,omitempty"`
	SecondFinanceTax                   *string `json:"SecondFinanceTax,omitempty" xml:"SecondFinanceTax,omitempty"`
	SecondFinanceTaxCertificateImgName *string `json:"SecondFinanceTaxCertificateImgName,omitempty" xml:"SecondFinanceTaxCertificateImgName,omitempty"`
	SecondFinanceTaxCertificateImgUrl  *string `json:"SecondFinanceTaxCertificateImgUrl,omitempty" xml:"SecondFinanceTaxCertificateImgUrl,omitempty"`
	Tax                                *string `json:"Tax,omitempty" xml:"Tax,omitempty"`
	Version                            *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s FindFinanceTaxResponseBodyFinanceVersion) String() string {
	return dara.Prettify(s)
}

func (s FindFinanceTaxResponseBodyFinanceVersion) GoString() string {
	return s.String()
}

func (s *FindFinanceTaxResponseBodyFinanceVersion) GetFinanceTaxCertificateImgName() *string {
	return s.FinanceTaxCertificateImgName
}

func (s *FindFinanceTaxResponseBodyFinanceVersion) GetFinanceTaxCertificateImgUrl() *string {
	return s.FinanceTaxCertificateImgUrl
}

func (s *FindFinanceTaxResponseBodyFinanceVersion) GetSecondFinanceTax() *string {
	return s.SecondFinanceTax
}

func (s *FindFinanceTaxResponseBodyFinanceVersion) GetSecondFinanceTaxCertificateImgName() *string {
	return s.SecondFinanceTaxCertificateImgName
}

func (s *FindFinanceTaxResponseBodyFinanceVersion) GetSecondFinanceTaxCertificateImgUrl() *string {
	return s.SecondFinanceTaxCertificateImgUrl
}

func (s *FindFinanceTaxResponseBodyFinanceVersion) GetTax() *string {
	return s.Tax
}

func (s *FindFinanceTaxResponseBodyFinanceVersion) GetVersion() *string {
	return s.Version
}

func (s *FindFinanceTaxResponseBodyFinanceVersion) SetFinanceTaxCertificateImgName(v string) *FindFinanceTaxResponseBodyFinanceVersion {
	s.FinanceTaxCertificateImgName = &v
	return s
}

func (s *FindFinanceTaxResponseBodyFinanceVersion) SetFinanceTaxCertificateImgUrl(v string) *FindFinanceTaxResponseBodyFinanceVersion {
	s.FinanceTaxCertificateImgUrl = &v
	return s
}

func (s *FindFinanceTaxResponseBodyFinanceVersion) SetSecondFinanceTax(v string) *FindFinanceTaxResponseBodyFinanceVersion {
	s.SecondFinanceTax = &v
	return s
}

func (s *FindFinanceTaxResponseBodyFinanceVersion) SetSecondFinanceTaxCertificateImgName(v string) *FindFinanceTaxResponseBodyFinanceVersion {
	s.SecondFinanceTaxCertificateImgName = &v
	return s
}

func (s *FindFinanceTaxResponseBodyFinanceVersion) SetSecondFinanceTaxCertificateImgUrl(v string) *FindFinanceTaxResponseBodyFinanceVersion {
	s.SecondFinanceTaxCertificateImgUrl = &v
	return s
}

func (s *FindFinanceTaxResponseBodyFinanceVersion) SetTax(v string) *FindFinanceTaxResponseBodyFinanceVersion {
	s.Tax = &v
	return s
}

func (s *FindFinanceTaxResponseBodyFinanceVersion) SetVersion(v string) *FindFinanceTaxResponseBodyFinanceVersion {
	s.Version = &v
	return s
}

func (s *FindFinanceTaxResponseBodyFinanceVersion) Validate() error {
	return dara.Validate(s)
}
