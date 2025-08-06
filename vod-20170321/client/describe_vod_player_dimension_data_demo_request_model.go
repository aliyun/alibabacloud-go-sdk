// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodPlayerDimensionDataDemoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeVodPlayerDimensionDataDemoRequest
	GetAppId() *string
	SetDimension(v string) *DescribeVodPlayerDimensionDataDemoRequest
	GetDimension() *string
}

type DescribeVodPlayerDimensionDataDemoRequest struct {
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
}

func (s DescribeVodPlayerDimensionDataDemoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodPlayerDimensionDataDemoRequest) GoString() string {
	return s.String()
}

func (s *DescribeVodPlayerDimensionDataDemoRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeVodPlayerDimensionDataDemoRequest) GetDimension() *string {
	return s.Dimension
}

func (s *DescribeVodPlayerDimensionDataDemoRequest) SetAppId(v string) *DescribeVodPlayerDimensionDataDemoRequest {
	s.AppId = &v
	return s
}

func (s *DescribeVodPlayerDimensionDataDemoRequest) SetDimension(v string) *DescribeVodPlayerDimensionDataDemoRequest {
	s.Dimension = &v
	return s
}

func (s *DescribeVodPlayerDimensionDataDemoRequest) Validate() error {
	return dara.Validate(s)
}
