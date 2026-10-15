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
	// The time when the template reached the current desired state, in UTC and RFC 3339 format. This parameter is empty when the template is still being built or updated.
	//
	// example:
	//
	// 2026-08-20T08:35:12Z
	FinishedAt *string `json:"finishedAt,omitempty" xml:"finishedAt,omitempty"`
	// The reason why the template entered a failure state. This parameter is returned only when state is error or updateFailed.
	Reason *PublicTemplateStatusReason `json:"reason,omitempty" xml:"reason,omitempty"`
	// The template status. Valid values:
	//
	// - building: The template is being built.
	//
	// - ready: The template is active.
	//
	// - updating: The template is being updated.
	//
	// - updateFailed: The template failed to be updated.
	//
	// - error: The template failed to be built.
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
