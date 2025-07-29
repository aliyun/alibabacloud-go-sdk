// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserKubeConfigStatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListUserKubeConfigStatesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListUserKubeConfigStatesRequest
	GetPageSize() *int32
}

type ListUserKubeConfigStatesRequest struct {
	// The page number.
	//
	// 	- Valid values: ≥ 1.
	//
	// 	- Default value: 1.
	//
	// example:
	//
	// 2
	PageNumber *int32 `json:"page_number,omitempty" xml:"page_number,omitempty"`
	// The number of entries per page.
	//
	// 	- Value values: 1 to 100.
	//
	// 	- Default value: 50.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"page_size,omitempty" xml:"page_size,omitempty"`
}

func (s ListUserKubeConfigStatesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUserKubeConfigStatesRequest) GoString() string {
	return s.String()
}

func (s *ListUserKubeConfigStatesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListUserKubeConfigStatesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListUserKubeConfigStatesRequest) SetPageNumber(v int32) *ListUserKubeConfigStatesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListUserKubeConfigStatesRequest) SetPageSize(v int32) *ListUserKubeConfigStatesRequest {
	s.PageSize = &v
	return s
}

func (s *ListUserKubeConfigStatesRequest) Validate() error {
	return dara.Validate(s)
}
