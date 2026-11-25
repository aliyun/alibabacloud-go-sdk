// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAiEngineEventResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAttackSources(v []*DescribeAiEngineEventResponseBodyAttackSources) *DescribeAiEngineEventResponseBody
	GetAttackSources() []*DescribeAiEngineEventResponseBodyAttackSources
	SetBypassRequestCount(v int64) *DescribeAiEngineEventResponseBody
	GetBypassRequestCount() *int64
	SetDomains(v []*string) *DescribeAiEngineEventResponseBody
	GetDomains() []*string
	SetEventId(v string) *DescribeAiEngineEventResponseBody
	GetEventId() *string
	SetEventName(v string) *DescribeAiEngineEventResponseBody
	GetEventName() *string
	SetEventNodeCount(v int32) *DescribeAiEngineEventResponseBody
	GetEventNodeCount() *int32
	SetEventSummary(v string) *DescribeAiEngineEventResponseBody
	GetEventSummary() *string
	SetEvidenceRequestCount(v int64) *DescribeAiEngineEventResponseBody
	GetEvidenceRequestCount() *int64
	SetFirstSeenAt(v int64) *DescribeAiEngineEventResponseBody
	GetFirstSeenAt() *int64
	SetHandleStatus(v string) *DescribeAiEngineEventResponseBody
	GetHandleStatus() *string
	SetLastSeenAt(v int64) *DescribeAiEngineEventResponseBody
	GetLastSeenAt() *int64
	SetRelatedEventIds(v string) *DescribeAiEngineEventResponseBody
	GetRelatedEventIds() *string
	SetReportMarkdown(v string) *DescribeAiEngineEventResponseBody
	GetReportMarkdown() *string
	SetRequestId(v string) *DescribeAiEngineEventResponseBody
	GetRequestId() *string
	SetResources(v []*string) *DescribeAiEngineEventResponseBody
	GetResources() []*string
	SetRiskLevel(v string) *DescribeAiEngineEventResponseBody
	GetRiskLevel() *string
}

type DescribeAiEngineEventResponseBody struct {
	// The grouped list of attack sources for the event.
	AttackSources []*DescribeAiEngineEventResponseBodyAttackSources `json:"AttackSources,omitempty" xml:"AttackSources,omitempty" type:"Repeated"`
	// The statistical value of suspected bypass evidence requests for the event.
	//
	// example:
	//
	// 1
	BypassRequestCount *int64 `json:"BypassRequestCount,omitempty" xml:"BypassRequestCount,omitempty"`
	// The domain names involved in the event.
	Domains []*string `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Repeated"`
	// The event ID.
	//
	// example:
	//
	// EVT-0bda6****6b5cwwz
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// The name of the event in the specified language.
	//
	// example:
	//
	// Suspected injection attack
	EventName *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	// The number of nodes included in the event.
	//
	// example:
	//
	// 1
	EventNodeCount *int32 `json:"EventNodeCount,omitempty" xml:"EventNodeCount,omitempty"`
	// The summary of the event in the specified language.
	//
	// example:
	//
	// Requests suspected of injection attacks were detected.
	EventSummary *string `json:"EventSummary,omitempty" xml:"EventSummary,omitempty"`
	// The statistical value of evidence requests for the event.
	//
	// example:
	//
	// 200
	EvidenceRequestCount *int64 `json:"EvidenceRequestCount,omitempty" xml:"EvidenceRequestCount,omitempty"`
	// The time when the event was first detected. The time is a UNIX timestamp in UTC. Unit: milliseconds.
	//
	// example:
	//
	// 1792973400000
	FirstSeenAt *int64 `json:"FirstSeenAt,omitempty" xml:"FirstSeenAt,omitempty"`
	// The current manual handling status. Valid values:
	//
	// - **unhandled**: Unhandled.
	//
	// - **handling**: Handling.
	//
	// - **handled**: Handled.
	//
	// example:
	//
	// unhandled
	HandleStatus *string `json:"HandleStatus,omitempty" xml:"HandleStatus,omitempty"`
	// The time when the event was last detected. The time is a UNIX timestamp in UTC. Unit: milliseconds.
	//
	// example:
	//
	// 1792973700000
	LastSeenAt *int64 `json:"LastSeenAt,omitempty" xml:"LastSeenAt,omitempty"`
	// The IDs of the related events. Multiple IDs are separated by commas (,).
	//
	// example:
	//
	// EVT-0bda6****6b5c174,EVT-0bda6****6b5c175
	RelatedEventIds *string `json:"RelatedEventIds,omitempty" xml:"RelatedEventIds,omitempty"`
	// The event analysis report in Markdown format in the specified language.
	//
	// example:
	//
	// # Event analysis report
	//
	// Requests suspected of injection attacks were detected.
	ReportMarkdown *string `json:"ReportMarkdown,omitempty" xml:"ReportMarkdown,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// request-demo
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The IDs of the protected objects involved in the event.
	Resources []*string `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
	// The current risk level. Valid values:
	//
	// - **critical**: Critical.
	//
	// - **high**: High.
	//
	// - **medium**: Medium.
	//
	// - **low**: Low.
	//
	// example:
	//
	// high
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
}

