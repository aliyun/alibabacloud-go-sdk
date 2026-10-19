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
	// > This parameter is returned if a single query does not return all results. You can pass the value of this parameter in a subsequent query to continue the query.
	//
	// example:
	//
	// 212db86sca4384811e0b5e8707ec2****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F3F44BE3-5419-4B61-9BAC-E66E295A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The tag information.
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

type DescribeTagsResponseBodyTags struct {
	// The tag key.
	//
	// example:
	//
	// size
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag values.
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
