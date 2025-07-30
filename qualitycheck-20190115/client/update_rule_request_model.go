// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *UpdateRuleRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *UpdateRuleRequest
	GetJsonStr() *string
}

type UpdateRuleRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s UpdateRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateRuleRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *UpdateRuleRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *UpdateRuleRequest) SetBaseMeAgentId(v int64) *UpdateRuleRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *UpdateRuleRequest) SetJsonStr(v string) *UpdateRuleRequest {
	s.JsonStr = &v
	return s
}

func (s *UpdateRuleRequest) Validate() error {
	return dara.Validate(s)
}
