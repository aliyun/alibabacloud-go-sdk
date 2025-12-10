// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListImagesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLabels(v string) *ListImagesRequest
	GetLabels() *string
	SetName(v string) *ListImagesRequest
	GetName() *string
	SetOrder(v string) *ListImagesRequest
	GetOrder() *string
	SetPageNumber(v int32) *ListImagesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListImagesRequest
	GetPageSize() *int32
	SetSortBy(v string) *ListImagesRequest
	GetSortBy() *string
	SetVerbose(v bool) *ListImagesRequest
	GetVerbose() *bool
}

type ListImagesRequest struct {
	// example:
	//
	// Framework="Tensorflow 1.0",Framework="Tensorflow 2.0",Platform="GPU"
	Labels     *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Order      *string `json:"Order,omitempty" xml:"Order,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// GmtCreateTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// example:
	//
	// true
	Verbose *bool `json:"Verbose,omitempty" xml:"Verbose,omitempty"`
}

func (s ListImagesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListImagesRequest) GoString() string {
	return s.String()
}

func (s *ListImagesRequest) GetLabels() *string {
	return s.Labels
}

func (s *ListImagesRequest) GetName() *string {
	return s.Name
}

func (s *ListImagesRequest) GetOrder() *string {
	return s.Order
}

func (s *ListImagesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListImagesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListImagesRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListImagesRequest) GetVerbose() *bool {
	return s.Verbose
}

func (s *ListImagesRequest) SetLabels(v string) *ListImagesRequest {
	s.Labels = &v
	return s
}

func (s *ListImagesRequest) SetName(v string) *ListImagesRequest {
	s.Name = &v
	return s
}

func (s *ListImagesRequest) SetOrder(v string) *ListImagesRequest {
	s.Order = &v
	return s
}

func (s *ListImagesRequest) SetPageNumber(v int32) *ListImagesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListImagesRequest) SetPageSize(v int32) *ListImagesRequest {
	s.PageSize = &v
	return s
}

func (s *ListImagesRequest) SetSortBy(v string) *ListImagesRequest {
	s.SortBy = &v
	return s
}

func (s *ListImagesRequest) SetVerbose(v bool) *ListImagesRequest {
	s.Verbose = &v
	return s
}

func (s *ListImagesRequest) Validate() error {
	return dara.Validate(s)
}
