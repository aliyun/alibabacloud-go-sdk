// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourcesForTagOptionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListResourcesForTagOptionResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListResourcesForTagOptionResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListResourcesForTagOptionResponseBody
	GetRequestId() *string
	SetResourceDetails(v []*ListResourcesForTagOptionResponseBodyResourceDetails) *ListResourcesForTagOptionResponseBody
	GetResourceDetails() []*ListResourcesForTagOptionResponseBodyResourceDetails
	SetTotalCount(v int32) *ListResourcesForTagOptionResponseBody
	GetTotalCount() *int32
}

type ListResourcesForTagOptionResponseBody struct {
	// The page number of the returned page.
	//
	// Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// Valid values: 1 to 100 Minimum value: 1. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// E465F21C-8712-5794-A754-5E4D9152****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// An array that consists of the associated resources.
	ResourceDetails []*ListResourcesForTagOptionResponseBodyResourceDetails `json:"ResourceDetails,omitempty" xml:"ResourceDetails,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListResourcesForTagOptionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListResourcesForTagOptionResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourcesForTagOptionResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListResourcesForTagOptionResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListResourcesForTagOptionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListResourcesForTagOptionResponseBody) GetResourceDetails() []*ListResourcesForTagOptionResponseBodyResourceDetails {
	return s.ResourceDetails
}

func (s *ListResourcesForTagOptionResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListResourcesForTagOptionResponseBody) SetPageNumber(v int32) *ListResourcesForTagOptionResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListResourcesForTagOptionResponseBody) SetPageSize(v int32) *ListResourcesForTagOptionResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListResourcesForTagOptionResponseBody) SetRequestId(v string) *ListResourcesForTagOptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListResourcesForTagOptionResponseBody) SetResourceDetails(v []*ListResourcesForTagOptionResponseBodyResourceDetails) *ListResourcesForTagOptionResponseBody {
	s.ResourceDetails = v
	return s
}

func (s *ListResourcesForTagOptionResponseBody) SetTotalCount(v int32) *ListResourcesForTagOptionResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListResourcesForTagOptionResponseBody) Validate() error {
	if s.ResourceDetails != nil {
		for _, item := range s.ResourceDetails {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListResourcesForTagOptionResponseBodyResourceDetails struct {
	// The time when the resource was created.
	//
	// The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-11-04T08:07:04.281986714Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the resource.
	//
	// The value must be 1 to 128 characters in length.
	//
	// example:
	//
	// The description of the resource.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the resource.
	//
	// example:
	//
	// acs:servicecatalog:cn-hangzhou:146611588617****:product/prod-bp18r7q127****
	ResourceArn *string `json:"ResourceArn,omitempty" xml:"ResourceArn,omitempty"`
	// The ID of the resource with which the tag option is associated.
	//
	// example:
	//
	// port-bp15p96922****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The name of the resource.
	//
	// example:
	//
	// DEMO-ECS
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
}

func (s ListResourcesForTagOptionResponseBodyResourceDetails) String() string {
	return dara.Prettify(s)
}

func (s ListResourcesForTagOptionResponseBodyResourceDetails) GoString() string {
	return s.String()
}

func (s *ListResourcesForTagOptionResponseBodyResourceDetails) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListResourcesForTagOptionResponseBodyResourceDetails) GetDescription() *string {
	return s.Description
}

func (s *ListResourcesForTagOptionResponseBodyResourceDetails) GetResourceArn() *string {
	return s.ResourceArn
}

func (s *ListResourcesForTagOptionResponseBodyResourceDetails) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListResourcesForTagOptionResponseBodyResourceDetails) GetResourceName() *string {
	return s.ResourceName
}

func (s *ListResourcesForTagOptionResponseBodyResourceDetails) SetCreateTime(v string) *ListResourcesForTagOptionResponseBodyResourceDetails {
	s.CreateTime = &v
	return s
}

func (s *ListResourcesForTagOptionResponseBodyResourceDetails) SetDescription(v string) *ListResourcesForTagOptionResponseBodyResourceDetails {
	s.Description = &v
	return s
}

func (s *ListResourcesForTagOptionResponseBodyResourceDetails) SetResourceArn(v string) *ListResourcesForTagOptionResponseBodyResourceDetails {
	s.ResourceArn = &v
	return s
}

func (s *ListResourcesForTagOptionResponseBodyResourceDetails) SetResourceId(v string) *ListResourcesForTagOptionResponseBodyResourceDetails {
	s.ResourceId = &v
	return s
}

func (s *ListResourcesForTagOptionResponseBodyResourceDetails) SetResourceName(v string) *ListResourcesForTagOptionResponseBodyResourceDetails {
	s.ResourceName = &v
	return s
}

func (s *ListResourcesForTagOptionResponseBodyResourceDetails) Validate() error {
	return dara.Validate(s)
}
