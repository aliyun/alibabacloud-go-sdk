// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudOutboundPreviewObReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCnos(v string) *CloudOutboundPreviewObReportRequest
	GetCnos() *string
	SetEndHour(v int64) *CloudOutboundPreviewObReportRequest
	GetEndHour() *int64
	SetEndTime(v string) *CloudOutboundPreviewObReportRequest
	GetEndTime() *string
	SetEnterpriseId(v int64) *CloudOutboundPreviewObReportRequest
	GetEnterpriseId() *int64
	SetLimit(v int64) *CloudOutboundPreviewObReportRequest
	GetLimit() *int64
	SetStart(v int64) *CloudOutboundPreviewObReportRequest
	GetStart() *int64
	SetStartHour(v int64) *CloudOutboundPreviewObReportRequest
	GetStartHour() *int64
	SetStartTime(v string) *CloudOutboundPreviewObReportRequest
	GetStartTime() *string
	SetStatisticMethod(v int64) *CloudOutboundPreviewObReportRequest
	GetStatisticMethod() *int64
	SetTimeRangeType(v int64) *CloudOutboundPreviewObReportRequest
	GetTimeRangeType() *int64
}

type CloudOutboundPreviewObReportRequest struct {
	// 座席号；说明：根据座席工号查询指定座席的预览外呼统计支持按照多个座席工号进行查询，多个座席工号之间使用英文逗号","分隔，默认查询账户下所有座席的预览外呼统计
	//
	// This parameter is required.
	//
	// example:
	//
	// 3008
	Cnos *string `json:"Cnos,omitempty" xml:"Cnos,omitempty"`
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
	// 6000001
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
	// 统计方法；说明：0:分时1：分日2：汇总
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

func (s CloudOutboundPreviewObReportRequest) String() string {
	return dara.Prettify(s)
}

func (s CloudOutboundPreviewObReportRequest) GoString() string {
	return s.String()
}

func (s *CloudOutboundPreviewObReportRequest) GetCnos() *string {
	return s.Cnos
}

func (s *CloudOutboundPreviewObReportRequest) GetEndHour() *int64 {
	return s.EndHour
}

func (s *CloudOutboundPreviewObReportRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *CloudOutboundPreviewObReportRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *CloudOutboundPreviewObReportRequest) GetLimit() *int64 {
	return s.Limit
}

func (s *CloudOutboundPreviewObReportRequest) GetStart() *int64 {
	return s.Start
}

func (s *CloudOutboundPreviewObReportRequest) GetStartHour() *int64 {
	return s.StartHour
}

func (s *CloudOutboundPreviewObReportRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *CloudOutboundPreviewObReportRequest) GetStatisticMethod() *int64 {
	return s.StatisticMethod
}

func (s *CloudOutboundPreviewObReportRequest) GetTimeRangeType() *int64 {
	return s.TimeRangeType
}

func (s *CloudOutboundPreviewObReportRequest) SetCnos(v string) *CloudOutboundPreviewObReportRequest {
	s.Cnos = &v
	return s
}

func (s *CloudOutboundPreviewObReportRequest) SetEndHour(v int64) *CloudOutboundPreviewObReportRequest {
	s.EndHour = &v
	return s
}

func (s *CloudOutboundPreviewObReportRequest) SetEndTime(v string) *CloudOutboundPreviewObReportRequest {
	s.EndTime = &v
	return s
}

func (s *CloudOutboundPreviewObReportRequest) SetEnterpriseId(v int64) *CloudOutboundPreviewObReportRequest {
	s.EnterpriseId = &v
	return s
}

func (s *CloudOutboundPreviewObReportRequest) SetLimit(v int64) *CloudOutboundPreviewObReportRequest {
	s.Limit = &v
	return s
}

func (s *CloudOutboundPreviewObReportRequest) SetStart(v int64) *CloudOutboundPreviewObReportRequest {
	s.Start = &v
	return s
}

func (s *CloudOutboundPreviewObReportRequest) SetStartHour(v int64) *CloudOutboundPreviewObReportRequest {
	s.StartHour = &v
	return s
}

func (s *CloudOutboundPreviewObReportRequest) SetStartTime(v string) *CloudOutboundPreviewObReportRequest {
	s.StartTime = &v
	return s
}

func (s *CloudOutboundPreviewObReportRequest) SetStatisticMethod(v int64) *CloudOutboundPreviewObReportRequest {
	s.StatisticMethod = &v
	return s
}

func (s *CloudOutboundPreviewObReportRequest) SetTimeRangeType(v int64) *CloudOutboundPreviewObReportRequest {
	s.TimeRangeType = &v
	return s
}

func (s *CloudOutboundPreviewObReportRequest) Validate() error {
	return dara.Validate(s)
}
