// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindAxgRequest interface {
	dara.Model
	String() string
	GoString() string
	SetASRModelId(v string) *BindAxgRequest
	GetASRModelId() *string
	SetASRStatus(v bool) *BindAxgRequest
	GetASRStatus() *bool
	SetCallDisplayType(v int32) *BindAxgRequest
	GetCallDisplayType() *int32
	SetCallRestrict(v string) *BindAxgRequest
	GetCallRestrict() *string
	SetExpectCity(v string) *BindAxgRequest
	GetExpectCity() *string
	SetExpiration(v string) *BindAxgRequest
	GetExpiration() *string
	SetGroupId(v string) *BindAxgRequest
	GetGroupId() *string
	SetIsRecordingEnabled(v bool) *BindAxgRequest
	GetIsRecordingEnabled() *bool
	SetOutId(v string) *BindAxgRequest
	GetOutId() *string
	SetOutOrderId(v string) *BindAxgRequest
	GetOutOrderId() *string
	SetOwnerId(v int64) *BindAxgRequest
	GetOwnerId() *int64
	SetPhoneNoA(v string) *BindAxgRequest
	GetPhoneNoA() *string
	SetPhoneNoB(v string) *BindAxgRequest
	GetPhoneNoB() *string
	SetPhoneNoX(v string) *BindAxgRequest
	GetPhoneNoX() *string
	SetPoolKey(v string) *BindAxgRequest
	GetPoolKey() *string
	SetResourceOwnerAccount(v string) *BindAxgRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *BindAxgRequest
	GetResourceOwnerId() *int64
	SetRingConfig(v string) *BindAxgRequest
	GetRingConfig() *string
}

