// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddLiveAudioAuditNotifyConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallback(v string) *AddLiveAudioAuditNotifyConfigRequest
	GetCallback() *string
	SetCallbackTemplate(v string) *AddLiveAudioAuditNotifyConfigRequest
	GetCallbackTemplate() *string
	SetDomainName(v string) *AddLiveAudioAuditNotifyConfigRequest
	GetDomainName() *string
	SetOwnerId(v int64) *AddLiveAudioAuditNotifyConfigRequest
	GetOwnerId() *int64
	SetRegionId(v string) *AddLiveAudioAuditNotifyConfigRequest
	GetRegionId() *string
}

type AddLiveAudioAuditNotifyConfigRequest struct {
	// The callback URL. This URL is used to receive callback notifications about violations in audio.
	//
	// example:
	//
	// https://demo.aliyundoc.com/callback
	Callback *string `json:"Callback,omitempty" xml:"Callback,omitempty"`
	// The callback template. Valid values:
	//
	// 	- **{DomainName}**: the streaming domain.
	//
	// 	- **{AppName}**: the name of the application to which the live stream belongs.
	//
	// 	- **{StreamName}**: the name of the live stream.
	//
	// 	- **{Timestamp}**: the time when the callback is returned. The value of this field is a UNIX timestamp. Unit: seconds.
	//
	// 	- **{Result}**: the moderation results.
	//
	// example:
	//
	// {"domain":{DomainName},"app":{AppName},"stream":{StreamName},"timestamp":{Timestamp},"result":{Result}}
	CallbackTemplate *string `json:"CallbackTemplate,omitempty" xml:"CallbackTemplate,omitempty"`
	// The main streaming domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AddLiveAudioAuditNotifyConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s AddLiveAudioAuditNotifyConfigRequest) GoString() string {
	return s.String()
}

func (s *AddLiveAudioAuditNotifyConfigRequest) GetCallback() *string {
	return s.Callback
}

func (s *AddLiveAudioAuditNotifyConfigRequest) GetCallbackTemplate() *string {
	return s.CallbackTemplate
}

func (s *AddLiveAudioAuditNotifyConfigRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *AddLiveAudioAuditNotifyConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddLiveAudioAuditNotifyConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AddLiveAudioAuditNotifyConfigRequest) SetCallback(v string) *AddLiveAudioAuditNotifyConfigRequest {
	s.Callback = &v
	return s
}

func (s *AddLiveAudioAuditNotifyConfigRequest) SetCallbackTemplate(v string) *AddLiveAudioAuditNotifyConfigRequest {
	s.CallbackTemplate = &v
	return s
}

func (s *AddLiveAudioAuditNotifyConfigRequest) SetDomainName(v string) *AddLiveAudioAuditNotifyConfigRequest {
	s.DomainName = &v
	return s
}

func (s *AddLiveAudioAuditNotifyConfigRequest) SetOwnerId(v int64) *AddLiveAudioAuditNotifyConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *AddLiveAudioAuditNotifyConfigRequest) SetRegionId(v string) *AddLiveAudioAuditNotifyConfigRequest {
	s.RegionId = &v
	return s
}

func (s *AddLiveAudioAuditNotifyConfigRequest) Validate() error {
	return dara.Validate(s)
}
