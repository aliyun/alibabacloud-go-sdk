// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSdlStatisticRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *DescribeSdlStatisticRequest
	GetEndTime() *int64
	SetLang(v string) *DescribeSdlStatisticRequest
	GetLang() *string
	SetStartTime(v int64) *DescribeSdlStatisticRequest
	GetStartTime() *int64
}

type DescribeSdlStatisticRequest struct {
	// example:
	//
	// 1748916368
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 1656664560
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeSdlStatisticRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSdlStatisticRequest) GoString() string {
	return s.String()
}

func (s *DescribeSdlStatisticRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeSdlStatisticRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeSdlStatisticRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeSdlStatisticRequest) SetEndTime(v int64) *DescribeSdlStatisticRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSdlStatisticRequest) SetLang(v string) *DescribeSdlStatisticRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSdlStatisticRequest) SetStartTime(v int64) *DescribeSdlStatisticRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSdlStatisticRequest) Validate() error {
	return dara.Validate(s)
}
