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
	// The end time for the query. Format: YYYY-MM-DD
	//
	// If you do not specify this parameter, the default value is the time when you perform the query.
	//
	// example:
	//
	// 2019-07-04
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The language in which you want the values of some response parameters to be returned. These response parameters support multiple languages.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The start time for the query. Format: YYYY-MM-DD
	//
	// You can query the user information of the last 90 days only. `Set the parameter to a value no earlier than 90 days from the current time`.
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
