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
	// The list of cloud computer templates.
	Bundles []*DescribeBundlesResponseBodyBundles `json:"Bundles,omitempty" xml:"Bundles,omitempty" type:"Repeated"`
	// The token for the next query. If NextToken is empty, no more results exist.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6lu3PTF6h3zE8egwlYuv8M8
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
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
	// The cloud computer template ID.
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
	// The cloud computer templatetype.
	//
	// example:
	//
	// SYSTEM
	BundleType *string `json:"BundleType,omitempty" xml:"BundleType,omitempty"`
	// The time when the cloud computer template was created. The time is in the ISO 8601 standard in UTC.
	//
	// example:
	//
	// 2021-09-30T06:09Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The data cloud disk type.
	//
	// example:
	//
	// cloud_efficiency
	DataDiskCategory *string `json:"DataDiskCategory,omitempty" xml:"DataDiskCategory,omitempty"`
	// The description of the cloud computer template.
	//
	// example:
	//
	// Template for daily office use
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The cloud computer specifications.
	//
	// example:
	//
	// ecd.graphics.2xlarge
	DesktopType *string `json:"DesktopType,omitempty" xml:"DesktopType,omitempty"`
	// The details of the cloud computer specifications.
	DesktopTypeAttribute *DescribeBundlesResponseBodyBundlesDesktopTypeAttribute `json:"DesktopTypeAttribute,omitempty" xml:"DesktopTypeAttribute,omitempty" type:"Struct"`
	// The cloud computer instance family.
	//
	// example:
	//
	// eds.general
	DesktopTypeFamily *string `json:"DesktopTypeFamily,omitempty" xml:"DesktopTypeFamily,omitempty"`
	// The disk information.
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
	// The image status.
	//
	// example:
	//
	// Available
	ImageStatus *string `json:"ImageStatus,omitempty" xml:"ImageStatus,omitempty"`
	// The language of the image operating system.
	//
	// example:
	//
	// en-US
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The operating system type.
	//
	// example:
	//
	// Windows
	OsType *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	// The operating system platform information.
	//
	// example:
	//
	// Windows Server 2019
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// The protocol type.
	//
	// example:
	//
	// ASP
	ProtocolType *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	// The session type.
	//
	// example:
	//
	// 0
	SessionType *string `json:"SessionType,omitempty" xml:"SessionType,omitempty"`
	// The stock status. This parameter is returned when `CheckStock` is set to `True`.
	//
	// example:
	//
	// Sufficient
	StockState *string `json:"StockState,omitempty" xml:"StockState,omitempty"`
	// The system cloud disk type.
	//
	// example:
	//
	// cloud_efficiency
	SystemDiskCategory *string `json:"SystemDiskCategory,omitempty" xml:"SystemDiskCategory,omitempty"`
	// Specifies whether to enable disk encryption.
	//
	// example:
	//
	// false
	VolumeEncryptionEnabled *bool `json:"VolumeEncryptionEnabled,omitempty" xml:"VolumeEncryptionEnabled,omitempty"`
	// The ID of the KMS key used when disk encryption is enabled.
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
	// 4096
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
	// The disk performance level.
	//
	// example:
	//
	// PL0
	DiskPerformanceLevel *string `json:"DiskPerformanceLevel,omitempty" xml:"DiskPerformanceLevel,omitempty"`
	// The disk size. Unit: GiB.
	//
	// example:
	//
	// 150
	DiskSize *int32 `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	// The disk type.
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
