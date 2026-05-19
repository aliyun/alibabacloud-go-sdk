// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEmbodiedAIPlatformsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *DescribeEmbodiedAIPlatformsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeEmbodiedAIPlatformsResponseBody
	GetNextToken() *string
	SetPageNumber(v int32) *DescribeEmbodiedAIPlatformsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeEmbodiedAIPlatformsResponseBody
	GetPageSize() *int32
	SetPlatforms(v []*DescribeEmbodiedAIPlatformsResponseBodyPlatforms) *DescribeEmbodiedAIPlatformsResponseBody
	GetPlatforms() []*DescribeEmbodiedAIPlatformsResponseBodyPlatforms
	SetRequestId(v string) *DescribeEmbodiedAIPlatformsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeEmbodiedAIPlatformsResponseBody
	GetTotalCount() *int64
}

type DescribeEmbodiedAIPlatformsResponseBody struct {
	// example:
	//
	// 30
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 298a7d5473b128dfe0b5e8707e******
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 30
	PageSize  *int32                                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Platforms []*DescribeEmbodiedAIPlatformsResponseBodyPlatforms `json:"Platforms,omitempty" xml:"Platforms,omitempty" type:"Repeated"`
	// example:
	//
	// B47EED99-BFA5-529D-8D85-A6642421D390
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 50
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeEmbodiedAIPlatformsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEmbodiedAIPlatformsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEmbodiedAIPlatformsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeEmbodiedAIPlatformsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeEmbodiedAIPlatformsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeEmbodiedAIPlatformsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeEmbodiedAIPlatformsResponseBody) GetPlatforms() []*DescribeEmbodiedAIPlatformsResponseBodyPlatforms {
	return s.Platforms
}

func (s *DescribeEmbodiedAIPlatformsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEmbodiedAIPlatformsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeEmbodiedAIPlatformsResponseBody) SetMaxResults(v int32) *DescribeEmbodiedAIPlatformsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeEmbodiedAIPlatformsResponseBody) SetNextToken(v string) *DescribeEmbodiedAIPlatformsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeEmbodiedAIPlatformsResponseBody) SetPageNumber(v int32) *DescribeEmbodiedAIPlatformsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeEmbodiedAIPlatformsResponseBody) SetPageSize(v int32) *DescribeEmbodiedAIPlatformsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeEmbodiedAIPlatformsResponseBody) SetPlatforms(v []*DescribeEmbodiedAIPlatformsResponseBodyPlatforms) *DescribeEmbodiedAIPlatformsResponseBody {
	s.Platforms = v
	return s
}

