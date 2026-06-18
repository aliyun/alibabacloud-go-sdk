// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRenewalRateListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFiscalYearAndQuarter(v string) *GetRenewalRateListRequest
	GetFiscalYearAndQuarter() *string
}

type GetRenewalRateListRequest struct {
	// The fiscal year and quarter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2025Q4
	FiscalYearAndQuarter *string `json:"FiscalYearAndQuarter,omitempty" xml:"FiscalYearAndQuarter,omitempty"`
}

func (s GetRenewalRateListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRenewalRateListRequest) GoString() string {
	return s.String()
}

func (s *GetRenewalRateListRequest) GetFiscalYearAndQuarter() *string {
	return s.FiscalYearAndQuarter
}

func (s *GetRenewalRateListRequest) SetFiscalYearAndQuarter(v string) *GetRenewalRateListRequest {
	s.FiscalYearAndQuarter = &v
	return s
}

func (s *GetRenewalRateListRequest) Validate() error {
	return dara.Validate(s)
}
