// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentResourcesOutput interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*AgentResource) *ListAgentResourcesOutput
	GetItems() []*AgentResource
	SetPageNumber(v int32) *ListAgentResourcesOutput
	GetPageNumber() *int32
	SetPageSize(v int32) *ListAgentResourcesOutput
	GetPageSize() *int32
	SetTotal(v int64) *ListAgentResourcesOutput
	GetTotal() *int64
}

type ListAgentResourcesOutput struct {
	// 智能体资源的列表
	Items []*AgentResource `json:"items" xml:"items" type:"Repeated"`
	// 当前页码，从 1 开始计数
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 每页返回的记录数量
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 符合条件的资源总数
	//
	// example:
	//
	// 100
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListAgentResourcesOutput) String() string {
	return dara.Prettify(s)
}

func (s ListAgentResourcesOutput) GoString() string {
	return s.String()
}

func (s *ListAgentResourcesOutput) GetItems() []*AgentResource {
	return s.Items
}

func (s *ListAgentResourcesOutput) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAgentResourcesOutput) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAgentResourcesOutput) GetTotal() *int64 {
	return s.Total
}

func (s *ListAgentResourcesOutput) SetItems(v []*AgentResource) *ListAgentResourcesOutput {
	s.Items = v
	return s
}

func (s *ListAgentResourcesOutput) SetPageNumber(v int32) *ListAgentResourcesOutput {
	s.PageNumber = &v
	return s
}

func (s *ListAgentResourcesOutput) SetPageSize(v int32) *ListAgentResourcesOutput {
	s.PageSize = &v
	return s
}

func (s *ListAgentResourcesOutput) SetTotal(v int64) *ListAgentResourcesOutput {
	s.Total = &v
	return s
}

func (s *ListAgentResourcesOutput) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
