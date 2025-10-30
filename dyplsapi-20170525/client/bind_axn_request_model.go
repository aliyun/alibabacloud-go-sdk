// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindAxnRequest interface {
	dara.Model
	String() string
	GoString() string
	SetASRModelId(v string) *BindAxnRequest
	GetASRModelId() *string
	SetASRStatus(v bool) *BindAxnRequest
	GetASRStatus() *bool
	SetCallDisplayType(v int32) *BindAxnRequest
	GetCallDisplayType() *int32
	SetCallRestrict(v string) *BindAxnRequest
	GetCallRestrict() *string
	SetCallTimeout(v int32) *BindAxnRequest
	GetCallTimeout() *int32
	SetExpectCity(v string) *BindAxnRequest
	GetExpectCity() *string
	SetExpiration(v string) *BindAxnRequest
	GetExpiration() *string
	SetExtend(v string) *BindAxnRequest
	GetExtend() *string
	SetIsRecordingEnabled(v bool) *BindAxnRequest
	GetIsRecordingEnabled() *bool
	SetNoType(v string) *BindAxnRequest
	GetNoType() *string
	SetOutId(v string) *BindAxnRequest
	GetOutId() *string
	SetOutOrderId(v string) *BindAxnRequest
	GetOutOrderId() *string
	SetOwnerId(v int64) *BindAxnRequest
	GetOwnerId() *int64
	SetPhoneNoA(v string) *BindAxnRequest
	GetPhoneNoA() *string
	SetPhoneNoB(v string) *BindAxnRequest
	GetPhoneNoB() *string
	SetPhoneNoX(v string) *BindAxnRequest
	GetPhoneNoX() *string
	SetPoolKey(v string) *BindAxnRequest
	GetPoolKey() *string
	SetResourceOwnerAccount(v string) *BindAxnRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *BindAxnRequest
	GetResourceOwnerId() *int64
	SetRingConfig(v string) *BindAxnRequest
	GetRingConfig() *string
}

