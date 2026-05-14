// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceModelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceModelList(v []*ListInstanceModelResponseBodyInstanceModelList) *ListInstanceModelResponseBody
	GetInstanceModelList() []*ListInstanceModelResponseBodyInstanceModelList
	SetRequestId(v string) *ListInstanceModelResponseBody
	GetRequestId() *string
}

type ListInstanceModelResponseBody struct {
	InstanceModelList []*ListInstanceModelResponseBodyInstanceModelList `json:"instanceModelList,omitempty" xml:"instanceModelList,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// 2C2ECDC1-FBAD-14A5-AA4A-96BC787FBDBC
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListInstanceModelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceModelResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstanceModelResponseBody) GetInstanceModelList() []*ListInstanceModelResponseBodyInstanceModelList {
	return s.InstanceModelList
}

func (s *ListInstanceModelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInstanceModelResponseBody) SetInstanceModelList(v []*ListInstanceModelResponseBodyInstanceModelList) *ListInstanceModelResponseBody {
	s.InstanceModelList = v
	return s
}

func (s *ListInstanceModelResponseBody) SetRequestId(v string) *ListInstanceModelResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstanceModelResponseBody) Validate() error {
	if s.InstanceModelList != nil {
		for _, item := range s.InstanceModelList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListInstanceModelResponseBodyInstanceModelList struct {
	// example:
	//
	// hologram_aicombo_public_cn-77xxx
	AiInstanceId *string `json:"aiInstanceId,omitempty" xml:"aiInstanceId,omitempty"`
	// example:
	//
	// small-8core-30G-24G
	AiSpec      *string `json:"aiSpec,omitempty" xml:"aiSpec,omitempty"`
	AutoRenewal *bool   `json:"autoRenewal,omitempty" xml:"autoRenewal,omitempty"`
	// example:
	//
	// PostPaid
	ChargeType *string `json:"chargeType,omitempty" xml:"chargeType,omitempty"`
	// example:
	//
	// hologram_aipostpay_public_cn
	CommodityCode *string `json:"commodityCode,omitempty" xml:"commodityCode,omitempty"`
	Cpu           *int64  `json:"cpu,omitempty" xml:"cpu,omitempty"`
	CpuUsed       *int64  `json:"cpuUsed,omitempty" xml:"cpuUsed,omitempty"`
	// example:
	//
	// 2026-01-28T07:44:27.535Z
	ExpirationTime *string `json:"expirationTime,omitempty" xml:"expirationTime,omitempty"`
	Gpu            *int64  `json:"gpu,omitempty" xml:"gpu,omitempty"`
	GpuMemory      *int64  `json:"gpuMemory,omitempty" xml:"gpuMemory,omitempty"`
	GpuMemoryUsed  *int64  `json:"gpuMemoryUsed,omitempty" xml:"gpuMemoryUsed,omitempty"`
	GpuUsed        *int64  `json:"gpuUsed,omitempty" xml:"gpuUsed,omitempty"`
	// example:
	//
	// hgpostcn-cn-yi34hlzdx003
	HoloInstanceId   *string `json:"holoInstanceId,omitempty" xml:"holoInstanceId,omitempty"`
	HoloInstanceName *string `json:"holoInstanceName,omitempty" xml:"holoInstanceName,omitempty"`
	Memory           *int64  `json:"memory,omitempty" xml:"memory,omitempty"`
	MemoryUsed       *int64  `json:"memoryUsed,omitempty" xml:"memoryUsed,omitempty"`
	NodeCount        *int64  `json:"nodeCount,omitempty" xml:"nodeCount,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// example:
	//
	// small
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	// example:
	//
	// ResourceReady
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListInstanceModelResponseBodyInstanceModelList) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceModelResponseBodyInstanceModelList) GoString() string {
	return s.String()
}

func (s *ListInstanceModelResponseBodyInstanceModelList) GetAiInstanceId() *string {
	return s.AiInstanceId
}

func (s *ListInstanceModelResponseBodyInstanceModelList) GetAiSpec() *string {
	return s.AiSpec
}

func (s *ListInstanceModelResponseBodyInstanceModelList) GetAutoRenewal() *bool {
	return s.AutoRenewal
}

func (s *ListInstanceModelResponseBodyInstanceModelList) GetChargeType() *string {
	return s.ChargeType
}

func (s *ListInstanceModelResponseBodyInstanceModelList) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *ListInstanceModelResponseBodyInstanceModelList) GetCpu() *int64 {
	return s.Cpu
}

func (s *ListInstanceModelResponseBodyInstanceModelList) GetCpuUsed() *int64 {
	return s.CpuUsed
}

func (s *ListInstanceModelResponseBodyInstanceModelList) GetExpirationTime() *string {
	return s.ExpirationTime
}

func (s *ListInstanceModelResponseBodyInstanceModelList) GetGpu() *int64 {
	return s.Gpu
}

func (s *ListInstanceModelResponseBodyInstanceModelList) GetGpuMemory() *int64 {
	return s.GpuMemory
}

func (s *ListInstanceModelResponseBodyInstanceModelList) GetGpuMemoryUsed() *int64 {
	return s.GpuMemoryUsed
}

