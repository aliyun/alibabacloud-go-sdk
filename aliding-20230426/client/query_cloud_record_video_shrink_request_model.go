// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCloudRecordVideoShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantContextShrink(v string) *QueryCloudRecordVideoShrinkRequest
	GetTenantContextShrink() *string
	SetConferenceId(v string) *QueryCloudRecordVideoShrinkRequest
	GetConferenceId() *string
}

type QueryCloudRecordVideoShrinkRequest struct {
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 61289fxxx
	ConferenceId *string `json:"conferenceId,omitempty" xml:"conferenceId,omitempty"`
}

func (s QueryCloudRecordVideoShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryCloudRecordVideoShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryCloudRecordVideoShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *QueryCloudRecordVideoShrinkRequest) GetConferenceId() *string {
	return s.ConferenceId
}

func (s *QueryCloudRecordVideoShrinkRequest) SetTenantContextShrink(v string) *QueryCloudRecordVideoShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *QueryCloudRecordVideoShrinkRequest) SetConferenceId(v string) *QueryCloudRecordVideoShrinkRequest {
	s.ConferenceId = &v
	return s
}

func (s *QueryCloudRecordVideoShrinkRequest) Validate() error {
	return dara.Validate(s)
}
