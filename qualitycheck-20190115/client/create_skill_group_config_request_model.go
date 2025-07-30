// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSkillGroupConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *CreateSkillGroupConfigRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *CreateSkillGroupConfigRequest
	GetJsonStr() *string
}

type CreateSkillGroupConfigRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"skillGroupFrom":0,"qualityCheckType":0,"modelId":746,"name":"test","rid":"2493","vocabId":"267","skillGroupList":[{"skillGroupId":"0903","skillGroupName":"0903"}]}
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s CreateSkillGroupConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSkillGroupConfigRequest) GoString() string {
	return s.String()
}

func (s *CreateSkillGroupConfigRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *CreateSkillGroupConfigRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *CreateSkillGroupConfigRequest) SetBaseMeAgentId(v int64) *CreateSkillGroupConfigRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *CreateSkillGroupConfigRequest) SetJsonStr(v string) *CreateSkillGroupConfigRequest {
	s.JsonStr = &v
	return s
}

func (s *CreateSkillGroupConfigRequest) Validate() error {
	return dara.Validate(s)
}
