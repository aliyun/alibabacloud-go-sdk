// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeThreatEventResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeThreatEventResponseBody
	GetRequestId() *string
	SetThreatEvents(v []*DescribeThreatEventResponseBodyThreatEvents) *DescribeThreatEventResponseBody
	GetThreatEvents() []*DescribeThreatEventResponseBodyThreatEvents
	SetTotalCount(v int64) *DescribeThreatEventResponseBody
	GetTotalCount() *int64
}

type DescribeThreatEventResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 98106632-6865-5600-A834-3D909***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of notable security events.
	ThreatEvents []*DescribeThreatEventResponseBodyThreatEvents `json:"ThreatEvents,omitempty" xml:"ThreatEvents,omitempty" type:"Repeated"`
	// The total number of security events that match the query conditions.
	//
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeThreatEventResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeThreatEventResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeThreatEventResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeThreatEventResponseBody) GetThreatEvents() []*DescribeThreatEventResponseBodyThreatEvents {
	return s.ThreatEvents
}

func (s *DescribeThreatEventResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeThreatEventResponseBody) SetRequestId(v string) *DescribeThreatEventResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeThreatEventResponseBody) SetThreatEvents(v []*DescribeThreatEventResponseBodyThreatEvents) *DescribeThreatEventResponseBody {
	s.ThreatEvents = v
	return s
}

func (s *DescribeThreatEventResponseBody) SetTotalCount(v int64) *DescribeThreatEventResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeThreatEventResponseBody) Validate() error {
	if s.ThreatEvents != nil {
		for _, item := range s.ThreatEvents {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeThreatEventResponseBodyThreatEvents struct {
	// The percentage of attack requests that were blocked by WAF.
	//
	// example:
	//
	// 100%
	BlockRate *string `json:"BlockRate,omitempty" xml:"BlockRate,omitempty"`
	// The time when the last attack occurred. This value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1768406400000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the security event.
	//
	// example:
	//
	// f439994c8ab39f84eced33490f0c4388
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// The severity level of the security event. Valid values:
	//
	// - **critical**
	//
	// - **high**
	//
	// - **medium**
	//
	// - **low**
	//
	// example:
	//
	// high
	EventLevel *string `json:"EventLevel,omitempty" xml:"EventLevel,omitempty"`
	// The source IP address of the attack.
	//
	// > A security event may have multiple source IP addresses. This operation returns only one of them.
	//
	// example:
	//
	// 3.3.3.3
	EventSrc *string `json:"EventSrc,omitempty" xml:"EventSrc,omitempty"`
	// The type of the security event. Valid values:
	//
	// - **MultipleDomainDirscan**: a directory and file scan against multiple domain names.
	//
	// - **SingleDomainDirscan**: a directory and file scan against a single domain name.
	//
	// - **MultipleDomainWebscan**: a web vulnerability scan against multiple domain names.
	//
	// - **SingleDomainWebscan**: a web vulnerability scan against a single domain name.
	//
	// - **MultipleDomainWebattack**: a web vulnerability attack against multiple domain names.
	//
	// - **SingleDomainWebattack**: a web vulnerability attack against a single domain name.
	//
	// - **SingleURLWebattack**: a web vulnerability attack against a specific URL.
	//
	// - **SingleURLSqlattack**: an SQL injection attack against a specific URL.
	//
	// - **SingleURLXssattack**: an XSS attack against a specific URL.
	//
	// - **WebshellUpload**: an attack that attempts to upload backdoor trojans.
	//
	// - **RandomVulnTest**: a random web vulnerability probe.
	//
	// example:
	//
	// Event_InternalLoginWeakPasswd
	EventTag *string `json:"EventTag,omitempty" xml:"EventTag,omitempty"`
	// The protected object that is the target of the attack.
	//
	// > A security event may have multiple protected objects as targets. This operation returns only one of them.
	//
	// example:
	//
	// test.aliyundemo.com-waf
	EventTarget *string `json:"EventTarget,omitempty" xml:"EventTarget,omitempty"`
}

func (s DescribeThreatEventResponseBodyThreatEvents) String() string {
	return dara.Prettify(s)
}

func (s DescribeThreatEventResponseBodyThreatEvents) GoString() string {
	return s.String()
}

func (s *DescribeThreatEventResponseBodyThreatEvents) GetBlockRate() *string {
	return s.BlockRate
}

func (s *DescribeThreatEventResponseBodyThreatEvents) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeThreatEventResponseBodyThreatEvents) GetEventId() *string {
	return s.EventId
}

func (s *DescribeThreatEventResponseBodyThreatEvents) GetEventLevel() *string {
	return s.EventLevel
}

func (s *DescribeThreatEventResponseBodyThreatEvents) GetEventSrc() *string {
	return s.EventSrc
}

func (s *DescribeThreatEventResponseBodyThreatEvents) GetEventTag() *string {
	return s.EventTag
}

func (s *DescribeThreatEventResponseBodyThreatEvents) GetEventTarget() *string {
	return s.EventTarget
}

func (s *DescribeThreatEventResponseBodyThreatEvents) SetBlockRate(v string) *DescribeThreatEventResponseBodyThreatEvents {
	s.BlockRate = &v
	return s
}

func (s *DescribeThreatEventResponseBodyThreatEvents) SetEndTime(v int64) *DescribeThreatEventResponseBodyThreatEvents {
	s.EndTime = &v
	return s
}

func (s *DescribeThreatEventResponseBodyThreatEvents) SetEventId(v string) *DescribeThreatEventResponseBodyThreatEvents {
	s.EventId = &v
	return s
}

func (s *DescribeThreatEventResponseBodyThreatEvents) SetEventLevel(v string) *DescribeThreatEventResponseBodyThreatEvents {
	s.EventLevel = &v
	return s
}

func (s *DescribeThreatEventResponseBodyThreatEvents) SetEventSrc(v string) *DescribeThreatEventResponseBodyThreatEvents {
	s.EventSrc = &v
	return s
}

func (s *DescribeThreatEventResponseBodyThreatEvents) SetEventTag(v string) *DescribeThreatEventResponseBodyThreatEvents {
	s.EventTag = &v
	return s
}

func (s *DescribeThreatEventResponseBodyThreatEvents) SetEventTarget(v string) *DescribeThreatEventResponseBodyThreatEvents {
	s.EventTarget = &v
	return s
}

func (s *DescribeThreatEventResponseBodyThreatEvents) Validate() error {
	return dara.Validate(s)
}
