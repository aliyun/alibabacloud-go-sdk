// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigXShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallAbility(v string) *ConfigXShrinkRequest
	GetCallAbility() *string
	SetCallerParentId(v int64) *ConfigXShrinkRequest
	GetCallerParentId() *int64
	SetCustomerPoolKey(v string) *ConfigXShrinkRequest
	GetCustomerPoolKey() *string
	SetGNFlag(v string) *ConfigXShrinkRequest
	GetGNFlag() *string
	SetOwnerId(v int64) *ConfigXShrinkRequest
	GetOwnerId() *int64
	SetReqId(v string) *ConfigXShrinkRequest
	GetReqId() *string
	SetResourceOwnerAccount(v string) *ConfigXShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ConfigXShrinkRequest
	GetResourceOwnerId() *int64
	SetSequenceCallsShrink(v string) *ConfigXShrinkRequest
	GetSequenceCallsShrink() *string
	SetSequenceMode(v string) *ConfigXShrinkRequest
	GetSequenceMode() *string
	SetSmsAbility(v string) *ConfigXShrinkRequest
	GetSmsAbility() *string
	SetSmsSignMode(v string) *ConfigXShrinkRequest
	GetSmsSignMode() *string
	SetTelX(v string) *ConfigXShrinkRequest
	GetTelX() *string
}

type ConfigXShrinkRequest struct {
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
	SequenceCallsShrink *string `json:"SequenceCalls,omitempty" xml:"SequenceCalls,omitempty"`
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

func (s ConfigXShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ConfigXShrinkRequest) GoString() string {
	return s.String()
}

func (s *ConfigXShrinkRequest) GetCallAbility() *string {
	return s.CallAbility
}

func (s *ConfigXShrinkRequest) GetCallerParentId() *int64 {
	return s.CallerParentId
}

func (s *ConfigXShrinkRequest) GetCustomerPoolKey() *string {
	return s.CustomerPoolKey
}

func (s *ConfigXShrinkRequest) GetGNFlag() *string {
	return s.GNFlag
}

func (s *ConfigXShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ConfigXShrinkRequest) GetReqId() *string {
	return s.ReqId
}

func (s *ConfigXShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ConfigXShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ConfigXShrinkRequest) GetSequenceCallsShrink() *string {
	return s.SequenceCallsShrink
}

func (s *ConfigXShrinkRequest) GetSequenceMode() *string {
	return s.SequenceMode
}

func (s *ConfigXShrinkRequest) GetSmsAbility() *string {
	return s.SmsAbility
}

func (s *ConfigXShrinkRequest) GetSmsSignMode() *string {
	return s.SmsSignMode
}

func (s *ConfigXShrinkRequest) GetTelX() *string {
	return s.TelX
}

func (s *ConfigXShrinkRequest) SetCallAbility(v string) *ConfigXShrinkRequest {
	s.CallAbility = &v
	return s
}

func (s *ConfigXShrinkRequest) SetCallerParentId(v int64) *ConfigXShrinkRequest {
	s.CallerParentId = &v
	return s
}

func (s *ConfigXShrinkRequest) SetCustomerPoolKey(v string) *ConfigXShrinkRequest {
	s.CustomerPoolKey = &v
	return s
}

func (s *ConfigXShrinkRequest) SetGNFlag(v string) *ConfigXShrinkRequest {
	s.GNFlag = &v
	return s
}

func (s *ConfigXShrinkRequest) SetOwnerId(v int64) *ConfigXShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *ConfigXShrinkRequest) SetReqId(v string) *ConfigXShrinkRequest {
	s.ReqId = &v
	return s
}

func (s *ConfigXShrinkRequest) SetResourceOwnerAccount(v string) *ConfigXShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ConfigXShrinkRequest) SetResourceOwnerId(v int64) *ConfigXShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ConfigXShrinkRequest) SetSequenceCallsShrink(v string) *ConfigXShrinkRequest {
	s.SequenceCallsShrink = &v
	return s
}

func (s *ConfigXShrinkRequest) SetSequenceMode(v string) *ConfigXShrinkRequest {
	s.SequenceMode = &v
	return s
}

func (s *ConfigXShrinkRequest) SetSmsAbility(v string) *ConfigXShrinkRequest {
	s.SmsAbility = &v
	return s
}

func (s *ConfigXShrinkRequest) SetSmsSignMode(v string) *ConfigXShrinkRequest {
	s.SmsSignMode = &v
	return s
}

func (s *ConfigXShrinkRequest) SetTelX(v string) *ConfigXShrinkRequest {
	s.TelX = &v
	return s
}

func (s *ConfigXShrinkRequest) Validate() error {
	return dara.Validate(s)
}
