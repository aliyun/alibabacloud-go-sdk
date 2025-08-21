// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSnapshotRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateSnapshotRequest
	GetDescription() *string
	SetDiskId(v string) *CreateSnapshotRequest
	GetDiskId() *string
	SetEnsRegionId(v string) *CreateSnapshotRequest
	GetEnsRegionId() *string
	SetInstanceBillingCycle(v string) *CreateSnapshotRequest
	GetInstanceBillingCycle() *string
	SetSnapshotName(v string) *CreateSnapshotRequest
	GetSnapshotName() *string
}

type CreateSnapshotRequest struct {
	// The description of the snapshot. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// By default, this parameter is left empty.
	//
	// example:
	//
	// testDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the cloud disk.
	//
	// This parameter is required.
	//
	// example:
	//
	// d-bp1s5fnvk4gn2tws0****
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// The ID of the edge node.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shenzhen-3
	EnsRegionId          *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	InstanceBillingCycle *string `json:"InstanceBillingCycle,omitempty" xml:"InstanceBillingCycle,omitempty"`
	// The name of the snapshot. The name must be 2 to 128 characters in length. It must start with a letter and cannot start with `http://` or `https://`. It can contain letters, digits, colons (:), underscores (_), and hyphens (-).
	//
	// example:
	//
	// testSnapshotName
	SnapshotName *string `json:"SnapshotName,omitempty" xml:"SnapshotName,omitempty"`
}

func (s CreateSnapshotRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSnapshotRequest) GoString() string {
	return s.String()
}

func (s *CreateSnapshotRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateSnapshotRequest) GetDiskId() *string {
	return s.DiskId
}

func (s *CreateSnapshotRequest) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *CreateSnapshotRequest) GetInstanceBillingCycle() *string {
	return s.InstanceBillingCycle
}

func (s *CreateSnapshotRequest) GetSnapshotName() *string {
	return s.SnapshotName
}

func (s *CreateSnapshotRequest) SetDescription(v string) *CreateSnapshotRequest {
	s.Description = &v
	return s
}

func (s *CreateSnapshotRequest) SetDiskId(v string) *CreateSnapshotRequest {
	s.DiskId = &v
	return s
}

func (s *CreateSnapshotRequest) SetEnsRegionId(v string) *CreateSnapshotRequest {
	s.EnsRegionId = &v
	return s
}

func (s *CreateSnapshotRequest) SetInstanceBillingCycle(v string) *CreateSnapshotRequest {
	s.InstanceBillingCycle = &v
	return s
}

func (s *CreateSnapshotRequest) SetSnapshotName(v string) *CreateSnapshotRequest {
	s.SnapshotName = &v
	return s
}

func (s *CreateSnapshotRequest) Validate() error {
	return dara.Validate(s)
}
