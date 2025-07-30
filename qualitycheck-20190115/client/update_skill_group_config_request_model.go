// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSkillGroupConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *UpdateSkillGroupConfigRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *UpdateSkillGroupConfigRequest
	GetJsonStr() *string
}

type UpdateSkillGroupConfigRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"skillGroupFrom":0,"name":"test","qualityCheckType":0,"rid":"2493,4098","vocabId":267,"skillGroupList":[{"skillGroupId":"090311","skillGroupName":"09031"}],"id":553}
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s UpdateSkillGroupConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSkillGroupConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateSkillGroupConfigRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *UpdateSkillGroupConfigRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *UpdateSkillGroupConfigRequest) SetBaseMeAgentId(v int64) *UpdateSkillGroupConfigRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *UpdateSkillGroupConfigRequest) SetJsonStr(v string) *UpdateSkillGroupConfigRequest {
	s.JsonStr = &v
	return s
}

func (s *UpdateSkillGroupConfigRequest) Validate() error {
	return dara.Validate(s)
}
