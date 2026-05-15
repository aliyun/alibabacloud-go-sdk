// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPageQueryAgentListNewRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v int64) *PageQueryAgentListNewRequest
	GetAgentId() *int64
	SetAgentName(v string) *PageQueryAgentListNewRequest
	GetAgentName() *string
	SetIsAvailable(v bool) *PageQueryAgentListNewRequest
	GetIsAvailable() *bool
	SetPageIndex(v int64) *PageQueryAgentListNewRequest
	GetPageIndex() *int64
	SetPageNo(v int64) *PageQueryAgentListNewRequest
	GetPageNo() *int64
	SetPageSize(v int64) *PageQueryAgentListNewRequest
	GetPageSize() *int64
}

type PageQueryAgentListNewRequest struct {
	// Agent ID
	//
	// example:
	//
	// 12345
	AgentId *int64 `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// example:
	//
	// 示例值示例值
	AgentName *string `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
	// example:
	//
	// true
	IsAvailable *bool `json:"IsAvailable,omitempty" xml:"IsAvailable,omitempty"`
	// example:
	//
	// 1
	PageIndex *int64 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s PageQueryAgentListNewRequest) String() string {
	return dara.Prettify(s)
}

func (s PageQueryAgentListNewRequest) GoString() string {
	return s.String()
}

func (s *PageQueryAgentListNewRequest) GetAgentId() *int64 {
	return s.AgentId
}

func (s *PageQueryAgentListNewRequest) GetAgentName() *string {
	return s.AgentName
}

func (s *PageQueryAgentListNewRequest) GetIsAvailable() *bool {
	return s.IsAvailable
}

func (s *PageQueryAgentListNewRequest) GetPageIndex() *int64 {
	return s.PageIndex
}

func (s *PageQueryAgentListNewRequest) GetPageNo() *int64 {
	return s.PageNo
}

func (s *PageQueryAgentListNewRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *PageQueryAgentListNewRequest) SetAgentId(v int64) *PageQueryAgentListNewRequest {
	s.AgentId = &v
	return s
}

func (s *PageQueryAgentListNewRequest) SetAgentName(v string) *PageQueryAgentListNewRequest {
	s.AgentName = &v
	return s
}

func (s *PageQueryAgentListNewRequest) SetIsAvailable(v bool) *PageQueryAgentListNewRequest {
	s.IsAvailable = &v
	return s
}

func (s *PageQueryAgentListNewRequest) SetPageIndex(v int64) *PageQueryAgentListNewRequest {
	s.PageIndex = &v
	return s
}

func (s *PageQueryAgentListNewRequest) SetPageNo(v int64) *PageQueryAgentListNewRequest {
	s.PageNo = &v
	return s
}

func (s *PageQueryAgentListNewRequest) SetPageSize(v int64) *PageQueryAgentListNewRequest {
	s.PageSize = &v
	return s
}

func (s *PageQueryAgentListNewRequest) Validate() error {
	return dara.Validate(s)
}
