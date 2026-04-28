// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudQueryPredictiveCallCdrRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentName(v string) *CloudQueryPredictiveCallCdrRequest
	GetAgentName() *string
	SetCity(v string) *CloudQueryPredictiveCallCdrRequest
	GetCity() *string
	SetClid(v string) *CloudQueryPredictiveCallCdrRequest
	GetClid() *string
	SetCno(v string) *CloudQueryPredictiveCallCdrRequest
	GetCno() *string
	SetCustomerNumber(v string) *CloudQueryPredictiveCallCdrRequest
	GetCustomerNumber() *string
	SetDisplayNumber(v string) *CloudQueryPredictiveCallCdrRequest
	GetDisplayNumber() *string
	SetDownGrade(v int64) *CloudQueryPredictiveCallCdrRequest
	GetDownGrade() *int64
	SetEndTime(v int64) *CloudQueryPredictiveCallCdrRequest
	GetEndTime() *int64
	SetEnterpriseId(v int64) *CloudQueryPredictiveCallCdrRequest
	GetEnterpriseId() *int64
	SetGno(v string) *CloudQueryPredictiveCallCdrRequest
	GetGno() *string
	SetIsRealAnswer(v int64) *CloudQueryPredictiveCallCdrRequest
	GetIsRealAnswer() *int64
	SetLimit(v int64) *CloudQueryPredictiveCallCdrRequest
	GetLimit() *int64
	SetProvince(v string) *CloudQueryPredictiveCallCdrRequest
	GetProvince() *string
	SetRequestUniqueId(v string) *CloudQueryPredictiveCallCdrRequest
	GetRequestUniqueId() *string
	SetStart(v int64) *CloudQueryPredictiveCallCdrRequest
	GetStart() *int64
	SetStartTime(v int64) *CloudQueryPredictiveCallCdrRequest
	GetStartTime() *int64
	SetStatus(v int64) *CloudQueryPredictiveCallCdrRequest
	GetStatus() *int64
	SetTaskFileId(v int64) *CloudQueryPredictiveCallCdrRequest
	GetTaskFileId() *int64
	SetTaskId(v int64) *CloudQueryPredictiveCallCdrRequest
	GetTaskId() *int64
	SetTimeRangeType(v string) *CloudQueryPredictiveCallCdrRequest
	GetTimeRangeType() *string
	SetUniqueId(v string) *CloudQueryPredictiveCallCdrRequest
	GetUniqueId() *string
	SetUserFieldValue(v string) *CloudQueryPredictiveCallCdrRequest
	GetUserFieldValue() *string
	SetUserFieldkey(v string) *CloudQueryPredictiveCallCdrRequest
	GetUserFieldkey() *string
}

