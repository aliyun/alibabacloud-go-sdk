// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSyncEcsHostTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEcsRegions(v *DescribeSyncEcsHostTaskResponseBodyEcsRegions) *DescribeSyncEcsHostTaskResponseBody
	GetEcsRegions() *DescribeSyncEcsHostTaskResponseBodyEcsRegions
	SetRegions(v *DescribeSyncEcsHostTaskResponseBodyRegions) *DescribeSyncEcsHostTaskResponseBody
	GetRegions() *DescribeSyncEcsHostTaskResponseBodyRegions
	SetRequestId(v string) *DescribeSyncEcsHostTaskResponseBody
	GetRequestId() *string
	SetStatus(v string) *DescribeSyncEcsHostTaskResponseBody
	GetStatus() *string
	SetSuccess(v bool) *DescribeSyncEcsHostTaskResponseBody
	GetSuccess() *bool
	SetZoneId(v string) *DescribeSyncEcsHostTaskResponseBody
	GetZoneId() *string
}

type DescribeSyncEcsHostTaskResponseBody struct {
	// The synchronized regions where the ECS instances are deployed.
	EcsRegions *DescribeSyncEcsHostTaskResponseBodyEcsRegions `json:"EcsRegions,omitempty" xml:"EcsRegions,omitempty" type:"Struct"`
	// The synchronized region IDs of the ECS instances.
	Regions *DescribeSyncEcsHostTaskResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 75446CC1-FC9A-4595-8D96-089D73D7A63D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether hostname automatic synchronization is enabled. Valid values:
	//
	// 	- ON: Hostname automatic synchronization is enabled. After this feature is enabled, the system automatically reads the hostnames of the Elastic Compute Service (ECS) instances in the specified regions and updates Domain Name System (DNS) records at an interval of 1 minute.
	//
	// 	- OFF: Hostname automatic synchronization is disabled.
	//
	// example:
	//
	// ON
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the task was successful. Valid values:
	//
	// 	- True
	//
	// 	- False
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The zone ID. This ID uniquely identifies the zone.
	//
	// example:
	//
	// pvtz-test-id-2989149d628c56****
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeSyncEcsHostTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSyncEcsHostTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSyncEcsHostTaskResponseBody) GetEcsRegions() *DescribeSyncEcsHostTaskResponseBodyEcsRegions {
	return s.EcsRegions
}

func (s *DescribeSyncEcsHostTaskResponseBody) GetRegions() *DescribeSyncEcsHostTaskResponseBodyRegions {
	return s.Regions
}

func (s *DescribeSyncEcsHostTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSyncEcsHostTaskResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeSyncEcsHostTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeSyncEcsHostTaskResponseBody) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeSyncEcsHostTaskResponseBody) SetEcsRegions(v *DescribeSyncEcsHostTaskResponseBodyEcsRegions) *DescribeSyncEcsHostTaskResponseBody {
	s.EcsRegions = v
	return s
}

func (s *DescribeSyncEcsHostTaskResponseBody) SetRegions(v *DescribeSyncEcsHostTaskResponseBodyRegions) *DescribeSyncEcsHostTaskResponseBody {
	s.Regions = v
	return s
}

func (s *DescribeSyncEcsHostTaskResponseBody) SetRequestId(v string) *DescribeSyncEcsHostTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSyncEcsHostTaskResponseBody) SetStatus(v string) *DescribeSyncEcsHostTaskResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeSyncEcsHostTaskResponseBody) SetSuccess(v bool) *DescribeSyncEcsHostTaskResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeSyncEcsHostTaskResponseBody) SetZoneId(v string) *DescribeSyncEcsHostTaskResponseBody {
	s.ZoneId = &v
	return s
}

