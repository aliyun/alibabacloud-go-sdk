// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSummariesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCenterSummary(v *ListSummariesResponseBodyCenterSummary) *ListSummariesResponseBody
	GetCenterSummary() *ListSummariesResponseBodyCenterSummary
	SetRegionSummaries(v []*ListSummariesResponseBodyRegionSummaries) *ListSummariesResponseBody
	GetRegionSummaries() []*ListSummariesResponseBodyRegionSummaries
	SetRequestId(v string) *ListSummariesResponseBody
	GetRequestId() *string
}

type ListSummariesResponseBody struct {
	CenterSummary   *ListSummariesResponseBodyCenterSummary     `json:"CenterSummary,omitempty" xml:"CenterSummary,omitempty" type:"Struct"`
	RegionSummaries []*ListSummariesResponseBodyRegionSummaries `json:"RegionSummaries,omitempty" xml:"RegionSummaries,omitempty" type:"Repeated"`
	RequestId       *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListSummariesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSummariesResponseBody) GoString() string {
	return s.String()
}

func (s *ListSummariesResponseBody) GetCenterSummary() *ListSummariesResponseBodyCenterSummary {
	return s.CenterSummary
}

func (s *ListSummariesResponseBody) GetRegionSummaries() []*ListSummariesResponseBodyRegionSummaries {
	return s.RegionSummaries
}

func (s *ListSummariesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSummariesResponseBody) SetCenterSummary(v *ListSummariesResponseBodyCenterSummary) *ListSummariesResponseBody {
	s.CenterSummary = v
	return s
}

func (s *ListSummariesResponseBody) SetRegionSummaries(v []*ListSummariesResponseBodyRegionSummaries) *ListSummariesResponseBody {
	s.RegionSummaries = v
	return s
}

func (s *ListSummariesResponseBody) SetRequestId(v string) *ListSummariesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSummariesResponseBody) Validate() error {
	if s.CenterSummary != nil {
		if err := s.CenterSummary.Validate(); err != nil {
			return err
		}
	}
	if s.RegionSummaries != nil {
		for _, item := range s.RegionSummaries {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSummariesResponseBodyCenterSummary struct {
	RegisteredResourceTypeCount *int32  `json:"RegisteredResourceTypeCount,omitempty" xml:"RegisteredResourceTypeCount,omitempty"`
	TemplateCount               *string `json:"TemplateCount,omitempty" xml:"TemplateCount,omitempty"`
}

func (s ListSummariesResponseBodyCenterSummary) String() string {
	return dara.Prettify(s)
}

func (s ListSummariesResponseBodyCenterSummary) GoString() string {
	return s.String()
}

func (s *ListSummariesResponseBodyCenterSummary) GetRegisteredResourceTypeCount() *int32 {
	return s.RegisteredResourceTypeCount
}

func (s *ListSummariesResponseBodyCenterSummary) GetTemplateCount() *string {
	return s.TemplateCount
}

func (s *ListSummariesResponseBodyCenterSummary) SetRegisteredResourceTypeCount(v int32) *ListSummariesResponseBodyCenterSummary {
	s.RegisteredResourceTypeCount = &v
	return s
}

func (s *ListSummariesResponseBodyCenterSummary) SetTemplateCount(v string) *ListSummariesResponseBodyCenterSummary {
	s.TemplateCount = &v
	return s
}

func (s *ListSummariesResponseBodyCenterSummary) Validate() error {
	return dara.Validate(s)
}

type ListSummariesResponseBodyRegionSummaries struct {
	RegionId             *string                                                 `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StackCount           *string                                                 `json:"StackCount,omitempty" xml:"StackCount,omitempty"`
	StackDetails         []*ListSummariesResponseBodyRegionSummariesStackDetails `json:"StackDetails,omitempty" xml:"StackDetails,omitempty" type:"Repeated"`
	StackGroupCount      *string                                                 `json:"StackGroupCount,omitempty" xml:"StackGroupCount,omitempty"`
	TemplateScratchCount *int32                                                  `json:"TemplateScratchCount,omitempty" xml:"TemplateScratchCount,omitempty"`
}

func (s ListSummariesResponseBodyRegionSummaries) String() string {
	return dara.Prettify(s)
}

func (s ListSummariesResponseBodyRegionSummaries) GoString() string {
	return s.String()
}

func (s *ListSummariesResponseBodyRegionSummaries) GetRegionId() *string {
	return s.RegionId
}

func (s *ListSummariesResponseBodyRegionSummaries) GetStackCount() *string {
	return s.StackCount
}

func (s *ListSummariesResponseBodyRegionSummaries) GetStackDetails() []*ListSummariesResponseBodyRegionSummariesStackDetails {
	return s.StackDetails
}

func (s *ListSummariesResponseBodyRegionSummaries) GetStackGroupCount() *string {
	return s.StackGroupCount
}

func (s *ListSummariesResponseBodyRegionSummaries) GetTemplateScratchCount() *int32 {
	return s.TemplateScratchCount
}

func (s *ListSummariesResponseBodyRegionSummaries) SetRegionId(v string) *ListSummariesResponseBodyRegionSummaries {
	s.RegionId = &v
	return s
}

func (s *ListSummariesResponseBodyRegionSummaries) SetStackCount(v string) *ListSummariesResponseBodyRegionSummaries {
	s.StackCount = &v
	return s
}

func (s *ListSummariesResponseBodyRegionSummaries) SetStackDetails(v []*ListSummariesResponseBodyRegionSummariesStackDetails) *ListSummariesResponseBodyRegionSummaries {
	s.StackDetails = v
	return s
}

func (s *ListSummariesResponseBodyRegionSummaries) SetStackGroupCount(v string) *ListSummariesResponseBodyRegionSummaries {
	s.StackGroupCount = &v
	return s
}

func (s *ListSummariesResponseBodyRegionSummaries) SetTemplateScratchCount(v int32) *ListSummariesResponseBodyRegionSummaries {
	s.TemplateScratchCount = &v
	return s
}

func (s *ListSummariesResponseBodyRegionSummaries) Validate() error {
	if s.StackDetails != nil {
		for _, item := range s.StackDetails {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSummariesResponseBodyRegionSummariesStackDetails struct {
	BriefStatus *string `json:"BriefStatus,omitempty" xml:"BriefStatus,omitempty"`
	Count       *string `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s ListSummariesResponseBodyRegionSummariesStackDetails) String() string {
	return dara.Prettify(s)
}

func (s ListSummariesResponseBodyRegionSummariesStackDetails) GoString() string {
	return s.String()
}

func (s *ListSummariesResponseBodyRegionSummariesStackDetails) GetBriefStatus() *string {
	return s.BriefStatus
}

func (s *ListSummariesResponseBodyRegionSummariesStackDetails) GetCount() *string {
	return s.Count
}

func (s *ListSummariesResponseBodyRegionSummariesStackDetails) SetBriefStatus(v string) *ListSummariesResponseBodyRegionSummariesStackDetails {
	s.BriefStatus = &v
	return s
}

func (s *ListSummariesResponseBodyRegionSummariesStackDetails) SetCount(v string) *ListSummariesResponseBodyRegionSummariesStackDetails {
	s.Count = &v
	return s
}

func (s *ListSummariesResponseBodyRegionSummariesStackDetails) Validate() error {
	return dara.Validate(s)
}
