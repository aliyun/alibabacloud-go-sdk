// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudDiskTypesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeCloudDiskTypesResponseBody
	GetRequestId() *string
	SetSupportResources(v *DescribeCloudDiskTypesResponseBodySupportResources) *DescribeCloudDiskTypesResponseBody
	GetSupportResources() *DescribeCloudDiskTypesResponseBodySupportResources
}

type DescribeCloudDiskTypesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 77990CEE-B714-5702-BDE6-943F702277DD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The specifications of resources that you can purchase.
	SupportResources *DescribeCloudDiskTypesResponseBodySupportResources `json:"SupportResources,omitempty" xml:"SupportResources,omitempty" type:"Struct"`
}

func (s DescribeCloudDiskTypesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudDiskTypesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCloudDiskTypesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCloudDiskTypesResponseBody) GetSupportResources() *DescribeCloudDiskTypesResponseBodySupportResources {
	return s.SupportResources
}

func (s *DescribeCloudDiskTypesResponseBody) SetRequestId(v string) *DescribeCloudDiskTypesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCloudDiskTypesResponseBody) SetSupportResources(v *DescribeCloudDiskTypesResponseBodySupportResources) *DescribeCloudDiskTypesResponseBody {
	s.SupportResources = v
	return s
}

func (s *DescribeCloudDiskTypesResponseBody) Validate() error {
	if s.SupportResources != nil {
		if err := s.SupportResources.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCloudDiskTypesResponseBodySupportResources struct {
	SupportResource []*DescribeCloudDiskTypesResponseBodySupportResourcesSupportResource `json:"SupportResource,omitempty" xml:"SupportResource,omitempty" type:"Repeated"`
}

func (s DescribeCloudDiskTypesResponseBodySupportResources) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudDiskTypesResponseBodySupportResources) GoString() string {
	return s.String()
}

func (s *DescribeCloudDiskTypesResponseBodySupportResources) GetSupportResource() []*DescribeCloudDiskTypesResponseBodySupportResourcesSupportResource {
	return s.SupportResource
}

func (s *DescribeCloudDiskTypesResponseBodySupportResources) SetSupportResource(v []*DescribeCloudDiskTypesResponseBodySupportResourcesSupportResource) *DescribeCloudDiskTypesResponseBodySupportResources {
	s.SupportResource = v
	return s
}

func (s *DescribeCloudDiskTypesResponseBodySupportResources) Validate() error {
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

type DescribeCloudDiskTypesResponseBodySupportResourcesSupportResource struct {
	// The category of the disk.
	//
	// 	- cloud_efficiency: ultra disk.
	//
	// 	- cloud_ssd: all-flash disk.
	//
	// 	- local_hdd: local HDD.
	//
	// 	- local_ssd: local SSD.
	//
	// example:
	//
	// cloud_efficiency
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The ID of the edge node.
	//
	// example:
	//
	// cn-guangzhou-10
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
}

func (s DescribeCloudDiskTypesResponseBodySupportResourcesSupportResource) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudDiskTypesResponseBodySupportResourcesSupportResource) GoString() string {
	return s.String()
}

func (s *DescribeCloudDiskTypesResponseBodySupportResourcesSupportResource) GetCategory() *string {
	return s.Category
}

func (s *DescribeCloudDiskTypesResponseBodySupportResourcesSupportResource) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *DescribeCloudDiskTypesResponseBodySupportResourcesSupportResource) SetCategory(v string) *DescribeCloudDiskTypesResponseBodySupportResourcesSupportResource {
	s.Category = &v
	return s
}

func (s *DescribeCloudDiskTypesResponseBodySupportResourcesSupportResource) SetEnsRegionId(v string) *DescribeCloudDiskTypesResponseBodySupportResourcesSupportResource {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeCloudDiskTypesResponseBodySupportResourcesSupportResource) Validate() error {
	return dara.Validate(s)
}
