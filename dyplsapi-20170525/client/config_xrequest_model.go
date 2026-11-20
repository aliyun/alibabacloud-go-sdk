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
	SequenceCalls []*ConfigXRequestSequenceCalls `json:"SequenceCalls,omitempty" xml:"SequenceCalls,omitempty" type:"Repeated"`
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
	// The playback code for the sequential ringing prompt tone. The format is the same as callNoPlayCode.
	//
	// example:
	//
	// 01
	SequenceCallNoPlayCode *string `json:"SequenceCallNoPlayCode,omitempty" xml:"SequenceCallNoPlayCode,omitempty"`
	// The callee number for sequential ringing.
	//
	// example:
	//
	// 18*******33
	SequenceCalledNo *string `json:"SequenceCalledNo,omitempty" xml:"SequenceCalledNo,omitempty"`
	// The playback code for the caller and callee after the call is connected. The format is the same as calledPlayCode.
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
