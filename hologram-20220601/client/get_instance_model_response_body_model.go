// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceModelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAiInstanceId(v string) *GetInstanceModelResponseBody
	GetAiInstanceId() *string
	SetAutoRenewal(v bool) *GetInstanceModelResponseBody
	GetAutoRenewal() *bool
	SetChargeType(v string) *GetInstanceModelResponseBody
	GetChargeType() *string
	SetCommodityCode(v string) *GetInstanceModelResponseBody
	GetCommodityCode() *string
	SetCpu(v int64) *GetInstanceModelResponseBody
	GetCpu() *int64
	SetCpuUsed(v int64) *GetInstanceModelResponseBody
	GetCpuUsed() *int64
	SetExpirationTime(v string) *GetInstanceModelResponseBody
	GetExpirationTime() *string
	SetGpu(v int64) *GetInstanceModelResponseBody
	GetGpu() *int64
	SetGpuMemory(v int64) *GetInstanceModelResponseBody
	GetGpuMemory() *int64
	SetGpuMemoryUsed(v int64) *GetInstanceModelResponseBody
	GetGpuMemoryUsed() *int64
	SetGpuUsed(v int64) *GetInstanceModelResponseBody
	GetGpuUsed() *int64
	SetMemory(v int64) *GetInstanceModelResponseBody
	GetMemory() *int64
	SetMemoryUsed(v int64) *GetInstanceModelResponseBody
	GetMemoryUsed() *int64
	SetModelServiceList(v []*GetInstanceModelResponseBodyModelServiceList) *GetInstanceModelResponseBody
	GetModelServiceList() []*GetInstanceModelResponseBodyModelServiceList
	SetNodeCount(v int64) *GetInstanceModelResponseBody
	GetNodeCount() *int64
	SetRegionId(v string) *GetInstanceModelResponseBody
	GetRegionId() *string
	SetRequestId(v string) *GetInstanceModelResponseBody
	GetRequestId() *string
	SetResourceType(v string) *GetInstanceModelResponseBody
	GetResourceType() *string
	SetStatus(v string) *GetInstanceModelResponseBody
	GetStatus() *string
}

