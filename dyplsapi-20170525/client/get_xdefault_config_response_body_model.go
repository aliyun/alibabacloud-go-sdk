// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetXDefaultConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetXDefaultConfigResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *GetXDefaultConfigResponseBody
	GetCode() *string
	SetData(v *GetXDefaultConfigResponseBodyData) *GetXDefaultConfigResponseBody
	GetData() *GetXDefaultConfigResponseBodyData
	SetMessage(v string) *GetXDefaultConfigResponseBody
	GetMessage() *string
	SetSuccess(v bool) *GetXDefaultConfigResponseBody
	GetSuccess() *bool
}

type GetXDefaultConfigResponseBody struct {
	// example:
	//
	// 0
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// 返回状态码 0000表示成功 其他表示失败
	//
	// example:
	//
	// 0000
	Code *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetXDefaultConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s GetXDefaultConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetXDefaultConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetXDefaultConfigResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetXDefaultConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetXDefaultConfigResponseBody) GetData() *GetXDefaultConfigResponseBodyData {
	return s.Data
}

func (s *GetXDefaultConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetXDefaultConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetXDefaultConfigResponseBody) SetAccessDeniedDetail(v string) *GetXDefaultConfigResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetXDefaultConfigResponseBody) SetCode(v string) *GetXDefaultConfigResponseBody {
	s.Code = &v
	return s
}

func (s *GetXDefaultConfigResponseBody) SetData(v *GetXDefaultConfigResponseBodyData) *GetXDefaultConfigResponseBody {
	s.Data = v
	return s
}

func (s *GetXDefaultConfigResponseBody) SetMessage(v string) *GetXDefaultConfigResponseBody {
	s.Message = &v
	return s
}

func (s *GetXDefaultConfigResponseBody) SetSuccess(v bool) *GetXDefaultConfigResponseBody {
	s.Success = &v
	return s
}

