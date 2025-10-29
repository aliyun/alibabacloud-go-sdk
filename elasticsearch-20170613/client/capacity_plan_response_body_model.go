// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCapacityPlanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CapacityPlanResponseBody
	GetRequestId() *string
	SetResult(v *CapacityPlanResponseBodyResult) *CapacityPlanResponseBody
	GetResult() *CapacityPlanResponseBodyResult
}

type CapacityPlanResponseBody struct {
	// ID of the current request.
	//
	// example:
	//
	// E91B7129-A669-4D9D-A743-F90A0FF1F5EF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returned result of the request.
	Result *CapacityPlanResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s CapacityPlanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CapacityPlanResponseBody) GoString() string {
	return s.String()
}

func (s *CapacityPlanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CapacityPlanResponseBody) GetResult() *CapacityPlanResponseBodyResult {
	return s.Result
}

func (s *CapacityPlanResponseBody) SetRequestId(v string) *CapacityPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *CapacityPlanResponseBody) SetResult(v *CapacityPlanResponseBodyResult) *CapacityPlanResponseBody {
	s.Result = v
	return s
}

func (s *CapacityPlanResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CapacityPlanResponseBodyResult struct {
	// Extended configuration information.
	ExtendConfigs []*CapacityPlanResponseBodyResultExtendConfigs `json:"ExtendConfigs,omitempty" xml:"ExtendConfigs,omitempty" type:"Repeated"`
	// Edition type, with values meaning as follows:
	//
	// - advanced: Enhanced Edition
	//
	// - x-pack: Commercial Edition
	//
	// - community: Community Edition
	//
	// example:
	//
	// advanced
	InstanceCategory *string `json:"InstanceCategory,omitempty" xml:"InstanceCategory,omitempty"`
	// Node information.
	NodeConfigurations []*CapacityPlanResponseBodyResultNodeConfigurations `json:"NodeConfigurations,omitempty" xml:"NodeConfigurations,omitempty" type:"Repeated"`
	// Based on the capacity planning calculation, there is no default value. The meanings of the values are as follows:
	//
	// - true: Represents an oversized cluster, indicating that the number of data nodes calculated by the capacity planning exceeds the threshold of 50.
	//
	// - false: The number of data nodes calculated by the capacity planning is within 50.
	//
	// example:
	//
	// true
	OversizedCluster *bool `json:"OversizedCluster,omitempty" xml:"OversizedCluster,omitempty"`
}

func (s CapacityPlanResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s CapacityPlanResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CapacityPlanResponseBodyResult) GetExtendConfigs() []*CapacityPlanResponseBodyResultExtendConfigs {
	return s.ExtendConfigs
}

func (s *CapacityPlanResponseBodyResult) GetInstanceCategory() *string {
	return s.InstanceCategory
}

func (s *CapacityPlanResponseBodyResult) GetNodeConfigurations() []*CapacityPlanResponseBodyResultNodeConfigurations {
	return s.NodeConfigurations
}

func (s *CapacityPlanResponseBodyResult) GetOversizedCluster() *bool {
	return s.OversizedCluster
}

func (s *CapacityPlanResponseBodyResult) SetExtendConfigs(v []*CapacityPlanResponseBodyResultExtendConfigs) *CapacityPlanResponseBodyResult {
	s.ExtendConfigs = v
	return s
}

func (s *CapacityPlanResponseBodyResult) SetInstanceCategory(v string) *CapacityPlanResponseBodyResult {
	s.InstanceCategory = &v
	return s
}

func (s *CapacityPlanResponseBodyResult) SetNodeConfigurations(v []*CapacityPlanResponseBodyResultNodeConfigurations) *CapacityPlanResponseBodyResult {
	s.NodeConfigurations = v
	return s
}

func (s *CapacityPlanResponseBodyResult) SetOversizedCluster(v bool) *CapacityPlanResponseBodyResult {
	s.OversizedCluster = &v
	return s
}

func (s *CapacityPlanResponseBodyResult) Validate() error {
	if s.ExtendConfigs != nil {
		for _, item := range s.ExtendConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.NodeConfigurations != nil {
		for _, item := range s.NodeConfigurations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CapacityPlanResponseBodyResultExtendConfigs struct {
	// Configuration type, with a single value: sharedDisk.
	//
	// > This extendConfigs attribute may appear when the planned instance type is Advanced.
	//
	// example:
	//
	// sharedDisk
	ConfigType *string `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	// Disk size, in GiB.
	//
	// example:
	//
	// 2048
	Disk *int64 `json:"Disk,omitempty" xml:"Disk,omitempty"`
	// Disk type, with a single value: CPFS_PREMIUM.
	//
	// > This extendConfigs attribute may appear when the planned instance type is Advanced.
	//
	// example:
	//
	// CPFS_PREMIUM
	DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
}

func (s CapacityPlanResponseBodyResultExtendConfigs) String() string {
	return dara.Prettify(s)
}

func (s CapacityPlanResponseBodyResultExtendConfigs) GoString() string {
	return s.String()
}

func (s *CapacityPlanResponseBodyResultExtendConfigs) GetConfigType() *string {
	return s.ConfigType
}

func (s *CapacityPlanResponseBodyResultExtendConfigs) GetDisk() *int64 {
	return s.Disk
}

func (s *CapacityPlanResponseBodyResultExtendConfigs) GetDiskType() *string {
	return s.DiskType
}

func (s *CapacityPlanResponseBodyResultExtendConfigs) SetConfigType(v string) *CapacityPlanResponseBodyResultExtendConfigs {
	s.ConfigType = &v
	return s
}

func (s *CapacityPlanResponseBodyResultExtendConfigs) SetDisk(v int64) *CapacityPlanResponseBodyResultExtendConfigs {
	s.Disk = &v
	return s
}

func (s *CapacityPlanResponseBodyResultExtendConfigs) SetDiskType(v string) *CapacityPlanResponseBodyResultExtendConfigs {
	s.DiskType = &v
	return s
}

func (s *CapacityPlanResponseBodyResultExtendConfigs) Validate() error {
	return dara.Validate(s)
}

type CapacityPlanResponseBodyResultNodeConfigurations struct {
	// Number of nodes.
	//
	// example:
	//
	// 10
	Amount *int64 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// Number of CPUs.
	//
	// example:
	//
	// 1
	Cpu *int64 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// Disk size, in GiB.
	//
	// example:
	//
	// 20
	Disk *int64 `json:"Disk,omitempty" xml:"Disk,omitempty"`
	// Disk type, with meanings as follows:
	//
	// - cloud_essd: ESSD Cloud Disk
	//
	// - cloud_ssd: SSD Cloud Disk
	//
	// - cloud_efficiency: Efficient Cloud Disk
	//
	// - local_ssd: Local SSD Disk
	//
	// - local_efficiency: Local Efficient Disk
	//
	// example:
	//
	// cloud_ssd
	DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	// Specified memory size for the current node role.
	//
	// example:
	//
	// 2
	Memory *int64 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// Node type, with supported types as follows:
	//
	// - WORKER: Data Node
	//
	// - WORKER_WARM: Cold Data Node
	//
	// - MASTER: Dedicated Master Node
	//
	// - KIBANA: Kibana Node
	//
	// - COORDINATING: Coordinator Node
	//
	// - ELASTIC_WORKER: Elastic Node
	//
	// example:
	//
	// WORKER
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
}

func (s CapacityPlanResponseBodyResultNodeConfigurations) String() string {
	return dara.Prettify(s)
}

func (s CapacityPlanResponseBodyResultNodeConfigurations) GoString() string {
	return s.String()
}

func (s *CapacityPlanResponseBodyResultNodeConfigurations) GetAmount() *int64 {
	return s.Amount
}

func (s *CapacityPlanResponseBodyResultNodeConfigurations) GetCpu() *int64 {
	return s.Cpu
}

func (s *CapacityPlanResponseBodyResultNodeConfigurations) GetDisk() *int64 {
	return s.Disk
}

func (s *CapacityPlanResponseBodyResultNodeConfigurations) GetDiskType() *string {
	return s.DiskType
}

func (s *CapacityPlanResponseBodyResultNodeConfigurations) GetMemory() *int64 {
	return s.Memory
}

func (s *CapacityPlanResponseBodyResultNodeConfigurations) GetNodeType() *string {
	return s.NodeType
}

func (s *CapacityPlanResponseBodyResultNodeConfigurations) SetAmount(v int64) *CapacityPlanResponseBodyResultNodeConfigurations {
	s.Amount = &v
	return s
}

func (s *CapacityPlanResponseBodyResultNodeConfigurations) SetCpu(v int64) *CapacityPlanResponseBodyResultNodeConfigurations {
	s.Cpu = &v
	return s
}

func (s *CapacityPlanResponseBodyResultNodeConfigurations) SetDisk(v int64) *CapacityPlanResponseBodyResultNodeConfigurations {
	s.Disk = &v
	return s
}

func (s *CapacityPlanResponseBodyResultNodeConfigurations) SetDiskType(v string) *CapacityPlanResponseBodyResultNodeConfigurations {
	s.DiskType = &v
	return s
}

func (s *CapacityPlanResponseBodyResultNodeConfigurations) SetMemory(v int64) *CapacityPlanResponseBodyResultNodeConfigurations {
	s.Memory = &v
	return s
}

func (s *CapacityPlanResponseBodyResultNodeConfigurations) SetNodeType(v string) *CapacityPlanResponseBodyResultNodeConfigurations {
	s.NodeType = &v
	return s
}

func (s *CapacityPlanResponseBodyResultNodeConfigurations) Validate() error {
	return dara.Validate(s)
}
