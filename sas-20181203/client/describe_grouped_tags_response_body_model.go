// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGroupedTagsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int32) *DescribeGroupedTagsResponseBody
	GetCount() *int32
	SetGroupedFileds(v []*DescribeGroupedTagsResponseBodyGroupedFileds) *DescribeGroupedTagsResponseBody
	GetGroupedFileds() []*DescribeGroupedTagsResponseBodyGroupedFileds
	SetHttpStatusCode(v int32) *DescribeGroupedTagsResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *DescribeGroupedTagsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeGroupedTagsResponseBody
	GetSuccess() *bool
}

type DescribeGroupedTagsResponseBody struct {
	// This parameter is deprecated.
	//
	// example:
	//
	// 0
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// An array that consists of the statistics of the asset tags.
	GroupedFileds []*DescribeGroupedTagsResponseBodyGroupedFileds `json:"GroupedFileds,omitempty" xml:"GroupedFileds,omitempty" type:"Repeated"`
	// The HTTP status code of the request.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 151F6EB6-D5F3-417A-AF7B-4D84975DB586
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- **true**: The request is successful.
	//
	// 	- **false**: The request fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeGroupedTagsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupedTagsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGroupedTagsResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *DescribeGroupedTagsResponseBody) GetGroupedFileds() []*DescribeGroupedTagsResponseBodyGroupedFileds {
	return s.GroupedFileds
}

func (s *DescribeGroupedTagsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeGroupedTagsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeGroupedTagsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeGroupedTagsResponseBody) SetCount(v int32) *DescribeGroupedTagsResponseBody {
	s.Count = &v
	return s
}

func (s *DescribeGroupedTagsResponseBody) SetGroupedFileds(v []*DescribeGroupedTagsResponseBodyGroupedFileds) *DescribeGroupedTagsResponseBody {
	s.GroupedFileds = v
	return s
}

func (s *DescribeGroupedTagsResponseBody) SetHttpStatusCode(v int32) *DescribeGroupedTagsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeGroupedTagsResponseBody) SetRequestId(v string) *DescribeGroupedTagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGroupedTagsResponseBody) SetSuccess(v bool) *DescribeGroupedTagsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeGroupedTagsResponseBody) Validate() error {
	if s.GroupedFileds != nil {
		for _, item := range s.GroupedFileds {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeGroupedTagsResponseBodyGroupedFileds struct {
	// The number of assets to which the tag is added.
	//
	// example:
	//
	// 152
	Count *string `json:"Count,omitempty" xml:"Count,omitempty"`
	// The name of the tag.
	//
	// example:
	//
	// InternetIp
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the tag.
	//
	// example:
	//
	// 3252366
	TagId *int32 `json:"TagId,omitempty" xml:"TagId,omitempty"`
}

func (s DescribeGroupedTagsResponseBodyGroupedFileds) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupedTagsResponseBodyGroupedFileds) GoString() string {
	return s.String()
}

func (s *DescribeGroupedTagsResponseBodyGroupedFileds) GetCount() *string {
	return s.Count
}

func (s *DescribeGroupedTagsResponseBodyGroupedFileds) GetName() *string {
	return s.Name
}

func (s *DescribeGroupedTagsResponseBodyGroupedFileds) GetTagId() *int32 {
	return s.TagId
}

func (s *DescribeGroupedTagsResponseBodyGroupedFileds) SetCount(v string) *DescribeGroupedTagsResponseBodyGroupedFileds {
	s.Count = &v
	return s
}

func (s *DescribeGroupedTagsResponseBodyGroupedFileds) SetName(v string) *DescribeGroupedTagsResponseBodyGroupedFileds {
	s.Name = &v
	return s
}

func (s *DescribeGroupedTagsResponseBodyGroupedFileds) SetTagId(v int32) *DescribeGroupedTagsResponseBodyGroupedFileds {
	s.TagId = &v
	return s
}

func (s *DescribeGroupedTagsResponseBodyGroupedFileds) Validate() error {
	return dara.Validate(s)
}
