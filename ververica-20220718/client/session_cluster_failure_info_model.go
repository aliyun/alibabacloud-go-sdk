// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSessionClusterFailureInfo interface {
	dara.Model
	String() string
	GoString() string
	SetFailedAt(v int64) *SessionClusterFailureInfo
	GetFailedAt() *int64
	SetMessage(v string) *SessionClusterFailureInfo
	GetMessage() *string
	SetReason(v string) *SessionClusterFailureInfo
	GetReason() *string
}

type SessionClusterFailureInfo struct {
	FailedAt *int64  `json:"failedAt,omitempty" xml:"failedAt,omitempty"`
	Message  *string `json:"message,omitempty" xml:"message,omitempty"`
	Reason   *string `json:"reason,omitempty" xml:"reason,omitempty"`
}

func (s SessionClusterFailureInfo) String() string {
	return dara.Prettify(s)
}

func (s SessionClusterFailureInfo) GoString() string {
	return s.String()
}

func (s *SessionClusterFailureInfo) GetFailedAt() *int64 {
	return s.FailedAt
}

func (s *SessionClusterFailureInfo) GetMessage() *string {
	return s.Message
}

func (s *SessionClusterFailureInfo) GetReason() *string {
	return s.Reason
}

func (s *SessionClusterFailureInfo) SetFailedAt(v int64) *SessionClusterFailureInfo {
	s.FailedAt = &v
	return s
}

func (s *SessionClusterFailureInfo) SetMessage(v string) *SessionClusterFailureInfo {
	s.Message = &v
	return s
}

func (s *SessionClusterFailureInfo) SetReason(v string) *SessionClusterFailureInfo {
	s.Reason = &v
	return s
}

func (s *SessionClusterFailureInfo) Validate() error {
	return dara.Validate(s)
}
