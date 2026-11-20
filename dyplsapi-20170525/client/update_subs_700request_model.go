// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSubs700Request interface {
	dara.Model
	String() string
	GoString() string
	SetAsrModelId(v string) *UpdateSubs700Request
	GetAsrModelId() *string
	SetAudio(v string) *UpdateSubs700Request
	GetAudio() *string
	SetCallRestrict(v string) *UpdateSubs700Request
	GetCallRestrict() *string
	SetDefaultA(v string) *UpdateSubs700Request
	GetDefaultA() *string
	SetExpiration(v string) *UpdateSubs700Request
	GetExpiration() *string
	SetGroupId(v string) *UpdateSubs700Request
	GetGroupId() *string
	SetIndustrialId(v string) *UpdateSubs700Request
	GetIndustrialId() *string
	SetNeedAsr(v bool) *UpdateSubs700Request
	GetNeedAsr() *bool
	SetNeedRecord(v bool) *UpdateSubs700Request
	GetNeedRecord() *bool
	SetOperateType(v string) *UpdateSubs700Request
	GetOperateType() *string
	SetOrderId(v string) *UpdateSubs700Request
	GetOrderId() *string
	SetOutId(v string) *UpdateSubs700Request
	GetOutId() *string
	SetOwnerId(v int64) *UpdateSubs700Request
	GetOwnerId() *int64
	SetPoolKey(v string) *UpdateSubs700Request
	GetPoolKey() *string
	SetResourceOwnerAccount(v string) *UpdateSubs700Request
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdateSubs700Request
	GetResourceOwnerId() *int64
	SetSubsId(v int64) *UpdateSubs700Request
	GetSubsId() *int64
	SetTelA(v string) *UpdateSubs700Request
	GetTelA() *string
	SetTelB(v string) *UpdateSubs700Request
	GetTelB() *string
	SetTelX(v string) *UpdateSubs700Request
	GetTelX() *string
}

