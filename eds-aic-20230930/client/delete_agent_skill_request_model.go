// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAgentSkillRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSkillIds(v []*string) *DeleteAgentSkillRequest
	GetSkillIds() []*string
}

type DeleteAgentSkillRequest struct {
	// The list of skill IDs.
	SkillIds []*string `json:"SkillIds,omitempty" xml:"SkillIds,omitempty" type:"Repeated"`
}

func (s DeleteAgentSkillRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAgentSkillRequest) GoString() string {
	return s.String()
}

func (s *DeleteAgentSkillRequest) GetSkillIds() []*string {
	return s.SkillIds
}

func (s *DeleteAgentSkillRequest) SetSkillIds(v []*string) *DeleteAgentSkillRequest {
	s.SkillIds = v
	return s
}

func (s *DeleteAgentSkillRequest) Validate() error {
	return dara.Validate(s)
}
