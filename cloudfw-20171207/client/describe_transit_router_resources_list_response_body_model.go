// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTransitRouterResourcesListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeTransitRouterResourcesListResponseBody
	GetRequestId() *string
	SetTransitRouterAttachedResources(v []*DescribeTransitRouterResourcesListResponseBodyTransitRouterAttachedResources) *DescribeTransitRouterResourcesListResponseBody
	GetTransitRouterAttachedResources() []*DescribeTransitRouterResourcesListResponseBodyTransitRouterAttachedResources
}

type DescribeTransitRouterResourcesListResponseBody struct {
	// example:
	//
	// A61A2516-0A22-5B3F-986B-3D4BF2A****
	RequestId                      *string                                                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TransitRouterAttachedResources []*DescribeTransitRouterResourcesListResponseBodyTransitRouterAttachedResources `json:"TransitRouterAttachedResources,omitempty" xml:"TransitRouterAttachedResources,omitempty" type:"Repeated"`
}

func (s DescribeTransitRouterResourcesListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTransitRouterResourcesListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTransitRouterResourcesListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTransitRouterResourcesListResponseBody) GetTransitRouterAttachedResources() []*DescribeTransitRouterResourcesListResponseBodyTransitRouterAttachedResources {
	return s.TransitRouterAttachedResources
}

func (s *DescribeTransitRouterResourcesListResponseBody) SetRequestId(v string) *DescribeTransitRouterResourcesListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTransitRouterResourcesListResponseBody) SetTransitRouterAttachedResources(v []*DescribeTransitRouterResourcesListResponseBodyTransitRouterAttachedResources) *DescribeTransitRouterResourcesListResponseBody {
	s.TransitRouterAttachedResources = v
	return s
}

func (s *DescribeTransitRouterResourcesListResponseBody) Validate() error {
	if s.TransitRouterAttachedResources != nil {
		for _, item := range s.TransitRouterAttachedResources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeTransitRouterResourcesListResponseBodyTransitRouterAttachedResources struct {
	// example:
	//
	// eas-r-8k1a6jjofkp0cq****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// example:
	//
	// test
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// example:
	//
	// VPC
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s DescribeTransitRouterResourcesListResponseBodyTransitRouterAttachedResources) String() string {
	return dara.Prettify(s)
}

func (s DescribeTransitRouterResourcesListResponseBodyTransitRouterAttachedResources) GoString() string {
	return s.String()
}

func (s *DescribeTransitRouterResourcesListResponseBodyTransitRouterAttachedResources) GetResourceId() *string {
	return s.ResourceId
}

func (s *DescribeTransitRouterResourcesListResponseBodyTransitRouterAttachedResources) GetResourceName() *string {
	return s.ResourceName
}

func (s *DescribeTransitRouterResourcesListResponseBodyTransitRouterAttachedResources) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeTransitRouterResourcesListResponseBodyTransitRouterAttachedResources) SetResourceId(v string) *DescribeTransitRouterResourcesListResponseBodyTransitRouterAttachedResources {
	s.ResourceId = &v
	return s
}

func (s *DescribeTransitRouterResourcesListResponseBodyTransitRouterAttachedResources) SetResourceName(v string) *DescribeTransitRouterResourcesListResponseBodyTransitRouterAttachedResources {
	s.ResourceName = &v
	return s
}

func (s *DescribeTransitRouterResourcesListResponseBodyTransitRouterAttachedResources) SetResourceType(v string) *DescribeTransitRouterResourcesListResponseBodyTransitRouterAttachedResources {
	s.ResourceType = &v
	return s
}

func (s *DescribeTransitRouterResourcesListResponseBodyTransitRouterAttachedResources) Validate() error {
	return dara.Validate(s)
}