type UpdateSubs700Request struct {
	// The ASR model ID. Valid values:
	//
	// - General-telephone-Mandarin Chinese model ID: 7ee372834d2f4cc7ac0d0ab2d0ae1aac.
	//
	// - General-telephone-Cantonese model ID: 9b4afe3d224e4fa3bf963697113e2f77.
	//
	// - General-telephone-regional accent model ID: 3655a6cb94764145a1f9a38a3c3b3b4b.
	//
	// example:
	//
	// 7ee372834d2f4cc7ac0d0ab2d0ae1aac
	AsrModelId *string `json:"AsrModelId,omitempty" xml:"AsrModelId,omitempty"`
	// Optional. The audio playback code.
	//
	// - AXB mode:
	//
	// ACallX_Apre: The audio playback code for the voice prompt played to Number A before the call is connected when Number A calls Number X.
	//
	// ACallX_Ain: The audio playback code for the voice prompt played to Number A after the call is connected when Number A calls Number X.
	//
	// ACallX_Bin: The audio playback code for the voice prompt played to Number B after the call is connected when Number A calls Number X.
	//
	// BCallX_Bpre: The audio playback code for the voice prompt played to Number B before the call is connected when Number B calls Number X.
	//
	// BCallX_Ain: The audio playback code for the voice prompt played to Number A after the call is connected when Number B calls Number X.
	//
	// BCallX_Bin: The audio playback code for the voice prompt played to Number B after the call is connected when Number B calls Number X.
	//
	// - GXB mode:
	//
	// GCallX_Gpre: The audio playback code for the voice prompt played to the caller before the call is connected when a number in Group G calls Number X.
	//
	// GCallX_Gin: The audio playback code for the voice prompt played to the caller after the call is connected when a number in Group G calls Number X.
	//
	// GCallX_Bin: The audio playback code for the voice prompt played to Number B after the call is connected when a number in Group G calls Number X.
	//
	// BCallX_Bpre: The audio playback code for the voice prompt played to Number B before the call is connected when Number B calls Number X.
	//
	// BCallX_Gin: The audio playback code for the voice prompt played to the callee after the call is connected when Number B calls Number X.
	//
	// BCallX_Bin: The audio playback code for the voice prompt played to Number B after the call is connected when Number B calls Number X.
	//
	// example:
	//
	// {
	//
	// "ACallX_Apre":"185"
	//
	// }
	Audio *string `json:"Audio,omitempty" xml:"Audio,omitempty"`
	// The new single-call restriction logic.
	//
	// example:
	//
	// CONTROL_AX_DISABLE
	CallRestrict *string `json:"CallRestrict,omitempty" xml:"CallRestrict,omitempty"`
	// The default Number A for the GXB private number. Note: The callback mode can be set at the number pool level. When the callback mode is set to the default mode, this field takes effect. When Number B calls Number X, the call is forwarded to this Number A by default. If Number A is not set, the call cannot be connected. To update Number A, call the number binding modification API.
	//
	// example:
	//
	// 138****0000
	DefaultA *string `json:"DefaultA,omitempty" xml:"DefaultA,omitempty"`
	// The new expiration time of the binding relationship.
	//
	// example:
	//
	// 2021-09-05 12:00:00
	Expiration *string `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	// The new group ID in the binding relationship.
	//
	// example:
	//
	// Sample value
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// Optional. The Industrial Internet identification code, which has a one-to-one mapping with the OrderId. The code supports characters such as periods (.) and forward slashes (/) to separate code domains. The code must not contain non-printable characters, such as spaces and line breaks. This field is required when the business user generates the Industrial Internet identification code.
	//
	// example:
	//
	// 700.100.1/12345678
	IndustrialId *string `json:"IndustrialId,omitempty" xml:"IndustrialId,omitempty"`
	// Specifies whether to enable the ASR feature. Valid values:
	//
	// - false: Disables the feature. This is the default value.
	//
	// - true: Enables the feature.
	//
	// example:
	//
	// false
	NeedAsr *bool `json:"NeedAsr,omitempty" xml:"NeedAsr,omitempty"`
	// The new media selection in the binding relationship.
	//
	// example:
	//
	// true
	NeedRecord *bool `json:"NeedRecord,omitempty" xml:"NeedRecord,omitempty"`
	// The operation type. Valid values:
	//
	// - updateNoA: Updates Number A.
	//
	// - updateNoB: Updates Number B.
	//
	// - updateExpire: Updates the validity period of the binding relationship.
	//
	// - updateGroup: Updates the Number G group.
	//
	// - updateNoDefaultA: Updates the default callback Number A in GXB mode.
	//
	// - updateCallRestrict: Sets the single-call restriction.
	//
	// - updateOutId: Updates the OutId field.
	//
	// - updateIsRecordingEnabled: Updates the recording status in the binding.
	//
	// - updateAsr: Updates the ASR status in the binding.
	//
	// - updateAudio: Updates the ringtone configuration in the binding.
	//
	// This parameter is required.
	//
	// example:
	//
	// updateNoA
	OperateType *string `json:"OperateType,omitempty" xml:"OperateType,omitempty"`
	// The order ID. This parameter is required in practice. Encrypt the order ID for reporting to the regulatory platform.
	//
	// example:
	//
	// 12345678
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The external business extension field. This parameter is returned in the call record receipt message.
	//
	// example:
	//
	// abcdef
	OutId   *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The key of the number pool.
	//
	// This parameter is required.
	//
	// example:
	//
	// FC10000016848****
	PoolKey              *string `json:"PoolKey,omitempty" xml:"PoolKey,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the binding relationship, which uniquely identifies the number binding relationship.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000*****
	SubsId *int64 `json:"SubsId,omitempty" xml:"SubsId,omitempty"`
	// The new Number A in the binding relationship.
	//
	// example:
	//
	// 138****0000
	TelA *string `json:"TelA,omitempty" xml:"TelA,omitempty"`
	// The new Number B in the binding relationship.
	//
	// example:
	//
	// 136****0000
	TelB *string `json:"TelB,omitempty" xml:"TelB,omitempty"`
	// The intermediate Number X. If the business user does not specify Number X, the application platform assigns a number from the number pool as Number X.
	//
	// This parameter is required.
	//
	// example:
	//
	// 700********0000
	TelX *string `json:"TelX,omitempty" xml:"TelX,omitempty"`
}

func (s UpdateSubs700Request) String() string {
	return dara.Prettify(s)
}

func (s UpdateSubs700Request) GoString() string {
	return s.String()
}

func (s *UpdateSubs700Request) GetAsrModelId() *string {
	return s.AsrModelId
}

func (s *UpdateSubs700Request) GetAudio() *string {
	return s.Audio
}

