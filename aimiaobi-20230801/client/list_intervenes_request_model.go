// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIntervenesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *ListIntervenesRequest
	GetAgentKey() *string
	SetInterveneType(v int32) *ListIntervenesRequest
	GetInterveneType() *int32
	SetPageIndex(v int32) *ListIntervenesRequest
	GetPageIndex() *int32
	SetPageSize(v int32) *ListIntervenesRequest
	GetPageSize() *int32
	SetQuery(v string) *ListIntervenesRequest
	GetQuery() *string
	SetRuleId(v int64) *ListIntervenesRequest
	GetRuleId() *int64
}

type ListIntervenesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 33a2658aaabf4c24b45d50e575125311_p_beebot_public
	AgentKey      *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	InterveneType *int32  `json:"InterveneType,omitempty" xml:"InterveneType,omitempty"`
	// example:
	//
	// 1
	PageIndex *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// example:
	//
	// 10
	PageSize *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Query    *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// example:
	//
	// mqtt_outRule_1679019634514
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s ListIntervenesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListIntervenesRequest) GoString() string {
	return s.String()
}

func (s *ListIntervenesRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *ListIntervenesRequest) GetInterveneType() *int32 {
	return s.InterveneType
}

func (s *ListIntervenesRequest) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *ListIntervenesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListIntervenesRequest) GetQuery() *string {
	return s.Query
}

func (s *ListIntervenesRequest) GetRuleId() *int64 {
	return s.RuleId
}

func (s *ListIntervenesRequest) SetAgentKey(v string) *ListIntervenesRequest {
	s.AgentKey = &v
	return s
}

func (s *ListIntervenesRequest) SetInterveneType(v int32) *ListIntervenesRequest {
	s.InterveneType = &v
	return s
}

func (s *ListIntervenesRequest) SetPageIndex(v int32) *ListIntervenesRequest {
	s.PageIndex = &v
	return s
}

func (s *ListIntervenesRequest) SetPageSize(v int32) *ListIntervenesRequest {
	s.PageSize = &v
	return s
}

func (s *ListIntervenesRequest) SetQuery(v string) *ListIntervenesRequest {
	s.Query = &v
	return s
}

func (s *ListIntervenesRequest) SetRuleId(v int64) *ListIntervenesRequest {
	s.RuleId = &v
	return s
}

func (s *ListIntervenesRequest) Validate() error {
	return dara.Validate(s)
}
