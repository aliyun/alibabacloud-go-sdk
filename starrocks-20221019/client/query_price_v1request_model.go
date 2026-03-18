// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPriceV1Request interface {
	dara.Model
	String() string
	GoString() string
	SetAgentNodeGroup(v *QueryPriceV1RequestAgentNodeGroup) *QueryPriceV1Request
	GetAgentNodeGroup() *QueryPriceV1RequestAgentNodeGroup
	SetBackendNodeGroups(v []*QueryPriceV1RequestBackendNodeGroups) *QueryPriceV1Request
	GetBackendNodeGroups() []*QueryPriceV1RequestBackendNodeGroups
	SetDuration(v int32) *QueryPriceV1Request
	GetDuration() *int32
	SetFrontendNodeGroups(v []*QueryPriceV1RequestFrontendNodeGroups) *QueryPriceV1Request
	GetFrontendNodeGroups() []*QueryPriceV1RequestFrontendNodeGroups
	SetObserverNodeGroups(v []*QueryPriceV1RequestObserverNodeGroups) *QueryPriceV1Request
	GetObserverNodeGroups() []*QueryPriceV1RequestObserverNodeGroups
	SetPackageType(v string) *QueryPriceV1Request
	GetPackageType() *string
	SetPayType(v string) *QueryPriceV1Request
	GetPayType() *string
	SetPricingCycle(v string) *QueryPriceV1Request
	GetPricingCycle() *string
	SetPromotionOptionNo(v string) *QueryPriceV1Request
	GetPromotionOptionNo() *string
	SetRegionId(v string) *QueryPriceV1Request
	GetRegionId() *string
	SetRunMode(v string) *QueryPriceV1Request
	GetRunMode() *string
}

type QueryPriceV1Request struct {
	AgentNodeGroup    *QueryPriceV1RequestAgentNodeGroup      `json:"AgentNodeGroup,omitempty" xml:"AgentNodeGroup,omitempty" type:"Struct"`
	BackendNodeGroups []*QueryPriceV1RequestBackendNodeGroups `json:"BackendNodeGroups,omitempty" xml:"BackendNodeGroups,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Duration           *int32                                   `json:"Duration,omitempty" xml:"Duration,omitempty"`
	FrontendNodeGroups []*QueryPriceV1RequestFrontendNodeGroups `json:"FrontendNodeGroups,omitempty" xml:"FrontendNodeGroups,omitempty" type:"Repeated"`
	ObserverNodeGroups []*QueryPriceV1RequestObserverNodeGroups `json:"ObserverNodeGroups,omitempty" xml:"ObserverNodeGroups,omitempty" type:"Repeated"`
	// example:
	//
	// official
	PackageType *string `json:"PackageType,omitempty" xml:"PackageType,omitempty"`
	// example:
	//
	// prePaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// example:
	//
	// Month
	PricingCycle *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	// example:
	//
	// youhuiquan_12378dfj6
	PromotionOptionNo *string `json:"PromotionOptionNo,omitempty" xml:"PromotionOptionNo,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// shared_data
	RunMode *string `json:"RunMode,omitempty" xml:"RunMode,omitempty"`
}

func (s QueryPriceV1Request) String() string {
	return dara.Prettify(s)
}

func (s QueryPriceV1Request) GoString() string {
	return s.String()
}

func (s *QueryPriceV1Request) GetAgentNodeGroup() *QueryPriceV1RequestAgentNodeGroup {
	return s.AgentNodeGroup
}

func (s *QueryPriceV1Request) GetBackendNodeGroups() []*QueryPriceV1RequestBackendNodeGroups {
	return s.BackendNodeGroups
}

func (s *QueryPriceV1Request) GetDuration() *int32 {
	return s.Duration
}

func (s *QueryPriceV1Request) GetFrontendNodeGroups() []*QueryPriceV1RequestFrontendNodeGroups {
	return s.FrontendNodeGroups
}

func (s *QueryPriceV1Request) GetObserverNodeGroups() []*QueryPriceV1RequestObserverNodeGroups {
	return s.ObserverNodeGroups
}

func (s *QueryPriceV1Request) GetPackageType() *string {
	return s.PackageType
}

func (s *QueryPriceV1Request) GetPayType() *string {
	return s.PayType
}

func (s *QueryPriceV1Request) GetPricingCycle() *string {
	return s.PricingCycle
}

func (s *QueryPriceV1Request) GetPromotionOptionNo() *string {
	return s.PromotionOptionNo
}

func (s *QueryPriceV1Request) GetRegionId() *string {
	return s.RegionId
}

func (s *QueryPriceV1Request) GetRunMode() *string {
	return s.RunMode
}

func (s *QueryPriceV1Request) SetAgentNodeGroup(v *QueryPriceV1RequestAgentNodeGroup) *QueryPriceV1Request {
	s.AgentNodeGroup = v
	return s
}

