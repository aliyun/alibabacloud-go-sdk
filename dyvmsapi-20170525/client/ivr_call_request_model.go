// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIvrCallRequest interface {
	dara.Model
	String() string
	GoString() string
	SetByeCode(v string) *IvrCallRequest
	GetByeCode() *string
	SetByeTtsParams(v string) *IvrCallRequest
	GetByeTtsParams() *string
	SetCalledNumber(v string) *IvrCallRequest
	GetCalledNumber() *string
	SetCalledShowNumber(v string) *IvrCallRequest
	GetCalledShowNumber() *string
	SetMenuKeyMap(v []*IvrCallRequestMenuKeyMap) *IvrCallRequest
	GetMenuKeyMap() []*IvrCallRequestMenuKeyMap
	SetOutId(v string) *IvrCallRequest
	GetOutId() *string
	SetOwnerId(v int64) *IvrCallRequest
	GetOwnerId() *int64
	SetPlayTimes(v int64) *IvrCallRequest
	GetPlayTimes() *int64
	SetResourceOwnerAccount(v string) *IvrCallRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *IvrCallRequest
	GetResourceOwnerId() *int64
	SetStartCode(v string) *IvrCallRequest
	GetStartCode() *string
	SetStartTtsParams(v string) *IvrCallRequest
	GetStartTtsParams() *string
	SetTimeout(v int32) *IvrCallRequest
	GetTimeout() *int32
}

