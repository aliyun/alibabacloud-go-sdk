// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClinkStatIbRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDate(v string) *ClinkStatIbRequest
	GetDate() *string
	SetDateEnd(v string) *ClinkStatIbRequest
	GetDateEnd() *string
	SetEndHour(v int64) *ClinkStatIbRequest
	GetEndHour() *int64
	SetEndMinute(v int64) *ClinkStatIbRequest
	GetEndMinute() *int64
	SetEnterpriseId(v int64) *ClinkStatIbRequest
	GetEnterpriseId() *int64
	SetFields(v string) *ClinkStatIbRequest
	GetFields() *string
	SetHotlines(v string) *ClinkStatIbRequest
	GetHotlines() *string
	SetOwnerId(v int64) *ClinkStatIbRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ClinkStatIbRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ClinkStatIbRequest
	GetResourceOwnerId() *int64
	SetStartHour(v int64) *ClinkStatIbRequest
	GetStartHour() *int64
	SetStartMinute(v int64) *ClinkStatIbRequest
	GetStartMinute() *int64
	SetStatisticMethod(v int64) *ClinkStatIbRequest
	GetStatisticMethod() *int64
}

type ClinkStatIbRequest struct {
	// 同步日期，时间格式(yyyyMMdd)
	//
	// This parameter is required.
	//
	// example:
	//
	// 20240303
	Date *string `json:"Date,omitempty" xml:"Date,omitempty"`
	// 同步日期截止，时间格式(yyyyMMdd) ，默认与date相同，查询时间范围最大支持6个月
	//
	// example:
	//
	// 20240303
	DateEnd *string `json:"DateEnd,omitempty" xml:"DateEnd,omitempty"`
	// 查询结束时间 (单位：小时，范围：0-23)，缺省值为23
	//
	// example:
	//
	// 23
	EndHour *int64 `json:"EndHour,omitempty" xml:"EndHour,omitempty"`
	// 查询结束分钟 (单位：分，范围：0、15、30、45、59)，缺省值为59 注：由startHour:startMinute和endHour:endMinute组成的开始时间和结束时间，开始时间不得晚于或等于结束时间。 只使用startHour或endHour时，请注意startMinute与endMinute的缺省值。
	//
	// example:
	//
	// 59
	EndMinute *int64 `json:"EndMinute,omitempty" xml:"EndMinute,omitempty"`
	// 呼叫中心 id
	//
	// This parameter is required.
	//
	// example:
	//
	// 8004970
	EnterpriseId *int64 `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	// 指定需要显示的字段 (默认全部)
	//
	// example:
	//
	// ibTotalCount,ibAnsweredCount
	Fields *string `json:"Fields,omitempty" xml:"Fields,omitempty"`
	// 热线号码，指定需要查询的热线号码 (默认全部热线号码)
	//
	// example:
	//
	// xxxx,xxxx
	Hotlines             *string `json:"Hotlines,omitempty" xml:"Hotlines,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 查询开始时间 (单位：小时，范围：0-23)，缺省值为0
	//
	// example:
	//
	// 0
	StartHour *int64 `json:"StartHour,omitempty" xml:"StartHour,omitempty"`
	// 查询开始分钟 (单位：分，范围：0、15、30、45)，缺省值为0
	//
	// example:
	//
	// 0
	StartMinute *int64 `json:"StartMinute,omitempty" xml:"StartMinute,omitempty"`
	// 统计方式 (默认为2) 取值范围为[2,3]; 2:汇总统计;3:分时统计 注：分时统计下只会返回存在队列工作情况的数据，若队列在该时段没有工作或来电，则不会返回该时段的数据
	//
	// example:
	//
	// 2
	StatisticMethod *int64 `json:"StatisticMethod,omitempty" xml:"StatisticMethod,omitempty"`
}

func (s ClinkStatIbRequest) String() string {
	return dara.Prettify(s)
}

func (s ClinkStatIbRequest) GoString() string {
	return s.String()
}

func (s *ClinkStatIbRequest) GetDate() *string {
	return s.Date
}

func (s *ClinkStatIbRequest) GetDateEnd() *string {
	return s.DateEnd
}

func (s *ClinkStatIbRequest) GetEndHour() *int64 {
	return s.EndHour
}

func (s *ClinkStatIbRequest) GetEndMinute() *int64 {
	return s.EndMinute
}

func (s *ClinkStatIbRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *ClinkStatIbRequest) GetFields() *string {
	return s.Fields
}

func (s *ClinkStatIbRequest) GetHotlines() *string {
	return s.Hotlines
}

func (s *ClinkStatIbRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ClinkStatIbRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ClinkStatIbRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ClinkStatIbRequest) GetStartHour() *int64 {
	return s.StartHour
}

func (s *ClinkStatIbRequest) GetStartMinute() *int64 {
	return s.StartMinute
}

func (s *ClinkStatIbRequest) GetStatisticMethod() *int64 {
	return s.StatisticMethod
}

func (s *ClinkStatIbRequest) SetDate(v string) *ClinkStatIbRequest {
	s.Date = &v
	return s
}

func (s *ClinkStatIbRequest) SetDateEnd(v string) *ClinkStatIbRequest {
	s.DateEnd = &v
	return s
}

func (s *ClinkStatIbRequest) SetEndHour(v int64) *ClinkStatIbRequest {
	s.EndHour = &v
	return s
}

func (s *ClinkStatIbRequest) SetEndMinute(v int64) *ClinkStatIbRequest {
	s.EndMinute = &v
	return s
}

func (s *ClinkStatIbRequest) SetEnterpriseId(v int64) *ClinkStatIbRequest {
	s.EnterpriseId = &v
	return s
}

func (s *ClinkStatIbRequest) SetFields(v string) *ClinkStatIbRequest {
	s.Fields = &v
	return s
}

func (s *ClinkStatIbRequest) SetHotlines(v string) *ClinkStatIbRequest {
	s.Hotlines = &v
	return s
}

func (s *ClinkStatIbRequest) SetOwnerId(v int64) *ClinkStatIbRequest {
	s.OwnerId = &v
	return s
}

func (s *ClinkStatIbRequest) SetResourceOwnerAccount(v string) *ClinkStatIbRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ClinkStatIbRequest) SetResourceOwnerId(v int64) *ClinkStatIbRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ClinkStatIbRequest) SetStartHour(v int64) *ClinkStatIbRequest {
	s.StartHour = &v
	return s
}

func (s *ClinkStatIbRequest) SetStartMinute(v int64) *ClinkStatIbRequest {
	s.StartMinute = &v
	return s
}

func (s *ClinkStatIbRequest) SetStatisticMethod(v int64) *ClinkStatIbRequest {
	s.StatisticMethod = &v
	return s
}

func (s *ClinkStatIbRequest) Validate() error {
	return dara.Validate(s)
}
