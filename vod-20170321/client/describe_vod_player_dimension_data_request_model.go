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
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// app-1000000
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The dimension type. Valid values:
	//
	// - Os: operating system.
	//
	// - AppVersion: application version.
	//
	// - SdkVersion: SDK version.
	//
	// - Codec: codec.
	//
	// - VideoType: video format.
	//
	// - Network: network type.
	//
	// - Isp: Internet service provider.
	//
	// - VideoDefinition: resolution.
	//
	// - Domain: domain name.
	//
	// - Country: country.
	//
	// - Province: province.
	//
	// - ErrorCode: error code.
	//
	// - IsHw: whether hardware decoding is used.
	//
	// This parameter is required.
	//
	// example:
	//
	// Os
	Dimension *string `json:"Dimension,omitempty" xml:"Dimension,omitempty"`
	// The end time of the query. Specify the time in the yyyy-mm-ddthh:mm:ssz format (UTC).
	//
	// example:
	//
	// 2025-06-05T15:59:59Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The region filter used when querying the Province or Isp dimension metadata. Valid values:
	//
	// - ALL (default): all regions.
	//
	// - CN: China.
	//
	// - OVERSEAS: outside China.
	//
	// example:
	//
	// ALL
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The start time of the query. Specify the time in the <i>yyyy-mm-dd</i>t<i>hh:mm:ss</i>z format (UTC).
	//
	// >
	//
	// > - Playback data from the last year is supported.
	//
	// > - The time range for a single query cannot exceed 31 days.
	//
	// > - The time interval is left-closed and right-open [StartTime, EndTime).
	//
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
