// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetXConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetXConfigResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *GetXConfigResponseBody
	GetCode() *string
	SetData(v *GetXConfigResponseBodyData) *GetXConfigResponseBody
	GetData() *GetXConfigResponseBodyData
	SetMessage(v string) *GetXConfigResponseBody
	GetMessage() *string
	SetSuccess(v bool) *GetXConfigResponseBody
	GetSuccess() *bool
}

type GetXConfigResponseBody struct {
	// example:
	//
	// 0
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// 返回状态码 0000表示成功 其他表示失败
	//
	// example:
	//
	// 0000
	Code *string                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetXConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// 返回信息
	//
	// example:
	//
	// 成功
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 返回是否成功 true  表示成功 false表示失败
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetXConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetXConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetXConfigResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetXConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetXConfigResponseBody) GetData() *GetXConfigResponseBodyData {
	return s.Data
}

func (s *GetXConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetXConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetXConfigResponseBody) SetAccessDeniedDetail(v string) *GetXConfigResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetXConfigResponseBody) SetCode(v string) *GetXConfigResponseBody {
	s.Code = &v
	return s
}

func (s *GetXConfigResponseBody) SetData(v *GetXConfigResponseBodyData) *GetXConfigResponseBody {
	s.Data = v
	return s
}

func (s *GetXConfigResponseBody) SetMessage(v string) *GetXConfigResponseBody {
	s.Message = &v
	return s
}

func (s *GetXConfigResponseBody) SetSuccess(v bool) *GetXConfigResponseBody {
	s.Success = &v
	return s
}