type BindAxgRequest struct {
	// The ID of the ASR model.
	//
	// example:
	//
	// 980abddb908f48e8b987cb2cd303****
	ASRModelId *string `json:"ASRModelId,omitempty" xml:"ASRModelId,omitempty"`
	// Specifies whether to enable automatic speech recognition (ASR). Valid values:
	//
	// 	- **False*	- (default): disables ASR.
	//
	// 	- **True**: enables ASR.
	//
	// example:
	//
	// False
	ASRStatus *bool `json:"ASRStatus,omitempty" xml:"ASRStatus,omitempty"`
	// Re-sets the phone number display logic in the AXG binding. Fixed value: **1**, indicating that phone number X is displayed on both the calling phone and the called phone.
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
	// The expiration time of the AXG binding. The value is accurate to seconds.
	//
	// >  The expiration time must be more than 1 minute later than the time when you call this API operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2019-09-05 12:00:00
	Expiration *string `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	// The group ID in the AXG binding. You can view the group ID by using either of the following methods:
	//
	// 	- On the **Number Pool Management*	- page in the [Phone Number Protection console](https://dypls.console.aliyun.com/dypls.htm#/account), filter AXG privacy numbers and click **Number Group G Management*	- to view the group IDs of number groups G.****
	//
	// 	- If the [CreateAxgGroup](https://help.aliyun.com/document_detail/110250.html) operation is called to create number group G, the value of **GroupId*	- in the response of the CreateAxgGroup operation is specified as the value of the request parameter **GroupId*	- of the BindAxg operation.
	//
	// >  Number group G must have one or more phone numbers.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// Specifies whether to record all calls made by the bound phone numbers.
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
	// Phone number A in the AXG binding. Phone number A can be set to a mobile phone number or a landline phone number. The landline phone number must be added with an area code, and no hyphen is required between the area code and the landline phone number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 139****0000
	PhoneNoA *string `json:"PhoneNoA,omitempty" xml:"PhoneNoA,omitempty"`
	// Phone number B in the AXG binding. If phone number A is used to call phone number X, the call is forwarded to phone number B. If you need to update phone number B, call the [UpdateSubscription](https://help.aliyun.com/document_detail/110253.html) operation.
	//
	// Phone number B can be set to a mobile phone number or a landline phone number. The landline phone number must be added with an area code, and no hyphen is required between the area code and the landline phone number.
	//
	// example:
	//
	// 139****0000
	PhoneNoB *string `json:"PhoneNoB,omitempty" xml:"PhoneNoB,omitempty"`
	// Phone number X in the AXG binding. If you do not specify this parameter, a random phone number is selected from the phone number pool based on the value of the **ExpectCity*	- parameter and is used as phone number X.
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
	// Sets the ringtone for enterprise Color Ring Back Tone (CRBT) in the AXG binding.
	//
	// 	- Ringtone setting (with a callback number) when phone number A is used to call phone number X in the AXG binding: AXGRing_AB
	//
	// 	- Ringtone setting (without a callback number) when phone number A is used to call phone number X in the AXG binding: AXGRing_A
	//
	// 	- Ringtone setting when a phone number in number group G is used to call phone number X in the AXG binding: AXGRing_G
	//
	// 	- Enterprise CRBT codes: Enterprise CRBT codes can be queried in the Phone Number Protection console. You can choose **Number Pool Management > Enterprise CRBT Management*	- to view and manage enterprise CRBT codes. You can also upload, delete, or perform other operations on enterprise CRBT codes.
	//
	// >  The bound enterprise CRBTs are preferentially used. If no enterprise CRBT is set or the setting does not take effect, the enterprise CRBTs at the phone number pool level are used.
	//
	// example:
	//
	// {"AXGRing_AB":"100000001","AXGRing_A":"100000002","AXGRing_G":"100000003"}
	RingConfig *string `json:"RingConfig,omitempty" xml:"RingConfig,omitempty"`
}

func (s BindAxgRequest) String() string {
	return dara.Prettify(s)
}

func (s BindAxgRequest) GoString() string {
	return s.String()
}

func (s *BindAxgRequest) GetASRModelId() *string {
	return s.ASRModelId
}

func (s *BindAxgRequest) GetASRStatus() *bool {
	return s.ASRStatus
}

func (s *BindAxgRequest) GetCallDisplayType() *int32 {
	return s.CallDisplayType
}

func (s *BindAxgRequest) GetCallRestrict() *string {
	return s.CallRestrict
}

func (s *BindAxgRequest) GetExpectCity() *string {
	return s.ExpectCity
}

func (s *BindAxgRequest) GetExpiration() *string {
	return s.Expiration
}

func (s *BindAxgRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *BindAxgRequest) GetIsRecordingEnabled() *bool {
	return s.IsRecordingEnabled
}

func (s *BindAxgRequest) GetOutId() *string {
	return s.OutId
}

func (s *BindAxgRequest) GetOutOrderId() *string {
	return s.OutOrderId
}

func (s *BindAxgRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *BindAxgRequest) GetPhoneNoA() *string {
	return s.PhoneNoA
}

func (s *BindAxgRequest) GetPhoneNoB() *string {
	return s.PhoneNoB
}

func (s *BindAxgRequest) GetPhoneNoX() *string {
	return s.PhoneNoX
}

func (s *BindAxgRequest) GetPoolKey() *string {
	return s.PoolKey
}

func (s *BindAxgRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *BindAxgRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *BindAxgRequest) GetRingConfig() *string {
	return s.RingConfig
}

func (s *BindAxgRequest) SetASRModelId(v string) *BindAxgRequest {
	s.ASRModelId = &v
	return s
}

func (s *BindAxgRequest) SetASRStatus(v bool) *BindAxgRequest {
	s.ASRStatus = &v
	return s
}

func (s *BindAxgRequest) SetCallDisplayType(v int32) *BindAxgRequest {
	s.CallDisplayType = &v
	return s
}

func (s *BindAxgRequest) SetCallRestrict(v string) *BindAxgRequest {
	s.CallRestrict = &v
	return s
}

func (s *BindAxgRequest) SetExpectCity(v string) *BindAxgRequest {
	s.ExpectCity = &v
	return s
}

func (s *BindAxgRequest) SetExpiration(v string) *BindAxgRequest {
	s.Expiration = &v
	return s
}

func (s *BindAxgRequest) SetGroupId(v string) *BindAxgRequest {
	s.GroupId = &v
	return s
}

func (s *BindAxgRequest) SetIsRecordingEnabled(v bool) *BindAxgRequest {
	s.IsRecordingEnabled = &v
	return s
}

func (s *BindAxgRequest) SetOutId(v string) *BindAxgRequest {
	s.OutId = &v
	return s
}

func (s *BindAxgRequest) SetOutOrderId(v string) *BindAxgRequest {
	s.OutOrderId = &v
	return s
}

func (s *BindAxgRequest) SetOwnerId(v int64) *BindAxgRequest {
	s.OwnerId = &v
	return s
}

func (s *BindAxgRequest) SetPhoneNoA(v string) *BindAxgRequest {
	s.PhoneNoA = &v
	return s
}

func (s *BindAxgRequest) SetPhoneNoB(v string) *BindAxgRequest {
	s.PhoneNoB = &v
	return s
}

func (s *BindAxgRequest) SetPhoneNoX(v string) *BindAxgRequest {
	s.PhoneNoX = &v
	return s
}

func (s *BindAxgRequest) SetPoolKey(v string) *BindAxgRequest {
	s.PoolKey = &v
	return s
}

func (s *BindAxgRequest) SetResourceOwnerAccount(v string) *BindAxgRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *BindAxgRequest) SetResourceOwnerId(v int64) *BindAxgRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *BindAxgRequest) SetRingConfig(v string) *BindAxgRequest {
	s.RingConfig = &v
	return s
}

func (s *BindAxgRequest) Validate() error {
	return dara.Validate(s)
}
