// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloseVideoConferenceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantContextShrink(v string) *CloseVideoConferenceShrinkRequest
	GetTenantContextShrink() *string
	SetConferenceId(v string) *CloseVideoConferenceShrinkRequest
	GetConferenceId() *string
}

type CloseVideoConferenceShrinkRequest struct {
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 607452e01401526ee39609e1
	ConferenceId *string `json:"conferenceId,omitempty" xml:"conferenceId,omitempty"`
}

func (s CloseVideoConferenceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CloseVideoConferenceShrinkRequest) GoString() string {
	return s.String()
}

func (s *CloseVideoConferenceShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *CloseVideoConferenceShrinkRequest) GetConferenceId() *string {
	return s.ConferenceId
}

func (s *CloseVideoConferenceShrinkRequest) SetTenantContextShrink(v string) *CloseVideoConferenceShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *CloseVideoConferenceShrinkRequest) SetConferenceId(v string) *CloseVideoConferenceShrinkRequest {
	s.ConferenceId = &v
	return s
}

func (s *CloseVideoConferenceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
