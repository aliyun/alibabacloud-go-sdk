// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudQueryTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CloudQueryTaskResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CloudQueryTaskResponseBody
	GetCode() *string
	SetData(v *CloudQueryTaskResponseBodyData) *CloudQueryTaskResponseBody
	GetData() *CloudQueryTaskResponseBodyData
	SetMessage(v string) *CloudQueryTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *CloudQueryTaskResponseBody
	GetRequestId() *string
}

type CloudQueryTaskResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CloudQueryTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// D9CB3933-9FE3-4870-BA8E-2BEE91B69D23
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloudQueryTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloudQueryTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CloudQueryTaskResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CloudQueryTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *CloudQueryTaskResponseBody) GetData() *CloudQueryTaskResponseBodyData {
	return s.Data
}

func (s *CloudQueryTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CloudQueryTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloudQueryTaskResponseBody) SetAccessDeniedDetail(v string) *CloudQueryTaskResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CloudQueryTaskResponseBody) SetCode(v string) *CloudQueryTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CloudQueryTaskResponseBody) SetData(v *CloudQueryTaskResponseBodyData) *CloudQueryTaskResponseBody {
	s.Data = v
	return s
}

func (s *CloudQueryTaskResponseBody) SetMessage(v string) *CloudQueryTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CloudQueryTaskResponseBody) SetRequestId(v string) *CloudQueryTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloudQueryTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CloudQueryTaskResponseBodyData struct {
	TaskProperties []*CloudQueryTaskResponseBodyDataTaskProperties `json:"TaskProperties,omitempty" xml:"TaskProperties,omitempty" type:"Repeated"`
}

func (s CloudQueryTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CloudQueryTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *CloudQueryTaskResponseBodyData) GetTaskProperties() []*CloudQueryTaskResponseBodyDataTaskProperties {
	return s.TaskProperties
}

func (s *CloudQueryTaskResponseBodyData) SetTaskProperties(v []*CloudQueryTaskResponseBodyDataTaskProperties) *CloudQueryTaskResponseBodyData {
	s.TaskProperties = v
	return s
}