func (s *QueryPriceV1Request) SetBackendNodeGroups(v []*QueryPriceV1RequestBackendNodeGroups) *QueryPriceV1Request {
	s.BackendNodeGroups = v
	return s
}

func (s *QueryPriceV1Request) SetDuration(v int32) *QueryPriceV1Request {
	s.Duration = &v
	return s
}

func (s *QueryPriceV1Request) SetFrontendNodeGroups(v []*QueryPriceV1RequestFrontendNodeGroups) *QueryPriceV1Request {
	s.FrontendNodeGroups = v
	return s
}

func (s *QueryPriceV1Request) SetObserverNodeGroups(v []*QueryPriceV1RequestObserverNodeGroups) *QueryPriceV1Request {
	s.ObserverNodeGroups = v
	return s
}

func (s *QueryPriceV1Request) SetPackageType(v string) *QueryPriceV1Request {
	s.PackageType = &v
	return s
}

func (s *QueryPriceV1Request) SetPayType(v string) *QueryPriceV1Request {
	s.PayType = &v
	return s
}

func (s *QueryPriceV1Request) SetPricingCycle(v string) *QueryPriceV1Request {
	s.PricingCycle = &v
	return s
}

func (s *QueryPriceV1Request) SetPromotionOptionNo(v string) *QueryPriceV1Request {
	s.PromotionOptionNo = &v
	return s
}

func (s *QueryPriceV1Request) SetRegionId(v string) *QueryPriceV1Request {
	s.RegionId = &v
	return s
}

func (s *QueryPriceV1Request) SetRunMode(v string) *QueryPriceV1Request {
	s.RunMode = &v
	return s
}

