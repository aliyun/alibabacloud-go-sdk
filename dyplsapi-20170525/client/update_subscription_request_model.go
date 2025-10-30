// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSubscriptionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetASRModelId(v string) *UpdateSubscriptionRequest
	GetASRModelId() *string
	SetASRStatus(v bool) *UpdateSubscriptionRequest
	GetASRStatus() *bool
	SetCallDisplayType(v int32) *UpdateSubscriptionRequest
	GetCallDisplayType() *int32
	SetCallRestrict(v string) *UpdateSubscriptionRequest
	GetCallRestrict() *string
	SetExpiration(v string) *UpdateSubscriptionRequest
	GetExpiration() *string
	SetGroupId(v string) *UpdateSubscriptionRequest
	GetGroupId() *string
	SetIsRecordingEnabled(v bool) *UpdateSubscriptionRequest
	GetIsRecordingEnabled() *bool
	SetOperateType(v string) *UpdateSubscriptionRequest
	GetOperateType() *string
	SetOutId(v string) *UpdateSubscriptionRequest
	GetOutId() *string
	SetOwnerId(v int64) *UpdateSubscriptionRequest
	GetOwnerId() *int64
	SetPhoneNoA(v string) *UpdateSubscriptionRequest
	GetPhoneNoA() *string
	SetPhoneNoB(v string) *UpdateSubscriptionRequest
	GetPhoneNoB() *string
	SetPhoneNoX(v string) *UpdateSubscriptionRequest
	GetPhoneNoX() *string
	SetPoolKey(v string) *UpdateSubscriptionRequest
	GetPoolKey() *string
	SetProductType(v string) *UpdateSubscriptionRequest
	GetProductType() *string
	SetResourceOwnerAccount(v string) *UpdateSubscriptionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdateSubscriptionRequest
	GetResourceOwnerId() *int64
	SetRingConfig(v string) *UpdateSubscriptionRequest
	GetRingConfig() *string
	SetSubsId(v string) *UpdateSubscriptionRequest
	GetSubsId() *string
}