func (s *UpdateSubs700Request) GetCallRestrict() *string {
	return s.CallRestrict
}

func (s *UpdateSubs700Request) GetDefaultA() *string {
	return s.DefaultA
}

func (s *UpdateSubs700Request) GetExpiration() *string {
	return s.Expiration
}

func (s *UpdateSubs700Request) GetGroupId() *string {
	return s.GroupId
}

func (s *UpdateSubs700Request) GetIndustrialId() *string {
	return s.IndustrialId
}

func (s *UpdateSubs700Request) GetNeedAsr() *bool {
	return s.NeedAsr
}

func (s *UpdateSubs700Request) GetNeedRecord() *bool {
	return s.NeedRecord
}

func (s *UpdateSubs700Request) GetOperateType() *string {
	return s.OperateType
}

func (s *UpdateSubs700Request) GetOrderId() *string {
	return s.OrderId
}

func (s *UpdateSubs700Request) GetOutId() *string {
	return s.OutId
}

func (s *UpdateSubs700Request) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateSubs700Request) GetPoolKey() *string {
	return s.PoolKey
}

func (s *UpdateSubs700Request) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateSubs700Request) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdateSubs700Request) GetSubsId() *int64 {
	return s.SubsId
}

func (s *UpdateSubs700Request) GetTelA() *string {
	return s.TelA
}

func (s *UpdateSubs700Request) GetTelB() *string {
	return s.TelB
}

func (s *UpdateSubs700Request) GetTelX() *string {
	return s.TelX
}

func (s *UpdateSubs700Request) SetAsrModelId(v string) *UpdateSubs700Request {
	s.AsrModelId = &v
	return s
}

func (s *UpdateSubs700Request) SetAudio(v string) *UpdateSubs700Request {
	s.Audio = &v
	return s
}

func (s *UpdateSubs700Request) SetCallRestrict(v string) *UpdateSubs700Request {
	s.CallRestrict = &v
	return s
}

func (s *UpdateSubs700Request) SetDefaultA(v string) *UpdateSubs700Request {
	s.DefaultA = &v
	return s
}

func (s *UpdateSubs700Request) SetExpiration(v string) *UpdateSubs700Request {
	s.Expiration = &v
	return s
}

func (s *UpdateSubs700Request) SetGroupId(v string) *UpdateSubs700Request {
	s.GroupId = &v
	return s
}

func (s *UpdateSubs700Request) SetIndustrialId(v string) *UpdateSubs700Request {
	s.IndustrialId = &v
	return s
}

func (s *UpdateSubs700Request) SetNeedAsr(v bool) *UpdateSubs700Request {
	s.NeedAsr = &v
	return s
}

func (s *UpdateSubs700Request) SetNeedRecord(v bool) *UpdateSubs700Request {
	s.NeedRecord = &v
	return s
}

func (s *UpdateSubs700Request) SetOperateType(v string) *UpdateSubs700Request {
	s.OperateType = &v
	return s
}

func (s *UpdateSubs700Request) SetOrderId(v string) *UpdateSubs700Request {
	s.OrderId = &v
	return s
}

func (s *UpdateSubs700Request) SetOutId(v string) *UpdateSubs700Request {
	s.OutId = &v
	return s
}

func (s *UpdateSubs700Request) SetOwnerId(v int64) *UpdateSubs700Request {
	s.OwnerId = &v
	return s
}

func (s *UpdateSubs700Request) SetPoolKey(v string) *UpdateSubs700Request {
	s.PoolKey = &v
	return s
}

func (s *UpdateSubs700Request) SetResourceOwnerAccount(v string) *UpdateSubs700Request {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateSubs700Request) SetResourceOwnerId(v int64) *UpdateSubs700Request {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateSubs700Request) SetSubsId(v int64) *UpdateSubs700Request {
	s.SubsId = &v
	return s
}

func (s *UpdateSubs700Request) SetTelA(v string) *UpdateSubs700Request {
	s.TelA = &v
	return s
}

func (s *UpdateSubs700Request) SetTelB(v string) *UpdateSubs700Request {
	s.TelB = &v
	return s
}

func (s *UpdateSubs700Request) SetTelX(v string) *UpdateSubs700Request {
	s.TelX = &v
	return s
}

func (s *UpdateSubs700Request) Validate() error {
	return dara.Validate(s)
}
