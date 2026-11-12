// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEffectiveImageByIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetEffectiveImageByIdResponseBody
	GetCode() *string
	SetData(v *GetEffectiveImageByIdResponseBodyData) *GetEffectiveImageByIdResponseBody
	GetData() *GetEffectiveImageByIdResponseBodyData
	SetMessage(v string) *GetEffectiveImageByIdResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetEffectiveImageByIdResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetEffectiveImageByIdResponseBody
	GetSuccess() *bool
}

type GetEffectiveImageByIdResponseBody struct {
	// The response code returned by the API. A value of 200 indicates success. If the call fails, the corresponding error code is returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The image information returned by the query.
	//
	// example:
	//
	// Object
	Data *GetEffectiveImageByIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message returned by the API.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID used to locate and troubleshoot the request.
	//
	// example:
	//
	// 50158E8B-992E-1286-B174-**********
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the API call is successful. Valid values: true: Successful. false: Failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetEffectiveImageByIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetEffectiveImageByIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetEffectiveImageByIdResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetEffectiveImageByIdResponseBody) GetData() *GetEffectiveImageByIdResponseBodyData {
	return s.Data
}

func (s *GetEffectiveImageByIdResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetEffectiveImageByIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetEffectiveImageByIdResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetEffectiveImageByIdResponseBody) SetCode(v string) *GetEffectiveImageByIdResponseBody {
	s.Code = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBody) SetData(v *GetEffectiveImageByIdResponseBodyData) *GetEffectiveImageByIdResponseBody {
	s.Data = v
	return s
}