func (s *DescribeSyncEcsHostTaskResponseBody) Validate() error {
	if s.EcsRegions != nil {
		if err := s.EcsRegions.Validate(); err != nil {
			return err
		}
	}
	if s.Regions != nil {
		if err := s.Regions.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSyncEcsHostTaskResponseBodyEcsRegions struct {
	EcsRegion []*DescribeSyncEcsHostTaskResponseBodyEcsRegionsEcsRegion `json:"EcsRegion,omitempty" xml:"EcsRegion,omitempty" type:"Repeated"`
}

func (s DescribeSyncEcsHostTaskResponseBodyEcsRegions) String() string {
	return dara.Prettify(s)
}

func (s DescribeSyncEcsHostTaskResponseBodyEcsRegions) GoString() string {
	return s.String()
}

func (s *DescribeSyncEcsHostTaskResponseBodyEcsRegions) GetEcsRegion() []*DescribeSyncEcsHostTaskResponseBodyEcsRegionsEcsRegion {
	return s.EcsRegion
}

func (s *DescribeSyncEcsHostTaskResponseBodyEcsRegions) SetEcsRegion(v []*DescribeSyncEcsHostTaskResponseBodyEcsRegionsEcsRegion) *DescribeSyncEcsHostTaskResponseBodyEcsRegions {
	s.EcsRegion = v
	return s
}

func (s *DescribeSyncEcsHostTaskResponseBodyEcsRegions) Validate() error {
	if s.EcsRegion != nil {
		for _, item := range s.EcsRegion {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSyncEcsHostTaskResponseBodyEcsRegionsEcsRegion struct {
	// The synchronized region IDs.
	RegionIds *DescribeSyncEcsHostTaskResponseBodyEcsRegionsEcsRegionRegionIds `json:"RegionIds,omitempty" xml:"RegionIds,omitempty" type:"Struct"`
	// The user ID to which the region belongs. This parameter is used in cross-account synchronization scenarios.
	//
	// example:
	//
	// 141339776561****
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DescribeSyncEcsHostTaskResponseBodyEcsRegionsEcsRegion) String() string {
	return dara.Prettify(s)
}

func (s DescribeSyncEcsHostTaskResponseBodyEcsRegionsEcsRegion) GoString() string {
	return s.String()
}

func (s *DescribeSyncEcsHostTaskResponseBodyEcsRegionsEcsRegion) GetRegionIds() *DescribeSyncEcsHostTaskResponseBodyEcsRegionsEcsRegionRegionIds {
	return s.RegionIds
}

func (s *DescribeSyncEcsHostTaskResponseBodyEcsRegionsEcsRegion) GetUserId() *int64 {
	return s.UserId
}

func (s *DescribeSyncEcsHostTaskResponseBodyEcsRegionsEcsRegion) SetRegionIds(v *DescribeSyncEcsHostTaskResponseBodyEcsRegionsEcsRegionRegionIds) *DescribeSyncEcsHostTaskResponseBodyEcsRegionsEcsRegion {
	s.RegionIds = v
	return s
}

func (s *DescribeSyncEcsHostTaskResponseBodyEcsRegionsEcsRegion) SetUserId(v int64) *DescribeSyncEcsHostTaskResponseBodyEcsRegionsEcsRegion {
	s.UserId = &v
	return s
}

func (s *DescribeSyncEcsHostTaskResponseBodyEcsRegionsEcsRegion) Validate() error {
	if s.RegionIds != nil {
		if err := s.RegionIds.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSyncEcsHostTaskResponseBodyEcsRegionsEcsRegionRegionIds struct {
	RegionId []*string `json:"RegionId,omitempty" xml:"RegionId,omitempty" type:"Repeated"`
}

func (s DescribeSyncEcsHostTaskResponseBodyEcsRegionsEcsRegionRegionIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeSyncEcsHostTaskResponseBodyEcsRegionsEcsRegionRegionIds) GoString() string {
	return s.String()
}

func (s *DescribeSyncEcsHostTaskResponseBodyEcsRegionsEcsRegionRegionIds) GetRegionId() []*string {
	return s.RegionId
}

func (s *DescribeSyncEcsHostTaskResponseBodyEcsRegionsEcsRegionRegionIds) SetRegionId(v []*string) *DescribeSyncEcsHostTaskResponseBodyEcsRegionsEcsRegionRegionIds {
	s.RegionId = v
	return s
}

func (s *DescribeSyncEcsHostTaskResponseBodyEcsRegionsEcsRegionRegionIds) Validate() error {
	return dara.Validate(s)
}

type DescribeSyncEcsHostTaskResponseBodyRegions struct {
	RegionId []*string `json:"RegionId,omitempty" xml:"RegionId,omitempty" type:"Repeated"`
}

func (s DescribeSyncEcsHostTaskResponseBodyRegions) String() string {
	return dara.Prettify(s)
}

func (s DescribeSyncEcsHostTaskResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeSyncEcsHostTaskResponseBodyRegions) GetRegionId() []*string {
	return s.RegionId
}

func (s *DescribeSyncEcsHostTaskResponseBodyRegions) SetRegionId(v []*string) *DescribeSyncEcsHostTaskResponseBodyRegions {
	s.RegionId = v
	return s
}

func (s *DescribeSyncEcsHostTaskResponseBodyRegions) Validate() error {
	return dara.Validate(s)
}
