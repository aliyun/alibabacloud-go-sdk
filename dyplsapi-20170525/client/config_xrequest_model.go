// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigXRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallAbility(v string) *ConfigXRequest
	GetCallAbility() *string
	SetCallerParentId(v int64) *ConfigXRequest
	GetCallerParentId() *int64
	SetCustomerPoolKey(v string) *ConfigXRequest
	GetCustomerPoolKey() *string
	SetGNFlag(v string) *ConfigXRequest
	GetGNFlag() *string
	SetOwnerId(v int64) *ConfigXRequest
	GetOwnerId() *int64
	SetReqId(v string) *ConfigXRequest
	GetReqId() *string
	SetResourceOwnerAccount(v string) *ConfigXRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ConfigXRequest
	GetResourceOwnerId() *int64
	SetSequenceCalls(v []*ConfigXRequestSequenceCalls) *ConfigXRequest
	GetSequenceCalls() []*ConfigXRequestSequenceCalls
	SetSequenceMode(v string) *ConfigXRequest
	GetSequenceMode() *string
	SetSmsAbility(v string) *ConfigXRequest
	GetSmsAbility() *string
	SetSmsSignMode(v string) *ConfigXRequest
	GetSmsSignMode() *string
	SetTelX(v string) *ConfigXRequest
	GetTelX() *string
}

type ConfigXRequest struct {
	// 开/关呼叫能力状态‘0’：禁用‘1’：开启
	//
	// example:
	//
	// 0
	CallAbility *string `json:"CallAbility,omitempty" xml:"CallAbility,omitempty"`
	// 客户uid
	//
	// example:
	//
	// 1898871967585852
	CallerParentId *int64 `json:"CallerParentId,omitempty" xml:"CallerParentId,omitempty"`
	// 号码池key
	//
	// This parameter is required.
	//
	// example:
	//
	// FC5**********************a1a
	CustomerPoolKey *string `json:"CustomerPoolKey,omitempty" xml:"CustomerPoolKey,omitempty"`
	// 是否透传来显为真实主叫：00-非透传：互相拨打时都显示工作号;10-透传：A客户为主叫时,B员工的来显为客户A号码;B员工为主叫时,A客户的来显为工作号;默认为 00
	//
	// example:
	//
	// 0
	GNFlag  *string `json:"GNFlag,omitempty" xml:"GNFlag,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 请求去重ID, reqId最大长度为20位,接入方需要保持原子性
	//
	// This parameter is required.
	//
	// example:
	//
	// 564**********879
	ReqId                *string `json:"ReqId,omitempty" xml:"ReqId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 顺振控制参数
	SequenceCalls []*ConfigXRequestSequenceCalls `json:"SequenceCalls,omitempty" xml:"SequenceCalls,omitempty" type:"Repeated"`
	// 顺振模式：0-不顺振（默认）1-有条件顺振，先接续calledNo指定被叫，如果该被叫未能接通，再顺振sequenceCalls号码列表2-无条件顺振，不接续calledNo指定被叫，直接顺振sequenceCalls号码列表
	//
	// example:
	//
	// 0
	SequenceMode *string `json:"SequenceMode,omitempty" xml:"SequenceMode,omitempty"`
	// 开/关短信功能状态‘0’：禁用；‘1’：开启；
	//
	// example:
	//
	// 0
	SmsAbility *string `json:"SmsAbility,omitempty" xml:"SmsAbility,omitempty"`
	// 是否透传来显为真实用户0：不透传; 1：透传默认：0不透传
	//
	// example:
	//
	// 0
	SmsSignMode *string `json:"SmsSignMode,omitempty" xml:"SmsSignMode,omitempty"`
	// X号码
	//
	// This parameter is required.
	//
	// example:
	//
	// 17*******22
	TelX *string `json:"TelX,omitempty" xml:"TelX,omitempty"`
}

func (s ConfigXRequest) String() string {
	return dara.Prettify(s)
}

func (s ConfigXRequest) GoString() string {
	return s.String()
}

func (s *ConfigXRequest) GetCallAbility() *string {
	return s.CallAbility
}

func (s *ConfigXRequest) GetCallerParentId() *int64 {
	return s.CallerParentId
}

func (s *ConfigXRequest) GetCustomerPoolKey() *string {
	return s.CustomerPoolKey
}

func (s *ConfigXRequest) GetGNFlag() *string {
	return s.GNFlag
}

func (s *ConfigXRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ConfigXRequest) GetReqId() *string {
	return s.ReqId
}

