// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListImageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListImageResponseBody
	GetCode() *string
	SetCount(v int32) *ListImageResponseBody
	GetCount() *int32
	SetData(v []*ListImageResponseBodyData) *ListImageResponseBody
	GetData() []*ListImageResponseBodyData
	SetMessage(v string) *ListImageResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *ListImageResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListImageResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListImageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListImageResponseBody
	GetSuccess() *bool
}

type ListImageResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 22
	Count *int32                       `json:"Count,omitempty" xml:"Count,omitempty"`
	Data  []*ListImageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 8737D130-BFD0-5D51-96F6-C08EB1139A25
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListImageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListImageResponseBody) GoString() string {
	return s.String()
}

func (s *ListImageResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListImageResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *ListImageResponseBody) GetData() []*ListImageResponseBodyData {
	return s.Data
}

func (s *ListImageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListImageResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListImageResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListImageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListImageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListImageResponseBody) SetCode(v string) *ListImageResponseBody {
	s.Code = &v
	return s
}

func (s *ListImageResponseBody) SetCount(v int32) *ListImageResponseBody {
	s.Count = &v
	return s
}

func (s *ListImageResponseBody) SetData(v []*ListImageResponseBodyData) *ListImageResponseBody {
	s.Data = v
	return s
}

func (s *ListImageResponseBody) SetMessage(v string) *ListImageResponseBody {
	s.Message = &v
	return s
}

func (s *ListImageResponseBody) SetPageNumber(v int32) *ListImageResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListImageResponseBody) SetPageSize(v int32) *ListImageResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListImageResponseBody) SetRequestId(v string) *ListImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListImageResponseBody) SetSuccess(v bool) *ListImageResponseBody {
	s.Success = &v
	return s
}

