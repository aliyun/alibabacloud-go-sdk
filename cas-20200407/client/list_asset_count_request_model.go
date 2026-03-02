// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAssetCountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int64) *ListAssetCountRequest
	GetCurrentPage() *int64
	SetEndDate(v int64) *ListAssetCountRequest
	GetEndDate() *int64
	SetShowSize(v int64) *ListAssetCountRequest
	GetShowSize() *int64
	SetStartDate(v int64) *ListAssetCountRequest
	GetStartDate() *int64
}

type ListAssetCountRequest struct {
	// example:
	//
	// 1
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// 2020-07-13
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// example:
	//
	// 1
	ShowSize *int64 `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
	// example:
	//
	// 2018-07-13
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s ListAssetCountRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAssetCountRequest) GoString() string {
	return s.String()
}

func (s *ListAssetCountRequest) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *ListAssetCountRequest) GetEndDate() *int64 {
	return s.EndDate
}

func (s *ListAssetCountRequest) GetShowSize() *int64 {
	return s.ShowSize
}

func (s *ListAssetCountRequest) GetStartDate() *int64 {
	return s.StartDate
}

func (s *ListAssetCountRequest) SetCurrentPage(v int64) *ListAssetCountRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListAssetCountRequest) SetEndDate(v int64) *ListAssetCountRequest {
	s.EndDate = &v
	return s
}

func (s *ListAssetCountRequest) SetShowSize(v int64) *ListAssetCountRequest {
	s.ShowSize = &v
	return s
}

func (s *ListAssetCountRequest) SetStartDate(v int64) *ListAssetCountRequest {
	s.StartDate = &v
	return s
}

func (s *ListAssetCountRequest) Validate() error {
	return dara.Validate(s)
}
