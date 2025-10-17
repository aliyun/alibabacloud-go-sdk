// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeChannelUserMetricsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMetricDatas(v []*DescribeChannelUserMetricsResponseBodyMetricDatas) *DescribeChannelUserMetricsResponseBody
	GetMetricDatas() []*DescribeChannelUserMetricsResponseBodyMetricDatas
	SetOverallData(v *DescribeChannelUserMetricsResponseBodyOverallData) *DescribeChannelUserMetricsResponseBody
	GetOverallData() *DescribeChannelUserMetricsResponseBodyOverallData
	SetRequestId(v string) *DescribeChannelUserMetricsResponseBody
	GetRequestId() *string
}

type DescribeChannelUserMetricsResponseBody struct {
	MetricDatas []*DescribeChannelUserMetricsResponseBodyMetricDatas `json:"MetricDatas,omitempty" xml:"MetricDatas,omitempty" type:"Repeated"`
	OverallData *DescribeChannelUserMetricsResponseBodyOverallData   `json:"OverallData,omitempty" xml:"OverallData,omitempty" type:"Struct"`
	// example:
	//
	// 231470C1-ACFB-4C9F-844F-4CFE1E3804C5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeChannelUserMetricsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeChannelUserMetricsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeChannelUserMetricsResponseBody) GetMetricDatas() []*DescribeChannelUserMetricsResponseBodyMetricDatas {
	return s.MetricDatas
}

func (s *DescribeChannelUserMetricsResponseBody) GetOverallData() *DescribeChannelUserMetricsResponseBodyOverallData {
	return s.OverallData
}

func (s *DescribeChannelUserMetricsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeChannelUserMetricsResponseBody) SetMetricDatas(v []*DescribeChannelUserMetricsResponseBodyMetricDatas) *DescribeChannelUserMetricsResponseBody {
	s.MetricDatas = v
	return s
}

func (s *DescribeChannelUserMetricsResponseBody) SetOverallData(v *DescribeChannelUserMetricsResponseBodyOverallData) *DescribeChannelUserMetricsResponseBody {
	s.OverallData = v
	return s
}