type CloudQueryPredictiveCallCdrRequest struct {
	// 座席姓名
	//
	// example:
	//
	// 示例值示例值
	AgentName *string `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
	// 客户电话归属城市；为空表示全部，如"成都"，此参数需要URLEncode
	//
	// example:
	//
	// 示例值示例值
	City *string `json:"City,omitempty" xml:"City,omitempty"`
	// 外显号码；外显的号码
	//
	// example:
	//
	// 02138276106
	Clid *string `json:"Clid,omitempty" xml:"Clid,omitempty"`
	// 座席工号，3-10位数字
	//
	// example:
	//
	// 1111
	Cno *string `json:"Cno,omitempty" xml:"Cno,omitempty"`
	// 客户电话；客户号码，如：18012345678
	//
	// example:
	//
	// 18012345678
	CustomerNumber *string `json:"CustomerNumber,omitempty" xml:"CustomerNumber,omitempty"`
	// 客户侧真实外显号码，当使用AXB场景进行外呼时，客户侧的真实外显号码为虚拟号
	//
	// example:
	//
	// 02138276106
	DisplayNumber *string `json:"DisplayNumber,omitempty" xml:"DisplayNumber,omitempty"`
	// 是否呼叫降级；0-否, 1-是
	//
	// example:
	//
	// 0
	DownGrade *int64 `json:"DownGrade,omitempty" xml:"DownGrade,omitempty"`
	// 查询结束时间；时间戳格式,精确到s。startTime与endTime不允许跨月，默认值取当前月份最后一天
	//
	// example:
	//
	// 1775024775
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 呼叫中心 id
	//
	// This parameter is required.
	//
	// example:
	//
	// 7000002
	EnterpriseId *int64 `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	// 座席所属外呼组 单个外呼组号2-20位，支持一个或多个座席工号查询，多个座席工号以英文逗号隔开，一次最多支持10个。不支持模糊匹配。
	//
	// example:
	//
	// WH13
	Gno *string `json:"Gno,omitempty" xml:"Gno,omitempty"`
	// 真人接听；取值说明：1-是、0-否
	//
	// example:
	//
	// 1
	IsRealAnswer *int64 `json:"IsRealAnswer,omitempty" xml:"IsRealAnswer,omitempty"`
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
	// 示例值示例值示例值
	Province *string `json:"Province,omitempty" xml:"Province,omitempty"`
	// 查询请求唯一id对应的记录信息，默认查询当前月，如requestUniqueId不属于当前月，查询时请传递查询参数 查询开始时间(startTime)
	//
	// example:
	//
	// req-uniq_1-162046xxxx.12
	RequestUniqueId *string `json:"RequestUniqueId,omitempty" xml:"RequestUniqueId,omitempty"`
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
	// 呼叫结果；参数说明：40:预测外呼, 客户未接听； 41:预测外呼, 客户接听； 42: 预测外呼,已呼叫； 43: 预测外呼, 双方接听
	//
	// example:
	//
	// 41
	Status *int64 `json:"Status,omitempty" xml:"Status,omitempty"`
	// 批次id
	//
	// example:
	//
	// 133
	TaskFileId *int64 `json:"TaskFileId,omitempty" xml:"TaskFileId,omitempty"`
	// 外呼任务id
	//
	// example:
	//
	// 31111
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
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
	// 用户自定义参数指定查询value；传递该参数时必须传递userFieldKey，仅查询一个字段，不支持模糊匹配，构造的自定义字段查询条件为:{"userFieldkey":"userFieldvalue"}，此参数需要URLEncode
	//
	// example:
	//
	// userFieldvalue
	UserFieldValue *string `json:"UserFieldValue,omitempty" xml:"UserFieldValue,omitempty"`
	// 用户自定义参数指定查询key；传递该参数时必须传递userFieldValue，仅查询一个字段，不支持模糊匹配，构造的自定义字段查询条件为:{"userFieldkey":"userFieldvalue"}，此参数需要URLEncode
	//
	// example:
	//
	// userFieldkey
	UserFieldkey *string `json:"UserFieldkey,omitempty" xml:"UserFieldkey,omitempty"`
}

func (s CloudQueryPredictiveCallCdrRequest) String() string {
	return dara.Prettify(s)
}

func (s CloudQueryPredictiveCallCdrRequest) GoString() string {
	return s.String()
}

func (s *CloudQueryPredictiveCallCdrRequest) GetAgentName() *string {
	return s.AgentName
}

func (s *CloudQueryPredictiveCallCdrRequest) GetCity() *string {
	return s.City
}

func (s *CloudQueryPredictiveCallCdrRequest) GetClid() *string {
	return s.Clid
}

func (s *CloudQueryPredictiveCallCdrRequest) GetCno() *string {
	return s.Cno
}

func (s *CloudQueryPredictiveCallCdrRequest) GetCustomerNumber() *string {
	return s.CustomerNumber
}

func (s *CloudQueryPredictiveCallCdrRequest) GetDisplayNumber() *string {
	return s.DisplayNumber
}

func (s *CloudQueryPredictiveCallCdrRequest) GetDownGrade() *int64 {
	return s.DownGrade
}

func (s *CloudQueryPredictiveCallCdrRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *CloudQueryPredictiveCallCdrRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *CloudQueryPredictiveCallCdrRequest) GetGno() *string {
	return s.Gno
}

func (s *CloudQueryPredictiveCallCdrRequest) GetIsRealAnswer() *int64 {
	return s.IsRealAnswer
}

func (s *CloudQueryPredictiveCallCdrRequest) GetLimit() *int64 {
	return s.Limit
}

func (s *CloudQueryPredictiveCallCdrRequest) GetProvince() *string {
	return s.Province
}

func (s *CloudQueryPredictiveCallCdrRequest) GetRequestUniqueId() *string {
	return s.RequestUniqueId
}

func (s *CloudQueryPredictiveCallCdrRequest) GetStart() *int64 {
	return s.Start
}

func (s *CloudQueryPredictiveCallCdrRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *CloudQueryPredictiveCallCdrRequest) GetStatus() *int64 {
	return s.Status
}

func (s *CloudQueryPredictiveCallCdrRequest) GetTaskFileId() *int64 {
	return s.TaskFileId
}

func (s *CloudQueryPredictiveCallCdrRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *CloudQueryPredictiveCallCdrRequest) GetTimeRangeType() *string {
	return s.TimeRangeType
}

func (s *CloudQueryPredictiveCallCdrRequest) GetUniqueId() *string {
	return s.UniqueId
}

