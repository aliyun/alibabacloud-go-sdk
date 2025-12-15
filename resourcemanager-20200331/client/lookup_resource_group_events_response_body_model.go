// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLookupResourceGroupEventsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEvents(v []*LookupResourceGroupEventsResponseBodyEvents) *LookupResourceGroupEventsResponseBody
	GetEvents() []*LookupResourceGroupEventsResponseBodyEvents
	SetNextToken(v string) *LookupResourceGroupEventsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *LookupResourceGroupEventsResponseBody
	GetRequestId() *string
}

type LookupResourceGroupEventsResponseBody struct {
	Events []*LookupResourceGroupEventsResponseBodyEvents `json:"Events,omitempty" xml:"Events,omitempty" type:"Repeated"`
	// example:
	//
	// TGlzdFJlc291cm****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// F7701451-340B-5CB3-AEA7-7D831F7F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s LookupResourceGroupEventsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s LookupResourceGroupEventsResponseBody) GoString() string {
	return s.String()
}

func (s *LookupResourceGroupEventsResponseBody) GetEvents() []*LookupResourceGroupEventsResponseBodyEvents {
	return s.Events
}

func (s *LookupResourceGroupEventsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *LookupResourceGroupEventsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *LookupResourceGroupEventsResponseBody) SetEvents(v []*LookupResourceGroupEventsResponseBodyEvents) *LookupResourceGroupEventsResponseBody {
	s.Events = v
	return s
}

func (s *LookupResourceGroupEventsResponseBody) SetNextToken(v string) *LookupResourceGroupEventsResponseBody {
	s.NextToken = &v
	return s
}

