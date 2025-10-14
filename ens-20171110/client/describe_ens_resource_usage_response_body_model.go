// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEnsResourceUsageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEnsResourceUsage(v []*DescribeEnsResourceUsageResponseBodyEnsResourceUsage) *DescribeEnsResourceUsageResponseBody
	GetEnsResourceUsage() []*DescribeEnsResourceUsageResponseBodyEnsResourceUsage
	SetRequestId(v string) *DescribeEnsResourceUsageResponseBody
	GetRequestId() *string
}

type DescribeEnsResourceUsageResponseBody struct {
	// The resource usage data.
	EnsResourceUsage []*DescribeEnsResourceUsageResponseBodyEnsResourceUsage `json:"EnsResourceUsage,omitempty" xml:"EnsResourceUsage,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E	 Request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeEnsResourceUsageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnsResourceUsageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEnsResourceUsageResponseBody) GetEnsResourceUsage() []*DescribeEnsResourceUsageResponseBodyEnsResourceUsage {
	return s.EnsResourceUsage
}

func (s *DescribeEnsResourceUsageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEnsResourceUsageResponseBody) SetEnsResourceUsage(v []*DescribeEnsResourceUsageResponseBodyEnsResourceUsage) *DescribeEnsResourceUsageResponseBody {
	s.EnsResourceUsage = v
	return s
}

func (s *DescribeEnsResourceUsageResponseBody) SetRequestId(v string) *DescribeEnsResourceUsageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEnsResourceUsageResponseBody) Validate() error {
	if s.EnsResourceUsage != nil {
		for _, item := range s.EnsResourceUsage {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeEnsResourceUsageResponseBodyEnsResourceUsage struct {
	// The number of edge services. This parameter is available only when you set the ServiceType parameter to 2.
	//
	// example:
	//
	// 2
	ComputeResourceCount *int32 `json:"ComputeResourceCount,omitempty" xml:"ComputeResourceCount,omitempty"`
	// The CPU usage. Unit: cores.
	//
	// example:
	//
	// 2
	CpuSum *int64 `json:"CpuSum,omitempty" xml:"CpuSum,omitempty"`
	// The number of data disks.
	//
	// example:
	//
	// 4
	DiskCount *int32 `json:"DiskCount,omitempty" xml:"DiskCount,omitempty"`
	// The number of stopped VMs.
	//
	// example:
	//
	// 7
	DownCount *int32 `json:"DownCount,omitempty" xml:"DownCount,omitempty"`
	// The number of expired VM instances.
	//
	// example:
	//
	// 1
	ExpiredCount *int32 `json:"ExpiredCount,omitempty" xml:"ExpiredCount,omitempty"`
	// The number of VM instances that are about to expire.
	//
	// example:
	//
	// 0
	ExpiringCount *int32 `json:"ExpiringCount,omitempty" xml:"ExpiringCount,omitempty"`
	// The number of GPUs.
	//
	// example:
	//
	// 6
	GpuSum *int64 `json:"GpuSum,omitempty" xml:"GpuSum,omitempty"`
	// The number of instances.
	//
	// example:
	//
	// 2
	InstanceCount *int32 `json:"InstanceCount,omitempty" xml:"InstanceCount,omitempty"`
	// The number of running instances.
	//
	// example:
	//
	// 19
	RunningCount *int32 `json:"RunningCount,omitempty" xml:"RunningCount,omitempty"`
	// The type of the service. Valid values:
	//
	// 	- 1: subscription instance.
	//
	// 	- 2: edge service instance.
	//
	// 	- 3: pay-as-you-go instance.
	//
	// example:
	//
	// 1
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// The total disk size.
	//
	// example:
	//
	// 5000
	StorageSum *int64 `json:"StorageSum,omitempty" xml:"StorageSum,omitempty"`
}

func (s DescribeEnsResourceUsageResponseBodyEnsResourceUsage) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnsResourceUsageResponseBodyEnsResourceUsage) GoString() string {
	return s.String()
}

func (s *DescribeEnsResourceUsageResponseBodyEnsResourceUsage) GetComputeResourceCount() *int32 {
	return s.ComputeResourceCount
}

func (s *DescribeEnsResourceUsageResponseBodyEnsResourceUsage) GetCpuSum() *int64 {
	return s.CpuSum
}

func (s *DescribeEnsResourceUsageResponseBodyEnsResourceUsage) GetDiskCount() *int32 {
	return s.DiskCount
}

func (s *DescribeEnsResourceUsageResponseBodyEnsResourceUsage) GetDownCount() *int32 {
	return s.DownCount
}

func (s *DescribeEnsResourceUsageResponseBodyEnsResourceUsage) GetExpiredCount() *int32 {
	return s.ExpiredCount
}

func (s *DescribeEnsResourceUsageResponseBodyEnsResourceUsage) GetExpiringCount() *int32 {
	return s.ExpiringCount
}

func (s *DescribeEnsResourceUsageResponseBodyEnsResourceUsage) GetGpuSum() *int64 {
	return s.GpuSum
}

func (s *DescribeEnsResourceUsageResponseBodyEnsResourceUsage) GetInstanceCount() *int32 {
	return s.InstanceCount
}

func (s *DescribeEnsResourceUsageResponseBodyEnsResourceUsage) GetRunningCount() *int32 {
	return s.RunningCount
}

func (s *DescribeEnsResourceUsageResponseBodyEnsResourceUsage) GetServiceType() *string {
	return s.ServiceType
}

func (s *DescribeEnsResourceUsageResponseBodyEnsResourceUsage) GetStorageSum() *int64 {
	return s.StorageSum
}

func (s *DescribeEnsResourceUsageResponseBodyEnsResourceUsage) SetComputeResourceCount(v int32) *DescribeEnsResourceUsageResponseBodyEnsResourceUsage {
	s.ComputeResourceCount = &v
	return s
}

func (s *DescribeEnsResourceUsageResponseBodyEnsResourceUsage) SetCpuSum(v int64) *DescribeEnsResourceUsageResponseBodyEnsResourceUsage {
	s.CpuSum = &v
	return s
}

func (s *DescribeEnsResourceUsageResponseBodyEnsResourceUsage) SetDiskCount(v int32) *DescribeEnsResourceUsageResponseBodyEnsResourceUsage {
	s.DiskCount = &v
	return s
}

func (s *DescribeEnsResourceUsageResponseBodyEnsResourceUsage) SetDownCount(v int32) *DescribeEnsResourceUsageResponseBodyEnsResourceUsage {
	s.DownCount = &v
	return s
}

func (s *DescribeEnsResourceUsageResponseBodyEnsResourceUsage) SetExpiredCount(v int32) *DescribeEnsResourceUsageResponseBodyEnsResourceUsage {
	s.ExpiredCount = &v
	return s
}

func (s *DescribeEnsResourceUsageResponseBodyEnsResourceUsage) SetExpiringCount(v int32) *DescribeEnsResourceUsageResponseBodyEnsResourceUsage {
	s.ExpiringCount = &v
	return s
}

func (s *DescribeEnsResourceUsageResponseBodyEnsResourceUsage) SetGpuSum(v int64) *DescribeEnsResourceUsageResponseBodyEnsResourceUsage {
	s.GpuSum = &v
	return s
}

func (s *DescribeEnsResourceUsageResponseBodyEnsResourceUsage) SetInstanceCount(v int32) *DescribeEnsResourceUsageResponseBodyEnsResourceUsage {
	s.InstanceCount = &v
	return s
}

func (s *DescribeEnsResourceUsageResponseBodyEnsResourceUsage) SetRunningCount(v int32) *DescribeEnsResourceUsageResponseBodyEnsResourceUsage {
	s.RunningCount = &v
	return s
}

func (s *DescribeEnsResourceUsageResponseBodyEnsResourceUsage) SetServiceType(v string) *DescribeEnsResourceUsageResponseBodyEnsResourceUsage {
	s.ServiceType = &v
	return s
}

func (s *DescribeEnsResourceUsageResponseBodyEnsResourceUsage) SetStorageSum(v int64) *DescribeEnsResourceUsageResponseBodyEnsResourceUsage {
	s.StorageSum = &v
	return s
}

func (s *DescribeEnsResourceUsageResponseBodyEnsResourceUsage) Validate() error {
	return dara.Validate(s)
}
