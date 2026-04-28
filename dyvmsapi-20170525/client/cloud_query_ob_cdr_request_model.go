// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudQueryObCdrRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentName(v string) *CloudQueryObCdrRequest
	GetAgentName() *string
	SetAgentNumber(v string) *CloudQueryObCdrRequest
	GetAgentNumber() *string
	SetCallType(v int64) *CloudQueryObCdrRequest
	GetCallType() *int64
	SetCity(v string) *CloudQueryObCdrRequest
	GetCity() *string
	SetClid(v string) *CloudQueryObCdrRequest
	GetClid() *string
	SetCno(v string) *CloudQueryObCdrRequest
	GetCno() *string
	SetCustomerNumber(v string) *CloudQueryObCdrRequest
	GetCustomerNumber() *string
	SetDownGrade(v int64) *CloudQueryObCdrRequest
	GetDownGrade() *int64
	SetEndTime(v int64) *CloudQueryObCdrRequest
	GetEndTime() *int64
	SetEnterpriseId(v int64) *CloudQueryObCdrRequest
	GetEnterpriseId() *int64
	SetGno(v string) *CloudQueryObCdrRequest
	GetGno() *string
	SetIsRealAnswer(v int64) *CloudQueryObCdrRequest
	GetIsRealAnswer() *int64
	SetLeftDisplayNumber(v string) *CloudQueryObCdrRequest
	GetLeftDisplayNumber() *string
	SetLimit(v int64) *CloudQueryObCdrRequest
	GetLimit() *int64
	SetOrder(v int64) *CloudQueryObCdrRequest
	GetOrder() *int64
	SetProvince(v string) *CloudQueryObCdrRequest
	GetProvince() *string
	SetRequestUniqueId(v string) *CloudQueryObCdrRequest
	GetRequestUniqueId() *string
	SetReturnCount(v int64) *CloudQueryObCdrRequest
	GetReturnCount() *int64
	SetStart(v int64) *CloudQueryObCdrRequest
	GetStart() *int64
	SetStartTime(v int64) *CloudQueryObCdrRequest
	GetStartTime() *int64
	SetStatus(v int64) *CloudQueryObCdrRequest
	GetStatus() *int64
	SetTimeRangeType(v string) *CloudQueryObCdrRequest
	GetTimeRangeType() *string
	SetUniqueId(v string) *CloudQueryObCdrRequest
	GetUniqueId() *string
	SetUserFieldValue(v string) *CloudQueryObCdrRequest
	GetUserFieldValue() *string
	SetUserFieldkey(v string) *CloudQueryObCdrRequest
	GetUserFieldkey() *string
}

