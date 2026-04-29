// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudGetTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CloudGetTaskResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CloudGetTaskResponseBody
	GetCode() *string
	SetData(v *CloudGetTaskResponseBodyData) *CloudGetTaskResponseBody
	GetData() *CloudGetTaskResponseBodyData
	SetMessage(v string) *CloudGetTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *CloudGetTaskResponseBody
	GetRequestId() *string
}

type CloudGetTaskResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CloudGetTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 7BF47617-7851-48F7-A3A1-2021342A78E2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloudGetTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloudGetTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CloudGetTaskResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CloudGetTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *CloudGetTaskResponseBody) GetData() *CloudGetTaskResponseBodyData {
	return s.Data
}

func (s *CloudGetTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CloudGetTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloudGetTaskResponseBody) SetAccessDeniedDetail(v string) *CloudGetTaskResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CloudGetTaskResponseBody) SetCode(v string) *CloudGetTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CloudGetTaskResponseBody) SetData(v *CloudGetTaskResponseBodyData) *CloudGetTaskResponseBody {
	s.Data = v
	return s
}

func (s *CloudGetTaskResponseBody) SetMessage(v string) *CloudGetTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CloudGetTaskResponseBody) SetRequestId(v string) *CloudGetTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloudGetTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CloudGetTaskResponseBodyData struct {
	// 任务配置信息
	TaskProperty *CloudGetTaskResponseBodyDataTaskProperty `json:"TaskProperty,omitempty" xml:"TaskProperty,omitempty" type:"Struct"`
}

func (s CloudGetTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CloudGetTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *CloudGetTaskResponseBodyData) GetTaskProperty() *CloudGetTaskResponseBodyDataTaskProperty {
	return s.TaskProperty
}

func (s *CloudGetTaskResponseBodyData) SetTaskProperty(v *CloudGetTaskResponseBodyDataTaskProperty) *CloudGetTaskResponseBodyData {
	s.TaskProperty = v
	return s
}

