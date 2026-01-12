// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMemoryCollectionsOutput interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*MemoryCollection) *ListMemoryCollectionsOutput
	GetItems() []*MemoryCollection
	SetPageNumber(v int32) *ListMemoryCollectionsOutput
	GetPageNumber() *int32
	SetPageSize(v int32) *ListMemoryCollectionsOutput
	GetPageSize() *int32
	SetTotal(v int64) *ListMemoryCollectionsOutput
	GetTotal() *int64
}

type ListMemoryCollectionsOutput struct {
	Items      []*MemoryCollection `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	PageNumber *int32              `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int32              `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Total      *int64              `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListMemoryCollectionsOutput) String() string {
	return dara.Prettify(s)
}

func (s ListMemoryCollectionsOutput) GoString() string {
	return s.String()
}

func (s *ListMemoryCollectionsOutput) GetItems() []*MemoryCollection {
	return s.Items
}

func (s *ListMemoryCollectionsOutput) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListMemoryCollectionsOutput) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMemoryCollectionsOutput) GetTotal() *int64 {
	return s.Total
}

func (s *ListMemoryCollectionsOutput) SetItems(v []*MemoryCollection) *ListMemoryCollectionsOutput {
	s.Items = v
	return s
}

func (s *ListMemoryCollectionsOutput) SetPageNumber(v int32) *ListMemoryCollectionsOutput {
	s.PageNumber = &v
	return s
}

func (s *ListMemoryCollectionsOutput) SetPageSize(v int32) *ListMemoryCollectionsOutput {
	s.PageSize = &v
	return s
}

func (s *ListMemoryCollectionsOutput) SetTotal(v int64) *ListMemoryCollectionsOutput {
	s.Total = &v
	return s
}

func (s *ListMemoryCollectionsOutput) Validate() error {
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
