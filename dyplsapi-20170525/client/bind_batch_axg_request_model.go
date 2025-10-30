// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindBatchAxgRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAxgBindList(v []*BindBatchAxgRequestAxgBindList) *BindBatchAxgRequest
	GetAxgBindList() []*BindBatchAxgRequestAxgBindList
	SetOwnerId(v int64) *BindBatchAxgRequest
	GetOwnerId() *int64
	SetPoolKey(v string) *BindBatchAxgRequest
	GetPoolKey() *string
	SetResourceOwnerAccount(v string) *BindBatchAxgRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *BindBatchAxgRequest
	GetResourceOwnerId() *int64
}

type BindBatchAxgRequest struct {
	// This parameter is required.
	AxgBindList []*BindBatchAxgRequestAxgBindList `json:"AxgBindList,omitempty" xml:"AxgBindList,omitempty" type:"Repeated"`
	OwnerId     *int64                            `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// FC2235****
	PoolKey              *string `json:"PoolKey,omitempty" xml:"PoolKey,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s BindBatchAxgRequest) String() string {
	return dara.Prettify(s)
}

func (s BindBatchAxgRequest) GoString() string {
	return s.String()
}

func (s *BindBatchAxgRequest) GetAxgBindList() []*BindBatchAxgRequestAxgBindList {
	return s.AxgBindList
}

func (s *BindBatchAxgRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *BindBatchAxgRequest) GetPoolKey() *string {
	return s.PoolKey
}

func (s *BindBatchAxgRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *BindBatchAxgRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *BindBatchAxgRequest) SetAxgBindList(v []*BindBatchAxgRequestAxgBindList) *BindBatchAxgRequest {
	s.AxgBindList = v
	return s
}

func (s *BindBatchAxgRequest) SetOwnerId(v int64) *BindBatchAxgRequest {
	s.OwnerId = &v
	return s
}

func (s *BindBatchAxgRequest) SetPoolKey(v string) *BindBatchAxgRequest {
	s.PoolKey = &v
	return s
}

