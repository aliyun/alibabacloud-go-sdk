// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTagsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *QueryTagsResponseBody
	GetRequestId() *string
	SetTagInfos(v *QueryTagsResponseBodyTagInfos) *QueryTagsResponseBody
	GetTagInfos() *QueryTagsResponseBodyTagInfos
}

type QueryTagsResponseBody struct {
	// example:
	//
	// D68AE5C6-8AAF-46C9-B627-3FDACD1A4168
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TagInfos  *QueryTagsResponseBodyTagInfos `json:"TagInfos,omitempty" xml:"TagInfos,omitempty" type:"Struct"`
}

func (s QueryTagsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryTagsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTagsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryTagsResponseBody) GetTagInfos() *QueryTagsResponseBodyTagInfos {
	return s.TagInfos
}

func (s *QueryTagsResponseBody) SetRequestId(v string) *QueryTagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryTagsResponseBody) SetTagInfos(v *QueryTagsResponseBodyTagInfos) *QueryTagsResponseBody {
	s.TagInfos = v
	return s
}

func (s *QueryTagsResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryTagsResponseBodyTagInfos struct {
	TagInfo []*QueryTagsResponseBodyTagInfosTagInfo `json:"TagInfo,omitempty" xml:"TagInfo,omitempty" type:"Repeated"`
}

func (s QueryTagsResponseBodyTagInfos) String() string {
	return dara.Prettify(s)
}

func (s QueryTagsResponseBodyTagInfos) GoString() string {
	return s.String()
}

func (s *QueryTagsResponseBodyTagInfos) GetTagInfo() []*QueryTagsResponseBodyTagInfosTagInfo {
	return s.TagInfo
}

func (s *QueryTagsResponseBodyTagInfos) SetTagInfo(v []*QueryTagsResponseBodyTagInfosTagInfo) *QueryTagsResponseBodyTagInfos {
	s.TagInfo = v
	return s
}

func (s *QueryTagsResponseBodyTagInfos) Validate() error {
	return dara.Validate(s)
}

type QueryTagsResponseBodyTagInfosTagInfo struct {
	// example:
	//
	// test_tag2
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
}

func (s QueryTagsResponseBodyTagInfosTagInfo) String() string {
	return dara.Prettify(s)
}

func (s QueryTagsResponseBodyTagInfosTagInfo) GoString() string {
	return s.String()
}

func (s *QueryTagsResponseBodyTagInfosTagInfo) GetTagName() *string {
	return s.TagName
}

func (s *QueryTagsResponseBodyTagInfosTagInfo) SetTagName(v string) *QueryTagsResponseBodyTagInfosTagInfo {
	s.TagName = &v
	return s
}

func (s *QueryTagsResponseBodyTagInfosTagInfo) Validate() error {
	return dara.Validate(s)
}
