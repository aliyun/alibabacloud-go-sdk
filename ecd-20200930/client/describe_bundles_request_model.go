// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBundlesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBundleId(v []*string) *DescribeBundlesRequest
	GetBundleId() []*string
	SetBundleType(v string) *DescribeBundlesRequest
	GetBundleType() *string
	SetCheckStock(v bool) *DescribeBundlesRequest
	GetCheckStock() *bool
	SetCpuCount(v int32) *DescribeBundlesRequest
	GetCpuCount() *int32
	SetDesktopTypeFamily(v string) *DescribeBundlesRequest
	GetDesktopTypeFamily() *string
	SetFotaChannel(v string) *DescribeBundlesRequest
	GetFotaChannel() *string
	SetFromDesktopGroup(v bool) *DescribeBundlesRequest
	GetFromDesktopGroup() *bool
	SetGpuCount(v float32) *DescribeBundlesRequest
	GetGpuCount() *float32
	SetGpuDriverType(v string) *DescribeBundlesRequest
	GetGpuDriverType() *string
	SetImageId(v []*string) *DescribeBundlesRequest
	GetImageId() []*string
	SetMaxResults(v int32) *DescribeBundlesRequest
	GetMaxResults() *int32
	SetMemorySize(v int32) *DescribeBundlesRequest
	GetMemorySize() *int32
	SetNextToken(v string) *DescribeBundlesRequest
	GetNextToken() *string
	SetOsType(v string) *DescribeBundlesRequest
	GetOsType() *string
	SetProtocolType(v string) *DescribeBundlesRequest
	GetProtocolType() *string
	SetRegionId(v string) *DescribeBundlesRequest
	GetRegionId() *string
	SetScope(v string) *DescribeBundlesRequest
	GetScope() *string
	SetSelectedBundle(v bool) *DescribeBundlesRequest
	GetSelectedBundle() *bool
	SetSessionType(v string) *DescribeBundlesRequest
	GetSessionType() *string
	SetSupportMultiSession(v bool) *DescribeBundlesRequest
	GetSupportMultiSession() *bool
	SetVolumeEncryptionEnabled(v bool) *DescribeBundlesRequest
	GetVolumeEncryptionEnabled() *bool
}

type DescribeBundlesRequest struct {
	// The IDs of the cloud computer templates. You can specify 1 to 100 IDs.
	//
	// example:
	//
	// bundle_ecd_graphics.2xlarge_s15d15_win2019
	BundleId []*string `json:"BundleId,omitempty" xml:"BundleId,omitempty" type:"Repeated"`
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
	// Specifies whether to query the inventory status of the cloud computer instance type.
	//
	// example:
	//
	// true
	CheckStock *bool `json:"CheckStock,omitempty" xml:"CheckStock,omitempty"`
	// The number of vCPUs contained in the cloud computer instance type.
	//
	// example:
	//
	// 2
	CpuCount *int32 `json:"CpuCount,omitempty" xml:"CpuCount,omitempty"`
	// The instance family of the cloud computers.
	//
	// Valid values:
	//
	// 	- eds.graphics: graphical instance families
	//
	// 	- eds.hf: instance families with high clock speeds
	//
	// 	- eds.general: general-purpose instance families
	//
	// example:
	//
	// eds.general
	DesktopTypeFamily *string `json:"DesktopTypeFamily,omitempty" xml:"DesktopTypeFamily,omitempty"`
	// >  This parameter is not available for public use.
	//
	// example:
	//
	// This parameter is now in invitational preview and unavailable.
	FotaChannel *string `json:"FotaChannel,omitempty" xml:"FotaChannel,omitempty"`
	// Specifies whether the cloud computers in the template belong to a cloud computer pool.
	//
	// Valid values:
	//
	// 	- true
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- false
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// false
	FromDesktopGroup *bool `json:"FromDesktopGroup,omitempty" xml:"FromDesktopGroup,omitempty"`
	// The number of GPUs contained in the cloud computer instance type.
	//
	// example:
	//
	// 1
	GpuCount *float32 `json:"GpuCount,omitempty" xml:"GpuCount,omitempty"`
	// The GPU driver type.
	//
	// Valid values:
	//
	// 	- T4
	//
	// 	- A10
	//
	// 	- G28
	//
	// 	- G39
	//
	// example:
	//
	// T4
	GpuDriverType *string `json:"GpuDriverType,omitempty" xml:"GpuDriverType,omitempty"`
	// The image IDs.
	ImageId []*string `json:"ImageId,omitempty" xml:"ImageId,omitempty" type:"Repeated"`
	// The number of entries to return on each page.
	//
	// Maximum value: 100.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The memory size of the cloud computer instance type. Unit: GiB.
	//
	// example:
	//
	// 4
	MemorySize *int32 `json:"MemorySize,omitempty" xml:"MemorySize,omitempty"`
	// The token that is used to start the next query.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
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
	// The protocol type.
	//
	// Valid values:
	//
	// 	- HDX: High-definition Experience (HDX) protocol
	//
	// 	- ASP: in-house Adaptive Streaming Protocol (ASP) (recommend)
	//
	// example:
	//
	// ASP
	ProtocolType *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	// The region ID. You can call the [DescribeRegions](~~DescribeRegions~~) operation to query the regions supported by Elastic Desktop Service (EDS).
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The scenario to use the image.
	//
	// example:
	//
	// FastBuy
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// The desktop template that is selected based on specific criteria.
	//
	// example:
	//
	// true
	SelectedBundle *bool `json:"SelectedBundle,omitempty" xml:"SelectedBundle,omitempty"`
	// The type of the session. Valide values:
	//
	// - SingleSession
	//
	// - MultipleSession
	//
	// example:
	//
	// SingleSession
	SessionType *string `json:"SessionType,omitempty" xml:"SessionType,omitempty"`
	// Specifies whether to return multi-session cloud computer templates. Default value: false.
	//
	// example:
	//
	// false
	SupportMultiSession *bool `json:"SupportMultiSession,omitempty" xml:"SupportMultiSession,omitempty"`
	// Specifies whether to enable disk encryption.
	//
	// example:
	//
	// false
	VolumeEncryptionEnabled *bool `json:"VolumeEncryptionEnabled,omitempty" xml:"VolumeEncryptionEnabled,omitempty"`
}

