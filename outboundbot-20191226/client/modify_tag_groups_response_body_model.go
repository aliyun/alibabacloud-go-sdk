// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyTagGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModifyTagGroupsResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ModifyTagGroupsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ModifyTagGroupsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyTagGroupsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyTagGroupsResponseBody
	GetSuccess() *bool
	SetTagGroups(v []*ModifyTagGroupsResponseBodyTagGroups) *ModifyTagGroupsResponseBody
	GetTagGroups() []*ModifyTagGroupsResponseBodyTagGroups
	SetTags(v []*ModifyTagGroupsResponseBodyTags) *ModifyTagGroupsResponseBody
	GetTags() []*ModifyTagGroupsResponseBodyTags
}

type ModifyTagGroupsResponseBody struct {
	// API status code
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// HTTP status code
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// API message
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID
	//
	// example:
	//
	// 5a865b03-d2b9-4ef9-be98-f21fa0d93744
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation succeeded.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// List of tag groups
	TagGroups []*ModifyTagGroupsResponseBodyTagGroups `json:"TagGroups,omitempty" xml:"TagGroups,omitempty" type:"Repeated"`
	// List of tags
	Tags []*ModifyTagGroupsResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ModifyTagGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyTagGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyTagGroupsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyTagGroupsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModifyTagGroupsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyTagGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyTagGroupsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyTagGroupsResponseBody) GetTagGroups() []*ModifyTagGroupsResponseBodyTagGroups {
	return s.TagGroups
}

func (s *ModifyTagGroupsResponseBody) GetTags() []*ModifyTagGroupsResponseBodyTags {
	return s.Tags
}

func (s *ModifyTagGroupsResponseBody) SetCode(v string) *ModifyTagGroupsResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyTagGroupsResponseBody) SetHttpStatusCode(v int32) *ModifyTagGroupsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyTagGroupsResponseBody) SetMessage(v string) *ModifyTagGroupsResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyTagGroupsResponseBody) SetRequestId(v string) *ModifyTagGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyTagGroupsResponseBody) SetSuccess(v bool) *ModifyTagGroupsResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyTagGroupsResponseBody) SetTagGroups(v []*ModifyTagGroupsResponseBodyTagGroups) *ModifyTagGroupsResponseBody {
	s.TagGroups = v
	return s
}

func (s *ModifyTagGroupsResponseBody) SetTags(v []*ModifyTagGroupsResponseBodyTags) *ModifyTagGroupsResponseBody {
	s.Tags = v
	return s
}

func (s *ModifyTagGroupsResponseBody) Validate() error {
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

type ModifyTagGroupsResponseBodyTagGroups struct {
	// Scenario ID
	//
	// example:
	//
	// 365b955d-6f4d-4ab5-a6e1-9a301307f4b1
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// Tag group name
	//
	// example:
	//
	// 标签组
	TagGroup *string `json:"TagGroup,omitempty" xml:"TagGroup,omitempty"`
	// Tag group ID
	//
	// example:
	//
	// 54629be9-0746-464a-ab59-4830242cf644
	TagGroupId *string `json:"TagGroupId,omitempty" xml:"TagGroupId,omitempty"`
	// Tag group index
	//
	// example:
	//
	// 1
	TagGroupIndex *int32 `json:"TagGroupIndex,omitempty" xml:"TagGroupIndex,omitempty"`
}

func (s ModifyTagGroupsResponseBodyTagGroups) String() string {
	return dara.Prettify(s)
}

func (s ModifyTagGroupsResponseBodyTagGroups) GoString() string {
	return s.String()
}

func (s *ModifyTagGroupsResponseBodyTagGroups) GetScriptId() *string {
	return s.ScriptId
}

func (s *ModifyTagGroupsResponseBodyTagGroups) GetTagGroup() *string {
	return s.TagGroup
}

func (s *ModifyTagGroupsResponseBodyTagGroups) GetTagGroupId() *string {
	return s.TagGroupId
}

func (s *ModifyTagGroupsResponseBodyTagGroups) GetTagGroupIndex() *int32 {
	return s.TagGroupIndex
}

func (s *ModifyTagGroupsResponseBodyTagGroups) SetScriptId(v string) *ModifyTagGroupsResponseBodyTagGroups {
	s.ScriptId = &v
	return s
}

func (s *ModifyTagGroupsResponseBodyTagGroups) SetTagGroup(v string) *ModifyTagGroupsResponseBodyTagGroups {
	s.TagGroup = &v
	return s
}

func (s *ModifyTagGroupsResponseBodyTagGroups) SetTagGroupId(v string) *ModifyTagGroupsResponseBodyTagGroups {
	s.TagGroupId = &v
	return s
}

func (s *ModifyTagGroupsResponseBodyTagGroups) SetTagGroupIndex(v int32) *ModifyTagGroupsResponseBodyTagGroups {
	s.TagGroupIndex = &v
	return s
}

func (s *ModifyTagGroupsResponseBodyTagGroups) Validate() error {
	return dara.Validate(s)
}

type ModifyTagGroupsResponseBodyTags struct {
	// Scenario ID
	//
	// example:
	//
	// 8a4c6d3d-5ed6-44ca-b779-16c20f8862be
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// Tag Group name
	//
	// example:
	//
	// 当前学历
	TagGroup *string `json:"TagGroup,omitempty" xml:"TagGroup,omitempty"`
	// Unique tag ID
	//
	// example:
	//
	// 19b23e92-4ee3-4129-8c2e-e1968670d887
	TagId *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
	// Tag index
	//
	// example:
	//
	// 1
	TagIndex *int32 `json:"TagIndex,omitempty" xml:"TagIndex,omitempty"`
	// Tag name
	//
	// example:
	//
	// 本科
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
}

func (s ModifyTagGroupsResponseBodyTags) String() string {
	return dara.Prettify(s)
}

func (s ModifyTagGroupsResponseBodyTags) GoString() string {
	return s.String()
}

func (s *ModifyTagGroupsResponseBodyTags) GetScriptId() *string {
	return s.ScriptId
}

func (s *ModifyTagGroupsResponseBodyTags) GetTagGroup() *string {
	return s.TagGroup
}

func (s *ModifyTagGroupsResponseBodyTags) GetTagId() *string {
	return s.TagId
}

func (s *ModifyTagGroupsResponseBodyTags) GetTagIndex() *int32 {
	return s.TagIndex
}

func (s *ModifyTagGroupsResponseBodyTags) GetTagName() *string {
	return s.TagName
}

func (s *ModifyTagGroupsResponseBodyTags) SetScriptId(v string) *ModifyTagGroupsResponseBodyTags {
	s.ScriptId = &v
	return s
}

func (s *ModifyTagGroupsResponseBodyTags) SetTagGroup(v string) *ModifyTagGroupsResponseBodyTags {
	s.TagGroup = &v
	return s
}

func (s *ModifyTagGroupsResponseBodyTags) SetTagId(v string) *ModifyTagGroupsResponseBodyTags {
	s.TagId = &v
	return s
}

func (s *ModifyTagGroupsResponseBodyTags) SetTagIndex(v int32) *ModifyTagGroupsResponseBodyTags {
	s.TagIndex = &v
	return s
}

func (s *ModifyTagGroupsResponseBodyTags) SetTagName(v string) *ModifyTagGroupsResponseBodyTags {
	s.TagName = &v
	return s
}

func (s *ModifyTagGroupsResponseBodyTags) Validate() error {
	return dara.Validate(s)
}
