// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMultimodalSearchTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ListMultimodalSearchTaskRequest
	GetDBClusterId() *string
	SetPageNumber(v int32) *ListMultimodalSearchTaskRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListMultimodalSearchTaskRequest
	GetPageSize() *int32
}

type ListMultimodalSearchTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pc-2ze454l20me07****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListMultimodalSearchTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMultimodalSearchTaskRequest) GoString() string {
	return s.String()
}

func (s *ListMultimodalSearchTaskRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ListMultimodalSearchTaskRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListMultimodalSearchTaskRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMultimodalSearchTaskRequest) SetDBClusterId(v string) *ListMultimodalSearchTaskRequest {
	s.DBClusterId = &v
	return s
}

func (s *ListMultimodalSearchTaskRequest) SetPageNumber(v int32) *ListMultimodalSearchTaskRequest {
	s.PageNumber = &v
	return s
}

func (s *ListMultimodalSearchTaskRequest) SetPageSize(v int32) *ListMultimodalSearchTaskRequest {
	s.PageSize = &v
	return s
}

func (s *ListMultimodalSearchTaskRequest) Validate() error {
	return dara.Validate(s)
}
