// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeReservedResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DescribeReservedResourceResponseBody
	GetCode() *int32
	SetImages(v *DescribeReservedResourceResponseBodyImages) *DescribeReservedResourceResponseBody
	GetImages() *DescribeReservedResourceResponseBodyImages
	SetRequestId(v string) *DescribeReservedResourceResponseBody
	GetRequestId() *string
	SetSupportResources(v *DescribeReservedResourceResponseBodySupportResources) *DescribeReservedResourceResponseBody
	GetSupportResources() *DescribeReservedResourceResponseBodySupportResources
}

type DescribeReservedResourceResponseBody struct {
	// The returned service code. 0 indicates that the request was successful.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information about the image.
	Images *DescribeReservedResourceResponseBodyImages `json:"Images,omitempty" xml:"Images,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 6666C5A5-75ED-422E-A022-7121FA18C968
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The resources.
	SupportResources *DescribeReservedResourceResponseBodySupportResources `json:"SupportResources,omitempty" xml:"SupportResources,omitempty" type:"Struct"`
}

func (s DescribeReservedResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeReservedResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeReservedResourceResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribeReservedResourceResponseBody) GetImages() *DescribeReservedResourceResponseBodyImages {
	return s.Images
}

func (s *DescribeReservedResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeReservedResourceResponseBody) GetSupportResources() *DescribeReservedResourceResponseBodySupportResources {
	return s.SupportResources
}

func (s *DescribeReservedResourceResponseBody) SetCode(v int32) *DescribeReservedResourceResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeReservedResourceResponseBody) SetImages(v *DescribeReservedResourceResponseBodyImages) *DescribeReservedResourceResponseBody {
	s.Images = v
	return s
}

func (s *DescribeReservedResourceResponseBody) SetRequestId(v string) *DescribeReservedResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeReservedResourceResponseBody) SetSupportResources(v *DescribeReservedResourceResponseBodySupportResources) *DescribeReservedResourceResponseBody {
	s.SupportResources = v
	return s
}

func (s *DescribeReservedResourceResponseBody) Validate() error {
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

type DescribeReservedResourceResponseBodyImages struct {
	Image []*DescribeReservedResourceResponseBodyImagesImage `json:"Image,omitempty" xml:"Image,omitempty" type:"Repeated"`
}

func (s DescribeReservedResourceResponseBodyImages) String() string {
	return dara.Prettify(s)
}

func (s DescribeReservedResourceResponseBodyImages) GoString() string {
	return s.String()
}

func (s *DescribeReservedResourceResponseBodyImages) GetImage() []*DescribeReservedResourceResponseBodyImagesImage {
	return s.Image
}

func (s *DescribeReservedResourceResponseBodyImages) SetImage(v []*DescribeReservedResourceResponseBodyImagesImage) *DescribeReservedResourceResponseBodyImages {
	s.Image = v
	return s
}

func (s *DescribeReservedResourceResponseBodyImages) Validate() error {
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

type DescribeReservedResourceResponseBodyImagesImage struct {
	// The ID of the image.
	//
	// example:
	//
	// centos_6_08_64_20G_alibase_****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The name of the image.
	//
	// example:
	//
	// centos_6_08_64_20G_alibase_****
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
}

func (s DescribeReservedResourceResponseBodyImagesImage) String() string {
	return dara.Prettify(s)
}

func (s DescribeReservedResourceResponseBodyImagesImage) GoString() string {
	return s.String()
}

func (s *DescribeReservedResourceResponseBodyImagesImage) GetImageId() *string {
	return s.ImageId
}

func (s *DescribeReservedResourceResponseBodyImagesImage) GetImageName() *string {
	return s.ImageName
}

func (s *DescribeReservedResourceResponseBodyImagesImage) SetImageId(v string) *DescribeReservedResourceResponseBodyImagesImage {
	s.ImageId = &v
	return s
}

func (s *DescribeReservedResourceResponseBodyImagesImage) SetImageName(v string) *DescribeReservedResourceResponseBodyImagesImage {
	s.ImageName = &v
	return s
}

func (s *DescribeReservedResourceResponseBodyImagesImage) Validate() error {
	return dara.Validate(s)
}

type DescribeReservedResourceResponseBodySupportResources struct {
	SupportResource []*DescribeReservedResourceResponseBodySupportResourcesSupportResource `json:"SupportResource,omitempty" xml:"SupportResource,omitempty" type:"Repeated"`
}

func (s DescribeReservedResourceResponseBodySupportResources) String() string {
	return dara.Prettify(s)
}

func (s DescribeReservedResourceResponseBodySupportResources) GoString() string {
	return s.String()
}

func (s *DescribeReservedResourceResponseBodySupportResources) GetSupportResource() []*DescribeReservedResourceResponseBodySupportResourcesSupportResource {
	return s.SupportResource
}

func (s *DescribeReservedResourceResponseBodySupportResources) SetSupportResource(v []*DescribeReservedResourceResponseBodySupportResourcesSupportResource) *DescribeReservedResourceResponseBodySupportResources {
	s.SupportResource = v
	return s
}

func (s *DescribeReservedResourceResponseBodySupportResources) Validate() error {
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

type DescribeReservedResourceResponseBodySupportResourcesSupportResource struct {
	// The sizes of data disks.
	DataDiskSizes *DescribeReservedResourceResponseBodySupportResourcesSupportResourceDataDiskSizes `json:"DataDiskSizes,omitempty" xml:"DataDiskSizes,omitempty" type:"Struct"`
	// The ID of the node.
	//
	// example:
	//
	// cn-beijing-cmcc
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// The specifications of instances.
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
	// The sizes of system disks.
	SystemDiskSizes *DescribeReservedResourceResponseBodySupportResourcesSupportResourceSystemDiskSizes `json:"SystemDiskSizes,omitempty" xml:"SystemDiskSizes,omitempty" type:"Struct"`
}

func (s DescribeReservedResourceResponseBodySupportResourcesSupportResource) String() string {
	return dara.Prettify(s)
}

func (s DescribeReservedResourceResponseBodySupportResourcesSupportResource) GoString() string {
	return s.String()
}

func (s *DescribeReservedResourceResponseBodySupportResourcesSupportResource) GetDataDiskSizes() *DescribeReservedResourceResponseBodySupportResourcesSupportResourceDataDiskSizes {
	return s.DataDiskSizes
}

func (s *DescribeReservedResourceResponseBodySupportResourcesSupportResource) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *DescribeReservedResourceResponseBodySupportResourcesSupportResource) GetInstanceSpec() *string {
	return s.InstanceSpec
}

func (s *DescribeReservedResourceResponseBodySupportResourcesSupportResource) GetSupportResourcesCount() *string {
	return s.SupportResourcesCount
}

func (s *DescribeReservedResourceResponseBodySupportResourcesSupportResource) GetSystemDiskSizes() *DescribeReservedResourceResponseBodySupportResourcesSupportResourceSystemDiskSizes {
	return s.SystemDiskSizes
}

func (s *DescribeReservedResourceResponseBodySupportResourcesSupportResource) SetDataDiskSizes(v *DescribeReservedResourceResponseBodySupportResourcesSupportResourceDataDiskSizes) *DescribeReservedResourceResponseBodySupportResourcesSupportResource {
	s.DataDiskSizes = v
	return s
}

func (s *DescribeReservedResourceResponseBodySupportResourcesSupportResource) SetEnsRegionId(v string) *DescribeReservedResourceResponseBodySupportResourcesSupportResource {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeReservedResourceResponseBodySupportResourcesSupportResource) SetInstanceSpec(v string) *DescribeReservedResourceResponseBodySupportResourcesSupportResource {
	s.InstanceSpec = &v
	return s
}

func (s *DescribeReservedResourceResponseBodySupportResourcesSupportResource) SetSupportResourcesCount(v string) *DescribeReservedResourceResponseBodySupportResourcesSupportResource {
	s.SupportResourcesCount = &v
	return s
}

func (s *DescribeReservedResourceResponseBodySupportResourcesSupportResource) SetSystemDiskSizes(v *DescribeReservedResourceResponseBodySupportResourcesSupportResourceSystemDiskSizes) *DescribeReservedResourceResponseBodySupportResourcesSupportResource {
	s.SystemDiskSizes = v
	return s
}

func (s *DescribeReservedResourceResponseBodySupportResourcesSupportResource) Validate() error {
	if s.DataDiskSizes != nil {
		if err := s.DataDiskSizes.Validate(); err != nil {
			return err
		}
	}
	if s.SystemDiskSizes != nil {
		if err := s.SystemDiskSizes.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeReservedResourceResponseBodySupportResourcesSupportResourceDataDiskSizes struct {
	DataDiskSize []*string `json:"DataDiskSize,omitempty" xml:"DataDiskSize,omitempty" type:"Repeated"`
}

func (s DescribeReservedResourceResponseBodySupportResourcesSupportResourceDataDiskSizes) String() string {
	return dara.Prettify(s)
}

func (s DescribeReservedResourceResponseBodySupportResourcesSupportResourceDataDiskSizes) GoString() string {
	return s.String()
}

func (s *DescribeReservedResourceResponseBodySupportResourcesSupportResourceDataDiskSizes) GetDataDiskSize() []*string {
	return s.DataDiskSize
}

func (s *DescribeReservedResourceResponseBodySupportResourcesSupportResourceDataDiskSizes) SetDataDiskSize(v []*string) *DescribeReservedResourceResponseBodySupportResourcesSupportResourceDataDiskSizes {
	s.DataDiskSize = v
	return s
}

func (s *DescribeReservedResourceResponseBodySupportResourcesSupportResourceDataDiskSizes) Validate() error {
	return dara.Validate(s)
}

type DescribeReservedResourceResponseBodySupportResourcesSupportResourceSystemDiskSizes struct {
	SystemDiskSize []*string `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty" type:"Repeated"`
}

func (s DescribeReservedResourceResponseBodySupportResourcesSupportResourceSystemDiskSizes) String() string {
	return dara.Prettify(s)
}

func (s DescribeReservedResourceResponseBodySupportResourcesSupportResourceSystemDiskSizes) GoString() string {
	return s.String()
}

func (s *DescribeReservedResourceResponseBodySupportResourcesSupportResourceSystemDiskSizes) GetSystemDiskSize() []*string {
	return s.SystemDiskSize
}

func (s *DescribeReservedResourceResponseBodySupportResourcesSupportResourceSystemDiskSizes) SetSystemDiskSize(v []*string) *DescribeReservedResourceResponseBodySupportResourcesSupportResourceSystemDiskSizes {
	s.SystemDiskSize = v
	return s
}

func (s *DescribeReservedResourceResponseBodySupportResourcesSupportResourceSystemDiskSizes) Validate() error {
	return dara.Validate(s)
}
