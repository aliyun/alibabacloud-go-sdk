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
	// The URL used to receive callbacks for audio that contains problematic content.
	//
	// example:
	//
	// https://demo.aliyundoc.com/callback
	Callback *string `json:"Callback,omitempty" xml:"Callback,omitempty"`
	// The callback template for automated reviews. The value is a JSON string that can contain the following variables:
	//
	// - **{DomainName}**: The streaming domain.
	//
	// - **{AppName}**: The AppName of the live stream.
	//
	// - **{StreamName}**: The stream name.
	//
	// - **{Timestamp}**: The UNIX timestamp of the callback. Unit: seconds.
	//
	// - **{Result}**: The review result.
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
