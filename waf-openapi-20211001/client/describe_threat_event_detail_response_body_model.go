// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeThreatEventDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeThreatEventDetailResponseBody
	GetRequestId() *string
	SetThreatEventDetail(v *DescribeThreatEventDetailResponseBodyThreatEventDetail) *DescribeThreatEventDetailResponseBody
	GetThreatEventDetail() *DescribeThreatEventDetailResponseBodyThreatEventDetail
}

type DescribeThreatEventDetailResponseBody struct {
	// example:
	//
	// D7861F61-5B61-46CE-A47C-6B1****
	RequestId         *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ThreatEventDetail *DescribeThreatEventDetailResponseBodyThreatEventDetail `json:"ThreatEventDetail,omitempty" xml:"ThreatEventDetail,omitempty" type:"Struct"`
}

func (s DescribeThreatEventDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeThreatEventDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeThreatEventDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeThreatEventDetailResponseBody) GetThreatEventDetail() *DescribeThreatEventDetailResponseBodyThreatEventDetail {
	return s.ThreatEventDetail
}

func (s *DescribeThreatEventDetailResponseBody) SetRequestId(v string) *DescribeThreatEventDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeThreatEventDetailResponseBody) SetThreatEventDetail(v *DescribeThreatEventDetailResponseBodyThreatEventDetail) *DescribeThreatEventDetailResponseBody {
	s.ThreatEventDetail = v
	return s
}

