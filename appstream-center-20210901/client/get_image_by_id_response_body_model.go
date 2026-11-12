// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetImageByIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetImageByIdResponseBody
	GetCode() *string
	SetData(v *GetImageByIdResponseBodyData) *GetImageByIdResponseBody
	GetData() *GetImageByIdResponseBodyData
	SetMessage(v string) *GetImageByIdResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetImageByIdResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetImageByIdResponseBody
	GetSuccess() *bool
}

type GetImageByIdResponseBody struct {
	// The response code. A value of 200 indicates success. If the call fails, the corresponding error code is returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The image information returned by the query.
	Data *GetImageByIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID, which is used to locate and troubleshoot the request.
	//
	// example:
	//
	// 8737D130-BFD0-5D51-96F6-C08EB1139A25
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the API call is successful. Valid values:
	//
	// - true: Successful.
	//
	// - false: Failed.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetImageByIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetImageByIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetImageByIdResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetImageByIdResponseBody) GetData() *GetImageByIdResponseBodyData {
	return s.Data
}

func (s *GetImageByIdResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetImageByIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetImageByIdResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetImageByIdResponseBody) SetCode(v string) *GetImageByIdResponseBody {
	s.Code = &v
	return s
}

func (s *GetImageByIdResponseBody) SetData(v *GetImageByIdResponseBodyData) *GetImageByIdResponseBody {
	s.Data = v
	return s
}

func (s *GetImageByIdResponseBody) SetMessage(v string) *GetImageByIdResponseBody {
	s.Message = &v
	return s
}

func (s *GetImageByIdResponseBody) SetRequestId(v string) *GetImageByIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetImageByIdResponseBody) SetSuccess(v bool) *GetImageByIdResponseBody {
	s.Success = &v
	return s
}

