// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveUserTagsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeLiveUserTagsResponseBody
	GetRequestId() *string
	SetTags(v []*DescribeLiveUserTagsResponseBodyTags) *DescribeLiveUserTagsResponseBody
	GetTags() []*DescribeLiveUserTagsResponseBodyTags
}

type DescribeLiveUserTagsResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 6B0256B6-2442-5BEF-B8D6-A0C28A801DFC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The tags.
	Tags []*DescribeLiveUserTagsResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s DescribeLiveUserTagsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveUserTagsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveUserTagsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveUserTagsResponseBody) GetTags() []*DescribeLiveUserTagsResponseBodyTags {
	return s.Tags
}

func (s *DescribeLiveUserTagsResponseBody) SetRequestId(v string) *DescribeLiveUserTagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveUserTagsResponseBody) SetTags(v []*DescribeLiveUserTagsResponseBodyTags) *DescribeLiveUserTagsResponseBody {
	s.Tags = v
	return s
}

func (s *DescribeLiveUserTagsResponseBody) Validate() error {
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

type DescribeLiveUserTagsResponseBodyTags struct {
	// The key of the tag.
	//
	// example:
	//
	// key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s DescribeLiveUserTagsResponseBodyTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveUserTagsResponseBodyTags) GoString() string {
	return s.String()
}

func (s *DescribeLiveUserTagsResponseBodyTags) GetKey() *string {
	return s.Key
}

func (s *DescribeLiveUserTagsResponseBodyTags) GetValue() []*string {
	return s.Value
}

func (s *DescribeLiveUserTagsResponseBodyTags) SetKey(v string) *DescribeLiveUserTagsResponseBodyTags {
	s.Key = &v
	return s
}

func (s *DescribeLiveUserTagsResponseBodyTags) SetValue(v []*string) *DescribeLiveUserTagsResponseBodyTags {
	s.Value = v
	return s
}

func (s *DescribeLiveUserTagsResponseBodyTags) Validate() error {
	return dara.Validate(s)
}