func (s *GetEffectiveImageByIdResponseBody) SetMessage(v string) *GetEffectiveImageByIdResponseBody {
	s.Message = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBody) SetRequestId(v string) *GetEffectiveImageByIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBody) SetSuccess(v bool) *GetEffectiveImageByIdResponseBody {
	s.Success = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetEffectiveImageByIdResponseBodyData struct {
	// The ID of the Alibaba Cloud account to which the image belongs. The account ID for a system image is 0.
	//
	// example:
	//
	// 100931896542****
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// The business line to which the image belongs. Valid values:
	//
	// - 1: Basic business.
	//
	// - 2: Personal Edition.
	//
	// - 3: Cloud Office.
	//
	// - 4: Cloud Streaming.
	//
	// - 10: Standard Edition.
	//
	// example:
	//
	// 1
	BizType *int32 `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// The image description.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The description document of the image in Markdown format.
	//
	// example:
	//
	// ## this is a test document
	Document *string `json:"Document,omitempty" xml:"Document,omitempty"`
	// The environment ID.
	//
	// example:
	//
	// eid-***
	EnvironmentId *string `json:"EnvironmentId,omitempty" xml:"EnvironmentId,omitempty"`
	// The creation time. The time is in RFC 3339 format. The Z in the example indicates the UTC time zone.
	//
	// example:
	//
	// 2026-03-10T01:33:37Z
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The last update time. The time follows the RFC 3339 format. The Z in the example indicates the UTC time zone.
	//
	// example:
	//
	// 2025-12-11T02:19:27Z
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The URL of the image icon.
	//
	// example:
	//
	// https://xxx.jpg
	ImageIconUrl *string `json:"ImageIconUrl,omitempty" xml:"ImageIconUrl,omitempty"`
	// The image ID.
	//
	// example:
	//
	// img-bp1234567890abcde
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The details of the image version that actually takes effect in each region. For the same region, the configured newer version takes precedence.
	ImageInfoRegionModelList []*GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList `json:"ImageInfoRegionModelList,omitempty" xml:"ImageInfoRegionModelList,omitempty" type:"Repeated"`
	// The image name.
	//
	// example:
	//
	// DemoImage
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// The list of region IDs supported by the image.
	ImageRegionList []*string `json:"ImageRegionList,omitempty" xml:"ImageRegionList,omitempty" type:"Repeated"`
	// The image type. Valid values:
	//
	// - System: System image.
	//
	// - User: Custom image.
	//
	// - Shared: Shared image.
	//
	// - Community: Community image.
	//
	// example:
	//
	// User
	ImageType *string `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	// The ID of the latest version of the image.
	//
	// example:
	//
	// iv-xxx
	LatestVersionId *string `json:"LatestVersionId,omitempty" xml:"LatestVersionId,omitempty"`
	// The data cloud disk size of the online version. Unit: GiB.
	//
	// example:
	//
	// 100
	OnlineVersionDataDiskSize *int32 `json:"OnlineVersionDataDiskSize,omitempty" xml:"OnlineVersionDataDiskSize,omitempty"`
	// The ID of the current online version of the image.
	//
	// example:
	//
	// iv-xxxx
	OnlineVersionId *string `json:"OnlineVersionId,omitempty" xml:"OnlineVersionId,omitempty"`
	// The system cloud disk size of the online version. Unit: GiB.
	//
	// example:
	//
	// 40
	OnlineVersionSystemDiskSize *int32 `json:"OnlineVersionSystemDiskSize,omitempty" xml:"OnlineVersionSystemDiskSize,omitempty"`
	// Indicates whether disk encryption is enabled for the online version. Valid values: true: Enabled. false: Disabled.
	OnlineVersionVolumeEncryptionEnabled *bool `json:"OnlineVersionVolumeEncryptionEnabled,omitempty" xml:"OnlineVersionVolumeEncryptionEnabled,omitempty"`
	// The KMS key ID used when disk encryption is enabled for the online version.
	//
	// example:
	//
	// a7b3c0c8-xxxx
	OnlineVersionVolumeEncryptionKey *string `json:"OnlineVersionVolumeEncryptionKey,omitempty" xml:"OnlineVersionVolumeEncryptionKey,omitempty"`
	// The operating system type of the image. Valid values:
	//
	// - Windows: Windows operating system.
	//
	// - Linux: Linux operating system.
	//
	// - Android: Android operating system.
	//
	// - Unknown: Unknown operating system.
	//
	// example:
	//
	// Windows
	OsType *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	// The image packaging type. Valid values:
	//
	// - Ecs: ECS image.
	//
	// - Vhd: Virtual disk image.
	//
	// - Container: Container image.
	//
	// - Ecs_Container: Combined ECS and container image.
	//
	// - K8s_Node: Kubernetes node image.
	//
	// example:
	//
	// ECS
	PackageType *string `json:"PackageType,omitempty" xml:"PackageType,omitempty"`
	// The parent image ID, which indicates the image inheritance relationship. System images do not have parent images.
	//
	// example:
	//
	// imgc-xxx
	ParentImageId *string `json:"ParentImageId,omitempty" xml:"ParentImageId,omitempty"`
	// The product type to which the image belongs, such as CloudDesktop (Cloud Desktop), CloudApp (Cloud App), or WuyingServer (WUYING Workspace Enterprise Edition workstation).
	//
	// example:
	//
	// WuyingServer
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The connection protocol used by the image. Valid values:
	//
	// - ASP: Alibaba Cloud Streaming Protocol.
	//
	// - HDX: HDX protocol.
	//
	// example:
	//
	// ASP
	ProtocolType *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	// The rating score of the image, out of 100.
	//
	// example:
	//
	// 95
	Rating *int32 `json:"Rating,omitempty" xml:"Rating,omitempty"`
	// The image status. Valid values:
	//
	// - INIT: Initializing.
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
	// The list of image tags.
	TagList []*GetEffectiveImageByIdResponseBodyDataTagList `json:"TagList,omitempty" xml:"TagList,omitempty" type:"Repeated"`
	// The image version ID.
	//
	// example:
	//
	// iv-agde4****
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s GetEffectiveImageByIdResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetEffectiveImageByIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetEffectiveImageByIdResponseBodyData) GetAliUid() *int64 {
	return s.AliUid
}

func (s *GetEffectiveImageByIdResponseBodyData) GetBizType() *int32 {
	return s.BizType
}

func (s *GetEffectiveImageByIdResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetEffectiveImageByIdResponseBodyData) GetDocument() *string {
	return s.Document
}

func (s *GetEffectiveImageByIdResponseBodyData) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *GetEffectiveImageByIdResponseBodyData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetEffectiveImageByIdResponseBodyData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetEffectiveImageByIdResponseBodyData) GetImageIconUrl() *string {
	return s.ImageIconUrl
}

func (s *GetEffectiveImageByIdResponseBodyData) GetImageId() *string {
	return s.ImageId
}

func (s *GetEffectiveImageByIdResponseBodyData) GetImageInfoRegionModelList() []*GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList {
	return s.ImageInfoRegionModelList
}

func (s *GetEffectiveImageByIdResponseBodyData) GetImageName() *string {
	return s.ImageName
}

func (s *GetEffectiveImageByIdResponseBodyData) GetImageRegionList() []*string {
	return s.ImageRegionList
}

func (s *GetEffectiveImageByIdResponseBodyData) GetImageType() *string {
	return s.ImageType
}

func (s *GetEffectiveImageByIdResponseBodyData) GetLatestVersionId() *string {
	return s.LatestVersionId
}

func (s *GetEffectiveImageByIdResponseBodyData) GetOnlineVersionDataDiskSize() *int32 {
	return s.OnlineVersionDataDiskSize
}

func (s *GetEffectiveImageByIdResponseBodyData) GetOnlineVersionId() *string {
	return s.OnlineVersionId
}

