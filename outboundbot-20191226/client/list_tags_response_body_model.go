// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTagsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListTagsResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListTagsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListTagsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListTagsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListTagsResponseBody
	GetSuccess() *bool
	SetTagGroups(v []*ListTagsResponseBodyTagGroups) *ListTagsResponseBody
	GetTagGroups() []*ListTagsResponseBodyTagGroups
	SetTags(v []*ListTagsResponseBodyTags) *ListTagsResponseBody
	GetTags() []*ListTagsResponseBodyTags
}

type ListTagsResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success   *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
	TagGroups []*ListTagsResponseBodyTagGroups `json:"TagGroups,omitempty" xml:"TagGroups,omitempty" type:"Repeated"`
	Tags      []*ListTagsResponseBodyTags      `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListTagsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTagsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListTagsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListTagsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListTagsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTagsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListTagsResponseBody) GetTagGroups() []*ListTagsResponseBodyTagGroups {
	return s.TagGroups
}

func (s *ListTagsResponseBody) GetTags() []*ListTagsResponseBodyTags {
	return s.Tags
}

func (s *ListTagsResponseBody) SetCode(v string) *ListTagsResponseBody {
	s.Code = &v
	return s
}

func (s *ListTagsResponseBody) SetHttpStatusCode(v int32) *ListTagsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListTagsResponseBody) SetMessage(v string) *ListTagsResponseBody {
	s.Message = &v
	return s
}

func (s *ListTagsResponseBody) SetRequestId(v string) *ListTagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagsResponseBody) SetSuccess(v bool) *ListTagsResponseBody {
	s.Success = &v
	return s
}

func (s *ListTagsResponseBody) SetTagGroups(v []*ListTagsResponseBodyTagGroups) *ListTagsResponseBody {
	s.TagGroups = v
	return s
}

func (s *ListTagsResponseBody) SetTags(v []*ListTagsResponseBodyTags) *ListTagsResponseBody {
	s.Tags = v
	return s
}

func (s *ListTagsResponseBody) Validate() error {
	if s.TagGroups != nil {
		for _, item := range s.TagGroups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
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

type ListTagsResponseBodyTagGroups struct {
	// example:
	//
	// 8a4c6d3d-5ed6-44ca-b779-16c20f8862be
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	TagGroup *string `json:"TagGroup,omitempty" xml:"TagGroup,omitempty"`
	// example:
	//
	// 38c03261-9fe8-4b9b-8c3b-983a60319012
	TagGroupId *string `json:"TagGroupId,omitempty" xml:"TagGroupId,omitempty"`
	// example:
	//
	// 1
	TagGroupIndex *int32 `json:"TagGroupIndex,omitempty" xml:"TagGroupIndex,omitempty"`
}

func (s ListTagsResponseBodyTagGroups) String() string {
	return dara.Prettify(s)
}

func (s ListTagsResponseBodyTagGroups) GoString() string {
	return s.String()
}

func (s *ListTagsResponseBodyTagGroups) GetScriptId() *string {
	return s.ScriptId
}

func (s *ListTagsResponseBodyTagGroups) GetTagGroup() *string {
	return s.TagGroup
}

func (s *ListTagsResponseBodyTagGroups) GetTagGroupId() *string {
	return s.TagGroupId
}

func (s *ListTagsResponseBodyTagGroups) GetTagGroupIndex() *int32 {
	return s.TagGroupIndex
}

func (s *ListTagsResponseBodyTagGroups) SetScriptId(v string) *ListTagsResponseBodyTagGroups {
	s.ScriptId = &v
	return s
}

func (s *ListTagsResponseBodyTagGroups) SetTagGroup(v string) *ListTagsResponseBodyTagGroups {
	s.TagGroup = &v
	return s
}

func (s *ListTagsResponseBodyTagGroups) SetTagGroupId(v string) *ListTagsResponseBodyTagGroups {
	s.TagGroupId = &v
	return s
}

func (s *ListTagsResponseBodyTagGroups) SetTagGroupIndex(v int32) *ListTagsResponseBodyTagGroups {
	s.TagGroupIndex = &v
	return s
}

func (s *ListTagsResponseBodyTagGroups) Validate() error {
	return dara.Validate(s)
}

type ListTagsResponseBodyTags struct {
	// example:
	//
	// 8a4c6d3d-5ed6-44ca-b779-16c20f8862be
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	TagGroup *string `json:"TagGroup,omitempty" xml:"TagGroup,omitempty"`
	// example:
	//
	// d62be647-6202-4b1f-9708-0baeec552635
	TagId *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
	// example:
	//
	// 1
	TagIndex *int32  `json:"TagIndex,omitempty" xml:"TagIndex,omitempty"`
	TagName  *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
}

func (s ListTagsResponseBodyTags) String() string {
	return dara.Prettify(s)
}

func (s ListTagsResponseBodyTags) GoString() string {
	return s.String()
}

func (s *ListTagsResponseBodyTags) GetScriptId() *string {
	return s.ScriptId
}

func (s *ListTagsResponseBodyTags) GetTagGroup() *string {
	return s.TagGroup
}

func (s *ListTagsResponseBodyTags) GetTagId() *string {
	return s.TagId
}

func (s *ListTagsResponseBodyTags) GetTagIndex() *int32 {
	return s.TagIndex
}

func (s *ListTagsResponseBodyTags) GetTagName() *string {
	return s.TagName
}

func (s *ListTagsResponseBodyTags) SetScriptId(v string) *ListTagsResponseBodyTags {
	s.ScriptId = &v
	return s
}

func (s *ListTagsResponseBodyTags) SetTagGroup(v string) *ListTagsResponseBodyTags {
	s.TagGroup = &v
	return s
}

func (s *ListTagsResponseBodyTags) SetTagId(v string) *ListTagsResponseBodyTags {
	s.TagId = &v
	return s
}

func (s *ListTagsResponseBodyTags) SetTagIndex(v int32) *ListTagsResponseBodyTags {
	s.TagIndex = &v
	return s
}

func (s *ListTagsResponseBodyTags) SetTagName(v string) *ListTagsResponseBodyTags {
	s.TagName = &v
	return s
}

func (s *ListTagsResponseBodyTags) Validate() error {
	return dara.Validate(s)
}
