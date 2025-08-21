// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHttpConfig interface {
	dara.Model
	String() string
	GoString() string
	SetCookie(v string) *HttpConfig
	GetCookie() *string
	SetCookieTimeout(v int32) *HttpConfig
	GetCookieTimeout() *int32
	SetIdleTimeout(v int32) *HttpConfig
	GetIdleTimeout() *int32
	SetRequestTimeout(v int32) *HttpConfig
	GetRequestTimeout() *int32
	SetScheduler(v string) *HttpConfig
	GetScheduler() *string
	SetServerCertificateId(v string) *HttpConfig
	GetServerCertificateId() *string
	SetStickySession(v string) *HttpConfig
	GetStickySession() *string
	SetStickySessionType(v string) *HttpConfig
	GetStickySessionType() *string
	SetXForwardedFor(v string) *HttpConfig
	GetXForwardedFor() *string
}

type HttpConfig struct {
	Cookie *string `json:"Cookie,omitempty" xml:"Cookie,omitempty"`
	// example:
	//
	// 500
	CookieTimeout       *int32  `json:"CookieTimeout,omitempty" xml:"CookieTimeout,omitempty"`
	IdleTimeout         *int32  `json:"IdleTimeout,omitempty" xml:"IdleTimeout,omitempty"`
	RequestTimeout      *int32  `json:"RequestTimeout,omitempty" xml:"RequestTimeout,omitempty"`
	Scheduler           *string `json:"Scheduler,omitempty" xml:"Scheduler,omitempty"`
	ServerCertificateId *string `json:"ServerCertificateId,omitempty" xml:"ServerCertificateId,omitempty"`
	// if can be null:
	// false
	StickySession *string `json:"StickySession,omitempty" xml:"StickySession,omitempty"`
	// example:
	//
	// insert
	StickySessionType *string `json:"StickySessionType,omitempty" xml:"StickySessionType,omitempty"`
	// example:
	//
	// on
	XForwardedFor *string `json:"XForwardedFor,omitempty" xml:"XForwardedFor,omitempty"`
}

func (s HttpConfig) String() string {
	return dara.Prettify(s)
}

func (s HttpConfig) GoString() string {
	return s.String()
}

func (s *HttpConfig) GetCookie() *string {
	return s.Cookie
}

func (s *HttpConfig) GetCookieTimeout() *int32 {
	return s.CookieTimeout
}

func (s *HttpConfig) GetIdleTimeout() *int32 {
	return s.IdleTimeout
}

func (s *HttpConfig) GetRequestTimeout() *int32 {
	return s.RequestTimeout
}

func (s *HttpConfig) GetScheduler() *string {
	return s.Scheduler
}

func (s *HttpConfig) GetServerCertificateId() *string {
	return s.ServerCertificateId
}

func (s *HttpConfig) GetStickySession() *string {
	return s.StickySession
}

func (s *HttpConfig) GetStickySessionType() *string {
	return s.StickySessionType
}

func (s *HttpConfig) GetXForwardedFor() *string {
	return s.XForwardedFor
}

func (s *HttpConfig) SetCookie(v string) *HttpConfig {
	s.Cookie = &v
	return s
}

func (s *HttpConfig) SetCookieTimeout(v int32) *HttpConfig {
	s.CookieTimeout = &v
	return s
}

func (s *HttpConfig) SetIdleTimeout(v int32) *HttpConfig {
	s.IdleTimeout = &v
	return s
}

func (s *HttpConfig) SetRequestTimeout(v int32) *HttpConfig {
	s.RequestTimeout = &v
	return s
}

func (s *HttpConfig) SetScheduler(v string) *HttpConfig {
	s.Scheduler = &v
	return s
}

func (s *HttpConfig) SetServerCertificateId(v string) *HttpConfig {
	s.ServerCertificateId = &v
	return s
}

func (s *HttpConfig) SetStickySession(v string) *HttpConfig {
	s.StickySession = &v
	return s
}

func (s *HttpConfig) SetStickySessionType(v string) *HttpConfig {
	s.StickySessionType = &v
	return s
}

func (s *HttpConfig) SetXForwardedFor(v string) *HttpConfig {
	s.XForwardedFor = &v
	return s
}

func (s *HttpConfig) Validate() error {
	return dara.Validate(s)
}
