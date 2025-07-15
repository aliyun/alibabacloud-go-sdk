// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLiveAudioAuditNotifyConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallback(v string) *UpdateLiveAudioAuditNotifyConfigRequest
	GetCallback() *string
	SetCallbackTemplate(v string) *UpdateLiveAudioAuditNotifyConfigRequest
	GetCallbackTemplate() *string
	SetDomainName(v string) *UpdateLiveAudioAuditNotifyConfigRequest
	GetDomainName() *string
	SetOwnerId(v int64) *UpdateLiveAudioAuditNotifyConfigRequest
	GetOwnerId() *int64
	SetRegionId(v string) *UpdateLiveAudioAuditNotifyConfigRequest
	GetRegionId() *string
}

type UpdateLiveAudioAuditNotifyConfigRequest struct {
	// The callback URL. This URL is used to receive callback notifications about violations in audio.
	//
	// example:
	//
	// http://guide.aliyundoc.com/callback
	Callback *string `json:"Callback,omitempty" xml:"Callback,omitempty"`
	// The callback template. Configure the following fields:
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

func (s UpdateLiveAudioAuditNotifyConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateLiveAudioAuditNotifyConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateLiveAudioAuditNotifyConfigRequest) GetCallback() *string {
	return s.Callback
}

func (s *UpdateLiveAudioAuditNotifyConfigRequest) GetCallbackTemplate() *string {
	return s.CallbackTemplate
}

func (s *UpdateLiveAudioAuditNotifyConfigRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *UpdateLiveAudioAuditNotifyConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateLiveAudioAuditNotifyConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateLiveAudioAuditNotifyConfigRequest) SetCallback(v string) *UpdateLiveAudioAuditNotifyConfigRequest {
	s.Callback = &v
	return s
}

func (s *UpdateLiveAudioAuditNotifyConfigRequest) SetCallbackTemplate(v string) *UpdateLiveAudioAuditNotifyConfigRequest {
	s.CallbackTemplate = &v
	return s
}

func (s *UpdateLiveAudioAuditNotifyConfigRequest) SetDomainName(v string) *UpdateLiveAudioAuditNotifyConfigRequest {
	s.DomainName = &v
	return s
}

func (s *UpdateLiveAudioAuditNotifyConfigRequest) SetOwnerId(v int64) *UpdateLiveAudioAuditNotifyConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateLiveAudioAuditNotifyConfigRequest) SetRegionId(v string) *UpdateLiveAudioAuditNotifyConfigRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateLiveAudioAuditNotifyConfigRequest) Validate() error {
	return dara.Validate(s)
}
