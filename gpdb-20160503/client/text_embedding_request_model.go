// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTextEmbeddingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *TextEmbeddingRequest
	GetDBInstanceId() *string
	SetDimension(v int32) *TextEmbeddingRequest
	GetDimension() *int32
	SetInput(v []*string) *TextEmbeddingRequest
	GetInput() []*string
	SetModel(v string) *TextEmbeddingRequest
	GetModel() *string
	SetOwnerId(v int64) *TextEmbeddingRequest
	GetOwnerId() *int64
	SetRegionId(v string) *TextEmbeddingRequest
	GetRegionId() *string
}

type TextEmbeddingRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The number of embedding dimensions. The default value is the number of dimensions supported by the embedding algorithm.
	//
	// >
	//
	// 	- The text-embedding-v3 supports 1024, 768, and 512 dimensions. Default value: 1024.
	//
	// example:
	//
	// 1024
	Dimension *int32 `json:"Dimension,omitempty" xml:"Dimension,omitempty"`
	// A list of text content to be embedded. The list length must not exceed 100.
	Input []*string `json:"Input,omitempty" xml:"Input,omitempty" type:"Repeated"`
	// The text embedding model. Valid values:
	//
	// 	- text-embedding-v1:1536 dimensions
	//
	// 	- text-embedding-v2:1536 dimensions
	//
	// 	- text-embedding-v3 (default):1024, 768, and 512 dimensions
	//
	// 	- text2vec: 1024 dimensions
	//
	// 	- m3e-base: 768 dimensions
	//
	// 	- m3e-small: 512 dimensions
	//
	// example:
	//
	// text-embedding-v2
	Model   *string `json:"Model,omitempty" xml:"Model,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s TextEmbeddingRequest) String() string {
	return dara.Prettify(s)
}

func (s TextEmbeddingRequest) GoString() string {
	return s.String()
}

func (s *TextEmbeddingRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *TextEmbeddingRequest) GetDimension() *int32 {
	return s.Dimension
}

func (s *TextEmbeddingRequest) GetInput() []*string {
	return s.Input
}

func (s *TextEmbeddingRequest) GetModel() *string {
	return s.Model
}

func (s *TextEmbeddingRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *TextEmbeddingRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *TextEmbeddingRequest) SetDBInstanceId(v string) *TextEmbeddingRequest {
	s.DBInstanceId = &v
	return s
}

func (s *TextEmbeddingRequest) SetDimension(v int32) *TextEmbeddingRequest {
	s.Dimension = &v
	return s
}

func (s *TextEmbeddingRequest) SetInput(v []*string) *TextEmbeddingRequest {
	s.Input = v
	return s
}

func (s *TextEmbeddingRequest) SetModel(v string) *TextEmbeddingRequest {
	s.Model = &v
	return s
}

func (s *TextEmbeddingRequest) SetOwnerId(v int64) *TextEmbeddingRequest {
	s.OwnerId = &v
	return s
}

func (s *TextEmbeddingRequest) SetRegionId(v string) *TextEmbeddingRequest {
	s.RegionId = &v
	return s
}

func (s *TextEmbeddingRequest) Validate() error {
	return dara.Validate(s)
}
