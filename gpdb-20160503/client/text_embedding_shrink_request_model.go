// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTextEmbeddingShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *TextEmbeddingShrinkRequest
	GetDBInstanceId() *string
	SetDimension(v int32) *TextEmbeddingShrinkRequest
	GetDimension() *int32
	SetInputShrink(v string) *TextEmbeddingShrinkRequest
	GetInputShrink() *string
	SetModel(v string) *TextEmbeddingShrinkRequest
	GetModel() *string
	SetOwnerId(v int64) *TextEmbeddingShrinkRequest
	GetOwnerId() *int64
	SetRegionId(v string) *TextEmbeddingShrinkRequest
	GetRegionId() *string
}

type TextEmbeddingShrinkRequest struct {
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
	InputShrink *string `json:"Input,omitempty" xml:"Input,omitempty"`
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

func (s TextEmbeddingShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s TextEmbeddingShrinkRequest) GoString() string {
	return s.String()
}

func (s *TextEmbeddingShrinkRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *TextEmbeddingShrinkRequest) GetDimension() *int32 {
	return s.Dimension
}

func (s *TextEmbeddingShrinkRequest) GetInputShrink() *string {
	return s.InputShrink
}

func (s *TextEmbeddingShrinkRequest) GetModel() *string {
	return s.Model
}

func (s *TextEmbeddingShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *TextEmbeddingShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *TextEmbeddingShrinkRequest) SetDBInstanceId(v string) *TextEmbeddingShrinkRequest {
	s.DBInstanceId = &v
	return s
}

func (s *TextEmbeddingShrinkRequest) SetDimension(v int32) *TextEmbeddingShrinkRequest {
	s.Dimension = &v
	return s
}

func (s *TextEmbeddingShrinkRequest) SetInputShrink(v string) *TextEmbeddingShrinkRequest {
	s.InputShrink = &v
	return s
}

func (s *TextEmbeddingShrinkRequest) SetModel(v string) *TextEmbeddingShrinkRequest {
	s.Model = &v
	return s
}

func (s *TextEmbeddingShrinkRequest) SetOwnerId(v int64) *TextEmbeddingShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *TextEmbeddingShrinkRequest) SetRegionId(v string) *TextEmbeddingShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *TextEmbeddingShrinkRequest) Validate() error {
	return dara.Validate(s)
}
