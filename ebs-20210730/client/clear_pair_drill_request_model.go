// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClearPairDrillRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDrillId(v string) *ClearPairDrillRequest
	GetDrillId() *string
	SetPairId(v string) *ClearPairDrillRequest
	GetPairId() *string
	SetRegionId(v string) *ClearPairDrillRequest
	GetRegionId() *string
}

type ClearPairDrillRequest struct {
	// The ID of the drill. You can call the [DescribePairDrills](https://help.aliyun.com/document_detail/2584480.html) operation to query the disaster recovery drills that were performed on replication pairs in a specific region.
	//
	// This parameter is required.
	//
	// example:
	//
	// drill-xxxx
	DrillId *string `json:"DrillId,omitempty" xml:"DrillId,omitempty"`
	// The ID of the replication pair. You can call the [DescribeDiskReplicaPairs](https://help.aliyun.com/document_detail/354206.html) operation to query the most recent list of replication pairs, including replication pair IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// pair-xxxx
	PairId *string `json:"PairId,omitempty" xml:"PairId,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/354276.html) operation to query the most recent list of regions in which async replication is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ClearPairDrillRequest) String() string {
	return dara.Prettify(s)
}

func (s ClearPairDrillRequest) GoString() string {
	return s.String()
}

func (s *ClearPairDrillRequest) GetDrillId() *string {
	return s.DrillId
}

func (s *ClearPairDrillRequest) GetPairId() *string {
	return s.PairId
}

func (s *ClearPairDrillRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ClearPairDrillRequest) SetDrillId(v string) *ClearPairDrillRequest {
	s.DrillId = &v
	return s
}

func (s *ClearPairDrillRequest) SetPairId(v string) *ClearPairDrillRequest {
	s.PairId = &v
	return s
}

func (s *ClearPairDrillRequest) SetRegionId(v string) *ClearPairDrillRequest {
	s.RegionId = &v
	return s
}

func (s *ClearPairDrillRequest) Validate() error {
	return dara.Validate(s)
}
