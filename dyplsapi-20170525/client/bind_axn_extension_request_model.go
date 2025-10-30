// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindAxnExtensionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetASRModelId(v string) *BindAxnExtensionRequest
	GetASRModelId() *string
	SetASRStatus(v bool) *BindAxnExtensionRequest
	GetASRStatus() *bool
	SetCallDisplayType(v int32) *BindAxnExtensionRequest
	GetCallDisplayType() *int32
	SetCallRestrict(v string) *BindAxnExtensionRequest
	GetCallRestrict() *string
	SetExpectCity(v string) *BindAxnExtensionRequest
	GetExpectCity() *string
	SetExpiration(v string) *BindAxnExtensionRequest
	GetExpiration() *string
	SetExtend(v string) *BindAxnExtensionRequest
	GetExtend() *string
	SetExtension(v string) *BindAxnExtensionRequest
	GetExtension() *string
	SetIsRecordingEnabled(v bool) *BindAxnExtensionRequest
	GetIsRecordingEnabled() *bool
	SetOutId(v string) *BindAxnExtensionRequest
	GetOutId() *string
	SetOutOrderId(v string) *BindAxnExtensionRequest
	GetOutOrderId() *string
	SetOwnerId(v int64) *BindAxnExtensionRequest
	GetOwnerId() *int64
	SetPhoneNoA(v string) *BindAxnExtensionRequest
	GetPhoneNoA() *string
	SetPhoneNoB(v string) *BindAxnExtensionRequest
	GetPhoneNoB() *string
	SetPhoneNoX(v string) *BindAxnExtensionRequest
	GetPhoneNoX() *string
	SetPoolKey(v string) *BindAxnExtensionRequest
	GetPoolKey() *string
	SetResourceOwnerAccount(v string) *BindAxnExtensionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *BindAxnExtensionRequest
	GetResourceOwnerId() *int64
	SetRingConfig(v string) *BindAxnExtensionRequest
	GetRingConfig() *string
}

