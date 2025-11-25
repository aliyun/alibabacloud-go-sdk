// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSdlEventStatisticRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *DescribeSdlEventStatisticRequest
	GetEndTime() *int64
	SetLang(v string) *DescribeSdlEventStatisticRequest
	GetLang() *string
	SetStartTime(v int64) *DescribeSdlEventStatisticRequest
	GetStartTime() *int64
}

type DescribeSdlEventStatisticRequest struct {
	// example:
	//
	// 1732586712
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 1656750960
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeSdlEventStatisticRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSdlEventStatisticRequest) GoString() string {
	return s.String()
}

func (s *DescribeSdlEventStatisticRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeSdlEventStatisticRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeSdlEventStatisticRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeSdlEventStatisticRequest) SetEndTime(v int64) *DescribeSdlEventStatisticRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSdlEventStatisticRequest) SetLang(v string) *DescribeSdlEventStatisticRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSdlEventStatisticRequest) SetStartTime(v int64) *DescribeSdlEventStatisticRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSdlEventStatisticRequest) Validate() error {
	return dara.Validate(s)
}