type BindAxnRequest struct {
	// The ID of the ASR model. On the [Automatic Speech Recognition (ASR) Model Management](https://dyplsnext.console.aliyun.com/?spm=5176.12818093.categories-n-products.ddypls.22e616d0a0tEFC#/asr) page, you can view the ID of the ASR model.
	//
	// example:
	//
	// 7ee372834d2f4cc7ac0d0ab2d0ae1aac
	ASRModelId *string `json:"ASRModelId,omitempty" xml:"ASRModelId,omitempty"`
	// Specifies whether to enable automatic speech recognition (ASR). Valid values:
	//
	// 	- **false*	- (default): disables ASR.
	//
	// 	- **true**: enables ASR.
	//
	// example:
	//
	// true
	ASRStatus *bool `json:"ASRStatus,omitempty" xml:"ASRStatus,omitempty"`
	// Re-sets the phone number display logic in the AXN binding. Fixed value: **1**, indicating that phone number X is displayed on both the calling phone and the called phone.
	//
	// >  Due to the regulatory restrictions imposed by carriers, the setting to display real phone numbers during calls does not take effect.
	//
	// example:
	//
	// 1
	CallDisplayType *int32 `json:"CallDisplayType,omitempty" xml:"CallDisplayType,omitempty"`
	// The status of one-way call restrictions. Valid values:
	//
	// 	- **CONTROL_AX_DISABLE**: Phone number A cannot be used to call phone number X.
	//
	// 	- **CONTROL_BX_DISABLE**: Phone number B cannot be used to call phone number X.
	//
	// example:
	//
	// CONTROL_AX_DISABLE
	CallRestrict *string `json:"CallRestrict,omitempty" xml:"CallRestrict,omitempty"`
	// The maximum ringing duration for each number in sequential ringing. Unit: seconds.
	//
	// example:
	//
	// 10
	CallTimeout *int32 `json:"CallTimeout,omitempty" xml:"CallTimeout,omitempty"`
	// Specifies the city to which phone number X to be selected belongs.
	//
	// 	- If no phone number for the specified city is available in the current phone number pool or this parameter is not specified, a phone number that belongs to another city is randomly selected from the current phone number pool and assigned as phone number X.
	//
	// 	- If **Number X Assignment Mode*	- is set to **Strict Matching Mode*	- and no phone number meets the requirement, the system displays an allocation error.
	//
	// example:
	//
	// hangzhou
	ExpectCity *string `json:"ExpectCity,omitempty" xml:"ExpectCity,omitempty"`
	// The expiration time of the AXN binding. Unit: seconds.
	//
	// >  The expiration time must be more than 60 seconds later than the time when you call this API operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2021-09-05 12:00:00
	Expiration *string `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	Extend     *string `json:"Extend,omitempty" xml:"Extend,omitempty"`
	// Specifies whether to record all calls made by the bound phone numbers. Valid values:
	//
	// 	- **true**
	//
	// 	- **false*	- (default)
	//
	// example:
	//
	// true
	IsRecordingEnabled *bool `json:"IsRecordingEnabled,omitempty" xml:"IsRecordingEnabled,omitempty"`
	// The type of the phone number.
	//
	// >  This parameter is applicable to the key accounts of Alibaba Cloud. This parameter can be ignored for Alibaba Cloud users.
	//
	// example:
	//
	// AXB_170
	NoType *string `json:"NoType,omitempty" xml:"NoType,omitempty"`
	// The extension field for the external business. This parameter is returned in a call record receipt.
	//
	// example:
	//
	// abcdef
	OutId *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	// The ID of the external business.
	//
	// example:
	//
	// 34553330****
	OutOrderId *string `json:"OutOrderId,omitempty" xml:"OutOrderId,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Phone number A in the AXN binding. Phone number A can be set to a mobile phone number or a landline phone number. The landline phone number must be added with an area code, and no hyphen is required between the area code and the landline phone number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 139****0000
	PhoneNoA *string `json:"PhoneNoA,omitempty" xml:"PhoneNoA,omitempty"`
	// Phone number B in the AXN binding. If phone number A is used to call phone number X, the call is forwarded to phone number B. Phone number B can be set to a mobile phone number or a landline phone number. The landline phone number must be added with an area code, and no hyphen is required between the area code and the landline phone number.
	//
	// >  If phone number B is not specified in the AXN binding, the system automatically generates a nonexistent number. If phone number A is used to call phone number X, the nonexistent number is returned. If you need to update phone number B, call the [UpdateSubscription](https://help.aliyun.com/document_detail/110253.html) operation.
	//
	// example:
	//
	// 138****0000
	PhoneNoB *string `json:"PhoneNoB,omitempty" xml:"PhoneNoB,omitempty"`
	// Phone number X in the AXN binding. Phone number X is the phone number that you purchased in the [Phone Number Protection console](https://dypls.console.aliyun.com/dypls.htm#/account) or by using the [BuySecretNo](https://help.aliyun.com/document_detail/110266.html) operation before you bind a phone number. Phone number X is used to forward calls.
	//
	// >  If you do not specify this parameter, a random phone number is selected from the phone number pool based on the value of the ExpectCity parameter and is used as phone number X.
	//
	// example:
	//
	// 139****0000
	PhoneNoX *string `json:"PhoneNoX,omitempty" xml:"PhoneNoX,omitempty"`
	// The key of the phone number pool. Log on to the [Phone Number Protection console ](https://dypls.console.aliyun.com/dypls.htm#/account)and view the key of the phone number pool on the **Number Pool Management*	- page.
	//
	// example:
	//
	// FC2256****
	PoolKey              *string `json:"PoolKey,omitempty" xml:"PoolKey,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Sets the ringtone code for enterprise Color Ring Back Tone (CRBT) in the AXN extension binding.
	//
	// 	- Ringtone setting (with a callback number) when phone number A is used to call phone number X in the AXN extension binding: AXNRing_AB
	//
	// 	- Ringtone setting (without a callback number) when phone number A is used to call phone number X in the AXN extension binding: AXNRing_A
	//
	// 	- Ringtone setting when an N-side number is used to call phone number X in the AXN extension binding: AXNRing_N
	//
	// Enterprise CRBT codes: Enterprise CRBT codes can be queried in the [Phone Number Protection console](https://dypls.console.aliyun.com/dypls.htm#/account). You can choose **Number Pool Management > Enterprise CRBT Management*	- to view enterprise CRBT codes. You can also upload, delete, or perform other operations on enterprise CRBT codes.
	//
	// >  The bound enterprise CRBTs are preferentially used. If no enterprise CRBT is set or the setting does not take effect, the enterprise CRBTs at the phone number pool level are used.
	//
	// example:
	//
	// {"AXNRing_N":"100000001","AXNRing_A":"100000001"}
	RingConfig *string `json:"RingConfig,omitempty" xml:"RingConfig,omitempty"`
}

func (s BindAxnRequest) String() string {
	return dara.Prettify(s)
}

func (s BindAxnRequest) GoString() string {
	return s.String()
}

func (s *BindAxnRequest) GetASRModelId() *string {
	return s.ASRModelId
}

func (s *BindAxnRequest) GetASRStatus() *bool {
	return s.ASRStatus
}

func (s *BindAxnRequest) GetCallDisplayType() *int32 {
	return s.CallDisplayType
}

func (s *BindAxnRequest) GetCallRestrict() *string {
	return s.CallRestrict
}

func (s *BindAxnRequest) GetCallTimeout() *int32 {
	return s.CallTimeout
}

func (s *BindAxnRequest) GetExpectCity() *string {
	return s.ExpectCity
}

func (s *BindAxnRequest) GetExpiration() *string {
	return s.Expiration
}

func (s *BindAxnRequest) GetExtend() *string {
	return s.Extend
}

func (s *BindAxnRequest) GetIsRecordingEnabled() *bool {
	return s.IsRecordingEnabled
}

func (s *BindAxnRequest) GetNoType() *string {
	return s.NoType
}

func (s *BindAxnRequest) GetOutId() *string {
	return s.OutId
}

func (s *BindAxnRequest) GetOutOrderId() *string {
	return s.OutOrderId
}

func (s *BindAxnRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *BindAxnRequest) GetPhoneNoA() *string {
	return s.PhoneNoA
}

func (s *BindAxnRequest) GetPhoneNoB() *string {
	return s.PhoneNoB
}

func (s *BindAxnRequest) GetPhoneNoX() *string {
	return s.PhoneNoX
}

func (s *BindAxnRequest) GetPoolKey() *string {
	return s.PoolKey
}

func (s *BindAxnRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *BindAxnRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *BindAxnRequest) GetRingConfig() *string {
	return s.RingConfig
}

func (s *BindAxnRequest) SetASRModelId(v string) *BindAxnRequest {
	s.ASRModelId = &v
	return s
}

func (s *BindAxnRequest) SetASRStatus(v bool) *BindAxnRequest {
	s.ASRStatus = &v
	return s
}

func (s *BindAxnRequest) SetCallDisplayType(v int32) *BindAxnRequest {
	s.CallDisplayType = &v
	return s
}

func (s *BindAxnRequest) SetCallRestrict(v string) *BindAxnRequest {
	s.CallRestrict = &v
	return s
}

func (s *BindAxnRequest) SetCallTimeout(v int32) *BindAxnRequest {
	s.CallTimeout = &v
	return s
}

func (s *BindAxnRequest) SetExpectCity(v string) *BindAxnRequest {
	s.ExpectCity = &v
	return s
}

func (s *BindAxnRequest) SetExpiration(v string) *BindAxnRequest {
	s.Expiration = &v
	return s
}

func (s *BindAxnRequest) SetExtend(v string) *BindAxnRequest {
	s.Extend = &v
	return s
}

func (s *BindAxnRequest) SetIsRecordingEnabled(v bool) *BindAxnRequest {
	s.IsRecordingEnabled = &v
	return s
}

func (s *BindAxnRequest) SetNoType(v string) *BindAxnRequest {
	s.NoType = &v
	return s
}

func (s *BindAxnRequest) SetOutId(v string) *BindAxnRequest {
	s.OutId = &v
	return s
}

func (s *BindAxnRequest) SetOutOrderId(v string) *BindAxnRequest {
	s.OutOrderId = &v
	return s
}

func (s *BindAxnRequest) SetOwnerId(v int64) *BindAxnRequest {
	s.OwnerId = &v
	return s
}

func (s *BindAxnRequest) SetPhoneNoA(v string) *BindAxnRequest {
	s.PhoneNoA = &v
	return s
}

func (s *BindAxnRequest) SetPhoneNoB(v string) *BindAxnRequest {
	s.PhoneNoB = &v
	return s
}

func (s *BindAxnRequest) SetPhoneNoX(v string) *BindAxnRequest {
	s.PhoneNoX = &v
	return s
}

func (s *BindAxnRequest) SetPoolKey(v string) *BindAxnRequest {
	s.PoolKey = &v
	return s
}

func (s *BindAxnRequest) SetResourceOwnerAccount(v string) *BindAxnRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *BindAxnRequest) SetResourceOwnerId(v int64) *BindAxnRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *BindAxnRequest) SetRingConfig(v string) *BindAxnRequest {
	s.RingConfig = &v
	return s
}

func (s *BindAxnRequest) Validate() error {
	return dara.Validate(s)
}
