// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClearReplicaGroupDrillRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDrillId(v string) *ClearReplicaGroupDrillRequest
	GetDrillId() *string
	SetGroupId(v string) *ClearReplicaGroupDrillRequest
	GetGroupId() *string
	SetRegionId(v string) *ClearReplicaGroupDrillRequest
	GetRegionId() *string
}

type ClearReplicaGroupDrillRequest struct {
	// The ID of the drill. You can call the [DescribeReplicaGroupDrills](https://help.aliyun.com/document_detail/2584481.html) operation to query disaster recovery drills that were performed on replication pairs in a specific region.
	//
	// This parameter is required.
	//
	// example:
	//
	// pg-drill-xxxx
	DrillId *string `json:"DrillId,omitempty" xml:"DrillId,omitempty"`
	// The ID of the replication pair-consistent group. You can call the [DescribeDiskReplicaGroups](https://help.aliyun.com/document_detail/426614.html) operation to query the most recent list of replication pair-consistent groups, including group IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// pg-xxxx
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/354276.html) operation to query the most recent list of regions in which async replication is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ClearReplicaGroupDrillRequest) String() string {
	return dara.Prettify(s)
}

func (s ClearReplicaGroupDrillRequest) GoString() string {
	return s.String()
}

func (s *ClearReplicaGroupDrillRequest) GetDrillId() *string {
	return s.DrillId
}

func (s *ClearReplicaGroupDrillRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *ClearReplicaGroupDrillRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ClearReplicaGroupDrillRequest) SetDrillId(v string) *ClearReplicaGroupDrillRequest {
	s.DrillId = &v
	return s
}

func (s *ClearReplicaGroupDrillRequest) SetGroupId(v string) *ClearReplicaGroupDrillRequest {
	s.GroupId = &v
	return s
}

func (s *ClearReplicaGroupDrillRequest) SetRegionId(v string) *ClearReplicaGroupDrillRequest {
	s.RegionId = &v
	return s
}

func (s *ClearReplicaGroupDrillRequest) Validate() error {
	return dara.Validate(s)
}
