// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRCAvailableResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAvailableZones(v []*DescribeRCAvailableResourceResponseBodyAvailableZones) *DescribeRCAvailableResourceResponseBody
	GetAvailableZones() []*DescribeRCAvailableResourceResponseBodyAvailableZones
	SetRequestId(v string) *DescribeRCAvailableResourceResponseBody
	GetRequestId() *string
}

type DescribeRCAvailableResourceResponseBody struct {
	AvailableZones []*DescribeRCAvailableResourceResponseBodyAvailableZones `json:"AvailableZones,omitempty" xml:"AvailableZones,omitempty" type:"Repeated"`
	RequestId      *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRCAvailableResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCAvailableResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRCAvailableResourceResponseBody) GetAvailableZones() []*DescribeRCAvailableResourceResponseBodyAvailableZones {
	return s.AvailableZones
}

func (s *DescribeRCAvailableResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRCAvailableResourceResponseBody) SetAvailableZones(v []*DescribeRCAvailableResourceResponseBodyAvailableZones) *DescribeRCAvailableResourceResponseBody {
	s.AvailableZones = v
	return s
}

func (s *DescribeRCAvailableResourceResponseBody) SetRequestId(v string) *DescribeRCAvailableResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRCAvailableResourceResponseBody) Validate() error {
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

type DescribeRCAvailableResourceResponseBodyAvailableZones struct {
	AvailableResources []*DescribeRCAvailableResourceResponseBodyAvailableZonesAvailableResources `json:"AvailableResources,omitempty" xml:"AvailableResources,omitempty" type:"Repeated"`
	RegionId           *string                                                                    `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Status             *string                                                                    `json:"Status,omitempty" xml:"Status,omitempty"`
	StatusCategory     *string                                                                    `json:"StatusCategory,omitempty" xml:"StatusCategory,omitempty"`
	ZoneId             *string                                                                    `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeRCAvailableResourceResponseBodyAvailableZones) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCAvailableResourceResponseBodyAvailableZones) GoString() string {
	return s.String()
}

func (s *DescribeRCAvailableResourceResponseBodyAvailableZones) GetAvailableResources() []*DescribeRCAvailableResourceResponseBodyAvailableZonesAvailableResources {
	return s.AvailableResources
}

func (s *DescribeRCAvailableResourceResponseBodyAvailableZones) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRCAvailableResourceResponseBodyAvailableZones) GetStatus() *string {
	return s.Status
}

func (s *DescribeRCAvailableResourceResponseBodyAvailableZones) GetStatusCategory() *string {
	return s.StatusCategory
}

func (s *DescribeRCAvailableResourceResponseBodyAvailableZones) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeRCAvailableResourceResponseBodyAvailableZones) SetAvailableResources(v []*DescribeRCAvailableResourceResponseBodyAvailableZonesAvailableResources) *DescribeRCAvailableResourceResponseBodyAvailableZones {
	s.AvailableResources = v
	return s
}

func (s *DescribeRCAvailableResourceResponseBodyAvailableZones) SetRegionId(v string) *DescribeRCAvailableResourceResponseBodyAvailableZones {
	s.RegionId = &v
	return s
}

func (s *DescribeRCAvailableResourceResponseBodyAvailableZones) SetStatus(v string) *DescribeRCAvailableResourceResponseBodyAvailableZones {
	s.Status = &v
	return s
}

func (s *DescribeRCAvailableResourceResponseBodyAvailableZones) SetStatusCategory(v string) *DescribeRCAvailableResourceResponseBodyAvailableZones {
	s.StatusCategory = &v
	return s
}

func (s *DescribeRCAvailableResourceResponseBodyAvailableZones) SetZoneId(v string) *DescribeRCAvailableResourceResponseBodyAvailableZones {
	s.ZoneId = &v
	return s
}

