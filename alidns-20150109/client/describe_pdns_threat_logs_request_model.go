// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePdnsThreatLogsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndDate(v string) *DescribePdnsThreatLogsRequest
	GetEndDate() *string
	SetKeyword(v string) *DescribePdnsThreatLogsRequest
	GetKeyword() *string
	SetLang(v string) *DescribePdnsThreatLogsRequest
	GetLang() *string
	SetPageNumber(v int64) *DescribePdnsThreatLogsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribePdnsThreatLogsRequest
	GetPageSize() *int64
	SetStartDate(v string) *DescribePdnsThreatLogsRequest
	GetStartDate() *string
	SetThreatLevel(v string) *DescribePdnsThreatLogsRequest
	GetThreatLevel() *string
	SetThreatSourceIp(v string) *DescribePdnsThreatLogsRequest
	GetThreatSourceIp() *string
	SetThreatType(v string) *DescribePdnsThreatLogsRequest
	GetThreatType() *string
}

type DescribePdnsThreatLogsRequest struct {
	EndDate        *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	Keyword        *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	Lang           *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	PageNumber     *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StartDate      *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	ThreatLevel    *string `json:"ThreatLevel,omitempty" xml:"ThreatLevel,omitempty"`
	ThreatSourceIp *string `json:"ThreatSourceIp,omitempty" xml:"ThreatSourceIp,omitempty"`
	ThreatType     *string `json:"ThreatType,omitempty" xml:"ThreatType,omitempty"`
}

func (s DescribePdnsThreatLogsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePdnsThreatLogsRequest) GoString() string {
	return s.String()
}

func (s *DescribePdnsThreatLogsRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *DescribePdnsThreatLogsRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *DescribePdnsThreatLogsRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribePdnsThreatLogsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribePdnsThreatLogsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribePdnsThreatLogsRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *DescribePdnsThreatLogsRequest) GetThreatLevel() *string {
	return s.ThreatLevel
}

func (s *DescribePdnsThreatLogsRequest) GetThreatSourceIp() *string {
	return s.ThreatSourceIp
}

func (s *DescribePdnsThreatLogsRequest) GetThreatType() *string {
	return s.ThreatType
}

func (s *DescribePdnsThreatLogsRequest) SetEndDate(v string) *DescribePdnsThreatLogsRequest {
	s.EndDate = &v
	return s
}

func (s *DescribePdnsThreatLogsRequest) SetKeyword(v string) *DescribePdnsThreatLogsRequest {
	s.Keyword = &v
	return s
}

func (s *DescribePdnsThreatLogsRequest) SetLang(v string) *DescribePdnsThreatLogsRequest {
	s.Lang = &v
	return s
}

func (s *DescribePdnsThreatLogsRequest) SetPageNumber(v int64) *DescribePdnsThreatLogsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribePdnsThreatLogsRequest) SetPageSize(v int64) *DescribePdnsThreatLogsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribePdnsThreatLogsRequest) SetStartDate(v string) *DescribePdnsThreatLogsRequest {
	s.StartDate = &v
	return s
}

func (s *DescribePdnsThreatLogsRequest) SetThreatLevel(v string) *DescribePdnsThreatLogsRequest {
	s.ThreatLevel = &v
	return s
}

func (s *DescribePdnsThreatLogsRequest) SetThreatSourceIp(v string) *DescribePdnsThreatLogsRequest {
	s.ThreatSourceIp = &v
	return s
}

func (s *DescribePdnsThreatLogsRequest) SetThreatType(v string) *DescribePdnsThreatLogsRequest {
	s.ThreatType = &v
	return s
}

func (s *DescribePdnsThreatLogsRequest) Validate() error {
	return dara.Validate(s)
}
