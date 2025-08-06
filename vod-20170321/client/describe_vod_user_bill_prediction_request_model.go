// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodUserBillPredictionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArea(v string) *DescribeVodUserBillPredictionRequest
	GetArea() *string
	SetDimension(v string) *DescribeVodUserBillPredictionRequest
	GetDimension() *string
	SetOwnerId(v int64) *DescribeVodUserBillPredictionRequest
	GetOwnerId() *int64
}

type DescribeVodUserBillPredictionRequest struct {
	Area *string `json:"Area,omitempty" xml:"Area,omitempty"`
	// This parameter is required.
	Dimension *string `json:"Dimension,omitempty" xml:"Dimension,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribeVodUserBillPredictionRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodUserBillPredictionRequest) GoString() string {
	return s.String()
}

func (s *DescribeVodUserBillPredictionRequest) GetArea() *string {
	return s.Area
}

func (s *DescribeVodUserBillPredictionRequest) GetDimension() *string {
	return s.Dimension
}

func (s *DescribeVodUserBillPredictionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVodUserBillPredictionRequest) SetArea(v string) *DescribeVodUserBillPredictionRequest {
	s.Area = &v
	return s
}

func (s *DescribeVodUserBillPredictionRequest) SetDimension(v string) *DescribeVodUserBillPredictionRequest {
	s.Dimension = &v
	return s
}

func (s *DescribeVodUserBillPredictionRequest) SetOwnerId(v int64) *DescribeVodUserBillPredictionRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVodUserBillPredictionRequest) Validate() error {
	return dara.Validate(s)
}
