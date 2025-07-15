// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLiveRecordNotifyConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *UpdateLiveRecordNotifyConfigRequest
	GetDomainName() *string
	SetNeedStatusNotify(v bool) *UpdateLiveRecordNotifyConfigRequest
	GetNeedStatusNotify() *bool
	SetNotifyUrl(v string) *UpdateLiveRecordNotifyConfigRequest
	GetNotifyUrl() *string
	SetOnDemandUrl(v string) *UpdateLiveRecordNotifyConfigRequest
	GetOnDemandUrl() *string
	SetOwnerId(v int64) *UpdateLiveRecordNotifyConfigRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *UpdateLiveRecordNotifyConfigRequest
	GetSecurityToken() *string
}

type UpdateLiveRecordNotifyConfigRequest struct {
	// The main streaming domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// Specifies whether to enable callbacks for recording status. Valid values:
	//
	// 	- **true**
	//
	// 	- **false*	- (default)
	//
	// example:
	//
	// false
	NeedStatusNotify *bool `json:"NeedStatusNotify,omitempty" xml:"NeedStatusNotify,omitempty"`
	// The callback URL that is used to receive notifications about recording events and status.
	//
	// >
	//
	// 	- The URL must start with `http://` or `https://`.
	//
	// 	- You must use URLEncoder for encoding. This way, the system can identify Chinese characters, spaces, and special characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// http://learn.aliyundoc.com/examplecallback.action
	NotifyUrl *string `json:"NotifyUrl,omitempty" xml:"NotifyUrl,omitempty"`
	// The callback URL for on-demand recordings.
	//
	// >
	//
	// 	- The URL must start with `http://` or `https://`.
	//
	// 	- You must use URLEncoder for encoding. This way, the system can identify Chinese characters, spaces, and special characters.
	//
	// example:
	//
	// http://guide.aliyundoc.com/ondemandcallback.action
	OnDemandUrl   *string `json:"OnDemandUrl,omitempty" xml:"OnDemandUrl,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s UpdateLiveRecordNotifyConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateLiveRecordNotifyConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateLiveRecordNotifyConfigRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *UpdateLiveRecordNotifyConfigRequest) GetNeedStatusNotify() *bool {
	return s.NeedStatusNotify
}

func (s *UpdateLiveRecordNotifyConfigRequest) GetNotifyUrl() *string {
	return s.NotifyUrl
}

func (s *UpdateLiveRecordNotifyConfigRequest) GetOnDemandUrl() *string {
	return s.OnDemandUrl
}

func (s *UpdateLiveRecordNotifyConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateLiveRecordNotifyConfigRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *UpdateLiveRecordNotifyConfigRequest) SetDomainName(v string) *UpdateLiveRecordNotifyConfigRequest {
	s.DomainName = &v
	return s
}

func (s *UpdateLiveRecordNotifyConfigRequest) SetNeedStatusNotify(v bool) *UpdateLiveRecordNotifyConfigRequest {
	s.NeedStatusNotify = &v
	return s
}

func (s *UpdateLiveRecordNotifyConfigRequest) SetNotifyUrl(v string) *UpdateLiveRecordNotifyConfigRequest {
	s.NotifyUrl = &v
	return s
}

func (s *UpdateLiveRecordNotifyConfigRequest) SetOnDemandUrl(v string) *UpdateLiveRecordNotifyConfigRequest {
	s.OnDemandUrl = &v
	return s
}

func (s *UpdateLiveRecordNotifyConfigRequest) SetOwnerId(v int64) *UpdateLiveRecordNotifyConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateLiveRecordNotifyConfigRequest) SetSecurityToken(v string) *UpdateLiveRecordNotifyConfigRequest {
	s.SecurityToken = &v
	return s
}

func (s *UpdateLiveRecordNotifyConfigRequest) Validate() error {
	return dara.Validate(s)
}
