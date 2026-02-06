// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSkillDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *GetSkillDetailRequest
	GetGroupId() *string
	SetSkillId(v string) *GetSkillDetailRequest
	GetSkillId() *string
	SetSourceIdOfAssistantId(v string) *GetSkillDetailRequest
	GetSourceIdOfAssistantId() *string
}

type GetSkillDetailRequest struct {
	// example:
	//
	// xxx
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// example:
	//
	// 8f6a2111-3828-4a9f-a3ce-51ce73c6ec9b
	SkillId *string `json:"SkillId,omitempty" xml:"SkillId,omitempty"`
	// example:
	//
	// xxx
	SourceIdOfAssistantId *string `json:"SourceIdOfAssistantId,omitempty" xml:"SourceIdOfAssistantId,omitempty"`
}

func (s GetSkillDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSkillDetailRequest) GoString() string {
	return s.String()
}

func (s *GetSkillDetailRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *GetSkillDetailRequest) GetSkillId() *string {
	return s.SkillId
}

func (s *GetSkillDetailRequest) GetSourceIdOfAssistantId() *string {
	return s.SourceIdOfAssistantId
}

func (s *GetSkillDetailRequest) SetGroupId(v string) *GetSkillDetailRequest {
	s.GroupId = &v
	return s
}

func (s *GetSkillDetailRequest) SetSkillId(v string) *GetSkillDetailRequest {
	s.SkillId = &v
	return s
}

func (s *GetSkillDetailRequest) SetSourceIdOfAssistantId(v string) *GetSkillDetailRequest {
	s.SourceIdOfAssistantId = &v
	return s
}

func (s *GetSkillDetailRequest) Validate() error {
	return dara.Validate(s)
}
