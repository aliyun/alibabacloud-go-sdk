// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSingleCallByVoiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCalledNumber(v string) *SingleCallByVoiceRequest
	GetCalledNumber() *string
	SetCalledShowNumber(v string) *SingleCallByVoiceRequest
	GetCalledShowNumber() *string
	SetOutId(v string) *SingleCallByVoiceRequest
	GetOutId() *string
	SetOwnerId(v int64) *SingleCallByVoiceRequest
	GetOwnerId() *int64
	SetPlayTimes(v int32) *SingleCallByVoiceRequest
	GetPlayTimes() *int32
	SetResourceOwnerAccount(v string) *SingleCallByVoiceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SingleCallByVoiceRequest
	GetResourceOwnerId() *int64
	SetSpeed(v int32) *SingleCallByVoiceRequest
	GetSpeed() *int32
	SetVoiceCode(v string) *SingleCallByVoiceRequest
	GetVoiceCode() *string
	SetVolume(v int32) *SingleCallByVoiceRequest
	GetVolume() *int32
}

type SingleCallByVoiceRequest struct {
	// The number for receiving voice notifications.
	//
	// Number format:
	//
	// 	- In the Chinese mainland:
	//
	//     	- Mobile phone number, for example, 159\\*\\*\\*\\*0000.
	//
	//     	- Landline number, for example, 0571\\*\\*\\*\\*5678.
	//
	// 	- Outside the Chinese mainland: country code + phone number, for example, 85200\\*\\*\\*\\*00.
	//
	// >
	//
	// 	- You can specify only one called number for a request. For more information, see [How to use voice notifications in the Chinese mainland](https://help.aliyun.com/document_detail/150016.html) or [How to use voice notifications in regions outside the Chinese mainland](https://help.aliyun.com/document_detail/268810.html).
	//
	// 	- Voice notifications are sent to a called number at the following frequency: one time per minute, five times per hour, and 20 times per 24 hours.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1590****000
	CalledNumber *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	// The number displayed to the called party.
	//
	// 	- You do not need to specify this parameter if you use a voice notification file that uses the common outbound call mode. For more information, see [FAQ about the common outbound call mode](https://help.aliyun.com/document_detail/172104.html).
	//
	// 	- If you use a voice notification file that uses the dedicated outbound call mode, you must specify a number that you purchased. You can specify only one number. You can log on to the [Voice Messaging Service console](https://dyvms.console.aliyun.com/overview/home) and choose **Real Number Service*	- > **Real Number Management*	- to view the number that you purchased.
	//
	// example:
	//
	// 0571****5678
	CalledShowNumber *string `json:"CalledShowNumber,omitempty" xml:"CalledShowNumber,omitempty"`
	// The ID reserved for the caller. This ID is returned to the caller in a receipt message.
	//
	// The value must be of the STRING type and 1 to 15 bytes in length.
	//
	// example:
	//
	// 22596****
	OutId   *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of times the voice notification file is played. Valid values: 1 to 3.
	//
	// example:
	//
	// 3
	PlayTimes            *int32  `json:"PlayTimes,omitempty" xml:"PlayTimes,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The playback speed of the voice notification file. Valid values: -500 to 500.
	//
	// example:
	//
	// 100
	Speed *int32 `json:"Speed,omitempty" xml:"Speed,omitempty"`
	// The voice ID of the voice notification file.
	//
	// You can log on to the [Voice Messaging Service console](https://dyvms.console.aliyun.com/overview/home), choose **Voice Messages*	- > **Voice Notifications*	- or **Voice File Management**, and then click the **Voice Notification Files*	- tab to view the **voice ID**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2d4c-4e78-8d2a-afbb06cf****.wav
	VoiceCode *string `json:"VoiceCode,omitempty" xml:"VoiceCode,omitempty"`
	// The playback volume of the voice notification file. Valid values: 0 to 100. Default value: 100.
	//
	// example:
	//
	// 100
	Volume *int32 `json:"Volume,omitempty" xml:"Volume,omitempty"`
}

func (s SingleCallByVoiceRequest) String() string {
	return dara.Prettify(s)
}

func (s SingleCallByVoiceRequest) GoString() string {
	return s.String()
}

func (s *SingleCallByVoiceRequest) GetCalledNumber() *string {
	return s.CalledNumber
}

func (s *SingleCallByVoiceRequest) GetCalledShowNumber() *string {
	return s.CalledShowNumber
}

func (s *SingleCallByVoiceRequest) GetOutId() *string {
	return s.OutId
}

func (s *SingleCallByVoiceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SingleCallByVoiceRequest) GetPlayTimes() *int32 {
	return s.PlayTimes
}

func (s *SingleCallByVoiceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SingleCallByVoiceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SingleCallByVoiceRequest) GetSpeed() *int32 {
	return s.Speed
}

func (s *SingleCallByVoiceRequest) GetVoiceCode() *string {
	return s.VoiceCode
}

func (s *SingleCallByVoiceRequest) GetVolume() *int32 {
	return s.Volume
}

func (s *SingleCallByVoiceRequest) SetCalledNumber(v string) *SingleCallByVoiceRequest {
	s.CalledNumber = &v
	return s
}

func (s *SingleCallByVoiceRequest) SetCalledShowNumber(v string) *SingleCallByVoiceRequest {
	s.CalledShowNumber = &v
	return s
}

func (s *SingleCallByVoiceRequest) SetOutId(v string) *SingleCallByVoiceRequest {
	s.OutId = &v
	return s
}

func (s *SingleCallByVoiceRequest) SetOwnerId(v int64) *SingleCallByVoiceRequest {
	s.OwnerId = &v
	return s
}

func (s *SingleCallByVoiceRequest) SetPlayTimes(v int32) *SingleCallByVoiceRequest {
	s.PlayTimes = &v
	return s
}

func (s *SingleCallByVoiceRequest) SetResourceOwnerAccount(v string) *SingleCallByVoiceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SingleCallByVoiceRequest) SetResourceOwnerId(v int64) *SingleCallByVoiceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SingleCallByVoiceRequest) SetSpeed(v int32) *SingleCallByVoiceRequest {
	s.Speed = &v
	return s
}

func (s *SingleCallByVoiceRequest) SetVoiceCode(v string) *SingleCallByVoiceRequest {
	s.VoiceCode = &v
	return s
}

func (s *SingleCallByVoiceRequest) SetVolume(v int32) *SingleCallByVoiceRequest {
	s.Volume = &v
	return s
}

func (s *SingleCallByVoiceRequest) Validate() error {
	return dara.Validate(s)
}
