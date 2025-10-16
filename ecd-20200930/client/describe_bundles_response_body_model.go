// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBundlesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBundles(v []*DescribeBundlesResponseBodyBundles) *DescribeBundlesResponseBody
	GetBundles() []*DescribeBundlesResponseBodyBundles
	SetNextToken(v string) *DescribeBundlesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeBundlesResponseBody
	GetRequestId() *string
}

type DescribeBundlesResponseBody struct {
	// The cloud computer templates.
	Bundles []*DescribeBundlesResponseBodyBundles `json:"Bundles,omitempty" xml:"Bundles,omitempty" type:"Repeated"`
	// The token that is used for the next query. If this parameter is empty, all results have been returned.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6lu3PTF6h3zE8egwlYuv8M8
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// BCC854D8-5D1E-46D3-96EF-797A5DD36789
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeBundlesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBundlesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBundlesResponseBody) GetBundles() []*DescribeBundlesResponseBodyBundles {
	return s.Bundles
}

func (s *DescribeBundlesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeBundlesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBundlesResponseBody) SetBundles(v []*DescribeBundlesResponseBodyBundles) *DescribeBundlesResponseBody {
	s.Bundles = v
	return s
}

func (s *DescribeBundlesResponseBody) SetNextToken(v string) *DescribeBundlesResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeBundlesResponseBody) SetRequestId(v string) *DescribeBundlesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBundlesResponseBody) Validate() error {
	if s.Bundles != nil {
		for _, item := range s.Bundles {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeBundlesResponseBodyBundles struct {
	// The ID of the cloud computer template.
	//
	// example:
	//
	// bundle_ecd_graphics.2xlarge_s15d15_win2019
	BundleId *string `json:"BundleId,omitempty" xml:"BundleId,omitempty"`
	// The name of the cloud computer template.
	//
	// example:
	//
	// Advanced graphics with Windows 2019
	BundleName *string `json:"BundleName,omitempty" xml:"BundleName,omitempty"`
	// The type of the cloud computer template.
	//
	// Valid values:
	//
	// 	- SYSTEM: system template
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- CUSTOM: custom template
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// SYSTEM
	BundleType *string `json:"BundleType,omitempty" xml:"BundleType,omitempty"`
	// The time when the cloud computer template was created.
	//
	// example:
	//
	// 2021-09-30T06:09Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The category of the data disk.
	//
	// Valid values:
	//
	// 	- cloud_efficiency: the ultra disk
	//
	// 	- cloud_auto: the standard SSD.
	//
	// 	- cloud_essd: the ESSD. Take note that only specific cloud computer types support ESSDs.
	//
	// example:
	//
	// cloud_efficiency
	DataDiskCategory *string `json:"DataDiskCategory,omitempty" xml:"DataDiskCategory,omitempty"`
	// The description of the cloud computer template.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The instance type of the cloud computer.
	//
	// example:
	//
	// ecd.graphics.2xlarge
	DesktopType *string `json:"DesktopType,omitempty" xml:"DesktopType,omitempty"`
	// The details of the cloud computer instance type.
	DesktopTypeAttribute *DescribeBundlesResponseBodyBundlesDesktopTypeAttribute `json:"DesktopTypeAttribute,omitempty" xml:"DesktopTypeAttribute,omitempty" type:"Struct"`
	// The instance family of the cloud computer.
	//
	// Valid values:
	//
	// 	- eds.graphics: graphical instance family
	//
	// 	- eds.hf: instance family with a high clock speed
	//
	// 	- eds.general: general-purpose instance family
	//
	// example:
	//
	// eds.general
	DesktopTypeFamily *string `json:"DesktopTypeFamily,omitempty" xml:"DesktopTypeFamily,omitempty"`
	// Details of the disks.
	Disks []*DescribeBundlesResponseBodyBundlesDisks `json:"Disks,omitempty" xml:"Disks,omitempty" type:"Repeated"`
	// The image ID.
	//
	// example:
	//
	// desktopimage-windows-server-2019-64-ch-vgpu
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The image name.
	//
	// example:
	//
	// Windows server 2019 Chinese
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// The status of the image.
	//
	// example:
	//
	// Available
	ImageStatus *string `json:"ImageStatus,omitempty" xml:"ImageStatus,omitempty"`
	// The OS language of the image.
	//
	// Valid values:
	//
	// 	- en-US: English
	//
	// 	- zh-HK: Chinese, Traditional (Hong Kong, China)
	//
	// 	- zh-CN: Simplified Chinese
	//
	// 	- ja-JP: Japanese
	//
	// example:
	//
	// en-US
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The type of the OS.
	//
	// Valid values:
	//
	// 	- Linux
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Windows
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// Windows
	OsType *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	// The OS.
	//
	// Valid values:
	//
	// 	- Ubuntu
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Windows Server 2022
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- UOS
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- CentOS
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Windows Server 2019
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Windows Server 2016
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// Windows Server 2019
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// The protocol type.
	//
	// Valid values:
	//
	// 	- HDX: HDX protocol
	//
	// 	- ASP: in-house ASP
	//
	// example:
	//
	// ASP
	ProtocolType *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	// The session type.
	//
	// Valid values:
	//
	// 	- 0: single-session
	//
	// 	- 1: multi-session
	//
	// example:
	//
	// 0
	SessionType *string `json:"SessionType,omitempty" xml:"SessionType,omitempty"`
	// The inventory status of the cloud computer instance type. This parameter is returned only if you set the `CheckStock` parameter to `true`.
	//
	// example:
	//
	// Sufficient
	StockState *string `json:"StockState,omitempty" xml:"StockState,omitempty"`
	// The category of the system disk.
	//
	// Valid values:
	//
	// 	- cloud_efficiency: the ultra disk
	//
	// 	- cloud_auto: the standard SSD.
	//
	// 	- cloud_essd: the Enterprise SSD (ESSD). Take note that only specific cloud computer types support ESSDs.
	//
	// example:
	//
	// cloud_efficiency
	SystemDiskCategory *string `json:"SystemDiskCategory,omitempty" xml:"SystemDiskCategory,omitempty"`
	// Indicates whether disk encryption is enabled.
	//
	// example:
	//
	// false
	VolumeEncryptionEnabled *bool `json:"VolumeEncryptionEnabled,omitempty" xml:"VolumeEncryptionEnabled,omitempty"`
	// The ID of the Key Management Service (KMS) key that is used when disk encryption is enabled.
	//
	// example:
	//
	// 08c33a6f-4e0a-4a1b-a3fa-7ddfa1d4****
	VolumeEncryptionKey *string `json:"VolumeEncryptionKey,omitempty" xml:"VolumeEncryptionKey,omitempty"`
}

func (s DescribeBundlesResponseBodyBundles) String() string {
	return dara.Prettify(s)
}

func (s DescribeBundlesResponseBodyBundles) GoString() string {
	return s.String()
}

func (s *DescribeBundlesResponseBodyBundles) GetBundleId() *string {
	return s.BundleId
}

func (s *DescribeBundlesResponseBodyBundles) GetBundleName() *string {
	return s.BundleName
}

func (s *DescribeBundlesResponseBodyBundles) GetBundleType() *string {
	return s.BundleType
}

func (s *DescribeBundlesResponseBodyBundles) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeBundlesResponseBodyBundles) GetDataDiskCategory() *string {
	return s.DataDiskCategory
}

func (s *DescribeBundlesResponseBodyBundles) GetDescription() *string {
	return s.Description
}

func (s *DescribeBundlesResponseBodyBundles) GetDesktopType() *string {
	return s.DesktopType
}

func (s *DescribeBundlesResponseBodyBundles) GetDesktopTypeAttribute() *DescribeBundlesResponseBodyBundlesDesktopTypeAttribute {
	return s.DesktopTypeAttribute
}

func (s *DescribeBundlesResponseBodyBundles) GetDesktopTypeFamily() *string {
	return s.DesktopTypeFamily
}

func (s *DescribeBundlesResponseBodyBundles) GetDisks() []*DescribeBundlesResponseBodyBundlesDisks {
	return s.Disks
}

func (s *DescribeBundlesResponseBodyBundles) GetImageId() *string {
	return s.ImageId
}

func (s *DescribeBundlesResponseBodyBundles) GetImageName() *string {
	return s.ImageName
}

func (s *DescribeBundlesResponseBodyBundles) GetImageStatus() *string {
	return s.ImageStatus
}

func (s *DescribeBundlesResponseBodyBundles) GetLanguage() *string {
	return s.Language
}

func (s *DescribeBundlesResponseBodyBundles) GetOsType() *string {
	return s.OsType
}

func (s *DescribeBundlesResponseBodyBundles) GetPlatform() *string {
	return s.Platform
}

func (s *DescribeBundlesResponseBodyBundles) GetProtocolType() *string {
	return s.ProtocolType
}

func (s *DescribeBundlesResponseBodyBundles) GetSessionType() *string {
	return s.SessionType
}

func (s *DescribeBundlesResponseBodyBundles) GetStockState() *string {
	return s.StockState
}

func (s *DescribeBundlesResponseBodyBundles) GetSystemDiskCategory() *string {
	return s.SystemDiskCategory
}

func (s *DescribeBundlesResponseBodyBundles) GetVolumeEncryptionEnabled() *bool {
	return s.VolumeEncryptionEnabled
}

func (s *DescribeBundlesResponseBodyBundles) GetVolumeEncryptionKey() *string {
	return s.VolumeEncryptionKey
}

func (s *DescribeBundlesResponseBodyBundles) SetBundleId(v string) *DescribeBundlesResponseBodyBundles {
	s.BundleId = &v
	return s
}

func (s *DescribeBundlesResponseBodyBundles) SetBundleName(v string) *DescribeBundlesResponseBodyBundles {
	s.BundleName = &v
	return s
}

func (s *DescribeBundlesResponseBodyBundles) SetBundleType(v string) *DescribeBundlesResponseBodyBundles {
	s.BundleType = &v
	return s
}

func (s *DescribeBundlesResponseBodyBundles) SetCreationTime(v string) *DescribeBundlesResponseBodyBundles {
	s.CreationTime = &v
	return s
}

func (s *DescribeBundlesResponseBodyBundles) SetDataDiskCategory(v string) *DescribeBundlesResponseBodyBundles {
	s.DataDiskCategory = &v
	return s
}

func (s *DescribeBundlesResponseBodyBundles) SetDescription(v string) *DescribeBundlesResponseBodyBundles {
	s.Description = &v
	return s
}

func (s *DescribeBundlesResponseBodyBundles) SetDesktopType(v string) *DescribeBundlesResponseBodyBundles {
	s.DesktopType = &v
	return s
}

func (s *DescribeBundlesResponseBodyBundles) SetDesktopTypeAttribute(v *DescribeBundlesResponseBodyBundlesDesktopTypeAttribute) *DescribeBundlesResponseBodyBundles {
	s.DesktopTypeAttribute = v
	return s
}

func (s *DescribeBundlesResponseBodyBundles) SetDesktopTypeFamily(v string) *DescribeBundlesResponseBodyBundles {
	s.DesktopTypeFamily = &v
	return s
}

func (s *DescribeBundlesResponseBodyBundles) SetDisks(v []*DescribeBundlesResponseBodyBundlesDisks) *DescribeBundlesResponseBodyBundles {
	s.Disks = v
	return s
}

func (s *DescribeBundlesResponseBodyBundles) SetImageId(v string) *DescribeBundlesResponseBodyBundles {
	s.ImageId = &v
	return s
}

func (s *DescribeBundlesResponseBodyBundles) SetImageName(v string) *DescribeBundlesResponseBodyBundles {
	s.ImageName = &v
	return s
}

func (s *DescribeBundlesResponseBodyBundles) SetImageStatus(v string) *DescribeBundlesResponseBodyBundles {
	s.ImageStatus = &v
	return s
}

func (s *DescribeBundlesResponseBodyBundles) SetLanguage(v string) *DescribeBundlesResponseBodyBundles {
	s.Language = &v
	return s
}

func (s *DescribeBundlesResponseBodyBundles) SetOsType(v string) *DescribeBundlesResponseBodyBundles {
	s.OsType = &v
	return s
}

func (s *DescribeBundlesResponseBodyBundles) SetPlatform(v string) *DescribeBundlesResponseBodyBundles {
	s.Platform = &v
	return s
}

func (s *DescribeBundlesResponseBodyBundles) SetProtocolType(v string) *DescribeBundlesResponseBodyBundles {
	s.ProtocolType = &v
	return s
}

func (s *DescribeBundlesResponseBodyBundles) SetSessionType(v string) *DescribeBundlesResponseBodyBundles {
	s.SessionType = &v
	return s
}

func (s *DescribeBundlesResponseBodyBundles) SetStockState(v string) *DescribeBundlesResponseBodyBundles {
	s.StockState = &v
	return s
}

func (s *DescribeBundlesResponseBodyBundles) SetSystemDiskCategory(v string) *DescribeBundlesResponseBodyBundles {
	s.SystemDiskCategory = &v
	return s
}

func (s *DescribeBundlesResponseBodyBundles) SetVolumeEncryptionEnabled(v bool) *DescribeBundlesResponseBodyBundles {
	s.VolumeEncryptionEnabled = &v
	return s
}

func (s *DescribeBundlesResponseBodyBundles) SetVolumeEncryptionKey(v string) *DescribeBundlesResponseBodyBundles {
	s.VolumeEncryptionKey = &v
	return s
}

func (s *DescribeBundlesResponseBodyBundles) Validate() error {
	if s.DesktopTypeAttribute != nil {
		if err := s.DesktopTypeAttribute.Validate(); err != nil {
			return err
		}
	}
	if s.Disks != nil {
		for _, item := range s.Disks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeBundlesResponseBodyBundlesDesktopTypeAttribute struct {
	// The number of vCPUs.
	//
	// example:
	//
	// 10
	CpuCount *int32 `json:"CpuCount,omitempty" xml:"CpuCount,omitempty"`
	// The number of GPUs.
	//
	// example:
	//
	// 0.5
	GpuCount *float32 `json:"GpuCount,omitempty" xml:"GpuCount,omitempty"`
	// The GPU type.
	//
	// example:
	//
	// NVIDIA T4
	GpuSpec *string `json:"GpuSpec,omitempty" xml:"GpuSpec,omitempty"`
	// The memory size. Unit: MiB.
	//
	// example:
	//
	// 47104
	MemorySize *int32 `json:"MemorySize,omitempty" xml:"MemorySize,omitempty"`
}

func (s DescribeBundlesResponseBodyBundlesDesktopTypeAttribute) String() string {
	return dara.Prettify(s)
}

func (s DescribeBundlesResponseBodyBundlesDesktopTypeAttribute) GoString() string {
	return s.String()
}

func (s *DescribeBundlesResponseBodyBundlesDesktopTypeAttribute) GetCpuCount() *int32 {
	return s.CpuCount
}

func (s *DescribeBundlesResponseBodyBundlesDesktopTypeAttribute) GetGpuCount() *float32 {
	return s.GpuCount
}

func (s *DescribeBundlesResponseBodyBundlesDesktopTypeAttribute) GetGpuSpec() *string {
	return s.GpuSpec
}

func (s *DescribeBundlesResponseBodyBundlesDesktopTypeAttribute) GetMemorySize() *int32 {
	return s.MemorySize
}

func (s *DescribeBundlesResponseBodyBundlesDesktopTypeAttribute) SetCpuCount(v int32) *DescribeBundlesResponseBodyBundlesDesktopTypeAttribute {
	s.CpuCount = &v
	return s
}

func (s *DescribeBundlesResponseBodyBundlesDesktopTypeAttribute) SetGpuCount(v float32) *DescribeBundlesResponseBodyBundlesDesktopTypeAttribute {
	s.GpuCount = &v
	return s
}

func (s *DescribeBundlesResponseBodyBundlesDesktopTypeAttribute) SetGpuSpec(v string) *DescribeBundlesResponseBodyBundlesDesktopTypeAttribute {
	s.GpuSpec = &v
	return s
}

func (s *DescribeBundlesResponseBodyBundlesDesktopTypeAttribute) SetMemorySize(v int32) *DescribeBundlesResponseBodyBundlesDesktopTypeAttribute {
	s.MemorySize = &v
	return s
}

func (s *DescribeBundlesResponseBodyBundlesDesktopTypeAttribute) Validate() error {
	return dara.Validate(s)
}

type DescribeBundlesResponseBodyBundlesDisks struct {
	// The PL of the disk.
	//
	// Valid values:
	//
	// 	- PL1
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- PL0
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- PL3
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- PL2
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// PL0
	DiskPerformanceLevel *string `json:"DiskPerformanceLevel,omitempty" xml:"DiskPerformanceLevel,omitempty"`
	// The size of the disk. Unit: GiB.
	//
	// example:
	//
	// 150
	DiskSize *int32 `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	// The type of the disk.
	//
	// Valid values:
	//
	// 	- SYSTEM: system disk
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- DATA: data disk
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// SYSTEM
	DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
}

func (s DescribeBundlesResponseBodyBundlesDisks) String() string {
	return dara.Prettify(s)
}

func (s DescribeBundlesResponseBodyBundlesDisks) GoString() string {
	return s.String()
}

func (s *DescribeBundlesResponseBodyBundlesDisks) GetDiskPerformanceLevel() *string {
	return s.DiskPerformanceLevel
}

func (s *DescribeBundlesResponseBodyBundlesDisks) GetDiskSize() *int32 {
	return s.DiskSize
}

func (s *DescribeBundlesResponseBodyBundlesDisks) GetDiskType() *string {
	return s.DiskType
}

func (s *DescribeBundlesResponseBodyBundlesDisks) SetDiskPerformanceLevel(v string) *DescribeBundlesResponseBodyBundlesDisks {
	s.DiskPerformanceLevel = &v
	return s
}

func (s *DescribeBundlesResponseBodyBundlesDisks) SetDiskSize(v int32) *DescribeBundlesResponseBodyBundlesDisks {
	s.DiskSize = &v
	return s
}

func (s *DescribeBundlesResponseBodyBundlesDisks) SetDiskType(v string) *DescribeBundlesResponseBodyBundlesDisks {
	s.DiskType = &v
	return s
}

func (s *DescribeBundlesResponseBodyBundlesDisks) Validate() error {
	return dara.Validate(s)
}
