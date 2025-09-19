// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSession interface {
	dara.Model
	String() string
	GoString() string
	SetContainerId(v string) *Session
	GetContainerId() *string
	SetCreatedTime(v string) *Session
	GetCreatedTime() *string
	SetFunctionName(v string) *Session
	GetFunctionName() *string
	SetLastModifiedTime(v string) *Session
	GetLastModifiedTime() *string
	SetQualifier(v string) *Session
	GetQualifier() *string
	SetSessionAffinityType(v string) *Session
	GetSessionAffinityType() *string
	SetSessionId(v string) *Session
	GetSessionId() *string
	SetSessionIdleTimeoutInSeconds(v int64) *Session
	GetSessionIdleTimeoutInSeconds() *int64
	SetSessionStatus(v string) *Session
	GetSessionStatus() *string
	SetSessionTTLInSeconds(v int64) *Session
	GetSessionTTLInSeconds() *int64
}

type Session struct {
	// example:
	//
	// c-68999e02-16a1955c-d2a03d1ccs
	ContainerId *string `json:"containerId,omitempty" xml:"containerId,omitempty"`
	// example:
	//
	// 2025-04-01T08:15:27Z
	CreatedTime *string `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	// example:
	//
	// functionName1
	FunctionName *string `json:"functionName,omitempty" xml:"functionName,omitempty"`
	// example:
	//
	// 2025-04-01T18:15:27Z
	LastModifiedTime *string `json:"lastModifiedTime,omitempty" xml:"lastModifiedTime,omitempty"`
	// example:
	//
	// AliasName1
	Qualifier *string `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
	// example:
	//
	// HEADER_FIELD
	SessionAffinityType *string `json:"sessionAffinityType,omitempty" xml:"sessionAffinityType,omitempty"`
	// example:
	//
	// 81f70ae156904eb9b7d43e12f511fe58
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// example:
	//
	// 1800
	SessionIdleTimeoutInSeconds *int64 `json:"sessionIdleTimeoutInSeconds,omitempty" xml:"sessionIdleTimeoutInSeconds,omitempty"`
	// example:
	//
	// Active
	SessionStatus *string `json:"sessionStatus,omitempty" xml:"sessionStatus,omitempty"`
	// example:
	//
	// 21600
	SessionTTLInSeconds *int64 `json:"sessionTTLInSeconds,omitempty" xml:"sessionTTLInSeconds,omitempty"`
}

func (s Session) String() string {
	return dara.Prettify(s)
}

func (s Session) GoString() string {
	return s.String()
}

func (s *Session) GetContainerId() *string {
	return s.ContainerId
}

func (s *Session) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *Session) GetFunctionName() *string {
	return s.FunctionName
}

func (s *Session) GetLastModifiedTime() *string {
	return s.LastModifiedTime
}

func (s *Session) GetQualifier() *string {
	return s.Qualifier
}

func (s *Session) GetSessionAffinityType() *string {
	return s.SessionAffinityType
}

func (s *Session) GetSessionId() *string {
	return s.SessionId
}

func (s *Session) GetSessionIdleTimeoutInSeconds() *int64 {
	return s.SessionIdleTimeoutInSeconds
}

func (s *Session) GetSessionStatus() *string {
	return s.SessionStatus
}

func (s *Session) GetSessionTTLInSeconds() *int64 {
	return s.SessionTTLInSeconds
}

func (s *Session) SetContainerId(v string) *Session {
	s.ContainerId = &v
	return s
}

func (s *Session) SetCreatedTime(v string) *Session {
	s.CreatedTime = &v
	return s
}

func (s *Session) SetFunctionName(v string) *Session {
	s.FunctionName = &v
	return s
}

func (s *Session) SetLastModifiedTime(v string) *Session {
	s.LastModifiedTime = &v
	return s
}

func (s *Session) SetQualifier(v string) *Session {
	s.Qualifier = &v
	return s
}

func (s *Session) SetSessionAffinityType(v string) *Session {
	s.SessionAffinityType = &v
	return s
}

func (s *Session) SetSessionId(v string) *Session {
	s.SessionId = &v
	return s
}

func (s *Session) SetSessionIdleTimeoutInSeconds(v int64) *Session {
	s.SessionIdleTimeoutInSeconds = &v
	return s
}

func (s *Session) SetSessionStatus(v string) *Session {
	s.SessionStatus = &v
	return s
}

func (s *Session) SetSessionTTLInSeconds(v int64) *Session {
	s.SessionTTLInSeconds = &v
	return s
}

func (s *Session) Validate() error {
	return dara.Validate(s)
}
