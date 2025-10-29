// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTagResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v *ListTagResourcesResponseBodyHeaders) *ListTagResourcesResponseBody
	GetHeaders() *ListTagResourcesResponseBodyHeaders
	SetPageSize(v int32) *ListTagResourcesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListTagResourcesResponseBody
	GetRequestId() *string
	SetTagResources(v *ListTagResourcesResponseBodyTagResources) *ListTagResourcesResponseBody
	GetTagResources() *ListTagResourcesResponseBodyTagResources
}

type ListTagResourcesResponseBody struct {
	// The labels of the resource.
	Headers *ListTagResourcesResponseBodyHeaders `json:"Headers,omitempty" xml:"Headers,omitempty" type:"Struct"`
	// The number of resources to query.
	//
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// A list of resources that have tags.
	//
	// example:
	//
	// F99407AB-2FA9-489E-A259-40CF6D******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The type of the resource. Fixed to `ALIYUN::ELASTICSEARCH::INSTANCE`.
	TagResources *ListTagResourcesResponseBodyTagResources `json:"TagResources,omitempty" xml:"TagResources,omitempty" type:"Struct"`
}

func (s ListTagResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBody) GetHeaders() *ListTagResourcesResponseBodyHeaders {
	return s.Headers
}

func (s *ListTagResourcesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTagResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTagResourcesResponseBody) GetTagResources() *ListTagResourcesResponseBodyTagResources {
	return s.TagResources
}

func (s *ListTagResourcesResponseBody) SetHeaders(v *ListTagResourcesResponseBodyHeaders) *ListTagResourcesResponseBody {
	s.Headers = v
	return s
}

func (s *ListTagResourcesResponseBody) SetPageSize(v int32) *ListTagResourcesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetRequestId(v string) *ListTagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetTagResources(v *ListTagResourcesResponseBodyTagResources) *ListTagResourcesResponseBody {
	s.TagResources = v
	return s
}

func (s *ListTagResourcesResponseBody) Validate() error {
	if s.Headers != nil {
		if err := s.Headers.Validate(); err != nil {
			return err
		}
	}
	if s.TagResources != nil {
		if err := s.TagResources.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListTagResourcesResponseBodyHeaders struct {
	// The value of the tag.
	//
	// example:
	//
	// 10
	XTotalCount *int32 `json:"X-Total-Count,omitempty" xml:"X-Total-Count,omitempty"`
}

func (s ListTagResourcesResponseBodyHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListTagResourcesResponseBodyHeaders) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyHeaders) GetXTotalCount() *int32 {
	return s.XTotalCount
}

func (s *ListTagResourcesResponseBodyHeaders) SetXTotalCount(v int32) *ListTagResourcesResponseBodyHeaders {
	s.XTotalCount = &v
	return s
}

func (s *ListTagResourcesResponseBodyHeaders) Validate() error {
	return dara.Validate(s)
}

type ListTagResourcesResponseBodyTagResources struct {
	// Indicates the ID of a resource.
	TagResource []*ListTagResourcesResponseBodyTagResourcesTagResource `json:"TagResource,omitempty" xml:"TagResource,omitempty" type:"Repeated"`
}

func (s ListTagResourcesResponseBodyTagResources) String() string {
	return dara.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResources) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResources) GetTagResource() []*ListTagResourcesResponseBodyTagResourcesTagResource {
	return s.TagResource
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagResource(v []*ListTagResourcesResponseBodyTagResourcesTagResource) *ListTagResourcesResponseBodyTagResources {
	s.TagResource = v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) Validate() error {
	if s.TagResource != nil {
		for _, item := range s.TagResource {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTagResourcesResponseBodyTagResourcesTagResource struct {
	// example:
	//
	// es-cn-oew1q8bev0002****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// example:
	//
	// ALIYUN::ELASTICSEARCH::INSTANCE
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// env
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag key.
	//
	// example:
	//
	// dev
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListTagResourcesResponseBodyTagResourcesTagResource) String() string {
	return dara.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResourcesTagResource) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) GetTagKey() *string {
	return s.TagKey
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) GetTagValue() *string {
	return s.TagValue
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetResourceId(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.ResourceId = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetResourceType(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetTagKey(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.TagKey = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetTagValue(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.TagValue = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) Validate() error {
	return dara.Validate(s)
}
