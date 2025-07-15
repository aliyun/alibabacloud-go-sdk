// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetLiveStreamsNotifyUrlConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *SetLiveStreamsNotifyUrlConfigRequest
	GetDomainName() *string
	SetExceptionNotifyUrl(v string) *SetLiveStreamsNotifyUrlConfigRequest
	GetExceptionNotifyUrl() *string
	SetNotifyAuthKey(v string) *SetLiveStreamsNotifyUrlConfigRequest
	GetNotifyAuthKey() *string
	SetNotifyReqAuth(v string) *SetLiveStreamsNotifyUrlConfigRequest
	GetNotifyReqAuth() *string
	SetNotifyUrl(v string) *SetLiveStreamsNotifyUrlConfigRequest
	GetNotifyUrl() *string
	SetOwnerId(v int64) *SetLiveStreamsNotifyUrlConfigRequest
	GetOwnerId() *int64
	SetSwitchNotifyUrl(v string) *SetLiveStreamsNotifyUrlConfigRequest
	GetSwitchNotifyUrl() *string
}

type SetLiveStreamsNotifyUrlConfigRequest struct {
	// The ingest domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// demo.aliyundoc.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// Exception event callback URL.
	//
	// example:
	//
	// https://4a7e5f08.r37.cpolar.top/live/Rsssd/call-back/streamStart
	ExceptionNotifyUrl *string `json:"ExceptionNotifyUrl,omitempty" xml:"ExceptionNotifyUrl,omitempty"`
	// The authentication key.
	//
	// >  This parameter is required if you set the NotifyReqAuth parameter to **yes**.
	//
	// Value requirements:
	//
	// 	- The key must be 16 to 64 characters in length.
	//
	// 	- The key can contain letters and digits.
	//
	// example:
	//
	// 123456
	NotifyAuthKey *string `json:"NotifyAuthKey,omitempty" xml:"NotifyAuthKey,omitempty"`
	// Specifies whether to enable callback authentication. Valid values:
	//
	// 	- **yes**: enables callback authentication. If you set this parameter to **yes**, you must also specify the NotifyAuthKey parameter.
	//
	// 	- **no**: disables callback authentication.
	//
	// >  If you do not specify this parameter, the default value **no*	- is used.
	//
	// For information about the authentication logic, see **Authentication for stream ingest callbacks**.
	//
	// example:
	//
	// yes
	NotifyReqAuth *string `json:"NotifyReqAuth,omitempty" xml:"NotifyReqAuth,omitempty"`
	// The URL to which the stream ingest callbacks are sent.
	//
	// example:
	//
	// http://guide.aliyundoc.com/notify
	NotifyUrl       *string `json:"NotifyUrl,omitempty" xml:"NotifyUrl,omitempty"`
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SwitchNotifyUrl *string `json:"SwitchNotifyUrl,omitempty" xml:"SwitchNotifyUrl,omitempty"`
}

func (s SetLiveStreamsNotifyUrlConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s SetLiveStreamsNotifyUrlConfigRequest) GoString() string {
	return s.String()
}

func (s *SetLiveStreamsNotifyUrlConfigRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *SetLiveStreamsNotifyUrlConfigRequest) GetExceptionNotifyUrl() *string {
	return s.ExceptionNotifyUrl
}

func (s *SetLiveStreamsNotifyUrlConfigRequest) GetNotifyAuthKey() *string {
	return s.NotifyAuthKey
}

func (s *SetLiveStreamsNotifyUrlConfigRequest) GetNotifyReqAuth() *string {
	return s.NotifyReqAuth
}

func (s *SetLiveStreamsNotifyUrlConfigRequest) GetNotifyUrl() *string {
	return s.NotifyUrl
}

func (s *SetLiveStreamsNotifyUrlConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SetLiveStreamsNotifyUrlConfigRequest) GetSwitchNotifyUrl() *string {
	return s.SwitchNotifyUrl
}

func (s *SetLiveStreamsNotifyUrlConfigRequest) SetDomainName(v string) *SetLiveStreamsNotifyUrlConfigRequest {
	s.DomainName = &v
	return s
}

func (s *SetLiveStreamsNotifyUrlConfigRequest) SetExceptionNotifyUrl(v string) *SetLiveStreamsNotifyUrlConfigRequest {
	s.ExceptionNotifyUrl = &v
	return s
}

func (s *SetLiveStreamsNotifyUrlConfigRequest) SetNotifyAuthKey(v string) *SetLiveStreamsNotifyUrlConfigRequest {
	s.NotifyAuthKey = &v
	return s
}

func (s *SetLiveStreamsNotifyUrlConfigRequest) SetNotifyReqAuth(v string) *SetLiveStreamsNotifyUrlConfigRequest {
	s.NotifyReqAuth = &v
	return s
}

func (s *SetLiveStreamsNotifyUrlConfigRequest) SetNotifyUrl(v string) *SetLiveStreamsNotifyUrlConfigRequest {
	s.NotifyUrl = &v
	return s
}

func (s *SetLiveStreamsNotifyUrlConfigRequest) SetOwnerId(v int64) *SetLiveStreamsNotifyUrlConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *SetLiveStreamsNotifyUrlConfigRequest) SetSwitchNotifyUrl(v string) *SetLiveStreamsNotifyUrlConfigRequest {
	s.SwitchNotifyUrl = &v
	return s
}

func (s *SetLiveStreamsNotifyUrlConfigRequest) Validate() error {
	return dara.Validate(s)
}
