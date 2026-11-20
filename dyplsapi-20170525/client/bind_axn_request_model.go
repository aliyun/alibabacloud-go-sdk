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
	// The ASR model ID. Valid values:
	//
	// - 7ee372834d2f4cc7ac0d0ab2d0ae1aac: General-Telephone-Mandarin Chinese model.
	//
	// - 9b4afe3d224e4fa3bf963697113e2f77: General-Telephone-Cantonese model.
	//
	// - 3655a6cb94764145a1f9a38a3c3b3b4b: General-Telephone-Regional accent model.
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
	// Resets the number display logic in the binding relationship. Set the value to **1**, which indicates that the intermediate Number X is displayed for both the caller and the callee.
	//
	// > Due to carrier regulatory restrictions, the setting to display the real number during a call does not take effect.
	//
	// example:
	//
	// 1
	CallDisplayType *int32 `json:"CallDisplayType,omitempty" xml:"CallDisplayType,omitempty"`
	// The single-call restriction status. Valid values:
	//
	// - **CONTROL_AX_DISABLE**: Number A cannot call Number X.
	//
	// - **CONTROL_BX_DISABLE**: Number B cannot call Number X.
	//
	// example:
	//
	// CONTROL_AX_DISABLE
	CallRestrict *string `json:"CallRestrict,omitempty" xml:"CallRestrict,omitempty"`
	// The ring duration. Unit: seconds.
	//
	// example:
	//
	// 10
	CallTimeout *int32 `json:"CallTimeout,omitempty" xml:"CallTimeout,omitempty"`
	// The city specified for selecting Number X.
	//
	// > - If no available numbers in the specified city exist in the current number pool, or if this parameter is not specified, a number from another city is randomly assigned from the current number pool as Number X.
	//
	// > - If the **Number X allocation mode*	- is set to **strict match mode**, the system returns an allocation error when no eligible numbers are available.
	//
	// example:
	//
	// Shanghai
	ExpectCity *string `json:"ExpectCity,omitempty" xml:"ExpectCity,omitempty"`
	// The expiration time of the binding relationship. Unit: seconds.
	//
	// > The expiration time must be at least 60 seconds later than the current time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2021-09-05 12:00:00
	Expiration *string `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	Extend     *string `json:"Extend,omitempty" xml:"Extend,omitempty"`
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
	// The number type.
	//
	// > This parameter applies to former Alibaba Dayu customers. Alibaba Cloud users can ignore this parameter.
	//
	// example:
	//
	// AXB_170
	NoType *string `json:"NoType,omitempty" xml:"NoType,omitempty"`
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
	// 34553330****
	OutOrderId *string `json:"OutOrderId,omitempty" xml:"OutOrderId,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Number A of the AXN privacy number. Number A can be a mobile phone number or a landline number. If it is a landline number, the area code must be included, and no hyphen is required between the area code and the phone number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 139****0000
	PhoneNoA *string `json:"PhoneNoA,omitempty" xml:"PhoneNoA,omitempty"`
	// Number B of the AXN privacy number.
	//
	// Number B can be a mobile phone number or a landline number. If it is a landline number, the area code must be included, and no hyphen is required between the area code and the phone number.
	//
	// > - When Number A dials Number X, the call is forwarded to Number B by default. If Number B is not set, the call cannot be connected.
	//
	// > - When Number A dials Number X, the call is forwarded to Number B only when the **callback mode for Number A*	- is set to **default mode**. Otherwise, the call cannot be connected.
	//
	// > - To modify the **callback mode for Number A**, go to [Number Pool Management](https://dyplsnext.console.aliyun.com/number-pool?resType=1) > filter **AXN privacy number*	- > **Number Pool Management*	- > **Callback mode for Number A*	- > **Adjust Configuration**.
	//
	// > - To update Number B, call the [UpdateSubscription](https://www.alibabacloud.com/help/en/pnp/developer-reference/api-dyplsapi-2017-05-25-updatesubscription) operation.
	//
	// example:
	//
	// 138****0000
	PhoneNoB *string `json:"PhoneNoB,omitempty" xml:"PhoneNoB,omitempty"`
	// Number X of the AXN privacy number.
	//
	// Number X is a privacy number that you have [purchased](https://dyplsnext.console.aliyun.com/numberbuy?spm=a2c4g.11186623.0.0.7e6449bbtiHP8m) in the console.
	//
	// > If Number X is not specified, a number is randomly selected from the specified number pool based on the ExpectCity parameter as Number X.
	//
	// example:
	//
	// 139****1111
	PhoneNoX *string `json:"PhoneNoX,omitempty" xml:"PhoneNoX,omitempty"`
	// The key of the number pool. You can obtain the **number pool key*	- on the [Number Pool Management](https://dyplsnext.console.aliyun.com/number-pool?spm=a2c4g.11186623.0.0.7e6449bbtiHP8m) page in the Privacy Protection console.
	//
	// example:
	//
	// FC2256****
	PoolKey              *string `json:"PoolKey,omitempty" xml:"PoolKey,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Sets the corporate color ringback tone (CRBT) played during a dial-up to Number X after the AXN privacy number is attached.
	//
	// Specify the **ringtone field*	- and **CRBT encoding*	- for this parameter.
	//
	// **Ringtone field*	- valid values:
	//
	// - AXNRing_AB: The ringtone Settings before Number A dials Number X and the call is connected (with a callback number) in AXN mode.
	//
	// - AXNRing_A: The ringtone Settings before Number A dials Number X and the call is connected (without a callback number) in AXN mode.
	//
	// - AXNRing_N: The ringtone Settings before Number N dials Number X and the call is connected in AXN mode.
	//
	// - AXNRing_A_PICKUP: The ringtone Settings after Number A dials Number X and the call is connected in AXN mode.
	//
	// - AXNRing_N_PICKUP: The ringtone Settings after Number N dials Number X and the call is connected in AXN mode.
	//
	// **CRBT encoding**:
	//
	// You can obtain the corporate CRBT encoding on the [Corporate CRBT Management](https://dyplsnext.console.aliyun.com/number-pool/ringtone?resType=1) page in the Privacy Protection console.
	//
	// **Example**:
	//
	// To set the ringtone played before Number A dials Number X and the call is connected (without a callback number), enter "AXNRing_A": "1000****1".
	//
	// > The corporate CRBT configured at the attachment level takes precedence. If it is not set or does not take effect, the CRBT configured at the number pool level is used.
	//
	// example:
	//
	// {"AXNRing_N":"1000****1","AXNRing_A":"1000****1"}
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
