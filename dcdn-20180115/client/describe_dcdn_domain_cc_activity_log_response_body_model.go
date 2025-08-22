// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainCcActivityLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetActivityLog(v []*DescribeDcdnDomainCcActivityLogResponseBodyActivityLog) *DescribeDcdnDomainCcActivityLogResponseBody
	GetActivityLog() []*DescribeDcdnDomainCcActivityLogResponseBodyActivityLog
	SetPageIndex(v int64) *DescribeDcdnDomainCcActivityLogResponseBody
	GetPageIndex() *int64
	SetPageSize(v int64) *DescribeDcdnDomainCcActivityLogResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *DescribeDcdnDomainCcActivityLogResponseBody
	GetRequestId() *string
	SetTotal(v int64) *DescribeDcdnDomainCcActivityLogResponseBody
	GetTotal() *int64
}

type DescribeDcdnDomainCcActivityLogResponseBody struct {
	// The log data of the event that triggered rate limiting.
	ActivityLog []*DescribeDcdnDomainCcActivityLogResponseBodyActivityLog `json:"ActivityLog,omitempty" xml:"ActivityLog,omitempty" type:"Repeated"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageIndex *int64 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 30
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 3C6CCEC4-6B88-4D4A-93E4-D47B3D92CF8F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 20
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeDcdnDomainCcActivityLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainCcActivityLogResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainCcActivityLogResponseBody) GetActivityLog() []*DescribeDcdnDomainCcActivityLogResponseBodyActivityLog {
	return s.ActivityLog
}

func (s *DescribeDcdnDomainCcActivityLogResponseBody) GetPageIndex() *int64 {
	return s.PageIndex
}

func (s *DescribeDcdnDomainCcActivityLogResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeDcdnDomainCcActivityLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnDomainCcActivityLogResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *DescribeDcdnDomainCcActivityLogResponseBody) SetActivityLog(v []*DescribeDcdnDomainCcActivityLogResponseBodyActivityLog) *DescribeDcdnDomainCcActivityLogResponseBody {
	s.ActivityLog = v
	return s
}

func (s *DescribeDcdnDomainCcActivityLogResponseBody) SetPageIndex(v int64) *DescribeDcdnDomainCcActivityLogResponseBody {
	s.PageIndex = &v
	return s
}

func (s *DescribeDcdnDomainCcActivityLogResponseBody) SetPageSize(v int64) *DescribeDcdnDomainCcActivityLogResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDcdnDomainCcActivityLogResponseBody) SetRequestId(v string) *DescribeDcdnDomainCcActivityLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnDomainCcActivityLogResponseBody) SetTotal(v int64) *DescribeDcdnDomainCcActivityLogResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeDcdnDomainCcActivityLogResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainCcActivityLogResponseBodyActivityLog struct {
	// The action that was triggered.
	//
	// example:
	//
	// deny
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The accelerated domain name whose ICP filing status you want to update.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The name of the rule that was triggered.
	//
	// example:
	//
	// test2
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The timestamp of the data returned.
	//
	// example:
	//
	// 2015-12-10T20:00:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	// The object that triggered the blocking event.
	//
	// example:
	//
	// IP
	TriggerObject *string `json:"TriggerObject,omitempty" xml:"TriggerObject,omitempty"`
	// The period of time during which rate limiting remains effective.
	//
	// example:
	//
	// 300
	Ttl *int64 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	// The value of the trigger for rate limiting.
	//
	// example:
	//
	// 10.10.10.10
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDcdnDomainCcActivityLogResponseBodyActivityLog) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainCcActivityLogResponseBodyActivityLog) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainCcActivityLogResponseBodyActivityLog) GetAction() *string {
	return s.Action
}

func (s *DescribeDcdnDomainCcActivityLogResponseBodyActivityLog) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnDomainCcActivityLogResponseBodyActivityLog) GetRuleName() *string {
	return s.RuleName
}

func (s *DescribeDcdnDomainCcActivityLogResponseBodyActivityLog) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeDcdnDomainCcActivityLogResponseBodyActivityLog) GetTriggerObject() *string {
	return s.TriggerObject
}

func (s *DescribeDcdnDomainCcActivityLogResponseBodyActivityLog) GetTtl() *int64 {
	return s.Ttl
}

func (s *DescribeDcdnDomainCcActivityLogResponseBodyActivityLog) GetValue() *string {
	return s.Value
}

func (s *DescribeDcdnDomainCcActivityLogResponseBodyActivityLog) SetAction(v string) *DescribeDcdnDomainCcActivityLogResponseBodyActivityLog {
	s.Action = &v
	return s
}

func (s *DescribeDcdnDomainCcActivityLogResponseBodyActivityLog) SetDomainName(v string) *DescribeDcdnDomainCcActivityLogResponseBodyActivityLog {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainCcActivityLogResponseBodyActivityLog) SetRuleName(v string) *DescribeDcdnDomainCcActivityLogResponseBodyActivityLog {
	s.RuleName = &v
	return s
}

func (s *DescribeDcdnDomainCcActivityLogResponseBodyActivityLog) SetTimeStamp(v string) *DescribeDcdnDomainCcActivityLogResponseBodyActivityLog {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDcdnDomainCcActivityLogResponseBodyActivityLog) SetTriggerObject(v string) *DescribeDcdnDomainCcActivityLogResponseBodyActivityLog {
	s.TriggerObject = &v
	return s
}

func (s *DescribeDcdnDomainCcActivityLogResponseBodyActivityLog) SetTtl(v int64) *DescribeDcdnDomainCcActivityLogResponseBodyActivityLog {
	s.Ttl = &v
	return s
}

func (s *DescribeDcdnDomainCcActivityLogResponseBodyActivityLog) SetValue(v string) *DescribeDcdnDomainCcActivityLogResponseBodyActivityLog {
	s.Value = &v
	return s
}

func (s *DescribeDcdnDomainCcActivityLogResponseBodyActivityLog) Validate() error {
	return dara.Validate(s)
}
