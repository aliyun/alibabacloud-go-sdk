// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCodeInterpreterSessionOut interface {
	dara.Model
	String() string
	GoString() string
	SetCodeInterpreterId(v string) *CodeInterpreterSessionOut
	GetCodeInterpreterId() *string
	SetCodeInterpreterName(v string) *CodeInterpreterSessionOut
	GetCodeInterpreterName() *string
	SetCreatedAt(v string) *CodeInterpreterSessionOut
	GetCreatedAt() *string
	SetLastUpdatedAt(v string) *CodeInterpreterSessionOut
	GetLastUpdatedAt() *string
	SetSessionId(v string) *CodeInterpreterSessionOut
	GetSessionId() *string
	SetSessionIdleTimeoutSeconds(v int32) *CodeInterpreterSessionOut
	GetSessionIdleTimeoutSeconds() *int32
	SetStatus(v string) *CodeInterpreterSessionOut
	GetStatus() *string
}

type CodeInterpreterSessionOut struct {
	// The Unique Identifier of the associated code interpreter
	//
	// This parameter is required.
	CodeInterpreterId *string `json:"codeInterpreterId,omitempty" xml:"codeInterpreterId,omitempty"`
	// The name of the code interpreter session
	CodeInterpreterName *string `json:"codeInterpreterName,omitempty" xml:"codeInterpreterName,omitempty"`
	// The creation time of the code interpreter session, in ISO 8601 format
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The last update time of the code interpreter session, in ISO 8601 format
	LastUpdatedAt *string `json:"lastUpdatedAt,omitempty" xml:"lastUpdatedAt,omitempty"`
	// The Unique Identifier of the code interpreter session
	//
	// This parameter is required.
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// The idle timeout duration of the code interpreter session, in seconds. After the instance receives no session requests, it enters an idle state, which is billed under the idle billing model. If the idle duration exceeds this timeout, the session automatically expires and can no longer be used.
	SessionIdleTimeoutSeconds *int32 `json:"sessionIdleTimeoutSeconds,omitempty" xml:"sessionIdleTimeoutSeconds,omitempty"`
	// The current status of the code interpreter session
	//
	// This parameter is required.
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s CodeInterpreterSessionOut) String() string {
	return dara.Prettify(s)
}

func (s CodeInterpreterSessionOut) GoString() string {
	return s.String()
}

func (s *CodeInterpreterSessionOut) GetCodeInterpreterId() *string {
	return s.CodeInterpreterId
}

func (s *CodeInterpreterSessionOut) GetCodeInterpreterName() *string {
	return s.CodeInterpreterName
}

func (s *CodeInterpreterSessionOut) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *CodeInterpreterSessionOut) GetLastUpdatedAt() *string {
	return s.LastUpdatedAt
}

func (s *CodeInterpreterSessionOut) GetSessionId() *string {
	return s.SessionId
}

func (s *CodeInterpreterSessionOut) GetSessionIdleTimeoutSeconds() *int32 {
	return s.SessionIdleTimeoutSeconds
}

func (s *CodeInterpreterSessionOut) GetStatus() *string {
	return s.Status
}

func (s *CodeInterpreterSessionOut) SetCodeInterpreterId(v string) *CodeInterpreterSessionOut {
	s.CodeInterpreterId = &v
	return s
}

func (s *CodeInterpreterSessionOut) SetCodeInterpreterName(v string) *CodeInterpreterSessionOut {
	s.CodeInterpreterName = &v
	return s
}

func (s *CodeInterpreterSessionOut) SetCreatedAt(v string) *CodeInterpreterSessionOut {
	s.CreatedAt = &v
	return s
}

func (s *CodeInterpreterSessionOut) SetLastUpdatedAt(v string) *CodeInterpreterSessionOut {
	s.LastUpdatedAt = &v
	return s
}

func (s *CodeInterpreterSessionOut) SetSessionId(v string) *CodeInterpreterSessionOut {
	s.SessionId = &v
	return s
}

func (s *CodeInterpreterSessionOut) SetSessionIdleTimeoutSeconds(v int32) *CodeInterpreterSessionOut {
	s.SessionIdleTimeoutSeconds = &v
	return s
}

func (s *CodeInterpreterSessionOut) SetStatus(v string) *CodeInterpreterSessionOut {
	s.Status = &v
	return s
}

func (s *CodeInterpreterSessionOut) Validate() error {
	return dara.Validate(s)
}
