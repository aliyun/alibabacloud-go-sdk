// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainCcActivityLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDomainCcActivityLogRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeDomainCcActivityLogRequest
	GetEndTime() *string
	SetPageNumber(v int64) *DescribeDomainCcActivityLogRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeDomainCcActivityLogRequest
	GetPageSize() *int64
	SetRuleName(v string) *DescribeDomainCcActivityLogRequest
	GetRuleName() *string
	SetStartTime(v string) *DescribeDomainCcActivityLogRequest
	GetStartTime() *string
	SetTriggerObject(v string) *DescribeDomainCcActivityLogRequest
	GetTriggerObject() *string
	SetValue(v string) *DescribeDomainCcActivityLogRequest
	GetValue() *string
}

type DescribeDomainCcActivityLogRequest struct {
	// The accelerated domain name. You can specify multiple domain names and separate them with commas (,).
	//
	// If you do not specify this parameter, data of all accelerated domain names under your account is queried.
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
	// 2018-12-10T21:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The page number of the page to return. Default value: **1**.
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
	// A custom rule name. Valid values:
	//
	// 	- default_normal: rule for the normal mode
	//
	// 	- default_attack: rule for the emergency mode
	//
	// If you leave this parameter empty, events that triggered rate limiting based on all rules are queried.
	//
	// example:
	//
	// test2
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// The minimum time granularity of data collection is 5 minutes.
	//
	// If you leave this parameter empty, the data collected over the last 24 hours is queried.
	//
	// example:
	//
	// 2018-12-10T20:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The trigger of rate limiting by which you want to query data.
	//
	// If you leave this parameter empty, all events that triggered rate limiting are queried.
	//
	// example:
	//
	// IP
	TriggerObject *string `json:"TriggerObject,omitempty" xml:"TriggerObject,omitempty"`
	// The value of the trigger.
	//
	// If you leave this parameter empty, all events recorded for the trigger are queried.
	//
	// example:
	//
	// 1.2.3.4
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDomainCcActivityLogRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainCcActivityLogRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainCcActivityLogRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDomainCcActivityLogRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDomainCcActivityLogRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeDomainCcActivityLogRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeDomainCcActivityLogRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *DescribeDomainCcActivityLogRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDomainCcActivityLogRequest) GetTriggerObject() *string {
	return s.TriggerObject
}

func (s *DescribeDomainCcActivityLogRequest) GetValue() *string {
	return s.Value
}

func (s *DescribeDomainCcActivityLogRequest) SetDomainName(v string) *DescribeDomainCcActivityLogRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainCcActivityLogRequest) SetEndTime(v string) *DescribeDomainCcActivityLogRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainCcActivityLogRequest) SetPageNumber(v int64) *DescribeDomainCcActivityLogRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDomainCcActivityLogRequest) SetPageSize(v int64) *DescribeDomainCcActivityLogRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDomainCcActivityLogRequest) SetRuleName(v string) *DescribeDomainCcActivityLogRequest {
	s.RuleName = &v
	return s
}

func (s *DescribeDomainCcActivityLogRequest) SetStartTime(v string) *DescribeDomainCcActivityLogRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainCcActivityLogRequest) SetTriggerObject(v string) *DescribeDomainCcActivityLogRequest {
	s.TriggerObject = &v
	return s
}

func (s *DescribeDomainCcActivityLogRequest) SetValue(v string) *DescribeDomainCcActivityLogRequest {
	s.Value = &v
	return s
}

func (s *DescribeDomainCcActivityLogRequest) Validate() error {
	return dara.Validate(s)
}
