// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindAxb700Request interface {
	dara.Model
	String() string
	GoString() string
	SetAsrModelId(v string) *BindAxb700Request
	GetAsrModelId() *string
	SetAudio(v string) *BindAxb700Request
	GetAudio() *string
	SetCallRestrict(v string) *BindAxb700Request
	GetCallRestrict() *string
	SetCallTimeout(v int64) *BindAxb700Request
	GetCallTimeout() *int64
	SetDtmfConfig(v string) *BindAxb700Request
	GetDtmfConfig() *string
	SetExpiration(v string) *BindAxb700Request
	GetExpiration() *string
	SetIndustrialId(v string) *BindAxb700Request
	GetIndustrialId() *string
	SetNeedAsr(v bool) *BindAxb700Request
	GetNeedAsr() *bool
	SetNeedRecord(v bool) *BindAxb700Request
	GetNeedRecord() *bool
	SetOrderId(v string) *BindAxb700Request
	GetOrderId() *string
	SetOutId(v string) *BindAxb700Request
	GetOutId() *string
	SetOutOrderId(v string) *BindAxb700Request
	GetOutOrderId() *string
	SetOwnerId(v int64) *BindAxb700Request
	GetOwnerId() *int64
	SetPoolKey(v string) *BindAxb700Request
	GetPoolKey() *string
	SetRecType(v string) *BindAxb700Request
	GetRecType() *string
	SetResourceOwnerAccount(v string) *BindAxb700Request
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *BindAxb700Request
	GetResourceOwnerId() *int64
	SetTelA(v string) *BindAxb700Request
	GetTelA() *string
	SetTelB(v string) *BindAxb700Request
	GetTelB() *string
	SetTelX(v string) *BindAxb700Request
	GetTelX() *string
}

