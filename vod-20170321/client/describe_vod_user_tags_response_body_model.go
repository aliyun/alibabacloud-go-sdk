// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodUserTagsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeVodUserTagsResponseBody
	GetRequestId() *string
	SetTags(v []*DescribeVodUserTagsResponseBodyTags) *DescribeVodUserTagsResponseBody
	GetTags() []*DescribeVodUserTagsResponseBodyTags
}

type DescribeVodUserTagsResponseBody struct {
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Tags      []*DescribeVodUserTagsResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s DescribeVodUserTagsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodUserTagsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVodUserTagsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVodUserTagsResponseBody) GetTags() []*DescribeVodUserTagsResponseBodyTags {
	return s.Tags
}

func (s *DescribeVodUserTagsResponseBody) SetRequestId(v string) *DescribeVodUserTagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVodUserTagsResponseBody) SetTags(v []*DescribeVodUserTagsResponseBodyTags) *DescribeVodUserTagsResponseBody {
	s.Tags = v
	return s
}

func (s *DescribeVodUserTagsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeVodUserTagsResponseBodyTags struct {
	Key   *string   `json:"Key,omitempty" xml:"Key,omitempty"`
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s DescribeVodUserTagsResponseBodyTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodUserTagsResponseBodyTags) GoString() string {
	return s.String()
}

func (s *DescribeVodUserTagsResponseBodyTags) GetKey() *string {
	return s.Key
}

func (s *DescribeVodUserTagsResponseBodyTags) GetValue() []*string {
	return s.Value
}

func (s *DescribeVodUserTagsResponseBodyTags) SetKey(v string) *DescribeVodUserTagsResponseBodyTags {
	s.Key = &v
	return s
}

func (s *DescribeVodUserTagsResponseBodyTags) SetValue(v []*string) *DescribeVodUserTagsResponseBodyTags {
	s.Value = v
	return s
}

func (s *DescribeVodUserTagsResponseBodyTags) Validate() error {
	return dara.Validate(s)
}