func (s *LookupResourceGroupEventsResponseBody) SetRequestId(v string) *LookupResourceGroupEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *LookupResourceGroupEventsResponseBody) Validate() error {
	if s.Events != nil {
		for _, item := range s.Events {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type LookupResourceGroupEventsResponseBodyEvents struct {
	// example:
	//
	// Add
	ChangeType *string `json:"ChangeType,omitempty" xml:"ChangeType,omitempty"`
	// example:
	//
	// 2025-12-04T18:35:17Z
	EventTime *string `json:"EventTime,omitempty" xml:"EventTime,omitempty"`
	// example:
	//
	// cn-qingdao
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// ProjectA
	ResourceGroupDisplayName *string `json:"ResourceGroupDisplayName,omitempty" xml:"ResourceGroupDisplayName,omitempty"`
	// example:
	//
	// rg-acfm2sohr74****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// i-wz9fpfe64****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// example:
	//
	// instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// ecs
	Service                 *string                                                             `json:"Service,omitempty" xml:"Service,omitempty"`
	SourceResourceGroupInfo *LookupResourceGroupEventsResponseBodyEventsSourceResourceGroupInfo `json:"SourceResourceGroupInfo,omitempty" xml:"SourceResourceGroupInfo,omitempty" type:"Struct"`
	TargetResourceGroupInfo *LookupResourceGroupEventsResponseBodyEventsTargetResourceGroupInfo `json:"TargetResourceGroupInfo,omitempty" xml:"TargetResourceGroupInfo,omitempty" type:"Struct"`
}

func (s LookupResourceGroupEventsResponseBodyEvents) String() string {
	return dara.Prettify(s)
}

func (s LookupResourceGroupEventsResponseBodyEvents) GoString() string {
	return s.String()
}

func (s *LookupResourceGroupEventsResponseBodyEvents) GetChangeType() *string {
	return s.ChangeType
}

func (s *LookupResourceGroupEventsResponseBodyEvents) GetEventTime() *string {
	return s.EventTime
}

func (s *LookupResourceGroupEventsResponseBodyEvents) GetRegionId() *string {
	return s.RegionId
}

func (s *LookupResourceGroupEventsResponseBodyEvents) GetResourceGroupDisplayName() *string {
	return s.ResourceGroupDisplayName
}

func (s *LookupResourceGroupEventsResponseBodyEvents) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *LookupResourceGroupEventsResponseBodyEvents) GetResourceId() *string {
	return s.ResourceId
}

func (s *LookupResourceGroupEventsResponseBodyEvents) GetResourceType() *string {
	return s.ResourceType
}

func (s *LookupResourceGroupEventsResponseBodyEvents) GetService() *string {
	return s.Service
}

func (s *LookupResourceGroupEventsResponseBodyEvents) GetSourceResourceGroupInfo() *LookupResourceGroupEventsResponseBodyEventsSourceResourceGroupInfo {
	return s.SourceResourceGroupInfo
}

func (s *LookupResourceGroupEventsResponseBodyEvents) GetTargetResourceGroupInfo() *LookupResourceGroupEventsResponseBodyEventsTargetResourceGroupInfo {
	return s.TargetResourceGroupInfo
}

func (s *LookupResourceGroupEventsResponseBodyEvents) SetChangeType(v string) *LookupResourceGroupEventsResponseBodyEvents {
	s.ChangeType = &v
	return s
}

func (s *LookupResourceGroupEventsResponseBodyEvents) SetEventTime(v string) *LookupResourceGroupEventsResponseBodyEvents {
	s.EventTime = &v
	return s
}

func (s *LookupResourceGroupEventsResponseBodyEvents) SetRegionId(v string) *LookupResourceGroupEventsResponseBodyEvents {
	s.RegionId = &v
	return s
}

func (s *LookupResourceGroupEventsResponseBodyEvents) SetResourceGroupDisplayName(v string) *LookupResourceGroupEventsResponseBodyEvents {
	s.ResourceGroupDisplayName = &v
	return s
}

func (s *LookupResourceGroupEventsResponseBodyEvents) SetResourceGroupId(v string) *LookupResourceGroupEventsResponseBodyEvents {
	s.ResourceGroupId = &v
	return s
}

func (s *LookupResourceGroupEventsResponseBodyEvents) SetResourceId(v string) *LookupResourceGroupEventsResponseBodyEvents {
	s.ResourceId = &v
	return s
}

func (s *LookupResourceGroupEventsResponseBodyEvents) SetResourceType(v string) *LookupResourceGroupEventsResponseBodyEvents {
	s.ResourceType = &v
	return s
}

func (s *LookupResourceGroupEventsResponseBodyEvents) SetService(v string) *LookupResourceGroupEventsResponseBodyEvents {
	s.Service = &v
	return s
}

func (s *LookupResourceGroupEventsResponseBodyEvents) SetSourceResourceGroupInfo(v *LookupResourceGroupEventsResponseBodyEventsSourceResourceGroupInfo) *LookupResourceGroupEventsResponseBodyEvents {
	s.SourceResourceGroupInfo = v
	return s
}

func (s *LookupResourceGroupEventsResponseBodyEvents) SetTargetResourceGroupInfo(v *LookupResourceGroupEventsResponseBodyEventsTargetResourceGroupInfo) *LookupResourceGroupEventsResponseBodyEvents {
	s.TargetResourceGroupInfo = v
	return s
}

func (s *LookupResourceGroupEventsResponseBodyEvents) Validate() error {
	if s.SourceResourceGroupInfo != nil {
		if err := s.SourceResourceGroupInfo.Validate(); err != nil {
			return err
		}
	}
	if s.TargetResourceGroupInfo != nil {
		if err := s.TargetResourceGroupInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type LookupResourceGroupEventsResponseBodyEventsSourceResourceGroupInfo struct {
	// example:
	//
	// TestGroupA
	ResourceGroupDisplayName *string `json:"ResourceGroupDisplayName,omitempty" xml:"ResourceGroupDisplayName,omitempty"`
	// example:
	//
	// rg-aekz25pfurj****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s LookupResourceGroupEventsResponseBodyEventsSourceResourceGroupInfo) String() string {
	return dara.Prettify(s)
}

func (s LookupResourceGroupEventsResponseBodyEventsSourceResourceGroupInfo) GoString() string {
	return s.String()
}

func (s *LookupResourceGroupEventsResponseBodyEventsSourceResourceGroupInfo) GetResourceGroupDisplayName() *string {
	return s.ResourceGroupDisplayName
}

func (s *LookupResourceGroupEventsResponseBodyEventsSourceResourceGroupInfo) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *LookupResourceGroupEventsResponseBodyEventsSourceResourceGroupInfo) SetResourceGroupDisplayName(v string) *LookupResourceGroupEventsResponseBodyEventsSourceResourceGroupInfo {
	s.ResourceGroupDisplayName = &v
	return s
}

func (s *LookupResourceGroupEventsResponseBodyEventsSourceResourceGroupInfo) SetResourceGroupId(v string) *LookupResourceGroupEventsResponseBodyEventsSourceResourceGroupInfo {
	s.ResourceGroupId = &v
	return s
}

func (s *LookupResourceGroupEventsResponseBodyEventsSourceResourceGroupInfo) Validate() error {
	return dara.Validate(s)
}

type LookupResourceGroupEventsResponseBodyEventsTargetResourceGroupInfo struct {
	// example:
	//
	// TestGroupB
	ResourceGroupDisplayName *string `json:"ResourceGroupDisplayName,omitempty" xml:"ResourceGroupDisplayName,omitempty"`
	// example:
	//
	// rg-acfmwfrxcre****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s LookupResourceGroupEventsResponseBodyEventsTargetResourceGroupInfo) String() string {
	return dara.Prettify(s)
}

func (s LookupResourceGroupEventsResponseBodyEventsTargetResourceGroupInfo) GoString() string {
	return s.String()
}

func (s *LookupResourceGroupEventsResponseBodyEventsTargetResourceGroupInfo) GetResourceGroupDisplayName() *string {
	return s.ResourceGroupDisplayName
}

func (s *LookupResourceGroupEventsResponseBodyEventsTargetResourceGroupInfo) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *LookupResourceGroupEventsResponseBodyEventsTargetResourceGroupInfo) SetResourceGroupDisplayName(v string) *LookupResourceGroupEventsResponseBodyEventsTargetResourceGroupInfo {
	s.ResourceGroupDisplayName = &v
	return s
}

func (s *LookupResourceGroupEventsResponseBodyEventsTargetResourceGroupInfo) SetResourceGroupId(v string) *LookupResourceGroupEventsResponseBodyEventsTargetResourceGroupInfo {
	s.ResourceGroupId = &v
	return s
}

func (s *LookupResourceGroupEventsResponseBodyEventsTargetResourceGroupInfo) Validate() error {
	return dara.Validate(s)
}
