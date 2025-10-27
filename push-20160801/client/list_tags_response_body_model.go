// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTagsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListTagsResponseBody
	GetRequestId() *string
	SetTagInfos(v *ListTagsResponseBodyTagInfos) *ListTagsResponseBody
	GetTagInfos() *ListTagsResponseBodyTagInfos
}

type ListTagsResponseBody struct {
	// example:
	//
	// 6EEF262B-EA7D-41DC-89B9-20F3D1E28194
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TagInfos  *ListTagsResponseBodyTagInfos `json:"TagInfos,omitempty" xml:"TagInfos,omitempty" type:"Struct"`
}

func (s ListTagsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTagsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTagsResponseBody) GetTagInfos() *ListTagsResponseBodyTagInfos {
	return s.TagInfos
}

func (s *ListTagsResponseBody) SetRequestId(v string) *ListTagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagsResponseBody) SetTagInfos(v *ListTagsResponseBodyTagInfos) *ListTagsResponseBody {
	s.TagInfos = v
	return s
}

func (s *ListTagsResponseBody) Validate() error {
	if s.TagInfos != nil {
		if err := s.TagInfos.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListTagsResponseBodyTagInfos struct {
	TagInfo []*ListTagsResponseBodyTagInfosTagInfo `json:"TagInfo,omitempty" xml:"TagInfo,omitempty" type:"Repeated"`
}

func (s ListTagsResponseBodyTagInfos) String() string {
	return dara.Prettify(s)
}

func (s ListTagsResponseBodyTagInfos) GoString() string {
	return s.String()
}

func (s *ListTagsResponseBodyTagInfos) GetTagInfo() []*ListTagsResponseBodyTagInfosTagInfo {
	return s.TagInfo
}

func (s *ListTagsResponseBodyTagInfos) SetTagInfo(v []*ListTagsResponseBodyTagInfosTagInfo) *ListTagsResponseBodyTagInfos {
	s.TagInfo = v
	return s
}

func (s *ListTagsResponseBodyTagInfos) Validate() error {
	if s.TagInfo != nil {
		for _, item := range s.TagInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTagsResponseBodyTagInfosTagInfo struct {
	// example:
	//
	// test_tag2
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
}

func (s ListTagsResponseBodyTagInfosTagInfo) String() string {
	return dara.Prettify(s)
}

func (s ListTagsResponseBodyTagInfosTagInfo) GoString() string {
	return s.String()
}

func (s *ListTagsResponseBodyTagInfosTagInfo) GetTagName() *string {
	return s.TagName
}

func (s *ListTagsResponseBodyTagInfosTagInfo) SetTagName(v string) *ListTagsResponseBodyTagInfosTagInfo {
	s.TagName = &v
	return s
}

func (s *ListTagsResponseBodyTagInfosTagInfo) Validate() error {
	return dara.Validate(s)
}