func (s *ListImageResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListImageResponseBodyData struct {
	// example:
	//
	// 123456789
	AliUid  *int64                              `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	AppList []*ListImageResponseBodyDataAppList `json:"AppList,omitempty" xml:"AppList,omitempty" type:"Repeated"`
	// example:
	//
	// imgc-xxx
	BaseImageId *string `json:"BaseImageId,omitempty" xml:"BaseImageId,omitempty"`
	// example:
	//
	// iv-xxx
	BaseImageVersion *string `json:"BaseImageVersion,omitempty" xml:"BaseImageVersion,omitempty"`
	// example:
	//
	// 1
	BizType *int32 `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// example:
	//
	// true
	CompatibleMode *bool `json:"CompatibleMode,omitempty" xml:"CompatibleMode,omitempty"`
	// example:
	//
	// 100
	DataDiskSize *int32    `json:"DataDiskSize,omitempty" xml:"DataDiskSize,omitempty"`
	Description  *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	Distro       *string   `json:"Distro,omitempty" xml:"Distro,omitempty"`
	DriverList   []*string `json:"DriverList,omitempty" xml:"DriverList,omitempty" type:"Repeated"`
	FeatureList  []*string `json:"FeatureList,omitempty" xml:"FeatureList,omitempty" type:"Repeated"`
	FotaChannel  *string   `json:"FotaChannel,omitempty" xml:"FotaChannel,omitempty"`
	// example:
	//
	// 2.3.0-xxx
	FotaVersion *string `json:"FotaVersion,omitempty" xml:"FotaVersion,omitempty"`
	// example:
	//
	// 2025-04-25 15:13:57
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 2025-04-25 15:13:57
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// BUILDER_MANUAL
	ImageCreateMode *string `json:"ImageCreateMode,omitempty" xml:"ImageCreateMode,omitempty"`
	// example:
	//
	// imgc-xxxx
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// example:
	//
	// DemoImage
	ImageName                 *string                                               `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	ImageRegionDistributeList []*ListImageResponseBodyDataImageRegionDistributeList `json:"ImageRegionDistributeList,omitempty" xml:"ImageRegionDistributeList,omitempty" type:"Repeated"`
	ImageRegionList           []*string                                             `json:"ImageRegionList,omitempty" xml:"ImageRegionList,omitempty" type:"Repeated"`
	// example:
	//
	// User
	ImageType *string `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	// example:
	//
	// zh-CN
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// example:
	//
	// iv-xxx
	LatestVersionId *string `json:"LatestVersionId,omitempty" xml:"LatestVersionId,omitempty"`
	OnlineVersion   *bool   `json:"OnlineVersion,omitempty" xml:"OnlineVersion,omitempty"`
	// example:
	//
	// iv-xxxx
	OnlineVersionId *string `json:"OnlineVersionId,omitempty" xml:"OnlineVersionId,omitempty"`
	// example:
	//
	// Windows
	OsType *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	// example:
	//
	// ECS
	PackageType *string `json:"PackageType,omitempty" xml:"PackageType,omitempty"`
	// example:
	//
	// imgc-xxx
	ParentImageId *string `json:"ParentImageId,omitempty" xml:"ParentImageId,omitempty"`
	// example:
	//
	// iv-xxx
	ParentImageVersion *string `json:"ParentImageVersion,omitempty" xml:"ParentImageVersion,omitempty"`
	// example:
	//
	// Windows
	Platform *int32 `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// example:
	//
	// Windows Server 2022
	PlatformName *string `json:"PlatformName,omitempty" xml:"PlatformName,omitempty"`
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// example:
	//
	// ASP
	ProtocolType *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	// example:
	//
	// ["eds.cpu.category"]
	ResourceInstanceCategory *string `json:"ResourceInstanceCategory,omitempty" xml:"ResourceInstanceCategory,omitempty"`
	// example:
	//
	// SINGLE_SESSION
	SessionType *string `json:"SessionType,omitempty" xml:"SessionType,omitempty"`
	// example:
	//
	// INIT
	Status                *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	SupportedLanguageList []*string `json:"SupportedLanguageList,omitempty" xml:"SupportedLanguageList,omitempty" type:"Repeated"`
	// example:
	//
	// 40
	SystemDiskSize *int32 `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
	// example:
	//
	// iv-xxx
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	// example:
	//
	// v0.1.0
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
	// example:
	//
	// true
	VolumeEncryptionEnabled *bool `json:"VolumeEncryptionEnabled,omitempty" xml:"VolumeEncryptionEnabled,omitempty"`
	// example:
	//
	// a7b3c0c8-xxxx
	VolumeEncryptionKey *string `json:"VolumeEncryptionKey,omitempty" xml:"VolumeEncryptionKey,omitempty"`
}

func (s ListImageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListImageResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListImageResponseBodyData) GetAliUid() *int64 {
	return s.AliUid
}

func (s *ListImageResponseBodyData) GetAppList() []*ListImageResponseBodyDataAppList {
	return s.AppList
}

func (s *ListImageResponseBodyData) GetBaseImageId() *string {
	return s.BaseImageId
}

func (s *ListImageResponseBodyData) GetBaseImageVersion() *string {
	return s.BaseImageVersion
}

func (s *ListImageResponseBodyData) GetBizType() *int32 {
	return s.BizType
}

func (s *ListImageResponseBodyData) GetCompatibleMode() *bool {
	return s.CompatibleMode
}

func (s *ListImageResponseBodyData) GetDataDiskSize() *int32 {
	return s.DataDiskSize
}

func (s *ListImageResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *ListImageResponseBodyData) GetDistro() *string {
	return s.Distro
}

func (s *ListImageResponseBodyData) GetDriverList() []*string {
	return s.DriverList
}

func (s *ListImageResponseBodyData) GetFeatureList() []*string {
	return s.FeatureList
}

func (s *ListImageResponseBodyData) GetFotaChannel() *string {
	return s.FotaChannel
}

func (s *ListImageResponseBodyData) GetFotaVersion() *string {
	return s.FotaVersion
}

func (s *ListImageResponseBodyData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListImageResponseBodyData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListImageResponseBodyData) GetImageCreateMode() *string {
	return s.ImageCreateMode
}

func (s *ListImageResponseBodyData) GetImageId() *string {
	return s.ImageId
}

func (s *ListImageResponseBodyData) GetImageName() *string {
	return s.ImageName
}

func (s *ListImageResponseBodyData) GetImageRegionDistributeList() []*ListImageResponseBodyDataImageRegionDistributeList {
	return s.ImageRegionDistributeList
}

func (s *ListImageResponseBodyData) GetImageRegionList() []*string {
	return s.ImageRegionList
}

func (s *ListImageResponseBodyData) GetImageType() *string {
	return s.ImageType
}

func (s *ListImageResponseBodyData) GetLanguage() *string {
	return s.Language
}

func (s *ListImageResponseBodyData) GetLatestVersionId() *string {
	return s.LatestVersionId
}

func (s *ListImageResponseBodyData) GetOnlineVersion() *bool {
	return s.OnlineVersion
}

func (s *ListImageResponseBodyData) GetOnlineVersionId() *string {
	return s.OnlineVersionId
}

func (s *ListImageResponseBodyData) GetOsType() *string {
	return s.OsType
}

func (s *ListImageResponseBodyData) GetPackageType() *string {
	return s.PackageType
}

func (s *ListImageResponseBodyData) GetParentImageId() *string {
	return s.ParentImageId
}

func (s *ListImageResponseBodyData) GetParentImageVersion() *string {
	return s.ParentImageVersion
}

func (s *ListImageResponseBodyData) GetPlatform() *int32 {
	return s.Platform
}

func (s *ListImageResponseBodyData) GetPlatformName() *string {
	return s.PlatformName
}

func (s *ListImageResponseBodyData) GetProductType() *string {
	return s.ProductType
}

func (s *ListImageResponseBodyData) GetProtocolType() *string {
	return s.ProtocolType
}

func (s *ListImageResponseBodyData) GetResourceInstanceCategory() *string {
	return s.ResourceInstanceCategory
}

func (s *ListImageResponseBodyData) GetSessionType() *string {
	return s.SessionType
}

func (s *ListImageResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *ListImageResponseBodyData) GetSupportedLanguageList() []*string {
	return s.SupportedLanguageList
}

func (s *ListImageResponseBodyData) GetSystemDiskSize() *int32 {
	return s.SystemDiskSize
}

func (s *ListImageResponseBodyData) GetVersionId() *string {
	return s.VersionId
}

func (s *ListImageResponseBodyData) GetVersionName() *string {
	return s.VersionName
}

func (s *ListImageResponseBodyData) GetVolumeEncryptionEnabled() *bool {
	return s.VolumeEncryptionEnabled
}

func (s *ListImageResponseBodyData) GetVolumeEncryptionKey() *string {
	return s.VolumeEncryptionKey
}

func (s *ListImageResponseBodyData) SetAliUid(v int64) *ListImageResponseBodyData {
	s.AliUid = &v
	return s
}

func (s *ListImageResponseBodyData) SetAppList(v []*ListImageResponseBodyDataAppList) *ListImageResponseBodyData {
	s.AppList = v
	return s
}

func (s *ListImageResponseBodyData) SetBaseImageId(v string) *ListImageResponseBodyData {
	s.BaseImageId = &v
	return s
}

func (s *ListImageResponseBodyData) SetBaseImageVersion(v string) *ListImageResponseBodyData {
	s.BaseImageVersion = &v
	return s
}

func (s *ListImageResponseBodyData) SetBizType(v int32) *ListImageResponseBodyData {
	s.BizType = &v
	return s
}

func (s *ListImageResponseBodyData) SetCompatibleMode(v bool) *ListImageResponseBodyData {
	s.CompatibleMode = &v
	return s
}

func (s *ListImageResponseBodyData) SetDataDiskSize(v int32) *ListImageResponseBodyData {
	s.DataDiskSize = &v
	return s
}

func (s *ListImageResponseBodyData) SetDescription(v string) *ListImageResponseBodyData {
	s.Description = &v
	return s
}

func (s *ListImageResponseBodyData) SetDistro(v string) *ListImageResponseBodyData {
	s.Distro = &v
	return s
}

func (s *ListImageResponseBodyData) SetDriverList(v []*string) *ListImageResponseBodyData {
	s.DriverList = v
	return s
}

func (s *ListImageResponseBodyData) SetFeatureList(v []*string) *ListImageResponseBodyData {
	s.FeatureList = v
	return s
}

func (s *ListImageResponseBodyData) SetFotaChannel(v string) *ListImageResponseBodyData {
	s.FotaChannel = &v
	return s
}

func (s *ListImageResponseBodyData) SetFotaVersion(v string) *ListImageResponseBodyData {
	s.FotaVersion = &v
	return s
}

func (s *ListImageResponseBodyData) SetGmtCreate(v string) *ListImageResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *ListImageResponseBodyData) SetGmtModified(v string) *ListImageResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *ListImageResponseBodyData) SetImageCreateMode(v string) *ListImageResponseBodyData {
	s.ImageCreateMode = &v
	return s
}