func (s *CloudGetTaskResponseBodyData) Validate() error {
	if s.TaskProperty != nil {
		if err := s.TaskProperty.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CloudGetTaskResponseBodyDataTaskProperty struct {
	// adjustStep
	//
	// example:
	//
	// 5
	AdjustStep *string `json:"AdjustStep,omitempty" xml:"AdjustStep,omitempty"`
	// 座席超时时间，单位秒
	//
	// example:
	//
	// 10
	AgentTimeout *string `json:"AgentTimeout,omitempty" xml:"AgentTimeout,omitempty"`
	// 已废弃。座席利用率
	//
	// example:
	//
	// deprecated
	AgentUtilization *string `json:"AgentUtilization,omitempty" xml:"AgentUtilization,omitempty"`
	// 初始化预计客户接通率
	//
	// example:
	//
	// 100
	AnswerRate *string `json:"AnswerRate,omitempty" xml:"AnswerRate,omitempty"`
	// 自动完成，0.关闭 1.开启
	//
	// example:
	//
	// 0
	AutoComplete *string `json:"AutoComplete,omitempty" xml:"AutoComplete,omitempty"`
	// 定时开始，0.关闭 1.开启
	//
	// example:
	//
	// 0
	AutoStart *string `json:"AutoStart,omitempty" xml:"AutoStart,omitempty"`
	// 定时开始日期，格式：yyyy-MM-dd，如：2017-02-11
	//
	// example:
	//
	// 2026-02-11
	AutoStartDay *string `json:"AutoStartDay,omitempty" xml:"AutoStartDay,omitempty"`
	// 定时开始时间，格式：HH:mm:ss，如：08:00:00
	//
	// example:
	//
	// 08:00:00
	AutoStartTime *string `json:"AutoStartTime,omitempty" xml:"AutoStartTime,omitempty"`
	// 定时完成，0.关闭 1.开启
	//
	// example:
	//
	// 0
	AutoStop *string `json:"AutoStop,omitempty" xml:"AutoStop,omitempty"`
	// 定时完成日期，格式：yyyy-MM-dd，如：2017-02-11
	//
	// example:
	//
	// 2026-02-11
	AutoStopDay *string `json:"AutoStopDay,omitempty" xml:"AutoStopDay,omitempty"`
	// 定时完成时间，格式：HH:mm:ss，如：17:00:00
	//
	// example:
	//
	// 17:00:00
	AutoStopTime *string `json:"AutoStopTime,omitempty" xml:"AutoStopTime,omitempty"`
	// 呼叫方式，0.连续呼叫 1.间隔呼叫
	//
	// example:
	//
	// 0
	AutoTaskType *string `json:"AutoTaskType,omitempty" xml:"AutoTaskType,omitempty"`
	// 间隔呼叫时间段，呼叫的时间段配置。参数格式：时间条件编号。支持多个，多个使用英文","逗号隔开，如：9,22。
	//
	// example:
	//
	// 9,22
	AutoTriggerTimeStrategy *string `json:"AutoTriggerTimeStrategy,omitempty" xml:"AutoTriggerTimeStrategy,omitempty"`
	// 座席当日接听数限制，JsonArray格式[{"cnos":["2000","2001"],"limit":"5"}]
	//
	// example:
	//
	// [{"cnos":["2000","2001"],"limit":"5"}]
	CallLimitStrategy *string `json:"CallLimitStrategy,omitempty" xml:"CallLimitStrategy,omitempty"`
	// 座席分配规则，1.随机 2.顺序 3.座席未进线最大时（从上一次呼叫结束到当前的总时长） 4.座席状态[空闲]最长时长（座席最近一次切换为空闲状态的持续时长）
	//
	// example:
	//
	// 2
	CallStrategy *string `json:"CallStrategy,omitempty" xml:"CallStrategy,omitempty"`
	// 座席通道变量，JsonArray格式[{"key1":"value1"},{"key2":"value2"}]
	//
	// example:
	//
	// [{"key1":"value1"},{"key2":"value2"}]
	CallVariables *string `json:"CallVariables,omitempty" xml:"CallVariables,omitempty"`
	// 已废弃。给企业推送任务状态的URL地址
	//
	// example:
	//
	// deprecated
	CallbackUrl *string `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
	// 已废弃。热线号码组
	//
	// example:
	//
	// deprecated
	ClidGroup *string `json:"ClidGroup,omitempty" xml:"ClidGroup,omitempty"`
	// 外显规则，Json格式，外显号码选择规则：默认区号，是否匹配省会等；如：{"defaultAreaCode":"010", "isMatchCapital":"1"}
	//
	// example:
	//
	// {"defaultAreaCode":"010", "isMatchCapital":"1"}
	ClidProperty *string `json:"ClidProperty,omitempty" xml:"ClidProperty,omitempty"`
	// 座席工号列表
	//
	// example:
	//
	// 111,222
	Cnos *string `json:"Cnos,omitempty" xml:"Cnos,omitempty"`
	// 最大并发限制
	//
	// example:
	//
	// 0
	Concurrency *string `json:"Concurrency,omitempty" xml:"Concurrency,omitempty"`
	// 任务创建时间 ，格式为： yyyy-MM-dd HH:mm:ss
	//
	// example:
	//
	// 2025-05-10 14:07:22
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 已废弃。推送重试次数 最大5次
	//
	// example:
	//
	// 5
	CurlRetryTimes *string `json:"CurlRetryTimes,omitempty" xml:"CurlRetryTimes,omitempty"`
	// 客户侧外显规则，1.随机 2.按区号
	//
	// example:
	//
	// 2
	CustomerClidType *string `json:"CustomerClidType,omitempty" xml:"CustomerClidType,omitempty"`
	// 外显号码比例配置内容，JsonArray格式[{"number":"123,345","weight":"5"},{"number":"567,789","weight":"5"}]
	//
	// example:
	//
	// [{"number":"123,345","weight":"5"},{"number":"567,789","weight":"5"}]
	CustomerClidWeight *string `json:"CustomerClidWeight,omitempty" xml:"CustomerClidWeight,omitempty"`
	// 外显号码比例配置开关，0.关闭 1.开启
	//
	// example:
	//
	// 0
	CustomerClidWeightFlag *string `json:"CustomerClidWeightFlag,omitempty" xml:"CustomerClidWeightFlag,omitempty"`
	// 客户侧外显号码列表
	//
	// example:
	//
	// ""
	CustomerClids *string `json:"CustomerClids,omitempty" xml:"CustomerClids,omitempty"`
	// 客户侧外显号码类型，1.外显固话 2.外显手机号 4.外显号码池 5.外显导航
	//
	// example:
	//
	// 1
	CustomerClidsCategory *string `json:"CustomerClidsCategory,omitempty" xml:"CustomerClidsCategory,omitempty"`
	// 客户侧等待音
	//
	// example:
	//
	// no
	CustomerMoh *string `json:"CustomerMoh,omitempty" xml:"CustomerMoh,omitempty"`
	// 客户超时时间，单位秒
	//
	// example:
	//
	// 30
	CustomerTimeout *string `json:"CustomerTimeout,omitempty" xml:"CustomerTimeout,omitempty"`
	// 客户侧提示音
	//
	// example:
	//
	// 示例值
	CustomerVoice *string `json:"CustomerVoice,omitempty" xml:"CustomerVoice,omitempty"`
	// 任务描述
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 企业编号
	//
	// example:
	//
	// 7000002
	EnterpriseId *string `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	// 定时完成时强制结束任务。开启配置后，当任务到定时完成时间时无论任务中的号码是否已经呼完，均将任务状态设置为结束。适用于对数据有呼叫时间限制的场景。注：任务在结束状态时无法再次启动。
	//
	// example:
	//
	// 0
	ForceEndFlag *string `json:"ForceEndFlag,omitempty" xml:"ForceEndFlag,omitempty"`
	// 外呼任务Id
	//
	// example:
	//
	// 11059
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 随机呼叫，0.关闭 1.开启
	//
	// example:
	//
	// 0
	IsRandom *string `json:"IsRandom,omitempty" xml:"IsRandom,omitempty"`
	// 暂停后重新预热，0.关闭，1开启
	//
	// example:
	//
	// 0
	IsRewarm *string `json:"IsRewarm,omitempty" xml:"IsRewarm,omitempty"`
	// 语音导航Id
	//
	// example:
	//
	// 33
	IvrId *string `json:"IvrId,omitempty" xml:"IvrId,omitempty"`
	// 座席最大等待时间，单位秒
	//
	// example:
	//
	// 10
	MaxWaitTime *string `json:"MaxWaitTime,omitempty" xml:"MaxWaitTime,omitempty"`
	// 最小可用座席数
	//
	// example:
	//
	// 1
	MinAvailableAgentCount *string `json:"MinAvailableAgentCount,omitempty" xml:"MinAvailableAgentCount,omitempty"`
	// 任务名称
	//
	// example:
	//
	// name1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// ""
	OverCallLimitCnos *string `json:"OverCallLimitCnos,omitempty" xml:"OverCallLimitCnos,omitempty"`
	// 任务结束时间
	//
	// example:
	//
	// 示例值示例值示例值
	OverTime *string `json:"OverTime,omitempty" xml:"OverTime,omitempty"`
	// 任务暂停时长(针对自动启动时间段) 单位分钟
	//
	// example:
	//
	// 0
	PauseDuration *string `json:"PauseDuration,omitempty" xml:"PauseDuration,omitempty"`
	// 任务暂停时间点
	//
	// example:
	//
	// 2026-04-14 16:51:11
	PauseTime *string `json:"PauseTime,omitempty" xml:"PauseTime,omitempty"`
	// 超呼率
	//
	// example:
	//
	// 100
	PredictAdjust *string `json:"PredictAdjust,omitempty" xml:"PredictAdjust,omitempty"`
	// 骚扰率，支持小数，精确到小数点两位
	//
	// example:
	//
	// 3.0
	Quotiety *string `json:"Quotiety,omitempty" xml:"Quotiety,omitempty"`
	// 重试策略，「基础模式」数据结构为Json格式```{"retry":2,"strategy":[{"round":1,"time":"1-1-1"},{"round":2,"time":"2-2-2"}]}```格式说明：retry是重试次数，round是第几次重试，time是第几次重试间隔的时间：1-1-1，第一个1是天，第二个1是小时，第三个1是分钟「高级模式」高级模式需要开启「号码状态识别」服务后。目前只支持「自动外呼」任务模式。数据结构为JsonArray格式 ```[{"condition":{"sipCause":[710,719]},"retry":1,"sort":1,"strategy":[{"round":1,"time":"0-0-10"}]},{"condition":{"sipCause":[800]},"retry":1,"sort":2,"strategy":[{"round":2,"time":"0-0-10"}]}]```格式说明：condition是重试条件，sort是重试条件的顺序，其余字段与基础模式一致
	//
	// example:
	//
	// {"retry":2,"strategy":[{"round":1,"time":"1-1-1"},{"round":2,"time":"2-2-2"}]}
	RetryStrategy *string `json:"RetryStrategy,omitempty" xml:"RetryStrategy,omitempty"`
	// 重试仅当天生效，0.关闭 1.开启，删除待重试的数据和待呼叫的数据 2.开启，删除待重试的数据 3.开启，删除待呼叫的数据
	//
	// example:
	//
	// 0
	RetryStrategyOnlyToday *string `json:"RetryStrategyOnlyToday,omitempty" xml:"RetryStrategyOnlyToday,omitempty"`
	// 已废弃
	//
	// example:
	//
	// deprecated
	StandbyCnos *string `json:"StandbyCnos,omitempty" xml:"StandbyCnos,omitempty"`
	// 任务开始时间
	//
	// example:
	//
	// 2025-05-10 14:11:15
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 任务状态，0.初始 1.运行中 2.暂停 3.结束
	//
	// example:
	//
	// 3
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 任务状态描述
	//
	// example:
	//
	// 示例值示例值
	StatusDescription *string `json:"StatusDescription,omitempty" xml:"StatusDescription,omitempty"`
	// 已废弃。任务级呼叫属性 json格式{"key1":"value1","key2":"value2"}，仅做呼叫控制用，不放通道变量
	//
	// example:
	//
	// deprecated
	TaskFields *string `json:"TaskFields,omitempty" xml:"TaskFields,omitempty"`
	// 禁止呼叫时间限制，在特定的时间段内禁止呼叫。参数格式：时间条件编号。支持多个，多个使用英文","逗号隔开，如：9,22。
	//
	// example:
	//
	// 9,22
	TimeStrategy *string `json:"TimeStrategy,omitempty" xml:"TimeStrategy,omitempty"`
	// 任务类型，1.预测外呼任务 2.自动外呼任务
	//
	// example:
	//
	// 1
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// 任务自定义字段，JsonArray格式[{"key1":"value1"},{"key2":"value2"}]
	//
	// example:
	//
	// [{"key1":"value1"},{"key2":"value2"}]
	UserFields *string `json:"UserFields,omitempty" xml:"UserFields,omitempty"`
	// 任务预热时间，单位秒
	//
	// example:
	//
	// 300
	WarmUpDuration *string `json:"WarmUpDuration,omitempty" xml:"WarmUpDuration,omitempty"`
	// 座席整理时间，单位秒
	//
	// example:
	//
	// 30
	Wrapup *string `json:"Wrapup,omitempty" xml:"Wrapup,omitempty"`
}

func (s CloudGetTaskResponseBodyDataTaskProperty) String() string {
	return dara.Prettify(s)
}

func (s CloudGetTaskResponseBodyDataTaskProperty) GoString() string {
	return s.String()
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) GetAdjustStep() *string {
	return s.AdjustStep
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) GetAgentTimeout() *string {
	return s.AgentTimeout
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) GetAgentUtilization() *string {
	return s.AgentUtilization
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) GetAnswerRate() *string {
	return s.AnswerRate
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) GetAutoComplete() *string {
	return s.AutoComplete
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) GetAutoStart() *string {
	return s.AutoStart
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) GetAutoStartDay() *string {
	return s.AutoStartDay
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) GetAutoStartTime() *string {
	return s.AutoStartTime
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) GetAutoStop() *string {
	return s.AutoStop
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) GetAutoStopDay() *string {
	return s.AutoStopDay
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) GetAutoStopTime() *string {
	return s.AutoStopTime
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) GetAutoTaskType() *string {
	return s.AutoTaskType
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) GetAutoTriggerTimeStrategy() *string {
	return s.AutoTriggerTimeStrategy
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) GetCallLimitStrategy() *string {
	return s.CallLimitStrategy
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) GetCallStrategy() *string {
	return s.CallStrategy
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) GetCallVariables() *string {
	return s.CallVariables
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) GetCallbackUrl() *string {
	return s.CallbackUrl
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) GetClidGroup() *string {
	return s.ClidGroup
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) GetClidProperty() *string {
	return s.ClidProperty
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) GetCnos() *string {
	return s.Cnos
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) GetConcurrency() *string {
	return s.Concurrency
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) GetCurlRetryTimes() *string {
	return s.CurlRetryTimes
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) GetCustomerClidType() *string {
	return s.CustomerClidType
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) GetCustomerClidWeight() *string {
	return s.CustomerClidWeight
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) GetCustomerClidWeightFlag() *string {
	return s.CustomerClidWeightFlag
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) GetCustomerClids() *string {
	return s.CustomerClids
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) GetCustomerClidsCategory() *string {
	return s.CustomerClidsCategory
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) GetCustomerMoh() *string {
	return s.CustomerMoh
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) GetCustomerTimeout() *string {
	return s.CustomerTimeout
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) GetCustomerVoice() *string {
	return s.CustomerVoice
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) GetDescription() *string {
	return s.Description
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) GetEnterpriseId() *string {
	return s.EnterpriseId
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) GetForceEndFlag() *string {
	return s.ForceEndFlag
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) GetId() *string {
	return s.Id
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) GetIsRandom() *string {
	return s.IsRandom
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) GetIsRewarm() *string {
	return s.IsRewarm
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) GetIvrId() *string {
	return s.IvrId
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) GetMaxWaitTime() *string {
	return s.MaxWaitTime
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) GetMinAvailableAgentCount() *string {
	return s.MinAvailableAgentCount
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) GetName() *string {
	return s.Name
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) GetOverCallLimitCnos() *string {
	return s.OverCallLimitCnos
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) GetOverTime() *string {
	return s.OverTime
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) GetPauseDuration() *string {
	return s.PauseDuration
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) GetPauseTime() *string {
	return s.PauseTime
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) GetPredictAdjust() *string {
	return s.PredictAdjust
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) GetQuotiety() *string {
	return s.Quotiety
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) GetRetryStrategy() *string {
	return s.RetryStrategy
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) GetRetryStrategyOnlyToday() *string {
	return s.RetryStrategyOnlyToday
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) GetStandbyCnos() *string {
	return s.StandbyCnos
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) GetStartTime() *string {
	return s.StartTime
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) GetStatus() *string {
	return s.Status
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) GetStatusDescription() *string {
	return s.StatusDescription
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) GetTaskFields() *string {
	return s.TaskFields
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) GetTimeStrategy() *string {
	return s.TimeStrategy
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) GetType() *string {
	return s.Type
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) GetUserFields() *string {
	return s.UserFields
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) GetWarmUpDuration() *string {
	return s.WarmUpDuration
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) GetWrapup() *string {
	return s.Wrapup
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) SetAdjustStep(v string) *CloudGetTaskResponseBodyDataTaskProperty {
	s.AdjustStep = &v
	return s
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) SetAgentTimeout(v string) *CloudGetTaskResponseBodyDataTaskProperty {
	s.AgentTimeout = &v
	return s
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) SetAgentUtilization(v string) *CloudGetTaskResponseBodyDataTaskProperty {
	s.AgentUtilization = &v
	return s
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) SetAnswerRate(v string) *CloudGetTaskResponseBodyDataTaskProperty {
	s.AnswerRate = &v
	return s
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) SetAutoComplete(v string) *CloudGetTaskResponseBodyDataTaskProperty {
	s.AutoComplete = &v
	return s
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) SetAutoStart(v string) *CloudGetTaskResponseBodyDataTaskProperty {
	s.AutoStart = &v
	return s
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) SetAutoStartDay(v string) *CloudGetTaskResponseBodyDataTaskProperty {
	s.AutoStartDay = &v
	return s
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) SetAutoStartTime(v string) *CloudGetTaskResponseBodyDataTaskProperty {
	s.AutoStartTime = &v
	return s
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) SetAutoStop(v string) *CloudGetTaskResponseBodyDataTaskProperty {
	s.AutoStop = &v
	return s
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) SetAutoStopDay(v string) *CloudGetTaskResponseBodyDataTaskProperty {
	s.AutoStopDay = &v
	return s
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) SetAutoStopTime(v string) *CloudGetTaskResponseBodyDataTaskProperty {
	s.AutoStopTime = &v
	return s
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) SetAutoTaskType(v string) *CloudGetTaskResponseBodyDataTaskProperty {
	s.AutoTaskType = &v
	return s
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) SetAutoTriggerTimeStrategy(v string) *CloudGetTaskResponseBodyDataTaskProperty {
	s.AutoTriggerTimeStrategy = &v
	return s
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) SetCallLimitStrategy(v string) *CloudGetTaskResponseBodyDataTaskProperty {
	s.CallLimitStrategy = &v
	return s
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) SetCallStrategy(v string) *CloudGetTaskResponseBodyDataTaskProperty {
	s.CallStrategy = &v
	return s
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) SetCallVariables(v string) *CloudGetTaskResponseBodyDataTaskProperty {
	s.CallVariables = &v
	return s
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) SetCallbackUrl(v string) *CloudGetTaskResponseBodyDataTaskProperty {
	s.CallbackUrl = &v
	return s
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) SetClidGroup(v string) *CloudGetTaskResponseBodyDataTaskProperty {
	s.ClidGroup = &v
	return s
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) SetClidProperty(v string) *CloudGetTaskResponseBodyDataTaskProperty {
	s.ClidProperty = &v
	return s
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) SetCnos(v string) *CloudGetTaskResponseBodyDataTaskProperty {
	s.Cnos = &v
	return s
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) SetConcurrency(v string) *CloudGetTaskResponseBodyDataTaskProperty {
	s.Concurrency = &v
	return s
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) SetCreateTime(v string) *CloudGetTaskResponseBodyDataTaskProperty {
	s.CreateTime = &v
	return s
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) SetCurlRetryTimes(v string) *CloudGetTaskResponseBodyDataTaskProperty {
	s.CurlRetryTimes = &v
	return s
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) SetCustomerClidType(v string) *CloudGetTaskResponseBodyDataTaskProperty {
	s.CustomerClidType = &v
	return s
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) SetCustomerClidWeight(v string) *CloudGetTaskResponseBodyDataTaskProperty {
	s.CustomerClidWeight = &v
	return s
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) SetCustomerClidWeightFlag(v string) *CloudGetTaskResponseBodyDataTaskProperty {
	s.CustomerClidWeightFlag = &v
	return s
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) SetCustomerClids(v string) *CloudGetTaskResponseBodyDataTaskProperty {
	s.CustomerClids = &v
	return s
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) SetCustomerClidsCategory(v string) *CloudGetTaskResponseBodyDataTaskProperty {
	s.CustomerClidsCategory = &v
	return s
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) SetCustomerMoh(v string) *CloudGetTaskResponseBodyDataTaskProperty {
	s.CustomerMoh = &v
	return s
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) SetCustomerTimeout(v string) *CloudGetTaskResponseBodyDataTaskProperty {
	s.CustomerTimeout = &v
	return s
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) SetCustomerVoice(v string) *CloudGetTaskResponseBodyDataTaskProperty {
	s.CustomerVoice = &v
	return s
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) SetDescription(v string) *CloudGetTaskResponseBodyDataTaskProperty {
	s.Description = &v
	return s
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) SetEnterpriseId(v string) *CloudGetTaskResponseBodyDataTaskProperty {
	s.EnterpriseId = &v
	return s
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) SetForceEndFlag(v string) *CloudGetTaskResponseBodyDataTaskProperty {
	s.ForceEndFlag = &v
	return s
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) SetId(v string) *CloudGetTaskResponseBodyDataTaskProperty {
	s.Id = &v
	return s
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) SetIsRandom(v string) *CloudGetTaskResponseBodyDataTaskProperty {
	s.IsRandom = &v
	return s
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) SetIsRewarm(v string) *CloudGetTaskResponseBodyDataTaskProperty {
	s.IsRewarm = &v
	return s
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) SetIvrId(v string) *CloudGetTaskResponseBodyDataTaskProperty {
	s.IvrId = &v
	return s
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) SetMaxWaitTime(v string) *CloudGetTaskResponseBodyDataTaskProperty {
	s.MaxWaitTime = &v
	return s
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) SetMinAvailableAgentCount(v string) *CloudGetTaskResponseBodyDataTaskProperty {
	s.MinAvailableAgentCount = &v
	return s
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) SetName(v string) *CloudGetTaskResponseBodyDataTaskProperty {
	s.Name = &v
	return s
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) SetOverCallLimitCnos(v string) *CloudGetTaskResponseBodyDataTaskProperty {
	s.OverCallLimitCnos = &v
	return s
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) SetOverTime(v string) *CloudGetTaskResponseBodyDataTaskProperty {
	s.OverTime = &v
	return s
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) SetPauseDuration(v string) *CloudGetTaskResponseBodyDataTaskProperty {
	s.PauseDuration = &v
	return s
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) SetPauseTime(v string) *CloudGetTaskResponseBodyDataTaskProperty {
	s.PauseTime = &v
	return s
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) SetPredictAdjust(v string) *CloudGetTaskResponseBodyDataTaskProperty {
	s.PredictAdjust = &v
	return s
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) SetQuotiety(v string) *CloudGetTaskResponseBodyDataTaskProperty {
	s.Quotiety = &v
	return s
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) SetRetryStrategy(v string) *CloudGetTaskResponseBodyDataTaskProperty {
	s.RetryStrategy = &v
	return s
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) SetRetryStrategyOnlyToday(v string) *CloudGetTaskResponseBodyDataTaskProperty {
	s.RetryStrategyOnlyToday = &v
	return s
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) SetStandbyCnos(v string) *CloudGetTaskResponseBodyDataTaskProperty {
	s.StandbyCnos = &v
	return s
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) SetStartTime(v string) *CloudGetTaskResponseBodyDataTaskProperty {
	s.StartTime = &v
	return s
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) SetStatus(v string) *CloudGetTaskResponseBodyDataTaskProperty {
	s.Status = &v
	return s
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) SetStatusDescription(v string) *CloudGetTaskResponseBodyDataTaskProperty {
	s.StatusDescription = &v
	return s
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) SetTaskFields(v string) *CloudGetTaskResponseBodyDataTaskProperty {
	s.TaskFields = &v
	return s
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) SetTimeStrategy(v string) *CloudGetTaskResponseBodyDataTaskProperty {
	s.TimeStrategy = &v
	return s
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) SetType(v string) *CloudGetTaskResponseBodyDataTaskProperty {
	s.Type = &v
	return s
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) SetUserFields(v string) *CloudGetTaskResponseBodyDataTaskProperty {
	s.UserFields = &v
	return s
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) SetWarmUpDuration(v string) *CloudGetTaskResponseBodyDataTaskProperty {
	s.WarmUpDuration = &v
	return s
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) SetWrapup(v string) *CloudGetTaskResponseBodyDataTaskProperty {
	s.Wrapup = &v
	return s
}

func (s *CloudGetTaskResponseBodyDataTaskProperty) Validate() error {
	return dara.Validate(s)
}