func (s *GetEffectiveImageByIdResponseBodyData) GetOnlineVersionSystemDiskSize() *int32 {
	return s.OnlineVersionSystemDiskSize
}

func (s *GetEffectiveImageByIdResponseBodyData) GetOnlineVersionVolumeEncryptionEnabled() *bool {
	return s.OnlineVersionVolumeEncryptionEnabled
}

func (s *GetEffectiveImageByIdResponseBodyData) GetOnlineVersionVolumeEncryptionKey() *string {
	return s.OnlineVersionVolumeEncryptionKey
}

func (s *GetEffectiveImageByIdResponseBodyData) GetOsType() *string {
	return s.OsType
}

func (s *GetEffectiveImageByIdResponseBodyData) GetPackageType() *string {
	return s.PackageType
}

func (s *GetEffectiveImageByIdResponseBodyData) GetParentImageId() *string {
	return s.ParentImageId
}

func (s *GetEffectiveImageByIdResponseBodyData) GetProductType() *string {
	return s.ProductType
}

func (s *GetEffectiveImageByIdResponseBodyData) GetProtocolType() *string {
	return s.ProtocolType
}

func (s *GetEffectiveImageByIdResponseBodyData) GetRating() *int32 {
	return s.Rating
}

func (s *GetEffectiveImageByIdResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetEffectiveImageByIdResponseBodyData) GetTagList() []*GetEffectiveImageByIdResponseBodyDataTagList {
	return s.TagList
}

func (s *GetEffectiveImageByIdResponseBodyData) GetVersionId() *string {
	return s.VersionId
}

