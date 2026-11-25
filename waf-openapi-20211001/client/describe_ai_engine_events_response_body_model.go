// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAiEngineEventsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*DescribeAiEngineEventsResponseBodyItems) *DescribeAiEngineEventsResponseBody
	GetItems() []*DescribeAiEngineEventsResponseBodyItems
	SetMaxResults(v int32) *DescribeAiEngineEventsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeAiEngineEventsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeAiEngineEventsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeAiEngineEventsResponseBody
	GetTotalCount() *int64
}

type DescribeAiEngineEventsResponseBody struct {
	// The list of events.
	Items []*DescribeAiEngineEventsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The number of entries to return per page in a paging query. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token for the next page. If a next page exists, this field returns a value.
	//
	// > If this parameter returns a value, a next page exists. You can use the returned **NextToken*	- as a request parameter to retrieve the next page until no value is returned, which indicates that all data has been retrieved.
	//
	// example:
	//
	// token-demo
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// request-demo
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of events that meet the query conditions.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAiEngineEventsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAiEngineEventsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAiEngineEventsResponseBody) GetItems() []*DescribeAiEngineEventsResponseBodyItems {
	return s.Items
}

func (s *DescribeAiEngineEventsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeAiEngineEventsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeAiEngineEventsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAiEngineEventsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeAiEngineEventsResponseBody) SetItems(v []*DescribeAiEngineEventsResponseBodyItems) *DescribeAiEngineEventsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeAiEngineEventsResponseBody) SetMaxResults(v int32) *DescribeAiEngineEventsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeAiEngineEventsResponseBody) SetNextToken(v string) *DescribeAiEngineEventsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeAiEngineEventsResponseBody) SetRequestId(v string) *DescribeAiEngineEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAiEngineEventsResponseBody) SetTotalCount(v int64) *DescribeAiEngineEventsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeAiEngineEventsResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAiEngineEventsResponseBodyItems struct {
	// The list of attack source groups for the event.
	AttackSources []*DescribeAiEngineEventsResponseBodyItemsAttackSources `json:"AttackSources,omitempty" xml:"AttackSources,omitempty" type:"Repeated"`
	// The statistics of suspected bypass evidence requests for the event.
	//
	// example:
	//
	// 1
	BypassRequestCount *int64 `json:"BypassRequestCount,omitempty" xml:"BypassRequestCount,omitempty"`
	// The list of domain names involved in the event.
	Domains []*string `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Repeated"`
	// The ID of the event.
	//
	// example:
	//
	// EVT-20261025-109
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// The name of the event in the selected language.
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
	// The statistics of evidence requests for the event.
	//
	// example:
	//
	// 200
	EvidenceRequestCount *int64 `json:"EvidenceRequestCount,omitempty" xml:"EvidenceRequestCount,omitempty"`
	// The time when the event was first detected. The value is a UNIX timestamp in UTC. Unit: milliseconds.
	//
	// example:
	//
	// 1792973400000
	FirstSeenAt *int64 `json:"FirstSeenAt,omitempty" xml:"FirstSeenAt,omitempty"`
	// The current manual handling status. Valid values:
	//
	// - **unhandled**: Unhandled.
	//
	// - **handling**: In progress.
	//
	// - **handled**: Handled.
	//
	// example:
	//
	// unhandled
	HandleStatus *string `json:"HandleStatus,omitempty" xml:"HandleStatus,omitempty"`
	// The time when the event was last detected. The value is a UNIX timestamp in UTC. Unit: milliseconds.
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
	// The list of protected objects involved in the event.
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

func (s DescribeAiEngineEventsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeAiEngineEventsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeAiEngineEventsResponseBodyItems) GetAttackSources() []*DescribeAiEngineEventsResponseBodyItemsAttackSources {
	return s.AttackSources
}

func (s *DescribeAiEngineEventsResponseBodyItems) GetBypassRequestCount() *int64 {
	return s.BypassRequestCount
}

func (s *DescribeAiEngineEventsResponseBodyItems) GetDomains() []*string {
	return s.Domains
}

func (s *DescribeAiEngineEventsResponseBodyItems) GetEventId() *string {
	return s.EventId
}

func (s *DescribeAiEngineEventsResponseBodyItems) GetEventName() *string {
	return s.EventName
}

func (s *DescribeAiEngineEventsResponseBodyItems) GetEventNodeCount() *int32 {
	return s.EventNodeCount
}

func (s *DescribeAiEngineEventsResponseBodyItems) GetEvidenceRequestCount() *int64 {
	return s.EvidenceRequestCount
}

func (s *DescribeAiEngineEventsResponseBodyItems) GetFirstSeenAt() *int64 {
	return s.FirstSeenAt
}

func (s *DescribeAiEngineEventsResponseBodyItems) GetHandleStatus() *string {
	return s.HandleStatus
}

func (s *DescribeAiEngineEventsResponseBodyItems) GetLastSeenAt() *int64 {
	return s.LastSeenAt
}

func (s *DescribeAiEngineEventsResponseBodyItems) GetRelatedEventIds() *string {
	return s.RelatedEventIds
}

func (s *DescribeAiEngineEventsResponseBodyItems) GetResources() []*string {
	return s.Resources
}

func (s *DescribeAiEngineEventsResponseBodyItems) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *DescribeAiEngineEventsResponseBodyItems) SetAttackSources(v []*DescribeAiEngineEventsResponseBodyItemsAttackSources) *DescribeAiEngineEventsResponseBodyItems {
	s.AttackSources = v
	return s
}