func (s *GetImageByIdResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetImageByIdResponseBodyData struct {
	// The list of adapted applications for the image version.
	AdaptorInfoBOList []*GetImageByIdResponseBodyDataAdaptorInfoBOList `json:"AdaptorInfoBOList,omitempty" xml:"AdaptorInfoBOList,omitempty" type:"Repeated"`
	// The ID of the Alibaba Cloud account to which the image belongs. The account ID for a system image is 0.
	//
	// example:
	//
	// 123456789
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// The ID of the base image.
	//
	// example:
	//
	// imgc-xxx
	BaseImageId *string `json:"BaseImageId,omitempty" xml:"BaseImageId,omitempty"`
	// The ID of the base image version.
	//
	// example:
	//
	// iv-xxx
	BaseImageVersion *string `json:"BaseImageVersion,omitempty" xml:"BaseImageVersion,omitempty"`
	// The business line to which the image belongs. Valid values:
	//
	// - 1: basic business
	//
	// - 2: personal edition
	//
	// - 3: cloud office
	//
	// - 4: cloud streaming
	//
	// - 10: standard edition
	//
	// example:
	//
	// 1
	BizType *int32 `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// The name of the OSS bucket where the image files are stored.
	//
	// example:
	//
	// appstream-example
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// The description of the image.
	//
	// example:
	//
	// Description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The list of feature tags supported by the image.
	FeatureList []*string `json:"FeatureList,omitempty" xml:"FeatureList,omitempty" type:"Repeated"`
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
	// The last update time.
	//
	// example:
	//
	// 2025-04-25 15:13:57
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The image creation method. For example, BUILDER_MANUAL indicates manual creation.
	//
	// example:
	//
	// BUILDER_MANUAL
	ImageCreateMode *string `json:"ImageCreateMode,omitempty" xml:"ImageCreateMode,omitempty"`
	// The image extended information, including disk specifications and GPU driver information. For VHD or container images, the extended information of the base image and base version is used.
	ImageExtendInfoModel *GetImageByIdResponseBodyDataImageExtendInfoModel `json:"ImageExtendInfoModel,omitempty" xml:"ImageExtendInfoModel,omitempty" type:"Struct"`
	// The image ID.
	//
	// example:
	//
	// imgc-xxxx
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The name of the image.
	//
	// example:
	//
	// DemoImage
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// The list of region distribution information for the image version.
	ImageRegionDistributeList []*GetImageByIdResponseBodyDataImageRegionDistributeList `json:"ImageRegionDistributeList,omitempty" xml:"ImageRegionDistributeList,omitempty" type:"Repeated"`
	// The list of region IDs supported by the image.
	ImageRegionList []*string `json:"ImageRegionList,omitempty" xml:"ImageRegionList,omitempty" type:"Repeated"`
	// The source of the image. For example, center indicates the central image service.
	//
	// example:
	//
	// center
	ImageSource *string `json:"ImageSource,omitempty" xml:"ImageSource,omitempty"`
	// The image type. Valid values:
	//
	// - System: system image
	//
	// - User: custom image
	//
	// - Shared: shared image
	//
	// - Community: community image
	//
	// example:
	//
	// User
	ImageType *string `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	// The image language. If the encapsulation type is VHD or Container, this property is inherited from the image with the ECS encapsulation type in the image combination.
	//
	// example:
	//
	// zh-CN
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The ID of the latest image version.
	//
	// example:
	//
	// iv-xxx
	LatestVersionId *string `json:"LatestVersionId,omitempty" xml:"LatestVersionId,omitempty"`
	// Indicates whether the returned version is the active version. Valid values:
	//
	// - true: Yes.
	//
	// - false: No.
	OnlineVersion *bool `json:"OnlineVersion,omitempty" xml:"OnlineVersion,omitempty"`
	// The ID of the current online version of the image.
	//
	// example:
	//
	// iv-xxxx
	OnlineVersionId *string `json:"OnlineVersionId,omitempty" xml:"OnlineVersionId,omitempty"`
	// The operating system type of the image. Valid values:
	//
	// - Windows: Windows system
	//
	// - Linux: Linux system
	//
	// - Android: Android system
	//
	// - Unknown: unknown system
	//
	// example:
	//
	// Windows
	OsType *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	// The image package type. Valid values:
	//
	// - Ecs: ECS image.
	//
	// - Vhd: Virtual disk image.
	//
	// - Container: Container image.
	//
	// - Ecs_Container: ECS and container combined image.
	//
	// - K8s_Node: Kubernetes node image.
	//
	// example:
	//
	// ECS
	PackageType *string `json:"PackageType,omitempty" xml:"PackageType,omitempty"`
	// The ID of the parent image, which indicates the image inheritance relationship. System images do not have parent images.
	//
	// example:
	//
	// imgc-xxx
	ParentImageId *string `json:"ParentImageId,omitempty" xml:"ParentImageId,omitempty"`
	// The ID of the parent image version.
	//
	// example:
	//
	// iv-xxx
	ParentImageVersion *string `json:"ParentImageVersion,omitempty" xml:"ParentImageVersion,omitempty"`
	// The operating system platform ID. For example, 33 indicates Windows Server 2022.
	//
	// example:
	//
	// 33
	Platform *int32 `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// The name of the operating system platform, which corresponds to Platform.
	//
	// example:
	//
	// Windows Server 2022
	PlatformName *string `json:"PlatformName,omitempty" xml:"PlatformName,omitempty"`
	// The product type to which the image belongs, such as CloudDesktop (Cloud Desktop), CloudApp (Cloud Application), or WuyingServer (WUYING Workspace Enterprise Edition workstation).
	//
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The extended properties of the image version, in key-value pair format.
	//
	// example:
	//
	// {}
	PropertyMap map[string]*string `json:"PropertyMap,omitempty" xml:"PropertyMap,omitempty"`
	// The connection protocol used by the image. Valid values:
	//
	// - ASP: Alibaba Cloud Streaming Protocol
	//
	// - HDX: HDX protocol
	//
	// example:
	//
	// ASP
	ProtocolType *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	// The image status. Valid values:
	//
	// - INIT: Initialized.
	//
	// - CREATING: Creating.
	//
	// - AVAILABLE: Available.
	//
	// - CREATE_FAILED: Creation failed.
	//
	// - DELETE: Deleted.
	//
	// example:
	//
	// INIT
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The image version ID.
	//
	// example:
	//
	// iv-xxx
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s GetImageByIdResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetImageByIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetImageByIdResponseBodyData) GetAdaptorInfoBOList() []*GetImageByIdResponseBodyDataAdaptorInfoBOList {
	return s.AdaptorInfoBOList
}

func (s *GetImageByIdResponseBodyData) GetAliUid() *int64 {
	return s.AliUid
}

func (s *GetImageByIdResponseBodyData) GetBaseImageId() *string {
	return s.BaseImageId
}

func (s *GetImageByIdResponseBodyData) GetBaseImageVersion() *string {
	return s.BaseImageVersion
}

func (s *GetImageByIdResponseBodyData) GetBizType() *int32 {
	return s.BizType
}

func (s *GetImageByIdResponseBodyData) GetBucketName() *string {
	return s.BucketName
}

func (s *GetImageByIdResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetImageByIdResponseBodyData) GetFeatureList() []*string {
	return s.FeatureList
}

func (s *GetImageByIdResponseBodyData) GetFotaVersion() *string {
	return s.FotaVersion
}

func (s *GetImageByIdResponseBodyData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetImageByIdResponseBodyData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetImageByIdResponseBodyData) GetImageCreateMode() *string {
	return s.ImageCreateMode
}

func (s *GetImageByIdResponseBodyData) GetImageExtendInfoModel() *GetImageByIdResponseBodyDataImageExtendInfoModel {
	return s.ImageExtendInfoModel
}

func (s *GetImageByIdResponseBodyData) GetImageId() *string {
	return s.ImageId
}

func (s *GetImageByIdResponseBodyData) GetImageName() *string {
	return s.ImageName
}

func (s *GetImageByIdResponseBodyData) GetImageRegionDistributeList() []*GetImageByIdResponseBodyDataImageRegionDistributeList {
	return s.ImageRegionDistributeList
}

func (s *GetImageByIdResponseBodyData) GetImageRegionList() []*string {
	return s.ImageRegionList
}

func (s *GetImageByIdResponseBodyData) GetImageSource() *string {
	return s.ImageSource
}

func (s *GetImageByIdResponseBodyData) GetImageType() *string {
	return s.ImageType
}

func (s *GetImageByIdResponseBodyData) GetLanguage() *string {
	return s.Language
}

func (s *GetImageByIdResponseBodyData) GetLatestVersionId() *string {
	return s.LatestVersionId
}

func (s *GetImageByIdResponseBodyData) GetOnlineVersion() *bool {
	return s.OnlineVersion
}

func (s *GetImageByIdResponseBodyData) GetOnlineVersionId() *string {
	return s.OnlineVersionId
}

func (s *GetImageByIdResponseBodyData) GetOsType() *string {
	return s.OsType
}

func (s *GetImageByIdResponseBodyData) GetPackageType() *string {
	return s.PackageType
}

func (s *GetImageByIdResponseBodyData) GetParentImageId() *string {
	return s.ParentImageId
}

func (s *GetImageByIdResponseBodyData) GetParentImageVersion() *string {
	return s.ParentImageVersion
}

func (s *GetImageByIdResponseBodyData) GetPlatform() *int32 {
	return s.Platform
}

func (s *GetImageByIdResponseBodyData) GetPlatformName() *string {
	return s.PlatformName
}

func (s *GetImageByIdResponseBodyData) GetProductType() *string {
	return s.ProductType
}

func (s *GetImageByIdResponseBodyData) GetPropertyMap() map[string]*string {
	return s.PropertyMap
}

func (s *GetImageByIdResponseBodyData) GetProtocolType() *string {
	return s.ProtocolType
}

func (s *GetImageByIdResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetImageByIdResponseBodyData) GetVersionId() *string {
	return s.VersionId
}

func (s *GetImageByIdResponseBodyData) SetAdaptorInfoBOList(v []*GetImageByIdResponseBodyDataAdaptorInfoBOList) *GetImageByIdResponseBodyData {
	s.AdaptorInfoBOList = v
	return s
}

func (s *GetImageByIdResponseBodyData) SetAliUid(v int64) *GetImageByIdResponseBodyData {
	s.AliUid = &v
	return s
}

func (s *GetImageByIdResponseBodyData) SetBaseImageId(v string) *GetImageByIdResponseBodyData {
	s.BaseImageId = &v
	return s
}

func (s *GetImageByIdResponseBodyData) SetBaseImageVersion(v string) *GetImageByIdResponseBodyData {
	s.BaseImageVersion = &v
	return s
}

func (s *GetImageByIdResponseBodyData) SetBizType(v int32) *GetImageByIdResponseBodyData {
	s.BizType = &v
	return s
}

func (s *GetImageByIdResponseBodyData) SetBucketName(v string) *GetImageByIdResponseBodyData {
	s.BucketName = &v
	return s
}

func (s *GetImageByIdResponseBodyData) SetDescription(v string) *GetImageByIdResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetImageByIdResponseBodyData) SetFeatureList(v []*string) *GetImageByIdResponseBodyData {
	s.FeatureList = v
	return s
}

func (s *GetImageByIdResponseBodyData) SetFotaVersion(v string) *GetImageByIdResponseBodyData {
	s.FotaVersion = &v
	return s
}

func (s *GetImageByIdResponseBodyData) SetGmtCreate(v string) *GetImageByIdResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *GetImageByIdResponseBodyData) SetGmtModified(v string) *GetImageByIdResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *GetImageByIdResponseBodyData) SetImageCreateMode(v string) *GetImageByIdResponseBodyData {
	s.ImageCreateMode = &v
	return s
}

func (s *GetImageByIdResponseBodyData) SetImageExtendInfoModel(v *GetImageByIdResponseBodyDataImageExtendInfoModel) *GetImageByIdResponseBodyData {
	s.ImageExtendInfoModel = v
	return s
}

func (s *GetImageByIdResponseBodyData) SetImageId(v string) *GetImageByIdResponseBodyData {
	s.ImageId = &v
	return s
}

func (s *GetImageByIdResponseBodyData) SetImageName(v string) *GetImageByIdResponseBodyData {
	s.ImageName = &v
	return s
}

func (s *GetImageByIdResponseBodyData) SetImageRegionDistributeList(v []*GetImageByIdResponseBodyDataImageRegionDistributeList) *GetImageByIdResponseBodyData {
	s.ImageRegionDistributeList = v
	return s
}

func (s *GetImageByIdResponseBodyData) SetImageRegionList(v []*string) *GetImageByIdResponseBodyData {
	s.ImageRegionList = v
	return s
}

func (s *GetImageByIdResponseBodyData) SetImageSource(v string) *GetImageByIdResponseBodyData {
	s.ImageSource = &v
	return s
}

func (s *GetImageByIdResponseBodyData) SetImageType(v string) *GetImageByIdResponseBodyData {
	s.ImageType = &v
	return s
}

func (s *GetImageByIdResponseBodyData) SetLanguage(v string) *GetImageByIdResponseBodyData {
	s.Language = &v
	return s
}

func (s *GetImageByIdResponseBodyData) SetLatestVersionId(v string) *GetImageByIdResponseBodyData {
	s.LatestVersionId = &v
	return s
}

func (s *GetImageByIdResponseBodyData) SetOnlineVersion(v bool) *GetImageByIdResponseBodyData {
	s.OnlineVersion = &v
	return s
}

func (s *GetImageByIdResponseBodyData) SetOnlineVersionId(v string) *GetImageByIdResponseBodyData {
	s.OnlineVersionId = &v
	return s
}

func (s *GetImageByIdResponseBodyData) SetOsType(v string) *GetImageByIdResponseBodyData {
	s.OsType = &v
	return s
}

func (s *GetImageByIdResponseBodyData) SetPackageType(v string) *GetImageByIdResponseBodyData {
	s.PackageType = &v
	return s
}

func (s *GetImageByIdResponseBodyData) SetParentImageId(v string) *GetImageByIdResponseBodyData {
	s.ParentImageId = &v
	return s
}

func (s *GetImageByIdResponseBodyData) SetParentImageVersion(v string) *GetImageByIdResponseBodyData {
	s.ParentImageVersion = &v
	return s
}

func (s *GetImageByIdResponseBodyData) SetPlatform(v int32) *GetImageByIdResponseBodyData {
	s.Platform = &v
	return s
}

func (s *GetImageByIdResponseBodyData) SetPlatformName(v string) *GetImageByIdResponseBodyData {
	s.PlatformName = &v
	return s
}

func (s *GetImageByIdResponseBodyData) SetProductType(v string) *GetImageByIdResponseBodyData {
	s.ProductType = &v
	return s
}

func (s *GetImageByIdResponseBodyData) SetPropertyMap(v map[string]*string) *GetImageByIdResponseBodyData {
	s.PropertyMap = v
	return s
}

func (s *GetImageByIdResponseBodyData) SetProtocolType(v string) *GetImageByIdResponseBodyData {
	s.ProtocolType = &v
	return s
}

func (s *GetImageByIdResponseBodyData) SetStatus(v string) *GetImageByIdResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetImageByIdResponseBodyData) SetVersionId(v string) *GetImageByIdResponseBodyData {
	s.VersionId = &v
	return s
}

func (s *GetImageByIdResponseBodyData) Validate() error {
	if s.AdaptorInfoBOList != nil {
		for _, item := range s.AdaptorInfoBOList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ImageExtendInfoModel != nil {
		if err := s.ImageExtendInfoModel.Validate(); err != nil {
			return err
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

type GetImageByIdResponseBodyDataAdaptorInfoBOList struct {
	// The adapter ID.
	//
	// example:
	//
	// adt-13ff3****
	AdaptorId *string `json:"AdaptorId,omitempty" xml:"AdaptorId,omitempty"`
	// The adaptor message.
	//
	// example:
	//
	// success
	AdaptorMessage *string `json:"AdaptorMessage,omitempty" xml:"AdaptorMessage,omitempty"`
	// The execution mode.
	//
	// example:
	//
	// MANUAL
	AdaptorMode *string `json:"AdaptorMode,omitempty" xml:"AdaptorMode,omitempty"`
	// The execution step.
	//
	// example:
	//
	// FINISH_ADAPT
	AdaptorStep *string `json:"AdaptorStep,omitempty" xml:"AdaptorStep,omitempty"`
	// The numeric ID of the application.
	//
	// example:
	//
	// 12345
	Aid *int64 `json:"Aid,omitempty" xml:"Aid,omitempty"`
	// The ID of the Alibaba Cloud account to which the application belongs.
	//
	// example:
	//
	// 123456789
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// The application icon.
	//
	// example:
	//
	// https://****.jpg
	AppIconUrl *string `json:"AppIconUrl,omitempty" xml:"AppIconUrl,omitempty"`
	// The application ID.
	//
	// example:
	//
	// app-0aae4rgle910jh689
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The application name.
	//
	// example:
	//
	// idb-ent-console
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The application script.
	//
	// example:
	//
	// {"param":"test", "startCmd":"test.exe"}
	AppScript *string `json:"AppScript,omitempty" xml:"AppScript,omitempty"`
	// The theme color.
	//
	// example:
	//
	// #FFFFFF
	AppThemeColor *string `json:"AppThemeColor,omitempty" xml:"AppThemeColor,omitempty"`
	// The application type.
	//
	// example:
	//
	// USER_APP
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// The application version ID.
	//
	// example:
	//
	// apv-****
	AppVersionId *string `json:"AppVersionId,omitempty" xml:"AppVersionId,omitempty"`
	// The version name.
	//
	// example:
	//
	// 1.0
	AppVersionName *string `json:"AppVersionName,omitempty" xml:"AppVersionName,omitempty"`
	// The version number.
	//
	// example:
	//
	// 1.0.0
	AppVersionNumber *string `json:"AppVersionNumber,omitempty" xml:"AppVersionNumber,omitempty"`
	// The ID of the business line to which the application belongs.
	//
	// example:
	//
	// 1
	BizType *int32 `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// The name of the OSS bucket where the application files are stored.
	//
	// example:
	//
	// appstream-example
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// The level-2 category of the application.
	//
	// example:
	//
	// 123
	CategoryId *int32 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// The level-1 category of the application.
	//
	// example:
	//
	// 1
	CategoryType *int32 `json:"CategoryType,omitempty" xml:"CategoryType,omitempty"`
	// The control parameters that can be configured by tenants and administrators.
	//
	// example:
	//
	// {}
	ControlParam *string `json:"ControlParam,omitempty" xml:"ControlParam,omitempty"`
	// The ID of the disk where the application files are stored.
	//
	// example:
	//
	// d-bp1example****
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// The path of the executable file after the cloud application is installed.
	//
	// example:
	//
	// abc.text
	ExecuteFileAddress *string `json:"ExecuteFileAddress,omitempty" xml:"ExecuteFileAddress,omitempty"`
	// The extended information of the adapted application.
	//
	// example:
	//
	// {}
	Extend *string `json:"Extend,omitempty" xml:"Extend,omitempty"`
	// The file name.
	//
	// example:
	//
	// filename_install.exe
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The file path.
	//
	// example:
	//
	// /tenant/123
	FilePath *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	// The file size.
	//
	// example:
	//
	// 102400
	FileSize *int64 `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2025-04-25 15:13:57
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The last update time.
	//
	// example:
	//
	// 2025-04-25 15:13:57
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The optimization parameters configured in the administrator backend.
	//
	// example:
	//
	// {}
	OptimizeParam *string `json:"OptimizeParam,omitempty" xml:"OptimizeParam,omitempty"`
	// The original application type.
	//
	// example:
	//
	// CLIENT_BASE
	OriginAppType *string `json:"OriginAppType,omitempty" xml:"OriginAppType,omitempty"`
	// The operating system type supported by the application.
	//
	// example:
	//
	// Windows
	OsType *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	// The product type to which the application belongs.
	//
	// example:
	//
	// CloudDesktop
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The runtime environment.
	//
	// example:
	//
	// WINDOWS_MULTI_SESSION
	RuntimeEnv *string `json:"RuntimeEnv,omitempty" xml:"RuntimeEnv,omitempty"`
	// Specifies whether to skip installation. This parameter can be configured by tenants and administrators.
	//
	// example:
	//
	// {}
	SilenceFlag *string `json:"SilenceFlag,omitempty" xml:"SilenceFlag,omitempty"`
	// The installation parameters that can be configured by tenants and administrators.
	//
	// example:
	//
	// {}
	SilenceParam *string `json:"SilenceParam,omitempty" xml:"SilenceParam,omitempty"`
	// The application adaptation status. For example, ADAPT_DONE indicates that the adaptation is complete.
	//
	// example:
	//
	// ADAPT_DONE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The sub-application type.
	//
	// example:
	//
	// SAAS
	SubAppType *string `json:"SubAppType,omitempty" xml:"SubAppType,omitempty"`
}

func (s GetImageByIdResponseBodyDataAdaptorInfoBOList) String() string {
	return dara.Prettify(s)
}

func (s GetImageByIdResponseBodyDataAdaptorInfoBOList) GoString() string {
	return s.String()
}

func (s *GetImageByIdResponseBodyDataAdaptorInfoBOList) GetAdaptorId() *string {
	return s.AdaptorId
}

func (s *GetImageByIdResponseBodyDataAdaptorInfoBOList) GetAdaptorMessage() *string {
	return s.AdaptorMessage
}

func (s *GetImageByIdResponseBodyDataAdaptorInfoBOList) GetAdaptorMode() *string {
	return s.AdaptorMode
}

func (s *GetImageByIdResponseBodyDataAdaptorInfoBOList) GetAdaptorStep() *string {
	return s.AdaptorStep
}

func (s *GetImageByIdResponseBodyDataAdaptorInfoBOList) GetAid() *int64 {
	return s.Aid
}

func (s *GetImageByIdResponseBodyDataAdaptorInfoBOList) GetAliUid() *int64 {
	return s.AliUid
}

func (s *GetImageByIdResponseBodyDataAdaptorInfoBOList) GetAppIconUrl() *string {
	return s.AppIconUrl
}

func (s *GetImageByIdResponseBodyDataAdaptorInfoBOList) GetAppId() *string {
	return s.AppId
}

func (s *GetImageByIdResponseBodyDataAdaptorInfoBOList) GetAppName() *string {
	return s.AppName
}

func (s *GetImageByIdResponseBodyDataAdaptorInfoBOList) GetAppScript() *string {
	return s.AppScript
}

func (s *GetImageByIdResponseBodyDataAdaptorInfoBOList) GetAppThemeColor() *string {
	return s.AppThemeColor
}

func (s *GetImageByIdResponseBodyDataAdaptorInfoBOList) GetAppType() *string {
	return s.AppType
}

func (s *GetImageByIdResponseBodyDataAdaptorInfoBOList) GetAppVersionId() *string {
	return s.AppVersionId
}

func (s *GetImageByIdResponseBodyDataAdaptorInfoBOList) GetAppVersionName() *string {
	return s.AppVersionName
}

func (s *GetImageByIdResponseBodyDataAdaptorInfoBOList) GetAppVersionNumber() *string {
	return s.AppVersionNumber
}

func (s *GetImageByIdResponseBodyDataAdaptorInfoBOList) GetBizType() *int32 {
	return s.BizType
}

func (s *GetImageByIdResponseBodyDataAdaptorInfoBOList) GetBucketName() *string {
	return s.BucketName
}

func (s *GetImageByIdResponseBodyDataAdaptorInfoBOList) GetCategoryId() *int32 {
	return s.CategoryId
}

func (s *GetImageByIdResponseBodyDataAdaptorInfoBOList) GetCategoryType() *int32 {
	return s.CategoryType
}

func (s *GetImageByIdResponseBodyDataAdaptorInfoBOList) GetControlParam() *string {
	return s.ControlParam
}

func (s *GetImageByIdResponseBodyDataAdaptorInfoBOList) GetDiskId() *string {
	return s.DiskId
}

func (s *GetImageByIdResponseBodyDataAdaptorInfoBOList) GetExecuteFileAddress() *string {
	return s.ExecuteFileAddress
}

func (s *GetImageByIdResponseBodyDataAdaptorInfoBOList) GetExtend() *string {
	return s.Extend
}

func (s *GetImageByIdResponseBodyDataAdaptorInfoBOList) GetFileName() *string {
	return s.FileName
}

func (s *GetImageByIdResponseBodyDataAdaptorInfoBOList) GetFilePath() *string {
	return s.FilePath
}

func (s *GetImageByIdResponseBodyDataAdaptorInfoBOList) GetFileSize() *int64 {
	return s.FileSize
}

func (s *GetImageByIdResponseBodyDataAdaptorInfoBOList) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetImageByIdResponseBodyDataAdaptorInfoBOList) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetImageByIdResponseBodyDataAdaptorInfoBOList) GetOptimizeParam() *string {
	return s.OptimizeParam
}

func (s *GetImageByIdResponseBodyDataAdaptorInfoBOList) GetOriginAppType() *string {
	return s.OriginAppType
}

func (s *GetImageByIdResponseBodyDataAdaptorInfoBOList) GetOsType() *string {
	return s.OsType
}

func (s *GetImageByIdResponseBodyDataAdaptorInfoBOList) GetProductType() *string {
	return s.ProductType
}

func (s *GetImageByIdResponseBodyDataAdaptorInfoBOList) GetRuntimeEnv() *string {
	return s.RuntimeEnv
}

func (s *GetImageByIdResponseBodyDataAdaptorInfoBOList) GetSilenceFlag() *string {
	return s.SilenceFlag
}

func (s *GetImageByIdResponseBodyDataAdaptorInfoBOList) GetSilenceParam() *string {
	return s.SilenceParam
}

func (s *GetImageByIdResponseBodyDataAdaptorInfoBOList) GetStatus() *string {
	return s.Status
}

func (s *GetImageByIdResponseBodyDataAdaptorInfoBOList) GetSubAppType() *string {
	return s.SubAppType
}

func (s *GetImageByIdResponseBodyDataAdaptorInfoBOList) SetAdaptorId(v string) *GetImageByIdResponseBodyDataAdaptorInfoBOList {
	s.AdaptorId = &v
	return s
}

func (s *GetImageByIdResponseBodyDataAdaptorInfoBOList) SetAdaptorMessage(v string) *GetImageByIdResponseBodyDataAdaptorInfoBOList {
	s.AdaptorMessage = &v
	return s
}

func (s *GetImageByIdResponseBodyDataAdaptorInfoBOList) SetAdaptorMode(v string) *GetImageByIdResponseBodyDataAdaptorInfoBOList {
	s.AdaptorMode = &v
	return s
}

func (s *GetImageByIdResponseBodyDataAdaptorInfoBOList) SetAdaptorStep(v string) *GetImageByIdResponseBodyDataAdaptorInfoBOList {
	s.AdaptorStep = &v
	return s
}

func (s *GetImageByIdResponseBodyDataAdaptorInfoBOList) SetAid(v int64) *GetImageByIdResponseBodyDataAdaptorInfoBOList {
	s.Aid = &v
	return s
}

func (s *GetImageByIdResponseBodyDataAdaptorInfoBOList) SetAliUid(v int64) *GetImageByIdResponseBodyDataAdaptorInfoBOList {
	s.AliUid = &v
	return s
}

func (s *GetImageByIdResponseBodyDataAdaptorInfoBOList) SetAppIconUrl(v string) *GetImageByIdResponseBodyDataAdaptorInfoBOList {
	s.AppIconUrl = &v
	return s
}

func (s *GetImageByIdResponseBodyDataAdaptorInfoBOList) SetAppId(v string) *GetImageByIdResponseBodyDataAdaptorInfoBOList {
	s.AppId = &v
	return s
}

func (s *GetImageByIdResponseBodyDataAdaptorInfoBOList) SetAppName(v string) *GetImageByIdResponseBodyDataAdaptorInfoBOList {
	s.AppName = &v
	return s
}

func (s *GetImageByIdResponseBodyDataAdaptorInfoBOList) SetAppScript(v string) *GetImageByIdResponseBodyDataAdaptorInfoBOList {
	s.AppScript = &v
	return s
}

func (s *GetImageByIdResponseBodyDataAdaptorInfoBOList) SetAppThemeColor(v string) *GetImageByIdResponseBodyDataAdaptorInfoBOList {
	s.AppThemeColor = &v
	return s
}

func (s *GetImageByIdResponseBodyDataAdaptorInfoBOList) SetAppType(v string) *GetImageByIdResponseBodyDataAdaptorInfoBOList {
	s.AppType = &v
	return s
}

func (s *GetImageByIdResponseBodyDataAdaptorInfoBOList) SetAppVersionId(v string) *GetImageByIdResponseBodyDataAdaptorInfoBOList {
	s.AppVersionId = &v
	return s
}

func (s *GetImageByIdResponseBodyDataAdaptorInfoBOList) SetAppVersionName(v string) *GetImageByIdResponseBodyDataAdaptorInfoBOList {
	s.AppVersionName = &v
	return s
}

func (s *GetImageByIdResponseBodyDataAdaptorInfoBOList) SetAppVersionNumber(v string) *GetImageByIdResponseBodyDataAdaptorInfoBOList {
	s.AppVersionNumber = &v
	return s
}

func (s *GetImageByIdResponseBodyDataAdaptorInfoBOList) SetBizType(v int32) *GetImageByIdResponseBodyDataAdaptorInfoBOList {
	s.BizType = &v
	return s
}

func (s *GetImageByIdResponseBodyDataAdaptorInfoBOList) SetBucketName(v string) *GetImageByIdResponseBodyDataAdaptorInfoBOList {
	s.BucketName = &v
	return s
}

func (s *GetImageByIdResponseBodyDataAdaptorInfoBOList) SetCategoryId(v int32) *GetImageByIdResponseBodyDataAdaptorInfoBOList {
	s.CategoryId = &v
	return s
}

func (s *GetImageByIdResponseBodyDataAdaptorInfoBOList) SetCategoryType(v int32) *GetImageByIdResponseBodyDataAdaptorInfoBOList {
	s.CategoryType = &v
	return s
}

func (s *GetImageByIdResponseBodyDataAdaptorInfoBOList) SetControlParam(v string) *GetImageByIdResponseBodyDataAdaptorInfoBOList {
	s.ControlParam = &v
	return s
}

func (s *GetImageByIdResponseBodyDataAdaptorInfoBOList) SetDiskId(v string) *GetImageByIdResponseBodyDataAdaptorInfoBOList {
	s.DiskId = &v
	return s
}

func (s *GetImageByIdResponseBodyDataAdaptorInfoBOList) SetExecuteFileAddress(v string) *GetImageByIdResponseBodyDataAdaptorInfoBOList {
	s.ExecuteFileAddress = &v
	return s
}

func (s *GetImageByIdResponseBodyDataAdaptorInfoBOList) SetExtend(v string) *GetImageByIdResponseBodyDataAdaptorInfoBOList {
	s.Extend = &v
	return s
}

func (s *GetImageByIdResponseBodyDataAdaptorInfoBOList) SetFileName(v string) *GetImageByIdResponseBodyDataAdaptorInfoBOList {
	s.FileName = &v
	return s
}

func (s *GetImageByIdResponseBodyDataAdaptorInfoBOList) SetFilePath(v string) *GetImageByIdResponseBodyDataAdaptorInfoBOList {
	s.FilePath = &v
	return s
}

func (s *GetImageByIdResponseBodyDataAdaptorInfoBOList) SetFileSize(v int64) *GetImageByIdResponseBodyDataAdaptorInfoBOList {
	s.FileSize = &v
	return s
}

func (s *GetImageByIdResponseBodyDataAdaptorInfoBOList) SetGmtCreate(v string) *GetImageByIdResponseBodyDataAdaptorInfoBOList {
	s.GmtCreate = &v
	return s
}

func (s *GetImageByIdResponseBodyDataAdaptorInfoBOList) SetGmtModified(v string) *GetImageByIdResponseBodyDataAdaptorInfoBOList {
	s.GmtModified = &v
	return s
}

func (s *GetImageByIdResponseBodyDataAdaptorInfoBOList) SetOptimizeParam(v string) *GetImageByIdResponseBodyDataAdaptorInfoBOList {
	s.OptimizeParam = &v
	return s
}

func (s *GetImageByIdResponseBodyDataAdaptorInfoBOList) SetOriginAppType(v string) *GetImageByIdResponseBodyDataAdaptorInfoBOList {
	s.OriginAppType = &v
	return s
}

func (s *GetImageByIdResponseBodyDataAdaptorInfoBOList) SetOsType(v string) *GetImageByIdResponseBodyDataAdaptorInfoBOList {
	s.OsType = &v
	return s
}

func (s *GetImageByIdResponseBodyDataAdaptorInfoBOList) SetProductType(v string) *GetImageByIdResponseBodyDataAdaptorInfoBOList {
	s.ProductType = &v
	return s
}

func (s *GetImageByIdResponseBodyDataAdaptorInfoBOList) SetRuntimeEnv(v string) *GetImageByIdResponseBodyDataAdaptorInfoBOList {
	s.RuntimeEnv = &v
	return s
}

func (s *GetImageByIdResponseBodyDataAdaptorInfoBOList) SetSilenceFlag(v string) *GetImageByIdResponseBodyDataAdaptorInfoBOList {
	s.SilenceFlag = &v
	return s
}

func (s *GetImageByIdResponseBodyDataAdaptorInfoBOList) SetSilenceParam(v string) *GetImageByIdResponseBodyDataAdaptorInfoBOList {
	s.SilenceParam = &v
	return s
}

func (s *GetImageByIdResponseBodyDataAdaptorInfoBOList) SetStatus(v string) *GetImageByIdResponseBodyDataAdaptorInfoBOList {
	s.Status = &v
	return s
}

func (s *GetImageByIdResponseBodyDataAdaptorInfoBOList) SetSubAppType(v string) *GetImageByIdResponseBodyDataAdaptorInfoBOList {
	s.SubAppType = &v
	return s
}

func (s *GetImageByIdResponseBodyDataAdaptorInfoBOList) Validate() error {
	return dara.Validate(s)
}

type GetImageByIdResponseBodyDataImageExtendInfoModel struct {
	// The ID of the Alibaba Cloud account to which the image belongs. The account ID for a system image is 0.
	//
	// example:
	//
	// 123456789
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// The size of the data cloud disk. Unit: GiB.
	//
	// example:
	//
	// 100
	DataDiskSize *int32 `json:"DataDiskSize,omitempty" xml:"DataDiskSize,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2025-04-25 15:13:57
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The last update time.
	//
	// example:
	//
	// 2025-04-25 15:13:57
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The GPU category.
	//
	// example:
	//
	// 1
	GpuCategory *int32 `json:"GpuCategory,omitempty" xml:"GpuCategory,omitempty"`
	// The GPU driver version.
	//
	// example:
	//
	// 12
	GpuDriverVersion *string `json:"GpuDriverVersion,omitempty" xml:"GpuDriverVersion,omitempty"`
	// The image ID.
	//
	// example:
	//
	// imgc-xxxx
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The image package type. Valid values:
	//
	// - Ecs: ECS image.
	//
	// - Vhd: Virtual disk image.
	//
	// - Container: Container image.
	//
	// - Ecs_Container: ECS and container combined image.
	//
	// - K8s_Node: Kubernetes node image.
	//
	// example:
	//
	// ECS
	PackageType *string `json:"PackageType,omitempty" xml:"PackageType,omitempty"`
	// The size of the system cloud disk. Unit: GiB.
	//
	// example:
	//
	// 40
	SystemDiskSize *int32 `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
	// The image version ID.
	//
	// example:
	//
	// iv-xxx
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s GetImageByIdResponseBodyDataImageExtendInfoModel) String() string {
	return dara.Prettify(s)
}

func (s GetImageByIdResponseBodyDataImageExtendInfoModel) GoString() string {
	return s.String()
}

func (s *GetImageByIdResponseBodyDataImageExtendInfoModel) GetAliUid() *int64 {
	return s.AliUid
}

func (s *GetImageByIdResponseBodyDataImageExtendInfoModel) GetDataDiskSize() *int32 {
	return s.DataDiskSize
}

func (s *GetImageByIdResponseBodyDataImageExtendInfoModel) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetImageByIdResponseBodyDataImageExtendInfoModel) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetImageByIdResponseBodyDataImageExtendInfoModel) GetGpuCategory() *int32 {
	return s.GpuCategory
}

func (s *GetImageByIdResponseBodyDataImageExtendInfoModel) GetGpuDriverVersion() *string {
	return s.GpuDriverVersion
}

func (s *GetImageByIdResponseBodyDataImageExtendInfoModel) GetImageId() *string {
	return s.ImageId
}

func (s *GetImageByIdResponseBodyDataImageExtendInfoModel) GetPackageType() *string {
	return s.PackageType
}

func (s *GetImageByIdResponseBodyDataImageExtendInfoModel) GetSystemDiskSize() *int32 {
	return s.SystemDiskSize
}

func (s *GetImageByIdResponseBodyDataImageExtendInfoModel) GetVersionId() *string {
	return s.VersionId
}

func (s *GetImageByIdResponseBodyDataImageExtendInfoModel) SetAliUid(v int64) *GetImageByIdResponseBodyDataImageExtendInfoModel {
	s.AliUid = &v
	return s
}

func (s *GetImageByIdResponseBodyDataImageExtendInfoModel) SetDataDiskSize(v int32) *GetImageByIdResponseBodyDataImageExtendInfoModel {
	s.DataDiskSize = &v
	return s
}

func (s *GetImageByIdResponseBodyDataImageExtendInfoModel) SetGmtCreate(v string) *GetImageByIdResponseBodyDataImageExtendInfoModel {
	s.GmtCreate = &v
	return s
}

func (s *GetImageByIdResponseBodyDataImageExtendInfoModel) SetGmtModified(v string) *GetImageByIdResponseBodyDataImageExtendInfoModel {
	s.GmtModified = &v
	return s
}

func (s *GetImageByIdResponseBodyDataImageExtendInfoModel) SetGpuCategory(v int32) *GetImageByIdResponseBodyDataImageExtendInfoModel {
	s.GpuCategory = &v
	return s
}

func (s *GetImageByIdResponseBodyDataImageExtendInfoModel) SetGpuDriverVersion(v string) *GetImageByIdResponseBodyDataImageExtendInfoModel {
	s.GpuDriverVersion = &v
	return s
}

func (s *GetImageByIdResponseBodyDataImageExtendInfoModel) SetImageId(v string) *GetImageByIdResponseBodyDataImageExtendInfoModel {
	s.ImageId = &v
	return s
}

func (s *GetImageByIdResponseBodyDataImageExtendInfoModel) SetPackageType(v string) *GetImageByIdResponseBodyDataImageExtendInfoModel {
	s.PackageType = &v
	return s
}

func (s *GetImageByIdResponseBodyDataImageExtendInfoModel) SetSystemDiskSize(v int32) *GetImageByIdResponseBodyDataImageExtendInfoModel {
	s.SystemDiskSize = &v
	return s
}

func (s *GetImageByIdResponseBodyDataImageExtendInfoModel) SetVersionId(v string) *GetImageByIdResponseBodyDataImageExtendInfoModel {
	s.VersionId = &v
	return s
}

func (s *GetImageByIdResponseBodyDataImageExtendInfoModel) Validate() error {
	return dara.Validate(s)
}

type GetImageByIdResponseBodyDataImageRegionDistributeList struct {
	// The ID of the Alibaba Cloud account to which the image belongs. The account ID for a system image is 0.
	//
	// example:
	//
	// 123456789
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// The ID of the ECS image corresponding to the region.
	//
	// example:
	//
	// m-bp1example****
	EcsImageId *string `json:"EcsImageId,omitempty" xml:"EcsImageId,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2025-04-25 15:13:57
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The last update time.
	//
	// example:
	//
	// 2025-04-25 15:13:57
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The image ID.
	//
	// example:
	//
	// imgc-xxxx
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The processing information reported by the region distribution task.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The distribution progress of the image in the region, expressed as a percentage.
	//
	// example:
	//
	// 100
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The ID of the region where the image is distributed.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The region distribution status. Valid values:
	//
	// - CREATING: Distributing.
	//
	// - AVAILABLE: Available.
	//
	// - CREATE_FAILED: Distribution failed.
	//
	// - DELETE: Deleted.
	//
	// example:
	//
	// AVAILABLE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The logical ID of the image on the unit side, which can be different from the central image ID.
	//
	// example:
	//
	// m-1bdedef****
	UnitImageId *string `json:"UnitImageId,omitempty" xml:"UnitImageId,omitempty"`
	// The image version ID.
	//
	// example:
	//
	// iv-xxx
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s GetImageByIdResponseBodyDataImageRegionDistributeList) String() string {
	return dara.Prettify(s)
}

func (s GetImageByIdResponseBodyDataImageRegionDistributeList) GoString() string {
	return s.String()
}

func (s *GetImageByIdResponseBodyDataImageRegionDistributeList) GetAliUid() *int64 {
	return s.AliUid
}

func (s *GetImageByIdResponseBodyDataImageRegionDistributeList) GetEcsImageId() *string {
	return s.EcsImageId
}

func (s *GetImageByIdResponseBodyDataImageRegionDistributeList) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetImageByIdResponseBodyDataImageRegionDistributeList) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetImageByIdResponseBodyDataImageRegionDistributeList) GetImageId() *string {
	return s.ImageId
}