func (s *GetXDefaultConfigResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetXDefaultConfigResponseBodyData struct {
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
	ReachJson []*GetXDefaultConfigResponseBodyDataReachJson `json:"ReachJson,omitempty" xml:"ReachJson,omitempty" type:"Repeated"`
	// 顺振控制参数
	SequenceCall []*GetXDefaultConfigResponseBodyDataSequenceCall `json:"SequenceCall,omitempty" xml:"SequenceCall,omitempty" type:"Repeated"`
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

func (s GetXDefaultConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetXDefaultConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetXDefaultConfigResponseBodyData) GetCallAbility() *string {
	return s.CallAbility
}

func (s *GetXDefaultConfigResponseBodyData) GetGNFlag() *string {
	return s.GNFlag
}

func (s *GetXDefaultConfigResponseBodyData) GetReachJson() []*GetXDefaultConfigResponseBodyDataReachJson {
	return s.ReachJson
}

func (s *GetXDefaultConfigResponseBodyData) GetSequenceCall() []*GetXDefaultConfigResponseBodyDataSequenceCall {
	return s.SequenceCall
}

func (s *GetXDefaultConfigResponseBodyData) GetSequenceEndTime() *string {
	return s.SequenceEndTime
}

func (s *GetXDefaultConfigResponseBodyData) GetSequenceStartTime() *string {
	return s.SequenceStartTime
}

func (s *GetXDefaultConfigResponseBodyData) GetSmsAbility() *string {
	return s.SmsAbility
}

func (s *GetXDefaultConfigResponseBodyData) GetSmsSignMode() *string {
	return s.SmsSignMode
}

func (s *GetXDefaultConfigResponseBodyData) SetCallAbility(v string) *GetXDefaultConfigResponseBodyData {
	s.CallAbility = &v
	return s
}

func (s *GetXDefaultConfigResponseBodyData) SetGNFlag(v string) *GetXDefaultConfigResponseBodyData {
	s.GNFlag = &v
	return s
}

func (s *GetXDefaultConfigResponseBodyData) SetReachJson(v []*GetXDefaultConfigResponseBodyDataReachJson) *GetXDefaultConfigResponseBodyData {
	s.ReachJson = v
	return s
}

func (s *GetXDefaultConfigResponseBodyData) SetSequenceCall(v []*GetXDefaultConfigResponseBodyDataSequenceCall) *GetXDefaultConfigResponseBodyData {
	s.SequenceCall = v
	return s
}

func (s *GetXDefaultConfigResponseBodyData) SetSequenceEndTime(v string) *GetXDefaultConfigResponseBodyData {
	s.SequenceEndTime = &v
	return s
}

func (s *GetXDefaultConfigResponseBodyData) SetSequenceStartTime(v string) *GetXDefaultConfigResponseBodyData {
	s.SequenceStartTime = &v
	return s
}

func (s *GetXDefaultConfigResponseBodyData) SetSmsAbility(v string) *GetXDefaultConfigResponseBodyData {
	s.SmsAbility = &v
	return s
}

func (s *GetXDefaultConfigResponseBodyData) SetSmsSignMode(v string) *GetXDefaultConfigResponseBodyData {
	s.SmsSignMode = &v
	return s
}

func (s *GetXDefaultConfigResponseBodyData) Validate() error {
	if s.ReachJson != nil {
		for _, item := range s.ReachJson {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SequenceCall != nil {
		for _, item := range s.SequenceCall {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetXDefaultConfigResponseBodyDataReachJson struct {
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

func (s GetXDefaultConfigResponseBodyDataReachJson) String() string {
	return dara.Prettify(s)
}

func (s GetXDefaultConfigResponseBodyDataReachJson) GoString() string {
	return s.String()
}

func (s *GetXDefaultConfigResponseBodyDataReachJson) GetCallDir() *string {
	return s.CallDir
}

func (s *GetXDefaultConfigResponseBodyDataReachJson) GetCallStatus() *string {
	return s.CallStatus
}

func (s *GetXDefaultConfigResponseBodyDataReachJson) GetReceiveDir() *string {
	return s.ReceiveDir
}

func (s *GetXDefaultConfigResponseBodyDataReachJson) GetRuleId() *string {
	return s.RuleId
}

func (s *GetXDefaultConfigResponseBodyDataReachJson) GetRuleName() *string {
	return s.RuleName
}

func (s *GetXDefaultConfigResponseBodyDataReachJson) GetRuleType() *string {
	return s.RuleType
}

func (s *GetXDefaultConfigResponseBodyDataReachJson) GetTempId() *string {
	return s.TempId
}

func (s *GetXDefaultConfigResponseBodyDataReachJson) SetCallDir(v string) *GetXDefaultConfigResponseBodyDataReachJson {
	s.CallDir = &v
	return s
}

func (s *GetXDefaultConfigResponseBodyDataReachJson) SetCallStatus(v string) *GetXDefaultConfigResponseBodyDataReachJson {
	s.CallStatus = &v
	return s
}

func (s *GetXDefaultConfigResponseBodyDataReachJson) SetReceiveDir(v string) *GetXDefaultConfigResponseBodyDataReachJson {
	s.ReceiveDir = &v
	return s
}

func (s *GetXDefaultConfigResponseBodyDataReachJson) SetRuleId(v string) *GetXDefaultConfigResponseBodyDataReachJson {
	s.RuleId = &v
	return s
}

func (s *GetXDefaultConfigResponseBodyDataReachJson) SetRuleName(v string) *GetXDefaultConfigResponseBodyDataReachJson {
	s.RuleName = &v
	return s
}

func (s *GetXDefaultConfigResponseBodyDataReachJson) SetRuleType(v string) *GetXDefaultConfigResponseBodyDataReachJson {
	s.RuleType = &v
	return s
}

func (s *GetXDefaultConfigResponseBodyDataReachJson) SetTempId(v string) *GetXDefaultConfigResponseBodyDataReachJson {
	s.TempId = &v
	return s
}

func (s *GetXDefaultConfigResponseBodyDataReachJson) Validate() error {
	return dara.Validate(s)
}

type GetXDefaultConfigResponseBodyDataSequenceCall struct {
	// 顺振提示音放音编号，格式如callNoPlayCode
	//
	// example:
	//
	// 示例值
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

func (s GetXDefaultConfigResponseBodyDataSequenceCall) String() string {
	return dara.Prettify(s)
}

func (s GetXDefaultConfigResponseBodyDataSequenceCall) GoString() string {
	return s.String()
}

func (s *GetXDefaultConfigResponseBodyDataSequenceCall) GetSequenceCallNoPlayCode() *string {
	return s.SequenceCallNoPlayCode
}

func (s *GetXDefaultConfigResponseBodyDataSequenceCall) GetSequenceCalledNo() *string {
	return s.SequenceCalledNo
}

func (s *GetXDefaultConfigResponseBodyDataSequenceCall) GetSequenceCalledPlayCode() *string {
	return s.SequenceCalledPlayCode
}

func (s *GetXDefaultConfigResponseBodyDataSequenceCall) SetSequenceCallNoPlayCode(v string) *GetXDefaultConfigResponseBodyDataSequenceCall {
	s.SequenceCallNoPlayCode = &v
	return s
}

func (s *GetXDefaultConfigResponseBodyDataSequenceCall) SetSequenceCalledNo(v string) *GetXDefaultConfigResponseBodyDataSequenceCall {
	s.SequenceCalledNo = &v
	return s
}

func (s *GetXDefaultConfigResponseBodyDataSequenceCall) SetSequenceCalledPlayCode(v string) *GetXDefaultConfigResponseBodyDataSequenceCall {
	s.SequenceCalledPlayCode = &v
	return s
}

func (s *GetXDefaultConfigResponseBodyDataSequenceCall) Validate() error {
	return dara.Validate(s)
}
