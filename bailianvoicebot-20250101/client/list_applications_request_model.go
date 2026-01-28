// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessUnitId(v string) *ListApplicationsRequest
	GetBusinessUnitId() *string
	SetPageNumber(v int32) *ListApplicationsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListApplicationsRequest
	GetPageSize() *int32
	SetSearchPattern(v string) *ListApplicationsRequest
	GetSearchPattern() *string
}

type ListApplicationsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// llm-c11iig67g863rih8
	BusinessUnitId *string `json:"BusinessUnitId,omitempty" xml:"BusinessUnitId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SearchPattern *string `json:"SearchPattern,omitempty" xml:"SearchPattern,omitempty"`
}

func (s ListApplicationsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsRequest) GoString() string {
	return s.String()
}

func (s *ListApplicationsRequest) GetBusinessUnitId() *string {
	return s.BusinessUnitId
}

func (s *ListApplicationsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListApplicationsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListApplicationsRequest) GetSearchPattern() *string {
	return s.SearchPattern
}

func (s *ListApplicationsRequest) SetBusinessUnitId(v string) *ListApplicationsRequest {
	s.BusinessUnitId = &v
	return s
}

func (s *ListApplicationsRequest) SetPageNumber(v int32) *ListApplicationsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListApplicationsRequest) SetPageSize(v int32) *ListApplicationsRequest {
	s.PageSize = &v
	return s
}

func (s *ListApplicationsRequest) SetSearchPattern(v string) *ListApplicationsRequest {
	s.SearchPattern = &v
	return s
}

func (s *ListApplicationsRequest) Validate() error {
	return dara.Validate(s)
}
