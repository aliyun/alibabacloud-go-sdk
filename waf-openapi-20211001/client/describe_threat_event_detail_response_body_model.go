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
	// The request ID.
	//
	// example:
	//
	// D7861F61-5B61-46CE-A47C-6B1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details of the security event.
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
	// The time of the most recent attack. This value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1749916800000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The number of attacks that were blocked in the security event.
	//
	// example:
	//
	// 20
	EventBlock *string `json:"EventBlock,omitempty" xml:"EventBlock,omitempty"`
	// The total number of attacks in the security event.
	//
	// example:
	//
	// 20
	EventCnt *string `json:"EventCnt,omitempty" xml:"EventCnt,omitempty"`
	// The filter condition for viewing logs. The value is a JSON object in the string format.
	//
	// example:
	//
	// {"end_ts": 1766637714, "start_ts": 1764096746, "condition": {"real_client_ip": ["78.153.140.179", "78.153.140.203", "78.153.140.177", "78.153.140.178", "78.153.140.151"]}}
	EventCondition *string `json:"EventCondition,omitempty" xml:"EventCondition,omitempty"`
	// The threat intelligence associated with the event. The value is a JSON array in the string format.
	//
	// example:
	//
	// ["CVE-2020-14882","DDoS Attack"]
	EventIntelligence *string `json:"EventIntelligence,omitempty" xml:"EventIntelligence,omitempty"`
	// The severity level of the event. Valid values:
	//
	// - **critical**: Critical severity.
	//
	// - **high**: High severity.
	//
	// - **medium**: Medium severity.
	//
	// - **low**: Low severity.
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
	// XX.XX.XX.XX
	EventSrc *string `json:"EventSrc,omitempty" xml:"EventSrc,omitempty"`
	// The country of the source IP address of the attack.
	//
	// example:
	//
	// GB
	EventSrcCountry *string `json:"EventSrcCountry,omitempty" xml:"EventSrcCountry,omitempty"`
	// The region of the source IP address of the attack.
	//
	// example:
	//
	// GB-ENG
	EventSrcRegion *string `json:"EventSrcRegion,omitempty" xml:"EventSrcRegion,omitempty"`
	// The security suggestion. Valid values:
	//
	// - **ProtectInterface**: The attack target appears to be a backend management address. If the address has specific access patterns, configure custom rules in the access control module to restrict access.
	//
	// - **BlockArea**: Monitor the attack source region. If the attack source region is different from your normal business regions, configure a location blacklist or an IP address blacklist in the access control module to restrict access.
	//
	// - **SwitchBlock**: The current protection rule is in Alert mode. To ensure business security, switch to Block mode. Before you switch, check for false positives.
	//
	// - **FixBug**: Check the attack target for security vulnerabilities. If any vulnerabilities exist, fix them promptly to prevent exploitation.
	//
	// - **SwitchStrict**: If it does not affect your normal business, change the policies of modules, such as protection rules and scan protection, to a stricter mode. Before you change the policies, check for false positives.
	//
	// - **ProtectFile**: Check the target domain name for sensitive files or paths to prevent them from being detected and exploited.
	//
	// - **BlockIP**: The source IP address has a high degree of maliciousness. Keep monitoring it. If it does not affect your normal business, use an IP address blacklist to block access from the malicious IP address.
	//
	// - **KeepConcerned**: No threats are found. Continue to monitor the situation.
	//
	// example:
	//
	// FixBug
	EventSuggest *string `json:"EventSuggest,omitempty" xml:"EventSuggest,omitempty"`
	// The name of the event. Valid values:
	//
	// - **MultipleDomainDirscan**: A directory and file scan for multiple domain names.
	//
	// - **SingleDomainDirscan**: A directory and file scan for a single domain name.
	//
	// - **MultipleDomainWebscan**: A web vulnerability scan for multiple domain names.
	//
	// - **SingleDomainWebscan**: A web vulnerability scan for a single domain name.
	//
	// - **MultipleDomainWebattack**: A web vulnerability attack on multiple domain names.
	//
	// - **SingleDomainWebattack**: A web vulnerability attack on a single domain name.
	//
	// - **SingleURLWebattack**: A web vulnerability attack on a specific URL.
	//
	// - **SingleURLSqlattack**: An SQL injection attack on a specific URL.
	//
	// - **SingleURLXssattack**: A cross-site scripting (XSS) attack on a specific URL.
	//
	// - **WebshellUpload**: An attempt to upload backdoor trojans.
	//
	// - **RandomVulnTest**: A random web vulnerability probe.
	//
	// example:
	//
	// MultipleDomainWebattack
	EventTag *string `json:"EventTag,omitempty" xml:"EventTag,omitempty"`
	// Indicates whether the event is a persistent attack. Valid values:
	//
	// - **0**: The event is not a persistent attack.
	//
	// - **1**: The event is a persistent attack.
	//
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
