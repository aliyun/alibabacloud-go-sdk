// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServerRelatedGlobalAccelerationInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGlobalAccelerationInstances(v *DescribeServerRelatedGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstances) *DescribeServerRelatedGlobalAccelerationInstancesResponseBody
	GetGlobalAccelerationInstances() *DescribeServerRelatedGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstances
	SetRequestId(v string) *DescribeServerRelatedGlobalAccelerationInstancesResponseBody
	GetRequestId() *string
}

type DescribeServerRelatedGlobalAccelerationInstancesResponseBody struct {
	GlobalAccelerationInstances *DescribeServerRelatedGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstances `json:"GlobalAccelerationInstances,omitempty" xml:"GlobalAccelerationInstances,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// A8252014-D8DE-4D85-AF35-AFEXXXXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeServerRelatedGlobalAccelerationInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeServerRelatedGlobalAccelerationInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeServerRelatedGlobalAccelerationInstancesResponseBody) GetGlobalAccelerationInstances() *DescribeServerRelatedGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstances {
	return s.GlobalAccelerationInstances
}

func (s *DescribeServerRelatedGlobalAccelerationInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeServerRelatedGlobalAccelerationInstancesResponseBody) SetGlobalAccelerationInstances(v *DescribeServerRelatedGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstances) *DescribeServerRelatedGlobalAccelerationInstancesResponseBody {
	s.GlobalAccelerationInstances = v
	return s
}

func (s *DescribeServerRelatedGlobalAccelerationInstancesResponseBody) SetRequestId(v string) *DescribeServerRelatedGlobalAccelerationInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeServerRelatedGlobalAccelerationInstancesResponseBody) Validate() error {
	if s.GlobalAccelerationInstances != nil {
		if err := s.GlobalAccelerationInstances.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeServerRelatedGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstances struct {
	GlobalAccelerationInstance []*DescribeServerRelatedGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstance `json:"GlobalAccelerationInstance,omitempty" xml:"GlobalAccelerationInstance,omitempty" type:"Repeated"`
}

func (s DescribeServerRelatedGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstances) String() string {
	return dara.Prettify(s)
}

func (s DescribeServerRelatedGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstances) GoString() string {
	return s.String()
}

func (s *DescribeServerRelatedGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstances) GetGlobalAccelerationInstance() []*DescribeServerRelatedGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstance {
	return s.GlobalAccelerationInstance
}

func (s *DescribeServerRelatedGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstances) SetGlobalAccelerationInstance(v []*DescribeServerRelatedGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstance) *DescribeServerRelatedGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstances {
	s.GlobalAccelerationInstance = v
	return s
}

func (s *DescribeServerRelatedGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstances) Validate() error {
	if s.GlobalAccelerationInstance != nil {
		for _, item := range s.GlobalAccelerationInstance {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeServerRelatedGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstance struct {
	GlobalAccelerationInstanceId *string `json:"GlobalAccelerationInstanceId,omitempty" xml:"GlobalAccelerationInstanceId,omitempty"`
	IpAddress                    *string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
	RegionId                     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ServerIpAddress              *string `json:"ServerIpAddress,omitempty" xml:"ServerIpAddress,omitempty"`
}

func (s DescribeServerRelatedGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstance) String() string {
	return dara.Prettify(s)
}

func (s DescribeServerRelatedGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstance) GoString() string {
	return s.String()
}

func (s *DescribeServerRelatedGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstance) GetGlobalAccelerationInstanceId() *string {
	return s.GlobalAccelerationInstanceId
}

func (s *DescribeServerRelatedGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstance) GetIpAddress() *string {
	return s.IpAddress
}

func (s *DescribeServerRelatedGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstance) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeServerRelatedGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstance) GetServerIpAddress() *string {
	return s.ServerIpAddress
}

func (s *DescribeServerRelatedGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstance) SetGlobalAccelerationInstanceId(v string) *DescribeServerRelatedGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstance {
	s.GlobalAccelerationInstanceId = &v
	return s
}

func (s *DescribeServerRelatedGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstance) SetIpAddress(v string) *DescribeServerRelatedGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstance {
	s.IpAddress = &v
	return s
}

func (s *DescribeServerRelatedGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstance) SetRegionId(v string) *DescribeServerRelatedGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstance {
	s.RegionId = &v
	return s
}

func (s *DescribeServerRelatedGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstance) SetServerIpAddress(v string) *DescribeServerRelatedGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstance {
	s.ServerIpAddress = &v
	return s
}

func (s *DescribeServerRelatedGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstance) Validate() error {
	return dara.Validate(s)
}
