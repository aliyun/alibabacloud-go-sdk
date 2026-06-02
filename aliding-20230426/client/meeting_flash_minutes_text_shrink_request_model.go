// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMeetingFlashMinutesTextShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantContextShrink(v string) *MeetingFlashMinutesTextShrinkRequest
	GetTenantContextShrink() *string
	SetConferenceId(v string) *MeetingFlashMinutesTextShrinkRequest
	GetConferenceId() *string
	SetMaxResults(v int32) *MeetingFlashMinutesTextShrinkRequest
	GetMaxResults() *int32
	SetNextToken(v string) *MeetingFlashMinutesTextShrinkRequest
	GetNextToken() *string
}

type MeetingFlashMinutesTextShrinkRequest struct {
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 61289fxxx
	ConferenceId *string `json:"conferenceId,omitempty" xml:"conferenceId,omitempty"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// 0
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s MeetingFlashMinutesTextShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s MeetingFlashMinutesTextShrinkRequest) GoString() string {
	return s.String()
}

func (s *MeetingFlashMinutesTextShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *MeetingFlashMinutesTextShrinkRequest) GetConferenceId() *string {
	return s.ConferenceId
}

func (s *MeetingFlashMinutesTextShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *MeetingFlashMinutesTextShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *MeetingFlashMinutesTextShrinkRequest) SetTenantContextShrink(v string) *MeetingFlashMinutesTextShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *MeetingFlashMinutesTextShrinkRequest) SetConferenceId(v string) *MeetingFlashMinutesTextShrinkRequest {
	s.ConferenceId = &v
	return s
}

func (s *MeetingFlashMinutesTextShrinkRequest) SetMaxResults(v int32) *MeetingFlashMinutesTextShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *MeetingFlashMinutesTextShrinkRequest) SetNextToken(v string) *MeetingFlashMinutesTextShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *MeetingFlashMinutesTextShrinkRequest) Validate() error {
	return dara.Validate(s)
}
