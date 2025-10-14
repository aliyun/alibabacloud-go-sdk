// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAvailableResourceInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetImages(v *DescribeAvailableResourceInfoResponseBodyImages) *DescribeAvailableResourceInfoResponseBody
	GetImages() *DescribeAvailableResourceInfoResponseBodyImages
	SetRequestId(v string) *DescribeAvailableResourceInfoResponseBody
	GetRequestId() *string
	SetSupportResources(v *DescribeAvailableResourceInfoResponseBodySupportResources) *DescribeAvailableResourceInfoResponseBody
	GetSupportResources() *DescribeAvailableResourceInfoResponseBodySupportResources
}

type DescribeAvailableResourceInfoResponseBody struct {
	// The information about the image.
	Images *DescribeAvailableResourceInfoResponseBodyImages `json:"Images,omitempty" xml:"Images,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 8629F679-B51D-4194-A1CC-5D8F504C362B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The specifications of resources that you can purchase.
	SupportResources *DescribeAvailableResourceInfoResponseBodySupportResources `json:"SupportResources,omitempty" xml:"SupportResources,omitempty" type:"Struct"`
}

func (s DescribeAvailableResourceInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableResourceInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceInfoResponseBody) GetImages() *DescribeAvailableResourceInfoResponseBodyImages {
	return s.Images
}

func (s *DescribeAvailableResourceInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAvailableResourceInfoResponseBody) GetSupportResources() *DescribeAvailableResourceInfoResponseBodySupportResources {
	return s.SupportResources
}

func (s *DescribeAvailableResourceInfoResponseBody) SetImages(v *DescribeAvailableResourceInfoResponseBodyImages) *DescribeAvailableResourceInfoResponseBody {
	s.Images = v
	return s
}

func (s *DescribeAvailableResourceInfoResponseBody) SetRequestId(v string) *DescribeAvailableResourceInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAvailableResourceInfoResponseBody) SetSupportResources(v *DescribeAvailableResourceInfoResponseBodySupportResources) *DescribeAvailableResourceInfoResponseBody {
	s.SupportResources = v
	return s
}

func (s *DescribeAvailableResourceInfoResponseBody) Validate() error {
	if s.Images != nil {
		if err := s.Images.Validate(); err != nil {
			return err
		}
	}
	if s.SupportResources != nil {
		if err := s.SupportResources.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAvailableResourceInfoResponseBodyImages struct {
	Image []*DescribeAvailableResourceInfoResponseBodyImagesImage `json:"Image,omitempty" xml:"Image,omitempty" type:"Repeated"`
}

func (s DescribeAvailableResourceInfoResponseBodyImages) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableResourceInfoResponseBodyImages) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceInfoResponseBodyImages) GetImage() []*DescribeAvailableResourceInfoResponseBodyImagesImage {
	return s.Image
}

func (s *DescribeAvailableResourceInfoResponseBodyImages) SetImage(v []*DescribeAvailableResourceInfoResponseBodyImagesImage) *DescribeAvailableResourceInfoResponseBodyImages {
	s.Image = v
	return s
}

func (s *DescribeAvailableResourceInfoResponseBodyImages) Validate() error {
	if s.Image != nil {
		for _, item := range s.Image {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAvailableResourceInfoResponseBodyImagesImage struct {
	// The ID of the image.
	//
	// example:
	//
	// centos_6_08_64_20G_a****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The name of the image.
	//
	// example:
	//
	// centos_6_08_64_20G_a****
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// The size of the image. Unit: GB.
	//
	// example:
	//
	// 20
	ImageSize *int32 `json:"ImageSize,omitempty" xml:"ImageSize,omitempty"`
}

func (s DescribeAvailableResourceInfoResponseBodyImagesImage) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableResourceInfoResponseBodyImagesImage) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceInfoResponseBodyImagesImage) GetImageId() *string {
	return s.ImageId
}

func (s *DescribeAvailableResourceInfoResponseBodyImagesImage) GetImageName() *string {
	return s.ImageName
}

func (s *DescribeAvailableResourceInfoResponseBodyImagesImage) GetImageSize() *int32 {
	return s.ImageSize
}

func (s *DescribeAvailableResourceInfoResponseBodyImagesImage) SetImageId(v string) *DescribeAvailableResourceInfoResponseBodyImagesImage {
	s.ImageId = &v
	return s
}

func (s *DescribeAvailableResourceInfoResponseBodyImagesImage) SetImageName(v string) *DescribeAvailableResourceInfoResponseBodyImagesImage {
	s.ImageName = &v
	return s
}

func (s *DescribeAvailableResourceInfoResponseBodyImagesImage) SetImageSize(v int32) *DescribeAvailableResourceInfoResponseBodyImagesImage {
	s.ImageSize = &v
	return s
}

func (s *DescribeAvailableResourceInfoResponseBodyImagesImage) Validate() error {
	return dara.Validate(s)
}

type DescribeAvailableResourceInfoResponseBodySupportResources struct {
	SupportResource []*DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResource `json:"SupportResource,omitempty" xml:"SupportResource,omitempty" type:"Repeated"`
}

func (s DescribeAvailableResourceInfoResponseBodySupportResources) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableResourceInfoResponseBodySupportResources) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceInfoResponseBodySupportResources) GetSupportResource() []*DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResource {
	return s.SupportResource
}

func (s *DescribeAvailableResourceInfoResponseBodySupportResources) SetSupportResource(v []*DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResource) *DescribeAvailableResourceInfoResponseBodySupportResources {
	s.SupportResource = v
	return s
}

func (s *DescribeAvailableResourceInfoResponseBodySupportResources) Validate() error {
	if s.SupportResource != nil {
		for _, item := range s.SupportResource {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResource struct {
	// Bandwidth billing method.
	BandwidthTypes *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceBandwidthTypes `json:"BandwidthTypes,omitempty" xml:"BandwidthTypes,omitempty" type:"Struct"`
	// The maximum capacity of a data disk. Unit: GB.
	//
	// example:
	//
	// 200
	DataDiskMaxSize *int32 `json:"DataDiskMaxSize,omitempty" xml:"DataDiskMaxSize,omitempty"`
	// The minimum data disk size. Unit: GiB.
	//
	// example:
	//
	// 100
	DataDiskMinSize *int32 `json:"DataDiskMinSize,omitempty" xml:"DataDiskMinSize,omitempty"`
	// node ID
	EnsRegionIds *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceEnsRegionIds `json:"EnsRegionIds,omitempty" xml:"EnsRegionIds,omitempty" type:"Struct"`
	// The supplementary information about the edge nodes.
	EnsRegionIdsExtends *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceEnsRegionIdsExtends `json:"EnsRegionIdsExtends,omitempty" xml:"EnsRegionIdsExtends,omitempty" type:"Struct"`
	InstanceSpeces      *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceInstanceSpeces      `json:"InstanceSpeces,omitempty" xml:"InstanceSpeces,omitempty" type:"Struct"`
	// Operator
	Isp *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceIsp `json:"Isp,omitempty" xml:"Isp,omitempty" type:"Struct"`
	// The maximum size of the system disk. Unit: GiB.
	//
	// example:
	//
	// 100
	SystemDiskMaxSize *int32 `json:"SystemDiskMaxSize,omitempty" xml:"SystemDiskMaxSize,omitempty"`
	// The minimum capacity of a system disk. Unit: GB.
	//
	// example:
	//
	// 20
	SystemDiskMinSize *int32 `json:"SystemDiskMinSize,omitempty" xml:"SystemDiskMinSize,omitempty"`
}

func (s DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResource) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResource) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResource) GetBandwidthTypes() *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceBandwidthTypes {
	return s.BandwidthTypes
}

func (s *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResource) GetDataDiskMaxSize() *int32 {
	return s.DataDiskMaxSize
}

func (s *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResource) GetDataDiskMinSize() *int32 {
	return s.DataDiskMinSize
}

func (s *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResource) GetEnsRegionIds() *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceEnsRegionIds {
	return s.EnsRegionIds
}

func (s *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResource) GetEnsRegionIdsExtends() *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceEnsRegionIdsExtends {
	return s.EnsRegionIdsExtends
}

func (s *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResource) GetInstanceSpeces() *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceInstanceSpeces {
	return s.InstanceSpeces
}

func (s *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResource) GetIsp() *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceIsp {
	return s.Isp
}

func (s *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResource) GetSystemDiskMaxSize() *int32 {
	return s.SystemDiskMaxSize
}

func (s *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResource) GetSystemDiskMinSize() *int32 {
	return s.SystemDiskMinSize
}

func (s *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResource) SetBandwidthTypes(v *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceBandwidthTypes) *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResource {
	s.BandwidthTypes = v
	return s
}

func (s *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResource) SetDataDiskMaxSize(v int32) *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResource {
	s.DataDiskMaxSize = &v
	return s
}

func (s *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResource) SetDataDiskMinSize(v int32) *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResource {
	s.DataDiskMinSize = &v
	return s
}

func (s *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResource) SetEnsRegionIds(v *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceEnsRegionIds) *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResource {
	s.EnsRegionIds = v
	return s
}

func (s *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResource) SetEnsRegionIdsExtends(v *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceEnsRegionIdsExtends) *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResource {
	s.EnsRegionIdsExtends = v
	return s
}

func (s *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResource) SetInstanceSpeces(v *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceInstanceSpeces) *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResource {
	s.InstanceSpeces = v
	return s
}

func (s *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResource) SetIsp(v *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceIsp) *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResource {
	s.Isp = v
	return s
}

func (s *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResource) SetSystemDiskMaxSize(v int32) *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResource {
	s.SystemDiskMaxSize = &v
	return s
}

func (s *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResource) SetSystemDiskMinSize(v int32) *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResource {
	s.SystemDiskMinSize = &v
	return s
}

func (s *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResource) Validate() error {
	if s.BandwidthTypes != nil {
		if err := s.BandwidthTypes.Validate(); err != nil {
			return err
		}
	}
	if s.EnsRegionIds != nil {
		if err := s.EnsRegionIds.Validate(); err != nil {
			return err
		}
	}
	if s.EnsRegionIdsExtends != nil {
		if err := s.EnsRegionIdsExtends.Validate(); err != nil {
			return err
		}
	}
	if s.InstanceSpeces != nil {
		if err := s.InstanceSpeces.Validate(); err != nil {
			return err
		}
	}
	if s.Isp != nil {
		if err := s.Isp.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceBandwidthTypes struct {
	BandwidthType []*string `json:"BandwidthType,omitempty" xml:"BandwidthType,omitempty" type:"Repeated"`
}

func (s DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceBandwidthTypes) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceBandwidthTypes) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceBandwidthTypes) GetBandwidthType() []*string {
	return s.BandwidthType
}

func (s *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceBandwidthTypes) SetBandwidthType(v []*string) *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceBandwidthTypes {
	s.BandwidthType = v
	return s
}

func (s *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceBandwidthTypes) Validate() error {
	return dara.Validate(s)
}

type DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceEnsRegionIds struct {
	EnsRegionId []*string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty" type:"Repeated"`
}

