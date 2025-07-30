// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTaskAssignRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *ListTaskAssignRulesRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *ListTaskAssignRulesRequest
	GetJsonStr() *string
}

type ListTaskAssignRulesRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"pageNumber":1,"pageSize":10}
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s ListTaskAssignRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTaskAssignRulesRequest) GoString() string {
	return s.String()
}

func (s *ListTaskAssignRulesRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *ListTaskAssignRulesRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *ListTaskAssignRulesRequest) SetBaseMeAgentId(v int64) *ListTaskAssignRulesRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *ListTaskAssignRulesRequest) SetJsonStr(v string) *ListTaskAssignRulesRequest {
	s.JsonStr = &v
	return s
}

func (s *ListTaskAssignRulesRequest) Validate() error {
	return dara.Validate(s)
}
