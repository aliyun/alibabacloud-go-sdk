// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInterAuthStatisticsGlobalOverviewResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeInterAuthStatisticsGlobalOverviewResponseBodyData) *DescribeInterAuthStatisticsGlobalOverviewResponseBody
	GetData() *DescribeInterAuthStatisticsGlobalOverviewResponseBodyData
	SetRequestId(v string) *DescribeInterAuthStatisticsGlobalOverviewResponseBody
	GetRequestId() *string
}

type DescribeInterAuthStatisticsGlobalOverviewResponseBody struct {
	Data *DescribeInterAuthStatisticsGlobalOverviewResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 389DFFA3-77A5-4A9E-BF3D-147C6F98A5BA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInterAuthStatisticsGlobalOverviewResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInterAuthStatisticsGlobalOverviewResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInterAuthStatisticsGlobalOverviewResponseBody) GetData() *DescribeInterAuthStatisticsGlobalOverviewResponseBodyData {
	return s.Data
}

func (s *DescribeInterAuthStatisticsGlobalOverviewResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInterAuthStatisticsGlobalOverviewResponseBody) SetData(v *DescribeInterAuthStatisticsGlobalOverviewResponseBodyData) *DescribeInterAuthStatisticsGlobalOverviewResponseBody {
	s.Data = v
	return s
}

func (s *DescribeInterAuthStatisticsGlobalOverviewResponseBody) SetRequestId(v string) *DescribeInterAuthStatisticsGlobalOverviewResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInterAuthStatisticsGlobalOverviewResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeInterAuthStatisticsGlobalOverviewResponseBodyData struct {
	// example:
	//
	// 72
	AvgSuccessRatio *int64 `json:"AvgSuccessRatio,omitempty" xml:"AvgSuccessRatio,omitempty"`
	// example:
	//
	// 75
	AvgSuccessRatioTrend *int64 `json:"AvgSuccessRatioTrend,omitempty" xml:"AvgSuccessRatioTrend,omitempty"`
	// example:
	//
	// 18349
	TotalResolveCount *int64 `json:"TotalResolveCount,omitempty" xml:"TotalResolveCount,omitempty"`
	// example:
	//
	// 2341
	TotalResolveCountTrend *int64 `json:"TotalResolveCountTrend,omitempty" xml:"TotalResolveCountTrend,omitempty"`
}

func (s DescribeInterAuthStatisticsGlobalOverviewResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeInterAuthStatisticsGlobalOverviewResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeInterAuthStatisticsGlobalOverviewResponseBodyData) GetAvgSuccessRatio() *int64 {
	return s.AvgSuccessRatio
}

func (s *DescribeInterAuthStatisticsGlobalOverviewResponseBodyData) GetAvgSuccessRatioTrend() *int64 {
	return s.AvgSuccessRatioTrend
}

func (s *DescribeInterAuthStatisticsGlobalOverviewResponseBodyData) GetTotalResolveCount() *int64 {
	return s.TotalResolveCount
}

func (s *DescribeInterAuthStatisticsGlobalOverviewResponseBodyData) GetTotalResolveCountTrend() *int64 {
	return s.TotalResolveCountTrend
}

func (s *DescribeInterAuthStatisticsGlobalOverviewResponseBodyData) SetAvgSuccessRatio(v int64) *DescribeInterAuthStatisticsGlobalOverviewResponseBodyData {
	s.AvgSuccessRatio = &v
	return s
}

func (s *DescribeInterAuthStatisticsGlobalOverviewResponseBodyData) SetAvgSuccessRatioTrend(v int64) *DescribeInterAuthStatisticsGlobalOverviewResponseBodyData {
	s.AvgSuccessRatioTrend = &v
	return s
}

func (s *DescribeInterAuthStatisticsGlobalOverviewResponseBodyData) SetTotalResolveCount(v int64) *DescribeInterAuthStatisticsGlobalOverviewResponseBodyData {
	s.TotalResolveCount = &v
	return s
}

func (s *DescribeInterAuthStatisticsGlobalOverviewResponseBodyData) SetTotalResolveCountTrend(v int64) *DescribeInterAuthStatisticsGlobalOverviewResponseBodyData {
	s.TotalResolveCountTrend = &v
	return s
}

func (s *DescribeInterAuthStatisticsGlobalOverviewResponseBodyData) Validate() error {
	return dara.Validate(s)
}
