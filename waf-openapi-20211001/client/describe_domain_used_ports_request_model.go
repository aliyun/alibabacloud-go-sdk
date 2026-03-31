// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainUsedPortsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeDomainUsedPortsRequest
	GetInstanceId() *string
	SetRegionId(v string) *DescribeDomainUsedPortsRequest
	GetRegionId() *string
}

type DescribeDomainUsedPortsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// waf_elasticity-cn-0xldbqt****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDomainUsedPortsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainUsedPortsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainUsedPortsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeDomainUsedPortsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDomainUsedPortsRequest) SetInstanceId(v string) *DescribeDomainUsedPortsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDomainUsedPortsRequest) SetRegionId(v string) *DescribeDomainUsedPortsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDomainUsedPortsRequest) Validate() error {
	return dara.Validate(s)
}
