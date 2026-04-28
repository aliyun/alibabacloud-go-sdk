// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudQueryIbCdrRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCalleeNumber(v string) *CloudQueryIbCdrRequest
	GetCalleeNumber() *string
	SetCity(v string) *CloudQueryIbCdrRequest
	GetCity() *string
	SetCno(v string) *CloudQueryIbCdrRequest
	GetCno() *string
	SetCustomerNumber(v string) *CloudQueryIbCdrRequest
	GetCustomerNumber() *string
	SetEndTime(v int64) *CloudQueryIbCdrRequest
	GetEndTime() *int64
	SetEnterpriseId(v int64) *CloudQueryIbCdrRequest
	GetEnterpriseId() *int64
	SetHotline(v string) *CloudQueryIbCdrRequest
	GetHotline() *string
	SetJoinQueueCode(v int64) *CloudQueryIbCdrRequest
	GetJoinQueueCode() *int64
	SetLeaveQueueCode(v int64) *CloudQueryIbCdrRequest
	GetLeaveQueueCode() *int64
	SetLimit(v int64) *CloudQueryIbCdrRequest
	GetLimit() *int64
	SetProvince(v string) *CloudQueryIbCdrRequest
	GetProvince() *string
	SetQno(v string) *CloudQueryIbCdrRequest
	GetQno() *string
	SetStart(v int64) *CloudQueryIbCdrRequest
	GetStart() *int64
	SetStartTime(v int64) *CloudQueryIbCdrRequest
	GetStartTime() *int64
	SetStatus(v int64) *CloudQueryIbCdrRequest
	GetStatus() *int64
	SetTimeRangeType(v string) *CloudQueryIbCdrRequest
	GetTimeRangeType() *string
	SetUniqueId(v string) *CloudQueryIbCdrRequest
	GetUniqueId() *string
	SetUserFieldValue(v string) *CloudQueryIbCdrRequest
	GetUserFieldValue() *string
	SetUserFieldkey(v string) *CloudQueryIbCdrRequest
	GetUserFieldkey() *string
}