func (s *QueryPriceV1Request) Validate() error {
	if s.AgentNodeGroup != nil {
		if err := s.AgentNodeGroup.Validate(); err != nil {
			return err
		}
	}
	if s.BackendNodeGroups != nil {
		for _, item := range s.BackendNodeGroups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.FrontendNodeGroups != nil {
		for _, item := range s.FrontendNodeGroups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ObserverNodeGroups != nil {
		for _, item := range s.ObserverNodeGroups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryPriceV1RequestAgentNodeGroup struct {
	// example:
	//
	// 2
	Cu *int32 `json:"cu,omitempty" xml:"cu,omitempty"`
}

func (s QueryPriceV1RequestAgentNodeGroup) String() string {
	return dara.Prettify(s)
}

func (s QueryPriceV1RequestAgentNodeGroup) GoString() string {
	return s.String()
}

func (s *QueryPriceV1RequestAgentNodeGroup) GetCu() *int32 {
	return s.Cu
}

func (s *QueryPriceV1RequestAgentNodeGroup) SetCu(v int32) *QueryPriceV1RequestAgentNodeGroup {
	s.Cu = &v
	return s
}

func (s *QueryPriceV1RequestAgentNodeGroup) Validate() error {
	return dara.Validate(s)
}

type QueryPriceV1RequestBackendNodeGroups struct {
	// example:
	//
	// 8
	Cu *string `json:"cu,omitempty" xml:"cu,omitempty"`
	// example:
	//
	// 1
	DiskNumber *int32 `json:"diskNumber,omitempty" xml:"diskNumber,omitempty"`
	// example:
	//
	// local_ssd_4_4xlarge
	LocalStorageInstanceType *string `json:"localStorageInstanceType,omitempty" xml:"localStorageInstanceType,omitempty"`
	// example:
	//
	// 3
	ResidentNodeNumber *int32 `json:"residentNodeNumber,omitempty" xml:"residentNodeNumber,omitempty"`
	// example:
	//
	// standard
	SpecType *string `json:"specType,omitempty" xml:"specType,omitempty"`
	// example:
	//
	// pl1
	StoragePerformanceLevel *string `json:"storagePerformanceLevel,omitempty" xml:"storagePerformanceLevel,omitempty"`
	// example:
	//
	// 100
	StorageSize *int32 `json:"storageSize,omitempty" xml:"storageSize,omitempty"`
}

func (s QueryPriceV1RequestBackendNodeGroups) String() string {
	return dara.Prettify(s)
}

func (s QueryPriceV1RequestBackendNodeGroups) GoString() string {
	return s.String()
}

func (s *QueryPriceV1RequestBackendNodeGroups) GetCu() *string {
	return s.Cu
}

func (s *QueryPriceV1RequestBackendNodeGroups) GetDiskNumber() *int32 {
	return s.DiskNumber
}

func (s *QueryPriceV1RequestBackendNodeGroups) GetLocalStorageInstanceType() *string {
	return s.LocalStorageInstanceType
}

func (s *QueryPriceV1RequestBackendNodeGroups) GetResidentNodeNumber() *int32 {
	return s.ResidentNodeNumber
}

func (s *QueryPriceV1RequestBackendNodeGroups) GetSpecType() *string {
	return s.SpecType
}

func (s *QueryPriceV1RequestBackendNodeGroups) GetStoragePerformanceLevel() *string {
	return s.StoragePerformanceLevel
}

func (s *QueryPriceV1RequestBackendNodeGroups) GetStorageSize() *int32 {
	return s.StorageSize
}

func (s *QueryPriceV1RequestBackendNodeGroups) SetCu(v string) *QueryPriceV1RequestBackendNodeGroups {
	s.Cu = &v
	return s
}

func (s *QueryPriceV1RequestBackendNodeGroups) SetDiskNumber(v int32) *QueryPriceV1RequestBackendNodeGroups {
	s.DiskNumber = &v
	return s
}

func (s *QueryPriceV1RequestBackendNodeGroups) SetLocalStorageInstanceType(v string) *QueryPriceV1RequestBackendNodeGroups {
	s.LocalStorageInstanceType = &v
	return s
}

func (s *QueryPriceV1RequestBackendNodeGroups) SetResidentNodeNumber(v int32) *QueryPriceV1RequestBackendNodeGroups {
	s.ResidentNodeNumber = &v
	return s
}

func (s *QueryPriceV1RequestBackendNodeGroups) SetSpecType(v string) *QueryPriceV1RequestBackendNodeGroups {
	s.SpecType = &v
	return s
}

func (s *QueryPriceV1RequestBackendNodeGroups) SetStoragePerformanceLevel(v string) *QueryPriceV1RequestBackendNodeGroups {
	s.StoragePerformanceLevel = &v
	return s
}

func (s *QueryPriceV1RequestBackendNodeGroups) SetStorageSize(v int32) *QueryPriceV1RequestBackendNodeGroups {
	s.StorageSize = &v
	return s
}

func (s *QueryPriceV1RequestBackendNodeGroups) Validate() error {
	return dara.Validate(s)
}

type QueryPriceV1RequestFrontendNodeGroups struct {
	// example:
	//
	// 8
	Cu *string `json:"cu,omitempty" xml:"cu,omitempty"`
	// example:
	//
	// 1
	DiskNumber *int32 `json:"diskNumber,omitempty" xml:"diskNumber,omitempty"`
	// example:
	//
	// null
	LocalStorageInstanceType *string `json:"localStorageInstanceType,omitempty" xml:"localStorageInstanceType,omitempty"`
	// example:
	//
	// 3
	ResidentNodeNumber *int32 `json:"residentNodeNumber,omitempty" xml:"residentNodeNumber,omitempty"`
	// example:
	//
	// standard
	SpecType *string `json:"specType,omitempty" xml:"specType,omitempty"`
	// example:
	//
	// pl1
	StoragePerformanceLevel *string `json:"storagePerformanceLevel,omitempty" xml:"storagePerformanceLevel,omitempty"`
	// example:
	//
	// 100
	StorageSize *int32 `json:"storageSize,omitempty" xml:"storageSize,omitempty"`
}

func (s QueryPriceV1RequestFrontendNodeGroups) String() string {
	return dara.Prettify(s)
}

func (s QueryPriceV1RequestFrontendNodeGroups) GoString() string {
	return s.String()
}

func (s *QueryPriceV1RequestFrontendNodeGroups) GetCu() *string {
	return s.Cu
}

func (s *QueryPriceV1RequestFrontendNodeGroups) GetDiskNumber() *int32 {
	return s.DiskNumber
}

func (s *QueryPriceV1RequestFrontendNodeGroups) GetLocalStorageInstanceType() *string {
	return s.LocalStorageInstanceType
}

func (s *QueryPriceV1RequestFrontendNodeGroups) GetResidentNodeNumber() *int32 {
	return s.ResidentNodeNumber
}

func (s *QueryPriceV1RequestFrontendNodeGroups) GetSpecType() *string {
	return s.SpecType
}

func (s *QueryPriceV1RequestFrontendNodeGroups) GetStoragePerformanceLevel() *string {
	return s.StoragePerformanceLevel
}

func (s *QueryPriceV1RequestFrontendNodeGroups) GetStorageSize() *int32 {
	return s.StorageSize
}

func (s *QueryPriceV1RequestFrontendNodeGroups) SetCu(v string) *QueryPriceV1RequestFrontendNodeGroups {
	s.Cu = &v
	return s
}

func (s *QueryPriceV1RequestFrontendNodeGroups) SetDiskNumber(v int32) *QueryPriceV1RequestFrontendNodeGroups {
	s.DiskNumber = &v
	return s
}

func (s *QueryPriceV1RequestFrontendNodeGroups) SetLocalStorageInstanceType(v string) *QueryPriceV1RequestFrontendNodeGroups {
	s.LocalStorageInstanceType = &v
	return s
}

func (s *QueryPriceV1RequestFrontendNodeGroups) SetResidentNodeNumber(v int32) *QueryPriceV1RequestFrontendNodeGroups {
	s.ResidentNodeNumber = &v
	return s
}

func (s *QueryPriceV1RequestFrontendNodeGroups) SetSpecType(v string) *QueryPriceV1RequestFrontendNodeGroups {
	s.SpecType = &v
	return s
}

func (s *QueryPriceV1RequestFrontendNodeGroups) SetStoragePerformanceLevel(v string) *QueryPriceV1RequestFrontendNodeGroups {
	s.StoragePerformanceLevel = &v
	return s
}

func (s *QueryPriceV1RequestFrontendNodeGroups) SetStorageSize(v int32) *QueryPriceV1RequestFrontendNodeGroups {
	s.StorageSize = &v
	return s
}

func (s *QueryPriceV1RequestFrontendNodeGroups) Validate() error {
	return dara.Validate(s)
}

type QueryPriceV1RequestObserverNodeGroups struct {
	// example:
	//
	// 8
	Cu *string `json:"cu,omitempty" xml:"cu,omitempty"`
	// example:
	//
	// 1
	DiskNumber *int32 `json:"diskNumber,omitempty" xml:"diskNumber,omitempty"`
	// example:
	//
	// null
	LocalStorageInstanceType *string `json:"localStorageInstanceType,omitempty" xml:"localStorageInstanceType,omitempty"`
	// example:
	//
	// 3
	ResidentNodeNumber *int32 `json:"residentNodeNumber,omitempty" xml:"residentNodeNumber,omitempty"`
	// example:
	//
	// standard
	SpecType *string `json:"specType,omitempty" xml:"specType,omitempty"`
	// example:
	//
	// pl1
	StoragePerformanceLevel *string `json:"storagePerformanceLevel,omitempty" xml:"storagePerformanceLevel,omitempty"`
	// example:
	//
	// 100
	StorageSize *int32 `json:"storageSize,omitempty" xml:"storageSize,omitempty"`
}

func (s QueryPriceV1RequestObserverNodeGroups) String() string {
	return dara.Prettify(s)
}

func (s QueryPriceV1RequestObserverNodeGroups) GoString() string {
	return s.String()
}

func (s *QueryPriceV1RequestObserverNodeGroups) GetCu() *string {
	return s.Cu
}

func (s *QueryPriceV1RequestObserverNodeGroups) GetDiskNumber() *int32 {
	return s.DiskNumber
}

func (s *QueryPriceV1RequestObserverNodeGroups) GetLocalStorageInstanceType() *string {
	return s.LocalStorageInstanceType
}

func (s *QueryPriceV1RequestObserverNodeGroups) GetResidentNodeNumber() *int32 {
	return s.ResidentNodeNumber
}

func (s *QueryPriceV1RequestObserverNodeGroups) GetSpecType() *string {
	return s.SpecType
}

func (s *QueryPriceV1RequestObserverNodeGroups) GetStoragePerformanceLevel() *string {
	return s.StoragePerformanceLevel
}

func (s *QueryPriceV1RequestObserverNodeGroups) GetStorageSize() *int32 {
	return s.StorageSize
}

func (s *QueryPriceV1RequestObserverNodeGroups) SetCu(v string) *QueryPriceV1RequestObserverNodeGroups {
	s.Cu = &v
	return s
}

func (s *QueryPriceV1RequestObserverNodeGroups) SetDiskNumber(v int32) *QueryPriceV1RequestObserverNodeGroups {
	s.DiskNumber = &v
	return s
}

func (s *QueryPriceV1RequestObserverNodeGroups) SetLocalStorageInstanceType(v string) *QueryPriceV1RequestObserverNodeGroups {
	s.LocalStorageInstanceType = &v
	return s
}

func (s *QueryPriceV1RequestObserverNodeGroups) SetResidentNodeNumber(v int32) *QueryPriceV1RequestObserverNodeGroups {
	s.ResidentNodeNumber = &v
	return s
}

func (s *QueryPriceV1RequestObserverNodeGroups) SetSpecType(v string) *QueryPriceV1RequestObserverNodeGroups {
	s.SpecType = &v
	return s
}

func (s *QueryPriceV1RequestObserverNodeGroups) SetStoragePerformanceLevel(v string) *QueryPriceV1RequestObserverNodeGroups {
	s.StoragePerformanceLevel = &v
	return s
}

func (s *QueryPriceV1RequestObserverNodeGroups) SetStorageSize(v int32) *QueryPriceV1RequestObserverNodeGroups {
	s.StorageSize = &v
	return s
}

func (s *QueryPriceV1RequestObserverNodeGroups) Validate() error {
	return dara.Validate(s)
}
