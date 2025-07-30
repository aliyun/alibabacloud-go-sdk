// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *ListRulesRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *ListRulesRequest
	GetJsonStr() *string
}

type ListRulesRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"pageNumber":1,"pageSize":10}
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s ListRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRulesRequest) GoString() string {
	return s.String()
}

func (s *ListRulesRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *ListRulesRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *ListRulesRequest) SetBaseMeAgentId(v int64) *ListRulesRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *ListRulesRequest) SetJsonStr(v string) *ListRulesRequest {
	s.JsonStr = &v
	return s
}

func (s *ListRulesRequest) Validate() error {
	return dara.Validate(s)
}
