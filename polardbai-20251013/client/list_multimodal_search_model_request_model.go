// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMultimodalSearchModelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ListMultimodalSearchModelRequest
	GetDBClusterId() *string
	SetPageNumber(v int32) *ListMultimodalSearchModelRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListMultimodalSearchModelRequest
	GetPageSize() *int32
}

type ListMultimodalSearchModelRequest struct {
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
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListMultimodalSearchModelRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMultimodalSearchModelRequest) GoString() string {
	return s.String()
}

func (s *ListMultimodalSearchModelRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ListMultimodalSearchModelRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListMultimodalSearchModelRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMultimodalSearchModelRequest) SetDBClusterId(v string) *ListMultimodalSearchModelRequest {
	s.DBClusterId = &v
	return s
}

func (s *ListMultimodalSearchModelRequest) SetPageNumber(v int32) *ListMultimodalSearchModelRequest {
	s.PageNumber = &v
	return s
}

func (s *ListMultimodalSearchModelRequest) SetPageSize(v int32) *ListMultimodalSearchModelRequest {
	s.PageSize = &v
	return s
}

func (s *ListMultimodalSearchModelRequest) Validate() error {
	return dara.Validate(s)
}