func (s *ListImageResponseBodyData) SetImageId(v string) *ListImageResponseBodyData {
	s.ImageId = &v
	return s
}

func (s *ListImageResponseBodyData) SetImageName(v string) *ListImageResponseBodyData {
	s.ImageName = &v
	return s
}

func (s *ListImageResponseBodyData) SetImageRegionDistributeList(v []*ListImageResponseBodyDataImageRegionDistributeList) *ListImageResponseBodyData {
	s.ImageRegionDistributeList = v
	return s
}

func (s *ListImageResponseBodyData) SetImageRegionList(v []*string) *ListImageResponseBodyData {
	s.ImageRegionList = v
	return s
}

func (s *ListImageResponseBodyData) SetImageType(v string) *ListImageResponseBodyData {
	s.ImageType = &v
	return s
}

func (s *ListImageResponseBodyData) SetLanguage(v string) *ListImageResponseBodyData {
	s.Language = &v
	return s
}

func (s *ListImageResponseBodyData) SetLatestVersionId(v string) *ListImageResponseBodyData {
	s.LatestVersionId = &v
	return s
}

func (s *ListImageResponseBodyData) SetOnlineVersion(v bool) *ListImageResponseBodyData {
	s.OnlineVersion = &v
	return s
}

func (s *ListImageResponseBodyData) SetOnlineVersionId(v string) *ListImageResponseBodyData {
	s.OnlineVersionId = &v
	return s
}

