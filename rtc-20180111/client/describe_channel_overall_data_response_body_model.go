// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeChannelOverallDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCallInfo(v *DescribeChannelOverallDataResponseBodyCallInfo) *DescribeChannelOverallDataResponseBody
	GetCallInfo() *DescribeChannelOverallDataResponseBodyCallInfo
	SetMetricDatas(v []*DescribeChannelOverallDataResponseBodyMetricDatas) *DescribeChannelOverallDataResponseBody
	GetMetricDatas() []*DescribeChannelOverallDataResponseBodyMetricDatas
	SetOverallData(v *DescribeChannelOverallDataResponseBodyOverallData) *DescribeChannelOverallDataResponseBody
	GetOverallData() *DescribeChannelOverallDataResponseBodyOverallData
	SetRequestId(v string) *DescribeChannelOverallDataResponseBody
	GetRequestId() *string
}

type DescribeChannelOverallDataResponseBody struct {
	CallInfo    *DescribeChannelOverallDataResponseBodyCallInfo      `json:"CallInfo,omitempty" xml:"CallInfo,omitempty" type:"Struct"`
	MetricDatas []*DescribeChannelOverallDataResponseBodyMetricDatas `json:"MetricDatas,omitempty" xml:"MetricDatas,omitempty" type:"Repeated"`
	OverallData *DescribeChannelOverallDataResponseBodyOverallData   `json:"OverallData,omitempty" xml:"OverallData,omitempty" type:"Struct"`
	// example:
	//
	// 231470C1-ACFB-4C9F-844F-4CFE1E3804C5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeChannelOverallDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeChannelOverallDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeChannelOverallDataResponseBody) GetCallInfo() *DescribeChannelOverallDataResponseBodyCallInfo {
	return s.CallInfo
}

func (s *DescribeChannelOverallDataResponseBody) GetMetricDatas() []*DescribeChannelOverallDataResponseBodyMetricDatas {
	return s.MetricDatas
}

func (s *DescribeChannelOverallDataResponseBody) GetOverallData() *DescribeChannelOverallDataResponseBodyOverallData {
	return s.OverallData
}

func (s *DescribeChannelOverallDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeChannelOverallDataResponseBody) SetCallInfo(v *DescribeChannelOverallDataResponseBodyCallInfo) *DescribeChannelOverallDataResponseBody {
	s.CallInfo = v
	return s
}

func (s *DescribeChannelOverallDataResponseBody) SetMetricDatas(v []*DescribeChannelOverallDataResponseBodyMetricDatas) *DescribeChannelOverallDataResponseBody {
	s.MetricDatas = v
	return s
}

func (s *DescribeChannelOverallDataResponseBody) SetOverallData(v *DescribeChannelOverallDataResponseBodyOverallData) *DescribeChannelOverallDataResponseBody {
	s.OverallData = v
	return s
}