func (s *ListInstanceModelResponseBodyInstanceModelList) GetGpuUsed() *int64 {
	return s.GpuUsed
}

func (s *ListInstanceModelResponseBodyInstanceModelList) GetHoloInstanceId() *string {
	return s.HoloInstanceId
}

func (s *ListInstanceModelResponseBodyInstanceModelList) GetHoloInstanceName() *string {
	return s.HoloInstanceName
}

func (s *ListInstanceModelResponseBodyInstanceModelList) GetMemory() *int64 {
	return s.Memory
}

func (s *ListInstanceModelResponseBodyInstanceModelList) GetMemoryUsed() *int64 {
	return s.MemoryUsed
}

func (s *ListInstanceModelResponseBodyInstanceModelList) GetNodeCount() *int64 {
	return s.NodeCount
}

func (s *ListInstanceModelResponseBodyInstanceModelList) GetRegionId() *string {
	return s.RegionId
}

func (s *ListInstanceModelResponseBodyInstanceModelList) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListInstanceModelResponseBodyInstanceModelList) GetStatus() *string {
	return s.Status
}

func (s *ListInstanceModelResponseBodyInstanceModelList) SetAiInstanceId(v string) *ListInstanceModelResponseBodyInstanceModelList {
	s.AiInstanceId = &v
	return s
}

func (s *ListInstanceModelResponseBodyInstanceModelList) SetAiSpec(v string) *ListInstanceModelResponseBodyInstanceModelList {
	s.AiSpec = &v
	return s
}

func (s *ListInstanceModelResponseBodyInstanceModelList) SetAutoRenewal(v bool) *ListInstanceModelResponseBodyInstanceModelList {
	s.AutoRenewal = &v
	return s
}

func (s *ListInstanceModelResponseBodyInstanceModelList) SetChargeType(v string) *ListInstanceModelResponseBodyInstanceModelList {
	s.ChargeType = &v
	return s
}

func (s *ListInstanceModelResponseBodyInstanceModelList) SetCommodityCode(v string) *ListInstanceModelResponseBodyInstanceModelList {
	s.CommodityCode = &v
	return s
}

func (s *ListInstanceModelResponseBodyInstanceModelList) SetCpu(v int64) *ListInstanceModelResponseBodyInstanceModelList {
	s.Cpu = &v
	return s
}

func (s *ListInstanceModelResponseBodyInstanceModelList) SetCpuUsed(v int64) *ListInstanceModelResponseBodyInstanceModelList {
	s.CpuUsed = &v
	return s
}

func (s *ListInstanceModelResponseBodyInstanceModelList) SetExpirationTime(v string) *ListInstanceModelResponseBodyInstanceModelList {
	s.ExpirationTime = &v
	return s
}

func (s *ListInstanceModelResponseBodyInstanceModelList) SetGpu(v int64) *ListInstanceModelResponseBodyInstanceModelList {
	s.Gpu = &v
	return s
}

func (s *ListInstanceModelResponseBodyInstanceModelList) SetGpuMemory(v int64) *ListInstanceModelResponseBodyInstanceModelList {
	s.GpuMemory = &v
	return s
}

func (s *ListInstanceModelResponseBodyInstanceModelList) SetGpuMemoryUsed(v int64) *ListInstanceModelResponseBodyInstanceModelList {
	s.GpuMemoryUsed = &v
	return s
}

func (s *ListInstanceModelResponseBodyInstanceModelList) SetGpuUsed(v int64) *ListInstanceModelResponseBodyInstanceModelList {
	s.GpuUsed = &v
	return s
}

func (s *ListInstanceModelResponseBodyInstanceModelList) SetHoloInstanceId(v string) *ListInstanceModelResponseBodyInstanceModelList {
	s.HoloInstanceId = &v
	return s
}

func (s *ListInstanceModelResponseBodyInstanceModelList) SetHoloInstanceName(v string) *ListInstanceModelResponseBodyInstanceModelList {
	s.HoloInstanceName = &v
	return s
}

func (s *ListInstanceModelResponseBodyInstanceModelList) SetMemory(v int64) *ListInstanceModelResponseBodyInstanceModelList {
	s.Memory = &v
	return s
}

func (s *ListInstanceModelResponseBodyInstanceModelList) SetMemoryUsed(v int64) *ListInstanceModelResponseBodyInstanceModelList {
	s.MemoryUsed = &v
	return s
}

func (s *ListInstanceModelResponseBodyInstanceModelList) SetNodeCount(v int64) *ListInstanceModelResponseBodyInstanceModelList {
	s.NodeCount = &v
	return s
}

func (s *ListInstanceModelResponseBodyInstanceModelList) SetRegionId(v string) *ListInstanceModelResponseBodyInstanceModelList {
	s.RegionId = &v
	return s
}

func (s *ListInstanceModelResponseBodyInstanceModelList) SetResourceType(v string) *ListInstanceModelResponseBodyInstanceModelList {
	s.ResourceType = &v
	return s
}

func (s *ListInstanceModelResponseBodyInstanceModelList) SetStatus(v string) *ListInstanceModelResponseBodyInstanceModelList {
	s.Status = &v
	return s
}

func (s *ListInstanceModelResponseBodyInstanceModelList) Validate() error {
	return dara.Validate(s)
}
