// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopySnapshotShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDestinationRegionIdsShrink(v string) *CopySnapshotShrinkRequest
	GetDestinationRegionIdsShrink() *string
	SetDestinationSnapshotDescription(v string) *CopySnapshotShrinkRequest
	GetDestinationSnapshotDescription() *string
	SetDestinationSnapshotName(v string) *CopySnapshotShrinkRequest
	GetDestinationSnapshotName() *string
	SetInstanceBillingCycle(v string) *CopySnapshotShrinkRequest
	GetInstanceBillingCycle() *string
	SetSnapshotId(v string) *CopySnapshotShrinkRequest
	GetSnapshotId() *string
}

type CopySnapshotShrinkRequest struct {
	// The IDs of destination nodes.
	//
	// This parameter is required.
	DestinationRegionIdsShrink *string `json:"DestinationRegionIds,omitempty" xml:"DestinationRegionIds,omitempty"`
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

func (s CopySnapshotShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CopySnapshotShrinkRequest) GoString() string {
	return s.String()
}

func (s *CopySnapshotShrinkRequest) GetDestinationRegionIdsShrink() *string {
	return s.DestinationRegionIdsShrink
}

func (s *CopySnapshotShrinkRequest) GetDestinationSnapshotDescription() *string {
	return s.DestinationSnapshotDescription
}

func (s *CopySnapshotShrinkRequest) GetDestinationSnapshotName() *string {
	return s.DestinationSnapshotName
}

func (s *CopySnapshotShrinkRequest) GetInstanceBillingCycle() *string {
	return s.InstanceBillingCycle
}

func (s *CopySnapshotShrinkRequest) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *CopySnapshotShrinkRequest) SetDestinationRegionIdsShrink(v string) *CopySnapshotShrinkRequest {
	s.DestinationRegionIdsShrink = &v
	return s
}

func (s *CopySnapshotShrinkRequest) SetDestinationSnapshotDescription(v string) *CopySnapshotShrinkRequest {
	s.DestinationSnapshotDescription = &v
	return s
}

func (s *CopySnapshotShrinkRequest) SetDestinationSnapshotName(v string) *CopySnapshotShrinkRequest {
	s.DestinationSnapshotName = &v
	return s
}

func (s *CopySnapshotShrinkRequest) SetInstanceBillingCycle(v string) *CopySnapshotShrinkRequest {
	s.InstanceBillingCycle = &v
	return s
}

func (s *CopySnapshotShrinkRequest) SetSnapshotId(v string) *CopySnapshotShrinkRequest {
	s.SnapshotId = &v
	return s
}

func (s *CopySnapshotShrinkRequest) Validate() error {
	return dara.Validate(s)
}
