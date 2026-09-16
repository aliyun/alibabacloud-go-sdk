// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLumaCatalogsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentName(v string) *ListLumaCatalogsRequest
	GetAgentName() *string
	SetLimit(v int32) *ListLumaCatalogsRequest
	GetLimit() *int32
	SetNextToken(v string) *ListLumaCatalogsRequest
	GetNextToken() *string
}

type ListLumaCatalogsRequest struct {
	// The name of the Agent.
	//
	// This parameter is required.
	//
	// example:
	//
	// my_agent
	AgentName *string `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
	// 每页返回的最大数据条数。取值范围 1~100，不传时默认 100。每条记录都需回源查询一次元数据，因此该值同时限制单次调用的回源次数
	//
	// example:
	//
	// 10
	Limit *int32 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// 分页查询的起始Token。首次查询不传或传 "0"；后续翻页使用上一次响应中返回的 NextToken 值
	//
	// example:
	//
	// 0
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListLumaCatalogsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListLumaCatalogsRequest) GoString() string {
	return s.String()
}

func (s *ListLumaCatalogsRequest) GetAgentName() *string {
	return s.AgentName
}

func (s *ListLumaCatalogsRequest) GetLimit() *int32 {
	return s.Limit
}

func (s *ListLumaCatalogsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListLumaCatalogsRequest) SetAgentName(v string) *ListLumaCatalogsRequest {
	s.AgentName = &v
	return s
}

func (s *ListLumaCatalogsRequest) SetLimit(v int32) *ListLumaCatalogsRequest {
	s.Limit = &v
	return s
}

func (s *ListLumaCatalogsRequest) SetNextToken(v string) *ListLumaCatalogsRequest {
	s.NextToken = &v
	return s
}

func (s *ListLumaCatalogsRequest) Validate() error {
	return dara.Validate(s)
}
