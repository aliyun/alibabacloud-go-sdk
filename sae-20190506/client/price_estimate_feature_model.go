// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPriceEstimateFeature interface {
	dara.Model
	String() string
	GoString() string
	SetAppCount(v int64) *PriceEstimateFeature
	GetAppCount() *int64
	SetAppType(v string) *PriceEstimateFeature
	GetAppType() *string
	SetCpuCore(v float32) *PriceEstimateFeature
	GetCpuCore() *float32
	SetCpuStrategy(v string) *PriceEstimateFeature
	GetCpuStrategy() *string
	SetCpuUtilLevel(v string) *PriceEstimateFeature
	GetCpuUtilLevel() *string
	SetCpuUtilMetrics(v []*float32) *PriceEstimateFeature
	GetCpuUtilMetrics() []*float32
	SetEnableCpuIdle(v bool) *PriceEstimateFeature
	GetEnableCpuIdle() *bool
	SetEnvType(v string) *PriceEstimateFeature
	GetEnvType() *string
	SetEphemeralStorageGiB(v int64) *PriceEstimateFeature
	GetEphemeralStorageGiB() *int64
	SetHighLoadInstanceCount(v int64) *PriceEstimateFeature
	GetHighLoadInstanceCount() *int64
	SetHighLoadQps(v float32) *PriceEstimateFeature
	GetHighLoadQps() *float32
	SetHighLoadSeconds(v int64) *PriceEstimateFeature
	GetHighLoadSeconds() *int64
	SetInstanceQps(v float32) *PriceEstimateFeature
	GetInstanceQps() *float32
	SetInternetOutboundGiB(v float32) *PriceEstimateFeature
	GetInternetOutboundGiB() *float32
	SetLowLoadInstanceCount(v int64) *PriceEstimateFeature
	GetLowLoadInstanceCount() *int64
	SetLowLoadQps(v float32) *PriceEstimateFeature
	GetLowLoadQps() *float32
	SetLowLoadSeconds(v int64) *PriceEstimateFeature
	GetLowLoadSeconds() *int64
	SetMaxInstanceCount(v int64) *PriceEstimateFeature
	GetMaxInstanceCount() *int64
	SetMemoryGiB(v float32) *PriceEstimateFeature
	GetMemoryGiB() *float32
	SetMinInstanceCount(v int64) *PriceEstimateFeature
	GetMinInstanceCount() *int64
	SetNewSaeVersion(v string) *PriceEstimateFeature
	GetNewSaeVersion() *string
	SetNoneLoadInstanceCount(v int64) *PriceEstimateFeature
	GetNoneLoadInstanceCount() *int64
	SetNoneLoadSeconds(v int64) *PriceEstimateFeature
	GetNoneLoadSeconds() *int64
	SetRegionId(v string) *PriceEstimateFeature
	GetRegionId() *string
	SetResourceType(v string) *PriceEstimateFeature
	GetResourceType() *string
}

