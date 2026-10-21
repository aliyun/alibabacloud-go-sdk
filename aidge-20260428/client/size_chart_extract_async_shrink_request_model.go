// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSizeChartExtractAsyncShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *SizeChartExtractAsyncShrinkRequest
	GetBizId() *string
	SetColumnNameListShrink(v string) *SizeChartExtractAsyncShrinkRequest
	GetColumnNameListShrink() *string
	SetImageUrl(v string) *SizeChartExtractAsyncShrinkRequest
	GetImageUrl() *string
	SetLanguageModel(v string) *SizeChartExtractAsyncShrinkRequest
	GetLanguageModel() *string
}

type SizeChartExtractAsyncShrinkRequest struct {
	// The business order number of the caller. It is recorded when the request is accepted and returned as is when you query the asynchronous task result.
	//
	// example:
	//
	// order-20260928-001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// The list of column names to extract, such as Size, Bust, and Length.
	ColumnNameListShrink *string `json:"ColumnNameList,omitempty" xml:"ColumnNameList,omitempty"`
	// The image URL. The image must be in the JPG, JPEG, PNG, BMP, or WEBP format. The resolution must be between 256 × 256 pixels and 3000 × 3000 pixels, and the file size cannot exceed 10 MB.
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

func (s SizeChartExtractAsyncShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SizeChartExtractAsyncShrinkRequest) GoString() string {
	return s.String()
}

func (s *SizeChartExtractAsyncShrinkRequest) GetBizId() *string {
	return s.BizId
}

func (s *SizeChartExtractAsyncShrinkRequest) GetColumnNameListShrink() *string {
	return s.ColumnNameListShrink
}

func (s *SizeChartExtractAsyncShrinkRequest) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *SizeChartExtractAsyncShrinkRequest) GetLanguageModel() *string {
	return s.LanguageModel
}

func (s *SizeChartExtractAsyncShrinkRequest) SetBizId(v string) *SizeChartExtractAsyncShrinkRequest {
	s.BizId = &v
	return s
}

func (s *SizeChartExtractAsyncShrinkRequest) SetColumnNameListShrink(v string) *SizeChartExtractAsyncShrinkRequest {
	s.ColumnNameListShrink = &v
	return s
}

func (s *SizeChartExtractAsyncShrinkRequest) SetImageUrl(v string) *SizeChartExtractAsyncShrinkRequest {
	s.ImageUrl = &v
	return s
}

func (s *SizeChartExtractAsyncShrinkRequest) SetLanguageModel(v string) *SizeChartExtractAsyncShrinkRequest {
	s.LanguageModel = &v
	return s
}

func (s *SizeChartExtractAsyncShrinkRequest) Validate() error {
	return dara.Validate(s)
}