func (s *DescribeRCAvailableResourceResponseBodyAvailableZones) Validate() error {
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

type DescribeRCAvailableResourceResponseBodyAvailableZonesAvailableResources struct {
	SupportedResources []*DescribeRCAvailableResourceResponseBodyAvailableZonesAvailableResourcesSupportedResources `json:"SupportedResources,omitempty" xml:"SupportedResources,omitempty" type:"Repeated"`
	Type               *string                                                                                      `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeRCAvailableResourceResponseBodyAvailableZonesAvailableResources) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCAvailableResourceResponseBodyAvailableZonesAvailableResources) GoString() string {
	return s.String()
}

func (s *DescribeRCAvailableResourceResponseBodyAvailableZonesAvailableResources) GetSupportedResources() []*DescribeRCAvailableResourceResponseBodyAvailableZonesAvailableResourcesSupportedResources {
	return s.SupportedResources
}

func (s *DescribeRCAvailableResourceResponseBodyAvailableZonesAvailableResources) GetType() *string {
	return s.Type
}

func (s *DescribeRCAvailableResourceResponseBodyAvailableZonesAvailableResources) SetSupportedResources(v []*DescribeRCAvailableResourceResponseBodyAvailableZonesAvailableResourcesSupportedResources) *DescribeRCAvailableResourceResponseBodyAvailableZonesAvailableResources {
	s.SupportedResources = v
	return s
}

func (s *DescribeRCAvailableResourceResponseBodyAvailableZonesAvailableResources) SetType(v string) *DescribeRCAvailableResourceResponseBodyAvailableZonesAvailableResources {
	s.Type = &v
	return s
}

func (s *DescribeRCAvailableResourceResponseBodyAvailableZonesAvailableResources) Validate() error {
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

type DescribeRCAvailableResourceResponseBodyAvailableZonesAvailableResourcesSupportedResources struct {
	Max            *int32  `json:"Max,omitempty" xml:"Max,omitempty"`
	Min            *int32  `json:"Min,omitempty" xml:"Min,omitempty"`
	QuotaStatus    *string `json:"QuotaStatus,omitempty" xml:"QuotaStatus,omitempty"`
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
	StatusCategory *string `json:"StatusCategory,omitempty" xml:"StatusCategory,omitempty"`
	Unit           *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	Value          *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeRCAvailableResourceResponseBodyAvailableZonesAvailableResourcesSupportedResources) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCAvailableResourceResponseBodyAvailableZonesAvailableResourcesSupportedResources) GoString() string {
	return s.String()
}

func (s *DescribeRCAvailableResourceResponseBodyAvailableZonesAvailableResourcesSupportedResources) GetMax() *int32 {
	return s.Max
}

func (s *DescribeRCAvailableResourceResponseBodyAvailableZonesAvailableResourcesSupportedResources) GetMin() *int32 {
	return s.Min
}

func (s *DescribeRCAvailableResourceResponseBodyAvailableZonesAvailableResourcesSupportedResources) GetQuotaStatus() *string {
	return s.QuotaStatus
}

func (s *DescribeRCAvailableResourceResponseBodyAvailableZonesAvailableResourcesSupportedResources) GetStatus() *string {
	return s.Status
}

func (s *DescribeRCAvailableResourceResponseBodyAvailableZonesAvailableResourcesSupportedResources) GetStatusCategory() *string {
	return s.StatusCategory
}

func (s *DescribeRCAvailableResourceResponseBodyAvailableZonesAvailableResourcesSupportedResources) GetUnit() *string {
	return s.Unit
}

func (s *DescribeRCAvailableResourceResponseBodyAvailableZonesAvailableResourcesSupportedResources) GetValue() *string {
	return s.Value
}

func (s *DescribeRCAvailableResourceResponseBodyAvailableZonesAvailableResourcesSupportedResources) SetMax(v int32) *DescribeRCAvailableResourceResponseBodyAvailableZonesAvailableResourcesSupportedResources {
	s.Max = &v
	return s
}

func (s *DescribeRCAvailableResourceResponseBodyAvailableZonesAvailableResourcesSupportedResources) SetMin(v int32) *DescribeRCAvailableResourceResponseBodyAvailableZonesAvailableResourcesSupportedResources {
	s.Min = &v
	return s
}

func (s *DescribeRCAvailableResourceResponseBodyAvailableZonesAvailableResourcesSupportedResources) SetQuotaStatus(v string) *DescribeRCAvailableResourceResponseBodyAvailableZonesAvailableResourcesSupportedResources {
	s.QuotaStatus = &v
	return s
}

func (s *DescribeRCAvailableResourceResponseBodyAvailableZonesAvailableResourcesSupportedResources) SetStatus(v string) *DescribeRCAvailableResourceResponseBodyAvailableZonesAvailableResourcesSupportedResources {
	s.Status = &v
	return s
}

func (s *DescribeRCAvailableResourceResponseBodyAvailableZonesAvailableResourcesSupportedResources) SetStatusCategory(v string) *DescribeRCAvailableResourceResponseBodyAvailableZonesAvailableResourcesSupportedResources {
	s.StatusCategory = &v
	return s
}

func (s *DescribeRCAvailableResourceResponseBodyAvailableZonesAvailableResourcesSupportedResources) SetUnit(v string) *DescribeRCAvailableResourceResponseBodyAvailableZonesAvailableResourcesSupportedResources {
	s.Unit = &v
	return s
}

func (s *DescribeRCAvailableResourceResponseBodyAvailableZonesAvailableResourcesSupportedResources) SetValue(v string) *DescribeRCAvailableResourceResponseBodyAvailableZonesAvailableResourcesSupportedResources {
	s.Value = &v
	return s
}

func (s *DescribeRCAvailableResourceResponseBodyAvailableZonesAvailableResourcesSupportedResources) Validate() error {
	return dara.Validate(s)
}