type PriceEstimateFeature struct {
	// example:
	//
	// 1
	AppCount *int64 `json:"AppCount,omitempty" xml:"AppCount,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Web/MicroService
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	CpuCore *float32 `json:"CpuCore,omitempty" xml:"CpuCore,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Request/Always
	CpuStrategy *string `json:"CpuStrategy,omitempty" xml:"CpuStrategy,omitempty"`
	// example:
	//
	// L1
	CpuUtilLevel   *string    `json:"CpuUtilLevel,omitempty" xml:"CpuUtilLevel,omitempty"`
	CpuUtilMetrics []*float32 `json:"CpuUtilMetrics,omitempty" xml:"CpuUtilMetrics,omitempty" type:"Repeated"`
	// example:
	//
	// true
	EnableCpuIdle *bool `json:"EnableCpuIdle,omitempty" xml:"EnableCpuIdle,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Test/Production
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// example:
	//
	// 30
	EphemeralStorageGiB *int64 `json:"EphemeralStorageGiB,omitempty" xml:"EphemeralStorageGiB,omitempty"`
	// example:
	//
	// 3
	HighLoadInstanceCount *int64 `json:"HighLoadInstanceCount,omitempty" xml:"HighLoadInstanceCount,omitempty"`
	// example:
	//
	// 5
	HighLoadQps *float32 `json:"HighLoadQps,omitempty" xml:"HighLoadQps,omitempty"`
	// example:
	//
	// 3600
	HighLoadSeconds *int64 `json:"HighLoadSeconds,omitempty" xml:"HighLoadSeconds,omitempty"`
	// example:
	//
	// 2
	InstanceQps *float32 `json:"InstanceQps,omitempty" xml:"InstanceQps,omitempty"`
	// example:
	//
	// 24
	InternetOutboundGiB *float32 `json:"InternetOutboundGiB,omitempty" xml:"InternetOutboundGiB,omitempty"`
	// example:
	//
	// 1
	LowLoadInstanceCount *int64 `json:"LowLoadInstanceCount,omitempty" xml:"LowLoadInstanceCount,omitempty"`
	// example:
	//
	// 2
	LowLoadQps *float32 `json:"LowLoadQps,omitempty" xml:"LowLoadQps,omitempty"`
	// example:
	//
	// 3600
	LowLoadSeconds *int64 `json:"LowLoadSeconds,omitempty" xml:"LowLoadSeconds,omitempty"`
	// example:
	//
	// 10
	MaxInstanceCount *int64 `json:"MaxInstanceCount,omitempty" xml:"MaxInstanceCount,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	MemoryGiB *float32 `json:"MemoryGiB,omitempty" xml:"MemoryGiB,omitempty"`
	// example:
	//
	// 1
	MinInstanceCount *int64 `json:"MinInstanceCount,omitempty" xml:"MinInstanceCount,omitempty"`
	// example:
	//
	// std
	NewSaeVersion *string `json:"NewSaeVersion,omitempty" xml:"NewSaeVersion,omitempty"`
	// example:
	//
	// 0
	NoneLoadInstanceCount *int64 `json:"NoneLoadInstanceCount,omitempty" xml:"NoneLoadInstanceCount,omitempty"`
	// example:
	//
	// 79200
	NoneLoadSeconds *int64 `json:"NoneLoadSeconds,omitempty" xml:"NoneLoadSeconds,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// haiguang
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s PriceEstimateFeature) String() string {
	return dara.Prettify(s)
}

func (s PriceEstimateFeature) GoString() string {
	return s.String()
}

func (s *PriceEstimateFeature) GetAppCount() *int64 {
	return s.AppCount
}

func (s *PriceEstimateFeature) GetAppType() *string {
	return s.AppType
}

func (s *PriceEstimateFeature) GetCpuCore() *float32 {
	return s.CpuCore
}

func (s *PriceEstimateFeature) GetCpuStrategy() *string {
	return s.CpuStrategy
}

func (s *PriceEstimateFeature) GetCpuUtilLevel() *string {
	return s.CpuUtilLevel
}

func (s *PriceEstimateFeature) GetCpuUtilMetrics() []*float32 {
	return s.CpuUtilMetrics
}

func (s *PriceEstimateFeature) GetEnableCpuIdle() *bool {
	return s.EnableCpuIdle
}

func (s *PriceEstimateFeature) GetEnvType() *string {
	return s.EnvType
}

func (s *PriceEstimateFeature) GetEphemeralStorageGiB() *int64 {
	return s.EphemeralStorageGiB
}

func (s *PriceEstimateFeature) GetHighLoadInstanceCount() *int64 {
	return s.HighLoadInstanceCount
}

func (s *PriceEstimateFeature) GetHighLoadQps() *float32 {
	return s.HighLoadQps
}

func (s *PriceEstimateFeature) GetHighLoadSeconds() *int64 {
	return s.HighLoadSeconds
}

func (s *PriceEstimateFeature) GetInstanceQps() *float32 {
	return s.InstanceQps
}

func (s *PriceEstimateFeature) GetInternetOutboundGiB() *float32 {
	return s.InternetOutboundGiB
}

func (s *PriceEstimateFeature) GetLowLoadInstanceCount() *int64 {
	return s.LowLoadInstanceCount
}

func (s *PriceEstimateFeature) GetLowLoadQps() *float32 {
	return s.LowLoadQps
}

func (s *PriceEstimateFeature) GetLowLoadSeconds() *int64 {
	return s.LowLoadSeconds
}

func (s *PriceEstimateFeature) GetMaxInstanceCount() *int64 {
	return s.MaxInstanceCount
}

func (s *PriceEstimateFeature) GetMemoryGiB() *float32 {
	return s.MemoryGiB
}

func (s *PriceEstimateFeature) GetMinInstanceCount() *int64 {
	return s.MinInstanceCount
}

func (s *PriceEstimateFeature) GetNewSaeVersion() *string {
	return s.NewSaeVersion
}

func (s *PriceEstimateFeature) GetNoneLoadInstanceCount() *int64 {
	return s.NoneLoadInstanceCount
}

func (s *PriceEstimateFeature) GetNoneLoadSeconds() *int64 {
	return s.NoneLoadSeconds
}

func (s *PriceEstimateFeature) GetRegionId() *string {
	return s.RegionId
}

func (s *PriceEstimateFeature) GetResourceType() *string {
	return s.ResourceType
}

func (s *PriceEstimateFeature) SetAppCount(v int64) *PriceEstimateFeature {
	s.AppCount = &v
	return s
}

func (s *PriceEstimateFeature) SetAppType(v string) *PriceEstimateFeature {
	s.AppType = &v
	return s
}

func (s *PriceEstimateFeature) SetCpuCore(v float32) *PriceEstimateFeature {
	s.CpuCore = &v
	return s
}

func (s *PriceEstimateFeature) SetCpuStrategy(v string) *PriceEstimateFeature {
	s.CpuStrategy = &v
	return s
}

func (s *PriceEstimateFeature) SetCpuUtilLevel(v string) *PriceEstimateFeature {
	s.CpuUtilLevel = &v
	return s
}

func (s *PriceEstimateFeature) SetCpuUtilMetrics(v []*float32) *PriceEstimateFeature {
	s.CpuUtilMetrics = v
	return s
}

func (s *PriceEstimateFeature) SetEnableCpuIdle(v bool) *PriceEstimateFeature {
	s.EnableCpuIdle = &v
	return s
}

func (s *PriceEstimateFeature) SetEnvType(v string) *PriceEstimateFeature {
	s.EnvType = &v
	return s
}

func (s *PriceEstimateFeature) SetEphemeralStorageGiB(v int64) *PriceEstimateFeature {
	s.EphemeralStorageGiB = &v
	return s
}

func (s *PriceEstimateFeature) SetHighLoadInstanceCount(v int64) *PriceEstimateFeature {
	s.HighLoadInstanceCount = &v
	return s
}

func (s *PriceEstimateFeature) SetHighLoadQps(v float32) *PriceEstimateFeature {
	s.HighLoadQps = &v
	return s
}

func (s *PriceEstimateFeature) SetHighLoadSeconds(v int64) *PriceEstimateFeature {
	s.HighLoadSeconds = &v
	return s
}

func (s *PriceEstimateFeature) SetInstanceQps(v float32) *PriceEstimateFeature {
	s.InstanceQps = &v
	return s
}

func (s *PriceEstimateFeature) SetInternetOutboundGiB(v float32) *PriceEstimateFeature {
	s.InternetOutboundGiB = &v
	return s
}

func (s *PriceEstimateFeature) SetLowLoadInstanceCount(v int64) *PriceEstimateFeature {
	s.LowLoadInstanceCount = &v
	return s
}

func (s *PriceEstimateFeature) SetLowLoadQps(v float32) *PriceEstimateFeature {
	s.LowLoadQps = &v
	return s
}

func (s *PriceEstimateFeature) SetLowLoadSeconds(v int64) *PriceEstimateFeature {
	s.LowLoadSeconds = &v
	return s
}

func (s *PriceEstimateFeature) SetMaxInstanceCount(v int64) *PriceEstimateFeature {
	s.MaxInstanceCount = &v
	return s
}

func (s *PriceEstimateFeature) SetMemoryGiB(v float32) *PriceEstimateFeature {
	s.MemoryGiB = &v
	return s
}

func (s *PriceEstimateFeature) SetMinInstanceCount(v int64) *PriceEstimateFeature {
	s.MinInstanceCount = &v
	return s
}

func (s *PriceEstimateFeature) SetNewSaeVersion(v string) *PriceEstimateFeature {
	s.NewSaeVersion = &v
	return s
}

func (s *PriceEstimateFeature) SetNoneLoadInstanceCount(v int64) *PriceEstimateFeature {
	s.NoneLoadInstanceCount = &v
	return s
}

func (s *PriceEstimateFeature) SetNoneLoadSeconds(v int64) *PriceEstimateFeature {
	s.NoneLoadSeconds = &v
	return s
}

func (s *PriceEstimateFeature) SetRegionId(v string) *PriceEstimateFeature {
	s.RegionId = &v
	return s
}

func (s *PriceEstimateFeature) SetResourceType(v string) *PriceEstimateFeature {
	s.ResourceType = &v
	return s
}

func (s *PriceEstimateFeature) Validate() error {
	return dara.Validate(s)
}