func (s *BindBatchAxgRequest) SetResourceOwnerAccount(v string) *BindBatchAxgRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *BindBatchAxgRequest) SetResourceOwnerId(v int64) *BindBatchAxgRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *BindBatchAxgRequest) Validate() error {
	if s.AxgBindList != nil {
		for _, item := range s.AxgBindList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type BindBatchAxgRequestAxgBindList struct {
	// example:
	//
	// 7ee372834d2f4cc7ac0d0ab2d0ae1aac
	ASRModelId *string `json:"ASRModelId,omitempty" xml:"ASRModelId,omitempty"`
	// example:
	//
	// true
	ASRStatus *bool `json:"ASRStatus,omitempty" xml:"ASRStatus,omitempty"`
	// example:
	//
	// 1
	CallDisplayType *int32 `json:"CallDisplayType,omitempty" xml:"CallDisplayType,omitempty"`
	// example:
	//
	// CONTROL_AX_DISABLE
	CallRestrict *string `json:"CallRestrict,omitempty" xml:"CallRestrict,omitempty"`
	ExpectCity   *string `json:"ExpectCity,omitempty" xml:"ExpectCity,omitempty"`
	// example:
	//
	// 2022-07-11 01:05:15
	Expiration *string `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1234
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// example:
	//
	// False
	IsRecordingEnabled *bool `json:"IsRecordingEnabled,omitempty" xml:"IsRecordingEnabled,omitempty"`
	// example:
	//
	// 18223ad447910fd
	OutId *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	// example:
	//
	// 20220824021816883677
	OutOrderId *string `json:"OutOrderId,omitempty" xml:"OutOrderId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 13333333333
	PhoneNoA *string `json:"PhoneNoA,omitempty" xml:"PhoneNoA,omitempty"`
	// example:
	//
	// 13333333333
	PhoneNoB *string `json:"PhoneNoB,omitempty" xml:"PhoneNoB,omitempty"`
	// example:
	//
	// 13333333333
	PhoneNoX *string `json:"PhoneNoX,omitempty" xml:"PhoneNoX,omitempty"`
	// example:
	//
	// {\\"AXBRing_B\\":\\"100000002\\",\\"AXBRing_A\\":\\"100000001\\"}
	RingConfig *string `json:"RingConfig,omitempty" xml:"RingConfig,omitempty"`
}

func (s BindBatchAxgRequestAxgBindList) String() string {
	return dara.Prettify(s)
}

func (s BindBatchAxgRequestAxgBindList) GoString() string {
	return s.String()
}

func (s *BindBatchAxgRequestAxgBindList) GetASRModelId() *string {
	return s.ASRModelId
}

func (s *BindBatchAxgRequestAxgBindList) GetASRStatus() *bool {
	return s.ASRStatus
}

func (s *BindBatchAxgRequestAxgBindList) GetCallDisplayType() *int32 {
	return s.CallDisplayType
}

func (s *BindBatchAxgRequestAxgBindList) GetCallRestrict() *string {
	return s.CallRestrict
}

func (s *BindBatchAxgRequestAxgBindList) GetExpectCity() *string {
	return s.ExpectCity
}

func (s *BindBatchAxgRequestAxgBindList) GetExpiration() *string {
	return s.Expiration
}

func (s *BindBatchAxgRequestAxgBindList) GetGroupId() *string {
	return s.GroupId
}

func (s *BindBatchAxgRequestAxgBindList) GetIsRecordingEnabled() *bool {
	return s.IsRecordingEnabled
}

func (s *BindBatchAxgRequestAxgBindList) GetOutId() *string {
	return s.OutId
}

func (s *BindBatchAxgRequestAxgBindList) GetOutOrderId() *string {
	return s.OutOrderId
}

func (s *BindBatchAxgRequestAxgBindList) GetPhoneNoA() *string {
	return s.PhoneNoA
}

func (s *BindBatchAxgRequestAxgBindList) GetPhoneNoB() *string {
	return s.PhoneNoB
}

func (s *BindBatchAxgRequestAxgBindList) GetPhoneNoX() *string {
	return s.PhoneNoX
}

func (s *BindBatchAxgRequestAxgBindList) GetRingConfig() *string {
	return s.RingConfig
}

func (s *BindBatchAxgRequestAxgBindList) SetASRModelId(v string) *BindBatchAxgRequestAxgBindList {
	s.ASRModelId = &v
	return s
}

func (s *BindBatchAxgRequestAxgBindList) SetASRStatus(v bool) *BindBatchAxgRequestAxgBindList {
	s.ASRStatus = &v
	return s
}

func (s *BindBatchAxgRequestAxgBindList) SetCallDisplayType(v int32) *BindBatchAxgRequestAxgBindList {
	s.CallDisplayType = &v
	return s
}

func (s *BindBatchAxgRequestAxgBindList) SetCallRestrict(v string) *BindBatchAxgRequestAxgBindList {
	s.CallRestrict = &v
	return s
}

func (s *BindBatchAxgRequestAxgBindList) SetExpectCity(v string) *BindBatchAxgRequestAxgBindList {
	s.ExpectCity = &v
	return s
}

func (s *BindBatchAxgRequestAxgBindList) SetExpiration(v string) *BindBatchAxgRequestAxgBindList {
	s.Expiration = &v
	return s
}

func (s *BindBatchAxgRequestAxgBindList) SetGroupId(v string) *BindBatchAxgRequestAxgBindList {
	s.GroupId = &v
	return s
}

func (s *BindBatchAxgRequestAxgBindList) SetIsRecordingEnabled(v bool) *BindBatchAxgRequestAxgBindList {
	s.IsRecordingEnabled = &v
	return s
}

func (s *BindBatchAxgRequestAxgBindList) SetOutId(v string) *BindBatchAxgRequestAxgBindList {
	s.OutId = &v
	return s
}

func (s *BindBatchAxgRequestAxgBindList) SetOutOrderId(v string) *BindBatchAxgRequestAxgBindList {
	s.OutOrderId = &v
	return s
}

func (s *BindBatchAxgRequestAxgBindList) SetPhoneNoA(v string) *BindBatchAxgRequestAxgBindList {
	s.PhoneNoA = &v
	return s
}

func (s *BindBatchAxgRequestAxgBindList) SetPhoneNoB(v string) *BindBatchAxgRequestAxgBindList {
	s.PhoneNoB = &v
	return s
}

func (s *BindBatchAxgRequestAxgBindList) SetPhoneNoX(v string) *BindBatchAxgRequestAxgBindList {
	s.PhoneNoX = &v
	return s
}

func (s *BindBatchAxgRequestAxgBindList) SetRingConfig(v string) *BindBatchAxgRequestAxgBindList {
	s.RingConfig = &v
	return s
}

func (s *BindBatchAxgRequestAxgBindList) Validate() error {
	return dara.Validate(s)
}