type UpdateSubscriptionRequest struct {
	// The ID of the ASR model.
	//
	// example:
	//
	// 980abddb908f48e8b987cb2cd303****
	ASRModelId *string `json:"ASRModelId,omitempty" xml:"ASRModelId,omitempty"`
	// Specifies whether to enable automatic speech recognition (ASR). Valid values:
	//
	// 	- **false*	- (default): disables ASR.
	//
	// 	- **true**: enables ASR.
	//
	// example:
	//
	// false
	ASRStatus *bool `json:"ASRStatus,omitempty" xml:"ASRStatus,omitempty"`
	// Re-sets the phone number display logic in the phone number binding. Fixed value: **1**, indicating that phone number X is displayed on both the calling phone and the called phone.
	//
	// >  Due to the regulatory restrictions imposed by carriers, the setting to display real phone numbers during calls does not take effect.
	//
	// example:
	//
	// 1
	CallDisplayType *int32 `json:"CallDisplayType,omitempty" xml:"CallDisplayType,omitempty"`
	// One-way call restrictions. Valid values:
	//
	// 	- **CONTROL_AX_DISABLE**: Phone number A cannot be used to call phone number X.
	//
	// 	- **CONTROL_BX_DISABLE**: Phone number B cannot be used to call phone number X.
	//
	// 	- **CONTROL_CLEAR_DISABLE**: The call restrictions are cleared.
	//
	// >  This parameter is required when **OperateType*	- is set to **updateCallRestrict**.
	//
	// example:
	//
	// CONTROL_BX_DISABLE
	CallRestrict *string `json:"CallRestrict,omitempty" xml:"CallRestrict,omitempty"`
	// Re-sets the expiration time of the phone number binding.
	//
	// >
	//
	// 	- This parameter is required when **OperateType*	- is set to **updateExpire**.
	//
	// 	- The expiration time must be more than 1 minute later than the time when you call this API operation.
	//
	// example:
	//
	// 2019-09-05 12:00:00
	Expiration *string `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	// The ID of number group G in the phone number binding.
	//
	// >  This parameter is required when **OperateType*	- is set to **updateAxgGroup**.
	//
	// example:
	//
	// 1234
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// Re-sets the recording status in the phone number binding.
	//
	// >  This parameter does not have a default value. If you do not specify this parameter, the value of the corresponding field is not updated.
	//
	// example:
	//
	// true
	IsRecordingEnabled *bool `json:"IsRecordingEnabled,omitempty" xml:"IsRecordingEnabled,omitempty"`
	// The operation to modify the phone number binding. Valid values:
	//
	// 	- **updateNoA**: modifies phone number A.
	//
	// 	- **updateNoB**: modifies phone number B.
	//
	// 	- **updateExpire**: modifies the validity period of the binding.
	//
	// 	- **updateAxgGroup**: modifies number group G.
	//
	// 	- **updateCallRestrict**: modifies one-way call restrictions.
	//
	// 	- **updateCallDisplayType**: updates the number display logic for calls.
	//
	// 	- **updateOutId**: updates the value of the OutId parameter.
	//
	// 	- **updateIsRecordingEnabled**: updates the status of the recording feature in the binding.
	//
	// This parameter is required.
	//
	// example:
	//
	// updateNoA
	OperateType *string `json:"OperateType,omitempty" xml:"OperateType,omitempty"`
	// Re-sets the value of the OutId parameter in the phone number binding.
	//
	// example:
	//
	// abcdef
	OutId   *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Phone number A in the phone number binding.
	//
	// >  This parameter is required when **OperateType*	- is set to **updateNoA**.
	//
	// example:
	//
	// 1390000****
	PhoneNoA *string `json:"PhoneNoA,omitempty" xml:"PhoneNoA,omitempty"`
	// Phone number B in the phone number binding.
	//
	// >  This parameter is required when **OperateType*	- is set to **updateNoB**.
	//
	// example:
	//
	// 1390000****
	PhoneNoB *string `json:"PhoneNoB,omitempty" xml:"PhoneNoB,omitempty"`
	// Phone number X in the phone number binding.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1390000****
	PhoneNoX *string `json:"PhoneNoX,omitempty" xml:"PhoneNoX,omitempty"`
	// The key of the phone number pool.
	//
	// Log on to the [Phone Number Protection console](https://dypls.console.aliyun.com/dypls.htm#/account) and view the key of the phone number pool on the **Number Pool Management*	- page.
	//
	// >  This parameter is required when **ProductType*	- is left empty.
	//
	// example:
	//
	// FC122356****
	PoolKey *string `json:"PoolKey,omitempty" xml:"PoolKey,omitempty"`
	// The product type. Valid values:
	//
	// 	- **AXB_170**
	//
	// 	- **AXN_170**
	//
	// 	- **AXN_95**
	//
	// 	- **AXN_EXTENSION_REUSE**
	//
	// >
	//
	// 	- This parameter is applicable to the original key accounts of Alibaba Cloud. This parameter can be ignored for Alibaba Cloud users.
	//
	// 	- This parameter is required when **PoolKey*	- is left empty.
	//
	// example:
	//
	// AXB_170
	ProductType          *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Updates the ringtone code for enterprise Color Ring Back Tone (CRBT) in the phone number binding.
	//
	// AXB product:
	//
	// 	- Ringtone setting when phone number A is used to call phone number X in the AXB binding: AXBRing_A
	//
	// 	- Ringtone setting when phone number B is used to call phone number X in the AXB binding: AXBRing_B
	//
	// AXN product:
	//
	// 	- Ringtone setting (with a callback number) when phone number A is used to call phone number X in the AXN extension binding: AXNRing_AB
	//
	// 	- Ringtone setting (without a callback number) when phone number A is used to call phone number X in the AXN extension binding: AXNRing_A
	//
	// 	- Ringtone setting when phone number N is used to call phone number X in the AXN extension binding: AXNRing_N
	//
	// AXG product:
	//
	// 	- Ringtone setting (with a callback number) when phone number A is used to call phone number X in the AXG binding: AXGRing_AB
	//
	// 	- Ringtone setting (without a callback number) when phone number A is used to call phone number X in the AXG binding: AXGRing_A
	//
	// 	- Ringtone setting when a phone number in number group G is used to call phone number X in the AXG binding: AXGRing_G
	//
	// Enterprise CRBT codes: Enterprise CRBT codes can be queried in the [Phone Number Protection console](https://dypls.console.aliyun.com/dypls.htm#/account). You can choose **Number Pool Management*	- > **Enterprise CRBT Management*	- to view and manage enterprise CRBT codes. You can also upload, delete, or perform other operations on enterprise CRBT codes.
	//
	// >  The bound enterprise CRBTs are preferentially used. If no enterprise CRBT is set or the setting does not take effect, the enterprise CRBTs at the phone number pool level are used.
	//
	// example:
	//
	// {"AXBRing_B":"100000001","AXBRing_A":"100000001"}
	RingConfig *string `json:"RingConfig,omitempty" xml:"RingConfig,omitempty"`
	// The binding ID.
	//
	// Log on to the [Phone Number Protection console](https://dypls.console.aliyun.com/dypls.htm#/account), choose **Number and Number Pool*	- > **Number Management**. On the Number Management page, select the desired record and click Details to view the binding ID. Alternatively, you can view the value of the **SubsId*	- parameter returned by an API operation for a phone number binding such as BindAxb. The value of this parameter indicates a binding ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100000076879****
	SubsId *string `json:"SubsId,omitempty" xml:"SubsId,omitempty"`
}

func (s UpdateSubscriptionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSubscriptionRequest) GoString() string {
	return s.String()
}

func (s *UpdateSubscriptionRequest) GetASRModelId() *string {
	return s.ASRModelId
}

func (s *UpdateSubscriptionRequest) GetASRStatus() *bool {
	return s.ASRStatus
}

func (s *UpdateSubscriptionRequest) GetCallDisplayType() *int32 {
	return s.CallDisplayType
}

func (s *UpdateSubscriptionRequest) GetCallRestrict() *string {
	return s.CallRestrict
}

func (s *UpdateSubscriptionRequest) GetExpiration() *string {
	return s.Expiration
}

func (s *UpdateSubscriptionRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *UpdateSubscriptionRequest) GetIsRecordingEnabled() *bool {
	return s.IsRecordingEnabled
}

func (s *UpdateSubscriptionRequest) GetOperateType() *string {
	return s.OperateType
}

func (s *UpdateSubscriptionRequest) GetOutId() *string {
	return s.OutId
}

func (s *UpdateSubscriptionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateSubscriptionRequest) GetPhoneNoA() *string {
	return s.PhoneNoA
}

func (s *UpdateSubscriptionRequest) GetPhoneNoB() *string {
	return s.PhoneNoB
}

func (s *UpdateSubscriptionRequest) GetPhoneNoX() *string {
	return s.PhoneNoX
}

func (s *UpdateSubscriptionRequest) GetPoolKey() *string {
	return s.PoolKey
}

func (s *UpdateSubscriptionRequest) GetProductType() *string {
	return s.ProductType
}

func (s *UpdateSubscriptionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateSubscriptionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdateSubscriptionRequest) GetRingConfig() *string {
	return s.RingConfig
}

func (s *UpdateSubscriptionRequest) GetSubsId() *string {
	return s.SubsId
}

func (s *UpdateSubscriptionRequest) SetASRModelId(v string) *UpdateSubscriptionRequest {
	s.ASRModelId = &v
	return s
}

func (s *UpdateSubscriptionRequest) SetASRStatus(v bool) *UpdateSubscriptionRequest {
	s.ASRStatus = &v
	return s
}

func (s *UpdateSubscriptionRequest) SetCallDisplayType(v int32) *UpdateSubscriptionRequest {
	s.CallDisplayType = &v
	return s
}

func (s *UpdateSubscriptionRequest) SetCallRestrict(v string) *UpdateSubscriptionRequest {
	s.CallRestrict = &v
	return s
}

func (s *UpdateSubscriptionRequest) SetExpiration(v string) *UpdateSubscriptionRequest {
	s.Expiration = &v
	return s
}

func (s *UpdateSubscriptionRequest) SetGroupId(v string) *UpdateSubscriptionRequest {
	s.GroupId = &v
	return s
}

func (s *UpdateSubscriptionRequest) SetIsRecordingEnabled(v bool) *UpdateSubscriptionRequest {
	s.IsRecordingEnabled = &v
	return s
}

func (s *UpdateSubscriptionRequest) SetOperateType(v string) *UpdateSubscriptionRequest {
	s.OperateType = &v
	return s
}

func (s *UpdateSubscriptionRequest) SetOutId(v string) *UpdateSubscriptionRequest {
	s.OutId = &v
	return s
}

func (s *UpdateSubscriptionRequest) SetOwnerId(v int64) *UpdateSubscriptionRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateSubscriptionRequest) SetPhoneNoA(v string) *UpdateSubscriptionRequest {
	s.PhoneNoA = &v
	return s
}

func (s *UpdateSubscriptionRequest) SetPhoneNoB(v string) *UpdateSubscriptionRequest {
	s.PhoneNoB = &v
	return s
}

func (s *UpdateSubscriptionRequest) SetPhoneNoX(v string) *UpdateSubscriptionRequest {
	s.PhoneNoX = &v
	return s
}

func (s *UpdateSubscriptionRequest) SetPoolKey(v string) *UpdateSubscriptionRequest {
	s.PoolKey = &v
	return s
}

func (s *UpdateSubscriptionRequest) SetProductType(v string) *UpdateSubscriptionRequest {
	s.ProductType = &v
	return s
}

func (s *UpdateSubscriptionRequest) SetResourceOwnerAccount(v string) *UpdateSubscriptionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateSubscriptionRequest) SetResourceOwnerId(v int64) *UpdateSubscriptionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateSubscriptionRequest) SetRingConfig(v string) *UpdateSubscriptionRequest {
	s.RingConfig = &v
	return s
}

func (s *UpdateSubscriptionRequest) SetSubsId(v string) *UpdateSubscriptionRequest {
	s.SubsId = &v
	return s
}

func (s *UpdateSubscriptionRequest) Validate() error {
	return dara.Validate(s)
}