func (s DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceEnsRegionIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceEnsRegionIds) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceEnsRegionIds) GetEnsRegionId() []*string {
	return s.EnsRegionId
}

func (s *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceEnsRegionIds) SetEnsRegionId(v []*string) *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceEnsRegionIds {
	s.EnsRegionId = v
	return s
}

func (s *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceEnsRegionIds) Validate() error {
	return dara.Validate(s)
}

type DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceEnsRegionIdsExtends struct {
	EnsRegionId []*DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceEnsRegionIdsExtendsEnsRegionId `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty" type:"Repeated"`
}

func (s DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceEnsRegionIdsExtends) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceEnsRegionIdsExtends) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceEnsRegionIdsExtends) GetEnsRegionId() []*DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceEnsRegionIdsExtendsEnsRegionId {
	return s.EnsRegionId
}

func (s *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceEnsRegionIdsExtends) SetEnsRegionId(v []*DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceEnsRegionIdsExtendsEnsRegionId) *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceEnsRegionIdsExtends {
	s.EnsRegionId = v
	return s
}

func (s *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceEnsRegionIdsExtends) Validate() error {
	if s.EnsRegionId != nil {
		for _, item := range s.EnsRegionId {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceEnsRegionIdsExtendsEnsRegionId struct {
	// The region.
	//
	// example:
	//
	// EastChina
	Area *string `json:"Area,omitempty" xml:"Area,omitempty"`
	// The name. This parameter is empty by default.
	//
	// example:
	//
	// EnName
	EnName *string `json:"EnName,omitempty" xml:"EnName,omitempty"`
	// The ID of the edge node.
	//
	// example:
	//
	// cn-chengdu-telecom-4
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// The information about the Internet service provider (ISP).
	//
	// example:
	//
	// unicom
	Isp *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
	// The name of the edge node.
	//
	// example:
	//
	// Taizhou Telecom, China Unicom, and China Mobile
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The province.
	//
	// example:
	//
	// Zhejiang Province
	Province *string `json:"Province,omitempty" xml:"Province,omitempty"`
}

func (s DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceEnsRegionIdsExtendsEnsRegionId) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceEnsRegionIdsExtendsEnsRegionId) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceEnsRegionIdsExtendsEnsRegionId) GetArea() *string {
	return s.Area
}

func (s *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceEnsRegionIdsExtendsEnsRegionId) GetEnName() *string {
	return s.EnName
}

func (s *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceEnsRegionIdsExtendsEnsRegionId) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceEnsRegionIdsExtendsEnsRegionId) GetIsp() *string {
	return s.Isp
}

