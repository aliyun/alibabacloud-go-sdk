// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnUserSecDropByMinuteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDcdnUserSecDropByMinuteRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeDcdnUserSecDropByMinuteRequest
	GetEndTime() *string
	SetLang(v string) *DescribeDcdnUserSecDropByMinuteRequest
	GetLang() *string
	SetObject(v string) *DescribeDcdnUserSecDropByMinuteRequest
	GetObject() *string
	SetPageNumber(v int64) *DescribeDcdnUserSecDropByMinuteRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeDcdnUserSecDropByMinuteRequest
	GetPageSize() *int64
	SetRuleName(v string) *DescribeDcdnUserSecDropByMinuteRequest
	GetRuleName() *string
	SetSecFunc(v string) *DescribeDcdnUserSecDropByMinuteRequest
	GetSecFunc() *string
	SetStartTime(v string) *DescribeDcdnUserSecDropByMinuteRequest
	GetStartTime() *string
}

type DescribeDcdnUserSecDropByMinuteRequest struct {
	// The domain name.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC. Example: 2006-01-02T15:05:04Z.
	//
	// > The end time must be later than the start time.
	//
	// example:
	//
	// 2006-01-02T15:05:04Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The language. Valid values: en and zh. Default value: en
	//
	// This parameter is required.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The object that triggered rate limiting.
	//
	// example:
	//
	// robot_fingerprint_ai
	Object *string `json:"Object,omitempty" xml:"Object,omitempty"`
	// The number of the page to return. Pages start from page 1.
	//
	// example:
	//
	// 10
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Maximum value: 100.
	//
	// example:
	//
	// 2
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The rule that was triggered.
	//
	// example:
	//
	// robot_ai
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The name of the security feature.
	//
	// This parameter is required.
	//
	// example:
	//
	// robot
	SecFunc *string `json:"SecFunc,omitempty" xml:"SecFunc,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC. Example: 2006-01-02T15:04:04Z.
	//
	// example:
	//
	// 2006-01-02T15:04:04Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnUserSecDropByMinuteRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnUserSecDropByMinuteRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserSecDropByMinuteRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnUserSecDropByMinuteRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDcdnUserSecDropByMinuteRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeDcdnUserSecDropByMinuteRequest) GetObject() *string {
	return s.Object
}

func (s *DescribeDcdnUserSecDropByMinuteRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeDcdnUserSecDropByMinuteRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeDcdnUserSecDropByMinuteRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *DescribeDcdnUserSecDropByMinuteRequest) GetSecFunc() *string {
	return s.SecFunc
}

func (s *DescribeDcdnUserSecDropByMinuteRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDcdnUserSecDropByMinuteRequest) SetDomainName(v string) *DescribeDcdnUserSecDropByMinuteRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnUserSecDropByMinuteRequest) SetEndTime(v string) *DescribeDcdnUserSecDropByMinuteRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnUserSecDropByMinuteRequest) SetLang(v string) *DescribeDcdnUserSecDropByMinuteRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDcdnUserSecDropByMinuteRequest) SetObject(v string) *DescribeDcdnUserSecDropByMinuteRequest {
	s.Object = &v
	return s
}

func (s *DescribeDcdnUserSecDropByMinuteRequest) SetPageNumber(v int64) *DescribeDcdnUserSecDropByMinuteRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDcdnUserSecDropByMinuteRequest) SetPageSize(v int64) *DescribeDcdnUserSecDropByMinuteRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDcdnUserSecDropByMinuteRequest) SetRuleName(v string) *DescribeDcdnUserSecDropByMinuteRequest {
	s.RuleName = &v
	return s
}

func (s *DescribeDcdnUserSecDropByMinuteRequest) SetSecFunc(v string) *DescribeDcdnUserSecDropByMinuteRequest {
	s.SecFunc = &v
	return s
}

func (s *DescribeDcdnUserSecDropByMinuteRequest) SetStartTime(v string) *DescribeDcdnUserSecDropByMinuteRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnUserSecDropByMinuteRequest) Validate() error {
	return dara.Validate(s)
}