func (s DescribeBundlesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeBundlesRequest) GoString() string {
	return s.String()
}

func (s *DescribeBundlesRequest) GetBundleId() []*string {
	return s.BundleId
}

func (s *DescribeBundlesRequest) GetBundleType() *string {
	return s.BundleType
}

func (s *DescribeBundlesRequest) GetCheckStock() *bool {
	return s.CheckStock
}

func (s *DescribeBundlesRequest) GetCpuCount() *int32 {
	return s.CpuCount
}

func (s *DescribeBundlesRequest) GetDesktopTypeFamily() *string {
	return s.DesktopTypeFamily
}

func (s *DescribeBundlesRequest) GetFotaChannel() *string {
	return s.FotaChannel
}

func (s *DescribeBundlesRequest) GetFromDesktopGroup() *bool {
	return s.FromDesktopGroup
}

func (s *DescribeBundlesRequest) GetGpuCount() *float32 {
	return s.GpuCount
}

func (s *DescribeBundlesRequest) GetGpuDriverType() *string {
	return s.GpuDriverType
}

func (s *DescribeBundlesRequest) GetImageId() []*string {
	return s.ImageId
}

func (s *DescribeBundlesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeBundlesRequest) GetMemorySize() *int32 {
	return s.MemorySize
}

func (s *DescribeBundlesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeBundlesRequest) GetOsType() *string {
	return s.OsType
}

func (s *DescribeBundlesRequest) GetProtocolType() *string {
	return s.ProtocolType
}

func (s *DescribeBundlesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeBundlesRequest) GetScope() *string {
	return s.Scope
}

func (s *DescribeBundlesRequest) GetSelectedBundle() *bool {
	return s.SelectedBundle
}

func (s *DescribeBundlesRequest) GetSessionType() *string {
	return s.SessionType
}

func (s *DescribeBundlesRequest) GetSupportMultiSession() *bool {
	return s.SupportMultiSession
}

func (s *DescribeBundlesRequest) GetVolumeEncryptionEnabled() *bool {
	return s.VolumeEncryptionEnabled
}

func (s *DescribeBundlesRequest) SetBundleId(v []*string) *DescribeBundlesRequest {
	s.BundleId = v
	return s
}

func (s *DescribeBundlesRequest) SetBundleType(v string) *DescribeBundlesRequest {
	s.BundleType = &v
	return s
}

func (s *DescribeBundlesRequest) SetCheckStock(v bool) *DescribeBundlesRequest {
	s.CheckStock = &v
	return s
}

func (s *DescribeBundlesRequest) SetCpuCount(v int32) *DescribeBundlesRequest {
	s.CpuCount = &v
	return s
}

func (s *DescribeBundlesRequest) SetDesktopTypeFamily(v string) *DescribeBundlesRequest {
	s.DesktopTypeFamily = &v
	return s
}

func (s *DescribeBundlesRequest) SetFotaChannel(v string) *DescribeBundlesRequest {
	s.FotaChannel = &v
	return s
}

func (s *DescribeBundlesRequest) SetFromDesktopGroup(v bool) *DescribeBundlesRequest {
	s.FromDesktopGroup = &v
	return s
}

func (s *DescribeBundlesRequest) SetGpuCount(v float32) *DescribeBundlesRequest {
	s.GpuCount = &v
	return s
}

func (s *DescribeBundlesRequest) SetGpuDriverType(v string) *DescribeBundlesRequest {
	s.GpuDriverType = &v
	return s
}

func (s *DescribeBundlesRequest) SetImageId(v []*string) *DescribeBundlesRequest {
	s.ImageId = v
	return s
}

func (s *DescribeBundlesRequest) SetMaxResults(v int32) *DescribeBundlesRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeBundlesRequest) SetMemorySize(v int32) *DescribeBundlesRequest {
	s.MemorySize = &v
	return s
}

func (s *DescribeBundlesRequest) SetNextToken(v string) *DescribeBundlesRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeBundlesRequest) SetOsType(v string) *DescribeBundlesRequest {
	s.OsType = &v
	return s
}

func (s *DescribeBundlesRequest) SetProtocolType(v string) *DescribeBundlesRequest {
	s.ProtocolType = &v
	return s
}

func (s *DescribeBundlesRequest) SetRegionId(v string) *DescribeBundlesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeBundlesRequest) SetScope(v string) *DescribeBundlesRequest {
	s.Scope = &v
	return s
}

func (s *DescribeBundlesRequest) SetSelectedBundle(v bool) *DescribeBundlesRequest {
	s.SelectedBundle = &v
	return s
}

func (s *DescribeBundlesRequest) SetSessionType(v string) *DescribeBundlesRequest {
	s.SessionType = &v
	return s
}

func (s *DescribeBundlesRequest) SetSupportMultiSession(v bool) *DescribeBundlesRequest {
	s.SupportMultiSession = &v
	return s
}

func (s *DescribeBundlesRequest) SetVolumeEncryptionEnabled(v bool) *DescribeBundlesRequest {
	s.VolumeEncryptionEnabled = &v
	return s
}

func (s *DescribeBundlesRequest) Validate() error {
	return dara.Validate(s)
}