func (s *DescribeChannelUserMetricsResponseBody) SetRequestId(v string) *DescribeChannelUserMetricsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeChannelUserMetricsResponseBody) Validate() error {
	if s.MetricDatas != nil {
		for _, item := range s.MetricDatas {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.OverallData != nil {
		if err := s.OverallData.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeChannelUserMetricsResponseBodyMetricDatas struct {
	Nodes []*DescribeChannelUserMetricsResponseBodyMetricDatasNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	// example:
	//
	// ALL_NUM
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeChannelUserMetricsResponseBodyMetricDatas) String() string {
	return dara.Prettify(s)
}

func (s DescribeChannelUserMetricsResponseBodyMetricDatas) GoString() string {
	return s.String()
}

func (s *DescribeChannelUserMetricsResponseBodyMetricDatas) GetNodes() []*DescribeChannelUserMetricsResponseBodyMetricDatasNodes {
	return s.Nodes
}

func (s *DescribeChannelUserMetricsResponseBodyMetricDatas) GetType() *string {
	return s.Type
}

func (s *DescribeChannelUserMetricsResponseBodyMetricDatas) SetNodes(v []*DescribeChannelUserMetricsResponseBodyMetricDatasNodes) *DescribeChannelUserMetricsResponseBodyMetricDatas {
	s.Nodes = v
	return s
}

func (s *DescribeChannelUserMetricsResponseBodyMetricDatas) SetType(v string) *DescribeChannelUserMetricsResponseBodyMetricDatas {
	s.Type = &v
	return s
}

func (s *DescribeChannelUserMetricsResponseBodyMetricDatas) Validate() error {
	if s.Nodes != nil {
		for _, item := range s.Nodes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeChannelUserMetricsResponseBodyMetricDatasNodes struct {
	Ext map[string]interface{} `json:"Ext,omitempty" xml:"Ext,omitempty"`
	// example:
	//
	// 1612418625
	X *string `json:"X,omitempty" xml:"X,omitempty"`
	// example:
	//
	// 123
	Y *string `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s DescribeChannelUserMetricsResponseBodyMetricDatasNodes) String() string {
	return dara.Prettify(s)
}

func (s DescribeChannelUserMetricsResponseBodyMetricDatasNodes) GoString() string {
	return s.String()
}

func (s *DescribeChannelUserMetricsResponseBodyMetricDatasNodes) GetExt() map[string]interface{} {
	return s.Ext
}

func (s *DescribeChannelUserMetricsResponseBodyMetricDatasNodes) GetX() *string {
	return s.X
}

func (s *DescribeChannelUserMetricsResponseBodyMetricDatasNodes) GetY() *string {
	return s.Y
}

func (s *DescribeChannelUserMetricsResponseBodyMetricDatasNodes) SetExt(v map[string]interface{}) *DescribeChannelUserMetricsResponseBodyMetricDatasNodes {
	s.Ext = v
	return s
}

func (s *DescribeChannelUserMetricsResponseBodyMetricDatasNodes) SetX(v string) *DescribeChannelUserMetricsResponseBodyMetricDatasNodes {
	s.X = &v
	return s
}

func (s *DescribeChannelUserMetricsResponseBodyMetricDatasNodes) SetY(v string) *DescribeChannelUserMetricsResponseBodyMetricDatasNodes {
	s.Y = &v
	return s
}

func (s *DescribeChannelUserMetricsResponseBodyMetricDatasNodes) Validate() error {
	return dara.Validate(s)
}

type DescribeChannelUserMetricsResponseBodyOverallData struct {
	// example:
	//
	// 0
	TotalBadExpNum *int64 `json:"TotalBadExpNum,omitempty" xml:"TotalBadExpNum,omitempty"`
	// example:
	//
	// 0
	TotalJoinFailNum *int64 `json:"TotalJoinFailNum,omitempty" xml:"TotalJoinFailNum,omitempty"`
	// example:
	//
	// 1
	TotalPubUserNum *int64 `json:"TotalPubUserNum,omitempty" xml:"TotalPubUserNum,omitempty"`
	// example:
	//
	// 3
	TotalSubUserNum *int64 `json:"TotalSubUserNum,omitempty" xml:"TotalSubUserNum,omitempty"`
	// example:
	//
	// 5
	TotalUserNum *int64 `json:"TotalUserNum,omitempty" xml:"TotalUserNum,omitempty"`
}

func (s DescribeChannelUserMetricsResponseBodyOverallData) String() string {
	return dara.Prettify(s)
}

func (s DescribeChannelUserMetricsResponseBodyOverallData) GoString() string {
	return s.String()
}

func (s *DescribeChannelUserMetricsResponseBodyOverallData) GetTotalBadExpNum() *int64 {
	return s.TotalBadExpNum
}

func (s *DescribeChannelUserMetricsResponseBodyOverallData) GetTotalJoinFailNum() *int64 {
	return s.TotalJoinFailNum
}

func (s *DescribeChannelUserMetricsResponseBodyOverallData) GetTotalPubUserNum() *int64 {
	return s.TotalPubUserNum
}

func (s *DescribeChannelUserMetricsResponseBodyOverallData) GetTotalSubUserNum() *int64 {
	return s.TotalSubUserNum
}

func (s *DescribeChannelUserMetricsResponseBodyOverallData) GetTotalUserNum() *int64 {
	return s.TotalUserNum
}

func (s *DescribeChannelUserMetricsResponseBodyOverallData) SetTotalBadExpNum(v int64) *DescribeChannelUserMetricsResponseBodyOverallData {
	s.TotalBadExpNum = &v
	return s
}

func (s *DescribeChannelUserMetricsResponseBodyOverallData) SetTotalJoinFailNum(v int64) *DescribeChannelUserMetricsResponseBodyOverallData {
	s.TotalJoinFailNum = &v
	return s
}

func (s *DescribeChannelUserMetricsResponseBodyOverallData) SetTotalPubUserNum(v int64) *DescribeChannelUserMetricsResponseBodyOverallData {
	s.TotalPubUserNum = &v
	return s
}

func (s *DescribeChannelUserMetricsResponseBodyOverallData) SetTotalSubUserNum(v int64) *DescribeChannelUserMetricsResponseBodyOverallData {
	s.TotalSubUserNum = &v
	return s
}

func (s *DescribeChannelUserMetricsResponseBodyOverallData) SetTotalUserNum(v int64) *DescribeChannelUserMetricsResponseBodyOverallData {
	s.TotalUserNum = &v
	return s
}

func (s *DescribeChannelUserMetricsResponseBodyOverallData) Validate() error {
	return dara.Validate(s)
}