type CloudQueryIbCdrRequest struct {
	// 座席号码
	//
	// example:
	//
	// 41008502
	CalleeNumber *string `json:"CalleeNumber,omitempty" xml:"CalleeNumber,omitempty"`
	// 客户电话归属城市；为空表示全部，如"成都"，此参数需要URLEncode
	//
	// example:
	//
	// 示例值
	City *string `json:"City,omitempty" xml:"City,omitempty"`
	// 座席工号,此字段支持传入多个座席工号，使用英文半角逗号隔开
	//
	// example:
	//
	// 1111
	Cno *string `json:"Cno,omitempty" xml:"Cno,omitempty"`
	// 客户号码
	//
	// example:
	//
	// 17750247753
	CustomerNumber *string `json:"CustomerNumber,omitempty" xml:"CustomerNumber,omitempty"`
	// 查询结束时间；时间戳格式,精确到s。startTime与endTime不允许跨月，默认值取当前月份最后一天
	//
	// example:
	//
	// 1775030413
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 呼叫中心 id
	//
	// This parameter is required.
	//
	// example:
	//
	// 7000002
	EnterpriseId *int64 `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	// 热线号码
	//
	// example:
	//
	// 10003221
	Hotline *string `json:"Hotline,omitempty" xml:"Hotline,omitempty"`
	// 未进入队列原因 参数说明： 4: 无座席未进入队列 , 3: 队列满未进入队列
	//
	// example:
	//
	// 3
	JoinQueueCode *int64 `json:"JoinQueueCode,omitempty" xml:"JoinQueueCode,omitempty"`
	// 离开队列原因 参数说明： 2: 队列中放弃 , 3: 队列中超时溢出 , 4: 队列中无座席溢出
	//
	// example:
	//
	// 2
	LeaveQueueCode *int64 `json:"LeaveQueueCode,omitempty" xml:"LeaveQueueCode,omitempty"`
	// 需要取出的数据条数；大于0，最大不能超过1000，默认10
	//
	// example:
	//
	// 10
	Limit *int64 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// 客户电话归属省份；为空表示全部，如"四川"，此参数需要URLEncode
	//
	// example:
	//
	// 示例值
	Province *string `json:"Province,omitempty" xml:"Province,omitempty"`
	// 队列号 参数说明：-1: 未进入队列
	//
	// example:
	//
	// 566
	Qno *string `json:"Qno,omitempty" xml:"Qno,omitempty"`
	// 从第几条开始取；大于等于0，默认0
	//
	// example:
	//
	// 0
	Start *int64 `json:"Start,omitempty" xml:"Start,omitempty"`
	// 查询开始时间；时间戳格式,精确到s。startTime与endTime不允许跨月，默认值取当前月份第一天
	//
	// example:
	//
	// 1775024775
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 呼叫结果；参数说明： 1:座席接听 2:已呼叫座席，座席未接听 3:系统接听 4:系统未接听-IVR配置错误 5:系统未接听-停机 6:系统未接听-欠费 7:系统未接听-黑名单 8:系统未接听-未注册 9:系统未接听-彩铃 11:系统未接听-呼叫超出营帐中设置的最大限制 12:系统未接听-客户呼入系统后在系统未应答前挂机 13:其他错误 14:未进入队列 15:队列中放弃 16:队列中超时溢出 17:队列中无座席溢出 18:无座席未进入队列 19:队列满未进入队列
	//
	// example:
	//
	// 1
	Status *int64 `json:"Status,omitempty" xml:"Status,omitempty"`
	// 查询时间开始结束范围生效类型；通话记录startTime和endTime时间范围生效类型，当startTime和endTime不为空时生效；取值：1.呼叫开始时间 2.呼叫结束时间； 默认为1
	//
	// example:
	//
	// 1
	TimeRangeType *string `json:"TimeRangeType,omitempty" xml:"TimeRangeType,omitempty"`
	// 如果uniqueId值不为空，则以下其它参数无效，即仅查询电话唯一标识为uniqueId的记录信息
	//
	// example:
	//
	// uniq_1-162046xxxx.12
	UniqueId *string `json:"UniqueId,omitempty" xml:"UniqueId,omitempty"`
	// 用户自定义参数指定查询value；传递该参数时必须传递userFieldKey，仅查询一个字段，不支持模糊匹配， 构造的自定义字段查询条件为:{"userFieldkey":"userFieldvalue"}，此参数需要URLEncode
	//
	// example:
	//
	// userFieldvalue
	UserFieldValue *string `json:"UserFieldValue,omitempty" xml:"UserFieldValue,omitempty"`
	// 用户自定义参数指定查询key；传递该参数时必须传递userFieldValue，仅查询一个字段，不支持模糊匹配， 构造的自定义字段查询条件为:{"userFieldkey":"userFieldvalue"}，此参数需要URLEncode
	//
	// example:
	//
	// userFieldkey
	UserFieldkey *string `json:"UserFieldkey,omitempty" xml:"UserFieldkey,omitempty"`
}

func (s CloudQueryIbCdrRequest) String() string {
	return dara.Prettify(s)
}

func (s CloudQueryIbCdrRequest) GoString() string {
	return s.String()
}

func (s *CloudQueryIbCdrRequest) GetCalleeNumber() *string {
	return s.CalleeNumber
}

func (s *CloudQueryIbCdrRequest) GetCity() *string {
	return s.City
}

func (s *CloudQueryIbCdrRequest) GetCno() *string {
	return s.Cno
}

func (s *CloudQueryIbCdrRequest) GetCustomerNumber() *string {
	return s.CustomerNumber
}

func (s *CloudQueryIbCdrRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *CloudQueryIbCdrRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *CloudQueryIbCdrRequest) GetHotline() *string {
	return s.Hotline
}

func (s *CloudQueryIbCdrRequest) GetJoinQueueCode() *int64 {
	return s.JoinQueueCode
}

func (s *CloudQueryIbCdrRequest) GetLeaveQueueCode() *int64 {
	return s.LeaveQueueCode
}

func (s *CloudQueryIbCdrRequest) GetLimit() *int64 {
	return s.Limit
}

func (s *CloudQueryIbCdrRequest) GetProvince() *string {
	return s.Province
}

func (s *CloudQueryIbCdrRequest) GetQno() *string {
	return s.Qno
}

func (s *CloudQueryIbCdrRequest) GetStart() *int64 {
	return s.Start
}

func (s *CloudQueryIbCdrRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *CloudQueryIbCdrRequest) GetStatus() *int64 {
	return s.Status
}

func (s *CloudQueryIbCdrRequest) GetTimeRangeType() *string {
	return s.TimeRangeType
}

func (s *CloudQueryIbCdrRequest) GetUniqueId() *string {
	return s.UniqueId
}

func (s *CloudQueryIbCdrRequest) GetUserFieldValue() *string {
	return s.UserFieldValue
}

func (s *CloudQueryIbCdrRequest) GetUserFieldkey() *string {
	return s.UserFieldkey
}

func (s *CloudQueryIbCdrRequest) SetCalleeNumber(v string) *CloudQueryIbCdrRequest {
	s.CalleeNumber = &v
	return s
}

func (s *CloudQueryIbCdrRequest) SetCity(v string) *CloudQueryIbCdrRequest {
	s.City = &v
	return s
}

func (s *CloudQueryIbCdrRequest) SetCno(v string) *CloudQueryIbCdrRequest {
	s.Cno = &v
	return s
}

func (s *CloudQueryIbCdrRequest) SetCustomerNumber(v string) *CloudQueryIbCdrRequest {
	s.CustomerNumber = &v
	return s
}

func (s *CloudQueryIbCdrRequest) SetEndTime(v int64) *CloudQueryIbCdrRequest {
	s.EndTime = &v
	return s
}

func (s *CloudQueryIbCdrRequest) SetEnterpriseId(v int64) *CloudQueryIbCdrRequest {
	s.EnterpriseId = &v
	return s
}

func (s *CloudQueryIbCdrRequest) SetHotline(v string) *CloudQueryIbCdrRequest {
	s.Hotline = &v
	return s
}

func (s *CloudQueryIbCdrRequest) SetJoinQueueCode(v int64) *CloudQueryIbCdrRequest {
	s.JoinQueueCode = &v
	return s
}

func (s *CloudQueryIbCdrRequest) SetLeaveQueueCode(v int64) *CloudQueryIbCdrRequest {
	s.LeaveQueueCode = &v
	return s
}

func (s *CloudQueryIbCdrRequest) SetLimit(v int64) *CloudQueryIbCdrRequest {
	s.Limit = &v
	return s
}

func (s *CloudQueryIbCdrRequest) SetProvince(v string) *CloudQueryIbCdrRequest {
	s.Province = &v
	return s
}

func (s *CloudQueryIbCdrRequest) SetQno(v string) *CloudQueryIbCdrRequest {
	s.Qno = &v
	return s
}

func (s *CloudQueryIbCdrRequest) SetStart(v int64) *CloudQueryIbCdrRequest {
	s.Start = &v
	return s
}

func (s *CloudQueryIbCdrRequest) SetStartTime(v int64) *CloudQueryIbCdrRequest {
	s.StartTime = &v
	return s
}

func (s *CloudQueryIbCdrRequest) SetStatus(v int64) *CloudQueryIbCdrRequest {
	s.Status = &v
	return s
}

func (s *CloudQueryIbCdrRequest) SetTimeRangeType(v string) *CloudQueryIbCdrRequest {
	s.TimeRangeType = &v
	return s
}

func (s *CloudQueryIbCdrRequest) SetUniqueId(v string) *CloudQueryIbCdrRequest {
	s.UniqueId = &v
	return s
}

func (s *CloudQueryIbCdrRequest) SetUserFieldValue(v string) *CloudQueryIbCdrRequest {
	s.UserFieldValue = &v
	return s
}

func (s *CloudQueryIbCdrRequest) SetUserFieldkey(v string) *CloudQueryIbCdrRequest {
	s.UserFieldkey = &v
	return s
}

func (s *CloudQueryIbCdrRequest) Validate() error {
	return dara.Validate(s)
}
