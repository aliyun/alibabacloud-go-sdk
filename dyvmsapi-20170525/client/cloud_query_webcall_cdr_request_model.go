// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudQueryWebcallCdrRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCalleeClid(v string) *CloudQueryWebcallCdrRequest
	GetCalleeClid() *string
	SetCalleeDisplayNumber(v string) *CloudQueryWebcallCdrRequest
	GetCalleeDisplayNumber() *string
	SetCalleeNumber(v string) *CloudQueryWebcallCdrRequest
	GetCalleeNumber() *string
	SetCity(v string) *CloudQueryWebcallCdrRequest
	GetCity() *string
	SetClid(v string) *CloudQueryWebcallCdrRequest
	GetClid() *string
	SetCno(v string) *CloudQueryWebcallCdrRequest
	GetCno() *string
	SetCustomerNumber(v string) *CloudQueryWebcallCdrRequest
	GetCustomerNumber() *string
	SetDisplayNumber(v string) *CloudQueryWebcallCdrRequest
	GetDisplayNumber() *string
	SetEndTime(v int64) *CloudQueryWebcallCdrRequest
	GetEndTime() *int64
	SetEnterpriseId(v int64) *CloudQueryWebcallCdrRequest
	GetEnterpriseId() *int64
	SetGno(v string) *CloudQueryWebcallCdrRequest
	GetGno() *string
	SetIsRealAnswer(v int64) *CloudQueryWebcallCdrRequest
	GetIsRealAnswer() *int64
	SetLimit(v int64) *CloudQueryWebcallCdrRequest
	GetLimit() *int64
	SetProvince(v string) *CloudQueryWebcallCdrRequest
	GetProvince() *string
	SetRequestUniqueId(v string) *CloudQueryWebcallCdrRequest
	GetRequestUniqueId() *string
	SetReturnCount(v int64) *CloudQueryWebcallCdrRequest
	GetReturnCount() *int64
	SetStart(v int64) *CloudQueryWebcallCdrRequest
	GetStart() *int64
	SetStartTime(v int64) *CloudQueryWebcallCdrRequest
	GetStartTime() *int64
	SetStatus(v int64) *CloudQueryWebcallCdrRequest
	GetStatus() *int64
	SetTimeRangeType(v string) *CloudQueryWebcallCdrRequest
	GetTimeRangeType() *string
	SetUniqueId(v string) *CloudQueryWebcallCdrRequest
	GetUniqueId() *string
	SetUserFieldValue(v string) *CloudQueryWebcallCdrRequest
	GetUserFieldValue() *string
	SetUserFieldkey(v string) *CloudQueryWebcallCdrRequest
	GetUserFieldkey() *string
}

type CloudQueryWebcallCdrRequest struct {
	// 第二侧外显号码
	//
	// example:
	//
	// 02138276106
	CalleeClid *string `json:"CalleeClid,omitempty" xml:"CalleeClid,omitempty"`
	// 第二侧真实外显号码
	//
	// example:
	//
	// 02138276106
	CalleeDisplayNumber *string `json:"CalleeDisplayNumber,omitempty" xml:"CalleeDisplayNumber,omitempty"`
	// 座席号码
	//
	// example:
	//
	// 01041009283
	CalleeNumber *string `json:"CalleeNumber,omitempty" xml:"CalleeNumber,omitempty"`
	// 客户电话归属城市；为空表示全部，如"成都"，此参数需要URLEncode
	//
	// example:
	//
	// 示例值示例值
	City *string `json:"City,omitempty" xml:"City,omitempty"`
	// 外显号码
	//
	// example:
	//
	// 02138276106
	Clid *string `json:"Clid,omitempty" xml:"Clid,omitempty"`
	// 座席工号
	//
	// example:
	//
	// 1111
	Cno *string `json:"Cno,omitempty" xml:"Cno,omitempty"`
	// 客户号码
	//
	// example:
	//
	// 01042003255
	CustomerNumber *string `json:"CustomerNumber,omitempty" xml:"CustomerNumber,omitempty"`
	// 真实外显号码，当使用AXB场景进行外呼时，真实外显号码为虚拟号
	//
	// example:
	//
	// 02138276106
	DisplayNumber *string `json:"DisplayNumber,omitempty" xml:"DisplayNumber,omitempty"`
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
	// 233
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
	// 示例值示例值
	Province *string `json:"Province,omitempty" xml:"Province,omitempty"`
	// 查询请求唯一id对应的记录信息，默认查询当前月，如requestUniqueId不属于当前月，查询时请传递查询参数 查询开始时间(startTime)
	//
	// example:
	//
	// uniq_1-162046xxxx.12
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
	// 呼叫结果；参数说明： 20 webcall, TTS合成失败； 21 webcall, 客户未接 ； 22:webcall, 客户接听； 23: webcall； 已呼叫； 24: webcall, 双方接听；
	//
	// example:
	//
	// 24
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
	// userFieldvalue
	UserFieldValue *string `json:"UserFieldValue,omitempty" xml:"UserFieldValue,omitempty"`
	// 用户自定义参数指定查询key；传递该参数时必须传递userFieldValue，仅查询一个字段，不支持模糊匹配，构造的自定义字段查询条件为:{"userFieldkey":"userFieldvalue"}，此参数需要URLEncode
	//
	// example:
	//
	// userFieldkey
	UserFieldkey *string `json:"UserFieldkey,omitempty" xml:"UserFieldkey,omitempty"`
}

