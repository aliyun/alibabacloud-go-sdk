// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchPhysicalDtsJobToCloudRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDtsInstanceId(v string) *SwitchPhysicalDtsJobToCloudRequest
	GetDtsInstanceId() *string
	SetDtsJobId(v string) *SwitchPhysicalDtsJobToCloudRequest
	GetDtsJobId() *string
	SetRegionId(v string) *SwitchPhysicalDtsJobToCloudRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *SwitchPhysicalDtsJobToCloudRequest
	GetResourceGroupId() *string
	SetSynchronizationDirection(v string) *SwitchPhysicalDtsJobToCloudRequest
	GetSynchronizationDirection() *string
}

type SwitchPhysicalDtsJobToCloudRequest struct {
	// The ID of the migration, synchronization, or change tracking instance.
	//
	// example:
	//
	// dtsl3m1213ye7l****
	DtsInstanceId *string `json:"DtsInstanceId,omitempty" xml:"DtsInstanceId,omitempty"`
	// The ID of the data migration or synchronization task. You can call the **DescribeDtsJobs*	- operation to query the ID.
	//
	// example:
	//
	// l5512es7w15****
	DtsJobId *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	// The region ID. Specify this parameter to indicate the region where the instance resides. For more information, see the list of supported regions.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-aekz4us4iruleja
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The synchronization direction. Valid values:
	//
	// - **Forward**: forward.
	//
	// - **Reverse**: reverse.
	//
	// > - Default value: **Forward**.
	//
	// - You can set this parameter to **Reverse*	- to release the reverse synchronization link only when the topology of the data synchronization instance is two-way synchronization.
	//
	// example:
	//
	// Forward
	SynchronizationDirection *string `json:"SynchronizationDirection,omitempty" xml:"SynchronizationDirection,omitempty"`
}

func (s SwitchPhysicalDtsJobToCloudRequest) String() string {
	return dara.Prettify(s)
}

func (s SwitchPhysicalDtsJobToCloudRequest) GoString() string {
	return s.String()
}

func (s *SwitchPhysicalDtsJobToCloudRequest) GetDtsInstanceId() *string {
	return s.DtsInstanceId
}

func (s *SwitchPhysicalDtsJobToCloudRequest) GetDtsJobId() *string {
	return s.DtsJobId
}

func (s *SwitchPhysicalDtsJobToCloudRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SwitchPhysicalDtsJobToCloudRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *SwitchPhysicalDtsJobToCloudRequest) GetSynchronizationDirection() *string {
	return s.SynchronizationDirection
}

func (s *SwitchPhysicalDtsJobToCloudRequest) SetDtsInstanceId(v string) *SwitchPhysicalDtsJobToCloudRequest {
	s.DtsInstanceId = &v
	return s
}

func (s *SwitchPhysicalDtsJobToCloudRequest) SetDtsJobId(v string) *SwitchPhysicalDtsJobToCloudRequest {
	s.DtsJobId = &v
	return s
}

func (s *SwitchPhysicalDtsJobToCloudRequest) SetRegionId(v string) *SwitchPhysicalDtsJobToCloudRequest {
	s.RegionId = &v
	return s
}

func (s *SwitchPhysicalDtsJobToCloudRequest) SetResourceGroupId(v string) *SwitchPhysicalDtsJobToCloudRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *SwitchPhysicalDtsJobToCloudRequest) SetSynchronizationDirection(v string) *SwitchPhysicalDtsJobToCloudRequest {
	s.SynchronizationDirection = &v
	return s
}

func (s *SwitchPhysicalDtsJobToCloudRequest) Validate() error {
	return dara.Validate(s)
}
