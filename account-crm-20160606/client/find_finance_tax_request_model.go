// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFindFinanceTaxRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHId(v int64) *FindFinanceTaxRequest
	GetHId() *int64
	SetTaxVersion(v string) *FindFinanceTaxRequest
	GetTaxVersion() *string
}

type FindFinanceTaxRequest struct {
	// This parameter is required.
	HId        *int64  `json:"HId,omitempty" xml:"HId,omitempty"`
	TaxVersion *string `json:"TaxVersion,omitempty" xml:"TaxVersion,omitempty"`
}

func (s FindFinanceTaxRequest) String() string {
	return dara.Prettify(s)
}

func (s FindFinanceTaxRequest) GoString() string {
	return s.String()
}

func (s *FindFinanceTaxRequest) GetHId() *int64 {
	return s.HId
}

func (s *FindFinanceTaxRequest) GetTaxVersion() *string {
	return s.TaxVersion
}

func (s *FindFinanceTaxRequest) SetHId(v int64) *FindFinanceTaxRequest {
	s.HId = &v
	return s
}

func (s *FindFinanceTaxRequest) SetTaxVersion(v string) *FindFinanceTaxRequest {
	s.TaxVersion = &v
	return s
}

func (s *FindFinanceTaxRequest) Validate() error {
	return dara.Validate(s)
}