func (s *DescribeEmbodiedAIPlatformsResponseBody) SetRequestId(v string) *DescribeEmbodiedAIPlatformsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEmbodiedAIPlatformsResponseBody) SetTotalCount(v int64) *DescribeEmbodiedAIPlatformsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeEmbodiedAIPlatformsResponseBody) Validate() error {
	if s.Platforms != nil {
		for _, item := range s.Platforms {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeEmbodiedAIPlatformsResponseBodyPlatforms struct {
	// example:
	//
	// 2025-12-01 14:55:36
	CreateTime *int64                                                     `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	EapConfig  *DescribeEmbodiedAIPlatformsResponseBodyPlatformsEapConfig `json:"EapConfig,omitempty" xml:"EapConfig,omitempty" type:"Struct"`
	// example:
	//
	// adb-lake-cn-beijing-5q1w******
	OssBucketName *string `json:"OssBucketName,omitempty" xml:"OssBucketName,omitempty"`
	// example:
	//
	// platform1
	PlatformName   *string                                                         `json:"PlatformName,omitempty" xml:"PlatformName,omitempty"`
	RayConfig      *DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayConfig      `json:"RayConfig,omitempty" xml:"RayConfig,omitempty" type:"Struct"`
	RayTrainConfig *DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayTrainConfig `json:"RayTrainConfig,omitempty" xml:"RayTrainConfig,omitempty" type:"Struct"`
	// example:
	//
	// running
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s DescribeEmbodiedAIPlatformsResponseBodyPlatforms) String() string {
	return dara.Prettify(s)
}

func (s DescribeEmbodiedAIPlatformsResponseBodyPlatforms) GoString() string {
	return s.String()
}

func (s *DescribeEmbodiedAIPlatformsResponseBodyPlatforms) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *DescribeEmbodiedAIPlatformsResponseBodyPlatforms) GetEapConfig() *DescribeEmbodiedAIPlatformsResponseBodyPlatformsEapConfig {
	return s.EapConfig
}

func (s *DescribeEmbodiedAIPlatformsResponseBodyPlatforms) GetOssBucketName() *string {
	return s.OssBucketName
}

func (s *DescribeEmbodiedAIPlatformsResponseBodyPlatforms) GetPlatformName() *string {
	return s.PlatformName
}

func (s *DescribeEmbodiedAIPlatformsResponseBodyPlatforms) GetRayConfig() *DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayConfig {
	return s.RayConfig
}

func (s *DescribeEmbodiedAIPlatformsResponseBodyPlatforms) GetRayTrainConfig() *DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayTrainConfig {
	return s.RayTrainConfig
}

func (s *DescribeEmbodiedAIPlatformsResponseBodyPlatforms) GetState() *string {
	return s.State
}

func (s *DescribeEmbodiedAIPlatformsResponseBodyPlatforms) SetCreateTime(v int64) *DescribeEmbodiedAIPlatformsResponseBodyPlatforms {
	s.CreateTime = &v
	return s
}

func (s *DescribeEmbodiedAIPlatformsResponseBodyPlatforms) SetEapConfig(v *DescribeEmbodiedAIPlatformsResponseBodyPlatformsEapConfig) *DescribeEmbodiedAIPlatformsResponseBodyPlatforms {
	s.EapConfig = v
	return s
}

func (s *DescribeEmbodiedAIPlatformsResponseBodyPlatforms) SetOssBucketName(v string) *DescribeEmbodiedAIPlatformsResponseBodyPlatforms {
	s.OssBucketName = &v
	return s
}

func (s *DescribeEmbodiedAIPlatformsResponseBodyPlatforms) SetPlatformName(v string) *DescribeEmbodiedAIPlatformsResponseBodyPlatforms {
	s.PlatformName = &v
	return s
}

func (s *DescribeEmbodiedAIPlatformsResponseBodyPlatforms) SetRayConfig(v *DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayConfig) *DescribeEmbodiedAIPlatformsResponseBodyPlatforms {
	s.RayConfig = v
	return s
}

func (s *DescribeEmbodiedAIPlatformsResponseBodyPlatforms) SetRayTrainConfig(v *DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayTrainConfig) *DescribeEmbodiedAIPlatformsResponseBodyPlatforms {
	s.RayTrainConfig = v
	return s
}

func (s *DescribeEmbodiedAIPlatformsResponseBodyPlatforms) SetState(v string) *DescribeEmbodiedAIPlatformsResponseBodyPlatforms {
	s.State = &v
	return s
}

func (s *DescribeEmbodiedAIPlatformsResponseBodyPlatforms) Validate() error {
	if s.EapConfig != nil {
		if err := s.EapConfig.Validate(); err != nil {
			return err
		}
	}
	if s.RayConfig != nil {
		if err := s.RayConfig.Validate(); err != nil {
			return err
		}
	}
	if s.RayTrainConfig != nil {
		if err := s.RayTrainConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeEmbodiedAIPlatformsResponseBodyPlatformsEapConfig struct {
	// example:
	//
	// http://amv-2z******-***-roboto.ads.example.com:80
	WebserverAddress *string `json:"WebserverAddress,omitempty" xml:"WebserverAddress,omitempty"`
	// example:
	//
	// large
	WebserverSpecName *string `json:"WebserverSpecName,omitempty" xml:"WebserverSpecName,omitempty"`
}

func (s DescribeEmbodiedAIPlatformsResponseBodyPlatformsEapConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeEmbodiedAIPlatformsResponseBodyPlatformsEapConfig) GoString() string {
	return s.String()
}

func (s *DescribeEmbodiedAIPlatformsResponseBodyPlatformsEapConfig) GetWebserverAddress() *string {
	return s.WebserverAddress
}

func (s *DescribeEmbodiedAIPlatformsResponseBodyPlatformsEapConfig) GetWebserverSpecName() *string {
	return s.WebserverSpecName
}

func (s *DescribeEmbodiedAIPlatformsResponseBodyPlatformsEapConfig) SetWebserverAddress(v string) *DescribeEmbodiedAIPlatformsResponseBodyPlatformsEapConfig {
	s.WebserverAddress = &v
	return s
}

func (s *DescribeEmbodiedAIPlatformsResponseBodyPlatformsEapConfig) SetWebserverSpecName(v string) *DescribeEmbodiedAIPlatformsResponseBodyPlatformsEapConfig {
	s.WebserverSpecName = &v
	return s
}

func (s *DescribeEmbodiedAIPlatformsResponseBodyPlatformsEapConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayConfig struct {
	// example:
	//
	// BASIC
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// example:
	//
	// 100G
	HeadDiskCapacity *string `json:"HeadDiskCapacity,omitempty" xml:"HeadDiskCapacity,omitempty"`
	// example:
	//
	// large
	HeadSpec *string `json:"HeadSpec,omitempty" xml:"HeadSpec,omitempty"`
	// example:
	//
	// CPU
	HeadSpecType *string `json:"HeadSpecType,omitempty" xml:"HeadSpecType,omitempty"`
	// example:
	//
	// http://ray-cluster-address.example.com
	RayClusterAddress *string `json:"RayClusterAddress,omitempty" xml:"RayClusterAddress,omitempty"`
	// example:
	//
	// http://ray-dashboard-address.example.com
	RayDashboardAddress *string `json:"RayDashboardAddress,omitempty" xml:"RayDashboardAddress,omitempty"`
	// example:
	//
	// http://ray-grafana-address.example.com
	RayGrafanaAddress *string                                                                  `json:"RayGrafanaAddress,omitempty" xml:"RayGrafanaAddress,omitempty"`
	WorkerGroups      []*DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayConfigWorkerGroups `json:"WorkerGroups,omitempty" xml:"WorkerGroups,omitempty" type:"Repeated"`
}

func (s DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayConfig) GoString() string {
	return s.String()
}

func (s *DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayConfig) GetCategory() *string {
	return s.Category
}

func (s *DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayConfig) GetHeadDiskCapacity() *string {
	return s.HeadDiskCapacity
}

func (s *DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayConfig) GetHeadSpec() *string {
	return s.HeadSpec
}

func (s *DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayConfig) GetHeadSpecType() *string {
	return s.HeadSpecType
}

func (s *DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayConfig) GetRayClusterAddress() *string {
	return s.RayClusterAddress
}

func (s *DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayConfig) GetRayDashboardAddress() *string {
	return s.RayDashboardAddress
}

func (s *DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayConfig) GetRayGrafanaAddress() *string {
	return s.RayGrafanaAddress
}

func (s *DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayConfig) GetWorkerGroups() []*DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayConfigWorkerGroups {
	return s.WorkerGroups
}

func (s *DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayConfig) SetCategory(v string) *DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayConfig {
	s.Category = &v
	return s
}

func (s *DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayConfig) SetHeadDiskCapacity(v string) *DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayConfig {
	s.HeadDiskCapacity = &v
	return s
}

func (s *DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayConfig) SetHeadSpec(v string) *DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayConfig {
	s.HeadSpec = &v
	return s
}

func (s *DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayConfig) SetHeadSpecType(v string) *DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayConfig {
	s.HeadSpecType = &v
	return s
}

func (s *DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayConfig) SetRayClusterAddress(v string) *DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayConfig {
	s.RayClusterAddress = &v
	return s
}

func (s *DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayConfig) SetRayDashboardAddress(v string) *DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayConfig {
	s.RayDashboardAddress = &v
	return s
}

func (s *DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayConfig) SetRayGrafanaAddress(v string) *DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayConfig {
	s.RayGrafanaAddress = &v
	return s
}

func (s *DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayConfig) SetWorkerGroups(v []*DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayConfigWorkerGroups) *DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayConfig {
	s.WorkerGroups = v
	return s
}

func (s *DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayConfig) Validate() error {
	if s.WorkerGroups != nil {
		for _, item := range s.WorkerGroups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayConfigWorkerGroups struct {
	// example:
	//
	// 1
	AllocateUnit *string `json:"AllocateUnit,omitempty" xml:"AllocateUnit,omitempty"`
	// example:
	//
	// worker1
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// example:
	//
	// 2
	MaxWorkerQuantity *int32 `json:"MaxWorkerQuantity,omitempty" xml:"MaxWorkerQuantity,omitempty"`
	// example:
	//
	// 1
	MinWorkerQuantity *int32 `json:"MinWorkerQuantity,omitempty" xml:"MinWorkerQuantity,omitempty"`
	// example:
	//
	// 100G
	WorkerDiskCapacity *string `json:"WorkerDiskCapacity,omitempty" xml:"WorkerDiskCapacity,omitempty"`
	// example:
	//
	// large
	WorkerSpecName *string `json:"WorkerSpecName,omitempty" xml:"WorkerSpecName,omitempty"`
	// example:
	//
	// CPU
	WorkerSpecType *string `json:"WorkerSpecType,omitempty" xml:"WorkerSpecType,omitempty"`
}

func (s DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayConfigWorkerGroups) String() string {
	return dara.Prettify(s)
}

func (s DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayConfigWorkerGroups) GoString() string {
	return s.String()
}

func (s *DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayConfigWorkerGroups) GetAllocateUnit() *string {
	return s.AllocateUnit
}

func (s *DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayConfigWorkerGroups) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayConfigWorkerGroups) GetMaxWorkerQuantity() *int32 {
	return s.MaxWorkerQuantity
}

func (s *DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayConfigWorkerGroups) GetMinWorkerQuantity() *int32 {
	return s.MinWorkerQuantity
}

func (s *DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayConfigWorkerGroups) GetWorkerDiskCapacity() *string {
	return s.WorkerDiskCapacity
}

func (s *DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayConfigWorkerGroups) GetWorkerSpecName() *string {
	return s.WorkerSpecName
}

func (s *DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayConfigWorkerGroups) GetWorkerSpecType() *string {
	return s.WorkerSpecType
}

func (s *DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayConfigWorkerGroups) SetAllocateUnit(v string) *DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayConfigWorkerGroups {
	s.AllocateUnit = &v
	return s
}

func (s *DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayConfigWorkerGroups) SetGroupName(v string) *DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayConfigWorkerGroups {
	s.GroupName = &v
	return s
}

func (s *DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayConfigWorkerGroups) SetMaxWorkerQuantity(v int32) *DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayConfigWorkerGroups {
	s.MaxWorkerQuantity = &v
	return s
}

func (s *DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayConfigWorkerGroups) SetMinWorkerQuantity(v int32) *DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayConfigWorkerGroups {
	s.MinWorkerQuantity = &v
	return s
}

func (s *DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayConfigWorkerGroups) SetWorkerDiskCapacity(v string) *DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayConfigWorkerGroups {
	s.WorkerDiskCapacity = &v
	return s
}

func (s *DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayConfigWorkerGroups) SetWorkerSpecName(v string) *DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayConfigWorkerGroups {
	s.WorkerSpecName = &v
	return s
}

func (s *DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayConfigWorkerGroups) SetWorkerSpecType(v string) *DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayConfigWorkerGroups {
	s.WorkerSpecType = &v
	return s
}

func (s *DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayConfigWorkerGroups) Validate() error {
	return dara.Validate(s)
}

type DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayTrainConfig struct {
	CpuAcu         *int64                                                                        `json:"CpuAcu,omitempty" xml:"CpuAcu,omitempty"`
	GpuSpecs       []*DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayTrainConfigGpuSpecs     `json:"GpuSpecs,omitempty" xml:"GpuSpecs,omitempty" type:"Repeated"`
	TerminalConfig *DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayTrainConfigTerminalConfig `json:"TerminalConfig,omitempty" xml:"TerminalConfig,omitempty" type:"Struct"`
}

func (s DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayTrainConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayTrainConfig) GoString() string {
	return s.String()
}

func (s *DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayTrainConfig) GetCpuAcu() *int64 {
	return s.CpuAcu
}

func (s *DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayTrainConfig) GetGpuSpecs() []*DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayTrainConfigGpuSpecs {
	return s.GpuSpecs
}

func (s *DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayTrainConfig) GetTerminalConfig() *DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayTrainConfigTerminalConfig {
	return s.TerminalConfig
}

func (s *DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayTrainConfig) SetCpuAcu(v int64) *DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayTrainConfig {
	s.CpuAcu = &v
	return s
}

func (s *DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayTrainConfig) SetGpuSpecs(v []*DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayTrainConfigGpuSpecs) *DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayTrainConfig {
	s.GpuSpecs = v
	return s
}

func (s *DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayTrainConfig) SetTerminalConfig(v *DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayTrainConfigTerminalConfig) *DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayTrainConfig {
	s.TerminalConfig = v
	return s
}

func (s *DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayTrainConfig) Validate() error {
	if s.GpuSpecs != nil {
		for _, item := range s.GpuSpecs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TerminalConfig != nil {
		if err := s.TerminalConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayTrainConfigGpuSpecs struct {
	// example:
	//
	// "1"
	AllocateUnit *string `json:"AllocateUnit,omitempty" xml:"AllocateUnit,omitempty"`
	// example:
	//
	// 1
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// example:
	//
	// ADB.MLTensor.2
	SpecName *string `json:"SpecName,omitempty" xml:"SpecName,omitempty"`
}

func (s DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayTrainConfigGpuSpecs) String() string {
	return dara.Prettify(s)
}

func (s DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayTrainConfigGpuSpecs) GoString() string {
	return s.String()
}

func (s *DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayTrainConfigGpuSpecs) GetAllocateUnit() *string {
	return s.AllocateUnit
}

func (s *DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayTrainConfigGpuSpecs) GetCount() *int64 {
	return s.Count
}

func (s *DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayTrainConfigGpuSpecs) GetSpecName() *string {
	return s.SpecName
}

func (s *DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayTrainConfigGpuSpecs) SetAllocateUnit(v string) *DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayTrainConfigGpuSpecs {
	s.AllocateUnit = &v
	return s
}

func (s *DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayTrainConfigGpuSpecs) SetCount(v int64) *DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayTrainConfigGpuSpecs {
	s.Count = &v
	return s
}

func (s *DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayTrainConfigGpuSpecs) SetSpecName(v string) *DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayTrainConfigGpuSpecs {
	s.SpecName = &v
	return s
}

func (s *DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayTrainConfigGpuSpecs) Validate() error {
	return dara.Validate(s)
}

type DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayTrainConfigTerminalConfig struct {
	AcrConfig *DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayTrainConfigTerminalConfigAcrConfig `json:"AcrConfig,omitempty" xml:"AcrConfig,omitempty" type:"Struct"`
}

func (s DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayTrainConfigTerminalConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayTrainConfigTerminalConfig) GoString() string {
	return s.String()
}

func (s *DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayTrainConfigTerminalConfig) GetAcrConfig() *DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayTrainConfigTerminalConfigAcrConfig {
	return s.AcrConfig
}

func (s *DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayTrainConfigTerminalConfig) SetAcrConfig(v *DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayTrainConfigTerminalConfigAcrConfig) *DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayTrainConfigTerminalConfig {
	s.AcrConfig = v
	return s
}

func (s *DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayTrainConfigTerminalConfig) Validate() error {
	if s.AcrConfig != nil {
		if err := s.AcrConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayTrainConfigTerminalConfigAcrConfig struct {
	// example:
	//
	// cri-***
	InstanceId *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Namespaces []*string `json:"Namespaces,omitempty" xml:"Namespaces,omitempty" type:"Repeated"`
	// example:
	//
	// example-vpc.example-region.cr.aliyuncs.com
	Registry *string `json:"Registry,omitempty" xml:"Registry,omitempty"`
}

func (s DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayTrainConfigTerminalConfigAcrConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayTrainConfigTerminalConfigAcrConfig) GoString() string {
	return s.String()
}

func (s *DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayTrainConfigTerminalConfigAcrConfig) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayTrainConfigTerminalConfigAcrConfig) GetNamespaces() []*string {
	return s.Namespaces
}

func (s *DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayTrainConfigTerminalConfigAcrConfig) GetRegistry() *string {
	return s.Registry
}

func (s *DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayTrainConfigTerminalConfigAcrConfig) SetInstanceId(v string) *DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayTrainConfigTerminalConfigAcrConfig {
	s.InstanceId = &v
	return s
}

func (s *DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayTrainConfigTerminalConfigAcrConfig) SetNamespaces(v []*string) *DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayTrainConfigTerminalConfigAcrConfig {
	s.Namespaces = v
	return s
}

func (s *DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayTrainConfigTerminalConfigAcrConfig) SetRegistry(v string) *DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayTrainConfigTerminalConfigAcrConfig {
	s.Registry = &v
	return s
}

func (s *DescribeEmbodiedAIPlatformsResponseBodyPlatformsRayTrainConfigTerminalConfigAcrConfig) Validate() error {
	return dara.Validate(s)
}