func (s *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceEnsRegionIdsExtendsEnsRegionId) GetName() *string {
	return s.Name
}

func (s *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceEnsRegionIdsExtendsEnsRegionId) GetProvince() *string {
	return s.Province
}

func (s *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceEnsRegionIdsExtendsEnsRegionId) SetArea(v string) *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceEnsRegionIdsExtendsEnsRegionId {
	s.Area = &v
	return s
}

func (s *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceEnsRegionIdsExtendsEnsRegionId) SetEnName(v string) *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceEnsRegionIdsExtendsEnsRegionId {
	s.EnName = &v
	return s
}

func (s *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceEnsRegionIdsExtendsEnsRegionId) SetEnsRegionId(v string) *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceEnsRegionIdsExtendsEnsRegionId {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceEnsRegionIdsExtendsEnsRegionId) SetIsp(v string) *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceEnsRegionIdsExtendsEnsRegionId {
	s.Isp = &v
	return s
}

func (s *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceEnsRegionIdsExtendsEnsRegionId) SetName(v string) *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceEnsRegionIdsExtendsEnsRegionId {
	s.Name = &v
	return s
}

func (s *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceEnsRegionIdsExtendsEnsRegionId) SetProvince(v string) *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceEnsRegionIdsExtendsEnsRegionId {
	s.Province = &v
	return s
}

func (s *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceEnsRegionIdsExtendsEnsRegionId) Validate() error {
	return dara.Validate(s)
}

type DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceInstanceSpeces struct {
	InstanceSpec []*string `json:"InstanceSpec,omitempty" xml:"InstanceSpec,omitempty" type:"Repeated"`
}

func (s DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceInstanceSpeces) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceInstanceSpeces) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceInstanceSpeces) GetInstanceSpec() []*string {
	return s.InstanceSpec
}

func (s *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceInstanceSpeces) SetInstanceSpec(v []*string) *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceInstanceSpeces {
	s.InstanceSpec = v
	return s
}

func (s *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceInstanceSpeces) Validate() error {
	return dara.Validate(s)
}

type DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceIsp struct {
	Isp []*string `json:"Isp,omitempty" xml:"Isp,omitempty" type:"Repeated"`
}

func (s DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceIsp) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceIsp) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceIsp) GetIsp() []*string {
	return s.Isp
}

func (s *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceIsp) SetIsp(v []*string) *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceIsp {
	s.Isp = v
	return s
}

func (s *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceIsp) Validate() error {
	return dara.Validate(s)
}
