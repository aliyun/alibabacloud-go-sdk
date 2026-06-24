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
	// The error code.
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
	// The returned data object.
	Data []*ListImageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The message returned for the API request.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The page number of the returned data.
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
	// Indicates whether the call was successful.
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
	// The application configurations.
	AppList []*ListImageResponseBodyDataAppList `json:"AppList,omitempty" xml:"AppList,omitempty" type:"Repeated"`
	// The base image ID.
	//
	// example:
	//
	// imgc-xxx
	BaseImageId *string `json:"BaseImageId,omitempty" xml:"BaseImageId,omitempty"`
	// The base image version.
	//
	// example:
	//
	// iv-xxx
	BaseImageVersion *string `json:"BaseImageVersion,omitempty" xml:"BaseImageVersion,omitempty"`
	// The business type.
	//
	// example:
	//
	// 1
	BizType *int32 `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// Indicates whether the compatibility mode is enabled.
	//
	// example:
	//
	// true
	CompatibleMode *bool `json:"CompatibleMode,omitempty" xml:"CompatibleMode,omitempty"`
	// The data cloud disk size. Unit: GiB.
	//
	// example:
	//
	// 100
	DataDiskSize *int32 `json:"DataDiskSize,omitempty" xml:"DataDiskSize,omitempty"`
	// The image description.
	//
	// example:
	//
	// 描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The distribution name.
	//
	// example:
	//
	// Ubuntu 22.04 64位
	Distro *string `json:"Distro,omitempty" xml:"Distro,omitempty"`
	// The list of driver information.
	DriverList []*string `json:"DriverList,omitempty" xml:"DriverList,omitempty" type:"Repeated"`
	// example:
	//
	// env-164c321f405ca84143e4b730dbe4
	EnvironmentId *string `json:"EnvironmentId,omitempty" xml:"EnvironmentId,omitempty"`
	// The list of image feature tags.
	FeatureList []*string `json:"FeatureList,omitempty" xml:"FeatureList,omitempty" type:"Repeated"`
	// > This parameter is not publicly available.
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
	// example:
	//
	// BUILDER_MANUAL
	ImageCreateMode *string `json:"ImageCreateMode,omitempty" xml:"ImageCreateMode,omitempty"`
	ImageIconUrl    *string `json:"ImageIconUrl,omitempty" xml:"ImageIconUrl,omitempty"`
	// The image ID. System image IDs are meaningful, while custom image IDs are automatically generated.
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
	// The effective region information for overlay layers.
	ImageRegionDistributeList []*ListImageResponseBodyDataImageRegionDistributeList `json:"ImageRegionDistributeList,omitempty" xml:"ImageRegionDistributeList,omitempty" type:"Repeated"`
	// The regions.
	ImageRegionList []*string `json:"ImageRegionList,omitempty" xml:"ImageRegionList,omitempty" type:"Repeated"`
	// The image type.
	//
	// example:
	//
	// User
	ImageType *string `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	// The image language. If the package type is VHD or Container, this property is inherited from the ECS-packaged image in the image combination.
	//
	// example:
	//
	// zh-CN
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The latest sub-version of the image. An image consists of multiple sub-versions.
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
	// The sub-version from which the current image reads the primary image information. An image consists of multiple sub-versions.
	//
	// example:
	//
	// iv-xxxx
	OnlineVersionId *string `json:"OnlineVersionId,omitempty" xml:"OnlineVersionId,omitempty"`
	// The image type.
	//
	// example:
	//
	// Windows
	OsType *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	// The image package type.
	//
	// example:
	//
	// ECS
	PackageType *string `json:"PackageType,omitempty" xml:"PackageType,omitempty"`
	// The parent image ID. This parameter indicates only the inheritance relationship. System images do not have a parent image.
	//
	// example:
	//
	// imgc-xxx
	ParentImageId *string `json:"ParentImageId,omitempty" xml:"ParentImageId,omitempty"`
	// The parent image version.
	//
	// example:
	//
	// iv-xxx
	ParentImageVersion *string `json:"ParentImageVersion,omitempty" xml:"ParentImageVersion,omitempty"`
	// The operating system platform of the image.
	//
	// > If the package type is VHD or Container, this property is inherited from the ECS-packaged image in the image combination.
	//
	// example:
	//
	// Windows
	Platform *int32 `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// The system platform name.
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
	// The resource type supported by the image.
	//
	// example:
	//
	// ["eds.cpu.category"]
	ResourceInstanceCategory *string `json:"ResourceInstanceCategory,omitempty" xml:"ResourceInstanceCategory,omitempty"`
	// example:
	//
	// AIGC
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// The session type.
	//
	// example:
	//
	// SINGLE_SESSION
	SessionType  *string                                  `json:"SessionType,omitempty" xml:"SessionType,omitempty"`
	SnapshotList []*ListImageResponseBodyDataSnapshotList `json:"SnapshotList,omitempty" xml:"SnapshotList,omitempty" type:"Repeated"`
	// The image status.
	//
	// example:
	//
	// INIT
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The list of supported languages.
	SupportedLanguageList []*string `json:"SupportedLanguageList,omitempty" xml:"SupportedLanguageList,omitempty" type:"Repeated"`
	// The system cloud disk size. Unit: GiB.
	//
	// > The system cloud disk size cannot be smaller than the image file.
	//
	// example:
	//
	// 40
	SystemDiskSize *int32                              `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
	TagList        []*ListImageResponseBodyDataTagList `json:"TagList,omitempty" xml:"TagList,omitempty" type:"Repeated"`
	// The image version.
	//
	// example:
	//
	// iv-xxx
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	// The version name.
	//
	// example:
	//
	// v0.1.0
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
	// Indicates whether cloud disk encryption is enabled.
	//
	// example:
	//
	// true
	VolumeEncryptionEnabled *bool `json:"VolumeEncryptionEnabled,omitempty" xml:"VolumeEncryptionEnabled,omitempty"`
	// The KMS key ID used when cloud disk encryption is enabled.
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

