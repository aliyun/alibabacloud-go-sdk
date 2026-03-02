// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotUpdateJobStatus interface {
	dara.Model
	String() string
	GoString() string
	SetFailure(v *HotUpdateJobFailureInfo) *HotUpdateJobStatus
	GetFailure() *HotUpdateJobFailureInfo
	SetRequestId(v string) *HotUpdateJobStatus
	GetRequestId() *string
	SetStatus(v string) *HotUpdateJobStatus
	GetStatus() *string
}

type HotUpdateJobStatus struct {
	// The error message of the dynamical update.
	Failure *HotUpdateJobFailureInfo `json:"failure,omitempty" xml:"failure,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CBC799F0-AS7S-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The status of the dynamic update.
	//
	// example:
	//
	// HOT_UPDATED
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s HotUpdateJobStatus) String() string {
	return dara.Prettify(s)
}

func (s HotUpdateJobStatus) GoString() string {
	return s.String()
}

func (s *HotUpdateJobStatus) GetFailure() *HotUpdateJobFailureInfo {
	return s.Failure
}

func (s *HotUpdateJobStatus) GetRequestId() *string {
	return s.RequestId
}

func (s *HotUpdateJobStatus) GetStatus() *string {
	return s.Status
}

func (s *HotUpdateJobStatus) SetFailure(v *HotUpdateJobFailureInfo) *HotUpdateJobStatus {
	s.Failure = v
	return s
}

func (s *HotUpdateJobStatus) SetRequestId(v string) *HotUpdateJobStatus {
	s.RequestId = &v
	return s
}

func (s *HotUpdateJobStatus) SetStatus(v string) *HotUpdateJobStatus {
	s.Status = &v
	return s
}

func (s *HotUpdateJobStatus) Validate() error {
	if s.Failure != nil {
		if err := s.Failure.Validate(); err != nil {
			return err
		}
	}
	return nil
}
