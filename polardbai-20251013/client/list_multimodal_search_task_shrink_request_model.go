// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMultimodalSearchTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ListMultimodalSearchTaskShrinkRequest
	GetDBClusterId() *string
	SetDatasetIdsShrink(v string) *ListMultimodalSearchTaskShrinkRequest
	GetDatasetIdsShrink() *string
	SetInputField(v string) *ListMultimodalSearchTaskShrinkRequest
	GetInputField() *string
	SetModelMode(v string) *ListMultimodalSearchTaskShrinkRequest
	GetModelMode() *string
	SetPageNumber(v int32) *ListMultimodalSearchTaskShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListMultimodalSearchTaskShrinkRequest
	GetPageSize() *int32
}

type ListMultimodalSearchTaskShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pc-2ze454l20me07****
	DBClusterId      *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	DatasetIdsShrink *string `json:"DatasetIds,omitempty" xml:"DatasetIds,omitempty"`
	// example:
	//
	// 红绿灯
	InputField *string `json:"InputField,omitempty" xml:"InputField,omitempty"`
	// example:
	//
	// flash，plus
	ModelMode *string `json:"ModelMode,omitempty" xml:"ModelMode,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListMultimodalSearchTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMultimodalSearchTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListMultimodalSearchTaskShrinkRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ListMultimodalSearchTaskShrinkRequest) GetDatasetIdsShrink() *string {
	return s.DatasetIdsShrink
}

func (s *ListMultimodalSearchTaskShrinkRequest) GetInputField() *string {
	return s.InputField
}

func (s *ListMultimodalSearchTaskShrinkRequest) GetModelMode() *string {
	return s.ModelMode
}

func (s *ListMultimodalSearchTaskShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListMultimodalSearchTaskShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMultimodalSearchTaskShrinkRequest) SetDBClusterId(v string) *ListMultimodalSearchTaskShrinkRequest {
	s.DBClusterId = &v
	return s
}

func (s *ListMultimodalSearchTaskShrinkRequest) SetDatasetIdsShrink(v string) *ListMultimodalSearchTaskShrinkRequest {
	s.DatasetIdsShrink = &v
	return s
}

func (s *ListMultimodalSearchTaskShrinkRequest) SetInputField(v string) *ListMultimodalSearchTaskShrinkRequest {
	s.InputField = &v
	return s
}

func (s *ListMultimodalSearchTaskShrinkRequest) SetModelMode(v string) *ListMultimodalSearchTaskShrinkRequest {
	s.ModelMode = &v
	return s
}

func (s *ListMultimodalSearchTaskShrinkRequest) SetPageNumber(v int32) *ListMultimodalSearchTaskShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListMultimodalSearchTaskShrinkRequest) SetPageSize(v int32) *ListMultimodalSearchTaskShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListMultimodalSearchTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}
