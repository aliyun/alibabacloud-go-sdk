// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudDiskAvailableResourceInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeCloudDiskAvailableResourceInfoResponseBody
	GetRequestId() *string
	SetSupportResources(v *DescribeCloudDiskAvailableResourceInfoResponseBodySupportResources) *DescribeCloudDiskAvailableResourceInfoResponseBody
	GetSupportResources() *DescribeCloudDiskAvailableResourceInfoResponseBodySupportResources
}

type DescribeCloudDiskAvailableResourceInfoResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0AE4F26E-7527-569F-A987-E3CF269A3C11
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The specifications of resources that you can purchase.
	SupportResources *DescribeCloudDiskAvailableResourceInfoResponseBodySupportResources `json:"SupportResources,omitempty" xml:"SupportResources,omitempty" type:"Struct"`
}

func (s DescribeCloudDiskAvailableResourceInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudDiskAvailableResourceInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCloudDiskAvailableResourceInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCloudDiskAvailableResourceInfoResponseBody) GetSupportResources() *DescribeCloudDiskAvailableResourceInfoResponseBodySupportResources {
	return s.SupportResources
}

func (s *DescribeCloudDiskAvailableResourceInfoResponseBody) SetRequestId(v string) *DescribeCloudDiskAvailableResourceInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCloudDiskAvailableResourceInfoResponseBody) SetSupportResources(v *DescribeCloudDiskAvailableResourceInfoResponseBodySupportResources) *DescribeCloudDiskAvailableResourceInfoResponseBody {
	s.SupportResources = v
	return s
}

func (s *DescribeCloudDiskAvailableResourceInfoResponseBody) Validate() error {
	if s.SupportResources != nil {
		if err := s.SupportResources.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCloudDiskAvailableResourceInfoResponseBodySupportResources struct {
	SupportResource []*DescribeCloudDiskAvailableResourceInfoResponseBodySupportResourcesSupportResource `json:"SupportResource,omitempty" xml:"SupportResource,omitempty" type:"Repeated"`
}

func (s DescribeCloudDiskAvailableResourceInfoResponseBodySupportResources) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudDiskAvailableResourceInfoResponseBodySupportResources) GoString() string {
	return s.String()
}

func (s *DescribeCloudDiskAvailableResourceInfoResponseBodySupportResources) GetSupportResource() []*DescribeCloudDiskAvailableResourceInfoResponseBodySupportResourcesSupportResource {
	return s.SupportResource
}

func (s *DescribeCloudDiskAvailableResourceInfoResponseBodySupportResources) SetSupportResource(v []*DescribeCloudDiskAvailableResourceInfoResponseBodySupportResourcesSupportResource) *DescribeCloudDiskAvailableResourceInfoResponseBodySupportResources {
	s.SupportResource = v
	return s
}

func (s *DescribeCloudDiskAvailableResourceInfoResponseBodySupportResources) Validate() error {
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

type DescribeCloudDiskAvailableResourceInfoResponseBodySupportResourcesSupportResource struct {
	// Node product capability.
	Ability *DescribeCloudDiskAvailableResourceInfoResponseBodySupportResourcesSupportResourceAbility `json:"Ability,omitempty" xml:"Ability,omitempty" type:"Struct"`
	// The number of disks that you can purchase.
	//
	// example:
	//
	// 2
	CanBuyCount *int64 `json:"CanBuyCount,omitempty" xml:"CanBuyCount,omitempty"`
	// The type of the disk.
	//
	// 	- cloud_efficiency:ultra disk.
	//
	// 	- cloud_ssd:all-flash disk.
	//
	// 	- local_hdd:local HDD.
	//
	// 	- local_ssd:local SSD.
	//
	// example:
	//
	// cloud_ssd
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The default size of the disk. Unit: GiB.
	//
	// example:
	//
	// 20
	DefaultDiskSize *int64 `json:"DefaultDiskSize,omitempty" xml:"DefaultDiskSize,omitempty"`
	// The maximum size of the disk. Unit: GiB.
	//
	// example:
	//
	// 80
	DiskMaxSize *int64 `json:"DiskMaxSize,omitempty" xml:"DiskMaxSize,omitempty"`
	// The minimum size of the disk size. Unit: GiB.
	//
	// example:
	//
	// 20
	DiskMinSize *int64 `json:"DiskMinSize,omitempty" xml:"DiskMinSize,omitempty"`
	// The ID of the edge node.
	//
	// example:
	//
	// cn-beijing-cmcc
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// The name of the task node.
	//
	// example:
	//
	// Beijing Mobile
	EnsRegionName *string `json:"EnsRegionName,omitempty" xml:"EnsRegionName,omitempty"`
}

func (s DescribeCloudDiskAvailableResourceInfoResponseBodySupportResourcesSupportResource) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudDiskAvailableResourceInfoResponseBodySupportResourcesSupportResource) GoString() string {
	return s.String()
}

func (s *DescribeCloudDiskAvailableResourceInfoResponseBodySupportResourcesSupportResource) GetAbility() *DescribeCloudDiskAvailableResourceInfoResponseBodySupportResourcesSupportResourceAbility {
	return s.Ability
}

func (s *DescribeCloudDiskAvailableResourceInfoResponseBodySupportResourcesSupportResource) GetCanBuyCount() *int64 {
	return s.CanBuyCount
}

func (s *DescribeCloudDiskAvailableResourceInfoResponseBodySupportResourcesSupportResource) GetCategory() *string {
	return s.Category
}

func (s *DescribeCloudDiskAvailableResourceInfoResponseBodySupportResourcesSupportResource) GetDefaultDiskSize() *int64 {
	return s.DefaultDiskSize
}

func (s *DescribeCloudDiskAvailableResourceInfoResponseBodySupportResourcesSupportResource) GetDiskMaxSize() *int64 {
	return s.DiskMaxSize
}

func (s *DescribeCloudDiskAvailableResourceInfoResponseBodySupportResourcesSupportResource) GetDiskMinSize() *int64 {
	return s.DiskMinSize
}

func (s *DescribeCloudDiskAvailableResourceInfoResponseBodySupportResourcesSupportResource) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *DescribeCloudDiskAvailableResourceInfoResponseBodySupportResourcesSupportResource) GetEnsRegionName() *string {
	return s.EnsRegionName
}

