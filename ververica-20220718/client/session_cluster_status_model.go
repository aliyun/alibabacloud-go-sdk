// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSessionClusterStatus interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentSessionClusterStatus(v string) *SessionClusterStatus
	GetCurrentSessionClusterStatus() *string
	SetFailure(v *SessionClusterFailureInfo) *SessionClusterStatus
	GetFailure() *SessionClusterFailureInfo
	SetRunning(v *SessionClusterRunningInfo) *SessionClusterStatus
	GetRunning() *SessionClusterRunningInfo
}

type SessionClusterStatus struct {
	CurrentSessionClusterStatus *string                    `json:"currentSessionClusterStatus,omitempty" xml:"currentSessionClusterStatus,omitempty"`
	Failure                     *SessionClusterFailureInfo `json:"failure,omitempty" xml:"failure,omitempty"`
	Running                     *SessionClusterRunningInfo `json:"running,omitempty" xml:"running,omitempty"`
}

func (s SessionClusterStatus) String() string {
	return dara.Prettify(s)
}

func (s SessionClusterStatus) GoString() string {
	return s.String()
}

func (s *SessionClusterStatus) GetCurrentSessionClusterStatus() *string {
	return s.CurrentSessionClusterStatus
}

func (s *SessionClusterStatus) GetFailure() *SessionClusterFailureInfo {
	return s.Failure
}

func (s *SessionClusterStatus) GetRunning() *SessionClusterRunningInfo {
	return s.Running
}

func (s *SessionClusterStatus) SetCurrentSessionClusterStatus(v string) *SessionClusterStatus {
	s.CurrentSessionClusterStatus = &v
	return s
}

func (s *SessionClusterStatus) SetFailure(v *SessionClusterFailureInfo) *SessionClusterStatus {
	s.Failure = v
	return s
}

func (s *SessionClusterStatus) SetRunning(v *SessionClusterRunningInfo) *SessionClusterStatus {
	s.Running = v
	return s
}

func (s *SessionClusterStatus) Validate() error {
	if s.Failure != nil {
		if err := s.Failure.Validate(); err != nil {
			return err
		}
	}
	if s.Running != nil {
		if err := s.Running.Validate(); err != nil {
			return err
		}
	}
	return nil
}
