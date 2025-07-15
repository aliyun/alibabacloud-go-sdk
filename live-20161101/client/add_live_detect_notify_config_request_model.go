// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddLiveDetectNotifyConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *AddLiveDetectNotifyConfigRequest
	GetDomainName() *string
	SetNotifyUrl(v string) *AddLiveDetectNotifyConfigRequest
	GetNotifyUrl() *string
	SetOwnerId(v int64) *AddLiveDetectNotifyConfigRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *AddLiveDetectNotifyConfigRequest
	GetSecurityToken() *string
}

type AddLiveDetectNotifyConfigRequest struct {
	// The main streaming domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The callback URL that is used to receive callback notifications about violations, such as pornographic content and politically sensitive content, detected in live streams.
	//
	// This parameter is required.
	//
	// example:
	//
	// http://demo.aliyundoc.com/examplecallback.action
	NotifyUrl     *string `json:"NotifyUrl,omitempty" xml:"NotifyUrl,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s AddLiveDetectNotifyConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s AddLiveDetectNotifyConfigRequest) GoString() string {
	return s.String()
}

func (s *AddLiveDetectNotifyConfigRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *AddLiveDetectNotifyConfigRequest) GetNotifyUrl() *string {
	return s.NotifyUrl
}

func (s *AddLiveDetectNotifyConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddLiveDetectNotifyConfigRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *AddLiveDetectNotifyConfigRequest) SetDomainName(v string) *AddLiveDetectNotifyConfigRequest {
	s.DomainName = &v
	return s
}

func (s *AddLiveDetectNotifyConfigRequest) SetNotifyUrl(v string) *AddLiveDetectNotifyConfigRequest {
	s.NotifyUrl = &v
	return s
}

func (s *AddLiveDetectNotifyConfigRequest) SetOwnerId(v int64) *AddLiveDetectNotifyConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *AddLiveDetectNotifyConfigRequest) SetSecurityToken(v string) *AddLiveDetectNotifyConfigRequest {
	s.SecurityToken = &v
	return s
}

func (s *AddLiveDetectNotifyConfigRequest) Validate() error {
	return dara.Validate(s)
}