type BindAxnExtensionRequest struct {
	// The ID of the ASR model. On the [Automatic Speech Recognition (ASR) Model Management](https://dyplsnext.console.aliyun.com/?spm=5176.12818093.categories-n-products.ddypls.22e616d0a0tEFC#/asr) page, you can view the ID of the ASR model.
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
	// True
	ASRStatus *bool `json:"ASRStatus,omitempty" xml:"ASRStatus,omitempty"`
	// Re-sets the phone number display logic in the AXN extension binding. Fixed value: **1**, indicating that phone number X is displayed on both the calling phone and the called phone.
	//
	// >  Due to the regulatory restrictions imposed by carriers, the setting to display real phone numbers during calls does not take effect.
	//
	// example:
	//
	// 1
	CallDisplayType *int32 `json:"CallDisplayType,omitempty" xml:"CallDisplayType,omitempty"`
	// The status of call restrictions. Valid values:
	//
	// 	- **CONTROL_AX_DISABLE**: Phone number A cannot be used to call phone number X.
	//
	// 	- **CONTROL_BX_DISABLE**: Phone number B cannot be used to call phone number X.
	//
	// example:
	//
	// CONTROL_AX_DISABLE
	CallRestrict *string `json:"CallRestrict,omitempty" xml:"CallRestrict,omitempty"`
	// Specifies the city to which phone number X to be selected belongs.
	//
	// 	- If no phone number for the specified city is available in the current phone number pool or this parameter is not specified, a phone number that belongs to another city is randomly selected from the current phone number pool and assigned as phone number X.
	//
	// 	- If Number X Assignment Mode is set to Strict Matching Mode and no phone number meets the requirement, the system displays an allocation error.
	//
	// example:
	//
	// hangzhou
	ExpectCity *string `json:"ExpectCity,omitempty" xml:"ExpectCity,omitempty"`
	// The expiration time of the AXN extension binding. The value is accurate to seconds.
	//
	// >  The expiration time must be more than 1 minute later than the time when you call this API operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2019-09-05 12:00:00
	Expiration *string `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	Extend     *string `json:"Extend,omitempty" xml:"Extend,omitempty"`
	// The extension of phone number X. The extension is 1 to 3 digits in length.
	//
	// >  If you specify Extension, you must also specify PhoneNoX.
	//
	// example:
	//
	// 130
	Extension *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
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
	// abcdef
	OutOrderId *string `json:"OutOrderId,omitempty" xml:"OutOrderId,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Phone number A in the AXN extension binding. Phone number A can be set to a mobile phone number or a landline phone number. The landline phone number must be added with an area code, and no hyphen is required between the area code and the landline phone number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 139****0000
	PhoneNoA *string `json:"PhoneNoA,omitempty" xml:"PhoneNoA,omitempty"`
	// Phone number B in the AXN extension binding. When phone number A is used to call phone number X, the call is forwarded to phone number B. If you need to update phone number B, call the [UpdateSubscription](https://help.aliyun.com/document_detail/110253.html) operation.
	//
	// Phone number B can be set to a mobile phone number or a landline phone number. The landline phone number must be added with an area code, and no hyphen is required between the area code and the landline phone number.
	//
	// example:
	//
	// 139****0000
	PhoneNoB *string `json:"PhoneNoB,omitempty" xml:"PhoneNoB,omitempty"`
	// Phone number X in the AXN extension binding. If you do not specify this parameter, a random phone number is selected from the phone number pool based on the value of the **ExpectCity*	- parameter and is used as phone number X.
	//
	// >  Phone number X is the phone number that you purchased in the Phone Number Protection console or by using the [BuySecretNo](https://help.aliyun.com/document_detail/110266.html) operation before you bind a phone number. Phone number X is used to forward calls.
	//
	// example:
	//
	// 139****0000
	PhoneNoX *string `json:"PhoneNoX,omitempty" xml:"PhoneNoX,omitempty"`
	// The key of the phone number pool. Log on to the [Phone Number Protection console](https://dypls.console.aliyun.com/dypls.htm#/account) and view the key of the phone number pool on the **Number Pool Management*	- page.
	//
	// example:
	//
	// FC123456
	PoolKey              *string `json:"PoolKey,omitempty" xml:"PoolKey,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Sets the ringtone for enterprise Color Ring Back Tone (CRBT) in the AXN extension binding.
	//
	// 	- Ringtone setting (with a callback number) when phone number A is used to call phone number X in the AXN extension binding: AXNRing_AB
	//
	// 	- Ringtone setting (without a callback number) when phone number A is used to call phone number X in the AXN extension binding: AXNRing_A
	//
	// 	- Ringtone setting when an N-side number is used to call phone number X in the AXN extension binding: AXNRing_N
	//
	// Enterprise CRBT codes: Enterprise CRBT codes can be queried in the Phone Number Protection console. You can choose **Number Pool Management > Enterprise CRBT Management*	- to view and manage enterprise CRBT codes. You can also upload, delete, or perform other operations on enterprise CRBT codes.
	//
	// >  The bound enterprise CRBTs are preferentially used. If no enterprise CRBT is set or the setting does not take effect, the enterprise CRBTs at the phone number pool level are used.
	//
	// example:
	//
	// {"AXNRing_N":"100000001","AXNRing_A":"100000001"}
	RingConfig *string `json:"RingConfig,omitempty" xml:"RingConfig,omitempty"`
}

func (s BindAxnExtensionRequest) String() string {
	return dara.Prettify(s)
}

func (s BindAxnExtensionRequest) GoString() string {
	return s.String()
}

func (s *BindAxnExtensionRequest) GetASRModelId() *string {
	return s.ASRModelId
}

func (s *BindAxnExtensionRequest) GetASRStatus() *bool {
	return s.ASRStatus
}

func (s *BindAxnExtensionRequest) GetCallDisplayType() *int32 {
	return s.CallDisplayType
}

func (s *BindAxnExtensionRequest) GetCallRestrict() *string {
	return s.CallRestrict
}

func (s *BindAxnExtensionRequest) GetExpectCity() *string {
	return s.ExpectCity
}

func (s *BindAxnExtensionRequest) GetExpiration() *string {
	return s.Expiration
}

func (s *BindAxnExtensionRequest) GetExtend() *string {
	return s.Extend
}

func (s *BindAxnExtensionRequest) GetExtension() *string {
	return s.Extension
}

func (s *BindAxnExtensionRequest) GetIsRecordingEnabled() *bool {
	return s.IsRecordingEnabled
}

func (s *BindAxnExtensionRequest) GetOutId() *string {
	return s.OutId
}

func (s *BindAxnExtensionRequest) GetOutOrderId() *string {
	return s.OutOrderId
}

func (s *BindAxnExtensionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *BindAxnExtensionRequest) GetPhoneNoA() *string {
	return s.PhoneNoA
}

func (s *BindAxnExtensionRequest) GetPhoneNoB() *string {
	return s.PhoneNoB
}

func (s *BindAxnExtensionRequest) GetPhoneNoX() *string {
	return s.PhoneNoX
}

func (s *BindAxnExtensionRequest) GetPoolKey() *string {
	return s.PoolKey
}

func (s *BindAxnExtensionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *BindAxnExtensionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *BindAxnExtensionRequest) GetRingConfig() *string {
	return s.RingConfig
}

func (s *BindAxnExtensionRequest) SetASRModelId(v string) *BindAxnExtensionRequest {
	s.ASRModelId = &v
	return s
}

func (s *BindAxnExtensionRequest) SetASRStatus(v bool) *BindAxnExtensionRequest {
	s.ASRStatus = &v
	return s
}

func (s *BindAxnExtensionRequest) SetCallDisplayType(v int32) *BindAxnExtensionRequest {
	s.CallDisplayType = &v
	return s
}

func (s *BindAxnExtensionRequest) SetCallRestrict(v string) *BindAxnExtensionRequest {
	s.CallRestrict = &v
	return s
}

func (s *BindAxnExtensionRequest) SetExpectCity(v string) *BindAxnExtensionRequest {
	s.ExpectCity = &v
	return s
}

func (s *BindAxnExtensionRequest) SetExpiration(v string) *BindAxnExtensionRequest {
	s.Expiration = &v
	return s
}

func (s *BindAxnExtensionRequest) SetExtend(v string) *BindAxnExtensionRequest {
	s.Extend = &v
	return s
}

func (s *BindAxnExtensionRequest) SetExtension(v string) *BindAxnExtensionRequest {
	s.Extension = &v
	return s
}

func (s *BindAxnExtensionRequest) SetIsRecordingEnabled(v bool) *BindAxnExtensionRequest {
	s.IsRecordingEnabled = &v
	return s
}

func (s *BindAxnExtensionRequest) SetOutId(v string) *BindAxnExtensionRequest {
	s.OutId = &v
	return s
}

func (s *BindAxnExtensionRequest) SetOutOrderId(v string) *BindAxnExtensionRequest {
	s.OutOrderId = &v
	return s
}

func (s *BindAxnExtensionRequest) SetOwnerId(v int64) *BindAxnExtensionRequest {
	s.OwnerId = &v
	return s
}

func (s *BindAxnExtensionRequest) SetPhoneNoA(v string) *BindAxnExtensionRequest {
	s.PhoneNoA = &v
	return s
}

func (s *BindAxnExtensionRequest) SetPhoneNoB(v string) *BindAxnExtensionRequest {
	s.PhoneNoB = &v
	return s
}

func (s *BindAxnExtensionRequest) SetPhoneNoX(v string) *BindAxnExtensionRequest {
	s.PhoneNoX = &v
	return s
}

func (s *BindAxnExtensionRequest) SetPoolKey(v string) *BindAxnExtensionRequest {
	s.PoolKey = &v
	return s
}

func (s *BindAxnExtensionRequest) SetResourceOwnerAccount(v string) *BindAxnExtensionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *BindAxnExtensionRequest) SetResourceOwnerId(v int64) *BindAxnExtensionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *BindAxnExtensionRequest) SetRingConfig(v string) *BindAxnExtensionRequest {
	s.RingConfig = &v
	return s
}

func (s *BindAxnExtensionRequest) Validate() error {
	return dara.Validate(s)
}
