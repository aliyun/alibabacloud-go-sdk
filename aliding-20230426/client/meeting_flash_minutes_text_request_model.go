// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMeetingFlashMinutesTextRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantContext(v *MeetingFlashMinutesTextRequestTenantContext) *MeetingFlashMinutesTextRequest
	GetTenantContext() *MeetingFlashMinutesTextRequestTenantContext
	SetConferenceId(v string) *MeetingFlashMinutesTextRequest
	GetConferenceId() *string
	SetMaxResults(v int32) *MeetingFlashMinutesTextRequest
	GetMaxResults() *int32
	SetNextToken(v string) *MeetingFlashMinutesTextRequest
	GetNextToken() *string
}

type MeetingFlashMinutesTextRequest struct {
	TenantContext *MeetingFlashMinutesTextRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
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

func (s MeetingFlashMinutesTextRequest) String() string {
	return dara.Prettify(s)
}

func (s MeetingFlashMinutesTextRequest) GoString() string {
	return s.String()
}

func (s *MeetingFlashMinutesTextRequest) GetTenantContext() *MeetingFlashMinutesTextRequestTenantContext {
	return s.TenantContext
}

func (s *MeetingFlashMinutesTextRequest) GetConferenceId() *string {
	return s.ConferenceId
}

func (s *MeetingFlashMinutesTextRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *MeetingFlashMinutesTextRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *MeetingFlashMinutesTextRequest) SetTenantContext(v *MeetingFlashMinutesTextRequestTenantContext) *MeetingFlashMinutesTextRequest {
	s.TenantContext = v
	return s
}

func (s *MeetingFlashMinutesTextRequest) SetConferenceId(v string) *MeetingFlashMinutesTextRequest {
	s.ConferenceId = &v
	return s
}

func (s *MeetingFlashMinutesTextRequest) SetMaxResults(v int32) *MeetingFlashMinutesTextRequest {
	s.MaxResults = &v
	return s
}

func (s *MeetingFlashMinutesTextRequest) SetNextToken(v string) *MeetingFlashMinutesTextRequest {
	s.NextToken = &v
	return s
}

func (s *MeetingFlashMinutesTextRequest) Validate() error {
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type MeetingFlashMinutesTextRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s MeetingFlashMinutesTextRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s MeetingFlashMinutesTextRequestTenantContext) GoString() string {
	return s.String()
}

func (s *MeetingFlashMinutesTextRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *MeetingFlashMinutesTextRequestTenantContext) SetTenantId(v string) *MeetingFlashMinutesTextRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *MeetingFlashMinutesTextRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
