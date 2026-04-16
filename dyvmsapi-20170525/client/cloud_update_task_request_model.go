// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudUpdateTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentGroup(v string) *CloudUpdateTaskRequest
	GetAgentGroup() *string
	SetAgentTimeout(v int64) *CloudUpdateTaskRequest
	GetAgentTimeout() *int64
	SetAnswerRate(v int64) *CloudUpdateTaskRequest
	GetAnswerRate() *int64
	SetAutoComplete(v int64) *CloudUpdateTaskRequest
	GetAutoComplete() *int64
	SetAutoStart(v int64) *CloudUpdateTaskRequest
	GetAutoStart() *int64
	SetAutoStartDay(v string) *CloudUpdateTaskRequest
	GetAutoStartDay() *string
	SetAutoStartTime(v string) *CloudUpdateTaskRequest
	GetAutoStartTime() *string
	SetAutoStop(v int64) *CloudUpdateTaskRequest
	GetAutoStop() *int64
	SetAutoStopDay(v string) *CloudUpdateTaskRequest
	GetAutoStopDay() *string
	SetAutoStopTime(v string) *CloudUpdateTaskRequest
	GetAutoStopTime() *string
	SetAutoTaskType(v int64) *CloudUpdateTaskRequest
	GetAutoTaskType() *int64
	SetAutoTriggerTimeStrategy(v string) *CloudUpdateTaskRequest
	GetAutoTriggerTimeStrategy() *string
	SetCallLimitStrategy(v string) *CloudUpdateTaskRequest
	GetCallLimitStrategy() *string
	SetCallPriorityStrategy(v string) *CloudUpdateTaskRequest
	GetCallPriorityStrategy() *string
	SetCallRouteStrategy(v int64) *CloudUpdateTaskRequest
	GetCallRouteStrategy() *int64
	SetCallStrategy(v string) *CloudUpdateTaskRequest
	GetCallStrategy() *string
	SetCallVariables(v string) *CloudUpdateTaskRequest
	GetCallVariables() *string
	SetClidProperty(v string) *CloudUpdateTaskRequest
	GetClidProperty() *string
	SetCnos(v string) *CloudUpdateTaskRequest
	GetCnos() *string
	SetConcurrency(v int64) *CloudUpdateTaskRequest
	GetConcurrency() *int64
	SetCustomerClidType(v int64) *CloudUpdateTaskRequest
	GetCustomerClidType() *int64
	SetCustomerClidWeight(v string) *CloudUpdateTaskRequest
	GetCustomerClidWeight() *string
	SetCustomerClidWeightFlag(v int64) *CloudUpdateTaskRequest
	GetCustomerClidWeightFlag() *int64
	SetCustomerClids(v string) *CloudUpdateTaskRequest
	GetCustomerClids() *string
	SetCustomerClidsCategory(v int64) *CloudUpdateTaskRequest
	GetCustomerClidsCategory() *int64
	SetCustomerClidsGroup(v string) *CloudUpdateTaskRequest
	GetCustomerClidsGroup() *string
	SetCustomerMoh(v string) *CloudUpdateTaskRequest
	GetCustomerMoh() *string
	SetCustomerTimeout(v int64) *CloudUpdateTaskRequest
	GetCustomerTimeout() *int64
	SetCustomerVoice(v string) *CloudUpdateTaskRequest
	GetCustomerVoice() *string
	SetDescription(v string) *CloudUpdateTaskRequest
	GetDescription() *string
	SetEnterpriseId(v int64) *CloudUpdateTaskRequest
	GetEnterpriseId() *int64
	SetForceEndFlag(v int64) *CloudUpdateTaskRequest
	GetForceEndFlag() *int64
	SetIsRewarm(v int64) *CloudUpdateTaskRequest
	GetIsRewarm() *int64
	SetIvrId(v int64) *CloudUpdateTaskRequest
	GetIvrId() *int64
	SetIvrName(v string) *CloudUpdateTaskRequest
	GetIvrName() *string
	SetMaxWaitTime(v int64) *CloudUpdateTaskRequest
	GetMaxWaitTime() *int64
	SetMinAvailableAgentCount(v int64) *CloudUpdateTaskRequest
	GetMinAvailableAgentCount() *int64
	SetName(v string) *CloudUpdateTaskRequest
	GetName() *string
	SetOwnerId(v int64) *CloudUpdateTaskRequest
	GetOwnerId() *int64
	SetPredictAdjust(v int64) *CloudUpdateTaskRequest
	GetPredictAdjust() *int64
	SetQuotiety(v float64) *CloudUpdateTaskRequest
	GetQuotiety() *float64
	SetResourceOwnerAccount(v string) *CloudUpdateTaskRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CloudUpdateTaskRequest
	GetResourceOwnerId() *int64
	SetRetryStrategy(v string) *CloudUpdateTaskRequest
	GetRetryStrategy() *string
	SetRetryStrategyOnlyToday(v int64) *CloudUpdateTaskRequest
	GetRetryStrategyOnlyToday() *int64
	SetRetryStrategyTimeType(v int64) *CloudUpdateTaskRequest
	GetRetryStrategyTimeType() *int64
	SetTaskId(v string) *CloudUpdateTaskRequest
	GetTaskId() *string
	SetTimeStrategy(v string) *CloudUpdateTaskRequest
	GetTimeStrategy() *string
	SetUserFields(v string) *CloudUpdateTaskRequest
	GetUserFields() *string
	SetWarmUpDuration(v int64) *CloudUpdateTaskRequest
	GetWarmUpDuration() *int64
	SetWrapup(v int64) *CloudUpdateTaskRequest
	GetWrapup() *int64
}

