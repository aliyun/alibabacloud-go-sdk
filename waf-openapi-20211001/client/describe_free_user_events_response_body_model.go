// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFreeUserEventsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEvent(v []*DescribeFreeUserEventsResponseBodyEvent) *DescribeFreeUserEventsResponseBody
	GetEvent() []*DescribeFreeUserEventsResponseBodyEvent
	SetRequestId(v string) *DescribeFreeUserEventsResponseBody
	GetRequestId() *string
}

type DescribeFreeUserEventsResponseBody struct {
	// The security events on which basic detection is performed.
	Event []*DescribeFreeUserEventsResponseBodyEvent `json:"Event,omitempty" xml:"Event,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 8A2DF88D-90C2-56E9-B8D5-36BB9646791C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeFreeUserEventsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeFreeUserEventsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFreeUserEventsResponseBody) GetEvent() []*DescribeFreeUserEventsResponseBodyEvent {
	return s.Event
}

func (s *DescribeFreeUserEventsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeFreeUserEventsResponseBody) SetEvent(v []*DescribeFreeUserEventsResponseBodyEvent) *DescribeFreeUserEventsResponseBody {
	s.Event = v
	return s
}

func (s *DescribeFreeUserEventsResponseBody) SetRequestId(v string) *DescribeFreeUserEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFreeUserEventsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeFreeUserEventsResponseBodyEvent struct {
	// The API.
	//
	// example:
	//
	// /api/login
	ApiFormat *string `json:"ApiFormat,omitempty" xml:"ApiFormat,omitempty"`
	// The attacker IP address.
	//
	// example:
	//
	// 104.234.140.**
	AttackIP *string `json:"AttackIP,omitempty" xml:"AttackIP,omitempty"`
	// The time at which the attack was launched. The value is a UNIX timestamp displayed in UTC. Unit: seconds.
	//
	// example:
	//
	// 1683703260
	AttackTime *int64 `json:"AttackTime,omitempty" xml:"AttackTime,omitempty"`
	// The domain name of the API.
	//
	// example:
	//
	// www.***.cn
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The severity level of the security event. Valid values:
	//
	// 	- **high**
	//
	// 	- **medium**
	//
	// 	- **low**
	//
	// example:
	//
	// high
	EventLevel *string `json:"EventLevel,omitempty" xml:"EventLevel,omitempty"`
	// The type of the security event.
	//
	// >  You can call the [DescribeApisecRules](https://help.aliyun.com/document_detail/2859155.html) operation to query the supported types of security events.
	//
	// example:
	//
	// Event_DataTraverse
	EventTag *string `json:"EventTag,omitempty" xml:"EventTag,omitempty"`
}

func (s DescribeFreeUserEventsResponseBodyEvent) String() string {
	return dara.Prettify(s)
}

func (s DescribeFreeUserEventsResponseBodyEvent) GoString() string {
	return s.String()
}

func (s *DescribeFreeUserEventsResponseBodyEvent) GetApiFormat() *string {
	return s.ApiFormat
}

func (s *DescribeFreeUserEventsResponseBodyEvent) GetAttackIP() *string {
	return s.AttackIP
}

func (s *DescribeFreeUserEventsResponseBodyEvent) GetAttackTime() *int64 {
	return s.AttackTime
}

func (s *DescribeFreeUserEventsResponseBodyEvent) GetDomain() *string {
	return s.Domain
}

func (s *DescribeFreeUserEventsResponseBodyEvent) GetEventLevel() *string {
	return s.EventLevel
}

func (s *DescribeFreeUserEventsResponseBodyEvent) GetEventTag() *string {
	return s.EventTag
}

func (s *DescribeFreeUserEventsResponseBodyEvent) SetApiFormat(v string) *DescribeFreeUserEventsResponseBodyEvent {
	s.ApiFormat = &v
	return s
}

func (s *DescribeFreeUserEventsResponseBodyEvent) SetAttackIP(v string) *DescribeFreeUserEventsResponseBodyEvent {
	s.AttackIP = &v
	return s
}

func (s *DescribeFreeUserEventsResponseBodyEvent) SetAttackTime(v int64) *DescribeFreeUserEventsResponseBodyEvent {
	s.AttackTime = &v
	return s
}

func (s *DescribeFreeUserEventsResponseBodyEvent) SetDomain(v string) *DescribeFreeUserEventsResponseBodyEvent {
	s.Domain = &v
	return s
}

func (s *DescribeFreeUserEventsResponseBodyEvent) SetEventLevel(v string) *DescribeFreeUserEventsResponseBodyEvent {
	s.EventLevel = &v
	return s
}

func (s *DescribeFreeUserEventsResponseBodyEvent) SetEventTag(v string) *DescribeFreeUserEventsResponseBodyEvent {
	s.EventTag = &v
	return s
}

func (s *DescribeFreeUserEventsResponseBodyEvent) Validate() error {
	return dara.Validate(s)
}
