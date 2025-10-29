// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBrowserSessionListOut interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*BrowserSessionOut) *BrowserSessionListOut
	GetItems() []*BrowserSessionOut
	SetPageNumber(v int32) *BrowserSessionListOut
	GetPageNumber() *int32
	SetPageSize(v int32) *BrowserSessionListOut
	GetPageSize() *int32
	SetTotal(v int64) *BrowserSessionListOut
	GetTotal() *int64
}

type BrowserSessionListOut struct {
	// example:
	//
	// []
	Items []*BrowserSessionOut `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 100
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s BrowserSessionListOut) String() string {
	return dara.Prettify(s)
}

func (s BrowserSessionListOut) GoString() string {
	return s.String()
}

func (s *BrowserSessionListOut) GetItems() []*BrowserSessionOut {
	return s.Items
}

func (s *BrowserSessionListOut) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *BrowserSessionListOut) GetPageSize() *int32 {
	return s.PageSize
}

func (s *BrowserSessionListOut) GetTotal() *int64 {
	return s.Total
}

func (s *BrowserSessionListOut) SetItems(v []*BrowserSessionOut) *BrowserSessionListOut {
	s.Items = v
	return s
}

func (s *BrowserSessionListOut) SetPageNumber(v int32) *BrowserSessionListOut {
	s.PageNumber = &v
	return s
}

func (s *BrowserSessionListOut) SetPageSize(v int32) *BrowserSessionListOut {
	s.PageSize = &v
	return s
}

func (s *BrowserSessionListOut) SetTotal(v int64) *BrowserSessionListOut {
	s.Total = &v
	return s
}

func (s *BrowserSessionListOut) Validate() error {
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
