// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePdnsAccountSummaryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndDate(v string) *DescribePdnsAccountSummaryRequest
	GetEndDate() *string
	SetLang(v string) *DescribePdnsAccountSummaryRequest
	GetLang() *string
	SetStartDate(v string) *DescribePdnsAccountSummaryRequest
	GetStartDate() *string
}

type DescribePdnsAccountSummaryRequest struct {
	EndDate   *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	Lang      *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s DescribePdnsAccountSummaryRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePdnsAccountSummaryRequest) GoString() string {
	return s.String()
}

func (s *DescribePdnsAccountSummaryRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *DescribePdnsAccountSummaryRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribePdnsAccountSummaryRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *DescribePdnsAccountSummaryRequest) SetEndDate(v string) *DescribePdnsAccountSummaryRequest {
	s.EndDate = &v
	return s
}

func (s *DescribePdnsAccountSummaryRequest) SetLang(v string) *DescribePdnsAccountSummaryRequest {
	s.Lang = &v
	return s
}

func (s *DescribePdnsAccountSummaryRequest) SetStartDate(v string) *DescribePdnsAccountSummaryRequest {
	s.StartDate = &v
	return s
}

func (s *DescribePdnsAccountSummaryRequest) Validate() error {
	return dara.Validate(s)
}
