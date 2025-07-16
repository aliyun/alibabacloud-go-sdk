// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartCloudRecordShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMode(v string) *StartCloudRecordShrinkRequest
	GetMode() *string
	SetSmallWindowPosition(v string) *StartCloudRecordShrinkRequest
	GetSmallWindowPosition() *string
	SetTenantContextShrink(v string) *StartCloudRecordShrinkRequest
	GetTenantContextShrink() *string
	SetConferenceId(v string) *StartCloudRecordShrinkRequest
	GetConferenceId() *string
}

type StartCloudRecordShrinkRequest struct {
	// example:
	//
	// speech
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// example:
	//
	// relative_right
	SmallWindowPosition *string `json:"SmallWindowPosition,omitempty" xml:"SmallWindowPosition,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 61289fxxx
	ConferenceId *string `json:"conferenceId,omitempty" xml:"conferenceId,omitempty"`
}

func (s StartCloudRecordShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s StartCloudRecordShrinkRequest) GoString() string {
	return s.String()
}

func (s *StartCloudRecordShrinkRequest) GetMode() *string {
	return s.Mode
}

func (s *StartCloudRecordShrinkRequest) GetSmallWindowPosition() *string {
	return s.SmallWindowPosition
}

func (s *StartCloudRecordShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *StartCloudRecordShrinkRequest) GetConferenceId() *string {
	return s.ConferenceId
}

func (s *StartCloudRecordShrinkRequest) SetMode(v string) *StartCloudRecordShrinkRequest {
	s.Mode = &v
	return s
}

func (s *StartCloudRecordShrinkRequest) SetSmallWindowPosition(v string) *StartCloudRecordShrinkRequest {
	s.SmallWindowPosition = &v
	return s
}

func (s *StartCloudRecordShrinkRequest) SetTenantContextShrink(v string) *StartCloudRecordShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *StartCloudRecordShrinkRequest) SetConferenceId(v string) *StartCloudRecordShrinkRequest {
	s.ConferenceId = &v
	return s
}

func (s *StartCloudRecordShrinkRequest) Validate() error {
	return dara.Validate(s)
}
