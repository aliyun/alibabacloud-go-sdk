// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSizeChartDetectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageUrl(v string) *SizeChartDetectRequest
	GetImageUrl() *string
	SetThreshold(v float64) *SizeChartDetectRequest
	GetThreshold() *float64
}

type SizeChartDetectRequest struct {
	// The URL of the image to detect.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://img.alicdn.com/imgextra/i4/1822632490/O1CN01pbENR21UGT9w3a6gU_!!1822632490.jpg
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// The detection threshold. Valid values: 0 to 1.
	//
	// example:
	//
	// 1
	Threshold *float64 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s SizeChartDetectRequest) String() string {
	return dara.Prettify(s)
}

func (s SizeChartDetectRequest) GoString() string {
	return s.String()
}

func (s *SizeChartDetectRequest) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *SizeChartDetectRequest) GetThreshold() *float64 {
	return s.Threshold
}

func (s *SizeChartDetectRequest) SetImageUrl(v string) *SizeChartDetectRequest {
	s.ImageUrl = &v
	return s
}

func (s *SizeChartDetectRequest) SetThreshold(v float64) *SizeChartDetectRequest {
	s.Threshold = &v
	return s
}

func (s *SizeChartDetectRequest) Validate() error {
	return dara.Validate(s)
}