func (s CloudQueryWebcallCdrRequest) String() string {
	return dara.Prettify(s)
}

func (s CloudQueryWebcallCdrRequest) GoString() string {
	return s.String()
}

func (s *CloudQueryWebcallCdrRequest) GetCalleeClid() *string {
	return s.CalleeClid
}

func (s *CloudQueryWebcallCdrRequest) GetCalleeDisplayNumber() *string {
	return s.CalleeDisplayNumber
}

func (s *CloudQueryWebcallCdrRequest) GetCalleeNumber() *string {
	return s.CalleeNumber
}

func (s *CloudQueryWebcallCdrRequest) GetCity() *string {
	return s.City
}

func (s *CloudQueryWebcallCdrRequest) GetClid() *string {
	return s.Clid
}

func (s *CloudQueryWebcallCdrRequest) GetCno() *string {
	return s.Cno
}

func (s *CloudQueryWebcallCdrRequest) GetCustomerNumber() *string {
	return s.CustomerNumber
}

func (s *CloudQueryWebcallCdrRequest) GetDisplayNumber() *string {
	return s.DisplayNumber
}

func (s *CloudQueryWebcallCdrRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *CloudQueryWebcallCdrRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *CloudQueryWebcallCdrRequest) GetGno() *string {
	return s.Gno
}

func (s *CloudQueryWebcallCdrRequest) GetIsRealAnswer() *int64 {
	return s.IsRealAnswer
}

func (s *CloudQueryWebcallCdrRequest) GetLimit() *int64 {
	return s.Limit
}

func (s *CloudQueryWebcallCdrRequest) GetProvince() *string {
	return s.Province
}

func (s *CloudQueryWebcallCdrRequest) GetRequestUniqueId() *string {
	return s.RequestUniqueId
}

func (s *CloudQueryWebcallCdrRequest) GetReturnCount() *int64 {
	return s.ReturnCount
}

func (s *CloudQueryWebcallCdrRequest) GetStart() *int64 {
	return s.Start
}

func (s *CloudQueryWebcallCdrRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *CloudQueryWebcallCdrRequest) GetStatus() *int64 {
	return s.Status
}

func (s *CloudQueryWebcallCdrRequest) GetTimeRangeType() *string {
	return s.TimeRangeType
}

func (s *CloudQueryWebcallCdrRequest) GetUniqueId() *string {
	return s.UniqueId
}

func (s *CloudQueryWebcallCdrRequest) GetUserFieldValue() *string {
	return s.UserFieldValue
}

func (s *CloudQueryWebcallCdrRequest) GetUserFieldkey() *string {
	return s.UserFieldkey
}