func (s *GetImageByIdResponseBodyDataImageRegionDistributeList) GetMessage() *string {
	return s.Message
}

func (s *GetImageByIdResponseBodyDataImageRegionDistributeList) GetProgress() *string {
	return s.Progress
}

func (s *GetImageByIdResponseBodyDataImageRegionDistributeList) GetRegionId() *string {
	return s.RegionId
}

func (s *GetImageByIdResponseBodyDataImageRegionDistributeList) GetStatus() *string {
	return s.Status
}

func (s *GetImageByIdResponseBodyDataImageRegionDistributeList) GetUnitImageId() *string {
	return s.UnitImageId
}

func (s *GetImageByIdResponseBodyDataImageRegionDistributeList) GetVersionId() *string {
	return s.VersionId
}

func (s *GetImageByIdResponseBodyDataImageRegionDistributeList) SetAliUid(v int64) *GetImageByIdResponseBodyDataImageRegionDistributeList {
	s.AliUid = &v
	return s
}

func (s *GetImageByIdResponseBodyDataImageRegionDistributeList) SetEcsImageId(v string) *GetImageByIdResponseBodyDataImageRegionDistributeList {
	s.EcsImageId = &v
	return s
}

func (s *GetImageByIdResponseBodyDataImageRegionDistributeList) SetGmtCreate(v string) *GetImageByIdResponseBodyDataImageRegionDistributeList {
	s.GmtCreate = &v
	return s
}

