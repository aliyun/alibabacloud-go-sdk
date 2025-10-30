// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindAxbRequest interface {
	dara.Model
	String() string
	GoString() string
	SetASRModelId(v string) *BindAxbRequest
	GetASRModelId() *string
	SetASRStatus(v bool) *BindAxbRequest
	GetASRStatus() *bool
	SetCallDisplayType(v int32) *BindAxbRequest
	GetCallDisplayType() *int32
	SetCallRestrict(v string) *BindAxbRequest
	GetCallRestrict() *string
	SetCallTimeout(v int32) *BindAxbRequest
	GetCallTimeout() *int32
	SetDtmfConfig(v string) *BindAxbRequest
	GetDtmfConfig() *string
	SetExpectCity(v string) *BindAxbRequest
	GetExpectCity() *string
	SetExpiration(v string) *BindAxbRequest
	GetExpiration() *string
	SetIsRecordingEnabled(v bool) *BindAxbRequest
	GetIsRecordingEnabled() *bool
	SetOutId(v string) *BindAxbRequest
	GetOutId() *string
	SetOutOrderId(v string) *BindAxbRequest
	GetOutOrderId() *string
	SetOwnerId(v int64) *BindAxbRequest
	GetOwnerId() *int64
	SetPhoneNoA(v string) *BindAxbRequest
	GetPhoneNoA() *string
	SetPhoneNoB(v string) *BindAxbRequest
	GetPhoneNoB() *string
	SetPhoneNoX(v string) *BindAxbRequest
	GetPhoneNoX() *string
	SetPoolKey(v string) *BindAxbRequest
	GetPoolKey() *string
	SetResourceOwnerAccount(v string) *BindAxbRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *BindAxbRequest
	GetResourceOwnerId() *int64
	SetRingConfig(v string) *BindAxbRequest
	GetRingConfig() *string
}