func (s *DescribeChannelOverallDataResponseBody) SetRequestId(v string) *DescribeChannelOverallDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeChannelOverallDataResponseBody) Validate() error {
	if s.CallInfo != nil {
		if err := s.CallInfo.Validate(); err != nil {
			return err
		}
	}
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

type DescribeChannelOverallDataResponseBodyCallInfo struct {
	// example:
	//
	// rjdhtnqy
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// IN
	CallStatus *string `json:"CallStatus,omitempty" xml:"CallStatus,omitempty"`
	// example:
	//
	// 123456
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// example:
	//
	// 1615860711
	CreatedTs *int64 `json:"CreatedTs,omitempty" xml:"CreatedTs,omitempty"`
	// example:
	//
	// 1615860811
	DestroyedTs *int64 `json:"DestroyedTs,omitempty" xml:"DestroyedTs,omitempty"`
	// example:
	//
	// 100
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
}

func (s DescribeChannelOverallDataResponseBodyCallInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeChannelOverallDataResponseBodyCallInfo) GoString() string {
	return s.String()
}

func (s *DescribeChannelOverallDataResponseBodyCallInfo) GetAppId() *string {
	return s.AppId
}

func (s *DescribeChannelOverallDataResponseBodyCallInfo) GetCallStatus() *string {
	return s.CallStatus
}

func (s *DescribeChannelOverallDataResponseBodyCallInfo) GetChannelId() *string {
	return s.ChannelId
}

func (s *DescribeChannelOverallDataResponseBodyCallInfo) GetCreatedTs() *int64 {
	return s.CreatedTs
}

func (s *DescribeChannelOverallDataResponseBodyCallInfo) GetDestroyedTs() *int64 {
	return s.DestroyedTs
}

func (s *DescribeChannelOverallDataResponseBodyCallInfo) GetDuration() *int64 {
	return s.Duration
}

func (s *DescribeChannelOverallDataResponseBodyCallInfo) SetAppId(v string) *DescribeChannelOverallDataResponseBodyCallInfo {
	s.AppId = &v
	return s
}

func (s *DescribeChannelOverallDataResponseBodyCallInfo) SetCallStatus(v string) *DescribeChannelOverallDataResponseBodyCallInfo {
	s.CallStatus = &v
	return s
}

func (s *DescribeChannelOverallDataResponseBodyCallInfo) SetChannelId(v string) *DescribeChannelOverallDataResponseBodyCallInfo {
	s.ChannelId = &v
	return s
}

func (s *DescribeChannelOverallDataResponseBodyCallInfo) SetCreatedTs(v int64) *DescribeChannelOverallDataResponseBodyCallInfo {
	s.CreatedTs = &v
	return s
}

func (s *DescribeChannelOverallDataResponseBodyCallInfo) SetDestroyedTs(v int64) *DescribeChannelOverallDataResponseBodyCallInfo {
	s.DestroyedTs = &v
	return s
}

func (s *DescribeChannelOverallDataResponseBodyCallInfo) SetDuration(v int64) *DescribeChannelOverallDataResponseBodyCallInfo {
	s.Duration = &v
	return s
}

func (s *DescribeChannelOverallDataResponseBodyCallInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeChannelOverallDataResponseBodyMetricDatas struct {
	Nodes []*DescribeChannelOverallDataResponseBodyMetricDatasNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	// example:
	//
	// CALL_QUALITY
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeChannelOverallDataResponseBodyMetricDatas) String() string {
	return dara.Prettify(s)
}

func (s DescribeChannelOverallDataResponseBodyMetricDatas) GoString() string {
	return s.String()
}

func (s *DescribeChannelOverallDataResponseBodyMetricDatas) GetNodes() []*DescribeChannelOverallDataResponseBodyMetricDatasNodes {
	return s.Nodes
}

func (s *DescribeChannelOverallDataResponseBodyMetricDatas) GetType() *string {
	return s.Type
}

func (s *DescribeChannelOverallDataResponseBodyMetricDatas) SetNodes(v []*DescribeChannelOverallDataResponseBodyMetricDatasNodes) *DescribeChannelOverallDataResponseBodyMetricDatas {
	s.Nodes = v
	return s
}

func (s *DescribeChannelOverallDataResponseBodyMetricDatas) SetType(v string) *DescribeChannelOverallDataResponseBodyMetricDatas {
	s.Type = &v
	return s
}

func (s *DescribeChannelOverallDataResponseBodyMetricDatas) Validate() error {
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

type DescribeChannelOverallDataResponseBodyMetricDatasNodes struct {
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

func (s DescribeChannelOverallDataResponseBodyMetricDatasNodes) String() string {
	return dara.Prettify(s)
}

func (s DescribeChannelOverallDataResponseBodyMetricDatasNodes) GoString() string {
	return s.String()
}

func (s *DescribeChannelOverallDataResponseBodyMetricDatasNodes) GetExt() map[string]interface{} {
	return s.Ext
}

func (s *DescribeChannelOverallDataResponseBodyMetricDatasNodes) GetX() *string {
	return s.X
}

func (s *DescribeChannelOverallDataResponseBodyMetricDatasNodes) GetY() *string {
	return s.Y
}

func (s *DescribeChannelOverallDataResponseBodyMetricDatasNodes) SetExt(v map[string]interface{}) *DescribeChannelOverallDataResponseBodyMetricDatasNodes {
	s.Ext = v
	return s
}

func (s *DescribeChannelOverallDataResponseBodyMetricDatasNodes) SetX(v string) *DescribeChannelOverallDataResponseBodyMetricDatasNodes {
	s.X = &v
	return s
}

func (s *DescribeChannelOverallDataResponseBodyMetricDatasNodes) SetY(v string) *DescribeChannelOverallDataResponseBodyMetricDatasNodes {
	s.Y = &v
	return s
}

func (s *DescribeChannelOverallDataResponseBodyMetricDatasNodes) Validate() error {
	return dara.Validate(s)
}

type DescribeChannelOverallDataResponseBodyOverallData struct {
	// example:
	//
	// 0.5
	ConnAvgTime *float32 `json:"ConnAvgTime,omitempty" xml:"ConnAvgTime,omitempty"`
	// example:
	//
	// 0.91
	FiveSecJoinRate *float32 `json:"FiveSecJoinRate,omitempty" xml:"FiveSecJoinRate,omitempty"`
	// example:
	//
	// 0.02
	TotalAudioStuckRate *float32 `json:"TotalAudioStuckRate,omitempty" xml:"TotalAudioStuckRate,omitempty"`
	// example:
	//
	// 0.02
	TotalVideoStuckRate *float32 `json:"TotalVideoStuckRate,omitempty" xml:"TotalVideoStuckRate,omitempty"`
	// example:
	//
	// 0.02
	TotalVideoVagueRate *float32 `json:"TotalVideoVagueRate,omitempty" xml:"TotalVideoVagueRate,omitempty"`
}

func (s DescribeChannelOverallDataResponseBodyOverallData) String() string {
	return dara.Prettify(s)
}

func (s DescribeChannelOverallDataResponseBodyOverallData) GoString() string {
	return s.String()
}

func (s *DescribeChannelOverallDataResponseBodyOverallData) GetConnAvgTime() *float32 {
	return s.ConnAvgTime
}

func (s *DescribeChannelOverallDataResponseBodyOverallData) GetFiveSecJoinRate() *float32 {
	return s.FiveSecJoinRate
}

func (s *DescribeChannelOverallDataResponseBodyOverallData) GetTotalAudioStuckRate() *float32 {
	return s.TotalAudioStuckRate
}

func (s *DescribeChannelOverallDataResponseBodyOverallData) GetTotalVideoStuckRate() *float32 {
	return s.TotalVideoStuckRate
}

func (s *DescribeChannelOverallDataResponseBodyOverallData) GetTotalVideoVagueRate() *float32 {
	return s.TotalVideoVagueRate
}

func (s *DescribeChannelOverallDataResponseBodyOverallData) SetConnAvgTime(v float32) *DescribeChannelOverallDataResponseBodyOverallData {
	s.ConnAvgTime = &v
	return s
}

func (s *DescribeChannelOverallDataResponseBodyOverallData) SetFiveSecJoinRate(v float32) *DescribeChannelOverallDataResponseBodyOverallData {
	s.FiveSecJoinRate = &v
	return s
}

func (s *DescribeChannelOverallDataResponseBodyOverallData) SetTotalAudioStuckRate(v float32) *DescribeChannelOverallDataResponseBodyOverallData {
	s.TotalAudioStuckRate = &v
	return s
}

func (s *DescribeChannelOverallDataResponseBodyOverallData) SetTotalVideoStuckRate(v float32) *DescribeChannelOverallDataResponseBodyOverallData {
	s.TotalVideoStuckRate = &v
	return s
}

func (s *DescribeChannelOverallDataResponseBodyOverallData) SetTotalVideoVagueRate(v float32) *DescribeChannelOverallDataResponseBodyOverallData {
	s.TotalVideoVagueRate = &v
	return s
}

func (s *DescribeChannelOverallDataResponseBodyOverallData) Validate() error {
	return dara.Validate(s)
}
