// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataQualityTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCatalog(v string) *ListDataQualityTemplatesRequest
	GetCatalog() *string
	SetName(v string) *ListDataQualityTemplatesRequest
	GetName() *string
	SetPageNumber(v int32) *ListDataQualityTemplatesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDataQualityTemplatesRequest
	GetPageSize() *int32
	SetProjectId(v int64) *ListDataQualityTemplatesRequest
	GetProjectId() *int64
}

type ListDataQualityTemplatesRequest struct {
	Catalog *string `json:"Catalog,omitempty" xml:"Catalog,omitempty"`
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ListDataQualityTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataQualityTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListDataQualityTemplatesRequest) GetCatalog() *string {
	return s.Catalog
}

func (s *ListDataQualityTemplatesRequest) GetName() *string {
	return s.Name
}

func (s *ListDataQualityTemplatesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDataQualityTemplatesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDataQualityTemplatesRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListDataQualityTemplatesRequest) SetCatalog(v string) *ListDataQualityTemplatesRequest {
	s.Catalog = &v
	return s
}

func (s *ListDataQualityTemplatesRequest) SetName(v string) *ListDataQualityTemplatesRequest {
	s.Name = &v
	return s
}

func (s *ListDataQualityTemplatesRequest) SetPageNumber(v int32) *ListDataQualityTemplatesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDataQualityTemplatesRequest) SetPageSize(v int32) *ListDataQualityTemplatesRequest {
	s.PageSize = &v
	return s
}

func (s *ListDataQualityTemplatesRequest) SetProjectId(v int64) *ListDataQualityTemplatesRequest {
	s.ProjectId = &v
	return s
}

func (s *ListDataQualityTemplatesRequest) Validate() error {
	return dara.Validate(s)
}
