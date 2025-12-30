// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMultimodalEmbeddingModelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ListMultimodalEmbeddingModelRequest
	GetDBClusterId() *string
	SetPageNumber(v int32) *ListMultimodalEmbeddingModelRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListMultimodalEmbeddingModelRequest
	GetPageSize() *int32
}

type ListMultimodalEmbeddingModelRequest struct {
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
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListMultimodalEmbeddingModelRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMultimodalEmbeddingModelRequest) GoString() string {
	return s.String()
}

func (s *ListMultimodalEmbeddingModelRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ListMultimodalEmbeddingModelRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListMultimodalEmbeddingModelRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMultimodalEmbeddingModelRequest) SetDBClusterId(v string) *ListMultimodalEmbeddingModelRequest {
	s.DBClusterId = &v
	return s
}

func (s *ListMultimodalEmbeddingModelRequest) SetPageNumber(v int32) *ListMultimodalEmbeddingModelRequest {
	s.PageNumber = &v
	return s
}

func (s *ListMultimodalEmbeddingModelRequest) SetPageSize(v int32) *ListMultimodalEmbeddingModelRequest {
	s.PageSize = &v
	return s
}

func (s *ListMultimodalEmbeddingModelRequest) Validate() error {
	return dara.Validate(s)
}
