// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublicTemplateStatus interface {
	dara.Model
	String() string
	GoString() string
	SetFinishedAt(v string) *PublicTemplateStatus
	GetFinishedAt() *string
	SetReason(v *PublicTemplateStatusReason) *PublicTemplateStatus
	GetReason() *PublicTemplateStatusReason
	SetState(v string) *PublicTemplateStatus
	GetState() *string
}

type PublicTemplateStatus struct {
	// The time when the build is completed.
	//
	// example:
	//
	// 2026-08-28T12:00:00.000Z
	FinishedAt *string `json:"finishedAt,omitempty" xml:"finishedAt,omitempty"`
	// The reason for the build failure.
	Reason *PublicTemplateStatusReason `json:"reason,omitempty" xml:"reason,omitempty"`
	// The build status.
	//
	// example:
	//
	// ready
	State *string `json:"state,omitempty" xml:"state,omitempty"`
}

func (s PublicTemplateStatus) String() string {
	return dara.Prettify(s)
}

func (s PublicTemplateStatus) GoString() string {
	return s.String()
}

func (s *PublicTemplateStatus) GetFinishedAt() *string {
	return s.FinishedAt
}

func (s *PublicTemplateStatus) GetReason() *PublicTemplateStatusReason {
	return s.Reason
}

func (s *PublicTemplateStatus) GetState() *string {
	return s.State
}

func (s *PublicTemplateStatus) SetFinishedAt(v string) *PublicTemplateStatus {
	s.FinishedAt = &v
	return s
}

func (s *PublicTemplateStatus) SetReason(v *PublicTemplateStatusReason) *PublicTemplateStatus {
	s.Reason = v
	return s
}

func (s *PublicTemplateStatus) SetState(v string) *PublicTemplateStatus {
	s.State = &v
	return s
}

func (s *PublicTemplateStatus) Validate() error {
	if s.Reason != nil {
		if err := s.Reason.Validate(); err != nil {
			return err
		}
	}
	return nil
}