func (s *ListImageResponseBodyData) SetOsType(v string) *ListImageResponseBodyData {
	s.OsType = &v
	return s
}

func (s *ListImageResponseBodyData) SetPackageType(v string) *ListImageResponseBodyData {
	s.PackageType = &v
	return s
}

func (s *ListImageResponseBodyData) SetParentImageId(v string) *ListImageResponseBodyData {
	s.ParentImageId = &v
	return s
}

func (s *ListImageResponseBodyData) SetParentImageVersion(v string) *ListImageResponseBodyData {
	s.ParentImageVersion = &v
	return s
}

func (s *ListImageResponseBodyData) SetPlatform(v int32) *ListImageResponseBodyData {
	s.Platform = &v
	return s
}

func (s *ListImageResponseBodyData) SetPlatformName(v string) *ListImageResponseBodyData {
	s.PlatformName = &v
	return s
}

func (s *ListImageResponseBodyData) SetProductType(v string) *ListImageResponseBodyData {
	s.ProductType = &v
	return s
}

func (s *ListImageResponseBodyData) SetProtocolType(v string) *ListImageResponseBodyData {
	s.ProtocolType = &v
	return s
}

func (s *ListImageResponseBodyData) SetResourceInstanceCategory(v string) *ListImageResponseBodyData {
	s.ResourceInstanceCategory = &v
	return s
}

func (s *ListImageResponseBodyData) SetSessionType(v string) *ListImageResponseBodyData {
	s.SessionType = &v
	return s
}

func (s *ListImageResponseBodyData) SetStatus(v string) *ListImageResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListImageResponseBodyData) SetSupportedLanguageList(v []*string) *ListImageResponseBodyData {
	s.SupportedLanguageList = v
	return s
}

