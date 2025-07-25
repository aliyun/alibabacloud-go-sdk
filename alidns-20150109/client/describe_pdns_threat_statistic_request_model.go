// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePdnsThreatStatisticRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndDate(v string) *DescribePdnsThreatStatisticRequest
	GetEndDate() *string
	SetLang(v string) *DescribePdnsThreatStatisticRequest
	GetLang() *string
	SetStartDate(v string) *DescribePdnsThreatStatisticRequest
	GetStartDate() *string
	SetThreatSourceIp(v string) *DescribePdnsThreatStatisticRequest
	GetThreatSourceIp() *string
}

type DescribePdnsThreatStatisticRequest struct {
	EndDate        *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	Lang           *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	StartDate      *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	ThreatSourceIp *string `json:"ThreatSourceIp,omitempty" xml:"ThreatSourceIp,omitempty"`
}

func (s DescribePdnsThreatStatisticRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePdnsThreatStatisticRequest) GoString() string {
	return s.String()
}

func (s *DescribePdnsThreatStatisticRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *DescribePdnsThreatStatisticRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribePdnsThreatStatisticRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *DescribePdnsThreatStatisticRequest) GetThreatSourceIp() *string {
	return s.ThreatSourceIp
}

func (s *DescribePdnsThreatStatisticRequest) SetEndDate(v string) *DescribePdnsThreatStatisticRequest {
	s.EndDate = &v
	return s
}

func (s *DescribePdnsThreatStatisticRequest) SetLang(v string) *DescribePdnsThreatStatisticRequest {
	s.Lang = &v
	return s
}

func (s *DescribePdnsThreatStatisticRequest) SetStartDate(v string) *DescribePdnsThreatStatisticRequest {
	s.StartDate = &v
	return s
}

func (s *DescribePdnsThreatStatisticRequest) SetThreatSourceIp(v string) *DescribePdnsThreatStatisticRequest {
	s.ThreatSourceIp = &v
	return s
}

func (s *DescribePdnsThreatStatisticRequest) Validate() error {
	return dara.Validate(s)
}