func (s DescribeAiEngineEventResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAiEngineEventResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAiEngineEventResponseBody) GetAttackSources() []*DescribeAiEngineEventResponseBodyAttackSources {
	return s.AttackSources
}

func (s *DescribeAiEngineEventResponseBody) GetBypassRequestCount() *int64 {
	return s.BypassRequestCount
}

func (s *DescribeAiEngineEventResponseBody) GetDomains() []*string {
	return s.Domains
}

func (s *DescribeAiEngineEventResponseBody) GetEventId() *string {
	return s.EventId
}

func (s *DescribeAiEngineEventResponseBody) GetEventName() *string {
	return s.EventName
}

func (s *DescribeAiEngineEventResponseBody) GetEventNodeCount() *int32 {
	return s.EventNodeCount
}

func (s *DescribeAiEngineEventResponseBody) GetEventSummary() *string {
	return s.EventSummary
}

func (s *DescribeAiEngineEventResponseBody) GetEvidenceRequestCount() *int64 {
	return s.EvidenceRequestCount
}

func (s *DescribeAiEngineEventResponseBody) GetFirstSeenAt() *int64 {
	return s.FirstSeenAt
}

func (s *DescribeAiEngineEventResponseBody) GetHandleStatus() *string {
	return s.HandleStatus
}

func (s *DescribeAiEngineEventResponseBody) GetLastSeenAt() *int64 {
	return s.LastSeenAt
}

func (s *DescribeAiEngineEventResponseBody) GetRelatedEventIds() *string {
	return s.RelatedEventIds
}

func (s *DescribeAiEngineEventResponseBody) GetReportMarkdown() *string {
	return s.ReportMarkdown
}

func (s *DescribeAiEngineEventResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAiEngineEventResponseBody) GetResources() []*string {
	return s.Resources
}

func (s *DescribeAiEngineEventResponseBody) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *DescribeAiEngineEventResponseBody) SetAttackSources(v []*DescribeAiEngineEventResponseBodyAttackSources) *DescribeAiEngineEventResponseBody {
	s.AttackSources = v
	return s
}

func (s *DescribeAiEngineEventResponseBody) SetBypassRequestCount(v int64) *DescribeAiEngineEventResponseBody {
	s.BypassRequestCount = &v
	return s
}

func (s *DescribeAiEngineEventResponseBody) SetDomains(v []*string) *DescribeAiEngineEventResponseBody {
	s.Domains = v
	return s
}

func (s *DescribeAiEngineEventResponseBody) SetEventId(v string) *DescribeAiEngineEventResponseBody {
	s.EventId = &v
	return s
}

func (s *DescribeAiEngineEventResponseBody) SetEventName(v string) *DescribeAiEngineEventResponseBody {
	s.EventName = &v
	return s
}

func (s *DescribeAiEngineEventResponseBody) SetEventNodeCount(v int32) *DescribeAiEngineEventResponseBody {
	s.EventNodeCount = &v
	return s
}

func (s *DescribeAiEngineEventResponseBody) SetEventSummary(v string) *DescribeAiEngineEventResponseBody {
	s.EventSummary = &v
	return s
}

func (s *DescribeAiEngineEventResponseBody) SetEvidenceRequestCount(v int64) *DescribeAiEngineEventResponseBody {
	s.EvidenceRequestCount = &v
	return s
}

func (s *DescribeAiEngineEventResponseBody) SetFirstSeenAt(v int64) *DescribeAiEngineEventResponseBody {
	s.FirstSeenAt = &v
	return s
}