type BindAxbRequest struct {
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
	// false
	ASRStatus *bool `json:"ASRStatus,omitempty" xml:"ASRStatus,omitempty"`
	// Re-sets the phone number display logic in the AXB binding. Fixed value: **1**, indicating that phone number X is displayed on both the calling phone and the called phone.
	//
	// >  Due to the regulatory restrictions imposed by carriers, the setting to display real phone numbers during calls does not take effect.
	//
	// example:
	//
	// 1
	CallDisplayType *int32 `json:"CallDisplayType,omitempty" xml:"CallDisplayType,omitempty"`
	// The status of the one-way call restriction. Valid values:
	//
	// 	- **CONTROL_AX_DISABLE**: Phone number A cannot be used to call phone number X.
	//
	// 	- **CONTROL_BX_DISABLE**: Phone number B cannot be used to call phone number X.
	//
	// example:
	//
	// CONTROL_AX_DISABLE
	CallRestrict *string `json:"CallRestrict,omitempty" xml:"CallRestrict,omitempty"`
	// The maximum ringing duration for each number in sequential ringing. Unit: seconds. The value ranges from 5 to 20.
	//
	// example:
	//
	// 10
	CallTimeout *int32 `json:"CallTimeout,omitempty" xml:"CallTimeout,omitempty"`
	// Specifies the dual tone multiple frequency (DTMF) key configuration in the AXB binding. The following content can be configured:
	//
	// 	- endCallIvrPhoneNo: for whom the audio is played, user A or user B.
	//
	// 	- waitingDtmfTime: the maximum waiting time after the first audio is played. The maximum waiting time is 30 seconds.
	//
	// 	- maxLoop: the maximum number of loop playback times of the first audio if the DTMF key is not matched. The maximum number of loop playback times is 5.
	//
	// 	- step1File: the name of the first audio.
	//
	// 	- step2File: the name of the second audio.
	//
	// 	- validKey: the valid key values, such as 1,2. Only two valid key values can be set, and the key values are separated by a comma (,).
	//
	// 	- waitingEndCall: The waiting duration to hang up a call. The waiting duration is allowed by a carrier. The maximum waiting duration is 10 seconds.
	//
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
	// Specifies the city to which phone number X to be selected belongs.
	//
	// 	- If no phone number for the specified city is available in the current phone number pool or this parameter is not specified, a phone number that belongs to another city is randomly selected from the current phone number pool and assigned as phone number X.
	//
	// 	- If**Number X Assignment Mode*	- is set to **Strict Matching Mode*	- and no phone number meets the requirement, the system displays an allocation error.
	//
	// example:
	//
	// hangzhou
	ExpectCity *string `json:"ExpectCity,omitempty" xml:"ExpectCity,omitempty"`
	// The expiration time of the AXB binding.
	//
	// >  The expiration time must be more than 1 minute later than the time when you call this API operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2021-09-05 12:00:00
	Expiration *string `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
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
	// 34553330****
	OutOrderId *string `json:"OutOrderId,omitempty" xml:"OutOrderId,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Phone number A in the AXB binding.
	//
	// Phone number A can be set to a mobile phone number or a landline phone number. The landline phone number must be added with an area code, and no hyphen is required between the area code and the landline phone number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 139****0000
	PhoneNoA *string `json:"PhoneNoA,omitempty" xml:"PhoneNoA,omitempty"`
	// Phone number B in the AXB binding. If phone number A is used to call phone number X, the call is forwarded to phone number B. Phone number B can be set to a mobile phone number or a landline phone number. The landline phone number must be added with an area code, and no hyphen is required between the area code and the landline phone number.
	//
	// >  If you need to update phone number B, call the [UpdateSubscription](https://help.aliyun.com/document_detail/110253.html) operation.
	//
	// example:
	//
	// 139****0000
	PhoneNoB *string `json:"PhoneNoB,omitempty" xml:"PhoneNoB,omitempty"`
	// Phone number X in the AXB binding.
	//
	// Phone number X is the phone number that you purchased in the [Phone Number Protection console](https://dypls.console.aliyun.com/dypls.htm#/account) or by using the [BuySecretNo](https://help.aliyun.com/document_detail/110266.html) operation before you bind a phone number. Phone number X is used to forward calls.
	//
	// If you do not specify this parameter, a random phone number is selected from the phone number pool based on the value of the ExpectCity parameter and is used as phone number X.
	//
	// example:
	//
	// 139****0000
	PhoneNoX *string `json:"PhoneNoX,omitempty" xml:"PhoneNoX,omitempty"`
	// The key of the phone number pool.
	//
	// Log on to the [Phone Number Protection console](https://dypls.console.aliyun.com/dypls.htm#/account) and view the key of the phone number pool on the **Number Pool Management*	- page.
	//
	// example:
	//
	// FC5526*****
	PoolKey              *string `json:"PoolKey,omitempty" xml:"PoolKey,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Sets the ringtone code for enterprise Color Ring Back Tone (CRBT) in the AXB binding.
	//
	// 	- Ringtone setting when phone number A is used to call phone number X in the AXB binding: AXBRing_A
	//
	// 	- Ringtone setting when phone number B is used to call phone number X in the AXB binding: AXBRing_B
	//
	// Enterprise CRBT codes: Enterprise CRBT codes can be queried in the [Phone Number Protection console](https://dypls.console.aliyun.com/dypls.htm#/account). You can choose **Number Pool Management*	- > **Enterprise CRBT Management*	- to view enterprise CRBT codes. You can also upload, delete, or perform other operations on enterprise CRBT codes.
	//
	// >  The bound enterprise CRBTs are preferentially used. If no enterprise CRBT is set or the setting does not take effect, the enterprise CRBTs at the phone number pool level are used.
	//
	// example:
	//
	// {"AXBRing_B":"1000****1","AXBRing_A":"1000****1"}
	RingConfig *string `json:"RingConfig,omitempty" xml:"RingConfig,omitempty"`
}

func (s BindAxbRequest) String() string {
	return dara.Prettify(s)
}

func (s BindAxbRequest) GoString() string {
	return s.String()
}

func (s *BindAxbRequest) GetASRModelId() *string {
	return s.ASRModelId
}

func (s *BindAxbRequest) GetASRStatus() *bool {
	return s.ASRStatus
}

func (s *BindAxbRequest) GetCallDisplayType() *int32 {
	return s.CallDisplayType
}

func (s *BindAxbRequest) GetCallRestrict() *string {
	return s.CallRestrict
}

func (s *BindAxbRequest) GetCallTimeout() *int32 {
	return s.CallTimeout
}

func (s *BindAxbRequest) GetDtmfConfig() *string {
	return s.DtmfConfig
}

func (s *BindAxbRequest) GetExpectCity() *string {
	return s.ExpectCity
}

func (s *BindAxbRequest) GetExpiration() *string {
	return s.Expiration
}

func (s *BindAxbRequest) GetIsRecordingEnabled() *bool {
	return s.IsRecordingEnabled
}

func (s *BindAxbRequest) GetOutId() *string {
	return s.OutId
}

func (s *BindAxbRequest) GetOutOrderId() *string {
	return s.OutOrderId
}

func (s *BindAxbRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *BindAxbRequest) GetPhoneNoA() *string {
	return s.PhoneNoA
}

func (s *BindAxbRequest) GetPhoneNoB() *string {
	return s.PhoneNoB
}

func (s *BindAxbRequest) GetPhoneNoX() *string {
	return s.PhoneNoX
}

func (s *BindAxbRequest) GetPoolKey() *string {
	return s.PoolKey
}

func (s *BindAxbRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *BindAxbRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *BindAxbRequest) GetRingConfig() *string {
	return s.RingConfig
}

func (s *BindAxbRequest) SetASRModelId(v string) *BindAxbRequest {
	s.ASRModelId = &v
	return s
}

func (s *BindAxbRequest) SetASRStatus(v bool) *BindAxbRequest {
	s.ASRStatus = &v
	return s
}

func (s *BindAxbRequest) SetCallDisplayType(v int32) *BindAxbRequest {
	s.CallDisplayType = &v
	return s
}

func (s *BindAxbRequest) SetCallRestrict(v string) *BindAxbRequest {
	s.CallRestrict = &v
	return s
}

func (s *BindAxbRequest) SetCallTimeout(v int32) *BindAxbRequest {
	s.CallTimeout = &v
	return s
}

func (s *BindAxbRequest) SetDtmfConfig(v string) *BindAxbRequest {
	s.DtmfConfig = &v
	return s
}

func (s *BindAxbRequest) SetExpectCity(v string) *BindAxbRequest {
	s.ExpectCity = &v
	return s
}

func (s *BindAxbRequest) SetExpiration(v string) *BindAxbRequest {
	s.Expiration = &v
	return s
}

func (s *BindAxbRequest) SetIsRecordingEnabled(v bool) *BindAxbRequest {
	s.IsRecordingEnabled = &v
	return s
}

func (s *BindAxbRequest) SetOutId(v string) *BindAxbRequest {
	s.OutId = &v
	return s
}

func (s *BindAxbRequest) SetOutOrderId(v string) *BindAxbRequest {
	s.OutOrderId = &v
	return s
}

func (s *BindAxbRequest) SetOwnerId(v int64) *BindAxbRequest {
	s.OwnerId = &v
	return s
}

func (s *BindAxbRequest) SetPhoneNoA(v string) *BindAxbRequest {
	s.PhoneNoA = &v
	return s
}

func (s *BindAxbRequest) SetPhoneNoB(v string) *BindAxbRequest {
	s.PhoneNoB = &v
	return s
}

func (s *BindAxbRequest) SetPhoneNoX(v string) *BindAxbRequest {
	s.PhoneNoX = &v
	return s
}

func (s *BindAxbRequest) SetPoolKey(v string) *BindAxbRequest {
	s.PoolKey = &v
	return s
}

func (s *BindAxbRequest) SetResourceOwnerAccount(v string) *BindAxbRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *BindAxbRequest) SetResourceOwnerId(v int64) *BindAxbRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *BindAxbRequest) SetRingConfig(v string) *BindAxbRequest {
	s.RingConfig = &v
	return s
}

func (s *BindAxbRequest) Validate() error {
	return dara.Validate(s)
}
