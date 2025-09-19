// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeContainerTagsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeContainerTagsResponseBody
	GetRequestId() *string
	SetTagValues(v []*string) *DescribeContainerTagsResponseBody
	GetTagValues() []*string
}

type DescribeContainerTagsResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 028CF634-5268-5660-9575-48C9ED6BF880
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// An array that consists of the attributes of container assets.
	TagValues []*string `json:"TagValues,omitempty" xml:"TagValues,omitempty" type:"Repeated"`
}

func (s DescribeContainerTagsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeContainerTagsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeContainerTagsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeContainerTagsResponseBody) GetTagValues() []*string {
	return s.TagValues
}

func (s *DescribeContainerTagsResponseBody) SetRequestId(v string) *DescribeContainerTagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeContainerTagsResponseBody) SetTagValues(v []*string) *DescribeContainerTagsResponseBody {
	s.TagValues = v
	return s
}

func (s *DescribeContainerTagsResponseBody) Validate() error {
	return dara.Validate(s)
}