func (s *CloudQueryPredictiveCallCdrRequest) GetUserFieldValue() *string {
	return s.UserFieldValue
}

func (s *CloudQueryPredictiveCallCdrRequest) GetUserFieldkey() *string {
	return s.UserFieldkey
}

func (s *CloudQueryPredictiveCallCdrRequest) SetAgentName(v string) *CloudQueryPredictiveCallCdrRequest {
	s.AgentName = &v
	return s
}

func (s *CloudQueryPredictiveCallCdrRequest) SetCity(v string) *CloudQueryPredictiveCallCdrRequest {
	s.City = &v
	return s
}

func (s *CloudQueryPredictiveCallCdrRequest) SetClid(v string) *CloudQueryPredictiveCallCdrRequest {
	s.Clid = &v
	return s
}

func (s *CloudQueryPredictiveCallCdrRequest) SetCno(v string) *CloudQueryPredictiveCallCdrRequest {
	s.Cno = &v
	return s
}

func (s *CloudQueryPredictiveCallCdrRequest) SetCustomerNumber(v string) *CloudQueryPredictiveCallCdrRequest {
	s.CustomerNumber = &v
	return s
}

func (s *CloudQueryPredictiveCallCdrRequest) SetDisplayNumber(v string) *CloudQueryPredictiveCallCdrRequest {
	s.DisplayNumber = &v
	return s
}

func (s *CloudQueryPredictiveCallCdrRequest) SetDownGrade(v int64) *CloudQueryPredictiveCallCdrRequest {
	s.DownGrade = &v
	return s
}

func (s *CloudQueryPredictiveCallCdrRequest) SetEndTime(v int64) *CloudQueryPredictiveCallCdrRequest {
	s.EndTime = &v
	return s
}

func (s *CloudQueryPredictiveCallCdrRequest) SetEnterpriseId(v int64) *CloudQueryPredictiveCallCdrRequest {
	s.EnterpriseId = &v
	return s
}

func (s *CloudQueryPredictiveCallCdrRequest) SetGno(v string) *CloudQueryPredictiveCallCdrRequest {
	s.Gno = &v
	return s
}

func (s *CloudQueryPredictiveCallCdrRequest) SetIsRealAnswer(v int64) *CloudQueryPredictiveCallCdrRequest {
	s.IsRealAnswer = &v
	return s
}

func (s *CloudQueryPredictiveCallCdrRequest) SetLimit(v int64) *CloudQueryPredictiveCallCdrRequest {
	s.Limit = &v
	return s
}

func (s *CloudQueryPredictiveCallCdrRequest) SetProvince(v string) *CloudQueryPredictiveCallCdrRequest {
	s.Province = &v
	return s
}

func (s *CloudQueryPredictiveCallCdrRequest) SetRequestUniqueId(v string) *CloudQueryPredictiveCallCdrRequest {
	s.RequestUniqueId = &v
	return s
}

func (s *CloudQueryPredictiveCallCdrRequest) SetStart(v int64) *CloudQueryPredictiveCallCdrRequest {
	s.Start = &v
	return s
}

func (s *CloudQueryPredictiveCallCdrRequest) SetStartTime(v int64) *CloudQueryPredictiveCallCdrRequest {
	s.StartTime = &v
	return s
}

func (s *CloudQueryPredictiveCallCdrRequest) SetStatus(v int64) *CloudQueryPredictiveCallCdrRequest {
	s.Status = &v
	return s
}

func (s *CloudQueryPredictiveCallCdrRequest) SetTaskFileId(v int64) *CloudQueryPredictiveCallCdrRequest {
	s.TaskFileId = &v
	return s
}

func (s *CloudQueryPredictiveCallCdrRequest) SetTaskId(v int64) *CloudQueryPredictiveCallCdrRequest {
	s.TaskId = &v
	return s
}

func (s *CloudQueryPredictiveCallCdrRequest) SetTimeRangeType(v string) *CloudQueryPredictiveCallCdrRequest {
	s.TimeRangeType = &v
	return s
}

func (s *CloudQueryPredictiveCallCdrRequest) SetUniqueId(v string) *CloudQueryPredictiveCallCdrRequest {
	s.UniqueId = &v
	return s
}

func (s *CloudQueryPredictiveCallCdrRequest) SetUserFieldValue(v string) *CloudQueryPredictiveCallCdrRequest {
	s.UserFieldValue = &v
	return s
}

func (s *CloudQueryPredictiveCallCdrRequest) SetUserFieldkey(v string) *CloudQueryPredictiveCallCdrRequest {
	s.UserFieldkey = &v
	return s
}

func (s *CloudQueryPredictiveCallCdrRequest) Validate() error {
	return dara.Validate(s)
}