type CloudQueryObCdrRequest struct {
	// 座席姓名
	//
	// example:
	//
	// name3
	AgentName *string `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
	// 座席号码
	//
	// example:
	//
	// 41008502
	AgentNumber *string `json:"AgentNumber,omitempty" xml:"AgentNumber,omitempty"`
	// 呼叫类型；为空表示全部，可选值为：4:预览外呼 6:主叫外呼 9:内部呼叫
	//
	// example:
	//
	// 4
	CallType *int64 `json:"CallType,omitempty" xml:"CallType,omitempty"`
	// 客户电话归属城市；为空表示全部，如"成都"，此参数需要URLEncode
	//
	// example:
	//
	// 示例值
	City *string `json:"City,omitempty" xml:"City,omitempty"`
	// 客户侧外显号码过滤条件
	//
	// example:
	//
	// 4100
	Clid *string `json:"Clid,omitempty" xml:"Clid,omitempty"`
	// 座席工号，单个座席工号3-10位，支持一个或多个座席工号查询，多个座席工号以英文逗号隔开，一次最多支持100个。不支持模糊匹配。
	//
	// example:
	//
	// 1111
	Cno *string `json:"Cno,omitempty" xml:"Cno,omitempty"`
	// 客户号码，支持一个或多个客户号码查询，多个客户号码以英文逗号隔开，一次最多支持100个。不支持模糊匹配。
	//
	// example:
	//
	// 13322223333
	CustomerNumber *string `json:"CustomerNumber,omitempty" xml:"CustomerNumber,omitempty"`
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
	// 座席所属外呼组 单个外呼组号2-20位，支持一个或多个座席工号查询，多个座席工号以英文逗号隔开，一次最多支持10个。不支持模糊匹配。
	//
	// example:
	//
	// 455
	Gno *string `json:"Gno,omitempty" xml:"Gno,omitempty"`
	// 真人接听；取值说明：1-是、0-否
	//
	// example:
	//
	// 1
	IsRealAnswer *int64 `json:"IsRealAnswer,omitempty" xml:"IsRealAnswer,omitempty"`
	// 客户侧真实外显号码过滤条件，当使用虚拟号进行AXB外呼且开关打开时，实际过滤的客户侧真实外显号码是虚拟号
	//
	// example:
	//
	// 2000
	LeftDisplayNumber *string `json:"LeftDisplayNumber,omitempty" xml:"LeftDisplayNumber,omitempty"`
	// 需要取出的数据条数；大于0，最大不能超过1000，默认10
	//
	// example:
	//
	// 10
	Limit *int64 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// 排序方式；取值说明：0: 按照记录产生的时序顺序排序，1：按照记录产生的时序逆序排序，2: 按照结束时间的时序顺序排序，3：按照结束时间的时序逆序排序
	//
	// example:
	//
	// 0
	Order *int64 `json:"Order,omitempty" xml:"Order,omitempty"`
	// 客户电话归属省份；为空表示全部，如"四川"，此参数需要URLEncode
	//
	// example:
	//
	// 示例值示例值
	Province *string `json:"Province,omitempty" xml:"Province,omitempty"`
	// 查询请求唯一id对应的记录信息，默认查询当前月，如requestUniqueId不属于当前月，查询时请传递查询参数 查询开始时间(startTime)
	//
	// example:
	//
	// req_uniq_1-162046xxxx.12
	RequestUniqueId *string `json:"RequestUniqueId,omitempty" xml:"RequestUniqueId,omitempty"`
	// 是否返回总条数；取值说明：0: 不返回，1：返回，不传默认为1
	//
	// example:
	//
	// 1
	ReturnCount *int64 `json:"ReturnCount,omitempty" xml:"ReturnCount,omitempty"`
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
	// 呼叫结果；参数说明：30 座席未接听; 31 座席接听,未呼叫客户; 32 座席接听,客户未接听; 33 双方接听; 50 主叫外呼接听; 51 主叫外呼,客户未接听; 52 主叫外呼,双方接听。
	//
	// example:
	//
	// 33
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
	// 用户自定义参数指定查询value；传递该参数时必须传递userFieldKey，仅查询一个字段，不支持模糊匹配，构造的自定义字段查询条件为:{"userFieldkey":"userFieldvalue"}，此参数需要URLEncode
	//
	// example:
	//
	// {"userFieldkey":"userFieldvalue"}
	UserFieldValue *string `json:"UserFieldValue,omitempty" xml:"UserFieldValue,omitempty"`
	// 用户自定义参数指定查询key；传递该参数时必须传递userFieldValue，仅查询一个字段，不支持模糊匹配，构造的自定义字段查询条件为:{"userFieldkey":"userFieldvalue"}，此参数需要URLEncode
	//
	// example:
	//
	// {"userFieldkey":"userFieldvalue"}
	UserFieldkey *string `json:"UserFieldkey,omitempty" xml:"UserFieldkey,omitempty"`
}

func (s CloudQueryObCdrRequest) String() string {
	return dara.Prettify(s)
}

func (s CloudQueryObCdrRequest) GoString() string {
	return s.String()
}

func (s *CloudQueryObCdrRequest) GetAgentName() *string {
	return s.AgentName
}

func (s *CloudQueryObCdrRequest) GetAgentNumber() *string {
	return s.AgentNumber
}

func (s *CloudQueryObCdrRequest) GetCallType() *int64 {
	return s.CallType
}

func (s *CloudQueryObCdrRequest) GetCity() *string {
	return s.City
}

func (s *CloudQueryObCdrRequest) GetClid() *string {
	return s.Clid
}

func (s *CloudQueryObCdrRequest) GetCno() *string {
	return s.Cno
}

func (s *CloudQueryObCdrRequest) GetCustomerNumber() *string {
	return s.CustomerNumber
}

func (s *CloudQueryObCdrRequest) GetDownGrade() *int64 {
	return s.DownGrade
}

func (s *CloudQueryObCdrRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *CloudQueryObCdrRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *CloudQueryObCdrRequest) GetGno() *string {
	return s.Gno
}

func (s *CloudQueryObCdrRequest) GetIsRealAnswer() *int64 {
	return s.IsRealAnswer
}

func (s *CloudQueryObCdrRequest) GetLeftDisplayNumber() *string {
	return s.LeftDisplayNumber
}

func (s *CloudQueryObCdrRequest) GetLimit() *int64 {
	return s.Limit
}

func (s *CloudQueryObCdrRequest) GetOrder() *int64 {
	return s.Order
}

func (s *CloudQueryObCdrRequest) GetProvince() *string {
	return s.Province
}

func (s *CloudQueryObCdrRequest) GetRequestUniqueId() *string {
	return s.RequestUniqueId
}

func (s *CloudQueryObCdrRequest) GetReturnCount() *int64 {
	return s.ReturnCount
}

func (s *CloudQueryObCdrRequest) GetStart() *int64 {
	return s.Start
}

func (s *CloudQueryObCdrRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *CloudQueryObCdrRequest) GetStatus() *int64 {
	return s.Status
}

func (s *CloudQueryObCdrRequest) GetTimeRangeType() *string {
	return s.TimeRangeType
}

func (s *CloudQueryObCdrRequest) GetUniqueId() *string {
	return s.UniqueId
}

func (s *CloudQueryObCdrRequest) GetUserFieldValue() *string {
	return s.UserFieldValue
}

func (s *CloudQueryObCdrRequest) GetUserFieldkey() *string {
	return s.UserFieldkey
}

func (s *CloudQueryObCdrRequest) SetAgentName(v string) *CloudQueryObCdrRequest {
	s.AgentName = &v
	return s
}

func (s *CloudQueryObCdrRequest) SetAgentNumber(v string) *CloudQueryObCdrRequest {
	s.AgentNumber = &v
	return s
}

func (s *CloudQueryObCdrRequest) SetCallType(v int64) *CloudQueryObCdrRequest {
	s.CallType = &v
	return s
}

func (s *CloudQueryObCdrRequest) SetCity(v string) *CloudQueryObCdrRequest {
	s.City = &v
	return s
}

func (s *CloudQueryObCdrRequest) SetClid(v string) *CloudQueryObCdrRequest {
	s.Clid = &v
	return s
}

func (s *CloudQueryObCdrRequest) SetCno(v string) *CloudQueryObCdrRequest {
	s.Cno = &v
	return s
}

func (s *CloudQueryObCdrRequest) SetCustomerNumber(v string) *CloudQueryObCdrRequest {
	s.CustomerNumber = &v
	return s
}

func (s *CloudQueryObCdrRequest) SetDownGrade(v int64) *CloudQueryObCdrRequest {
	s.DownGrade = &v
	return s
}

func (s *CloudQueryObCdrRequest) SetEndTime(v int64) *CloudQueryObCdrRequest {
	s.EndTime = &v
	return s
}

func (s *CloudQueryObCdrRequest) SetEnterpriseId(v int64) *CloudQueryObCdrRequest {
	s.EnterpriseId = &v
	return s
}

func (s *CloudQueryObCdrRequest) SetGno(v string) *CloudQueryObCdrRequest {
	s.Gno = &v
	return s
}

func (s *CloudQueryObCdrRequest) SetIsRealAnswer(v int64) *CloudQueryObCdrRequest {
	s.IsRealAnswer = &v
	return s
}

func (s *CloudQueryObCdrRequest) SetLeftDisplayNumber(v string) *CloudQueryObCdrRequest {
	s.LeftDisplayNumber = &v
	return s
}

func (s *CloudQueryObCdrRequest) SetLimit(v int64) *CloudQueryObCdrRequest {
	s.Limit = &v
	return s
}

func (s *CloudQueryObCdrRequest) SetOrder(v int64) *CloudQueryObCdrRequest {
	s.Order = &v
	return s
}

func (s *CloudQueryObCdrRequest) SetProvince(v string) *CloudQueryObCdrRequest {
	s.Province = &v
	return s
}

func (s *CloudQueryObCdrRequest) SetRequestUniqueId(v string) *CloudQueryObCdrRequest {
	s.RequestUniqueId = &v
	return s
}

func (s *CloudQueryObCdrRequest) SetReturnCount(v int64) *CloudQueryObCdrRequest {
	s.ReturnCount = &v
	return s
}

func (s *CloudQueryObCdrRequest) SetStart(v int64) *CloudQueryObCdrRequest {
	s.Start = &v
	return s
}

func (s *CloudQueryObCdrRequest) SetStartTime(v int64) *CloudQueryObCdrRequest {
	s.StartTime = &v
	return s
}

func (s *CloudQueryObCdrRequest) SetStatus(v int64) *CloudQueryObCdrRequest {
	s.Status = &v
	return s
}

func (s *CloudQueryObCdrRequest) SetTimeRangeType(v string) *CloudQueryObCdrRequest {
	s.TimeRangeType = &v
	return s
}

func (s *CloudQueryObCdrRequest) SetUniqueId(v string) *CloudQueryObCdrRequest {
	s.UniqueId = &v
	return s
}

func (s *CloudQueryObCdrRequest) SetUserFieldValue(v string) *CloudQueryObCdrRequest {
	s.UserFieldValue = &v
	return s
}

func (s *CloudQueryObCdrRequest) SetUserFieldkey(v string) *CloudQueryObCdrRequest {
	s.UserFieldkey = &v
	return s
}

func (s *CloudQueryObCdrRequest) Validate() error {
	return dara.Validate(s)
}
