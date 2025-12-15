// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRegionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRegionIds(v *DescribeRegionsResponseBodyRegionIds) *DescribeRegionsResponseBody
	GetRegionIds() *DescribeRegionsResponseBodyRegionIds
	SetRequestId(v string) *DescribeRegionsResponseBody
	GetRequestId() *string
}

type DescribeRegionsResponseBody struct {
	// The **region IDs**.
	RegionIds *DescribeRegionsResponseBodyRegionIds `json:"RegionIds,omitempty" xml:"RegionIds,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// AD425AD3-CC7B-4EE2-A5CB-2F61BA73****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRegionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) GetRegionIds() *DescribeRegionsResponseBodyRegionIds {
	return s.RegionIds
}

func (s *DescribeRegionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRegionsResponseBody) SetRegionIds(v *DescribeRegionsResponseBodyRegionIds) *DescribeRegionsResponseBody {
	s.RegionIds = v
	return s
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRegionsResponseBody) Validate() error {
	if s.RegionIds != nil {
		if err := s.RegionIds.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRegionsResponseBodyRegionIds struct {
	KVStoreRegion []*DescribeRegionsResponseBodyRegionIdsKVStoreRegion `json:"KVStoreRegion,omitempty" xml:"KVStoreRegion,omitempty" type:"Repeated"`
}

func (s DescribeRegionsResponseBodyRegionIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegionIds) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegionIds) GetKVStoreRegion() []*DescribeRegionsResponseBodyRegionIdsKVStoreRegion {
	return s.KVStoreRegion
}

func (s *DescribeRegionsResponseBodyRegionIds) SetKVStoreRegion(v []*DescribeRegionsResponseBodyRegionIdsKVStoreRegion) *DescribeRegionsResponseBodyRegionIds {
	s.KVStoreRegion = v
	return s
}

func (s *DescribeRegionsResponseBodyRegionIds) Validate() error {
	if s.KVStoreRegion != nil {
		for _, item := range s.KVStoreRegion {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRegionsResponseBodyRegionIdsKVStoreRegion struct {
	// The name of the region.
	//
	// example:
	//
	// China (Hangzhou)
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// The endpoint of the region.
	//
	// example:
	//
	// r-kvstore.aliyuncs.com
	RegionEndpoint *string `json:"RegionEndpoint,omitempty" xml:"RegionEndpoint,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The zone IDs.
	ZoneIdList *DescribeRegionsResponseBodyRegionIdsKVStoreRegionZoneIdList `json:"ZoneIdList,omitempty" xml:"ZoneIdList,omitempty" type:"Struct"`
	// The IDs of the zones in the region.
	//
	// example:
	//
	// cn-hangzhou-h
	ZoneIds *string `json:"ZoneIds,omitempty" xml:"ZoneIds,omitempty"`
}

func (s DescribeRegionsResponseBodyRegionIdsKVStoreRegion) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegionIdsKVStoreRegion) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegionIdsKVStoreRegion) GetLocalName() *string {
	return s.LocalName
}

func (s *DescribeRegionsResponseBodyRegionIdsKVStoreRegion) GetRegionEndpoint() *string {
	return s.RegionEndpoint
}

func (s *DescribeRegionsResponseBodyRegionIdsKVStoreRegion) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRegionsResponseBodyRegionIdsKVStoreRegion) GetZoneIdList() *DescribeRegionsResponseBodyRegionIdsKVStoreRegionZoneIdList {
	return s.ZoneIdList
}

func (s *DescribeRegionsResponseBodyRegionIdsKVStoreRegion) GetZoneIds() *string {
	return s.ZoneIds
}

func (s *DescribeRegionsResponseBodyRegionIdsKVStoreRegion) SetLocalName(v string) *DescribeRegionsResponseBodyRegionIdsKVStoreRegion {
	s.LocalName = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionIdsKVStoreRegion) SetRegionEndpoint(v string) *DescribeRegionsResponseBodyRegionIdsKVStoreRegion {
	s.RegionEndpoint = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionIdsKVStoreRegion) SetRegionId(v string) *DescribeRegionsResponseBodyRegionIdsKVStoreRegion {
	s.RegionId = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionIdsKVStoreRegion) SetZoneIdList(v *DescribeRegionsResponseBodyRegionIdsKVStoreRegionZoneIdList) *DescribeRegionsResponseBodyRegionIdsKVStoreRegion {
	s.ZoneIdList = v
	return s
}

func (s *DescribeRegionsResponseBodyRegionIdsKVStoreRegion) SetZoneIds(v string) *DescribeRegionsResponseBodyRegionIdsKVStoreRegion {
	s.ZoneIds = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionIdsKVStoreRegion) Validate() error {
	if s.ZoneIdList != nil {
		if err := s.ZoneIdList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRegionsResponseBodyRegionIdsKVStoreRegionZoneIdList struct {
	ZoneId []*string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty" type:"Repeated"`
}

func (s DescribeRegionsResponseBodyRegionIdsKVStoreRegionZoneIdList) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegionIdsKVStoreRegionZoneIdList) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegionIdsKVStoreRegionZoneIdList) GetZoneId() []*string {
	return s.ZoneId
}

func (s *DescribeRegionsResponseBodyRegionIdsKVStoreRegionZoneIdList) SetZoneId(v []*string) *DescribeRegionsResponseBodyRegionIdsKVStoreRegionZoneIdList {
	s.ZoneId = v
	return s
}

func (s *DescribeRegionsResponseBodyRegionIdsKVStoreRegionZoneIdList) Validate() error {
	return dara.Validate(s)
}
