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
	// The end of the time range to query. Specify the time in the YYYY-MM-DD format.
	//
	// The default value is the day when you perform the operation.
	//
	// example:
	//
	// 2019-07-04
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The language type.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The beginning of the time range to query. Specify the time in the YYYY-MM-DD format.
	//
	// You can query only the DNS records of the latest 90 days.`The value of StartDate must be greater than or equal to the difference between the current date and 90`.
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
