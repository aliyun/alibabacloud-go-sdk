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
	// Specifies whether to enable the call capability. Valid values:
	//
	// - 0: Disable.
	//
	// - 1: Enable.
	//
	// example:
	//
	// 0
	CallAbility *string `json:"CallAbility,omitempty" xml:"CallAbility,omitempty"`
	// The customer UID.
	//
	// example:
	//
	// 1898871967585852
	CallerParentId *int64 `json:"CallerParentId,omitempty" xml:"CallerParentId,omitempty"`
	// The key of the number pool.
	//
	// This parameter is required.
	//
	// example:
	//
	// FC5**********************a1a
	CustomerPoolKey *string `json:"CustomerPoolKey,omitempty" xml:"CustomerPoolKey,omitempty"`
	// Specifies whether to pass through the caller ID to display the actual caller. Valid values:
	//
	// - 00: No pass-through. The work number is displayed when either party calls the other.
	//
	// - 10: Pass-through. When Customer A is the caller, the caller ID displayed to Employee B is the number of Customer A. When Employee B is the caller, the caller ID displayed to Customer A is the work number.
	//
	// Default value: 00.
	//
	// example:
	//
	// 0
	GNFlag  *string `json:"GNFlag,omitempty" xml:"GNFlag,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The request deduplication ID. The maximum length of reqId is 20 characters. The caller must ensure atomicity.
	//
	// This parameter is required.
	//
	// example:
	//
	// 564**********879
	ReqId                *string `json:"ReqId,omitempty" xml:"ReqId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The sequential ringing control parameters.
	SequenceCallsShrink *string `json:"SequenceCalls,omitempty" xml:"SequenceCalls,omitempty"`
	// The sequential ringing mode. Valid values:
	//
	// - 0: No sequential ringing. This is the default value.
	//
	// - 1: Conditional sequential ringing. The system first connects to the callee specified by calledNo. If the callee cannot be connected, the system sequentially rings the numbers in the sequenceCalls list.
	//
	// - 2: Unconditional sequential ringing. The system does not connect to the callee specified by calledNo and directly sequentially rings the numbers in the sequenceCalls list.
	//
	// example:
	//
	// 0
	SequenceMode *string `json:"SequenceMode,omitempty" xml:"SequenceMode,omitempty"`
	// Specifies whether to enable the SMS feature. Valid values:
	//
	// - 0: Disable.
	//
	// - 1: Enable.
	//
	// example:
	//
	// 0
	SmsAbility *string `json:"SmsAbility,omitempty" xml:"SmsAbility,omitempty"`
	// Specifies whether to pass through the SMS signature to display the actual user. Valid values:
	//
	// - 0: No pass-through.
	//
	// - 1: Pass-through.
	//
	// Default value: 0.
	//
	// example:
	//
	// 0
	SmsSignMode *string `json:"SmsSignMode,omitempty" xml:"SmsSignMode,omitempty"`
	// The X number.
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