func (s *DescribeAiEngineEventsResponseBodyItems) SetBypassRequestCount(v int64) *DescribeAiEngineEventsResponseBodyItems {
	s.BypassRequestCount = &v
	return s
}

func (s *DescribeAiEngineEventsResponseBodyItems) SetDomains(v []*string) *DescribeAiEngineEventsResponseBodyItems {
	s.Domains = v
	return s
}

func (s *DescribeAiEngineEventsResponseBodyItems) SetEventId(v string) *DescribeAiEngineEventsResponseBodyItems {
	s.EventId = &v
	return s
}

func (s *DescribeAiEngineEventsResponseBodyItems) SetEventName(v string) *DescribeAiEngineEventsResponseBodyItems {
	s.EventName = &v
	return s
}

func (s *DescribeAiEngineEventsResponseBodyItems) SetEventNodeCount(v int32) *DescribeAiEngineEventsResponseBodyItems {
	s.EventNodeCount = &v
	return s
}

func (s *DescribeAiEngineEventsResponseBodyItems) SetEvidenceRequestCount(v int64) *DescribeAiEngineEventsResponseBodyItems {
	s.EvidenceRequestCount = &v
	return s
}

func (s *DescribeAiEngineEventsResponseBodyItems) SetFirstSeenAt(v int64) *DescribeAiEngineEventsResponseBodyItems {
	s.FirstSeenAt = &v
	return s
}

func (s *DescribeAiEngineEventsResponseBodyItems) SetHandleStatus(v string) *DescribeAiEngineEventsResponseBodyItems {
	s.HandleStatus = &v
	return s
}

func (s *DescribeAiEngineEventsResponseBodyItems) SetLastSeenAt(v int64) *DescribeAiEngineEventsResponseBodyItems {
	s.LastSeenAt = &v
	return s
}

func (s *DescribeAiEngineEventsResponseBodyItems) SetRelatedEventIds(v string) *DescribeAiEngineEventsResponseBodyItems {
	s.RelatedEventIds = &v
	return s
}

func (s *DescribeAiEngineEventsResponseBodyItems) SetResources(v []*string) *DescribeAiEngineEventsResponseBodyItems {
	s.Resources = v
	return s
}

func (s *DescribeAiEngineEventsResponseBodyItems) SetRiskLevel(v string) *DescribeAiEngineEventsResponseBodyItems {
	s.RiskLevel = &v
	return s
}

func (s *DescribeAiEngineEventsResponseBodyItems) Validate() error {
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

type DescribeAiEngineEventsResponseBodyItemsAttackSources struct {
	// The list of attack sources of the specified type.
	Sources []*DescribeAiEngineEventsResponseBodyItemsAttackSourcesSources `json:"Sources,omitempty" xml:"Sources,omitempty" type:"Repeated"`
	// The type of the attack source. Valid values:
	//
	// - **ip**: IP address.
	//
	// example:
	//
	// ip
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeAiEngineEventsResponseBodyItemsAttackSources) String() string {
	return dara.Prettify(s)
}

func (s DescribeAiEngineEventsResponseBodyItemsAttackSources) GoString() string {
	return s.String()
}

func (s *DescribeAiEngineEventsResponseBodyItemsAttackSources) GetSources() []*DescribeAiEngineEventsResponseBodyItemsAttackSourcesSources {
	return s.Sources
}

func (s *DescribeAiEngineEventsResponseBodyItemsAttackSources) GetType() *string {
	return s.Type
}

func (s *DescribeAiEngineEventsResponseBodyItemsAttackSources) SetSources(v []*DescribeAiEngineEventsResponseBodyItemsAttackSourcesSources) *DescribeAiEngineEventsResponseBodyItemsAttackSources {
	s.Sources = v
	return s
}

func (s *DescribeAiEngineEventsResponseBodyItemsAttackSources) SetType(v string) *DescribeAiEngineEventsResponseBodyItemsAttackSources {
	s.Type = &v
	return s
}

func (s *DescribeAiEngineEventsResponseBodyItemsAttackSources) Validate() error {
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

type DescribeAiEngineEventsResponseBodyItemsAttackSourcesSources struct {
	// The display attribute of the attack source, such as the region code.
	//
	// example:
	//
	// CN-510100
	Attribute *string `json:"Attribute,omitempty" xml:"Attribute,omitempty"`
	// The value of the attack source. If the attack source type is IP address (the **Type*	- parameter is set to **ip**), the value is an IP address.
	//
	// example:
	//
	// 203.0.113.109
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeAiEngineEventsResponseBodyItemsAttackSourcesSources) String() string {
	return dara.Prettify(s)
}

func (s DescribeAiEngineEventsResponseBodyItemsAttackSourcesSources) GoString() string {
	return s.String()
}

func (s *DescribeAiEngineEventsResponseBodyItemsAttackSourcesSources) GetAttribute() *string {
	return s.Attribute
}

func (s *DescribeAiEngineEventsResponseBodyItemsAttackSourcesSources) GetValue() *string {
	return s.Value
}

func (s *DescribeAiEngineEventsResponseBodyItemsAttackSourcesSources) SetAttribute(v string) *DescribeAiEngineEventsResponseBodyItemsAttackSourcesSources {
	s.Attribute = &v
	return s
}

func (s *DescribeAiEngineEventsResponseBodyItemsAttackSourcesSources) SetValue(v string) *DescribeAiEngineEventsResponseBodyItemsAttackSourcesSources {
	s.Value = &v
	return s
}

func (s *DescribeAiEngineEventsResponseBodyItemsAttackSourcesSources) Validate() error {
	return dara.Validate(s)
}
