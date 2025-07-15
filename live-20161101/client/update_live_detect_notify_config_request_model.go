// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLiveDetectNotifyConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *UpdateLiveDetectNotifyConfigRequest
	GetDomainName() *string
	SetNotifyUrl(v string) *UpdateLiveDetectNotifyConfigRequest
	GetNotifyUrl() *string
	SetOwnerId(v int64) *UpdateLiveDetectNotifyConfigRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *UpdateLiveDetectNotifyConfigRequest
	GetSecurityToken() *string
}

type UpdateLiveDetectNotifyConfigRequest struct {
	// The main streaming domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The callback URL that is used to receive callback notifications about violations, such as pornographic content, detected in live streams.
	//
	// This parameter is required.
	//
	// example:
	//
	// http://guide.aliyundoc.com/examplecallback.action
	NotifyUrl     *string `json:"NotifyUrl,omitempty" xml:"NotifyUrl,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s UpdateLiveDetectNotifyConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateLiveDetectNotifyConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateLiveDetectNotifyConfigRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *UpdateLiveDetectNotifyConfigRequest) GetNotifyUrl() *string {
	return s.NotifyUrl
}

func (s *UpdateLiveDetectNotifyConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateLiveDetectNotifyConfigRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *UpdateLiveDetectNotifyConfigRequest) SetDomainName(v string) *UpdateLiveDetectNotifyConfigRequest {
	s.DomainName = &v
	return s
}

func (s *UpdateLiveDetectNotifyConfigRequest) SetNotifyUrl(v string) *UpdateLiveDetectNotifyConfigRequest {
	s.NotifyUrl = &v
	return s
}

func (s *UpdateLiveDetectNotifyConfigRequest) SetOwnerId(v int64) *UpdateLiveDetectNotifyConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateLiveDetectNotifyConfigRequest) SetSecurityToken(v string) *UpdateLiveDetectNotifyConfigRequest {
	s.SecurityToken = &v
	return s
}

func (s *UpdateLiveDetectNotifyConfigRequest) Validate() error {
	return dara.Validate(s)
}
