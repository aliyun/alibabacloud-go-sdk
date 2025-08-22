// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnUserTagsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeDcdnUserTagsResponseBody
	GetRequestId() *string
	SetTags(v []*DescribeDcdnUserTagsResponseBodyTags) *DescribeDcdnUserTagsResponseBody
	GetTags() []*DescribeDcdnUserTagsResponseBodyTags
}

type DescribeDcdnUserTagsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 34AB41F1-04A5-496F-8C8D-634BDBE6A9FB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The tag.
	Tags []*DescribeDcdnUserTagsResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s DescribeDcdnUserTagsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnUserTagsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserTagsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnUserTagsResponseBody) GetTags() []*DescribeDcdnUserTagsResponseBodyTags {
	return s.Tags
}

func (s *DescribeDcdnUserTagsResponseBody) SetRequestId(v string) *DescribeDcdnUserTagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnUserTagsResponseBody) SetTags(v []*DescribeDcdnUserTagsResponseBodyTags) *DescribeDcdnUserTagsResponseBody {
	s.Tags = v
	return s
}

func (s *DescribeDcdnUserTagsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnUserTagsResponseBodyTags struct {
	// The tag key.
	//
	// example:
	//
	// region
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag that you want to query.
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s DescribeDcdnUserTagsResponseBodyTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnUserTagsResponseBodyTags) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserTagsResponseBodyTags) GetKey() *string {
	return s.Key
}

func (s *DescribeDcdnUserTagsResponseBodyTags) GetValue() []*string {
	return s.Value
}

func (s *DescribeDcdnUserTagsResponseBodyTags) SetKey(v string) *DescribeDcdnUserTagsResponseBodyTags {
	s.Key = &v
	return s
}

func (s *DescribeDcdnUserTagsResponseBodyTags) SetValue(v []*string) *DescribeDcdnUserTagsResponseBodyTags {
	s.Value = v
	return s
}

func (s *DescribeDcdnUserTagsResponseBodyTags) Validate() error {
	return dara.Validate(s)
}
