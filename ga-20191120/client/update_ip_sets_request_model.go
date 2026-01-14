// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateIpSetsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIpSets(v []*UpdateIpSetsRequestIpSets) *UpdateIpSetsRequest
	GetIpSets() []*UpdateIpSetsRequestIpSets
	SetRegionId(v string) *UpdateIpSetsRequest
	GetRegionId() *string
}

type UpdateIpSetsRequest struct {
	// The acceleration regions.
	//
	// This parameter is required.
	IpSets []*UpdateIpSetsRequestIpSets `json:"IpSets,omitempty" xml:"IpSets,omitempty" type:"Repeated"`
	// The region ID of the Global Accelerator (GA) instance. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateIpSetsRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateIpSetsRequest) GoString() string {
	return s.String()
}

func (s *UpdateIpSetsRequest) GetIpSets() []*UpdateIpSetsRequestIpSets {
	return s.IpSets
}

func (s *UpdateIpSetsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateIpSetsRequest) SetIpSets(v []*UpdateIpSetsRequestIpSets) *UpdateIpSetsRequest {
	s.IpSets = v
	return s
}

func (s *UpdateIpSetsRequest) SetRegionId(v string) *UpdateIpSetsRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateIpSetsRequest) Validate() error {
	if s.IpSets != nil {
		for _, item := range s.IpSets {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateIpSetsRequestIpSets struct {
	// The new bandwidth that you want to allocate to the acceleration regions. Unit: Mbit/s.
	//
	// You must allocate at least 2 Mbit/s of bandwidth to each acceleration region. You can specify the bandwidth for up to 100 acceleration regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The IDs of the acceleration regions that you want to modify.
	//
	// You can specify the IDs of up to 100 acceleration regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// ips-bp11c9mpphtb1xkds****
	IpSetId *string `json:"IpSetId,omitempty" xml:"IpSetId,omitempty"`
}

func (s UpdateIpSetsRequestIpSets) String() string {
	return dara.Prettify(s)
}

func (s UpdateIpSetsRequestIpSets) GoString() string {
	return s.String()
}

func (s *UpdateIpSetsRequestIpSets) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *UpdateIpSetsRequestIpSets) GetIpSetId() *string {
	return s.IpSetId
}

func (s *UpdateIpSetsRequestIpSets) SetBandwidth(v int32) *UpdateIpSetsRequestIpSets {
	s.Bandwidth = &v
	return s
}

func (s *UpdateIpSetsRequestIpSets) SetIpSetId(v string) *UpdateIpSetsRequestIpSets {
	s.IpSetId = &v
	return s
}

func (s *UpdateIpSetsRequestIpSets) Validate() error {
	return dara.Validate(s)
}
