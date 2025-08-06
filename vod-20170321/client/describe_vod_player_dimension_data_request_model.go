// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodPlayerDimensionDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeVodPlayerDimensionDataRequest
	GetAppId() *string
	SetDimension(v string) *DescribeVodPlayerDimensionDataRequest
	GetDimension() *string
	SetEndTime(v string) *DescribeVodPlayerDimensionDataRequest
	GetEndTime() *string
	SetRegion(v string) *DescribeVodPlayerDimensionDataRequest
	GetRegion() *string
	SetStartTime(v string) *DescribeVodPlayerDimensionDataRequest
	GetStartTime() *string
}

type DescribeVodPlayerDimensionDataRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// app-1000000
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Os
	Dimension *string `json:"Dimension,omitempty" xml:"Dimension,omitempty"`
	// example:
	//
	// 2025-06-05T15:59:59Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// ALL
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// example:
	//
	// 2025-06-24T00:55:06Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeVodPlayerDimensionDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodPlayerDimensionDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeVodPlayerDimensionDataRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeVodPlayerDimensionDataRequest) GetDimension() *string {
	return s.Dimension
}

func (s *DescribeVodPlayerDimensionDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVodPlayerDimensionDataRequest) GetRegion() *string {
	return s.Region
}

func (s *DescribeVodPlayerDimensionDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVodPlayerDimensionDataRequest) SetAppId(v string) *DescribeVodPlayerDimensionDataRequest {
	s.AppId = &v
	return s
}

func (s *DescribeVodPlayerDimensionDataRequest) SetDimension(v string) *DescribeVodPlayerDimensionDataRequest {
	s.Dimension = &v
	return s
}

func (s *DescribeVodPlayerDimensionDataRequest) SetEndTime(v string) *DescribeVodPlayerDimensionDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeVodPlayerDimensionDataRequest) SetRegion(v string) *DescribeVodPlayerDimensionDataRequest {
	s.Region = &v
	return s
}

func (s *DescribeVodPlayerDimensionDataRequest) SetStartTime(v string) *DescribeVodPlayerDimensionDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeVodPlayerDimensionDataRequest) Validate() error {
	return dara.Validate(s)
}