func (s *DescribeAiEngineEventResponseBody) SetHandleStatus(v string) *DescribeAiEngineEventResponseBody {
	s.HandleStatus = &v
	return s
}

func (s *DescribeAiEngineEventResponseBody) SetLastSeenAt(v int64) *DescribeAiEngineEventResponseBody {
	s.LastSeenAt = &v
	return s
}

func (s *DescribeAiEngineEventResponseBody) SetRelatedEventIds(v string) *DescribeAiEngineEventResponseBody {
	s.RelatedEventIds = &v
	return s
}

func (s *DescribeAiEngineEventResponseBody) SetReportMarkdown(v string) *DescribeAiEngineEventResponseBody {
	s.ReportMarkdown = &v
	return s
}

func (s *DescribeAiEngineEventResponseBody) SetRequestId(v string) *DescribeAiEngineEventResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAiEngineEventResponseBody) SetResources(v []*string) *DescribeAiEngineEventResponseBody {
	s.Resources = v
	return s
}

func (s *DescribeAiEngineEventResponseBody) SetRiskLevel(v string) *DescribeAiEngineEventResponseBody {
	s.RiskLevel = &v
	return s
}

func (s *DescribeAiEngineEventResponseBody) Validate() error {
	if s.AttackSources != nil {
		for _, item := range s.AttackSources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAiEngineEventResponseBodyAttackSources struct {
	// The list of sources of this type.
	Sources []*DescribeAiEngineEventResponseBodyAttackSourcesSources `json:"Sources,omitempty" xml:"Sources,omitempty" type:"Repeated"`
	// The type of the attack source. Valid values:
	//
	// - **ip**: IP address.
	//
	// example:
	//
	// ip
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeAiEngineEventResponseBodyAttackSources) String() string {
	return dara.Prettify(s)
}

func (s DescribeAiEngineEventResponseBodyAttackSources) GoString() string {
	return s.String()
}

func (s *DescribeAiEngineEventResponseBodyAttackSources) GetSources() []*DescribeAiEngineEventResponseBodyAttackSourcesSources {
	return s.Sources
}

func (s *DescribeAiEngineEventResponseBodyAttackSources) GetType() *string {
	return s.Type
}

func (s *DescribeAiEngineEventResponseBodyAttackSources) SetSources(v []*DescribeAiEngineEventResponseBodyAttackSourcesSources) *DescribeAiEngineEventResponseBodyAttackSources {
	s.Sources = v
	return s
}

func (s *DescribeAiEngineEventResponseBodyAttackSources) SetType(v string) *DescribeAiEngineEventResponseBodyAttackSources {
	s.Type = &v
	return s
}

func (s *DescribeAiEngineEventResponseBodyAttackSources) Validate() error {
	if s.Sources != nil {
		for _, item := range s.Sources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAiEngineEventResponseBodyAttackSourcesSources struct {
	// The display attribute of the source, such as the region code.
	//
	// example:
	//
	// CN-510100
	Attribute *string `json:"Attribute,omitempty" xml:"Attribute,omitempty"`
	// The value of the attack source. If the attack source type is IP address (**Type*	- is set to **ip**), this parameter indicates the IP address.
	//
	// example:
	//
	// 203.0.113.109
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeAiEngineEventResponseBodyAttackSourcesSources) String() string {
	return dara.Prettify(s)
}

func (s DescribeAiEngineEventResponseBodyAttackSourcesSources) GoString() string {
	return s.String()
}

func (s *DescribeAiEngineEventResponseBodyAttackSourcesSources) GetAttribute() *string {
	return s.Attribute
}

func (s *DescribeAiEngineEventResponseBodyAttackSourcesSources) GetValue() *string {
	return s.Value
}

func (s *DescribeAiEngineEventResponseBodyAttackSourcesSources) SetAttribute(v string) *DescribeAiEngineEventResponseBodyAttackSourcesSources {
	s.Attribute = &v
	return s
}

func (s *DescribeAiEngineEventResponseBodyAttackSourcesSources) SetValue(v string) *DescribeAiEngineEventResponseBodyAttackSourcesSources {
	s.Value = &v
	return s
}

func (s *DescribeAiEngineEventResponseBodyAttackSourcesSources) Validate() error {
	return dara.Validate(s)
}
