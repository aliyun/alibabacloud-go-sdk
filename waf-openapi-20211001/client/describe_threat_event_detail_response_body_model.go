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
	// The ID of the request.
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
	// The most recent attack time. The value is a UNIX timestamp in UTC. Unit: seconds.
	//
	// example:
	//
	// 1749916800000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The number of blocked attacks.
	//
	// example:
	//
	// 20
	EventBlock *string `json:"EventBlock,omitempty" xml:"EventBlock,omitempty"`
	// The total number of attacks.
	//
	// example:
	//
	// 20
	EventCnt *string `json:"EventCnt,omitempty" xml:"EventCnt,omitempty"`
	// The filter condition for the log query feature. The value is a string converted from a **JSON*	- structure constructed with a series of parameters.
	//
	// example:
	//
	// {"end_ts": 1766637714, "start_ts": 1764096746, "condition": {"real_client_ip": ["78.153.140.179", "78.153.140.203", "78.153.140.177", "78.153.140.178", "78.153.140.151"]}}
	EventCondition *string `json:"EventCondition,omitempty" xml:"EventCondition,omitempty"`
	// The threat intelligence associated with the event. The value is a string converted from a **JSON*	- array.
	//
	// example:
	//
	// ["CVE-2020-14882","DDoS Attack"]
	EventIntelligence *string `json:"EventIntelligence,omitempty" xml:"EventIntelligence,omitempty"`
	// The severity level of the event. Valid values:
	//
	// - **critical**: critical.
	//
	// - **high**: high.
	//
	// - **medium**: medium.
	//
	// - **low**: low.
	//
	// example:
	//
	// high
	EventLevel *string `json:"EventLevel,omitempty" xml:"EventLevel,omitempty"`
	// The attack source IP addresses.
	//
	// > A security event may have multiple attack source IP addresses. This operation returns only one.
	//
	// example:
	//
	// XX.XX.XX.XX
	EventSrc *string `json:"EventSrc,omitempty" xml:"EventSrc,omitempty"`
	// The country to which the attack source IP addresses belong.
	//
	// example:
	//
	// GB
	EventSrcCountry *string `json:"EventSrcCountry,omitempty" xml:"EventSrcCountry,omitempty"`
	// The region to which the attack source IP addresses belong.
	//
	// example:
	//
	// GB-ENG
	EventSrcRegion *string `json:"EventSrcRegion,omitempty" xml:"EventSrcRegion,omitempty"`
	// The security suggestion. Valid values:
	//
	// - **ProtectInterface**: The Attack Target is suspected to be a backend management address. If the address has fixed access characteristics, configure custom rules in the access control module to restrict access.
	//
	// - **BlockArea**: Monitor the source region of the attack. If the attack source differs from the normal regional distribution of your service, configure Location Blacklist or IP blacklist rules in the access control module to restrict access.
	//
	// - **SwitchBlock**: The current rule protection is in alerting mode. To ensure service security, switch to block mode. Before switching, check for false positives.
	//
	// - **FixBug**: Check whether the Attack Target has a security vulnerability. If the security vulnerability exists, fix it promptly to prevent exploitation by attackers.
	//
	// - **SwitchStrict**: If normal service is not affected, adjust the policy of modules such as rule protection and scan protection to a stricter pattern. Before adjusting the policy, check for false positives.
	//
	// - **ProtectFile**: Check whether sensitive files or paths exist under the target domain name to prevent detection and exploitation by attackers.
	//
	// - **BlockIP**: The source IP address of the attack is highly malicious. Monitor the IP address. If normal service is not affected, block access from the malicious IP address by using an IP blacklist.
	//
	// - **KeepConcerned**: No risks are detected. Continue to monitor the event.
	//
	// example:
	//
	// FixBug
	EventSuggest *string `json:"EventSuggest,omitempty" xml:"EventSuggest,omitempty"`
	// The name of the event. Valid values:
	//
	// - **MultipleDomainDirscan**: directory and file scanning targeting multiple domain names
	//
	// - **SingleDomainDirscan**: directory and file scanning targeting a single domain name
	//
	// - **MultipleDomainWebscan**: web vulnerability scanning targeting multiple domain names
	//
	// - **SingleDomainWebscan**: web vulnerability scanning targeting a single domain name
	//
	// - **MultipleDomainWebattack**: web vulnerability attacks targeting multiple domain names
	//
	// - **SingleDomainWebattack**: web vulnerability attacks targeting a single domain name
	//
	// - **SingleURLWebattack**: web vulnerability attacks targeting a specific URL
	//
	// - **SingleURLSqlattack**: SQL injection attacks targeting a specific URL
	//
	// - **SingleURLXssattack**: XSS attacks targeting a specific URL
	//
	// - **WebshellUpload**: attack behavior attempting to upload backdoor trojans
	//
	// - **RandomVulnTest**: random web vulnerability detection behavior
	//
	// example:
	//
	// MultipleDomainWebattack
	EventTag *string `json:"EventTag,omitempty" xml:"EventTag,omitempty"`
	// Indicates whether the attack event is a targeted persistent attack event. Valid values:
	//
	// - **0**: non-persistent attack.
	//
	// - **1**: persistent attack.
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
