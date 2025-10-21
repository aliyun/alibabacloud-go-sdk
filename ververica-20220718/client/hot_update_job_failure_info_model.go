// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotUpdateJobFailureInfo interface {
	dara.Model
	String() string
	GoString() string
	SetFailureSeverity(v string) *HotUpdateJobFailureInfo
	GetFailureSeverity() *string
	SetMessage(v string) *HotUpdateJobFailureInfo
	GetMessage() *string
	SetReason(v string) *HotUpdateJobFailureInfo
	GetReason() *string
}

type HotUpdateJobFailureInfo struct {
	FailureSeverity *string `json:"failureSeverity,omitempty" xml:"failureSeverity,omitempty"`
	Message         *string `json:"message,omitempty" xml:"message,omitempty"`
	Reason          *string `json:"reason,omitempty" xml:"reason,omitempty"`
}

func (s HotUpdateJobFailureInfo) String() string {
	return dara.Prettify(s)
}

func (s HotUpdateJobFailureInfo) GoString() string {
	return s.String()
}

func (s *HotUpdateJobFailureInfo) GetFailureSeverity() *string {
	return s.FailureSeverity
}

func (s *HotUpdateJobFailureInfo) GetMessage() *string {
	return s.Message
}

func (s *HotUpdateJobFailureInfo) GetReason() *string {
	return s.Reason
}

func (s *HotUpdateJobFailureInfo) SetFailureSeverity(v string) *HotUpdateJobFailureInfo {
	s.FailureSeverity = &v
	return s
}

func (s *HotUpdateJobFailureInfo) SetMessage(v string) *HotUpdateJobFailureInfo {
	s.Message = &v
	return s
}

func (s *HotUpdateJobFailureInfo) SetReason(v string) *HotUpdateJobFailureInfo {
	s.Reason = &v
	return s
}

func (s *HotUpdateJobFailureInfo) Validate() error {
	return dara.Validate(s)
}
