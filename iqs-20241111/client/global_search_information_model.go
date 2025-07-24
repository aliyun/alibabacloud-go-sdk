// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGlobalSearchInformation interface {
	dara.Model
	String() string
	GoString() string
	SetSearchTime(v int64) *GlobalSearchInformation
	GetSearchTime() *int64
	SetTotal(v int64) *GlobalSearchInformation
	GetTotal() *int64
}

type GlobalSearchInformation struct {
	SearchTime *int64 `json:"searchTime,omitempty" xml:"searchTime,omitempty"`
	Total      *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s GlobalSearchInformation) String() string {
	return dara.Prettify(s)
}

func (s GlobalSearchInformation) GoString() string {
	return s.String()
}

func (s *GlobalSearchInformation) GetSearchTime() *int64 {
	return s.SearchTime
}

func (s *GlobalSearchInformation) GetTotal() *int64 {
	return s.Total
}

func (s *GlobalSearchInformation) SetSearchTime(v int64) *GlobalSearchInformation {
	s.SearchTime = &v
	return s
}

func (s *GlobalSearchInformation) SetTotal(v int64) *GlobalSearchInformation {
	s.Total = &v
	return s
}

func (s *GlobalSearchInformation) Validate() error {
	return dara.Validate(s)
}
