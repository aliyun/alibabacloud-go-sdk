// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSingleCallByTtsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCalledNumber(v string) *SingleCallByTtsRequest
	GetCalledNumber() *string
	SetCalledShowNumber(v string) *SingleCallByTtsRequest
	GetCalledShowNumber() *string
	SetOutId(v string) *SingleCallByTtsRequest
	GetOutId() *string
	SetOwnerId(v int64) *SingleCallByTtsRequest
	GetOwnerId() *int64
	SetPlayTimes(v int32) *SingleCallByTtsRequest
	GetPlayTimes() *int32
	SetResourceOwnerAccount(v string) *SingleCallByTtsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SingleCallByTtsRequest
	GetResourceOwnerId() *int64
	SetSpeed(v int32) *SingleCallByTtsRequest
	GetSpeed() *int32
	SetTtsCode(v string) *SingleCallByTtsRequest
	GetTtsCode() *string
	SetTtsParam(v string) *SingleCallByTtsRequest
	GetTtsParam() *string
	SetVolume(v int32) *SingleCallByTtsRequest
	GetVolume() *int32
}

type SingleCallByTtsRequest struct {
	// The mobile phone number that receives voice notifications.
	//
	// 	- Number format in the Chinese mainland:
	//
	//     	- Mobile phone number, for example, 159\\*\\*\\*\\*0000.
	//
	//     	- Landline number, for example, 0571\\*\\*\\*\\*5678.
	//
	// 	- Number format outside the Chinese mainland: country code + phone number, for example, 85200\\*\\*\\*\\*00.
	//
	// >
	//
	// 	- Each request supports only one called number. For more information, see [How to use voice notifications in the Chinese mainland](https://help.aliyun.com/document_detail/150016.html) or [How to use voice verification codes in regions outside the Chinese mainland](https://help.aliyun.com/document_detail/270044.html).
	//
	// 	- Voice verification codes are sent to a called number at the following frequency: one time per minute, five times per hour, and 20 times per 24 hours.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1590****000
	CalledNumber *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	// The number displayed to the called party.
	//
	// 	- You do not need to specify this parameter if you use the text-to-speech (TTS) notification template or voice verification code template for outbound calls in the common mode. For more information, see [FAQ about the common outbound call mode](https://help.aliyun.com/document_detail/172104.html).
	//
	// 	- If you use the TTS notification template or voice verification code template for outbound calls in the dedicated mode, you must specify a number you purchased and only one number can be specified. You can log on to the [Voice Messaging Service console](https://dyvms.console.aliyun.com/overview/home) and choose **Voice Numbers*	- > **Real Number Management*	- to view the number you purchased.
	//
	// example:
	//
	// 0571****5678
	CalledShowNumber *string `json:"CalledShowNumber,omitempty" xml:"CalledShowNumber,omitempty"`
	// The custom ID that is reserved for the caller of the operation when the request is initiated. This ID is returned to the caller in a receipt message.
	//
	// The value is of the STRING type and must be 1 to 15 bytes in length.
	//
	// example:
	//
	// 225869*****
	OutId   *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of times a voice notification is played back in a call. Valid values: 1 to 3. Default value: 3.
	//
	// example:
	//
	// 3
	PlayTimes            *int32  `json:"PlayTimes,omitempty" xml:"PlayTimes,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The playback speed. Valid value: -500 to 500.
	//
	// example:
	//
	// 5
	Speed *int32 `json:"Speed,omitempty" xml:"Speed,omitempty"`
	// The ID of the approved TTS notification template or voice verification code template.
	//
	// You can log on to the [Voice Messaging Service console](https://dyvms.console.aliyun.com/overview/home), and choose **Voice Messages*	- > **Voice Verification Codes*	- or choose **Voice Messages*	- > **Voice Notifications*	- to view the **template ID**.
	//
	// > The account to which the TTS template belongs must be the same as the account that is used to call the SingleCallByTts operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// TTS_100****
	TtsCode *string `json:"TtsCode,omitempty" xml:"TtsCode,omitempty"`
	// The variables in the template, in the JSON format.
	//
	// > The variables in the template must be less than 250 characters in length. The length of each single variable is not limited. These variables do not support URLs. The variables in the verification code template support only digits and letters.
	//
	// example:
	//
	// {"AckNum":"123456"}
	TtsParam *string `json:"TtsParam,omitempty" xml:"TtsParam,omitempty"`
	// The playback volume of the voice notification. Valid values: 0 to 100. Default value: 100.
	//
	// example:
	//
	// 100
	Volume *int32 `json:"Volume,omitempty" xml:"Volume,omitempty"`
}

func (s SingleCallByTtsRequest) String() string {
	return dara.Prettify(s)
}

func (s SingleCallByTtsRequest) GoString() string {
	return s.String()
}

func (s *SingleCallByTtsRequest) GetCalledNumber() *string {
	return s.CalledNumber
}

func (s *SingleCallByTtsRequest) GetCalledShowNumber() *string {
	return s.CalledShowNumber
}

func (s *SingleCallByTtsRequest) GetOutId() *string {
	return s.OutId
}

func (s *SingleCallByTtsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SingleCallByTtsRequest) GetPlayTimes() *int32 {
	return s.PlayTimes
}

func (s *SingleCallByTtsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SingleCallByTtsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SingleCallByTtsRequest) GetSpeed() *int32 {
	return s.Speed
}

func (s *SingleCallByTtsRequest) GetTtsCode() *string {
	return s.TtsCode
}

func (s *SingleCallByTtsRequest) GetTtsParam() *string {
	return s.TtsParam
}

func (s *SingleCallByTtsRequest) GetVolume() *int32 {
	return s.Volume
}

func (s *SingleCallByTtsRequest) SetCalledNumber(v string) *SingleCallByTtsRequest {
	s.CalledNumber = &v
	return s
}

func (s *SingleCallByTtsRequest) SetCalledShowNumber(v string) *SingleCallByTtsRequest {
	s.CalledShowNumber = &v
	return s
}

func (s *SingleCallByTtsRequest) SetOutId(v string) *SingleCallByTtsRequest {
	s.OutId = &v
	return s
}

func (s *SingleCallByTtsRequest) SetOwnerId(v int64) *SingleCallByTtsRequest {
	s.OwnerId = &v
	return s
}

func (s *SingleCallByTtsRequest) SetPlayTimes(v int32) *SingleCallByTtsRequest {
	s.PlayTimes = &v
	return s
}

func (s *SingleCallByTtsRequest) SetResourceOwnerAccount(v string) *SingleCallByTtsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SingleCallByTtsRequest) SetResourceOwnerId(v int64) *SingleCallByTtsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SingleCallByTtsRequest) SetSpeed(v int32) *SingleCallByTtsRequest {
	s.Speed = &v
	return s
}

func (s *SingleCallByTtsRequest) SetTtsCode(v string) *SingleCallByTtsRequest {
	s.TtsCode = &v
	return s
}

func (s *SingleCallByTtsRequest) SetTtsParam(v string) *SingleCallByTtsRequest {
	s.TtsParam = &v
	return s
}

func (s *SingleCallByTtsRequest) SetVolume(v int32) *SingleCallByTtsRequest {
	s.Volume = &v
	return s
}

func (s *SingleCallByTtsRequest) Validate() error {
	return dara.Validate(s)
}
