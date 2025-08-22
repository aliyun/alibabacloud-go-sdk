// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainCcActivityLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDcdnDomainCcActivityLogRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeDcdnDomainCcActivityLogRequest
	GetEndTime() *string
	SetPageNumber(v int64) *DescribeDcdnDomainCcActivityLogRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeDcdnDomainCcActivityLogRequest
	GetPageSize() *int64
	SetRuleName(v string) *DescribeDcdnDomainCcActivityLogRequest
	GetRuleName() *string
	SetStartTime(v string) *DescribeDcdnDomainCcActivityLogRequest
	GetStartTime() *string
	SetTriggerObject(v string) *DescribeDcdnDomainCcActivityLogRequest
	GetTriggerObject() *string
	SetValue(v string) *DescribeDcdnDomainCcActivityLogRequest
	GetValue() *string
}

type DescribeDcdnDomainCcActivityLogRequest struct {
	// The accelerated domain name. You can specify one or more domain names. Separate multiple domain names with commas (,).
	//
	// If you leave this parameter empty, the data of all domain names is queried.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// The end time must be later than the start time.
	//
	// example:
	//
	// 2015-12-10T21:05:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The page number of the page returned. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Default value: **30**.
	//
	// example:
	//
	// 30
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the defense rule.
	//
	// 	- default_normal in normal mode
	//
	// 	- default_attack in emergency mode
	//
	// 	- A custom rule name in custom mode. Example: test2.
	//
	// If you leave this parameter empty, events that triggered rate limiting based on all rules are queried.
	//
	// example:
	//
	// test2
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// The resolution of the queried data is 5 minutes.
	//
	// If you leave this parameter empty, the data collected over the last 24 hours is queried.
	//
	// example:
	//
	// 2015-12-10T20:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The trigger of rate limiting by which you want to query data.
	//
	// If you leave this parameter empty, all events that triggered rate limiting are queried.
	//
	// example:
	//
	// IP
	TriggerObject *string `json:"TriggerObject,omitempty" xml:"TriggerObject,omitempty"`
	// The value of the object that triggered rate limiting.
	//
	// If you leave this parameter empty, events that triggered rate limiting based on all rules are queried.
	//
	// example:
	//
	// 10.10.10.10
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDcdnDomainCcActivityLogRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainCcActivityLogRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainCcActivityLogRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnDomainCcActivityLogRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDcdnDomainCcActivityLogRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeDcdnDomainCcActivityLogRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeDcdnDomainCcActivityLogRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *DescribeDcdnDomainCcActivityLogRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDcdnDomainCcActivityLogRequest) GetTriggerObject() *string {
	return s.TriggerObject
}

func (s *DescribeDcdnDomainCcActivityLogRequest) GetValue() *string {
	return s.Value
}

func (s *DescribeDcdnDomainCcActivityLogRequest) SetDomainName(v string) *DescribeDcdnDomainCcActivityLogRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainCcActivityLogRequest) SetEndTime(v string) *DescribeDcdnDomainCcActivityLogRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainCcActivityLogRequest) SetPageNumber(v int64) *DescribeDcdnDomainCcActivityLogRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDcdnDomainCcActivityLogRequest) SetPageSize(v int64) *DescribeDcdnDomainCcActivityLogRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDcdnDomainCcActivityLogRequest) SetRuleName(v string) *DescribeDcdnDomainCcActivityLogRequest {
	s.RuleName = &v
	return s
}

func (s *DescribeDcdnDomainCcActivityLogRequest) SetStartTime(v string) *DescribeDcdnDomainCcActivityLogRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnDomainCcActivityLogRequest) SetTriggerObject(v string) *DescribeDcdnDomainCcActivityLogRequest {
	s.TriggerObject = &v
	return s
}

func (s *DescribeDcdnDomainCcActivityLogRequest) SetValue(v string) *DescribeDcdnDomainCcActivityLogRequest {
	s.Value = &v
	return s
}

func (s *DescribeDcdnDomainCcActivityLogRequest) Validate() error {
	return dara.Validate(s)
}