func (s *ConfigXRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ConfigXRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ConfigXRequest) GetSequenceCalls() []*ConfigXRequestSequenceCalls {
	return s.SequenceCalls
}

func (s *ConfigXRequest) GetSequenceMode() *string {
	return s.SequenceMode
}

func (s *ConfigXRequest) GetSmsAbility() *string {
	return s.SmsAbility
}

func (s *ConfigXRequest) GetSmsSignMode() *string {
	return s.SmsSignMode
}

func (s *ConfigXRequest) GetTelX() *string {
	return s.TelX
}

func (s *ConfigXRequest) SetCallAbility(v string) *ConfigXRequest {
	s.CallAbility = &v
	return s
}

func (s *ConfigXRequest) SetCallerParentId(v int64) *ConfigXRequest {
	s.CallerParentId = &v
	return s
}

func (s *ConfigXRequest) SetCustomerPoolKey(v string) *ConfigXRequest {
	s.CustomerPoolKey = &v
	return s
}

func (s *ConfigXRequest) SetGNFlag(v string) *ConfigXRequest {
	s.GNFlag = &v
	return s
}

func (s *ConfigXRequest) SetOwnerId(v int64) *ConfigXRequest {
	s.OwnerId = &v
	return s
}

func (s *ConfigXRequest) SetReqId(v string) *ConfigXRequest {
	s.ReqId = &v
	return s
}

func (s *ConfigXRequest) SetResourceOwnerAccount(v string) *ConfigXRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ConfigXRequest) SetResourceOwnerId(v int64) *ConfigXRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ConfigXRequest) SetSequenceCalls(v []*ConfigXRequestSequenceCalls) *ConfigXRequest {
	s.SequenceCalls = v
	return s
}

func (s *ConfigXRequest) SetSequenceMode(v string) *ConfigXRequest {
	s.SequenceMode = &v
	return s
}

func (s *ConfigXRequest) SetSmsAbility(v string) *ConfigXRequest {
	s.SmsAbility = &v
	return s
}

func (s *ConfigXRequest) SetSmsSignMode(v string) *ConfigXRequest {
	s.SmsSignMode = &v
	return s
}

func (s *ConfigXRequest) SetTelX(v string) *ConfigXRequest {
	s.TelX = &v
	return s
}

func (s *ConfigXRequest) Validate() error {
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

type ConfigXRequestSequenceCalls struct {
	// 顺振提示音放音编号，格式如callNoPlayCode
	//
	// example:
	//
	// 01
	SequenceCallNoPlayCode *string `json:"SequenceCallNoPlayCode,omitempty" xml:"SequenceCallNoPlayCode,omitempty"`
	// 顺振被叫号码
	//
	// example:
	//
	// 18*******33
	SequenceCalledNo *string `json:"SequenceCalledNo,omitempty" xml:"SequenceCalledNo,omitempty"`
	// 接通后主被叫放音编号，格式如calledPlayCode
	//
	// example:
	//
	// 02
	SequenceCalledPlayCode *string `json:"SequenceCalledPlayCode,omitempty" xml:"SequenceCalledPlayCode,omitempty"`
}

func (s ConfigXRequestSequenceCalls) String() string {
	return dara.Prettify(s)
}

func (s ConfigXRequestSequenceCalls) GoString() string {
	return s.String()
}

func (s *ConfigXRequestSequenceCalls) GetSequenceCallNoPlayCode() *string {
	return s.SequenceCallNoPlayCode
}

func (s *ConfigXRequestSequenceCalls) GetSequenceCalledNo() *string {
	return s.SequenceCalledNo
}

func (s *ConfigXRequestSequenceCalls) GetSequenceCalledPlayCode() *string {
	return s.SequenceCalledPlayCode
}

func (s *ConfigXRequestSequenceCalls) SetSequenceCallNoPlayCode(v string) *ConfigXRequestSequenceCalls {
	s.SequenceCallNoPlayCode = &v
	return s
}

func (s *ConfigXRequestSequenceCalls) SetSequenceCalledNo(v string) *ConfigXRequestSequenceCalls {
	s.SequenceCalledNo = &v
	return s
}

func (s *ConfigXRequestSequenceCalls) SetSequenceCalledPlayCode(v string) *ConfigXRequestSequenceCalls {
	s.SequenceCalledPlayCode = &v
	return s
}

func (s *ConfigXRequestSequenceCalls) Validate() error {
	return dara.Validate(s)
}
