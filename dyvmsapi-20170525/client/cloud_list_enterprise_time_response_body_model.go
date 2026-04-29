// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudListEnterpriseTimeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CloudListEnterpriseTimeResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CloudListEnterpriseTimeResponseBody
	GetCode() *string
	SetData(v *CloudListEnterpriseTimeResponseBodyData) *CloudListEnterpriseTimeResponseBody
	GetData() *CloudListEnterpriseTimeResponseBodyData
	SetMessage(v string) *CloudListEnterpriseTimeResponseBody
	GetMessage() *string
	SetRequestId(v string) *CloudListEnterpriseTimeResponseBody
	GetRequestId() *string
}

type CloudListEnterpriseTimeResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CloudListEnterpriseTimeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// F655A8D5-B967-440B-8683-DAD6FF8DE990
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloudListEnterpriseTimeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloudListEnterpriseTimeResponseBody) GoString() string {
	return s.String()
}

func (s *CloudListEnterpriseTimeResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CloudListEnterpriseTimeResponseBody) GetCode() *string {
	return s.Code
}

func (s *CloudListEnterpriseTimeResponseBody) GetData() *CloudListEnterpriseTimeResponseBodyData {
	return s.Data
}

func (s *CloudListEnterpriseTimeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CloudListEnterpriseTimeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloudListEnterpriseTimeResponseBody) SetAccessDeniedDetail(v string) *CloudListEnterpriseTimeResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CloudListEnterpriseTimeResponseBody) SetCode(v string) *CloudListEnterpriseTimeResponseBody {
	s.Code = &v
	return s
}

func (s *CloudListEnterpriseTimeResponseBody) SetData(v *CloudListEnterpriseTimeResponseBodyData) *CloudListEnterpriseTimeResponseBody {
	s.Data = v
	return s
}

func (s *CloudListEnterpriseTimeResponseBody) SetMessage(v string) *CloudListEnterpriseTimeResponseBody {
	s.Message = &v
	return s
}

func (s *CloudListEnterpriseTimeResponseBody) SetRequestId(v string) *CloudListEnterpriseTimeResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloudListEnterpriseTimeResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CloudListEnterpriseTimeResponseBodyData struct {
	// 时间条件设置列表
	List []*CloudListEnterpriseTimeResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
}

func (s CloudListEnterpriseTimeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CloudListEnterpriseTimeResponseBodyData) GoString() string {
	return s.String()
}

func (s *CloudListEnterpriseTimeResponseBodyData) GetList() []*CloudListEnterpriseTimeResponseBodyDataList {
	return s.List
}

func (s *CloudListEnterpriseTimeResponseBodyData) SetList(v []*CloudListEnterpriseTimeResponseBodyDataList) *CloudListEnterpriseTimeResponseBodyData {
	s.List = v
	return s
}

