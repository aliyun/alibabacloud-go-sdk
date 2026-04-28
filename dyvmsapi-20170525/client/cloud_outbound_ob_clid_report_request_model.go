// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudOutboundObClidReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAreaCodes(v string) *CloudOutboundObClidReportRequest
	GetAreaCodes() *string
	SetEndHour(v int64) *CloudOutboundObClidReportRequest
	GetEndHour() *int64
	SetEndTime(v string) *CloudOutboundObClidReportRequest
	GetEndTime() *string
	SetEnterpriseId(v int64) *CloudOutboundObClidReportRequest
	GetEnterpriseId() *int64
	SetLimit(v int64) *CloudOutboundObClidReportRequest
	GetLimit() *int64
	SetStart(v int64) *CloudOutboundObClidReportRequest
	GetStart() *int64
	SetStartHour(v int64) *CloudOutboundObClidReportRequest
	GetStartHour() *int64
	SetStartTime(v string) *CloudOutboundObClidReportRequest
	GetStartTime() *string
	SetStatisticMethod(v int64) *CloudOutboundObClidReportRequest
	GetStatisticMethod() *int64
	SetTimeRangeType(v int64) *CloudOutboundObClidReportRequest
	GetTimeRangeType() *int64
}

type CloudOutboundObClidReportRequest struct {
	// 说明：根据区号查询指定区域的预览外呼被叫号码统计支持按照多个区号进行查询，多个区号之间使用英文逗号","分隔，默认查询账户下所有地区的预览外呼被叫号码统计
	//
	// example:
	//
	// 010
	AreaCodes *string `json:"AreaCodes,omitempty" xml:"AreaCodes,omitempty"`
	// 统计时段结束时间；取值：0~23,默认值为24时
	//
	// example:
	//
	// 23
	EndHour *int64 `json:"EndHour,omitempty" xml:"EndHour,omitempty"`
	// 说明：统计日期的结束时间，格式：yyyy-MM-dd
	//
	// This parameter is required.
	//
	// example:
	//
	// 2017-06-13
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 呼叫中心 id
	//
	// This parameter is required.
	//
	// example:
	//
	// 7000002
	EnterpriseId *int64 `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	// 查询条数；取值：最大不能超过1000，默认10
	//
	// example:
	//
	// 10
	Limit *int64 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// 查询起始位置；取值：大于等于0，默认0
	//
	// example:
	//
	// 0
	Start *int64 `json:"Start,omitempty" xml:"Start,omitempty"`
	// 统计时段开始时间；取值：0~23,默认值为0时
	//
	// example:
	//
	// 0
	StartHour *int64 `json:"StartHour,omitempty" xml:"StartHour,omitempty"`
	// 说明：统计日期的开始时间，格式：yyyy-MM-dd
	//
	// This parameter is required.
	//
	// example:
	//
	// 2017-06-13
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 统计方法；说明：0：分时1：分日2：汇总11：中继群组汇总
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	StatisticMethod *int64 `json:"StatisticMethod,omitempty" xml:"StatisticMethod,omitempty"`
	// 统计类型；说明：统计报表时间类型，1：日报表2：周报表3：月报表4：自定义时间
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	TimeRangeType *int64 `json:"TimeRangeType,omitempty" xml:"TimeRangeType,omitempty"`
}

func (s CloudOutboundObClidReportRequest) String() string {
	return dara.Prettify(s)
}

func (s CloudOutboundObClidReportRequest) GoString() string {
	return s.String()
}

func (s *CloudOutboundObClidReportRequest) GetAreaCodes() *string {
	return s.AreaCodes
}

func (s *CloudOutboundObClidReportRequest) GetEndHour() *int64 {
	return s.EndHour
}

func (s *CloudOutboundObClidReportRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *CloudOutboundObClidReportRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *CloudOutboundObClidReportRequest) GetLimit() *int64 {
	return s.Limit
}

func (s *CloudOutboundObClidReportRequest) GetStart() *int64 {
	return s.Start
}

func (s *CloudOutboundObClidReportRequest) GetStartHour() *int64 {
	return s.StartHour
}

func (s *CloudOutboundObClidReportRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *CloudOutboundObClidReportRequest) GetStatisticMethod() *int64 {
	return s.StatisticMethod
}

func (s *CloudOutboundObClidReportRequest) GetTimeRangeType() *int64 {
	return s.TimeRangeType
}

func (s *CloudOutboundObClidReportRequest) SetAreaCodes(v string) *CloudOutboundObClidReportRequest {
	s.AreaCodes = &v
	return s
}

func (s *CloudOutboundObClidReportRequest) SetEndHour(v int64) *CloudOutboundObClidReportRequest {
	s.EndHour = &v
	return s
}

func (s *CloudOutboundObClidReportRequest) SetEndTime(v string) *CloudOutboundObClidReportRequest {
	s.EndTime = &v
	return s
}

func (s *CloudOutboundObClidReportRequest) SetEnterpriseId(v int64) *CloudOutboundObClidReportRequest {
	s.EnterpriseId = &v
	return s
}

func (s *CloudOutboundObClidReportRequest) SetLimit(v int64) *CloudOutboundObClidReportRequest {
	s.Limit = &v
	return s
}

func (s *CloudOutboundObClidReportRequest) SetStart(v int64) *CloudOutboundObClidReportRequest {
	s.Start = &v
	return s
}

func (s *CloudOutboundObClidReportRequest) SetStartHour(v int64) *CloudOutboundObClidReportRequest {
	s.StartHour = &v
	return s
}

func (s *CloudOutboundObClidReportRequest) SetStartTime(v string) *CloudOutboundObClidReportRequest {
	s.StartTime = &v
	return s
}

func (s *CloudOutboundObClidReportRequest) SetStatisticMethod(v int64) *CloudOutboundObClidReportRequest {
	s.StatisticMethod = &v
	return s
}

func (s *CloudOutboundObClidReportRequest) SetTimeRangeType(v int64) *CloudOutboundObClidReportRequest {
	s.TimeRangeType = &v
	return s
}

func (s *CloudOutboundObClidReportRequest) Validate() error {
	return dara.Validate(s)
}
