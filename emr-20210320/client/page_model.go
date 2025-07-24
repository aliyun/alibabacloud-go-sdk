// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPage interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*string) *Page
	GetItems() []*string
	SetMaxResults(v int64) *Page
	GetMaxResults() *int64
	SetNextToken(v string) *Page
	GetNextToken() *string
	SetTotalCount(v int64) *Page
	GetTotalCount() *int64
}

type Page struct {
	Items      []*string `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	MaxResults *int64    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	TotalCount *int64    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s Page) String() string {
	return dara.Prettify(s)
}

func (s Page) GoString() string {
	return s.String()
}

func (s *Page) GetItems() []*string {
	return s.Items
}

func (s *Page) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *Page) GetNextToken() *string {
	return s.NextToken
}

func (s *Page) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *Page) SetItems(v []*string) *Page {
	s.Items = v
	return s
}

func (s *Page) SetMaxResults(v int64) *Page {
	s.MaxResults = &v
	return s
}

func (s *Page) SetNextToken(v string) *Page {
	s.NextToken = &v
	return s
}

func (s *Page) SetTotalCount(v int64) *Page {
	s.TotalCount = &v
	return s
}

func (s *Page) Validate() error {
	return dara.Validate(s)
}
