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
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string   `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	Dimension    *int32    `json:"Dimension,omitempty" xml:"Dimension,omitempty"`
	Input        []*string `json:"Input,omitempty" xml:"Input,omitempty" type:"Repeated"`
	// example:
	//
	// text-embedding-v2
	Model   *string `json:"Model,omitempty" xml:"Model,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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
