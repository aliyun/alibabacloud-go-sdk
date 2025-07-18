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
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	Dimension    *int32  `json:"Dimension,omitempty" xml:"Dimension,omitempty"`
	InputShrink  *string `json:"Input,omitempty" xml:"Input,omitempty"`
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
