// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyTagGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ModifyTagGroupsRequest
	GetInstanceId() *string
	SetScriptId(v string) *ModifyTagGroupsRequest
	GetScriptId() *string
	SetTagGroups(v string) *ModifyTagGroupsRequest
	GetTagGroups() *string
	SetTags(v string) *ModifyTagGroupsRequest
	GetTags() *string
}

type ModifyTagGroupsRequest struct {
	// Instance ID
	//
	// This parameter is required.
	//
	// example:
	//
	// cd560e89-0459-4c8a-ad98-47d713e4abd6
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Script ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 365b955d-6f4d-4ab5-a6e1-9a301307f4b1
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// Tag Group information
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"TagName":"111","TagGroup":"多层次","ScriptId":"b4d0dcc8-892d-4323-8c9d-3568e5faa62f","showInput":true,"Id":"a683fa32-91c5-457e-9ddf-aa8549d14ce0"}]
	TagGroups *string `json:"TagGroups,omitempty" xml:"TagGroups,omitempty"`
	// List of tags
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"TagGroup":"多层次","ScriptId":"b4d0dcc8-892d-43234-987c9d-3568e5faa62f","TagGroupIndex":0,"Id":"56728a30-c392-453a-a287-31af8301150f"}]
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s ModifyTagGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyTagGroupsRequest) GoString() string {
	return s.String()
}

func (s *ModifyTagGroupsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyTagGroupsRequest) GetScriptId() *string {
	return s.ScriptId
}

func (s *ModifyTagGroupsRequest) GetTagGroups() *string {
	return s.TagGroups
}

func (s *ModifyTagGroupsRequest) GetTags() *string {
	return s.Tags
}

func (s *ModifyTagGroupsRequest) SetInstanceId(v string) *ModifyTagGroupsRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyTagGroupsRequest) SetScriptId(v string) *ModifyTagGroupsRequest {
	s.ScriptId = &v
	return s
}

func (s *ModifyTagGroupsRequest) SetTagGroups(v string) *ModifyTagGroupsRequest {
	s.TagGroups = &v
	return s
}

func (s *ModifyTagGroupsRequest) SetTags(v string) *ModifyTagGroupsRequest {
	s.Tags = &v
	return s
}

func (s *ModifyTagGroupsRequest) Validate() error {
	return dara.Validate(s)
}
