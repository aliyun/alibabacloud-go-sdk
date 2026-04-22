// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEmbodiedAIPlatformResourceUsageInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAcuDetails(v *GetEmbodiedAIPlatformResourceUsageInfoResponseBodyAcuDetails) *GetEmbodiedAIPlatformResourceUsageInfoResponseBody
	GetAcuDetails() *GetEmbodiedAIPlatformResourceUsageInfoResponseBodyAcuDetails
	SetGpuDetails(v []*GetEmbodiedAIPlatformResourceUsageInfoResponseBodyGpuDetails) *GetEmbodiedAIPlatformResourceUsageInfoResponseBody
	GetGpuDetails() []*GetEmbodiedAIPlatformResourceUsageInfoResponseBodyGpuDetails
	SetMaxRegisteredDevices(v int64) *GetEmbodiedAIPlatformResourceUsageInfoResponseBody
	GetMaxRegisteredDevices() *int64
	SetRegisteredDeviceCount(v int64) *GetEmbodiedAIPlatformResourceUsageInfoResponseBody
	GetRegisteredDeviceCount() *int64
	SetRequestId(v string) *GetEmbodiedAIPlatformResourceUsageInfoResponseBody
	GetRequestId() *string
	SetSlbTraffic(v *GetEmbodiedAIPlatformResourceUsageInfoResponseBodySlbTraffic) *GetEmbodiedAIPlatformResourceUsageInfoResponseBody
	GetSlbTraffic() *GetEmbodiedAIPlatformResourceUsageInfoResponseBodySlbTraffic
	SetStorageUsage(v *GetEmbodiedAIPlatformResourceUsageInfoResponseBodyStorageUsage) *GetEmbodiedAIPlatformResourceUsageInfoResponseBody
	GetStorageUsage() *GetEmbodiedAIPlatformResourceUsageInfoResponseBodyStorageUsage
}

type GetEmbodiedAIPlatformResourceUsageInfoResponseBody struct {
	AcuDetails *GetEmbodiedAIPlatformResourceUsageInfoResponseBodyAcuDetails   `json:"AcuDetails,omitempty" xml:"AcuDetails,omitempty" type:"Struct"`
	GpuDetails []*GetEmbodiedAIPlatformResourceUsageInfoResponseBodyGpuDetails `json:"GpuDetails,omitempty" xml:"GpuDetails,omitempty" type:"Repeated"`
	// example:
	//
	// 3
	MaxRegisteredDevices *int64 `json:"MaxRegisteredDevices,omitempty" xml:"MaxRegisteredDevices,omitempty"`
	// example:
	//
	// 1
	RegisteredDeviceCount *int64 `json:"RegisteredDeviceCount,omitempty" xml:"RegisteredDeviceCount,omitempty"`
	// Id of the request
	//
	// example:
	//
	// B47EED99-BFA5-529D-8D85-A6642421D390
	RequestId    *string                                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SlbTraffic   *GetEmbodiedAIPlatformResourceUsageInfoResponseBodySlbTraffic   `json:"SlbTraffic,omitempty" xml:"SlbTraffic,omitempty" type:"Struct"`
	StorageUsage *GetEmbodiedAIPlatformResourceUsageInfoResponseBodyStorageUsage `json:"StorageUsage,omitempty" xml:"StorageUsage,omitempty" type:"Struct"`
}

func (s GetEmbodiedAIPlatformResourceUsageInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetEmbodiedAIPlatformResourceUsageInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetEmbodiedAIPlatformResourceUsageInfoResponseBody) GetAcuDetails() *GetEmbodiedAIPlatformResourceUsageInfoResponseBodyAcuDetails {
	return s.AcuDetails
}

func (s *GetEmbodiedAIPlatformResourceUsageInfoResponseBody) GetGpuDetails() []*GetEmbodiedAIPlatformResourceUsageInfoResponseBodyGpuDetails {
	return s.GpuDetails
}

func (s *GetEmbodiedAIPlatformResourceUsageInfoResponseBody) GetMaxRegisteredDevices() *int64 {
	return s.MaxRegisteredDevices
}

func (s *GetEmbodiedAIPlatformResourceUsageInfoResponseBody) GetRegisteredDeviceCount() *int64 {
	return s.RegisteredDeviceCount
}

func (s *GetEmbodiedAIPlatformResourceUsageInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetEmbodiedAIPlatformResourceUsageInfoResponseBody) GetSlbTraffic() *GetEmbodiedAIPlatformResourceUsageInfoResponseBodySlbTraffic {
	return s.SlbTraffic
}

func (s *GetEmbodiedAIPlatformResourceUsageInfoResponseBody) GetStorageUsage() *GetEmbodiedAIPlatformResourceUsageInfoResponseBodyStorageUsage {
	return s.StorageUsage
}

