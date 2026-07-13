// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAtiAgentRegisterInfosRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentDisplayName(v string) *ListAtiAgentRegisterInfosRequest
	GetAgentDisplayName() *string
	SetAgentHost(v string) *ListAtiAgentRegisterInfosRequest
	GetAgentHost() *string
	SetAgentId(v string) *ListAtiAgentRegisterInfosRequest
	GetAgentId() *string
	SetAgentVersion(v string) *ListAtiAgentRegisterInfosRequest
	GetAgentVersion() *string
	SetClientToken(v string) *ListAtiAgentRegisterInfosRequest
	GetClientToken() *string
	SetMaxResults(v int32) *ListAtiAgentRegisterInfosRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListAtiAgentRegisterInfosRequest
	GetNextToken() *string
	SetPageNumber(v int32) *ListAtiAgentRegisterInfosRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListAtiAgentRegisterInfosRequest
	GetPageSize() *int32
	SetStatus(v string) *ListAtiAgentRegisterInfosRequest
	GetStatus() *string
}

type ListAtiAgentRegisterInfosRequest struct {
	// The display name of the Agent.
	//
	// example:
	//
	// 测试Agent
	AgentDisplayName *string `json:"AgentDisplayName,omitempty" xml:"AgentDisplayName,omitempty"`
	// The host address of the Agent.
	//
	// example:
	//
	// www.example.com
	AgentHost *string `json:"AgentHost,omitempty" xml:"AgentHost,omitempty"`
	// The Agent ID, which is uniformly assigned by CNNIC after real-name verification through CNNIC. The AgentID serves as the unique identifier that binds the Agent to the real-name verified registrant.
	//
	// example:
	//
	// liuq@azt400
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// The version of the Agent.
	//
	// example:
	//
	// 1.0.0
	AgentVersion *string `json:"AgentVersion,omitempty" xml:"AgentVersion,omitempty"`
	// Ensures the idempotency of the request. Generate a unique parameter value from your client to ensure that the value is unique across different requests. ClientToken supports only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// eyJhbGciOiJIUzI1NiIsInR5cC.....
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The maximum number of entries to return in this request.
	//
	// example:
	//
	// 100
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token for the next query.
	//
	// example:
	//
	// 4698691
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The current page number. Minimum value: 1. Default value: 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The page size for the paged query. This parameter specifies the number of entries per page for paging.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The status of the Agent.
	//
	// example:
	//
	// 活跃
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListAtiAgentRegisterInfosRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAtiAgentRegisterInfosRequest) GoString() string {
	return s.String()
}

func (s *ListAtiAgentRegisterInfosRequest) GetAgentDisplayName() *string {
	return s.AgentDisplayName
}

func (s *ListAtiAgentRegisterInfosRequest) GetAgentHost() *string {
	return s.AgentHost
}

func (s *ListAtiAgentRegisterInfosRequest) GetAgentId() *string {
	return s.AgentId
}

func (s *ListAtiAgentRegisterInfosRequest) GetAgentVersion() *string {
	return s.AgentVersion
}

func (s *ListAtiAgentRegisterInfosRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ListAtiAgentRegisterInfosRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAtiAgentRegisterInfosRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAtiAgentRegisterInfosRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAtiAgentRegisterInfosRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAtiAgentRegisterInfosRequest) GetStatus() *string {
	return s.Status
}

func (s *ListAtiAgentRegisterInfosRequest) SetAgentDisplayName(v string) *ListAtiAgentRegisterInfosRequest {
	s.AgentDisplayName = &v
	return s
}

func (s *ListAtiAgentRegisterInfosRequest) SetAgentHost(v string) *ListAtiAgentRegisterInfosRequest {
	s.AgentHost = &v
	return s
}

func (s *ListAtiAgentRegisterInfosRequest) SetAgentId(v string) *ListAtiAgentRegisterInfosRequest {
	s.AgentId = &v
	return s
}

func (s *ListAtiAgentRegisterInfosRequest) SetAgentVersion(v string) *ListAtiAgentRegisterInfosRequest {
	s.AgentVersion = &v
	return s
}

func (s *ListAtiAgentRegisterInfosRequest) SetClientToken(v string) *ListAtiAgentRegisterInfosRequest {
	s.ClientToken = &v
	return s
}

func (s *ListAtiAgentRegisterInfosRequest) SetMaxResults(v int32) *ListAtiAgentRegisterInfosRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAtiAgentRegisterInfosRequest) SetNextToken(v string) *ListAtiAgentRegisterInfosRequest {
	s.NextToken = &v
	return s
}

func (s *ListAtiAgentRegisterInfosRequest) SetPageNumber(v int32) *ListAtiAgentRegisterInfosRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAtiAgentRegisterInfosRequest) SetPageSize(v int32) *ListAtiAgentRegisterInfosRequest {
	s.PageSize = &v
	return s
}

func (s *ListAtiAgentRegisterInfosRequest) SetStatus(v string) *ListAtiAgentRegisterInfosRequest {
	s.Status = &v
	return s
}

func (s *ListAtiAgentRegisterInfosRequest) Validate() error {
	return dara.Validate(s)
}
