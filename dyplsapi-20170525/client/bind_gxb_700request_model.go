// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindGxb700Request interface {
	dara.Model
	String() string
	GoString() string
	SetAsrModelId(v string) *BindGxb700Request
	GetAsrModelId() *string
	SetAudio(v string) *BindGxb700Request
	GetAudio() *string
	SetCallRestrict(v string) *BindGxb700Request
	GetCallRestrict() *string
	SetCallTimeout(v int64) *BindGxb700Request
	GetCallTimeout() *int64
	SetDefaultA(v string) *BindGxb700Request
	GetDefaultA() *string
	SetDtmfConfig(v string) *BindGxb700Request
	GetDtmfConfig() *string
	SetExpiration(v string) *BindGxb700Request
	GetExpiration() *string
	SetGroupId(v string) *BindGxb700Request
	GetGroupId() *string
	SetIndustrialId(v string) *BindGxb700Request
	GetIndustrialId() *string
	SetNeedAsr(v bool) *BindGxb700Request
	GetNeedAsr() *bool
	SetNeedRecord(v bool) *BindGxb700Request
	GetNeedRecord() *bool
	SetOrderId(v string) *BindGxb700Request
	GetOrderId() *string
	SetOutId(v string) *BindGxb700Request
	GetOutId() *string
	SetOutOrderId(v string) *BindGxb700Request
	GetOutOrderId() *string
	SetOwnerId(v int64) *BindGxb700Request
	GetOwnerId() *int64
	SetPoolKey(v string) *BindGxb700Request
	GetPoolKey() *string
	SetRecType(v string) *BindGxb700Request
	GetRecType() *string
	SetResourceOwnerAccount(v string) *BindGxb700Request
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *BindGxb700Request
	GetResourceOwnerId() *int64
	SetTelB(v string) *BindGxb700Request
	GetTelB() *string
	SetTelX(v string) *BindGxb700Request
	GetTelX() *string
}

