// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRCResourcesModificationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAvailableZones(v []*DescribeRCResourcesModificationResponseBodyAvailableZones) *DescribeRCResourcesModificationResponseBody
	GetAvailableZones() []*DescribeRCResourcesModificationResponseBodyAvailableZones
	SetRequestId(v string) *DescribeRCResourcesModificationResponseBody
	GetRequestId() *string
}

type DescribeRCResourcesModificationResponseBody struct {
	AvailableZones []*DescribeRCResourcesModificationResponseBodyAvailableZones `json:"AvailableZones,omitempty" xml:"AvailableZones,omitempty" type:"Repeated"`
	RequestId      *string                                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRCResourcesModificationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCResourcesModificationResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRCResourcesModificationResponseBody) GetAvailableZones() []*DescribeRCResourcesModificationResponseBodyAvailableZones {
	return s.AvailableZones
}

func (s *DescribeRCResourcesModificationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRCResourcesModificationResponseBody) SetAvailableZones(v []*DescribeRCResourcesModificationResponseBodyAvailableZones) *DescribeRCResourcesModificationResponseBody {
	s.AvailableZones = v
	return s
}

func (s *DescribeRCResourcesModificationResponseBody) SetRequestId(v string) *DescribeRCResourcesModificationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRCResourcesModificationResponseBody) Validate() error {
	if s.AvailableZones != nil {
		for _, item := range s.AvailableZones {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRCResourcesModificationResponseBodyAvailableZones struct {
	AvailableResources []*DescribeRCResourcesModificationResponseBodyAvailableZonesAvailableResources `json:"AvailableResources,omitempty" xml:"AvailableResources,omitempty" type:"Repeated"`
	RegionId           *string                                                                        `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Status             *string                                                                        `json:"Status,omitempty" xml:"Status,omitempty"`
	StatusCategory     *string                                                                        `json:"StatusCategory,omitempty" xml:"StatusCategory,omitempty"`
	ZoneId             *string                                                                        `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeRCResourcesModificationResponseBodyAvailableZones) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCResourcesModificationResponseBodyAvailableZones) GoString() string {
	return s.String()
}

func (s *DescribeRCResourcesModificationResponseBodyAvailableZones) GetAvailableResources() []*DescribeRCResourcesModificationResponseBodyAvailableZonesAvailableResources {
	return s.AvailableResources
}

func (s *DescribeRCResourcesModificationResponseBodyAvailableZones) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRCResourcesModificationResponseBodyAvailableZones) GetStatus() *string {
	return s.Status
}

func (s *DescribeRCResourcesModificationResponseBodyAvailableZones) GetStatusCategory() *string {
	return s.StatusCategory
}

func (s *DescribeRCResourcesModificationResponseBodyAvailableZones) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeRCResourcesModificationResponseBodyAvailableZones) SetAvailableResources(v []*DescribeRCResourcesModificationResponseBodyAvailableZonesAvailableResources) *DescribeRCResourcesModificationResponseBodyAvailableZones {
	s.AvailableResources = v
	return s
}

func (s *DescribeRCResourcesModificationResponseBodyAvailableZones) SetRegionId(v string) *DescribeRCResourcesModificationResponseBodyAvailableZones {
	s.RegionId = &v
	return s
}

func (s *DescribeRCResourcesModificationResponseBodyAvailableZones) SetStatus(v string) *DescribeRCResourcesModificationResponseBodyAvailableZones {
	s.Status = &v
	return s
}

func (s *DescribeRCResourcesModificationResponseBodyAvailableZones) SetStatusCategory(v string) *DescribeRCResourcesModificationResponseBodyAvailableZones {
	s.StatusCategory = &v
	return s
}

func (s *DescribeRCResourcesModificationResponseBodyAvailableZones) SetZoneId(v string) *DescribeRCResourcesModificationResponseBodyAvailableZones {
	s.ZoneId = &v
	return s
}

