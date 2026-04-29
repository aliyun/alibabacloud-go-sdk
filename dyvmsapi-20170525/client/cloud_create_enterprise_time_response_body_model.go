// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudCreateEnterpriseTimeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CloudCreateEnterpriseTimeResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CloudCreateEnterpriseTimeResponseBody
	GetCode() *string
	SetData(v *CloudCreateEnterpriseTimeResponseBodyData) *CloudCreateEnterpriseTimeResponseBody
	GetData() *CloudCreateEnterpriseTimeResponseBodyData
	SetMessage(v string) *CloudCreateEnterpriseTimeResponseBody
	GetMessage() *string
	SetRequestId(v string) *CloudCreateEnterpriseTimeResponseBody
	GetRequestId() *string
}

type CloudCreateEnterpriseTimeResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CloudCreateEnterpriseTimeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// ED18A5AE-9BBC-5851-1111-8E5B8C23CEDE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloudCreateEnterpriseTimeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloudCreateEnterpriseTimeResponseBody) GoString() string {
	return s.String()
}

func (s *CloudCreateEnterpriseTimeResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CloudCreateEnterpriseTimeResponseBody) GetCode() *string {
	return s.Code
}

func (s *CloudCreateEnterpriseTimeResponseBody) GetData() *CloudCreateEnterpriseTimeResponseBodyData {
	return s.Data
}

func (s *CloudCreateEnterpriseTimeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CloudCreateEnterpriseTimeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloudCreateEnterpriseTimeResponseBody) SetAccessDeniedDetail(v string) *CloudCreateEnterpriseTimeResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CloudCreateEnterpriseTimeResponseBody) SetCode(v string) *CloudCreateEnterpriseTimeResponseBody {
	s.Code = &v
	return s
}

func (s *CloudCreateEnterpriseTimeResponseBody) SetData(v *CloudCreateEnterpriseTimeResponseBodyData) *CloudCreateEnterpriseTimeResponseBody {
	s.Data = v
	return s
}

func (s *CloudCreateEnterpriseTimeResponseBody) SetMessage(v string) *CloudCreateEnterpriseTimeResponseBody {
	s.Message = &v
	return s
}

