// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGtmRecoveryPlanAvailableConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstances(v *DescribeGtmRecoveryPlanAvailableConfigResponseBodyInstances) *DescribeGtmRecoveryPlanAvailableConfigResponseBody
	GetInstances() *DescribeGtmRecoveryPlanAvailableConfigResponseBodyInstances
	SetRequestId(v string) *DescribeGtmRecoveryPlanAvailableConfigResponseBody
	GetRequestId() *string
}

type DescribeGtmRecoveryPlanAvailableConfigResponseBody struct {
	// The instances.
	Instances *DescribeGtmRecoveryPlanAvailableConfigResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// F8F8EF50-8B7F-4702-B294-97170A423403
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeGtmRecoveryPlanAvailableConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeGtmRecoveryPlanAvailableConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGtmRecoveryPlanAvailableConfigResponseBody) GetInstances() *DescribeGtmRecoveryPlanAvailableConfigResponseBodyInstances {
	return s.Instances
}

func (s *DescribeGtmRecoveryPlanAvailableConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeGtmRecoveryPlanAvailableConfigResponseBody) SetInstances(v *DescribeGtmRecoveryPlanAvailableConfigResponseBodyInstances) *DescribeGtmRecoveryPlanAvailableConfigResponseBody {
	s.Instances = v
	return s
}