func (s *CloudListEnterpriseTimeResponseBodyData) Validate() error {
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

type CloudListEnterpriseTimeResponseBodyDataList struct {
	// 创建时间，格式为： yyyy-MM-dd HH:mm:ss
	//
	// example:
	//
	// 2026-02-20 10:00:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 星期 周一：2、周二：3、周三：4、周四：5、周五：6、周六：7、周日：1、使用","进行分割、type=2时为空串
	//
	// example:
	//
	// 1,2,3
	DayOfWeek *string `json:"DayOfWeek,omitempty" xml:"DayOfWeek,omitempty"`
	// 结束时间 格式为 HH:mm startTime不能大于endTime
	//
	// example:
	//
	// 20:00
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
	// 12314
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 时间条件名称 、同一呼叫中心下不能重复
	//
	// example:
	//
	// time-setting-name
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
	// 10:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 1
	TimeType *string `json:"TimeType,omitempty" xml:"TimeType,omitempty"`
	// 终止日期 格式为：yyyy-MM-dd、type=1时为空串 fromDay不能大于toDay
	//
	// example:
	//
	// 2026-04-22
	ToDay *string `json:"ToDay,omitempty" xml:"ToDay,omitempty"`
	// 时间条件类型 1:按照星期配置 2:按照固定时间配置
	//
	// example:
	//
	// 1
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CloudListEnterpriseTimeResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s CloudListEnterpriseTimeResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *CloudListEnterpriseTimeResponseBodyDataList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CloudListEnterpriseTimeResponseBodyDataList) GetDayOfWeek() *string {
	return s.DayOfWeek
}

func (s *CloudListEnterpriseTimeResponseBodyDataList) GetEndTime() *string {
	return s.EndTime
}

func (s *CloudListEnterpriseTimeResponseBodyDataList) GetEnterpriseId() *string {
	return s.EnterpriseId
}

func (s *CloudListEnterpriseTimeResponseBodyDataList) GetFromDay() *string {
	return s.FromDay
}

func (s *CloudListEnterpriseTimeResponseBodyDataList) GetId() *string {
	return s.Id
}

func (s *CloudListEnterpriseTimeResponseBodyDataList) GetName() *string {
	return s.Name
}

func (s *CloudListEnterpriseTimeResponseBodyDataList) GetPriority() *string {
	return s.Priority
}

func (s *CloudListEnterpriseTimeResponseBodyDataList) GetStartTime() *string {
	return s.StartTime
}

func (s *CloudListEnterpriseTimeResponseBodyDataList) GetTimeType() *string {
	return s.TimeType
}

func (s *CloudListEnterpriseTimeResponseBodyDataList) GetToDay() *string {
	return s.ToDay
}

func (s *CloudListEnterpriseTimeResponseBodyDataList) GetType() *string {
	return s.Type
}

func (s *CloudListEnterpriseTimeResponseBodyDataList) SetCreateTime(v string) *CloudListEnterpriseTimeResponseBodyDataList {
	s.CreateTime = &v
	return s
}

func (s *CloudListEnterpriseTimeResponseBodyDataList) SetDayOfWeek(v string) *CloudListEnterpriseTimeResponseBodyDataList {
	s.DayOfWeek = &v
	return s
}

func (s *CloudListEnterpriseTimeResponseBodyDataList) SetEndTime(v string) *CloudListEnterpriseTimeResponseBodyDataList {
	s.EndTime = &v
	return s
}

func (s *CloudListEnterpriseTimeResponseBodyDataList) SetEnterpriseId(v string) *CloudListEnterpriseTimeResponseBodyDataList {
	s.EnterpriseId = &v
	return s
}

func (s *CloudListEnterpriseTimeResponseBodyDataList) SetFromDay(v string) *CloudListEnterpriseTimeResponseBodyDataList {
	s.FromDay = &v
	return s
}

func (s *CloudListEnterpriseTimeResponseBodyDataList) SetId(v string) *CloudListEnterpriseTimeResponseBodyDataList {
	s.Id = &v
	return s
}

func (s *CloudListEnterpriseTimeResponseBodyDataList) SetName(v string) *CloudListEnterpriseTimeResponseBodyDataList {
	s.Name = &v
	return s
}

func (s *CloudListEnterpriseTimeResponseBodyDataList) SetPriority(v string) *CloudListEnterpriseTimeResponseBodyDataList {
	s.Priority = &v
	return s
}

func (s *CloudListEnterpriseTimeResponseBodyDataList) SetStartTime(v string) *CloudListEnterpriseTimeResponseBodyDataList {
	s.StartTime = &v
	return s
}

func (s *CloudListEnterpriseTimeResponseBodyDataList) SetTimeType(v string) *CloudListEnterpriseTimeResponseBodyDataList {
	s.TimeType = &v
	return s
}

func (s *CloudListEnterpriseTimeResponseBodyDataList) SetToDay(v string) *CloudListEnterpriseTimeResponseBodyDataList {
	s.ToDay = &v
	return s
}

func (s *CloudListEnterpriseTimeResponseBodyDataList) SetType(v string) *CloudListEnterpriseTimeResponseBodyDataList {
	s.Type = &v
	return s
}

func (s *CloudListEnterpriseTimeResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