func (s *CloudCreateEnterpriseTimeResponseBody) SetRequestId(v string) *CloudCreateEnterpriseTimeResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloudCreateEnterpriseTimeResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CloudCreateEnterpriseTimeResponseBodyData struct {
	// 创建时间，格式为： yyyy-MM-dd HH:mm:ss
	//
	// example:
	//
	// 2026-04-20 10:00:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 星期 周一：2、周二：3、周三：4、周四：5、周五：6、周六：7、周日：1、使用","进行分割、type=2时为空串
	//
	// example:
	//
	// 2,3,4
	DayOfWeek *string `json:"DayOfWeek,omitempty" xml:"DayOfWeek,omitempty"`
	// 结束时间 格式为 HH:mm startTime不能大于endTime
	//
	// example:
	//
	// 19:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 呼叫中心编号
	//
	// example:
	//
	// 7000002
	EnterpriseId *string `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	// 起始日期 格式为：yyyy-MM-dd、type=1时为空串 fromDay不能大于toDay
	//
	// example:
	//
	// 2026-04-20
	FromDay *string `json:"FromDay,omitempty" xml:"FromDay,omitempty"`
	// 时间条件id
	//
	// example:
	//
	// 236
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 时间条件名称 、同一呼叫中心下不能重复
	//
	// example:
	//
	// demo01
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 优先级、同一呼叫中心下不能重复，值从1开始，值越小优先级越高
	//
	// example:
	//
	// 1
	Priority *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// 开始时间 格式为 HH:mm startTime不能大于endTime
	//
	// example:
	//
	// 06:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 时间类型 1.连续 2.间隔
	//
	// example:
	//
	// 1
	TimeType *string `json:"TimeType,omitempty" xml:"TimeType,omitempty"`
	// 终止日期 格式为：yyyy-MM-dd、type=1时为空串 fromDay不能大于toDay
	//
	// example:
	//
	// 2026-04-25
	ToDay *string `json:"ToDay,omitempty" xml:"ToDay,omitempty"`
	// 时间条件类型 1:按照星期配置 2:按照固定时间配置
	//
	// example:
	//
	// 1
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CloudCreateEnterpriseTimeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CloudCreateEnterpriseTimeResponseBodyData) GoString() string {
	return s.String()
}

func (s *CloudCreateEnterpriseTimeResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CloudCreateEnterpriseTimeResponseBodyData) GetDayOfWeek() *string {
	return s.DayOfWeek
}

func (s *CloudCreateEnterpriseTimeResponseBodyData) GetEndTime() *string {
	return s.EndTime
}

func (s *CloudCreateEnterpriseTimeResponseBodyData) GetEnterpriseId() *string {
	return s.EnterpriseId
}

func (s *CloudCreateEnterpriseTimeResponseBodyData) GetFromDay() *string {
	return s.FromDay
}

func (s *CloudCreateEnterpriseTimeResponseBodyData) GetId() *string {
	return s.Id
}

func (s *CloudCreateEnterpriseTimeResponseBodyData) GetName() *string {
	return s.Name
}

func (s *CloudCreateEnterpriseTimeResponseBodyData) GetPriority() *string {
	return s.Priority
}

func (s *CloudCreateEnterpriseTimeResponseBodyData) GetStartTime() *string {
	return s.StartTime
}

func (s *CloudCreateEnterpriseTimeResponseBodyData) GetTimeType() *string {
	return s.TimeType
}

func (s *CloudCreateEnterpriseTimeResponseBodyData) GetToDay() *string {
	return s.ToDay
}

func (s *CloudCreateEnterpriseTimeResponseBodyData) GetType() *string {
	return s.Type
}

func (s *CloudCreateEnterpriseTimeResponseBodyData) SetCreateTime(v string) *CloudCreateEnterpriseTimeResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *CloudCreateEnterpriseTimeResponseBodyData) SetDayOfWeek(v string) *CloudCreateEnterpriseTimeResponseBodyData {
	s.DayOfWeek = &v
	return s
}

func (s *CloudCreateEnterpriseTimeResponseBodyData) SetEndTime(v string) *CloudCreateEnterpriseTimeResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *CloudCreateEnterpriseTimeResponseBodyData) SetEnterpriseId(v string) *CloudCreateEnterpriseTimeResponseBodyData {
	s.EnterpriseId = &v
	return s
}

func (s *CloudCreateEnterpriseTimeResponseBodyData) SetFromDay(v string) *CloudCreateEnterpriseTimeResponseBodyData {
	s.FromDay = &v
	return s
}

func (s *CloudCreateEnterpriseTimeResponseBodyData) SetId(v string) *CloudCreateEnterpriseTimeResponseBodyData {
	s.Id = &v
	return s
}

func (s *CloudCreateEnterpriseTimeResponseBodyData) SetName(v string) *CloudCreateEnterpriseTimeResponseBodyData {
	s.Name = &v
	return s
}

func (s *CloudCreateEnterpriseTimeResponseBodyData) SetPriority(v string) *CloudCreateEnterpriseTimeResponseBodyData {
	s.Priority = &v
	return s
}

func (s *CloudCreateEnterpriseTimeResponseBodyData) SetStartTime(v string) *CloudCreateEnterpriseTimeResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *CloudCreateEnterpriseTimeResponseBodyData) SetTimeType(v string) *CloudCreateEnterpriseTimeResponseBodyData {
	s.TimeType = &v
	return s
}

func (s *CloudCreateEnterpriseTimeResponseBodyData) SetToDay(v string) *CloudCreateEnterpriseTimeResponseBodyData {
	s.ToDay = &v
	return s
}

func (s *CloudCreateEnterpriseTimeResponseBodyData) SetType(v string) *CloudCreateEnterpriseTimeResponseBodyData {
	s.Type = &v
	return s
}

func (s *CloudCreateEnterpriseTimeResponseBodyData) Validate() error {
	return dara.Validate(s)
}
