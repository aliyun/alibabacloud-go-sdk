// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainCcActivityLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetActivityLog(v []*DescribeDomainCcActivityLogResponseBodyActivityLog) *DescribeDomainCcActivityLogResponseBody
	GetActivityLog() []*DescribeDomainCcActivityLogResponseBodyActivityLog
	SetPageIndex(v int64) *DescribeDomainCcActivityLogResponseBody
	GetPageIndex() *int64
	SetPageSize(v int64) *DescribeDomainCcActivityLogResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *DescribeDomainCcActivityLogResponseBody
	GetRequestId() *string
	SetTotal(v int64) *DescribeDomainCcActivityLogResponseBody
	GetTotal() *int64
}

type DescribeDomainCcActivityLogResponseBody struct {
	// The list of rate limiting logs.
	ActivityLog []*DescribeDomainCcActivityLogResponseBodyActivityLog `json:"ActivityLog,omitempty" xml:"ActivityLog,omitempty" type:"Repeated"`
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
	// The ID of the request.
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

func (s DescribeDomainCcActivityLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainCcActivityLogResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainCcActivityLogResponseBody) GetActivityLog() []*DescribeDomainCcActivityLogResponseBodyActivityLog {
	return s.ActivityLog
}

func (s *DescribeDomainCcActivityLogResponseBody) GetPageIndex() *int64 {
	return s.PageIndex
}

func (s *DescribeDomainCcActivityLogResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeDomainCcActivityLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainCcActivityLogResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *DescribeDomainCcActivityLogResponseBody) SetActivityLog(v []*DescribeDomainCcActivityLogResponseBodyActivityLog) *DescribeDomainCcActivityLogResponseBody {
	s.ActivityLog = v
	return s
}

func (s *DescribeDomainCcActivityLogResponseBody) SetPageIndex(v int64) *DescribeDomainCcActivityLogResponseBody {
	s.PageIndex = &v
	return s
}

func (s *DescribeDomainCcActivityLogResponseBody) SetPageSize(v int64) *DescribeDomainCcActivityLogResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDomainCcActivityLogResponseBody) SetRequestId(v string) *DescribeDomainCcActivityLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainCcActivityLogResponseBody) SetTotal(v int64) *DescribeDomainCcActivityLogResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeDomainCcActivityLogResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainCcActivityLogResponseBodyActivityLog struct {
	// The action that was triggered.
	//
	// example:
	//
	// deny
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The accelerated domain name.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The name of the rule based on which rate limiting was triggered.
	//
	// example:
	//
	// test
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The timestamp of the data returned.
	//
	// example:
	//
	// 2015-12-10T20:00:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	// The trigger of rate limiting.
	//
	// example:
	//
	// Ip
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
	// 1.2.3.4
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDomainCcActivityLogResponseBodyActivityLog) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainCcActivityLogResponseBodyActivityLog) GoString() string {
	return s.String()
}

func (s *DescribeDomainCcActivityLogResponseBodyActivityLog) GetAction() *string {
	return s.Action
}

func (s *DescribeDomainCcActivityLogResponseBodyActivityLog) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDomainCcActivityLogResponseBodyActivityLog) GetRuleName() *string {
	return s.RuleName
}

func (s *DescribeDomainCcActivityLogResponseBodyActivityLog) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeDomainCcActivityLogResponseBodyActivityLog) GetTriggerObject() *string {
	return s.TriggerObject
}

func (s *DescribeDomainCcActivityLogResponseBodyActivityLog) GetTtl() *int64 {
	return s.Ttl
}

func (s *DescribeDomainCcActivityLogResponseBodyActivityLog) GetValue() *string {
	return s.Value
}

func (s *DescribeDomainCcActivityLogResponseBodyActivityLog) SetAction(v string) *DescribeDomainCcActivityLogResponseBodyActivityLog {
	s.Action = &v
	return s
}

func (s *DescribeDomainCcActivityLogResponseBodyActivityLog) SetDomainName(v string) *DescribeDomainCcActivityLogResponseBodyActivityLog {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainCcActivityLogResponseBodyActivityLog) SetRuleName(v string) *DescribeDomainCcActivityLogResponseBodyActivityLog {
	s.RuleName = &v
	return s
}

func (s *DescribeDomainCcActivityLogResponseBodyActivityLog) SetTimeStamp(v string) *DescribeDomainCcActivityLogResponseBodyActivityLog {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDomainCcActivityLogResponseBodyActivityLog) SetTriggerObject(v string) *DescribeDomainCcActivityLogResponseBodyActivityLog {
	s.TriggerObject = &v
	return s
}

func (s *DescribeDomainCcActivityLogResponseBodyActivityLog) SetTtl(v int64) *DescribeDomainCcActivityLogResponseBodyActivityLog {
	s.Ttl = &v
	return s
}

func (s *DescribeDomainCcActivityLogResponseBodyActivityLog) SetValue(v string) *DescribeDomainCcActivityLogResponseBodyActivityLog {
	s.Value = &v
	return s
}

func (s *DescribeDomainCcActivityLogResponseBodyActivityLog) Validate() error {
	return dara.Validate(s)
}
