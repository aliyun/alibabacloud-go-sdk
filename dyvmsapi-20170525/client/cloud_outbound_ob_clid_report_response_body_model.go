// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudOutboundObClidReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CloudOutboundObClidReportResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CloudOutboundObClidReportResponseBody
	GetCode() *string
	SetData(v *CloudOutboundObClidReportResponseBodyData) *CloudOutboundObClidReportResponseBody
	GetData() *CloudOutboundObClidReportResponseBodyData
	SetMessage(v string) *CloudOutboundObClidReportResponseBody
	GetMessage() *string
	SetRequestId(v string) *CloudOutboundObClidReportResponseBody
	GetRequestId() *string
}

type CloudOutboundObClidReportResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CloudOutboundObClidReportResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 9FF70B74-1B3C-44C1-ACDF-8DF962988F0E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloudOutboundObClidReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloudOutboundObClidReportResponseBody) GoString() string {
	return s.String()
}

func (s *CloudOutboundObClidReportResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CloudOutboundObClidReportResponseBody) GetCode() *string {
	return s.Code
}

func (s *CloudOutboundObClidReportResponseBody) GetData() *CloudOutboundObClidReportResponseBodyData {
	return s.Data
}

func (s *CloudOutboundObClidReportResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CloudOutboundObClidReportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloudOutboundObClidReportResponseBody) SetAccessDeniedDetail(v string) *CloudOutboundObClidReportResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CloudOutboundObClidReportResponseBody) SetCode(v string) *CloudOutboundObClidReportResponseBody {
	s.Code = &v
	return s
}

func (s *CloudOutboundObClidReportResponseBody) SetData(v *CloudOutboundObClidReportResponseBodyData) *CloudOutboundObClidReportResponseBody {
	s.Data = v
	return s
}

func (s *CloudOutboundObClidReportResponseBody) SetMessage(v string) *CloudOutboundObClidReportResponseBody {
	s.Message = &v
	return s
}

func (s *CloudOutboundObClidReportResponseBody) SetRequestId(v string) *CloudOutboundObClidReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloudOutboundObClidReportResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CloudOutboundObClidReportResponseBodyData struct {
	// 返回数据
	List []*CloudOutboundObClidReportResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
}

func (s CloudOutboundObClidReportResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CloudOutboundObClidReportResponseBodyData) GoString() string {
	return s.String()
}

func (s *CloudOutboundObClidReportResponseBodyData) GetList() []*CloudOutboundObClidReportResponseBodyDataList {
	return s.List
}

func (s *CloudOutboundObClidReportResponseBodyData) SetList(v []*CloudOutboundObClidReportResponseBodyDataList) *CloudOutboundObClidReportResponseBodyData {
	s.List = v
	return s
}