func (s *GetEmbodiedAIPlatformResourceUsageInfoResponseBody) SetAcuDetails(v *GetEmbodiedAIPlatformResourceUsageInfoResponseBodyAcuDetails) *GetEmbodiedAIPlatformResourceUsageInfoResponseBody {
	s.AcuDetails = v
	return s
}

func (s *GetEmbodiedAIPlatformResourceUsageInfoResponseBody) SetGpuDetails(v []*GetEmbodiedAIPlatformResourceUsageInfoResponseBodyGpuDetails) *GetEmbodiedAIPlatformResourceUsageInfoResponseBody {
	s.GpuDetails = v
	return s
}

func (s *GetEmbodiedAIPlatformResourceUsageInfoResponseBody) SetMaxRegisteredDevices(v int64) *GetEmbodiedAIPlatformResourceUsageInfoResponseBody {
	s.MaxRegisteredDevices = &v
	return s
}

func (s *GetEmbodiedAIPlatformResourceUsageInfoResponseBody) SetRegisteredDeviceCount(v int64) *GetEmbodiedAIPlatformResourceUsageInfoResponseBody {
	s.RegisteredDeviceCount = &v
	return s
}

func (s *GetEmbodiedAIPlatformResourceUsageInfoResponseBody) SetRequestId(v string) *GetEmbodiedAIPlatformResourceUsageInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEmbodiedAIPlatformResourceUsageInfoResponseBody) SetSlbTraffic(v *GetEmbodiedAIPlatformResourceUsageInfoResponseBodySlbTraffic) *GetEmbodiedAIPlatformResourceUsageInfoResponseBody {
	s.SlbTraffic = v
	return s
}

func (s *GetEmbodiedAIPlatformResourceUsageInfoResponseBody) SetStorageUsage(v *GetEmbodiedAIPlatformResourceUsageInfoResponseBodyStorageUsage) *GetEmbodiedAIPlatformResourceUsageInfoResponseBody {
	s.StorageUsage = v
	return s
}

