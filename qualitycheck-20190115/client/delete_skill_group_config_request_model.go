// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSkillGroupConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *DeleteSkillGroupConfigRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *DeleteSkillGroupConfigRequest
	GetJsonStr() *string
}

type DeleteSkillGroupConfigRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"id":552}
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s DeleteSkillGroupConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSkillGroupConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteSkillGroupConfigRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *DeleteSkillGroupConfigRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *DeleteSkillGroupConfigRequest) SetBaseMeAgentId(v int64) *DeleteSkillGroupConfigRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *DeleteSkillGroupConfigRequest) SetJsonStr(v string) *DeleteSkillGroupConfigRequest {
	s.JsonStr = &v
	return s
}

func (s *DeleteSkillGroupConfigRequest) Validate() error {
	return dara.Validate(s)
}
