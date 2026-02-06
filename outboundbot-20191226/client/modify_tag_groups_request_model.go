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
	// This parameter is required.
	//
	// example:
	//
	// cd560e89-0459-4c8a-ad98-47d713e4abd6
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 365b955d-6f4d-4ab5-a6e1-9a301307f4b1
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// This parameter is required.
	TagGroups *string `json:"TagGroups,omitempty" xml:"TagGroups,omitempty"`
	// This parameter is required.
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