func (s *DescribeRCResourcesModificationResponseBodyAvailableZones) Validate() error {
	if s.AvailableResources != nil {
		for _, item := range s.AvailableResources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRCResourcesModificationResponseBodyAvailableZonesAvailableResources struct {
	SupportedResources []*DescribeRCResourcesModificationResponseBodyAvailableZonesAvailableResourcesSupportedResources `json:"SupportedResources,omitempty" xml:"SupportedResources,omitempty" type:"Repeated"`
	Type               *string                                                                                          `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeRCResourcesModificationResponseBodyAvailableZonesAvailableResources) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCResourcesModificationResponseBodyAvailableZonesAvailableResources) GoString() string {
	return s.String()
}

func (s *DescribeRCResourcesModificationResponseBodyAvailableZonesAvailableResources) GetSupportedResources() []*DescribeRCResourcesModificationResponseBodyAvailableZonesAvailableResourcesSupportedResources {
	return s.SupportedResources
}

func (s *DescribeRCResourcesModificationResponseBodyAvailableZonesAvailableResources) GetType() *string {
	return s.Type
}

func (s *DescribeRCResourcesModificationResponseBodyAvailableZonesAvailableResources) SetSupportedResources(v []*DescribeRCResourcesModificationResponseBodyAvailableZonesAvailableResourcesSupportedResources) *DescribeRCResourcesModificationResponseBodyAvailableZonesAvailableResources {
	s.SupportedResources = v
	return s
}

func (s *DescribeRCResourcesModificationResponseBodyAvailableZonesAvailableResources) SetType(v string) *DescribeRCResourcesModificationResponseBodyAvailableZonesAvailableResources {
	s.Type = &v
	return s
}

func (s *DescribeRCResourcesModificationResponseBodyAvailableZonesAvailableResources) Validate() error {
	if s.SupportedResources != nil {
		for _, item := range s.SupportedResources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRCResourcesModificationResponseBodyAvailableZonesAvailableResourcesSupportedResources struct {
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
	StatusCategory *string `json:"StatusCategory,omitempty" xml:"StatusCategory,omitempty"`
	Value          *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeRCResourcesModificationResponseBodyAvailableZonesAvailableResourcesSupportedResources) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCResourcesModificationResponseBodyAvailableZonesAvailableResourcesSupportedResources) GoString() string {
	return s.String()
}

func (s *DescribeRCResourcesModificationResponseBodyAvailableZonesAvailableResourcesSupportedResources) GetStatus() *string {
	return s.Status
}

func (s *DescribeRCResourcesModificationResponseBodyAvailableZonesAvailableResourcesSupportedResources) GetStatusCategory() *string {
	return s.StatusCategory
}

func (s *DescribeRCResourcesModificationResponseBodyAvailableZonesAvailableResourcesSupportedResources) GetValue() *string {
	return s.Value
}

func (s *DescribeRCResourcesModificationResponseBodyAvailableZonesAvailableResourcesSupportedResources) SetStatus(v string) *DescribeRCResourcesModificationResponseBodyAvailableZonesAvailableResourcesSupportedResources {
	s.Status = &v
	return s
}

func (s *DescribeRCResourcesModificationResponseBodyAvailableZonesAvailableResourcesSupportedResources) SetStatusCategory(v string) *DescribeRCResourcesModificationResponseBodyAvailableZonesAvailableResourcesSupportedResources {
	s.StatusCategory = &v
	return s
}

func (s *DescribeRCResourcesModificationResponseBodyAvailableZonesAvailableResourcesSupportedResources) SetValue(v string) *DescribeRCResourcesModificationResponseBodyAvailableZonesAvailableResourcesSupportedResources {
	s.Value = &v
	return s
}

func (s *DescribeRCResourcesModificationResponseBodyAvailableZonesAvailableResourcesSupportedResources) Validate() error {
	return dara.Validate(s)
}
