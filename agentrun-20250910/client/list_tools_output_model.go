// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListToolsOutput interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ToolListItem) *ListToolsOutput
	GetData() []*ToolListItem
	SetPageNum(v int32) *ListToolsOutput
	GetPageNum() *int32
	SetPageSize(v int32) *ListToolsOutput
	GetPageSize() *int32
	SetTotal(v int32) *ListToolsOutput
	GetTotal() *int32
}

type ListToolsOutput struct {
	Data     []*ToolListItem `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	PageNum  *int32          `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
	PageSize *int32          `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Total    *int32          `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListToolsOutput) String() string {
	return dara.Prettify(s)
}

func (s ListToolsOutput) GoString() string {
	return s.String()
}

func (s *ListToolsOutput) GetData() []*ToolListItem {
	return s.Data
}

func (s *ListToolsOutput) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListToolsOutput) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListToolsOutput) GetTotal() *int32 {
	return s.Total
}

func (s *ListToolsOutput) SetData(v []*ToolListItem) *ListToolsOutput {
	s.Data = v
	return s
}

func (s *ListToolsOutput) SetPageNum(v int32) *ListToolsOutput {
	s.PageNum = &v
	return s
}

func (s *ListToolsOutput) SetPageSize(v int32) *ListToolsOutput {
	s.PageSize = &v
	return s
}

func (s *ListToolsOutput) SetTotal(v int32) *ListToolsOutput {
	s.Total = &v
	return s
}

func (s *ListToolsOutput) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