func (s *CloudQueryWebcallCdrRequest) SetCalleeClid(v string) *CloudQueryWebcallCdrRequest {
	s.CalleeClid = &v
	return s
}

func (s *CloudQueryWebcallCdrRequest) SetCalleeDisplayNumber(v string) *CloudQueryWebcallCdrRequest {
	s.CalleeDisplayNumber = &v
	return s
}

func (s *CloudQueryWebcallCdrRequest) SetCalleeNumber(v string) *CloudQueryWebcallCdrRequest {
	s.CalleeNumber = &v
	return s
}

func (s *CloudQueryWebcallCdrRequest) SetCity(v string) *CloudQueryWebcallCdrRequest {
	s.City = &v
	return s
}

func (s *CloudQueryWebcallCdrRequest) SetClid(v string) *CloudQueryWebcallCdrRequest {
	s.Clid = &v
	return s
}

func (s *CloudQueryWebcallCdrRequest) SetCno(v string) *CloudQueryWebcallCdrRequest {
	s.Cno = &v
	return s
}

func (s *CloudQueryWebcallCdrRequest) SetCustomerNumber(v string) *CloudQueryWebcallCdrRequest {
	s.CustomerNumber = &v
	return s
}

func (s *CloudQueryWebcallCdrRequest) SetDisplayNumber(v string) *CloudQueryWebcallCdrRequest {
	s.DisplayNumber = &v
	return s
}

func (s *CloudQueryWebcallCdrRequest) SetEndTime(v int64) *CloudQueryWebcallCdrRequest {
	s.EndTime = &v
	return s
}

func (s *CloudQueryWebcallCdrRequest) SetEnterpriseId(v int64) *CloudQueryWebcallCdrRequest {
	s.EnterpriseId = &v
	return s
}

func (s *CloudQueryWebcallCdrRequest) SetGno(v string) *CloudQueryWebcallCdrRequest {
	s.Gno = &v
	return s
}

func (s *CloudQueryWebcallCdrRequest) SetIsRealAnswer(v int64) *CloudQueryWebcallCdrRequest {
	s.IsRealAnswer = &v
	return s
}

func (s *CloudQueryWebcallCdrRequest) SetLimit(v int64) *CloudQueryWebcallCdrRequest {
	s.Limit = &v
	return s
}

func (s *CloudQueryWebcallCdrRequest) SetProvince(v string) *CloudQueryWebcallCdrRequest {
	s.Province = &v
	return s
}

func (s *CloudQueryWebcallCdrRequest) SetRequestUniqueId(v string) *CloudQueryWebcallCdrRequest {
	s.RequestUniqueId = &v
	return s
}

func (s *CloudQueryWebcallCdrRequest) SetReturnCount(v int64) *CloudQueryWebcallCdrRequest {
	s.ReturnCount = &v
	return s
}

func (s *CloudQueryWebcallCdrRequest) SetStart(v int64) *CloudQueryWebcallCdrRequest {
	s.Start = &v
	return s
}

func (s *CloudQueryWebcallCdrRequest) SetStartTime(v int64) *CloudQueryWebcallCdrRequest {
	s.StartTime = &v
	return s
}

func (s *CloudQueryWebcallCdrRequest) SetStatus(v int64) *CloudQueryWebcallCdrRequest {
	s.Status = &v
	return s
}

func (s *CloudQueryWebcallCdrRequest) SetTimeRangeType(v string) *CloudQueryWebcallCdrRequest {
	s.TimeRangeType = &v
	return s
}

func (s *CloudQueryWebcallCdrRequest) SetUniqueId(v string) *CloudQueryWebcallCdrRequest {
	s.UniqueId = &v
	return s
}

func (s *CloudQueryWebcallCdrRequest) SetUserFieldValue(v string) *CloudQueryWebcallCdrRequest {
	s.UserFieldValue = &v
	return s
}

func (s *CloudQueryWebcallCdrRequest) SetUserFieldkey(v string) *CloudQueryWebcallCdrRequest {
	s.UserFieldkey = &v
	return s
}

func (s *CloudQueryWebcallCdrRequest) Validate() error {
	return dara.Validate(s)
}