type IvrCallRequest struct {
	// The end voice.
	//
	// 	- If you use a voice notification file, this parameter specifies the voice ID. You can log on to the [Voice Messaging Service console](https://dyvms.console.aliyun.com/overview/home), choose **Voice Messages*	- > **Voice Notifications**, and then click the **Voice Notification Files*	- tab to view the voice ID.
	//
	// 	- If you use a TTS template, this parameter specifies the template ID. You can log on to the [Voice Messaging Service console](https://dyvms.console.aliyun.com/overview/home), choose **Voice Messages*	- > **Voice Notifications**, and then click the **TTS Template*	- tab to view the template ID.
	//
	// > The value of the ByeCode parameter must be of the same type as the value of the StartCode parameter. This means that both parameters must specify voice IDs or TTS template IDs.
	//
	// example:
	//
	// TTS_1234****
	ByeCode *string `json:"ByeCode,omitempty" xml:"ByeCode,omitempty"`
	// The variables in the TTS template, in the JSON format.
	//
	// > This parameter is required when the ByeCode parameter is set to the ID of a TTS template that contains variables.
	//
	// example:
	//
	// {"name":"xxx","code":"123"}
	ByeTtsParams *string `json:"ByeTtsParams,omitempty" xml:"ByeTtsParams,omitempty"`
	// The called number.
	//
	// Only phone numbers in the Chinese mainland are supported. Each request supports only one called number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1590****000
	CalledNumber *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	// The calling number.
	//
	// The value must be a number you purchased. Each request supports only one calling number. In most cases, a calling number is configured with the maximum number of concurrent requests. New requests fail if the maximum number of concurrent requests is reached. You can log on to the [Voice Messaging Service console](https://dyvms.console.aliyun.com/overview/home) and choose **Real Number Service > Real Number Management*	- to view the number you purchased.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0571****5678
	CalledShowNumber *string `json:"CalledShowNumber,omitempty" xml:"CalledShowNumber,omitempty"`
	// The information about the key pressed by the subscriber.
	MenuKeyMap []*IvrCallRequestMenuKeyMap `json:"MenuKeyMap,omitempty" xml:"MenuKeyMap,omitempty" type:"Repeated"`
	// The ID that is reserved for the caller of the operation. This ID is returned to the caller in a receipt message.
	//
	// The value is of the STRING type and must be 1 to 15 bytes in length.
	//
	// example:
	//
	// PR0210428****
	OutId   *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of replay times. Valid values: 1 to 3.
	//
	// example:
	//
	// 3
	PlayTimes            *int64  `json:"PlayTimes,omitempty" xml:"PlayTimes,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The voice that is played when the call begins.
	//
	// 	- If you use a voice notification file, this parameter specifies the voice ID. You can log on to the [Voice Messaging Service console](https://dyvms.console.aliyun.com/overview/home), choose **Voice Messages*	- > Voice Notifications, and then click the **Voice Notification Files*	- tab to view the voice ID.
	//
	// 	- If you use a text-to-speech (TTS) template, this parameter specifies the template ID. You can log on to the [Voice Messaging Service console](https://dyvms.console.aliyun.com/overview/home), choose **Voice Messages*	- > **Voice Notifications**, and then click the **TTS Template*	- tab to view the template ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// TTS_1234****
	StartCode *string `json:"StartCode,omitempty" xml:"StartCode,omitempty"`
	// The variables in the TTS template, in the JSON format.
	//
	// > This parameter is required when the StartCode parameter is set to the ID of a TTS template that contains variables.
	//
	// example:
	//
	// {"name":"xxx","code":"123"}
	StartTtsParams *string `json:"StartTtsParams,omitempty" xml:"StartTtsParams,omitempty"`
	// The timeout period for the subscriber to press a key. Unit: milliseconds.
	//
	// example:
	//
	// 3000
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s IvrCallRequest) String() string {
	return dara.Prettify(s)
}

func (s IvrCallRequest) GoString() string {
	return s.String()
}

func (s *IvrCallRequest) GetByeCode() *string {
	return s.ByeCode
}

func (s *IvrCallRequest) GetByeTtsParams() *string {
	return s.ByeTtsParams
}

func (s *IvrCallRequest) GetCalledNumber() *string {
	return s.CalledNumber
}

func (s *IvrCallRequest) GetCalledShowNumber() *string {
	return s.CalledShowNumber
}

func (s *IvrCallRequest) GetMenuKeyMap() []*IvrCallRequestMenuKeyMap {
	return s.MenuKeyMap
}

func (s *IvrCallRequest) GetOutId() *string {
	return s.OutId
}

func (s *IvrCallRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *IvrCallRequest) GetPlayTimes() *int64 {
	return s.PlayTimes
}

func (s *IvrCallRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *IvrCallRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *IvrCallRequest) GetStartCode() *string {
	return s.StartCode
}

func (s *IvrCallRequest) GetStartTtsParams() *string {
	return s.StartTtsParams
}

func (s *IvrCallRequest) GetTimeout() *int32 {
	return s.Timeout
}

func (s *IvrCallRequest) SetByeCode(v string) *IvrCallRequest {
	s.ByeCode = &v
	return s
}

func (s *IvrCallRequest) SetByeTtsParams(v string) *IvrCallRequest {
	s.ByeTtsParams = &v
	return s
}

func (s *IvrCallRequest) SetCalledNumber(v string) *IvrCallRequest {
	s.CalledNumber = &v
	return s
}

func (s *IvrCallRequest) SetCalledShowNumber(v string) *IvrCallRequest {
	s.CalledShowNumber = &v
	return s
}

func (s *IvrCallRequest) SetMenuKeyMap(v []*IvrCallRequestMenuKeyMap) *IvrCallRequest {
	s.MenuKeyMap = v
	return s
}

func (s *IvrCallRequest) SetOutId(v string) *IvrCallRequest {
	s.OutId = &v
	return s
}

func (s *IvrCallRequest) SetOwnerId(v int64) *IvrCallRequest {
	s.OwnerId = &v
	return s
}

func (s *IvrCallRequest) SetPlayTimes(v int64) *IvrCallRequest {
	s.PlayTimes = &v
	return s
}

func (s *IvrCallRequest) SetResourceOwnerAccount(v string) *IvrCallRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *IvrCallRequest) SetResourceOwnerId(v int64) *IvrCallRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *IvrCallRequest) SetStartCode(v string) *IvrCallRequest {
	s.StartCode = &v
	return s
}

func (s *IvrCallRequest) SetStartTtsParams(v string) *IvrCallRequest {
	s.StartTtsParams = &v
	return s
}

func (s *IvrCallRequest) SetTimeout(v int32) *IvrCallRequest {
	s.Timeout = &v
	return s
}

func (s *IvrCallRequest) Validate() error {
	return dara.Validate(s)
}

type IvrCallRequestMenuKeyMap struct {
	// The voice that corresponds to the key specified by the **MenuKeyMap.N.Key*	- parameter.
	//
	// 	- If you use a voice notification file, this parameter specifies the voice ID. You can log on to the [Voice Messaging Service console](https://dyvms.console.aliyun.com/overview/home), choose **Voice Messages*	- > **Voice Notifications**, and then click the **Voice Notification Files*	- tab to view the voice ID.
	//
	// 	- If you use a TTS template, this parameter specifies the template ID. You can log on to the [Voice Messaging Service console](https://dyvms.console.aliyun.com/overview/home), choose **Voice Messages*	- > **Voice Notifications**, and then click the **TTS Template*	- tab to view the template ID.
	//
	// example:
	//
	// TTS_1235****
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The key that can be pressed by the subscriber.
	//
	// example:
	//
	// 1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The variables in the TTS template, in the JSON format.
	//
	// >
	//
	// 	- This parameter specifies the substitution relationship of the variables in the TTS template if the value of the **MenuKeyMap.N.Code*	- parameter is set to the ID of the TTS template.
	//
	// 	- This parameter is required if the value of the **MenuKeyMap.N.Code*	- parameter is set to the ID of a TTS template that contains variables.
	//
	// example:
	//
	// {"name":"xxx","code":"123"}
	TtsParams *string `json:"TtsParams,omitempty" xml:"TtsParams,omitempty"`
}

func (s IvrCallRequestMenuKeyMap) String() string {
	return dara.Prettify(s)
}

func (s IvrCallRequestMenuKeyMap) GoString() string {
	return s.String()
}

func (s *IvrCallRequestMenuKeyMap) GetCode() *string {
	return s.Code
}

func (s *IvrCallRequestMenuKeyMap) GetKey() *string {
	return s.Key
}

func (s *IvrCallRequestMenuKeyMap) GetTtsParams() *string {
	return s.TtsParams
}

func (s *IvrCallRequestMenuKeyMap) SetCode(v string) *IvrCallRequestMenuKeyMap {
	s.Code = &v
	return s
}

func (s *IvrCallRequestMenuKeyMap) SetKey(v string) *IvrCallRequestMenuKeyMap {
	s.Key = &v
	return s
}

func (s *IvrCallRequestMenuKeyMap) SetTtsParams(v string) *IvrCallRequestMenuKeyMap {
	s.TtsParams = &v
	return s
}

func (s *IvrCallRequestMenuKeyMap) Validate() error {
	return dara.Validate(s)
}