type BindGxb700Request struct {
	// The ASR model ID. Valid values:
	//
	// - General-Phone-Chinese Mandarin model ID: 7ee372834d2f4cc7ac0d0ab2d0ae1aac
	//
	// - General-Phone-Chinese Cantonese model ID: 9b4afe3d224e4fa3bf963697113e2f77
	//
	// - General-Phone-Regional Accent model ID: 3655a6cb94764145a1f9a38a3c3b3b4b.
	//
	// example:
	//
	// 7ee372834d2f4cc7ac0d0ab2d0ae1aac
	AsrModelId *string `json:"AsrModelId,omitempty" xml:"AsrModelId,omitempty"`
	// The audio prompt code.
	//
	//
	// - GCallX_Gpre: the audio prompt code played to the caller before the call is connected when a number in group G calls number X.
	//
	// - GCallX_Gin: the audio prompt code played to the caller after the call is connected when a number in group G calls number X.
	//
	// - GCallX_Bin: the audio prompt code played to user B after the call is connected when a number in group G calls number X.
	//
	// - BCallX_Bpre: the audio prompt code played to user B before the call is connected when user B calls number X.
	//
	// - BCallX_Gin: the audio prompt code played to the callee after the call is connected when number X is called.
	//
	// - BCallX_Bin: the audio prompt code played to user B after the call is connected when user B calls number X.
	//
	// example:
	//
	// {
	//
	// "GCallX_Gpre":"185"
	//
	// }
	Audio *string `json:"Audio,omitempty" xml:"Audio,omitempty"`
	// The one-way call restriction logic. Valid values:
	//
	// - CONTROL_BX_DISABLE: number B cannot call number X.
	//
	// - CONTROL_GX_DISABLE: numbers in number group G cannot call number X.
	//
	// example:
	//
	// CONTROL_AX_DISABLE
	CallRestrict *string `json:"CallRestrict,omitempty" xml:"CallRestrict,omitempty"`
	// The sequential ringing duration. Unit: seconds. Valid values: 5 to 20.
	//
	// example:
	//
	// 10
	CallTimeout *int64 `json:"CallTimeout,omitempty" xml:"CallTimeout,omitempty"`
	// Phone number A.
	//
	// The DefaultA number for the GXB private number.
	//
	// > Note
	//
	// The callback mode can be configured at the number pool level. When the callback mode is set to default, this field takes effect. When user B calls number X, the call is forwarded to this A number by default. If number A is not configured, the call cannot be connected.
	//
	// To update number A, call the number binding modification operation.
	//
	// example:
	//
	// 138****0000
	DefaultA *string `json:"DefaultA,omitempty" xml:"DefaultA,omitempty"`
	// The DTMF key configuration, used when setting up AXB intermediate number binding. The following items can be configured:
	//
	// - endCallIvrPhoneNo: the party to play the audio prompt to, A or B.
	//
	// - waitingDtmfTime: the maximum wait time after the first audio prompt is played. Maximum value: 30 seconds.
	//
	// - maxLoop dtmf: the maximum number of times the first audio file is played in a loop when the key input does not match. Maximum value: 5.
	//
	// - step1File: the name of the first audio file.
	//
	// - step2File: the name of the second audio file.
	//
	// - validKey: the valid key values, such as 1,2. Only two valid key values can be configured, separated by ",".
	//
	// - waitingEndCall: the carrier hang-up wait time. Maximum value: 10 seconds.
	//
	// example:
	//
	// {     "endCallIvrPhoneNo":"A",     "waitingDtmfTime":10,     "maxLoop":3,     "step1File":"62ab72f8-4750-4234-859e-e8d678c0cad3-flow_tts_test_1.wav",     "step2File":"62ab72f8-4750-4234-859e-e8d678c0cad3-flow_tts_test_2.wav",     "validKey":"1,2",     "waitingEndCall":2 }
	DtmfConfig *string `json:"DtmfConfig,omitempty" xml:"DtmfConfig,omitempty"`
	// The expiration time of the number binding.
	//
	// > Note
	//
	// The expiration time must be at least 1 minute later than the current binding time, accurate to the second.
	//
	// You can manually unbind the number at any time before the binding expires.
	//
	// If you unbind a number that is in an active call, the current call is not affected.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2021-09-05 12:00:00
	Expiration *string `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	// The ID of number group G.
	//
	// Number group G is created by using the "Create G Number Group" message.
	//
	// example:
	//
	// 10000*****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// - The Industrial Internet identifier code, which has a one-to-one mapping with OrderID.
	//
	// - The code supports characters such as "." and "/" as domain separators. The code must not contain non-printable characters such as spaces or line breaks.
	//
	// - This field is required when the business user generates the Industrial Internet identifier code independently.
	//
	// example:
	//
	// 700.100.1/12345678
	IndustrialId *string `json:"IndustrialId,omitempty" xml:"IndustrialId,omitempty"`
	// Specifies whether to enable the ASR service. Valid values:
	//
	// - false: Disabled (default value).
	//
	// - true: Enabled.
	//
	// example:
	//
	// true
	NeedAsr *bool `json:"NeedAsr,omitempty" xml:"NeedAsr,omitempty"`
	// Specifies whether the business user requires all call media files generated by this number binding relationship. Valid values:
	//
	// - true: Enable recording.
	//
	// - false: Disable recording (default value).
	//
	// example:
	//
	// false
	NeedRecord *bool `json:"NeedRecord,omitempty" xml:"NeedRecord,omitempty"`
	// The business order code.
	//
	// - This code has a one-to-one mapping with the actual business order and must be unique.
	//
	// - The code supports characters such as "." and "/" as domain separators. The code must not contain non-printable characters such as spaces or line breaks.
	//
	// - This field is required when the business user submits a business order code and the platform generates an Industrial Internet identifier code.
	//
	// example:
	//
	// 12345678
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// An external business extension field. This parameter is returned in the call record receipt message.
	//
	// example:
	//
	// abcdef
	OutId *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	// The external business ID.
	//
	// example:
	//
	// abcdefgh
	OutOrderId *string `json:"OutOrderId,omitempty" xml:"OutOrderId,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The key of the China 700 number pool.
	//
	// This parameter is required.
	//
	// example:
	//
	// FC10000016848****
	PoolKey *string `json:"PoolKey,omitempty" xml:"PoolKey,omitempty"`
	// The recording type. Valid values: mp3 and wav.
	//
	// example:
	//
	// mp3
	RecType              *string `json:"RecType,omitempty" xml:"RecType,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Phone number of user B.
	//
	// This parameter is required.
	//
	// example:
	//
	// 136****0000
	TelB *string `json:"TelB,omitempty" xml:"TelB,omitempty"`
	// The intermediate number X. If the business user does not specify an X number, the platform assigns a number from the number pool as the X number.
	//
	// example:
	//
	// 700********0000
	TelX *string `json:"TelX,omitempty" xml:"TelX,omitempty"`
}

func (s BindGxb700Request) String() string {
	return dara.Prettify(s)
}

func (s BindGxb700Request) GoString() string {
	return s.String()
}

func (s *BindGxb700Request) GetAsrModelId() *string {
	return s.AsrModelId
}

func (s *BindGxb700Request) GetAudio() *string {
	return s.Audio
}

func (s *BindGxb700Request) GetCallRestrict() *string {
	return s.CallRestrict
}

func (s *BindGxb700Request) GetCallTimeout() *int64 {
	return s.CallTimeout
}

func (s *BindGxb700Request) GetDefaultA() *string {
	return s.DefaultA
}

func (s *BindGxb700Request) GetDtmfConfig() *string {
	return s.DtmfConfig
}

func (s *BindGxb700Request) GetExpiration() *string {
	return s.Expiration
}

func (s *BindGxb700Request) GetGroupId() *string {
	return s.GroupId
}

func (s *BindGxb700Request) GetIndustrialId() *string {
	return s.IndustrialId
}

func (s *BindGxb700Request) GetNeedAsr() *bool {
	return s.NeedAsr
}

func (s *BindGxb700Request) GetNeedRecord() *bool {
	return s.NeedRecord
}

func (s *BindGxb700Request) GetOrderId() *string {
	return s.OrderId
}

func (s *BindGxb700Request) GetOutId() *string {
	return s.OutId
}

func (s *BindGxb700Request) GetOutOrderId() *string {
	return s.OutOrderId
}

func (s *BindGxb700Request) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *BindGxb700Request) GetPoolKey() *string {
	return s.PoolKey
}

func (s *BindGxb700Request) GetRecType() *string {
	return s.RecType
}

func (s *BindGxb700Request) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *BindGxb700Request) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *BindGxb700Request) GetTelB() *string {
	return s.TelB
}

func (s *BindGxb700Request) GetTelX() *string {
	return s.TelX
}

func (s *BindGxb700Request) SetAsrModelId(v string) *BindGxb700Request {
	s.AsrModelId = &v
	return s
}

func (s *BindGxb700Request) SetAudio(v string) *BindGxb700Request {
	s.Audio = &v
	return s
}

func (s *BindGxb700Request) SetCallRestrict(v string) *BindGxb700Request {
	s.CallRestrict = &v
	return s
}

func (s *BindGxb700Request) SetCallTimeout(v int64) *BindGxb700Request {
	s.CallTimeout = &v
	return s
}

func (s *BindGxb700Request) SetDefaultA(v string) *BindGxb700Request {
	s.DefaultA = &v
	return s
}

func (s *BindGxb700Request) SetDtmfConfig(v string) *BindGxb700Request {
	s.DtmfConfig = &v
	return s
}

func (s *BindGxb700Request) SetExpiration(v string) *BindGxb700Request {
	s.Expiration = &v
	return s
}

func (s *BindGxb700Request) SetGroupId(v string) *BindGxb700Request {
	s.GroupId = &v
	return s
}

func (s *BindGxb700Request) SetIndustrialId(v string) *BindGxb700Request {
	s.IndustrialId = &v
	return s
}

func (s *BindGxb700Request) SetNeedAsr(v bool) *BindGxb700Request {
	s.NeedAsr = &v
	return s
}

func (s *BindGxb700Request) SetNeedRecord(v bool) *BindGxb700Request {
	s.NeedRecord = &v
	return s
}

func (s *BindGxb700Request) SetOrderId(v string) *BindGxb700Request {
	s.OrderId = &v
	return s
}

func (s *BindGxb700Request) SetOutId(v string) *BindGxb700Request {
	s.OutId = &v
	return s
}

func (s *BindGxb700Request) SetOutOrderId(v string) *BindGxb700Request {
	s.OutOrderId = &v
	return s
}

func (s *BindGxb700Request) SetOwnerId(v int64) *BindGxb700Request {
	s.OwnerId = &v
	return s
}

func (s *BindGxb700Request) SetPoolKey(v string) *BindGxb700Request {
	s.PoolKey = &v
	return s
}

func (s *BindGxb700Request) SetRecType(v string) *BindGxb700Request {
	s.RecType = &v
	return s
}

func (s *BindGxb700Request) SetResourceOwnerAccount(v string) *BindGxb700Request {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *BindGxb700Request) SetResourceOwnerId(v int64) *BindGxb700Request {
	s.ResourceOwnerId = &v
	return s
}

func (s *BindGxb700Request) SetTelB(v string) *BindGxb700Request {
	s.TelB = &v
	return s
}

func (s *BindGxb700Request) SetTelX(v string) *BindGxb700Request {
	s.TelX = &v
	return s
}

func (s *BindGxb700Request) Validate() error {
	return dara.Validate(s)
}