type CloudUpdateTaskRequest struct {
	// 外呼组号；callGroupType=2时必填。注：一个外呼组可以绑定到多个任务，但只能同时运行一个任务中
	//
	// example:
	//
	// 6,7
	AgentGroup *string `json:"AgentGroup,omitempty" xml:"AgentGroup,omitempty"`
	// 座席超时时间；单位秒数，默认10秒，取值范围5 ~ 60
	//
	// example:
	//
	// 50
	AgentTimeout *int64 `json:"AgentTimeout,omitempty" xml:"AgentTimeout,omitempty"`
	// 初始化预计客户接通率；默认50，取值范围1～100，预热阶段的呼叫，按照此接通率计算呼叫频率
	//
	// example:
	//
	// 50
	AnswerRate *int64 `json:"AnswerRate,omitempty" xml:"AnswerRate,omitempty"`
	// 0.关闭 1.开启，默认开启。开启：任务中的号码呼叫完后，任务自动完成，状态设置为结束，关闭：任务中的号码呼叫完后，任务自动暂停，状态设置为暂停注：任务在结束状态时无法再次启动。
	//
	// example:
	//
	// 1
	AutoComplete *int64 `json:"AutoComplete,omitempty" xml:"AutoComplete,omitempty"`
	// 定时开始；0.关闭 1.开启，默认关闭。值为1时autoStartDay和autoStartTime参数才生效，并且至少设置其中的一个
	//
	// example:
	//
	// 0
	AutoStart *int64 `json:"AutoStart,omitempty" xml:"AutoStart,omitempty"`
	// 定时开始日期；autoStart=1时生效。参数不传入时默认在当天的autoStartTime定时启动。格式：yyyy-MM-dd，如：2017-01-01
	//
	// example:
	//
	// 2026-01-01
	AutoStartDay *string `json:"AutoStartDay,omitempty" xml:"AutoStartDay,omitempty"`
	// 定时开始时间；autoStart=1时生效。参数不传入时默认在autoStartDay的00:00:00定时启动。格式：HH:mm:ss，如：08:00:00
	//
	// example:
	//
	// 08:00:00
	AutoStartTime *string `json:"AutoStartTime,omitempty" xml:"AutoStartTime,omitempty"`
	// 定时完成；0.关闭 1.开启，默认关闭。值为1时autoStopDay和autoStopTime参数才生效，并且至少设置其中的一个
	//
	// example:
	//
	// 0
	AutoStop *int64 `json:"AutoStop,omitempty" xml:"AutoStop,omitempty"`
	// 定时完成日期；autoStop=1时生效。参数不传入时默认在当天的autoStopTime定时完成。格式：yyyy-MM-dd，如：2017-01-01
	//
	// example:
	//
	// 示例值
	AutoStopDay *string `json:"AutoStopDay,omitempty" xml:"AutoStopDay,omitempty"`
	// 定时完成时间；autoStop=1时生效。参数不传入时默认在autoStopDay的23:59:59定时完成。格式：HH:mm:ss，如：17:00:00
	//
	// example:
	//
	// 17:00:00
	AutoStopTime *string `json:"AutoStopTime,omitempty" xml:"AutoStopTime,omitempty"`
	// 0.连续呼叫 1.间隔呼叫，默认连续呼叫。配合定时开始（autoStart）和定时结束（autoStop）参数使用，如任务需要在每天的特定时间段内呼叫则开启间隔呼叫方式
	//
	// example:
	//
	// 0
	AutoTaskType *int64 `json:"AutoTaskType,omitempty" xml:"AutoTaskType,omitempty"`
	// 间隔呼叫时间段；autoTaskType=1时必填。呼叫的时间段配置。参数格式：时间条件编号。支持多个，多个使用英文","逗号隔开，如：9,22。注：时间条件列表可通过查询时间条件设置列表接口获取
	//
	// example:
	//
	// 9,22
	AutoTriggerTimeStrategy *string `json:"AutoTriggerTimeStrategy,omitempty" xml:"AutoTriggerTimeStrategy,omitempty"`
	// 座席当日接听的通话个数上限，使用方式详见下方说明；JsonArray格式 [{"cnos":["2000","2001"],"limit":"5"}] 注：使用座席当日接听数限制功能需开启企业配置，功能开启并且配置上限后系统才开始统计当日接听数
	//
	// example:
	//
	// [{"cnos":["2000","2001"],"limit":"5"}]
	CallLimitStrategy *string `json:"CallLimitStrategy,omitempty" xml:"CallLimitStrategy,omitempty"`
	// 数据结构为Json格式```{"strategy":[{"sort":1,"type":"retryCall", "desc":0},{"sort":2,"type":"firstCall","orderType":0}]}```格式说明：sort是重试号码和未呼叫号码的呼叫顺序 type：retryCall重试号码、firstCall未呼叫号码当type=retryCall，desc：0.优先呼叫待重呼轮次数值较小的号码，即轮次靠前（如第1轮、第2轮）的号码先被呼叫 1.优先呼叫待重呼轮次数值较大的号码，即轮次靠后（如第5轮、第4轮）的号码先被呼叫当type=firstCall时，orderType：随机呼叫，0顺序(优先级) 1随机 2顺序(导入时间)
	//
	// example:
	//
	// {"strategy":[{"sort":1,"type":"retryCall", "desc":0},{"sort":2,"type":"firstCall","orderType":0}]}
	CallPriorityStrategy *string `json:"CallPriorityStrategy,omitempty" xml:"CallPriorityStrategy,omitempty"`
	// 1.直连座席 2.AI转人工 直连座席模式：客户接听后直接分配座席，无空闲座席可溢出语音导航（需配置语音导航），保障高优先级客户直连体验 AI转人工：客户接入后，优先进入语音导航中配置的智能体机器人，按交互结果决定是否转接座席，提升自动化覆盖率，降低人工成本
	//
	// example:
	//
	// 1
	CallRouteStrategy *int64 `json:"CallRouteStrategy,omitempty" xml:"CallRouteStrategy,omitempty"`
	// 座席分配规则；1.随机 2.顺序 3.座席未进线最大时（从上一次呼叫结束到当前的总时长） 4.座席状态[空闲]最长时长（座席最近一次切换为空闲状态的持续时长），默认随机
	//
	// example:
	//
	// 1
	CallStrategy *string `json:"CallStrategy,omitempty" xml:"CallStrategy,omitempty"`
	// 座席通道变量。JsonArray格式；最大支持5个变量。默认空值， 示例：[{"name":"abcdefg","value":"${cdr_enterprise_id}"}] 注：单个字段值的长度不能超过1000个字符。变量会以SIP_HEADER的形式设置到座席侧通道
	//
	// example:
	//
	// [{"name":"abcdefg","value":"${cdr_enterprise_id}"}]
	CallVariables *string `json:"CallVariables,omitempty" xml:"CallVariables,omitempty"`
	// 外显规则；customerClidType=2时生效。Json格式，外显号码选择规则：默认区号，是否匹配省会等；如：{"defaultAreaCode":"010", "isMatchCapital":"1"}。不建议使用
	//
	// example:
	//
	// {"defaultAreaCode":"010", "isMatchCapital":"1"}
	ClidProperty *string `json:"ClidProperty,omitempty" xml:"ClidProperty,omitempty"`
	// 座席工号列表；callGroupType=1时必填。支持多个座席工号，多个之间使用英文逗号\\",\\"分隔注：一个座席只能在一个运行中的任务中
	//
	// example:
	//
	// 1111,2222
	Cnos *string `json:"Cnos,omitempty" xml:"Cnos,omitempty"`
	// 任务维度限制最大并发数，实际并发不会超过最大并发限制。type=1时，配置成0表示不限制，座席数量少于10时建议配置该参数
	//
	// example:
	//
	// 0
	Concurrency *int64 `json:"Concurrency,omitempty" xml:"Concurrency,omitempty"`
	// 客户侧外显规则；customerClidsCategory=1或2时生效。1.随机 2.按区号，默认随机。不建议使用
	//
	// example:
	//
	// 1
	CustomerClidType *int64 `json:"CustomerClidType,omitempty" xml:"CustomerClidType,omitempty"`
	// 外显号码比例配置内容；customerClidsCategory=1且customerClidWeightFlag=1时生效。JsonArray格式 [{"number":"123,345","weight":"5"}, {"number":"567,789","weight":"5"}]，不建议使用
	//
	// example:
	//
	// [{"number":"123,345","weight":"5"}, {"number":"567,789","weight":"5"}]
	CustomerClidWeight *string `json:"CustomerClidWeight,omitempty" xml:"CustomerClidWeight,omitempty"`
	// 外显号码比例配置开关；customerClidsCategory=1时生效。0.关闭 1.开启，默认关闭。不建议使用
	//
	// example:
	//
	// 0
	CustomerClidWeightFlag *int64 `json:"CustomerClidWeightFlag,omitempty" xml:"CustomerClidWeightFlag,omitempty"`
	// 客户侧外显号码列表；customerClidsCategory=1或2时必填。支持多个，多个直接使用英文逗号\\",\\"分隔。不建议使用
	//
	// example:
	//
	// 64203667,23434294
	CustomerClids *string `json:"CustomerClids,omitempty" xml:"CustomerClids,omitempty"`
	// 客户侧外显号码类型；1.外显固话 2.外显手机号 4.外显号码池 5.外显导航注：推荐使用外显导航，可以灵活调整和控制外显策略。其他类型将逐步下线
	//
	// example:
	//
	// 1
	CustomerClidsCategory *int64 `json:"CustomerClidsCategory,omitempty" xml:"CustomerClidsCategory,omitempty"`
	// 客户侧外显号码池名称或外显导航标识；customerClidsCategory=4或5时必填。customerClidsCategory=4时customerClidsGroup传号码池名称。customerClidsCategory=5时customerClidsGroup传外显导航标识
	//
	// example:
	//
	// 示例值
	CustomerClidsGroup *string `json:"CustomerClidsGroup,omitempty" xml:"CustomerClidsGroup,omitempty"`
	// 客户侧等待音；客户接听后呼叫座席，等待座席接听时播放的语音，默认为空白音。支持公共语音和企业语音，值为语音文件的path，如:60000011533526918819
	//
	// example:
	//
	// 60000011533526918819
	CustomerMoh *string `json:"CustomerMoh,omitempty" xml:"CustomerMoh,omitempty"`
	// 客户超时时间；单位秒数，默认30秒，取值范围5 ~ 60
	//
	// example:
	//
	// 55
	CustomerTimeout *int64 `json:"CustomerTimeout,omitempty" xml:"CustomerTimeout,omitempty"`
	// 客户侧提示音；客户接听后呼转座席前播放的语音，提示音播放完成后再找座席，默认无提示音。支持公共语音和企业语音，值为语音文件的path，如:60000011533526918819
	//
	// example:
	//
	// 60000011533526918819
	CustomerVoice *string `json:"CustomerVoice,omitempty" xml:"CustomerVoice,omitempty"`
	// 任务描述；需进行UTF-8格式的URLEncode编码，长度小于200字
	//
	// example:
	//
	// 示例值示例值示例值
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 呼叫中心 id
	//
	// This parameter is required.
	//
	// example:
	//
	// 7000002
	EnterpriseId *int64 `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	// 定时完成时强制结束任务；autoStop=1时生效。0.关闭 1.开启，默认关闭。开启配置后，当任务到定时完成时间时无论任务中的号码是否已经呼完，均将任务状态设置为结束。适用于对数据有呼叫时间限制的场景。注：任务在结束状态时无法再次启动。
	//
	// example:
	//
	// 0
	ForceEndFlag *int64 `json:"ForceEndFlag,omitempty" xml:"ForceEndFlag,omitempty"`
	// 暂停后重新预热；0.关闭，1开启，默认开启，任务暂停后是否需要重新预热
	//
	// example:
	//
	// 1
	IsRewarm *int64 `json:"IsRewarm,omitempty" xml:"IsRewarm,omitempty"`
	// 语音导航Id；参数ivrId和ivrName二选一，同时传入时ivrId生效，在创建自动外呼任务时ivrId或ivrName必选其一。预测外呼使用场景：如果客户接听后未在特定时间内（默认2秒）找到可用座席，通话将溢出到语音导航继续排队找座席。同时支持在呼转座席，座席未接听时也溢出到语音导航，需要开启企业配置生效自动外呼使用场景：客户接听后，转到的语音导航
	//
	// example:
	//
	// 13
	IvrId *int64 `json:"IvrId,omitempty" xml:"IvrId,omitempty"`
	// 语音导航名称；参数ivrId和ivrName二选一，同时传入时ivrId生效，在创建自动外呼任务时ivrId或ivrName必选其一。
	//
	// example:
	//
	// IvrName1
	IvrName *string `json:"IvrName,omitempty" xml:"IvrName,omitempty"`
	// 座席最长等待时间；默认40秒，最小10，最大600，允许座席空闲的最大时间。任务执行过程中，如果座席平均空闲时间大于maxWaitTime，每次呼叫个数会增加
	//
	// example:
	//
	// 40
	MaxWaitTime *int64 `json:"MaxWaitTime,omitempty" xml:"MaxWaitTime,omitempty"`
	// 最小可用座席数；默认为10，取值范围1 ~ 10。任务内可用座席数小于当前值时，任务自动暂停，避免骚扰
	//
	// example:
	//
	// 10
	MinAvailableAgentCount *int64 `json:"MinAvailableAgentCount,omitempty" xml:"MinAvailableAgentCount,omitempty"`
	// 任务名称；需进行UTF-8格式的URLEncode编码，长度小于50字
	//
	// example:
	//
	// 示例值示例值
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 超呼率；默认值为100，取值范围50～400
	//
	// example:
	//
	// 100
	PredictAdjust *int64 `json:"PredictAdjust,omitempty" xml:"PredictAdjust,omitempty"`
	// 骚扰率；默认值为1，取值范围为大于0，小于等于20，支持小数，精确到小数点两位。值越小呼叫的号码越少，值越大呼叫的号码越多，座席利用率越高
	//
	// example:
	//
	// 19.00
	Quotiety             *float64 `json:"Quotiety,omitempty" xml:"Quotiety,omitempty"`
	ResourceOwnerAccount *string  `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64   `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 重试策略；「基础模式」数据结构为Json格式```{"retry":2,"strategy":[{"round":1,"time":"1-1-1"},{"round":2,"time":"2-2-2"}]}```格式说明：retry是重试次数，round是第几次重试，time是第几次重试间隔的时间：1-1-1，第一个1是天，第二个1是小时，第三个1是分钟「高级模式」高级模式需要开启「号码状态识别」服务数据结构为JsonArray格式 ```[{"condition":{"sipCause":[710,719]},"retry":1,"sort":1,"strategy":[{"round":1,"time":"0-0-10"}]},{"condition":{"sipCause":[800]},"retry":1,"sort":2,"strategy":[{"round":2,"time":"0-0-10"}]}]```格式说明：condition是重试条件，sort是重试条件的顺序，其余字段与基础模式一致 注：基础模式和高级模式的区分根据传入的参数格式自动识别。基于首次呼叫时间重试时，每轮的重试时间必须大于上一轮的时间
	//
	// example:
	//
	// {"retry":2,"strategy":[{"round":1,"time":"1-1-1"},{"round":2,"time":"2-2-2"}]}
	RetryStrategy *string `json:"RetryStrategy,omitempty" xml:"RetryStrategy,omitempty"`
	// 任务仅当天生效；0.关闭 1.开启，删除待重试的数据和待呼叫的数据 2.开启，删除待重试的数据 3.开启，删除待呼叫的数据，默认关闭
	//
	// example:
	//
	// 0
	RetryStrategyOnlyToday *int64 `json:"RetryStrategyOnlyToday,omitempty" xml:"RetryStrategyOnlyToday,omitempty"`
	// 重试策略时间类型；配置重试策略时生效。 1.基于首次呼叫时间 2.基于上次呼叫时间，默认基于首次呼叫时间
	//
	// example:
	//
	// 1
	RetryStrategyTimeType *int64 `json:"RetryStrategyTimeType,omitempty" xml:"RetryStrategyTimeType,omitempty"`
	// 外呼任务Id
	//
	// This parameter is required.
	//
	// example:
	//
	// 133224
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// 禁呼时间；在特定的时间段内禁止呼叫。参数格式：时间条件编号。支持多个，多个使用英文","逗号隔开，如：9,22。注：时间条件列表可通过查询时间条件设置列表接口获取
	//
	// example:
	//
	// 9,22
	TimeStrategy *string `json:"TimeStrategy,omitempty" xml:"TimeStrategy,omitempty"`
	// 任务自定义字段。JsonArray格式；数组的每个元素只支持传递一组键值对，传递多组只取第一组，如 {"key", "value"}。 默认空值，示例：[{"name":"1234"},{"name1":"12345"}] 注：单个字段值的长度不能超过1000个字符
	//
	// example:
	//
	// [{"name":"1234"},{"name1":"12345"}]
	UserFields *string `json:"UserFields,omitempty" xml:"UserFields,omitempty"`
	// 任务预热时间；默认300秒, 取值范围60 ～ 600
	//
	// example:
	//
	// 370
	WarmUpDuration *int64 `json:"WarmUpDuration,omitempty" xml:"WarmUpDuration,omitempty"`
	// 座席整理时间；默认30秒，取值范围1～10800，整理时间会影响每次呼叫的个数，值越大，呼叫号码个数会减小
	//
	// example:
	//
	// 30
	Wrapup *int64 `json:"Wrapup,omitempty" xml:"Wrapup,omitempty"`
}

func (s CloudUpdateTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CloudUpdateTaskRequest) GoString() string {
	return s.String()
}

func (s *CloudUpdateTaskRequest) GetAgentGroup() *string {
	return s.AgentGroup
}

func (s *CloudUpdateTaskRequest) GetAgentTimeout() *int64 {
	return s.AgentTimeout
}

func (s *CloudUpdateTaskRequest) GetAnswerRate() *int64 {
	return s.AnswerRate
}

func (s *CloudUpdateTaskRequest) GetAutoComplete() *int64 {
	return s.AutoComplete
}

func (s *CloudUpdateTaskRequest) GetAutoStart() *int64 {
	return s.AutoStart
}

func (s *CloudUpdateTaskRequest) GetAutoStartDay() *string {
	return s.AutoStartDay
}

func (s *CloudUpdateTaskRequest) GetAutoStartTime() *string {
	return s.AutoStartTime
}

func (s *CloudUpdateTaskRequest) GetAutoStop() *int64 {
	return s.AutoStop
}

func (s *CloudUpdateTaskRequest) GetAutoStopDay() *string {
	return s.AutoStopDay
}

func (s *CloudUpdateTaskRequest) GetAutoStopTime() *string {
	return s.AutoStopTime
}

func (s *CloudUpdateTaskRequest) GetAutoTaskType() *int64 {
	return s.AutoTaskType
}

func (s *CloudUpdateTaskRequest) GetAutoTriggerTimeStrategy() *string {
	return s.AutoTriggerTimeStrategy
}

func (s *CloudUpdateTaskRequest) GetCallLimitStrategy() *string {
	return s.CallLimitStrategy
}

func (s *CloudUpdateTaskRequest) GetCallPriorityStrategy() *string {
	return s.CallPriorityStrategy
}

func (s *CloudUpdateTaskRequest) GetCallRouteStrategy() *int64 {
	return s.CallRouteStrategy
}

func (s *CloudUpdateTaskRequest) GetCallStrategy() *string {
	return s.CallStrategy
}

func (s *CloudUpdateTaskRequest) GetCallVariables() *string {
	return s.CallVariables
}

func (s *CloudUpdateTaskRequest) GetClidProperty() *string {
	return s.ClidProperty
}

func (s *CloudUpdateTaskRequest) GetCnos() *string {
	return s.Cnos
}

func (s *CloudUpdateTaskRequest) GetConcurrency() *int64 {
	return s.Concurrency
}

func (s *CloudUpdateTaskRequest) GetCustomerClidType() *int64 {
	return s.CustomerClidType
}

func (s *CloudUpdateTaskRequest) GetCustomerClidWeight() *string {
	return s.CustomerClidWeight
}

func (s *CloudUpdateTaskRequest) GetCustomerClidWeightFlag() *int64 {
	return s.CustomerClidWeightFlag
}

func (s *CloudUpdateTaskRequest) GetCustomerClids() *string {
	return s.CustomerClids
}

func (s *CloudUpdateTaskRequest) GetCustomerClidsCategory() *int64 {
	return s.CustomerClidsCategory
}

func (s *CloudUpdateTaskRequest) GetCustomerClidsGroup() *string {
	return s.CustomerClidsGroup
}

func (s *CloudUpdateTaskRequest) GetCustomerMoh() *string {
	return s.CustomerMoh
}

func (s *CloudUpdateTaskRequest) GetCustomerTimeout() *int64 {
	return s.CustomerTimeout
}

func (s *CloudUpdateTaskRequest) GetCustomerVoice() *string {
	return s.CustomerVoice
}

func (s *CloudUpdateTaskRequest) GetDescription() *string {
	return s.Description
}

func (s *CloudUpdateTaskRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *CloudUpdateTaskRequest) GetForceEndFlag() *int64 {
	return s.ForceEndFlag
}

func (s *CloudUpdateTaskRequest) GetIsRewarm() *int64 {
	return s.IsRewarm
}

func (s *CloudUpdateTaskRequest) GetIvrId() *int64 {
	return s.IvrId
}

func (s *CloudUpdateTaskRequest) GetIvrName() *string {
	return s.IvrName
}

func (s *CloudUpdateTaskRequest) GetMaxWaitTime() *int64 {
	return s.MaxWaitTime
}

func (s *CloudUpdateTaskRequest) GetMinAvailableAgentCount() *int64 {
	return s.MinAvailableAgentCount
}

func (s *CloudUpdateTaskRequest) GetName() *string {
	return s.Name
}

func (s *CloudUpdateTaskRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CloudUpdateTaskRequest) GetPredictAdjust() *int64 {
	return s.PredictAdjust
}

func (s *CloudUpdateTaskRequest) GetQuotiety() *float64 {
	return s.Quotiety
}

func (s *CloudUpdateTaskRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CloudUpdateTaskRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CloudUpdateTaskRequest) GetRetryStrategy() *string {
	return s.RetryStrategy
}

func (s *CloudUpdateTaskRequest) GetRetryStrategyOnlyToday() *int64 {
	return s.RetryStrategyOnlyToday
}

func (s *CloudUpdateTaskRequest) GetRetryStrategyTimeType() *int64 {
	return s.RetryStrategyTimeType
}

func (s *CloudUpdateTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *CloudUpdateTaskRequest) GetTimeStrategy() *string {
	return s.TimeStrategy
}

func (s *CloudUpdateTaskRequest) GetUserFields() *string {
	return s.UserFields
}

func (s *CloudUpdateTaskRequest) GetWarmUpDuration() *int64 {
	return s.WarmUpDuration
}

func (s *CloudUpdateTaskRequest) GetWrapup() *int64 {
	return s.Wrapup
}

func (s *CloudUpdateTaskRequest) SetAgentGroup(v string) *CloudUpdateTaskRequest {
	s.AgentGroup = &v
	return s
}

func (s *CloudUpdateTaskRequest) SetAgentTimeout(v int64) *CloudUpdateTaskRequest {
	s.AgentTimeout = &v
	return s
}

func (s *CloudUpdateTaskRequest) SetAnswerRate(v int64) *CloudUpdateTaskRequest {
	s.AnswerRate = &v
	return s
}

func (s *CloudUpdateTaskRequest) SetAutoComplete(v int64) *CloudUpdateTaskRequest {
	s.AutoComplete = &v
	return s
}

func (s *CloudUpdateTaskRequest) SetAutoStart(v int64) *CloudUpdateTaskRequest {
	s.AutoStart = &v
	return s
}

func (s *CloudUpdateTaskRequest) SetAutoStartDay(v string) *CloudUpdateTaskRequest {
	s.AutoStartDay = &v
	return s
}

func (s *CloudUpdateTaskRequest) SetAutoStartTime(v string) *CloudUpdateTaskRequest {
	s.AutoStartTime = &v
	return s
}

func (s *CloudUpdateTaskRequest) SetAutoStop(v int64) *CloudUpdateTaskRequest {
	s.AutoStop = &v
	return s
}

func (s *CloudUpdateTaskRequest) SetAutoStopDay(v string) *CloudUpdateTaskRequest {
	s.AutoStopDay = &v
	return s
}

func (s *CloudUpdateTaskRequest) SetAutoStopTime(v string) *CloudUpdateTaskRequest {
	s.AutoStopTime = &v
	return s
}

func (s *CloudUpdateTaskRequest) SetAutoTaskType(v int64) *CloudUpdateTaskRequest {
	s.AutoTaskType = &v
	return s
}

func (s *CloudUpdateTaskRequest) SetAutoTriggerTimeStrategy(v string) *CloudUpdateTaskRequest {
	s.AutoTriggerTimeStrategy = &v
	return s
}

func (s *CloudUpdateTaskRequest) SetCallLimitStrategy(v string) *CloudUpdateTaskRequest {
	s.CallLimitStrategy = &v
	return s
}

func (s *CloudUpdateTaskRequest) SetCallPriorityStrategy(v string) *CloudUpdateTaskRequest {
	s.CallPriorityStrategy = &v
	return s
}

func (s *CloudUpdateTaskRequest) SetCallRouteStrategy(v int64) *CloudUpdateTaskRequest {
	s.CallRouteStrategy = &v
	return s
}

func (s *CloudUpdateTaskRequest) SetCallStrategy(v string) *CloudUpdateTaskRequest {
	s.CallStrategy = &v
	return s
}

func (s *CloudUpdateTaskRequest) SetCallVariables(v string) *CloudUpdateTaskRequest {
	s.CallVariables = &v
	return s
}

func (s *CloudUpdateTaskRequest) SetClidProperty(v string) *CloudUpdateTaskRequest {
	s.ClidProperty = &v
	return s
}

func (s *CloudUpdateTaskRequest) SetCnos(v string) *CloudUpdateTaskRequest {
	s.Cnos = &v
	return s
}

func (s *CloudUpdateTaskRequest) SetConcurrency(v int64) *CloudUpdateTaskRequest {
	s.Concurrency = &v
	return s
}

func (s *CloudUpdateTaskRequest) SetCustomerClidType(v int64) *CloudUpdateTaskRequest {
	s.CustomerClidType = &v
	return s
}

func (s *CloudUpdateTaskRequest) SetCustomerClidWeight(v string) *CloudUpdateTaskRequest {
	s.CustomerClidWeight = &v
	return s
}

func (s *CloudUpdateTaskRequest) SetCustomerClidWeightFlag(v int64) *CloudUpdateTaskRequest {
	s.CustomerClidWeightFlag = &v
	return s
}

func (s *CloudUpdateTaskRequest) SetCustomerClids(v string) *CloudUpdateTaskRequest {
	s.CustomerClids = &v
	return s
}

func (s *CloudUpdateTaskRequest) SetCustomerClidsCategory(v int64) *CloudUpdateTaskRequest {
	s.CustomerClidsCategory = &v
	return s
}

func (s *CloudUpdateTaskRequest) SetCustomerClidsGroup(v string) *CloudUpdateTaskRequest {
	s.CustomerClidsGroup = &v
	return s
}

func (s *CloudUpdateTaskRequest) SetCustomerMoh(v string) *CloudUpdateTaskRequest {
	s.CustomerMoh = &v
	return s
}

func (s *CloudUpdateTaskRequest) SetCustomerTimeout(v int64) *CloudUpdateTaskRequest {
	s.CustomerTimeout = &v
	return s
}

func (s *CloudUpdateTaskRequest) SetCustomerVoice(v string) *CloudUpdateTaskRequest {
	s.CustomerVoice = &v
	return s
}

func (s *CloudUpdateTaskRequest) SetDescription(v string) *CloudUpdateTaskRequest {
	s.Description = &v
	return s
}

func (s *CloudUpdateTaskRequest) SetEnterpriseId(v int64) *CloudUpdateTaskRequest {
	s.EnterpriseId = &v
	return s
}

func (s *CloudUpdateTaskRequest) SetForceEndFlag(v int64) *CloudUpdateTaskRequest {
	s.ForceEndFlag = &v
	return s
}

func (s *CloudUpdateTaskRequest) SetIsRewarm(v int64) *CloudUpdateTaskRequest {
	s.IsRewarm = &v
	return s
}

func (s *CloudUpdateTaskRequest) SetIvrId(v int64) *CloudUpdateTaskRequest {
	s.IvrId = &v
	return s
}

func (s *CloudUpdateTaskRequest) SetIvrName(v string) *CloudUpdateTaskRequest {
	s.IvrName = &v
	return s
}

func (s *CloudUpdateTaskRequest) SetMaxWaitTime(v int64) *CloudUpdateTaskRequest {
	s.MaxWaitTime = &v
	return s
}

func (s *CloudUpdateTaskRequest) SetMinAvailableAgentCount(v int64) *CloudUpdateTaskRequest {
	s.MinAvailableAgentCount = &v
	return s
}

func (s *CloudUpdateTaskRequest) SetName(v string) *CloudUpdateTaskRequest {
	s.Name = &v
	return s
}

func (s *CloudUpdateTaskRequest) SetOwnerId(v int64) *CloudUpdateTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *CloudUpdateTaskRequest) SetPredictAdjust(v int64) *CloudUpdateTaskRequest {
	s.PredictAdjust = &v
	return s
}

func (s *CloudUpdateTaskRequest) SetQuotiety(v float64) *CloudUpdateTaskRequest {
	s.Quotiety = &v
	return s
}

func (s *CloudUpdateTaskRequest) SetResourceOwnerAccount(v string) *CloudUpdateTaskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CloudUpdateTaskRequest) SetResourceOwnerId(v int64) *CloudUpdateTaskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CloudUpdateTaskRequest) SetRetryStrategy(v string) *CloudUpdateTaskRequest {
	s.RetryStrategy = &v
	return s
}

func (s *CloudUpdateTaskRequest) SetRetryStrategyOnlyToday(v int64) *CloudUpdateTaskRequest {
	s.RetryStrategyOnlyToday = &v
	return s
}

func (s *CloudUpdateTaskRequest) SetRetryStrategyTimeType(v int64) *CloudUpdateTaskRequest {
	s.RetryStrategyTimeType = &v
	return s
}

func (s *CloudUpdateTaskRequest) SetTaskId(v string) *CloudUpdateTaskRequest {
	s.TaskId = &v
	return s
}

func (s *CloudUpdateTaskRequest) SetTimeStrategy(v string) *CloudUpdateTaskRequest {
	s.TimeStrategy = &v
	return s
}

func (s *CloudUpdateTaskRequest) SetUserFields(v string) *CloudUpdateTaskRequest {
	s.UserFields = &v
	return s
}

func (s *CloudUpdateTaskRequest) SetWarmUpDuration(v int64) *CloudUpdateTaskRequest {
	s.WarmUpDuration = &v
	return s
}

func (s *CloudUpdateTaskRequest) SetWrapup(v int64) *CloudUpdateTaskRequest {
	s.Wrapup = &v
	return s
}

func (s *CloudUpdateTaskRequest) Validate() error {
	return dara.Validate(s)
}
