// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateFinanceTaxRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFinanceTax(v string) *OperateFinanceTaxRequest
	GetFinanceTax() *string
	SetFinanceTaxCertificateImgName(v string) *OperateFinanceTaxRequest
	GetFinanceTaxCertificateImgName() *string
	SetHId(v int64) *OperateFinanceTaxRequest
	GetHId() *int64
	SetSecondFinanceTax(v string) *OperateFinanceTaxRequest
	GetSecondFinanceTax() *string
	SetSecondFinanceTaxCertificateImgName(v string) *OperateFinanceTaxRequest
	GetSecondFinanceTaxCertificateImgName() *string
	SetSecondFinanceTaxCertificateImgUrl(v string) *OperateFinanceTaxRequest
	GetSecondFinanceTaxCertificateImgUrl() *string
	SetFinanceTaxCertificateImgUrl(v string) *OperateFinanceTaxRequest
	GetFinanceTaxCertificateImgUrl() *string
}

type OperateFinanceTaxRequest struct {
	// This parameter is required.
	FinanceTax                   *string `json:"FinanceTax,omitempty" xml:"FinanceTax,omitempty"`
	FinanceTaxCertificateImgName *string `json:"FinanceTaxCertificateImgName,omitempty" xml:"FinanceTaxCertificateImgName,omitempty"`
	// This parameter is required.
	HId                                *int64  `json:"HId,omitempty" xml:"HId,omitempty"`
	SecondFinanceTax                   *string `json:"SecondFinanceTax,omitempty" xml:"SecondFinanceTax,omitempty"`
	SecondFinanceTaxCertificateImgName *string `json:"SecondFinanceTaxCertificateImgName,omitempty" xml:"SecondFinanceTaxCertificateImgName,omitempty"`
	SecondFinanceTaxCertificateImgUrl  *string `json:"SecondFinanceTaxCertificateImgUrl,omitempty" xml:"SecondFinanceTaxCertificateImgUrl,omitempty"`
	FinanceTaxCertificateImgUrl        *string `json:"financeTaxCertificateImgUrl,omitempty" xml:"financeTaxCertificateImgUrl,omitempty"`
}

func (s OperateFinanceTaxRequest) String() string {
	return dara.Prettify(s)
}

func (s OperateFinanceTaxRequest) GoString() string {
	return s.String()
}

func (s *OperateFinanceTaxRequest) GetFinanceTax() *string {
	return s.FinanceTax
}

func (s *OperateFinanceTaxRequest) GetFinanceTaxCertificateImgName() *string {
	return s.FinanceTaxCertificateImgName
}

func (s *OperateFinanceTaxRequest) GetHId() *int64 {
	return s.HId
}

func (s *OperateFinanceTaxRequest) GetSecondFinanceTax() *string {
	return s.SecondFinanceTax
}

func (s *OperateFinanceTaxRequest) GetSecondFinanceTaxCertificateImgName() *string {
	return s.SecondFinanceTaxCertificateImgName
}

func (s *OperateFinanceTaxRequest) GetSecondFinanceTaxCertificateImgUrl() *string {
	return s.SecondFinanceTaxCertificateImgUrl
}

func (s *OperateFinanceTaxRequest) GetFinanceTaxCertificateImgUrl() *string {
	return s.FinanceTaxCertificateImgUrl
}

func (s *OperateFinanceTaxRequest) SetFinanceTax(v string) *OperateFinanceTaxRequest {
	s.FinanceTax = &v
	return s
}

func (s *OperateFinanceTaxRequest) SetFinanceTaxCertificateImgName(v string) *OperateFinanceTaxRequest {
	s.FinanceTaxCertificateImgName = &v
	return s
}

func (s *OperateFinanceTaxRequest) SetHId(v int64) *OperateFinanceTaxRequest {
	s.HId = &v
	return s
}

func (s *OperateFinanceTaxRequest) SetSecondFinanceTax(v string) *OperateFinanceTaxRequest {
	s.SecondFinanceTax = &v
	return s
}

func (s *OperateFinanceTaxRequest) SetSecondFinanceTaxCertificateImgName(v string) *OperateFinanceTaxRequest {
	s.SecondFinanceTaxCertificateImgName = &v
	return s
}

func (s *OperateFinanceTaxRequest) SetSecondFinanceTaxCertificateImgUrl(v string) *OperateFinanceTaxRequest {
	s.SecondFinanceTaxCertificateImgUrl = &v
	return s
}

func (s *OperateFinanceTaxRequest) SetFinanceTaxCertificateImgUrl(v string) *OperateFinanceTaxRequest {
	s.FinanceTaxCertificateImgUrl = &v
	return s
}

func (s *OperateFinanceTaxRequest) Validate() error {
	return dara.Validate(s)
}
