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
	// The list of AXG bindings.
	//
	// This parameter is required.
	AxgBindList []*BindBatchAxgRequestAxgBindList `json:"AxgBindList,omitempty" xml:"AxgBindList,omitempty" type:"Repeated"`
	OwnerId     *int64                            `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The key of the number pool.
	//
	// You can log on to the Phone Number Protection console and obtain the key on the [Number Pool Management](https://dyplsnext.console.aliyun.com/number-pool) page.
	//
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
	// The ASR model ID. Valid values:
	//
	// - The ID of the telephony model for Mandarin: 7ee372834d2f4cc7ac0d0ab2d0ae1aac
	//
	// - The ID of the telephony model for Cantonese: 9b4afe3d224e4fa3bf963697113e2f77
	//
	// - The ID of the telephony model for local accents: 3655a6cb94764145a1f9a38a3c3b3b4b
	//
	// example:
	//
	// 7ee372834d2f4cc7ac0d0ab2d0ae1aac
	ASRModelId *string `json:"ASRModelId,omitempty" xml:"ASRModelId,omitempty"`
	// The automatic speech recognition (ASR) status. Valid values:
	//
	// - **false*	- (default): The ASR feature is disabled.
	//
	// - **true**: The ASR feature is enabled.
	//
	// example:
	//
	// true
	ASRStatus *bool `json:"ASRStatus,omitempty" xml:"ASRStatus,omitempty"`
	// Resets the number display logic for the binding. Set the value to **1**. This indicates that the intermediate number (Number X) is displayed to both the calling and called parties.
	//
	// > Due to carrier regulations, the feature of displaying real numbers is not supported.
	//
	// example:
	//
	// 1
	CallDisplayType *int32 `json:"CallDisplayType,omitempty" xml:"CallDisplayType,omitempty"`
	// The one-way call restriction. Valid values:
	//
	// - **CONTROL_AX_DISABLE**: Number A cannot call Number X.
	//
	// - **CONTROL_BX_DISABLE**: Number B cannot call Number X.
	//
	// example:
	//
	// CONTROL_AX_DISABLE
	CallRestrict *string `json:"CallRestrict,omitempty" xml:"CallRestrict,omitempty"`
	// The city to which Number X belongs.
	//
	// > - If no numbers are available in the specified city, or if this parameter is not specified, a number from another city in the current number pool is randomly selected and used as Number X.
	//
	// >
	//
	// > - If **Number X Assignment Mode*	- is set to **Strict Match Mode*	- and no eligible number is found, an allocation error is reported.
	//
	// example:
	//
	// 北京
	ExpectCity *string `json:"ExpectCity,omitempty" xml:"ExpectCity,omitempty"`
	// The expiration time of the binding. The value is accurate to seconds.
	//
	// > - The expiration time must be more than 1 minute later than the current time.
	//
	// >
	//
	// > - This parameter is required.
	//
	// example:
	//
	// 2022-07-11 01:05:15
	Expiration *string `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	// The group ID (GID) in the AXG binding.
	//
	// You can view the GID in the following ways:
	//
	// - Log on to the Phone Number Protection console. On the [Number Pool Management](https://dyplsnext.console.aliyun.com/number-pool) page, select **AXG Privacy Number*	- and click **G Number Group Management*	- to view the GIDs of all G number groups.
	//
	// - If you use the [CreateAxgGroup](https://help.aliyun.com/document_detail/110250.html) operation to create a G number group, the value of the **GroupId*	- parameter in the response is the GID.
	//
	// > The G number group must contain at least one phone number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 200001234****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// Specifies whether to record all calls made in the binding. Valid values:
	//
	// - **true**: Call recording is enabled.
	//
	// - **false*	- (default): Call recording is disabled.
	//
	// example:
	//
	// false
	IsRecordingEnabled *bool `json:"IsRecordingEnabled,omitempty" xml:"IsRecordingEnabled,omitempty"`
	// The extension field. This parameter is carried in call record messages.
	//
	// example:
	//
	// 1***********0fd
	OutId *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	// The external business ID.
	//
	// example:
	//
	// 20***************677
	OutOrderId *string `json:"OutOrderId,omitempty" xml:"OutOrderId,omitempty"`
	// Number A in the AXG binding. Number A can be a mobile phone number or a landline number. If you use a landline number, add the area code, and do not add a hyphen between the area code and the landline number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 139****0000
	PhoneNoA *string `json:"PhoneNoA,omitempty" xml:"PhoneNoA,omitempty"`
	// Number B in the AXG binding.
	//
	// Number B can be a mobile phone number or a landline number. If you use a landline number, add the area code, and do not add a hyphen between the area code and the landline number.
	//
	// > - When Number A calls Number X, the call is transferred to Number B by default. If Number B is not specified, the call fails.
	//
	// >
	//
	// > - When Number A calls Number X, if **A Callback Mode*	- is set to **Default Mode**, the call is transferred to Number B. Otherwise, the call fails.
	//
	// >
	//
	// > - To modify **A Callback Mode**, on the [Number Pool Management](https://dyplsnext.console.aliyun.com/number-pool?resType=1) page, select **AXG Privacy Number*	- and choose **Pool Management*	- > **A Callback Mode*	- > **Adjust Configuration**.
	//
	// >
	//
	// > - To update Number B, call the [UpdateSubscription](https://help.aliyun.com/zh/pnp/developer-reference/api-dyplsapi-2017-05-25-updatesubscription?spm=a2c4g.11186623.0.i16) operation.
	//
	// example:
	//
	// 138****0000
	PhoneNoB *string `json:"PhoneNoB,omitempty" xml:"PhoneNoB,omitempty"`
	// Number X in the AXG binding.
	//
	// Number X must be a [purchased privacy number](https://dyplsnext.console.aliyun.com/numberbuy?spm=a2c4g.11186623.0.0.7e6449bbtiHP8m). You can log on to the Phone Number Protection console, go to the [Number Management](https://dyplsnext.console.aliyun.com/number/list) page, and then select **AXG Privacy Number*	- to obtain Number X.
	//
	// > If you do not specify Number X, a random number is selected from the number pool based on the **ExpectCity*	- parameter and used as Number X.
	//
	// example:
	//
	// 139****1111
	PhoneNoX *string `json:"PhoneNoX,omitempty" xml:"PhoneNoX,omitempty"`
	// The corporate ringback tone (CRBT) that is played when Number X in an AXG binding is called.
	//
	// You must specify the **ringtone field*	- and the **CRBT code**.
	//
	// Valid values for the **ringtone field**:
	//
	// - `AXGRing_AB`: the ringtone for calls from Number A to Number X (with a callback number).
	//
	// - `AXGRing_A`: the ringtone for calls from Number A to Number X (without a callback number).
	//
	// - `AXGRing_G`: the ringtone for calls from a number in Group G to Number X.
	//
	// **CRBT code**:
	//
	// You can log on to the Phone Number Protection console and obtain the CRBT code on the [Corporate CRBT Management](https://dyplsnext.console.aliyun.com/number-pool/ringtone?resType=1) page.
	//
	// **Example**:
	//
	// To set the ringtone for calls from Number A to Number X (without a callback number), specify `"AXGRing_A": "1000****1"`.
	//
	// > The CRBT that is set at the binding level takes precedence over the CRBT that is set at the number pool level.
	//
	// example:
	//
	// {"AXGRing_AB":"1000****2","AXGRing_A":"1000****1"}
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
