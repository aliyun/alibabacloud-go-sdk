// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeContainerStatisticsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeContainerStatisticsResponseBodyData) *DescribeContainerStatisticsResponseBody
	GetData() *DescribeContainerStatisticsResponseBodyData
	SetRequestId(v string) *DescribeContainerStatisticsResponseBody
	GetRequestId() *string
}

type DescribeContainerStatisticsResponseBody struct {
	// The alert statistics of container assets.
	Data *DescribeContainerStatisticsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 21DA46CA-2DCE-4FF6-907D-D5DBBB7518C8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeContainerStatisticsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeContainerStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeContainerStatisticsResponseBody) GetData() *DescribeContainerStatisticsResponseBodyData {
	return s.Data
}

func (s *DescribeContainerStatisticsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeContainerStatisticsResponseBody) SetData(v *DescribeContainerStatisticsResponseBodyData) *DescribeContainerStatisticsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeContainerStatisticsResponseBody) SetRequestId(v string) *DescribeContainerStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeContainerStatisticsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeContainerStatisticsResponseBodyData struct {
	// The number of alerts whose risk level is **Reminder**.
	//
	// example:
	//
	// 1
	RemindAlarmCount *int32 `json:"RemindAlarmCount,omitempty" xml:"RemindAlarmCount,omitempty"`
	// The number of alerts whose risk level is **Urgent**.
	//
	// example:
	//
	// 2
	SeriousAlarmCount *int32 `json:"SeriousAlarmCount,omitempty" xml:"SeriousAlarmCount,omitempty"`
	// The number of alerts whose risk level is **Suspicious**.
	//
	// example:
	//
	// 3
	SuspiciousAlarmCount *int32 `json:"SuspiciousAlarmCount,omitempty" xml:"SuspiciousAlarmCount,omitempty"`
	// The total number of alerts that are generated in the current container cluster.
	//
	// example:
	//
	// 6
	TotalAlarmCount *int32 `json:"TotalAlarmCount,omitempty" xml:"TotalAlarmCount,omitempty"`
	// The total number of nodes in the current container cluster.
	//
	// example:
	//
	// 12
	TotalNode *int32 `json:"TotalNode,omitempty" xml:"TotalNode,omitempty"`
	// The number of nodes on which alerts are generated in the current container cluster.
	//
	// example:
	//
	// 4
	HasRiskNode *int32 `json:"hasRiskNode,omitempty" xml:"hasRiskNode,omitempty"`
}

func (s DescribeContainerStatisticsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeContainerStatisticsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeContainerStatisticsResponseBodyData) GetRemindAlarmCount() *int32 {
	return s.RemindAlarmCount
}

func (s *DescribeContainerStatisticsResponseBodyData) GetSeriousAlarmCount() *int32 {
	return s.SeriousAlarmCount
}

func (s *DescribeContainerStatisticsResponseBodyData) GetSuspiciousAlarmCount() *int32 {
	return s.SuspiciousAlarmCount
}

func (s *DescribeContainerStatisticsResponseBodyData) GetTotalAlarmCount() *int32 {
	return s.TotalAlarmCount
}

func (s *DescribeContainerStatisticsResponseBodyData) GetTotalNode() *int32 {
	return s.TotalNode
}

func (s *DescribeContainerStatisticsResponseBodyData) GetHasRiskNode() *int32 {
	return s.HasRiskNode
}

func (s *DescribeContainerStatisticsResponseBodyData) SetRemindAlarmCount(v int32) *DescribeContainerStatisticsResponseBodyData {
	s.RemindAlarmCount = &v
	return s
}

func (s *DescribeContainerStatisticsResponseBodyData) SetSeriousAlarmCount(v int32) *DescribeContainerStatisticsResponseBodyData {
	s.SeriousAlarmCount = &v
	return s
}

func (s *DescribeContainerStatisticsResponseBodyData) SetSuspiciousAlarmCount(v int32) *DescribeContainerStatisticsResponseBodyData {
	s.SuspiciousAlarmCount = &v
	return s
}

func (s *DescribeContainerStatisticsResponseBodyData) SetTotalAlarmCount(v int32) *DescribeContainerStatisticsResponseBodyData {
	s.TotalAlarmCount = &v
	return s
}

func (s *DescribeContainerStatisticsResponseBodyData) SetTotalNode(v int32) *DescribeContainerStatisticsResponseBodyData {
	s.TotalNode = &v
	return s
}

func (s *DescribeContainerStatisticsResponseBodyData) SetHasRiskNode(v int32) *DescribeContainerStatisticsResponseBodyData {
	s.HasRiskNode = &v
	return s
}

func (s *DescribeContainerStatisticsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
