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
	// The ASR model ID. Valid values:
	//
	// - General-Telephone-Mandarin Chinese model ID: 7ee372834d2f4cc7ac0d0ab2d0ae1aac
	//
	// - General-Telephone-Cantonese model ID: 9b4afe3d224e4fa3bf963697113e2f77
	//
	// - General-Telephone-Regional accent model ID: 3655a6cb94764145a1f9a38a3c3b3b4b
	//
	// example:
	//
	// 7ee372834d2f4cc7ac0d0ab2d0ae1aac
	ASRModelId *string `json:"ASRModelId,omitempty" xml:"ASRModelId,omitempty"`
	// The ASR status. Valid values:
	//
	// - **false**: Disabled. This is the default value.
	//
	// - **true**: Enabled.
	//
	// example:
	//
	// true
	ASRStatus *bool `json:"ASRStatus,omitempty" xml:"ASRStatus,omitempty"`
	// Resets the number display logic in the binding relationship. Set the value to **1**, which indicates that the intermediate number X is displayed for both the caller and the callee.
	//
	// > Due to carrier regulatory restrictions, the setting to display the real number during a call does not take effect.
	//
	// example:
	//
	// 1
	CallDisplayType *int32 `json:"CallDisplayType,omitempty" xml:"CallDisplayType,omitempty"`
	// The call restriction status. Valid values:
	//
	// - **CONTROL_AX_DISABLE**: Number A cannot call number X.
	//
	// - **CONTROL_BX_DISABLE**: Number B cannot call number X.
	//
	// example:
	//
	// CONTROL_AX_DISABLE
	CallRestrict *string `json:"CallRestrict,omitempty" xml:"CallRestrict,omitempty"`
	// The city where you want to select number X.
	//
	// > - If no available numbers exist in the specified city in the current number pool, or if this parameter is not specified, the system randomly assigns a number from another city in the current number pool as number X.
	//
	// > - If the **number X allocation mode*	- is set to **strict match mode**, the system returns an allocation error when no eligible numbers are available.
	//
	// example:
	//
	// Beijing
	ExpectCity *string `json:"ExpectCity,omitempty" xml:"ExpectCity,omitempty"`
	// The expiration time of the binding relationship. The time is accurate to seconds.
	//
	// > The expiration time must be at least 1 minute later than the current time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2019-09-05 12:00:00
	Expiration *string `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	Extend     *string `json:"Extend,omitempty" xml:"Extend,omitempty"`
	// The extension number of number X. The extension number must be 1 to 3 digits in length.
	//
	// > If you specify an extension number, you must also specify number X. This means that both the Extension and PhoneNoX parameters must be specified.
	//
	// example:
	//
	// 130
	Extension *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// Specifies whether to record all calls generated for this binding relationship. Valid values:
	//
	// - **true**: Records calls.
	//
	// - **false**: Does not record calls. This is the default value.
	//
	// example:
	//
	// true
	IsRecordingEnabled *bool `json:"IsRecordingEnabled,omitempty" xml:"IsRecordingEnabled,omitempty"`
	// The external business extension field. This parameter is returned in the call record receipt message.
	//
	// example:
	//
	// abcdef
	OutId *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	// The external business ID.
	//
	// example:
	//
	// abcdef
	OutOrderId *string `json:"OutOrderId,omitempty" xml:"OutOrderId,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number A of the AXN extension number.
	//
	// Number A can be a mobile phone number or a landline number. For a landline number, you must add the area code. Do not add a hyphen between the area code and the phone number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 139****0000
	PhoneNoA *string `json:"PhoneNoA,omitempty" xml:"PhoneNoA,omitempty"`
	// The number B of the AXN extension number.
	//
	// Number B can be a mobile phone number or a landline number. For a landline number, you must add the area code. Do not add a hyphen between the area code and the phone number.
	//
	// > - When number A calls number X, the call is forwarded to number B by default. If number B is not specified, the call cannot be connected.
	//
	// > - When number A calls number X, the call is forwarded to number B only when the **callback mode for number A*	- is set to **default mode**. If it is not set to **default mode**, the call cannot be connected.
	//
	// > - To modify the **callback mode for number A**, go to [Number Pool Management](https://dyplsnext.console.aliyun.com/number-pool?resType=1) > filter by **AXN Extension Number*	- > **Number Pool Management*	- > **Callback Mode for Number A*	- > **Adjust Configuration**.
	//
	// > - To update number B, call the [UpdateSubscription](https://www.alibabacloud.com/help/en/pnp/developer-reference/api-dyplsapi-2017-05-25-updatesubscription) operation.
	//
	// example:
	//
	// 138****0000
	PhoneNoB *string `json:"PhoneNoB,omitempty" xml:"PhoneNoB,omitempty"`
	// The number X of the AXN extension number.
	//
	// Number X is a privacy number that you have [purchased](https://dyplsnext.console.aliyun.com/numberbuy?spm=a2c4g.11186623.0.0.7e6449bbtiHP8m) in the console.
	//
	// > If you do not specify number X, the system randomly assigns a number from the specified number pool as number X based on the ExpectCity parameter.
	//
	// example:
	//
	// 139****1111
	PhoneNoX *string `json:"PhoneNoX,omitempty" xml:"PhoneNoX,omitempty"`
	// The key of the number pool.
	//
	// You can obtain the **number pool key*	- on the [Number Pool Management](https://dyplsnext.console.aliyun.com/number-pool?spm=a2c4g.11186623.0.0.7e6449bbtiHP8m) page in the Privacy Protection console.
	//
	// example:
	//
	// FC10000011111****
	PoolKey              *string `json:"PoolKey,omitempty" xml:"PoolKey,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Set the corporate ringback tone for dial-up to number X after the AXN extension number is bound.
	//
	// Specify the **ringtone field*	- and **ringback tone encoding*	- for this parameter.
	//
	// Valid values for the **ringtone field**:
	//
	// - AXNRing_AB: The ringtone settings before number A calls number X and the call is connected (with a callback number) in the AXN extension number scenario.
	//
	// - AXNRing_A: The ringtone settings before number A calls number X and the call is connected (without a callback number) in the AXN extension number scenario.
	//
	// - AXNRing_N: The ringtone settings before the number on side N calls number X and the call is connected in the AXN extension number scenario.
	//
	// - AXNRing_A_PICKUP: The ringtone settings after number A calls number X and the call is connected in the AXN extension number scenario.
	//
	// - AXNRing_N_PICKUP: The ringtone settings after the number on side N calls number X and the call is connected in the AXN extension number scenario.
	//
	// **Ringback tone encoding**:
	//
	// You can obtain the corporate ringback tone encoding on the [Corporate Ringback Tone Management](https://dyplsnext.console.aliyun.com/number-pool/ringtone?resType=1) page in the Privacy Protection console.
	//
	// **Example**:
	//
	// To set the ringtone before number A calls number X and the call is connected (without a callback number), enter "AXNRing_A": "1000****1".
	//
	// example:
	//
	// {"AXNRing_N":"1000****1","AXNRing_A":"1000****1"}
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
