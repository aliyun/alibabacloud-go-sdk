// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKnowledgeBasesOutput interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*KnowledgeBase) *ListKnowledgeBasesOutput
	GetItems() []*KnowledgeBase
	SetPageNumber(v int32) *ListKnowledgeBasesOutput
	GetPageNumber() *int32
	SetPageSize(v int32) *ListKnowledgeBasesOutput
	GetPageSize() *int32
	SetTotal(v int64) *ListKnowledgeBasesOutput
	GetTotal() *int64
}

type ListKnowledgeBasesOutput struct {
	Items      []*KnowledgeBase `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	PageNumber *int32           `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int32           `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Total      *int64           `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListKnowledgeBasesOutput) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeBasesOutput) GoString() string {
	return s.String()
}

func (s *ListKnowledgeBasesOutput) GetItems() []*KnowledgeBase {
	return s.Items
}

func (s *ListKnowledgeBasesOutput) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListKnowledgeBasesOutput) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListKnowledgeBasesOutput) GetTotal() *int64 {
	return s.Total
}

func (s *ListKnowledgeBasesOutput) SetItems(v []*KnowledgeBase) *ListKnowledgeBasesOutput {
	s.Items = v
	return s
}

func (s *ListKnowledgeBasesOutput) SetPageNumber(v int32) *ListKnowledgeBasesOutput {
	s.PageNumber = &v
	return s
}

func (s *ListKnowledgeBasesOutput) SetPageSize(v int32) *ListKnowledgeBasesOutput {
	s.PageSize = &v
	return s
}

func (s *ListKnowledgeBasesOutput) SetTotal(v int64) *ListKnowledgeBasesOutput {
	s.Total = &v
	return s
}

func (s *ListKnowledgeBasesOutput) Validate() error {
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
