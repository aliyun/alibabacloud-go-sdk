// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLumaNamespacesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentName(v string) *ListLumaNamespacesRequest
	GetAgentName() *string
	SetCatalog(v string) *ListLumaNamespacesRequest
	GetCatalog() *string
	SetLimit(v int32) *ListLumaNamespacesRequest
	GetLimit() *int32
	SetNextToken(v string) *ListLumaNamespacesRequest
	GetNextToken() *string
}

type ListLumaNamespacesRequest struct {
	// The name of the Agent.
	//
	// This parameter is required.
	//
	// example:
	//
	// my_agent
	AgentName *string `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
	// The name of the data catalog bound to the Agent. You can call ListLumaCatalogs to obtain the catalog name.
	//
	// This parameter is required.
	//
	// example:
	//
	// my_catalog
	Catalog *string `json:"Catalog,omitempty" xml:"Catalog,omitempty"`
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

func (s ListLumaNamespacesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListLumaNamespacesRequest) GoString() string {
	return s.String()
}

func (s *ListLumaNamespacesRequest) GetAgentName() *string {
	return s.AgentName
}

func (s *ListLumaNamespacesRequest) GetCatalog() *string {
	return s.Catalog
}

func (s *ListLumaNamespacesRequest) GetLimit() *int32 {
	return s.Limit
}

func (s *ListLumaNamespacesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListLumaNamespacesRequest) SetAgentName(v string) *ListLumaNamespacesRequest {
	s.AgentName = &v
	return s
}

func (s *ListLumaNamespacesRequest) SetCatalog(v string) *ListLumaNamespacesRequest {
	s.Catalog = &v
	return s
}

func (s *ListLumaNamespacesRequest) SetLimit(v int32) *ListLumaNamespacesRequest {
	s.Limit = &v
	return s
}

func (s *ListLumaNamespacesRequest) SetNextToken(v string) *ListLumaNamespacesRequest {
	s.NextToken = &v
	return s
}

func (s *ListLumaNamespacesRequest) Validate() error {
	return dara.Validate(s)
}
