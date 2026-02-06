// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSkillsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupIds(v []*string) *GetSkillsRequest
	GetGroupIds() []*string
	SetSkillIds(v []*string) *GetSkillsRequest
	GetSkillIds() []*string
	SetSourceIdOfAssistantId(v string) *GetSkillsRequest
	GetSourceIdOfAssistantId() *string
}

type GetSkillsRequest struct {
	GroupIds []*string `json:"GroupIds,omitempty" xml:"GroupIds,omitempty" type:"Repeated"`
	SkillIds []*string `json:"SkillIds,omitempty" xml:"SkillIds,omitempty" type:"Repeated"`
	// example:
	//
	// xxx
	SourceIdOfAssistantId *string `json:"SourceIdOfAssistantId,omitempty" xml:"SourceIdOfAssistantId,omitempty"`
}

func (s GetSkillsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSkillsRequest) GoString() string {
	return s.String()
}

func (s *GetSkillsRequest) GetGroupIds() []*string {
	return s.GroupIds
}

func (s *GetSkillsRequest) GetSkillIds() []*string {
	return s.SkillIds
}

func (s *GetSkillsRequest) GetSourceIdOfAssistantId() *string {
	return s.SourceIdOfAssistantId
}

func (s *GetSkillsRequest) SetGroupIds(v []*string) *GetSkillsRequest {
	s.GroupIds = v
	return s
}

func (s *GetSkillsRequest) SetSkillIds(v []*string) *GetSkillsRequest {
	s.SkillIds = v
	return s
}

func (s *GetSkillsRequest) SetSourceIdOfAssistantId(v string) *GetSkillsRequest {
	s.SourceIdOfAssistantId = &v
	return s
}

func (s *GetSkillsRequest) Validate() error {
	return dara.Validate(s)
}
