// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchInformation interface {
	dara.Model
	String() string
	GoString() string
	SetSearchTime(v int64) *SearchInformation
	GetSearchTime() *int64
	SetTotal(v int64) *SearchInformation
	GetTotal() *int64
}

type SearchInformation struct {
	SearchTime *int64 `json:"searchTime,omitempty" xml:"searchTime,omitempty"`
	Total      *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s SearchInformation) String() string {
	return dara.Prettify(s)
}

func (s SearchInformation) GoString() string {
	return s.String()
}

func (s *SearchInformation) GetSearchTime() *int64 {
	return s.SearchTime
}

func (s *SearchInformation) GetTotal() *int64 {
	return s.Total
}

func (s *SearchInformation) SetSearchTime(v int64) *SearchInformation {
	s.SearchTime = &v
	return s
}

func (s *SearchInformation) SetTotal(v int64) *SearchInformation {
	s.Total = &v
	return s
}

func (s *SearchInformation) Validate() error {
	return dara.Validate(s)
}
