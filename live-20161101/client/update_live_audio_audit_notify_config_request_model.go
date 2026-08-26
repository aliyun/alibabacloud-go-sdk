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
	// The webhook URL for receiving callback notifications about audio that contains violations.
	//
	// example:
	//
	// http://guide.aliyundoc.com/callback
	Callback *string `json:"Callback,omitempty" xml:"Callback,omitempty"`
	// The callback template for automated review. The value can contain the following variables:
	//
	// - **{DomainName}**: The streaming domain.
	//
	// - **{AppName}**: The AppName of the stream.
	//
	// - **{StreamName}**: The stream name.
	//
	// - **{Timestamp}**: The UNIX timestamp when the callback is generated. Unit: seconds.
	//
	// - **{Result}**: The detection result.
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
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
