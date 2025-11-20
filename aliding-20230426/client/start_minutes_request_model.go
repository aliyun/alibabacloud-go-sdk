// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartMinutesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantContext(v *StartMinutesRequestTenantContext) *StartMinutesRequest
	GetTenantContext() *StartMinutesRequestTenantContext
	SetConferenceId(v string) *StartMinutesRequest
	GetConferenceId() *string
	SetOwnerUserId(v string) *StartMinutesRequest
	GetOwnerUserId() *string
	SetRecordAudio(v bool) *StartMinutesRequest
	GetRecordAudio() *bool
}

type StartMinutesRequest struct {
	TenantContext *StartMinutesRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 61289fxxx
	ConferenceId *string `json:"conferenceId,omitempty" xml:"conferenceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 012345
	OwnerUserId *string `json:"ownerUserId,omitempty" xml:"ownerUserId,omitempty"`
	// example:
	//
	// false
	RecordAudio *bool `json:"recordAudio,omitempty" xml:"recordAudio,omitempty"`
}

func (s StartMinutesRequest) String() string {
	return dara.Prettify(s)
}

func (s StartMinutesRequest) GoString() string {
	return s.String()
}

func (s *StartMinutesRequest) GetTenantContext() *StartMinutesRequestTenantContext {
	return s.TenantContext
}

func (s *StartMinutesRequest) GetConferenceId() *string {
	return s.ConferenceId
}

func (s *StartMinutesRequest) GetOwnerUserId() *string {
	return s.OwnerUserId
}

func (s *StartMinutesRequest) GetRecordAudio() *bool {
	return s.RecordAudio
}

func (s *StartMinutesRequest) SetTenantContext(v *StartMinutesRequestTenantContext) *StartMinutesRequest {
	s.TenantContext = v
	return s
}

func (s *StartMinutesRequest) SetConferenceId(v string) *StartMinutesRequest {
	s.ConferenceId = &v
	return s
}

func (s *StartMinutesRequest) SetOwnerUserId(v string) *StartMinutesRequest {
	s.OwnerUserId = &v
	return s
}

func (s *StartMinutesRequest) SetRecordAudio(v bool) *StartMinutesRequest {
	s.RecordAudio = &v
	return s
}

func (s *StartMinutesRequest) Validate() error {
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type StartMinutesRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s StartMinutesRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s StartMinutesRequestTenantContext) GoString() string {
	return s.String()
}

func (s *StartMinutesRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *StartMinutesRequestTenantContext) SetTenantId(v string) *StartMinutesRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *StartMinutesRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
