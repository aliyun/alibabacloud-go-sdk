// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTagRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *CreateTagRequest
	GetInstanceId() *string
	SetScriptId(v string) *CreateTagRequest
	GetScriptId() *string
	SetTagGroup(v string) *CreateTagRequest
	GetTagGroup() *string
	SetTagName(v string) *CreateTagRequest
	GetTagName() *string
}

type CreateTagRequest struct {
	// Instance ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 174952ab-9825-4cc9-a5e2-de82d7fa4cdd
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Scenario ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 0fe7f71c-8771-42ef-9bb1-19aa16ae7120
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// Tag group information
	//
	// This parameter is required.
	//
	// example:
	//
	// 目标学历
	TagGroup *string `json:"TagGroup,omitempty" xml:"TagGroup,omitempty"`
	// Tag name
	//
	// This parameter is required.
	//
	// example:
	//
	// 研究生
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
}

func (s CreateTagRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTagRequest) GoString() string {
	return s.String()
}

func (s *CreateTagRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateTagRequest) GetScriptId() *string {
	return s.ScriptId
}

func (s *CreateTagRequest) GetTagGroup() *string {
	return s.TagGroup
}

func (s *CreateTagRequest) GetTagName() *string {
	return s.TagName
}

func (s *CreateTagRequest) SetInstanceId(v string) *CreateTagRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateTagRequest) SetScriptId(v string) *CreateTagRequest {
	s.ScriptId = &v
	return s
}

func (s *CreateTagRequest) SetTagGroup(v string) *CreateTagRequest {
	s.TagGroup = &v
	return s
}

func (s *CreateTagRequest) SetTagName(v string) *CreateTagRequest {
	s.TagName = &v
	return s
}

func (s *CreateTagRequest) Validate() error {
	return dara.Validate(s)
}
