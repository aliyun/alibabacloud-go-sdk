// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSkillGroupConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *GetSkillGroupConfigRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *GetSkillGroupConfigRequest
	GetJsonStr() *string
}

type GetSkillGroupConfigRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s GetSkillGroupConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSkillGroupConfigRequest) GoString() string {
	return s.String()
}

func (s *GetSkillGroupConfigRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *GetSkillGroupConfigRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *GetSkillGroupConfigRequest) SetBaseMeAgentId(v int64) *GetSkillGroupConfigRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *GetSkillGroupConfigRequest) SetJsonStr(v string) *GetSkillGroupConfigRequest {
	s.JsonStr = &v
	return s
}

func (s *GetSkillGroupConfigRequest) Validate() error {
	return dara.Validate(s)
}
