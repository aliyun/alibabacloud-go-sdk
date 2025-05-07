// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddAxnTrackNoRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The private number in the AXN binding, that is, phone number X.
	//
	// You can call the [BindAxn](https://help.aliyun.com/document_detail/110258.html) operation to obtain the value of PhoneNoX.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1700000****
	PhoneNoX *string `json:"PhoneNoX,omitempty" xml:"PhoneNoX,omitempty"`
	// The key of the phone number pool. Log on to the [Phone Number Protection console](https://dypls.console.aliyun.com/dypls.htm#/account) and view the key of the phone number pool on the **Number Pool Management*	- page.
	//
	// This parameter is required.
	//
	// example:
	//
	// FC2235****
	PoolKey              *string `json:"PoolKey,omitempty" xml:"PoolKey,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The binding ID.
	//
	// You can call the [BindAxn](https://help.aliyun.com/document_detail/110258.html) operation to obtain the value of SubsId.
	//
	// This parameter is required.
	//
	// example:
	//
	// 15678890****
	SubsId *string `json:"SubsId,omitempty" xml:"SubsId,omitempty"`
	// The tracking number.
	//
	// This parameter is required.
	//
	// example:
	//
	// abcde*****
	TrackNo *string `json:"trackNo,omitempty" xml:"trackNo,omitempty"`
}

func (s AddAxnTrackNoRequest) String() string {
	return tea.Prettify(s)
}

func (s AddAxnTrackNoRequest) GoString() string {
	return s.String()
}

func (s *AddAxnTrackNoRequest) SetOwnerId(v int64) *AddAxnTrackNoRequest {
	s.OwnerId = &v
	return s
}

func (s *AddAxnTrackNoRequest) SetPhoneNoX(v string) *AddAxnTrackNoRequest {
	s.PhoneNoX = &v
	return s
}

func (s *AddAxnTrackNoRequest) SetPoolKey(v string) *AddAxnTrackNoRequest {
	s.PoolKey = &v
	return s
}

func (s *AddAxnTrackNoRequest) SetResourceOwnerAccount(v string) *AddAxnTrackNoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AddAxnTrackNoRequest) SetResourceOwnerId(v int64) *AddAxnTrackNoRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AddAxnTrackNoRequest) SetSubsId(v string) *AddAxnTrackNoRequest {
	s.SubsId = &v
	return s
}

func (s *AddAxnTrackNoRequest) SetTrackNo(v string) *AddAxnTrackNoRequest {
	s.TrackNo = &v
	return s
}

type AddAxnTrackNoResponseBody struct {
	// The response code.
	//
	// 	- The value OK indicates that the request was successful.
	//
	// 	- Other values indicate that the request failed. For more information, see [Error codes](https://help.aliyun.com/document_detail/109196.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1A4CADEF-8516-5890-9177-A1A4D97F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddAxnTrackNoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddAxnTrackNoResponseBody) GoString() string {
	return s.String()
}

func (s *AddAxnTrackNoResponseBody) SetCode(v string) *AddAxnTrackNoResponseBody {
	s.Code = &v
	return s
}

func (s *AddAxnTrackNoResponseBody) SetMessage(v string) *AddAxnTrackNoResponseBody {
	s.Message = &v
	return s
}

func (s *AddAxnTrackNoResponseBody) SetRequestId(v string) *AddAxnTrackNoResponseBody {
	s.RequestId = &v
	return s
}

type AddAxnTrackNoResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddAxnTrackNoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddAxnTrackNoResponse) String() string {
	return tea.Prettify(s)
}

func (s AddAxnTrackNoResponse) GoString() string {
	return s.String()
}

func (s *AddAxnTrackNoResponse) SetHeaders(v map[string]*string) *AddAxnTrackNoResponse {
	s.Headers = v
	return s
}

func (s *AddAxnTrackNoResponse) SetStatusCode(v int32) *AddAxnTrackNoResponse {
	s.StatusCode = &v
	return s
}

func (s *AddAxnTrackNoResponse) SetBody(v *AddAxnTrackNoResponseBody) *AddAxnTrackNoResponse {
	s.Body = v
	return s
}

type AddSecretBlacklistRequest struct {
	// The numbers in the blacklist. A point-to-point blacklist has a pair of numbers separated by a colon (:). A number pool blacklist has only one single number.
	//
	// >  The asterisks (\\*) in the sample value are not wildcards.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1825638****:1825248****
	BlackNo *string `json:"BlackNo,omitempty" xml:"BlackNo,omitempty"`
	// The blacklist type. Valid values:
	//
	// 	- **POINT_TO_POINT_BLACK**: point-to-point blacklist.
	//
	// 	- **PARTNER_GLOBAL_NUMBER_BLACK**: number pool blacklist.
	//
	// This parameter is required.
	//
	// example:
	//
	// POINT_TO_POINT_BLACK
	BlackType *string `json:"BlackType,omitempty" xml:"BlackType,omitempty"`
	// The key of the phone number pool.
	//
	// Log on to the [Phone Number Protection console](https://dypls.console.aliyun.com/dypls.htm#/account) and view the key of the phone number pool on the Number Pool Management page.
	//
	// This parameter is required.
	//
	// example:
	//
	// FC2235****
	PoolKey *string `json:"PoolKey,omitempty" xml:"PoolKey,omitempty"`
	// The blacklist remarks.
	//
	// example:
	//
	// Customer service feedback
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The control on the call direction.
	//
	// 	- **PHONEA_REJECT**: The first number in the blacklist can be used to call the second number, but the second number cannot be used to call the first number.
	//
	// 	- **PHONEB_REJECT**: The first number in the blacklist cannot be used to call the second number, but the second number can be used to call the first number.
	//
	// 	- If this parameter is left empty, the two numbers in the blacklist cannot be used to call each other.
	//
	// >  This parameter is available only for a point-to-point blacklist.
	//
	// Valid values:
	//
	// 	- DUPLEX_REJECT
	//
	// 	- PHONEA_REJECT
	//
	// 	- PHONEB_REJECT
	//
	// example:
	//
	// PHONEA_REJECT
	WayControl *string `json:"WayControl,omitempty" xml:"WayControl,omitempty"`
}

func (s AddSecretBlacklistRequest) String() string {
	return tea.Prettify(s)
}

func (s AddSecretBlacklistRequest) GoString() string {
	return s.String()
}

func (s *AddSecretBlacklistRequest) SetBlackNo(v string) *AddSecretBlacklistRequest {
	s.BlackNo = &v
	return s
}

func (s *AddSecretBlacklistRequest) SetBlackType(v string) *AddSecretBlacklistRequest {
	s.BlackType = &v
	return s
}

func (s *AddSecretBlacklistRequest) SetPoolKey(v string) *AddSecretBlacklistRequest {
	s.PoolKey = &v
	return s
}

func (s *AddSecretBlacklistRequest) SetRemark(v string) *AddSecretBlacklistRequest {
	s.Remark = &v
	return s
}

func (s *AddSecretBlacklistRequest) SetWayControl(v string) *AddSecretBlacklistRequest {
	s.WayControl = &v
	return s
}

type AddSecretBlacklistResponseBody struct {
	// The response code.
	//
	// 	- The value OK indicates that the request was successful.
	//
	// 	- Other status codes indicate that the request failed. For more information, see [Error codes](https://help.aliyun.com/document_detail/109196.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F655A8D5-B967-440B-8683-DAD6FF8DE990
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddSecretBlacklistResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddSecretBlacklistResponseBody) GoString() string {
	return s.String()
}

func (s *AddSecretBlacklistResponseBody) SetCode(v string) *AddSecretBlacklistResponseBody {
	s.Code = &v
	return s
}

func (s *AddSecretBlacklistResponseBody) SetMessage(v string) *AddSecretBlacklistResponseBody {
	s.Message = &v
	return s
}

func (s *AddSecretBlacklistResponseBody) SetRequestId(v string) *AddSecretBlacklistResponseBody {
	s.RequestId = &v
	return s
}

type AddSecretBlacklistResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddSecretBlacklistResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddSecretBlacklistResponse) String() string {
	return tea.Prettify(s)
}

func (s AddSecretBlacklistResponse) GoString() string {
	return s.String()
}

func (s *AddSecretBlacklistResponse) SetHeaders(v map[string]*string) *AddSecretBlacklistResponse {
	s.Headers = v
	return s
}

func (s *AddSecretBlacklistResponse) SetStatusCode(v int32) *AddSecretBlacklistResponse {
	s.StatusCode = &v
	return s
}

func (s *AddSecretBlacklistResponse) SetBody(v *AddSecretBlacklistResponseBody) *AddSecretBlacklistResponse {
	s.Body = v
	return s
}

type BindAXBCallRequest struct {
	// authId绑定关系BX唯一id
	//
	// This parameter is required.
	//
	// example:
	//
	// 12353
	AuthId *string `json:"AuthId,omitempty" xml:"AuthId,omitempty"`
	// 客户uid
	//
	// example:
	//
	// -
	CallerParentId *int64 `json:"CallerParentId,omitempty" xml:"CallerParentId,omitempty"`
	// 号码池key
	//
	// This parameter is required.
	//
	// example:
	//
	// FC5**********************a1a
	CustomerPoolKey *string `json:"CustomerPoolKey,omitempty" xml:"CustomerPoolKey,omitempty"`
	// 绑定关系过期失效时间： 取值必须大于0； 单位：秒，ct_time + expiration = 自动解绑时间
	//
	// This parameter is required.
	//
	// example:
	//
	// 7200
	Expiration *int64 `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	OwnerId    *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 请求去重ID, reqId最大长度为20位,接入方需要保持原子性
	//
	// This parameter is required.
	//
	// example:
	//
	// 564**********879
	ReqId                *string `json:"ReqId,omitempty" xml:"ReqId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 客户A号码
	//
	// This parameter is required.
	//
	// example:
	//
	// 13*******43
	TelA *string `json:"TelA,omitempty" xml:"TelA,omitempty"`
	// 客户自定义参数回调带回
	//
	// example:
	//
	// 000
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s BindAXBCallRequest) String() string {
	return tea.Prettify(s)
}

func (s BindAXBCallRequest) GoString() string {
	return s.String()
}

func (s *BindAXBCallRequest) SetAuthId(v string) *BindAXBCallRequest {
	s.AuthId = &v
	return s
}

func (s *BindAXBCallRequest) SetCallerParentId(v int64) *BindAXBCallRequest {
	s.CallerParentId = &v
	return s
}

func (s *BindAXBCallRequest) SetCustomerPoolKey(v string) *BindAXBCallRequest {
	s.CustomerPoolKey = &v
	return s
}

func (s *BindAXBCallRequest) SetExpiration(v int64) *BindAXBCallRequest {
	s.Expiration = &v
	return s
}

func (s *BindAXBCallRequest) SetOwnerId(v int64) *BindAXBCallRequest {
	s.OwnerId = &v
	return s
}

func (s *BindAXBCallRequest) SetReqId(v string) *BindAXBCallRequest {
	s.ReqId = &v
	return s
}

func (s *BindAXBCallRequest) SetResourceOwnerAccount(v string) *BindAXBCallRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *BindAXBCallRequest) SetResourceOwnerId(v int64) *BindAXBCallRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *BindAXBCallRequest) SetTelA(v string) *BindAXBCallRequest {
	s.TelA = &v
	return s
}

func (s *BindAXBCallRequest) SetUserData(v string) *BindAXBCallRequest {
	s.UserData = &v
	return s
}

type BindAXBCallResponseBody struct {
	// example:
	//
	// 0
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// 返回状态码 0000表示成功 其他表示失败
	//
	// example:
	//
	// 0000
	Code *string                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *BindAXBCallResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// 返回信息
	//
	// example:
	//
	// 成功
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 返回是否成功 true  表示成功 false表示失败
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BindAXBCallResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BindAXBCallResponseBody) GoString() string {
	return s.String()
}

func (s *BindAXBCallResponseBody) SetAccessDeniedDetail(v string) *BindAXBCallResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *BindAXBCallResponseBody) SetCode(v string) *BindAXBCallResponseBody {
	s.Code = &v
	return s
}

func (s *BindAXBCallResponseBody) SetData(v *BindAXBCallResponseBodyData) *BindAXBCallResponseBody {
	s.Data = v
	return s
}

func (s *BindAXBCallResponseBody) SetMessage(v string) *BindAXBCallResponseBody {
	s.Message = &v
	return s
}

func (s *BindAXBCallResponseBody) SetSuccess(v bool) *BindAXBCallResponseBody {
	s.Success = &v
	return s
}

type BindAXBCallResponseBodyData struct {
	// 绑定关系ID
	//
	// example:
	//
	// 476567566
	BindId *string `json:"BindId,omitempty" xml:"BindId,omitempty"`
}

func (s BindAXBCallResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s BindAXBCallResponseBodyData) GoString() string {
	return s.String()
}

func (s *BindAXBCallResponseBodyData) SetBindId(v string) *BindAXBCallResponseBodyData {
	s.BindId = &v
	return s
}

type BindAXBCallResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BindAXBCallResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BindAXBCallResponse) String() string {
	return tea.Prettify(s)
}

func (s BindAXBCallResponse) GoString() string {
	return s.String()
}

func (s *BindAXBCallResponse) SetHeaders(v map[string]*string) *BindAXBCallResponse {
	s.Headers = v
	return s
}

func (s *BindAXBCallResponse) SetStatusCode(v int32) *BindAXBCallResponse {
	s.StatusCode = &v
	return s
}

func (s *BindAXBCallResponse) SetBody(v *BindAXBCallResponseBody) *BindAXBCallResponse {
	s.Body = v
	return s
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
	return tea.Prettify(s)
}

func (s BindAxbRequest) GoString() string {
	return s.String()
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

type BindAxbResponseBody struct {
	// The response code.
	//
	// 	- The value OK indicates that the request was successful.
	//
	// 	- Other status codes indicate that the request failed. For more information, see [Error codes](https://help.aliyun.com/document_detail/109196.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9297B722-A016-43FB-B51A-E54050D9369D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information returned after the phone numbers were bound.
	SecretBindDTO *BindAxbResponseBodySecretBindDTO `json:"SecretBindDTO,omitempty" xml:"SecretBindDTO,omitempty" type:"Struct"`
}

func (s BindAxbResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BindAxbResponseBody) GoString() string {
	return s.String()
}

func (s *BindAxbResponseBody) SetCode(v string) *BindAxbResponseBody {
	s.Code = &v
	return s
}

func (s *BindAxbResponseBody) SetMessage(v string) *BindAxbResponseBody {
	s.Message = &v
	return s
}

func (s *BindAxbResponseBody) SetRequestId(v string) *BindAxbResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindAxbResponseBody) SetSecretBindDTO(v *BindAxbResponseBodySecretBindDTO) *BindAxbResponseBody {
	s.SecretBindDTO = v
	return s
}

type BindAxbResponseBodySecretBindDTO struct {
	// The extension of the phone number.
	//
	// >  The BindAxb operation does not involve an extension. Ignore this parameter.
	//
	// example:
	//
	// 130
	Extension *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// The private number, that is, phone number X.
	//
	// example:
	//
	// 139****0000
	SecretNo *string `json:"SecretNo,omitempty" xml:"SecretNo,omitempty"`
	// The binding ID.
	//
	// example:
	//
	// 1**************3
	SubsId *string `json:"SubsId,omitempty" xml:"SubsId,omitempty"`
}

func (s BindAxbResponseBodySecretBindDTO) String() string {
	return tea.Prettify(s)
}

func (s BindAxbResponseBodySecretBindDTO) GoString() string {
	return s.String()
}

func (s *BindAxbResponseBodySecretBindDTO) SetExtension(v string) *BindAxbResponseBodySecretBindDTO {
	s.Extension = &v
	return s
}

func (s *BindAxbResponseBodySecretBindDTO) SetSecretNo(v string) *BindAxbResponseBodySecretBindDTO {
	s.SecretNo = &v
	return s
}

func (s *BindAxbResponseBodySecretBindDTO) SetSubsId(v string) *BindAxbResponseBodySecretBindDTO {
	s.SubsId = &v
	return s
}

type BindAxbResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BindAxbResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BindAxbResponse) String() string {
	return tea.Prettify(s)
}

func (s BindAxbResponse) GoString() string {
	return s.String()
}

func (s *BindAxbResponse) SetHeaders(v map[string]*string) *BindAxbResponse {
	s.Headers = v
	return s
}

func (s *BindAxbResponse) SetStatusCode(v int32) *BindAxbResponse {
	s.StatusCode = &v
	return s
}

func (s *BindAxbResponse) SetBody(v *BindAxbResponseBody) *BindAxbResponse {
	s.Body = v
	return s
}

type BindAxbFixedLineRequest struct {
	// 主叫侧放音编码，AXB业务时必须设置。  放音编码必须包含3个场景的编码。按照“B->X,A->X,其他号码->X”的顺序填写编码，编码之间以逗号分隔。  比如：“1,2,3”表示B->X放音编号为1，A->X放音编号为2， 其他号码->X放音编号为3
	//
	// This parameter is required.
	//
	// example:
	//
	// 1,2,3
	Anucode *string `json:"Anucode,omitempty" xml:"Anucode,omitempty"`
	// 接通后被叫侧放音编码,接通后被叫侧放音编码  被叫放音编码必须包含2个场景的编码。按照“A被叫,B被叫”的顺序填写编码，编码之间以逗号分隔。  比如：“1,2”表示A号码为被叫侧接听时的放音编号为1，B号码为被叫侧接听时的放音编号为2。
	//
	// example:
	//
	// 示例值示例值
	Anucodecalled *string `json:"Anucodecalled,omitempty" xml:"Anucodecalled,omitempty"`
	// 应用id
	//
	// This parameter is required.
	//
	// example:
	//
	// ALPT_1234
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 区号,去掉“0” 例如：北京（10）；在平台分配X号码模式中，平台从号码池中分配该地区的X号码，避免产生呼叫长途费。
	//
	// example:
	//
	// 10
	Areacode *string `json:"Areacode,omitempty" xml:"Areacode,omitempty"`
	// 绑定类型，值如下： mode101：客户携带X号码 mode102：平台分配X号码
	//
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值
	BindType *string `json:"BindType,omitempty" xml:"BindType,omitempty"`
	// 过期时间,单位：秒，必须为数字 0：不会自动解绑 非0：自动解绑周期
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	Expiration *string `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	// 扩展参数
	Extra *BindAxbFixedLineRequestExtra `json:"Extra,omitempty" xml:"Extra,omitempty" type:"Struct"`
	// 消息请求标识
	//
	// This parameter is required.
	//
	// example:
	//
	// 15433678436
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 接入商自有字段，不能超过100个长度
	//
	// example:
	//
	// remark
	Remark               *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 绑定时间，格式为yyyyMMddHHmmss。时间采用北京时间，24小时制。
	//
	// This parameter is required.
	//
	// example:
	//
	// 20161114143116
	Subts *string `json:"Subts,omitempty" xml:"Subts,omitempty"`
	// 接通后主叫侧的放音编码,接通后主叫侧放音编码  接通后主叫侧放音编码必须包含2个场景的编码。按照“A被叫,B被叫”的顺序填写编码，编码之间以逗号分隔。  比如：“1,2”表示A号码为被叫时主叫侧的放音编号为1，B号码为被叫时主叫侧的放音编号为2。
	//
	// example:
	//
	// 示例值示例值示例值
	TAnucodeConnect *string `json:"TAnucodeConnect,omitempty" xml:"TAnucodeConnect,omitempty"`
	// 真实号码
	//
	// This parameter is required.
	//
	// example:
	//
	// 18456713271
	TelA *string `json:"TelA,omitempty" xml:"TelA,omitempty"`
	// 对端号码
	//
	// This parameter is required.
	//
	// example:
	//
	// 18971362645
	TelB *string `json:"TelB,omitempty" xml:"TelB,omitempty"`
	// X号码； 平台分配号码模式下，该参数可不带，系统忽略该参数  格式为yyyyMMddHHmmss。时间采用北京时间，24小时制。
	//
	// example:
	//
	// 19767562345
	TelX *string `json:"TelX,omitempty" xml:"TelX,omitempty"`
	// 业务时间戳，格式yyyyMMddHHmmssSSS，时间采用北京时间，24小时制，精确至毫秒
	//
	// This parameter is required.
	//
	// example:
	//
	// 20161114143116001
	Ts *string `json:"Ts,omitempty" xml:"Ts,omitempty"`
}

func (s BindAxbFixedLineRequest) String() string {
	return tea.Prettify(s)
}

func (s BindAxbFixedLineRequest) GoString() string {
	return s.String()
}

func (s *BindAxbFixedLineRequest) SetAnucode(v string) *BindAxbFixedLineRequest {
	s.Anucode = &v
	return s
}

func (s *BindAxbFixedLineRequest) SetAnucodecalled(v string) *BindAxbFixedLineRequest {
	s.Anucodecalled = &v
	return s
}

func (s *BindAxbFixedLineRequest) SetAppId(v string) *BindAxbFixedLineRequest {
	s.AppId = &v
	return s
}

func (s *BindAxbFixedLineRequest) SetAreacode(v string) *BindAxbFixedLineRequest {
	s.Areacode = &v
	return s
}

func (s *BindAxbFixedLineRequest) SetBindType(v string) *BindAxbFixedLineRequest {
	s.BindType = &v
	return s
}

func (s *BindAxbFixedLineRequest) SetExpiration(v string) *BindAxbFixedLineRequest {
	s.Expiration = &v
	return s
}

func (s *BindAxbFixedLineRequest) SetExtra(v *BindAxbFixedLineRequestExtra) *BindAxbFixedLineRequest {
	s.Extra = v
	return s
}

func (s *BindAxbFixedLineRequest) SetOrderId(v string) *BindAxbFixedLineRequest {
	s.OrderId = &v
	return s
}

func (s *BindAxbFixedLineRequest) SetOwnerId(v int64) *BindAxbFixedLineRequest {
	s.OwnerId = &v
	return s
}

func (s *BindAxbFixedLineRequest) SetRemark(v string) *BindAxbFixedLineRequest {
	s.Remark = &v
	return s
}

func (s *BindAxbFixedLineRequest) SetResourceOwnerAccount(v string) *BindAxbFixedLineRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *BindAxbFixedLineRequest) SetResourceOwnerId(v int64) *BindAxbFixedLineRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *BindAxbFixedLineRequest) SetSubts(v string) *BindAxbFixedLineRequest {
	s.Subts = &v
	return s
}

func (s *BindAxbFixedLineRequest) SetTAnucodeConnect(v string) *BindAxbFixedLineRequest {
	s.TAnucodeConnect = &v
	return s
}

func (s *BindAxbFixedLineRequest) SetTelA(v string) *BindAxbFixedLineRequest {
	s.TelA = &v
	return s
}

func (s *BindAxbFixedLineRequest) SetTelB(v string) *BindAxbFixedLineRequest {
	s.TelB = &v
	return s
}

func (s *BindAxbFixedLineRequest) SetTelX(v string) *BindAxbFixedLineRequest {
	s.TelX = &v
	return s
}

func (s *BindAxbFixedLineRequest) SetTs(v string) *BindAxbFixedLineRequest {
	s.Ts = &v
	return s
}

type BindAxbFixedLineRequestExtra struct {
	// 录音控制，默认是0（不开通录音功能）。 0：不录音 1：录音
	//
	// example:
	//
	// 示例值示例值
	Callrecording *string `json:"Callrecording,omitempty" xml:"Callrecording,omitempty"`
}

func (s BindAxbFixedLineRequestExtra) String() string {
	return tea.Prettify(s)
}

func (s BindAxbFixedLineRequestExtra) GoString() string {
	return s.String()
}

func (s *BindAxbFixedLineRequestExtra) SetCallrecording(v string) *BindAxbFixedLineRequestExtra {
	s.Callrecording = &v
	return s
}

type BindAxbFixedLineShrinkRequest struct {
	// 主叫侧放音编码，AXB业务时必须设置。  放音编码必须包含3个场景的编码。按照“B->X,A->X,其他号码->X”的顺序填写编码，编码之间以逗号分隔。  比如：“1,2,3”表示B->X放音编号为1，A->X放音编号为2， 其他号码->X放音编号为3
	//
	// This parameter is required.
	//
	// example:
	//
	// 1,2,3
	Anucode *string `json:"Anucode,omitempty" xml:"Anucode,omitempty"`
	// 接通后被叫侧放音编码,接通后被叫侧放音编码  被叫放音编码必须包含2个场景的编码。按照“A被叫,B被叫”的顺序填写编码，编码之间以逗号分隔。  比如：“1,2”表示A号码为被叫侧接听时的放音编号为1，B号码为被叫侧接听时的放音编号为2。
	//
	// example:
	//
	// 示例值示例值
	Anucodecalled *string `json:"Anucodecalled,omitempty" xml:"Anucodecalled,omitempty"`
	// 应用id
	//
	// This parameter is required.
	//
	// example:
	//
	// ALPT_1234
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 区号,去掉“0” 例如：北京（10）；在平台分配X号码模式中，平台从号码池中分配该地区的X号码，避免产生呼叫长途费。
	//
	// example:
	//
	// 10
	Areacode *string `json:"Areacode,omitempty" xml:"Areacode,omitempty"`
	// 绑定类型，值如下： mode101：客户携带X号码 mode102：平台分配X号码
	//
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值
	BindType *string `json:"BindType,omitempty" xml:"BindType,omitempty"`
	// 过期时间,单位：秒，必须为数字 0：不会自动解绑 非0：自动解绑周期
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	Expiration *string `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	// 扩展参数
	ExtraShrink *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	// 消息请求标识
	//
	// This parameter is required.
	//
	// example:
	//
	// 15433678436
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 接入商自有字段，不能超过100个长度
	//
	// example:
	//
	// remark
	Remark               *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 绑定时间，格式为yyyyMMddHHmmss。时间采用北京时间，24小时制。
	//
	// This parameter is required.
	//
	// example:
	//
	// 20161114143116
	Subts *string `json:"Subts,omitempty" xml:"Subts,omitempty"`
	// 接通后主叫侧的放音编码,接通后主叫侧放音编码  接通后主叫侧放音编码必须包含2个场景的编码。按照“A被叫,B被叫”的顺序填写编码，编码之间以逗号分隔。  比如：“1,2”表示A号码为被叫时主叫侧的放音编号为1，B号码为被叫时主叫侧的放音编号为2。
	//
	// example:
	//
	// 示例值示例值示例值
	TAnucodeConnect *string `json:"TAnucodeConnect,omitempty" xml:"TAnucodeConnect,omitempty"`
	// 真实号码
	//
	// This parameter is required.
	//
	// example:
	//
	// 18456713271
	TelA *string `json:"TelA,omitempty" xml:"TelA,omitempty"`
	// 对端号码
	//
	// This parameter is required.
	//
	// example:
	//
	// 18971362645
	TelB *string `json:"TelB,omitempty" xml:"TelB,omitempty"`
	// X号码； 平台分配号码模式下，该参数可不带，系统忽略该参数  格式为yyyyMMddHHmmss。时间采用北京时间，24小时制。
	//
	// example:
	//
	// 19767562345
	TelX *string `json:"TelX,omitempty" xml:"TelX,omitempty"`
	// 业务时间戳，格式yyyyMMddHHmmssSSS，时间采用北京时间，24小时制，精确至毫秒
	//
	// This parameter is required.
	//
	// example:
	//
	// 20161114143116001
	Ts *string `json:"Ts,omitempty" xml:"Ts,omitempty"`
}

func (s BindAxbFixedLineShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s BindAxbFixedLineShrinkRequest) GoString() string {
	return s.String()
}

func (s *BindAxbFixedLineShrinkRequest) SetAnucode(v string) *BindAxbFixedLineShrinkRequest {
	s.Anucode = &v
	return s
}

func (s *BindAxbFixedLineShrinkRequest) SetAnucodecalled(v string) *BindAxbFixedLineShrinkRequest {
	s.Anucodecalled = &v
	return s
}

func (s *BindAxbFixedLineShrinkRequest) SetAppId(v string) *BindAxbFixedLineShrinkRequest {
	s.AppId = &v
	return s
}

func (s *BindAxbFixedLineShrinkRequest) SetAreacode(v string) *BindAxbFixedLineShrinkRequest {
	s.Areacode = &v
	return s
}

func (s *BindAxbFixedLineShrinkRequest) SetBindType(v string) *BindAxbFixedLineShrinkRequest {
	s.BindType = &v
	return s
}

func (s *BindAxbFixedLineShrinkRequest) SetExpiration(v string) *BindAxbFixedLineShrinkRequest {
	s.Expiration = &v
	return s
}

func (s *BindAxbFixedLineShrinkRequest) SetExtraShrink(v string) *BindAxbFixedLineShrinkRequest {
	s.ExtraShrink = &v
	return s
}

func (s *BindAxbFixedLineShrinkRequest) SetOrderId(v string) *BindAxbFixedLineShrinkRequest {
	s.OrderId = &v
	return s
}

func (s *BindAxbFixedLineShrinkRequest) SetOwnerId(v int64) *BindAxbFixedLineShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *BindAxbFixedLineShrinkRequest) SetRemark(v string) *BindAxbFixedLineShrinkRequest {
	s.Remark = &v
	return s
}

func (s *BindAxbFixedLineShrinkRequest) SetResourceOwnerAccount(v string) *BindAxbFixedLineShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *BindAxbFixedLineShrinkRequest) SetResourceOwnerId(v int64) *BindAxbFixedLineShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *BindAxbFixedLineShrinkRequest) SetSubts(v string) *BindAxbFixedLineShrinkRequest {
	s.Subts = &v
	return s
}

func (s *BindAxbFixedLineShrinkRequest) SetTAnucodeConnect(v string) *BindAxbFixedLineShrinkRequest {
	s.TAnucodeConnect = &v
	return s
}

func (s *BindAxbFixedLineShrinkRequest) SetTelA(v string) *BindAxbFixedLineShrinkRequest {
	s.TelA = &v
	return s
}

func (s *BindAxbFixedLineShrinkRequest) SetTelB(v string) *BindAxbFixedLineShrinkRequest {
	s.TelB = &v
	return s
}

func (s *BindAxbFixedLineShrinkRequest) SetTelX(v string) *BindAxbFixedLineShrinkRequest {
	s.TelX = &v
	return s
}

func (s *BindAxbFixedLineShrinkRequest) SetTs(v string) *BindAxbFixedLineShrinkRequest {
	s.Ts = &v
	return s
}

type BindAxbFixedLineResponseBody struct {
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// 响应码 0-成功
	//
	// example:
	//
	// 0
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 响应内容
	Data *BindAxbFixedLineResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// 响应消息
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 649E9EB5-9436-53CF-B41A-C4F0433212E7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 是否处理成功  true-成功
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BindAxbFixedLineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BindAxbFixedLineResponseBody) GoString() string {
	return s.String()
}

func (s *BindAxbFixedLineResponseBody) SetAccessDeniedDetail(v string) *BindAxbFixedLineResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *BindAxbFixedLineResponseBody) SetCode(v string) *BindAxbFixedLineResponseBody {
	s.Code = &v
	return s
}

func (s *BindAxbFixedLineResponseBody) SetData(v *BindAxbFixedLineResponseBodyData) *BindAxbFixedLineResponseBody {
	s.Data = v
	return s
}

func (s *BindAxbFixedLineResponseBody) SetMessage(v string) *BindAxbFixedLineResponseBody {
	s.Message = &v
	return s
}

func (s *BindAxbFixedLineResponseBody) SetRequestId(v string) *BindAxbFixedLineResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindAxbFixedLineResponseBody) SetSuccess(v bool) *BindAxbFixedLineResponseBody {
	s.Success = &v
	return s
}

type BindAxbFixedLineResponseBodyData struct {
	// 绑定id
	//
	// example:
	//
	// 示例值
	Subid *string `json:"Subid,omitempty" xml:"Subid,omitempty"`
	// X号码
	//
	// example:
	//
	// 示例值
	TelX *string `json:"TelX,omitempty" xml:"TelX,omitempty"`
}

func (s BindAxbFixedLineResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s BindAxbFixedLineResponseBodyData) GoString() string {
	return s.String()
}

func (s *BindAxbFixedLineResponseBodyData) SetSubid(v string) *BindAxbFixedLineResponseBodyData {
	s.Subid = &v
	return s
}

func (s *BindAxbFixedLineResponseBodyData) SetTelX(v string) *BindAxbFixedLineResponseBodyData {
	s.TelX = &v
	return s
}

type BindAxbFixedLineResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BindAxbFixedLineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BindAxbFixedLineResponse) String() string {
	return tea.Prettify(s)
}

func (s BindAxbFixedLineResponse) GoString() string {
	return s.String()
}

func (s *BindAxbFixedLineResponse) SetHeaders(v map[string]*string) *BindAxbFixedLineResponse {
	s.Headers = v
	return s
}

func (s *BindAxbFixedLineResponse) SetStatusCode(v int32) *BindAxbFixedLineResponse {
	s.StatusCode = &v
	return s
}

func (s *BindAxbFixedLineResponse) SetBody(v *BindAxbFixedLineResponseBody) *BindAxbFixedLineResponse {
	s.Body = v
	return s
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
	return tea.Prettify(s)
}

func (s BindAxgRequest) GoString() string {
	return s.String()
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

type BindAxgResponseBody struct {
	// The response code.
	//
	// 	- The value OK indicates that the request was successful.
	//
	// 	- Other values indicate that the request failed. For more information, see [Error codes](https://help.aliyun.com/document_detail/109196.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F655A8D5-B967-440B-8683-DAD6FF8DE990
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information returned after the phone numbers were bound.
	SecretBindDTO *BindAxgResponseBodySecretBindDTO `json:"SecretBindDTO,omitempty" xml:"SecretBindDTO,omitempty" type:"Struct"`
}

func (s BindAxgResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BindAxgResponseBody) GoString() string {
	return s.String()
}

func (s *BindAxgResponseBody) SetCode(v string) *BindAxgResponseBody {
	s.Code = &v
	return s
}

func (s *BindAxgResponseBody) SetMessage(v string) *BindAxgResponseBody {
	s.Message = &v
	return s
}

func (s *BindAxgResponseBody) SetRequestId(v string) *BindAxgResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindAxgResponseBody) SetSecretBindDTO(v *BindAxgResponseBodySecretBindDTO) *BindAxgResponseBody {
	s.SecretBindDTO = v
	return s
}

type BindAxgResponseBodySecretBindDTO struct {
	// The extension of the phone number.
	//
	// >  The BindAxg operation does not involve an extension. Ignore this parameter.
	//
	// example:
	//
	// 139****0000
	Extension *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// The private number, that is, phone number X.
	//
	// example:
	//
	// 139****0000
	SecretNo *string `json:"SecretNo,omitempty" xml:"SecretNo,omitempty"`
	// The binding ID.
	//
	// example:
	//
	// 1************3
	SubsId *string `json:"SubsId,omitempty" xml:"SubsId,omitempty"`
}

func (s BindAxgResponseBodySecretBindDTO) String() string {
	return tea.Prettify(s)
}

func (s BindAxgResponseBodySecretBindDTO) GoString() string {
	return s.String()
}

func (s *BindAxgResponseBodySecretBindDTO) SetExtension(v string) *BindAxgResponseBodySecretBindDTO {
	s.Extension = &v
	return s
}

func (s *BindAxgResponseBodySecretBindDTO) SetSecretNo(v string) *BindAxgResponseBodySecretBindDTO {
	s.SecretNo = &v
	return s
}

func (s *BindAxgResponseBodySecretBindDTO) SetSubsId(v string) *BindAxgResponseBodySecretBindDTO {
	s.SubsId = &v
	return s
}

type BindAxgResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BindAxgResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BindAxgResponse) String() string {
	return tea.Prettify(s)
}

func (s BindAxgResponse) GoString() string {
	return s.String()
}

func (s *BindAxgResponse) SetHeaders(v map[string]*string) *BindAxgResponse {
	s.Headers = v
	return s
}

func (s *BindAxgResponse) SetStatusCode(v int32) *BindAxgResponse {
	s.StatusCode = &v
	return s
}

func (s *BindAxgResponse) SetBody(v *BindAxgResponseBody) *BindAxgResponse {
	s.Body = v
	return s
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
	return tea.Prettify(s)
}

func (s BindAxnRequest) GoString() string {
	return s.String()
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

type BindAxnResponseBody struct {
	// The response code.
	//
	// 	- The value OK indicates that the request was successful.
	//
	// 	- Other values indicate that the request failed. For more information, see [Error codes](https://help.aliyun.com/document_detail/109196.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F655A8D5-B967-440B-8683-DAD6FF8DE990
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information returned after the phone numbers were bound.
	SecretBindDTO *BindAxnResponseBodySecretBindDTO `json:"SecretBindDTO,omitempty" xml:"SecretBindDTO,omitempty" type:"Struct"`
}

func (s BindAxnResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BindAxnResponseBody) GoString() string {
	return s.String()
}

func (s *BindAxnResponseBody) SetCode(v string) *BindAxnResponseBody {
	s.Code = &v
	return s
}

func (s *BindAxnResponseBody) SetMessage(v string) *BindAxnResponseBody {
	s.Message = &v
	return s
}

func (s *BindAxnResponseBody) SetRequestId(v string) *BindAxnResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindAxnResponseBody) SetSecretBindDTO(v *BindAxnResponseBodySecretBindDTO) *BindAxnResponseBody {
	s.SecretBindDTO = v
	return s
}

type BindAxnResponseBodySecretBindDTO struct {
	// The extension of the phone number.
	//
	// >  The BindAxn operation does not involve an extension. Ignore this parameter.
	//
	// example:
	//
	// 130
	Extension *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// The private number, that is, phone number X.
	//
	// example:
	//
	// 139****0000
	SecretNo *string `json:"SecretNo,omitempty" xml:"SecretNo,omitempty"`
	// The binding ID.
	//
	// example:
	//
	// 1***************3
	SubsId *string `json:"SubsId,omitempty" xml:"SubsId,omitempty"`
}

func (s BindAxnResponseBodySecretBindDTO) String() string {
	return tea.Prettify(s)
}

func (s BindAxnResponseBodySecretBindDTO) GoString() string {
	return s.String()
}

func (s *BindAxnResponseBodySecretBindDTO) SetExtension(v string) *BindAxnResponseBodySecretBindDTO {
	s.Extension = &v
	return s
}

func (s *BindAxnResponseBodySecretBindDTO) SetSecretNo(v string) *BindAxnResponseBodySecretBindDTO {
	s.SecretNo = &v
	return s
}

func (s *BindAxnResponseBodySecretBindDTO) SetSubsId(v string) *BindAxnResponseBodySecretBindDTO {
	s.SubsId = &v
	return s
}

type BindAxnResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BindAxnResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BindAxnResponse) String() string {
	return tea.Prettify(s)
}

func (s BindAxnResponse) GoString() string {
	return s.String()
}

func (s *BindAxnResponse) SetHeaders(v map[string]*string) *BindAxnResponse {
	s.Headers = v
	return s
}

func (s *BindAxnResponse) SetStatusCode(v int32) *BindAxnResponse {
	s.StatusCode = &v
	return s
}

func (s *BindAxnResponse) SetBody(v *BindAxnResponseBody) *BindAxnResponse {
	s.Body = v
	return s
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
	return tea.Prettify(s)
}

func (s BindAxnExtensionRequest) GoString() string {
	return s.String()
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

type BindAxnExtensionResponseBody struct {
	// The response code.
	//
	// 	- The value OK indicates that the request was successful.
	//
	// 	- Other values indicate that the request failed. For more information, see [Error codes](https://help.aliyun.com/document_detail/109196.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9297B722-A016-43FB-B51A-E54050D9369D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information returned after the phone numbers were bound.
	SecretBindDTO *BindAxnExtensionResponseBodySecretBindDTO `json:"SecretBindDTO,omitempty" xml:"SecretBindDTO,omitempty" type:"Struct"`
}

func (s BindAxnExtensionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BindAxnExtensionResponseBody) GoString() string {
	return s.String()
}

func (s *BindAxnExtensionResponseBody) SetCode(v string) *BindAxnExtensionResponseBody {
	s.Code = &v
	return s
}

func (s *BindAxnExtensionResponseBody) SetMessage(v string) *BindAxnExtensionResponseBody {
	s.Message = &v
	return s
}

func (s *BindAxnExtensionResponseBody) SetRequestId(v string) *BindAxnExtensionResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindAxnExtensionResponseBody) SetSecretBindDTO(v *BindAxnExtensionResponseBodySecretBindDTO) *BindAxnExtensionResponseBody {
	s.SecretBindDTO = v
	return s
}

type BindAxnExtensionResponseBodySecretBindDTO struct {
	// The extension of the phone number.
	//
	// example:
	//
	// 130
	Extension *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// The private number, that is, phone number X.
	//
	// example:
	//
	// 139*****0000
	SecretNo *string `json:"SecretNo,omitempty" xml:"SecretNo,omitempty"`
	// The binding ID.
	//
	// example:
	//
	// 1***************3
	SubsId *string `json:"SubsId,omitempty" xml:"SubsId,omitempty"`
}

func (s BindAxnExtensionResponseBodySecretBindDTO) String() string {
	return tea.Prettify(s)
}

func (s BindAxnExtensionResponseBodySecretBindDTO) GoString() string {
	return s.String()
}

func (s *BindAxnExtensionResponseBodySecretBindDTO) SetExtension(v string) *BindAxnExtensionResponseBodySecretBindDTO {
	s.Extension = &v
	return s
}

func (s *BindAxnExtensionResponseBodySecretBindDTO) SetSecretNo(v string) *BindAxnExtensionResponseBodySecretBindDTO {
	s.SecretNo = &v
	return s
}

func (s *BindAxnExtensionResponseBodySecretBindDTO) SetSubsId(v string) *BindAxnExtensionResponseBodySecretBindDTO {
	s.SubsId = &v
	return s
}

type BindAxnExtensionResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BindAxnExtensionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BindAxnExtensionResponse) String() string {
	return tea.Prettify(s)
}

func (s BindAxnExtensionResponse) GoString() string {
	return s.String()
}

func (s *BindAxnExtensionResponse) SetHeaders(v map[string]*string) *BindAxnExtensionResponse {
	s.Headers = v
	return s
}

func (s *BindAxnExtensionResponse) SetStatusCode(v int32) *BindAxnExtensionResponse {
	s.StatusCode = &v
	return s
}

func (s *BindAxnExtensionResponse) SetBody(v *BindAxnExtensionResponseBody) *BindAxnExtensionResponse {
	s.Body = v
	return s
}

type BindAxnExtensionFixedLineRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 10001,10002,10003
	Anucode *string `json:"Anucode,omitempty" xml:"Anucode,omitempty"`
	// 被叫侧放音编码  被叫放音编码必须包含2个场景的编码。按照“A被叫,其他号码被叫”的顺序填写编码，编码之间以逗号分隔。  比如：“1,2”表示A号码为被叫侧接听时的放音编号为1，其他号码为被叫侧接听时的放音编号为2
	//
	// example:
	//
	// 10001,10002
	Anucodecalled *string `json:"Anucodecalled,omitempty" xml:"Anucodecalled,omitempty"`
	// 业务id标识，由阿里云分配给客户侧
	//
	// This parameter is required.
	//
	// example:
	//
	// alitest
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 去掉“0” 例如：北京（10）；在平台分配X号码模式中，平台从号码池中分配该地区的X号码，避免产生呼叫长途费
	//
	// example:
	//
	// 10
	Areacode *string `json:"Areacode,omitempty" xml:"Areacode,omitempty"`
	// 绑定类型，值如下： mode101：客户携带X号码 mode102：平台分配X号码
	//
	// This parameter is required.
	//
	// example:
	//
	// mode101
	BindType *string `json:"BindType,omitempty" xml:"BindType,omitempty"`
	// 单位：秒，必须为数字 0：不会自动解绑 非0：自动解绑周期
	//
	// This parameter is required.
	//
	// example:
	//
	// 60
	Expiration *string `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	// 额外字段
	Extraaxx *BindAxnExtensionFixedLineRequestExtraaxx `json:"Extraaxx,omitempty" xml:"Extraaxx,omitempty" type:"Struct"`
	// 消息请求唯一标识
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345dkwkd99d
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 接入商自有字段，最大100字符长度
	//
	// example:
	//
	// 12444
	Remark               *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 格式为yyyyMMddHHmmss。时间采用北京时间，24小时制。
	//
	// This parameter is required.
	//
	// example:
	//
	// 20250421141723
	Subts *string `json:"Subts,omitempty" xml:"Subts,omitempty"`
	// 接通后主叫侧放音编码  接通后主叫侧放音编码必须包含2个场景的编码。按照“A被叫,其他号码被叫”的顺序填写编码，编码之间以逗号分隔。  比如：“1,2”表示A号码为被叫时主叫侧的放音编号为1，其他号码为被叫时主叫侧的放音编号为2
	//
	// example:
	//
	// 10001,10002
	TAnucodeConnect *string `json:"TAnucodeConnect,omitempty" xml:"TAnucodeConnect,omitempty"`
	// A号码
	//
	// This parameter is required.
	//
	// example:
	//
	// 15500001111放音编码必须包含3个场景的编码。按照“B->X,A->X,其他号码->X”的顺序填写编码，编码之间以逗号分隔。  AXN分机号业务的放音编码,B->X和其他号码->X的编码一致  比如：“1,2,3”表示B->X放音编号为1，A->X放音编号为2， 其他号码->X放音编号为3
	TelA *string `json:"TelA,omitempty" xml:"TelA,omitempty"`
	// X号码；平台分配号码模式下，该参数可不带，系统忽略该参数
	//
	// example:
	//
	// 0571409312
	TelX *string `json:"TelX,omitempty" xml:"TelX,omitempty"`
	// 分机号；平台分配号码模式下，该参数可不带，系统忽略该参数
	//
	// example:
	//
	// 1009
	TelXext *string `json:"TelXext,omitempty" xml:"TelXext,omitempty"`
	// 格式yyyyMMddHHmmssSSS，时间采用北京时间，24小时制，精确至毫秒
	//
	// This parameter is required.
	//
	// example:
	//
	// 20250421141723098
	Ts *string `json:"Ts,omitempty" xml:"Ts,omitempty"`
}

func (s BindAxnExtensionFixedLineRequest) String() string {
	return tea.Prettify(s)
}

func (s BindAxnExtensionFixedLineRequest) GoString() string {
	return s.String()
}

func (s *BindAxnExtensionFixedLineRequest) SetAnucode(v string) *BindAxnExtensionFixedLineRequest {
	s.Anucode = &v
	return s
}

func (s *BindAxnExtensionFixedLineRequest) SetAnucodecalled(v string) *BindAxnExtensionFixedLineRequest {
	s.Anucodecalled = &v
	return s
}

func (s *BindAxnExtensionFixedLineRequest) SetAppId(v string) *BindAxnExtensionFixedLineRequest {
	s.AppId = &v
	return s
}

func (s *BindAxnExtensionFixedLineRequest) SetAreacode(v string) *BindAxnExtensionFixedLineRequest {
	s.Areacode = &v
	return s
}

func (s *BindAxnExtensionFixedLineRequest) SetBindType(v string) *BindAxnExtensionFixedLineRequest {
	s.BindType = &v
	return s
}

func (s *BindAxnExtensionFixedLineRequest) SetExpiration(v string) *BindAxnExtensionFixedLineRequest {
	s.Expiration = &v
	return s
}

func (s *BindAxnExtensionFixedLineRequest) SetExtraaxx(v *BindAxnExtensionFixedLineRequestExtraaxx) *BindAxnExtensionFixedLineRequest {
	s.Extraaxx = v
	return s
}

func (s *BindAxnExtensionFixedLineRequest) SetOrderId(v string) *BindAxnExtensionFixedLineRequest {
	s.OrderId = &v
	return s
}

func (s *BindAxnExtensionFixedLineRequest) SetOwnerId(v int64) *BindAxnExtensionFixedLineRequest {
	s.OwnerId = &v
	return s
}

func (s *BindAxnExtensionFixedLineRequest) SetRemark(v string) *BindAxnExtensionFixedLineRequest {
	s.Remark = &v
	return s
}

func (s *BindAxnExtensionFixedLineRequest) SetResourceOwnerAccount(v string) *BindAxnExtensionFixedLineRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *BindAxnExtensionFixedLineRequest) SetResourceOwnerId(v int64) *BindAxnExtensionFixedLineRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *BindAxnExtensionFixedLineRequest) SetSubts(v string) *BindAxnExtensionFixedLineRequest {
	s.Subts = &v
	return s
}

func (s *BindAxnExtensionFixedLineRequest) SetTAnucodeConnect(v string) *BindAxnExtensionFixedLineRequest {
	s.TAnucodeConnect = &v
	return s
}

func (s *BindAxnExtensionFixedLineRequest) SetTelA(v string) *BindAxnExtensionFixedLineRequest {
	s.TelA = &v
	return s
}

func (s *BindAxnExtensionFixedLineRequest) SetTelX(v string) *BindAxnExtensionFixedLineRequest {
	s.TelX = &v
	return s
}

func (s *BindAxnExtensionFixedLineRequest) SetTelXext(v string) *BindAxnExtensionFixedLineRequest {
	s.TelXext = &v
	return s
}

func (s *BindAxnExtensionFixedLineRequest) SetTs(v string) *BindAxnExtensionFixedLineRequest {
	s.Ts = &v
	return s
}

type BindAxnExtensionFixedLineRequestExtraaxx struct {
	// A通过X呼叫，即A主叫X，仅下列值有效。默认是0。 0：不能外呼 1：接续最近的B号码
	//
	// example:
	//
	// 0
	Callback *string `json:"Callback,omitempty" xml:"Callback,omitempty"`
	// 录音控制，仅下列值有效。默认是0（不开通录音功能）。 0：不录音 1：录音
	//
	// example:
	//
	// 1
	Callrecording *string `json:"Callrecording,omitempty" xml:"Callrecording,omitempty"`
}

func (s BindAxnExtensionFixedLineRequestExtraaxx) String() string {
	return tea.Prettify(s)
}

func (s BindAxnExtensionFixedLineRequestExtraaxx) GoString() string {
	return s.String()
}

func (s *BindAxnExtensionFixedLineRequestExtraaxx) SetCallback(v string) *BindAxnExtensionFixedLineRequestExtraaxx {
	s.Callback = &v
	return s
}

func (s *BindAxnExtensionFixedLineRequestExtraaxx) SetCallrecording(v string) *BindAxnExtensionFixedLineRequestExtraaxx {
	s.Callrecording = &v
	return s
}

type BindAxnExtensionFixedLineShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 10001,10002,10003
	Anucode *string `json:"Anucode,omitempty" xml:"Anucode,omitempty"`
	// 被叫侧放音编码  被叫放音编码必须包含2个场景的编码。按照“A被叫,其他号码被叫”的顺序填写编码，编码之间以逗号分隔。  比如：“1,2”表示A号码为被叫侧接听时的放音编号为1，其他号码为被叫侧接听时的放音编号为2
	//
	// example:
	//
	// 10001,10002
	Anucodecalled *string `json:"Anucodecalled,omitempty" xml:"Anucodecalled,omitempty"`
	// 业务id标识，由阿里云分配给客户侧
	//
	// This parameter is required.
	//
	// example:
	//
	// alitest
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 去掉“0” 例如：北京（10）；在平台分配X号码模式中，平台从号码池中分配该地区的X号码，避免产生呼叫长途费
	//
	// example:
	//
	// 10
	Areacode *string `json:"Areacode,omitempty" xml:"Areacode,omitempty"`
	// 绑定类型，值如下： mode101：客户携带X号码 mode102：平台分配X号码
	//
	// This parameter is required.
	//
	// example:
	//
	// mode101
	BindType *string `json:"BindType,omitempty" xml:"BindType,omitempty"`
	// 单位：秒，必须为数字 0：不会自动解绑 非0：自动解绑周期
	//
	// This parameter is required.
	//
	// example:
	//
	// 60
	Expiration *string `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	// 额外字段
	ExtraaxxShrink *string `json:"Extraaxx,omitempty" xml:"Extraaxx,omitempty"`
	// 消息请求唯一标识
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345dkwkd99d
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 接入商自有字段，最大100字符长度
	//
	// example:
	//
	// 12444
	Remark               *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 格式为yyyyMMddHHmmss。时间采用北京时间，24小时制。
	//
	// This parameter is required.
	//
	// example:
	//
	// 20250421141723
	Subts *string `json:"Subts,omitempty" xml:"Subts,omitempty"`
	// 接通后主叫侧放音编码  接通后主叫侧放音编码必须包含2个场景的编码。按照“A被叫,其他号码被叫”的顺序填写编码，编码之间以逗号分隔。  比如：“1,2”表示A号码为被叫时主叫侧的放音编号为1，其他号码为被叫时主叫侧的放音编号为2
	//
	// example:
	//
	// 10001,10002
	TAnucodeConnect *string `json:"TAnucodeConnect,omitempty" xml:"TAnucodeConnect,omitempty"`
	// A号码
	//
	// This parameter is required.
	//
	// example:
	//
	// 15500001111放音编码必须包含3个场景的编码。按照“B->X,A->X,其他号码->X”的顺序填写编码，编码之间以逗号分隔。  AXN分机号业务的放音编码,B->X和其他号码->X的编码一致  比如：“1,2,3”表示B->X放音编号为1，A->X放音编号为2， 其他号码->X放音编号为3
	TelA *string `json:"TelA,omitempty" xml:"TelA,omitempty"`
	// X号码；平台分配号码模式下，该参数可不带，系统忽略该参数
	//
	// example:
	//
	// 0571409312
	TelX *string `json:"TelX,omitempty" xml:"TelX,omitempty"`
	// 分机号；平台分配号码模式下，该参数可不带，系统忽略该参数
	//
	// example:
	//
	// 1009
	TelXext *string `json:"TelXext,omitempty" xml:"TelXext,omitempty"`
	// 格式yyyyMMddHHmmssSSS，时间采用北京时间，24小时制，精确至毫秒
	//
	// This parameter is required.
	//
	// example:
	//
	// 20250421141723098
	Ts *string `json:"Ts,omitempty" xml:"Ts,omitempty"`
}

func (s BindAxnExtensionFixedLineShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s BindAxnExtensionFixedLineShrinkRequest) GoString() string {
	return s.String()
}

func (s *BindAxnExtensionFixedLineShrinkRequest) SetAnucode(v string) *BindAxnExtensionFixedLineShrinkRequest {
	s.Anucode = &v
	return s
}

func (s *BindAxnExtensionFixedLineShrinkRequest) SetAnucodecalled(v string) *BindAxnExtensionFixedLineShrinkRequest {
	s.Anucodecalled = &v
	return s
}

func (s *BindAxnExtensionFixedLineShrinkRequest) SetAppId(v string) *BindAxnExtensionFixedLineShrinkRequest {
	s.AppId = &v
	return s
}

func (s *BindAxnExtensionFixedLineShrinkRequest) SetAreacode(v string) *BindAxnExtensionFixedLineShrinkRequest {
	s.Areacode = &v
	return s
}

func (s *BindAxnExtensionFixedLineShrinkRequest) SetBindType(v string) *BindAxnExtensionFixedLineShrinkRequest {
	s.BindType = &v
	return s
}

func (s *BindAxnExtensionFixedLineShrinkRequest) SetExpiration(v string) *BindAxnExtensionFixedLineShrinkRequest {
	s.Expiration = &v
	return s
}

func (s *BindAxnExtensionFixedLineShrinkRequest) SetExtraaxxShrink(v string) *BindAxnExtensionFixedLineShrinkRequest {
	s.ExtraaxxShrink = &v
	return s
}

func (s *BindAxnExtensionFixedLineShrinkRequest) SetOrderId(v string) *BindAxnExtensionFixedLineShrinkRequest {
	s.OrderId = &v
	return s
}

func (s *BindAxnExtensionFixedLineShrinkRequest) SetOwnerId(v int64) *BindAxnExtensionFixedLineShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *BindAxnExtensionFixedLineShrinkRequest) SetRemark(v string) *BindAxnExtensionFixedLineShrinkRequest {
	s.Remark = &v
	return s
}

func (s *BindAxnExtensionFixedLineShrinkRequest) SetResourceOwnerAccount(v string) *BindAxnExtensionFixedLineShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *BindAxnExtensionFixedLineShrinkRequest) SetResourceOwnerId(v int64) *BindAxnExtensionFixedLineShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *BindAxnExtensionFixedLineShrinkRequest) SetSubts(v string) *BindAxnExtensionFixedLineShrinkRequest {
	s.Subts = &v
	return s
}

func (s *BindAxnExtensionFixedLineShrinkRequest) SetTAnucodeConnect(v string) *BindAxnExtensionFixedLineShrinkRequest {
	s.TAnucodeConnect = &v
	return s
}

func (s *BindAxnExtensionFixedLineShrinkRequest) SetTelA(v string) *BindAxnExtensionFixedLineShrinkRequest {
	s.TelA = &v
	return s
}

func (s *BindAxnExtensionFixedLineShrinkRequest) SetTelX(v string) *BindAxnExtensionFixedLineShrinkRequest {
	s.TelX = &v
	return s
}

func (s *BindAxnExtensionFixedLineShrinkRequest) SetTelXext(v string) *BindAxnExtensionFixedLineShrinkRequest {
	s.TelXext = &v
	return s
}

func (s *BindAxnExtensionFixedLineShrinkRequest) SetTs(v string) *BindAxnExtensionFixedLineShrinkRequest {
	s.Ts = &v
	return s
}

type BindAxnExtensionFixedLineResponseBody struct {
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// 响应码 0：成功，其它失败，具体见文档
	//
	// example:
	//
	// 0
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 响应信息
	Data *BindAxnExtensionFixedLineResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// 描述信息
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// AE2D6997-643A-59CB-9B3C-918572E5CEAA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BindAxnExtensionFixedLineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BindAxnExtensionFixedLineResponseBody) GoString() string {
	return s.String()
}

func (s *BindAxnExtensionFixedLineResponseBody) SetAccessDeniedDetail(v string) *BindAxnExtensionFixedLineResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *BindAxnExtensionFixedLineResponseBody) SetCode(v string) *BindAxnExtensionFixedLineResponseBody {
	s.Code = &v
	return s
}

func (s *BindAxnExtensionFixedLineResponseBody) SetData(v *BindAxnExtensionFixedLineResponseBodyData) *BindAxnExtensionFixedLineResponseBody {
	s.Data = v
	return s
}

func (s *BindAxnExtensionFixedLineResponseBody) SetMessage(v string) *BindAxnExtensionFixedLineResponseBody {
	s.Message = &v
	return s
}

func (s *BindAxnExtensionFixedLineResponseBody) SetRequestId(v string) *BindAxnExtensionFixedLineResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindAxnExtensionFixedLineResponseBody) SetSuccess(v string) *BindAxnExtensionFixedLineResponseBody {
	s.Success = &v
	return s
}

type BindAxnExtensionFixedLineResponseBodyData struct {
	// 绑定id
	//
	// example:
	//
	// GHX0534X202504221531579290029
	Subid *string `json:"Subid,omitempty" xml:"Subid,omitempty"`
	// 隐私号码
	//
	// example:
	//
	// 0571409312
	TelX *string `json:"TelX,omitempty" xml:"TelX,omitempty"`
	// 分机号，只有4位
	//
	// example:
	//
	// 1001
	TelXext *string `json:"TelXext,omitempty" xml:"TelXext,omitempty"`
}

func (s BindAxnExtensionFixedLineResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s BindAxnExtensionFixedLineResponseBodyData) GoString() string {
	return s.String()
}

func (s *BindAxnExtensionFixedLineResponseBodyData) SetSubid(v string) *BindAxnExtensionFixedLineResponseBodyData {
	s.Subid = &v
	return s
}

func (s *BindAxnExtensionFixedLineResponseBodyData) SetTelX(v string) *BindAxnExtensionFixedLineResponseBodyData {
	s.TelX = &v
	return s
}

func (s *BindAxnExtensionFixedLineResponseBodyData) SetTelXext(v string) *BindAxnExtensionFixedLineResponseBodyData {
	s.TelXext = &v
	return s
}

type BindAxnExtensionFixedLineResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BindAxnExtensionFixedLineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BindAxnExtensionFixedLineResponse) String() string {
	return tea.Prettify(s)
}

func (s BindAxnExtensionFixedLineResponse) GoString() string {
	return s.String()
}

func (s *BindAxnExtensionFixedLineResponse) SetHeaders(v map[string]*string) *BindAxnExtensionFixedLineResponse {
	s.Headers = v
	return s
}

func (s *BindAxnExtensionFixedLineResponse) SetStatusCode(v int32) *BindAxnExtensionFixedLineResponse {
	s.StatusCode = &v
	return s
}

func (s *BindAxnExtensionFixedLineResponse) SetBody(v *BindAxnExtensionFixedLineResponseBody) *BindAxnExtensionFixedLineResponse {
	s.Body = v
	return s
}

type BindAxnFixedLineRequest struct {
	// 放音编码必须包含3个场景的编码。按照“B->X,A->X,其他号码->X”的顺序填写编码，编码之间以逗号分隔。  比如：“1,2,3”表示B->X放音编号为1，A->X放音编号为2， 其他号码->X放音编号为3
	//
	// This parameter is required.
	//
	// example:
	//
	// 10001,10002,10003
	Anucode *string `json:"Anucode,omitempty" xml:"Anucode,omitempty"`
	// 被叫侧放音编码  被叫放音编码必须包含2个场景的编码。按照“A被叫,其他号码被叫”的顺序填写编码，编码之间以逗号分隔。  比如：“1,2”表示A号码为被叫侧接听时的放音编号为1，其他号码为被叫侧接听时的放音编号为2
	//
	// example:
	//
	// 10001,10002
	Anucodecalled *string `json:"Anucodecalled,omitempty" xml:"Anucodecalled,omitempty"`
	// 业务id标识，由阿里云分配给客户侧
	//
	// This parameter is required.
	//
	// example:
	//
	// alitest
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 去掉“0” 例如：北京（10）；在平台分配X号码模式中，平台从号码池中分配该地区的X号码，避免产生呼叫长途费
	//
	// example:
	//
	// 10
	Areacode *string `json:"Areacode,omitempty" xml:"Areacode,omitempty"`
	// 绑定类型，值如下： mode101：客户携带X号码 mode102：平台分配X号码
	//
	// This parameter is required.
	//
	// example:
	//
	// mode101
	BindType *string `json:"BindType,omitempty" xml:"BindType,omitempty"`
	// 位：秒，必须为数字 0：不会自动解绑 非0：自动解绑周期
	//
	// This parameter is required.
	//
	// example:
	//
	// 示例值
	Expiration *string `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	// 扩展参数
	Extra *BindAxnFixedLineRequestExtra `json:"Extra,omitempty" xml:"Extra,omitempty" type:"Struct"`
	// 消息请求唯一标识
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345dkwkd99d
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 接入商自有字段，最大100字符长度
	//
	// example:
	//
	// 12444
	Remark               *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 格式为yyyyMMddHHmmss。时间采用北京时间，24小时制。
	//
	// This parameter is required.
	//
	// example:
	//
	// 20250421141723098
	Subts *string `json:"Subts,omitempty" xml:"Subts,omitempty"`
	// 接通后主叫侧放音编码  接通后主叫侧放音编码必须包含2个场景的编码。按照“A被叫,其他号码被叫”的顺序填写编码，编码之间以逗号分隔。  比如：“1,2”表示A号码为被叫时主叫侧的放音编号为1，其他号码被叫为被叫时主叫侧的放音编号为2
	//
	// example:
	//
	// 10001,10002
	TAnucodeConnect *string `json:"TAnucodeConnect,omitempty" xml:"TAnucodeConnect,omitempty"`
	// A号码
	//
	// This parameter is required.
	//
	// example:
	//
	// 15500001111
	TelA *string `json:"TelA,omitempty" xml:"TelA,omitempty"`
	// N号码
	//
	// example:
	//
	// 15500002222
	TelB *string `json:"TelB,omitempty" xml:"TelB,omitempty"`
	// X号码； 平台分配号码模式下，该参数可不带，系统忽略该参数
	//
	// example:
	//
	// 0571409312
	TelX *string `json:"TelX,omitempty" xml:"TelX,omitempty"`
	// 格式yyyyMMddHHmmssSSS，时间采用北京时间，24小时制，精确至毫秒
	//
	// This parameter is required.
	//
	// example:
	//
	// 20250421141723098
	Ts *string `json:"Ts,omitempty" xml:"Ts,omitempty"`
}

func (s BindAxnFixedLineRequest) String() string {
	return tea.Prettify(s)
}

func (s BindAxnFixedLineRequest) GoString() string {
	return s.String()
}

func (s *BindAxnFixedLineRequest) SetAnucode(v string) *BindAxnFixedLineRequest {
	s.Anucode = &v
	return s
}

func (s *BindAxnFixedLineRequest) SetAnucodecalled(v string) *BindAxnFixedLineRequest {
	s.Anucodecalled = &v
	return s
}

func (s *BindAxnFixedLineRequest) SetAppId(v string) *BindAxnFixedLineRequest {
	s.AppId = &v
	return s
}

func (s *BindAxnFixedLineRequest) SetAreacode(v string) *BindAxnFixedLineRequest {
	s.Areacode = &v
	return s
}

func (s *BindAxnFixedLineRequest) SetBindType(v string) *BindAxnFixedLineRequest {
	s.BindType = &v
	return s
}

func (s *BindAxnFixedLineRequest) SetExpiration(v string) *BindAxnFixedLineRequest {
	s.Expiration = &v
	return s
}

func (s *BindAxnFixedLineRequest) SetExtra(v *BindAxnFixedLineRequestExtra) *BindAxnFixedLineRequest {
	s.Extra = v
	return s
}

func (s *BindAxnFixedLineRequest) SetOrderId(v string) *BindAxnFixedLineRequest {
	s.OrderId = &v
	return s
}

func (s *BindAxnFixedLineRequest) SetOwnerId(v int64) *BindAxnFixedLineRequest {
	s.OwnerId = &v
	return s
}

func (s *BindAxnFixedLineRequest) SetRemark(v string) *BindAxnFixedLineRequest {
	s.Remark = &v
	return s
}

func (s *BindAxnFixedLineRequest) SetResourceOwnerAccount(v string) *BindAxnFixedLineRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *BindAxnFixedLineRequest) SetResourceOwnerId(v int64) *BindAxnFixedLineRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *BindAxnFixedLineRequest) SetSubts(v string) *BindAxnFixedLineRequest {
	s.Subts = &v
	return s
}

func (s *BindAxnFixedLineRequest) SetTAnucodeConnect(v string) *BindAxnFixedLineRequest {
	s.TAnucodeConnect = &v
	return s
}

func (s *BindAxnFixedLineRequest) SetTelA(v string) *BindAxnFixedLineRequest {
	s.TelA = &v
	return s
}

func (s *BindAxnFixedLineRequest) SetTelB(v string) *BindAxnFixedLineRequest {
	s.TelB = &v
	return s
}

func (s *BindAxnFixedLineRequest) SetTelX(v string) *BindAxnFixedLineRequest {
	s.TelX = &v
	return s
}

func (s *BindAxnFixedLineRequest) SetTs(v string) *BindAxnFixedLineRequest {
	s.Ts = &v
	return s
}

type BindAxnFixedLineRequestExtra struct {
	// A通过X呼叫，即A主叫X，仅下列值有效。默认是0。 0：不能外呼 1：接续最近的N号码 2：回拨固定号码：telB
	//
	// example:
	//
	// 0
	Callback *string `json:"Callback,omitempty" xml:"Callback,omitempty"`
	// 仅下列值有效。默认是0（不开通录音功能）。 0：不录音 1：录音
	//
	// example:
	//
	// 0
	Callrecording *string `json:"Callrecording,omitempty" xml:"Callrecording,omitempty"`
}

func (s BindAxnFixedLineRequestExtra) String() string {
	return tea.Prettify(s)
}

func (s BindAxnFixedLineRequestExtra) GoString() string {
	return s.String()
}

func (s *BindAxnFixedLineRequestExtra) SetCallback(v string) *BindAxnFixedLineRequestExtra {
	s.Callback = &v
	return s
}

func (s *BindAxnFixedLineRequestExtra) SetCallrecording(v string) *BindAxnFixedLineRequestExtra {
	s.Callrecording = &v
	return s
}

type BindAxnFixedLineShrinkRequest struct {
	// 放音编码必须包含3个场景的编码。按照“B->X,A->X,其他号码->X”的顺序填写编码，编码之间以逗号分隔。  比如：“1,2,3”表示B->X放音编号为1，A->X放音编号为2， 其他号码->X放音编号为3
	//
	// This parameter is required.
	//
	// example:
	//
	// 10001,10002,10003
	Anucode *string `json:"Anucode,omitempty" xml:"Anucode,omitempty"`
	// 被叫侧放音编码  被叫放音编码必须包含2个场景的编码。按照“A被叫,其他号码被叫”的顺序填写编码，编码之间以逗号分隔。  比如：“1,2”表示A号码为被叫侧接听时的放音编号为1，其他号码为被叫侧接听时的放音编号为2
	//
	// example:
	//
	// 10001,10002
	Anucodecalled *string `json:"Anucodecalled,omitempty" xml:"Anucodecalled,omitempty"`
	// 业务id标识，由阿里云分配给客户侧
	//
	// This parameter is required.
	//
	// example:
	//
	// alitest
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 去掉“0” 例如：北京（10）；在平台分配X号码模式中，平台从号码池中分配该地区的X号码，避免产生呼叫长途费
	//
	// example:
	//
	// 10
	Areacode *string `json:"Areacode,omitempty" xml:"Areacode,omitempty"`
	// 绑定类型，值如下： mode101：客户携带X号码 mode102：平台分配X号码
	//
	// This parameter is required.
	//
	// example:
	//
	// mode101
	BindType *string `json:"BindType,omitempty" xml:"BindType,omitempty"`
	// 位：秒，必须为数字 0：不会自动解绑 非0：自动解绑周期
	//
	// This parameter is required.
	//
	// example:
	//
	// 示例值
	Expiration *string `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	// 扩展参数
	ExtraShrink *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	// 消息请求唯一标识
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345dkwkd99d
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 接入商自有字段，最大100字符长度
	//
	// example:
	//
	// 12444
	Remark               *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 格式为yyyyMMddHHmmss。时间采用北京时间，24小时制。
	//
	// This parameter is required.
	//
	// example:
	//
	// 20250421141723098
	Subts *string `json:"Subts,omitempty" xml:"Subts,omitempty"`
	// 接通后主叫侧放音编码  接通后主叫侧放音编码必须包含2个场景的编码。按照“A被叫,其他号码被叫”的顺序填写编码，编码之间以逗号分隔。  比如：“1,2”表示A号码为被叫时主叫侧的放音编号为1，其他号码被叫为被叫时主叫侧的放音编号为2
	//
	// example:
	//
	// 10001,10002
	TAnucodeConnect *string `json:"TAnucodeConnect,omitempty" xml:"TAnucodeConnect,omitempty"`
	// A号码
	//
	// This parameter is required.
	//
	// example:
	//
	// 15500001111
	TelA *string `json:"TelA,omitempty" xml:"TelA,omitempty"`
	// N号码
	//
	// example:
	//
	// 15500002222
	TelB *string `json:"TelB,omitempty" xml:"TelB,omitempty"`
	// X号码； 平台分配号码模式下，该参数可不带，系统忽略该参数
	//
	// example:
	//
	// 0571409312
	TelX *string `json:"TelX,omitempty" xml:"TelX,omitempty"`
	// 格式yyyyMMddHHmmssSSS，时间采用北京时间，24小时制，精确至毫秒
	//
	// This parameter is required.
	//
	// example:
	//
	// 20250421141723098
	Ts *string `json:"Ts,omitempty" xml:"Ts,omitempty"`
}

func (s BindAxnFixedLineShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s BindAxnFixedLineShrinkRequest) GoString() string {
	return s.String()
}

func (s *BindAxnFixedLineShrinkRequest) SetAnucode(v string) *BindAxnFixedLineShrinkRequest {
	s.Anucode = &v
	return s
}

func (s *BindAxnFixedLineShrinkRequest) SetAnucodecalled(v string) *BindAxnFixedLineShrinkRequest {
	s.Anucodecalled = &v
	return s
}

func (s *BindAxnFixedLineShrinkRequest) SetAppId(v string) *BindAxnFixedLineShrinkRequest {
	s.AppId = &v
	return s
}

func (s *BindAxnFixedLineShrinkRequest) SetAreacode(v string) *BindAxnFixedLineShrinkRequest {
	s.Areacode = &v
	return s
}

func (s *BindAxnFixedLineShrinkRequest) SetBindType(v string) *BindAxnFixedLineShrinkRequest {
	s.BindType = &v
	return s
}

func (s *BindAxnFixedLineShrinkRequest) SetExpiration(v string) *BindAxnFixedLineShrinkRequest {
	s.Expiration = &v
	return s
}

func (s *BindAxnFixedLineShrinkRequest) SetExtraShrink(v string) *BindAxnFixedLineShrinkRequest {
	s.ExtraShrink = &v
	return s
}

func (s *BindAxnFixedLineShrinkRequest) SetOrderId(v string) *BindAxnFixedLineShrinkRequest {
	s.OrderId = &v
	return s
}

func (s *BindAxnFixedLineShrinkRequest) SetOwnerId(v int64) *BindAxnFixedLineShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *BindAxnFixedLineShrinkRequest) SetRemark(v string) *BindAxnFixedLineShrinkRequest {
	s.Remark = &v
	return s
}

func (s *BindAxnFixedLineShrinkRequest) SetResourceOwnerAccount(v string) *BindAxnFixedLineShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *BindAxnFixedLineShrinkRequest) SetResourceOwnerId(v int64) *BindAxnFixedLineShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *BindAxnFixedLineShrinkRequest) SetSubts(v string) *BindAxnFixedLineShrinkRequest {
	s.Subts = &v
	return s
}

func (s *BindAxnFixedLineShrinkRequest) SetTAnucodeConnect(v string) *BindAxnFixedLineShrinkRequest {
	s.TAnucodeConnect = &v
	return s
}

func (s *BindAxnFixedLineShrinkRequest) SetTelA(v string) *BindAxnFixedLineShrinkRequest {
	s.TelA = &v
	return s
}

func (s *BindAxnFixedLineShrinkRequest) SetTelB(v string) *BindAxnFixedLineShrinkRequest {
	s.TelB = &v
	return s
}

func (s *BindAxnFixedLineShrinkRequest) SetTelX(v string) *BindAxnFixedLineShrinkRequest {
	s.TelX = &v
	return s
}

func (s *BindAxnFixedLineShrinkRequest) SetTs(v string) *BindAxnFixedLineShrinkRequest {
	s.Ts = &v
	return s
}

type BindAxnFixedLineResponseBody struct {
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// 响应码 0：成功，其它失败，具体见文档
	//
	// example:
	//
	// 0
	Code *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *BindAxnFixedLineResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// 描述信息
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 4D690962-08CE-5D38-A65A-AB247D7DF7A2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BindAxnFixedLineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BindAxnFixedLineResponseBody) GoString() string {
	return s.String()
}

func (s *BindAxnFixedLineResponseBody) SetAccessDeniedDetail(v string) *BindAxnFixedLineResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *BindAxnFixedLineResponseBody) SetCode(v string) *BindAxnFixedLineResponseBody {
	s.Code = &v
	return s
}

func (s *BindAxnFixedLineResponseBody) SetData(v *BindAxnFixedLineResponseBodyData) *BindAxnFixedLineResponseBody {
	s.Data = v
	return s
}

func (s *BindAxnFixedLineResponseBody) SetMessage(v string) *BindAxnFixedLineResponseBody {
	s.Message = &v
	return s
}

func (s *BindAxnFixedLineResponseBody) SetRequestId(v string) *BindAxnFixedLineResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindAxnFixedLineResponseBody) SetSuccess(v bool) *BindAxnFixedLineResponseBody {
	s.Success = &v
	return s
}

type BindAxnFixedLineResponseBodyData struct {
	// 绑定id
	//
	// example:
	//
	// GHX0534X202504221531579290029-2-1-aliaxn
	Subid *string `json:"Subid,omitempty" xml:"Subid,omitempty"`
	// 隐私号码
	//
	// example:
	//
	// 0571409312
	TelX *string `json:"TelX,omitempty" xml:"TelX,omitempty"`
}

func (s BindAxnFixedLineResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s BindAxnFixedLineResponseBodyData) GoString() string {
	return s.String()
}

func (s *BindAxnFixedLineResponseBodyData) SetSubid(v string) *BindAxnFixedLineResponseBodyData {
	s.Subid = &v
	return s
}

func (s *BindAxnFixedLineResponseBodyData) SetTelX(v string) *BindAxnFixedLineResponseBodyData {
	s.TelX = &v
	return s
}

type BindAxnFixedLineResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BindAxnFixedLineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BindAxnFixedLineResponse) String() string {
	return tea.Prettify(s)
}

func (s BindAxnFixedLineResponse) GoString() string {
	return s.String()
}

func (s *BindAxnFixedLineResponse) SetHeaders(v map[string]*string) *BindAxnFixedLineResponse {
	s.Headers = v
	return s
}

func (s *BindAxnFixedLineResponse) SetStatusCode(v int32) *BindAxnFixedLineResponse {
	s.StatusCode = &v
	return s
}

func (s *BindAxnFixedLineResponse) SetBody(v *BindAxnFixedLineResponseBody) *BindAxnFixedLineResponse {
	s.Body = v
	return s
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
	return tea.Prettify(s)
}

func (s BindBatchAxgRequest) GoString() string {
	return s.String()
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
	return tea.Prettify(s)
}

func (s BindBatchAxgRequestAxgBindList) GoString() string {
	return s.String()
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

type BindBatchAxgShrinkRequest struct {
	// This parameter is required.
	AxgBindListShrink *string `json:"AxgBindList,omitempty" xml:"AxgBindList,omitempty"`
	OwnerId           *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// FC2235****
	PoolKey              *string `json:"PoolKey,omitempty" xml:"PoolKey,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s BindBatchAxgShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s BindBatchAxgShrinkRequest) GoString() string {
	return s.String()
}

func (s *BindBatchAxgShrinkRequest) SetAxgBindListShrink(v string) *BindBatchAxgShrinkRequest {
	s.AxgBindListShrink = &v
	return s
}

func (s *BindBatchAxgShrinkRequest) SetOwnerId(v int64) *BindBatchAxgShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *BindBatchAxgShrinkRequest) SetPoolKey(v string) *BindBatchAxgShrinkRequest {
	s.PoolKey = &v
	return s
}

func (s *BindBatchAxgShrinkRequest) SetResourceOwnerAccount(v string) *BindBatchAxgShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *BindBatchAxgShrinkRequest) SetResourceOwnerId(v int64) *BindBatchAxgShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

type BindBatchAxgResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 5DCCA8CD-6C0A-50B4-A496-B1D2AB48A1C3
	RequestId      *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SecretBindList *BindBatchAxgResponseBodySecretBindList `json:"SecretBindList,omitempty" xml:"SecretBindList,omitempty" type:"Struct"`
}

func (s BindBatchAxgResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BindBatchAxgResponseBody) GoString() string {
	return s.String()
}

func (s *BindBatchAxgResponseBody) SetCode(v string) *BindBatchAxgResponseBody {
	s.Code = &v
	return s
}

func (s *BindBatchAxgResponseBody) SetMessage(v string) *BindBatchAxgResponseBody {
	s.Message = &v
	return s
}

func (s *BindBatchAxgResponseBody) SetRequestId(v string) *BindBatchAxgResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindBatchAxgResponseBody) SetSecretBindList(v *BindBatchAxgResponseBodySecretBindList) *BindBatchAxgResponseBody {
	s.SecretBindList = v
	return s
}

type BindBatchAxgResponseBodySecretBindList struct {
	SecretBind []*BindBatchAxgResponseBodySecretBindListSecretBind `json:"SecretBind,omitempty" xml:"SecretBind,omitempty" type:"Repeated"`
}

func (s BindBatchAxgResponseBodySecretBindList) String() string {
	return tea.Prettify(s)
}

func (s BindBatchAxgResponseBodySecretBindList) GoString() string {
	return s.String()
}

func (s *BindBatchAxgResponseBodySecretBindList) SetSecretBind(v []*BindBatchAxgResponseBodySecretBindListSecretBind) *BindBatchAxgResponseBodySecretBindList {
	s.SecretBind = v
	return s
}

type BindBatchAxgResponseBodySecretBindListSecretBind struct {
	// example:
	//
	// isv.INVALID_PARAMETERS
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 257
	Extension *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// example:
	//
	// 1234
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// example:
	//
	// ringConfig invalid
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 13333333333
	PhoneNoA *string `json:"PhoneNoA,omitempty" xml:"PhoneNoA,omitempty"`
	// example:
	//
	// 13333333333
	SecretNo *string `json:"SecretNo,omitempty" xml:"SecretNo,omitempty"`
	// example:
	//
	// 1000085060515673
	SubsId *string `json:"SubsId,omitempty" xml:"SubsId,omitempty"`
}

func (s BindBatchAxgResponseBodySecretBindListSecretBind) String() string {
	return tea.Prettify(s)
}

func (s BindBatchAxgResponseBodySecretBindListSecretBind) GoString() string {
	return s.String()
}

func (s *BindBatchAxgResponseBodySecretBindListSecretBind) SetCode(v string) *BindBatchAxgResponseBodySecretBindListSecretBind {
	s.Code = &v
	return s
}

func (s *BindBatchAxgResponseBodySecretBindListSecretBind) SetExtension(v string) *BindBatchAxgResponseBodySecretBindListSecretBind {
	s.Extension = &v
	return s
}

func (s *BindBatchAxgResponseBodySecretBindListSecretBind) SetGroupId(v string) *BindBatchAxgResponseBodySecretBindListSecretBind {
	s.GroupId = &v
	return s
}

func (s *BindBatchAxgResponseBodySecretBindListSecretBind) SetMessage(v string) *BindBatchAxgResponseBodySecretBindListSecretBind {
	s.Message = &v
	return s
}

func (s *BindBatchAxgResponseBodySecretBindListSecretBind) SetPhoneNoA(v string) *BindBatchAxgResponseBodySecretBindListSecretBind {
	s.PhoneNoA = &v
	return s
}

func (s *BindBatchAxgResponseBodySecretBindListSecretBind) SetSecretNo(v string) *BindBatchAxgResponseBodySecretBindListSecretBind {
	s.SecretNo = &v
	return s
}

func (s *BindBatchAxgResponseBodySecretBindListSecretBind) SetSubsId(v string) *BindBatchAxgResponseBodySecretBindListSecretBind {
	s.SubsId = &v
	return s
}

type BindBatchAxgResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BindBatchAxgResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BindBatchAxgResponse) String() string {
	return tea.Prettify(s)
}

func (s BindBatchAxgResponse) GoString() string {
	return s.String()
}

func (s *BindBatchAxgResponse) SetHeaders(v map[string]*string) *BindBatchAxgResponse {
	s.Headers = v
	return s
}

func (s *BindBatchAxgResponse) SetStatusCode(v int32) *BindBatchAxgResponse {
	s.StatusCode = &v
	return s
}

func (s *BindBatchAxgResponse) SetBody(v *BindBatchAxgResponseBody) *BindBatchAxgResponse {
	s.Body = v
	return s
}

type BindXBRequest struct {
	// 客户uid
	//
	// example:
	//
	// -
	CallerParentId *int64 `json:"CallerParentId,omitempty" xml:"CallerParentId,omitempty"`
	// 号码池key
	//
	// This parameter is required.
	//
	// example:
	//
	// FC5**********************a1a
	CustomerPoolKey *string `json:"CustomerPoolKey,omitempty" xml:"CustomerPoolKey,omitempty"`
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 请求去重ID, reqId最大长度为20位,接入方需要保持原子性
	//
	// This parameter is required.
	//
	// example:
	//
	// 564**********879
	ReqId                *string `json:"ReqId,omitempty" xml:"ReqId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 员工真实号码
	//
	// This parameter is required.
	//
	// example:
	//
	// 18*******22
	TelB *string `json:"TelB,omitempty" xml:"TelB,omitempty"`
	// X号码
	//
	// This parameter is required.
	//
	// example:
	//
	// 17*******22
	TelX *string `json:"TelX,omitempty" xml:"TelX,omitempty"`
	// 客户自定义参数回调带回
	//
	// example:
	//
	// 000
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s BindXBRequest) String() string {
	return tea.Prettify(s)
}

func (s BindXBRequest) GoString() string {
	return s.String()
}

func (s *BindXBRequest) SetCallerParentId(v int64) *BindXBRequest {
	s.CallerParentId = &v
	return s
}

func (s *BindXBRequest) SetCustomerPoolKey(v string) *BindXBRequest {
	s.CustomerPoolKey = &v
	return s
}

func (s *BindXBRequest) SetOwnerId(v int64) *BindXBRequest {
	s.OwnerId = &v
	return s
}

func (s *BindXBRequest) SetReqId(v string) *BindXBRequest {
	s.ReqId = &v
	return s
}

func (s *BindXBRequest) SetResourceOwnerAccount(v string) *BindXBRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *BindXBRequest) SetResourceOwnerId(v int64) *BindXBRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *BindXBRequest) SetTelB(v string) *BindXBRequest {
	s.TelB = &v
	return s
}

func (s *BindXBRequest) SetTelX(v string) *BindXBRequest {
	s.TelX = &v
	return s
}

func (s *BindXBRequest) SetUserData(v string) *BindXBRequest {
	s.UserData = &v
	return s
}

type BindXBResponseBody struct {
	// example:
	//
	// 0
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// 返回状态码 0000表示成功 其他表示失败
	//
	// example:
	//
	// 0000
	Code *string                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *BindXBResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// 返回信息
	//
	// example:
	//
	// 成功
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 返回是否成功 true  表示成功 false表示失败
	//
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BindXBResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BindXBResponseBody) GoString() string {
	return s.String()
}

func (s *BindXBResponseBody) SetAccessDeniedDetail(v string) *BindXBResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *BindXBResponseBody) SetCode(v string) *BindXBResponseBody {
	s.Code = &v
	return s
}

func (s *BindXBResponseBody) SetData(v *BindXBResponseBodyData) *BindXBResponseBody {
	s.Data = v
	return s
}

func (s *BindXBResponseBody) SetMessage(v string) *BindXBResponseBody {
	s.Message = &v
	return s
}

func (s *BindXBResponseBody) SetSuccess(v bool) *BindXBResponseBody {
	s.Success = &v
	return s
}

type BindXBResponseBodyData struct {
	// 工作号关系绑定的唯一标识
	//
	// example:
	//
	// 4353453456
	AuthId *string `json:"AuthId,omitempty" xml:"AuthId,omitempty"`
	// X号码
	//
	// example:
	//
	// 18640577897
	TelX *string `json:"TelX,omitempty" xml:"TelX,omitempty"`
}

func (s BindXBResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s BindXBResponseBodyData) GoString() string {
	return s.String()
}

func (s *BindXBResponseBodyData) SetAuthId(v string) *BindXBResponseBodyData {
	s.AuthId = &v
	return s
}

func (s *BindXBResponseBodyData) SetTelX(v string) *BindXBResponseBodyData {
	s.TelX = &v
	return s
}

type BindXBResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BindXBResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BindXBResponse) String() string {
	return tea.Prettify(s)
}

func (s BindXBResponse) GoString() string {
	return s.String()
}

func (s *BindXBResponse) SetHeaders(v map[string]*string) *BindXBResponse {
	s.Headers = v
	return s
}

func (s *BindXBResponse) SetStatusCode(v int32) *BindXBResponse {
	s.StatusCode = &v
	return s
}

func (s *BindXBResponse) SetBody(v *BindXBResponseBody) *BindXBResponse {
	s.Body = v
	return s
}

type BuySecretNoRequest struct {
	// Specifies the home location of the phone number.
	//
	// >
	//
	// 	- The home location can be set only to a location in the Chinese mainland.
	//
	// 	- A phone number that starts with 95 does not have a home location. If you purchase a phone number that starts with 95, set this parameter to **Nationwide**.
	//
	// This parameter is required.
	//
	// example:
	//
	// hangzhou
	City *string `json:"City,omitempty" xml:"City,omitempty"`
	// Specifies whether to add the phone number to the pool of numbers that will be displayed during calls.
	//
	// >  This parameter takes effect only for customers who have enabled the number display feature.
	//
	// example:
	//
	// true
	DisplayPool *bool  `json:"DisplayPool,omitempty" xml:"DisplayPool,omitempty"`
	OwnerId     *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The key of the phone number pool. Log on to the [Phone Number Protection console](https://dypls.console.aliyun.com/dypls.htm#/account) and view the key of the phone number pool on the **Number Pool Management*	- page.
	//
	// This parameter is required.
	//
	// example:
	//
	// FC123456
	PoolKey              *string `json:"PoolKey,omitempty" xml:"PoolKey,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The prefix of the phone number. If you specify the value of **SecretNo*	- when you purchase a phone number, a phone number starting with the specified prefix is selected.
	//
	// >  You can specify up to 18 digits of the phone number prefix.
	//
	// example:
	//
	// 130
	SecretNo *string `json:"SecretNo,omitempty" xml:"SecretNo,omitempty"`
	// The type of the phone number. Valid values:
	//
	// 	- **1**: a phone number assigned by a virtual network operator, that is, a phone number that belongs to the 170 or 171 number segment.
	//
	// 	- **2**: a phone number provided by a carrier.
	//
	// 	- **3**: a phone number that starts with 95.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	SpecId *int64 `json:"SpecId,omitempty" xml:"SpecId,omitempty"`
}

func (s BuySecretNoRequest) String() string {
	return tea.Prettify(s)
}

func (s BuySecretNoRequest) GoString() string {
	return s.String()
}

func (s *BuySecretNoRequest) SetCity(v string) *BuySecretNoRequest {
	s.City = &v
	return s
}

func (s *BuySecretNoRequest) SetDisplayPool(v bool) *BuySecretNoRequest {
	s.DisplayPool = &v
	return s
}

func (s *BuySecretNoRequest) SetOwnerId(v int64) *BuySecretNoRequest {
	s.OwnerId = &v
	return s
}

func (s *BuySecretNoRequest) SetPoolKey(v string) *BuySecretNoRequest {
	s.PoolKey = &v
	return s
}

func (s *BuySecretNoRequest) SetResourceOwnerAccount(v string) *BuySecretNoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *BuySecretNoRequest) SetResourceOwnerId(v int64) *BuySecretNoRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *BuySecretNoRequest) SetSecretNo(v string) *BuySecretNoRequest {
	s.SecretNo = &v
	return s
}

func (s *BuySecretNoRequest) SetSpecId(v int64) *BuySecretNoRequest {
	s.SpecId = &v
	return s
}

type BuySecretNoResponseBody struct {
	// The response code.
	//
	// 	- The value OK indicates that the request was successful.
	//
	// 	- Other values indicate that the request failed. For more information, see [Error codes](https://help.aliyun.com/document_detail/109196.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2D1AEB96-96D0-454E-B0DC-AE2A8DF08020
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information returned after the operation was called.
	SecretBuyInfoDTO *BuySecretNoResponseBodySecretBuyInfoDTO `json:"SecretBuyInfoDTO,omitempty" xml:"SecretBuyInfoDTO,omitempty" type:"Struct"`
}

func (s BuySecretNoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BuySecretNoResponseBody) GoString() string {
	return s.String()
}

func (s *BuySecretNoResponseBody) SetCode(v string) *BuySecretNoResponseBody {
	s.Code = &v
	return s
}

func (s *BuySecretNoResponseBody) SetMessage(v string) *BuySecretNoResponseBody {
	s.Message = &v
	return s
}

func (s *BuySecretNoResponseBody) SetRequestId(v string) *BuySecretNoResponseBody {
	s.RequestId = &v
	return s
}

func (s *BuySecretNoResponseBody) SetSecretBuyInfoDTO(v *BuySecretNoResponseBodySecretBuyInfoDTO) *BuySecretNoResponseBody {
	s.SecretBuyInfoDTO = v
	return s
}

type BuySecretNoResponseBodySecretBuyInfoDTO struct {
	// The private number, that is, phone number X.
	//
	// example:
	//
	// 1390000****
	SecretNo *string `json:"SecretNo,omitempty" xml:"SecretNo,omitempty"`
}

func (s BuySecretNoResponseBodySecretBuyInfoDTO) String() string {
	return tea.Prettify(s)
}

func (s BuySecretNoResponseBodySecretBuyInfoDTO) GoString() string {
	return s.String()
}

func (s *BuySecretNoResponseBodySecretBuyInfoDTO) SetSecretNo(v string) *BuySecretNoResponseBodySecretBuyInfoDTO {
	s.SecretNo = &v
	return s
}

type BuySecretNoResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BuySecretNoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BuySecretNoResponse) String() string {
	return tea.Prettify(s)
}

func (s BuySecretNoResponse) GoString() string {
	return s.String()
}

func (s *BuySecretNoResponse) SetHeaders(v map[string]*string) *BuySecretNoResponse {
	s.Headers = v
	return s
}

func (s *BuySecretNoResponse) SetStatusCode(v int32) *BuySecretNoResponse {
	s.StatusCode = &v
	return s
}

func (s *BuySecretNoResponse) SetBody(v *BuySecretNoResponseBody) *BuySecretNoResponse {
	s.Body = v
	return s
}

type CancelPickUpWaybillRequest struct {
	// The cancellation reason.
	//
	// This parameter is required.
	//
	// example:
	//
	// {\\"action\\":\\"UPDATE_DESC\\",\\"value\\":\\"The courier is unable to pick up the package.\\"}
	CancelDesc *string `json:"CancelDesc,omitempty" xml:"CancelDesc,omitempty"`
	// The ID of the external order.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1145678823****
	OuterOrderCode       *string `json:"OuterOrderCode,omitempty" xml:"OuterOrderCode,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CancelPickUpWaybillRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelPickUpWaybillRequest) GoString() string {
	return s.String()
}

func (s *CancelPickUpWaybillRequest) SetCancelDesc(v string) *CancelPickUpWaybillRequest {
	s.CancelDesc = &v
	return s
}

func (s *CancelPickUpWaybillRequest) SetOuterOrderCode(v string) *CancelPickUpWaybillRequest {
	s.OuterOrderCode = &v
	return s
}

func (s *CancelPickUpWaybillRequest) SetOwnerId(v int64) *CancelPickUpWaybillRequest {
	s.OwnerId = &v
	return s
}

func (s *CancelPickUpWaybillRequest) SetResourceOwnerAccount(v string) *CancelPickUpWaybillRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CancelPickUpWaybillRequest) SetResourceOwnerId(v int64) *CancelPickUpWaybillRequest {
	s.ResourceOwnerId = &v
	return s
}

type CancelPickUpWaybillResponseBody struct {
	// The response code.
	//
	// 	- The value OK indicates that the request was successful.
	//
	// 	- Other status codes indicate that the request failed. For more information, see [Error codes](https://help.aliyun.com/document_detail/109196.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *CancelPickUpWaybillResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9FC30594-3841-43AD-9008-03393BCB5CD2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelPickUpWaybillResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelPickUpWaybillResponseBody) GoString() string {
	return s.String()
}

func (s *CancelPickUpWaybillResponseBody) SetCode(v string) *CancelPickUpWaybillResponseBody {
	s.Code = &v
	return s
}

func (s *CancelPickUpWaybillResponseBody) SetData(v *CancelPickUpWaybillResponseBodyData) *CancelPickUpWaybillResponseBody {
	s.Data = v
	return s
}

func (s *CancelPickUpWaybillResponseBody) SetMessage(v string) *CancelPickUpWaybillResponseBody {
	s.Message = &v
	return s
}

func (s *CancelPickUpWaybillResponseBody) SetRequestId(v string) *CancelPickUpWaybillResponseBody {
	s.RequestId = &v
	return s
}

type CancelPickUpWaybillResponseBodyData struct {
	// The error code.
	//
	// example:
	//
	// none
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// none
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// The cancellation result.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Indicates whether the cancellation was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CancelPickUpWaybillResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CancelPickUpWaybillResponseBodyData) GoString() string {
	return s.String()
}

func (s *CancelPickUpWaybillResponseBodyData) SetErrorCode(v string) *CancelPickUpWaybillResponseBodyData {
	s.ErrorCode = &v
	return s
}

func (s *CancelPickUpWaybillResponseBodyData) SetErrorMsg(v string) *CancelPickUpWaybillResponseBodyData {
	s.ErrorMsg = &v
	return s
}

func (s *CancelPickUpWaybillResponseBodyData) SetMessage(v string) *CancelPickUpWaybillResponseBodyData {
	s.Message = &v
	return s
}

func (s *CancelPickUpWaybillResponseBodyData) SetSuccess(v bool) *CancelPickUpWaybillResponseBodyData {
	s.Success = &v
	return s
}

type CancelPickUpWaybillResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelPickUpWaybillResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelPickUpWaybillResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelPickUpWaybillResponse) GoString() string {
	return s.String()
}

func (s *CancelPickUpWaybillResponse) SetHeaders(v map[string]*string) *CancelPickUpWaybillResponse {
	s.Headers = v
	return s
}

func (s *CancelPickUpWaybillResponse) SetStatusCode(v int32) *CancelPickUpWaybillResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelPickUpWaybillResponse) SetBody(v *CancelPickUpWaybillResponseBody) *CancelPickUpWaybillResponse {
	s.Body = v
	return s
}

type ConfigXRequest struct {
	// 开/关呼叫能力状态‘0’：禁用‘1’：开启
	//
	// example:
	//
	// 0
	CallAbility *string `json:"CallAbility,omitempty" xml:"CallAbility,omitempty"`
	// 客户uid
	//
	// example:
	//
	// 1898871967585852
	CallerParentId *int64 `json:"CallerParentId,omitempty" xml:"CallerParentId,omitempty"`
	// 号码池key
	//
	// This parameter is required.
	//
	// example:
	//
	// FC5**********************a1a
	CustomerPoolKey *string `json:"CustomerPoolKey,omitempty" xml:"CustomerPoolKey,omitempty"`
	// 是否透传来显为真实主叫：00-非透传：互相拨打时都显示工作号;10-透传：A客户为主叫时,B员工的来显为客户A号码;B员工为主叫时,A客户的来显为工作号;默认为 00
	//
	// example:
	//
	// 0
	GNFlag  *string `json:"GNFlag,omitempty" xml:"GNFlag,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 请求去重ID, reqId最大长度为20位,接入方需要保持原子性
	//
	// This parameter is required.
	//
	// example:
	//
	// 564**********879
	ReqId                *string `json:"ReqId,omitempty" xml:"ReqId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 顺振控制参数
	SequenceCalls []*ConfigXRequestSequenceCalls `json:"SequenceCalls,omitempty" xml:"SequenceCalls,omitempty" type:"Repeated"`
	// 顺振模式：0-不顺振（默认）1-有条件顺振，先接续calledNo指定被叫，如果该被叫未能接通，再顺振sequenceCalls号码列表2-无条件顺振，不接续calledNo指定被叫，直接顺振sequenceCalls号码列表
	//
	// example:
	//
	// 0
	SequenceMode *string `json:"SequenceMode,omitempty" xml:"SequenceMode,omitempty"`
	// 开/关短信功能状态‘0’：禁用；‘1’：开启；
	//
	// example:
	//
	// 0
	SmsAbility *string `json:"SmsAbility,omitempty" xml:"SmsAbility,omitempty"`
	// 是否透传来显为真实用户0：不透传; 1：透传默认：0不透传
	//
	// example:
	//
	// 0
	SmsSignMode *string `json:"SmsSignMode,omitempty" xml:"SmsSignMode,omitempty"`
	// X号码
	//
	// This parameter is required.
	//
	// example:
	//
	// 17*******22
	TelX *string `json:"TelX,omitempty" xml:"TelX,omitempty"`
}

func (s ConfigXRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfigXRequest) GoString() string {
	return s.String()
}

func (s *ConfigXRequest) SetCallAbility(v string) *ConfigXRequest {
	s.CallAbility = &v
	return s
}

func (s *ConfigXRequest) SetCallerParentId(v int64) *ConfigXRequest {
	s.CallerParentId = &v
	return s
}

func (s *ConfigXRequest) SetCustomerPoolKey(v string) *ConfigXRequest {
	s.CustomerPoolKey = &v
	return s
}

func (s *ConfigXRequest) SetGNFlag(v string) *ConfigXRequest {
	s.GNFlag = &v
	return s
}

func (s *ConfigXRequest) SetOwnerId(v int64) *ConfigXRequest {
	s.OwnerId = &v
	return s
}

func (s *ConfigXRequest) SetReqId(v string) *ConfigXRequest {
	s.ReqId = &v
	return s
}

func (s *ConfigXRequest) SetResourceOwnerAccount(v string) *ConfigXRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ConfigXRequest) SetResourceOwnerId(v int64) *ConfigXRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ConfigXRequest) SetSequenceCalls(v []*ConfigXRequestSequenceCalls) *ConfigXRequest {
	s.SequenceCalls = v
	return s
}

func (s *ConfigXRequest) SetSequenceMode(v string) *ConfigXRequest {
	s.SequenceMode = &v
	return s
}

func (s *ConfigXRequest) SetSmsAbility(v string) *ConfigXRequest {
	s.SmsAbility = &v
	return s
}

func (s *ConfigXRequest) SetSmsSignMode(v string) *ConfigXRequest {
	s.SmsSignMode = &v
	return s
}

func (s *ConfigXRequest) SetTelX(v string) *ConfigXRequest {
	s.TelX = &v
	return s
}

type ConfigXRequestSequenceCalls struct {
	// 顺振提示音放音编号，格式如callNoPlayCode
	//
	// example:
	//
	// 01
	SequenceCallNoPlayCode *string `json:"SequenceCallNoPlayCode,omitempty" xml:"SequenceCallNoPlayCode,omitempty"`
	// 顺振被叫号码
	//
	// example:
	//
	// 18*******33
	SequenceCalledNo *string `json:"SequenceCalledNo,omitempty" xml:"SequenceCalledNo,omitempty"`
	// 接通后主被叫放音编号，格式如calledPlayCode
	//
	// example:
	//
	// 02
	SequenceCalledPlayCode *string `json:"SequenceCalledPlayCode,omitempty" xml:"SequenceCalledPlayCode,omitempty"`
}

func (s ConfigXRequestSequenceCalls) String() string {
	return tea.Prettify(s)
}

func (s ConfigXRequestSequenceCalls) GoString() string {
	return s.String()
}

func (s *ConfigXRequestSequenceCalls) SetSequenceCallNoPlayCode(v string) *ConfigXRequestSequenceCalls {
	s.SequenceCallNoPlayCode = &v
	return s
}

func (s *ConfigXRequestSequenceCalls) SetSequenceCalledNo(v string) *ConfigXRequestSequenceCalls {
	s.SequenceCalledNo = &v
	return s
}

func (s *ConfigXRequestSequenceCalls) SetSequenceCalledPlayCode(v string) *ConfigXRequestSequenceCalls {
	s.SequenceCalledPlayCode = &v
	return s
}

type ConfigXShrinkRequest struct {
	// 开/关呼叫能力状态‘0’：禁用‘1’：开启
	//
	// example:
	//
	// 0
	CallAbility *string `json:"CallAbility,omitempty" xml:"CallAbility,omitempty"`
	// 客户uid
	//
	// example:
	//
	// 1898871967585852
	CallerParentId *int64 `json:"CallerParentId,omitempty" xml:"CallerParentId,omitempty"`
	// 号码池key
	//
	// This parameter is required.
	//
	// example:
	//
	// FC5**********************a1a
	CustomerPoolKey *string `json:"CustomerPoolKey,omitempty" xml:"CustomerPoolKey,omitempty"`
	// 是否透传来显为真实主叫：00-非透传：互相拨打时都显示工作号;10-透传：A客户为主叫时,B员工的来显为客户A号码;B员工为主叫时,A客户的来显为工作号;默认为 00
	//
	// example:
	//
	// 0
	GNFlag  *string `json:"GNFlag,omitempty" xml:"GNFlag,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 请求去重ID, reqId最大长度为20位,接入方需要保持原子性
	//
	// This parameter is required.
	//
	// example:
	//
	// 564**********879
	ReqId                *string `json:"ReqId,omitempty" xml:"ReqId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 顺振控制参数
	SequenceCallsShrink *string `json:"SequenceCalls,omitempty" xml:"SequenceCalls,omitempty"`
	// 顺振模式：0-不顺振（默认）1-有条件顺振，先接续calledNo指定被叫，如果该被叫未能接通，再顺振sequenceCalls号码列表2-无条件顺振，不接续calledNo指定被叫，直接顺振sequenceCalls号码列表
	//
	// example:
	//
	// 0
	SequenceMode *string `json:"SequenceMode,omitempty" xml:"SequenceMode,omitempty"`
	// 开/关短信功能状态‘0’：禁用；‘1’：开启；
	//
	// example:
	//
	// 0
	SmsAbility *string `json:"SmsAbility,omitempty" xml:"SmsAbility,omitempty"`
	// 是否透传来显为真实用户0：不透传; 1：透传默认：0不透传
	//
	// example:
	//
	// 0
	SmsSignMode *string `json:"SmsSignMode,omitempty" xml:"SmsSignMode,omitempty"`
	// X号码
	//
	// This parameter is required.
	//
	// example:
	//
	// 17*******22
	TelX *string `json:"TelX,omitempty" xml:"TelX,omitempty"`
}

func (s ConfigXShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfigXShrinkRequest) GoString() string {
	return s.String()
}

func (s *ConfigXShrinkRequest) SetCallAbility(v string) *ConfigXShrinkRequest {
	s.CallAbility = &v
	return s
}

func (s *ConfigXShrinkRequest) SetCallerParentId(v int64) *ConfigXShrinkRequest {
	s.CallerParentId = &v
	return s
}

func (s *ConfigXShrinkRequest) SetCustomerPoolKey(v string) *ConfigXShrinkRequest {
	s.CustomerPoolKey = &v
	return s
}

func (s *ConfigXShrinkRequest) SetGNFlag(v string) *ConfigXShrinkRequest {
	s.GNFlag = &v
	return s
}

func (s *ConfigXShrinkRequest) SetOwnerId(v int64) *ConfigXShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *ConfigXShrinkRequest) SetReqId(v string) *ConfigXShrinkRequest {
	s.ReqId = &v
	return s
}

func (s *ConfigXShrinkRequest) SetResourceOwnerAccount(v string) *ConfigXShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ConfigXShrinkRequest) SetResourceOwnerId(v int64) *ConfigXShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ConfigXShrinkRequest) SetSequenceCallsShrink(v string) *ConfigXShrinkRequest {
	s.SequenceCallsShrink = &v
	return s
}

func (s *ConfigXShrinkRequest) SetSequenceMode(v string) *ConfigXShrinkRequest {
	s.SequenceMode = &v
	return s
}

func (s *ConfigXShrinkRequest) SetSmsAbility(v string) *ConfigXShrinkRequest {
	s.SmsAbility = &v
	return s
}

func (s *ConfigXShrinkRequest) SetSmsSignMode(v string) *ConfigXShrinkRequest {
	s.SmsSignMode = &v
	return s
}

func (s *ConfigXShrinkRequest) SetTelX(v string) *ConfigXShrinkRequest {
	s.TelX = &v
	return s
}

type ConfigXResponseBody struct {
	// example:
	//
	// 0
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 0000
	Code      *string                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ConfigXResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ConfigXResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConfigXResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigXResponseBody) SetAccessDeniedDetail(v string) *ConfigXResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ConfigXResponseBody) SetCode(v string) *ConfigXResponseBody {
	s.Code = &v
	return s
}

func (s *ConfigXResponseBody) SetData(v *ConfigXResponseBodyData) *ConfigXResponseBody {
	s.Data = v
	return s
}

func (s *ConfigXResponseBody) SetMessage(v string) *ConfigXResponseBody {
	s.Message = &v
	return s
}

func (s *ConfigXResponseBody) SetRequestId(v string) *ConfigXResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfigXResponseBody) SetSuccess(v bool) *ConfigXResponseBody {
	s.Success = &v
	return s
}

type ConfigXResponseBodyData struct {
	// 返回状态码 0000表示成功 其他表示失败
	//
	// example:
	//
	// 0000
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 返回信息
	//
	// example:
	//
	// 成功
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 返回是否成功 true  表示成功 false表示失败
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ConfigXResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ConfigXResponseBodyData) GoString() string {
	return s.String()
}

func (s *ConfigXResponseBodyData) SetCode(v string) *ConfigXResponseBodyData {
	s.Code = &v
	return s
}

func (s *ConfigXResponseBodyData) SetMessage(v string) *ConfigXResponseBodyData {
	s.Message = &v
	return s
}

func (s *ConfigXResponseBodyData) SetSuccess(v bool) *ConfigXResponseBodyData {
	s.Success = &v
	return s
}

type ConfigXResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfigXResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfigXResponse) String() string {
	return tea.Prettify(s)
}

func (s ConfigXResponse) GoString() string {
	return s.String()
}

func (s *ConfigXResponse) SetHeaders(v map[string]*string) *ConfigXResponse {
	s.Headers = v
	return s
}

func (s *ConfigXResponse) SetStatusCode(v int32) *ConfigXResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigXResponse) SetBody(v *ConfigXResponseBody) *ConfigXResponse {
	s.Body = v
	return s
}

type CreateAxgGroupRequest struct {
	// The name of number group G. If the name of number group G is not specified, the ID of number group G is used as the name of number group G.
	//
	// >  The value must be 1 to 128 characters in length.
	//
	// example:
	//
	// Aliyun
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The phone numbers that you add to number group G. Separate multiple phone numbers with commas (,). A maximum of 200 phone numbers can be added to number group G.
	//
	// example:
	//
	// 1390000****,1380000****
	Numbers *string `json:"Numbers,omitempty" xml:"Numbers,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The key of the phone number pool. Log on to the [Phone Number Protection console](https://dypls.console.aliyun.com/dypls.htm#/account) and view the key of the phone number pool on the **Number Pool Management*	- page.
	//
	// This parameter is required.
	//
	// example:
	//
	// FC123456
	PoolKey *string `json:"PoolKey,omitempty" xml:"PoolKey,omitempty"`
	// The remarks of number group G. The value must be 0 to 256 characters in length.
	//
	// example:
	//
	// test
	Remark               *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateAxgGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAxgGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateAxgGroupRequest) SetName(v string) *CreateAxgGroupRequest {
	s.Name = &v
	return s
}

func (s *CreateAxgGroupRequest) SetNumbers(v string) *CreateAxgGroupRequest {
	s.Numbers = &v
	return s
}

func (s *CreateAxgGroupRequest) SetOwnerId(v int64) *CreateAxgGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateAxgGroupRequest) SetPoolKey(v string) *CreateAxgGroupRequest {
	s.PoolKey = &v
	return s
}

func (s *CreateAxgGroupRequest) SetRemark(v string) *CreateAxgGroupRequest {
	s.Remark = &v
	return s
}

func (s *CreateAxgGroupRequest) SetResourceOwnerAccount(v string) *CreateAxgGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateAxgGroupRequest) SetResourceOwnerId(v int64) *CreateAxgGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

type CreateAxgGroupResponseBody struct {
	// The response code.
	//
	// 	- The value OK indicates that the request was successful.
	//
	// 	- Other values indicate that the request failed. For more information, see [Error codes](https://help.aliyun.com/document_detail/109196.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of number group G. The value of this parameter is required when the [BindAxg](https://help.aliyun.com/document_detail/110249.html) operation is called to add an AXG binding.
	//
	// example:
	//
	// 2000000000001
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 635C0FDA-9EBC-43D7-B368-9F583C08A126
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAxgGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAxgGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAxgGroupResponseBody) SetCode(v string) *CreateAxgGroupResponseBody {
	s.Code = &v
	return s
}

func (s *CreateAxgGroupResponseBody) SetGroupId(v int64) *CreateAxgGroupResponseBody {
	s.GroupId = &v
	return s
}

func (s *CreateAxgGroupResponseBody) SetMessage(v string) *CreateAxgGroupResponseBody {
	s.Message = &v
	return s
}

func (s *CreateAxgGroupResponseBody) SetRequestId(v string) *CreateAxgGroupResponseBody {
	s.RequestId = &v
	return s
}

type CreateAxgGroupResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAxgGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAxgGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAxgGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateAxgGroupResponse) SetHeaders(v map[string]*string) *CreateAxgGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateAxgGroupResponse) SetStatusCode(v int32) *CreateAxgGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAxgGroupResponse) SetBody(v *CreateAxgGroupResponseBody) *CreateAxgGroupResponse {
	s.Body = v
	return s
}

type CreateFixedNoAReportRequest struct {
	// 所属a号码组id
	//
	// This parameter is required.
	//
	// example:
	//
	// 12
	ANoWhiteGroupId *int64 `json:"ANoWhiteGroupId,omitempty" xml:"ANoWhiteGroupId,omitempty"`
	// 姓名
	//
	// This parameter is required.
	//
	// example:
	//
	// 赵**
	CustName *string `json:"CustName,omitempty" xml:"CustName,omitempty"`
	// 经办人/法人电话
	//
	// example:
	//
	// 130*****8888
	CustPhoneNo *string `json:"CustPhoneNo,omitempty" xml:"CustPhoneNo,omitempty"`
	// 固话客户类型 1:法人,2:经办人
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CustType *int64 `json:"CustType,omitempty" xml:"CustType,omitempty"`
	// 证件号码
	//
	// This parameter is required.
	//
	// example:
	//
	// 370*********
	DocumentNumber *string `json:"DocumentNumber,omitempty" xml:"DocumentNumber,omitempty"`
	// 证件类型 填写1表示身份证
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DocumentType *int64 `json:"DocumentType,omitempty" xml:"DocumentType,omitempty"`
	// 固话申请时服务大厅工单号
	//
	// example:
	//
	// 45***12
	FixedLineWorkId *string `json:"FixedLineWorkId,omitempty" xml:"FixedLineWorkId,omitempty"`
	// a号码(固话)
	//
	// This parameter is required.
	//
	// example:
	//
	// 89*****1234
	FixedNoA *string `json:"FixedNoA,omitempty" xml:"FixedNoA,omitempty"`
	// 正面人像照片 OSS 路径地址
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456/test1719383196033.jpg
	IdCardAlivePhoto *string `json:"IdCardAlivePhoto,omitempty" xml:"IdCardAlivePhoto,omitempty"`
	// 身份证反面照片oss路径地址
	//
	// example:
	//
	// 123456/test1719383196032.jpg
	IdCardBackPhoto *string `json:"IdCardBackPhoto,omitempty" xml:"IdCardBackPhoto,omitempty"`
	// 身份证正面照片oss路径地址
	//
	// example:
	//
	// 123456/test1719383196031.jpg
	IdCardFrontPhoto *string `json:"IdCardFrontPhoto,omitempty" xml:"IdCardFrontPhoto,omitempty"`
	OwnerId          *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 备注（客户自己的业务备注，可编辑）
	//
	// example:
	//
	// ***报备
	Remark               *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateFixedNoAReportRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFixedNoAReportRequest) GoString() string {
	return s.String()
}

func (s *CreateFixedNoAReportRequest) SetANoWhiteGroupId(v int64) *CreateFixedNoAReportRequest {
	s.ANoWhiteGroupId = &v
	return s
}

func (s *CreateFixedNoAReportRequest) SetCustName(v string) *CreateFixedNoAReportRequest {
	s.CustName = &v
	return s
}

func (s *CreateFixedNoAReportRequest) SetCustPhoneNo(v string) *CreateFixedNoAReportRequest {
	s.CustPhoneNo = &v
	return s
}

func (s *CreateFixedNoAReportRequest) SetCustType(v int64) *CreateFixedNoAReportRequest {
	s.CustType = &v
	return s
}

func (s *CreateFixedNoAReportRequest) SetDocumentNumber(v string) *CreateFixedNoAReportRequest {
	s.DocumentNumber = &v
	return s
}

func (s *CreateFixedNoAReportRequest) SetDocumentType(v int64) *CreateFixedNoAReportRequest {
	s.DocumentType = &v
	return s
}

func (s *CreateFixedNoAReportRequest) SetFixedLineWorkId(v string) *CreateFixedNoAReportRequest {
	s.FixedLineWorkId = &v
	return s
}

func (s *CreateFixedNoAReportRequest) SetFixedNoA(v string) *CreateFixedNoAReportRequest {
	s.FixedNoA = &v
	return s
}

func (s *CreateFixedNoAReportRequest) SetIdCardAlivePhoto(v string) *CreateFixedNoAReportRequest {
	s.IdCardAlivePhoto = &v
	return s
}

func (s *CreateFixedNoAReportRequest) SetIdCardBackPhoto(v string) *CreateFixedNoAReportRequest {
	s.IdCardBackPhoto = &v
	return s
}

func (s *CreateFixedNoAReportRequest) SetIdCardFrontPhoto(v string) *CreateFixedNoAReportRequest {
	s.IdCardFrontPhoto = &v
	return s
}

func (s *CreateFixedNoAReportRequest) SetOwnerId(v int64) *CreateFixedNoAReportRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateFixedNoAReportRequest) SetRemark(v string) *CreateFixedNoAReportRequest {
	s.Remark = &v
	return s
}

func (s *CreateFixedNoAReportRequest) SetResourceOwnerAccount(v string) *CreateFixedNoAReportRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateFixedNoAReportRequest) SetResourceOwnerId(v int64) *CreateFixedNoAReportRequest {
	s.ResourceOwnerId = &v
	return s
}

type CreateFixedNoAReportResponseBody struct {
	// example:
	//
	// 0
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// 请求状态码
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// A号码报备结果结构体
	Data *CreateFixedNoAReportResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// 失败错误提示
	//
	// example:
	//
	// 示例值示例值
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 返回id
	//
	// example:
	//
	// 1D73E648-0978-18A5-B089-3BB2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 请求是否成功
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateFixedNoAReportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFixedNoAReportResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFixedNoAReportResponseBody) SetAccessDeniedDetail(v string) *CreateFixedNoAReportResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CreateFixedNoAReportResponseBody) SetCode(v string) *CreateFixedNoAReportResponseBody {
	s.Code = &v
	return s
}

func (s *CreateFixedNoAReportResponseBody) SetData(v *CreateFixedNoAReportResponseBodyData) *CreateFixedNoAReportResponseBody {
	s.Data = v
	return s
}

func (s *CreateFixedNoAReportResponseBody) SetMessage(v string) *CreateFixedNoAReportResponseBody {
	s.Message = &v
	return s
}

func (s *CreateFixedNoAReportResponseBody) SetRequestId(v string) *CreateFixedNoAReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFixedNoAReportResponseBody) SetSuccess(v bool) *CreateFixedNoAReportResponseBody {
	s.Success = &v
	return s
}

type CreateFixedNoAReportResponseBodyData struct {
	// 创建结果
	//
	// example:
	//
	// true
	CreateResult *bool `json:"CreateResult,omitempty" xml:"CreateResult,omitempty"`
	// 创建类型枚举，1为成功，负数为创建失败
	//
	// example:
	//
	// 1
	FailType *int64 `json:"FailType,omitempty" xml:"FailType,omitempty"`
}

func (s CreateFixedNoAReportResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateFixedNoAReportResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateFixedNoAReportResponseBodyData) SetCreateResult(v bool) *CreateFixedNoAReportResponseBodyData {
	s.CreateResult = &v
	return s
}

func (s *CreateFixedNoAReportResponseBodyData) SetFailType(v int64) *CreateFixedNoAReportResponseBodyData {
	s.FailType = &v
	return s
}

type CreateFixedNoAReportResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateFixedNoAReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateFixedNoAReportResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFixedNoAReportResponse) GoString() string {
	return s.String()
}

func (s *CreateFixedNoAReportResponse) SetHeaders(v map[string]*string) *CreateFixedNoAReportResponse {
	s.Headers = v
	return s
}

func (s *CreateFixedNoAReportResponse) SetStatusCode(v int32) *CreateFixedNoAReportResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFixedNoAReportResponse) SetBody(v *CreateFixedNoAReportResponseBody) *CreateFixedNoAReportResponse {
	s.Body = v
	return s
}

type CreatePhoneNoAReportRequest struct {
	// 所属a号码组id
	//
	// This parameter is required.
	//
	// example:
	//
	// 19
	ANoWhiteGroupId *int64 `json:"ANoWhiteGroupId,omitempty" xml:"ANoWhiteGroupId,omitempty"`
	// 姓名
	//
	// This parameter is required.
	//
	// example:
	//
	// 赵**
	CustName *string `json:"CustName,omitempty" xml:"CustName,omitempty"`
	// 证件号码
	//
	// This parameter is required.
	//
	// example:
	//
	// 370*********
	DocumentNumber *string `json:"DocumentNumber,omitempty" xml:"DocumentNumber,omitempty"`
	// 证件类型 填写1表示身份证
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DocumentType *int64 `json:"DocumentType,omitempty" xml:"DocumentType,omitempty"`
	// 半身照oss路径地址
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456/test1719383196033.jpg示例值
	IdCardAlivePhoto *string `json:"IdCardAlivePhoto,omitempty" xml:"IdCardAlivePhoto,omitempty"`
	// 身份证反面照片oss路径地址
	//
	// example:
	//
	// 123456/test1719383196032.jpg
	IdCardBackPhoto *string `json:"IdCardBackPhoto,omitempty" xml:"IdCardBackPhoto,omitempty"`
	// 身份证正面照片oss路径地址
	//
	// example:
	//
	// 123456/test1719383196031.jpg
	IdCardFrontPhoto *string `json:"IdCardFrontPhoto,omitempty" xml:"IdCardFrontPhoto,omitempty"`
	OwnerId          *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 手机号
	//
	// This parameter is required.
	//
	// example:
	//
	// 130*****1234
	PhoneNoA *string `json:"PhoneNoA,omitempty" xml:"PhoneNoA,omitempty"`
	// 备注（客户自己的业务备注，可编辑）
	//
	// example:
	//
	// ***报备
	Remark               *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreatePhoneNoAReportRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePhoneNoAReportRequest) GoString() string {
	return s.String()
}

func (s *CreatePhoneNoAReportRequest) SetANoWhiteGroupId(v int64) *CreatePhoneNoAReportRequest {
	s.ANoWhiteGroupId = &v
	return s
}

func (s *CreatePhoneNoAReportRequest) SetCustName(v string) *CreatePhoneNoAReportRequest {
	s.CustName = &v
	return s
}

func (s *CreatePhoneNoAReportRequest) SetDocumentNumber(v string) *CreatePhoneNoAReportRequest {
	s.DocumentNumber = &v
	return s
}

func (s *CreatePhoneNoAReportRequest) SetDocumentType(v int64) *CreatePhoneNoAReportRequest {
	s.DocumentType = &v
	return s
}

func (s *CreatePhoneNoAReportRequest) SetIdCardAlivePhoto(v string) *CreatePhoneNoAReportRequest {
	s.IdCardAlivePhoto = &v
	return s
}

func (s *CreatePhoneNoAReportRequest) SetIdCardBackPhoto(v string) *CreatePhoneNoAReportRequest {
	s.IdCardBackPhoto = &v
	return s
}

func (s *CreatePhoneNoAReportRequest) SetIdCardFrontPhoto(v string) *CreatePhoneNoAReportRequest {
	s.IdCardFrontPhoto = &v
	return s
}

func (s *CreatePhoneNoAReportRequest) SetOwnerId(v int64) *CreatePhoneNoAReportRequest {
	s.OwnerId = &v
	return s
}

func (s *CreatePhoneNoAReportRequest) SetPhoneNoA(v string) *CreatePhoneNoAReportRequest {
	s.PhoneNoA = &v
	return s
}

func (s *CreatePhoneNoAReportRequest) SetRemark(v string) *CreatePhoneNoAReportRequest {
	s.Remark = &v
	return s
}

func (s *CreatePhoneNoAReportRequest) SetResourceOwnerAccount(v string) *CreatePhoneNoAReportRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreatePhoneNoAReportRequest) SetResourceOwnerId(v int64) *CreatePhoneNoAReportRequest {
	s.ResourceOwnerId = &v
	return s
}

type CreatePhoneNoAReportResponseBody struct {
	// example:
	//
	// 0
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// 请求状态码
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// A号码报备结果结构体
	Data *CreatePhoneNoAReportResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// 失败错误提示
	//
	// example:
	//
	// 手机号码***已存在
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 返回id
	//
	// example:
	//
	// 1D73E648-0978-18A5-B089-3BB2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 请求是否成功
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreatePhoneNoAReportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreatePhoneNoAReportResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePhoneNoAReportResponseBody) SetAccessDeniedDetail(v string) *CreatePhoneNoAReportResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CreatePhoneNoAReportResponseBody) SetCode(v string) *CreatePhoneNoAReportResponseBody {
	s.Code = &v
	return s
}

func (s *CreatePhoneNoAReportResponseBody) SetData(v *CreatePhoneNoAReportResponseBodyData) *CreatePhoneNoAReportResponseBody {
	s.Data = v
	return s
}

func (s *CreatePhoneNoAReportResponseBody) SetMessage(v string) *CreatePhoneNoAReportResponseBody {
	s.Message = &v
	return s
}

func (s *CreatePhoneNoAReportResponseBody) SetRequestId(v string) *CreatePhoneNoAReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePhoneNoAReportResponseBody) SetSuccess(v bool) *CreatePhoneNoAReportResponseBody {
	s.Success = &v
	return s
}

type CreatePhoneNoAReportResponseBodyData struct {
	// 创建结果
	//
	// example:
	//
	// true
	CreateResult *bool `json:"CreateResult,omitempty" xml:"CreateResult,omitempty"`
	// 创建类型枚举，1为成功，负数为创建失败
	//
	// example:
	//
	// 1
	FailType *int64 `json:"FailType,omitempty" xml:"FailType,omitempty"`
}

func (s CreatePhoneNoAReportResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreatePhoneNoAReportResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreatePhoneNoAReportResponseBodyData) SetCreateResult(v bool) *CreatePhoneNoAReportResponseBodyData {
	s.CreateResult = &v
	return s
}

func (s *CreatePhoneNoAReportResponseBodyData) SetFailType(v int64) *CreatePhoneNoAReportResponseBodyData {
	s.FailType = &v
	return s
}

type CreatePhoneNoAReportResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePhoneNoAReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePhoneNoAReportResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePhoneNoAReportResponse) GoString() string {
	return s.String()
}

func (s *CreatePhoneNoAReportResponse) SetHeaders(v map[string]*string) *CreatePhoneNoAReportResponse {
	s.Headers = v
	return s
}

func (s *CreatePhoneNoAReportResponse) SetStatusCode(v int32) *CreatePhoneNoAReportResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePhoneNoAReportResponse) SetBody(v *CreatePhoneNoAReportResponseBody) *CreatePhoneNoAReportResponse {
	s.Body = v
	return s
}

type CreatePickUpWaybillRequest struct {
	// The end time of the door-to-door pickup in the appointment. The value of **AppointGotEndTime*	- is the value of **EndTime*	- of **AppointTimes*	- in **CpTimeSelectList*	- returned by the [CreatePickUpWaybillPreQuery](~~CreatePickUpWaybillPreQuery~~#resultMapping) operation.
	//
	// >  This parameter is required when **BizType*	- is set to **1**.
	//
	// example:
	//
	// 2021-01-01 12:00:00
	AppointGotEndTime *string `json:"AppointGotEndTime,omitempty" xml:"AppointGotEndTime,omitempty"`
	// The start time of the door-to-door pickup in the appointment. The value of **AppointGotStartTime*	- is the value of **StartTime*	- of **AppointTimes*	- in **CpTimeSelectList*	- returned by the [CreatePickUpWaybillPreQuery](~~CreatePickUpWaybillPreQuery~~#resultMapping) operation.
	//
	// >  This parameter is required when **BizType*	- is set to **1**.
	//
	// example:
	//
	// 2021-01-01 10:00:00
	AppointGotStartTime *string `json:"AppointGotStartTime,omitempty" xml:"AppointGotStartTime,omitempty"`
	// The pickup mode. Valid values:
	//
	// 	- **0*	- (default): real-time order.
	//
	// 	- **1**: appointment order.
	//
	// example:
	//
	// 0
	BizType *int32 `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// The address of the consignee.
	//
	// This parameter is required.
	ConsigneeAddress *CreatePickUpWaybillRequestConsigneeAddress `json:"ConsigneeAddress,omitempty" xml:"ConsigneeAddress,omitempty" type:"Struct"`
	// The mobile phone number of the consignee.
	//
	// >  Either ConsigneeMobile or ConsigneePhone must be set.
	//
	// example:
	//
	// 1580000****
	ConsigneeMobile *string `json:"ConsigneeMobile,omitempty" xml:"ConsigneeMobile,omitempty"`
	// The name of the consignee.
	//
	// This parameter is required.
	//
	// example:
	//
	// Li
	ConsigneeName *string `json:"ConsigneeName,omitempty" xml:"ConsigneeName,omitempty"`
	// The landline phone number of the consignee.
	//
	// >  Either ConsigneeMobile or ConsigneePhone must be set.
	//
	// example:
	//
	// 0570000****
	ConsigneePhone *string `json:"ConsigneePhone,omitempty" xml:"ConsigneePhone,omitempty"`
	// The code of the courier company.
	//
	// example:
	//
	// YTO
	CpCode *string `json:"CpCode,omitempty" xml:"CpCode,omitempty"`
	// The items.
	GoodsInfos []*CreatePickUpWaybillRequestGoodsInfos `json:"GoodsInfos,omitempty" xml:"GoodsInfos,omitempty" type:"Repeated"`
	// The external channel sources.
	//
	// This parameter is required.
	//
	// example:
	//
	// YUN_DIAN_SHANG
	OrderChannels *string `json:"OrderChannels,omitempty" xml:"OrderChannels,omitempty"`
	// The ID of the external order.
	//
	// This parameter is required.
	//
	// example:
	//
	// 143234234266****
	OuterOrderCode *string `json:"OuterOrderCode,omitempty" xml:"OuterOrderCode,omitempty"`
	// The additional information about the order. The additional information will be printed on the order.
	//
	// example:
	//
	// fragile
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The address of the sender.
	//
	// This parameter is required.
	SendAddress *CreatePickUpWaybillRequestSendAddress `json:"SendAddress,omitempty" xml:"SendAddress,omitempty" type:"Struct"`
	// The mobile phone number of the sender.
	//
	// >  Either SendMobile or SendPhone must be set.
	//
	// example:
	//
	// 1596714****
	SendMobile *string `json:"SendMobile,omitempty" xml:"SendMobile,omitempty"`
	// The name of the sender.
	//
	// This parameter is required.
	//
	// example:
	//
	// Wang
	SendName *string `json:"SendName,omitempty" xml:"SendName,omitempty"`
	// The landline phone number of the sender.
	//
	// >  Either SendMobile or SendPhone must be set.
	//
	// example:
	//
	// 05718845****
	SendPhone *string `json:"SendPhone,omitempty" xml:"SendPhone,omitempty"`
}

func (s CreatePickUpWaybillRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePickUpWaybillRequest) GoString() string {
	return s.String()
}

func (s *CreatePickUpWaybillRequest) SetAppointGotEndTime(v string) *CreatePickUpWaybillRequest {
	s.AppointGotEndTime = &v
	return s
}

func (s *CreatePickUpWaybillRequest) SetAppointGotStartTime(v string) *CreatePickUpWaybillRequest {
	s.AppointGotStartTime = &v
	return s
}

func (s *CreatePickUpWaybillRequest) SetBizType(v int32) *CreatePickUpWaybillRequest {
	s.BizType = &v
	return s
}

func (s *CreatePickUpWaybillRequest) SetConsigneeAddress(v *CreatePickUpWaybillRequestConsigneeAddress) *CreatePickUpWaybillRequest {
	s.ConsigneeAddress = v
	return s
}

func (s *CreatePickUpWaybillRequest) SetConsigneeMobile(v string) *CreatePickUpWaybillRequest {
	s.ConsigneeMobile = &v
	return s
}

func (s *CreatePickUpWaybillRequest) SetConsigneeName(v string) *CreatePickUpWaybillRequest {
	s.ConsigneeName = &v
	return s
}

func (s *CreatePickUpWaybillRequest) SetConsigneePhone(v string) *CreatePickUpWaybillRequest {
	s.ConsigneePhone = &v
	return s
}

func (s *CreatePickUpWaybillRequest) SetCpCode(v string) *CreatePickUpWaybillRequest {
	s.CpCode = &v
	return s
}

func (s *CreatePickUpWaybillRequest) SetGoodsInfos(v []*CreatePickUpWaybillRequestGoodsInfos) *CreatePickUpWaybillRequest {
	s.GoodsInfos = v
	return s
}

func (s *CreatePickUpWaybillRequest) SetOrderChannels(v string) *CreatePickUpWaybillRequest {
	s.OrderChannels = &v
	return s
}

func (s *CreatePickUpWaybillRequest) SetOuterOrderCode(v string) *CreatePickUpWaybillRequest {
	s.OuterOrderCode = &v
	return s
}

func (s *CreatePickUpWaybillRequest) SetRemark(v string) *CreatePickUpWaybillRequest {
	s.Remark = &v
	return s
}

func (s *CreatePickUpWaybillRequest) SetSendAddress(v *CreatePickUpWaybillRequestSendAddress) *CreatePickUpWaybillRequest {
	s.SendAddress = v
	return s
}

func (s *CreatePickUpWaybillRequest) SetSendMobile(v string) *CreatePickUpWaybillRequest {
	s.SendMobile = &v
	return s
}

func (s *CreatePickUpWaybillRequest) SetSendName(v string) *CreatePickUpWaybillRequest {
	s.SendName = &v
	return s
}

func (s *CreatePickUpWaybillRequest) SetSendPhone(v string) *CreatePickUpWaybillRequest {
	s.SendPhone = &v
	return s
}

type CreatePickUpWaybillRequestConsigneeAddress struct {
	// The detailed address of the consignee.
	//
	// This parameter is required.
	//
	// example:
	//
	// XX community
	AddressDetail *string `json:"AddressDetail,omitempty" xml:"AddressDetail,omitempty"`
	// The district where the consignee is located.
	//
	// This parameter is required.
	//
	// example:
	//
	// xihu
	AreaName *string `json:"AreaName,omitempty" xml:"AreaName,omitempty"`
	// The city where the consignee is located.
	//
	// This parameter is required.
	//
	// example:
	//
	// hangzhou
	CityName *string `json:"CityName,omitempty" xml:"CityName,omitempty"`
	// The province where the consignee is located.
	//
	// This parameter is required.
	//
	// example:
	//
	// zhejiang
	ProvinceName *string `json:"ProvinceName,omitempty" xml:"ProvinceName,omitempty"`
	// The street where the consignee is located.
	//
	// example:
	//
	// XX Street
	TownName *string `json:"TownName,omitempty" xml:"TownName,omitempty"`
}

func (s CreatePickUpWaybillRequestConsigneeAddress) String() string {
	return tea.Prettify(s)
}

func (s CreatePickUpWaybillRequestConsigneeAddress) GoString() string {
	return s.String()
}

func (s *CreatePickUpWaybillRequestConsigneeAddress) SetAddressDetail(v string) *CreatePickUpWaybillRequestConsigneeAddress {
	s.AddressDetail = &v
	return s
}

func (s *CreatePickUpWaybillRequestConsigneeAddress) SetAreaName(v string) *CreatePickUpWaybillRequestConsigneeAddress {
	s.AreaName = &v
	return s
}

func (s *CreatePickUpWaybillRequestConsigneeAddress) SetCityName(v string) *CreatePickUpWaybillRequestConsigneeAddress {
	s.CityName = &v
	return s
}

func (s *CreatePickUpWaybillRequestConsigneeAddress) SetProvinceName(v string) *CreatePickUpWaybillRequestConsigneeAddress {
	s.ProvinceName = &v
	return s
}

func (s *CreatePickUpWaybillRequestConsigneeAddress) SetTownName(v string) *CreatePickUpWaybillRequestConsigneeAddress {
	s.TownName = &v
	return s
}

type CreatePickUpWaybillRequestGoodsInfos struct {
	// The item name.
	//
	// example:
	//
	// zhang
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The item quantity.
	//
	// example:
	//
	// 1
	Quantity *string `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	// The item weight. Unit: gram.
	//
	// example:
	//
	// 1000
	Weight *string `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s CreatePickUpWaybillRequestGoodsInfos) String() string {
	return tea.Prettify(s)
}

func (s CreatePickUpWaybillRequestGoodsInfos) GoString() string {
	return s.String()
}

func (s *CreatePickUpWaybillRequestGoodsInfos) SetName(v string) *CreatePickUpWaybillRequestGoodsInfos {
	s.Name = &v
	return s
}

func (s *CreatePickUpWaybillRequestGoodsInfos) SetQuantity(v string) *CreatePickUpWaybillRequestGoodsInfos {
	s.Quantity = &v
	return s
}

func (s *CreatePickUpWaybillRequestGoodsInfos) SetWeight(v string) *CreatePickUpWaybillRequestGoodsInfos {
	s.Weight = &v
	return s
}

type CreatePickUpWaybillRequestSendAddress struct {
	// The detailed address of the sender.
	//
	// This parameter is required.
	//
	// example:
	//
	// XX community
	AddressDetail *string `json:"AddressDetail,omitempty" xml:"AddressDetail,omitempty"`
	// The district where the sender is located.
	//
	// This parameter is required.
	//
	// example:
	//
	// wenjiang
	AreaName *string `json:"AreaName,omitempty" xml:"AreaName,omitempty"`
	// The city where the sender is located.
	//
	// This parameter is required.
	//
	// example:
	//
	// chengdu
	CityName *string `json:"CityName,omitempty" xml:"CityName,omitempty"`
	// The province where the sender is located.
	//
	// This parameter is required.
	//
	// example:
	//
	// Sichuan
	ProvinceName *string `json:"ProvinceName,omitempty" xml:"ProvinceName,omitempty"`
	// The street where the sender is located.
	//
	// example:
	//
	// XX Street
	TownName *string `json:"TownName,omitempty" xml:"TownName,omitempty"`
}

func (s CreatePickUpWaybillRequestSendAddress) String() string {
	return tea.Prettify(s)
}

func (s CreatePickUpWaybillRequestSendAddress) GoString() string {
	return s.String()
}

func (s *CreatePickUpWaybillRequestSendAddress) SetAddressDetail(v string) *CreatePickUpWaybillRequestSendAddress {
	s.AddressDetail = &v
	return s
}

func (s *CreatePickUpWaybillRequestSendAddress) SetAreaName(v string) *CreatePickUpWaybillRequestSendAddress {
	s.AreaName = &v
	return s
}

func (s *CreatePickUpWaybillRequestSendAddress) SetCityName(v string) *CreatePickUpWaybillRequestSendAddress {
	s.CityName = &v
	return s
}

func (s *CreatePickUpWaybillRequestSendAddress) SetProvinceName(v string) *CreatePickUpWaybillRequestSendAddress {
	s.ProvinceName = &v
	return s
}

func (s *CreatePickUpWaybillRequestSendAddress) SetTownName(v string) *CreatePickUpWaybillRequestSendAddress {
	s.TownName = &v
	return s
}

type CreatePickUpWaybillShrinkRequest struct {
	// The end time of the door-to-door pickup in the appointment. The value of **AppointGotEndTime*	- is the value of **EndTime*	- of **AppointTimes*	- in **CpTimeSelectList*	- returned by the [CreatePickUpWaybillPreQuery](~~CreatePickUpWaybillPreQuery~~#resultMapping) operation.
	//
	// >  This parameter is required when **BizType*	- is set to **1**.
	//
	// example:
	//
	// 2021-01-01 12:00:00
	AppointGotEndTime *string `json:"AppointGotEndTime,omitempty" xml:"AppointGotEndTime,omitempty"`
	// The start time of the door-to-door pickup in the appointment. The value of **AppointGotStartTime*	- is the value of **StartTime*	- of **AppointTimes*	- in **CpTimeSelectList*	- returned by the [CreatePickUpWaybillPreQuery](~~CreatePickUpWaybillPreQuery~~#resultMapping) operation.
	//
	// >  This parameter is required when **BizType*	- is set to **1**.
	//
	// example:
	//
	// 2021-01-01 10:00:00
	AppointGotStartTime *string `json:"AppointGotStartTime,omitempty" xml:"AppointGotStartTime,omitempty"`
	// The pickup mode. Valid values:
	//
	// 	- **0*	- (default): real-time order.
	//
	// 	- **1**: appointment order.
	//
	// example:
	//
	// 0
	BizType *int32 `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// The address of the consignee.
	//
	// This parameter is required.
	ConsigneeAddressShrink *string `json:"ConsigneeAddress,omitempty" xml:"ConsigneeAddress,omitempty"`
	// The mobile phone number of the consignee.
	//
	// >  Either ConsigneeMobile or ConsigneePhone must be set.
	//
	// example:
	//
	// 1580000****
	ConsigneeMobile *string `json:"ConsigneeMobile,omitempty" xml:"ConsigneeMobile,omitempty"`
	// The name of the consignee.
	//
	// This parameter is required.
	//
	// example:
	//
	// Li
	ConsigneeName *string `json:"ConsigneeName,omitempty" xml:"ConsigneeName,omitempty"`
	// The landline phone number of the consignee.
	//
	// >  Either ConsigneeMobile or ConsigneePhone must be set.
	//
	// example:
	//
	// 0570000****
	ConsigneePhone *string `json:"ConsigneePhone,omitempty" xml:"ConsigneePhone,omitempty"`
	// The code of the courier company.
	//
	// example:
	//
	// YTO
	CpCode *string `json:"CpCode,omitempty" xml:"CpCode,omitempty"`
	// The items.
	GoodsInfosShrink *string `json:"GoodsInfos,omitempty" xml:"GoodsInfos,omitempty"`
	// The external channel sources.
	//
	// This parameter is required.
	//
	// example:
	//
	// YUN_DIAN_SHANG
	OrderChannels *string `json:"OrderChannels,omitempty" xml:"OrderChannels,omitempty"`
	// The ID of the external order.
	//
	// This parameter is required.
	//
	// example:
	//
	// 143234234266****
	OuterOrderCode *string `json:"OuterOrderCode,omitempty" xml:"OuterOrderCode,omitempty"`
	// The additional information about the order. The additional information will be printed on the order.
	//
	// example:
	//
	// fragile
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The address of the sender.
	//
	// This parameter is required.
	SendAddressShrink *string `json:"SendAddress,omitempty" xml:"SendAddress,omitempty"`
	// The mobile phone number of the sender.
	//
	// >  Either SendMobile or SendPhone must be set.
	//
	// example:
	//
	// 1596714****
	SendMobile *string `json:"SendMobile,omitempty" xml:"SendMobile,omitempty"`
	// The name of the sender.
	//
	// This parameter is required.
	//
	// example:
	//
	// Wang
	SendName *string `json:"SendName,omitempty" xml:"SendName,omitempty"`
	// The landline phone number of the sender.
	//
	// >  Either SendMobile or SendPhone must be set.
	//
	// example:
	//
	// 05718845****
	SendPhone *string `json:"SendPhone,omitempty" xml:"SendPhone,omitempty"`
}

func (s CreatePickUpWaybillShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePickUpWaybillShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreatePickUpWaybillShrinkRequest) SetAppointGotEndTime(v string) *CreatePickUpWaybillShrinkRequest {
	s.AppointGotEndTime = &v
	return s
}

func (s *CreatePickUpWaybillShrinkRequest) SetAppointGotStartTime(v string) *CreatePickUpWaybillShrinkRequest {
	s.AppointGotStartTime = &v
	return s
}

func (s *CreatePickUpWaybillShrinkRequest) SetBizType(v int32) *CreatePickUpWaybillShrinkRequest {
	s.BizType = &v
	return s
}

func (s *CreatePickUpWaybillShrinkRequest) SetConsigneeAddressShrink(v string) *CreatePickUpWaybillShrinkRequest {
	s.ConsigneeAddressShrink = &v
	return s
}

func (s *CreatePickUpWaybillShrinkRequest) SetConsigneeMobile(v string) *CreatePickUpWaybillShrinkRequest {
	s.ConsigneeMobile = &v
	return s
}

func (s *CreatePickUpWaybillShrinkRequest) SetConsigneeName(v string) *CreatePickUpWaybillShrinkRequest {
	s.ConsigneeName = &v
	return s
}

func (s *CreatePickUpWaybillShrinkRequest) SetConsigneePhone(v string) *CreatePickUpWaybillShrinkRequest {
	s.ConsigneePhone = &v
	return s
}

func (s *CreatePickUpWaybillShrinkRequest) SetCpCode(v string) *CreatePickUpWaybillShrinkRequest {
	s.CpCode = &v
	return s
}

func (s *CreatePickUpWaybillShrinkRequest) SetGoodsInfosShrink(v string) *CreatePickUpWaybillShrinkRequest {
	s.GoodsInfosShrink = &v
	return s
}

func (s *CreatePickUpWaybillShrinkRequest) SetOrderChannels(v string) *CreatePickUpWaybillShrinkRequest {
	s.OrderChannels = &v
	return s
}

func (s *CreatePickUpWaybillShrinkRequest) SetOuterOrderCode(v string) *CreatePickUpWaybillShrinkRequest {
	s.OuterOrderCode = &v
	return s
}

func (s *CreatePickUpWaybillShrinkRequest) SetRemark(v string) *CreatePickUpWaybillShrinkRequest {
	s.Remark = &v
	return s
}

func (s *CreatePickUpWaybillShrinkRequest) SetSendAddressShrink(v string) *CreatePickUpWaybillShrinkRequest {
	s.SendAddressShrink = &v
	return s
}

func (s *CreatePickUpWaybillShrinkRequest) SetSendMobile(v string) *CreatePickUpWaybillShrinkRequest {
	s.SendMobile = &v
	return s
}

func (s *CreatePickUpWaybillShrinkRequest) SetSendName(v string) *CreatePickUpWaybillShrinkRequest {
	s.SendName = &v
	return s
}

func (s *CreatePickUpWaybillShrinkRequest) SetSendPhone(v string) *CreatePickUpWaybillShrinkRequest {
	s.SendPhone = &v
	return s
}

type CreatePickUpWaybillResponseBody struct {
	// The returned data.
	Data *CreatePickUpWaybillResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9FC30594-3841-43AD-9008-03393BCB5CD2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreatePickUpWaybillResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreatePickUpWaybillResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePickUpWaybillResponseBody) SetData(v *CreatePickUpWaybillResponseBodyData) *CreatePickUpWaybillResponseBody {
	s.Data = v
	return s
}

func (s *CreatePickUpWaybillResponseBody) SetHttpStatusCode(v int32) *CreatePickUpWaybillResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreatePickUpWaybillResponseBody) SetMessage(v string) *CreatePickUpWaybillResponseBody {
	s.Message = &v
	return s
}

func (s *CreatePickUpWaybillResponseBody) SetRequestId(v string) *CreatePickUpWaybillResponseBody {
	s.RequestId = &v
	return s
}

type CreatePickUpWaybillResponseBodyData struct {
	// The code of the courier company.
	//
	// example:
	//
	// YTO
	CpCode *string `json:"CpCode,omitempty" xml:"CpCode,omitempty"`
	// The error code.
	//
	// example:
	//
	// 200
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// none
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// The pickup code.
	//
	// example:
	//
	// 3524
	GotCode *string `json:"GotCode,omitempty" xml:"GotCode,omitempty"`
	// The order ID.
	//
	// example:
	//
	// 77312345629****
	MailNo *string `json:"MailNo,omitempty" xml:"MailNo,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreatePickUpWaybillResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreatePickUpWaybillResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreatePickUpWaybillResponseBodyData) SetCpCode(v string) *CreatePickUpWaybillResponseBodyData {
	s.CpCode = &v
	return s
}

func (s *CreatePickUpWaybillResponseBodyData) SetErrorCode(v string) *CreatePickUpWaybillResponseBodyData {
	s.ErrorCode = &v
	return s
}

func (s *CreatePickUpWaybillResponseBodyData) SetErrorMsg(v string) *CreatePickUpWaybillResponseBodyData {
	s.ErrorMsg = &v
	return s
}

func (s *CreatePickUpWaybillResponseBodyData) SetGotCode(v string) *CreatePickUpWaybillResponseBodyData {
	s.GotCode = &v
	return s
}

func (s *CreatePickUpWaybillResponseBodyData) SetMailNo(v string) *CreatePickUpWaybillResponseBodyData {
	s.MailNo = &v
	return s
}

func (s *CreatePickUpWaybillResponseBodyData) SetSuccess(v string) *CreatePickUpWaybillResponseBodyData {
	s.Success = &v
	return s
}

type CreatePickUpWaybillResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePickUpWaybillResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePickUpWaybillResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePickUpWaybillResponse) GoString() string {
	return s.String()
}

func (s *CreatePickUpWaybillResponse) SetHeaders(v map[string]*string) *CreatePickUpWaybillResponse {
	s.Headers = v
	return s
}

func (s *CreatePickUpWaybillResponse) SetStatusCode(v int32) *CreatePickUpWaybillResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePickUpWaybillResponse) SetBody(v *CreatePickUpWaybillResponseBody) *CreatePickUpWaybillResponse {
	s.Body = v
	return s
}

type CreatePickUpWaybillPreQueryRequest struct {
	// The consignee information.
	//
	// This parameter is required.
	ConsigneeInfo *CreatePickUpWaybillPreQueryRequestConsigneeInfo `json:"ConsigneeInfo,omitempty" xml:"ConsigneeInfo,omitempty" type:"Struct"`
	// The code of the courier company. If no courier company is specified, the system allocates a courier company.
	//
	// example:
	//
	// YTO
	CpCode *string `json:"CpCode,omitempty" xml:"CpCode,omitempty"`
	// The identifier of the external channel source. It cannot contain underscores.
	//
	// This parameter is required.
	//
	// example:
	//
	// Test
	OrderChannels *string `json:"OrderChannels,omitempty" xml:"OrderChannels,omitempty"`
	// The order number of the access system.
	//
	// example:
	//
	// 787DFHHDS989****
	OuterOrderCode *string `json:"OuterOrderCode,omitempty" xml:"OuterOrderCode,omitempty"`
	// The estimated weight. Unit: gram.
	//
	// >  If you need to query the estimated price, this parameter is required.
	//
	// example:
	//
	// 2000
	PreWeight *string `json:"PreWeight,omitempty" xml:"PreWeight,omitempty"`
	// The sender information.
	//
	// This parameter is required.
	SenderInfo *CreatePickUpWaybillPreQueryRequestSenderInfo `json:"SenderInfo,omitempty" xml:"SenderInfo,omitempty" type:"Struct"`
}

func (s CreatePickUpWaybillPreQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePickUpWaybillPreQueryRequest) GoString() string {
	return s.String()
}

func (s *CreatePickUpWaybillPreQueryRequest) SetConsigneeInfo(v *CreatePickUpWaybillPreQueryRequestConsigneeInfo) *CreatePickUpWaybillPreQueryRequest {
	s.ConsigneeInfo = v
	return s
}

func (s *CreatePickUpWaybillPreQueryRequest) SetCpCode(v string) *CreatePickUpWaybillPreQueryRequest {
	s.CpCode = &v
	return s
}

func (s *CreatePickUpWaybillPreQueryRequest) SetOrderChannels(v string) *CreatePickUpWaybillPreQueryRequest {
	s.OrderChannels = &v
	return s
}

func (s *CreatePickUpWaybillPreQueryRequest) SetOuterOrderCode(v string) *CreatePickUpWaybillPreQueryRequest {
	s.OuterOrderCode = &v
	return s
}

func (s *CreatePickUpWaybillPreQueryRequest) SetPreWeight(v string) *CreatePickUpWaybillPreQueryRequest {
	s.PreWeight = &v
	return s
}

func (s *CreatePickUpWaybillPreQueryRequest) SetSenderInfo(v *CreatePickUpWaybillPreQueryRequestSenderInfo) *CreatePickUpWaybillPreQueryRequest {
	s.SenderInfo = v
	return s
}

type CreatePickUpWaybillPreQueryRequestConsigneeInfo struct {
	// The address of the consignee.
	AddressInfo *CreatePickUpWaybillPreQueryRequestConsigneeInfoAddressInfo `json:"AddressInfo,omitempty" xml:"AddressInfo,omitempty" type:"Struct"`
	// The mobile phone number of the consignee.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1390000****
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// The name of the consignee.
	//
	// This parameter is required.
	//
	// example:
	//
	// Li
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreatePickUpWaybillPreQueryRequestConsigneeInfo) String() string {
	return tea.Prettify(s)
}

func (s CreatePickUpWaybillPreQueryRequestConsigneeInfo) GoString() string {
	return s.String()
}

func (s *CreatePickUpWaybillPreQueryRequestConsigneeInfo) SetAddressInfo(v *CreatePickUpWaybillPreQueryRequestConsigneeInfoAddressInfo) *CreatePickUpWaybillPreQueryRequestConsigneeInfo {
	s.AddressInfo = v
	return s
}

func (s *CreatePickUpWaybillPreQueryRequestConsigneeInfo) SetMobile(v string) *CreatePickUpWaybillPreQueryRequestConsigneeInfo {
	s.Mobile = &v
	return s
}

func (s *CreatePickUpWaybillPreQueryRequestConsigneeInfo) SetName(v string) *CreatePickUpWaybillPreQueryRequestConsigneeInfo {
	s.Name = &v
	return s
}

type CreatePickUpWaybillPreQueryRequestConsigneeInfoAddressInfo struct {
	// The detailed address of the consignee.
	//
	// example:
	//
	// XX community
	AddressDetail *string `json:"AddressDetail,omitempty" xml:"AddressDetail,omitempty"`
	// The district where the consignee is located.
	//
	// example:
	//
	// chang,an
	AreaName *string `json:"AreaName,omitempty" xml:"AreaName,omitempty"`
	// The city where the consignee is located.
	//
	// example:
	//
	// Xi,an
	CityName *string `json:"CityName,omitempty" xml:"CityName,omitempty"`
	// The province where the consignee is located.
	//
	// example:
	//
	// Shanxi
	ProvinceName *string `json:"ProvinceName,omitempty" xml:"ProvinceName,omitempty"`
	// The street where the consignee is located.
	//
	// example:
	//
	// XX Street
	TownName *string `json:"TownName,omitempty" xml:"TownName,omitempty"`
}

func (s CreatePickUpWaybillPreQueryRequestConsigneeInfoAddressInfo) String() string {
	return tea.Prettify(s)
}

func (s CreatePickUpWaybillPreQueryRequestConsigneeInfoAddressInfo) GoString() string {
	return s.String()
}

func (s *CreatePickUpWaybillPreQueryRequestConsigneeInfoAddressInfo) SetAddressDetail(v string) *CreatePickUpWaybillPreQueryRequestConsigneeInfoAddressInfo {
	s.AddressDetail = &v
	return s
}

func (s *CreatePickUpWaybillPreQueryRequestConsigneeInfoAddressInfo) SetAreaName(v string) *CreatePickUpWaybillPreQueryRequestConsigneeInfoAddressInfo {
	s.AreaName = &v
	return s
}

func (s *CreatePickUpWaybillPreQueryRequestConsigneeInfoAddressInfo) SetCityName(v string) *CreatePickUpWaybillPreQueryRequestConsigneeInfoAddressInfo {
	s.CityName = &v
	return s
}

func (s *CreatePickUpWaybillPreQueryRequestConsigneeInfoAddressInfo) SetProvinceName(v string) *CreatePickUpWaybillPreQueryRequestConsigneeInfoAddressInfo {
	s.ProvinceName = &v
	return s
}

func (s *CreatePickUpWaybillPreQueryRequestConsigneeInfoAddressInfo) SetTownName(v string) *CreatePickUpWaybillPreQueryRequestConsigneeInfoAddressInfo {
	s.TownName = &v
	return s
}

type CreatePickUpWaybillPreQueryRequestSenderInfo struct {
	// The address of the sender.
	AddressInfo *CreatePickUpWaybillPreQueryRequestSenderInfoAddressInfo `json:"AddressInfo,omitempty" xml:"AddressInfo,omitempty" type:"Struct"`
	// The mobile phone number of the sender.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1390000****
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// The name of the sender.
	//
	// This parameter is required.
	//
	// example:
	//
	// Wang
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreatePickUpWaybillPreQueryRequestSenderInfo) String() string {
	return tea.Prettify(s)
}

func (s CreatePickUpWaybillPreQueryRequestSenderInfo) GoString() string {
	return s.String()
}

func (s *CreatePickUpWaybillPreQueryRequestSenderInfo) SetAddressInfo(v *CreatePickUpWaybillPreQueryRequestSenderInfoAddressInfo) *CreatePickUpWaybillPreQueryRequestSenderInfo {
	s.AddressInfo = v
	return s
}

func (s *CreatePickUpWaybillPreQueryRequestSenderInfo) SetMobile(v string) *CreatePickUpWaybillPreQueryRequestSenderInfo {
	s.Mobile = &v
	return s
}

func (s *CreatePickUpWaybillPreQueryRequestSenderInfo) SetName(v string) *CreatePickUpWaybillPreQueryRequestSenderInfo {
	s.Name = &v
	return s
}

type CreatePickUpWaybillPreQueryRequestSenderInfoAddressInfo struct {
	// The detailed address of the sender.
	//
	// example:
	//
	// XX community
	AddressDetail *string `json:"AddressDetail,omitempty" xml:"AddressDetail,omitempty"`
	// The district where the sender is located.
	//
	// example:
	//
	// xihu
	AreaName *string `json:"AreaName,omitempty" xml:"AreaName,omitempty"`
	// The city where the sender is located.
	//
	// example:
	//
	// hangzhou
	CityName *string `json:"CityName,omitempty" xml:"CityName,omitempty"`
	// The province where the sender is located.
	//
	// example:
	//
	// zhejiang
	ProvinceName *string `json:"ProvinceName,omitempty" xml:"ProvinceName,omitempty"`
	// The street where the sender is located.
	//
	// example:
	//
	// XX Street
	TownName *string `json:"TownName,omitempty" xml:"TownName,omitempty"`
}

func (s CreatePickUpWaybillPreQueryRequestSenderInfoAddressInfo) String() string {
	return tea.Prettify(s)
}

func (s CreatePickUpWaybillPreQueryRequestSenderInfoAddressInfo) GoString() string {
	return s.String()
}

func (s *CreatePickUpWaybillPreQueryRequestSenderInfoAddressInfo) SetAddressDetail(v string) *CreatePickUpWaybillPreQueryRequestSenderInfoAddressInfo {
	s.AddressDetail = &v
	return s
}

func (s *CreatePickUpWaybillPreQueryRequestSenderInfoAddressInfo) SetAreaName(v string) *CreatePickUpWaybillPreQueryRequestSenderInfoAddressInfo {
	s.AreaName = &v
	return s
}

func (s *CreatePickUpWaybillPreQueryRequestSenderInfoAddressInfo) SetCityName(v string) *CreatePickUpWaybillPreQueryRequestSenderInfoAddressInfo {
	s.CityName = &v
	return s
}

func (s *CreatePickUpWaybillPreQueryRequestSenderInfoAddressInfo) SetProvinceName(v string) *CreatePickUpWaybillPreQueryRequestSenderInfoAddressInfo {
	s.ProvinceName = &v
	return s
}

func (s *CreatePickUpWaybillPreQueryRequestSenderInfoAddressInfo) SetTownName(v string) *CreatePickUpWaybillPreQueryRequestSenderInfoAddressInfo {
	s.TownName = &v
	return s
}

type CreatePickUpWaybillPreQueryShrinkRequest struct {
	// The consignee information.
	//
	// This parameter is required.
	ConsigneeInfoShrink *string `json:"ConsigneeInfo,omitempty" xml:"ConsigneeInfo,omitempty"`
	// The code of the courier company. If no courier company is specified, the system allocates a courier company.
	//
	// example:
	//
	// YTO
	CpCode *string `json:"CpCode,omitempty" xml:"CpCode,omitempty"`
	// The identifier of the external channel source. It cannot contain underscores.
	//
	// This parameter is required.
	//
	// example:
	//
	// Test
	OrderChannels *string `json:"OrderChannels,omitempty" xml:"OrderChannels,omitempty"`
	// The order number of the access system.
	//
	// example:
	//
	// 787DFHHDS989****
	OuterOrderCode *string `json:"OuterOrderCode,omitempty" xml:"OuterOrderCode,omitempty"`
	// The estimated weight. Unit: gram.
	//
	// >  If you need to query the estimated price, this parameter is required.
	//
	// example:
	//
	// 2000
	PreWeight *string `json:"PreWeight,omitempty" xml:"PreWeight,omitempty"`
	// The sender information.
	//
	// This parameter is required.
	SenderInfoShrink *string `json:"SenderInfo,omitempty" xml:"SenderInfo,omitempty"`
}

func (s CreatePickUpWaybillPreQueryShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePickUpWaybillPreQueryShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreatePickUpWaybillPreQueryShrinkRequest) SetConsigneeInfoShrink(v string) *CreatePickUpWaybillPreQueryShrinkRequest {
	s.ConsigneeInfoShrink = &v
	return s
}

func (s *CreatePickUpWaybillPreQueryShrinkRequest) SetCpCode(v string) *CreatePickUpWaybillPreQueryShrinkRequest {
	s.CpCode = &v
	return s
}

func (s *CreatePickUpWaybillPreQueryShrinkRequest) SetOrderChannels(v string) *CreatePickUpWaybillPreQueryShrinkRequest {
	s.OrderChannels = &v
	return s
}

func (s *CreatePickUpWaybillPreQueryShrinkRequest) SetOuterOrderCode(v string) *CreatePickUpWaybillPreQueryShrinkRequest {
	s.OuterOrderCode = &v
	return s
}

func (s *CreatePickUpWaybillPreQueryShrinkRequest) SetPreWeight(v string) *CreatePickUpWaybillPreQueryShrinkRequest {
	s.PreWeight = &v
	return s
}

func (s *CreatePickUpWaybillPreQueryShrinkRequest) SetSenderInfoShrink(v string) *CreatePickUpWaybillPreQueryShrinkRequest {
	s.SenderInfoShrink = &v
	return s
}

type CreatePickUpWaybillPreQueryResponseBody struct {
	// The result set.
	Data *CreatePickUpWaybillPreQueryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9FC30594-3841-43AD-9008-03393BCB5CD2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreatePickUpWaybillPreQueryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreatePickUpWaybillPreQueryResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePickUpWaybillPreQueryResponseBody) SetData(v *CreatePickUpWaybillPreQueryResponseBodyData) *CreatePickUpWaybillPreQueryResponseBody {
	s.Data = v
	return s
}

func (s *CreatePickUpWaybillPreQueryResponseBody) SetHttpStatusCode(v int32) *CreatePickUpWaybillPreQueryResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreatePickUpWaybillPreQueryResponseBody) SetMessage(v string) *CreatePickUpWaybillPreQueryResponseBody {
	s.Message = &v
	return s
}

func (s *CreatePickUpWaybillPreQueryResponseBody) SetRequestId(v string) *CreatePickUpWaybillPreQueryResponseBody {
	s.RequestId = &v
	return s
}

type CreatePickUpWaybillPreQueryResponseBodyData struct {
	// The response code.
	//
	// example:
	//
	// 0
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information about whether the courier company can accept the order.
	CpTimeSelectList []*CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectList `json:"CpTimeSelectList,omitempty" xml:"CpTimeSelectList,omitempty" type:"Repeated"`
	// The error code.
	//
	// example:
	//
	// Success
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// none
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// The response content.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreatePickUpWaybillPreQueryResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreatePickUpWaybillPreQueryResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreatePickUpWaybillPreQueryResponseBodyData) SetCode(v string) *CreatePickUpWaybillPreQueryResponseBodyData {
	s.Code = &v
	return s
}

func (s *CreatePickUpWaybillPreQueryResponseBodyData) SetCpTimeSelectList(v []*CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectList) *CreatePickUpWaybillPreQueryResponseBodyData {
	s.CpTimeSelectList = v
	return s
}

func (s *CreatePickUpWaybillPreQueryResponseBodyData) SetErrorCode(v string) *CreatePickUpWaybillPreQueryResponseBodyData {
	s.ErrorCode = &v
	return s
}

func (s *CreatePickUpWaybillPreQueryResponseBodyData) SetErrorMsg(v string) *CreatePickUpWaybillPreQueryResponseBodyData {
	s.ErrorMsg = &v
	return s
}

func (s *CreatePickUpWaybillPreQueryResponseBodyData) SetMessage(v string) *CreatePickUpWaybillPreQueryResponseBodyData {
	s.Message = &v
	return s
}

func (s *CreatePickUpWaybillPreQueryResponseBodyData) SetSuccess(v bool) *CreatePickUpWaybillPreQueryResponseBodyData {
	s.Success = &v
	return s
}

type CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectList struct {
	// The available time for the scheduled pickup. If the current courier company cannot accept the scheduled pickup, this field is left empty.
	AppointTimes []*CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectListAppointTimes `json:"AppointTimes,omitempty" xml:"AppointTimes,omitempty" type:"Repeated"`
	// The estimated price. Unit: CNY. The value is accurate to two decimal places. The value of this parameter is displayed if an estimated weight is specified.
	//
	// example:
	//
	// 12.50
	PrePrice *string `json:"PrePrice,omitempty" xml:"PrePrice,omitempty"`
	// The information about whether the real-time order can be selected.
	RealTime *CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectListRealTime `json:"RealTime,omitempty" xml:"RealTime,omitempty" type:"Struct"`
}

func (s CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectList) String() string {
	return tea.Prettify(s)
}

func (s CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectList) GoString() string {
	return s.String()
}

func (s *CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectList) SetAppointTimes(v []*CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectListAppointTimes) *CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectList {
	s.AppointTimes = v
	return s
}

func (s *CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectList) SetPrePrice(v string) *CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectList {
	s.PrePrice = &v
	return s
}

func (s *CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectList) SetRealTime(v *CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectListRealTime) *CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectList {
	s.RealTime = v
	return s
}

type CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectListAppointTimes struct {
	// The date in the YYYY-MM-DD format.
	//
	// example:
	//
	// 2022-04-28
	Date *string `json:"Date,omitempty" xml:"Date,omitempty"`
	// Indicates whether the date is selectable.
	//
	// example:
	//
	// true
	DateSelectable *bool `json:"DateSelectable,omitempty" xml:"DateSelectable,omitempty"`
	// The time range for the scheduled pickup for this date.
	TimeList []*CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectListAppointTimesTimeList `json:"TimeList,omitempty" xml:"TimeList,omitempty" type:"Repeated"`
}

func (s CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectListAppointTimes) String() string {
	return tea.Prettify(s)
}

func (s CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectListAppointTimes) GoString() string {
	return s.String()
}

func (s *CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectListAppointTimes) SetDate(v string) *CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectListAppointTimes {
	s.Date = &v
	return s
}

func (s *CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectListAppointTimes) SetDateSelectable(v bool) *CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectListAppointTimes {
	s.DateSelectable = &v
	return s
}

func (s *CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectListAppointTimes) SetTimeList(v []*CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectListAppointTimesTimeList) *CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectListAppointTimes {
	s.TimeList = v
	return s
}

type CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectListAppointTimesTimeList struct {
	// The end of the time range.
	//
	// example:
	//
	// 12:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The tip displayed when the scheduled pickup is not available.
	//
	// example:
	//
	// Appointment Full
	SelectDisableTip *string `json:"SelectDisableTip,omitempty" xml:"SelectDisableTip,omitempty"`
	// Indicates whether the time range can be selected for the scheduled pickup.
	//
	// example:
	//
	// true
	Selectable *bool `json:"Selectable,omitempty" xml:"Selectable,omitempty"`
	// The beginning of the time range.
	//
	// example:
	//
	// 10:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectListAppointTimesTimeList) String() string {
	return tea.Prettify(s)
}

func (s CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectListAppointTimesTimeList) GoString() string {
	return s.String()
}

func (s *CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectListAppointTimesTimeList) SetEndTime(v string) *CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectListAppointTimesTimeList {
	s.EndTime = &v
	return s
}

func (s *CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectListAppointTimesTimeList) SetSelectDisableTip(v string) *CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectListAppointTimesTimeList {
	s.SelectDisableTip = &v
	return s
}

func (s *CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectListAppointTimesTimeList) SetSelectable(v bool) *CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectListAppointTimesTimeList {
	s.Selectable = &v
	return s
}

func (s *CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectListAppointTimesTimeList) SetStartTime(v string) *CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectListAppointTimesTimeList {
	s.StartTime = &v
	return s
}

type CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectListRealTime struct {
	// The name of the real-time order type.
	//
	// example:
	//
	// Aliyun
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The tip displayed when the real-time order cannot be placed.
	//
	// example:
	//
	// Exceeding the real-time ordering time range.
	SelectDisableTip *string `json:"SelectDisableTip,omitempty" xml:"SelectDisableTip,omitempty"`
	// Indicates whether the real-time order can be placed after the deadline for placing a real-time order is reached.
	//
	// example:
	//
	// false
	Selectable *bool `json:"Selectable,omitempty" xml:"Selectable,omitempty"`
}

func (s CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectListRealTime) String() string {
	return tea.Prettify(s)
}

func (s CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectListRealTime) GoString() string {
	return s.String()
}

func (s *CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectListRealTime) SetName(v string) *CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectListRealTime {
	s.Name = &v
	return s
}

func (s *CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectListRealTime) SetSelectDisableTip(v string) *CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectListRealTime {
	s.SelectDisableTip = &v
	return s
}

func (s *CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectListRealTime) SetSelectable(v bool) *CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectListRealTime {
	s.Selectable = &v
	return s
}

type CreatePickUpWaybillPreQueryResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePickUpWaybillPreQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePickUpWaybillPreQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePickUpWaybillPreQueryResponse) GoString() string {
	return s.String()
}

func (s *CreatePickUpWaybillPreQueryResponse) SetHeaders(v map[string]*string) *CreatePickUpWaybillPreQueryResponse {
	s.Headers = v
	return s
}

func (s *CreatePickUpWaybillPreQueryResponse) SetStatusCode(v int32) *CreatePickUpWaybillPreQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePickUpWaybillPreQueryResponse) SetBody(v *CreatePickUpWaybillPreQueryResponseBody) *CreatePickUpWaybillPreQueryResponse {
	s.Body = v
	return s
}

type CreateSmsSignRequest struct {
	// 收信人号码
	//
	// This parameter is required.
	//
	// example:
	//
	// 178****3614
	CalledNo *string `json:"CalledNo,omitempty" xml:"CalledNo,omitempty"`
	// 客户uid
	//
	// example:
	//
	// -
	CallerParentId *int64 `json:"CallerParentId,omitempty" xml:"CallerParentId,omitempty"`
	// 发信人号码
	//
	// This parameter is required.
	//
	// example:
	//
	// 150****6539
	CallingNo *string `json:"CallingNo,omitempty" xml:"CallingNo,omitempty"`
	// 号码池key
	//
	// This parameter is required.
	//
	// example:
	//
	// FC5**********************a1a
	CustomerPoolKey *string `json:"CustomerPoolKey,omitempty" xml:"CustomerPoolKey,omitempty"`
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 请求去重ID, reqId最大长度为20位,接入方需要保持原子性
	//
	// This parameter is required.
	//
	// example:
	//
	// 564**********879
	ReqId                *string `json:"ReqId,omitempty" xml:"ReqId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateSmsSignRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSmsSignRequest) GoString() string {
	return s.String()
}

func (s *CreateSmsSignRequest) SetCalledNo(v string) *CreateSmsSignRequest {
	s.CalledNo = &v
	return s
}

func (s *CreateSmsSignRequest) SetCallerParentId(v int64) *CreateSmsSignRequest {
	s.CallerParentId = &v
	return s
}

func (s *CreateSmsSignRequest) SetCallingNo(v string) *CreateSmsSignRequest {
	s.CallingNo = &v
	return s
}

func (s *CreateSmsSignRequest) SetCustomerPoolKey(v string) *CreateSmsSignRequest {
	s.CustomerPoolKey = &v
	return s
}

func (s *CreateSmsSignRequest) SetOwnerId(v int64) *CreateSmsSignRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateSmsSignRequest) SetReqId(v string) *CreateSmsSignRequest {
	s.ReqId = &v
	return s
}

func (s *CreateSmsSignRequest) SetResourceOwnerAccount(v string) *CreateSmsSignRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateSmsSignRequest) SetResourceOwnerId(v int64) *CreateSmsSignRequest {
	s.ResourceOwnerId = &v
	return s
}

type CreateSmsSignResponseBody struct {
	// example:
	//
	// 0
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// 返回状态码 0000表示成功 其他表示失败
	//
	// example:
	//
	// 0000
	Code *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CreateSmsSignResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// 返回信息
	//
	// example:
	//
	// 成功
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 返回是否成功 true  表示成功 false表示失败
	//
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateSmsSignResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSmsSignResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSmsSignResponseBody) SetAccessDeniedDetail(v string) *CreateSmsSignResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CreateSmsSignResponseBody) SetCode(v string) *CreateSmsSignResponseBody {
	s.Code = &v
	return s
}

func (s *CreateSmsSignResponseBody) SetData(v *CreateSmsSignResponseBodyData) *CreateSmsSignResponseBody {
	s.Data = v
	return s
}

func (s *CreateSmsSignResponseBody) SetMessage(v string) *CreateSmsSignResponseBody {
	s.Message = &v
	return s
}

func (s *CreateSmsSignResponseBody) SetSuccess(v bool) *CreateSmsSignResponseBody {
	s.Success = &v
	return s
}

type CreateSmsSignResponseBodyData struct {
	// 短信接收者号码签名串(加到短信内容中供解析真实被叫号码)
	//
	// example:
	//
	// sign23343466
	CalledNoSign *string `json:"CalledNoSign,omitempty" xml:"CalledNoSign,omitempty"`
}

func (s CreateSmsSignResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateSmsSignResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateSmsSignResponseBodyData) SetCalledNoSign(v string) *CreateSmsSignResponseBodyData {
	s.CalledNoSign = &v
	return s
}

type CreateSmsSignResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSmsSignResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSmsSignResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSmsSignResponse) GoString() string {
	return s.String()
}

func (s *CreateSmsSignResponse) SetHeaders(v map[string]*string) *CreateSmsSignResponse {
	s.Headers = v
	return s
}

func (s *CreateSmsSignResponse) SetStatusCode(v int32) *CreateSmsSignResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSmsSignResponse) SetBody(v *CreateSmsSignResponseBody) *CreateSmsSignResponse {
	s.Body = v
	return s
}

type DeleteAxbBindFixedLineRequest struct {
	// 应用id
	//
	// This parameter is required.
	//
	// example:
	//
	// ALPT_1234
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 消息请求标识
	//
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值示例值
	OrderId              *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 绑定id
	//
	// This parameter is required.
	//
	// example:
	//
	// A93IOELD93
	SubId *string `json:"SubId,omitempty" xml:"SubId,omitempty"`
	// 业务时间戳，格式yyyyMMddHHmmssSSS，时间采用北京时间，24小时制，精确至毫秒
	//
	// This parameter is required.
	//
	// example:
	//
	// 20161114143116001
	Ts *string `json:"Ts,omitempty" xml:"Ts,omitempty"`
}

func (s DeleteAxbBindFixedLineRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAxbBindFixedLineRequest) GoString() string {
	return s.String()
}

func (s *DeleteAxbBindFixedLineRequest) SetAppId(v string) *DeleteAxbBindFixedLineRequest {
	s.AppId = &v
	return s
}

func (s *DeleteAxbBindFixedLineRequest) SetOrderId(v string) *DeleteAxbBindFixedLineRequest {
	s.OrderId = &v
	return s
}

func (s *DeleteAxbBindFixedLineRequest) SetOwnerId(v int64) *DeleteAxbBindFixedLineRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteAxbBindFixedLineRequest) SetResourceOwnerAccount(v string) *DeleteAxbBindFixedLineRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteAxbBindFixedLineRequest) SetResourceOwnerId(v int64) *DeleteAxbBindFixedLineRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteAxbBindFixedLineRequest) SetSubId(v string) *DeleteAxbBindFixedLineRequest {
	s.SubId = &v
	return s
}

func (s *DeleteAxbBindFixedLineRequest) SetTs(v string) *DeleteAxbBindFixedLineRequest {
	s.Ts = &v
	return s
}

type DeleteAxbBindFixedLineResponseBody struct {
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// isp.SYSTEM_ERROR
	Code *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *DeleteAxbBindFixedLineResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// F036366A-0182-5066-A686-19C4C82F2D51
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteAxbBindFixedLineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAxbBindFixedLineResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAxbBindFixedLineResponseBody) SetAccessDeniedDetail(v string) *DeleteAxbBindFixedLineResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DeleteAxbBindFixedLineResponseBody) SetCode(v string) *DeleteAxbBindFixedLineResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteAxbBindFixedLineResponseBody) SetData(v *DeleteAxbBindFixedLineResponseBodyData) *DeleteAxbBindFixedLineResponseBody {
	s.Data = v
	return s
}

func (s *DeleteAxbBindFixedLineResponseBody) SetMessage(v string) *DeleteAxbBindFixedLineResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteAxbBindFixedLineResponseBody) SetRequestId(v string) *DeleteAxbBindFixedLineResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAxbBindFixedLineResponseBody) SetSuccess(v bool) *DeleteAxbBindFixedLineResponseBody {
	s.Success = &v
	return s
}

type DeleteAxbBindFixedLineResponseBodyData struct {
	// 响应码  0-成功
	//
	// example:
	//
	// 0
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 响应消息
	//
	// example:
	//
	// 示例值示例值示例值
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 是否处理成功  true-成功  false-失败
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteAxbBindFixedLineResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DeleteAxbBindFixedLineResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteAxbBindFixedLineResponseBodyData) SetCode(v string) *DeleteAxbBindFixedLineResponseBodyData {
	s.Code = &v
	return s
}

func (s *DeleteAxbBindFixedLineResponseBodyData) SetMessage(v string) *DeleteAxbBindFixedLineResponseBodyData {
	s.Message = &v
	return s
}

func (s *DeleteAxbBindFixedLineResponseBodyData) SetSuccess(v bool) *DeleteAxbBindFixedLineResponseBodyData {
	s.Success = &v
	return s
}

type DeleteAxbBindFixedLineResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAxbBindFixedLineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAxbBindFixedLineResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAxbBindFixedLineResponse) GoString() string {
	return s.String()
}

func (s *DeleteAxbBindFixedLineResponse) SetHeaders(v map[string]*string) *DeleteAxbBindFixedLineResponse {
	s.Headers = v
	return s
}

func (s *DeleteAxbBindFixedLineResponse) SetStatusCode(v int32) *DeleteAxbBindFixedLineResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAxbBindFixedLineResponse) SetBody(v *DeleteAxbBindFixedLineResponseBody) *DeleteAxbBindFixedLineResponse {
	s.Body = v
	return s
}

type DeleteAxgGroupRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1234
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// FC2235****
	PoolKey              *string `json:"PoolKey,omitempty" xml:"PoolKey,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteAxgGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAxgGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteAxgGroupRequest) SetGroupId(v int64) *DeleteAxgGroupRequest {
	s.GroupId = &v
	return s
}

func (s *DeleteAxgGroupRequest) SetOwnerId(v int64) *DeleteAxgGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteAxgGroupRequest) SetPoolKey(v string) *DeleteAxgGroupRequest {
	s.PoolKey = &v
	return s
}

func (s *DeleteAxgGroupRequest) SetResourceOwnerAccount(v string) *DeleteAxgGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteAxgGroupRequest) SetResourceOwnerId(v int64) *DeleteAxgGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

type DeleteAxgGroupResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// SUCCESS
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 9297B722-A016-43FB-B51A-E54050D9369D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAxgGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAxgGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAxgGroupResponseBody) SetCode(v string) *DeleteAxgGroupResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteAxgGroupResponseBody) SetMessage(v string) *DeleteAxgGroupResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteAxgGroupResponseBody) SetRequestId(v string) *DeleteAxgGroupResponseBody {
	s.RequestId = &v
	return s
}

type DeleteAxgGroupResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAxgGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAxgGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAxgGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteAxgGroupResponse) SetHeaders(v map[string]*string) *DeleteAxgGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteAxgGroupResponse) SetStatusCode(v int32) *DeleteAxgGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAxgGroupResponse) SetBody(v *DeleteAxgGroupResponseBody) *DeleteAxgGroupResponse {
	s.Body = v
	return s
}

type DeleteAxnBindFixedLineRequest struct {
	// 业务id标识，由阿里云分配给客户侧
	//
	// This parameter is required.
	//
	// example:
	//
	// alitest
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 消息请求唯一标识
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345dkwkd99d
	OrderId              *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 绑定id
	//
	// This parameter is required.
	//
	// example:
	//
	// GHX0534X202504221531579290029-2-1-aliaxn
	SubId *string `json:"SubId,omitempty" xml:"SubId,omitempty"`
	// 格式yyyyMMddHHmmssSSS，时间采用北京时间，24小时制，精确至毫秒
	//
	// This parameter is required.
	//
	// example:
	//
	// 20250421141723098
	Ts *string `json:"Ts,omitempty" xml:"Ts,omitempty"`
}

func (s DeleteAxnBindFixedLineRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAxnBindFixedLineRequest) GoString() string {
	return s.String()
}

func (s *DeleteAxnBindFixedLineRequest) SetAppId(v string) *DeleteAxnBindFixedLineRequest {
	s.AppId = &v
	return s
}

func (s *DeleteAxnBindFixedLineRequest) SetOrderId(v string) *DeleteAxnBindFixedLineRequest {
	s.OrderId = &v
	return s
}

func (s *DeleteAxnBindFixedLineRequest) SetOwnerId(v int64) *DeleteAxnBindFixedLineRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteAxnBindFixedLineRequest) SetResourceOwnerAccount(v string) *DeleteAxnBindFixedLineRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteAxnBindFixedLineRequest) SetResourceOwnerId(v int64) *DeleteAxnBindFixedLineRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteAxnBindFixedLineRequest) SetSubId(v string) *DeleteAxnBindFixedLineRequest {
	s.SubId = &v
	return s
}

func (s *DeleteAxnBindFixedLineRequest) SetTs(v string) *DeleteAxnBindFixedLineRequest {
	s.Ts = &v
	return s
}

type DeleteAxnBindFixedLineResponseBody struct {
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *DeleteAxnBindFixedLineResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// AE2D6997-643A-59CB-9B3C-918572E5CEAA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteAxnBindFixedLineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAxnBindFixedLineResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAxnBindFixedLineResponseBody) SetAccessDeniedDetail(v string) *DeleteAxnBindFixedLineResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DeleteAxnBindFixedLineResponseBody) SetCode(v string) *DeleteAxnBindFixedLineResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteAxnBindFixedLineResponseBody) SetData(v *DeleteAxnBindFixedLineResponseBodyData) *DeleteAxnBindFixedLineResponseBody {
	s.Data = v
	return s
}

func (s *DeleteAxnBindFixedLineResponseBody) SetMessage(v string) *DeleteAxnBindFixedLineResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteAxnBindFixedLineResponseBody) SetRequestId(v string) *DeleteAxnBindFixedLineResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAxnBindFixedLineResponseBody) SetSuccess(v bool) *DeleteAxnBindFixedLineResponseBody {
	s.Success = &v
	return s
}

type DeleteAxnBindFixedLineResponseBodyData struct {
	// 响应码 0：成功，其它失败，具体见文档
	//
	// example:
	//
	// 0
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 描述信息
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteAxnBindFixedLineResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DeleteAxnBindFixedLineResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteAxnBindFixedLineResponseBodyData) SetCode(v string) *DeleteAxnBindFixedLineResponseBodyData {
	s.Code = &v
	return s
}

func (s *DeleteAxnBindFixedLineResponseBodyData) SetMessage(v string) *DeleteAxnBindFixedLineResponseBodyData {
	s.Message = &v
	return s
}

func (s *DeleteAxnBindFixedLineResponseBodyData) SetSuccess(v bool) *DeleteAxnBindFixedLineResponseBodyData {
	s.Success = &v
	return s
}

type DeleteAxnBindFixedLineResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAxnBindFixedLineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAxnBindFixedLineResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAxnBindFixedLineResponse) GoString() string {
	return s.String()
}

func (s *DeleteAxnBindFixedLineResponse) SetHeaders(v map[string]*string) *DeleteAxnBindFixedLineResponse {
	s.Headers = v
	return s
}

func (s *DeleteAxnBindFixedLineResponse) SetStatusCode(v int32) *DeleteAxnBindFixedLineResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAxnBindFixedLineResponse) SetBody(v *DeleteAxnBindFixedLineResponseBody) *DeleteAxnBindFixedLineResponse {
	s.Body = v
	return s
}

type DeleteAxnExtensionBindFixedLineRequest struct {
	// 业务id标识，由阿里云分配给客户侧
	//
	// This parameter is required.
	//
	// example:
	//
	// alitest
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 消息请求唯一标识
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345dkwkd99d
	OrderId              *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 绑定id
	//
	// This parameter is required.
	//
	// example:
	//
	// GHX0534X202504221531579290029
	SubId *string `json:"SubId,omitempty" xml:"SubId,omitempty"`
	// 格式yyyyMMddHHmmssSSS，时间采用北京时间，24小时制，精确至毫秒
	//
	// This parameter is required.
	//
	// example:
	//
	// 20250421141723098
	Ts *string `json:"Ts,omitempty" xml:"Ts,omitempty"`
}

func (s DeleteAxnExtensionBindFixedLineRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAxnExtensionBindFixedLineRequest) GoString() string {
	return s.String()
}

func (s *DeleteAxnExtensionBindFixedLineRequest) SetAppId(v string) *DeleteAxnExtensionBindFixedLineRequest {
	s.AppId = &v
	return s
}

func (s *DeleteAxnExtensionBindFixedLineRequest) SetOrderId(v string) *DeleteAxnExtensionBindFixedLineRequest {
	s.OrderId = &v
	return s
}

func (s *DeleteAxnExtensionBindFixedLineRequest) SetOwnerId(v int64) *DeleteAxnExtensionBindFixedLineRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteAxnExtensionBindFixedLineRequest) SetResourceOwnerAccount(v string) *DeleteAxnExtensionBindFixedLineRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteAxnExtensionBindFixedLineRequest) SetResourceOwnerId(v int64) *DeleteAxnExtensionBindFixedLineRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteAxnExtensionBindFixedLineRequest) SetSubId(v string) *DeleteAxnExtensionBindFixedLineRequest {
	s.SubId = &v
	return s
}

func (s *DeleteAxnExtensionBindFixedLineRequest) SetTs(v string) *DeleteAxnExtensionBindFixedLineRequest {
	s.Ts = &v
	return s
}

type DeleteAxnExtensionBindFixedLineResponseBody struct {
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *DeleteAxnExtensionBindFixedLineResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// conflict with subs id=1000203635098305, phoneA conflict
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 3DA9D6DF-C5FA-5A0D-B6C2-547B1FD1F9B4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteAxnExtensionBindFixedLineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAxnExtensionBindFixedLineResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAxnExtensionBindFixedLineResponseBody) SetAccessDeniedDetail(v string) *DeleteAxnExtensionBindFixedLineResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DeleteAxnExtensionBindFixedLineResponseBody) SetCode(v string) *DeleteAxnExtensionBindFixedLineResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteAxnExtensionBindFixedLineResponseBody) SetData(v *DeleteAxnExtensionBindFixedLineResponseBodyData) *DeleteAxnExtensionBindFixedLineResponseBody {
	s.Data = v
	return s
}

func (s *DeleteAxnExtensionBindFixedLineResponseBody) SetMessage(v string) *DeleteAxnExtensionBindFixedLineResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteAxnExtensionBindFixedLineResponseBody) SetRequestId(v string) *DeleteAxnExtensionBindFixedLineResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAxnExtensionBindFixedLineResponseBody) SetSuccess(v bool) *DeleteAxnExtensionBindFixedLineResponseBody {
	s.Success = &v
	return s
}

type DeleteAxnExtensionBindFixedLineResponseBodyData struct {
	// 响应码 0：成功，其它失败，具体见文档
	//
	// example:
	//
	// 0
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 描述信息
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteAxnExtensionBindFixedLineResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DeleteAxnExtensionBindFixedLineResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteAxnExtensionBindFixedLineResponseBodyData) SetCode(v string) *DeleteAxnExtensionBindFixedLineResponseBodyData {
	s.Code = &v
	return s
}

func (s *DeleteAxnExtensionBindFixedLineResponseBodyData) SetMessage(v string) *DeleteAxnExtensionBindFixedLineResponseBodyData {
	s.Message = &v
	return s
}

func (s *DeleteAxnExtensionBindFixedLineResponseBodyData) SetSuccess(v bool) *DeleteAxnExtensionBindFixedLineResponseBodyData {
	s.Success = &v
	return s
}

type DeleteAxnExtensionBindFixedLineResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAxnExtensionBindFixedLineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAxnExtensionBindFixedLineResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAxnExtensionBindFixedLineResponse) GoString() string {
	return s.String()
}

func (s *DeleteAxnExtensionBindFixedLineResponse) SetHeaders(v map[string]*string) *DeleteAxnExtensionBindFixedLineResponse {
	s.Headers = v
	return s
}

func (s *DeleteAxnExtensionBindFixedLineResponse) SetStatusCode(v int32) *DeleteAxnExtensionBindFixedLineResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAxnExtensionBindFixedLineResponse) SetBody(v *DeleteAxnExtensionBindFixedLineResponseBody) *DeleteAxnExtensionBindFixedLineResponse {
	s.Body = v
	return s
}

type DeleteSecretAPhoneNoToCustRequest struct {
	// A号码组ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 51
	ANoWhiteGroupId *int64 `json:"ANoWhiteGroupId,omitempty" xml:"ANoWhiteGroupId,omitempty"`
	OwnerId         *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// A号码
	//
	// This parameter is required.
	//
	// example:
	//
	// 130*****1234
	PhoneNoA             *string `json:"PhoneNoA,omitempty" xml:"PhoneNoA,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteSecretAPhoneNoToCustRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSecretAPhoneNoToCustRequest) GoString() string {
	return s.String()
}

func (s *DeleteSecretAPhoneNoToCustRequest) SetANoWhiteGroupId(v int64) *DeleteSecretAPhoneNoToCustRequest {
	s.ANoWhiteGroupId = &v
	return s
}

func (s *DeleteSecretAPhoneNoToCustRequest) SetOwnerId(v int64) *DeleteSecretAPhoneNoToCustRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteSecretAPhoneNoToCustRequest) SetPhoneNoA(v string) *DeleteSecretAPhoneNoToCustRequest {
	s.PhoneNoA = &v
	return s
}

func (s *DeleteSecretAPhoneNoToCustRequest) SetResourceOwnerAccount(v string) *DeleteSecretAPhoneNoToCustRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteSecretAPhoneNoToCustRequest) SetResourceOwnerId(v int64) *DeleteSecretAPhoneNoToCustRequest {
	s.ResourceOwnerId = &v
	return s
}

type DeleteSecretAPhoneNoToCustResponseBody struct {
	// example:
	//
	// 0
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// 请求状态码
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 删除是否成功
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// 失败错误提示
	//
	// example:
	//
	// 号码组不存在
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 返回id
	//
	// example:
	//
	// 1D73E648-0978-18A5-B089-3BB2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 请求是否成功
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteSecretAPhoneNoToCustResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSecretAPhoneNoToCustResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSecretAPhoneNoToCustResponseBody) SetAccessDeniedDetail(v string) *DeleteSecretAPhoneNoToCustResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DeleteSecretAPhoneNoToCustResponseBody) SetCode(v string) *DeleteSecretAPhoneNoToCustResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteSecretAPhoneNoToCustResponseBody) SetData(v bool) *DeleteSecretAPhoneNoToCustResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteSecretAPhoneNoToCustResponseBody) SetMessage(v string) *DeleteSecretAPhoneNoToCustResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteSecretAPhoneNoToCustResponseBody) SetRequestId(v string) *DeleteSecretAPhoneNoToCustResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSecretAPhoneNoToCustResponseBody) SetSuccess(v bool) *DeleteSecretAPhoneNoToCustResponseBody {
	s.Success = &v
	return s
}

type DeleteSecretAPhoneNoToCustResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSecretAPhoneNoToCustResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSecretAPhoneNoToCustResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSecretAPhoneNoToCustResponse) GoString() string {
	return s.String()
}

func (s *DeleteSecretAPhoneNoToCustResponse) SetHeaders(v map[string]*string) *DeleteSecretAPhoneNoToCustResponse {
	s.Headers = v
	return s
}

func (s *DeleteSecretAPhoneNoToCustResponse) SetStatusCode(v int32) *DeleteSecretAPhoneNoToCustResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSecretAPhoneNoToCustResponse) SetBody(v *DeleteSecretAPhoneNoToCustResponseBody) *DeleteSecretAPhoneNoToCustResponse {
	s.Body = v
	return s
}

type DeleteSecretBlacklistRequest struct {
	// The phone numbers in the blacklist. A point-to-point blacklist has a pair of numbers separated by a colon (":"). A number pool blacklist or a platform blacklist has only one single number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 18252***383:18252***483
	BlackNo *string `json:"BlackNo,omitempty" xml:"BlackNo,omitempty"`
	// The blacklist type. Valid values:
	//
	// 	- **POINT_TO_POINT_BLACK**: point-to-point blacklist
	//
	// 	- **PARTNER_GLOBAL_NUMBER_BLACK**: number pool blacklist
	//
	// This parameter is required.
	//
	// example:
	//
	// POINT_TO_POINT_BLACK
	BlackType *string `json:"BlackType,omitempty" xml:"BlackType,omitempty"`
	// The key of the phone number pool. Log on to the [Phone Number Protection console](https://dypls.console.aliyun.com/dypls.htm#/account) and view the key of the phone number pool on the **Number Pool Management*	- page.
	//
	// This parameter is required.
	//
	// example:
	//
	// FC1232****
	PoolKey *string `json:"PoolKey,omitempty" xml:"PoolKey,omitempty"`
	// The remarks for the blacklist.
	//
	// example:
	//
	// fragile
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The control on the call direction.
	//
	// 	- **PHONEA_REJECT**: The first phone number in the blacklist can be used to call the second phone number, but the second phone number in the blacklist cannot be used to call the first phone number.
	//
	// 	- **PHONEB_REJECT**: The first phone number in the blacklist cannot be used to call the second phone number, but the second phone number in the blacklist can be used to call the first phone number.
	//
	// 	- If this parameter is not specified, the two phone numbers in the blacklist cannot be used to call each other.
	//
	// Valid values:
	//
	// 	- DUPLEX_REJECT
	//
	// 	- PHONEA_REJECT
	//
	// 	- PHONEB_REJECT
	//
	// example:
	//
	// PHONEA_REJECT
	WayControl *string `json:"WayControl,omitempty" xml:"WayControl,omitempty"`
}

func (s DeleteSecretBlacklistRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSecretBlacklistRequest) GoString() string {
	return s.String()
}

func (s *DeleteSecretBlacklistRequest) SetBlackNo(v string) *DeleteSecretBlacklistRequest {
	s.BlackNo = &v
	return s
}

func (s *DeleteSecretBlacklistRequest) SetBlackType(v string) *DeleteSecretBlacklistRequest {
	s.BlackType = &v
	return s
}

func (s *DeleteSecretBlacklistRequest) SetPoolKey(v string) *DeleteSecretBlacklistRequest {
	s.PoolKey = &v
	return s
}

func (s *DeleteSecretBlacklistRequest) SetRemark(v string) *DeleteSecretBlacklistRequest {
	s.Remark = &v
	return s
}

func (s *DeleteSecretBlacklistRequest) SetWayControl(v string) *DeleteSecretBlacklistRequest {
	s.WayControl = &v
	return s
}

type DeleteSecretBlacklistResponseBody struct {
	// The response code.
	//
	// 	- The value OK indicates that the request was successful.
	//
	// 	- Other status codes indicate that the request failed. For more information, see [Error codes](https://help.aliyun.com/document_detail/109196.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F655A8D5-B967-440B-8683-DAD6FF8DE990
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSecretBlacklistResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSecretBlacklistResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSecretBlacklistResponseBody) SetCode(v string) *DeleteSecretBlacklistResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteSecretBlacklistResponseBody) SetMessage(v string) *DeleteSecretBlacklistResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteSecretBlacklistResponseBody) SetRequestId(v string) *DeleteSecretBlacklistResponseBody {
	s.RequestId = &v
	return s
}

type DeleteSecretBlacklistResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSecretBlacklistResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSecretBlacklistResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSecretBlacklistResponse) GoString() string {
	return s.String()
}

func (s *DeleteSecretBlacklistResponse) SetHeaders(v map[string]*string) *DeleteSecretBlacklistResponse {
	s.Headers = v
	return s
}

func (s *DeleteSecretBlacklistResponse) SetStatusCode(v int32) *DeleteSecretBlacklistResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSecretBlacklistResponse) SetBody(v *DeleteSecretBlacklistResponseBody) *DeleteSecretBlacklistResponse {
	s.Body = v
	return s
}

type GetDyplsOSSInfoForUploadFileRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// phone_card
	BizType              *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetDyplsOSSInfoForUploadFileRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDyplsOSSInfoForUploadFileRequest) GoString() string {
	return s.String()
}

func (s *GetDyplsOSSInfoForUploadFileRequest) SetBizType(v string) *GetDyplsOSSInfoForUploadFileRequest {
	s.BizType = &v
	return s
}

func (s *GetDyplsOSSInfoForUploadFileRequest) SetOwnerId(v int64) *GetDyplsOSSInfoForUploadFileRequest {
	s.OwnerId = &v
	return s
}

func (s *GetDyplsOSSInfoForUploadFileRequest) SetResourceOwnerAccount(v string) *GetDyplsOSSInfoForUploadFileRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetDyplsOSSInfoForUploadFileRequest) SetResourceOwnerId(v int64) *GetDyplsOSSInfoForUploadFileRequest {
	s.ResourceOwnerId = &v
	return s
}

type GetDyplsOSSInfoForUploadFileResponseBody struct {
	// example:
	//
	// 0
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetDyplsOSSInfoForUploadFileResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// E2FD3B2F-5028-16E3-9F83-2F76F99B3873
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDyplsOSSInfoForUploadFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDyplsOSSInfoForUploadFileResponseBody) GoString() string {
	return s.String()
}

func (s *GetDyplsOSSInfoForUploadFileResponseBody) SetAccessDeniedDetail(v string) *GetDyplsOSSInfoForUploadFileResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetDyplsOSSInfoForUploadFileResponseBody) SetCode(v string) *GetDyplsOSSInfoForUploadFileResponseBody {
	s.Code = &v
	return s
}

func (s *GetDyplsOSSInfoForUploadFileResponseBody) SetData(v *GetDyplsOSSInfoForUploadFileResponseBodyData) *GetDyplsOSSInfoForUploadFileResponseBody {
	s.Data = v
	return s
}

func (s *GetDyplsOSSInfoForUploadFileResponseBody) SetMessage(v string) *GetDyplsOSSInfoForUploadFileResponseBody {
	s.Message = &v
	return s
}

func (s *GetDyplsOSSInfoForUploadFileResponseBody) SetRequestId(v string) *GetDyplsOSSInfoForUploadFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDyplsOSSInfoForUploadFileResponseBody) SetSuccess(v bool) *GetDyplsOSSInfoForUploadFileResponseBody {
	s.Success = &v
	return s
}

type GetDyplsOSSInfoForUploadFileResponseBodyData struct {
	// example:
	//
	// LTAI***pSvPz
	AccessKeyId *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	// example:
	//
	// 1744613007
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// example:
	//
	// https://alicom-**********-cn-zhangjiakou.aliyuncs.com
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// example:
	//
	// IjoiMjAyN*****9udGV
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// example:
	//
	// BXwW**********aoZH
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
	// example:
	//
	// 123456
	StartPath *string `json:"StartPath,omitempty" xml:"StartPath,omitempty"`
}

func (s GetDyplsOSSInfoForUploadFileResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetDyplsOSSInfoForUploadFileResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDyplsOSSInfoForUploadFileResponseBodyData) SetAccessKeyId(v string) *GetDyplsOSSInfoForUploadFileResponseBodyData {
	s.AccessKeyId = &v
	return s
}

func (s *GetDyplsOSSInfoForUploadFileResponseBodyData) SetExpireTime(v string) *GetDyplsOSSInfoForUploadFileResponseBodyData {
	s.ExpireTime = &v
	return s
}

func (s *GetDyplsOSSInfoForUploadFileResponseBodyData) SetHost(v string) *GetDyplsOSSInfoForUploadFileResponseBodyData {
	s.Host = &v
	return s
}

func (s *GetDyplsOSSInfoForUploadFileResponseBodyData) SetPolicy(v string) *GetDyplsOSSInfoForUploadFileResponseBodyData {
	s.Policy = &v
	return s
}

func (s *GetDyplsOSSInfoForUploadFileResponseBodyData) SetSignature(v string) *GetDyplsOSSInfoForUploadFileResponseBodyData {
	s.Signature = &v
	return s
}

func (s *GetDyplsOSSInfoForUploadFileResponseBodyData) SetStartPath(v string) *GetDyplsOSSInfoForUploadFileResponseBodyData {
	s.StartPath = &v
	return s
}

type GetDyplsOSSInfoForUploadFileResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDyplsOSSInfoForUploadFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDyplsOSSInfoForUploadFileResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDyplsOSSInfoForUploadFileResponse) GoString() string {
	return s.String()
}

func (s *GetDyplsOSSInfoForUploadFileResponse) SetHeaders(v map[string]*string) *GetDyplsOSSInfoForUploadFileResponse {
	s.Headers = v
	return s
}

func (s *GetDyplsOSSInfoForUploadFileResponse) SetStatusCode(v int32) *GetDyplsOSSInfoForUploadFileResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDyplsOSSInfoForUploadFileResponse) SetBody(v *GetDyplsOSSInfoForUploadFileResponseBody) *GetDyplsOSSInfoForUploadFileResponse {
	s.Body = v
	return s
}

type GetSecretAsrDetailRequest struct {
	// The ID of the call record.
	//
	// You can log on to the [Phone Number Protection console](https://dypls.console.aliyun.com/dypls.htm#/account) and view **Call Record ID*	- on the **Call Record Query*	- page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 225625****
	CallId *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	// The call initiation time in the call record.
	//
	// You can log on to the [Phone Number Protection console](https://dypls.console.aliyun.com/dypls.htm#/account). View **Call Initiated At*	- on the **Call Record Query*	- page, or view the call_time field in the Call Detail Record (CDR) receipt.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2019-03-05 12:00:00
	CallTime *string `json:"CallTime,omitempty" xml:"CallTime,omitempty"`
	// The key of the phone number pool.
	//
	// You can log on to the [Phone Number Protection console](https://dypls.console.aliyun.com/dypls.htm#/account) and view the key of the phone number pool on the **Number Pool Management*	- page.
	//
	// This parameter is required.
	//
	// example:
	//
	// FC2267****
	PoolKey *string `json:"PoolKey,omitempty" xml:"PoolKey,omitempty"`
}

func (s GetSecretAsrDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSecretAsrDetailRequest) GoString() string {
	return s.String()
}

func (s *GetSecretAsrDetailRequest) SetCallId(v string) *GetSecretAsrDetailRequest {
	s.CallId = &v
	return s
}

func (s *GetSecretAsrDetailRequest) SetCallTime(v string) *GetSecretAsrDetailRequest {
	s.CallTime = &v
	return s
}

func (s *GetSecretAsrDetailRequest) SetPoolKey(v string) *GetSecretAsrDetailRequest {
	s.PoolKey = &v
	return s
}

type GetSecretAsrDetailResponseBody struct {
	// The response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ASR details.
	Data *GetSecretAsrDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F655A8D5-B967-440B-8683-DAD6FF8DE990
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetSecretAsrDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSecretAsrDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetSecretAsrDetailResponseBody) SetCode(v string) *GetSecretAsrDetailResponseBody {
	s.Code = &v
	return s
}

func (s *GetSecretAsrDetailResponseBody) SetData(v *GetSecretAsrDetailResponseBodyData) *GetSecretAsrDetailResponseBody {
	s.Data = v
	return s
}

func (s *GetSecretAsrDetailResponseBody) SetMessage(v string) *GetSecretAsrDetailResponseBody {
	s.Message = &v
	return s
}

func (s *GetSecretAsrDetailResponseBody) SetRequestId(v string) *GetSecretAsrDetailResponseBody {
	s.RequestId = &v
	return s
}

type GetSecretAsrDetailResponseBodyData struct {
	// The total duration of the audio file that was recognized. Unit: milliseconds.
	//
	// example:
	//
	// 10944
	BizDuration *int64 `json:"BizDuration,omitempty" xml:"BizDuration,omitempty"`
	// The ID of the business process.
	//
	// example:
	//
	// 435ee78c7a019650@!FC100000074672458@!2020061522****
	BusinessId *string `json:"BusinessId,omitempty" xml:"BusinessId,omitempty"`
	// The business keyword.
	//
	// example:
	//
	// JCGTncltuNao****
	BusinessKey *string `json:"BusinessKey,omitempty" xml:"BusinessKey,omitempty"`
	// The status code. The status code 21050000 indicates that the request was successful.
	//
	// example:
	//
	// 21050000
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The description.
	//
	// example:
	//
	// SUCCESS
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 8d2329d407a83447a83be441681f4872ac74nE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ASR result.
	Sentences []*GetSecretAsrDetailResponseBodyDataSentences `json:"Sentences,omitempty" xml:"Sentences,omitempty" type:"Repeated"`
	// The type.
	//
	// example:
	//
	// asrResult
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetSecretAsrDetailResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetSecretAsrDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSecretAsrDetailResponseBodyData) SetBizDuration(v int64) *GetSecretAsrDetailResponseBodyData {
	s.BizDuration = &v
	return s
}

func (s *GetSecretAsrDetailResponseBodyData) SetBusinessId(v string) *GetSecretAsrDetailResponseBodyData {
	s.BusinessId = &v
	return s
}

func (s *GetSecretAsrDetailResponseBodyData) SetBusinessKey(v string) *GetSecretAsrDetailResponseBodyData {
	s.BusinessKey = &v
	return s
}

func (s *GetSecretAsrDetailResponseBodyData) SetCode(v string) *GetSecretAsrDetailResponseBodyData {
	s.Code = &v
	return s
}

func (s *GetSecretAsrDetailResponseBodyData) SetMsg(v string) *GetSecretAsrDetailResponseBodyData {
	s.Msg = &v
	return s
}

func (s *GetSecretAsrDetailResponseBodyData) SetRequestId(v string) *GetSecretAsrDetailResponseBodyData {
	s.RequestId = &v
	return s
}

func (s *GetSecretAsrDetailResponseBodyData) SetSentences(v []*GetSecretAsrDetailResponseBodyDataSentences) *GetSecretAsrDetailResponseBodyData {
	s.Sentences = v
	return s
}

func (s *GetSecretAsrDetailResponseBodyData) SetType(v string) *GetSecretAsrDetailResponseBodyData {
	s.Type = &v
	return s
}

type GetSecretAsrDetailResponseBodyDataSentences struct {
	// The start time offset of the sentence. Unit: milliseconds.
	//
	// example:
	//
	// 1020
	BeginTime *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// The ID of the audio track to which the sentence belongs.
	//
	// example:
	//
	// 0
	ChannelId *int32 `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// The emotion value. Value range: 1 to 10. The higher the value, the stronger the emotion.
	//
	// example:
	//
	// 5.7
	EmotionValue *string `json:"EmotionValue,omitempty" xml:"EmotionValue,omitempty"`
	// The end time offset of the sentence. Unit: milliseconds.
	//
	// example:
	//
	// 1770
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The silence duration between the current sentence and the previous sentence. Unit: seconds.
	//
	// example:
	//
	// 0
	SilenceDuration *int64 `json:"SilenceDuration,omitempty" xml:"SilenceDuration,omitempty"`
	// The average speech rate of the sentence. Unit: number of words per minute.
	//
	// example:
	//
	// 80
	SpeechRate *int32 `json:"SpeechRate,omitempty" xml:"SpeechRate,omitempty"`
	// The recognition result of the sentence.
	//
	// example:
	//
	// Hello
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s GetSecretAsrDetailResponseBodyDataSentences) String() string {
	return tea.Prettify(s)
}

func (s GetSecretAsrDetailResponseBodyDataSentences) GoString() string {
	return s.String()
}

func (s *GetSecretAsrDetailResponseBodyDataSentences) SetBeginTime(v int64) *GetSecretAsrDetailResponseBodyDataSentences {
	s.BeginTime = &v
	return s
}

func (s *GetSecretAsrDetailResponseBodyDataSentences) SetChannelId(v int32) *GetSecretAsrDetailResponseBodyDataSentences {
	s.ChannelId = &v
	return s
}

func (s *GetSecretAsrDetailResponseBodyDataSentences) SetEmotionValue(v string) *GetSecretAsrDetailResponseBodyDataSentences {
	s.EmotionValue = &v
	return s
}

func (s *GetSecretAsrDetailResponseBodyDataSentences) SetEndTime(v int64) *GetSecretAsrDetailResponseBodyDataSentences {
	s.EndTime = &v
	return s
}

func (s *GetSecretAsrDetailResponseBodyDataSentences) SetSilenceDuration(v int64) *GetSecretAsrDetailResponseBodyDataSentences {
	s.SilenceDuration = &v
	return s
}

func (s *GetSecretAsrDetailResponseBodyDataSentences) SetSpeechRate(v int32) *GetSecretAsrDetailResponseBodyDataSentences {
	s.SpeechRate = &v
	return s
}

func (s *GetSecretAsrDetailResponseBodyDataSentences) SetText(v string) *GetSecretAsrDetailResponseBodyDataSentences {
	s.Text = &v
	return s
}

type GetSecretAsrDetailResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSecretAsrDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSecretAsrDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSecretAsrDetailResponse) GoString() string {
	return s.String()
}

func (s *GetSecretAsrDetailResponse) SetHeaders(v map[string]*string) *GetSecretAsrDetailResponse {
	s.Headers = v
	return s
}

func (s *GetSecretAsrDetailResponse) SetStatusCode(v int32) *GetSecretAsrDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSecretAsrDetailResponse) SetBody(v *GetSecretAsrDetailResponseBody) *GetSecretAsrDetailResponse {
	s.Body = v
	return s
}

type GetTotalPublicUrlRequest struct {
	// The ID of the call record.
	//
	// Log on to the [Phone Number Protection console](https://dypls.console.aliyun.com/dypls.htm#/account) and view **Call Record ID*	- on the **Call Record Query*	- page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2568900****
	CallId *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	// The call initiation time in the call record.
	//
	// Log on to the [Phone Number Protection console](https://dypls.console.aliyun.com/dypls.htm#/account). View **Call Initiated At*	- on the **Call Record Query*	- page, or view the call_time field in the Call Detail Record (CDR) receipt.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2019-03-05 12:00:00
	CallTime *string `json:"CallTime,omitempty" xml:"CallTime,omitempty"`
	// Specifies whether the verification on the binding ID is required.
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	CheckSubs *bool  `json:"CheckSubs,omitempty" xml:"CheckSubs,omitempty"`
	OwnerId   *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The key of the phone number pool. Log on to the [Phone Number Protection console](https://dypls.console.aliyun.com/dypls.htm#/account) and view the key of the phone number pool on the **Number Pool Management*	- page.
	//
	// This parameter is required.
	//
	// example:
	//
	// FC12256****
	PartnerKey           *string `json:"PartnerKey,omitempty" xml:"PartnerKey,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetTotalPublicUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTotalPublicUrlRequest) GoString() string {
	return s.String()
}

func (s *GetTotalPublicUrlRequest) SetCallId(v string) *GetTotalPublicUrlRequest {
	s.CallId = &v
	return s
}

func (s *GetTotalPublicUrlRequest) SetCallTime(v string) *GetTotalPublicUrlRequest {
	s.CallTime = &v
	return s
}

func (s *GetTotalPublicUrlRequest) SetCheckSubs(v bool) *GetTotalPublicUrlRequest {
	s.CheckSubs = &v
	return s
}

func (s *GetTotalPublicUrlRequest) SetOwnerId(v int64) *GetTotalPublicUrlRequest {
	s.OwnerId = &v
	return s
}

func (s *GetTotalPublicUrlRequest) SetPartnerKey(v string) *GetTotalPublicUrlRequest {
	s.PartnerKey = &v
	return s
}

func (s *GetTotalPublicUrlRequest) SetResourceOwnerAccount(v string) *GetTotalPublicUrlRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetTotalPublicUrlRequest) SetResourceOwnerId(v int64) *GetTotalPublicUrlRequest {
	s.ResourceOwnerId = &v
	return s
}

type GetTotalPublicUrlResponseBody struct {
	// The response code.
	//
	// 	- The value OK indicates that the request was successful.
	//
	// 	- Other status codes indicate that the request failed. For more information, see [Error codes](https://help.aliyun.com/document_detail/109196.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The download URLs of the recording files.
	Data *GetTotalPublicUrlResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1AB3CEF7-DCBE-488C-9C33-D180982CE031
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetTotalPublicUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTotalPublicUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetTotalPublicUrlResponseBody) SetCode(v string) *GetTotalPublicUrlResponseBody {
	s.Code = &v
	return s
}

func (s *GetTotalPublicUrlResponseBody) SetData(v *GetTotalPublicUrlResponseBodyData) *GetTotalPublicUrlResponseBody {
	s.Data = v
	return s
}

func (s *GetTotalPublicUrlResponseBody) SetMessage(v string) *GetTotalPublicUrlResponseBody {
	s.Message = &v
	return s
}

func (s *GetTotalPublicUrlResponseBody) SetRequestId(v string) *GetTotalPublicUrlResponseBody {
	s.RequestId = &v
	return s
}

type GetTotalPublicUrlResponseBodyData struct {
	// The download URL of the recorded call.
	//
	// >  The download URL of the recorded call is valid for 30 days.
	//
	// example:
	//
	// http://secret-axb-reco****cn-shanghai.aliyuncs.com/1000000820257625_66647243838006067251551752068865.mp3?Expires=1551****07&OSSAccessKeyId=LTAIP00vvvv****v&Signature=tK6Yq9KusU4n%2BZQWX****4/WmEA%3D
	PhonePublicUrl *string `json:"PhonePublicUrl,omitempty" xml:"PhonePublicUrl,omitempty"`
	// The download URL of the recorded ringing tone.
	//
	// >  The download URL of the recorded ringing tone is valid for 30 days.
	//
	// example:
	//
	// http://secret-ab-reco****cn-shanghai.aliyuncs.com/1000000820257625_66647243838006067251551752068865.mp3?Expires=155175****&OSSAccessKeyId=LTAIP00vvv****vv&Signature=tK6Yq9KusU4n%2BZQW****g4/WmEA%3D
	RingPublicUrl *string `json:"RingPublicUrl,omitempty" xml:"RingPublicUrl,omitempty"`
}

func (s GetTotalPublicUrlResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetTotalPublicUrlResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTotalPublicUrlResponseBodyData) SetPhonePublicUrl(v string) *GetTotalPublicUrlResponseBodyData {
	s.PhonePublicUrl = &v
	return s
}

func (s *GetTotalPublicUrlResponseBodyData) SetRingPublicUrl(v string) *GetTotalPublicUrlResponseBodyData {
	s.RingPublicUrl = &v
	return s
}

type GetTotalPublicUrlResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTotalPublicUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTotalPublicUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTotalPublicUrlResponse) GoString() string {
	return s.String()
}

func (s *GetTotalPublicUrlResponse) SetHeaders(v map[string]*string) *GetTotalPublicUrlResponse {
	s.Headers = v
	return s
}

func (s *GetTotalPublicUrlResponse) SetStatusCode(v int32) *GetTotalPublicUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTotalPublicUrlResponse) SetBody(v *GetTotalPublicUrlResponseBody) *GetTotalPublicUrlResponse {
	s.Body = v
	return s
}

type GetXConfigRequest struct {
	// 客户uid
	//
	// example:
	//
	// -
	CallerParentId *int64 `json:"CallerParentId,omitempty" xml:"CallerParentId,omitempty"`
	// 号码池key
	//
	// This parameter is required.
	//
	// example:
	//
	// FC5**********************a1a
	CustomerPoolKey *string `json:"CustomerPoolKey,omitempty" xml:"CustomerPoolKey,omitempty"`
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 请求去重ID, reqId最大长度为20位,接入方需要保持原子性
	//
	// This parameter is required.
	//
	// example:
	//
	// 564**********879
	ReqId                *string `json:"ReqId,omitempty" xml:"ReqId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// X号码
	//
	// This parameter is required.
	//
	// example:
	//
	// 17*******22
	TelX *string `json:"TelX,omitempty" xml:"TelX,omitempty"`
}

func (s GetXConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s GetXConfigRequest) GoString() string {
	return s.String()
}

func (s *GetXConfigRequest) SetCallerParentId(v int64) *GetXConfigRequest {
	s.CallerParentId = &v
	return s
}

func (s *GetXConfigRequest) SetCustomerPoolKey(v string) *GetXConfigRequest {
	s.CustomerPoolKey = &v
	return s
}

func (s *GetXConfigRequest) SetOwnerId(v int64) *GetXConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *GetXConfigRequest) SetReqId(v string) *GetXConfigRequest {
	s.ReqId = &v
	return s
}

func (s *GetXConfigRequest) SetResourceOwnerAccount(v string) *GetXConfigRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetXConfigRequest) SetResourceOwnerId(v int64) *GetXConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetXConfigRequest) SetTelX(v string) *GetXConfigRequest {
	s.TelX = &v
	return s
}

type GetXConfigResponseBody struct {
	// example:
	//
	// 0
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// 返回状态码 0000表示成功 其他表示失败
	//
	// example:
	//
	// 0000
	Code *string                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetXConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// 返回信息
	//
	// example:
	//
	// 成功
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 返回是否成功 true  表示成功 false表示失败
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetXConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetXConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetXConfigResponseBody) SetAccessDeniedDetail(v string) *GetXConfigResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetXConfigResponseBody) SetCode(v string) *GetXConfigResponseBody {
	s.Code = &v
	return s
}

func (s *GetXConfigResponseBody) SetData(v *GetXConfigResponseBodyData) *GetXConfigResponseBody {
	s.Data = v
	return s
}

func (s *GetXConfigResponseBody) SetMessage(v string) *GetXConfigResponseBody {
	s.Message = &v
	return s
}

func (s *GetXConfigResponseBody) SetSuccess(v bool) *GetXConfigResponseBody {
	s.Success = &v
	return s
}

type GetXConfigResponseBodyData struct {
	// 开/关呼叫能力状态： ‘0’：禁用； ‘1’：开启；
	//
	// example:
	//
	// 0
	CallAbility *string `json:"CallAbility,omitempty" xml:"CallAbility,omitempty"`
	// 是否透传来显为真实主叫： 00-非透传：互相拨打时都显示工作号; 10-透传：A客户为主叫时,B员工的来显为客户A号码;B员工为主叫时,A客户的来显为工作号; 默认为 00
	//
	// example:
	//
	// 00
	GNFlag *string `json:"GNFlag,omitempty" xml:"GNFlag,omitempty"`
	// 企业名片规则控制参数
	ReachJsons []*GetXConfigResponseBodyDataReachJsons `json:"ReachJsons,omitempty" xml:"ReachJsons,omitempty" type:"Repeated"`
	// 顺振控制参数
	SequenceCalls []*GetXConfigResponseBodyDataSequenceCalls `json:"SequenceCalls,omitempty" xml:"SequenceCalls,omitempty" type:"Repeated"`
	// 顺振结束时间 格式：HH:mm:ss 18:00:00
	//
	// example:
	//
	// 09:00:00
	SequenceEndTime *string `json:"SequenceEndTime,omitempty" xml:"SequenceEndTime,omitempty"`
	// 顺振开启时间 格式：HH:mm:ss 09:00:00
	//
	// example:
	//
	// 09:00:00
	SequenceStartTime *string `json:"SequenceStartTime,omitempty" xml:"SequenceStartTime,omitempty"`
	// 开/关短信功能状态： ‘0’：禁用； ‘1’：开启；
	//
	// example:
	//
	// 0
	SmsAbility *string `json:"SmsAbility,omitempty" xml:"SmsAbility,omitempty"`
	// 是否透传来显为真实主叫： 00-非透传：互相拨打时都显示工作号; 10-透传：A客户为主叫时,B员工的来显为客户A号码;B员工为主叫时,A客户的来显为工作号; 默认为 00
	//
	// example:
	//
	// 0
	SmsSignMode *string `json:"SmsSignMode,omitempty" xml:"SmsSignMode,omitempty"`
}

func (s GetXConfigResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetXConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetXConfigResponseBodyData) SetCallAbility(v string) *GetXConfigResponseBodyData {
	s.CallAbility = &v
	return s
}

func (s *GetXConfigResponseBodyData) SetGNFlag(v string) *GetXConfigResponseBodyData {
	s.GNFlag = &v
	return s
}

func (s *GetXConfigResponseBodyData) SetReachJsons(v []*GetXConfigResponseBodyDataReachJsons) *GetXConfigResponseBodyData {
	s.ReachJsons = v
	return s
}

func (s *GetXConfigResponseBodyData) SetSequenceCalls(v []*GetXConfigResponseBodyDataSequenceCalls) *GetXConfigResponseBodyData {
	s.SequenceCalls = v
	return s
}

func (s *GetXConfigResponseBodyData) SetSequenceEndTime(v string) *GetXConfigResponseBodyData {
	s.SequenceEndTime = &v
	return s
}

func (s *GetXConfigResponseBodyData) SetSequenceStartTime(v string) *GetXConfigResponseBodyData {
	s.SequenceStartTime = &v
	return s
}

func (s *GetXConfigResponseBodyData) SetSmsAbility(v string) *GetXConfigResponseBodyData {
	s.SmsAbility = &v
	return s
}

func (s *GetXConfigResponseBodyData) SetSmsSignMode(v string) *GetXConfigResponseBodyData {
	s.SmsSignMode = &v
	return s
}

type GetXConfigResponseBodyDataReachJsons struct {
	// 呼叫方向 1:员工B呼叫客户A 2:客户A呼叫员工B
	//
	// example:
	//
	// 1
	CallDir *string `json:"CallDir,omitempty" xml:"CallDir,omitempty"`
	// 通话状态 1:通话振铃 2:接通前 3:接通后 4:通话结束 5:已接通6:未接通
	//
	// example:
	//
	// 1
	CallStatus *string `json:"CallStatus,omitempty" xml:"CallStatus,omitempty"`
	// 接收方向 1:主叫 2:被叫
	//
	// example:
	//
	// 1
	ReceiveDir *string `json:"ReceiveDir,omitempty" xml:"ReceiveDir,omitempty"`
	// 规则ID
	//
	// example:
	//
	// 345
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// 规则名称
	//
	// example:
	//
	// 企业名片-短信规则
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// 规则类型： 1：企业名片-短信 2：企业名片-闪信 3：企业名片-视频 4：企业名片-音频
	//
	// example:
	//
	// 1
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	// 模板ID
	//
	// example:
	//
	// 12345
	TempId *string `json:"TempId,omitempty" xml:"TempId,omitempty"`
}

func (s GetXConfigResponseBodyDataReachJsons) String() string {
	return tea.Prettify(s)
}

func (s GetXConfigResponseBodyDataReachJsons) GoString() string {
	return s.String()
}

func (s *GetXConfigResponseBodyDataReachJsons) SetCallDir(v string) *GetXConfigResponseBodyDataReachJsons {
	s.CallDir = &v
	return s
}

func (s *GetXConfigResponseBodyDataReachJsons) SetCallStatus(v string) *GetXConfigResponseBodyDataReachJsons {
	s.CallStatus = &v
	return s
}

func (s *GetXConfigResponseBodyDataReachJsons) SetReceiveDir(v string) *GetXConfigResponseBodyDataReachJsons {
	s.ReceiveDir = &v
	return s
}

func (s *GetXConfigResponseBodyDataReachJsons) SetRuleId(v string) *GetXConfigResponseBodyDataReachJsons {
	s.RuleId = &v
	return s
}

func (s *GetXConfigResponseBodyDataReachJsons) SetRuleName(v string) *GetXConfigResponseBodyDataReachJsons {
	s.RuleName = &v
	return s
}

func (s *GetXConfigResponseBodyDataReachJsons) SetRuleType(v string) *GetXConfigResponseBodyDataReachJsons {
	s.RuleType = &v
	return s
}

func (s *GetXConfigResponseBodyDataReachJsons) SetTempId(v string) *GetXConfigResponseBodyDataReachJsons {
	s.TempId = &v
	return s
}

type GetXConfigResponseBodyDataSequenceCalls struct {
	// 顺振提示音放音编号，格式如callNoPlayCode
	//
	// example:
	//
	// 示例值示例值示例值
	SequenceCallNoPlayCode *string `json:"SequenceCallNoPlayCode,omitempty" xml:"SequenceCallNoPlayCode,omitempty"`
	// 顺振被叫号码
	//
	// example:
	//
	// 示例值示例值示例值
	SequenceCalledNo *string `json:"SequenceCalledNo,omitempty" xml:"SequenceCalledNo,omitempty"`
	// 接通后主被叫放音编号，格式如calledPlayCode
	//
	// example:
	//
	// 示例值示例值示例值
	SequenceCalledPlayCode *string `json:"SequenceCalledPlayCode,omitempty" xml:"SequenceCalledPlayCode,omitempty"`
}

func (s GetXConfigResponseBodyDataSequenceCalls) String() string {
	return tea.Prettify(s)
}

func (s GetXConfigResponseBodyDataSequenceCalls) GoString() string {
	return s.String()
}

func (s *GetXConfigResponseBodyDataSequenceCalls) SetSequenceCallNoPlayCode(v string) *GetXConfigResponseBodyDataSequenceCalls {
	s.SequenceCallNoPlayCode = &v
	return s
}

func (s *GetXConfigResponseBodyDataSequenceCalls) SetSequenceCalledNo(v string) *GetXConfigResponseBodyDataSequenceCalls {
	s.SequenceCalledNo = &v
	return s
}

func (s *GetXConfigResponseBodyDataSequenceCalls) SetSequenceCalledPlayCode(v string) *GetXConfigResponseBodyDataSequenceCalls {
	s.SequenceCalledPlayCode = &v
	return s
}

type GetXConfigResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetXConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetXConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s GetXConfigResponse) GoString() string {
	return s.String()
}

func (s *GetXConfigResponse) SetHeaders(v map[string]*string) *GetXConfigResponse {
	s.Headers = v
	return s
}

func (s *GetXConfigResponse) SetStatusCode(v int32) *GetXConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetXConfigResponse) SetBody(v *GetXConfigResponseBody) *GetXConfigResponse {
	s.Body = v
	return s
}

type GetXDefaultConfigRequest struct {
	// 客户uid
	//
	// example:
	//
	// -
	CallerParentId *int64 `json:"CallerParentId,omitempty" xml:"CallerParentId,omitempty"`
	// 号码池key
	//
	// This parameter is required.
	//
	// example:
	//
	// FC5**********************a1a
	CustomerPoolKey *string `json:"CustomerPoolKey,omitempty" xml:"CustomerPoolKey,omitempty"`
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 请求去重ID, reqId最大长度为20位,接入方需要保持原子性
	//
	// This parameter is required.
	//
	// example:
	//
	// 564**********879
	ReqId                *string `json:"ReqId,omitempty" xml:"ReqId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// X号码
	//
	// This parameter is required.
	//
	// example:
	//
	// 17*******22
	TelX *string `json:"TelX,omitempty" xml:"TelX,omitempty"`
}

func (s GetXDefaultConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s GetXDefaultConfigRequest) GoString() string {
	return s.String()
}

func (s *GetXDefaultConfigRequest) SetCallerParentId(v int64) *GetXDefaultConfigRequest {
	s.CallerParentId = &v
	return s
}

func (s *GetXDefaultConfigRequest) SetCustomerPoolKey(v string) *GetXDefaultConfigRequest {
	s.CustomerPoolKey = &v
	return s
}

func (s *GetXDefaultConfigRequest) SetOwnerId(v int64) *GetXDefaultConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *GetXDefaultConfigRequest) SetReqId(v string) *GetXDefaultConfigRequest {
	s.ReqId = &v
	return s
}

func (s *GetXDefaultConfigRequest) SetResourceOwnerAccount(v string) *GetXDefaultConfigRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetXDefaultConfigRequest) SetResourceOwnerId(v int64) *GetXDefaultConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetXDefaultConfigRequest) SetTelX(v string) *GetXDefaultConfigRequest {
	s.TelX = &v
	return s
}

type GetXDefaultConfigResponseBody struct {
	// example:
	//
	// 0
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// 返回状态码 0000表示成功 其他表示失败
	//
	// example:
	//
	// 0000
	Code *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetXDefaultConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// 返回信息
	//
	// example:
	//
	// 成功
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 返回是否成功 true  表示成功 false表示失败
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetXDefaultConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetXDefaultConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetXDefaultConfigResponseBody) SetAccessDeniedDetail(v string) *GetXDefaultConfigResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetXDefaultConfigResponseBody) SetCode(v string) *GetXDefaultConfigResponseBody {
	s.Code = &v
	return s
}

func (s *GetXDefaultConfigResponseBody) SetData(v *GetXDefaultConfigResponseBodyData) *GetXDefaultConfigResponseBody {
	s.Data = v
	return s
}

func (s *GetXDefaultConfigResponseBody) SetMessage(v string) *GetXDefaultConfigResponseBody {
	s.Message = &v
	return s
}

func (s *GetXDefaultConfigResponseBody) SetSuccess(v bool) *GetXDefaultConfigResponseBody {
	s.Success = &v
	return s
}

type GetXDefaultConfigResponseBodyData struct {
	// 开/关呼叫能力状态： ‘0’：禁用； ‘1’：开启；
	//
	// example:
	//
	// 0
	CallAbility *string `json:"CallAbility,omitempty" xml:"CallAbility,omitempty"`
	// 是否透传来显为真实主叫： 00-非透传：互相拨打时都显示工作号; 10-透传：A客户为主叫时,B员工的来显为客户A号码;B员工为主叫时,A客户的来显为工作号; 默认为 00
	//
	// example:
	//
	// 00
	GNFlag *string `json:"GNFlag,omitempty" xml:"GNFlag,omitempty"`
	// 企业名片规则控制参数
	ReachJson []*GetXDefaultConfigResponseBodyDataReachJson `json:"ReachJson,omitempty" xml:"ReachJson,omitempty" type:"Repeated"`
	// 顺振控制参数
	SequenceCall []*GetXDefaultConfigResponseBodyDataSequenceCall `json:"SequenceCall,omitempty" xml:"SequenceCall,omitempty" type:"Repeated"`
	// 顺振结束时间 格式：HH:mm:ss 18:00:00
	//
	// example:
	//
	// 09:00:00
	SequenceEndTime *string `json:"SequenceEndTime,omitempty" xml:"SequenceEndTime,omitempty"`
	// 顺振开启时间 格式：HH:mm:ss 09:00:00
	//
	// example:
	//
	// 09:00:00
	SequenceStartTime *string `json:"SequenceStartTime,omitempty" xml:"SequenceStartTime,omitempty"`
	// 开/关短信功能状态： ‘0’：禁用； ‘1’：开启；
	//
	// example:
	//
	// 0
	SmsAbility *string `json:"SmsAbility,omitempty" xml:"SmsAbility,omitempty"`
	// 是否透传来显为真实主叫： 00-非透传：互相拨打时都显示工作号; 10-透传：A客户为主叫时,B员工的来显为客户A号码;B员工为主叫时,A客户的来显为工作号; 默认为 00
	//
	// example:
	//
	// 0
	SmsSignMode *string `json:"SmsSignMode,omitempty" xml:"SmsSignMode,omitempty"`
}

func (s GetXDefaultConfigResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetXDefaultConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetXDefaultConfigResponseBodyData) SetCallAbility(v string) *GetXDefaultConfigResponseBodyData {
	s.CallAbility = &v
	return s
}

func (s *GetXDefaultConfigResponseBodyData) SetGNFlag(v string) *GetXDefaultConfigResponseBodyData {
	s.GNFlag = &v
	return s
}

func (s *GetXDefaultConfigResponseBodyData) SetReachJson(v []*GetXDefaultConfigResponseBodyDataReachJson) *GetXDefaultConfigResponseBodyData {
	s.ReachJson = v
	return s
}

func (s *GetXDefaultConfigResponseBodyData) SetSequenceCall(v []*GetXDefaultConfigResponseBodyDataSequenceCall) *GetXDefaultConfigResponseBodyData {
	s.SequenceCall = v
	return s
}

func (s *GetXDefaultConfigResponseBodyData) SetSequenceEndTime(v string) *GetXDefaultConfigResponseBodyData {
	s.SequenceEndTime = &v
	return s
}

func (s *GetXDefaultConfigResponseBodyData) SetSequenceStartTime(v string) *GetXDefaultConfigResponseBodyData {
	s.SequenceStartTime = &v
	return s
}

func (s *GetXDefaultConfigResponseBodyData) SetSmsAbility(v string) *GetXDefaultConfigResponseBodyData {
	s.SmsAbility = &v
	return s
}

func (s *GetXDefaultConfigResponseBodyData) SetSmsSignMode(v string) *GetXDefaultConfigResponseBodyData {
	s.SmsSignMode = &v
	return s
}

type GetXDefaultConfigResponseBodyDataReachJson struct {
	// 呼叫方向 1:员工B呼叫客户A 2:客户A呼叫员工B
	//
	// example:
	//
	// 1
	CallDir *string `json:"CallDir,omitempty" xml:"CallDir,omitempty"`
	// 通话状态 1:通话振铃 2:接通前 3:接通后 4:通话结束 5:已接通6:未接通
	//
	// example:
	//
	// 1
	CallStatus *string `json:"CallStatus,omitempty" xml:"CallStatus,omitempty"`
	// 接收方向 1:主叫 2:被叫
	//
	// example:
	//
	// 1
	ReceiveDir *string `json:"ReceiveDir,omitempty" xml:"ReceiveDir,omitempty"`
	// 规则ID
	//
	// example:
	//
	// 345
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// 规则名称
	//
	// example:
	//
	// 企业名片-短信规则
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// 规则类型： 1：企业名片-短信 2：企业名片-闪信 3：企业名片-视频 4：企业名片-音频
	//
	// example:
	//
	// 1
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	// 模板ID
	//
	// example:
	//
	// 12345
	TempId *string `json:"TempId,omitempty" xml:"TempId,omitempty"`
}

func (s GetXDefaultConfigResponseBodyDataReachJson) String() string {
	return tea.Prettify(s)
}

func (s GetXDefaultConfigResponseBodyDataReachJson) GoString() string {
	return s.String()
}

func (s *GetXDefaultConfigResponseBodyDataReachJson) SetCallDir(v string) *GetXDefaultConfigResponseBodyDataReachJson {
	s.CallDir = &v
	return s
}

func (s *GetXDefaultConfigResponseBodyDataReachJson) SetCallStatus(v string) *GetXDefaultConfigResponseBodyDataReachJson {
	s.CallStatus = &v
	return s
}

func (s *GetXDefaultConfigResponseBodyDataReachJson) SetReceiveDir(v string) *GetXDefaultConfigResponseBodyDataReachJson {
	s.ReceiveDir = &v
	return s
}

func (s *GetXDefaultConfigResponseBodyDataReachJson) SetRuleId(v string) *GetXDefaultConfigResponseBodyDataReachJson {
	s.RuleId = &v
	return s
}

func (s *GetXDefaultConfigResponseBodyDataReachJson) SetRuleName(v string) *GetXDefaultConfigResponseBodyDataReachJson {
	s.RuleName = &v
	return s
}

func (s *GetXDefaultConfigResponseBodyDataReachJson) SetRuleType(v string) *GetXDefaultConfigResponseBodyDataReachJson {
	s.RuleType = &v
	return s
}

func (s *GetXDefaultConfigResponseBodyDataReachJson) SetTempId(v string) *GetXDefaultConfigResponseBodyDataReachJson {
	s.TempId = &v
	return s
}

type GetXDefaultConfigResponseBodyDataSequenceCall struct {
	// 顺振提示音放音编号，格式如callNoPlayCode
	//
	// example:
	//
	// 示例值
	SequenceCallNoPlayCode *string `json:"SequenceCallNoPlayCode,omitempty" xml:"SequenceCallNoPlayCode,omitempty"`
	// 顺振被叫号码
	//
	// example:
	//
	// 示例值示例值示例值
	SequenceCalledNo *string `json:"SequenceCalledNo,omitempty" xml:"SequenceCalledNo,omitempty"`
	// 接通后主被叫放音编号，格式如calledPlayCode
	//
	// example:
	//
	// 示例值示例值示例值
	SequenceCalledPlayCode *string `json:"SequenceCalledPlayCode,omitempty" xml:"SequenceCalledPlayCode,omitempty"`
}

func (s GetXDefaultConfigResponseBodyDataSequenceCall) String() string {
	return tea.Prettify(s)
}

func (s GetXDefaultConfigResponseBodyDataSequenceCall) GoString() string {
	return s.String()
}

func (s *GetXDefaultConfigResponseBodyDataSequenceCall) SetSequenceCallNoPlayCode(v string) *GetXDefaultConfigResponseBodyDataSequenceCall {
	s.SequenceCallNoPlayCode = &v
	return s
}

func (s *GetXDefaultConfigResponseBodyDataSequenceCall) SetSequenceCalledNo(v string) *GetXDefaultConfigResponseBodyDataSequenceCall {
	s.SequenceCalledNo = &v
	return s
}

func (s *GetXDefaultConfigResponseBodyDataSequenceCall) SetSequenceCalledPlayCode(v string) *GetXDefaultConfigResponseBodyDataSequenceCall {
	s.SequenceCalledPlayCode = &v
	return s
}

type GetXDefaultConfigResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetXDefaultConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetXDefaultConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s GetXDefaultConfigResponse) GoString() string {
	return s.String()
}

func (s *GetXDefaultConfigResponse) SetHeaders(v map[string]*string) *GetXDefaultConfigResponse {
	s.Headers = v
	return s
}

func (s *GetXDefaultConfigResponse) SetStatusCode(v int32) *GetXDefaultConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetXDefaultConfigResponse) SetBody(v *GetXDefaultConfigResponseBody) *GetXDefaultConfigResponse {
	s.Body = v
	return s
}

type ListXTelephonesRequest struct {
	// 客户uid
	//
	// example:
	//
	// -
	CallerParentId *int64 `json:"CallerParentId,omitempty" xml:"CallerParentId,omitempty"`
	// 号码池key
	//
	// example:
	//
	// FC5**********************a1a
	CustomerPoolKey *string `json:"CustomerPoolKey,omitempty" xml:"CustomerPoolKey,omitempty"`
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 页码从1开始
	//
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// 每页条数
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 请求去重ID, reqId最大长度为20位,接入方需要保持原子性
	//
	// This parameter is required.
	//
	// example:
	//
	// 564**********879
	ReqId                *string `json:"ReqId,omitempty" xml:"ReqId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ListXTelephonesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListXTelephonesRequest) GoString() string {
	return s.String()
}

func (s *ListXTelephonesRequest) SetCallerParentId(v int64) *ListXTelephonesRequest {
	s.CallerParentId = &v
	return s
}

func (s *ListXTelephonesRequest) SetCustomerPoolKey(v string) *ListXTelephonesRequest {
	s.CustomerPoolKey = &v
	return s
}

func (s *ListXTelephonesRequest) SetOwnerId(v int64) *ListXTelephonesRequest {
	s.OwnerId = &v
	return s
}

func (s *ListXTelephonesRequest) SetPageNo(v int64) *ListXTelephonesRequest {
	s.PageNo = &v
	return s
}

func (s *ListXTelephonesRequest) SetPageSize(v int64) *ListXTelephonesRequest {
	s.PageSize = &v
	return s
}

func (s *ListXTelephonesRequest) SetReqId(v string) *ListXTelephonesRequest {
	s.ReqId = &v
	return s
}

func (s *ListXTelephonesRequest) SetResourceOwnerAccount(v string) *ListXTelephonesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListXTelephonesRequest) SetResourceOwnerId(v int64) *ListXTelephonesRequest {
	s.ResourceOwnerId = &v
	return s
}

type ListXTelephonesResponseBody struct {
	// example:
	//
	// 0
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// 返回状态码 0000表示成功 其他表示失败
	//
	// example:
	//
	// 0000
	Code *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListXTelephonesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// 返回信息
	//
	// example:
	//
	// 成功
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 返回是否成功 true  表示成功 false表示失败
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListXTelephonesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListXTelephonesResponseBody) GoString() string {
	return s.String()
}

func (s *ListXTelephonesResponseBody) SetAccessDeniedDetail(v string) *ListXTelephonesResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListXTelephonesResponseBody) SetCode(v string) *ListXTelephonesResponseBody {
	s.Code = &v
	return s
}

func (s *ListXTelephonesResponseBody) SetData(v *ListXTelephonesResponseBodyData) *ListXTelephonesResponseBody {
	s.Data = v
	return s
}

func (s *ListXTelephonesResponseBody) SetMessage(v string) *ListXTelephonesResponseBody {
	s.Message = &v
	return s
}

func (s *ListXTelephonesResponseBody) SetSuccess(v bool) *ListXTelephonesResponseBody {
	s.Success = &v
	return s
}

type ListXTelephonesResponseBodyData struct {
	// 数据集合
	List []*ListXTelephonesResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// 页码
	//
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// 每页条数
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 符合查询条件的总数量
	//
	// example:
	//
	// 50
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListXTelephonesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListXTelephonesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListXTelephonesResponseBodyData) SetList(v []*ListXTelephonesResponseBodyDataList) *ListXTelephonesResponseBodyData {
	s.List = v
	return s
}

func (s *ListXTelephonesResponseBodyData) SetPageNo(v int64) *ListXTelephonesResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *ListXTelephonesResponseBodyData) SetPageSize(v int64) *ListXTelephonesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListXTelephonesResponseBodyData) SetTotal(v int64) *ListXTelephonesResponseBodyData {
	s.Total = &v
	return s
}

type ListXTelephonesResponseBodyDataList struct {
	// 绑定失败原因
	//
	// example:
	//
	// 绑定失败用户身份证黑名单
	AuthMsg *string `json:"AuthMsg,omitempty" xml:"AuthMsg,omitempty"`
	// 绑定时间
	//
	// example:
	//
	// 2024-08-29 17:23:58
	BindTime *string `json:"BindTime,omitempty" xml:"BindTime,omitempty"`
	// 购买时间
	//
	// example:
	//
	// 2024-08-29 17:23:58
	BuyTime *string `json:"BuyTime,omitempty" xml:"BuyTime,omitempty"`
	// 客户号码池key
	//
	// example:
	//
	// FC533e6eeb81f4400c87ef3745a21a1a
	CustomerPoolKey *string `json:"CustomerPoolKey,omitempty" xml:"CustomerPoolKey,omitempty"`
	// 号码池名称
	//
	// example:
	//
	// 测试号码池
	CustomerPoolName *string `json:"CustomerPoolName,omitempty" xml:"CustomerPoolName,omitempty"`
	// 释放时间
	//
	// example:
	//
	// 2024-08-29 17:23:58
	ReleaseTime *string `json:"ReleaseTime,omitempty" xml:"ReleaseTime,omitempty"`
	// 短信开通状态：0 未开通 1已开通
	//
	// example:
	//
	// 0
	SmsStatus *string `json:"SmsStatus,omitempty" xml:"SmsStatus,omitempty"`
	// X号码
	//
	// example:
	//
	// 17816876546
	Telephone *string `json:"Telephone,omitempty" xml:"Telephone,omitempty"`
	// 号码状态：0 空闲中 1 调拨完成待购买 2购买完成待认证  3 实名认证中  4 实名认证成功  5 认证失败  6 解绑中 7 解绑失败 8已释放 99 超时释放
	//
	// example:
	//
	// 0
	TelephoneStatus *string `json:"TelephoneStatus,omitempty" xml:"TelephoneStatus,omitempty"`
	// 解绑时间
	//
	// example:
	//
	// 2024-08-29 17:23:58
	UnbindTime *string `json:"UnbindTime,omitempty" xml:"UnbindTime,omitempty"`
}

func (s ListXTelephonesResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s ListXTelephonesResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListXTelephonesResponseBodyDataList) SetAuthMsg(v string) *ListXTelephonesResponseBodyDataList {
	s.AuthMsg = &v
	return s
}

func (s *ListXTelephonesResponseBodyDataList) SetBindTime(v string) *ListXTelephonesResponseBodyDataList {
	s.BindTime = &v
	return s
}

func (s *ListXTelephonesResponseBodyDataList) SetBuyTime(v string) *ListXTelephonesResponseBodyDataList {
	s.BuyTime = &v
	return s
}

func (s *ListXTelephonesResponseBodyDataList) SetCustomerPoolKey(v string) *ListXTelephonesResponseBodyDataList {
	s.CustomerPoolKey = &v
	return s
}

func (s *ListXTelephonesResponseBodyDataList) SetCustomerPoolName(v string) *ListXTelephonesResponseBodyDataList {
	s.CustomerPoolName = &v
	return s
}

func (s *ListXTelephonesResponseBodyDataList) SetReleaseTime(v string) *ListXTelephonesResponseBodyDataList {
	s.ReleaseTime = &v
	return s
}

func (s *ListXTelephonesResponseBodyDataList) SetSmsStatus(v string) *ListXTelephonesResponseBodyDataList {
	s.SmsStatus = &v
	return s
}

func (s *ListXTelephonesResponseBodyDataList) SetTelephone(v string) *ListXTelephonesResponseBodyDataList {
	s.Telephone = &v
	return s
}

func (s *ListXTelephonesResponseBodyDataList) SetTelephoneStatus(v string) *ListXTelephonesResponseBodyDataList {
	s.TelephoneStatus = &v
	return s
}

func (s *ListXTelephonesResponseBodyDataList) SetUnbindTime(v string) *ListXTelephonesResponseBodyDataList {
	s.UnbindTime = &v
	return s
}

type ListXTelephonesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListXTelephonesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListXTelephonesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListXTelephonesResponse) GoString() string {
	return s.String()
}

func (s *ListXTelephonesResponse) SetHeaders(v map[string]*string) *ListXTelephonesResponse {
	s.Headers = v
	return s
}

func (s *ListXTelephonesResponse) SetStatusCode(v int32) *ListXTelephonesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListXTelephonesResponse) SetBody(v *ListXTelephonesResponseBody) *ListXTelephonesResponse {
	s.Body = v
	return s
}

type LockSecretNoRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The key of the phone number pool.
	//
	// Log on to the [Phone Number Protection console](https://dypls.console.aliyun.com/dypls.htm#/account) and view the key of the phone number pool on the **Number Pool Management*	- page.
	//
	// This parameter is required.
	//
	// example:
	//
	// FC123****
	PoolKey              *string `json:"PoolKey,omitempty" xml:"PoolKey,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The private number that you want to lock. You must enter a complete mobile phone number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1300000****
	SecretNo *string `json:"SecretNo,omitempty" xml:"SecretNo,omitempty"`
}

func (s LockSecretNoRequest) String() string {
	return tea.Prettify(s)
}

func (s LockSecretNoRequest) GoString() string {
	return s.String()
}

func (s *LockSecretNoRequest) SetOwnerId(v int64) *LockSecretNoRequest {
	s.OwnerId = &v
	return s
}

func (s *LockSecretNoRequest) SetPoolKey(v string) *LockSecretNoRequest {
	s.PoolKey = &v
	return s
}

func (s *LockSecretNoRequest) SetResourceOwnerAccount(v string) *LockSecretNoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *LockSecretNoRequest) SetResourceOwnerId(v int64) *LockSecretNoRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *LockSecretNoRequest) SetSecretNo(v string) *LockSecretNoRequest {
	s.SecretNo = &v
	return s
}

type LockSecretNoResponseBody struct {
	// The response code.
	//
	// 	- The value OK indicates that the request was successful.
	//
	// 	- Other status codes indicate that the request failed. For more information, see [Error codes](https://help.aliyun.com/document_detail/109196.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2D1AEB96-96D0-454E-B0DC-AE2A8DF08020
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s LockSecretNoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s LockSecretNoResponseBody) GoString() string {
	return s.String()
}

func (s *LockSecretNoResponseBody) SetCode(v string) *LockSecretNoResponseBody {
	s.Code = &v
	return s
}

func (s *LockSecretNoResponseBody) SetMessage(v string) *LockSecretNoResponseBody {
	s.Message = &v
	return s
}

func (s *LockSecretNoResponseBody) SetRequestId(v string) *LockSecretNoResponseBody {
	s.RequestId = &v
	return s
}

type LockSecretNoResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *LockSecretNoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s LockSecretNoResponse) String() string {
	return tea.Prettify(s)
}

func (s LockSecretNoResponse) GoString() string {
	return s.String()
}

func (s *LockSecretNoResponse) SetHeaders(v map[string]*string) *LockSecretNoResponse {
	s.Headers = v
	return s
}

func (s *LockSecretNoResponse) SetStatusCode(v int32) *LockSecretNoResponse {
	s.StatusCode = &v
	return s
}

func (s *LockSecretNoResponse) SetBody(v *LockSecretNoResponseBody) *LockSecretNoResponse {
	s.Body = v
	return s
}

type OperateAxgGroupRequest struct {
	// The group ID in the AXG binding.
	//
	// You can view the group ID by using either of the following methods:
	//
	// 	- On the **Number Pool Management*	- page in the [Phone Number Protection console](https://dypls.console.aliyun.com/dypls.htm#/account), filter AXG private numbers and click **Number Group G Management*	- to view the group IDs of number groups G.****
	//
	// 	- If the [CreateAxgGroup](https://help.aliyun.com/document_detail/110250.html) operation is called to create number group G, the value of **GroupId*	- in the response of the CreateAxgGroup operation is specified as the value of the request parameter **GroupId*	- of the OperateAxgGroup operation.
	//
	// >  Number group G must have one or more phone numbers.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The phone numbers that you add to number group G. Separate the phone numbers with commas (,). You can add up to 200 phone numbers at a time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1390000****,1380000****
	Numbers *string `json:"Numbers,omitempty" xml:"Numbers,omitempty"`
	// The type of the operation on number group G. Valid values:
	//
	// 	- **addNumbers**: adds phone numbers to number group G.
	//
	// 	- **deleteNumbers**: deletes phone numbers from number group G.
	//
	// 	- **overwriteNumbers**: replaces all phone numbers in number group G. All the original phone numbers are deleted from number group G, and new phone numbers are added to number group G.
	//
	// >
	//
	// 	- When you replace all phone numbers in number group G, the quantity of original phone numbers in number group G cannot exceed 200.
	//
	// 	- You can add up to 200 phone numbers to number group G at a time.
	//
	// This parameter is required.
	//
	// example:
	//
	// addNumbers
	OperateType *string `json:"OperateType,omitempty" xml:"OperateType,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The key of the phone number pool. Log on to the [Phone Number Protection console](https://dypls.console.aliyun.com/dypls.htm#/account) and view the key of the phone number pool on the **Number Pool Management*	- page.
	//
	// This parameter is required.
	//
	// example:
	//
	// FC123456
	PoolKey              *string `json:"PoolKey,omitempty" xml:"PoolKey,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s OperateAxgGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s OperateAxgGroupRequest) GoString() string {
	return s.String()
}

func (s *OperateAxgGroupRequest) SetGroupId(v int64) *OperateAxgGroupRequest {
	s.GroupId = &v
	return s
}

func (s *OperateAxgGroupRequest) SetNumbers(v string) *OperateAxgGroupRequest {
	s.Numbers = &v
	return s
}

func (s *OperateAxgGroupRequest) SetOperateType(v string) *OperateAxgGroupRequest {
	s.OperateType = &v
	return s
}

func (s *OperateAxgGroupRequest) SetOwnerId(v int64) *OperateAxgGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *OperateAxgGroupRequest) SetPoolKey(v string) *OperateAxgGroupRequest {
	s.PoolKey = &v
	return s
}

func (s *OperateAxgGroupRequest) SetResourceOwnerAccount(v string) *OperateAxgGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *OperateAxgGroupRequest) SetResourceOwnerId(v int64) *OperateAxgGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

type OperateAxgGroupResponseBody struct {
	// The response code.
	//
	// 	- The value OK indicates that the request was successful.
	//
	// 	- Other values indicate that the request failed. For more information, see [Error codes](https://help.aliyun.com/document_detail/109196.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 986BCB6D-C9BF-42F9-91CE-3A9901233D36
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OperateAxgGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OperateAxgGroupResponseBody) GoString() string {
	return s.String()
}

func (s *OperateAxgGroupResponseBody) SetCode(v string) *OperateAxgGroupResponseBody {
	s.Code = &v
	return s
}

func (s *OperateAxgGroupResponseBody) SetMessage(v string) *OperateAxgGroupResponseBody {
	s.Message = &v
	return s
}

func (s *OperateAxgGroupResponseBody) SetRequestId(v string) *OperateAxgGroupResponseBody {
	s.RequestId = &v
	return s
}

type OperateAxgGroupResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OperateAxgGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OperateAxgGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s OperateAxgGroupResponse) GoString() string {
	return s.String()
}

func (s *OperateAxgGroupResponse) SetHeaders(v map[string]*string) *OperateAxgGroupResponse {
	s.Headers = v
	return s
}

func (s *OperateAxgGroupResponse) SetStatusCode(v int32) *OperateAxgGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *OperateAxgGroupResponse) SetBody(v *OperateAxgGroupResponseBody) *OperateAxgGroupResponse {
	s.Body = v
	return s
}

type OperateBlackNoRequest struct {
	// The phone number to be added to or deleted from the blacklist.
	//
	// This parameter is required.
	//
	// example:
	//
	// 150****0000
	BlackNo *string `json:"BlackNo,omitempty" xml:"BlackNo,omitempty"`
	// The type of the operation on the phone number. Valid values:
	//
	// 	- **AddBlack**: adds the phone number to the blacklist.
	//
	// 	- **DeleteBlack**: deletes the phone number from the blacklist.
	//
	// This parameter is required.
	//
	// example:
	//
	// AddBlack
	OperateType *string `json:"OperateType,omitempty" xml:"OperateType,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The key of the phone number pool. Log on to the [Phone Number Protection console](https://dypls.console.aliyun.com/dypls.htm#/account) and view the key of the phone number pool on the **Number Pool Management*	- page.
	//
	// This parameter is required.
	//
	// example:
	//
	// FC123456****
	PoolKey              *string `json:"PoolKey,omitempty" xml:"PoolKey,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The description.
	//
	// example:
	//
	// abcdef
	Tips *string `json:"Tips,omitempty" xml:"Tips,omitempty"`
}

func (s OperateBlackNoRequest) String() string {
	return tea.Prettify(s)
}

func (s OperateBlackNoRequest) GoString() string {
	return s.String()
}

func (s *OperateBlackNoRequest) SetBlackNo(v string) *OperateBlackNoRequest {
	s.BlackNo = &v
	return s
}

func (s *OperateBlackNoRequest) SetOperateType(v string) *OperateBlackNoRequest {
	s.OperateType = &v
	return s
}

func (s *OperateBlackNoRequest) SetOwnerId(v int64) *OperateBlackNoRequest {
	s.OwnerId = &v
	return s
}

func (s *OperateBlackNoRequest) SetPoolKey(v string) *OperateBlackNoRequest {
	s.PoolKey = &v
	return s
}

func (s *OperateBlackNoRequest) SetResourceOwnerAccount(v string) *OperateBlackNoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *OperateBlackNoRequest) SetResourceOwnerId(v int64) *OperateBlackNoRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *OperateBlackNoRequest) SetTips(v string) *OperateBlackNoRequest {
	s.Tips = &v
	return s
}

type OperateBlackNoResponseBody struct {
	// The response code.
	//
	// 	- The value OK indicates that the request was successful.
	//
	// 	- Other status codes indicate that the request failed. For more information, see [Error codes](https://help.aliyun.com/document_detail/109196.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2D1AEB96-96D0-454E-B0DC-AE2A8DF08020
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OperateBlackNoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OperateBlackNoResponseBody) GoString() string {
	return s.String()
}

func (s *OperateBlackNoResponseBody) SetCode(v string) *OperateBlackNoResponseBody {
	s.Code = &v
	return s
}

func (s *OperateBlackNoResponseBody) SetMessage(v string) *OperateBlackNoResponseBody {
	s.Message = &v
	return s
}

func (s *OperateBlackNoResponseBody) SetRequestId(v string) *OperateBlackNoResponseBody {
	s.RequestId = &v
	return s
}

type OperateBlackNoResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OperateBlackNoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OperateBlackNoResponse) String() string {
	return tea.Prettify(s)
}

func (s OperateBlackNoResponse) GoString() string {
	return s.String()
}

func (s *OperateBlackNoResponse) SetHeaders(v map[string]*string) *OperateBlackNoResponse {
	s.Headers = v
	return s
}

func (s *OperateBlackNoResponse) SetStatusCode(v int32) *OperateBlackNoResponse {
	s.StatusCode = &v
	return s
}

func (s *OperateBlackNoResponse) SetBody(v *OperateBlackNoResponseBody) *OperateBlackNoResponse {
	s.Body = v
	return s
}

type QueryAxbBindFixedLineRequest struct {
	// 应用id
	//
	// This parameter is required.
	//
	// example:
	//
	// ALPT_1234
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 业务id，消息请求标识
	//
	// This parameter is required.
	//
	// example:
	//
	// 3ererrrdrrrr
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// A/B号码，queryType=1时，必传
	//
	// example:
	//
	// 示例值
	Phone *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	// 查询类型 0：根据绑定id查询 1：根据X和A/B号码查询
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	QueryType            *string `json:"QueryType,omitempty" xml:"QueryType,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 绑定id， queryType=0时，必传
	//
	// example:
	//
	// A20304o0200303004j
	SubId *string `json:"SubId,omitempty" xml:"SubId,omitempty"`
	// x号码，queryType=1时，必传
	//
	// example:
	//
	// 示例值示例值示例值
	TelX *string `json:"TelX,omitempty" xml:"TelX,omitempty"`
	// 格式yyyyMMddHHmmssSSS，时间采用北京时间，24小时制，精确至毫秒
	//
	// This parameter is required.
	//
	// example:
	//
	// 20161114143116001
	Ts *string `json:"Ts,omitempty" xml:"Ts,omitempty"`
}

func (s QueryAxbBindFixedLineRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAxbBindFixedLineRequest) GoString() string {
	return s.String()
}

func (s *QueryAxbBindFixedLineRequest) SetAppId(v string) *QueryAxbBindFixedLineRequest {
	s.AppId = &v
	return s
}

func (s *QueryAxbBindFixedLineRequest) SetOrderId(v string) *QueryAxbBindFixedLineRequest {
	s.OrderId = &v
	return s
}

func (s *QueryAxbBindFixedLineRequest) SetOwnerId(v int64) *QueryAxbBindFixedLineRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryAxbBindFixedLineRequest) SetPhone(v string) *QueryAxbBindFixedLineRequest {
	s.Phone = &v
	return s
}

func (s *QueryAxbBindFixedLineRequest) SetQueryType(v string) *QueryAxbBindFixedLineRequest {
	s.QueryType = &v
	return s
}

func (s *QueryAxbBindFixedLineRequest) SetResourceOwnerAccount(v string) *QueryAxbBindFixedLineRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryAxbBindFixedLineRequest) SetResourceOwnerId(v int64) *QueryAxbBindFixedLineRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryAxbBindFixedLineRequest) SetSubId(v string) *QueryAxbBindFixedLineRequest {
	s.SubId = &v
	return s
}

func (s *QueryAxbBindFixedLineRequest) SetTelX(v string) *QueryAxbBindFixedLineRequest {
	s.TelX = &v
	return s
}

func (s *QueryAxbBindFixedLineRequest) SetTs(v string) *QueryAxbBindFixedLineRequest {
	s.Ts = &v
	return s
}

type QueryAxbBindFixedLineResponseBody struct {
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// 响应码
	//
	// example:
	//
	// 0
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 绑定信息
	Data *QueryAxbBindFixedLineResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// 描述信息
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 3179F199-C6C5-5963-85A6-21CBA2F47592
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 处理是否成功 true-成功
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryAxbBindFixedLineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAxbBindFixedLineResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAxbBindFixedLineResponseBody) SetAccessDeniedDetail(v string) *QueryAxbBindFixedLineResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *QueryAxbBindFixedLineResponseBody) SetCode(v string) *QueryAxbBindFixedLineResponseBody {
	s.Code = &v
	return s
}

func (s *QueryAxbBindFixedLineResponseBody) SetData(v *QueryAxbBindFixedLineResponseBodyData) *QueryAxbBindFixedLineResponseBody {
	s.Data = v
	return s
}

func (s *QueryAxbBindFixedLineResponseBody) SetMessage(v string) *QueryAxbBindFixedLineResponseBody {
	s.Message = &v
	return s
}

func (s *QueryAxbBindFixedLineResponseBody) SetRequestId(v string) *QueryAxbBindFixedLineResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryAxbBindFixedLineResponseBody) SetSuccess(v bool) *QueryAxbBindFixedLineResponseBody {
	s.Success = &v
	return s
}

type QueryAxbBindFixedLineResponseBodyData struct {
	// 接通前放音编码，放音编码必须包含3个场景的编码。按照“B->X,A->X,其他号码->X”的顺序填写编码，编码之间以逗号分隔。  比如：“1,2,3”表示B->X放音编号为1，A->X放音编号为2， 其他号码->X放音编号为3
	//
	// example:
	//
	// 10001,10002,10003
	Anucode *string `json:"Anucode,omitempty" xml:"Anucode,omitempty"`
	// 接通后被叫侧放音编码
	//
	// example:
	//
	// 10001,10002
	Anucodecalled *string `json:"Anucodecalled,omitempty" xml:"Anucodecalled,omitempty"`
	// 隐私号码区号
	//
	// example:
	//
	// 010
	Areacode *string `json:"Areacode,omitempty" xml:"Areacode,omitempty"`
	// 绑定过期时间
	//
	// example:
	//
	// 60
	Expiration *string `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	// 额外字段
	Extra *QueryAxbBindFixedLineResponseBodyDataExtra `json:"Extra,omitempty" xml:"Extra,omitempty" type:"Struct"`
	// 接入商自有字段，最大100字符长度
	//
	// example:
	//
	// 19394
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// 绑定id
	//
	// example:
	//
	// GHX0534X202504221531579290029-2-1-aliaxb
	Subid *string `json:"Subid,omitempty" xml:"Subid,omitempty"`
	// 绑定时间，格式为yyyyMMddHHmmss。时间采用北京时间，24小时制。
	//
	// example:
	//
	// 20250421141723
	Subts *string `json:"Subts,omitempty" xml:"Subts,omitempty"`
	// 接通后主叫侧放音编码
	//
	// example:
	//
	// 示例值示例值
	TAnucodeConnect *string `json:"TAnucodeConnect,omitempty" xml:"TAnucodeConnect,omitempty"`
	// A号码
	//
	// example:
	//
	// 15500001111
	TelA *string `json:"TelA,omitempty" xml:"TelA,omitempty"`
	// B号码
	//
	// example:
	//
	// 15500002222
	TelB *string `json:"TelB,omitempty" xml:"TelB,omitempty"`
	// 小号号码
	//
	// example:
	//
	// 19700002222
	TelX *string `json:"TelX,omitempty" xml:"TelX,omitempty"`
}

func (s QueryAxbBindFixedLineResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryAxbBindFixedLineResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryAxbBindFixedLineResponseBodyData) SetAnucode(v string) *QueryAxbBindFixedLineResponseBodyData {
	s.Anucode = &v
	return s
}

func (s *QueryAxbBindFixedLineResponseBodyData) SetAnucodecalled(v string) *QueryAxbBindFixedLineResponseBodyData {
	s.Anucodecalled = &v
	return s
}

func (s *QueryAxbBindFixedLineResponseBodyData) SetAreacode(v string) *QueryAxbBindFixedLineResponseBodyData {
	s.Areacode = &v
	return s
}

func (s *QueryAxbBindFixedLineResponseBodyData) SetExpiration(v string) *QueryAxbBindFixedLineResponseBodyData {
	s.Expiration = &v
	return s
}

func (s *QueryAxbBindFixedLineResponseBodyData) SetExtra(v *QueryAxbBindFixedLineResponseBodyDataExtra) *QueryAxbBindFixedLineResponseBodyData {
	s.Extra = v
	return s
}

func (s *QueryAxbBindFixedLineResponseBodyData) SetRemark(v string) *QueryAxbBindFixedLineResponseBodyData {
	s.Remark = &v
	return s
}

func (s *QueryAxbBindFixedLineResponseBodyData) SetSubid(v string) *QueryAxbBindFixedLineResponseBodyData {
	s.Subid = &v
	return s
}

func (s *QueryAxbBindFixedLineResponseBodyData) SetSubts(v string) *QueryAxbBindFixedLineResponseBodyData {
	s.Subts = &v
	return s
}

func (s *QueryAxbBindFixedLineResponseBodyData) SetTAnucodeConnect(v string) *QueryAxbBindFixedLineResponseBodyData {
	s.TAnucodeConnect = &v
	return s
}

func (s *QueryAxbBindFixedLineResponseBodyData) SetTelA(v string) *QueryAxbBindFixedLineResponseBodyData {
	s.TelA = &v
	return s
}

func (s *QueryAxbBindFixedLineResponseBodyData) SetTelB(v string) *QueryAxbBindFixedLineResponseBodyData {
	s.TelB = &v
	return s
}

func (s *QueryAxbBindFixedLineResponseBodyData) SetTelX(v string) *QueryAxbBindFixedLineResponseBodyData {
	s.TelX = &v
	return s
}

type QueryAxbBindFixedLineResponseBodyDataExtra struct {
	// 录音控制，仅下列值有效。默认是0（不开通录音功能）。 0：不录音 1：录音
	//
	// example:
	//
	// 1
	Callrecording *string `json:"Callrecording,omitempty" xml:"Callrecording,omitempty"`
}

func (s QueryAxbBindFixedLineResponseBodyDataExtra) String() string {
	return tea.Prettify(s)
}

func (s QueryAxbBindFixedLineResponseBodyDataExtra) GoString() string {
	return s.String()
}

func (s *QueryAxbBindFixedLineResponseBodyDataExtra) SetCallrecording(v string) *QueryAxbBindFixedLineResponseBodyDataExtra {
	s.Callrecording = &v
	return s
}

type QueryAxbBindFixedLineResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAxbBindFixedLineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryAxbBindFixedLineResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAxbBindFixedLineResponse) GoString() string {
	return s.String()
}

func (s *QueryAxbBindFixedLineResponse) SetHeaders(v map[string]*string) *QueryAxbBindFixedLineResponse {
	s.Headers = v
	return s
}

func (s *QueryAxbBindFixedLineResponse) SetStatusCode(v int32) *QueryAxbBindFixedLineResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAxbBindFixedLineResponse) SetBody(v *QueryAxbBindFixedLineResponseBody) *QueryAxbBindFixedLineResponse {
	s.Body = v
	return s
}

type QueryAxnBindFixedLineRequest struct {
	// 业务id标识，由阿里云分配给客户侧
	//
	// This parameter is required.
	//
	// example:
	//
	// alitest
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 消息请求唯一标识
	//
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// A号码，queryType=1时，必传
	//
	// example:
	//
	// 15500001111
	Phone *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	// 查询类型 0：根据绑定id查询 1：根据A号码查询 2：根据X查询
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	QueryType            *string `json:"QueryType,omitempty" xml:"QueryType,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 绑定id queryType=0时，必传
	//
	// example:
	//
	// GHX0534X202504221531579290029-2-1-aliaxn
	SubId *string `json:"SubId,omitempty" xml:"SubId,omitempty"`
	// 隐私号码
	//
	// example:
	//
	// 05718950
	TelX *string `json:"TelX,omitempty" xml:"TelX,omitempty"`
	// 格式yyyyMMddHHmmssSSS，时间采用北京时间，24小时制，精确至毫秒
	//
	// This parameter is required.
	//
	// example:
	//
	// 20250421141723098
	Ts *string `json:"Ts,omitempty" xml:"Ts,omitempty"`
}

func (s QueryAxnBindFixedLineRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAxnBindFixedLineRequest) GoString() string {
	return s.String()
}

func (s *QueryAxnBindFixedLineRequest) SetAppId(v string) *QueryAxnBindFixedLineRequest {
	s.AppId = &v
	return s
}

func (s *QueryAxnBindFixedLineRequest) SetOrderId(v string) *QueryAxnBindFixedLineRequest {
	s.OrderId = &v
	return s
}

func (s *QueryAxnBindFixedLineRequest) SetOwnerId(v int64) *QueryAxnBindFixedLineRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryAxnBindFixedLineRequest) SetPhone(v string) *QueryAxnBindFixedLineRequest {
	s.Phone = &v
	return s
}

func (s *QueryAxnBindFixedLineRequest) SetQueryType(v string) *QueryAxnBindFixedLineRequest {
	s.QueryType = &v
	return s
}

func (s *QueryAxnBindFixedLineRequest) SetResourceOwnerAccount(v string) *QueryAxnBindFixedLineRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryAxnBindFixedLineRequest) SetResourceOwnerId(v int64) *QueryAxnBindFixedLineRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryAxnBindFixedLineRequest) SetSubId(v string) *QueryAxnBindFixedLineRequest {
	s.SubId = &v
	return s
}

func (s *QueryAxnBindFixedLineRequest) SetTelX(v string) *QueryAxnBindFixedLineRequest {
	s.TelX = &v
	return s
}

func (s *QueryAxnBindFixedLineRequest) SetTs(v string) *QueryAxnBindFixedLineRequest {
	s.Ts = &v
	return s
}

type QueryAxnBindFixedLineResponseBody struct {
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// 响应码
	//
	// example:
	//
	// 0
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 绑定对象
	Data []*QueryAxnBindFixedLineResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// 描述信息
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// E8B9C3ED-D9BD-5E27-9588-6D84D3070160
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryAxnBindFixedLineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAxnBindFixedLineResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAxnBindFixedLineResponseBody) SetAccessDeniedDetail(v string) *QueryAxnBindFixedLineResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *QueryAxnBindFixedLineResponseBody) SetCode(v string) *QueryAxnBindFixedLineResponseBody {
	s.Code = &v
	return s
}

func (s *QueryAxnBindFixedLineResponseBody) SetData(v []*QueryAxnBindFixedLineResponseBodyData) *QueryAxnBindFixedLineResponseBody {
	s.Data = v
	return s
}

func (s *QueryAxnBindFixedLineResponseBody) SetMessage(v string) *QueryAxnBindFixedLineResponseBody {
	s.Message = &v
	return s
}

func (s *QueryAxnBindFixedLineResponseBody) SetRequestId(v string) *QueryAxnBindFixedLineResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryAxnBindFixedLineResponseBody) SetSuccess(v bool) *QueryAxnBindFixedLineResponseBody {
	s.Success = &v
	return s
}

type QueryAxnBindFixedLineResponseBodyData struct {
	// 接通前放音啊编码
	//
	// example:
	//
	// 10001,10002,10003
	Anucode *string `json:"Anucode,omitempty" xml:"Anucode,omitempty"`
	// 被叫侧放音编码
	//
	// example:
	//
	// 10001,10002
	Anucodecalled *string `json:"Anucodecalled,omitempty" xml:"Anucodecalled,omitempty"`
	// 隐私号码区号
	//
	// example:
	//
	// 010
	Areacode *string `json:"Areacode,omitempty" xml:"Areacode,omitempty"`
	// 过期时间
	//
	// example:
	//
	// 60
	Expiration *string `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	// 额外字段
	Extra *QueryAxnBindFixedLineResponseBodyDataExtra `json:"Extra,omitempty" xml:"Extra,omitempty" type:"Struct"`
	// 接入商自有字段，最大100字符长度
	//
	// example:
	//
	// 12444
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// 绑定id
	//
	// example:
	//
	// GHX0534X202504221531579290029-2-1-aliaxn
	Subid *string `json:"Subid,omitempty" xml:"Subid,omitempty"`
	// 格式为yyyyMMddHHmmss。时间采用北京时间，24小时制。
	//
	// example:
	//
	// 20250421141723
	Subts *string `json:"Subts,omitempty" xml:"Subts,omitempty"`
	// 接通后主叫侧放音编码
	//
	// example:
	//
	// 10001,10002
	TAnucodeConnect *string `json:"TAnucodeConnect,omitempty" xml:"TAnucodeConnect,omitempty"`
	// A号码
	//
	// example:
	//
	// 15500001111
	TelA *string `json:"TelA,omitempty" xml:"TelA,omitempty"`
	// N号码
	//
	// example:
	//
	// 15500002222
	TelB *string `json:"TelB,omitempty" xml:"TelB,omitempty"`
	// 隐私号码
	//
	// example:
	//
	// 057112345678
	TelX *string `json:"TelX,omitempty" xml:"TelX,omitempty"`
}

func (s QueryAxnBindFixedLineResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryAxnBindFixedLineResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryAxnBindFixedLineResponseBodyData) SetAnucode(v string) *QueryAxnBindFixedLineResponseBodyData {
	s.Anucode = &v
	return s
}

func (s *QueryAxnBindFixedLineResponseBodyData) SetAnucodecalled(v string) *QueryAxnBindFixedLineResponseBodyData {
	s.Anucodecalled = &v
	return s
}

func (s *QueryAxnBindFixedLineResponseBodyData) SetAreacode(v string) *QueryAxnBindFixedLineResponseBodyData {
	s.Areacode = &v
	return s
}

func (s *QueryAxnBindFixedLineResponseBodyData) SetExpiration(v string) *QueryAxnBindFixedLineResponseBodyData {
	s.Expiration = &v
	return s
}

func (s *QueryAxnBindFixedLineResponseBodyData) SetExtra(v *QueryAxnBindFixedLineResponseBodyDataExtra) *QueryAxnBindFixedLineResponseBodyData {
	s.Extra = v
	return s
}

func (s *QueryAxnBindFixedLineResponseBodyData) SetRemark(v string) *QueryAxnBindFixedLineResponseBodyData {
	s.Remark = &v
	return s
}

func (s *QueryAxnBindFixedLineResponseBodyData) SetSubid(v string) *QueryAxnBindFixedLineResponseBodyData {
	s.Subid = &v
	return s
}

func (s *QueryAxnBindFixedLineResponseBodyData) SetSubts(v string) *QueryAxnBindFixedLineResponseBodyData {
	s.Subts = &v
	return s
}

func (s *QueryAxnBindFixedLineResponseBodyData) SetTAnucodeConnect(v string) *QueryAxnBindFixedLineResponseBodyData {
	s.TAnucodeConnect = &v
	return s
}

func (s *QueryAxnBindFixedLineResponseBodyData) SetTelA(v string) *QueryAxnBindFixedLineResponseBodyData {
	s.TelA = &v
	return s
}

func (s *QueryAxnBindFixedLineResponseBodyData) SetTelB(v string) *QueryAxnBindFixedLineResponseBodyData {
	s.TelB = &v
	return s
}

func (s *QueryAxnBindFixedLineResponseBodyData) SetTelX(v string) *QueryAxnBindFixedLineResponseBodyData {
	s.TelX = &v
	return s
}

type QueryAxnBindFixedLineResponseBodyDataExtra struct {
	// A通过X呼叫，即A主叫X，仅下列值有效。默认是0。 0：不能外呼 1：接续最近的N号码 2：回拨固定号码：telB
	//
	// example:
	//
	// 1
	Callback *string `json:"Callback,omitempty" xml:"Callback,omitempty"`
	// 录音控制，仅下列值有效。默认是0（不开通录音功能）。 0：不录音 1：录音
	//
	// example:
	//
	// 1
	Callrecording *string `json:"Callrecording,omitempty" xml:"Callrecording,omitempty"`
}

func (s QueryAxnBindFixedLineResponseBodyDataExtra) String() string {
	return tea.Prettify(s)
}

func (s QueryAxnBindFixedLineResponseBodyDataExtra) GoString() string {
	return s.String()
}

func (s *QueryAxnBindFixedLineResponseBodyDataExtra) SetCallback(v string) *QueryAxnBindFixedLineResponseBodyDataExtra {
	s.Callback = &v
	return s
}

func (s *QueryAxnBindFixedLineResponseBodyDataExtra) SetCallrecording(v string) *QueryAxnBindFixedLineResponseBodyDataExtra {
	s.Callrecording = &v
	return s
}

type QueryAxnBindFixedLineResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAxnBindFixedLineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryAxnBindFixedLineResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAxnBindFixedLineResponse) GoString() string {
	return s.String()
}

func (s *QueryAxnBindFixedLineResponse) SetHeaders(v map[string]*string) *QueryAxnBindFixedLineResponse {
	s.Headers = v
	return s
}

func (s *QueryAxnBindFixedLineResponse) SetStatusCode(v int32) *QueryAxnBindFixedLineResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAxnBindFixedLineResponse) SetBody(v *QueryAxnBindFixedLineResponseBody) *QueryAxnBindFixedLineResponse {
	s.Body = v
	return s
}

type QueryAxnExtensionBindFixedLineRequest struct {
	// 业务id标识，由阿里云分配给客户侧
	//
	// This parameter is required.
	//
	// example:
	//
	// alitest
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 消息请求标识
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345dkwkd99d
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 查询类型 0：根据绑定id查询 1：根据A号码查询
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	QueryType            *string `json:"QueryType,omitempty" xml:"QueryType,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 绑定id，queryType=0时，必传
	//
	// example:
	//
	// 可参考绑定响应
	SubId *string `json:"SubId,omitempty" xml:"SubId,omitempty"`
	// A号码，queryType=1时，必传
	//
	// example:
	//
	// 15500001111
	TelA *string `json:"TelA,omitempty" xml:"TelA,omitempty"`
	// 格式yyyyMMddHHmmssSSS，时间采用北京时间，24小时制，精确至毫秒
	//
	// This parameter is required.
	//
	// example:
	//
	// 20250421141723098
	Ts *string `json:"Ts,omitempty" xml:"Ts,omitempty"`
}

func (s QueryAxnExtensionBindFixedLineRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAxnExtensionBindFixedLineRequest) GoString() string {
	return s.String()
}

func (s *QueryAxnExtensionBindFixedLineRequest) SetAppId(v string) *QueryAxnExtensionBindFixedLineRequest {
	s.AppId = &v
	return s
}

func (s *QueryAxnExtensionBindFixedLineRequest) SetOrderId(v string) *QueryAxnExtensionBindFixedLineRequest {
	s.OrderId = &v
	return s
}

func (s *QueryAxnExtensionBindFixedLineRequest) SetOwnerId(v int64) *QueryAxnExtensionBindFixedLineRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryAxnExtensionBindFixedLineRequest) SetQueryType(v string) *QueryAxnExtensionBindFixedLineRequest {
	s.QueryType = &v
	return s
}

func (s *QueryAxnExtensionBindFixedLineRequest) SetResourceOwnerAccount(v string) *QueryAxnExtensionBindFixedLineRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryAxnExtensionBindFixedLineRequest) SetResourceOwnerId(v int64) *QueryAxnExtensionBindFixedLineRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryAxnExtensionBindFixedLineRequest) SetSubId(v string) *QueryAxnExtensionBindFixedLineRequest {
	s.SubId = &v
	return s
}

func (s *QueryAxnExtensionBindFixedLineRequest) SetTelA(v string) *QueryAxnExtensionBindFixedLineRequest {
	s.TelA = &v
	return s
}

func (s *QueryAxnExtensionBindFixedLineRequest) SetTs(v string) *QueryAxnExtensionBindFixedLineRequest {
	s.Ts = &v
	return s
}

type QueryAxnExtensionBindFixedLineResponseBody struct {
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// 响应码
	//
	// example:
	//
	// 0
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 查询绑定对象集合，具体对象字段见绑定请求
	Data []*QueryAxnExtensionBindFixedLineResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// 描述信息
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 74EFA0E8-CFCA-54D9-BFE5-904F9FA88DBD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryAxnExtensionBindFixedLineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAxnExtensionBindFixedLineResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAxnExtensionBindFixedLineResponseBody) SetAccessDeniedDetail(v string) *QueryAxnExtensionBindFixedLineResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *QueryAxnExtensionBindFixedLineResponseBody) SetCode(v string) *QueryAxnExtensionBindFixedLineResponseBody {
	s.Code = &v
	return s
}

func (s *QueryAxnExtensionBindFixedLineResponseBody) SetData(v []*QueryAxnExtensionBindFixedLineResponseBodyData) *QueryAxnExtensionBindFixedLineResponseBody {
	s.Data = v
	return s
}

func (s *QueryAxnExtensionBindFixedLineResponseBody) SetMessage(v string) *QueryAxnExtensionBindFixedLineResponseBody {
	s.Message = &v
	return s
}

func (s *QueryAxnExtensionBindFixedLineResponseBody) SetRequestId(v string) *QueryAxnExtensionBindFixedLineResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryAxnExtensionBindFixedLineResponseBody) SetSuccess(v bool) *QueryAxnExtensionBindFixedLineResponseBody {
	s.Success = &v
	return s
}

type QueryAxnExtensionBindFixedLineResponseBodyData struct {
	// 放音编码
	//
	// example:
	//
	// 10001,10002,10003
	Anucode *string `json:"Anucode,omitempty" xml:"Anucode,omitempty"`
	// 被叫侧放音
	//
	// example:
	//
	// 10001,10002
	Anucodecalled *string `json:"Anucodecalled,omitempty" xml:"Anucodecalled,omitempty"`
	// 隐私号码区号
	//
	// example:
	//
	// 10
	Areacode *string `json:"Areacode,omitempty" xml:"Areacode,omitempty"`
	// 过期时间
	//
	// example:
	//
	// 60
	Expiration *string `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	// 额外字段
	Extraaxx *QueryAxnExtensionBindFixedLineResponseBodyDataExtraaxx `json:"Extraaxx,omitempty" xml:"Extraaxx,omitempty" type:"Struct"`
	// 接入商自有字段
	//
	// example:
	//
	// 12444
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// 绑定ID
	//
	// example:
	//
	// 可参考绑定响应
	Subid *string `json:"Subid,omitempty" xml:"Subid,omitempty"`
	// 绑定时间
	//
	// example:
	//
	// 20250421141723
	Subts *string `json:"Subts,omitempty" xml:"Subts,omitempty"`
	// 接通后主叫侧放音编码
	//
	// example:
	//
	// 10001,10002
	TAnucodeConnect *string `json:"TAnucodeConnect,omitempty" xml:"TAnucodeConnect,omitempty"`
	// A号码
	//
	// example:
	//
	// 15500001111
	TelA *string `json:"TelA,omitempty" xml:"TelA,omitempty"`
	// 小号号码
	//
	// example:
	//
	// 19700002222
	TelX *string `json:"TelX,omitempty" xml:"TelX,omitempty"`
	// 分机号
	//
	// example:
	//
	// 1009
	TelXext *string `json:"TelXext,omitempty" xml:"TelXext,omitempty"`
}

func (s QueryAxnExtensionBindFixedLineResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryAxnExtensionBindFixedLineResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryAxnExtensionBindFixedLineResponseBodyData) SetAnucode(v string) *QueryAxnExtensionBindFixedLineResponseBodyData {
	s.Anucode = &v
	return s
}

func (s *QueryAxnExtensionBindFixedLineResponseBodyData) SetAnucodecalled(v string) *QueryAxnExtensionBindFixedLineResponseBodyData {
	s.Anucodecalled = &v
	return s
}

func (s *QueryAxnExtensionBindFixedLineResponseBodyData) SetAreacode(v string) *QueryAxnExtensionBindFixedLineResponseBodyData {
	s.Areacode = &v
	return s
}

func (s *QueryAxnExtensionBindFixedLineResponseBodyData) SetExpiration(v string) *QueryAxnExtensionBindFixedLineResponseBodyData {
	s.Expiration = &v
	return s
}

func (s *QueryAxnExtensionBindFixedLineResponseBodyData) SetExtraaxx(v *QueryAxnExtensionBindFixedLineResponseBodyDataExtraaxx) *QueryAxnExtensionBindFixedLineResponseBodyData {
	s.Extraaxx = v
	return s
}

func (s *QueryAxnExtensionBindFixedLineResponseBodyData) SetRemark(v string) *QueryAxnExtensionBindFixedLineResponseBodyData {
	s.Remark = &v
	return s
}

func (s *QueryAxnExtensionBindFixedLineResponseBodyData) SetSubid(v string) *QueryAxnExtensionBindFixedLineResponseBodyData {
	s.Subid = &v
	return s
}

func (s *QueryAxnExtensionBindFixedLineResponseBodyData) SetSubts(v string) *QueryAxnExtensionBindFixedLineResponseBodyData {
	s.Subts = &v
	return s
}

func (s *QueryAxnExtensionBindFixedLineResponseBodyData) SetTAnucodeConnect(v string) *QueryAxnExtensionBindFixedLineResponseBodyData {
	s.TAnucodeConnect = &v
	return s
}

func (s *QueryAxnExtensionBindFixedLineResponseBodyData) SetTelA(v string) *QueryAxnExtensionBindFixedLineResponseBodyData {
	s.TelA = &v
	return s
}

func (s *QueryAxnExtensionBindFixedLineResponseBodyData) SetTelX(v string) *QueryAxnExtensionBindFixedLineResponseBodyData {
	s.TelX = &v
	return s
}

func (s *QueryAxnExtensionBindFixedLineResponseBodyData) SetTelXext(v string) *QueryAxnExtensionBindFixedLineResponseBodyData {
	s.TelXext = &v
	return s
}

type QueryAxnExtensionBindFixedLineResponseBodyDataExtraaxx struct {
	// 回拨控制
	//
	// example:
	//
	// 示例值示例值
	Callback *string `json:"Callback,omitempty" xml:"Callback,omitempty"`
	// 录音控制
	//
	// example:
	//
	// 示例值示例值
	Callrecording *string `json:"Callrecording,omitempty" xml:"Callrecording,omitempty"`
}

func (s QueryAxnExtensionBindFixedLineResponseBodyDataExtraaxx) String() string {
	return tea.Prettify(s)
}

func (s QueryAxnExtensionBindFixedLineResponseBodyDataExtraaxx) GoString() string {
	return s.String()
}

func (s *QueryAxnExtensionBindFixedLineResponseBodyDataExtraaxx) SetCallback(v string) *QueryAxnExtensionBindFixedLineResponseBodyDataExtraaxx {
	s.Callback = &v
	return s
}

func (s *QueryAxnExtensionBindFixedLineResponseBodyDataExtraaxx) SetCallrecording(v string) *QueryAxnExtensionBindFixedLineResponseBodyDataExtraaxx {
	s.Callrecording = &v
	return s
}

type QueryAxnExtensionBindFixedLineResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAxnExtensionBindFixedLineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryAxnExtensionBindFixedLineResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAxnExtensionBindFixedLineResponse) GoString() string {
	return s.String()
}

func (s *QueryAxnExtensionBindFixedLineResponse) SetHeaders(v map[string]*string) *QueryAxnExtensionBindFixedLineResponse {
	s.Headers = v
	return s
}

func (s *QueryAxnExtensionBindFixedLineResponse) SetStatusCode(v int32) *QueryAxnExtensionBindFixedLineResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAxnExtensionBindFixedLineResponse) SetBody(v *QueryAxnExtensionBindFixedLineResponseBody) *QueryAxnExtensionBindFixedLineResponse {
	s.Body = v
	return s
}

type QueryPhoneNoAByTrackNoRequest struct {
	// The cabinet number.
	//
	// example:
	//
	// 25689****
	CabinetNo *string `json:"CabinetNo,omitempty" xml:"CabinetNo,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Phone number X returned by the API operation for creating a binding.
	//
	// example:
	//
	// 1710000****
	PhoneNoX             *string `json:"PhoneNoX,omitempty" xml:"PhoneNoX,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tracking number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 22573****
	TrackNo *string `json:"trackNo,omitempty" xml:"trackNo,omitempty"`
}

func (s QueryPhoneNoAByTrackNoRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryPhoneNoAByTrackNoRequest) GoString() string {
	return s.String()
}

func (s *QueryPhoneNoAByTrackNoRequest) SetCabinetNo(v string) *QueryPhoneNoAByTrackNoRequest {
	s.CabinetNo = &v
	return s
}

func (s *QueryPhoneNoAByTrackNoRequest) SetOwnerId(v int64) *QueryPhoneNoAByTrackNoRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryPhoneNoAByTrackNoRequest) SetPhoneNoX(v string) *QueryPhoneNoAByTrackNoRequest {
	s.PhoneNoX = &v
	return s
}

func (s *QueryPhoneNoAByTrackNoRequest) SetResourceOwnerAccount(v string) *QueryPhoneNoAByTrackNoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryPhoneNoAByTrackNoRequest) SetResourceOwnerId(v int64) *QueryPhoneNoAByTrackNoRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryPhoneNoAByTrackNoRequest) SetTrackNo(v string) *QueryPhoneNoAByTrackNoRequest {
	s.TrackNo = &v
	return s
}

type QueryPhoneNoAByTrackNoResponseBody struct {
	// The response code.
	//
	// 	- The value OK indicates that the request was successful.
	//
	// 	- Other status codes indicate that the request failed. For more information, see [Error codes](https://help.aliyun.com/document_detail/109196.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The information returned after the phone numbers were bound.
	Module []*QueryPhoneNoAByTrackNoResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 8906582E-6722
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryPhoneNoAByTrackNoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryPhoneNoAByTrackNoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPhoneNoAByTrackNoResponseBody) SetCode(v string) *QueryPhoneNoAByTrackNoResponseBody {
	s.Code = &v
	return s
}

func (s *QueryPhoneNoAByTrackNoResponseBody) SetMessage(v string) *QueryPhoneNoAByTrackNoResponseBody {
	s.Message = &v
	return s
}

func (s *QueryPhoneNoAByTrackNoResponseBody) SetModule(v []*QueryPhoneNoAByTrackNoResponseBodyModule) *QueryPhoneNoAByTrackNoResponseBody {
	s.Module = v
	return s
}

func (s *QueryPhoneNoAByTrackNoResponseBody) SetRequestId(v string) *QueryPhoneNoAByTrackNoResponseBody {
	s.RequestId = &v
	return s
}

type QueryPhoneNoAByTrackNoResponseBodyModule struct {
	// The extension of phone number X.
	//
	// example:
	//
	// 130
	Extension *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// Phone number A.
	//
	// example:
	//
	// 1310000****
	PhoneNoA *string `json:"PhoneNoA,omitempty" xml:"PhoneNoA,omitempty"`
	// The private number, that is, phone number X.
	//
	// example:
	//
	// 1710000****
	PhoneNoX *string `json:"PhoneNoX,omitempty" xml:"PhoneNoX,omitempty"`
}

func (s QueryPhoneNoAByTrackNoResponseBodyModule) String() string {
	return tea.Prettify(s)
}

func (s QueryPhoneNoAByTrackNoResponseBodyModule) GoString() string {
	return s.String()
}

func (s *QueryPhoneNoAByTrackNoResponseBodyModule) SetExtension(v string) *QueryPhoneNoAByTrackNoResponseBodyModule {
	s.Extension = &v
	return s
}

func (s *QueryPhoneNoAByTrackNoResponseBodyModule) SetPhoneNoA(v string) *QueryPhoneNoAByTrackNoResponseBodyModule {
	s.PhoneNoA = &v
	return s
}

func (s *QueryPhoneNoAByTrackNoResponseBodyModule) SetPhoneNoX(v string) *QueryPhoneNoAByTrackNoResponseBodyModule {
	s.PhoneNoX = &v
	return s
}

type QueryPhoneNoAByTrackNoResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryPhoneNoAByTrackNoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryPhoneNoAByTrackNoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryPhoneNoAByTrackNoResponse) GoString() string {
	return s.String()
}

func (s *QueryPhoneNoAByTrackNoResponse) SetHeaders(v map[string]*string) *QueryPhoneNoAByTrackNoResponse {
	s.Headers = v
	return s
}

func (s *QueryPhoneNoAByTrackNoResponse) SetStatusCode(v int32) *QueryPhoneNoAByTrackNoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryPhoneNoAByTrackNoResponse) SetBody(v *QueryPhoneNoAByTrackNoResponseBody) *QueryPhoneNoAByTrackNoResponse {
	s.Body = v
	return s
}

type QueryRecordFileDownloadUrlRequest struct {
	// The ID of the call record. Log on to the [Phone Number Protection console](https://dypls.console.aliyun.com/dypls.htm#/account) and view **Call Record ID*	- on the **Call Record Query*	- page.
	//
	// This parameter is required.
	//
	// example:
	//
	// abcedf1234
	CallId *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	// The call initiation time in the call record. Log on to the [Phone Number Protection console](https://dypls.console.aliyun.com/dypls.htm#/account). View **Call Initiated At*	- on the **Call Record Query*	- page, or view the call_time field in the Call Detail Record (CDR) receipt.
	//
	// example:
	//
	// 2019-03-05 12:00:00
	CallTime *string `json:"CallTime,omitempty" xml:"CallTime,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The key of the phone number pool. Log on to the [Phone Number Protection console](https://dypls.console.aliyun.com/dypls.htm#/account) and view the key of the phone number pool on the **Number Pool Management*	- page.
	//
	// >  This parameter is required when **ProductType*	- is left empty.
	//
	// example:
	//
	// FC123456
	PoolKey *string `json:"PoolKey,omitempty" xml:"PoolKey,omitempty"`
	// The product type. Valid values:
	//
	// 	- **AXB_170**.
	//
	// 	- **AXN_170**.
	//
	// 	- **AXN_95**.
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
}

func (s QueryRecordFileDownloadUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordFileDownloadUrlRequest) GoString() string {
	return s.String()
}

func (s *QueryRecordFileDownloadUrlRequest) SetCallId(v string) *QueryRecordFileDownloadUrlRequest {
	s.CallId = &v
	return s
}

func (s *QueryRecordFileDownloadUrlRequest) SetCallTime(v string) *QueryRecordFileDownloadUrlRequest {
	s.CallTime = &v
	return s
}

func (s *QueryRecordFileDownloadUrlRequest) SetOwnerId(v int64) *QueryRecordFileDownloadUrlRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryRecordFileDownloadUrlRequest) SetPoolKey(v string) *QueryRecordFileDownloadUrlRequest {
	s.PoolKey = &v
	return s
}

func (s *QueryRecordFileDownloadUrlRequest) SetProductType(v string) *QueryRecordFileDownloadUrlRequest {
	s.ProductType = &v
	return s
}

func (s *QueryRecordFileDownloadUrlRequest) SetResourceOwnerAccount(v string) *QueryRecordFileDownloadUrlRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryRecordFileDownloadUrlRequest) SetResourceOwnerId(v int64) *QueryRecordFileDownloadUrlRequest {
	s.ResourceOwnerId = &v
	return s
}

type QueryRecordFileDownloadUrlResponseBody struct {
	// The response code.
	//
	// 	- The value OK indicates that the request was successful.
	//
	// 	- Other status codes indicate that the request failed. For more information, see [Error codes](https://help.aliyun.com/document_detail/109196.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The download URL of the recording file. The download URL is valid for 2 hours.
	//
	// example:
	//
	// http://secret-axb-reco****cn-shanghai.aliyuncs.com/1000000820257625_66647243838006067251551752068865.mp3?Expires=155175****&OSSAccessKeyId=LTAIP00vvvv****v&Signature=tK6Yq9KusU4n%2BZ****7lg4/WmEA%3D
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1AB3CEF7-DCBE-488C-9C33-D180982CE031
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryRecordFileDownloadUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordFileDownloadUrlResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRecordFileDownloadUrlResponseBody) SetCode(v string) *QueryRecordFileDownloadUrlResponseBody {
	s.Code = &v
	return s
}

func (s *QueryRecordFileDownloadUrlResponseBody) SetDownloadUrl(v string) *QueryRecordFileDownloadUrlResponseBody {
	s.DownloadUrl = &v
	return s
}

func (s *QueryRecordFileDownloadUrlResponseBody) SetMessage(v string) *QueryRecordFileDownloadUrlResponseBody {
	s.Message = &v
	return s
}

func (s *QueryRecordFileDownloadUrlResponseBody) SetRequestId(v string) *QueryRecordFileDownloadUrlResponseBody {
	s.RequestId = &v
	return s
}

type QueryRecordFileDownloadUrlResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryRecordFileDownloadUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryRecordFileDownloadUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordFileDownloadUrlResponse) GoString() string {
	return s.String()
}

func (s *QueryRecordFileDownloadUrlResponse) SetHeaders(v map[string]*string) *QueryRecordFileDownloadUrlResponse {
	s.Headers = v
	return s
}

func (s *QueryRecordFileDownloadUrlResponse) SetStatusCode(v int32) *QueryRecordFileDownloadUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryRecordFileDownloadUrlResponse) SetBody(v *QueryRecordFileDownloadUrlResponseBody) *QueryRecordFileDownloadUrlResponse {
	s.Body = v
	return s
}

type QuerySecretAPhoneNoToCustRequest struct {
	// 号码组ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 58
	ANoWhiteGroupId *int64 `json:"ANoWhiteGroupId,omitempty" xml:"ANoWhiteGroupId,omitempty"`
	OwnerId         *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// A号码
	//
	// This parameter is required.
	//
	// example:
	//
	// 130*****8888
	PhoneNoA             *string `json:"PhoneNoA,omitempty" xml:"PhoneNoA,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QuerySecretAPhoneNoToCustRequest) String() string {
	return tea.Prettify(s)
}

func (s QuerySecretAPhoneNoToCustRequest) GoString() string {
	return s.String()
}

func (s *QuerySecretAPhoneNoToCustRequest) SetANoWhiteGroupId(v int64) *QuerySecretAPhoneNoToCustRequest {
	s.ANoWhiteGroupId = &v
	return s
}

func (s *QuerySecretAPhoneNoToCustRequest) SetOwnerId(v int64) *QuerySecretAPhoneNoToCustRequest {
	s.OwnerId = &v
	return s
}

func (s *QuerySecretAPhoneNoToCustRequest) SetPhoneNoA(v string) *QuerySecretAPhoneNoToCustRequest {
	s.PhoneNoA = &v
	return s
}

func (s *QuerySecretAPhoneNoToCustRequest) SetResourceOwnerAccount(v string) *QuerySecretAPhoneNoToCustRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QuerySecretAPhoneNoToCustRequest) SetResourceOwnerId(v int64) *QuerySecretAPhoneNoToCustRequest {
	s.ResourceOwnerId = &v
	return s
}

type QuerySecretAPhoneNoToCustResponseBody struct {
	// example:
	//
	// 0
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// 请求状态码
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// A号码报备状态查询结构体
	Data *QuerySecretAPhoneNoToCustResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// 失败错误提示
	//
	// example:
	//
	// 号码组不存在
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 返回id
	//
	// example:
	//
	// 1D73E648-0978-18A5-B089-3BB2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 请求是否成功
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QuerySecretAPhoneNoToCustResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QuerySecretAPhoneNoToCustResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySecretAPhoneNoToCustResponseBody) SetAccessDeniedDetail(v string) *QuerySecretAPhoneNoToCustResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *QuerySecretAPhoneNoToCustResponseBody) SetCode(v string) *QuerySecretAPhoneNoToCustResponseBody {
	s.Code = &v
	return s
}

func (s *QuerySecretAPhoneNoToCustResponseBody) SetData(v *QuerySecretAPhoneNoToCustResponseBodyData) *QuerySecretAPhoneNoToCustResponseBody {
	s.Data = v
	return s
}

func (s *QuerySecretAPhoneNoToCustResponseBody) SetMessage(v string) *QuerySecretAPhoneNoToCustResponseBody {
	s.Message = &v
	return s
}

func (s *QuerySecretAPhoneNoToCustResponseBody) SetRequestId(v string) *QuerySecretAPhoneNoToCustResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySecretAPhoneNoToCustResponseBody) SetSuccess(v bool) *QuerySecretAPhoneNoToCustResponseBody {
	s.Success = &v
	return s
}

type QuerySecretAPhoneNoToCustResponseBodyData struct {
	// 所属a号码组id
	//
	// example:
	//
	// 19
	ANoWhiteGroupId *string `json:"ANoWhiteGroupId,omitempty" xml:"ANoWhiteGroupId,omitempty"`
	// 固话报备的经办人/法人电话
	//
	// example:
	//
	// 130*****8888
	CustPhoneNo *string `json:"CustPhoneNo,omitempty" xml:"CustPhoneNo,omitempty"`
	// 号码类型
	//
	// example:
	//
	// Mobile
	NoType *string `json:"NoType,omitempty" xml:"NoType,omitempty"`
	// A号码
	//
	// example:
	//
	// 130*****1234
	PhoneNoA *string `json:"PhoneNoA,omitempty" xml:"PhoneNoA,omitempty"`
	// 备注（客户自己的业务备注，可编辑）
	//
	// example:
	//
	// ***报备
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// 报备失败原因
	//
	// example:
	//
	// ["系统判断为不同人"]
	ReportResult *string `json:"ReportResult,omitempty" xml:"ReportResult,omitempty"`
	// 报备结果
	//
	// example:
	//
	// REVIEWING
	ReportStatus *string `json:"ReportStatus,omitempty" xml:"ReportStatus,omitempty"`
}

func (s QuerySecretAPhoneNoToCustResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QuerySecretAPhoneNoToCustResponseBodyData) GoString() string {
	return s.String()
}

func (s *QuerySecretAPhoneNoToCustResponseBodyData) SetANoWhiteGroupId(v string) *QuerySecretAPhoneNoToCustResponseBodyData {
	s.ANoWhiteGroupId = &v
	return s
}

func (s *QuerySecretAPhoneNoToCustResponseBodyData) SetCustPhoneNo(v string) *QuerySecretAPhoneNoToCustResponseBodyData {
	s.CustPhoneNo = &v
	return s
}

func (s *QuerySecretAPhoneNoToCustResponseBodyData) SetNoType(v string) *QuerySecretAPhoneNoToCustResponseBodyData {
	s.NoType = &v
	return s
}

func (s *QuerySecretAPhoneNoToCustResponseBodyData) SetPhoneNoA(v string) *QuerySecretAPhoneNoToCustResponseBodyData {
	s.PhoneNoA = &v
	return s
}

func (s *QuerySecretAPhoneNoToCustResponseBodyData) SetRemark(v string) *QuerySecretAPhoneNoToCustResponseBodyData {
	s.Remark = &v
	return s
}

func (s *QuerySecretAPhoneNoToCustResponseBodyData) SetReportResult(v string) *QuerySecretAPhoneNoToCustResponseBodyData {
	s.ReportResult = &v
	return s
}

func (s *QuerySecretAPhoneNoToCustResponseBodyData) SetReportStatus(v string) *QuerySecretAPhoneNoToCustResponseBodyData {
	s.ReportStatus = &v
	return s
}

type QuerySecretAPhoneNoToCustResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySecretAPhoneNoToCustResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuerySecretAPhoneNoToCustResponse) String() string {
	return tea.Prettify(s)
}

func (s QuerySecretAPhoneNoToCustResponse) GoString() string {
	return s.String()
}

func (s *QuerySecretAPhoneNoToCustResponse) SetHeaders(v map[string]*string) *QuerySecretAPhoneNoToCustResponse {
	s.Headers = v
	return s
}

func (s *QuerySecretAPhoneNoToCustResponse) SetStatusCode(v int32) *QuerySecretAPhoneNoToCustResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySecretAPhoneNoToCustResponse) SetBody(v *QuerySecretAPhoneNoToCustResponseBody) *QuerySecretAPhoneNoToCustResponse {
	s.Body = v
	return s
}

type QuerySecretNoDetailRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The key of the phone number pool.
	//
	// Log on to the [Phone Number Protection console](https://dypls.console.aliyun.com/dypls.htm#/account) and view the key of the phone number pool on the **Number Pool Management*	- page.
	//
	// This parameter is required.
	//
	// example:
	//
	// FC2258****
	PoolKey              *string `json:"PoolKey,omitempty" xml:"PoolKey,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The private number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1390000****
	SecretNo *string `json:"SecretNo,omitempty" xml:"SecretNo,omitempty"`
}

func (s QuerySecretNoDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s QuerySecretNoDetailRequest) GoString() string {
	return s.String()
}

func (s *QuerySecretNoDetailRequest) SetOwnerId(v int64) *QuerySecretNoDetailRequest {
	s.OwnerId = &v
	return s
}

func (s *QuerySecretNoDetailRequest) SetPoolKey(v string) *QuerySecretNoDetailRequest {
	s.PoolKey = &v
	return s
}

func (s *QuerySecretNoDetailRequest) SetResourceOwnerAccount(v string) *QuerySecretNoDetailRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QuerySecretNoDetailRequest) SetResourceOwnerId(v int64) *QuerySecretNoDetailRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QuerySecretNoDetailRequest) SetSecretNo(v string) *QuerySecretNoDetailRequest {
	s.SecretNo = &v
	return s
}

type QuerySecretNoDetailResponseBody struct {
	// The response code.
	//
	// 	- The value OK indicates that the request was successful.
	//
	// 	- Other status codes indicate that the request failed. For more information, see [Error codes](https://help.aliyun.com/document_detail/109196.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 066E6E47-04CB-4774-A976-4F73CB76D4A3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The attributes of the phone number.
	SecretNoInfoDTO *QuerySecretNoDetailResponseBodySecretNoInfoDTO `json:"SecretNoInfoDTO,omitempty" xml:"SecretNoInfoDTO,omitempty" type:"Struct"`
}

func (s QuerySecretNoDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QuerySecretNoDetailResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySecretNoDetailResponseBody) SetCode(v string) *QuerySecretNoDetailResponseBody {
	s.Code = &v
	return s
}

func (s *QuerySecretNoDetailResponseBody) SetMessage(v string) *QuerySecretNoDetailResponseBody {
	s.Message = &v
	return s
}

func (s *QuerySecretNoDetailResponseBody) SetRequestId(v string) *QuerySecretNoDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySecretNoDetailResponseBody) SetSecretNoInfoDTO(v *QuerySecretNoDetailResponseBodySecretNoInfoDTO) *QuerySecretNoDetailResponseBody {
	s.SecretNoInfoDTO = v
	return s
}

type QuerySecretNoDetailResponseBodySecretNoInfoDTO struct {
	// The verification status of the phone number. Valid values:
	//
	// 	- **0**: The phone number is not verified.
	//
	// 	- **1**: The phone number is verified.
	//
	// example:
	//
	// 0
	CertifyStatus *int32 `json:"CertifyStatus,omitempty" xml:"CertifyStatus,omitempty"`
	// The city.
	//
	// example:
	//
	// chengdu
	City *string `json:"City,omitempty" xml:"City,omitempty"`
	// The province.
	//
	// example:
	//
	// sichuan
	Province *string `json:"Province,omitempty" xml:"Province,omitempty"`
	// The time when the phone number was purchased.
	//
	// example:
	//
	// 2021-12-03 15:19:27
	PurchaseTime *string `json:"PurchaseTime,omitempty" xml:"PurchaseTime,omitempty"`
	// The status of the phone number. Valid values:
	//
	// 	- **0**: The phone number is not bound to other phone numbers.
	//
	// 	- **1**: The phone number is bound to other phone numbers.
	//
	// 	- **2**: The phone number is locked.
	//
	// 	- **3**: The phone number is frozen.
	//
	// example:
	//
	// 0
	SecretStatus *int64 `json:"SecretStatus,omitempty" xml:"SecretStatus,omitempty"`
	// The carrier to which the phone number belongs. Valid values:
	//
	// 	- **1**: China Mobile
	//
	// 	- **2**: China Unicom
	//
	// 	- **3**: China Telecom
	//
	// example:
	//
	// 1
	Vendor *int64 `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
}

func (s QuerySecretNoDetailResponseBodySecretNoInfoDTO) String() string {
	return tea.Prettify(s)
}

func (s QuerySecretNoDetailResponseBodySecretNoInfoDTO) GoString() string {
	return s.String()
}

func (s *QuerySecretNoDetailResponseBodySecretNoInfoDTO) SetCertifyStatus(v int32) *QuerySecretNoDetailResponseBodySecretNoInfoDTO {
	s.CertifyStatus = &v
	return s
}

func (s *QuerySecretNoDetailResponseBodySecretNoInfoDTO) SetCity(v string) *QuerySecretNoDetailResponseBodySecretNoInfoDTO {
	s.City = &v
	return s
}

func (s *QuerySecretNoDetailResponseBodySecretNoInfoDTO) SetProvince(v string) *QuerySecretNoDetailResponseBodySecretNoInfoDTO {
	s.Province = &v
	return s
}

func (s *QuerySecretNoDetailResponseBodySecretNoInfoDTO) SetPurchaseTime(v string) *QuerySecretNoDetailResponseBodySecretNoInfoDTO {
	s.PurchaseTime = &v
	return s
}

func (s *QuerySecretNoDetailResponseBodySecretNoInfoDTO) SetSecretStatus(v int64) *QuerySecretNoDetailResponseBodySecretNoInfoDTO {
	s.SecretStatus = &v
	return s
}

func (s *QuerySecretNoDetailResponseBodySecretNoInfoDTO) SetVendor(v int64) *QuerySecretNoDetailResponseBodySecretNoInfoDTO {
	s.Vendor = &v
	return s
}

type QuerySecretNoDetailResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySecretNoDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuerySecretNoDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s QuerySecretNoDetailResponse) GoString() string {
	return s.String()
}

func (s *QuerySecretNoDetailResponse) SetHeaders(v map[string]*string) *QuerySecretNoDetailResponse {
	s.Headers = v
	return s
}

func (s *QuerySecretNoDetailResponse) SetStatusCode(v int32) *QuerySecretNoDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySecretNoDetailResponse) SetBody(v *QuerySecretNoDetailResponseBody) *QuerySecretNoDetailResponse {
	s.Body = v
	return s
}

type QuerySecretNoRemainRequest struct {
	// The home location of the phone number.
	//
	// 	- If **SpecId*	- is set to 1 or 2, you can specify the **City*	- parameter to query the quantity of available phone numbers.
	//
	// 1.  You can enter a single city name to perform a query.
	//
	// 2.  You can enter National to query the quantity of remaining phone numbers available in the Chinese mainland for online purchase.
	//
	// 3.  You can enter National List to query the cities with available phone numbers and the quantities of remaining phone numbers in the Chinese mainland. Cities without available phone numbers will not be returned.
	//
	// 	- If **SpecId*	- is set to 3, home locations are not distinguished for phone numbers that start with 95 and only the quantity of all the remaining phone numbers that start with 95 and are available for online purchase can be queried. If SpecId is set to 3, **City*	- must be set to **Nationwide**.
	//
	// >  Home locations can be set to only locations in the Chinese mainland.
	//
	// This parameter is required.
	//
	// example:
	//
	// hangzhou
	City                 *string `json:"City,omitempty" xml:"City,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The prefix of the phone number. When you call the QuerySecretNoRemain operation with **SecretNo*	- specified, the quantity of remaining phone numbers with the specified prefix that are available for online purchase is queried.
	//
	// Up to 18 digits of a phone number prefix can be specified.
	//
	// example:
	//
	// 130
	SecretNo *string `json:"SecretNo,omitempty" xml:"SecretNo,omitempty"`
	// The type of the phone number. Valid values:
	//
	// 	- **1**: a phone number assigned by a virtual network operator, that is, a phone number that belongs to the 170 or 171 number segment.
	//
	// 	- **2**: a phone number provided by a carrier.
	//
	// 	- **3**: a phone number that starts with 95.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	SpecId *int64 `json:"SpecId,omitempty" xml:"SpecId,omitempty"`
}

func (s QuerySecretNoRemainRequest) String() string {
	return tea.Prettify(s)
}

func (s QuerySecretNoRemainRequest) GoString() string {
	return s.String()
}

func (s *QuerySecretNoRemainRequest) SetCity(v string) *QuerySecretNoRemainRequest {
	s.City = &v
	return s
}

func (s *QuerySecretNoRemainRequest) SetOwnerId(v int64) *QuerySecretNoRemainRequest {
	s.OwnerId = &v
	return s
}

func (s *QuerySecretNoRemainRequest) SetResourceOwnerAccount(v string) *QuerySecretNoRemainRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QuerySecretNoRemainRequest) SetResourceOwnerId(v int64) *QuerySecretNoRemainRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QuerySecretNoRemainRequest) SetSecretNo(v string) *QuerySecretNoRemainRequest {
	s.SecretNo = &v
	return s
}

func (s *QuerySecretNoRemainRequest) SetSpecId(v int64) *QuerySecretNoRemainRequest {
	s.SpecId = &v
	return s
}

type QuerySecretNoRemainResponseBody struct {
	// The response code.
	//
	// 	- The value OK indicates that the request was successful.
	//
	// 	- Other status codes indicate that the request failed. For more information, see [Error codes](https://help.aliyun.com/document_detail/109196.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9FC30594-3841-43AD-9008-03393BCB5CD2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information returned after the operation was called.
	SecretRemainDTO *QuerySecretNoRemainResponseBodySecretRemainDTO `json:"SecretRemainDTO,omitempty" xml:"SecretRemainDTO,omitempty" type:"Struct"`
}

func (s QuerySecretNoRemainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QuerySecretNoRemainResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySecretNoRemainResponseBody) SetCode(v string) *QuerySecretNoRemainResponseBody {
	s.Code = &v
	return s
}

func (s *QuerySecretNoRemainResponseBody) SetMessage(v string) *QuerySecretNoRemainResponseBody {
	s.Message = &v
	return s
}

func (s *QuerySecretNoRemainResponseBody) SetRequestId(v string) *QuerySecretNoRemainResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySecretNoRemainResponseBody) SetSecretRemainDTO(v *QuerySecretNoRemainResponseBodySecretRemainDTO) *QuerySecretNoRemainResponseBody {
	s.SecretRemainDTO = v
	return s
}

type QuerySecretNoRemainResponseBodySecretRemainDTO struct {
	// The quantity of remaining phone numbers available for online purchase.
	//
	// example:
	//
	// 0
	Amount *int64 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// The home location of the phone numbers.
	//
	// example:
	//
	// hangzhou
	City *string `json:"City,omitempty" xml:"City,omitempty"`
	// The information about remaining phone numbers available for online purchase.
	RemainDTOList *QuerySecretNoRemainResponseBodySecretRemainDTORemainDTOList `json:"RemainDTOList,omitempty" xml:"RemainDTOList,omitempty" type:"Struct"`
}

func (s QuerySecretNoRemainResponseBodySecretRemainDTO) String() string {
	return tea.Prettify(s)
}

func (s QuerySecretNoRemainResponseBodySecretRemainDTO) GoString() string {
	return s.String()
}

func (s *QuerySecretNoRemainResponseBodySecretRemainDTO) SetAmount(v int64) *QuerySecretNoRemainResponseBodySecretRemainDTO {
	s.Amount = &v
	return s
}

func (s *QuerySecretNoRemainResponseBodySecretRemainDTO) SetCity(v string) *QuerySecretNoRemainResponseBodySecretRemainDTO {
	s.City = &v
	return s
}

func (s *QuerySecretNoRemainResponseBodySecretRemainDTO) SetRemainDTOList(v *QuerySecretNoRemainResponseBodySecretRemainDTORemainDTOList) *QuerySecretNoRemainResponseBodySecretRemainDTO {
	s.RemainDTOList = v
	return s
}

type QuerySecretNoRemainResponseBodySecretRemainDTORemainDTOList struct {
	RemainDTO []*QuerySecretNoRemainResponseBodySecretRemainDTORemainDTOListRemainDTO `json:"remainDTO,omitempty" xml:"remainDTO,omitempty" type:"Repeated"`
}

func (s QuerySecretNoRemainResponseBodySecretRemainDTORemainDTOList) String() string {
	return tea.Prettify(s)
}

func (s QuerySecretNoRemainResponseBodySecretRemainDTORemainDTOList) GoString() string {
	return s.String()
}

func (s *QuerySecretNoRemainResponseBodySecretRemainDTORemainDTOList) SetRemainDTO(v []*QuerySecretNoRemainResponseBodySecretRemainDTORemainDTOListRemainDTO) *QuerySecretNoRemainResponseBodySecretRemainDTORemainDTOList {
	s.RemainDTO = v
	return s
}

type QuerySecretNoRemainResponseBodySecretRemainDTORemainDTOListRemainDTO struct {
	// The quantity of remaining phone numbers available for online purchase for the city.
	//
	// example:
	//
	// 120
	Amount *int64 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// The home location of the phone numbers.
	//
	// example:
	//
	// Wuhan
	City *string `json:"City,omitempty" xml:"City,omitempty"`
}

func (s QuerySecretNoRemainResponseBodySecretRemainDTORemainDTOListRemainDTO) String() string {
	return tea.Prettify(s)
}

func (s QuerySecretNoRemainResponseBodySecretRemainDTORemainDTOListRemainDTO) GoString() string {
	return s.String()
}

func (s *QuerySecretNoRemainResponseBodySecretRemainDTORemainDTOListRemainDTO) SetAmount(v int64) *QuerySecretNoRemainResponseBodySecretRemainDTORemainDTOListRemainDTO {
	s.Amount = &v
	return s
}

func (s *QuerySecretNoRemainResponseBodySecretRemainDTORemainDTOListRemainDTO) SetCity(v string) *QuerySecretNoRemainResponseBodySecretRemainDTORemainDTOListRemainDTO {
	s.City = &v
	return s
}

type QuerySecretNoRemainResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySecretNoRemainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuerySecretNoRemainResponse) String() string {
	return tea.Prettify(s)
}

func (s QuerySecretNoRemainResponse) GoString() string {
	return s.String()
}

func (s *QuerySecretNoRemainResponse) SetHeaders(v map[string]*string) *QuerySecretNoRemainResponse {
	s.Headers = v
	return s
}

func (s *QuerySecretNoRemainResponse) SetStatusCode(v int32) *QuerySecretNoRemainResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySecretNoRemainResponse) SetBody(v *QuerySecretNoRemainResponseBody) *QuerySecretNoRemainResponse {
	s.Body = v
	return s
}

type QuerySoundRecordRequest struct {
	// 本次呼叫唯一id
	//
	// This parameter is required.
	//
	// example:
	//
	// ac445343254
	CallId *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	// 客户uid
	//
	// example:
	//
	// -
	CallerParentId *int64 `json:"CallerParentId,omitempty" xml:"CallerParentId,omitempty"`
	// 号码池key
	//
	// This parameter is required.
	//
	// example:
	//
	// FC5**********************a1a
	CustomerPoolKey *string `json:"CustomerPoolKey,omitempty" xml:"CustomerPoolKey,omitempty"`
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 请求去重ID, reqId最大长度为20位,接入方需要保持原子性
	//
	// This parameter is required.
	//
	// example:
	//
	// 564**********879
	ReqId                *string `json:"ReqId,omitempty" xml:"ReqId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QuerySoundRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s QuerySoundRecordRequest) GoString() string {
	return s.String()
}

func (s *QuerySoundRecordRequest) SetCallId(v string) *QuerySoundRecordRequest {
	s.CallId = &v
	return s
}

func (s *QuerySoundRecordRequest) SetCallerParentId(v int64) *QuerySoundRecordRequest {
	s.CallerParentId = &v
	return s
}

func (s *QuerySoundRecordRequest) SetCustomerPoolKey(v string) *QuerySoundRecordRequest {
	s.CustomerPoolKey = &v
	return s
}

func (s *QuerySoundRecordRequest) SetOwnerId(v int64) *QuerySoundRecordRequest {
	s.OwnerId = &v
	return s
}

func (s *QuerySoundRecordRequest) SetReqId(v string) *QuerySoundRecordRequest {
	s.ReqId = &v
	return s
}

func (s *QuerySoundRecordRequest) SetResourceOwnerAccount(v string) *QuerySoundRecordRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QuerySoundRecordRequest) SetResourceOwnerId(v int64) *QuerySoundRecordRequest {
	s.ResourceOwnerId = &v
	return s
}

type QuerySoundRecordResponseBody struct {
	// example:
	//
	// 0
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// 返回状态码 0000表示成功 其他表示失败
	//
	// example:
	//
	// 0000
	Code *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *QuerySoundRecordResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// 返回信息
	//
	// example:
	//
	// 成功
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 返回是否成功 true  表示成功 false表示失败
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QuerySoundRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QuerySoundRecordResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySoundRecordResponseBody) SetAccessDeniedDetail(v string) *QuerySoundRecordResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *QuerySoundRecordResponseBody) SetCode(v string) *QuerySoundRecordResponseBody {
	s.Code = &v
	return s
}

func (s *QuerySoundRecordResponseBody) SetData(v *QuerySoundRecordResponseBodyData) *QuerySoundRecordResponseBody {
	s.Data = v
	return s
}

func (s *QuerySoundRecordResponseBody) SetMessage(v string) *QuerySoundRecordResponseBody {
	s.Message = &v
	return s
}

func (s *QuerySoundRecordResponseBody) SetSuccess(v bool) *QuerySoundRecordResponseBody {
	s.Success = &v
	return s
}

type QuerySoundRecordResponseBodyData struct {
	// 通话录音url路径，最大长度1000，有效期1小时
	//
	// example:
	//
	// http://www.oss.com/temepl/a.mp3
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
}

func (s QuerySoundRecordResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QuerySoundRecordResponseBodyData) GoString() string {
	return s.String()
}

func (s *QuerySoundRecordResponseBodyData) SetFileUrl(v string) *QuerySoundRecordResponseBodyData {
	s.FileUrl = &v
	return s
}

type QuerySoundRecordResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySoundRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuerySoundRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s QuerySoundRecordResponse) GoString() string {
	return s.String()
}

func (s *QuerySoundRecordResponse) SetHeaders(v map[string]*string) *QuerySoundRecordResponse {
	s.Headers = v
	return s
}

func (s *QuerySoundRecordResponse) SetStatusCode(v int32) *QuerySoundRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySoundRecordResponse) SetBody(v *QuerySoundRecordResponseBody) *QuerySoundRecordResponse {
	s.Body = v
	return s
}

type QuerySubsIdRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The private number in the binding, that is, phone number X.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1390000****
	PhoneNoX *string `json:"PhoneNoX,omitempty" xml:"PhoneNoX,omitempty"`
	// The key of the phone number pool.
	//
	// Log on to the [Phone Number Protection console](https://dyplsnext.console.aliyun.com/overview) and view the key of the phone number pool on the Number Pool Management page.
	//
	// example:
	//
	// FC123456
	PoolKey              *string `json:"PoolKey,omitempty" xml:"PoolKey,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QuerySubsIdRequest) String() string {
	return tea.Prettify(s)
}

func (s QuerySubsIdRequest) GoString() string {
	return s.String()
}

func (s *QuerySubsIdRequest) SetOwnerId(v int64) *QuerySubsIdRequest {
	s.OwnerId = &v
	return s
}

func (s *QuerySubsIdRequest) SetPhoneNoX(v string) *QuerySubsIdRequest {
	s.PhoneNoX = &v
	return s
}

func (s *QuerySubsIdRequest) SetPoolKey(v string) *QuerySubsIdRequest {
	s.PoolKey = &v
	return s
}

func (s *QuerySubsIdRequest) SetResourceOwnerAccount(v string) *QuerySubsIdRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QuerySubsIdRequest) SetResourceOwnerId(v int64) *QuerySubsIdRequest {
	s.ResourceOwnerId = &v
	return s
}

type QuerySubsIdResponseBody struct {
	// The response code. The value OK indicates that the request was successful.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E7F99446-8191-43C0-99B5-F58A6AEAD779
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The binding ID.
	//
	// example:
	//
	// 11111111****
	SubsId *string `json:"SubsId,omitempty" xml:"SubsId,omitempty"`
}

func (s QuerySubsIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QuerySubsIdResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySubsIdResponseBody) SetCode(v string) *QuerySubsIdResponseBody {
	s.Code = &v
	return s
}

func (s *QuerySubsIdResponseBody) SetMessage(v string) *QuerySubsIdResponseBody {
	s.Message = &v
	return s
}

func (s *QuerySubsIdResponseBody) SetRequestId(v string) *QuerySubsIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySubsIdResponseBody) SetSubsId(v string) *QuerySubsIdResponseBody {
	s.SubsId = &v
	return s
}

type QuerySubsIdResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySubsIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuerySubsIdResponse) String() string {
	return tea.Prettify(s)
}

func (s QuerySubsIdResponse) GoString() string {
	return s.String()
}

func (s *QuerySubsIdResponse) SetHeaders(v map[string]*string) *QuerySubsIdResponse {
	s.Headers = v
	return s
}

func (s *QuerySubsIdResponse) SetStatusCode(v int32) *QuerySubsIdResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySubsIdResponse) SetBody(v *QuerySubsIdResponseBody) *QuerySubsIdResponse {
	s.Body = v
	return s
}

type QuerySubscriptionDetailRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The private number in the binding, that is, phone number X.
	//
	// This parameter is required.
	//
	// example:
	//
	// 13900001234
	PhoneNoX *string `json:"PhoneNoX,omitempty" xml:"PhoneNoX,omitempty"`
	// The key of the phone number pool. Log on to the [Phone Number Protection console](https://dypls.console.aliyun.com/dypls.htm#/account) and view the key of the phone number pool on the **Number Pool Management*	- page.
	//
	// >  This parameter is required when **ProductType*	- is left empty.
	//
	// example:
	//
	// FC123456
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
	// The binding ID.
	//
	// Log on to the Phone Number Protection console, choose **Number and Number Pool*	- > **Number Management**. On the Number Management page, select the desired record and click Details to view the binding ID. Alternatively, you can view the value of the **SubsId*	- parameter returned by an API operation for a phone number binding such as [BindAxb](https://help.aliyun.com/document_detail/110248.html). The value of this parameter indicates a binding ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100000076879****
	SubsId *string `json:"SubsId,omitempty" xml:"SubsId,omitempty"`
}

func (s QuerySubscriptionDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s QuerySubscriptionDetailRequest) GoString() string {
	return s.String()
}

func (s *QuerySubscriptionDetailRequest) SetOwnerId(v int64) *QuerySubscriptionDetailRequest {
	s.OwnerId = &v
	return s
}

func (s *QuerySubscriptionDetailRequest) SetPhoneNoX(v string) *QuerySubscriptionDetailRequest {
	s.PhoneNoX = &v
	return s
}

func (s *QuerySubscriptionDetailRequest) SetPoolKey(v string) *QuerySubscriptionDetailRequest {
	s.PoolKey = &v
	return s
}

func (s *QuerySubscriptionDetailRequest) SetProductType(v string) *QuerySubscriptionDetailRequest {
	s.ProductType = &v
	return s
}

func (s *QuerySubscriptionDetailRequest) SetResourceOwnerAccount(v string) *QuerySubscriptionDetailRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QuerySubscriptionDetailRequest) SetResourceOwnerId(v int64) *QuerySubscriptionDetailRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QuerySubscriptionDetailRequest) SetSubsId(v string) *QuerySubscriptionDetailRequest {
	s.SubsId = &v
	return s
}

type QuerySubscriptionDetailResponseBody struct {
	// The response code.
	//
	// 	- The value OK indicates that the request was successful.
	//
	// 	- Other status codes indicate that the request failed. For more information, see [Error codes](https://help.aliyun.com/document_detail/109196.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 066E6E47-04CB-4774-A976-4F73CB76D4A3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information returned after the QuerySubscriptionDetail operation was called.
	SecretBindDetailDTO *QuerySubscriptionDetailResponseBodySecretBindDetailDTO `json:"SecretBindDetailDTO,omitempty" xml:"SecretBindDetailDTO,omitempty" type:"Struct"`
}

func (s QuerySubscriptionDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QuerySubscriptionDetailResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySubscriptionDetailResponseBody) SetCode(v string) *QuerySubscriptionDetailResponseBody {
	s.Code = &v
	return s
}

func (s *QuerySubscriptionDetailResponseBody) SetMessage(v string) *QuerySubscriptionDetailResponseBody {
	s.Message = &v
	return s
}

func (s *QuerySubscriptionDetailResponseBody) SetRequestId(v string) *QuerySubscriptionDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySubscriptionDetailResponseBody) SetSecretBindDetailDTO(v *QuerySubscriptionDetailResponseBodySecretBindDetailDTO) *QuerySubscriptionDetailResponseBody {
	s.SecretBindDetailDTO = v
	return s
}

type QuerySubscriptionDetailResponseBodySecretBindDetailDTO struct {
	// The ID of the ASR model.
	//
	// example:
	//
	// 123456
	ASRModelId *string `json:"ASRModelId,omitempty" xml:"ASRModelId,omitempty"`
	// Indicates whether automatic speech recognition (ASR) is enabled. Valid values:
	//
	// 	- **false**: ASR is disabled.
	//
	// 	- **true**: ASR is enabled.
	//
	// example:
	//
	// true
	ASRStatus *bool `json:"ASRStatus,omitempty" xml:"ASRStatus,omitempty"`
	// The status of one-way call restrictions. No value is returned for this parameter if one-way call restrictions are not set. Valid values:
	//
	// 	- **CONTROL_AX_DISABLE**: Phone number A cannot be used to call phone number X.
	//
	// 	- **CONTROL_BX_DISABLE**: Phone number B cannot be used to call phone number X.
	//
	// example:
	//
	// CONTROL_BX_DISABLE
	CallRestrict *string `json:"CallRestrict,omitempty" xml:"CallRestrict,omitempty"`
	// The expiration time of the binding.
	//
	// example:
	//
	// 2019-09-05 12:00:00
	ExpireDate *string `json:"ExpireDate,omitempty" xml:"ExpireDate,omitempty"`
	// The extension in the AXG extension binding.
	//
	// example:
	//
	// 130
	Extension *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// The creation time of the binding.
	//
	// example:
	//
	// 2019-03-05 12:00:00
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The ID of number group G in the binding.
	//
	// example:
	//
	// 2000000130001
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// Indicates whether all calls made by the bound phone numbers are recorded. Valid values:
	//
	// 	- **false**
	//
	// 	- **true**
	//
	// example:
	//
	// true
	NeedRecord *bool `json:"NeedRecord,omitempty" xml:"NeedRecord,omitempty"`
	// Phone number A in the binding.
	//
	// example:
	//
	// 13900001111
	PhoneNoA *string `json:"PhoneNoA,omitempty" xml:"PhoneNoA,omitempty"`
	// Phone number B in the binding.
	//
	// example:
	//
	// 13900002222
	PhoneNoB *string `json:"PhoneNoB,omitempty" xml:"PhoneNoB,omitempty"`
	// The private number in the binding, that is, phone number X.
	//
	// example:
	//
	// 13900001234
	PhoneNoX *string `json:"PhoneNoX,omitempty" xml:"PhoneNoX,omitempty"`
	// The binding status. Valid values:
	//
	// 	- **0**: The binding expired.
	//
	// 	- **1**: The binding is in effect.
	//
	// example:
	//
	// 1
	Status *int64 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The binding ID.
	//
	// example:
	//
	// 100000076879****
	SubsId *string `json:"SubsId,omitempty" xml:"SubsId,omitempty"`
}

func (s QuerySubscriptionDetailResponseBodySecretBindDetailDTO) String() string {
	return tea.Prettify(s)
}

func (s QuerySubscriptionDetailResponseBodySecretBindDetailDTO) GoString() string {
	return s.String()
}

func (s *QuerySubscriptionDetailResponseBodySecretBindDetailDTO) SetASRModelId(v string) *QuerySubscriptionDetailResponseBodySecretBindDetailDTO {
	s.ASRModelId = &v
	return s
}

func (s *QuerySubscriptionDetailResponseBodySecretBindDetailDTO) SetASRStatus(v bool) *QuerySubscriptionDetailResponseBodySecretBindDetailDTO {
	s.ASRStatus = &v
	return s
}

func (s *QuerySubscriptionDetailResponseBodySecretBindDetailDTO) SetCallRestrict(v string) *QuerySubscriptionDetailResponseBodySecretBindDetailDTO {
	s.CallRestrict = &v
	return s
}

func (s *QuerySubscriptionDetailResponseBodySecretBindDetailDTO) SetExpireDate(v string) *QuerySubscriptionDetailResponseBodySecretBindDetailDTO {
	s.ExpireDate = &v
	return s
}

func (s *QuerySubscriptionDetailResponseBodySecretBindDetailDTO) SetExtension(v string) *QuerySubscriptionDetailResponseBodySecretBindDetailDTO {
	s.Extension = &v
	return s
}

func (s *QuerySubscriptionDetailResponseBodySecretBindDetailDTO) SetGmtCreate(v string) *QuerySubscriptionDetailResponseBodySecretBindDetailDTO {
	s.GmtCreate = &v
	return s
}

func (s *QuerySubscriptionDetailResponseBodySecretBindDetailDTO) SetGroupId(v int64) *QuerySubscriptionDetailResponseBodySecretBindDetailDTO {
	s.GroupId = &v
	return s
}

func (s *QuerySubscriptionDetailResponseBodySecretBindDetailDTO) SetNeedRecord(v bool) *QuerySubscriptionDetailResponseBodySecretBindDetailDTO {
	s.NeedRecord = &v
	return s
}

func (s *QuerySubscriptionDetailResponseBodySecretBindDetailDTO) SetPhoneNoA(v string) *QuerySubscriptionDetailResponseBodySecretBindDetailDTO {
	s.PhoneNoA = &v
	return s
}

func (s *QuerySubscriptionDetailResponseBodySecretBindDetailDTO) SetPhoneNoB(v string) *QuerySubscriptionDetailResponseBodySecretBindDetailDTO {
	s.PhoneNoB = &v
	return s
}

func (s *QuerySubscriptionDetailResponseBodySecretBindDetailDTO) SetPhoneNoX(v string) *QuerySubscriptionDetailResponseBodySecretBindDetailDTO {
	s.PhoneNoX = &v
	return s
}

func (s *QuerySubscriptionDetailResponseBodySecretBindDetailDTO) SetStatus(v int64) *QuerySubscriptionDetailResponseBodySecretBindDetailDTO {
	s.Status = &v
	return s
}

func (s *QuerySubscriptionDetailResponseBodySecretBindDetailDTO) SetSubsId(v string) *QuerySubscriptionDetailResponseBodySecretBindDetailDTO {
	s.SubsId = &v
	return s
}

type QuerySubscriptionDetailResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySubscriptionDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuerySubscriptionDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s QuerySubscriptionDetailResponse) GoString() string {
	return s.String()
}

func (s *QuerySubscriptionDetailResponse) SetHeaders(v map[string]*string) *QuerySubscriptionDetailResponse {
	s.Headers = v
	return s
}

func (s *QuerySubscriptionDetailResponse) SetStatusCode(v int32) *QuerySubscriptionDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySubscriptionDetailResponse) SetBody(v *QuerySubscriptionDetailResponseBody) *QuerySubscriptionDetailResponse {
	s.Body = v
	return s
}

type ReleaseSecretNoRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The key of the phone number pool. Log on to the [Phone Number Protection console](https://dypls.console.aliyun.com/dypls.htm#/account) and view the key of the phone number pool on the **Number Pool Management*	- page.
	//
	// This parameter is required.
	//
	// example:
	//
	// FC123456
	PoolKey              *string `json:"PoolKey,omitempty" xml:"PoolKey,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The prefix of phone numbers. When you call the ReleaseSecretNo operation with **SecretNo*	- specified, the system performs fuzzy matching against phone numbers based on the prefix.
	//
	// >  Up to 18 digits of a phone number prefix can be specified.
	//
	// This parameter is required.
	//
	// example:
	//
	// 130
	SecretNo *string `json:"SecretNo,omitempty" xml:"SecretNo,omitempty"`
}

func (s ReleaseSecretNoRequest) String() string {
	return tea.Prettify(s)
}

func (s ReleaseSecretNoRequest) GoString() string {
	return s.String()
}

func (s *ReleaseSecretNoRequest) SetOwnerId(v int64) *ReleaseSecretNoRequest {
	s.OwnerId = &v
	return s
}

func (s *ReleaseSecretNoRequest) SetPoolKey(v string) *ReleaseSecretNoRequest {
	s.PoolKey = &v
	return s
}

func (s *ReleaseSecretNoRequest) SetResourceOwnerAccount(v string) *ReleaseSecretNoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ReleaseSecretNoRequest) SetResourceOwnerId(v int64) *ReleaseSecretNoRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ReleaseSecretNoRequest) SetSecretNo(v string) *ReleaseSecretNoRequest {
	s.SecretNo = &v
	return s
}

type ReleaseSecretNoResponseBody struct {
	// The response code.
	//
	// 	- The value OK indicates that the request was successful.
	//
	// 	- Other values indicate that the request failed. For more information, see [Error codes](https://help.aliyun.com/document_detail/109196.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 986BCB6D-C9BF-42F9-91CE-3A990121232
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReleaseSecretNoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReleaseSecretNoResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseSecretNoResponseBody) SetCode(v string) *ReleaseSecretNoResponseBody {
	s.Code = &v
	return s
}

func (s *ReleaseSecretNoResponseBody) SetMessage(v string) *ReleaseSecretNoResponseBody {
	s.Message = &v
	return s
}

func (s *ReleaseSecretNoResponseBody) SetRequestId(v string) *ReleaseSecretNoResponseBody {
	s.RequestId = &v
	return s
}

type ReleaseSecretNoResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReleaseSecretNoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReleaseSecretNoResponse) String() string {
	return tea.Prettify(s)
}

func (s ReleaseSecretNoResponse) GoString() string {
	return s.String()
}

func (s *ReleaseSecretNoResponse) SetHeaders(v map[string]*string) *ReleaseSecretNoResponse {
	s.Headers = v
	return s
}

func (s *ReleaseSecretNoResponse) SetStatusCode(v int32) *ReleaseSecretNoResponse {
	s.StatusCode = &v
	return s
}

func (s *ReleaseSecretNoResponse) SetBody(v *ReleaseSecretNoResponseBody) *ReleaseSecretNoResponse {
	s.Body = v
	return s
}

type UnBindAXBRequest struct {
	// bindId绑定关系AXB唯一id
	//
	// This parameter is required.
	//
	// example:
	//
	// 4534543
	BindId *string `json:"BindId,omitempty" xml:"BindId,omitempty"`
	// 客户uid
	//
	// example:
	//
	// -
	CallerParentId *int64 `json:"CallerParentId,omitempty" xml:"CallerParentId,omitempty"`
	// 号码池key
	//
	// This parameter is required.
	//
	// example:
	//
	// FC5**********************a1a
	CustomerPoolKey *string `json:"CustomerPoolKey,omitempty" xml:"CustomerPoolKey,omitempty"`
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 请求去重ID, reqId最大长度为20位,接入方需要保持原子性
	//
	// This parameter is required.
	//
	// example:
	//
	// 564**********879
	ReqId                *string `json:"ReqId,omitempty" xml:"ReqId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s UnBindAXBRequest) String() string {
	return tea.Prettify(s)
}

func (s UnBindAXBRequest) GoString() string {
	return s.String()
}

func (s *UnBindAXBRequest) SetBindId(v string) *UnBindAXBRequest {
	s.BindId = &v
	return s
}

func (s *UnBindAXBRequest) SetCallerParentId(v int64) *UnBindAXBRequest {
	s.CallerParentId = &v
	return s
}

func (s *UnBindAXBRequest) SetCustomerPoolKey(v string) *UnBindAXBRequest {
	s.CustomerPoolKey = &v
	return s
}

func (s *UnBindAXBRequest) SetOwnerId(v int64) *UnBindAXBRequest {
	s.OwnerId = &v
	return s
}

func (s *UnBindAXBRequest) SetReqId(v string) *UnBindAXBRequest {
	s.ReqId = &v
	return s
}

func (s *UnBindAXBRequest) SetResourceOwnerAccount(v string) *UnBindAXBRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UnBindAXBRequest) SetResourceOwnerId(v int64) *UnBindAXBRequest {
	s.ResourceOwnerId = &v
	return s
}

type UnBindAXBResponseBody struct {
	// example:
	//
	// 0
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 0000
	Code      *string                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *UnBindAXBResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UnBindAXBResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnBindAXBResponseBody) GoString() string {
	return s.String()
}

func (s *UnBindAXBResponseBody) SetAccessDeniedDetail(v string) *UnBindAXBResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UnBindAXBResponseBody) SetCode(v string) *UnBindAXBResponseBody {
	s.Code = &v
	return s
}

func (s *UnBindAXBResponseBody) SetData(v *UnBindAXBResponseBodyData) *UnBindAXBResponseBody {
	s.Data = v
	return s
}

func (s *UnBindAXBResponseBody) SetMessage(v string) *UnBindAXBResponseBody {
	s.Message = &v
	return s
}

func (s *UnBindAXBResponseBody) SetRequestId(v string) *UnBindAXBResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnBindAXBResponseBody) SetSuccess(v bool) *UnBindAXBResponseBody {
	s.Success = &v
	return s
}

type UnBindAXBResponseBodyData struct {
	// 返回状态码 0000表示成功 其他表示失败
	//
	// example:
	//
	// 0000
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 返回信息
	//
	// example:
	//
	// 成功
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 返回是否成功 true  表示成功 false表示失败
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UnBindAXBResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UnBindAXBResponseBodyData) GoString() string {
	return s.String()
}

func (s *UnBindAXBResponseBodyData) SetCode(v string) *UnBindAXBResponseBodyData {
	s.Code = &v
	return s
}

func (s *UnBindAXBResponseBodyData) SetMessage(v string) *UnBindAXBResponseBodyData {
	s.Message = &v
	return s
}

func (s *UnBindAXBResponseBodyData) SetSuccess(v bool) *UnBindAXBResponseBodyData {
	s.Success = &v
	return s
}

type UnBindAXBResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnBindAXBResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnBindAXBResponse) String() string {
	return tea.Prettify(s)
}

func (s UnBindAXBResponse) GoString() string {
	return s.String()
}

func (s *UnBindAXBResponse) SetHeaders(v map[string]*string) *UnBindAXBResponse {
	s.Headers = v
	return s
}

func (s *UnBindAXBResponse) SetStatusCode(v int32) *UnBindAXBResponse {
	s.StatusCode = &v
	return s
}

func (s *UnBindAXBResponse) SetBody(v *UnBindAXBResponseBody) *UnBindAXBResponse {
	s.Body = v
	return s
}

type UnBindXBRequest struct {
	// authId绑定关系BX唯一id
	//
	// This parameter is required.
	//
	// example:
	//
	// 34*****46
	AuthId *string `json:"AuthId,omitempty" xml:"AuthId,omitempty"`
	// 客户uid
	//
	// example:
	//
	// -
	CallerParentId *int64 `json:"CallerParentId,omitempty" xml:"CallerParentId,omitempty"`
	// 号码池key
	//
	// This parameter is required.
	//
	// example:
	//
	// FC5**********************a1a
	CustomerPoolKey *string `json:"CustomerPoolKey,omitempty" xml:"CustomerPoolKey,omitempty"`
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 请求去重ID, reqId最大长度为20位,接入方需要保持原子性
	//
	// This parameter is required.
	//
	// example:
	//
	// 564**********879
	ReqId                *string `json:"ReqId,omitempty" xml:"ReqId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// X号码
	//
	// This parameter is required.
	//
	// example:
	//
	// 17*******22
	TelX *string `json:"TelX,omitempty" xml:"TelX,omitempty"`
}

func (s UnBindXBRequest) String() string {
	return tea.Prettify(s)
}

func (s UnBindXBRequest) GoString() string {
	return s.String()
}

func (s *UnBindXBRequest) SetAuthId(v string) *UnBindXBRequest {
	s.AuthId = &v
	return s
}

func (s *UnBindXBRequest) SetCallerParentId(v int64) *UnBindXBRequest {
	s.CallerParentId = &v
	return s
}

func (s *UnBindXBRequest) SetCustomerPoolKey(v string) *UnBindXBRequest {
	s.CustomerPoolKey = &v
	return s
}

func (s *UnBindXBRequest) SetOwnerId(v int64) *UnBindXBRequest {
	s.OwnerId = &v
	return s
}

func (s *UnBindXBRequest) SetReqId(v string) *UnBindXBRequest {
	s.ReqId = &v
	return s
}

func (s *UnBindXBRequest) SetResourceOwnerAccount(v string) *UnBindXBRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UnBindXBRequest) SetResourceOwnerId(v int64) *UnBindXBRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UnBindXBRequest) SetTelX(v string) *UnBindXBRequest {
	s.TelX = &v
	return s
}

type UnBindXBResponseBody struct {
	// example:
	//
	// 0
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 0000
	Code      *string                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *UnBindXBResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UnBindXBResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnBindXBResponseBody) GoString() string {
	return s.String()
}

func (s *UnBindXBResponseBody) SetAccessDeniedDetail(v string) *UnBindXBResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UnBindXBResponseBody) SetCode(v string) *UnBindXBResponseBody {
	s.Code = &v
	return s
}

func (s *UnBindXBResponseBody) SetData(v *UnBindXBResponseBodyData) *UnBindXBResponseBody {
	s.Data = v
	return s
}

func (s *UnBindXBResponseBody) SetMessage(v string) *UnBindXBResponseBody {
	s.Message = &v
	return s
}

func (s *UnBindXBResponseBody) SetRequestId(v string) *UnBindXBResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnBindXBResponseBody) SetSuccess(v bool) *UnBindXBResponseBody {
	s.Success = &v
	return s
}

type UnBindXBResponseBodyData struct {
	// 返回状态码 0000表示成功 其他表示失败
	//
	// example:
	//
	// 0000
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 返回信息
	//
	// example:
	//
	// 成功
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 返回是否成功 true  表示成功 false表示失败
	//
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UnBindXBResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UnBindXBResponseBodyData) GoString() string {
	return s.String()
}

func (s *UnBindXBResponseBodyData) SetCode(v string) *UnBindXBResponseBodyData {
	s.Code = &v
	return s
}

func (s *UnBindXBResponseBodyData) SetMessage(v string) *UnBindXBResponseBodyData {
	s.Message = &v
	return s
}

func (s *UnBindXBResponseBodyData) SetSuccess(v bool) *UnBindXBResponseBodyData {
	s.Success = &v
	return s
}

type UnBindXBResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnBindXBResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnBindXBResponse) String() string {
	return tea.Prettify(s)
}

func (s UnBindXBResponse) GoString() string {
	return s.String()
}

func (s *UnBindXBResponse) SetHeaders(v map[string]*string) *UnBindXBResponse {
	s.Headers = v
	return s
}

func (s *UnBindXBResponse) SetStatusCode(v int32) *UnBindXBResponse {
	s.StatusCode = &v
	return s
}

func (s *UnBindXBResponse) SetBody(v *UnBindXBResponseBody) *UnBindXBResponse {
	s.Body = v
	return s
}

type UnbindSubscriptionRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The key of the phone number pool. Log on to the [Phone Number Protection console](https://dypls.console.aliyun.com/dypls.htm#/account) and view the key of the phone number pool on the **Number Pool Management*	- page.
	//
	// >  This parameter is required when **ProductType*	- is left empty.
	//
	// example:
	//
	// FC123456
	PoolKey *string `json:"PoolKey,omitempty" xml:"PoolKey,omitempty"`
	// The product type. Fixed value: **AXB_170**.
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
	// The private number, that is, phone number X specified in an API operation for a phone number binding such as [BindAXG](https://help.aliyun.com/document_detail/110249.html) or automatically assigned after such an operation is called.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1390000****
	SecretNo *string `json:"SecretNo,omitempty" xml:"SecretNo,omitempty"`
	// The binding ID.
	//
	// Log on to the Phone Number Protection console, choose **Number and Number Pool*	- > **Number Management**. On the Number Management page, select the desired record and click Details to view the binding ID. Alternatively, you can view the value of the **SubsId*	- parameter returned by an API operation for a phone number binding such as BindAxb. The value of this parameter indicates a binding ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1************2
	SubsId *string `json:"SubsId,omitempty" xml:"SubsId,omitempty"`
}

func (s UnbindSubscriptionRequest) String() string {
	return tea.Prettify(s)
}

func (s UnbindSubscriptionRequest) GoString() string {
	return s.String()
}

func (s *UnbindSubscriptionRequest) SetOwnerId(v int64) *UnbindSubscriptionRequest {
	s.OwnerId = &v
	return s
}

func (s *UnbindSubscriptionRequest) SetPoolKey(v string) *UnbindSubscriptionRequest {
	s.PoolKey = &v
	return s
}

func (s *UnbindSubscriptionRequest) SetProductType(v string) *UnbindSubscriptionRequest {
	s.ProductType = &v
	return s
}

func (s *UnbindSubscriptionRequest) SetResourceOwnerAccount(v string) *UnbindSubscriptionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UnbindSubscriptionRequest) SetResourceOwnerId(v int64) *UnbindSubscriptionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UnbindSubscriptionRequest) SetSecretNo(v string) *UnbindSubscriptionRequest {
	s.SecretNo = &v
	return s
}

func (s *UnbindSubscriptionRequest) SetSubsId(v string) *UnbindSubscriptionRequest {
	s.SubsId = &v
	return s
}

type UnbindSubscriptionResponseBody struct {
	// A deprecated parameter.
	//
	// example:
	//
	// true
	ChargeId *string `json:"ChargeId,omitempty" xml:"ChargeId,omitempty"`
	// The response code.
	//
	// 	- The value OK indicates that the request was successful.
	//
	// 	- Other values indicate that the request failed. For more information, see [Error codes](https://help.aliyun.com/document_detail/109196.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 986BCB6D-C9BF-42F9-91CE-3A9901233D36
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnbindSubscriptionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnbindSubscriptionResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindSubscriptionResponseBody) SetChargeId(v string) *UnbindSubscriptionResponseBody {
	s.ChargeId = &v
	return s
}

func (s *UnbindSubscriptionResponseBody) SetCode(v string) *UnbindSubscriptionResponseBody {
	s.Code = &v
	return s
}

func (s *UnbindSubscriptionResponseBody) SetMessage(v string) *UnbindSubscriptionResponseBody {
	s.Message = &v
	return s
}

func (s *UnbindSubscriptionResponseBody) SetRequestId(v string) *UnbindSubscriptionResponseBody {
	s.RequestId = &v
	return s
}

type UnbindSubscriptionResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnbindSubscriptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnbindSubscriptionResponse) String() string {
	return tea.Prettify(s)
}

func (s UnbindSubscriptionResponse) GoString() string {
	return s.String()
}

func (s *UnbindSubscriptionResponse) SetHeaders(v map[string]*string) *UnbindSubscriptionResponse {
	s.Headers = v
	return s
}

func (s *UnbindSubscriptionResponse) SetStatusCode(v int32) *UnbindSubscriptionResponse {
	s.StatusCode = &v
	return s
}

func (s *UnbindSubscriptionResponse) SetBody(v *UnbindSubscriptionResponseBody) *UnbindSubscriptionResponse {
	s.Body = v
	return s
}

type UnlockSecretNoRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The key of the phone number pool.
	//
	// Log on to the [Phone Number Protection console](https://dypls.console.aliyun.com/dypls.htm#/account) and view the key of the phone number pool on the **Number Pool Management*	- page.
	//
	// This parameter is required.
	//
	// example:
	//
	// FC2256****
	PoolKey              *string `json:"PoolKey,omitempty" xml:"PoolKey,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The private number that you want to unlock. You must enter a complete mobile phone number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1300000****
	SecretNo *string `json:"SecretNo,omitempty" xml:"SecretNo,omitempty"`
}

func (s UnlockSecretNoRequest) String() string {
	return tea.Prettify(s)
}

func (s UnlockSecretNoRequest) GoString() string {
	return s.String()
}

func (s *UnlockSecretNoRequest) SetOwnerId(v int64) *UnlockSecretNoRequest {
	s.OwnerId = &v
	return s
}

func (s *UnlockSecretNoRequest) SetPoolKey(v string) *UnlockSecretNoRequest {
	s.PoolKey = &v
	return s
}

func (s *UnlockSecretNoRequest) SetResourceOwnerAccount(v string) *UnlockSecretNoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UnlockSecretNoRequest) SetResourceOwnerId(v int64) *UnlockSecretNoRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UnlockSecretNoRequest) SetSecretNo(v string) *UnlockSecretNoRequest {
	s.SecretNo = &v
	return s
}

type UnlockSecretNoResponseBody struct {
	// The response code.
	//
	// 	- The value OK indicates that the request was successful.
	//
	// 	- Other status codes indicate that the request failed. For more information, see [Error codes](https://help.aliyun.com/document_detail/109196.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2D1AEB96-96D0-454E-B0DC-AE2A8DF08020
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnlockSecretNoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnlockSecretNoResponseBody) GoString() string {
	return s.String()
}

func (s *UnlockSecretNoResponseBody) SetCode(v string) *UnlockSecretNoResponseBody {
	s.Code = &v
	return s
}

func (s *UnlockSecretNoResponseBody) SetMessage(v string) *UnlockSecretNoResponseBody {
	s.Message = &v
	return s
}

func (s *UnlockSecretNoResponseBody) SetRequestId(v string) *UnlockSecretNoResponseBody {
	s.RequestId = &v
	return s
}

type UnlockSecretNoResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnlockSecretNoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnlockSecretNoResponse) String() string {
	return tea.Prettify(s)
}

func (s UnlockSecretNoResponse) GoString() string {
	return s.String()
}

func (s *UnlockSecretNoResponse) SetHeaders(v map[string]*string) *UnlockSecretNoResponse {
	s.Headers = v
	return s
}

func (s *UnlockSecretNoResponse) SetStatusCode(v int32) *UnlockSecretNoResponse {
	s.StatusCode = &v
	return s
}

func (s *UnlockSecretNoResponse) SetBody(v *UnlockSecretNoResponseBody) *UnlockSecretNoResponse {
	s.Body = v
	return s
}

type UpdateAxbBindFixedLineRequest struct {
	// 主叫侧放音编码
	//
	// example:
	//
	// 1,2,3
	Anucode *string `json:"Anucode,omitempty" xml:"Anucode,omitempty"`
	// 被叫侧放音编码
	//
	// example:
	//
	// 1,2
	Anucodecalled *string `json:"Anucodecalled,omitempty" xml:"Anucodecalled,omitempty"`
	// 应用id，请求和绑定时的appId必须一致
	//
	// This parameter is required.
	//
	// example:
	//
	// ALPT_1234
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 过期时间,单位：秒，必须为数字 0：不会自动解绑 非0：自动解绑周期
	//
	// example:
	//
	// 10
	Expiration *string `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	// 扩展参数
	Extra *UpdateAxbBindFixedLineRequestExtra `json:"Extra,omitempty" xml:"Extra,omitempty" type:"Struct"`
	// 消息请求标识
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234567890
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 接入商自有字段，不能超过100个字符
	//
	// example:
	//
	// remark
	Remark               *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 绑定id
	//
	// This parameter is required.
	//
	// example:
	//
	// A100X558X0000400023
	SubId *string `json:"SubId,omitempty" xml:"SubId,omitempty"`
	// 格式为yyyyMMddHHmmss。时间采用北京时间，24小时制。
	//
	// This parameter is required.
	//
	// example:
	//
	// 20150920190126
	Subts *string `json:"Subts,omitempty" xml:"Subts,omitempty"`
	// 接通后主叫侧的放音编码
	//
	// example:
	//
	// 1,2
	TAnucodeConnect *string `json:"TAnucodeConnect,omitempty" xml:"TAnucodeConnect,omitempty"`
	// 真实号码，telA,telB不允许同时更新
	//
	// example:
	//
	// 13900000000
	TelA *string `json:"TelA,omitempty" xml:"TelA,omitempty"`
	// 对端号码，telA,telB不允许同时更新
	//
	// example:
	//
	// 13005711234
	TelB *string `json:"TelB,omitempty" xml:"TelB,omitempty"`
	// 业务时间戳，格式yyyyMMddHHmmssSSS，时间采用北京时间，24小时制，精确至毫秒
	//
	// This parameter is required.
	//
	// example:
	//
	// 20161114143116001
	Ts *string `json:"Ts,omitempty" xml:"Ts,omitempty"`
}

func (s UpdateAxbBindFixedLineRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAxbBindFixedLineRequest) GoString() string {
	return s.String()
}

func (s *UpdateAxbBindFixedLineRequest) SetAnucode(v string) *UpdateAxbBindFixedLineRequest {
	s.Anucode = &v
	return s
}

func (s *UpdateAxbBindFixedLineRequest) SetAnucodecalled(v string) *UpdateAxbBindFixedLineRequest {
	s.Anucodecalled = &v
	return s
}

func (s *UpdateAxbBindFixedLineRequest) SetAppId(v string) *UpdateAxbBindFixedLineRequest {
	s.AppId = &v
	return s
}

func (s *UpdateAxbBindFixedLineRequest) SetExpiration(v string) *UpdateAxbBindFixedLineRequest {
	s.Expiration = &v
	return s
}

func (s *UpdateAxbBindFixedLineRequest) SetExtra(v *UpdateAxbBindFixedLineRequestExtra) *UpdateAxbBindFixedLineRequest {
	s.Extra = v
	return s
}

func (s *UpdateAxbBindFixedLineRequest) SetOrderId(v string) *UpdateAxbBindFixedLineRequest {
	s.OrderId = &v
	return s
}

func (s *UpdateAxbBindFixedLineRequest) SetOwnerId(v int64) *UpdateAxbBindFixedLineRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateAxbBindFixedLineRequest) SetRemark(v string) *UpdateAxbBindFixedLineRequest {
	s.Remark = &v
	return s
}

func (s *UpdateAxbBindFixedLineRequest) SetResourceOwnerAccount(v string) *UpdateAxbBindFixedLineRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateAxbBindFixedLineRequest) SetResourceOwnerId(v int64) *UpdateAxbBindFixedLineRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateAxbBindFixedLineRequest) SetSubId(v string) *UpdateAxbBindFixedLineRequest {
	s.SubId = &v
	return s
}

func (s *UpdateAxbBindFixedLineRequest) SetSubts(v string) *UpdateAxbBindFixedLineRequest {
	s.Subts = &v
	return s
}

func (s *UpdateAxbBindFixedLineRequest) SetTAnucodeConnect(v string) *UpdateAxbBindFixedLineRequest {
	s.TAnucodeConnect = &v
	return s
}

func (s *UpdateAxbBindFixedLineRequest) SetTelA(v string) *UpdateAxbBindFixedLineRequest {
	s.TelA = &v
	return s
}

func (s *UpdateAxbBindFixedLineRequest) SetTelB(v string) *UpdateAxbBindFixedLineRequest {
	s.TelB = &v
	return s
}

func (s *UpdateAxbBindFixedLineRequest) SetTs(v string) *UpdateAxbBindFixedLineRequest {
	s.Ts = &v
	return s
}

type UpdateAxbBindFixedLineRequestExtra struct {
	// 录音控制， 0：不录音 1：录音
	//
	// example:
	//
	// 0
	Callrecording *string `json:"Callrecording,omitempty" xml:"Callrecording,omitempty"`
}

func (s UpdateAxbBindFixedLineRequestExtra) String() string {
	return tea.Prettify(s)
}

func (s UpdateAxbBindFixedLineRequestExtra) GoString() string {
	return s.String()
}

func (s *UpdateAxbBindFixedLineRequestExtra) SetCallrecording(v string) *UpdateAxbBindFixedLineRequestExtra {
	s.Callrecording = &v
	return s
}

type UpdateAxbBindFixedLineShrinkRequest struct {
	// 主叫侧放音编码
	//
	// example:
	//
	// 1,2,3
	Anucode *string `json:"Anucode,omitempty" xml:"Anucode,omitempty"`
	// 被叫侧放音编码
	//
	// example:
	//
	// 1,2
	Anucodecalled *string `json:"Anucodecalled,omitempty" xml:"Anucodecalled,omitempty"`
	// 应用id，请求和绑定时的appId必须一致
	//
	// This parameter is required.
	//
	// example:
	//
	// ALPT_1234
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 过期时间,单位：秒，必须为数字 0：不会自动解绑 非0：自动解绑周期
	//
	// example:
	//
	// 10
	Expiration *string `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	// 扩展参数
	ExtraShrink *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	// 消息请求标识
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234567890
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 接入商自有字段，不能超过100个字符
	//
	// example:
	//
	// remark
	Remark               *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 绑定id
	//
	// This parameter is required.
	//
	// example:
	//
	// A100X558X0000400023
	SubId *string `json:"SubId,omitempty" xml:"SubId,omitempty"`
	// 格式为yyyyMMddHHmmss。时间采用北京时间，24小时制。
	//
	// This parameter is required.
	//
	// example:
	//
	// 20150920190126
	Subts *string `json:"Subts,omitempty" xml:"Subts,omitempty"`
	// 接通后主叫侧的放音编码
	//
	// example:
	//
	// 1,2
	TAnucodeConnect *string `json:"TAnucodeConnect,omitempty" xml:"TAnucodeConnect,omitempty"`
	// 真实号码，telA,telB不允许同时更新
	//
	// example:
	//
	// 13900000000
	TelA *string `json:"TelA,omitempty" xml:"TelA,omitempty"`
	// 对端号码，telA,telB不允许同时更新
	//
	// example:
	//
	// 13005711234
	TelB *string `json:"TelB,omitempty" xml:"TelB,omitempty"`
	// 业务时间戳，格式yyyyMMddHHmmssSSS，时间采用北京时间，24小时制，精确至毫秒
	//
	// This parameter is required.
	//
	// example:
	//
	// 20161114143116001
	Ts *string `json:"Ts,omitempty" xml:"Ts,omitempty"`
}

func (s UpdateAxbBindFixedLineShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAxbBindFixedLineShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateAxbBindFixedLineShrinkRequest) SetAnucode(v string) *UpdateAxbBindFixedLineShrinkRequest {
	s.Anucode = &v
	return s
}

func (s *UpdateAxbBindFixedLineShrinkRequest) SetAnucodecalled(v string) *UpdateAxbBindFixedLineShrinkRequest {
	s.Anucodecalled = &v
	return s
}

func (s *UpdateAxbBindFixedLineShrinkRequest) SetAppId(v string) *UpdateAxbBindFixedLineShrinkRequest {
	s.AppId = &v
	return s
}

func (s *UpdateAxbBindFixedLineShrinkRequest) SetExpiration(v string) *UpdateAxbBindFixedLineShrinkRequest {
	s.Expiration = &v
	return s
}

func (s *UpdateAxbBindFixedLineShrinkRequest) SetExtraShrink(v string) *UpdateAxbBindFixedLineShrinkRequest {
	s.ExtraShrink = &v
	return s
}

func (s *UpdateAxbBindFixedLineShrinkRequest) SetOrderId(v string) *UpdateAxbBindFixedLineShrinkRequest {
	s.OrderId = &v
	return s
}

func (s *UpdateAxbBindFixedLineShrinkRequest) SetOwnerId(v int64) *UpdateAxbBindFixedLineShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateAxbBindFixedLineShrinkRequest) SetRemark(v string) *UpdateAxbBindFixedLineShrinkRequest {
	s.Remark = &v
	return s
}

func (s *UpdateAxbBindFixedLineShrinkRequest) SetResourceOwnerAccount(v string) *UpdateAxbBindFixedLineShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateAxbBindFixedLineShrinkRequest) SetResourceOwnerId(v int64) *UpdateAxbBindFixedLineShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateAxbBindFixedLineShrinkRequest) SetSubId(v string) *UpdateAxbBindFixedLineShrinkRequest {
	s.SubId = &v
	return s
}

func (s *UpdateAxbBindFixedLineShrinkRequest) SetSubts(v string) *UpdateAxbBindFixedLineShrinkRequest {
	s.Subts = &v
	return s
}

func (s *UpdateAxbBindFixedLineShrinkRequest) SetTAnucodeConnect(v string) *UpdateAxbBindFixedLineShrinkRequest {
	s.TAnucodeConnect = &v
	return s
}

func (s *UpdateAxbBindFixedLineShrinkRequest) SetTelA(v string) *UpdateAxbBindFixedLineShrinkRequest {
	s.TelA = &v
	return s
}

func (s *UpdateAxbBindFixedLineShrinkRequest) SetTelB(v string) *UpdateAxbBindFixedLineShrinkRequest {
	s.TelB = &v
	return s
}

func (s *UpdateAxbBindFixedLineShrinkRequest) SetTs(v string) *UpdateAxbBindFixedLineShrinkRequest {
	s.Ts = &v
	return s
}

type UpdateAxbBindFixedLineResponseBody struct {
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *UpdateAxbBindFixedLineResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 73A5C73A-1D97-54B6-B47C-541BE59F84D9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateAxbBindFixedLineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAxbBindFixedLineResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAxbBindFixedLineResponseBody) SetAccessDeniedDetail(v string) *UpdateAxbBindFixedLineResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UpdateAxbBindFixedLineResponseBody) SetCode(v string) *UpdateAxbBindFixedLineResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateAxbBindFixedLineResponseBody) SetData(v *UpdateAxbBindFixedLineResponseBodyData) *UpdateAxbBindFixedLineResponseBody {
	s.Data = v
	return s
}

func (s *UpdateAxbBindFixedLineResponseBody) SetMessage(v string) *UpdateAxbBindFixedLineResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateAxbBindFixedLineResponseBody) SetRequestId(v string) *UpdateAxbBindFixedLineResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAxbBindFixedLineResponseBody) SetSuccess(v bool) *UpdateAxbBindFixedLineResponseBody {
	s.Success = &v
	return s
}

type UpdateAxbBindFixedLineResponseBodyData struct {
	// 响应码 0-成功
	//
	// example:
	//
	// 0
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 响应消息
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 是否处理成功  true-成功
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateAxbBindFixedLineResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UpdateAxbBindFixedLineResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateAxbBindFixedLineResponseBodyData) SetCode(v string) *UpdateAxbBindFixedLineResponseBodyData {
	s.Code = &v
	return s
}

func (s *UpdateAxbBindFixedLineResponseBodyData) SetMessage(v string) *UpdateAxbBindFixedLineResponseBodyData {
	s.Message = &v
	return s
}

func (s *UpdateAxbBindFixedLineResponseBodyData) SetSuccess(v bool) *UpdateAxbBindFixedLineResponseBodyData {
	s.Success = &v
	return s
}

type UpdateAxbBindFixedLineResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAxbBindFixedLineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAxbBindFixedLineResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAxbBindFixedLineResponse) GoString() string {
	return s.String()
}

func (s *UpdateAxbBindFixedLineResponse) SetHeaders(v map[string]*string) *UpdateAxbBindFixedLineResponse {
	s.Headers = v
	return s
}

func (s *UpdateAxbBindFixedLineResponse) SetStatusCode(v int32) *UpdateAxbBindFixedLineResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAxbBindFixedLineResponse) SetBody(v *UpdateAxbBindFixedLineResponseBody) *UpdateAxbBindFixedLineResponse {
	s.Body = v
	return s
}

type UpdateAxnBindFixedLineRequest struct {
	// 放音编码必须包含3个场景的编码。按照“B->X,A->X,其他号码->X”的顺序填写编码，编码之间以逗号分隔。  比如：“1,2,3”表示B->X放音编号为1，A->X放音编号为2， 其他号码->X放音编号为3。
	//
	// example:
	//
	// 示例值示例值示例值
	Anucode *string `json:"Anucode,omitempty" xml:"Anucode,omitempty"`
	// 被叫侧放音编码  被叫放音编码必须包含2个场景的编码。按照“A被叫,其他号码被叫”的顺序填写编码，编码之间以逗号分隔。  比如：“1,2”表示A号码为被叫侧接听时的放音编号为1，其他号码为被叫侧接听时的放音编号为2。
	//
	// example:
	//
	// 10001,10002
	Anucodecalled *string `json:"Anucodecalled,omitempty" xml:"Anucodecalled,omitempty"`
	// 业务id标识，由阿里云分配给客户侧
	//
	// This parameter is required.
	//
	// example:
	//
	// alitest
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 位：秒，必须为数字 0：不会自动解绑 非0：自动解绑周期
	//
	// example:
	//
	// 60
	Expiration *string `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	// 额外字段
	Extra *UpdateAxnBindFixedLineRequestExtra `json:"Extra,omitempty" xml:"Extra,omitempty" type:"Struct"`
	// 消息请求唯一标识
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345dkwkd99d
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 接入商自有字段，最大100字符长度
	//
	// example:
	//
	// 1234
	Remark               *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 绑定id
	//
	// This parameter is required.
	//
	// example:
	//
	// GHX0534X202504221531579290029-2-1-aliaxn
	SubId *string `json:"SubId,omitempty" xml:"SubId,omitempty"`
	// 格式为yyyyMMddHHmmss。时间采用北京时间，24小时制。
	//
	// This parameter is required.
	//
	// example:
	//
	// 20250421141723098
	Subts *string `json:"Subts,omitempty" xml:"Subts,omitempty"`
	// 接通后主叫侧放音编码  接通后主叫侧放音编码必须包含2个场景的编码。按照“A被叫,其他号码被叫”的顺序填写编码，编码之间以逗号分隔。  比如：“1,2”表示A号码为被叫时主叫侧的放音编号为1，其他号码被叫为被叫时主叫侧的放音编号为2
	//
	// example:
	//
	// 10001,10002
	TAnucodeConnect *string `json:"TAnucodeConnect,omitempty" xml:"TAnucodeConnect,omitempty"`
	// A号码
	//
	// example:
	//
	// 15500001111
	TelA *string `json:"TelA,omitempty" xml:"TelA,omitempty"`
	// B号码
	//
	// example:
	//
	// 15500002222
	TelB *string `json:"TelB,omitempty" xml:"TelB,omitempty"`
	// 格式yyyyMMddHHmmssSSS，时间采用北京时间，24小时制，精确至毫秒
	//
	// This parameter is required.
	//
	// example:
	//
	// 20250421141723098
	Ts *string `json:"Ts,omitempty" xml:"Ts,omitempty"`
}

func (s UpdateAxnBindFixedLineRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAxnBindFixedLineRequest) GoString() string {
	return s.String()
}

func (s *UpdateAxnBindFixedLineRequest) SetAnucode(v string) *UpdateAxnBindFixedLineRequest {
	s.Anucode = &v
	return s
}

func (s *UpdateAxnBindFixedLineRequest) SetAnucodecalled(v string) *UpdateAxnBindFixedLineRequest {
	s.Anucodecalled = &v
	return s
}

func (s *UpdateAxnBindFixedLineRequest) SetAppId(v string) *UpdateAxnBindFixedLineRequest {
	s.AppId = &v
	return s
}

func (s *UpdateAxnBindFixedLineRequest) SetExpiration(v string) *UpdateAxnBindFixedLineRequest {
	s.Expiration = &v
	return s
}

func (s *UpdateAxnBindFixedLineRequest) SetExtra(v *UpdateAxnBindFixedLineRequestExtra) *UpdateAxnBindFixedLineRequest {
	s.Extra = v
	return s
}

func (s *UpdateAxnBindFixedLineRequest) SetOrderId(v string) *UpdateAxnBindFixedLineRequest {
	s.OrderId = &v
	return s
}

func (s *UpdateAxnBindFixedLineRequest) SetOwnerId(v int64) *UpdateAxnBindFixedLineRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateAxnBindFixedLineRequest) SetRemark(v string) *UpdateAxnBindFixedLineRequest {
	s.Remark = &v
	return s
}

func (s *UpdateAxnBindFixedLineRequest) SetResourceOwnerAccount(v string) *UpdateAxnBindFixedLineRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateAxnBindFixedLineRequest) SetResourceOwnerId(v int64) *UpdateAxnBindFixedLineRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateAxnBindFixedLineRequest) SetSubId(v string) *UpdateAxnBindFixedLineRequest {
	s.SubId = &v
	return s
}

func (s *UpdateAxnBindFixedLineRequest) SetSubts(v string) *UpdateAxnBindFixedLineRequest {
	s.Subts = &v
	return s
}

func (s *UpdateAxnBindFixedLineRequest) SetTAnucodeConnect(v string) *UpdateAxnBindFixedLineRequest {
	s.TAnucodeConnect = &v
	return s
}

func (s *UpdateAxnBindFixedLineRequest) SetTelA(v string) *UpdateAxnBindFixedLineRequest {
	s.TelA = &v
	return s
}

func (s *UpdateAxnBindFixedLineRequest) SetTelB(v string) *UpdateAxnBindFixedLineRequest {
	s.TelB = &v
	return s
}

func (s *UpdateAxnBindFixedLineRequest) SetTs(v string) *UpdateAxnBindFixedLineRequest {
	s.Ts = &v
	return s
}

type UpdateAxnBindFixedLineRequestExtra struct {
	// A通过X呼叫，即A主叫X，仅下列值有效。默认是0。 0：不能外呼 1：接续最近的N号码 2：回拨固定号码：telB
	//
	// example:
	//
	// 0
	Callback *string `json:"Callback,omitempty" xml:"Callback,omitempty"`
	// 录音控制，仅下列值有效。默认是0（不开通录音功能）。 0：不录音 1：录音
	//
	// example:
	//
	// 1
	Callrecording *string `json:"Callrecording,omitempty" xml:"Callrecording,omitempty"`
}

func (s UpdateAxnBindFixedLineRequestExtra) String() string {
	return tea.Prettify(s)
}

func (s UpdateAxnBindFixedLineRequestExtra) GoString() string {
	return s.String()
}

func (s *UpdateAxnBindFixedLineRequestExtra) SetCallback(v string) *UpdateAxnBindFixedLineRequestExtra {
	s.Callback = &v
	return s
}

func (s *UpdateAxnBindFixedLineRequestExtra) SetCallrecording(v string) *UpdateAxnBindFixedLineRequestExtra {
	s.Callrecording = &v
	return s
}

type UpdateAxnBindFixedLineShrinkRequest struct {
	// 放音编码必须包含3个场景的编码。按照“B->X,A->X,其他号码->X”的顺序填写编码，编码之间以逗号分隔。  比如：“1,2,3”表示B->X放音编号为1，A->X放音编号为2， 其他号码->X放音编号为3。
	//
	// example:
	//
	// 示例值示例值示例值
	Anucode *string `json:"Anucode,omitempty" xml:"Anucode,omitempty"`
	// 被叫侧放音编码  被叫放音编码必须包含2个场景的编码。按照“A被叫,其他号码被叫”的顺序填写编码，编码之间以逗号分隔。  比如：“1,2”表示A号码为被叫侧接听时的放音编号为1，其他号码为被叫侧接听时的放音编号为2。
	//
	// example:
	//
	// 10001,10002
	Anucodecalled *string `json:"Anucodecalled,omitempty" xml:"Anucodecalled,omitempty"`
	// 业务id标识，由阿里云分配给客户侧
	//
	// This parameter is required.
	//
	// example:
	//
	// alitest
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 位：秒，必须为数字 0：不会自动解绑 非0：自动解绑周期
	//
	// example:
	//
	// 60
	Expiration *string `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	// 额外字段
	ExtraShrink *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	// 消息请求唯一标识
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345dkwkd99d
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 接入商自有字段，最大100字符长度
	//
	// example:
	//
	// 1234
	Remark               *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 绑定id
	//
	// This parameter is required.
	//
	// example:
	//
	// GHX0534X202504221531579290029-2-1-aliaxn
	SubId *string `json:"SubId,omitempty" xml:"SubId,omitempty"`
	// 格式为yyyyMMddHHmmss。时间采用北京时间，24小时制。
	//
	// This parameter is required.
	//
	// example:
	//
	// 20250421141723098
	Subts *string `json:"Subts,omitempty" xml:"Subts,omitempty"`
	// 接通后主叫侧放音编码  接通后主叫侧放音编码必须包含2个场景的编码。按照“A被叫,其他号码被叫”的顺序填写编码，编码之间以逗号分隔。  比如：“1,2”表示A号码为被叫时主叫侧的放音编号为1，其他号码被叫为被叫时主叫侧的放音编号为2
	//
	// example:
	//
	// 10001,10002
	TAnucodeConnect *string `json:"TAnucodeConnect,omitempty" xml:"TAnucodeConnect,omitempty"`
	// A号码
	//
	// example:
	//
	// 15500001111
	TelA *string `json:"TelA,omitempty" xml:"TelA,omitempty"`
	// B号码
	//
	// example:
	//
	// 15500002222
	TelB *string `json:"TelB,omitempty" xml:"TelB,omitempty"`
	// 格式yyyyMMddHHmmssSSS，时间采用北京时间，24小时制，精确至毫秒
	//
	// This parameter is required.
	//
	// example:
	//
	// 20250421141723098
	Ts *string `json:"Ts,omitempty" xml:"Ts,omitempty"`
}

func (s UpdateAxnBindFixedLineShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAxnBindFixedLineShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateAxnBindFixedLineShrinkRequest) SetAnucode(v string) *UpdateAxnBindFixedLineShrinkRequest {
	s.Anucode = &v
	return s
}

func (s *UpdateAxnBindFixedLineShrinkRequest) SetAnucodecalled(v string) *UpdateAxnBindFixedLineShrinkRequest {
	s.Anucodecalled = &v
	return s
}

func (s *UpdateAxnBindFixedLineShrinkRequest) SetAppId(v string) *UpdateAxnBindFixedLineShrinkRequest {
	s.AppId = &v
	return s
}

func (s *UpdateAxnBindFixedLineShrinkRequest) SetExpiration(v string) *UpdateAxnBindFixedLineShrinkRequest {
	s.Expiration = &v
	return s
}

func (s *UpdateAxnBindFixedLineShrinkRequest) SetExtraShrink(v string) *UpdateAxnBindFixedLineShrinkRequest {
	s.ExtraShrink = &v
	return s
}

func (s *UpdateAxnBindFixedLineShrinkRequest) SetOrderId(v string) *UpdateAxnBindFixedLineShrinkRequest {
	s.OrderId = &v
	return s
}

func (s *UpdateAxnBindFixedLineShrinkRequest) SetOwnerId(v int64) *UpdateAxnBindFixedLineShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateAxnBindFixedLineShrinkRequest) SetRemark(v string) *UpdateAxnBindFixedLineShrinkRequest {
	s.Remark = &v
	return s
}

func (s *UpdateAxnBindFixedLineShrinkRequest) SetResourceOwnerAccount(v string) *UpdateAxnBindFixedLineShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateAxnBindFixedLineShrinkRequest) SetResourceOwnerId(v int64) *UpdateAxnBindFixedLineShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateAxnBindFixedLineShrinkRequest) SetSubId(v string) *UpdateAxnBindFixedLineShrinkRequest {
	s.SubId = &v
	return s
}

func (s *UpdateAxnBindFixedLineShrinkRequest) SetSubts(v string) *UpdateAxnBindFixedLineShrinkRequest {
	s.Subts = &v
	return s
}

func (s *UpdateAxnBindFixedLineShrinkRequest) SetTAnucodeConnect(v string) *UpdateAxnBindFixedLineShrinkRequest {
	s.TAnucodeConnect = &v
	return s
}

func (s *UpdateAxnBindFixedLineShrinkRequest) SetTelA(v string) *UpdateAxnBindFixedLineShrinkRequest {
	s.TelA = &v
	return s
}

func (s *UpdateAxnBindFixedLineShrinkRequest) SetTelB(v string) *UpdateAxnBindFixedLineShrinkRequest {
	s.TelB = &v
	return s
}

func (s *UpdateAxnBindFixedLineShrinkRequest) SetTs(v string) *UpdateAxnBindFixedLineShrinkRequest {
	s.Ts = &v
	return s
}

type UpdateAxnBindFixedLineResponseBody struct {
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *UpdateAxnBindFixedLineResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 3FDD0A8F-34F1-5BD4-AF9F-CD90B3DE7C06
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateAxnBindFixedLineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAxnBindFixedLineResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAxnBindFixedLineResponseBody) SetAccessDeniedDetail(v string) *UpdateAxnBindFixedLineResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UpdateAxnBindFixedLineResponseBody) SetCode(v string) *UpdateAxnBindFixedLineResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateAxnBindFixedLineResponseBody) SetData(v *UpdateAxnBindFixedLineResponseBodyData) *UpdateAxnBindFixedLineResponseBody {
	s.Data = v
	return s
}

func (s *UpdateAxnBindFixedLineResponseBody) SetMessage(v string) *UpdateAxnBindFixedLineResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateAxnBindFixedLineResponseBody) SetRequestId(v string) *UpdateAxnBindFixedLineResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAxnBindFixedLineResponseBody) SetSuccess(v bool) *UpdateAxnBindFixedLineResponseBody {
	s.Success = &v
	return s
}

type UpdateAxnBindFixedLineResponseBodyData struct {
	// 响应码 0：成功，其它失败，具体见文档
	//
	// example:
	//
	// 0
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 描述信息
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateAxnBindFixedLineResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UpdateAxnBindFixedLineResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateAxnBindFixedLineResponseBodyData) SetCode(v string) *UpdateAxnBindFixedLineResponseBodyData {
	s.Code = &v
	return s
}

func (s *UpdateAxnBindFixedLineResponseBodyData) SetMessage(v string) *UpdateAxnBindFixedLineResponseBodyData {
	s.Message = &v
	return s
}

func (s *UpdateAxnBindFixedLineResponseBodyData) SetSuccess(v bool) *UpdateAxnBindFixedLineResponseBodyData {
	s.Success = &v
	return s
}

type UpdateAxnBindFixedLineResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAxnBindFixedLineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAxnBindFixedLineResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAxnBindFixedLineResponse) GoString() string {
	return s.String()
}

func (s *UpdateAxnBindFixedLineResponse) SetHeaders(v map[string]*string) *UpdateAxnBindFixedLineResponse {
	s.Headers = v
	return s
}

func (s *UpdateAxnBindFixedLineResponse) SetStatusCode(v int32) *UpdateAxnBindFixedLineResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAxnBindFixedLineResponse) SetBody(v *UpdateAxnBindFixedLineResponseBody) *UpdateAxnBindFixedLineResponse {
	s.Body = v
	return s
}

type UpdateAxnExtensionBindFixedLineRequest struct {
	// 放音编码必须包含3个场景的编码。按照“B->X,A->X,其他号码->X”的顺序填写编码，编码之间以逗号分隔。  AXN分机号业务的放音编码,B->X和其他号码->X的编码一致  比如：“1,2,3”表示B->X放音编号为1，A->X放音编号为2， 其他号码->X放音编号为3
	//
	// example:
	//
	// 10001,10002,10003
	Anucode *string `json:"Anucode,omitempty" xml:"Anucode,omitempty"`
	// 被叫侧放音编码  被叫放音编码必须包含2个场景的编码。按照“A被叫,其他号码被叫”的顺序填写编码，编码之间以逗号分隔。  比如：“1,2”表示A号码为被叫侧接听时的放音编号为1，其他号码为被叫侧接听时的放音编号为2
	//
	// example:
	//
	// 10001,10002
	Anucodecalled *string `json:"Anucodecalled,omitempty" xml:"Anucodecalled,omitempty"`
	// 业务id标识，由阿里云分配给客户侧
	//
	// This parameter is required.
	//
	// example:
	//
	// alitest
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 单位：秒，必须为数字 0：不会自动解绑 非0：自动解绑周期
	//
	// example:
	//
	// 60
	Expiration *string `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	// AXx的扩展参数项
	Extraaxx *UpdateAxnExtensionBindFixedLineRequestExtraaxx `json:"Extraaxx,omitempty" xml:"Extraaxx,omitempty" type:"Struct"`
	// 消息请求标识
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345dkwkd99d
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 接入商自有字段，最大100字符长度
	//
	// example:
	//
	// 1233
	Remark               *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 绑定id
	//
	// This parameter is required.
	//
	// example:
	//
	// 可参考绑定响应
	SubId *string `json:"SubId,omitempty" xml:"SubId,omitempty"`
	// 格式为yyyyMMddHHmmss。时间采用北京时间，24小时制。
	//
	// This parameter is required.
	//
	// example:
	//
	// 20250421141723098
	Subts *string `json:"Subts,omitempty" xml:"Subts,omitempty"`
	// 接通后主叫侧放音编码  接通后主叫侧放音编码必须包含2个场景的编码。按照“A被叫,其他号码被叫”的顺序填写编码，编码之间以逗号分隔。  比如：“1,2”表示A号码为被叫时主叫侧的放音编号为1，其他号码为被叫时主叫侧的放音编号为2
	//
	// example:
	//
	// 10001,10002
	TAnucodeConnect *string `json:"TAnucodeConnect,omitempty" xml:"TAnucodeConnect,omitempty"`
	// A号码
	//
	// This parameter is required.
	//
	// example:
	//
	// 15500001111
	TelA *string `json:"TelA,omitempty" xml:"TelA,omitempty"`
	// 格式yyyyMMddHHmmssSSS，时间采用北京时间，24小时制，精确至毫秒
	//
	// This parameter is required.
	//
	// example:
	//
	// 20250421141723098
	Ts *string `json:"Ts,omitempty" xml:"Ts,omitempty"`
}

func (s UpdateAxnExtensionBindFixedLineRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAxnExtensionBindFixedLineRequest) GoString() string {
	return s.String()
}

func (s *UpdateAxnExtensionBindFixedLineRequest) SetAnucode(v string) *UpdateAxnExtensionBindFixedLineRequest {
	s.Anucode = &v
	return s
}

func (s *UpdateAxnExtensionBindFixedLineRequest) SetAnucodecalled(v string) *UpdateAxnExtensionBindFixedLineRequest {
	s.Anucodecalled = &v
	return s
}

func (s *UpdateAxnExtensionBindFixedLineRequest) SetAppId(v string) *UpdateAxnExtensionBindFixedLineRequest {
	s.AppId = &v
	return s
}

func (s *UpdateAxnExtensionBindFixedLineRequest) SetExpiration(v string) *UpdateAxnExtensionBindFixedLineRequest {
	s.Expiration = &v
	return s
}

func (s *UpdateAxnExtensionBindFixedLineRequest) SetExtraaxx(v *UpdateAxnExtensionBindFixedLineRequestExtraaxx) *UpdateAxnExtensionBindFixedLineRequest {
	s.Extraaxx = v
	return s
}

func (s *UpdateAxnExtensionBindFixedLineRequest) SetOrderId(v string) *UpdateAxnExtensionBindFixedLineRequest {
	s.OrderId = &v
	return s
}

func (s *UpdateAxnExtensionBindFixedLineRequest) SetOwnerId(v int64) *UpdateAxnExtensionBindFixedLineRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateAxnExtensionBindFixedLineRequest) SetRemark(v string) *UpdateAxnExtensionBindFixedLineRequest {
	s.Remark = &v
	return s
}

func (s *UpdateAxnExtensionBindFixedLineRequest) SetResourceOwnerAccount(v string) *UpdateAxnExtensionBindFixedLineRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateAxnExtensionBindFixedLineRequest) SetResourceOwnerId(v int64) *UpdateAxnExtensionBindFixedLineRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateAxnExtensionBindFixedLineRequest) SetSubId(v string) *UpdateAxnExtensionBindFixedLineRequest {
	s.SubId = &v
	return s
}

func (s *UpdateAxnExtensionBindFixedLineRequest) SetSubts(v string) *UpdateAxnExtensionBindFixedLineRequest {
	s.Subts = &v
	return s
}

func (s *UpdateAxnExtensionBindFixedLineRequest) SetTAnucodeConnect(v string) *UpdateAxnExtensionBindFixedLineRequest {
	s.TAnucodeConnect = &v
	return s
}

func (s *UpdateAxnExtensionBindFixedLineRequest) SetTelA(v string) *UpdateAxnExtensionBindFixedLineRequest {
	s.TelA = &v
	return s
}

func (s *UpdateAxnExtensionBindFixedLineRequest) SetTs(v string) *UpdateAxnExtensionBindFixedLineRequest {
	s.Ts = &v
	return s
}

type UpdateAxnExtensionBindFixedLineRequestExtraaxx struct {
	// 可选。 A通过X呼叫，即A主叫X，仅下列值有效。默认是0。 0：不能外呼 1：接续最近的B号码
	//
	// example:
	//
	// 0
	Callback *string `json:"Callback,omitempty" xml:"Callback,omitempty"`
	// 录音控制，仅下列值有效。默认是0（不开通录音功能）。 0：不录音 1：录音
	//
	// example:
	//
	// 0
	Callrecording *string `json:"Callrecording,omitempty" xml:"Callrecording,omitempty"`
}

func (s UpdateAxnExtensionBindFixedLineRequestExtraaxx) String() string {
	return tea.Prettify(s)
}

func (s UpdateAxnExtensionBindFixedLineRequestExtraaxx) GoString() string {
	return s.String()
}

func (s *UpdateAxnExtensionBindFixedLineRequestExtraaxx) SetCallback(v string) *UpdateAxnExtensionBindFixedLineRequestExtraaxx {
	s.Callback = &v
	return s
}

func (s *UpdateAxnExtensionBindFixedLineRequestExtraaxx) SetCallrecording(v string) *UpdateAxnExtensionBindFixedLineRequestExtraaxx {
	s.Callrecording = &v
	return s
}

type UpdateAxnExtensionBindFixedLineShrinkRequest struct {
	// 放音编码必须包含3个场景的编码。按照“B->X,A->X,其他号码->X”的顺序填写编码，编码之间以逗号分隔。  AXN分机号业务的放音编码,B->X和其他号码->X的编码一致  比如：“1,2,3”表示B->X放音编号为1，A->X放音编号为2， 其他号码->X放音编号为3
	//
	// example:
	//
	// 10001,10002,10003
	Anucode *string `json:"Anucode,omitempty" xml:"Anucode,omitempty"`
	// 被叫侧放音编码  被叫放音编码必须包含2个场景的编码。按照“A被叫,其他号码被叫”的顺序填写编码，编码之间以逗号分隔。  比如：“1,2”表示A号码为被叫侧接听时的放音编号为1，其他号码为被叫侧接听时的放音编号为2
	//
	// example:
	//
	// 10001,10002
	Anucodecalled *string `json:"Anucodecalled,omitempty" xml:"Anucodecalled,omitempty"`
	// 业务id标识，由阿里云分配给客户侧
	//
	// This parameter is required.
	//
	// example:
	//
	// alitest
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 单位：秒，必须为数字 0：不会自动解绑 非0：自动解绑周期
	//
	// example:
	//
	// 60
	Expiration *string `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	// AXx的扩展参数项
	ExtraaxxShrink *string `json:"Extraaxx,omitempty" xml:"Extraaxx,omitempty"`
	// 消息请求标识
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345dkwkd99d
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 接入商自有字段，最大100字符长度
	//
	// example:
	//
	// 1233
	Remark               *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 绑定id
	//
	// This parameter is required.
	//
	// example:
	//
	// 可参考绑定响应
	SubId *string `json:"SubId,omitempty" xml:"SubId,omitempty"`
	// 格式为yyyyMMddHHmmss。时间采用北京时间，24小时制。
	//
	// This parameter is required.
	//
	// example:
	//
	// 20250421141723098
	Subts *string `json:"Subts,omitempty" xml:"Subts,omitempty"`
	// 接通后主叫侧放音编码  接通后主叫侧放音编码必须包含2个场景的编码。按照“A被叫,其他号码被叫”的顺序填写编码，编码之间以逗号分隔。  比如：“1,2”表示A号码为被叫时主叫侧的放音编号为1，其他号码为被叫时主叫侧的放音编号为2
	//
	// example:
	//
	// 10001,10002
	TAnucodeConnect *string `json:"TAnucodeConnect,omitempty" xml:"TAnucodeConnect,omitempty"`
	// A号码
	//
	// This parameter is required.
	//
	// example:
	//
	// 15500001111
	TelA *string `json:"TelA,omitempty" xml:"TelA,omitempty"`
	// 格式yyyyMMddHHmmssSSS，时间采用北京时间，24小时制，精确至毫秒
	//
	// This parameter is required.
	//
	// example:
	//
	// 20250421141723098
	Ts *string `json:"Ts,omitempty" xml:"Ts,omitempty"`
}

func (s UpdateAxnExtensionBindFixedLineShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAxnExtensionBindFixedLineShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateAxnExtensionBindFixedLineShrinkRequest) SetAnucode(v string) *UpdateAxnExtensionBindFixedLineShrinkRequest {
	s.Anucode = &v
	return s
}

func (s *UpdateAxnExtensionBindFixedLineShrinkRequest) SetAnucodecalled(v string) *UpdateAxnExtensionBindFixedLineShrinkRequest {
	s.Anucodecalled = &v
	return s
}

func (s *UpdateAxnExtensionBindFixedLineShrinkRequest) SetAppId(v string) *UpdateAxnExtensionBindFixedLineShrinkRequest {
	s.AppId = &v
	return s
}

func (s *UpdateAxnExtensionBindFixedLineShrinkRequest) SetExpiration(v string) *UpdateAxnExtensionBindFixedLineShrinkRequest {
	s.Expiration = &v
	return s
}

func (s *UpdateAxnExtensionBindFixedLineShrinkRequest) SetExtraaxxShrink(v string) *UpdateAxnExtensionBindFixedLineShrinkRequest {
	s.ExtraaxxShrink = &v
	return s
}

func (s *UpdateAxnExtensionBindFixedLineShrinkRequest) SetOrderId(v string) *UpdateAxnExtensionBindFixedLineShrinkRequest {
	s.OrderId = &v
	return s
}

func (s *UpdateAxnExtensionBindFixedLineShrinkRequest) SetOwnerId(v int64) *UpdateAxnExtensionBindFixedLineShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateAxnExtensionBindFixedLineShrinkRequest) SetRemark(v string) *UpdateAxnExtensionBindFixedLineShrinkRequest {
	s.Remark = &v
	return s
}

func (s *UpdateAxnExtensionBindFixedLineShrinkRequest) SetResourceOwnerAccount(v string) *UpdateAxnExtensionBindFixedLineShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateAxnExtensionBindFixedLineShrinkRequest) SetResourceOwnerId(v int64) *UpdateAxnExtensionBindFixedLineShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateAxnExtensionBindFixedLineShrinkRequest) SetSubId(v string) *UpdateAxnExtensionBindFixedLineShrinkRequest {
	s.SubId = &v
	return s
}

func (s *UpdateAxnExtensionBindFixedLineShrinkRequest) SetSubts(v string) *UpdateAxnExtensionBindFixedLineShrinkRequest {
	s.Subts = &v
	return s
}

func (s *UpdateAxnExtensionBindFixedLineShrinkRequest) SetTAnucodeConnect(v string) *UpdateAxnExtensionBindFixedLineShrinkRequest {
	s.TAnucodeConnect = &v
	return s
}

func (s *UpdateAxnExtensionBindFixedLineShrinkRequest) SetTelA(v string) *UpdateAxnExtensionBindFixedLineShrinkRequest {
	s.TelA = &v
	return s
}

func (s *UpdateAxnExtensionBindFixedLineShrinkRequest) SetTs(v string) *UpdateAxnExtensionBindFixedLineShrinkRequest {
	s.Ts = &v
	return s
}

type UpdateAxnExtensionBindFixedLineResponseBody struct {
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *UpdateAxnExtensionBindFixedLineResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 31031C54-7727-5057-9ED1-FA276B64205E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateAxnExtensionBindFixedLineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAxnExtensionBindFixedLineResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAxnExtensionBindFixedLineResponseBody) SetAccessDeniedDetail(v string) *UpdateAxnExtensionBindFixedLineResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UpdateAxnExtensionBindFixedLineResponseBody) SetCode(v string) *UpdateAxnExtensionBindFixedLineResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateAxnExtensionBindFixedLineResponseBody) SetData(v *UpdateAxnExtensionBindFixedLineResponseBodyData) *UpdateAxnExtensionBindFixedLineResponseBody {
	s.Data = v
	return s
}

func (s *UpdateAxnExtensionBindFixedLineResponseBody) SetMessage(v string) *UpdateAxnExtensionBindFixedLineResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateAxnExtensionBindFixedLineResponseBody) SetRequestId(v string) *UpdateAxnExtensionBindFixedLineResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAxnExtensionBindFixedLineResponseBody) SetSuccess(v bool) *UpdateAxnExtensionBindFixedLineResponseBody {
	s.Success = &v
	return s
}

type UpdateAxnExtensionBindFixedLineResponseBodyData struct {
	// 响应码 0：成功，其它失败，具体见文档
	//
	// example:
	//
	// 0
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 描述信息
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateAxnExtensionBindFixedLineResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UpdateAxnExtensionBindFixedLineResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateAxnExtensionBindFixedLineResponseBodyData) SetCode(v string) *UpdateAxnExtensionBindFixedLineResponseBodyData {
	s.Code = &v
	return s
}

func (s *UpdateAxnExtensionBindFixedLineResponseBodyData) SetMessage(v string) *UpdateAxnExtensionBindFixedLineResponseBodyData {
	s.Message = &v
	return s
}

func (s *UpdateAxnExtensionBindFixedLineResponseBodyData) SetSuccess(v bool) *UpdateAxnExtensionBindFixedLineResponseBodyData {
	s.Success = &v
	return s
}

type UpdateAxnExtensionBindFixedLineResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAxnExtensionBindFixedLineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAxnExtensionBindFixedLineResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAxnExtensionBindFixedLineResponse) GoString() string {
	return s.String()
}

func (s *UpdateAxnExtensionBindFixedLineResponse) SetHeaders(v map[string]*string) *UpdateAxnExtensionBindFixedLineResponse {
	s.Headers = v
	return s
}

func (s *UpdateAxnExtensionBindFixedLineResponse) SetStatusCode(v int32) *UpdateAxnExtensionBindFixedLineResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAxnExtensionBindFixedLineResponse) SetBody(v *UpdateAxnExtensionBindFixedLineResponseBody) *UpdateAxnExtensionBindFixedLineResponse {
	s.Body = v
	return s
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
	return tea.Prettify(s)
}

func (s UpdateSubscriptionRequest) GoString() string {
	return s.String()
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

type UpdateSubscriptionResponseBody struct {
	// The response code.
	//
	// 	- The value OK indicates that the request was successful.
	//
	// 	- Other values indicate that the request failed. For more information, see [Error codes](https://help.aliyun.com/document_detail/109196.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 986BCB6D-C9BF-42F9-91CE-3A9901233D36
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateSubscriptionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateSubscriptionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSubscriptionResponseBody) SetCode(v string) *UpdateSubscriptionResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateSubscriptionResponseBody) SetMessage(v string) *UpdateSubscriptionResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateSubscriptionResponseBody) SetRequestId(v string) *UpdateSubscriptionResponseBody {
	s.RequestId = &v
	return s
}

type UpdateSubscriptionResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSubscriptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSubscriptionResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateSubscriptionResponse) GoString() string {
	return s.String()
}

func (s *UpdateSubscriptionResponse) SetHeaders(v map[string]*string) *UpdateSubscriptionResponse {
	s.Headers = v
	return s
}

func (s *UpdateSubscriptionResponse) SetStatusCode(v int32) *UpdateSubscriptionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSubscriptionResponse) SetBody(v *UpdateSubscriptionResponseBody) *UpdateSubscriptionResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.SignatureAlgorithm = tea.String("v2")
	client.EndpointRule = tea.String("central")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("dyplsapi"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds a tracking number for a private number in the AXN binding.
//
// Description:
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 5,000 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - AddAxnTrackNoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddAxnTrackNoResponse
func (client *Client) AddAxnTrackNoWithOptions(request *AddAxnTrackNoRequest, runtime *util.RuntimeOptions) (_result *AddAxnTrackNoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNoX)) {
		query["PhoneNoX"] = request.PhoneNoX
	}

	if !tea.BoolValue(util.IsUnset(request.PoolKey)) {
		query["PoolKey"] = request.PoolKey
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SubsId)) {
		query["SubsId"] = request.SubsId
	}

	if !tea.BoolValue(util.IsUnset(request.TrackNo)) {
		query["trackNo"] = request.TrackNo
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddAxnTrackNo"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddAxnTrackNoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a tracking number for a private number in the AXN binding.
//
// Description:
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 5,000 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - AddAxnTrackNoRequest
//
// @return AddAxnTrackNoResponse
func (client *Client) AddAxnTrackNo(request *AddAxnTrackNoRequest) (_result *AddAxnTrackNoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddAxnTrackNoResponse{}
	_body, _err := client.AddAxnTrackNoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds a blacklist.
//
// Description:
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 100 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - AddSecretBlacklistRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddSecretBlacklistResponse
func (client *Client) AddSecretBlacklistWithOptions(request *AddSecretBlacklistRequest, runtime *util.RuntimeOptions) (_result *AddSecretBlacklistResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BlackNo)) {
		query["BlackNo"] = request.BlackNo
	}

	if !tea.BoolValue(util.IsUnset(request.BlackType)) {
		query["BlackType"] = request.BlackType
	}

	if !tea.BoolValue(util.IsUnset(request.PoolKey)) {
		query["PoolKey"] = request.PoolKey
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		query["Remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.WayControl)) {
		query["WayControl"] = request.WayControl
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddSecretBlacklist"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddSecretBlacklistResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a blacklist.
//
// Description:
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 100 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - AddSecretBlacklistRequest
//
// @return AddSecretBlacklistResponse
func (client *Client) AddSecretBlacklist(request *AddSecretBlacklistRequest) (_result *AddSecretBlacklistResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddSecretBlacklistResponse{}
	_body, _err := client.AddSecretBlacklistWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 调用本接口向工作号平台请求为员工B的工作号X建立呼叫绑定（B，X，A），允许B通过X呼叫客户A
//
// @param request - BindAXBCallRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindAXBCallResponse
func (client *Client) BindAXBCallWithOptions(request *BindAXBCallRequest, runtime *util.RuntimeOptions) (_result *BindAXBCallResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthId)) {
		query["AuthId"] = request.AuthId
	}

	if !tea.BoolValue(util.IsUnset(request.CallerParentId)) {
		query["CallerParentId"] = request.CallerParentId
	}

	if !tea.BoolValue(util.IsUnset(request.CustomerPoolKey)) {
		query["CustomerPoolKey"] = request.CustomerPoolKey
	}

	if !tea.BoolValue(util.IsUnset(request.Expiration)) {
		query["Expiration"] = request.Expiration
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ReqId)) {
		query["ReqId"] = request.ReqId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.TelA)) {
		query["TelA"] = request.TelA
	}

	if !tea.BoolValue(util.IsUnset(request.UserData)) {
		query["UserData"] = request.UserData
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BindAXBCall"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BindAXBCallResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 调用本接口向工作号平台请求为员工B的工作号X建立呼叫绑定（B，X，A），允许B通过X呼叫客户A
//
// @param request - BindAXBCallRequest
//
// @return BindAXBCallResponse
func (client *Client) BindAXBCall(request *BindAXBCallRequest) (_result *BindAXBCallResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BindAXBCallResponse{}
	_body, _err := client.BindAXBCallWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds an AXB binding.
//
// Description:
//
// Before you add an AXB binding, we recommend that you specify role A and role B in the AXB device certificate (ProductKey, DeviceName, and DeviceSecret) in your business scenario. For example, in a taxi-hailing scenario, role A is the passenger and role B is the driver.
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 5,000 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - BindAxbRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindAxbResponse
func (client *Client) BindAxbWithOptions(request *BindAxbRequest, runtime *util.RuntimeOptions) (_result *BindAxbResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ASRModelId)) {
		query["ASRModelId"] = request.ASRModelId
	}

	if !tea.BoolValue(util.IsUnset(request.ASRStatus)) {
		query["ASRStatus"] = request.ASRStatus
	}

	if !tea.BoolValue(util.IsUnset(request.CallDisplayType)) {
		query["CallDisplayType"] = request.CallDisplayType
	}

	if !tea.BoolValue(util.IsUnset(request.CallRestrict)) {
		query["CallRestrict"] = request.CallRestrict
	}

	if !tea.BoolValue(util.IsUnset(request.CallTimeout)) {
		query["CallTimeout"] = request.CallTimeout
	}

	if !tea.BoolValue(util.IsUnset(request.DtmfConfig)) {
		query["DtmfConfig"] = request.DtmfConfig
	}

	if !tea.BoolValue(util.IsUnset(request.ExpectCity)) {
		query["ExpectCity"] = request.ExpectCity
	}

	if !tea.BoolValue(util.IsUnset(request.Expiration)) {
		query["Expiration"] = request.Expiration
	}

	if !tea.BoolValue(util.IsUnset(request.IsRecordingEnabled)) {
		query["IsRecordingEnabled"] = request.IsRecordingEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.OutId)) {
		query["OutId"] = request.OutId
	}

	if !tea.BoolValue(util.IsUnset(request.OutOrderId)) {
		query["OutOrderId"] = request.OutOrderId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNoA)) {
		query["PhoneNoA"] = request.PhoneNoA
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNoB)) {
		query["PhoneNoB"] = request.PhoneNoB
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNoX)) {
		query["PhoneNoX"] = request.PhoneNoX
	}

	if !tea.BoolValue(util.IsUnset(request.PoolKey)) {
		query["PoolKey"] = request.PoolKey
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RingConfig)) {
		query["RingConfig"] = request.RingConfig
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BindAxb"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BindAxbResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds an AXB binding.
//
// Description:
//
// Before you add an AXB binding, we recommend that you specify role A and role B in the AXB device certificate (ProductKey, DeviceName, and DeviceSecret) in your business scenario. For example, in a taxi-hailing scenario, role A is the passenger and role B is the driver.
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 5,000 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - BindAxbRequest
//
// @return BindAxbResponse
func (client *Client) BindAxb(request *BindAxbRequest) (_result *BindAxbResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BindAxbResponse{}
	_body, _err := client.BindAxbWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 固话AxB绑定
//
// @param tmpReq - BindAxbFixedLineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindAxbFixedLineResponse
func (client *Client) BindAxbFixedLineWithOptions(tmpReq *BindAxbFixedLineRequest, runtime *util.RuntimeOptions) (_result *BindAxbFixedLineResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &BindAxbFixedLineShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Extra)) {
		request.ExtraShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Extra, tea.String("Extra"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Anucode)) {
		query["Anucode"] = request.Anucode
	}

	if !tea.BoolValue(util.IsUnset(request.Anucodecalled)) {
		query["Anucodecalled"] = request.Anucodecalled
	}

	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.Areacode)) {
		query["Areacode"] = request.Areacode
	}

	if !tea.BoolValue(util.IsUnset(request.BindType)) {
		query["BindType"] = request.BindType
	}

	if !tea.BoolValue(util.IsUnset(request.Expiration)) {
		query["Expiration"] = request.Expiration
	}

	if !tea.BoolValue(util.IsUnset(request.ExtraShrink)) {
		query["Extra"] = request.ExtraShrink
	}

	if !tea.BoolValue(util.IsUnset(request.OrderId)) {
		query["OrderId"] = request.OrderId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		query["Remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Subts)) {
		query["Subts"] = request.Subts
	}

	if !tea.BoolValue(util.IsUnset(request.TAnucodeConnect)) {
		query["TAnucodeConnect"] = request.TAnucodeConnect
	}

	if !tea.BoolValue(util.IsUnset(request.TelA)) {
		query["TelA"] = request.TelA
	}

	if !tea.BoolValue(util.IsUnset(request.TelB)) {
		query["TelB"] = request.TelB
	}

	if !tea.BoolValue(util.IsUnset(request.TelX)) {
		query["TelX"] = request.TelX
	}

	if !tea.BoolValue(util.IsUnset(request.Ts)) {
		query["Ts"] = request.Ts
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BindAxbFixedLine"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BindAxbFixedLineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 固话AxB绑定
//
// @param request - BindAxbFixedLineRequest
//
// @return BindAxbFixedLineResponse
func (client *Client) BindAxbFixedLine(request *BindAxbFixedLineRequest) (_result *BindAxbFixedLineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BindAxbFixedLineResponse{}
	_body, _err := client.BindAxbFixedLineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds an AXG binding.
//
// Description:
//
// An AXG protection solution can be configured to meet the requirements for grading users, limiting the scope of calls, and restricting order snatching. The letter G represents a phone number group to which you can add phone numbers as needed.
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 5,000 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - BindAxgRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindAxgResponse
func (client *Client) BindAxgWithOptions(request *BindAxgRequest, runtime *util.RuntimeOptions) (_result *BindAxgResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ASRModelId)) {
		query["ASRModelId"] = request.ASRModelId
	}

	if !tea.BoolValue(util.IsUnset(request.ASRStatus)) {
		query["ASRStatus"] = request.ASRStatus
	}

	if !tea.BoolValue(util.IsUnset(request.CallDisplayType)) {
		query["CallDisplayType"] = request.CallDisplayType
	}

	if !tea.BoolValue(util.IsUnset(request.CallRestrict)) {
		query["CallRestrict"] = request.CallRestrict
	}

	if !tea.BoolValue(util.IsUnset(request.ExpectCity)) {
		query["ExpectCity"] = request.ExpectCity
	}

	if !tea.BoolValue(util.IsUnset(request.Expiration)) {
		query["Expiration"] = request.Expiration
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.IsRecordingEnabled)) {
		query["IsRecordingEnabled"] = request.IsRecordingEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.OutId)) {
		query["OutId"] = request.OutId
	}

	if !tea.BoolValue(util.IsUnset(request.OutOrderId)) {
		query["OutOrderId"] = request.OutOrderId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNoA)) {
		query["PhoneNoA"] = request.PhoneNoA
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNoB)) {
		query["PhoneNoB"] = request.PhoneNoB
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNoX)) {
		query["PhoneNoX"] = request.PhoneNoX
	}

	if !tea.BoolValue(util.IsUnset(request.PoolKey)) {
		query["PoolKey"] = request.PoolKey
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RingConfig)) {
		query["RingConfig"] = request.RingConfig
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BindAxg"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BindAxgResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds an AXG binding.
//
// Description:
//
// An AXG protection solution can be configured to meet the requirements for grading users, limiting the scope of calls, and restricting order snatching. The letter G represents a phone number group to which you can add phone numbers as needed.
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 5,000 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - BindAxgRequest
//
// @return BindAxgResponse
func (client *Client) BindAxg(request *BindAxgRequest) (_result *BindAxgResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BindAxgResponse{}
	_body, _err := client.BindAxgWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds an AXN binding.
//
// Description:
//
// >  An AXN private number is a dedicated private number assigned to phone number A. When an N-side number is used to call phone number X, the call is forwarded to phone number A.
//
// @param request - BindAxnRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindAxnResponse
func (client *Client) BindAxnWithOptions(request *BindAxnRequest, runtime *util.RuntimeOptions) (_result *BindAxnResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ASRModelId)) {
		query["ASRModelId"] = request.ASRModelId
	}

	if !tea.BoolValue(util.IsUnset(request.ASRStatus)) {
		query["ASRStatus"] = request.ASRStatus
	}

	if !tea.BoolValue(util.IsUnset(request.CallDisplayType)) {
		query["CallDisplayType"] = request.CallDisplayType
	}

	if !tea.BoolValue(util.IsUnset(request.CallRestrict)) {
		query["CallRestrict"] = request.CallRestrict
	}

	if !tea.BoolValue(util.IsUnset(request.CallTimeout)) {
		query["CallTimeout"] = request.CallTimeout
	}

	if !tea.BoolValue(util.IsUnset(request.ExpectCity)) {
		query["ExpectCity"] = request.ExpectCity
	}

	if !tea.BoolValue(util.IsUnset(request.Expiration)) {
		query["Expiration"] = request.Expiration
	}

	if !tea.BoolValue(util.IsUnset(request.Extend)) {
		query["Extend"] = request.Extend
	}

	if !tea.BoolValue(util.IsUnset(request.IsRecordingEnabled)) {
		query["IsRecordingEnabled"] = request.IsRecordingEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.NoType)) {
		query["NoType"] = request.NoType
	}

	if !tea.BoolValue(util.IsUnset(request.OutId)) {
		query["OutId"] = request.OutId
	}

	if !tea.BoolValue(util.IsUnset(request.OutOrderId)) {
		query["OutOrderId"] = request.OutOrderId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNoA)) {
		query["PhoneNoA"] = request.PhoneNoA
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNoB)) {
		query["PhoneNoB"] = request.PhoneNoB
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNoX)) {
		query["PhoneNoX"] = request.PhoneNoX
	}

	if !tea.BoolValue(util.IsUnset(request.PoolKey)) {
		query["PoolKey"] = request.PoolKey
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RingConfig)) {
		query["RingConfig"] = request.RingConfig
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BindAxn"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BindAxnResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds an AXN binding.
//
// Description:
//
// >  An AXN private number is a dedicated private number assigned to phone number A. When an N-side number is used to call phone number X, the call is forwarded to phone number A.
//
// @param request - BindAxnRequest
//
// @return BindAxnResponse
func (client *Client) BindAxn(request *BindAxnRequest) (_result *BindAxnResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BindAxnResponse{}
	_body, _err := client.BindAxnWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds an AXN extension binding.
//
// Description:
//
// Before you add an AXN extension binding, confirm phone number A and phone number N in the business scenario. Phone number A belongs to a customer, and phone number X is the private number assigned to the customer. When any other phone number is used to call phone number X and the extension, the call is transferred to phone number A. When phone number A is used to call phone number X, the call is transferred to the default phone number B that is specified during the phone number binding.
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 5,000 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - BindAxnExtensionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindAxnExtensionResponse
func (client *Client) BindAxnExtensionWithOptions(request *BindAxnExtensionRequest, runtime *util.RuntimeOptions) (_result *BindAxnExtensionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ASRModelId)) {
		query["ASRModelId"] = request.ASRModelId
	}

	if !tea.BoolValue(util.IsUnset(request.ASRStatus)) {
		query["ASRStatus"] = request.ASRStatus
	}

	if !tea.BoolValue(util.IsUnset(request.CallDisplayType)) {
		query["CallDisplayType"] = request.CallDisplayType
	}

	if !tea.BoolValue(util.IsUnset(request.CallRestrict)) {
		query["CallRestrict"] = request.CallRestrict
	}

	if !tea.BoolValue(util.IsUnset(request.ExpectCity)) {
		query["ExpectCity"] = request.ExpectCity
	}

	if !tea.BoolValue(util.IsUnset(request.Expiration)) {
		query["Expiration"] = request.Expiration
	}

	if !tea.BoolValue(util.IsUnset(request.Extend)) {
		query["Extend"] = request.Extend
	}

	if !tea.BoolValue(util.IsUnset(request.Extension)) {
		query["Extension"] = request.Extension
	}

	if !tea.BoolValue(util.IsUnset(request.IsRecordingEnabled)) {
		query["IsRecordingEnabled"] = request.IsRecordingEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.OutId)) {
		query["OutId"] = request.OutId
	}

	if !tea.BoolValue(util.IsUnset(request.OutOrderId)) {
		query["OutOrderId"] = request.OutOrderId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNoA)) {
		query["PhoneNoA"] = request.PhoneNoA
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNoB)) {
		query["PhoneNoB"] = request.PhoneNoB
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNoX)) {
		query["PhoneNoX"] = request.PhoneNoX
	}

	if !tea.BoolValue(util.IsUnset(request.PoolKey)) {
		query["PoolKey"] = request.PoolKey
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RingConfig)) {
		query["RingConfig"] = request.RingConfig
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BindAxnExtension"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BindAxnExtensionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds an AXN extension binding.
//
// Description:
//
// Before you add an AXN extension binding, confirm phone number A and phone number N in the business scenario. Phone number A belongs to a customer, and phone number X is the private number assigned to the customer. When any other phone number is used to call phone number X and the extension, the call is transferred to phone number A. When phone number A is used to call phone number X, the call is transferred to the default phone number B that is specified during the phone number binding.
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 5,000 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - BindAxnExtensionRequest
//
// @return BindAxnExtensionResponse
func (client *Client) BindAxnExtension(request *BindAxnExtensionRequest) (_result *BindAxnExtensionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BindAxnExtensionResponse{}
	_body, _err := client.BindAxnExtensionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # AXN分机号-号码绑定
//
// @param tmpReq - BindAxnExtensionFixedLineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindAxnExtensionFixedLineResponse
func (client *Client) BindAxnExtensionFixedLineWithOptions(tmpReq *BindAxnExtensionFixedLineRequest, runtime *util.RuntimeOptions) (_result *BindAxnExtensionFixedLineResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &BindAxnExtensionFixedLineShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Extraaxx)) {
		request.ExtraaxxShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Extraaxx, tea.String("Extraaxx"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Anucode)) {
		query["Anucode"] = request.Anucode
	}

	if !tea.BoolValue(util.IsUnset(request.Anucodecalled)) {
		query["Anucodecalled"] = request.Anucodecalled
	}

	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.Areacode)) {
		query["Areacode"] = request.Areacode
	}

	if !tea.BoolValue(util.IsUnset(request.BindType)) {
		query["BindType"] = request.BindType
	}

	if !tea.BoolValue(util.IsUnset(request.Expiration)) {
		query["Expiration"] = request.Expiration
	}

	if !tea.BoolValue(util.IsUnset(request.ExtraaxxShrink)) {
		query["Extraaxx"] = request.ExtraaxxShrink
	}

	if !tea.BoolValue(util.IsUnset(request.OrderId)) {
		query["OrderId"] = request.OrderId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		query["Remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Subts)) {
		query["Subts"] = request.Subts
	}

	if !tea.BoolValue(util.IsUnset(request.TAnucodeConnect)) {
		query["TAnucodeConnect"] = request.TAnucodeConnect
	}

	if !tea.BoolValue(util.IsUnset(request.TelA)) {
		query["TelA"] = request.TelA
	}

	if !tea.BoolValue(util.IsUnset(request.TelX)) {
		query["TelX"] = request.TelX
	}

	if !tea.BoolValue(util.IsUnset(request.TelXext)) {
		query["TelXext"] = request.TelXext
	}

	if !tea.BoolValue(util.IsUnset(request.Ts)) {
		query["Ts"] = request.Ts
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BindAxnExtensionFixedLine"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BindAxnExtensionFixedLineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # AXN分机号-号码绑定
//
// @param request - BindAxnExtensionFixedLineRequest
//
// @return BindAxnExtensionFixedLineResponse
func (client *Client) BindAxnExtensionFixedLine(request *BindAxnExtensionFixedLineRequest) (_result *BindAxnExtensionFixedLineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BindAxnExtensionFixedLineResponse{}
	_body, _err := client.BindAxnExtensionFixedLineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # AXN模式绑定，分配X号码
//
// @param tmpReq - BindAxnFixedLineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindAxnFixedLineResponse
func (client *Client) BindAxnFixedLineWithOptions(tmpReq *BindAxnFixedLineRequest, runtime *util.RuntimeOptions) (_result *BindAxnFixedLineResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &BindAxnFixedLineShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Extra)) {
		request.ExtraShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Extra, tea.String("Extra"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Anucode)) {
		query["Anucode"] = request.Anucode
	}

	if !tea.BoolValue(util.IsUnset(request.Anucodecalled)) {
		query["Anucodecalled"] = request.Anucodecalled
	}

	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.Areacode)) {
		query["Areacode"] = request.Areacode
	}

	if !tea.BoolValue(util.IsUnset(request.BindType)) {
		query["BindType"] = request.BindType
	}

	if !tea.BoolValue(util.IsUnset(request.Expiration)) {
		query["Expiration"] = request.Expiration
	}

	if !tea.BoolValue(util.IsUnset(request.ExtraShrink)) {
		query["Extra"] = request.ExtraShrink
	}

	if !tea.BoolValue(util.IsUnset(request.OrderId)) {
		query["OrderId"] = request.OrderId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		query["Remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Subts)) {
		query["Subts"] = request.Subts
	}

	if !tea.BoolValue(util.IsUnset(request.TAnucodeConnect)) {
		query["TAnucodeConnect"] = request.TAnucodeConnect
	}

	if !tea.BoolValue(util.IsUnset(request.TelA)) {
		query["TelA"] = request.TelA
	}

	if !tea.BoolValue(util.IsUnset(request.TelB)) {
		query["TelB"] = request.TelB
	}

	if !tea.BoolValue(util.IsUnset(request.TelX)) {
		query["TelX"] = request.TelX
	}

	if !tea.BoolValue(util.IsUnset(request.Ts)) {
		query["Ts"] = request.Ts
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BindAxnFixedLine"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BindAxnFixedLineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # AXN模式绑定，分配X号码
//
// @param request - BindAxnFixedLineRequest
//
// @return BindAxnFixedLineResponse
func (client *Client) BindAxnFixedLine(request *BindAxnFixedLineRequest) (_result *BindAxnFixedLineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BindAxnFixedLineResponse{}
	_body, _err := client.BindAxnFixedLineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param tmpReq - BindBatchAxgRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindBatchAxgResponse
func (client *Client) BindBatchAxgWithOptions(tmpReq *BindBatchAxgRequest, runtime *util.RuntimeOptions) (_result *BindBatchAxgResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &BindBatchAxgShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.AxgBindList)) {
		request.AxgBindListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AxgBindList, tea.String("AxgBindList"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AxgBindListShrink)) {
		query["AxgBindList"] = request.AxgBindListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PoolKey)) {
		query["PoolKey"] = request.PoolKey
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BindBatchAxg"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BindBatchAxgResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - BindBatchAxgRequest
//
// @return BindBatchAxgResponse
func (client *Client) BindBatchAxg(request *BindBatchAxgRequest) (_result *BindBatchAxgResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BindBatchAxgResponse{}
	_body, _err := client.BindBatchAxgWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 平台指定工作号X 和员工号B建立关联，完成X 实名认证，绑定生效后，所有X 的呼叫都会转接到B
//
// @param request - BindXBRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindXBResponse
func (client *Client) BindXBWithOptions(request *BindXBRequest, runtime *util.RuntimeOptions) (_result *BindXBResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CallerParentId)) {
		query["CallerParentId"] = request.CallerParentId
	}

	if !tea.BoolValue(util.IsUnset(request.CustomerPoolKey)) {
		query["CustomerPoolKey"] = request.CustomerPoolKey
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ReqId)) {
		query["ReqId"] = request.ReqId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.TelB)) {
		query["TelB"] = request.TelB
	}

	if !tea.BoolValue(util.IsUnset(request.TelX)) {
		query["TelX"] = request.TelX
	}

	if !tea.BoolValue(util.IsUnset(request.UserData)) {
		query["UserData"] = request.UserData
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BindXB"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BindXBResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 平台指定工作号X 和员工号B建立关联，完成X 实名认证，绑定生效后，所有X 的呼叫都会转接到B
//
// @param request - BindXBRequest
//
// @return BindXBResponse
func (client *Client) BindXB(request *BindXBRequest) (_result *BindXBResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BindXBResponse{}
	_body, _err := client.BindXBWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Purchases a phone number.
//
// Description:
//
//	  After you create a phone number pool in the Phone Number Protection console, the phone number pool is empty by default. You must purchase phone numbers and add them to the phone number pool.
//
//		- Before you call this operation, make sure that you are familiar with the [pricing](https://help.aliyun.com/document_detail/59825.html) of Phone Number Protection.
//
//		- When purchasing a phone number, specify the home location. If no sufficient phone numbers are available for purchase in the home location, the purchase of the phone number fails. Before you call this operation to purchase a phone number, check the quantity of phone numbers available for purchase by using the [QuerySecretNoRemain](https://help.aliyun.com/document_detail/111699.html) operation.
//
//		- The account used to purchase a phone number must be an enterprise account that has passed real-name verification. For more information about how to perform real-name verification, see [Enterprise verification FAQs](https://help.aliyun.com/document_detail/37172.html).
//
// @param request - BuySecretNoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BuySecretNoResponse
func (client *Client) BuySecretNoWithOptions(request *BuySecretNoRequest, runtime *util.RuntimeOptions) (_result *BuySecretNoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.City)) {
		query["City"] = request.City
	}

	if !tea.BoolValue(util.IsUnset(request.DisplayPool)) {
		query["DisplayPool"] = request.DisplayPool
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PoolKey)) {
		query["PoolKey"] = request.PoolKey
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecretNo)) {
		query["SecretNo"] = request.SecretNo
	}

	if !tea.BoolValue(util.IsUnset(request.SpecId)) {
		query["SpecId"] = request.SpecId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BuySecretNo"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BuySecretNoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Purchases a phone number.
//
// Description:
//
//	  After you create a phone number pool in the Phone Number Protection console, the phone number pool is empty by default. You must purchase phone numbers and add them to the phone number pool.
//
//		- Before you call this operation, make sure that you are familiar with the [pricing](https://help.aliyun.com/document_detail/59825.html) of Phone Number Protection.
//
//		- When purchasing a phone number, specify the home location. If no sufficient phone numbers are available for purchase in the home location, the purchase of the phone number fails. Before you call this operation to purchase a phone number, check the quantity of phone numbers available for purchase by using the [QuerySecretNoRemain](https://help.aliyun.com/document_detail/111699.html) operation.
//
//		- The account used to purchase a phone number must be an enterprise account that has passed real-name verification. For more information about how to perform real-name verification, see [Enterprise verification FAQs](https://help.aliyun.com/document_detail/37172.html).
//
// @param request - BuySecretNoRequest
//
// @return BuySecretNoResponse
func (client *Client) BuySecretNo(request *BuySecretNoRequest) (_result *BuySecretNoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BuySecretNoResponse{}
	_body, _err := client.BuySecretNoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Cancels a door-to-door delivery order.
//
// Description:
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 100 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - CancelPickUpWaybillRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelPickUpWaybillResponse
func (client *Client) CancelPickUpWaybillWithOptions(request *CancelPickUpWaybillRequest, runtime *util.RuntimeOptions) (_result *CancelPickUpWaybillResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CancelDesc)) {
		query["CancelDesc"] = request.CancelDesc
	}

	if !tea.BoolValue(util.IsUnset(request.OuterOrderCode)) {
		query["OuterOrderCode"] = request.OuterOrderCode
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CancelPickUpWaybill"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelPickUpWaybillResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancels a door-to-door delivery order.
//
// Description:
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 100 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - CancelPickUpWaybillRequest
//
// @return CancelPickUpWaybillResponse
func (client *Client) CancelPickUpWaybill(request *CancelPickUpWaybillRequest) (_result *CancelPickUpWaybillResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CancelPickUpWaybillResponse{}
	_body, _err := client.CancelPickUpWaybillWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 配置X号码，单独对工作号的话音呼叫、企业名片等通信功能进行配置操作
//
// @param tmpReq - ConfigXRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfigXResponse
func (client *Client) ConfigXWithOptions(tmpReq *ConfigXRequest, runtime *util.RuntimeOptions) (_result *ConfigXResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ConfigXShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.SequenceCalls)) {
		request.SequenceCallsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SequenceCalls, tea.String("SequenceCalls"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CallAbility)) {
		query["CallAbility"] = request.CallAbility
	}

	if !tea.BoolValue(util.IsUnset(request.CallerParentId)) {
		query["CallerParentId"] = request.CallerParentId
	}

	if !tea.BoolValue(util.IsUnset(request.CustomerPoolKey)) {
		query["CustomerPoolKey"] = request.CustomerPoolKey
	}

	if !tea.BoolValue(util.IsUnset(request.GNFlag)) {
		query["GNFlag"] = request.GNFlag
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ReqId)) {
		query["ReqId"] = request.ReqId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SequenceCallsShrink)) {
		query["SequenceCalls"] = request.SequenceCallsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.SequenceMode)) {
		query["SequenceMode"] = request.SequenceMode
	}

	if !tea.BoolValue(util.IsUnset(request.SmsAbility)) {
		query["SmsAbility"] = request.SmsAbility
	}

	if !tea.BoolValue(util.IsUnset(request.SmsSignMode)) {
		query["SmsSignMode"] = request.SmsSignMode
	}

	if !tea.BoolValue(util.IsUnset(request.TelX)) {
		query["TelX"] = request.TelX
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ConfigX"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ConfigXResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 配置X号码，单独对工作号的话音呼叫、企业名片等通信功能进行配置操作
//
// @param request - ConfigXRequest
//
// @return ConfigXResponse
func (client *Client) ConfigX(request *ConfigXRequest) (_result *ConfigXResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ConfigXResponse{}
	_body, _err := client.ConfigXWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates number group G.
//
// Description:
//
// Before you add an AXG binding, you must purchase phone number X, create number group G, and then add phone numbers to number group G. If you do not add phone numbers to number group G after you create number group G, you can call the [OperateAxgGroup](https://help.aliyun.com/document_detail/110252.htm) operation to add phone numbers to number group G.
//
// >  Up to 2,000 number groups G can be added for a single phone number pool.
//
// @param request - CreateAxgGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAxgGroupResponse
func (client *Client) CreateAxgGroupWithOptions(request *CreateAxgGroupRequest, runtime *util.RuntimeOptions) (_result *CreateAxgGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Numbers)) {
		query["Numbers"] = request.Numbers
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PoolKey)) {
		query["PoolKey"] = request.PoolKey
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		query["Remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAxgGroup"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAxgGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates number group G.
//
// Description:
//
// Before you add an AXG binding, you must purchase phone number X, create number group G, and then add phone numbers to number group G. If you do not add phone numbers to number group G after you create number group G, you can call the [OperateAxgGroup](https://help.aliyun.com/document_detail/110252.htm) operation to add phone numbers to number group G.
//
// >  Up to 2,000 number groups G can be added for a single phone number pool.
//
// @param request - CreateAxgGroupRequest
//
// @return CreateAxgGroupResponse
func (client *Client) CreateAxgGroup(request *CreateAxgGroupRequest) (_result *CreateAxgGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAxgGroupResponse{}
	_body, _err := client.CreateAxgGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 通过API收集小号a号码固话
//
// @param request - CreateFixedNoAReportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFixedNoAReportResponse
func (client *Client) CreateFixedNoAReportWithOptions(request *CreateFixedNoAReportRequest, runtime *util.RuntimeOptions) (_result *CreateFixedNoAReportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ANoWhiteGroupId)) {
		query["ANoWhiteGroupId"] = request.ANoWhiteGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.CustName)) {
		query["CustName"] = request.CustName
	}

	if !tea.BoolValue(util.IsUnset(request.CustPhoneNo)) {
		query["CustPhoneNo"] = request.CustPhoneNo
	}

	if !tea.BoolValue(util.IsUnset(request.CustType)) {
		query["CustType"] = request.CustType
	}

	if !tea.BoolValue(util.IsUnset(request.DocumentNumber)) {
		query["DocumentNumber"] = request.DocumentNumber
	}

	if !tea.BoolValue(util.IsUnset(request.DocumentType)) {
		query["DocumentType"] = request.DocumentType
	}

	if !tea.BoolValue(util.IsUnset(request.FixedLineWorkId)) {
		query["FixedLineWorkId"] = request.FixedLineWorkId
	}

	if !tea.BoolValue(util.IsUnset(request.FixedNoA)) {
		query["FixedNoA"] = request.FixedNoA
	}

	if !tea.BoolValue(util.IsUnset(request.IdCardAlivePhoto)) {
		query["IdCardAlivePhoto"] = request.IdCardAlivePhoto
	}

	if !tea.BoolValue(util.IsUnset(request.IdCardBackPhoto)) {
		query["IdCardBackPhoto"] = request.IdCardBackPhoto
	}

	if !tea.BoolValue(util.IsUnset(request.IdCardFrontPhoto)) {
		query["IdCardFrontPhoto"] = request.IdCardFrontPhoto
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		query["Remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateFixedNoAReport"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateFixedNoAReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通过API收集小号a号码固话
//
// @param request - CreateFixedNoAReportRequest
//
// @return CreateFixedNoAReportResponse
func (client *Client) CreateFixedNoAReport(request *CreateFixedNoAReportRequest) (_result *CreateFixedNoAReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateFixedNoAReportResponse{}
	_body, _err := client.CreateFixedNoAReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 通过API收集小号a号码手机号
//
// @param request - CreatePhoneNoAReportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePhoneNoAReportResponse
func (client *Client) CreatePhoneNoAReportWithOptions(request *CreatePhoneNoAReportRequest, runtime *util.RuntimeOptions) (_result *CreatePhoneNoAReportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ANoWhiteGroupId)) {
		query["ANoWhiteGroupId"] = request.ANoWhiteGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.CustName)) {
		query["CustName"] = request.CustName
	}

	if !tea.BoolValue(util.IsUnset(request.DocumentNumber)) {
		query["DocumentNumber"] = request.DocumentNumber
	}

	if !tea.BoolValue(util.IsUnset(request.DocumentType)) {
		query["DocumentType"] = request.DocumentType
	}

	if !tea.BoolValue(util.IsUnset(request.IdCardAlivePhoto)) {
		query["IdCardAlivePhoto"] = request.IdCardAlivePhoto
	}

	if !tea.BoolValue(util.IsUnset(request.IdCardBackPhoto)) {
		query["IdCardBackPhoto"] = request.IdCardBackPhoto
	}

	if !tea.BoolValue(util.IsUnset(request.IdCardFrontPhoto)) {
		query["IdCardFrontPhoto"] = request.IdCardFrontPhoto
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNoA)) {
		query["PhoneNoA"] = request.PhoneNoA
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		query["Remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreatePhoneNoAReport"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreatePhoneNoAReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通过API收集小号a号码手机号
//
// @param request - CreatePhoneNoAReportRequest
//
// @return CreatePhoneNoAReportResponse
func (client *Client) CreatePhoneNoAReport(request *CreatePhoneNoAReportRequest) (_result *CreatePhoneNoAReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreatePhoneNoAReportResponse{}
	_body, _err := client.CreatePhoneNoAReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a door-to-door delivery order.
//
// Description:
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 100 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param tmpReq - CreatePickUpWaybillRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePickUpWaybillResponse
func (client *Client) CreatePickUpWaybillWithOptions(tmpReq *CreatePickUpWaybillRequest, runtime *util.RuntimeOptions) (_result *CreatePickUpWaybillResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreatePickUpWaybillShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ConsigneeAddress)) {
		request.ConsigneeAddressShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ConsigneeAddress, tea.String("ConsigneeAddress"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.GoodsInfos)) {
		request.GoodsInfosShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.GoodsInfos, tea.String("GoodsInfos"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.SendAddress)) {
		request.SendAddressShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SendAddress, tea.String("SendAddress"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppointGotEndTime)) {
		query["AppointGotEndTime"] = request.AppointGotEndTime
	}

	if !tea.BoolValue(util.IsUnset(request.AppointGotStartTime)) {
		query["AppointGotStartTime"] = request.AppointGotStartTime
	}

	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		query["BizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.ConsigneeAddressShrink)) {
		query["ConsigneeAddress"] = request.ConsigneeAddressShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ConsigneeMobile)) {
		query["ConsigneeMobile"] = request.ConsigneeMobile
	}

	if !tea.BoolValue(util.IsUnset(request.ConsigneeName)) {
		query["ConsigneeName"] = request.ConsigneeName
	}

	if !tea.BoolValue(util.IsUnset(request.ConsigneePhone)) {
		query["ConsigneePhone"] = request.ConsigneePhone
	}

	if !tea.BoolValue(util.IsUnset(request.CpCode)) {
		query["CpCode"] = request.CpCode
	}

	if !tea.BoolValue(util.IsUnset(request.GoodsInfosShrink)) {
		query["GoodsInfos"] = request.GoodsInfosShrink
	}

	if !tea.BoolValue(util.IsUnset(request.OrderChannels)) {
		query["OrderChannels"] = request.OrderChannels
	}

	if !tea.BoolValue(util.IsUnset(request.OuterOrderCode)) {
		query["OuterOrderCode"] = request.OuterOrderCode
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		query["Remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.SendAddressShrink)) {
		query["SendAddress"] = request.SendAddressShrink
	}

	if !tea.BoolValue(util.IsUnset(request.SendMobile)) {
		query["SendMobile"] = request.SendMobile
	}

	if !tea.BoolValue(util.IsUnset(request.SendName)) {
		query["SendName"] = request.SendName
	}

	if !tea.BoolValue(util.IsUnset(request.SendPhone)) {
		query["SendPhone"] = request.SendPhone
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreatePickUpWaybill"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreatePickUpWaybillResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a door-to-door delivery order.
//
// Description:
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 100 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - CreatePickUpWaybillRequest
//
// @return CreatePickUpWaybillResponse
func (client *Client) CreatePickUpWaybill(request *CreatePickUpWaybillRequest) (_result *CreatePickUpWaybillResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreatePickUpWaybillResponse{}
	_body, _err := client.CreatePickUpWaybillWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a pickup order.
//
// Description:
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 100 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param tmpReq - CreatePickUpWaybillPreQueryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePickUpWaybillPreQueryResponse
func (client *Client) CreatePickUpWaybillPreQueryWithOptions(tmpReq *CreatePickUpWaybillPreQueryRequest, runtime *util.RuntimeOptions) (_result *CreatePickUpWaybillPreQueryResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreatePickUpWaybillPreQueryShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ConsigneeInfo)) {
		request.ConsigneeInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ConsigneeInfo, tea.String("ConsigneeInfo"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.SenderInfo)) {
		request.SenderInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SenderInfo, tea.String("SenderInfo"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConsigneeInfoShrink)) {
		query["ConsigneeInfo"] = request.ConsigneeInfoShrink
	}

	if !tea.BoolValue(util.IsUnset(request.CpCode)) {
		query["CpCode"] = request.CpCode
	}

	if !tea.BoolValue(util.IsUnset(request.OrderChannels)) {
		query["OrderChannels"] = request.OrderChannels
	}

	if !tea.BoolValue(util.IsUnset(request.OuterOrderCode)) {
		query["OuterOrderCode"] = request.OuterOrderCode
	}

	if !tea.BoolValue(util.IsUnset(request.PreWeight)) {
		query["PreWeight"] = request.PreWeight
	}

	if !tea.BoolValue(util.IsUnset(request.SenderInfoShrink)) {
		query["SenderInfo"] = request.SenderInfoShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreatePickUpWaybillPreQuery"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreatePickUpWaybillPreQueryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a pickup order.
//
// Description:
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 100 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - CreatePickUpWaybillPreQueryRequest
//
// @return CreatePickUpWaybillPreQueryResponse
func (client *Client) CreatePickUpWaybillPreQuery(request *CreatePickUpWaybillPreQueryRequest) (_result *CreatePickUpWaybillPreQueryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreatePickUpWaybillPreQueryResponse{}
	_body, _err := client.CreatePickUpWaybillPreQueryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// B向A 发短信，客户端获取“短信标签”，尾部添加“标签”。通过“标签”解析被叫A，发短信到A。
//
// @param request - CreateSmsSignRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSmsSignResponse
func (client *Client) CreateSmsSignWithOptions(request *CreateSmsSignRequest, runtime *util.RuntimeOptions) (_result *CreateSmsSignResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CalledNo)) {
		query["CalledNo"] = request.CalledNo
	}

	if !tea.BoolValue(util.IsUnset(request.CallerParentId)) {
		query["CallerParentId"] = request.CallerParentId
	}

	if !tea.BoolValue(util.IsUnset(request.CallingNo)) {
		query["CallingNo"] = request.CallingNo
	}

	if !tea.BoolValue(util.IsUnset(request.CustomerPoolKey)) {
		query["CustomerPoolKey"] = request.CustomerPoolKey
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ReqId)) {
		query["ReqId"] = request.ReqId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSmsSign"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSmsSignResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// B向A 发短信，客户端获取“短信标签”，尾部添加“标签”。通过“标签”解析被叫A，发短信到A。
//
// @param request - CreateSmsSignRequest
//
// @return CreateSmsSignResponse
func (client *Client) CreateSmsSign(request *CreateSmsSignRequest) (_result *CreateSmsSignResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSmsSignResponse{}
	_body, _err := client.CreateSmsSignWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 固话AxB解绑
//
// @param request - DeleteAxbBindFixedLineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAxbBindFixedLineResponse
func (client *Client) DeleteAxbBindFixedLineWithOptions(request *DeleteAxbBindFixedLineRequest, runtime *util.RuntimeOptions) (_result *DeleteAxbBindFixedLineResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.OrderId)) {
		query["OrderId"] = request.OrderId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SubId)) {
		query["SubId"] = request.SubId
	}

	if !tea.BoolValue(util.IsUnset(request.Ts)) {
		query["Ts"] = request.Ts
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAxbBindFixedLine"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAxbBindFixedLineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 固话AxB解绑
//
// @param request - DeleteAxbBindFixedLineRequest
//
// @return DeleteAxbBindFixedLineResponse
func (client *Client) DeleteAxbBindFixedLine(request *DeleteAxbBindFixedLineRequest) (_result *DeleteAxbBindFixedLineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAxbBindFixedLineResponse{}
	_body, _err := client.DeleteAxbBindFixedLineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeleteAxgGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAxgGroupResponse
func (client *Client) DeleteAxgGroupWithOptions(request *DeleteAxgGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteAxgGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PoolKey)) {
		query["PoolKey"] = request.PoolKey
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAxgGroup"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAxgGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteAxgGroupRequest
//
// @return DeleteAxgGroupResponse
func (client *Client) DeleteAxgGroup(request *DeleteAxgGroupRequest) (_result *DeleteAxgGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAxgGroupResponse{}
	_body, _err := client.DeleteAxgGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 解绑已有Axn绑定
//
// @param request - DeleteAxnBindFixedLineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAxnBindFixedLineResponse
func (client *Client) DeleteAxnBindFixedLineWithOptions(request *DeleteAxnBindFixedLineRequest, runtime *util.RuntimeOptions) (_result *DeleteAxnBindFixedLineResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.OrderId)) {
		query["OrderId"] = request.OrderId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SubId)) {
		query["SubId"] = request.SubId
	}

	if !tea.BoolValue(util.IsUnset(request.Ts)) {
		query["Ts"] = request.Ts
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAxnBindFixedLine"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAxnBindFixedLineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 解绑已有Axn绑定
//
// @param request - DeleteAxnBindFixedLineRequest
//
// @return DeleteAxnBindFixedLineResponse
func (client *Client) DeleteAxnBindFixedLine(request *DeleteAxnBindFixedLineRequest) (_result *DeleteAxnBindFixedLineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAxnBindFixedLineResponse{}
	_body, _err := client.DeleteAxnBindFixedLineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 解绑已有AXN分机号绑定
//
// @param request - DeleteAxnExtensionBindFixedLineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAxnExtensionBindFixedLineResponse
func (client *Client) DeleteAxnExtensionBindFixedLineWithOptions(request *DeleteAxnExtensionBindFixedLineRequest, runtime *util.RuntimeOptions) (_result *DeleteAxnExtensionBindFixedLineResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.OrderId)) {
		query["OrderId"] = request.OrderId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SubId)) {
		query["SubId"] = request.SubId
	}

	if !tea.BoolValue(util.IsUnset(request.Ts)) {
		query["Ts"] = request.Ts
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAxnExtensionBindFixedLine"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAxnExtensionBindFixedLineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 解绑已有AXN分机号绑定
//
// @param request - DeleteAxnExtensionBindFixedLineRequest
//
// @return DeleteAxnExtensionBindFixedLineResponse
func (client *Client) DeleteAxnExtensionBindFixedLine(request *DeleteAxnExtensionBindFixedLineRequest) (_result *DeleteAxnExtensionBindFixedLineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAxnExtensionBindFixedLineResponse{}
	_body, _err := client.DeleteAxnExtensionBindFixedLineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # A号码报备数据删除
//
// @param request - DeleteSecretAPhoneNoToCustRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSecretAPhoneNoToCustResponse
func (client *Client) DeleteSecretAPhoneNoToCustWithOptions(request *DeleteSecretAPhoneNoToCustRequest, runtime *util.RuntimeOptions) (_result *DeleteSecretAPhoneNoToCustResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ANoWhiteGroupId)) {
		query["ANoWhiteGroupId"] = request.ANoWhiteGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNoA)) {
		query["PhoneNoA"] = request.PhoneNoA
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteSecretAPhoneNoToCust"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteSecretAPhoneNoToCustResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # A号码报备数据删除
//
// @param request - DeleteSecretAPhoneNoToCustRequest
//
// @return DeleteSecretAPhoneNoToCustResponse
func (client *Client) DeleteSecretAPhoneNoToCust(request *DeleteSecretAPhoneNoToCustRequest) (_result *DeleteSecretAPhoneNoToCustResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteSecretAPhoneNoToCustResponse{}
	_body, _err := client.DeleteSecretAPhoneNoToCustWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a blacklist.
//
// Description:
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 100 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DeleteSecretBlacklistRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSecretBlacklistResponse
func (client *Client) DeleteSecretBlacklistWithOptions(request *DeleteSecretBlacklistRequest, runtime *util.RuntimeOptions) (_result *DeleteSecretBlacklistResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BlackNo)) {
		query["BlackNo"] = request.BlackNo
	}

	if !tea.BoolValue(util.IsUnset(request.BlackType)) {
		query["BlackType"] = request.BlackType
	}

	if !tea.BoolValue(util.IsUnset(request.PoolKey)) {
		query["PoolKey"] = request.PoolKey
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		query["Remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.WayControl)) {
		query["WayControl"] = request.WayControl
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteSecretBlacklist"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteSecretBlacklistResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a blacklist.
//
// Description:
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 100 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DeleteSecretBlacklistRequest
//
// @return DeleteSecretBlacklistResponse
func (client *Client) DeleteSecretBlacklist(request *DeleteSecretBlacklistRequest) (_result *DeleteSecretBlacklistResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteSecretBlacklistResponse{}
	_body, _err := client.DeleteSecretBlacklistWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 隐私号上传文件，获取 OSS 信息
//
// @param request - GetDyplsOSSInfoForUploadFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDyplsOSSInfoForUploadFileResponse
func (client *Client) GetDyplsOSSInfoForUploadFileWithOptions(request *GetDyplsOSSInfoForUploadFileRequest, runtime *util.RuntimeOptions) (_result *GetDyplsOSSInfoForUploadFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		query["BizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDyplsOSSInfoForUploadFile"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDyplsOSSInfoForUploadFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 隐私号上传文件，获取 OSS 信息
//
// @param request - GetDyplsOSSInfoForUploadFileRequest
//
// @return GetDyplsOSSInfoForUploadFileResponse
func (client *Client) GetDyplsOSSInfoForUploadFile(request *GetDyplsOSSInfoForUploadFileRequest) (_result *GetDyplsOSSInfoForUploadFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDyplsOSSInfoForUploadFileResponse{}
	_body, _err := client.GetDyplsOSSInfoForUploadFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the details of the automatic speech recognition (ASR) result.
//
// Description:
//
// Before you call the GetSecretAsrDetail operation, set the ASRStatus parameter to true in the [BindAxn operation](https://help.aliyun.com/document_detail/400483.html). This ensures that you can obtain the ASR result properly.
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 100 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - GetSecretAsrDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSecretAsrDetailResponse
func (client *Client) GetSecretAsrDetailWithOptions(request *GetSecretAsrDetailRequest, runtime *util.RuntimeOptions) (_result *GetSecretAsrDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CallId)) {
		query["CallId"] = request.CallId
	}

	if !tea.BoolValue(util.IsUnset(request.CallTime)) {
		query["CallTime"] = request.CallTime
	}

	if !tea.BoolValue(util.IsUnset(request.PoolKey)) {
		query["PoolKey"] = request.PoolKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSecretAsrDetail"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSecretAsrDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the details of the automatic speech recognition (ASR) result.
//
// Description:
//
// Before you call the GetSecretAsrDetail operation, set the ASRStatus parameter to true in the [BindAxn operation](https://help.aliyun.com/document_detail/400483.html). This ensures that you can obtain the ASR result properly.
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 100 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - GetSecretAsrDetailRequest
//
// @return GetSecretAsrDetailResponse
func (client *Client) GetSecretAsrDetail(request *GetSecretAsrDetailRequest) (_result *GetSecretAsrDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSecretAsrDetailResponse{}
	_body, _err := client.GetSecretAsrDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the download URL of a recorded ringing tone.
//
// Description:
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 1,000 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - GetTotalPublicUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTotalPublicUrlResponse
func (client *Client) GetTotalPublicUrlWithOptions(request *GetTotalPublicUrlRequest, runtime *util.RuntimeOptions) (_result *GetTotalPublicUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CallId)) {
		query["CallId"] = request.CallId
	}

	if !tea.BoolValue(util.IsUnset(request.CallTime)) {
		query["CallTime"] = request.CallTime
	}

	if !tea.BoolValue(util.IsUnset(request.CheckSubs)) {
		query["CheckSubs"] = request.CheckSubs
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PartnerKey)) {
		query["PartnerKey"] = request.PartnerKey
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTotalPublicUrl"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTotalPublicUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the download URL of a recorded ringing tone.
//
// Description:
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 1,000 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - GetTotalPublicUrlRequest
//
// @return GetTotalPublicUrlResponse
func (client *Client) GetTotalPublicUrl(request *GetTotalPublicUrlRequest) (_result *GetTotalPublicUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTotalPublicUrlResponse{}
	_body, _err := client.GetTotalPublicUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取X号码配置信息
//
// @param request - GetXConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetXConfigResponse
func (client *Client) GetXConfigWithOptions(request *GetXConfigRequest, runtime *util.RuntimeOptions) (_result *GetXConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CallerParentId)) {
		query["CallerParentId"] = request.CallerParentId
	}

	if !tea.BoolValue(util.IsUnset(request.CustomerPoolKey)) {
		query["CustomerPoolKey"] = request.CustomerPoolKey
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ReqId)) {
		query["ReqId"] = request.ReqId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.TelX)) {
		query["TelX"] = request.TelX
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetXConfig"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetXConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取X号码配置信息
//
// @param request - GetXConfigRequest
//
// @return GetXConfigResponse
func (client *Client) GetXConfig(request *GetXConfigRequest) (_result *GetXConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetXConfigResponse{}
	_body, _err := client.GetXConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取X号码默认配置信息
//
// @param request - GetXDefaultConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetXDefaultConfigResponse
func (client *Client) GetXDefaultConfigWithOptions(request *GetXDefaultConfigRequest, runtime *util.RuntimeOptions) (_result *GetXDefaultConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CallerParentId)) {
		query["CallerParentId"] = request.CallerParentId
	}

	if !tea.BoolValue(util.IsUnset(request.CustomerPoolKey)) {
		query["CustomerPoolKey"] = request.CustomerPoolKey
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ReqId)) {
		query["ReqId"] = request.ReqId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.TelX)) {
		query["TelX"] = request.TelX
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetXDefaultConfig"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetXDefaultConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取X号码默认配置信息
//
// @param request - GetXDefaultConfigRequest
//
// @return GetXDefaultConfigResponse
func (client *Client) GetXDefaultConfig(request *GetXDefaultConfigRequest) (_result *GetXDefaultConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetXDefaultConfigResponse{}
	_body, _err := client.GetXDefaultConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询客户名下X号码列表
//
// @param request - ListXTelephonesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListXTelephonesResponse
func (client *Client) ListXTelephonesWithOptions(request *ListXTelephonesRequest, runtime *util.RuntimeOptions) (_result *ListXTelephonesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CallerParentId)) {
		query["CallerParentId"] = request.CallerParentId
	}

	if !tea.BoolValue(util.IsUnset(request.CustomerPoolKey)) {
		query["CustomerPoolKey"] = request.CustomerPoolKey
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ReqId)) {
		query["ReqId"] = request.ReqId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListXTelephones"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListXTelephonesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询客户名下X号码列表
//
// @param request - ListXTelephonesRequest
//
// @return ListXTelephonesResponse
func (client *Client) ListXTelephones(request *ListXTelephonesRequest) (_result *ListXTelephonesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListXTelephonesResponse{}
	_body, _err := client.ListXTelephonesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Locks a phone number.
//
// Description:
//
// After a phone number is locked, the locked phone number cannot be selected when you call an operation to create a binding.
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 500 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - LockSecretNoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return LockSecretNoResponse
func (client *Client) LockSecretNoWithOptions(request *LockSecretNoRequest, runtime *util.RuntimeOptions) (_result *LockSecretNoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PoolKey)) {
		query["PoolKey"] = request.PoolKey
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecretNo)) {
		query["SecretNo"] = request.SecretNo
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("LockSecretNo"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &LockSecretNoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Locks a phone number.
//
// Description:
//
// After a phone number is locked, the locked phone number cannot be selected when you call an operation to create a binding.
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 500 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - LockSecretNoRequest
//
// @return LockSecretNoResponse
func (client *Client) LockSecretNo(request *LockSecretNoRequest) (_result *LockSecretNoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &LockSecretNoResponse{}
	_body, _err := client.LockSecretNoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies number group G.
//
// Description:
//
// After you create number group G, you can call the OperateAxgGroup operation to modify number group G. For example, you can add phone numbers to number group G, delete phone numbers from number group G, and replace all phone numbers in number group G.
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 5,000 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - OperateAxgGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OperateAxgGroupResponse
func (client *Client) OperateAxgGroupWithOptions(request *OperateAxgGroupRequest, runtime *util.RuntimeOptions) (_result *OperateAxgGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Numbers)) {
		query["Numbers"] = request.Numbers
	}

	if !tea.BoolValue(util.IsUnset(request.OperateType)) {
		query["OperateType"] = request.OperateType
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PoolKey)) {
		query["PoolKey"] = request.PoolKey
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("OperateAxgGroup"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &OperateAxgGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies number group G.
//
// Description:
//
// After you create number group G, you can call the OperateAxgGroup operation to modify number group G. For example, you can add phone numbers to number group G, delete phone numbers from number group G, and replace all phone numbers in number group G.
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 5,000 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - OperateAxgGroupRequest
//
// @return OperateAxgGroupResponse
func (client *Client) OperateAxgGroup(request *OperateAxgGroupRequest) (_result *OperateAxgGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OperateAxgGroupResponse{}
	_body, _err := client.OperateAxgGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds a phone number to a blacklist or deletes a phone number from a blacklist.
//
// Description:
//
// The OperateBlackNo operation supports the following number pool types: AXN, AXN extension, and 95AXN.
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 1,000 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - OperateBlackNoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OperateBlackNoResponse
func (client *Client) OperateBlackNoWithOptions(request *OperateBlackNoRequest, runtime *util.RuntimeOptions) (_result *OperateBlackNoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BlackNo)) {
		query["BlackNo"] = request.BlackNo
	}

	if !tea.BoolValue(util.IsUnset(request.OperateType)) {
		query["OperateType"] = request.OperateType
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PoolKey)) {
		query["PoolKey"] = request.PoolKey
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Tips)) {
		query["Tips"] = request.Tips
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("OperateBlackNo"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &OperateBlackNoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a phone number to a blacklist or deletes a phone number from a blacklist.
//
// Description:
//
// The OperateBlackNo operation supports the following number pool types: AXN, AXN extension, and 95AXN.
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 1,000 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - OperateBlackNoRequest
//
// @return OperateBlackNoResponse
func (client *Client) OperateBlackNo(request *OperateBlackNoRequest) (_result *OperateBlackNoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OperateBlackNoResponse{}
	_body, _err := client.OperateBlackNoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 固话AxB查询
//
// @param request - QueryAxbBindFixedLineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryAxbBindFixedLineResponse
func (client *Client) QueryAxbBindFixedLineWithOptions(request *QueryAxbBindFixedLineRequest, runtime *util.RuntimeOptions) (_result *QueryAxbBindFixedLineResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.OrderId)) {
		query["OrderId"] = request.OrderId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Phone)) {
		query["Phone"] = request.Phone
	}

	if !tea.BoolValue(util.IsUnset(request.QueryType)) {
		query["QueryType"] = request.QueryType
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SubId)) {
		query["SubId"] = request.SubId
	}

	if !tea.BoolValue(util.IsUnset(request.TelX)) {
		query["TelX"] = request.TelX
	}

	if !tea.BoolValue(util.IsUnset(request.Ts)) {
		query["Ts"] = request.Ts
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryAxbBindFixedLine"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryAxbBindFixedLineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 固话AxB查询
//
// @param request - QueryAxbBindFixedLineRequest
//
// @return QueryAxbBindFixedLineResponse
func (client *Client) QueryAxbBindFixedLine(request *QueryAxbBindFixedLineRequest) (_result *QueryAxbBindFixedLineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryAxbBindFixedLineResponse{}
	_body, _err := client.QueryAxbBindFixedLineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询Axn绑定关系
//
// @param request - QueryAxnBindFixedLineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryAxnBindFixedLineResponse
func (client *Client) QueryAxnBindFixedLineWithOptions(request *QueryAxnBindFixedLineRequest, runtime *util.RuntimeOptions) (_result *QueryAxnBindFixedLineResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.OrderId)) {
		query["OrderId"] = request.OrderId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Phone)) {
		query["Phone"] = request.Phone
	}

	if !tea.BoolValue(util.IsUnset(request.QueryType)) {
		query["QueryType"] = request.QueryType
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SubId)) {
		query["SubId"] = request.SubId
	}

	if !tea.BoolValue(util.IsUnset(request.TelX)) {
		query["TelX"] = request.TelX
	}

	if !tea.BoolValue(util.IsUnset(request.Ts)) {
		query["Ts"] = request.Ts
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryAxnBindFixedLine"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryAxnBindFixedLineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询Axn绑定关系
//
// @param request - QueryAxnBindFixedLineRequest
//
// @return QueryAxnBindFixedLineResponse
func (client *Client) QueryAxnBindFixedLine(request *QueryAxnBindFixedLineRequest) (_result *QueryAxnBindFixedLineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryAxnBindFixedLineResponse{}
	_body, _err := client.QueryAxnBindFixedLineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询AXN分机号绑定关系
//
// @param request - QueryAxnExtensionBindFixedLineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryAxnExtensionBindFixedLineResponse
func (client *Client) QueryAxnExtensionBindFixedLineWithOptions(request *QueryAxnExtensionBindFixedLineRequest, runtime *util.RuntimeOptions) (_result *QueryAxnExtensionBindFixedLineResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.OrderId)) {
		query["OrderId"] = request.OrderId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.QueryType)) {
		query["QueryType"] = request.QueryType
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SubId)) {
		query["SubId"] = request.SubId
	}

	if !tea.BoolValue(util.IsUnset(request.TelA)) {
		query["TelA"] = request.TelA
	}

	if !tea.BoolValue(util.IsUnset(request.Ts)) {
		query["Ts"] = request.Ts
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryAxnExtensionBindFixedLine"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryAxnExtensionBindFixedLineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询AXN分机号绑定关系
//
// @param request - QueryAxnExtensionBindFixedLineRequest
//
// @return QueryAxnExtensionBindFixedLineResponse
func (client *Client) QueryAxnExtensionBindFixedLine(request *QueryAxnExtensionBindFixedLineRequest) (_result *QueryAxnExtensionBindFixedLineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryAxnExtensionBindFixedLineResponse{}
	_body, _err := client.QueryAxnExtensionBindFixedLineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details about a tracking number.
//
// Description:
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 5,000 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - QueryPhoneNoAByTrackNoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryPhoneNoAByTrackNoResponse
func (client *Client) QueryPhoneNoAByTrackNoWithOptions(request *QueryPhoneNoAByTrackNoRequest, runtime *util.RuntimeOptions) (_result *QueryPhoneNoAByTrackNoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CabinetNo)) {
		query["CabinetNo"] = request.CabinetNo
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNoX)) {
		query["PhoneNoX"] = request.PhoneNoX
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.TrackNo)) {
		query["trackNo"] = request.TrackNo
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryPhoneNoAByTrackNo"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryPhoneNoAByTrackNoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details about a tracking number.
//
// Description:
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 5,000 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - QueryPhoneNoAByTrackNoRequest
//
// @return QueryPhoneNoAByTrackNoResponse
func (client *Client) QueryPhoneNoAByTrackNo(request *QueryPhoneNoAByTrackNoRequest) (_result *QueryPhoneNoAByTrackNoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryPhoneNoAByTrackNoResponse{}
	_body, _err := client.QueryPhoneNoAByTrackNoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the download URL of a recording file.
//
// Description:
//
// If the recording feature is enabled for a binding, all calls made by the bound phone numbers are recorded. You can obtain the download URL of a recording file by calling the QueryRecordFileDownloadUrl operation and download the recording file.
//
// >  We recommend that you subscribe to [the recording status report SecretRecording](https://help.aliyun.com/document_detail/109198.html). The values of the response parameters in SecretRecording can be used as the values of the request parameters for downloading a recording file.
//
// ### [](#)Procedure for obtaining a recording file
//
// 1.  Specify the request parameter in an update or binding operation to enable the recording feature.
//
// 2.  Subscribe to recording message receipts in the Phone Number Protection console.
//
// 3.  After a recording message receipt is returned, call the QueryRecordFileDownloadUrl operation to obtain the download URL of the recording file, and download the recording file.
//
// >
//
//   - A download URL is valid for 2 hours. Download the recording file as soon as possible after obtaining a download URL.
//
//   - The storage period of recording files is 30 days. You can download only the recording files of calls recorded in the last 30 days.
//
// @param request - QueryRecordFileDownloadUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryRecordFileDownloadUrlResponse
func (client *Client) QueryRecordFileDownloadUrlWithOptions(request *QueryRecordFileDownloadUrlRequest, runtime *util.RuntimeOptions) (_result *QueryRecordFileDownloadUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CallId)) {
		query["CallId"] = request.CallId
	}

	if !tea.BoolValue(util.IsUnset(request.CallTime)) {
		query["CallTime"] = request.CallTime
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PoolKey)) {
		query["PoolKey"] = request.PoolKey
	}

	if !tea.BoolValue(util.IsUnset(request.ProductType)) {
		query["ProductType"] = request.ProductType
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryRecordFileDownloadUrl"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryRecordFileDownloadUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the download URL of a recording file.
//
// Description:
//
// If the recording feature is enabled for a binding, all calls made by the bound phone numbers are recorded. You can obtain the download URL of a recording file by calling the QueryRecordFileDownloadUrl operation and download the recording file.
//
// >  We recommend that you subscribe to [the recording status report SecretRecording](https://help.aliyun.com/document_detail/109198.html). The values of the response parameters in SecretRecording can be used as the values of the request parameters for downloading a recording file.
//
// ### [](#)Procedure for obtaining a recording file
//
// 1.  Specify the request parameter in an update or binding operation to enable the recording feature.
//
// 2.  Subscribe to recording message receipts in the Phone Number Protection console.
//
// 3.  After a recording message receipt is returned, call the QueryRecordFileDownloadUrl operation to obtain the download URL of the recording file, and download the recording file.
//
// >
//
//   - A download URL is valid for 2 hours. Download the recording file as soon as possible after obtaining a download URL.
//
//   - The storage period of recording files is 30 days. You can download only the recording files of calls recorded in the last 30 days.
//
// @param request - QueryRecordFileDownloadUrlRequest
//
// @return QueryRecordFileDownloadUrlResponse
func (client *Client) QueryRecordFileDownloadUrl(request *QueryRecordFileDownloadUrlRequest) (_result *QueryRecordFileDownloadUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryRecordFileDownloadUrlResponse{}
	_body, _err := client.QueryRecordFileDownloadUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # A号码报备状态查询
//
// @param request - QuerySecretAPhoneNoToCustRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySecretAPhoneNoToCustResponse
func (client *Client) QuerySecretAPhoneNoToCustWithOptions(request *QuerySecretAPhoneNoToCustRequest, runtime *util.RuntimeOptions) (_result *QuerySecretAPhoneNoToCustResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ANoWhiteGroupId)) {
		query["ANoWhiteGroupId"] = request.ANoWhiteGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNoA)) {
		query["PhoneNoA"] = request.PhoneNoA
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QuerySecretAPhoneNoToCust"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QuerySecretAPhoneNoToCustResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # A号码报备状态查询
//
// @param request - QuerySecretAPhoneNoToCustRequest
//
// @return QuerySecretAPhoneNoToCustResponse
func (client *Client) QuerySecretAPhoneNoToCust(request *QuerySecretAPhoneNoToCustRequest) (_result *QuerySecretAPhoneNoToCustResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QuerySecretAPhoneNoToCustResponse{}
	_body, _err := client.QuerySecretAPhoneNoToCustWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the attributes of a private number.
//
// Description:
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 1,000 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. We recommend that you take note of the limit when you call this operation.
//
// @param request - QuerySecretNoDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySecretNoDetailResponse
func (client *Client) QuerySecretNoDetailWithOptions(request *QuerySecretNoDetailRequest, runtime *util.RuntimeOptions) (_result *QuerySecretNoDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PoolKey)) {
		query["PoolKey"] = request.PoolKey
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecretNo)) {
		query["SecretNo"] = request.SecretNo
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QuerySecretNoDetail"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QuerySecretNoDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the attributes of a private number.
//
// Description:
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 1,000 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. We recommend that you take note of the limit when you call this operation.
//
// @param request - QuerySecretNoDetailRequest
//
// @return QuerySecretNoDetailResponse
func (client *Client) QuerySecretNoDetail(request *QuerySecretNoDetailRequest) (_result *QuerySecretNoDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QuerySecretNoDetailResponse{}
	_body, _err := client.QuerySecretNoDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the quantity of remaining phone numbers available for online purchase.
//
// Description:
//
// When purchasing a phone number, specify the home location. If no sufficient phone numbers are available for purchase in the home location, the purchase of the phone number fails. Before calling the [BuySecretNo](~~BuySecretNo~~) operation to purchase a phone number, call the [QuerySecretNoRemain](~~QuerySecretNoRemain~~) operation to query the quantity of remaining phone numbers available for online purchase.
//
// @param request - QuerySecretNoRemainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySecretNoRemainResponse
func (client *Client) QuerySecretNoRemainWithOptions(request *QuerySecretNoRemainRequest, runtime *util.RuntimeOptions) (_result *QuerySecretNoRemainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.City)) {
		query["City"] = request.City
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecretNo)) {
		query["SecretNo"] = request.SecretNo
	}

	if !tea.BoolValue(util.IsUnset(request.SpecId)) {
		query["SpecId"] = request.SpecId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QuerySecretNoRemain"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QuerySecretNoRemainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the quantity of remaining phone numbers available for online purchase.
//
// Description:
//
// When purchasing a phone number, specify the home location. If no sufficient phone numbers are available for purchase in the home location, the purchase of the phone number fails. Before calling the [BuySecretNo](~~BuySecretNo~~) operation to purchase a phone number, call the [QuerySecretNoRemain](~~QuerySecretNoRemain~~) operation to query the quantity of remaining phone numbers available for online purchase.
//
// @param request - QuerySecretNoRemainRequest
//
// @return QuerySecretNoRemainResponse
func (client *Client) QuerySecretNoRemain(request *QuerySecretNoRemainRequest) (_result *QuerySecretNoRemainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QuerySecretNoRemainResponse{}
	_body, _err := client.QuerySecretNoRemainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询通话录音链接
//
// @param request - QuerySoundRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySoundRecordResponse
func (client *Client) QuerySoundRecordWithOptions(request *QuerySoundRecordRequest, runtime *util.RuntimeOptions) (_result *QuerySoundRecordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CallId)) {
		query["CallId"] = request.CallId
	}

	if !tea.BoolValue(util.IsUnset(request.CallerParentId)) {
		query["CallerParentId"] = request.CallerParentId
	}

	if !tea.BoolValue(util.IsUnset(request.CustomerPoolKey)) {
		query["CustomerPoolKey"] = request.CustomerPoolKey
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ReqId)) {
		query["ReqId"] = request.ReqId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QuerySoundRecord"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QuerySoundRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询通话录音链接
//
// @param request - QuerySoundRecordRequest
//
// @return QuerySoundRecordResponse
func (client *Client) QuerySoundRecord(request *QuerySoundRecordRequest) (_result *QuerySoundRecordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QuerySoundRecordResponse{}
	_body, _err := client.QuerySoundRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries binding IDs.
//
// Description:
//
// You can query binding IDs by phone number X. In the AXB product, multiple bindings may exist for the same phone number X. In this case, multiple binding IDs may be obtained for the same phone number X.
//
// @param request - QuerySubsIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySubsIdResponse
func (client *Client) QuerySubsIdWithOptions(request *QuerySubsIdRequest, runtime *util.RuntimeOptions) (_result *QuerySubsIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNoX)) {
		query["PhoneNoX"] = request.PhoneNoX
	}

	if !tea.BoolValue(util.IsUnset(request.PoolKey)) {
		query["PoolKey"] = request.PoolKey
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QuerySubsId"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QuerySubsIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries binding IDs.
//
// Description:
//
// You can query binding IDs by phone number X. In the AXB product, multiple bindings may exist for the same phone number X. In this case, multiple binding IDs may be obtained for the same phone number X.
//
// @param request - QuerySubsIdRequest
//
// @return QuerySubsIdResponse
func (client *Client) QuerySubsId(request *QuerySubsIdRequest) (_result *QuerySubsIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QuerySubsIdResponse{}
	_body, _err := client.QuerySubsIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details about a phone number binding.
//
// Description:
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 5,000 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// ### [](#poolkeyproducttype)Limits on PoolKey and ProductType
//
// You must specify either PoolKey or ProductType. If both parameters are not specified, an error is reported when you call the QuerySubscriptionDetail operation. We recommend that you specify the ProductType parameter for the original key accounts of Alibaba Cloud and the PoolKey parameter for Alibaba Cloud users.
//
// @param request - QuerySubscriptionDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySubscriptionDetailResponse
func (client *Client) QuerySubscriptionDetailWithOptions(request *QuerySubscriptionDetailRequest, runtime *util.RuntimeOptions) (_result *QuerySubscriptionDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNoX)) {
		query["PhoneNoX"] = request.PhoneNoX
	}

	if !tea.BoolValue(util.IsUnset(request.PoolKey)) {
		query["PoolKey"] = request.PoolKey
	}

	if !tea.BoolValue(util.IsUnset(request.ProductType)) {
		query["ProductType"] = request.ProductType
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SubsId)) {
		query["SubsId"] = request.SubsId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QuerySubscriptionDetail"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QuerySubscriptionDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details about a phone number binding.
//
// Description:
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 5,000 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// ### [](#poolkeyproducttype)Limits on PoolKey and ProductType
//
// You must specify either PoolKey or ProductType. If both parameters are not specified, an error is reported when you call the QuerySubscriptionDetail operation. We recommend that you specify the ProductType parameter for the original key accounts of Alibaba Cloud and the PoolKey parameter for Alibaba Cloud users.
//
// @param request - QuerySubscriptionDetailRequest
//
// @return QuerySubscriptionDetailResponse
func (client *Client) QuerySubscriptionDetail(request *QuerySubscriptionDetailRequest) (_result *QuerySubscriptionDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QuerySubscriptionDetailResponse{}
	_body, _err := client.QuerySubscriptionDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Releases a phone number.
//
// Description:
//
//	  After a phone number is released, it will no longer be charged from the following month.
//
//		- Before you release a phone number, log on to the [Phone Number Protection console](https://dypls.console.aliyun.com/dypls.htm#/account) to check whether the phone number is bound to other phone numbers. The phone number can be released only if it is not bound to other phone numbers.
//
// @param request - ReleaseSecretNoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReleaseSecretNoResponse
func (client *Client) ReleaseSecretNoWithOptions(request *ReleaseSecretNoRequest, runtime *util.RuntimeOptions) (_result *ReleaseSecretNoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PoolKey)) {
		query["PoolKey"] = request.PoolKey
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecretNo)) {
		query["SecretNo"] = request.SecretNo
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ReleaseSecretNo"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ReleaseSecretNoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Releases a phone number.
//
// Description:
//
//	  After a phone number is released, it will no longer be charged from the following month.
//
//		- Before you release a phone number, log on to the [Phone Number Protection console](https://dypls.console.aliyun.com/dypls.htm#/account) to check whether the phone number is bound to other phone numbers. The phone number can be released only if it is not bound to other phone numbers.
//
// @param request - ReleaseSecretNoRequest
//
// @return ReleaseSecretNoResponse
func (client *Client) ReleaseSecretNo(request *ReleaseSecretNoRequest) (_result *ReleaseSecretNoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReleaseSecretNoResponse{}
	_body, _err := client.ReleaseSecretNoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 解除指定的呼叫绑定关系（A，X，B），解决呼叫绑定关系后，员工B不能通过工作号X呼叫到客户A。
//
// @param request - UnBindAXBRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnBindAXBResponse
func (client *Client) UnBindAXBWithOptions(request *UnBindAXBRequest, runtime *util.RuntimeOptions) (_result *UnBindAXBResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BindId)) {
		query["BindId"] = request.BindId
	}

	if !tea.BoolValue(util.IsUnset(request.CallerParentId)) {
		query["CallerParentId"] = request.CallerParentId
	}

	if !tea.BoolValue(util.IsUnset(request.CustomerPoolKey)) {
		query["CustomerPoolKey"] = request.CustomerPoolKey
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ReqId)) {
		query["ReqId"] = request.ReqId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UnBindAXB"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UnBindAXBResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 解除指定的呼叫绑定关系（A，X，B），解决呼叫绑定关系后，员工B不能通过工作号X呼叫到客户A。
//
// @param request - UnBindAXBRequest
//
// @return UnBindAXBResponse
func (client *Client) UnBindAXB(request *UnBindAXBRequest) (_result *UnBindAXBResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnBindAXBResponse{}
	_body, _err := client.UnBindAXBWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 调用本接口可取消工作号X与员工号码B的绑定。绑定解除后，对X的呼叫都不会转接给B。
//
// @param request - UnBindXBRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnBindXBResponse
func (client *Client) UnBindXBWithOptions(request *UnBindXBRequest, runtime *util.RuntimeOptions) (_result *UnBindXBResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthId)) {
		query["AuthId"] = request.AuthId
	}

	if !tea.BoolValue(util.IsUnset(request.CallerParentId)) {
		query["CallerParentId"] = request.CallerParentId
	}

	if !tea.BoolValue(util.IsUnset(request.CustomerPoolKey)) {
		query["CustomerPoolKey"] = request.CustomerPoolKey
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ReqId)) {
		query["ReqId"] = request.ReqId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.TelX)) {
		query["TelX"] = request.TelX
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UnBindXB"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UnBindXBResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 调用本接口可取消工作号X与员工号码B的绑定。绑定解除后，对X的呼叫都不会转接给B。
//
// @param request - UnBindXBRequest
//
// @return UnBindXBResponse
func (client *Client) UnBindXB(request *UnBindXBRequest) (_result *UnBindXBResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnBindXBResponse{}
	_body, _err := client.UnBindXBWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Unbinds a phone number.
//
// Description:
//
// Before releasing a phone number, you must call the UnbindSubscription operation to unbind the phone number.
//
// @param request - UnbindSubscriptionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnbindSubscriptionResponse
func (client *Client) UnbindSubscriptionWithOptions(request *UnbindSubscriptionRequest, runtime *util.RuntimeOptions) (_result *UnbindSubscriptionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PoolKey)) {
		query["PoolKey"] = request.PoolKey
	}

	if !tea.BoolValue(util.IsUnset(request.ProductType)) {
		query["ProductType"] = request.ProductType
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecretNo)) {
		query["SecretNo"] = request.SecretNo
	}

	if !tea.BoolValue(util.IsUnset(request.SubsId)) {
		query["SubsId"] = request.SubsId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UnbindSubscription"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UnbindSubscriptionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Unbinds a phone number.
//
// Description:
//
// Before releasing a phone number, you must call the UnbindSubscription operation to unbind the phone number.
//
// @param request - UnbindSubscriptionRequest
//
// @return UnbindSubscriptionResponse
func (client *Client) UnbindSubscription(request *UnbindSubscriptionRequest) (_result *UnbindSubscriptionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnbindSubscriptionResponse{}
	_body, _err := client.UnbindSubscriptionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Unlocks a phone number.
//
// Description:
//
// After a phone number is unlocked, you can reselect the unlocked phone number when you call an operation to create a binding.
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 500 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - UnlockSecretNoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnlockSecretNoResponse
func (client *Client) UnlockSecretNoWithOptions(request *UnlockSecretNoRequest, runtime *util.RuntimeOptions) (_result *UnlockSecretNoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PoolKey)) {
		query["PoolKey"] = request.PoolKey
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecretNo)) {
		query["SecretNo"] = request.SecretNo
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UnlockSecretNo"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UnlockSecretNoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Unlocks a phone number.
//
// Description:
//
// After a phone number is unlocked, you can reselect the unlocked phone number when you call an operation to create a binding.
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 500 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - UnlockSecretNoRequest
//
// @return UnlockSecretNoResponse
func (client *Client) UnlockSecretNo(request *UnlockSecretNoRequest) (_result *UnlockSecretNoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnlockSecretNoResponse{}
	_body, _err := client.UnlockSecretNoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 固话AxB绑定更新
//
// @param tmpReq - UpdateAxbBindFixedLineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAxbBindFixedLineResponse
func (client *Client) UpdateAxbBindFixedLineWithOptions(tmpReq *UpdateAxbBindFixedLineRequest, runtime *util.RuntimeOptions) (_result *UpdateAxbBindFixedLineResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateAxbBindFixedLineShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Extra)) {
		request.ExtraShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Extra, tea.String("Extra"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Anucode)) {
		query["Anucode"] = request.Anucode
	}

	if !tea.BoolValue(util.IsUnset(request.Anucodecalled)) {
		query["Anucodecalled"] = request.Anucodecalled
	}

	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.Expiration)) {
		query["Expiration"] = request.Expiration
	}

	if !tea.BoolValue(util.IsUnset(request.ExtraShrink)) {
		query["Extra"] = request.ExtraShrink
	}

	if !tea.BoolValue(util.IsUnset(request.OrderId)) {
		query["OrderId"] = request.OrderId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		query["Remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SubId)) {
		query["SubId"] = request.SubId
	}

	if !tea.BoolValue(util.IsUnset(request.Subts)) {
		query["Subts"] = request.Subts
	}

	if !tea.BoolValue(util.IsUnset(request.TAnucodeConnect)) {
		query["TAnucodeConnect"] = request.TAnucodeConnect
	}

	if !tea.BoolValue(util.IsUnset(request.TelA)) {
		query["TelA"] = request.TelA
	}

	if !tea.BoolValue(util.IsUnset(request.TelB)) {
		query["TelB"] = request.TelB
	}

	if !tea.BoolValue(util.IsUnset(request.Ts)) {
		query["Ts"] = request.Ts
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateAxbBindFixedLine"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateAxbBindFixedLineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 固话AxB绑定更新
//
// @param request - UpdateAxbBindFixedLineRequest
//
// @return UpdateAxbBindFixedLineResponse
func (client *Client) UpdateAxbBindFixedLine(request *UpdateAxbBindFixedLineRequest) (_result *UpdateAxbBindFixedLineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAxbBindFixedLineResponse{}
	_body, _err := client.UpdateAxbBindFixedLineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新Axn绑定关系
//
// @param tmpReq - UpdateAxnBindFixedLineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAxnBindFixedLineResponse
func (client *Client) UpdateAxnBindFixedLineWithOptions(tmpReq *UpdateAxnBindFixedLineRequest, runtime *util.RuntimeOptions) (_result *UpdateAxnBindFixedLineResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateAxnBindFixedLineShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Extra)) {
		request.ExtraShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Extra, tea.String("Extra"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Anucode)) {
		query["Anucode"] = request.Anucode
	}

	if !tea.BoolValue(util.IsUnset(request.Anucodecalled)) {
		query["Anucodecalled"] = request.Anucodecalled
	}

	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.Expiration)) {
		query["Expiration"] = request.Expiration
	}

	if !tea.BoolValue(util.IsUnset(request.ExtraShrink)) {
		query["Extra"] = request.ExtraShrink
	}

	if !tea.BoolValue(util.IsUnset(request.OrderId)) {
		query["OrderId"] = request.OrderId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		query["Remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SubId)) {
		query["SubId"] = request.SubId
	}

	if !tea.BoolValue(util.IsUnset(request.Subts)) {
		query["Subts"] = request.Subts
	}

	if !tea.BoolValue(util.IsUnset(request.TAnucodeConnect)) {
		query["TAnucodeConnect"] = request.TAnucodeConnect
	}

	if !tea.BoolValue(util.IsUnset(request.TelA)) {
		query["TelA"] = request.TelA
	}

	if !tea.BoolValue(util.IsUnset(request.TelB)) {
		query["TelB"] = request.TelB
	}

	if !tea.BoolValue(util.IsUnset(request.Ts)) {
		query["Ts"] = request.Ts
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateAxnBindFixedLine"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateAxnBindFixedLineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新Axn绑定关系
//
// @param request - UpdateAxnBindFixedLineRequest
//
// @return UpdateAxnBindFixedLineResponse
func (client *Client) UpdateAxnBindFixedLine(request *UpdateAxnBindFixedLineRequest) (_result *UpdateAxnBindFixedLineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAxnBindFixedLineResponse{}
	_body, _err := client.UpdateAxnBindFixedLineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新AXN分机号绑定关系
//
// @param tmpReq - UpdateAxnExtensionBindFixedLineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAxnExtensionBindFixedLineResponse
func (client *Client) UpdateAxnExtensionBindFixedLineWithOptions(tmpReq *UpdateAxnExtensionBindFixedLineRequest, runtime *util.RuntimeOptions) (_result *UpdateAxnExtensionBindFixedLineResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateAxnExtensionBindFixedLineShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Extraaxx)) {
		request.ExtraaxxShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Extraaxx, tea.String("Extraaxx"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Anucode)) {
		query["Anucode"] = request.Anucode
	}

	if !tea.BoolValue(util.IsUnset(request.Anucodecalled)) {
		query["Anucodecalled"] = request.Anucodecalled
	}

	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.Expiration)) {
		query["Expiration"] = request.Expiration
	}

	if !tea.BoolValue(util.IsUnset(request.ExtraaxxShrink)) {
		query["Extraaxx"] = request.ExtraaxxShrink
	}

	if !tea.BoolValue(util.IsUnset(request.OrderId)) {
		query["OrderId"] = request.OrderId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		query["Remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SubId)) {
		query["SubId"] = request.SubId
	}

	if !tea.BoolValue(util.IsUnset(request.Subts)) {
		query["Subts"] = request.Subts
	}

	if !tea.BoolValue(util.IsUnset(request.TAnucodeConnect)) {
		query["TAnucodeConnect"] = request.TAnucodeConnect
	}

	if !tea.BoolValue(util.IsUnset(request.TelA)) {
		query["TelA"] = request.TelA
	}

	if !tea.BoolValue(util.IsUnset(request.Ts)) {
		query["Ts"] = request.Ts
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateAxnExtensionBindFixedLine"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateAxnExtensionBindFixedLineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新AXN分机号绑定关系
//
// @param request - UpdateAxnExtensionBindFixedLineRequest
//
// @return UpdateAxnExtensionBindFixedLineResponse
func (client *Client) UpdateAxnExtensionBindFixedLine(request *UpdateAxnExtensionBindFixedLineRequest) (_result *UpdateAxnExtensionBindFixedLineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAxnExtensionBindFixedLineResponse{}
	_body, _err := client.UpdateAxnExtensionBindFixedLineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies a phone number binding.
//
// Description:
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 10,000 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - UpdateSubscriptionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSubscriptionResponse
func (client *Client) UpdateSubscriptionWithOptions(request *UpdateSubscriptionRequest, runtime *util.RuntimeOptions) (_result *UpdateSubscriptionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ASRModelId)) {
		query["ASRModelId"] = request.ASRModelId
	}

	if !tea.BoolValue(util.IsUnset(request.ASRStatus)) {
		query["ASRStatus"] = request.ASRStatus
	}

	if !tea.BoolValue(util.IsUnset(request.CallDisplayType)) {
		query["CallDisplayType"] = request.CallDisplayType
	}

	if !tea.BoolValue(util.IsUnset(request.CallRestrict)) {
		query["CallRestrict"] = request.CallRestrict
	}

	if !tea.BoolValue(util.IsUnset(request.Expiration)) {
		query["Expiration"] = request.Expiration
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.IsRecordingEnabled)) {
		query["IsRecordingEnabled"] = request.IsRecordingEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.OperateType)) {
		query["OperateType"] = request.OperateType
	}

	if !tea.BoolValue(util.IsUnset(request.OutId)) {
		query["OutId"] = request.OutId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNoA)) {
		query["PhoneNoA"] = request.PhoneNoA
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNoB)) {
		query["PhoneNoB"] = request.PhoneNoB
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNoX)) {
		query["PhoneNoX"] = request.PhoneNoX
	}

	if !tea.BoolValue(util.IsUnset(request.PoolKey)) {
		query["PoolKey"] = request.PoolKey
	}

	if !tea.BoolValue(util.IsUnset(request.ProductType)) {
		query["ProductType"] = request.ProductType
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RingConfig)) {
		query["RingConfig"] = request.RingConfig
	}

	if !tea.BoolValue(util.IsUnset(request.SubsId)) {
		query["SubsId"] = request.SubsId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateSubscription"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateSubscriptionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a phone number binding.
//
// Description:
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 10,000 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - UpdateSubscriptionRequest
//
// @return UpdateSubscriptionResponse
func (client *Client) UpdateSubscription(request *UpdateSubscriptionRequest) (_result *UpdateSubscriptionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateSubscriptionResponse{}
	_body, _err := client.UpdateSubscriptionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
