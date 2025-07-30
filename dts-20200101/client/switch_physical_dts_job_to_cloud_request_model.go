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
	// Migration, synchronization, or subscription instance ID.
	//
	// example:
	//
	// dtsl3m1213ye7l****
	DtsInstanceId *string `json:"DtsInstanceId,omitempty" xml:"DtsInstanceId,omitempty"`
	// Data migration or synchronization instance ID, which can be queried by calling the **describedtsjobs*	- interface.
	//
	// example:
	//
	// l5512es7w15****
	DtsJobId *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	// Region ID. Pass this parameter to specify the region where the instance is located. For more details, see the list of supported regions.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Resource group ID.
	//
	// example:
	//
	// rg-aekz4us4iruleja
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Synchronization direction, values: - **Forward**: Forward. - **Reverse**: Reverse.
	//
	// > - The default value is **Forward**. - **Reverse*	- can only be passed when the topology of the data synchronization instance is bidirectional, to release the reverse synchronization link.
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
