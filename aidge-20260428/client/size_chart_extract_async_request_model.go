// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSizeChartExtractAsyncRequest interface {
	dara.Model
	String() string
	GoString() string
	SetColumnNameList(v []*string) *SizeChartExtractAsyncRequest
	GetColumnNameList() []*string
	SetImageUrl(v string) *SizeChartExtractAsyncRequest
	GetImageUrl() *string
	SetLanguageModel(v string) *SizeChartExtractAsyncRequest
	GetLanguageModel() *string
}

type SizeChartExtractAsyncRequest struct {
	// The list of column names to extract, such as Size, Bust, and Length.
	ColumnNameList []*string `json:"ColumnNameList,omitempty" xml:"ColumnNameList,omitempty" type:"Repeated"`
	// The image URL. The image must be in the JPG, JPEG, PNG, BMP, or WEBP format. The resolution must be between 256 × 256 pixels and 3000 × 3000 pixels. The file size cannot exceed 10 MB.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://ae01.alicdn.com/kf/S342f0070dc9f4be09a6cbed34e90dc8fs.jpg
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// The output language format. If you do not specify this parameter, the original language is used. If the original image language is Chinese or English, specifying en outputs English, and specifying cn outputs Chinese.
	//
	// example:
	//
	// en
	LanguageModel *string `json:"LanguageModel,omitempty" xml:"LanguageModel,omitempty"`
}

func (s SizeChartExtractAsyncRequest) String() string {
	return dara.Prettify(s)
}

func (s SizeChartExtractAsyncRequest) GoString() string {
	return s.String()
}

func (s *SizeChartExtractAsyncRequest) GetColumnNameList() []*string {
	return s.ColumnNameList
}

func (s *SizeChartExtractAsyncRequest) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *SizeChartExtractAsyncRequest) GetLanguageModel() *string {
	return s.LanguageModel
}

func (s *SizeChartExtractAsyncRequest) SetColumnNameList(v []*string) *SizeChartExtractAsyncRequest {
	s.ColumnNameList = v
	return s
}

func (s *SizeChartExtractAsyncRequest) SetImageUrl(v string) *SizeChartExtractAsyncRequest {
	s.ImageUrl = &v
	return s
}

func (s *SizeChartExtractAsyncRequest) SetLanguageModel(v string) *SizeChartExtractAsyncRequest {
	s.LanguageModel = &v
	return s
}

func (s *SizeChartExtractAsyncRequest) Validate() error {
	return dara.Validate(s)
}