type BindAxb700Request struct {
	// example:
	//
	// 7ee372834d2f4cc7ac0d0ab2d0ae1aac
	AsrModelId *string `json:"AsrModelId,omitempty" xml:"AsrModelId,omitempty"`
	// example:
	//
	// {
	//
	// "ACallX_Apre":"185"
	//
	// }
	Audio *string `json:"Audio,omitempty" xml:"Audio,omitempty"`
	// example:
	//
	// CONTROL_AX_DISABLE
	CallRestrict *string `json:"CallRestrict,omitempty" xml:"CallRestrict,omitempty"`
	// example:
	//
	// 10
	CallTimeout *int64 `json:"CallTimeout,omitempty" xml:"CallTimeout,omitempty"`
	// example:
	//
	// {
	//
	//     "endCallIvrPhoneNo":"A",
	//
	//     "waitingDtmfTime":10,
	//
	//     "maxLoop":3,
	//
	//     "step1File":"62ab72f8-4750-4234-859e-e8d678c0cad3-flow_tts_test_1.wav",
	//
	//     "step2File":"62ab72f8-4750-4234-859e-e8d678c0cad3-flow_tts_test_2.wav",
	//
	//     "validKey":"1,2",
	//
	//     "waitingEndCall":2
	//
	// }
	DtmfConfig *string `json:"DtmfConfig,omitempty" xml:"DtmfConfig,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2021-09-05 12:00:00
	Expiration *string `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	// example:
	//
	// 700.100.1/12345678
	IndustrialId *string `json:"IndustrialId,omitempty" xml:"IndustrialId,omitempty"`
	// example:
	//
	// false
	NeedAsr *bool `json:"NeedAsr,omitempty" xml:"NeedAsr,omitempty"`
	// example:
	//
	// true
	NeedRecord *bool `json:"NeedRecord,omitempty" xml:"NeedRecord,omitempty"`
	// example:
	//
	// 12345678
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// abcdef
	OutId *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	// example:
	//
	// abcdefgh
	OutOrderId *string `json:"OutOrderId,omitempty" xml:"OutOrderId,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// FC2258****
	PoolKey *string `json:"PoolKey,omitempty" xml:"PoolKey,omitempty"`
	// example:
	//
	// mp3
	RecType              *string `json:"RecType,omitempty" xml:"RecType,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 139****0000
	TelA *string `json:"TelA,omitempty" xml:"TelA,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 136****0000
	TelB *string `json:"TelB,omitempty" xml:"TelB,omitempty"`
	// example:
	//
	// 700********0000
	TelX *string `json:"TelX,omitempty" xml:"TelX,omitempty"`
}

func (s BindAxb700Request) String() string {
	return dara.Prettify(s)
}

func (s BindAxb700Request) GoString() string {
	return s.String()
}

func (s *BindAxb700Request) GetAsrModelId() *string {
	return s.AsrModelId
}

func (s *BindAxb700Request) GetAudio() *string {
	return s.Audio
}

func (s *BindAxb700Request) GetCallRestrict() *string {
	return s.CallRestrict
}

func (s *BindAxb700Request) GetCallTimeout() *int64 {
	return s.CallTimeout
}

func (s *BindAxb700Request) GetDtmfConfig() *string {
	return s.DtmfConfig
}

func (s *BindAxb700Request) GetExpiration() *string {
	return s.Expiration
}

func (s *BindAxb700Request) GetIndustrialId() *string {
	return s.IndustrialId
}

func (s *BindAxb700Request) GetNeedAsr() *bool {
	return s.NeedAsr
}

func (s *BindAxb700Request) GetNeedRecord() *bool {
	return s.NeedRecord
}

func (s *BindAxb700Request) GetOrderId() *string {
	return s.OrderId
}

func (s *BindAxb700Request) GetOutId() *string {
	return s.OutId
}

func (s *BindAxb700Request) GetOutOrderId() *string {
	return s.OutOrderId
}

func (s *BindAxb700Request) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *BindAxb700Request) GetPoolKey() *string {
	return s.PoolKey
}

func (s *BindAxb700Request) GetRecType() *string {
	return s.RecType
}

func (s *BindAxb700Request) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *BindAxb700Request) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *BindAxb700Request) GetTelA() *string {
	return s.TelA
}

func (s *BindAxb700Request) GetTelB() *string {
	return s.TelB
}

func (s *BindAxb700Request) GetTelX() *string {
	return s.TelX
}

func (s *BindAxb700Request) SetAsrModelId(v string) *BindAxb700Request {
	s.AsrModelId = &v
	return s
}

func (s *BindAxb700Request) SetAudio(v string) *BindAxb700Request {
	s.Audio = &v
	return s
}

func (s *BindAxb700Request) SetCallRestrict(v string) *BindAxb700Request {
	s.CallRestrict = &v
	return s
}

func (s *BindAxb700Request) SetCallTimeout(v int64) *BindAxb700Request {
	s.CallTimeout = &v
	return s
}

func (s *BindAxb700Request) SetDtmfConfig(v string) *BindAxb700Request {
	s.DtmfConfig = &v
	return s
}

func (s *BindAxb700Request) SetExpiration(v string) *BindAxb700Request {
	s.Expiration = &v
	return s
}

func (s *BindAxb700Request) SetIndustrialId(v string) *BindAxb700Request {
	s.IndustrialId = &v
	return s
}

func (s *BindAxb700Request) SetNeedAsr(v bool) *BindAxb700Request {
	s.NeedAsr = &v
	return s
}

func (s *BindAxb700Request) SetNeedRecord(v bool) *BindAxb700Request {
	s.NeedRecord = &v
	return s
}

func (s *BindAxb700Request) SetOrderId(v string) *BindAxb700Request {
	s.OrderId = &v
	return s
}

func (s *BindAxb700Request) SetOutId(v string) *BindAxb700Request {
	s.OutId = &v
	return s
}

func (s *BindAxb700Request) SetOutOrderId(v string) *BindAxb700Request {
	s.OutOrderId = &v
	return s
}

func (s *BindAxb700Request) SetOwnerId(v int64) *BindAxb700Request {
	s.OwnerId = &v
	return s
}

func (s *BindAxb700Request) SetPoolKey(v string) *BindAxb700Request {
	s.PoolKey = &v
	return s
}

func (s *BindAxb700Request) SetRecType(v string) *BindAxb700Request {
	s.RecType = &v
	return s
}

func (s *BindAxb700Request) SetResourceOwnerAccount(v string) *BindAxb700Request {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *BindAxb700Request) SetResourceOwnerId(v int64) *BindAxb700Request {
	s.ResourceOwnerId = &v
	return s
}

func (s *BindAxb700Request) SetTelA(v string) *BindAxb700Request {
	s.TelA = &v
	return s
}

func (s *BindAxb700Request) SetTelB(v string) *BindAxb700Request {
	s.TelB = &v
	return s
}

func (s *BindAxb700Request) SetTelX(v string) *BindAxb700Request {
	s.TelX = &v
	return s
}

func (s *BindAxb700Request) Validate() error {
	return dara.Validate(s)
}
