// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBrowserSessionOut interface {
	dara.Model
	String() string
	GoString() string
	SetBrowserId(v string) *BrowserSessionOut
	GetBrowserId() *string
	SetBrowserName(v string) *BrowserSessionOut
	GetBrowserName() *string
	SetCreatedAt(v string) *BrowserSessionOut
	GetCreatedAt() *string
	SetLastUpdatedAt(v string) *BrowserSessionOut
	GetLastUpdatedAt() *string
	SetSessionId(v string) *BrowserSessionOut
	GetSessionId() *string
	SetSessionIdleTimeoutSeconds(v int32) *BrowserSessionOut
	GetSessionIdleTimeoutSeconds() *int32
	SetStatus(v string) *BrowserSessionOut
	GetStatus() *string
}

type BrowserSessionOut struct {
	// example:
	//
	// browser-1234567890abcdef
	BrowserId *string `json:"browserId,omitempty" xml:"browserId,omitempty"`
	// example:
	//
	// my-browser-session
	BrowserName *string `json:"browserName,omitempty" xml:"browserName,omitempty"`
	// example:
	//
	// 2025-01-10T10:30:00Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// example:
	//
	// 2025-01-10T11:45:00Z
	LastUpdatedAt *string `json:"lastUpdatedAt,omitempty" xml:"lastUpdatedAt,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// bs-1234567890abcdef
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// 会话空闲超时时间，单位为秒
	//
	// example:
	//
	// 3600
	SessionIdleTimeoutSeconds *int32 `json:"sessionIdleTimeoutSeconds,omitempty" xml:"sessionIdleTimeoutSeconds,omitempty"`
	// example:
	//
	// RUNNING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s BrowserSessionOut) String() string {
	return dara.Prettify(s)
}

func (s BrowserSessionOut) GoString() string {
	return s.String()
}

func (s *BrowserSessionOut) GetBrowserId() *string {
	return s.BrowserId
}

func (s *BrowserSessionOut) GetBrowserName() *string {
	return s.BrowserName
}

func (s *BrowserSessionOut) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *BrowserSessionOut) GetLastUpdatedAt() *string {
	return s.LastUpdatedAt
}

func (s *BrowserSessionOut) GetSessionId() *string {
	return s.SessionId
}

func (s *BrowserSessionOut) GetSessionIdleTimeoutSeconds() *int32 {
	return s.SessionIdleTimeoutSeconds
}

func (s *BrowserSessionOut) GetStatus() *string {
	return s.Status
}

func (s *BrowserSessionOut) SetBrowserId(v string) *BrowserSessionOut {
	s.BrowserId = &v
	return s
}

func (s *BrowserSessionOut) SetBrowserName(v string) *BrowserSessionOut {
	s.BrowserName = &v
	return s
}

func (s *BrowserSessionOut) SetCreatedAt(v string) *BrowserSessionOut {
	s.CreatedAt = &v
	return s
}

func (s *BrowserSessionOut) SetLastUpdatedAt(v string) *BrowserSessionOut {
	s.LastUpdatedAt = &v
	return s
}

func (s *BrowserSessionOut) SetSessionId(v string) *BrowserSessionOut {
	s.SessionId = &v
	return s
}

func (s *BrowserSessionOut) SetSessionIdleTimeoutSeconds(v int32) *BrowserSessionOut {
	s.SessionIdleTimeoutSeconds = &v
	return s
}

func (s *BrowserSessionOut) SetStatus(v string) *BrowserSessionOut {
	s.Status = &v
	return s
}

func (s *BrowserSessionOut) Validate() error {
	return dara.Validate(s)
}
