// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListImageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizRegionIdList(v []*string) *ListImageRequest
	GetBizRegionIdList() []*string
	SetBizType(v int32) *ListImageRequest
	GetBizType() *int32
	SetBizTypeList(v []*int32) *ListImageRequest
	GetBizTypeList() []*int32
	SetFeatureList(v []*string) *ListImageRequest
	GetFeatureList() []*string
	SetFotaVersion(v string) *ListImageRequest
	GetFotaVersion() *string
	SetImageId(v string) *ListImageRequest
	GetImageId() *string
	SetImageName(v string) *ListImageRequest
	GetImageName() *string
	SetImageType(v string) *ListImageRequest
	GetImageType() *string
	SetLanguageType(v string) *ListImageRequest
	GetLanguageType() *string
	SetOsType(v string) *ListImageRequest
	GetOsType() *string
	SetPackageType(v string) *ListImageRequest
	GetPackageType() *string
	SetPageNumber(v int32) *ListImageRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListImageRequest
	GetPageSize() *int32
	SetPlatformName(v string) *ListImageRequest
	GetPlatformName() *string
	SetPlatformNameList(v []*string) *ListImageRequest
	GetPlatformNameList() []*string
	SetProductType(v string) *ListImageRequest
	GetProductType() *string
	SetProductTypeList(v []*string) *ListImageRequest
	GetProductTypeList() []*string
	SetProtocolType(v string) *ListImageRequest
	GetProtocolType() *string
	SetResourceInstanceType(v string) *ListImageRequest
	GetResourceInstanceType() *string
	SetStatus(v string) *ListImageRequest
	GetStatus() *string
	SetTagList(v []*ListImageRequestTagList) *ListImageRequest
	GetTagList() []*ListImageRequestTagList
}

