// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListToolsOutputV2 interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*Tool) *ListToolsOutputV2
	GetItems() []*Tool
	SetPageNumber(v int32) *ListToolsOutputV2
	GetPageNumber() *int32
	SetPageSize(v int32) *ListToolsOutputV2
	GetPageSize() *int32
	SetTotal(v int32) *ListToolsOutputV2
	GetTotal() *int32
}

type ListToolsOutputV2 struct {
	// 当前页的工具详细信息列表
	Items []*Tool `json:"items" xml:"items" type:"Repeated"`
	// 当前页码，从 1 开始
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 每页返回的工具数量
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 符合条件的工具总数
	//
	// example:
	//
	// 100
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListToolsOutputV2) String() string {
	return dara.Prettify(s)
}

func (s ListToolsOutputV2) GoString() string {
	return s.String()
}

func (s *ListToolsOutputV2) GetItems() []*Tool {
	return s.Items
}

func (s *ListToolsOutputV2) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListToolsOutputV2) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListToolsOutputV2) GetTotal() *int32 {
	return s.Total
}

func (s *ListToolsOutputV2) SetItems(v []*Tool) *ListToolsOutputV2 {
	s.Items = v
	return s
}

func (s *ListToolsOutputV2) SetPageNumber(v int32) *ListToolsOutputV2 {
	s.PageNumber = &v
	return s
}

func (s *ListToolsOutputV2) SetPageSize(v int32) *ListToolsOutputV2 {
	s.PageSize = &v
	return s
}

func (s *ListToolsOutputV2) SetTotal(v int32) *ListToolsOutputV2 {
	s.Total = &v
	return s
}

func (s *ListToolsOutputV2) Validate() error {
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
