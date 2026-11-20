// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWarehouseDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetWarehouseDetailResponseBody
	GetRequestId() *string
	SetWarehouseDetail(v *GetWarehouseDetailResponseBodyWarehouseDetail) *GetWarehouseDetailResponseBody
	GetWarehouseDetail() *GetWarehouseDetailResponseBodyWarehouseDetail
}

type GetWarehouseDetailResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// D3AE84AB-0873-5FC7-A4C4-8CF869D2FA70
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The return value.
	WarehouseDetail *GetWarehouseDetailResponseBodyWarehouseDetail `json:"WarehouseDetail,omitempty" xml:"WarehouseDetail,omitempty" type:"Struct"`
}

func (s GetWarehouseDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetWarehouseDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetWarehouseDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetWarehouseDetailResponseBody) GetWarehouseDetail() *GetWarehouseDetailResponseBodyWarehouseDetail {
	return s.WarehouseDetail
}

func (s *GetWarehouseDetailResponseBody) SetRequestId(v string) *GetWarehouseDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWarehouseDetailResponseBody) SetWarehouseDetail(v *GetWarehouseDetailResponseBodyWarehouseDetail) *GetWarehouseDetailResponseBody {
	s.WarehouseDetail = v
	return s
}

func (s *GetWarehouseDetailResponseBody) Validate() error {
	if s.WarehouseDetail != nil {
		if err := s.WarehouseDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetWarehouseDetailResponseBodyWarehouseDetail struct {
	// The total auto-scaling elastic compute resources of the compute group.
	//
	// example:
	//
	// 32
	AutoElasticCpu *string `json:"AutoElasticCpu,omitempty" xml:"AutoElasticCpu,omitempty"`
	// The unallocated compute resources. The remaining unallocated compute resources for compute groups in the instance.
	//
	// example:
	//
	// 32
	RemainingCpu *string `json:"RemainingCpu,omitempty" xml:"RemainingCpu,omitempty"`
	// The reserved compute resources. The total compute resources of all running compute groups in the instance cannot exceed the reserved compute resources of the compute group.
	//
	// example:
	//
	// 64
	ReservedCpu *string `json:"ReservedCpu,omitempty" xml:"ReservedCpu,omitempty"`
	// The total time-based elastic compute resources of the compute group.
	//
	// example:
	//
	// 64
	TimedElasticCpu *string `json:"TimedElasticCpu,omitempty" xml:"TimedElasticCpu,omitempty"`
	// The warehouse list.
	WarehouseList []*GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList `json:"WarehouseList,omitempty" xml:"WarehouseList,omitempty" type:"Repeated"`
}

func (s GetWarehouseDetailResponseBodyWarehouseDetail) String() string {
	return dara.Prettify(s)
}

func (s GetWarehouseDetailResponseBodyWarehouseDetail) GoString() string {
	return s.String()
}

func (s *GetWarehouseDetailResponseBodyWarehouseDetail) GetAutoElasticCpu() *string {
	return s.AutoElasticCpu
}

func (s *GetWarehouseDetailResponseBodyWarehouseDetail) GetRemainingCpu() *string {
	return s.RemainingCpu
}

func (s *GetWarehouseDetailResponseBodyWarehouseDetail) GetReservedCpu() *string {
	return s.ReservedCpu
}

func (s *GetWarehouseDetailResponseBodyWarehouseDetail) GetTimedElasticCpu() *string {
	return s.TimedElasticCpu
}

func (s *GetWarehouseDetailResponseBodyWarehouseDetail) GetWarehouseList() []*GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList {
	return s.WarehouseList
}

func (s *GetWarehouseDetailResponseBodyWarehouseDetail) SetAutoElasticCpu(v string) *GetWarehouseDetailResponseBodyWarehouseDetail {
	s.AutoElasticCpu = &v
	return s
}

func (s *GetWarehouseDetailResponseBodyWarehouseDetail) SetRemainingCpu(v string) *GetWarehouseDetailResponseBodyWarehouseDetail {
	s.RemainingCpu = &v
	return s
}

func (s *GetWarehouseDetailResponseBodyWarehouseDetail) SetReservedCpu(v string) *GetWarehouseDetailResponseBodyWarehouseDetail {
	s.ReservedCpu = &v
	return s
}

func (s *GetWarehouseDetailResponseBodyWarehouseDetail) SetTimedElasticCpu(v string) *GetWarehouseDetailResponseBodyWarehouseDetail {
	s.TimedElasticCpu = &v
	return s
}

func (s *GetWarehouseDetailResponseBodyWarehouseDetail) SetWarehouseList(v []*GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList) *GetWarehouseDetailResponseBodyWarehouseDetail {
	s.WarehouseList = v
	return s
}

func (s *GetWarehouseDetailResponseBodyWarehouseDetail) Validate() error {
	if s.WarehouseList != nil {
		for _, item := range s.WarehouseList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList struct {
	// The auto-scaling type.
	//
	// example:
	//
	// Standard
	AutoScaleType *string `json:"AutoScaleType,omitempty" xml:"AutoScaleType,omitempty"`
	// The number of clusters.
	//
	// example:
	//
	// 2
	ClusterCount *string `json:"ClusterCount,omitempty" xml:"ClusterCount,omitempty"`
	// The cluster specifications.
	//
	// example:
	//
	// 32
	ClusterCpu *string `json:"ClusterCpu,omitempty" xml:"ClusterCpu,omitempty"`
	// cpu
	//
	// example:
	//
	// 32
	Cpu *int64 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// Indicates whether the compute group is the default compute group.
	//
	// example:
	//
	// true
	DefaultWarehouse *bool `json:"DefaultWarehouse,omitempty" xml:"DefaultWarehouse,omitempty"`
	// The elastic compute resources.
	//
	// example:
	//
	// 32
	ElasticCpu *int64 `json:"ElasticCpu,omitempty" xml:"ElasticCpu,omitempty"`
	// The elastic type.
	//
	// example:
	//
	// auto
	ElasticType *string `json:"ElasticType,omitempty" xml:"ElasticType,omitempty"`
	// id
	//
	// example:
	//
	// 2
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The number of reserved clusters.
	//
	// example:
	//
	// 2
	InitClusterCount *string `json:"InitClusterCount,omitempty" xml:"InitClusterCount,omitempty"`
	// The maximum number of clusters.
	//
	// example:
	//
	// 4
	MaxClusterCount *string `json:"MaxClusterCount,omitempty" xml:"MaxClusterCount,omitempty"`
	// The memory.
	//
	// example:
	//
	// 128
	Mem *int64 `json:"Mem,omitempty" xml:"Mem,omitempty"`
	// The name of the compute group.
	//
	// example:
	//
	// MyWarehouse
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The number of compute nodes.
	//
	// example:
	//
	// 2
	NodeCount *int64 `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	// The rebalance status of the compute group.
	//
	// example:
	//
	// DONE
	RebalanceStatus *string `json:"RebalanceStatus,omitempty" xml:"RebalanceStatus,omitempty"`
	// The status.
	//
	// example:
	//
	// kRunning
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList) String() string {
	return dara.Prettify(s)
}

func (s GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList) GoString() string {
	return s.String()
}

func (s *GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList) GetAutoScaleType() *string {
	return s.AutoScaleType
}

func (s *GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList) GetClusterCount() *string {
	return s.ClusterCount
}

func (s *GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList) GetClusterCpu() *string {
	return s.ClusterCpu
}

func (s *GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList) GetCpu() *int64 {
	return s.Cpu
}

func (s *GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList) GetDefaultWarehouse() *bool {
	return s.DefaultWarehouse
}

func (s *GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList) GetElasticCpu() *int64 {
	return s.ElasticCpu
}

func (s *GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList) GetElasticType() *string {
	return s.ElasticType
}

func (s *GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList) GetId() *int64 {
	return s.Id
}

func (s *GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList) GetInitClusterCount() *string {
	return s.InitClusterCount
}

func (s *GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList) GetMaxClusterCount() *string {
	return s.MaxClusterCount
}

func (s *GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList) GetMem() *int64 {
	return s.Mem
}

func (s *GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList) GetName() *string {
	return s.Name
}

func (s *GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList) GetNodeCount() *int64 {
	return s.NodeCount
}

func (s *GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList) GetRebalanceStatus() *string {
	return s.RebalanceStatus
}

func (s *GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList) GetStatus() *string {
	return s.Status
}

func (s *GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList) SetAutoScaleType(v string) *GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList {
	s.AutoScaleType = &v
	return s
}

func (s *GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList) SetClusterCount(v string) *GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList {
	s.ClusterCount = &v
	return s
}

func (s *GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList) SetClusterCpu(v string) *GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList {
	s.ClusterCpu = &v
	return s
}

func (s *GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList) SetCpu(v int64) *GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList {
	s.Cpu = &v
	return s
}

func (s *GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList) SetDefaultWarehouse(v bool) *GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList {
	s.DefaultWarehouse = &v
	return s
}

func (s *GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList) SetElasticCpu(v int64) *GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList {
	s.ElasticCpu = &v
	return s
}

func (s *GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList) SetElasticType(v string) *GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList {
	s.ElasticType = &v
	return s
}

func (s *GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList) SetId(v int64) *GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList {
	s.Id = &v
	return s
}

func (s *GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList) SetInitClusterCount(v string) *GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList {
	s.InitClusterCount = &v
	return s
}

func (s *GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList) SetMaxClusterCount(v string) *GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList {
	s.MaxClusterCount = &v
	return s
}

func (s *GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList) SetMem(v int64) *GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList {
	s.Mem = &v
	return s
}

func (s *GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList) SetName(v string) *GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList {
	s.Name = &v
	return s
}

func (s *GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList) SetNodeCount(v int64) *GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList {
	s.NodeCount = &v
	return s
}

func (s *GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList) SetRebalanceStatus(v string) *GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList {
	s.RebalanceStatus = &v
	return s
}

func (s *GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList) SetStatus(v string) *GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList {
	s.Status = &v
	return s
}

func (s *GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList) Validate() error {
	return dara.Validate(s)
}