func (s *CloudOutboundObClidReportResponseBodyData) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CloudOutboundObClidReportResponseBodyDataList struct {
	// 客户接听数
	//
	// example:
	//
	// 4
	AnsweredCount *string `json:"AnsweredCount,omitempty" xml:"AnsweredCount,omitempty"`
	// 平均通话时长
	//
	// example:
	//
	// 00:02:12
	AvgBridgeTime *string `json:"AvgBridgeTime,omitempty" xml:"AvgBridgeTime,omitempty"`
	// 平均等待时长
	//
	// example:
	//
	// 5
	AvgPreviewObWaitTime *int64 `json:"AvgPreviewObWaitTime,omitempty" xml:"AvgPreviewObWaitTime,omitempty"`
	// example:
	//
	// 示例值示例值
	CallType *string `json:"CallType,omitempty" xml:"CallType,omitempty"`
	// 城市
	//
	// example:
	//
	// 示例值示例值
	City *string `json:"City,omitempty" xml:"City,omitempty"`
	// 外显号码
	//
	// example:
	//
	// 01089192349
	Clid *string `json:"Clid,omitempty" xml:"Clid,omitempty"`
	// 号码接听率
	//
	// example:
	//
	// 20%
	ClidRate *string `json:"ClidRate,omitempty" xml:"ClidRate,omitempty"`
	// 创建日期
	//
	// example:
	//
	// 2026-04-20 10:00:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 客户侧计费分钟数
	//
	// example:
	//
	// 45
	CustomerBillMinute *int64 `json:"CustomerBillMinute,omitempty" xml:"CustomerBillMinute,omitempty"`
	// 时间范围
	//
	// example:
	//
	// 示例值示例值示例值
	DateTimeRange *string `json:"DateTimeRange,omitempty" xml:"DateTimeRange,omitempty"`
	// 报表日期
	//
	// example:
	//
	// 2026-04-20
	Day *string `json:"Day,omitempty" xml:"Day,omitempty"`
	// 企业Id
	//
	// example:
	//
	// 7000002
	EnterpriseId *int64 `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	// 小时
	//
	// example:
	//
	// 0
	Hour *string `json:"Hour,omitempty" xml:"Hour,omitempty"`
	// 最长通话时长
	//
	// example:
	//
	// 00:00:45
	MaxBridgeTime *string `json:"MaxBridgeTime,omitempty" xml:"MaxBridgeTime,omitempty"`
	// 最短通话时长
	//
	// example:
	//
	// 00:00:03
	MinBridgeTime *string `json:"MinBridgeTime,omitempty" xml:"MinBridgeTime,omitempty"`
	// 响铃数
	//
	// example:
	//
	// 19
	PreviewObCustomerRingingCount *int64 `json:"PreviewObCustomerRingingCount,omitempty" xml:"PreviewObCustomerRingingCount,omitempty"`
	// 响铃率
	//
	// example:
	//
	// 95%
	PreviewObCustomerRingingRate *string `json:"PreviewObCustomerRingingRate,omitempty" xml:"PreviewObCustomerRingingRate,omitempty"`
	// 省份
	//
	// example:
	//
	// 示例值示例值
	Province *string `json:"Province,omitempty" xml:"Province,omitempty"`
	// 格式化日期
	//
	// example:
	//
	// 示例值
	RowName *string `json:"RowName,omitempty" xml:"RowName,omitempty"`
	// 总通话时长
	//
	// example:
	//
	// 00:08:48
	TotalBridgeTime *string `json:"TotalBridgeTime,omitempty" xml:"TotalBridgeTime,omitempty"`
	// 外呼总数
	//
	// example:
	//
	// 20
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// 中继群组代号
	//
	// example:
	//
	// 10000
	TrunkGroupKey *string `json:"TrunkGroupKey,omitempty" xml:"TrunkGroupKey,omitempty"`
	// 有效通话平均时长
	//
	// example:
	//
	// 00:00:21
	ValidAvgBridgeTime *string `json:"ValidAvgBridgeTime,omitempty" xml:"ValidAvgBridgeTime,omitempty"`
	// 有效通话次数
	//
	// example:
	//
	// 2
	ValidCalls *string `json:"ValidCalls,omitempty" xml:"ValidCalls,omitempty"`
	// 有效通话总时长
	//
	// example:
	//
	// 00:00:42
	ValidTotalBridgeTime *string `json:"ValidTotalBridgeTime,omitempty" xml:"ValidTotalBridgeTime,omitempty"`
}

func (s CloudOutboundObClidReportResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s CloudOutboundObClidReportResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *CloudOutboundObClidReportResponseBodyDataList) GetAnsweredCount() *string {
	return s.AnsweredCount
}

func (s *CloudOutboundObClidReportResponseBodyDataList) GetAvgBridgeTime() *string {
	return s.AvgBridgeTime
}

func (s *CloudOutboundObClidReportResponseBodyDataList) GetAvgPreviewObWaitTime() *int64 {
	return s.AvgPreviewObWaitTime
}

func (s *CloudOutboundObClidReportResponseBodyDataList) GetCallType() *string {
	return s.CallType
}

func (s *CloudOutboundObClidReportResponseBodyDataList) GetCity() *string {
	return s.City
}

func (s *CloudOutboundObClidReportResponseBodyDataList) GetClid() *string {
	return s.Clid
}

func (s *CloudOutboundObClidReportResponseBodyDataList) GetClidRate() *string {
	return s.ClidRate
}

func (s *CloudOutboundObClidReportResponseBodyDataList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CloudOutboundObClidReportResponseBodyDataList) GetCustomerBillMinute() *int64 {
	return s.CustomerBillMinute
}

func (s *CloudOutboundObClidReportResponseBodyDataList) GetDateTimeRange() *string {
	return s.DateTimeRange
}

func (s *CloudOutboundObClidReportResponseBodyDataList) GetDay() *string {
	return s.Day
}

func (s *CloudOutboundObClidReportResponseBodyDataList) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *CloudOutboundObClidReportResponseBodyDataList) GetHour() *string {
	return s.Hour
}

func (s *CloudOutboundObClidReportResponseBodyDataList) GetMaxBridgeTime() *string {
	return s.MaxBridgeTime
}

func (s *CloudOutboundObClidReportResponseBodyDataList) GetMinBridgeTime() *string {
	return s.MinBridgeTime
}

func (s *CloudOutboundObClidReportResponseBodyDataList) GetPreviewObCustomerRingingCount() *int64 {
	return s.PreviewObCustomerRingingCount
}

func (s *CloudOutboundObClidReportResponseBodyDataList) GetPreviewObCustomerRingingRate() *string {
	return s.PreviewObCustomerRingingRate
}

func (s *CloudOutboundObClidReportResponseBodyDataList) GetProvince() *string {
	return s.Province
}

func (s *CloudOutboundObClidReportResponseBodyDataList) GetRowName() *string {
	return s.RowName
}

func (s *CloudOutboundObClidReportResponseBodyDataList) GetTotalBridgeTime() *string {
	return s.TotalBridgeTime
}

func (s *CloudOutboundObClidReportResponseBodyDataList) GetTotalCount() *string {
	return s.TotalCount
}

func (s *CloudOutboundObClidReportResponseBodyDataList) GetTrunkGroupKey() *string {
	return s.TrunkGroupKey
}

func (s *CloudOutboundObClidReportResponseBodyDataList) GetValidAvgBridgeTime() *string {
	return s.ValidAvgBridgeTime
}

func (s *CloudOutboundObClidReportResponseBodyDataList) GetValidCalls() *string {
	return s.ValidCalls
}

func (s *CloudOutboundObClidReportResponseBodyDataList) GetValidTotalBridgeTime() *string {
	return s.ValidTotalBridgeTime
}

func (s *CloudOutboundObClidReportResponseBodyDataList) SetAnsweredCount(v string) *CloudOutboundObClidReportResponseBodyDataList {
	s.AnsweredCount = &v
	return s
}

func (s *CloudOutboundObClidReportResponseBodyDataList) SetAvgBridgeTime(v string) *CloudOutboundObClidReportResponseBodyDataList {
	s.AvgBridgeTime = &v
	return s
}

func (s *CloudOutboundObClidReportResponseBodyDataList) SetAvgPreviewObWaitTime(v int64) *CloudOutboundObClidReportResponseBodyDataList {
	s.AvgPreviewObWaitTime = &v
	return s
}

func (s *CloudOutboundObClidReportResponseBodyDataList) SetCallType(v string) *CloudOutboundObClidReportResponseBodyDataList {
	s.CallType = &v
	return s
}

func (s *CloudOutboundObClidReportResponseBodyDataList) SetCity(v string) *CloudOutboundObClidReportResponseBodyDataList {
	s.City = &v
	return s
}

func (s *CloudOutboundObClidReportResponseBodyDataList) SetClid(v string) *CloudOutboundObClidReportResponseBodyDataList {
	s.Clid = &v
	return s
}

func (s *CloudOutboundObClidReportResponseBodyDataList) SetClidRate(v string) *CloudOutboundObClidReportResponseBodyDataList {
	s.ClidRate = &v
	return s
}

func (s *CloudOutboundObClidReportResponseBodyDataList) SetCreateTime(v string) *CloudOutboundObClidReportResponseBodyDataList {
	s.CreateTime = &v
	return s
}

func (s *CloudOutboundObClidReportResponseBodyDataList) SetCustomerBillMinute(v int64) *CloudOutboundObClidReportResponseBodyDataList {
	s.CustomerBillMinute = &v
	return s
}

func (s *CloudOutboundObClidReportResponseBodyDataList) SetDateTimeRange(v string) *CloudOutboundObClidReportResponseBodyDataList {
	s.DateTimeRange = &v
	return s
}

func (s *CloudOutboundObClidReportResponseBodyDataList) SetDay(v string) *CloudOutboundObClidReportResponseBodyDataList {
	s.Day = &v
	return s
}

func (s *CloudOutboundObClidReportResponseBodyDataList) SetEnterpriseId(v int64) *CloudOutboundObClidReportResponseBodyDataList {
	s.EnterpriseId = &v
	return s
}

func (s *CloudOutboundObClidReportResponseBodyDataList) SetHour(v string) *CloudOutboundObClidReportResponseBodyDataList {
	s.Hour = &v
	return s
}

func (s *CloudOutboundObClidReportResponseBodyDataList) SetMaxBridgeTime(v string) *CloudOutboundObClidReportResponseBodyDataList {
	s.MaxBridgeTime = &v
	return s
}

func (s *CloudOutboundObClidReportResponseBodyDataList) SetMinBridgeTime(v string) *CloudOutboundObClidReportResponseBodyDataList {
	s.MinBridgeTime = &v
	return s
}

func (s *CloudOutboundObClidReportResponseBodyDataList) SetPreviewObCustomerRingingCount(v int64) *CloudOutboundObClidReportResponseBodyDataList {
	s.PreviewObCustomerRingingCount = &v
	return s
}

func (s *CloudOutboundObClidReportResponseBodyDataList) SetPreviewObCustomerRingingRate(v string) *CloudOutboundObClidReportResponseBodyDataList {
	s.PreviewObCustomerRingingRate = &v
	return s
}

func (s *CloudOutboundObClidReportResponseBodyDataList) SetProvince(v string) *CloudOutboundObClidReportResponseBodyDataList {
	s.Province = &v
	return s
}

func (s *CloudOutboundObClidReportResponseBodyDataList) SetRowName(v string) *CloudOutboundObClidReportResponseBodyDataList {
	s.RowName = &v
	return s
}

func (s *CloudOutboundObClidReportResponseBodyDataList) SetTotalBridgeTime(v string) *CloudOutboundObClidReportResponseBodyDataList {
	s.TotalBridgeTime = &v
	return s
}

func (s *CloudOutboundObClidReportResponseBodyDataList) SetTotalCount(v string) *CloudOutboundObClidReportResponseBodyDataList {
	s.TotalCount = &v
	return s
}

func (s *CloudOutboundObClidReportResponseBodyDataList) SetTrunkGroupKey(v string) *CloudOutboundObClidReportResponseBodyDataList {
	s.TrunkGroupKey = &v
	return s
}

func (s *CloudOutboundObClidReportResponseBodyDataList) SetValidAvgBridgeTime(v string) *CloudOutboundObClidReportResponseBodyDataList {
	s.ValidAvgBridgeTime = &v
	return s
}

func (s *CloudOutboundObClidReportResponseBodyDataList) SetValidCalls(v string) *CloudOutboundObClidReportResponseBodyDataList {
	s.ValidCalls = &v
	return s
}

func (s *CloudOutboundObClidReportResponseBodyDataList) SetValidTotalBridgeTime(v string) *CloudOutboundObClidReportResponseBodyDataList {
	s.ValidTotalBridgeTime = &v
	return s
}

func (s *CloudOutboundObClidReportResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
