// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSizeChartExtractRequest interface {
	dara.Model
	String() string
	GoString() string
	SetColumnNameList(v []*string) *SizeChartExtractRequest
	GetColumnNameList() []*string
	SetImageUrl(v string) *SizeChartExtractRequest
	GetImageUrl() *string
	SetLanguageModel(v string) *SizeChartExtractRequest
	GetLanguageModel() *string
}

type SizeChartExtractRequest struct {
	// The list of column names to extract, such as Size, Bust, and Length.
	//
	// example:
	//
	// ["Size","Bust","Length"]
	ColumnNameList []*string `json:"ColumnNameList,omitempty" xml:"ColumnNameList,omitempty" type:"Repeated"`
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

func (s SizeChartExtractRequest) String() string {
	return dara.Prettify(s)
}

func (s SizeChartExtractRequest) GoString() string {
	return s.String()
}

func (s *SizeChartExtractRequest) GetColumnNameList() []*string {
	return s.ColumnNameList
}

func (s *SizeChartExtractRequest) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *SizeChartExtractRequest) GetLanguageModel() *string {
	return s.LanguageModel
}

func (s *SizeChartExtractRequest) SetColumnNameList(v []*string) *SizeChartExtractRequest {
	s.ColumnNameList = v
	return s
}

func (s *SizeChartExtractRequest) SetImageUrl(v string) *SizeChartExtractRequest {
	s.ImageUrl = &v
	return s
}

func (s *SizeChartExtractRequest) SetLanguageModel(v string) *SizeChartExtractRequest {
	s.LanguageModel = &v
	return s
}

func (s *SizeChartExtractRequest) Validate() error {
	return dara.Validate(s)
}