func (s *GetEffectiveImageByIdResponseBodyData) SetAliUid(v int64) *GetEffectiveImageByIdResponseBodyData {
	s.AliUid = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyData) SetBizType(v int32) *GetEffectiveImageByIdResponseBodyData {
	s.BizType = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyData) SetDescription(v string) *GetEffectiveImageByIdResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyData) SetDocument(v string) *GetEffectiveImageByIdResponseBodyData {
	s.Document = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyData) SetEnvironmentId(v string) *GetEffectiveImageByIdResponseBodyData {
	s.EnvironmentId = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyData) SetGmtCreate(v string) *GetEffectiveImageByIdResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyData) SetGmtModified(v string) *GetEffectiveImageByIdResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyData) SetImageIconUrl(v string) *GetEffectiveImageByIdResponseBodyData {
	s.ImageIconUrl = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyData) SetImageId(v string) *GetEffectiveImageByIdResponseBodyData {
	s.ImageId = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyData) SetImageInfoRegionModelList(v []*GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList) *GetEffectiveImageByIdResponseBodyData {
	s.ImageInfoRegionModelList = v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyData) SetImageName(v string) *GetEffectiveImageByIdResponseBodyData {
	s.ImageName = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyData) SetImageRegionList(v []*string) *GetEffectiveImageByIdResponseBodyData {
	s.ImageRegionList = v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyData) SetImageType(v string) *GetEffectiveImageByIdResponseBodyData {
	s.ImageType = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyData) SetLatestVersionId(v string) *GetEffectiveImageByIdResponseBodyData {
	s.LatestVersionId = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyData) SetOnlineVersionDataDiskSize(v int32) *GetEffectiveImageByIdResponseBodyData {
	s.OnlineVersionDataDiskSize = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyData) SetOnlineVersionId(v string) *GetEffectiveImageByIdResponseBodyData {
	s.OnlineVersionId = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyData) SetOnlineVersionSystemDiskSize(v int32) *GetEffectiveImageByIdResponseBodyData {
	s.OnlineVersionSystemDiskSize = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyData) SetOnlineVersionVolumeEncryptionEnabled(v bool) *GetEffectiveImageByIdResponseBodyData {
	s.OnlineVersionVolumeEncryptionEnabled = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyData) SetOnlineVersionVolumeEncryptionKey(v string) *GetEffectiveImageByIdResponseBodyData {
	s.OnlineVersionVolumeEncryptionKey = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyData) SetOsType(v string) *GetEffectiveImageByIdResponseBodyData {
	s.OsType = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyData) SetPackageType(v string) *GetEffectiveImageByIdResponseBodyData {
	s.PackageType = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyData) SetParentImageId(v string) *GetEffectiveImageByIdResponseBodyData {
	s.ParentImageId = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyData) SetProductType(v string) *GetEffectiveImageByIdResponseBodyData {
	s.ProductType = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyData) SetProtocolType(v string) *GetEffectiveImageByIdResponseBodyData {
	s.ProtocolType = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyData) SetRating(v int32) *GetEffectiveImageByIdResponseBodyData {
	s.Rating = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyData) SetStatus(v string) *GetEffectiveImageByIdResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyData) SetTagList(v []*GetEffectiveImageByIdResponseBodyDataTagList) *GetEffectiveImageByIdResponseBodyData {
	s.TagList = v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyData) SetVersionId(v string) *GetEffectiveImageByIdResponseBodyData {
	s.VersionId = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyData) Validate() error {
	if s.ImageInfoRegionModelList != nil {
		for _, item := range s.ImageInfoRegionModelList {
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

type GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList struct {
	// The list of adapted applications for the image version.
	AdaptorInfoBOList []*GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList `json:"AdaptorInfoBOList,omitempty" xml:"AdaptorInfoBOList,omitempty" type:"Repeated"`
	// The ID of the Alibaba Cloud account to which the image belongs. The account ID for a system image is 0.
	//
	// example:
	//
	// 13605346******
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// The application list, which displays only the application names and IDs.
	//
	// example:
	//
	// []
	AppList []*GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAppList `json:"AppList,omitempty" xml:"AppList,omitempty" type:"Repeated"`
	// The base image ID.
	//
	// example:
	//
	// imgc-dewgefwc******
	BaseImageId *string `json:"BaseImageId,omitempty" xml:"BaseImageId,omitempty"`
	// The base image version ID.
	//
	// example:
	//
	// iv-adbef134****
	BaseImageVersion *string `json:"BaseImageVersion,omitempty" xml:"BaseImageVersion,omitempty"`
	// The business line to which the image belongs. Valid values:
	//
	// - 1: Basic business.
	//
	// - 2: Personal Edition.
	//
	// - 3: Cloud Office.
	//
	// - 4: Cloud Streaming.
	//
	// - 10: Standard Edition.
	//
	// example:
	//
	// 1
	BizType *int32 `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// The name of the OSS bucket where the image file is stored.
	//
	// example:
	//
	// monitor-record-sh-twg2is
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// The image description.
	//
	// example:
	//
	// this is test image
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The list of feature tags supported by the image.
	FeatureList []*string `json:"FeatureList,omitempty" xml:"FeatureList,omitempty" type:"Repeated"`
	// The FOTA version.
	//
	// example:
	//
	// 3.0.11-R-20260624.195825
	FotaVersion *string `json:"FotaVersion,omitempty" xml:"FotaVersion,omitempty"`
	// The creation time. The time is in RFC 3339 format. The Z in the example indicates the UTC time zone.
	//
	// example:
	//
	// 2023-09-12T13:01:09Z
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The last update time. The time follows the RFC 3339 format. The Z in the example indicates the UTC time zone.
	//
	// example:
	//
	// 2026-07-16T07:39:43Z
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The image creation method. For example, BUILDER_MANUAL indicates manual creation.
	//
	// example:
	//
	// BUILDER_MANUAL
	ImageCreateMode *string `json:"ImageCreateMode,omitempty" xml:"ImageCreateMode,omitempty"`
	// The image extension information, including disk specifications and GPU driver information. VHD or Container images use the extension information of their base image and base version.
	//
	// example:
	//
	// {}
	ImageExtendInfoModel *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListImageExtendInfoModel `json:"ImageExtendInfoModel,omitempty" xml:"ImageExtendInfoModel,omitempty" type:"Struct"`
	// The image ID.
	//
	// example:
	//
	// imgc-0aae4rgk5yaa*****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The image name.
	//
	// example:
	//
	// image0625
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// The image distribution information for the current region.
	ImageRegionDistribute *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListImageRegionDistribute `json:"ImageRegionDistribute,omitempty" xml:"ImageRegionDistribute,omitempty" type:"Struct"`
	// The image source. For example, center indicates the central image service.
	//
	// example:
	//
	// center
	ImageSource *string `json:"ImageSource,omitempty" xml:"ImageSource,omitempty"`
	// The image type. Valid values:
	//
	// - System: System image.
	//
	// - User: Custom image.
	//
	// - Shared: Shared image.
	//
	// - Community: Community image.
	//
	// example:
	//
	// System
	ImageType *string `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	// The image language.
	//
	// example:
	//
	// zh-CN
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The ID of the latest version of the image.
	//
	// example:
	//
	// iv-1wder3ere****
	LatestVersionId *string `json:"LatestVersionId,omitempty" xml:"LatestVersionId,omitempty"`
	// Indicates whether the returned version is the active version. Valid values: true: Yes. false: No.
	//
	// example:
	//
	// true
	OnlineVersion *bool `json:"OnlineVersion,omitempty" xml:"OnlineVersion,omitempty"`
	// The ID of the current online version of the image.
	//
	// example:
	//
	// iv-1wder3ere****
	OnlineVersionId *string `json:"OnlineVersionId,omitempty" xml:"OnlineVersionId,omitempty"`
	// The operating system type of the image. Valid values:
	//
	// - Windows: Windows operating system.
	//
	// - Linux: Linux operating system.
	//
	// - Android: Android operating system.
	//
	// - Unknown: Unknown operating system.
	//
	// example:
	//
	// Windows
	OsType *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	// The image packaging type. Valid values:
	//
	// - Ecs: ECS image.
	//
	// - Vhd: Virtual disk image.
	//
	// - Container: Container image.
	//
	// - Ecs_Container: Combined ECS and container image.
	//
	// - K8s_Node: Kubernetes node image.
	//
	// example:
	//
	// ECS
	PackageType *string `json:"PackageType,omitempty" xml:"PackageType,omitempty"`
	// The parent image ID, which indicates the image inheritance relationship. System images do not have parent images.
	//
	// example:
	//
	// desktopimage-windows-11-64-asp
	ParentImageId *string `json:"ParentImageId,omitempty" xml:"ParentImageId,omitempty"`
	// The parent image version ID.
	//
	// example:
	//
	// iv-12ededgjj***
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
	// The product type to which the image belongs, such as CloudDesktop (Cloud Desktop), CloudApp (Cloud App), or WuyingServer (WUYING Workspace Enterprise Edition workstation).
	//
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The extended properties of the image version in key-value pairs.
	//
	// example:
	//
	// {}
	PropertyMap map[string]*string `json:"PropertyMap,omitempty" xml:"PropertyMap,omitempty"`
	// The connection protocol used by the image. Valid values:
	//
	// - ASP: Alibaba Cloud Streaming Protocol.
	//
	// - HDX: HDX protocol.
	//
	// example:
	//
	// ASP
	ProtocolType *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	// The ID of the region to which the image information belongs.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The image status. Valid values:
	//
	// - INIT: Initializing.
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

func (s GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList) String() string {
	return dara.Prettify(s)
}

func (s GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList) GoString() string {
	return s.String()
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList) GetAdaptorInfoBOList() []*GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList {
	return s.AdaptorInfoBOList
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList) GetAliUid() *int64 {
	return s.AliUid
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList) GetAppList() []*GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAppList {
	return s.AppList
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList) GetBaseImageId() *string {
	return s.BaseImageId
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList) GetBaseImageVersion() *string {
	return s.BaseImageVersion
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList) GetBizType() *int32 {
	return s.BizType
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList) GetBucketName() *string {
	return s.BucketName
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList) GetDescription() *string {
	return s.Description
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList) GetFeatureList() []*string {
	return s.FeatureList
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList) GetFotaVersion() *string {
	return s.FotaVersion
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList) GetImageCreateMode() *string {
	return s.ImageCreateMode
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList) GetImageExtendInfoModel() *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListImageExtendInfoModel {
	return s.ImageExtendInfoModel
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList) GetImageId() *string {
	return s.ImageId
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList) GetImageName() *string {
	return s.ImageName
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList) GetImageRegionDistribute() *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListImageRegionDistribute {
	return s.ImageRegionDistribute
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList) GetImageSource() *string {
	return s.ImageSource
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList) GetImageType() *string {
	return s.ImageType
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList) GetLanguage() *string {
	return s.Language
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList) GetLatestVersionId() *string {
	return s.LatestVersionId
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList) GetOnlineVersion() *bool {
	return s.OnlineVersion
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList) GetOnlineVersionId() *string {
	return s.OnlineVersionId
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList) GetOsType() *string {
	return s.OsType
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList) GetPackageType() *string {
	return s.PackageType
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList) GetParentImageId() *string {
	return s.ParentImageId
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList) GetParentImageVersion() *string {
	return s.ParentImageVersion
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList) GetPlatform() *int32 {
	return s.Platform
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList) GetPlatformName() *string {
	return s.PlatformName
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList) GetProductType() *string {
	return s.ProductType
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList) GetPropertyMap() map[string]*string {
	return s.PropertyMap
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList) GetProtocolType() *string {
	return s.ProtocolType
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList) GetRegionId() *string {
	return s.RegionId
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList) GetStatus() *string {
	return s.Status
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList) GetVersionId() *string {
	return s.VersionId
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList) SetAdaptorInfoBOList(v []*GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList) *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList {
	s.AdaptorInfoBOList = v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList) SetAliUid(v int64) *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList {
	s.AliUid = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList) SetAppList(v []*GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAppList) *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList {
	s.AppList = v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList) SetBaseImageId(v string) *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList {
	s.BaseImageId = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList) SetBaseImageVersion(v string) *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList {
	s.BaseImageVersion = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList) SetBizType(v int32) *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList {
	s.BizType = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList) SetBucketName(v string) *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList {
	s.BucketName = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList) SetDescription(v string) *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList {
	s.Description = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList) SetFeatureList(v []*string) *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList {
	s.FeatureList = v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList) SetFotaVersion(v string) *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList {
	s.FotaVersion = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList) SetGmtCreate(v string) *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList {
	s.GmtCreate = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList) SetGmtModified(v string) *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList {
	s.GmtModified = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList) SetImageCreateMode(v string) *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList {
	s.ImageCreateMode = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList) SetImageExtendInfoModel(v *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListImageExtendInfoModel) *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList {
	s.ImageExtendInfoModel = v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList) SetImageId(v string) *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList {
	s.ImageId = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList) SetImageName(v string) *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList {
	s.ImageName = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList) SetImageRegionDistribute(v *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListImageRegionDistribute) *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList {
	s.ImageRegionDistribute = v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList) SetImageSource(v string) *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList {
	s.ImageSource = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList) SetImageType(v string) *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList {
	s.ImageType = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList) SetLanguage(v string) *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList {
	s.Language = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList) SetLatestVersionId(v string) *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList {
	s.LatestVersionId = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList) SetOnlineVersion(v bool) *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList {
	s.OnlineVersion = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList) SetOnlineVersionId(v string) *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList {
	s.OnlineVersionId = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList) SetOsType(v string) *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList {
	s.OsType = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList) SetPackageType(v string) *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList {
	s.PackageType = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList) SetParentImageId(v string) *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList {
	s.ParentImageId = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList) SetParentImageVersion(v string) *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList {
	s.ParentImageVersion = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList) SetPlatform(v int32) *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList {
	s.Platform = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList) SetPlatformName(v string) *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList {
	s.PlatformName = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList) SetProductType(v string) *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList {
	s.ProductType = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList) SetPropertyMap(v map[string]*string) *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList {
	s.PropertyMap = v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList) SetProtocolType(v string) *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList {
	s.ProtocolType = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList) SetRegionId(v string) *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList {
	s.RegionId = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList) SetStatus(v string) *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList {
	s.Status = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList) SetVersionId(v string) *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList {
	s.VersionId = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelList) Validate() error {
	if s.AdaptorInfoBOList != nil {
		for _, item := range s.AdaptorInfoBOList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.AppList != nil {
		for _, item := range s.AppList {
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
	if s.ImageRegionDistribute != nil {
		if err := s.ImageRegionDistribute.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList struct {
	// The adapter ID.
	//
	// example:
	//
	// adt-13ff3****
	AdaptorId *string `json:"AdaptorId,omitempty" xml:"AdaptorId,omitempty"`
	// The adapter message.
	//
	// example:
	//
	// success
	AdaptorMessage *string `json:"AdaptorMessage,omitempty" xml:"AdaptorMessage,omitempty"`
	// The adaptation mode.
	//
	// example:
	//
	// MANUAL
	AdaptorMode *string `json:"AdaptorMode,omitempty" xml:"AdaptorMode,omitempty"`
	// The adaptation phase.
	//
	// example:
	//
	// FINISH_ADAPT
	AdaptorStep *string `json:"AdaptorStep,omitempty" xml:"AdaptorStep,omitempty"`
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
	// The business line ID of the application.
	//
	// example:
	//
	// 1
	BizType *int32 `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// The name of the OSS bucket where the application files are stored.
	//
	// example:
	//
	// xiaowutest-j
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// The level-2 application category, such as design software, office, or browser. This parameter is not available to the public.
	//
	// example:
	//
	// 123
	CategoryId *int32 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// The level-1 application category, such as cloud application or cloud game. This parameter is not available to the public.
	//
	// example:
	//
	// 1
	CategoryType *int32 `json:"CategoryType,omitempty" xml:"CategoryType,omitempty"`
	// The control parameter. This parameter can be configured by tenants and internal administrators and is not available to the public.
	//
	// example:
	//
	// {}
	ControlParam *string `json:"ControlParam,omitempty" xml:"ControlParam,omitempty"`
	// The path of the executable file. This parameter is not available to the public.
	//
	// example:
	//
	// abc.text
	ExecuteFileAddress *string `json:"ExecuteFileAddress,omitempty" xml:"ExecuteFileAddress,omitempty"`
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
	// The file size. Unit: KB.
	//
	// example:
	//
	// 102400
	FileSize *int64 `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	// The optimization parameter configured in the backend by internal administrators.
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
	// The operating system type applicable to the application.
	//
	// example:
	//
	// Windows
	OsType *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	// The application service port. For applications that require port-based access, specify the actual port number that the application listens on.
	//
	// example:
	//
	// 9000
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The product type of the application.
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
	// Specifies whether silent installation is enabled. This parameter can be configured by tenants and internal administrators and is not available to the public.
	//
	// example:
	//
	// {}
	SilenceFlag *string `json:"SilenceFlag,omitempty" xml:"SilenceFlag,omitempty"`
	// The silent installation parameter. This parameter is not available to the public.
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
	// The URL path suffix for application access. This suffix and the application service port together describe the access entry point.
	//
	// example:
	//
	// /index
	UrlSuffix *string `json:"UrlSuffix,omitempty" xml:"UrlSuffix,omitempty"`
}

func (s GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList) String() string {
	return dara.Prettify(s)
}

func (s GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList) GoString() string {
	return s.String()
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList) GetAdaptorId() *string {
	return s.AdaptorId
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList) GetAdaptorMessage() *string {
	return s.AdaptorMessage
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList) GetAdaptorMode() *string {
	return s.AdaptorMode
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList) GetAdaptorStep() *string {
	return s.AdaptorStep
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList) GetAppIconUrl() *string {
	return s.AppIconUrl
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList) GetAppId() *string {
	return s.AppId
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList) GetAppName() *string {
	return s.AppName
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList) GetAppScript() *string {
	return s.AppScript
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList) GetAppType() *string {
	return s.AppType
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList) GetAppVersionId() *string {
	return s.AppVersionId
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList) GetAppVersionName() *string {
	return s.AppVersionName
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList) GetAppVersionNumber() *string {
	return s.AppVersionNumber
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList) GetBizType() *int32 {
	return s.BizType
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList) GetBucketName() *string {
	return s.BucketName
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList) GetCategoryId() *int32 {
	return s.CategoryId
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList) GetCategoryType() *int32 {
	return s.CategoryType
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList) GetControlParam() *string {
	return s.ControlParam
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList) GetExecuteFileAddress() *string {
	return s.ExecuteFileAddress
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList) GetFileName() *string {
	return s.FileName
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList) GetFilePath() *string {
	return s.FilePath
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList) GetFileSize() *int64 {
	return s.FileSize
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList) GetOptimizeParam() *string {
	return s.OptimizeParam
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList) GetOriginAppType() *string {
	return s.OriginAppType
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList) GetOsType() *string {
	return s.OsType
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList) GetPort() *string {
	return s.Port
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList) GetProductType() *string {
	return s.ProductType
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList) GetRuntimeEnv() *string {
	return s.RuntimeEnv
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList) GetSilenceFlag() *string {
	return s.SilenceFlag
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList) GetSilenceParam() *string {
	return s.SilenceParam
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList) GetStatus() *string {
	return s.Status
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList) GetSubAppType() *string {
	return s.SubAppType
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList) GetUrlSuffix() *string {
	return s.UrlSuffix
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList) SetAdaptorId(v string) *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList {
	s.AdaptorId = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList) SetAdaptorMessage(v string) *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList {
	s.AdaptorMessage = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList) SetAdaptorMode(v string) *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList {
	s.AdaptorMode = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList) SetAdaptorStep(v string) *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList {
	s.AdaptorStep = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList) SetAppIconUrl(v string) *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList {
	s.AppIconUrl = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList) SetAppId(v string) *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList {
	s.AppId = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList) SetAppName(v string) *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList {
	s.AppName = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList) SetAppScript(v string) *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList {
	s.AppScript = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList) SetAppType(v string) *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList {
	s.AppType = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList) SetAppVersionId(v string) *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList {
	s.AppVersionId = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList) SetAppVersionName(v string) *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList {
	s.AppVersionName = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList) SetAppVersionNumber(v string) *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList {
	s.AppVersionNumber = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList) SetBizType(v int32) *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList {
	s.BizType = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList) SetBucketName(v string) *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList {
	s.BucketName = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList) SetCategoryId(v int32) *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList {
	s.CategoryId = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList) SetCategoryType(v int32) *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList {
	s.CategoryType = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList) SetControlParam(v string) *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList {
	s.ControlParam = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList) SetExecuteFileAddress(v string) *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList {
	s.ExecuteFileAddress = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList) SetFileName(v string) *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList {
	s.FileName = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList) SetFilePath(v string) *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList {
	s.FilePath = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList) SetFileSize(v int64) *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList {
	s.FileSize = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList) SetOptimizeParam(v string) *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList {
	s.OptimizeParam = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList) SetOriginAppType(v string) *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList {
	s.OriginAppType = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList) SetOsType(v string) *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList {
	s.OsType = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList) SetPort(v string) *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList {
	s.Port = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList) SetProductType(v string) *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList {
	s.ProductType = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList) SetRuntimeEnv(v string) *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList {
	s.RuntimeEnv = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList) SetSilenceFlag(v string) *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList {
	s.SilenceFlag = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList) SetSilenceParam(v string) *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList {
	s.SilenceParam = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList) SetStatus(v string) *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList {
	s.Status = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList) SetSubAppType(v string) *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList {
	s.SubAppType = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList) SetUrlSuffix(v string) *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList {
	s.UrlSuffix = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAdaptorInfoBOList) Validate() error {
	return dara.Validate(s)
}

type GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAppList struct {
	// The application ID.
	//
	// example:
	//
	// app-0aae4rxw78****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The application name.
	//
	// example:
	//
	// idb-ent-console
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
}

func (s GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAppList) String() string {
	return dara.Prettify(s)
}

func (s GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAppList) GoString() string {
	return s.String()
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAppList) GetAppId() *string {
	return s.AppId
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAppList) GetAppName() *string {
	return s.AppName
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAppList) SetAppId(v string) *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAppList {
	s.AppId = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAppList) SetAppName(v string) *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAppList {
	s.AppName = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListAppList) Validate() error {
	return dara.Validate(s)
}

type GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListImageExtendInfoModel struct {
	// The size of the data cloud disk. Unit: GiB.
	//
	// example:
	//
	// 100
	DataDiskSize *int32 `json:"DataDiskSize,omitempty" xml:"DataDiskSize,omitempty"`
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
	// The size of the system cloud disk. Unit: GiB.
	//
	// example:
	//
	// 50
	SystemDiskSize *int32 `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
}

func (s GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListImageExtendInfoModel) String() string {
	return dara.Prettify(s)
}

func (s GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListImageExtendInfoModel) GoString() string {
	return s.String()
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListImageExtendInfoModel) GetDataDiskSize() *int32 {
	return s.DataDiskSize
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListImageExtendInfoModel) GetGpuCategory() *int32 {
	return s.GpuCategory
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListImageExtendInfoModel) GetGpuDriverVersion() *string {
	return s.GpuDriverVersion
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListImageExtendInfoModel) GetSystemDiskSize() *int32 {
	return s.SystemDiskSize
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListImageExtendInfoModel) SetDataDiskSize(v int32) *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListImageExtendInfoModel {
	s.DataDiskSize = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListImageExtendInfoModel) SetGpuCategory(v int32) *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListImageExtendInfoModel {
	s.GpuCategory = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListImageExtendInfoModel) SetGpuDriverVersion(v string) *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListImageExtendInfoModel {
	s.GpuDriverVersion = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListImageExtendInfoModel) SetSystemDiskSize(v int32) *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListImageExtendInfoModel {
	s.SystemDiskSize = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListImageExtendInfoModel) Validate() error {
	return dara.Validate(s)
}

type GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListImageRegionDistribute struct {
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
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The region distribution status. Valid values: CREATING: distributing. AVAILABLE: available. CREATE_FAILED: distribution failed. DELETE: deleted.
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
}

func (s GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListImageRegionDistribute) String() string {
	return dara.Prettify(s)
}

func (s GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListImageRegionDistribute) GoString() string {
	return s.String()
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListImageRegionDistribute) GetMessage() *string {
	return s.Message
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListImageRegionDistribute) GetProgress() *string {
	return s.Progress
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListImageRegionDistribute) GetRegionId() *string {
	return s.RegionId
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListImageRegionDistribute) GetStatus() *string {
	return s.Status
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListImageRegionDistribute) GetUnitImageId() *string {
	return s.UnitImageId
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListImageRegionDistribute) SetMessage(v string) *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListImageRegionDistribute {
	s.Message = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListImageRegionDistribute) SetProgress(v string) *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListImageRegionDistribute {
	s.Progress = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListImageRegionDistribute) SetRegionId(v string) *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListImageRegionDistribute {
	s.RegionId = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListImageRegionDistribute) SetStatus(v string) *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListImageRegionDistribute {
	s.Status = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListImageRegionDistribute) SetUnitImageId(v string) *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListImageRegionDistribute {
	s.UnitImageId = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyDataImageInfoRegionModelListImageRegionDistribute) Validate() error {
	return dara.Validate(s)
}

type GetEffectiveImageByIdResponseBodyDataTagList struct {
	// The tag key.
	//
	// example:
	//
	// env
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// inner
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEffectiveImageByIdResponseBodyDataTagList) String() string {
	return dara.Prettify(s)
}

func (s GetEffectiveImageByIdResponseBodyDataTagList) GoString() string {
	return s.String()
}

func (s *GetEffectiveImageByIdResponseBodyDataTagList) GetKey() *string {
	return s.Key
}

func (s *GetEffectiveImageByIdResponseBodyDataTagList) GetValue() *string {
	return s.Value
}

func (s *GetEffectiveImageByIdResponseBodyDataTagList) SetKey(v string) *GetEffectiveImageByIdResponseBodyDataTagList {
	s.Key = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyDataTagList) SetValue(v string) *GetEffectiveImageByIdResponseBodyDataTagList {
	s.Value = &v
	return s
}

func (s *GetEffectiveImageByIdResponseBodyDataTagList) Validate() error {
	return dara.Validate(s)
}
