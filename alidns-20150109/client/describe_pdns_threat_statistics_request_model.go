// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePdnsThreatStatisticsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirection(v string) *DescribePdnsThreatStatisticsRequest
	GetDirection() *string
	SetDomainName(v string) *DescribePdnsThreatStatisticsRequest
	GetDomainName() *string
	SetEndDate(v string) *DescribePdnsThreatStatisticsRequest
	GetEndDate() *string
	SetLang(v string) *DescribePdnsThreatStatisticsRequest
	GetLang() *string
	SetOrderBy(v string) *DescribePdnsThreatStatisticsRequest
	GetOrderBy() *string
	SetPageNumber(v int64) *DescribePdnsThreatStatisticsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribePdnsThreatStatisticsRequest
	GetPageSize() *int64
	SetStartDate(v string) *DescribePdnsThreatStatisticsRequest
	GetStartDate() *string
	SetSubDomain(v string) *DescribePdnsThreatStatisticsRequest
	GetSubDomain() *string
	SetThreatLevel(v string) *DescribePdnsThreatStatisticsRequest
	GetThreatLevel() *string
	SetThreatSourceIp(v string) *DescribePdnsThreatStatisticsRequest
	GetThreatSourceIp() *string
	SetThreatType(v string) *DescribePdnsThreatStatisticsRequest
	GetThreatType() *string
	SetType(v string) *DescribePdnsThreatStatisticsRequest
	GetType() *string
}

type DescribePdnsThreatStatisticsRequest struct {
	Direction      *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	DomainName     *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndDate        *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	Lang           *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	OrderBy        *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	PageNumber     *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StartDate      *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	SubDomain      *string `json:"SubDomain,omitempty" xml:"SubDomain,omitempty"`
	ThreatLevel    *string `json:"ThreatLevel,omitempty" xml:"ThreatLevel,omitempty"`
	ThreatSourceIp *string `json:"ThreatSourceIp,omitempty" xml:"ThreatSourceIp,omitempty"`
	ThreatType     *string `json:"ThreatType,omitempty" xml:"ThreatType,omitempty"`
	Type           *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribePdnsThreatStatisticsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePdnsThreatStatisticsRequest) GoString() string {
	return s.String()
}

func (s *DescribePdnsThreatStatisticsRequest) GetDirection() *string {
	return s.Direction
}

func (s *DescribePdnsThreatStatisticsRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribePdnsThreatStatisticsRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *DescribePdnsThreatStatisticsRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribePdnsThreatStatisticsRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *DescribePdnsThreatStatisticsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribePdnsThreatStatisticsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribePdnsThreatStatisticsRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *DescribePdnsThreatStatisticsRequest) GetSubDomain() *string {
	return s.SubDomain
}

func (s *DescribePdnsThreatStatisticsRequest) GetThreatLevel() *string {
	return s.ThreatLevel
}

func (s *DescribePdnsThreatStatisticsRequest) GetThreatSourceIp() *string {
	return s.ThreatSourceIp
}

func (s *DescribePdnsThreatStatisticsRequest) GetThreatType() *string {
	return s.ThreatType
}

func (s *DescribePdnsThreatStatisticsRequest) GetType() *string {
	return s.Type
}

func (s *DescribePdnsThreatStatisticsRequest) SetDirection(v string) *DescribePdnsThreatStatisticsRequest {
	s.Direction = &v
	return s
}

func (s *DescribePdnsThreatStatisticsRequest) SetDomainName(v string) *DescribePdnsThreatStatisticsRequest {
	s.DomainName = &v
	return s
}

func (s *DescribePdnsThreatStatisticsRequest) SetEndDate(v string) *DescribePdnsThreatStatisticsRequest {
	s.EndDate = &v
	return s
}

func (s *DescribePdnsThreatStatisticsRequest) SetLang(v string) *DescribePdnsThreatStatisticsRequest {
	s.Lang = &v
	return s
}

func (s *DescribePdnsThreatStatisticsRequest) SetOrderBy(v string) *DescribePdnsThreatStatisticsRequest {
	s.OrderBy = &v
	return s
}

func (s *DescribePdnsThreatStatisticsRequest) SetPageNumber(v int64) *DescribePdnsThreatStatisticsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribePdnsThreatStatisticsRequest) SetPageSize(v int64) *DescribePdnsThreatStatisticsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribePdnsThreatStatisticsRequest) SetStartDate(v string) *DescribePdnsThreatStatisticsRequest {
	s.StartDate = &v
	return s
}

func (s *DescribePdnsThreatStatisticsRequest) SetSubDomain(v string) *DescribePdnsThreatStatisticsRequest {
	s.SubDomain = &v
	return s
}

func (s *DescribePdnsThreatStatisticsRequest) SetThreatLevel(v string) *DescribePdnsThreatStatisticsRequest {
	s.ThreatLevel = &v
	return s
}

func (s *DescribePdnsThreatStatisticsRequest) SetThreatSourceIp(v string) *DescribePdnsThreatStatisticsRequest {
	s.ThreatSourceIp = &v
	return s
}

func (s *DescribePdnsThreatStatisticsRequest) SetThreatType(v string) *DescribePdnsThreatStatisticsRequest {
	s.ThreatType = &v
	return s
}

func (s *DescribePdnsThreatStatisticsRequest) SetType(v string) *DescribePdnsThreatStatisticsRequest {
	s.Type = &v
	return s
}

func (s *DescribePdnsThreatStatisticsRequest) Validate() error {
	return dara.Validate(s)
}