func (s *GetEmbodiedAIPlatformResourceUsageInfoResponseBody) Validate() error {
	if s.AcuDetails != nil {
		if err := s.AcuDetails.Validate(); err != nil {
			return err
		}
	}
	if s.GpuDetails != nil {
		for _, item := range s.GpuDetails {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SlbTraffic != nil {
		if err := s.SlbTraffic.Validate(); err != nil {
			return err
		}
	}
	if s.StorageUsage != nil {
		if err := s.StorageUsage.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetEmbodiedAIPlatformResourceUsageInfoResponseBodyAcuDetails struct {
	EapWebserverACUCount *float64 `json:"EapWebserverACUCount,omitempty" xml:"EapWebserverACUCount,omitempty"`
	RayHeadACUCount      *float64 `json:"RayHeadACUCount,omitempty" xml:"RayHeadACUCount,omitempty"`
	RayWorkerACUCount    *float64 `json:"RayWorkerACUCount,omitempty" xml:"RayWorkerACUCount,omitempty"`
	TotalACUCount        *float64 `json:"TotalACUCount,omitempty" xml:"TotalACUCount,omitempty"`
}

func (s GetEmbodiedAIPlatformResourceUsageInfoResponseBodyAcuDetails) String() string {
	return dara.Prettify(s)
}

func (s GetEmbodiedAIPlatformResourceUsageInfoResponseBodyAcuDetails) GoString() string {
	return s.String()
}

func (s *GetEmbodiedAIPlatformResourceUsageInfoResponseBodyAcuDetails) GetEapWebserverACUCount() *float64 {
	return s.EapWebserverACUCount
}

func (s *GetEmbodiedAIPlatformResourceUsageInfoResponseBodyAcuDetails) GetRayHeadACUCount() *float64 {
	return s.RayHeadACUCount
}

func (s *GetEmbodiedAIPlatformResourceUsageInfoResponseBodyAcuDetails) GetRayWorkerACUCount() *float64 {
	return s.RayWorkerACUCount
}

func (s *GetEmbodiedAIPlatformResourceUsageInfoResponseBodyAcuDetails) GetTotalACUCount() *float64 {
	return s.TotalACUCount
}

func (s *GetEmbodiedAIPlatformResourceUsageInfoResponseBodyAcuDetails) SetEapWebserverACUCount(v float64) *GetEmbodiedAIPlatformResourceUsageInfoResponseBodyAcuDetails {
	s.EapWebserverACUCount = &v
	return s
}

func (s *GetEmbodiedAIPlatformResourceUsageInfoResponseBodyAcuDetails) SetRayHeadACUCount(v float64) *GetEmbodiedAIPlatformResourceUsageInfoResponseBodyAcuDetails {
	s.RayHeadACUCount = &v
	return s
}

func (s *GetEmbodiedAIPlatformResourceUsageInfoResponseBodyAcuDetails) SetRayWorkerACUCount(v float64) *GetEmbodiedAIPlatformResourceUsageInfoResponseBodyAcuDetails {
	s.RayWorkerACUCount = &v
	return s
}

func (s *GetEmbodiedAIPlatformResourceUsageInfoResponseBodyAcuDetails) SetTotalACUCount(v float64) *GetEmbodiedAIPlatformResourceUsageInfoResponseBodyAcuDetails {
	s.TotalACUCount = &v
	return s
}

func (s *GetEmbodiedAIPlatformResourceUsageInfoResponseBodyAcuDetails) Validate() error {
	return dara.Validate(s)
}

type GetEmbodiedAIPlatformResourceUsageInfoResponseBodyGpuDetails struct {
	AllocatedUnit *int32 `json:"AllocatedUnit,omitempty" xml:"AllocatedUnit,omitempty"`
	// example:
	//
	// ADB.MLLarge.2
	GpuModel *string `json:"GpuModel,omitempty" xml:"GpuModel,omitempty"`
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetEmbodiedAIPlatformResourceUsageInfoResponseBodyGpuDetails) String() string {
	return dara.Prettify(s)
}

func (s GetEmbodiedAIPlatformResourceUsageInfoResponseBodyGpuDetails) GoString() string {
	return s.String()
}

func (s *GetEmbodiedAIPlatformResourceUsageInfoResponseBodyGpuDetails) GetAllocatedUnit() *int32 {
	return s.AllocatedUnit
}

func (s *GetEmbodiedAIPlatformResourceUsageInfoResponseBodyGpuDetails) GetGpuModel() *string {
	return s.GpuModel
}

func (s *GetEmbodiedAIPlatformResourceUsageInfoResponseBodyGpuDetails) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *GetEmbodiedAIPlatformResourceUsageInfoResponseBodyGpuDetails) SetAllocatedUnit(v int32) *GetEmbodiedAIPlatformResourceUsageInfoResponseBodyGpuDetails {
	s.AllocatedUnit = &v
	return s
}

func (s *GetEmbodiedAIPlatformResourceUsageInfoResponseBodyGpuDetails) SetGpuModel(v string) *GetEmbodiedAIPlatformResourceUsageInfoResponseBodyGpuDetails {
	s.GpuModel = &v
	return s
}

func (s *GetEmbodiedAIPlatformResourceUsageInfoResponseBodyGpuDetails) SetTotalCount(v int32) *GetEmbodiedAIPlatformResourceUsageInfoResponseBodyGpuDetails {
	s.TotalCount = &v
	return s
}

func (s *GetEmbodiedAIPlatformResourceUsageInfoResponseBodyGpuDetails) Validate() error {
	return dara.Validate(s)
}

type GetEmbodiedAIPlatformResourceUsageInfoResponseBodySlbTraffic struct {
	// example:
	//
	// 0
	TotalBytesIn *int64 `json:"TotalBytesIn,omitempty" xml:"TotalBytesIn,omitempty"`
	// example:
	//
	// 0
	TotalBytesOut *int64 `json:"TotalBytesOut,omitempty" xml:"TotalBytesOut,omitempty"`
}

func (s GetEmbodiedAIPlatformResourceUsageInfoResponseBodySlbTraffic) String() string {
	return dara.Prettify(s)
}

func (s GetEmbodiedAIPlatformResourceUsageInfoResponseBodySlbTraffic) GoString() string {
	return s.String()
}

func (s *GetEmbodiedAIPlatformResourceUsageInfoResponseBodySlbTraffic) GetTotalBytesIn() *int64 {
	return s.TotalBytesIn
}

func (s *GetEmbodiedAIPlatformResourceUsageInfoResponseBodySlbTraffic) GetTotalBytesOut() *int64 {
	return s.TotalBytesOut
}

func (s *GetEmbodiedAIPlatformResourceUsageInfoResponseBodySlbTraffic) SetTotalBytesIn(v int64) *GetEmbodiedAIPlatformResourceUsageInfoResponseBodySlbTraffic {
	s.TotalBytesIn = &v
	return s
}

func (s *GetEmbodiedAIPlatformResourceUsageInfoResponseBodySlbTraffic) SetTotalBytesOut(v int64) *GetEmbodiedAIPlatformResourceUsageInfoResponseBodySlbTraffic {
	s.TotalBytesOut = &v
	return s
}

func (s *GetEmbodiedAIPlatformResourceUsageInfoResponseBodySlbTraffic) Validate() error {
	return dara.Validate(s)
}

type GetEmbodiedAIPlatformResourceUsageInfoResponseBodyStorageUsage struct {
	Nas *GetEmbodiedAIPlatformResourceUsageInfoResponseBodyStorageUsageNas `json:"Nas,omitempty" xml:"Nas,omitempty" type:"Struct"`
	Oss *GetEmbodiedAIPlatformResourceUsageInfoResponseBodyStorageUsageOss `json:"Oss,omitempty" xml:"Oss,omitempty" type:"Struct"`
}

func (s GetEmbodiedAIPlatformResourceUsageInfoResponseBodyStorageUsage) String() string {
	return dara.Prettify(s)
}

func (s GetEmbodiedAIPlatformResourceUsageInfoResponseBodyStorageUsage) GoString() string {
	return s.String()
}

func (s *GetEmbodiedAIPlatformResourceUsageInfoResponseBodyStorageUsage) GetNas() *GetEmbodiedAIPlatformResourceUsageInfoResponseBodyStorageUsageNas {
	return s.Nas
}

func (s *GetEmbodiedAIPlatformResourceUsageInfoResponseBodyStorageUsage) GetOss() *GetEmbodiedAIPlatformResourceUsageInfoResponseBodyStorageUsageOss {
	return s.Oss
}

func (s *GetEmbodiedAIPlatformResourceUsageInfoResponseBodyStorageUsage) SetNas(v *GetEmbodiedAIPlatformResourceUsageInfoResponseBodyStorageUsageNas) *GetEmbodiedAIPlatformResourceUsageInfoResponseBodyStorageUsage {
	s.Nas = v
	return s
}

func (s *GetEmbodiedAIPlatformResourceUsageInfoResponseBodyStorageUsage) SetOss(v *GetEmbodiedAIPlatformResourceUsageInfoResponseBodyStorageUsageOss) *GetEmbodiedAIPlatformResourceUsageInfoResponseBodyStorageUsage {
	s.Oss = v
	return s
}

func (s *GetEmbodiedAIPlatformResourceUsageInfoResponseBodyStorageUsage) Validate() error {
	if s.Nas != nil {
		if err := s.Nas.Validate(); err != nil {
			return err
		}
	}
	if s.Oss != nil {
		if err := s.Oss.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetEmbodiedAIPlatformResourceUsageInfoResponseBodyStorageUsageNas struct {
	// example:
	//
	// 0
	MeteredSize *int64 `json:"MeteredSize,omitempty" xml:"MeteredSize,omitempty"`
}

func (s GetEmbodiedAIPlatformResourceUsageInfoResponseBodyStorageUsageNas) String() string {
	return dara.Prettify(s)
}

func (s GetEmbodiedAIPlatformResourceUsageInfoResponseBodyStorageUsageNas) GoString() string {
	return s.String()
}

func (s *GetEmbodiedAIPlatformResourceUsageInfoResponseBodyStorageUsageNas) GetMeteredSize() *int64 {
	return s.MeteredSize
}

func (s *GetEmbodiedAIPlatformResourceUsageInfoResponseBodyStorageUsageNas) SetMeteredSize(v int64) *GetEmbodiedAIPlatformResourceUsageInfoResponseBodyStorageUsageNas {
	s.MeteredSize = &v
	return s
}

func (s *GetEmbodiedAIPlatformResourceUsageInfoResponseBodyStorageUsageNas) Validate() error {
	return dara.Validate(s)
}

type GetEmbodiedAIPlatformResourceUsageInfoResponseBodyStorageUsageOss struct {
	// example:
	//
	// 0
	StandardStorageSize *int64 `json:"StandardStorageSize,omitempty" xml:"StandardStorageSize,omitempty"`
}

func (s GetEmbodiedAIPlatformResourceUsageInfoResponseBodyStorageUsageOss) String() string {
	return dara.Prettify(s)
}

func (s GetEmbodiedAIPlatformResourceUsageInfoResponseBodyStorageUsageOss) GoString() string {
	return s.String()
}

func (s *GetEmbodiedAIPlatformResourceUsageInfoResponseBodyStorageUsageOss) GetStandardStorageSize() *int64 {
	return s.StandardStorageSize
}

func (s *GetEmbodiedAIPlatformResourceUsageInfoResponseBodyStorageUsageOss) SetStandardStorageSize(v int64) *GetEmbodiedAIPlatformResourceUsageInfoResponseBodyStorageUsageOss {
	s.StandardStorageSize = &v
	return s
}

func (s *GetEmbodiedAIPlatformResourceUsageInfoResponseBodyStorageUsageOss) Validate() error {
	return dara.Validate(s)
}
