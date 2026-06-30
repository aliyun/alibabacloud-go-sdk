// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTaskAssignRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *UpdateTaskAssignRuleRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *UpdateTaskAssignRuleRequest
	GetJsonStr() *string
}

type UpdateTaskAssignRuleRequest struct {
	// The workspace ID.
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// A complete JSON string. For details, see the parameter descriptions below.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"callType":"0","agents":[{"agentId":"202526561358712105","agentName":"agent"}],"reviewers":[{"reviewerId":"255746168704895558","reviewerName":"0917质检员"},{"reviewerId":"268370362815185444","reviewerName":"0710质检员"}],"durationMin":1,"durationMax":300,"rules":[{"rid":15659},{"rid":17075}],"skillGroups":[{"skillName":"客服组"}],"enabled":1,"ruleId":37,"updateType":0}
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s UpdateTaskAssignRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskAssignRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateTaskAssignRuleRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *UpdateTaskAssignRuleRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *UpdateTaskAssignRuleRequest) SetBaseMeAgentId(v int64) *UpdateTaskAssignRuleRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *UpdateTaskAssignRuleRequest) SetJsonStr(v string) *UpdateTaskAssignRuleRequest {
	s.JsonStr = &v
	return s
}

func (s *UpdateTaskAssignRuleRequest) Validate() error {
	return dara.Validate(s)
}
