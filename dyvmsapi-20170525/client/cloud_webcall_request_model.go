// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudWebcallRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAmd(v int64) *CloudWebcallRequest
	GetAmd() *int64
	SetClid(v string) *CloudWebcallRequest
	GetClid() *string
	SetClidAreaCode(v string) *CloudWebcallRequest
	GetClidAreaCode() *string
	SetClidGroup(v string) *CloudWebcallRequest
	GetClidGroup() *string
	SetClidList(v string) *CloudWebcallRequest
	GetClidList() *string
	SetCrnId(v string) *CloudWebcallRequest
	GetCrnId() *string
	SetDelay(v int64) *CloudWebcallRequest
	GetDelay() *int64
	SetEnterpriseId(v int64) *CloudWebcallRequest
	GetEnterpriseId() *int64
	SetExpirTime(v string) *CloudWebcallRequest
	GetExpirTime() *string
	SetIvrId(v int64) *CloudWebcallRequest
	GetIvrId() *int64
	SetMultiTelDelay(v int64) *CloudWebcallRequest
	GetMultiTelDelay() *int64
	SetMultiTelPush(v int64) *CloudWebcallRequest
	GetMultiTelPush() *int64
	SetRequestUniqueId(v string) *CloudWebcallRequest
	GetRequestUniqueId() *string
	SetRetry(v int64) *CloudWebcallRequest
	GetRetry() *int64
	SetRetryInterval(v int64) *CloudWebcallRequest
	GetRetryInterval() *int64
	SetTel(v string) *CloudWebcallRequest
	GetTel() *string
	SetTimeout(v int64) *CloudWebcallRequest
	GetTimeout() *int64
	SetUserField(v string) *CloudWebcallRequest
	GetUserField() *string
	SetVid(v string) *CloudWebcallRequest
	GetVid() *string
}

