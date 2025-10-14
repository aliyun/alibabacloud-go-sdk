// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAvailableResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DescribeAvailableResourceResponseBody
	GetCode() *int32
	SetImages(v *DescribeAvailableResourceResponseBodyImages) *DescribeAvailableResourceResponseBody
	GetImages() *DescribeAvailableResourceResponseBodyImages
	SetRequestId(v string) *DescribeAvailableResourceResponseBody
	GetRequestId() *string
	SetSupportResources(v *DescribeAvailableResourceResponseBodySupportResources) *DescribeAvailableResourceResponseBody
	GetSupportResources() *DescribeAvailableResourceResponseBodySupportResources
}

type DescribeAvailableResourceResponseBody struct {
	// The returned service code. 0 indicates that the request was successful.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The details of the images.
	Images *DescribeAvailableResourceResponseBodyImages `json:"Images,omitempty" xml:"Images,omitempty" type:"Struct"`
	// The ID of the request. This is a common parameter. Each request has a unique ID. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The specifications of resources that you can purchase.
	SupportResources *DescribeAvailableResourceResponseBodySupportResources `json:"SupportResources,omitempty" xml:"SupportResources,omitempty" type:"Struct"`
}

func (s DescribeAvailableResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribeAvailableResourceResponseBody) GetImages() *DescribeAvailableResourceResponseBodyImages {
	return s.Images
}

func (s *DescribeAvailableResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAvailableResourceResponseBody) GetSupportResources() *DescribeAvailableResourceResponseBodySupportResources {
	return s.SupportResources
}

func (s *DescribeAvailableResourceResponseBody) SetCode(v int32) *DescribeAvailableResourceResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAvailableResourceResponseBody) SetImages(v *DescribeAvailableResourceResponseBodyImages) *DescribeAvailableResourceResponseBody {
	s.Images = v
	return s
}

func (s *DescribeAvailableResourceResponseBody) SetRequestId(v string) *DescribeAvailableResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAvailableResourceResponseBody) SetSupportResources(v *DescribeAvailableResourceResponseBodySupportResources) *DescribeAvailableResourceResponseBody {
	s.SupportResources = v
	return s
}

func (s *DescribeAvailableResourceResponseBody) Validate() error {
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

type DescribeAvailableResourceResponseBodyImages struct {
	Image []*DescribeAvailableResourceResponseBodyImagesImage `json:"Image,omitempty" xml:"Image,omitempty" type:"Repeated"`
}

func (s DescribeAvailableResourceResponseBodyImages) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyImages) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyImages) GetImage() []*DescribeAvailableResourceResponseBodyImagesImage {
	return s.Image
}

func (s *DescribeAvailableResourceResponseBodyImages) SetImage(v []*DescribeAvailableResourceResponseBodyImagesImage) *DescribeAvailableResourceResponseBodyImages {
	s.Image = v
	return s
}

func (s *DescribeAvailableResourceResponseBodyImages) Validate() error {
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

type DescribeAvailableResourceResponseBodyImagesImage struct {
	// The ID of the image.
	//
	// example:
	//
	// centos_6_08_64_20G_alibase_20171208
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The name of the image.
	//
	// example:
	//
	// centos_6_08_64_20G_alibase_****
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
}

func (s DescribeAvailableResourceResponseBodyImagesImage) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyImagesImage) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyImagesImage) GetImageId() *string {
	return s.ImageId
}

func (s *DescribeAvailableResourceResponseBodyImagesImage) GetImageName() *string {
	return s.ImageName
}

func (s *DescribeAvailableResourceResponseBodyImagesImage) SetImageId(v string) *DescribeAvailableResourceResponseBodyImagesImage {
	s.ImageId = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodyImagesImage) SetImageName(v string) *DescribeAvailableResourceResponseBodyImagesImage {
	s.ImageName = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodyImagesImage) Validate() error {
	return dara.Validate(s)
}

type DescribeAvailableResourceResponseBodySupportResources struct {
	SupportResource []*DescribeAvailableResourceResponseBodySupportResourcesSupportResource `json:"SupportResource,omitempty" xml:"SupportResource,omitempty" type:"Repeated"`
}

func (s DescribeAvailableResourceResponseBodySupportResources) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodySupportResources) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodySupportResources) GetSupportResource() []*DescribeAvailableResourceResponseBodySupportResourcesSupportResource {
	return s.SupportResource
}

func (s *DescribeAvailableResourceResponseBodySupportResources) SetSupportResource(v []*DescribeAvailableResourceResponseBodySupportResourcesSupportResource) *DescribeAvailableResourceResponseBodySupportResources {
	s.SupportResource = v
	return s
}

func (s *DescribeAvailableResourceResponseBodySupportResources) Validate() error {
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

type DescribeAvailableResourceResponseBodySupportResourcesSupportResource struct {
	// The size of the data disk. Unit: GB.
	//
	// example:
	//
	// 500
	DataDiskSize *string `json:"DataDiskSize,omitempty" xml:"DataDiskSize,omitempty"`
	// The ID of the edge node.
	//
	// example:
	//
	// cn-beijing-cmcc
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// The specifications of the resource plan.
	//
	// example:
	//
	// ens.sn1.stiny
	InstanceSpec *string `json:"InstanceSpec,omitempty" xml:"InstanceSpec,omitempty"`
	// The number of resources that you can purchase.
	//
	// example:
	//
	// 9
	SupportResourcesCount *string `json:"SupportResourcesCount,omitempty" xml:"SupportResourcesCount,omitempty"`
	// The size of the system disk. Unit: GiB.
	//
	// example:
	//
	// 20
	SystemDiskSize *string `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
}

func (s DescribeAvailableResourceResponseBodySupportResourcesSupportResource) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodySupportResourcesSupportResource) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodySupportResourcesSupportResource) GetDataDiskSize() *string {
	return s.DataDiskSize
}

func (s *DescribeAvailableResourceResponseBodySupportResourcesSupportResource) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *DescribeAvailableResourceResponseBodySupportResourcesSupportResource) GetInstanceSpec() *string {
	return s.InstanceSpec
}

func (s *DescribeAvailableResourceResponseBodySupportResourcesSupportResource) GetSupportResourcesCount() *string {
	return s.SupportResourcesCount
}

func (s *DescribeAvailableResourceResponseBodySupportResourcesSupportResource) GetSystemDiskSize() *string {
	return s.SystemDiskSize
}

func (s *DescribeAvailableResourceResponseBodySupportResourcesSupportResource) SetDataDiskSize(v string) *DescribeAvailableResourceResponseBodySupportResourcesSupportResource {
	s.DataDiskSize = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodySupportResourcesSupportResource) SetEnsRegionId(v string) *DescribeAvailableResourceResponseBodySupportResourcesSupportResource {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodySupportResourcesSupportResource) SetInstanceSpec(v string) *DescribeAvailableResourceResponseBodySupportResourcesSupportResource {
	s.InstanceSpec = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodySupportResourcesSupportResource) SetSupportResourcesCount(v string) *DescribeAvailableResourceResponseBodySupportResourcesSupportResource {
	s.SupportResourcesCount = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodySupportResourcesSupportResource) SetSystemDiskSize(v string) *DescribeAvailableResourceResponseBodySupportResourcesSupportResource {
	s.SystemDiskSize = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodySupportResourcesSupportResource) Validate() error {
	return dara.Validate(s)
}
