// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModelConnectionsOutput interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*ModelConnection) *ListModelConnectionsOutput
	GetItems() []*ModelConnection
	SetPageNumber(v int) *ListModelConnectionsOutput
	GetPageNumber() *int
	SetPageSize(v int) *ListModelConnectionsOutput
	GetPageSize() *int
	SetTotal(v int) *ListModelConnectionsOutput
	GetTotal() *int
}

type ListModelConnectionsOutput struct {
	// 模型连接列表
	Items []*ModelConnection `json:"items" xml:"items" type:"Repeated"`
	// 当前页码
	//
	// example:
	//
	// 1
	PageNumber *int `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 每页记录数
	//
	// example:
	//
	// 20
	PageSize *int `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 符合条件的模型连接总数
	//
	// example:
	//
	// 10
	Total *int `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListModelConnectionsOutput) String() string {
	return dara.Prettify(s)
}

func (s ListModelConnectionsOutput) GoString() string {
	return s.String()
}

func (s *ListModelConnectionsOutput) GetItems() []*ModelConnection {
	return s.Items
}

func (s *ListModelConnectionsOutput) GetPageNumber() *int {
	return s.PageNumber
}

func (s *ListModelConnectionsOutput) GetPageSize() *int {
	return s.PageSize
}

func (s *ListModelConnectionsOutput) GetTotal() *int {
	return s.Total
}

func (s *ListModelConnectionsOutput) SetItems(v []*ModelConnection) *ListModelConnectionsOutput {
	s.Items = v
	return s
}

func (s *ListModelConnectionsOutput) SetPageNumber(v int) *ListModelConnectionsOutput {
	s.PageNumber = &v
	return s
}

func (s *ListModelConnectionsOutput) SetPageSize(v int) *ListModelConnectionsOutput {
	s.PageSize = &v
	return s
}

func (s *ListModelConnectionsOutput) SetTotal(v int) *ListModelConnectionsOutput {
	s.Total = &v
	return s
}

func (s *ListModelConnectionsOutput) Validate() error {
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