type CloudWebcallRequest struct {
	// 是否开启amd；自动应答检查（传真机等），1.开启 0.不开启 默认为0
	//
	// example:
	//
	// 0
	Amd *int64 `json:"Amd,omitempty" xml:"Amd,omitempty"`
	// 可传入企业中继号码或设置好的客户侧外显号码
	//
	// example:
	//
	// 41008502
	Clid *string `json:"Clid,omitempty" xml:"Clid,omitempty"`
	// 客户侧外显指定地区号码，优先级低于clid
	//
	// example:
	//
	// 010
	ClidAreaCode *string `json:"ClidAreaCode,omitempty" xml:"ClidAreaCode,omitempty"`
	// 客户侧外显号码组；使用clidGroup需要账号支持按标识路由，优先级低于clid
	//
	// example:
	//
	// clidGroup
	ClidGroup *string `json:"ClidGroup,omitempty" xml:"ClidGroup,omitempty"`
	// 指定本次外呼使用的客户侧外显号码集合，系统将根据号码调度策略选择可用号码进行外呼
	//
	// example:
	//
	// 41008502
	ClidList *string `json:"ClidList,omitempty" xml:"ClidList,omitempty"`
	// 外显导航标识
	//
	// example:
	//
	// 33
	CrnId *string `json:"CrnId,omitempty" xml:"CrnId,omitempty"`
	// 延迟时长；秒数，延迟多少秒呼叫 参数取值范围：0<=delay<=60 默认为0
	//
	// example:
	//
	// 20
	Delay *int64 `json:"Delay,omitempty" xml:"Delay,omitempty"`
	// 呼叫中心 id
	//
	// This parameter is required.
	//
	// example:
	//
	// 7000002
	EnterpriseId *int64 `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	// 过期时间，精确到s，yyyy-MM-dd HH:mm:ss
	//
	// example:
	//
	// 2026-04-20 10:00:00
	ExpirTime *string `json:"ExpirTime,omitempty" xml:"ExpirTime,omitempty"`
	// 回呼接通后进入的ivrId；如果不传此参数，使用后台配置的ivr
	//
	// example:
	//
	// 70
	IvrId *int64 `json:"IvrId,omitempty" xml:"IvrId,omitempty"`
	// 多号码时呼叫延时；tel多个号码时，号码之间的呼叫延迟，最大120秒，默认0
	//
	// example:
	//
	// 0
	MultiTelDelay *int64 `json:"MultiTelDelay,omitempty" xml:"MultiTelDelay,omitempty"`
	// 多号码，是否每个号码呼叫都挂机推送；取值说明 0：只挂机推送一次； 1：每次呼叫都触发挂机推送；默认0
	//
	// example:
	//
	// 0
	MultiTelPush *int64 `json:"MultiTelPush,omitempty" xml:"MultiTelPush,omitempty"`
	// 请求唯一标识；说明：长度不超过300个字节。1个汉字是3字节。此字段保存到通话记录requestUniqueId字段，后续接口查询使用。一次接口请求造成的多次呼叫requestUniqueId相同；如果是加密的号码，需要URLEncode后传入
	//
	// example:
	//
	// requset12345
	RequestUniqueId *string `json:"RequestUniqueId,omitempty" xml:"RequestUniqueId,omitempty"`
	// 重试次数；最大5次，使用的企业需要开启webcall自动重试开关
	//
	// example:
	//
	// 2
	Retry *int64 `json:"Retry,omitempty" xml:"Retry,omitempty"`
	// 重试间隔，单位分钟，最小0分钟，最大59分钟
	//
	// example:
	//
	// 3
	RetryInterval *int64 `json:"RetryInterval,omitempty" xml:"RetryInterval,omitempty"`
	// 电话号码；固话带区号，手机不加0；固话带分机的以\\"-\\"分隔；支持多个号码，多个直接用英文逗号’,’分隔；传多个号码时，若前面的号码没接通则会呼叫后面的号码，若接通则不会呼叫；[加密外呼](../字段定义/接口部分/加密外呼号码加密规则.md)；如果是加密的号码，需要URLEncode后传入
	//
	// This parameter is required.
	//
	// example:
	//
	// 13312345678
	Tel *string `json:"Tel,omitempty" xml:"Tel,omitempty"`
	// 呼叫客户超时时间；说明：参数取值范围 0<=timeout<=60；不传入，默认30(单位:s)
	//
	// example:
	//
	// 30
	Timeout *int64 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// 用户自定义字段；说明：长度不超过250个字节。1个汉字是3字节。此字段保存到通话记录userField字段，后续接口查询使用。该字段需“动态附带参数”paramNames有值时方可生效。
	//
	// example:
	//
	// UserField
	UserField *string `json:"UserField,omitempty" xml:"UserField,omitempty"`
	// 用哪种语言播放ivr提示音；1.普通话 2.粤语 4.标贝TTS 5.普通话-甜美女音 默认为1
	//
	// example:
	//
	// 1
	Vid *string `json:"Vid,omitempty" xml:"Vid,omitempty"`
}

func (s CloudWebcallRequest) String() string {
	return dara.Prettify(s)
}

func (s CloudWebcallRequest) GoString() string {
	return s.String()
}

func (s *CloudWebcallRequest) GetAmd() *int64 {
	return s.Amd
}

func (s *CloudWebcallRequest) GetClid() *string {
	return s.Clid
}

func (s *CloudWebcallRequest) GetClidAreaCode() *string {
	return s.ClidAreaCode
}

func (s *CloudWebcallRequest) GetClidGroup() *string {
	return s.ClidGroup
}

func (s *CloudWebcallRequest) GetClidList() *string {
	return s.ClidList
}

func (s *CloudWebcallRequest) GetCrnId() *string {
	return s.CrnId
}

func (s *CloudWebcallRequest) GetDelay() *int64 {
	return s.Delay
}

func (s *CloudWebcallRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *CloudWebcallRequest) GetExpirTime() *string {
	return s.ExpirTime
}

func (s *CloudWebcallRequest) GetIvrId() *int64 {
	return s.IvrId
}

func (s *CloudWebcallRequest) GetMultiTelDelay() *int64 {
	return s.MultiTelDelay
}

func (s *CloudWebcallRequest) GetMultiTelPush() *int64 {
	return s.MultiTelPush
}

func (s *CloudWebcallRequest) GetRequestUniqueId() *string {
	return s.RequestUniqueId
}

func (s *CloudWebcallRequest) GetRetry() *int64 {
	return s.Retry
}

func (s *CloudWebcallRequest) GetRetryInterval() *int64 {
	return s.RetryInterval
}

func (s *CloudWebcallRequest) GetTel() *string {
	return s.Tel
}

func (s *CloudWebcallRequest) GetTimeout() *int64 {
	return s.Timeout
}

func (s *CloudWebcallRequest) GetUserField() *string {
	return s.UserField
}

func (s *CloudWebcallRequest) GetVid() *string {
	return s.Vid
}

func (s *CloudWebcallRequest) SetAmd(v int64) *CloudWebcallRequest {
	s.Amd = &v
	return s
}

func (s *CloudWebcallRequest) SetClid(v string) *CloudWebcallRequest {
	s.Clid = &v
	return s
}

func (s *CloudWebcallRequest) SetClidAreaCode(v string) *CloudWebcallRequest {
	s.ClidAreaCode = &v
	return s
}

func (s *CloudWebcallRequest) SetClidGroup(v string) *CloudWebcallRequest {
	s.ClidGroup = &v
	return s
}

func (s *CloudWebcallRequest) SetClidList(v string) *CloudWebcallRequest {
	s.ClidList = &v
	return s
}

func (s *CloudWebcallRequest) SetCrnId(v string) *CloudWebcallRequest {
	s.CrnId = &v
	return s
}

func (s *CloudWebcallRequest) SetDelay(v int64) *CloudWebcallRequest {
	s.Delay = &v
	return s
}

func (s *CloudWebcallRequest) SetEnterpriseId(v int64) *CloudWebcallRequest {
	s.EnterpriseId = &v
	return s
}

func (s *CloudWebcallRequest) SetExpirTime(v string) *CloudWebcallRequest {
	s.ExpirTime = &v
	return s
}

func (s *CloudWebcallRequest) SetIvrId(v int64) *CloudWebcallRequest {
	s.IvrId = &v
	return s
}

func (s *CloudWebcallRequest) SetMultiTelDelay(v int64) *CloudWebcallRequest {
	s.MultiTelDelay = &v
	return s
}

func (s *CloudWebcallRequest) SetMultiTelPush(v int64) *CloudWebcallRequest {
	s.MultiTelPush = &v
	return s
}

func (s *CloudWebcallRequest) SetRequestUniqueId(v string) *CloudWebcallRequest {
	s.RequestUniqueId = &v
	return s
}

func (s *CloudWebcallRequest) SetRetry(v int64) *CloudWebcallRequest {
	s.Retry = &v
	return s
}

func (s *CloudWebcallRequest) SetRetryInterval(v int64) *CloudWebcallRequest {
	s.RetryInterval = &v
	return s
}

func (s *CloudWebcallRequest) SetTel(v string) *CloudWebcallRequest {
	s.Tel = &v
	return s
}

func (s *CloudWebcallRequest) SetTimeout(v int64) *CloudWebcallRequest {
	s.Timeout = &v
	return s
}

func (s *CloudWebcallRequest) SetUserField(v string) *CloudWebcallRequest {
	s.UserField = &v
	return s
}

func (s *CloudWebcallRequest) SetVid(v string) *CloudWebcallRequest {
	s.Vid = &v
	return s
}

func (s *CloudWebcallRequest) Validate() error {
	return dara.Validate(s)
}
