// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDohUserInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndDate(v string) *DescribeDohUserInfoRequest
	GetEndDate() *string
	SetLang(v string) *DescribeDohUserInfoRequest
	GetLang() *string
	SetStartDate(v string) *DescribeDohUserInfoRequest
	GetStartDate() *string
}

type DescribeDohUserInfoRequest struct {
	// The end date of the query. Use the \\`YYYY-MM-DD\\` format.
	//
	// If you do not specify this parameter, the current date is used.
	//
	// example:
	//
	// 2019-07-04
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The language.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The start date of the query. Use the \\`YYYY-MM-DD\\` format.
	//
	// You can query data from the last 90 days. The date must be within the last 90 days.
	//
	// example:
	//
	// 2019-07-04
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s DescribeDohUserInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDohUserInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeDohUserInfoRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *DescribeDohUserInfoRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeDohUserInfoRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *DescribeDohUserInfoRequest) SetEndDate(v string) *DescribeDohUserInfoRequest {
	s.EndDate = &v
	return s
}

func (s *DescribeDohUserInfoRequest) SetLang(v string) *DescribeDohUserInfoRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDohUserInfoRequest) SetStartDate(v string) *DescribeDohUserInfoRequest {
	s.StartDate = &v
	return s
}

func (s *DescribeDohUserInfoRequest) Validate() error {
	return dara.Validate(s)
}