func (s *GetImageByIdResponseBodyDataImageRegionDistributeList) SetGmtModified(v string) *GetImageByIdResponseBodyDataImageRegionDistributeList {
	s.GmtModified = &v
	return s
}

func (s *GetImageByIdResponseBodyDataImageRegionDistributeList) SetImageId(v string) *GetImageByIdResponseBodyDataImageRegionDistributeList {
	s.ImageId = &v
	return s
}

func (s *GetImageByIdResponseBodyDataImageRegionDistributeList) SetMessage(v string) *GetImageByIdResponseBodyDataImageRegionDistributeList {
	s.Message = &v
	return s
}

func (s *GetImageByIdResponseBodyDataImageRegionDistributeList) SetProgress(v string) *GetImageByIdResponseBodyDataImageRegionDistributeList {
	s.Progress = &v
	return s
}

func (s *GetImageByIdResponseBodyDataImageRegionDistributeList) SetRegionId(v string) *GetImageByIdResponseBodyDataImageRegionDistributeList {
	s.RegionId = &v
	return s
}

func (s *GetImageByIdResponseBodyDataImageRegionDistributeList) SetStatus(v string) *GetImageByIdResponseBodyDataImageRegionDistributeList {
	s.Status = &v
	return s
}

func (s *GetImageByIdResponseBodyDataImageRegionDistributeList) SetUnitImageId(v string) *GetImageByIdResponseBodyDataImageRegionDistributeList {
	s.UnitImageId = &v
	return s
}

func (s *GetImageByIdResponseBodyDataImageRegionDistributeList) SetVersionId(v string) *GetImageByIdResponseBodyDataImageRegionDistributeList {
	s.VersionId = &v
	return s
}

func (s *GetImageByIdResponseBodyDataImageRegionDistributeList) Validate() error {
	return dara.Validate(s)
}
