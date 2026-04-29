// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudCreateEnterpriseTimeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDayOfWeek(v string) *CloudCreateEnterpriseTimeRequest
	GetDayOfWeek() *string
	SetEndTime(v string) *CloudCreateEnterpriseTimeRequest
	GetEndTime() *string
	SetEnterpriseId(v int64) *CloudCreateEnterpriseTimeRequest
	GetEnterpriseId() *int64
	SetFromDay(v string) *CloudCreateEnterpriseTimeRequest
	GetFromDay() *string
	SetName(v string) *CloudCreateEnterpriseTimeRequest
	GetName() *string
	SetOwnerId(v int64) *CloudCreateEnterpriseTimeRequest
	GetOwnerId() *int64
	SetPriority(v int64) *CloudCreateEnterpriseTimeRequest
	GetPriority() *int64
	SetResourceOwnerAccount(v string) *CloudCreateEnterpriseTimeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CloudCreateEnterpriseTimeRequest
	GetResourceOwnerId() *int64
	SetStartTime(v string) *CloudCreateEnterpriseTimeRequest
	GetStartTime() *string
	SetTimeType(v int64) *CloudCreateEnterpriseTimeRequest
	GetTimeType() *int64
	SetToDay(v string) *CloudCreateEnterpriseTimeRequest
	GetToDay() *string
	SetType(v int64) *CloudCreateEnterpriseTimeRequest
	GetType() *int64
}

type CloudCreateEnterpriseTimeRequest struct {
	// 当type=1 时必选,周一：2、周二：3、周三：4、周四：5、周五：6、周六：7、周日：1。星期值以,分隔 例如： 2,3,4代表周一周二周三、type=2时为空串
	//
	// example:
	//
	// 2,3,4
	DayOfWeek *string `json:"DayOfWeek,omitempty" xml:"DayOfWeek,omitempty"`
	// 结束时间；格式为 HH:mm 例如、 19:00 startTime不能大于endTime
	//
	// This parameter is required.
	//
	// example:
	//
	// 19:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 呼叫中心 id
	//
	// This parameter is required.
	//
	// example:
	//
	// 7000002
	EnterpriseId *int64 `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	// 起始日期；当type=2 时必选，格式为：yyyy-MM-dd 、type=1时为空串 fromDay不能大于toDay
	//
	// example:
	//
	// 2026-04-20
	FromDay *string `json:"FromDay,omitempty" xml:"FromDay,omitempty"`
	// 时间条件名称；同一呼叫中心下不能重复
	//
	// This parameter is required.
	//
	// example:
	//
	// demo01
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 时间条件优先级；同一呼叫中心下不能重复，值从1开始，值越小优先级越高
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Priority             *int64  `json:"Priority,omitempty" xml:"Priority,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 开始时间；格式为 HH:mm 例如、 06:00 startTime不能大于endTime
	//
	// This parameter is required.
	//
	// example:
	//
	// 06:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 时间类型 1.连续 2.间隔
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	TimeType *int64 `json:"TimeType,omitempty" xml:"TimeType,omitempty"`
	// 终止日期；当type=2 时必选，格式为：yyyy-MM-dd 、type=1时为空串 fromDay不能大于toDay
	//
	// example:
	//
	// 2026-04-25
	ToDay *string `json:"ToDay,omitempty" xml:"ToDay,omitempty"`
	// 条件类型 1:按照星期配置 2:按照固定时间配置
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Type *int64 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CloudCreateEnterpriseTimeRequest) String() string {
	return dara.Prettify(s)
}

func (s CloudCreateEnterpriseTimeRequest) GoString() string {
	return s.String()
}

func (s *CloudCreateEnterpriseTimeRequest) GetDayOfWeek() *string {
	return s.DayOfWeek
}

func (s *CloudCreateEnterpriseTimeRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *CloudCreateEnterpriseTimeRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *CloudCreateEnterpriseTimeRequest) GetFromDay() *string {
	return s.FromDay
}

func (s *CloudCreateEnterpriseTimeRequest) GetName() *string {
	return s.Name
}

func (s *CloudCreateEnterpriseTimeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CloudCreateEnterpriseTimeRequest) GetPriority() *int64 {
	return s.Priority
}

func (s *CloudCreateEnterpriseTimeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CloudCreateEnterpriseTimeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CloudCreateEnterpriseTimeRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *CloudCreateEnterpriseTimeRequest) GetTimeType() *int64 {
	return s.TimeType
}

func (s *CloudCreateEnterpriseTimeRequest) GetToDay() *string {
	return s.ToDay
}

func (s *CloudCreateEnterpriseTimeRequest) GetType() *int64 {
	return s.Type
}

func (s *CloudCreateEnterpriseTimeRequest) SetDayOfWeek(v string) *CloudCreateEnterpriseTimeRequest {
	s.DayOfWeek = &v
	return s
}

func (s *CloudCreateEnterpriseTimeRequest) SetEndTime(v string) *CloudCreateEnterpriseTimeRequest {
	s.EndTime = &v
	return s
}

func (s *CloudCreateEnterpriseTimeRequest) SetEnterpriseId(v int64) *CloudCreateEnterpriseTimeRequest {
	s.EnterpriseId = &v
	return s
}

func (s *CloudCreateEnterpriseTimeRequest) SetFromDay(v string) *CloudCreateEnterpriseTimeRequest {
	s.FromDay = &v
	return s
}

func (s *CloudCreateEnterpriseTimeRequest) SetName(v string) *CloudCreateEnterpriseTimeRequest {
	s.Name = &v
	return s
}

func (s *CloudCreateEnterpriseTimeRequest) SetOwnerId(v int64) *CloudCreateEnterpriseTimeRequest {
	s.OwnerId = &v
	return s
}

func (s *CloudCreateEnterpriseTimeRequest) SetPriority(v int64) *CloudCreateEnterpriseTimeRequest {
	s.Priority = &v
	return s
}

func (s *CloudCreateEnterpriseTimeRequest) SetResourceOwnerAccount(v string) *CloudCreateEnterpriseTimeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CloudCreateEnterpriseTimeRequest) SetResourceOwnerId(v int64) *CloudCreateEnterpriseTimeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CloudCreateEnterpriseTimeRequest) SetStartTime(v string) *CloudCreateEnterpriseTimeRequest {
	s.StartTime = &v
	return s
}

func (s *CloudCreateEnterpriseTimeRequest) SetTimeType(v int64) *CloudCreateEnterpriseTimeRequest {
	s.TimeType = &v
	return s
}

func (s *CloudCreateEnterpriseTimeRequest) SetToDay(v string) *CloudCreateEnterpriseTimeRequest {
	s.ToDay = &v
	return s
}

func (s *CloudCreateEnterpriseTimeRequest) SetType(v int64) *CloudCreateEnterpriseTimeRequest {
	s.Type = &v
	return s
}

func (s *CloudCreateEnterpriseTimeRequest) Validate() error {
	return dara.Validate(s)
}