type GetInstanceModelResponseBody struct {
	// example:
	//
	// hologram_aicombo_public_cn-77xxx
	AiInstanceId *string `json:"aiInstanceId,omitempty" xml:"aiInstanceId,omitempty"`
	// example:
	//
	// true
	AutoRenewal *bool `json:"autoRenewal,omitempty" xml:"autoRenewal,omitempty"`
	// example:
	//
	// PrePaid
	ChargeType *string `json:"chargeType,omitempty" xml:"chargeType,omitempty"`
	// example:
	//
	// hologram_aipostpay_public_cn
	CommodityCode *string `json:"commodityCode,omitempty" xml:"commodityCode,omitempty"`
	// example:
	//
	// 32
	Cpu *int64 `json:"cpu,omitempty" xml:"cpu,omitempty"`
	// example:
	//
	// 32
	CpuUsed *int64 `json:"cpuUsed,omitempty" xml:"cpuUsed,omitempty"`
	// example:
	//
	// 2026-01-28T07:44:27.535Z
	ExpirationTime *string `json:"expirationTime,omitempty" xml:"expirationTime,omitempty"`
	// example:
	//
	// 4
	Gpu *int64 `json:"gpu,omitempty" xml:"gpu,omitempty"`
	// example:
	//
	// 128
	GpuMemory *int64 `json:"gpuMemory,omitempty" xml:"gpuMemory,omitempty"`
	// example:
	//
	// 64
	GpuMemoryUsed *int64 `json:"gpuMemoryUsed,omitempty" xml:"gpuMemoryUsed,omitempty"`
	// example:
	//
	// 2
	GpuUsed *int64 `json:"gpuUsed,omitempty" xml:"gpuUsed,omitempty"`
	// example:
	//
	// 128
	Memory *int64 `json:"memory,omitempty" xml:"memory,omitempty"`
	// example:
	//
	// 64
	MemoryUsed       *int64                                          `json:"memoryUsed,omitempty" xml:"memoryUsed,omitempty"`
	ModelServiceList []*GetInstanceModelResponseBodyModelServiceList `json:"modelServiceList,omitempty" xml:"modelServiceList,omitempty" type:"Repeated"`
	// example:
	//
	// 2
	NodeCount *int64 `json:"nodeCount,omitempty" xml:"nodeCount,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 819A7F0F-2951-540F-BD94-6A41ECF0281F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// middle
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	// example:
	//
	// ResourceReady
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetInstanceModelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceModelResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceModelResponseBody) GetAiInstanceId() *string {
	return s.AiInstanceId
}

func (s *GetInstanceModelResponseBody) GetAutoRenewal() *bool {
	return s.AutoRenewal
}

func (s *GetInstanceModelResponseBody) GetChargeType() *string {
	return s.ChargeType
}

func (s *GetInstanceModelResponseBody) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *GetInstanceModelResponseBody) GetCpu() *int64 {
	return s.Cpu
}

func (s *GetInstanceModelResponseBody) GetCpuUsed() *int64 {
	return s.CpuUsed
}

func (s *GetInstanceModelResponseBody) GetExpirationTime() *string {
	return s.ExpirationTime
}

func (s *GetInstanceModelResponseBody) GetGpu() *int64 {
	return s.Gpu
}

func (s *GetInstanceModelResponseBody) GetGpuMemory() *int64 {
	return s.GpuMemory
}

func (s *GetInstanceModelResponseBody) GetGpuMemoryUsed() *int64 {
	return s.GpuMemoryUsed
}

func (s *GetInstanceModelResponseBody) GetGpuUsed() *int64 {
	return s.GpuUsed
}

func (s *GetInstanceModelResponseBody) GetMemory() *int64 {
	return s.Memory
}

func (s *GetInstanceModelResponseBody) GetMemoryUsed() *int64 {
	return s.MemoryUsed
}

func (s *GetInstanceModelResponseBody) GetModelServiceList() []*GetInstanceModelResponseBodyModelServiceList {
	return s.ModelServiceList
}

func (s *GetInstanceModelResponseBody) GetNodeCount() *int64 {
	return s.NodeCount
}

func (s *GetInstanceModelResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *GetInstanceModelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInstanceModelResponseBody) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetInstanceModelResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetInstanceModelResponseBody) SetAiInstanceId(v string) *GetInstanceModelResponseBody {
	s.AiInstanceId = &v
	return s
}

func (s *GetInstanceModelResponseBody) SetAutoRenewal(v bool) *GetInstanceModelResponseBody {
	s.AutoRenewal = &v
	return s
}

func (s *GetInstanceModelResponseBody) SetChargeType(v string) *GetInstanceModelResponseBody {
	s.ChargeType = &v
	return s
}

func (s *GetInstanceModelResponseBody) SetCommodityCode(v string) *GetInstanceModelResponseBody {
	s.CommodityCode = &v
	return s
}

func (s *GetInstanceModelResponseBody) SetCpu(v int64) *GetInstanceModelResponseBody {
	s.Cpu = &v
	return s
}

func (s *GetInstanceModelResponseBody) SetCpuUsed(v int64) *GetInstanceModelResponseBody {
	s.CpuUsed = &v
	return s
}

func (s *GetInstanceModelResponseBody) SetExpirationTime(v string) *GetInstanceModelResponseBody {
	s.ExpirationTime = &v
	return s
}

func (s *GetInstanceModelResponseBody) SetGpu(v int64) *GetInstanceModelResponseBody {
	s.Gpu = &v
	return s
}

func (s *GetInstanceModelResponseBody) SetGpuMemory(v int64) *GetInstanceModelResponseBody {
	s.GpuMemory = &v
	return s
}

func (s *GetInstanceModelResponseBody) SetGpuMemoryUsed(v int64) *GetInstanceModelResponseBody {
	s.GpuMemoryUsed = &v
	return s
}

func (s *GetInstanceModelResponseBody) SetGpuUsed(v int64) *GetInstanceModelResponseBody {
	s.GpuUsed = &v
	return s
}

func (s *GetInstanceModelResponseBody) SetMemory(v int64) *GetInstanceModelResponseBody {
	s.Memory = &v
	return s
}

func (s *GetInstanceModelResponseBody) SetMemoryUsed(v int64) *GetInstanceModelResponseBody {
	s.MemoryUsed = &v
	return s
}

func (s *GetInstanceModelResponseBody) SetModelServiceList(v []*GetInstanceModelResponseBodyModelServiceList) *GetInstanceModelResponseBody {
	s.ModelServiceList = v
	return s
}

func (s *GetInstanceModelResponseBody) SetNodeCount(v int64) *GetInstanceModelResponseBody {
	s.NodeCount = &v
	return s
}

func (s *GetInstanceModelResponseBody) SetRegionId(v string) *GetInstanceModelResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetInstanceModelResponseBody) SetRequestId(v string) *GetInstanceModelResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceModelResponseBody) SetResourceType(v string) *GetInstanceModelResponseBody {
	s.ResourceType = &v
	return s
}

func (s *GetInstanceModelResponseBody) SetStatus(v string) *GetInstanceModelResponseBody {
	s.Status = &v
	return s
}

func (s *GetInstanceModelResponseBody) Validate() error {
	if s.ModelServiceList != nil {
		for _, item := range s.ModelServiceList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetInstanceModelResponseBodyModelServiceList struct {
	// example:
	//
	// sk-42f6c8xxxxxb
	ApiKey *string `json:"apiKey,omitempty" xml:"apiKey,omitempty"`
	// example:
	//
	// 32
	Cpu *int64 `json:"cpu,omitempty" xml:"cpu,omitempty"`
	// example:
	//
	// 2
	Gpu *int64 `json:"gpu,omitempty" xml:"gpu,omitempty"`
	// example:
	//
	// 32
	GpuMemory *int64 `json:"gpuMemory,omitempty" xml:"gpuMemory,omitempty"`
	// example:
	//
	// cn-beijing
	InstanceRegion *string `json:"instanceRegion,omitempty" xml:"instanceRegion,omitempty"`
	// example:
	//
	// 32
	Memory *int64 `json:"memory,omitempty" xml:"memory,omitempty"`
	// example:
	//
	// Failed
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// my_model
	ModelName *string `json:"modelName,omitempty" xml:"modelName,omitempty"`
	// example:
	//
	// {"timeout":600,"max_retries":10,"max_retry_delay":8,"initial_retry_delay":0.5}
	ModelParams *string `json:"modelParams,omitempty" xml:"modelParams,omitempty"`
	// example:
	//
	// qwen3.5-plus
	ModelType *string `json:"modelType,omitempty" xml:"modelType,omitempty"`
	// example:
	//
	// bailian
	Provider *string `json:"provider,omitempty" xml:"provider,omitempty"`
	// example:
	//
	// 2
	ServiceCount *int64 `json:"serviceCount,omitempty" xml:"serviceCount,omitempty"`
	// example:
	//
	// cn-beijing
	ServiceDeployRegion *string `json:"serviceDeployRegion,omitempty" xml:"serviceDeployRegion,omitempty"`
	// example:
	//
	// Running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// embedding
	TaskType *string `json:"taskType,omitempty" xml:"taskType,omitempty"`
	// example:
	//
	// v1.1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetInstanceModelResponseBodyModelServiceList) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceModelResponseBodyModelServiceList) GoString() string {
	return s.String()
}

func (s *GetInstanceModelResponseBodyModelServiceList) GetApiKey() *string {
	return s.ApiKey
}

func (s *GetInstanceModelResponseBodyModelServiceList) GetCpu() *int64 {
	return s.Cpu
}

func (s *GetInstanceModelResponseBodyModelServiceList) GetGpu() *int64 {
	return s.Gpu
}

func (s *GetInstanceModelResponseBodyModelServiceList) GetGpuMemory() *int64 {
	return s.GpuMemory
}

func (s *GetInstanceModelResponseBodyModelServiceList) GetInstanceRegion() *string {
	return s.InstanceRegion
}

func (s *GetInstanceModelResponseBodyModelServiceList) GetMemory() *int64 {
	return s.Memory
}

func (s *GetInstanceModelResponseBodyModelServiceList) GetMessage() *string {
	return s.Message
}

func (s *GetInstanceModelResponseBodyModelServiceList) GetModelName() *string {
	return s.ModelName
}

func (s *GetInstanceModelResponseBodyModelServiceList) GetModelParams() *string {
	return s.ModelParams
}

func (s *GetInstanceModelResponseBodyModelServiceList) GetModelType() *string {
	return s.ModelType
}

func (s *GetInstanceModelResponseBodyModelServiceList) GetProvider() *string {
	return s.Provider
}

func (s *GetInstanceModelResponseBodyModelServiceList) GetServiceCount() *int64 {
	return s.ServiceCount
}

func (s *GetInstanceModelResponseBodyModelServiceList) GetServiceDeployRegion() *string {
	return s.ServiceDeployRegion
}

func (s *GetInstanceModelResponseBodyModelServiceList) GetStatus() *string {
	return s.Status
}

func (s *GetInstanceModelResponseBodyModelServiceList) GetTaskType() *string {
	return s.TaskType
}

func (s *GetInstanceModelResponseBodyModelServiceList) GetVersion() *string {
	return s.Version
}

func (s *GetInstanceModelResponseBodyModelServiceList) SetApiKey(v string) *GetInstanceModelResponseBodyModelServiceList {
	s.ApiKey = &v
	return s
}

func (s *GetInstanceModelResponseBodyModelServiceList) SetCpu(v int64) *GetInstanceModelResponseBodyModelServiceList {
	s.Cpu = &v
	return s
}

func (s *GetInstanceModelResponseBodyModelServiceList) SetGpu(v int64) *GetInstanceModelResponseBodyModelServiceList {
	s.Gpu = &v
	return s
}

func (s *GetInstanceModelResponseBodyModelServiceList) SetGpuMemory(v int64) *GetInstanceModelResponseBodyModelServiceList {
	s.GpuMemory = &v
	return s
}

func (s *GetInstanceModelResponseBodyModelServiceList) SetInstanceRegion(v string) *GetInstanceModelResponseBodyModelServiceList {
	s.InstanceRegion = &v
	return s
}

func (s *GetInstanceModelResponseBodyModelServiceList) SetMemory(v int64) *GetInstanceModelResponseBodyModelServiceList {
	s.Memory = &v
	return s
}

func (s *GetInstanceModelResponseBodyModelServiceList) SetMessage(v string) *GetInstanceModelResponseBodyModelServiceList {
	s.Message = &v
	return s
}

func (s *GetInstanceModelResponseBodyModelServiceList) SetModelName(v string) *GetInstanceModelResponseBodyModelServiceList {
	s.ModelName = &v
	return s
}

func (s *GetInstanceModelResponseBodyModelServiceList) SetModelParams(v string) *GetInstanceModelResponseBodyModelServiceList {
	s.ModelParams = &v
	return s
}

func (s *GetInstanceModelResponseBodyModelServiceList) SetModelType(v string) *GetInstanceModelResponseBodyModelServiceList {
	s.ModelType = &v
	return s
}

func (s *GetInstanceModelResponseBodyModelServiceList) SetProvider(v string) *GetInstanceModelResponseBodyModelServiceList {
	s.Provider = &v
	return s
}

func (s *GetInstanceModelResponseBodyModelServiceList) SetServiceCount(v int64) *GetInstanceModelResponseBodyModelServiceList {
	s.ServiceCount = &v
	return s
}

func (s *GetInstanceModelResponseBodyModelServiceList) SetServiceDeployRegion(v string) *GetInstanceModelResponseBodyModelServiceList {
	s.ServiceDeployRegion = &v
	return s
}

func (s *GetInstanceModelResponseBodyModelServiceList) SetStatus(v string) *GetInstanceModelResponseBodyModelServiceList {
	s.Status = &v
	return s
}

func (s *GetInstanceModelResponseBodyModelServiceList) SetTaskType(v string) *GetInstanceModelResponseBodyModelServiceList {
	s.TaskType = &v
	return s
}

func (s *GetInstanceModelResponseBodyModelServiceList) SetVersion(v string) *GetInstanceModelResponseBodyModelServiceList {
	s.Version = &v
	return s
}

func (s *GetInstanceModelResponseBodyModelServiceList) Validate() error {
	return dara.Validate(s)
}