func (s *ListImageResponseBodyData) SetSystemDiskSize(v int32) *ListImageResponseBodyData {
	s.SystemDiskSize = &v
	return s
}

func (s *ListImageResponseBodyData) SetVersionId(v string) *ListImageResponseBodyData {
	s.VersionId = &v
	return s
}

func (s *ListImageResponseBodyData) SetVersionName(v string) *ListImageResponseBodyData {
	s.VersionName = &v
	return s
}

func (s *ListImageResponseBodyData) SetVolumeEncryptionEnabled(v bool) *ListImageResponseBodyData {
	s.VolumeEncryptionEnabled = &v
	return s
}

func (s *ListImageResponseBodyData) SetVolumeEncryptionKey(v string) *ListImageResponseBodyData {
	s.VolumeEncryptionKey = &v
	return s
}

func (s *ListImageResponseBodyData) Validate() error {
	if s.AppList != nil {
		for _, item := range s.AppList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ImageRegionDistributeList != nil {
		for _, item := range s.ImageRegionDistributeList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListImageResponseBodyDataAppList struct {
	// example:
	//
	// ca-xxx
	AppId   *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
}

func (s ListImageResponseBodyDataAppList) String() string {
	return dara.Prettify(s)
}

func (s ListImageResponseBodyDataAppList) GoString() string {
	return s.String()
}

func (s *ListImageResponseBodyDataAppList) GetAppId() *string {
	return s.AppId
}

func (s *ListImageResponseBodyDataAppList) GetAppName() *string {
	return s.AppName
}

func (s *ListImageResponseBodyDataAppList) SetAppId(v string) *ListImageResponseBodyDataAppList {
	s.AppId = &v
	return s
}

func (s *ListImageResponseBodyDataAppList) SetAppName(v string) *ListImageResponseBodyDataAppList {
	s.AppName = &v
	return s
}

func (s *ListImageResponseBodyDataAppList) Validate() error {
	return dara.Validate(s)
}

type ListImageResponseBodyDataImageRegionDistributeList struct {
	// example:
	//
	// imgc-xxx
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// example:
	//
	// 70%
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// INIT
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// iv-xxx
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s ListImageResponseBodyDataImageRegionDistributeList) String() string {
	return dara.Prettify(s)
}

func (s ListImageResponseBodyDataImageRegionDistributeList) GoString() string {
	return s.String()
}

func (s *ListImageResponseBodyDataImageRegionDistributeList) GetImageId() *string {
	return s.ImageId
}

func (s *ListImageResponseBodyDataImageRegionDistributeList) GetProgress() *string {
	return s.Progress
}

func (s *ListImageResponseBodyDataImageRegionDistributeList) GetRegionId() *string {
	return s.RegionId
}

func (s *ListImageResponseBodyDataImageRegionDistributeList) GetStatus() *string {
	return s.Status
}

func (s *ListImageResponseBodyDataImageRegionDistributeList) GetVersionId() *string {
	return s.VersionId
}

func (s *ListImageResponseBodyDataImageRegionDistributeList) SetImageId(v string) *ListImageResponseBodyDataImageRegionDistributeList {
	s.ImageId = &v
	return s
}

func (s *ListImageResponseBodyDataImageRegionDistributeList) SetProgress(v string) *ListImageResponseBodyDataImageRegionDistributeList {
	s.Progress = &v
	return s
}

func (s *ListImageResponseBodyDataImageRegionDistributeList) SetRegionId(v string) *ListImageResponseBodyDataImageRegionDistributeList {
	s.RegionId = &v
	return s
}

func (s *ListImageResponseBodyDataImageRegionDistributeList) SetStatus(v string) *ListImageResponseBodyDataImageRegionDistributeList {
	s.Status = &v
	return s
}

func (s *ListImageResponseBodyDataImageRegionDistributeList) SetVersionId(v string) *ListImageResponseBodyDataImageRegionDistributeList {
	s.VersionId = &v
	return s
}

func (s *ListImageResponseBodyDataImageRegionDistributeList) Validate() error {
	return dara.Validate(s)
}
