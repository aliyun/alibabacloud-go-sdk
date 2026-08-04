// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPrivateAccessTagsForDynamicRouteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDynamicRoutes(v []*ListPrivateAccessTagsForDynamicRouteResponseBodyDynamicRoutes) *ListPrivateAccessTagsForDynamicRouteResponseBody
	GetDynamicRoutes() []*ListPrivateAccessTagsForDynamicRouteResponseBodyDynamicRoutes
	SetRequestId(v string) *ListPrivateAccessTagsForDynamicRouteResponseBody
	GetRequestId() *string
}

type ListPrivateAccessTagsForDynamicRouteResponseBody struct {
	// The dynamic route list.
	DynamicRoutes []*ListPrivateAccessTagsForDynamicRouteResponseBodyDynamicRoutes `json:"DynamicRoutes,omitempty" xml:"DynamicRoutes,omitempty" type:"Repeated"`
	// The ID of this request.
	//
	// example:
	//
	// B608C6AE-623D-55C4-9454-601B88AE937E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListPrivateAccessTagsForDynamicRouteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPrivateAccessTagsForDynamicRouteResponseBody) GoString() string {
	return s.String()
}

func (s *ListPrivateAccessTagsForDynamicRouteResponseBody) GetDynamicRoutes() []*ListPrivateAccessTagsForDynamicRouteResponseBodyDynamicRoutes {
	return s.DynamicRoutes
}

func (s *ListPrivateAccessTagsForDynamicRouteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPrivateAccessTagsForDynamicRouteResponseBody) SetDynamicRoutes(v []*ListPrivateAccessTagsForDynamicRouteResponseBodyDynamicRoutes) *ListPrivateAccessTagsForDynamicRouteResponseBody {
	s.DynamicRoutes = v
	return s
}

func (s *ListPrivateAccessTagsForDynamicRouteResponseBody) SetRequestId(v string) *ListPrivateAccessTagsForDynamicRouteResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPrivateAccessTagsForDynamicRouteResponseBody) Validate() error {
	if s.DynamicRoutes != nil {
		for _, item := range s.DynamicRoutes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPrivateAccessTagsForDynamicRouteResponseBodyDynamicRoutes struct {
	// The dynamic route ID.
	//
	// example:
	//
	// dr-ca9fddfac7c6****
	DynamicRouteId *string `json:"DynamicRouteId,omitempty" xml:"DynamicRouteId,omitempty"`
	// A collection of private network access tags.
	Tags []*ListPrivateAccessTagsForDynamicRouteResponseBodyDynamicRoutesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListPrivateAccessTagsForDynamicRouteResponseBodyDynamicRoutes) String() string {
	return dara.Prettify(s)
}

func (s ListPrivateAccessTagsForDynamicRouteResponseBodyDynamicRoutes) GoString() string {
	return s.String()
}

func (s *ListPrivateAccessTagsForDynamicRouteResponseBodyDynamicRoutes) GetDynamicRouteId() *string {
	return s.DynamicRouteId
}

func (s *ListPrivateAccessTagsForDynamicRouteResponseBodyDynamicRoutes) GetTags() []*ListPrivateAccessTagsForDynamicRouteResponseBodyDynamicRoutesTags {
	return s.Tags
}

func (s *ListPrivateAccessTagsForDynamicRouteResponseBodyDynamicRoutes) SetDynamicRouteId(v string) *ListPrivateAccessTagsForDynamicRouteResponseBodyDynamicRoutes {
	s.DynamicRouteId = &v
	return s
}

func (s *ListPrivateAccessTagsForDynamicRouteResponseBodyDynamicRoutes) SetTags(v []*ListPrivateAccessTagsForDynamicRouteResponseBodyDynamicRoutesTags) *ListPrivateAccessTagsForDynamicRouteResponseBodyDynamicRoutes {
	s.Tags = v
	return s
}

func (s *ListPrivateAccessTagsForDynamicRouteResponseBodyDynamicRoutes) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPrivateAccessTagsForDynamicRouteResponseBodyDynamicRoutesTags struct {
	// The private network access tag creation time.
	//
	// example:
	//
	// 2022-10-23 14:02:56
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The private network access tag description.
	//
	// example:
	//
	// 这是一条被动态路由引用的内网访问标签
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The private network access tag name.
	//
	// example:
	//
	// tag_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The private network access tag ID.
	//
	// example:
	//
	// tag-d3f64e8bdd4a****
	TagId *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
	// The private network access tag type. Valid values:
	//
	// - **Default**: Default.
	//
	// - **Custom**: Custom.
	//
	// example:
	//
	// Custom
	TagType *string `json:"TagType,omitempty" xml:"TagType,omitempty"`
}

func (s ListPrivateAccessTagsForDynamicRouteResponseBodyDynamicRoutesTags) String() string {
	return dara.Prettify(s)
}

func (s ListPrivateAccessTagsForDynamicRouteResponseBodyDynamicRoutesTags) GoString() string {
	return s.String()
}

func (s *ListPrivateAccessTagsForDynamicRouteResponseBodyDynamicRoutesTags) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListPrivateAccessTagsForDynamicRouteResponseBodyDynamicRoutesTags) GetDescription() *string {
	return s.Description
}

func (s *ListPrivateAccessTagsForDynamicRouteResponseBodyDynamicRoutesTags) GetName() *string {
	return s.Name
}

func (s *ListPrivateAccessTagsForDynamicRouteResponseBodyDynamicRoutesTags) GetTagId() *string {
	return s.TagId
}

func (s *ListPrivateAccessTagsForDynamicRouteResponseBodyDynamicRoutesTags) GetTagType() *string {
	return s.TagType
}

func (s *ListPrivateAccessTagsForDynamicRouteResponseBodyDynamicRoutesTags) SetCreateTime(v string) *ListPrivateAccessTagsForDynamicRouteResponseBodyDynamicRoutesTags {
	s.CreateTime = &v
	return s
}

func (s *ListPrivateAccessTagsForDynamicRouteResponseBodyDynamicRoutesTags) SetDescription(v string) *ListPrivateAccessTagsForDynamicRouteResponseBodyDynamicRoutesTags {
	s.Description = &v
	return s
}

func (s *ListPrivateAccessTagsForDynamicRouteResponseBodyDynamicRoutesTags) SetName(v string) *ListPrivateAccessTagsForDynamicRouteResponseBodyDynamicRoutesTags {
	s.Name = &v
	return s
}

func (s *ListPrivateAccessTagsForDynamicRouteResponseBodyDynamicRoutesTags) SetTagId(v string) *ListPrivateAccessTagsForDynamicRouteResponseBodyDynamicRoutesTags {
	s.TagId = &v
	return s
}

func (s *ListPrivateAccessTagsForDynamicRouteResponseBodyDynamicRoutesTags) SetTagType(v string) *ListPrivateAccessTagsForDynamicRouteResponseBodyDynamicRoutesTags {
	s.TagType = &v
	return s
}

func (s *ListPrivateAccessTagsForDynamicRouteResponseBodyDynamicRoutesTags) Validate() error {
	return dara.Validate(s)
}
