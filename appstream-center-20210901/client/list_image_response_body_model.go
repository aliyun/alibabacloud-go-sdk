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
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 22
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The returned data.
	Data []*ListImageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The message that is returned for the request.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The page number returned.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 8737D130-BFD0-5D51-96F6-C08EB1139A25
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// Valid values:
	//
	// 	- true: The request is successful.
	//
	// 	- false: The request failed.
	//
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
	// The tenant ID.
	//
	// example:
	//
	// 123456789
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// The application configuration.
	AppList []*ListImageResponseBodyDataAppList `json:"AppList,omitempty" xml:"AppList,omitempty" type:"Repeated"`
	// The base image ID.
	//
	// example:
	//
	// imgc-xxx
	BaseImageId *string `json:"BaseImageId,omitempty" xml:"BaseImageId,omitempty"`
	// The version of the base image.
	//
	// example:
	//
	// iv-xxx
	BaseImageVersion *string `json:"BaseImageVersion,omitempty" xml:"BaseImageVersion,omitempty"`
	// The service type.
	//
	// example:
	//
	// 1
	BizType *int32 `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// Specifies whether to use the compatibility mode.
	//
	// example:
	//
	// true
	CompatibleMode *bool `json:"CompatibleMode,omitempty" xml:"CompatibleMode,omitempty"`
	// The size of the data disk. Unit: GiB.
	//
	// example:
	//
	// 100
	DataDiskSize *int32 `json:"DataDiskSize,omitempty" xml:"DataDiskSize,omitempty"`
	// The description of the image.
	//
	// example:
	//
	// 描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the release.
	//
	// example:
	//
	// Ubuntu 22.04 64位
	Distro *string `json:"Distro,omitempty" xml:"Distro,omitempty"`
	// The information about each driver.
	DriverList []*string `json:"DriverList,omitempty" xml:"DriverList,omitempty" type:"Repeated"`
	// example:
	//
	// env-164c321f405ca84143e4b730dbe4
	EnvironmentId *string `json:"EnvironmentId,omitempty" xml:"EnvironmentId,omitempty"`
	// The image feature tags.
	FeatureList []*string `json:"FeatureList,omitempty" xml:"FeatureList,omitempty" type:"Repeated"`
	// >  This parameter is not available for public use.
	//
	// example:
	//
	// 镜像支持的fota渠道，暂未启用
	FotaChannel *string `json:"FotaChannel,omitempty" xml:"FotaChannel,omitempty"`
	// The FOTA version.
	//
	// example:
	//
	// 2.3.0-xxx
	FotaVersion *string `json:"FotaVersion,omitempty" xml:"FotaVersion,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2025-04-25 15:13:57
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The update time.
	//
	// example:
	//
	// 2025-04-25 15:13:57
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The image creation type.
	//
	// Valid values:
	//
	// 	- BY_SNAPSHOT_ID: The image is created from the snapshot or contains only a system disk.
	//
	// 	- BUILDER_MANUAL: The image is manually generated.
	//
	// 	- INSTANCE_AUTO: The image is automatically generated based on an instance.
	//
	// 	- BY_INSTANCE_ID: The image is created from an instance or contains both a system disk and data disks.
	//
	// example:
	//
	// BUILDER_MANUAL
	ImageCreateMode *string `json:"ImageCreateMode,omitempty" xml:"ImageCreateMode,omitempty"`
	ImageIconUrl    *string `json:"ImageIconUrl,omitempty" xml:"ImageIconUrl,omitempty"`
	// System image IDs follow a descriptive, human-readable format, while custom image IDs are automatically generated by the system.
	//
	// example:
	//
	// imgc-xxxx
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The image name.
	//
	// example:
	//
	// DemoImage
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// Layered supported regions information.
	ImageRegionDistributeList []*ListImageResponseBodyDataImageRegionDistributeList `json:"ImageRegionDistributeList,omitempty" xml:"ImageRegionDistributeList,omitempty" type:"Repeated"`
	// The region ID.
	ImageRegionList []*string `json:"ImageRegionList,omitempty" xml:"ImageRegionList,omitempty" type:"Repeated"`
	// The type of the image.
	//
	// example:
	//
	// User
	ImageType *string `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	// The image language. When the packaging type is VHD or Container, the image inherits its properties from the ECS-type image within the same image bundle.
	//
	// example:
	//
	// zh-CN
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The latest sub-version of the image. (An image consists of multiple sub-versions.)
	//
	// example:
	//
	// iv-xxx
	LatestVersionId *string `json:"LatestVersionId,omitempty" xml:"LatestVersionId,omitempty"`
	// Indicates whether the current version is the active version.
	//
	// example:
	//
	// true
	OnlineVersion *bool `json:"OnlineVersion,omitempty" xml:"OnlineVersion,omitempty"`
	// This image tag currently points to the specified sub-version of the parent image. (An image consists of multiple sub-versions.)
	//
	// example:
	//
	// iv-xxxx
	OnlineVersionId *string `json:"OnlineVersionId,omitempty" xml:"OnlineVersionId,omitempty"`
	// The type of the image.
	//
	// example:
	//
	// Windows
	OsType *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	// The image encapsulation type.
	//
	// example:
	//
	// ECS
	PackageType *string `json:"PackageType,omitempty" xml:"PackageType,omitempty"`
	// The ID of the parent image from which this image is derived. Note: This field is only applicable to custom images, as system images do not have a parent.
	//
	// example:
	//
	// imgc-xxx
	ParentImageId *string `json:"ParentImageId,omitempty" xml:"ParentImageId,omitempty"`
	// The version of the parent image.
	//
	// example:
	//
	// iv-xxx
	ParentImageVersion *string `json:"ParentImageVersion,omitempty" xml:"ParentImageVersion,omitempty"`
	// The operating system platform of the image.
	//
	// >  When the packaging type is VHD or Container, the image inherits its properties from the ECS-type image within the same image bundle.
	//
	// example:
	//
	// Windows
	Platform *int32 `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// The name of the operating system platform.
	//
	// example:
	//
	// Windows Server 2022
	PlatformName *string `json:"PlatformName,omitempty" xml:"PlatformName,omitempty"`
	// The product type.
	//
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The protocol type.
	//
	// example:
	//
	// ASP
	ProtocolType *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	// example:
	//
	// 95
	Rating *int32 `json:"Rating,omitempty" xml:"Rating,omitempty"`
	// The types of resources that are supported by the images.
	//
	// example:
	//
	// ["eds.cpu.category"]
	ResourceInstanceCategory *string `json:"ResourceInstanceCategory,omitempty" xml:"ResourceInstanceCategory,omitempty"`
	// example:
	//
	// AIGC
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// The type of the session.
	//
	// Valid values:
	//
	// 	- SINGLE_SESSION
	//
	// 	- MULTIPLE_SESSION
	//
	// example:
	//
	// SINGLE_SESSION
	SessionType *string `json:"SessionType,omitempty" xml:"SessionType,omitempty"`
	// The state of the image.
	//
	// example:
	//
	// INIT
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The supported languages.
	SupportedLanguageList []*string `json:"SupportedLanguageList,omitempty" xml:"SupportedLanguageList,omitempty" type:"Repeated"`
	// The size of the system disk. Unit: GiB.
	//
	// >  The system disk must be at least as large as the image.
	//
	// example:
	//
	// 40
	SystemDiskSize *int32                              `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
	TagList        []*ListImageResponseBodyDataTagList `json:"TagList,omitempty" xml:"TagList,omitempty" type:"Repeated"`
	// The version of the image.
	//
	// example:
	//
	// iv-xxx
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	// The version number.
	//
	// example:
	//
	// v0.1.0
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
	// Specifies whether to enable disk encryption.
	//
	// example:
	//
	// true
	VolumeEncryptionEnabled *bool `json:"VolumeEncryptionEnabled,omitempty" xml:"VolumeEncryptionEnabled,omitempty"`
	// The ID of the Key Management Service (KMS) key that is used to encrypt the disk.
	//
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

func (s *ListImageResponseBodyData) GetEnvironmentId() *string {
	return s.EnvironmentId
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

func (s *ListImageResponseBodyData) GetImageIconUrl() *string {
	return s.ImageIconUrl
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

func (s *ListImageResponseBodyData) GetRating() *int32 {
	return s.Rating
}

func (s *ListImageResponseBodyData) GetResourceInstanceCategory() *string {
	return s.ResourceInstanceCategory
}

func (s *ListImageResponseBodyData) GetScene() *string {
	return s.Scene
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

func (s *ListImageResponseBodyData) GetTagList() []*ListImageResponseBodyDataTagList {
	return s.TagList
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

func (s *ListImageResponseBodyData) SetEnvironmentId(v string) *ListImageResponseBodyData {
	s.EnvironmentId = &v
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

func (s *ListImageResponseBodyData) SetImageIconUrl(v string) *ListImageResponseBodyData {
	s.ImageIconUrl = &v
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

func (s *ListImageResponseBodyData) SetRating(v int32) *ListImageResponseBodyData {
	s.Rating = &v
	return s
}

func (s *ListImageResponseBodyData) SetResourceInstanceCategory(v string) *ListImageResponseBodyData {
	s.ResourceInstanceCategory = &v
	return s
}

func (s *ListImageResponseBodyData) SetScene(v string) *ListImageResponseBodyData {
	s.Scene = &v
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

func (s *ListImageResponseBodyData) SetTagList(v []*ListImageResponseBodyDataTagList) *ListImageResponseBodyData {
	s.TagList = v
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
	if s.TagList != nil {
		for _, item := range s.TagList {
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
	// The ID of the application.
	//
	// example:
	//
	// ca-xxx
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The name of the application.
	//
	// example:
	//
	// 应用名称
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
	// System image IDs follow a descriptive, human-readable format, while custom image IDs are automatically generated by the system.
	//
	// example:
	//
	// imgc-xxx
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The progress.
	//
	// example:
	//
	// 70%
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The supported region.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The status.
	//
	// example:
	//
	// INIT
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The version of the image.
	//
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

type ListImageResponseBodyDataTagList struct {
	// example:
	//
	// 1630348213973321
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// inner
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListImageResponseBodyDataTagList) String() string {
	return dara.Prettify(s)
}

func (s ListImageResponseBodyDataTagList) GoString() string {
	return s.String()
}

func (s *ListImageResponseBodyDataTagList) GetKey() *string {
	return s.Key
}

func (s *ListImageResponseBodyDataTagList) GetValue() *string {
	return s.Value
}

func (s *ListImageResponseBodyDataTagList) SetKey(v string) *ListImageResponseBodyDataTagList {
	s.Key = &v
	return s
}

func (s *ListImageResponseBodyDataTagList) SetValue(v string) *ListImageResponseBodyDataTagList {
	s.Value = &v
	return s
}

func (s *ListImageResponseBodyDataTagList) Validate() error {
	return dara.Validate(s)
}