func (s *DescribeGtmRecoveryPlanAvailableConfigResponseBody) SetRequestId(v string) *DescribeGtmRecoveryPlanAvailableConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGtmRecoveryPlanAvailableConfigResponseBody) Validate() error {
	if s.Instances != nil {
		if err := s.Instances.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeGtmRecoveryPlanAvailableConfigResponseBodyInstances struct {
	Instance []*DescribeGtmRecoveryPlanAvailableConfigResponseBodyInstancesInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Repeated"`
}

func (s DescribeGtmRecoveryPlanAvailableConfigResponseBodyInstances) String() string {
	return dara.Prettify(s)
}

func (s DescribeGtmRecoveryPlanAvailableConfigResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *DescribeGtmRecoveryPlanAvailableConfigResponseBodyInstances) GetInstance() []*DescribeGtmRecoveryPlanAvailableConfigResponseBodyInstancesInstance {
	return s.Instance
}

func (s *DescribeGtmRecoveryPlanAvailableConfigResponseBodyInstances) SetInstance(v []*DescribeGtmRecoveryPlanAvailableConfigResponseBodyInstancesInstance) *DescribeGtmRecoveryPlanAvailableConfigResponseBodyInstances {
	s.Instance = v
	return s
}

func (s *DescribeGtmRecoveryPlanAvailableConfigResponseBodyInstances) Validate() error {
	if s.Instance != nil {
		for _, item := range s.Instance {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeGtmRecoveryPlanAvailableConfigResponseBodyInstancesInstance struct {
	// The address pools.
	AddrPools *DescribeGtmRecoveryPlanAvailableConfigResponseBodyInstancesInstanceAddrPools `json:"AddrPools,omitempty" xml:"AddrPools,omitempty" type:"Struct"`
	// The instance ID.
	//
	// example:
	//
	// instance-example
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance name.
	//
	// example:
	//
	// instance-name-example
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
}

func (s DescribeGtmRecoveryPlanAvailableConfigResponseBodyInstancesInstance) String() string {
	return dara.Prettify(s)
}

func (s DescribeGtmRecoveryPlanAvailableConfigResponseBodyInstancesInstance) GoString() string {
	return s.String()
}

func (s *DescribeGtmRecoveryPlanAvailableConfigResponseBodyInstancesInstance) GetAddrPools() *DescribeGtmRecoveryPlanAvailableConfigResponseBodyInstancesInstanceAddrPools {
	return s.AddrPools
}

func (s *DescribeGtmRecoveryPlanAvailableConfigResponseBodyInstancesInstance) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeGtmRecoveryPlanAvailableConfigResponseBodyInstancesInstance) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeGtmRecoveryPlanAvailableConfigResponseBodyInstancesInstance) SetAddrPools(v *DescribeGtmRecoveryPlanAvailableConfigResponseBodyInstancesInstanceAddrPools) *DescribeGtmRecoveryPlanAvailableConfigResponseBodyInstancesInstance {
	s.AddrPools = v
	return s
}

func (s *DescribeGtmRecoveryPlanAvailableConfigResponseBodyInstancesInstance) SetInstanceId(v string) *DescribeGtmRecoveryPlanAvailableConfigResponseBodyInstancesInstance {
	s.InstanceId = &v
	return s
}

func (s *DescribeGtmRecoveryPlanAvailableConfigResponseBodyInstancesInstance) SetInstanceName(v string) *DescribeGtmRecoveryPlanAvailableConfigResponseBodyInstancesInstance {
	s.InstanceName = &v
	return s
}

func (s *DescribeGtmRecoveryPlanAvailableConfigResponseBodyInstancesInstance) Validate() error {
	if s.AddrPools != nil {
		if err := s.AddrPools.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeGtmRecoveryPlanAvailableConfigResponseBodyInstancesInstanceAddrPools struct {
	AddrPool []*DescribeGtmRecoveryPlanAvailableConfigResponseBodyInstancesInstanceAddrPoolsAddrPool `json:"AddrPool,omitempty" xml:"AddrPool,omitempty" type:"Repeated"`
}

func (s DescribeGtmRecoveryPlanAvailableConfigResponseBodyInstancesInstanceAddrPools) String() string {
	return dara.Prettify(s)
}

func (s DescribeGtmRecoveryPlanAvailableConfigResponseBodyInstancesInstanceAddrPools) GoString() string {
	return s.String()
}

func (s *DescribeGtmRecoveryPlanAvailableConfigResponseBodyInstancesInstanceAddrPools) GetAddrPool() []*DescribeGtmRecoveryPlanAvailableConfigResponseBodyInstancesInstanceAddrPoolsAddrPool {
	return s.AddrPool
}

func (s *DescribeGtmRecoveryPlanAvailableConfigResponseBodyInstancesInstanceAddrPools) SetAddrPool(v []*DescribeGtmRecoveryPlanAvailableConfigResponseBodyInstancesInstanceAddrPoolsAddrPool) *DescribeGtmRecoveryPlanAvailableConfigResponseBodyInstancesInstanceAddrPools {
	s.AddrPool = v
	return s
}

func (s *DescribeGtmRecoveryPlanAvailableConfigResponseBodyInstancesInstanceAddrPools) Validate() error {
	if s.AddrPool != nil {
		for _, item := range s.AddrPool {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeGtmRecoveryPlanAvailableConfigResponseBodyInstancesInstanceAddrPoolsAddrPool struct {
	// The address pool ID.
	AddrPoolId *string `json:"AddrPoolId,omitempty" xml:"AddrPoolId,omitempty"`
	// The name of the address pool.
	//
	// example:
	//
	// hra0i9
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeGtmRecoveryPlanAvailableConfigResponseBodyInstancesInstanceAddrPoolsAddrPool) String() string {
	return dara.Prettify(s)
}

func (s DescribeGtmRecoveryPlanAvailableConfigResponseBodyInstancesInstanceAddrPoolsAddrPool) GoString() string {
	return s.String()
}

func (s *DescribeGtmRecoveryPlanAvailableConfigResponseBodyInstancesInstanceAddrPoolsAddrPool) GetAddrPoolId() *string {
	return s.AddrPoolId
}

func (s *DescribeGtmRecoveryPlanAvailableConfigResponseBodyInstancesInstanceAddrPoolsAddrPool) GetName() *string {
	return s.Name
}

func (s *DescribeGtmRecoveryPlanAvailableConfigResponseBodyInstancesInstanceAddrPoolsAddrPool) SetAddrPoolId(v string) *DescribeGtmRecoveryPlanAvailableConfigResponseBodyInstancesInstanceAddrPoolsAddrPool {
	s.AddrPoolId = &v
	return s
}

func (s *DescribeGtmRecoveryPlanAvailableConfigResponseBodyInstancesInstanceAddrPoolsAddrPool) SetName(v string) *DescribeGtmRecoveryPlanAvailableConfigResponseBodyInstancesInstanceAddrPoolsAddrPool {
	s.Name = &v
	return s
}

func (s *DescribeGtmRecoveryPlanAvailableConfigResponseBodyInstancesInstanceAddrPoolsAddrPool) Validate() error {
	return dara.Validate(s)
}