func (s *CloudQueryTaskResponseBodyData) Validate() error {
	if s.TaskProperties != nil {
		for _, item := range s.TaskProperties {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CloudQueryTaskResponseBodyDataTaskProperties struct {
	// 外呼组号
	//
	// example:
	//
	// WH124
	AgentGroup *string `json:"AgentGroup,omitempty" xml:"AgentGroup,omitempty"`
	// 座席超时时间，单位秒
	//
	// example:
	//
	// 10
	AgentTimeout *int64 `json:"AgentTimeout,omitempty" xml:"AgentTimeout,omitempty"`
	// 初始化预计客户接通率
	//
	// example:
	//
	// 50
	AnswerRate *int64 `json:"AnswerRate,omitempty" xml:"AnswerRate,omitempty"`
	// 自动完成，0.关闭 1.开启
	//
	// example:
	//
	// 0
	AutoComplete *int64 `json:"AutoComplete,omitempty" xml:"AutoComplete,omitempty"`
	// 定时开始，0.关闭 1.开启
	//
	// example:
	//
	// 0
	AutoStart *int64 `json:"AutoStart,omitempty" xml:"AutoStart,omitempty"`
	// 定时开始日期，格式：yyyy-MM-dd，如：2017-02-11
	//
	// example:
	//
	// 2017-02-11
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
	AutoStop *int64 `json:"AutoStop,omitempty" xml:"AutoStop,omitempty"`
	// 定时完成日期，格式：yyyy-MM-dd，如：2017-02-11
	//
	// example:
	//
	// 2017-02-11
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
	// 50
	AutoTaskType *int64 `json:"AutoTaskType,omitempty" xml:"AutoTaskType,omitempty"`
	// 间隔呼叫时间段，呼叫的时间段配置。参数格式：时间条件编号。支持多个，多个使用英文","逗号隔开，如：9,22
	//
	// example:
	//
	// 9,22
	AutoTriggerTimeStrategy *string `json:"AutoTriggerTimeStrategy,omitempty" xml:"AutoTriggerTimeStrategy,omitempty"`
	// 指定座席方式，1.座席工号列表 2.外呼组
	//
	// example:
	//
	// 1
	CallGroupType *int64 `json:"CallGroupType,omitempty" xml:"CallGroupType,omitempty"`
	// 座席当日接听数限制，JsonArray格式[{"cnos":["2000","2001"],"limit":"5"}]
	//
	// example:
	//
	// [{"cnos":["2000","2001"],"limit":"5"}]
	CallLimitStrategy *string `json:"CallLimitStrategy,omitempty" xml:"CallLimitStrategy,omitempty"`
	// 呼叫顺序，数据结构为Json格式```{"strategy":[{"sort":1,"type":"retryCall", "desc":0},{"sort":2,"type":"firstCall","orderType":0}]}```，格式说明：sort是重试号码和未呼叫号码的呼叫顺序 ，type：retryCall重试号码、firstCall未呼叫号码，当type=retryCall，desc：0.优先呼叫待重呼轮次数值较小的号码，即轮次靠前（如第1轮、第2轮）的号码先被呼叫 1.优先呼叫待重呼轮次数值较大的号码，即轮次靠后（如第5轮、第4轮）的号码先被呼叫，当type=firstCall时，orderType：随机呼叫，0顺序(优先级) 1随机 2顺序(导入时间)
	//
	// example:
	//
	// {"strategy":[{"sort":1,"type":"retryCall", "desc":0},{"sort":2,"type":"firstCall","orderType":0}]}
	CallPriorityStrategy *string `json:"CallPriorityStrategy,omitempty" xml:"CallPriorityStrategy,omitempty"`
	// 呼叫流转模式， 1.直连座席 2.AI转人工 直连座席模式：客户接听后直接分配座席，无空闲座席可溢出语音导航（需配置语音导航），保障高优先级客户直连体验 AI转人工：客户接入后，优先进入语音导航中配置的智能体机器人，按交互结果决定是否转接座席，提升自动化覆盖率，降低人工成本
	//
	// example:
	//
	// 1
	CallRouteStrategy *int64 `json:"CallRouteStrategy,omitempty" xml:"CallRouteStrategy,omitempty"`
	// 座席分配规则，1.随机 2.顺序 3.座席未进线最大时（从上一次呼叫结束到当前的总时长） 4.座席状态[空闲]最长时长（座席最近一次切换为空闲状态的持续时长）
	//
	// example:
	//
	// 1
	CallStrategy *int64 `json:"CallStrategy,omitempty" xml:"CallStrategy,omitempty"`
	// 座席通道变量
	//
	// example:
	//
	// [{"key1":"value1"},{"key2":"value2"}]
	CallVariables *string `json:"CallVariables,omitempty" xml:"CallVariables,omitempty"`
	// 外显规则，Json格式，外显号码选择规则：默认区号，是否匹配省会等；如：{"defaultAreaCode":"010", "isMatchCapital":"1"}
	//
	// example:
	//
	// 示例值
	ClidProperty *string `json:"ClidProperty,omitempty" xml:"ClidProperty,omitempty"`
	// 座席工号列表
	//
	// example:
	//
	// 1111,2222
	Cnos *string `json:"Cnos,omitempty" xml:"Cnos,omitempty"`
	// 最大并发限制
	//
	// example:
	//
	// 3
	Concurrency *int64 `json:"Concurrency,omitempty" xml:"Concurrency,omitempty"`
	// 任务创建时间 ，格式为： yyyy-MM-dd HH:mm:ss
	//
	// example:
	//
	// 2025-10-12 18:03:14
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 客户侧外显规则，1.随机 2.按区号
	//
	// example:
	//
	// 1
	CustomerClidType *int64 `json:"CustomerClidType,omitempty" xml:"CustomerClidType,omitempty"`
	// 外显号码比例配置内容
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
	CustomerClidWeightFlag *int64 `json:"CustomerClidWeightFlag,omitempty" xml:"CustomerClidWeightFlag,omitempty"`
	// 客户侧外显号码列表
	//
	// example:
	//
	// 02138276106
	CustomerClids *string `json:"CustomerClids,omitempty" xml:"CustomerClids,omitempty"`
	// 客户侧外显号码类型，1.外显固话 2.外显手机号  4.外显号码池 5.外显导航
	//
	// example:
	//
	// 1
	CustomerClidsCategory *int64 `json:"CustomerClidsCategory,omitempty" xml:"CustomerClidsCategory,omitempty"`
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
	// 17
	CustomerTimeout *int64 `json:"CustomerTimeout,omitempty" xml:"CustomerTimeout,omitempty"`
	// 客户侧提示音
	//
	// example:
	//
	// 示例值示例值
	CustomerVoice *string `json:"CustomerVoice,omitempty" xml:"CustomerVoice,omitempty"`
	// 任务描述
	//
	// example:
	//
	// desc
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 呼叫中心Id
	//
	// example:
	//
	// 7000002
	EnterpriseId *string `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	// 定时完成时强制结束任务，开启配置后，当任务到定时完成时间时无论任务中的号码是否已经呼完，均将任务状态设置为结束。适用于对数据有呼叫时间限制的场景。注：任务在结束状态时无法再次启动。 0：否，1：是
	//
	// example:
	//
	// 0
	ForceEndFlag *int64 `json:"ForceEndFlag,omitempty" xml:"ForceEndFlag,omitempty"`
	// 外呼任务Id
	//
	// example:
	//
	// 59
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// 暂停后重新预热，0.关闭，1开启
	//
	// example:
	//
	// 0
	IsRewarm *int64 `json:"IsRewarm,omitempty" xml:"IsRewarm,omitempty"`
	// 语音导航Id，预测外呼使用场景：直连座席模式下如果客户接听后未在特定时间内（默认2秒）找到可用座席，通话将溢出到语音导航继续排队找座席。同时支持在呼转座席，座席未接听时也溢出到语音导航，需要开启企业配置生效。AI转人工模式下，呼叫先转到语音导航 自动外呼使用场景：客户接听后，转到的语音导航
	//
	// example:
	//
	// 133
	IvrId *int64 `json:"IvrId,omitempty" xml:"IvrId,omitempty"`
	// 座席最大等待时间，单位秒
	//
	// example:
	//
	// 34
	MaxWaitTime *int64 `json:"MaxWaitTime,omitempty" xml:"MaxWaitTime,omitempty"`
	// 最小可用座席数
	//
	// example:
	//
	// 2
	MinAvailableAgentCount *int64 `json:"MinAvailableAgentCount,omitempty" xml:"MinAvailableAgentCount,omitempty"`
	// 任务名称
	//
	// example:
	//
	// name1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 超过呼叫限制的座席
	//
	// example:
	//
	// 0
	OverCallLimitCnos *string `json:"OverCallLimitCnos,omitempty" xml:"OverCallLimitCnos,omitempty"`
	// 任务结束时间
	//
	// example:
	//
	// 示例值示例值
	OverTime *string `json:"OverTime,omitempty" xml:"OverTime,omitempty"`
	// 任务暂停时长(针对自动启动时间段) 单位分钟
	//
	// example:
	//
	// 1
	PauseDuration *int64 `json:"PauseDuration,omitempty" xml:"PauseDuration,omitempty"`
	// 任务暂停时间点
	//
	// example:
	//
	// 2026-04-20 11:00:00
	PauseTime *string `json:"PauseTime,omitempty" xml:"PauseTime,omitempty"`
	// 超呼率
	//
	// example:
	//
	// 60
	PredictAdjust *int64 `json:"PredictAdjust,omitempty" xml:"PredictAdjust,omitempty"`
	// 骚扰率，支持小数，精确到小数点两位
	//
	// example:
	//
	// 1.0
	Quotiety *float64 `json:"Quotiety,omitempty" xml:"Quotiety,omitempty"`
	// 重试策略， 「基础模式」 数据结构为Json格式```{"retry":2,"strategy":[{"round":1,"time":"1-1-1"},{"round":2,"time":"2-2-2"}]}``` 格式说明：retry是重试次数，round是第几次重试，time是第几次重试间隔的时间：1-1-1，第一个1是天，第二个1是小时，第三个1是分钟 「高级模式」 高级模式需要开启「号码状态识别」服务后。目前只支持「自动外呼」任务模式。 数据结构为JsonArray格式 ```[{"condition":{"sipCause":[710,719]},"retry":1,"sort":1,"strategy":[{"round":1,"time":"0-0-10"}]},{"condition":{"sipCause":[800]},"retry":1,"sort":2,"strategy":[{"round":2,"time":"0-0-10"}]}]``` 格式说明：condition是重试条件，sort是重试条件的顺序，其余字段与基础模式一致
	//
	// example:
	//
	// {"retry":2,"strategy":[{"round":1,"time":"1-1-1"},{"round":2,"time":"2-2-2"}]}
	RetryStrategy *string `json:"RetryStrategy,omitempty" xml:"RetryStrategy,omitempty"`
	// 重试仅当天生效，0.关闭 1.开启，删除待重试的数据和待呼叫的数据 2.开启，删除待重试的数据 3.开启，删除待呼叫的数据
	//
	// example:
	//
	// 4
	RetryStrategyOnlyToday *int64 `json:"RetryStrategyOnlyToday,omitempty" xml:"RetryStrategyOnlyToday,omitempty"`
	// 重试策略时间类型，1.基于首次呼叫时间 2.基于上次呼叫时间
	//
	// example:
	//
	// 1
	RetryStrategyTimeType *int64 `json:"RetryStrategyTimeType,omitempty" xml:"RetryStrategyTimeType,omitempty"`
	// 任务开始时间
	//
	// example:
	//
	// 09:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 任务状态，0.初始 1.运行中 2.暂停 3.结束
	//
	// example:
	//
	// 2
	Status *int64 `json:"Status,omitempty" xml:"Status,omitempty"`
	// 任务状态描述
	//
	// example:
	//
	// ""
	StatusDescription *string `json:"StatusDescription,omitempty" xml:"StatusDescription,omitempty"`
	// 任务状态变更类型，0.系统变更 1.人为变更
	//
	// example:
	//
	// 0
	StatusTriggerType *int64 `json:"StatusTriggerType,omitempty" xml:"StatusTriggerType,omitempty"`
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
	Type *int64 `json:"Type,omitempty" xml:"Type,omitempty"`
	// 任务自定义字段
	//
	// example:
	//
	// [{"key1":"value1"},{"key2":"value2"}]
	UserFields *string `json:"UserFields,omitempty" xml:"UserFields,omitempty"`
	// 任务预热时间，单位秒
	//
	// example:
	//
	// 9
	WarmUpDuration *int64 `json:"WarmUpDuration,omitempty" xml:"WarmUpDuration,omitempty"`
	// 座席整理时间，单位秒
	//
	// example:
	//
	// 13
	Wrapup *int64 `json:"Wrapup,omitempty" xml:"Wrapup,omitempty"`
}

func (s CloudQueryTaskResponseBodyDataTaskProperties) String() string {
	return dara.Prettify(s)
}

func (s CloudQueryTaskResponseBodyDataTaskProperties) GoString() string {
	return s.String()
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) GetAgentGroup() *string {
	return s.AgentGroup
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) GetAgentTimeout() *int64 {
	return s.AgentTimeout
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) GetAnswerRate() *int64 {
	return s.AnswerRate
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) GetAutoComplete() *int64 {
	return s.AutoComplete
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) GetAutoStart() *int64 {
	return s.AutoStart
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) GetAutoStartDay() *string {
	return s.AutoStartDay
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) GetAutoStartTime() *string {
	return s.AutoStartTime
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) GetAutoStop() *int64 {
	return s.AutoStop
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) GetAutoStopDay() *string {
	return s.AutoStopDay
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) GetAutoStopTime() *string {
	return s.AutoStopTime
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) GetAutoTaskType() *int64 {
	return s.AutoTaskType
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) GetAutoTriggerTimeStrategy() *string {
	return s.AutoTriggerTimeStrategy
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) GetCallGroupType() *int64 {
	return s.CallGroupType
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) GetCallLimitStrategy() *string {
	return s.CallLimitStrategy
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) GetCallPriorityStrategy() *string {
	return s.CallPriorityStrategy
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) GetCallRouteStrategy() *int64 {
	return s.CallRouteStrategy
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) GetCallStrategy() *int64 {
	return s.CallStrategy
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) GetCallVariables() *string {
	return s.CallVariables
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) GetClidProperty() *string {
	return s.ClidProperty
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) GetCnos() *string {
	return s.Cnos
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) GetConcurrency() *int64 {
	return s.Concurrency
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) GetCustomerClidType() *int64 {
	return s.CustomerClidType
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) GetCustomerClidWeight() *string {
	return s.CustomerClidWeight
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) GetCustomerClidWeightFlag() *int64 {
	return s.CustomerClidWeightFlag
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) GetCustomerClids() *string {
	return s.CustomerClids
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) GetCustomerClidsCategory() *int64 {
	return s.CustomerClidsCategory
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) GetCustomerMoh() *string {
	return s.CustomerMoh
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) GetCustomerTimeout() *int64 {
	return s.CustomerTimeout
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) GetCustomerVoice() *string {
	return s.CustomerVoice
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) GetDescription() *string {
	return s.Description
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) GetEnterpriseId() *string {
	return s.EnterpriseId
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) GetForceEndFlag() *int64 {
	return s.ForceEndFlag
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) GetId() *int64 {
	return s.Id
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) GetIsRewarm() *int64 {
	return s.IsRewarm
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) GetIvrId() *int64 {
	return s.IvrId
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) GetMaxWaitTime() *int64 {
	return s.MaxWaitTime
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) GetMinAvailableAgentCount() *int64 {
	return s.MinAvailableAgentCount
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) GetName() *string {
	return s.Name
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) GetOverCallLimitCnos() *string {
	return s.OverCallLimitCnos
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) GetOverTime() *string {
	return s.OverTime
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) GetPauseDuration() *int64 {
	return s.PauseDuration
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) GetPauseTime() *string {
	return s.PauseTime
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) GetPredictAdjust() *int64 {
	return s.PredictAdjust
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) GetQuotiety() *float64 {
	return s.Quotiety
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) GetRetryStrategy() *string {
	return s.RetryStrategy
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) GetRetryStrategyOnlyToday() *int64 {
	return s.RetryStrategyOnlyToday
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) GetRetryStrategyTimeType() *int64 {
	return s.RetryStrategyTimeType
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) GetStartTime() *string {
	return s.StartTime
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) GetStatus() *int64 {
	return s.Status
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) GetStatusDescription() *string {
	return s.StatusDescription
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) GetStatusTriggerType() *int64 {
	return s.StatusTriggerType
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) GetTimeStrategy() *string {
	return s.TimeStrategy
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) GetType() *int64 {
	return s.Type
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) GetUserFields() *string {
	return s.UserFields
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) GetWarmUpDuration() *int64 {
	return s.WarmUpDuration
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) GetWrapup() *int64 {
	return s.Wrapup
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) SetAgentGroup(v string) *CloudQueryTaskResponseBodyDataTaskProperties {
	s.AgentGroup = &v
	return s
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) SetAgentTimeout(v int64) *CloudQueryTaskResponseBodyDataTaskProperties {
	s.AgentTimeout = &v
	return s
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) SetAnswerRate(v int64) *CloudQueryTaskResponseBodyDataTaskProperties {
	s.AnswerRate = &v
	return s
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) SetAutoComplete(v int64) *CloudQueryTaskResponseBodyDataTaskProperties {
	s.AutoComplete = &v
	return s
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) SetAutoStart(v int64) *CloudQueryTaskResponseBodyDataTaskProperties {
	s.AutoStart = &v
	return s
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) SetAutoStartDay(v string) *CloudQueryTaskResponseBodyDataTaskProperties {
	s.AutoStartDay = &v
	return s
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) SetAutoStartTime(v string) *CloudQueryTaskResponseBodyDataTaskProperties {
	s.AutoStartTime = &v
	return s
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) SetAutoStop(v int64) *CloudQueryTaskResponseBodyDataTaskProperties {
	s.AutoStop = &v
	return s
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) SetAutoStopDay(v string) *CloudQueryTaskResponseBodyDataTaskProperties {
	s.AutoStopDay = &v
	return s
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) SetAutoStopTime(v string) *CloudQueryTaskResponseBodyDataTaskProperties {
	s.AutoStopTime = &v
	return s
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) SetAutoTaskType(v int64) *CloudQueryTaskResponseBodyDataTaskProperties {
	s.AutoTaskType = &v
	return s
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) SetAutoTriggerTimeStrategy(v string) *CloudQueryTaskResponseBodyDataTaskProperties {
	s.AutoTriggerTimeStrategy = &v
	return s
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) SetCallGroupType(v int64) *CloudQueryTaskResponseBodyDataTaskProperties {
	s.CallGroupType = &v
	return s
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) SetCallLimitStrategy(v string) *CloudQueryTaskResponseBodyDataTaskProperties {
	s.CallLimitStrategy = &v
	return s
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) SetCallPriorityStrategy(v string) *CloudQueryTaskResponseBodyDataTaskProperties {
	s.CallPriorityStrategy = &v
	return s
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) SetCallRouteStrategy(v int64) *CloudQueryTaskResponseBodyDataTaskProperties {
	s.CallRouteStrategy = &v
	return s
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) SetCallStrategy(v int64) *CloudQueryTaskResponseBodyDataTaskProperties {
	s.CallStrategy = &v
	return s
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) SetCallVariables(v string) *CloudQueryTaskResponseBodyDataTaskProperties {
	s.CallVariables = &v
	return s
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) SetClidProperty(v string) *CloudQueryTaskResponseBodyDataTaskProperties {
	s.ClidProperty = &v
	return s
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) SetCnos(v string) *CloudQueryTaskResponseBodyDataTaskProperties {
	s.Cnos = &v
	return s
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) SetConcurrency(v int64) *CloudQueryTaskResponseBodyDataTaskProperties {
	s.Concurrency = &v
	return s
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) SetCreateTime(v string) *CloudQueryTaskResponseBodyDataTaskProperties {
	s.CreateTime = &v
	return s
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) SetCustomerClidType(v int64) *CloudQueryTaskResponseBodyDataTaskProperties {
	s.CustomerClidType = &v
	return s
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) SetCustomerClidWeight(v string) *CloudQueryTaskResponseBodyDataTaskProperties {
	s.CustomerClidWeight = &v
	return s
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) SetCustomerClidWeightFlag(v int64) *CloudQueryTaskResponseBodyDataTaskProperties {
	s.CustomerClidWeightFlag = &v
	return s
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) SetCustomerClids(v string) *CloudQueryTaskResponseBodyDataTaskProperties {
	s.CustomerClids = &v
	return s
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) SetCustomerClidsCategory(v int64) *CloudQueryTaskResponseBodyDataTaskProperties {
	s.CustomerClidsCategory = &v
	return s
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) SetCustomerMoh(v string) *CloudQueryTaskResponseBodyDataTaskProperties {
	s.CustomerMoh = &v
	return s
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) SetCustomerTimeout(v int64) *CloudQueryTaskResponseBodyDataTaskProperties {
	s.CustomerTimeout = &v
	return s
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) SetCustomerVoice(v string) *CloudQueryTaskResponseBodyDataTaskProperties {
	s.CustomerVoice = &v
	return s
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) SetDescription(v string) *CloudQueryTaskResponseBodyDataTaskProperties {
	s.Description = &v
	return s
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) SetEnterpriseId(v string) *CloudQueryTaskResponseBodyDataTaskProperties {
	s.EnterpriseId = &v
	return s
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) SetForceEndFlag(v int64) *CloudQueryTaskResponseBodyDataTaskProperties {
	s.ForceEndFlag = &v
	return s
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) SetId(v int64) *CloudQueryTaskResponseBodyDataTaskProperties {
	s.Id = &v
	return s
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) SetIsRewarm(v int64) *CloudQueryTaskResponseBodyDataTaskProperties {
	s.IsRewarm = &v
	return s
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) SetIvrId(v int64) *CloudQueryTaskResponseBodyDataTaskProperties {
	s.IvrId = &v
	return s
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) SetMaxWaitTime(v int64) *CloudQueryTaskResponseBodyDataTaskProperties {
	s.MaxWaitTime = &v
	return s
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) SetMinAvailableAgentCount(v int64) *CloudQueryTaskResponseBodyDataTaskProperties {
	s.MinAvailableAgentCount = &v
	return s
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) SetName(v string) *CloudQueryTaskResponseBodyDataTaskProperties {
	s.Name = &v
	return s
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) SetOverCallLimitCnos(v string) *CloudQueryTaskResponseBodyDataTaskProperties {
	s.OverCallLimitCnos = &v
	return s
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) SetOverTime(v string) *CloudQueryTaskResponseBodyDataTaskProperties {
	s.OverTime = &v
	return s
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) SetPauseDuration(v int64) *CloudQueryTaskResponseBodyDataTaskProperties {
	s.PauseDuration = &v
	return s
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) SetPauseTime(v string) *CloudQueryTaskResponseBodyDataTaskProperties {
	s.PauseTime = &v
	return s
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) SetPredictAdjust(v int64) *CloudQueryTaskResponseBodyDataTaskProperties {
	s.PredictAdjust = &v
	return s
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) SetQuotiety(v float64) *CloudQueryTaskResponseBodyDataTaskProperties {
	s.Quotiety = &v
	return s
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) SetRetryStrategy(v string) *CloudQueryTaskResponseBodyDataTaskProperties {
	s.RetryStrategy = &v
	return s
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) SetRetryStrategyOnlyToday(v int64) *CloudQueryTaskResponseBodyDataTaskProperties {
	s.RetryStrategyOnlyToday = &v
	return s
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) SetRetryStrategyTimeType(v int64) *CloudQueryTaskResponseBodyDataTaskProperties {
	s.RetryStrategyTimeType = &v
	return s
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) SetStartTime(v string) *CloudQueryTaskResponseBodyDataTaskProperties {
	s.StartTime = &v
	return s
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) SetStatus(v int64) *CloudQueryTaskResponseBodyDataTaskProperties {
	s.Status = &v
	return s
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) SetStatusDescription(v string) *CloudQueryTaskResponseBodyDataTaskProperties {
	s.StatusDescription = &v
	return s
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) SetStatusTriggerType(v int64) *CloudQueryTaskResponseBodyDataTaskProperties {
	s.StatusTriggerType = &v
	return s
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) SetTimeStrategy(v string) *CloudQueryTaskResponseBodyDataTaskProperties {
	s.TimeStrategy = &v
	return s
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) SetType(v int64) *CloudQueryTaskResponseBodyDataTaskProperties {
	s.Type = &v
	return s
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) SetUserFields(v string) *CloudQueryTaskResponseBodyDataTaskProperties {
	s.UserFields = &v
	return s
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) SetWarmUpDuration(v int64) *CloudQueryTaskResponseBodyDataTaskProperties {
	s.WarmUpDuration = &v
	return s
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) SetWrapup(v int64) *CloudQueryTaskResponseBodyDataTaskProperties {
	s.Wrapup = &v
	return s
}

func (s *CloudQueryTaskResponseBodyDataTaskProperties) Validate() error {
	return dara.Validate(s)
}
