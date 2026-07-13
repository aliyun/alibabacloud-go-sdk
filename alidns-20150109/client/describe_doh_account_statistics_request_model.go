// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDohAccountStatisticsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndDate(v string) *DescribeDohAccountStatisticsRequest
	GetEndDate() *string
	SetLang(v string) *DescribeDohAccountStatisticsRequest
	GetLang() *string
	SetStartDate(v string) *DescribeDohAccountStatisticsRequest
	GetStartDate() *string
}

type DescribeDohAccountStatisticsRequest struct {
	// The end date of the query. The format is YYYY-MM-DD.
	//
	// The default value is the current date.
	//
	// example:
	//
	// 2019-07-04
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The language of the request and response. The default value is **zh**. Valid values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The start date of the query. The format is YYYY-MM-DD.
	//
	// You can query data from the last 90 days only.
	//
	// example:
	//
	// 2019-07-04
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s DescribeDohAccountStatisticsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDohAccountStatisticsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDohAccountStatisticsRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *DescribeDohAccountStatisticsRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeDohAccountStatisticsRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *DescribeDohAccountStatisticsRequest) SetEndDate(v string) *DescribeDohAccountStatisticsRequest {
	s.EndDate = &v
	return s
}

func (s *DescribeDohAccountStatisticsRequest) SetLang(v string) *DescribeDohAccountStatisticsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDohAccountStatisticsRequest) SetStartDate(v string) *DescribeDohAccountStatisticsRequest {
	s.StartDate = &v
	return s
}

func (s *DescribeDohAccountStatisticsRequest) Validate() error {
	return dara.Validate(s)
}