type ListImageRequest struct {
	// The regions that are supported. The EDS images are centralized. Use this parameter to query the regions where the image is deployed.
	BizRegionIdList []*string `json:"BizRegionIdList,omitempty" xml:"BizRegionIdList,omitempty" type:"Repeated"`
	// The service type. This parameter is not available publicly.
	//
	// Valid value:
	//
	// 	- 1 (default)
	//
	// example:
	//
	// 1
	BizType *int32 `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// The list of all service types. It is not available publicly.
	BizTypeList []*int32 `json:"BizTypeList,omitempty" xml:"BizTypeList,omitempty" type:"Repeated"`
	// The features supported by the image.
	FeatureList []*string `json:"FeatureList,omitempty" xml:"FeatureList,omitempty" type:"Repeated"`
	// The image version.
	//
	// example:
	//
	// 2.0.3-xxxx
	FotaVersion *string `json:"FotaVersion,omitempty" xml:"FotaVersion,omitempty"`
	// The image ID.
	//
	// example:
	//
	// img-bp13mu****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The image name. Fuzzy match is supported.
	//
	// example:
	//
	// DemoImage
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// The type of the images.
	//
	// Valid values:
	//
	// 	- User: a custom image.
	//
	// 	- Shared: a shared image.
	//
	// 	- System: a system image.
	//
	// 	- Community: a community image.
	//
	// example:
	//
	// User
	ImageType *string `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	// The language.
	//
	// Valid values:
	//
	// 	- en-US: English.
	//
	// 	- zh-HK: Chinese, Traditional (Hong Kong, China).
	//
	// 	- zh-CN: Simplified Chinese.
	//
	// 	- ja-JP: Japanese.
	//
	// example:
	//
	// zh-CN
	LanguageType *string `json:"LanguageType,omitempty" xml:"LanguageType,omitempty"`
	// The OS type of the image.
	//
	// Valid values:
	//
	// 	- Linux
	//
	// 	- Unknown
	//
	// 	- Windows
	//
	// 	- Android
	//
	// example:
	//
	// Windows
	OsType *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	// The image encapsulation type.
	//
	// Valid values:
	//
	// 	- Ecs_Container: ECS and Docker image
	//
	// 	- Ecs: ECS image
	//
	// example:
	//
	// Ecs
	PackageType *string `json:"PackageType,omitempty" xml:"PackageType,omitempty"`
	// The current page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Maximum value: 100. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the operating system platform.
	//
	// Valid values:
	//
	// 	- Ubuntu
	//
	// 	- Debian
	//
	// 	- Windows Server 2022
	//
	// 	- Windows Server 2019
	//
	// 	- Windows Server 2016
	//
	// 	- Windows 11
	//
	// 	- Windows 10
	//
	// example:
	//
	// Windows Server 2019
	PlatformName *string `json:"PlatformName,omitempty" xml:"PlatformName,omitempty"`
	// The list of supported platform types. For valid values, refer to PlatformName above.
	PlatformNameList []*string `json:"PlatformNameList,omitempty" xml:"PlatformNameList,omitempty" type:"Repeated"`
	// The product type.
	//
	// Valid values:
	//
	// 	- CloudDesktop: Elastic Desktop Service
	//
	// 	- CloudApp: App Streaming
	//
	// 	- WuyingServer: Workstation
	//
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The list of products that are supported when the image supports multiple products.
	ProductTypeList []*string `json:"ProductTypeList,omitempty" xml:"ProductTypeList,omitempty" type:"Repeated"`
	// The protocol type of the image.
	//
	// Valid values:
	//
	// 	- HDX: the High-definition Experience (HDX) protocol
	//
	// 	- ASP: the Alibaba Cloud-developed ASP protocol
	//
	// example:
	//
	// ASP
	ProtocolType *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	// Find images with certain fixed specifications.
	//
	// example:
	//
	// eds.enterprise_office.2c4g
	ResourceInstanceType *string `json:"ResourceInstanceType,omitempty" xml:"ResourceInstanceType,omitempty"`
	// The status of the image. You can query images in the specified status. By default, all images in the Not Deleted state are queried.
	//
	// Valid values:
	//
	// 	- AVAILABLE: The image is available.
	//
	// 	- INIT: The image is being initialized.
	//
	// 	- CREATE_FAILED: The image failed to be created.
	//
	// 	- CREATING: The image is being created.
	//
	// example:
	//
	// INIT
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags to query.
	TagList []*ListImageRequestTagList `json:"TagList,omitempty" xml:"TagList,omitempty" type:"Repeated"`
}

func (s ListImageRequest) String() string {
	return dara.Prettify(s)
}

func (s ListImageRequest) GoString() string {
	return s.String()
}

func (s *ListImageRequest) GetBizRegionIdList() []*string {
	return s.BizRegionIdList
}

func (s *ListImageRequest) GetBizType() *int32 {
	return s.BizType
}

func (s *ListImageRequest) GetBizTypeList() []*int32 {
	return s.BizTypeList
}

func (s *ListImageRequest) GetFeatureList() []*string {
	return s.FeatureList
}

func (s *ListImageRequest) GetFotaVersion() *string {
	return s.FotaVersion
}

func (s *ListImageRequest) GetImageId() *string {
	return s.ImageId
}

func (s *ListImageRequest) GetImageName() *string {
	return s.ImageName
}

func (s *ListImageRequest) GetImageType() *string {
	return s.ImageType
}

func (s *ListImageRequest) GetLanguageType() *string {
	return s.LanguageType
}

func (s *ListImageRequest) GetOsType() *string {
	return s.OsType
}

func (s *ListImageRequest) GetPackageType() *string {
	return s.PackageType
}

func (s *ListImageRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListImageRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListImageRequest) GetPlatformName() *string {
	return s.PlatformName
}

func (s *ListImageRequest) GetPlatformNameList() []*string {
	return s.PlatformNameList
}

func (s *ListImageRequest) GetProductType() *string {
	return s.ProductType
}

func (s *ListImageRequest) GetProductTypeList() []*string {
	return s.ProductTypeList
}