func (s *DescribeThreatEventDetailResponseBody) Validate() error {
	if s.ThreatEventDetail != nil {
		if err := s.ThreatEventDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeThreatEventDetailResponseBodyThreatEventDetail struct {
	// example:
	//
	// 1749916800000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 20
	EventBlock *string `json:"EventBlock,omitempty" xml:"EventBlock,omitempty"`
	// example:
	//
	// 20
	EventCnt *string `json:"EventCnt,omitempty" xml:"EventCnt,omitempty"`
	// example:
	//
	// {"end_ts": 1766637714, "start_ts": 1764096746, "condition": {"real_client_ip": ["78.153.140.179", "78.153.140.203", "78.153.140.177", "78.153.140.178", "78.153.140.151"]}}
	EventCondition *string `json:"EventCondition,omitempty" xml:"EventCondition,omitempty"`
	// example:
	//
	// ["CVE-2020-14882","DDoS Attack"]
	EventIntelligence *string `json:"EventIntelligence,omitempty" xml:"EventIntelligence,omitempty"`
	// example:
	//
	// high
	EventLevel *string `json:"EventLevel,omitempty" xml:"EventLevel,omitempty"`
	// example:
	//
	// XX.XX.XX.XX
	EventSrc *string `json:"EventSrc,omitempty" xml:"EventSrc,omitempty"`
	// example:
	//
	// GB
	EventSrcCountry *string `json:"EventSrcCountry,omitempty" xml:"EventSrcCountry,omitempty"`
	// example:
	//
	// GB-ENG
	EventSrcRegion *string `json:"EventSrcRegion,omitempty" xml:"EventSrcRegion,omitempty"`
	// example:
	//
	// FixBug
	EventSuggest *string `json:"EventSuggest,omitempty" xml:"EventSuggest,omitempty"`
	// example:
	//
	// MultipleDomainWebattack
	EventTag *string `json:"EventTag,omitempty" xml:"EventTag,omitempty"`
	// example:
	//
	// 1
	IsPersistent *int64 `json:"IsPersistent,omitempty" xml:"IsPersistent,omitempty"`
}

func (s DescribeThreatEventDetailResponseBodyThreatEventDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeThreatEventDetailResponseBodyThreatEventDetail) GoString() string {
	return s.String()
}

func (s *DescribeThreatEventDetailResponseBodyThreatEventDetail) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeThreatEventDetailResponseBodyThreatEventDetail) GetEventBlock() *string {
	return s.EventBlock
}

func (s *DescribeThreatEventDetailResponseBodyThreatEventDetail) GetEventCnt() *string {
	return s.EventCnt
}

func (s *DescribeThreatEventDetailResponseBodyThreatEventDetail) GetEventCondition() *string {
	return s.EventCondition
}

func (s *DescribeThreatEventDetailResponseBodyThreatEventDetail) GetEventIntelligence() *string {
	return s.EventIntelligence
}

func (s *DescribeThreatEventDetailResponseBodyThreatEventDetail) GetEventLevel() *string {
	return s.EventLevel
}

func (s *DescribeThreatEventDetailResponseBodyThreatEventDetail) GetEventSrc() *string {
	return s.EventSrc
}

func (s *DescribeThreatEventDetailResponseBodyThreatEventDetail) GetEventSrcCountry() *string {
	return s.EventSrcCountry
}

func (s *DescribeThreatEventDetailResponseBodyThreatEventDetail) GetEventSrcRegion() *string {
	return s.EventSrcRegion
}

func (s *DescribeThreatEventDetailResponseBodyThreatEventDetail) GetEventSuggest() *string {
	return s.EventSuggest
}

func (s *DescribeThreatEventDetailResponseBodyThreatEventDetail) GetEventTag() *string {
	return s.EventTag
}

func (s *DescribeThreatEventDetailResponseBodyThreatEventDetail) GetIsPersistent() *int64 {
	return s.IsPersistent
}

func (s *DescribeThreatEventDetailResponseBodyThreatEventDetail) SetEndTime(v int64) *DescribeThreatEventDetailResponseBodyThreatEventDetail {
	s.EndTime = &v
	return s
}

func (s *DescribeThreatEventDetailResponseBodyThreatEventDetail) SetEventBlock(v string) *DescribeThreatEventDetailResponseBodyThreatEventDetail {
	s.EventBlock = &v
	return s
}

func (s *DescribeThreatEventDetailResponseBodyThreatEventDetail) SetEventCnt(v string) *DescribeThreatEventDetailResponseBodyThreatEventDetail {
	s.EventCnt = &v
	return s
}

func (s *DescribeThreatEventDetailResponseBodyThreatEventDetail) SetEventCondition(v string) *DescribeThreatEventDetailResponseBodyThreatEventDetail {
	s.EventCondition = &v
	return s
}

func (s *DescribeThreatEventDetailResponseBodyThreatEventDetail) SetEventIntelligence(v string) *DescribeThreatEventDetailResponseBodyThreatEventDetail {
	s.EventIntelligence = &v
	return s
}

func (s *DescribeThreatEventDetailResponseBodyThreatEventDetail) SetEventLevel(v string) *DescribeThreatEventDetailResponseBodyThreatEventDetail {
	s.EventLevel = &v
	return s
}

func (s *DescribeThreatEventDetailResponseBodyThreatEventDetail) SetEventSrc(v string) *DescribeThreatEventDetailResponseBodyThreatEventDetail {
	s.EventSrc = &v
	return s
}

func (s *DescribeThreatEventDetailResponseBodyThreatEventDetail) SetEventSrcCountry(v string) *DescribeThreatEventDetailResponseBodyThreatEventDetail {
	s.EventSrcCountry = &v
	return s
}

func (s *DescribeThreatEventDetailResponseBodyThreatEventDetail) SetEventSrcRegion(v string) *DescribeThreatEventDetailResponseBodyThreatEventDetail {
	s.EventSrcRegion = &v
	return s
}

func (s *DescribeThreatEventDetailResponseBodyThreatEventDetail) SetEventSuggest(v string) *DescribeThreatEventDetailResponseBodyThreatEventDetail {
	s.EventSuggest = &v
	return s
}

func (s *DescribeThreatEventDetailResponseBodyThreatEventDetail) SetEventTag(v string) *DescribeThreatEventDetailResponseBodyThreatEventDetail {
	s.EventTag = &v
	return s
}

func (s *DescribeThreatEventDetailResponseBodyThreatEventDetail) SetIsPersistent(v int64) *DescribeThreatEventDetailResponseBodyThreatEventDetail {
	s.IsPersistent = &v
	return s
}

func (s *DescribeThreatEventDetailResponseBodyThreatEventDetail) Validate() error {
	return dara.Validate(s)
}
