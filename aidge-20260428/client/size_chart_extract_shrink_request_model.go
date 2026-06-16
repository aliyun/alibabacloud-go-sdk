// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSizeChartExtractShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetColumnNameListShrink(v string) *SizeChartExtractShrinkRequest
	GetColumnNameListShrink() *string
	SetImageUrl(v string) *SizeChartExtractShrinkRequest
	GetImageUrl() *string
	SetLanguageModel(v string) *SizeChartExtractShrinkRequest
	GetLanguageModel() *string
}

type SizeChartExtractShrinkRequest struct {
	// The list of column names to extract, such as Size, Bust, and Length.
	//
	// example:
	//
	// ["Size","Bust","Length"]
	ColumnNameListShrink *string `json:"ColumnNameList,omitempty" xml:"ColumnNameList,omitempty"`
	// The URL of the size chart image to extract.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://example.com/size_chart.jpg
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// The language model that controls the output language, such as en and cn.
	//
	// example:
	//
	// en
	LanguageModel *string `json:"LanguageModel,omitempty" xml:"LanguageModel,omitempty"`
}

func (s SizeChartExtractShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SizeChartExtractShrinkRequest) GoString() string {
	return s.String()
}

func (s *SizeChartExtractShrinkRequest) GetColumnNameListShrink() *string {
	return s.ColumnNameListShrink
}

func (s *SizeChartExtractShrinkRequest) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *SizeChartExtractShrinkRequest) GetLanguageModel() *string {
	return s.LanguageModel
}

func (s *SizeChartExtractShrinkRequest) SetColumnNameListShrink(v string) *SizeChartExtractShrinkRequest {
	s.ColumnNameListShrink = &v
	return s
}

func (s *SizeChartExtractShrinkRequest) SetImageUrl(v string) *SizeChartExtractShrinkRequest {
	s.ImageUrl = &v
	return s
}

func (s *SizeChartExtractShrinkRequest) SetLanguageModel(v string) *SizeChartExtractShrinkRequest {
	s.LanguageModel = &v
	return s
}

func (s *SizeChartExtractShrinkRequest) Validate() error {
	return dara.Validate(s)
}