func (s *ListImageRequest) GetProtocolType() *string {
	return s.ProtocolType
}

func (s *ListImageRequest) GetResourceInstanceType() *string {
	return s.ResourceInstanceType
}

func (s *ListImageRequest) GetStatus() *string {
	return s.Status
}

func (s *ListImageRequest) GetTagList() []*ListImageRequestTagList {
	return s.TagList
}

func (s *ListImageRequest) SetBizRegionIdList(v []*string) *ListImageRequest {
	s.BizRegionIdList = v
	return s
}

func (s *ListImageRequest) SetBizType(v int32) *ListImageRequest {
	s.BizType = &v
	return s
}

func (s *ListImageRequest) SetBizTypeList(v []*int32) *ListImageRequest {
	s.BizTypeList = v
	return s
}

func (s *ListImageRequest) SetFeatureList(v []*string) *ListImageRequest {
	s.FeatureList = v
	return s
}

func (s *ListImageRequest) SetFotaVersion(v string) *ListImageRequest {
	s.FotaVersion = &v
	return s
}

func (s *ListImageRequest) SetImageId(v string) *ListImageRequest {
	s.ImageId = &v
	return s
}

func (s *ListImageRequest) SetImageName(v string) *ListImageRequest {
	s.ImageName = &v
	return s
}

func (s *ListImageRequest) SetImageType(v string) *ListImageRequest {
	s.ImageType = &v
	return s
}

func (s *ListImageRequest) SetLanguageType(v string) *ListImageRequest {
	s.LanguageType = &v
	return s
}

func (s *ListImageRequest) SetOsType(v string) *ListImageRequest {
	s.OsType = &v
	return s
}

func (s *ListImageRequest) SetPackageType(v string) *ListImageRequest {
	s.PackageType = &v
	return s
}

func (s *ListImageRequest) SetPageNumber(v int32) *ListImageRequest {
	s.PageNumber = &v
	return s
}

func (s *ListImageRequest) SetPageSize(v int32) *ListImageRequest {
	s.PageSize = &v
	return s
}

func (s *ListImageRequest) SetPlatformName(v string) *ListImageRequest {
	s.PlatformName = &v
	return s
}

func (s *ListImageRequest) SetPlatformNameList(v []*string) *ListImageRequest {
	s.PlatformNameList = v
	return s
}

func (s *ListImageRequest) SetProductType(v string) *ListImageRequest {
	s.ProductType = &v
	return s
}

func (s *ListImageRequest) SetProductTypeList(v []*string) *ListImageRequest {
	s.ProductTypeList = v
	return s
}

func (s *ListImageRequest) SetProtocolType(v string) *ListImageRequest {
	s.ProtocolType = &v
	return s
}

func (s *ListImageRequest) SetResourceInstanceType(v string) *ListImageRequest {
	s.ResourceInstanceType = &v
	return s
}

func (s *ListImageRequest) SetStatus(v string) *ListImageRequest {
	s.Status = &v
	return s
}

func (s *ListImageRequest) SetTagList(v []*ListImageRequestTagList) *ListImageRequest {
	s.TagList = v
	return s
}

func (s *ListImageRequest) Validate() error {
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

type ListImageRequestTagList struct {
	// The key of the custom tag.
	//
	// example:
	//
	// env
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the custom tag.
	//
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListImageRequestTagList) String() string {
	return dara.Prettify(s)
}

func (s ListImageRequestTagList) GoString() string {
	return s.String()
}

func (s *ListImageRequestTagList) GetKey() *string {
	return s.Key
}

func (s *ListImageRequestTagList) GetValue() *string {
	return s.Value
}

func (s *ListImageRequestTagList) SetKey(v string) *ListImageRequestTagList {
	s.Key = &v
	return s
}

func (s *ListImageRequestTagList) SetValue(v string) *ListImageRequestTagList {
	s.Value = &v
	return s
}

func (s *ListImageRequestTagList) Validate() error {
	return dara.Validate(s)
}