func (s *GetXConfigResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetXConfigResponseBodyData struct {
	// 开/关呼叫能力状态： ‘0’：禁用； ‘1’：开启；
	//
	// example:
	//
	// 0
	CallAbility *string `json:"CallAbility,omitempty" xml:"CallAbility,omitempty"`
	// 是否透传来显为真实主叫： 00-非透传：互相拨打时都显示工作号; 10-透传：A客户为主叫时,B员工的来显为客户A号码;B员工为主叫时,A客户的来显为工作号; 默认为 00
	//
	// example:
	//
	// 00
	GNFlag *string `json:"GNFlag,omitempty" xml:"GNFlag,omitempty"`
	// 企业名片规则控制参数
	ReachJsons []*GetXConfigResponseBodyDataReachJsons `json:"ReachJsons,omitempty" xml:"ReachJsons,omitempty" type:"Repeated"`
	// 顺振控制参数
	SequenceCalls []*GetXConfigResponseBodyDataSequenceCalls `json:"SequenceCalls,omitempty" xml:"SequenceCalls,omitempty" type:"Repeated"`
	// 顺振结束时间 格式：HH:mm:ss 18:00:00
	//
	// example:
	//
	// 09:00:00
	SequenceEndTime *string `json:"SequenceEndTime,omitempty" xml:"SequenceEndTime,omitempty"`
	// 顺振开启时间 格式：HH:mm:ss 09:00:00
	//
	// example:
	//
	// 09:00:00
	SequenceStartTime *string `json:"SequenceStartTime,omitempty" xml:"SequenceStartTime,omitempty"`
	// 开/关短信功能状态： ‘0’：禁用； ‘1’：开启；
	//
	// example:
	//
	// 0
	SmsAbility *string `json:"SmsAbility,omitempty" xml:"SmsAbility,omitempty"`
	// 是否透传来显为真实主叫： 00-非透传：互相拨打时都显示工作号; 10-透传：A客户为主叫时,B员工的来显为客户A号码;B员工为主叫时,A客户的来显为工作号; 默认为 00
	//
	// example:
	//
	// 0
	SmsSignMode *string `json:"SmsSignMode,omitempty" xml:"SmsSignMode,omitempty"`
}

func (s GetXConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetXConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetXConfigResponseBodyData) GetCallAbility() *string {
	return s.CallAbility
}

func (s *GetXConfigResponseBodyData) GetGNFlag() *string {
	return s.GNFlag
}

func (s *GetXConfigResponseBodyData) GetReachJsons() []*GetXConfigResponseBodyDataReachJsons {
	return s.ReachJsons
}

func (s *GetXConfigResponseBodyData) GetSequenceCalls() []*GetXConfigResponseBodyDataSequenceCalls {
	return s.SequenceCalls
}

func (s *GetXConfigResponseBodyData) GetSequenceEndTime() *string {
	return s.SequenceEndTime
}

func (s *GetXConfigResponseBodyData) GetSequenceStartTime() *string {
	return s.SequenceStartTime
}

func (s *GetXConfigResponseBodyData) GetSmsAbility() *string {
	return s.SmsAbility
}

func (s *GetXConfigResponseBodyData) GetSmsSignMode() *string {
	return s.SmsSignMode
}

func (s *GetXConfigResponseBodyData) SetCallAbility(v string) *GetXConfigResponseBodyData {
	s.CallAbility = &v
	return s
}

func (s *GetXConfigResponseBodyData) SetGNFlag(v string) *GetXConfigResponseBodyData {
	s.GNFlag = &v
	return s
}

func (s *GetXConfigResponseBodyData) SetReachJsons(v []*GetXConfigResponseBodyDataReachJsons) *GetXConfigResponseBodyData {
	s.ReachJsons = v
	return s
}

func (s *GetXConfigResponseBodyData) SetSequenceCalls(v []*GetXConfigResponseBodyDataSequenceCalls) *GetXConfigResponseBodyData {
	s.SequenceCalls = v
	return s
}

func (s *GetXConfigResponseBodyData) SetSequenceEndTime(v string) *GetXConfigResponseBodyData {
	s.SequenceEndTime = &v
	return s
}

func (s *GetXConfigResponseBodyData) SetSequenceStartTime(v string) *GetXConfigResponseBodyData {
	s.SequenceStartTime = &v
	return s
}

func (s *GetXConfigResponseBodyData) SetSmsAbility(v string) *GetXConfigResponseBodyData {
	s.SmsAbility = &v
	return s
}

func (s *GetXConfigResponseBodyData) SetSmsSignMode(v string) *GetXConfigResponseBodyData {
	s.SmsSignMode = &v
	return s
}

func (s *GetXConfigResponseBodyData) Validate() error {
	if s.ReachJsons != nil {
		for _, item := range s.ReachJsons {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SequenceCalls != nil {
		for _, item := range s.SequenceCalls {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetXConfigResponseBodyDataReachJsons struct {
	// 呼叫方向 1:员工B呼叫客户A 2:客户A呼叫员工B
	//
	// example:
	//
	// 1
	CallDir *string `json:"CallDir,omitempty" xml:"CallDir,omitempty"`
	// 通话状态 1:通话振铃 2:接通前 3:接通后 4:通话结束 5:已接通6:未接通
	//
	// example:
	//
	// 1
	CallStatus *string `json:"CallStatus,omitempty" xml:"CallStatus,omitempty"`
	// 接收方向 1:主叫 2:被叫
	//
	// example:
	//
	// 1
	ReceiveDir *string `json:"ReceiveDir,omitempty" xml:"ReceiveDir,omitempty"`
	// 规则ID
	//
	// example:
	//
	// 345
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// 规则名称
	//
	// example:
	//
	// 企业名片-短信规则
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// 规则类型： 1：企业名片-短信 2：企业名片-闪信 3：企业名片-视频 4：企业名片-音频
	//
	// example:
	//
	// 1
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	// 模板ID
	//
	// example:
	//
	// 12345
	TempId *string `json:"TempId,omitempty" xml:"TempId,omitempty"`
}

func (s GetXConfigResponseBodyDataReachJsons) String() string {
	return dara.Prettify(s)
}

func (s GetXConfigResponseBodyDataReachJsons) GoString() string {
	return s.String()
}

func (s *GetXConfigResponseBodyDataReachJsons) GetCallDir() *string {
	return s.CallDir
}

func (s *GetXConfigResponseBodyDataReachJsons) GetCallStatus() *string {
	return s.CallStatus
}

func (s *GetXConfigResponseBodyDataReachJsons) GetReceiveDir() *string {
	return s.ReceiveDir
}

func (s *GetXConfigResponseBodyDataReachJsons) GetRuleId() *string {
	return s.RuleId
}

func (s *GetXConfigResponseBodyDataReachJsons) GetRuleName() *string {
	return s.RuleName
}

func (s *GetXConfigResponseBodyDataReachJsons) GetRuleType() *string {
	return s.RuleType
}

func (s *GetXConfigResponseBodyDataReachJsons) GetTempId() *string {
	return s.TempId
}

func (s *GetXConfigResponseBodyDataReachJsons) SetCallDir(v string) *GetXConfigResponseBodyDataReachJsons {
	s.CallDir = &v
	return s
}

func (s *GetXConfigResponseBodyDataReachJsons) SetCallStatus(v string) *GetXConfigResponseBodyDataReachJsons {
	s.CallStatus = &v
	return s
}

func (s *GetXConfigResponseBodyDataReachJsons) SetReceiveDir(v string) *GetXConfigResponseBodyDataReachJsons {
	s.ReceiveDir = &v
	return s
}

func (s *GetXConfigResponseBodyDataReachJsons) SetRuleId(v string) *GetXConfigResponseBodyDataReachJsons {
	s.RuleId = &v
	return s
}

func (s *GetXConfigResponseBodyDataReachJsons) SetRuleName(v string) *GetXConfigResponseBodyDataReachJsons {
	s.RuleName = &v
	return s
}

func (s *GetXConfigResponseBodyDataReachJsons) SetRuleType(v string) *GetXConfigResponseBodyDataReachJsons {
	s.RuleType = &v
	return s
}

func (s *GetXConfigResponseBodyDataReachJsons) SetTempId(v string) *GetXConfigResponseBodyDataReachJsons {
	s.TempId = &v
	return s
}

func (s *GetXConfigResponseBodyDataReachJsons) Validate() error {
	return dara.Validate(s)
}

type GetXConfigResponseBodyDataSequenceCalls struct {
	// 顺振提示音放音编号，格式如callNoPlayCode
	//
	// example:
	//
	// 示例值示例值示例值
	SequenceCallNoPlayCode *string `json:"SequenceCallNoPlayCode,omitempty" xml:"SequenceCallNoPlayCode,omitempty"`
	// 顺振被叫号码
	//
	// example:
	//
	// 示例值示例值示例值
	SequenceCalledNo *string `json:"SequenceCalledNo,omitempty" xml:"SequenceCalledNo,omitempty"`
	// 接通后主被叫放音编号，格式如calledPlayCode
	//
	// example:
	//
	// 示例值示例值示例值
	SequenceCalledPlayCode *string `json:"SequenceCalledPlayCode,omitempty" xml:"SequenceCalledPlayCode,omitempty"`
}

func (s GetXConfigResponseBodyDataSequenceCalls) String() string {
	return dara.Prettify(s)
}

func (s GetXConfigResponseBodyDataSequenceCalls) GoString() string {
	return s.String()
}

func (s *GetXConfigResponseBodyDataSequenceCalls) GetSequenceCallNoPlayCode() *string {
	return s.SequenceCallNoPlayCode
}

func (s *GetXConfigResponseBodyDataSequenceCalls) GetSequenceCalledNo() *string {
	return s.SequenceCalledNo
}

func (s *GetXConfigResponseBodyDataSequenceCalls) GetSequenceCalledPlayCode() *string {
	return s.SequenceCalledPlayCode
}

func (s *GetXConfigResponseBodyDataSequenceCalls) SetSequenceCallNoPlayCode(v string) *GetXConfigResponseBodyDataSequenceCalls {
	s.SequenceCallNoPlayCode = &v
	return s
}

func (s *GetXConfigResponseBodyDataSequenceCalls) SetSequenceCalledNo(v string) *GetXConfigResponseBodyDataSequenceCalls {
	s.SequenceCalledNo = &v
	return s
}

func (s *GetXConfigResponseBodyDataSequenceCalls) SetSequenceCalledPlayCode(v string) *GetXConfigResponseBodyDataSequenceCalls {
	s.SequenceCalledPlayCode = &v
	return s
}

func (s *GetXConfigResponseBodyDataSequenceCalls) Validate() error {
	return dara.Validate(s)
}
