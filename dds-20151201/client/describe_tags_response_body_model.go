// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTagsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *DescribeTagsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeTagsResponseBody
	GetRequestId() *string
	SetTags(v []*DescribeTagsResponseBodyTags) *DescribeTagsResponseBody
	GetTags() []*DescribeTagsResponseBodyTags
}

type DescribeTagsResponseBody struct {
	// The token used to start the next query.
	//
	// >  If not all results are returned in the first query, this parameter is returned. You can pass in the value of this parameter in the next query.
	//
	// example:
	//
	// 212db86****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// EEDBE38F-5CF5-4316-AAC2-35817BA60D68
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Details about the tags.
	Tags []*DescribeTagsResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s DescribeTagsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTagsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeTagsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTagsResponseBody) GetTags() []*DescribeTagsResponseBodyTags {
	return s.Tags
}

func (s *DescribeTagsResponseBody) SetNextToken(v string) *DescribeTagsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeTagsResponseBody) SetRequestId(v string) *DescribeTagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTagsResponseBody) SetTags(v []*DescribeTagsResponseBodyTags) *DescribeTagsResponseBody {
	s.Tags = v
	return s
}

func (s *DescribeTagsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeTagsResponseBodyTags struct {
	// The key of the tag.
	//
	// example:
	//
	// newKey
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The values of the tags.
	TagValues []*string `json:"TagValues,omitempty" xml:"TagValues,omitempty" type:"Repeated"`
}

func (s DescribeTagsResponseBodyTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagsResponseBodyTags) GoString() string {
	return s.String()
}

func (s *DescribeTagsResponseBodyTags) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeTagsResponseBodyTags) GetTagValues() []*string {
	return s.TagValues
}

func (s *DescribeTagsResponseBodyTags) SetTagKey(v string) *DescribeTagsResponseBodyTags {
	s.TagKey = &v
	return s
}

func (s *DescribeTagsResponseBodyTags) SetTagValues(v []*string) *DescribeTagsResponseBodyTags {
	s.TagValues = v
	return s
}

func (s *DescribeTagsResponseBodyTags) Validate() error {
	return dara.Validate(s)
}
