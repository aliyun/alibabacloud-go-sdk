// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopySnapshotRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDestinationRegionIds(v []*string) *CopySnapshotRequest
	GetDestinationRegionIds() []*string
	SetDestinationSnapshotDescription(v string) *CopySnapshotRequest
	GetDestinationSnapshotDescription() *string
	SetDestinationSnapshotName(v string) *CopySnapshotRequest
	GetDestinationSnapshotName() *string
	SetInstanceBillingCycle(v string) *CopySnapshotRequest
	GetInstanceBillingCycle() *string
	SetSnapshotId(v string) *CopySnapshotRequest
	GetSnapshotId() *string
}

type CopySnapshotRequest struct {
	// The IDs of destination nodes.
	//
	// This parameter is required.
	DestinationRegionIds []*string `json:"DestinationRegionIds,omitempty" xml:"DestinationRegionIds,omitempty" type:"Repeated"`
	// The description of the snapshot. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// testDescription
	DestinationSnapshotDescription *string `json:"DestinationSnapshotDescription,omitempty" xml:"DestinationSnapshotDescription,omitempty"`
	// The name of the snapshot. The name must be 2 to 128 characters in length. It must start with a letter and cannot start with `http://` or `https://`. It can contain letters, digits, colons (:), underscores (_), and hyphens (-).
	//
	// example:
	//
	// testSnapshotName
	DestinationSnapshotName *string `json:"DestinationSnapshotName,omitempty" xml:"DestinationSnapshotName,omitempty"`
	InstanceBillingCycle    *string `json:"InstanceBillingCycle,omitempty" xml:"InstanceBillingCycle,omitempty"`
	// The ID of the source snapshot.
	//
	// This parameter is required.
	//
	// example:
	//
	// sp-bp1c0doj0taqyzzl****
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
}

func (s CopySnapshotRequest) String() string {
	return dara.Prettify(s)
}

func (s CopySnapshotRequest) GoString() string {
	return s.String()
}

func (s *CopySnapshotRequest) GetDestinationRegionIds() []*string {
	return s.DestinationRegionIds
}

func (s *CopySnapshotRequest) GetDestinationSnapshotDescription() *string {
	return s.DestinationSnapshotDescription
}

func (s *CopySnapshotRequest) GetDestinationSnapshotName() *string {
	return s.DestinationSnapshotName
}

func (s *CopySnapshotRequest) GetInstanceBillingCycle() *string {
	return s.InstanceBillingCycle
}

func (s *CopySnapshotRequest) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *CopySnapshotRequest) SetDestinationRegionIds(v []*string) *CopySnapshotRequest {
	s.DestinationRegionIds = v
	return s
}

func (s *CopySnapshotRequest) SetDestinationSnapshotDescription(v string) *CopySnapshotRequest {
	s.DestinationSnapshotDescription = &v
	return s
}

func (s *CopySnapshotRequest) SetDestinationSnapshotName(v string) *CopySnapshotRequest {
	s.DestinationSnapshotName = &v
	return s
}

func (s *CopySnapshotRequest) SetInstanceBillingCycle(v string) *CopySnapshotRequest {
	s.InstanceBillingCycle = &v
	return s
}

func (s *CopySnapshotRequest) SetSnapshotId(v string) *CopySnapshotRequest {
	s.SnapshotId = &v
	return s
}

func (s *CopySnapshotRequest) Validate() error {
	return dara.Validate(s)
}