func (s *ListImageResponseBodyData) GetSnapshotList() []*ListImageResponseBodyDataSnapshotList {
	return s.SnapshotList
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

func (s *ListImageResponseBodyData) SetSnapshotList(v []*ListImageResponseBodyDataSnapshotList) *ListImageResponseBodyData {
	s.SnapshotList = v
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
	if s.SnapshotList != nil {
		for _, item := range s.SnapshotList {
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
	// The application ID.
	//
	// example:
	//
	// ca-xxx
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The application name.
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
	// The image ID. System image IDs are meaningful, while custom image IDs are automatically generated.
	//
	// example:
	//
	// imgc-xxx
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The progress percentage.
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
	// The image version.
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

type ListImageResponseBodyDataSnapshotList struct {
	BindType     *string `json:"BindType,omitempty" xml:"BindType,omitempty"`
	DiskCategory *string `json:"DiskCategory,omitempty" xml:"DiskCategory,omitempty"`
	DiskSubType  *string `json:"DiskSubType,omitempty" xml:"DiskSubType,omitempty"`
	DiskType     *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	Size         *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
	SnapshotId   *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	VersionId    *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s ListImageResponseBodyDataSnapshotList) String() string {
	return dara.Prettify(s)
}

func (s ListImageResponseBodyDataSnapshotList) GoString() string {
	return s.String()
}

func (s *ListImageResponseBodyDataSnapshotList) GetBindType() *string {
	return s.BindType
}

func (s *ListImageResponseBodyDataSnapshotList) GetDiskCategory() *string {
	return s.DiskCategory
}

func (s *ListImageResponseBodyDataSnapshotList) GetDiskSubType() *string {
	return s.DiskSubType
}

func (s *ListImageResponseBodyDataSnapshotList) GetDiskType() *string {
	return s.DiskType
}

func (s *ListImageResponseBodyDataSnapshotList) GetSize() *int32 {
	return s.Size
}

func (s *ListImageResponseBodyDataSnapshotList) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *ListImageResponseBodyDataSnapshotList) GetVersionId() *string {
	return s.VersionId
}

func (s *ListImageResponseBodyDataSnapshotList) SetBindType(v string) *ListImageResponseBodyDataSnapshotList {
	s.BindType = &v
	return s
}

func (s *ListImageResponseBodyDataSnapshotList) SetDiskCategory(v string) *ListImageResponseBodyDataSnapshotList {
	s.DiskCategory = &v
	return s
}

func (s *ListImageResponseBodyDataSnapshotList) SetDiskSubType(v string) *ListImageResponseBodyDataSnapshotList {
	s.DiskSubType = &v
	return s
}

func (s *ListImageResponseBodyDataSnapshotList) SetDiskType(v string) *ListImageResponseBodyDataSnapshotList {
	s.DiskType = &v
	return s
}

func (s *ListImageResponseBodyDataSnapshotList) SetSize(v int32) *ListImageResponseBodyDataSnapshotList {
	s.Size = &v
	return s
}

func (s *ListImageResponseBodyDataSnapshotList) SetSnapshotId(v string) *ListImageResponseBodyDataSnapshotList {
	s.SnapshotId = &v
	return s
}

func (s *ListImageResponseBodyDataSnapshotList) SetVersionId(v string) *ListImageResponseBodyDataSnapshotList {
	s.VersionId = &v
	return s
}

func (s *ListImageResponseBodyDataSnapshotList) Validate() error {
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
