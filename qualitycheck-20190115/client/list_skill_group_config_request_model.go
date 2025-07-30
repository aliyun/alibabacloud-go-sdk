// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSkillGroupConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *ListSkillGroupConfigRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *ListSkillGroupConfigRequest
	GetJsonStr() *string
}

type ListSkillGroupConfigRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"pageNumber":1,"pageSize": 1}
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s ListSkillGroupConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSkillGroupConfigRequest) GoString() string {
	return s.String()
}

func (s *ListSkillGroupConfigRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *ListSkillGroupConfigRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *ListSkillGroupConfigRequest) SetBaseMeAgentId(v int64) *ListSkillGroupConfigRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *ListSkillGroupConfigRequest) SetJsonStr(v string) *ListSkillGroupConfigRequest {
	s.JsonStr = &v
	return s
}

func (s *ListSkillGroupConfigRequest) Validate() error {
	return dara.Validate(s)
}