func (s *DescribeCloudDiskAvailableResourceInfoResponseBodySupportResourcesSupportResource) SetAbility(v *DescribeCloudDiskAvailableResourceInfoResponseBodySupportResourcesSupportResourceAbility) *DescribeCloudDiskAvailableResourceInfoResponseBodySupportResourcesSupportResource {
	s.Ability = v
	return s
}

func (s *DescribeCloudDiskAvailableResourceInfoResponseBodySupportResourcesSupportResource) SetCanBuyCount(v int64) *DescribeCloudDiskAvailableResourceInfoResponseBodySupportResourcesSupportResource {
	s.CanBuyCount = &v
	return s
}

func (s *DescribeCloudDiskAvailableResourceInfoResponseBodySupportResourcesSupportResource) SetCategory(v string) *DescribeCloudDiskAvailableResourceInfoResponseBodySupportResourcesSupportResource {
	s.Category = &v
	return s
}

func (s *DescribeCloudDiskAvailableResourceInfoResponseBodySupportResourcesSupportResource) SetDefaultDiskSize(v int64) *DescribeCloudDiskAvailableResourceInfoResponseBodySupportResourcesSupportResource {
	s.DefaultDiskSize = &v
	return s
}

func (s *DescribeCloudDiskAvailableResourceInfoResponseBodySupportResourcesSupportResource) SetDiskMaxSize(v int64) *DescribeCloudDiskAvailableResourceInfoResponseBodySupportResourcesSupportResource {
	s.DiskMaxSize = &v
	return s
}

func (s *DescribeCloudDiskAvailableResourceInfoResponseBodySupportResourcesSupportResource) SetDiskMinSize(v int64) *DescribeCloudDiskAvailableResourceInfoResponseBodySupportResourcesSupportResource {
	s.DiskMinSize = &v
	return s
}

func (s *DescribeCloudDiskAvailableResourceInfoResponseBodySupportResourcesSupportResource) SetEnsRegionId(v string) *DescribeCloudDiskAvailableResourceInfoResponseBodySupportResourcesSupportResource {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeCloudDiskAvailableResourceInfoResponseBodySupportResourcesSupportResource) SetEnsRegionName(v string) *DescribeCloudDiskAvailableResourceInfoResponseBodySupportResourcesSupportResource {
	s.EnsRegionName = &v
	return s
}

func (s *DescribeCloudDiskAvailableResourceInfoResponseBodySupportResourcesSupportResource) Validate() error {
	if s.Ability != nil {
		if err := s.Ability.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCloudDiskAvailableResourceInfoResponseBodySupportResourcesSupportResourceAbility struct {
	Ability []*string `json:"Ability,omitempty" xml:"Ability,omitempty" type:"Repeated"`
}

func (s DescribeCloudDiskAvailableResourceInfoResponseBodySupportResourcesSupportResourceAbility) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudDiskAvailableResourceInfoResponseBodySupportResourcesSupportResourceAbility) GoString() string {
	return s.String()
}

func (s *DescribeCloudDiskAvailableResourceInfoResponseBodySupportResourcesSupportResourceAbility) GetAbility() []*string {
	return s.Ability
}

func (s *DescribeCloudDiskAvailableResourceInfoResponseBodySupportResourcesSupportResourceAbility) SetAbility(v []*string) *DescribeCloudDiskAvailableResourceInfoResponseBodySupportResourcesSupportResourceAbility {
	s.Ability = v
	return s
}

func (s *DescribeCloudDiskAvailableResourceInfoResponseBodySupportResourcesSupportResourceAbility) Validate() error {
	return dara.Validate(s)
}
