// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopCloudRecordShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantContextShrink(v string) *StopCloudRecordShrinkRequest
	GetTenantContextShrink() *string
	SetConferenceId(v string) *StopCloudRecordShrinkRequest
	GetConferenceId() *string
}

type StopCloudRecordShrinkRequest struct {
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 61289fxxx
	ConferenceId *string `json:"conferenceId,omitempty" xml:"conferenceId,omitempty"`
}

func (s StopCloudRecordShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s StopCloudRecordShrinkRequest) GoString() string {
	return s.String()
}

func (s *StopCloudRecordShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *StopCloudRecordShrinkRequest) GetConferenceId() *string {
	return s.ConferenceId
}

func (s *StopCloudRecordShrinkRequest) SetTenantContextShrink(v string) *StopCloudRecordShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *StopCloudRecordShrinkRequest) SetConferenceId(v string) *StopCloudRecordShrinkRequest {
	s.ConferenceId = &v
	return s
}

func (s *StopCloudRecordShrinkRequest) Validate() error {
	return dara.Validate(s)
}
