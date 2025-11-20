// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartCloudRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMode(v string) *StartCloudRecordRequest
	GetMode() *string
	SetSmallWindowPosition(v string) *StartCloudRecordRequest
	GetSmallWindowPosition() *string
	SetTenantContext(v *StartCloudRecordRequestTenantContext) *StartCloudRecordRequest
	GetTenantContext() *StartCloudRecordRequestTenantContext
	SetConferenceId(v string) *StartCloudRecordRequest
	GetConferenceId() *string
}

type StartCloudRecordRequest struct {
	// example:
	//
	// speech
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// example:
	//
	// relative_right
	SmallWindowPosition *string                               `json:"SmallWindowPosition,omitempty" xml:"SmallWindowPosition,omitempty"`
	TenantContext       *StartCloudRecordRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 61289fxxx
	ConferenceId *string `json:"conferenceId,omitempty" xml:"conferenceId,omitempty"`
}

func (s StartCloudRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s StartCloudRecordRequest) GoString() string {
	return s.String()
}

func (s *StartCloudRecordRequest) GetMode() *string {
	return s.Mode
}

func (s *StartCloudRecordRequest) GetSmallWindowPosition() *string {
	return s.SmallWindowPosition
}

func (s *StartCloudRecordRequest) GetTenantContext() *StartCloudRecordRequestTenantContext {
	return s.TenantContext
}

func (s *StartCloudRecordRequest) GetConferenceId() *string {
	return s.ConferenceId
}

func (s *StartCloudRecordRequest) SetMode(v string) *StartCloudRecordRequest {
	s.Mode = &v
	return s
}

func (s *StartCloudRecordRequest) SetSmallWindowPosition(v string) *StartCloudRecordRequest {
	s.SmallWindowPosition = &v
	return s
}

func (s *StartCloudRecordRequest) SetTenantContext(v *StartCloudRecordRequestTenantContext) *StartCloudRecordRequest {
	s.TenantContext = v
	return s
}

func (s *StartCloudRecordRequest) SetConferenceId(v string) *StartCloudRecordRequest {
	s.ConferenceId = &v
	return s
}

func (s *StartCloudRecordRequest) Validate() error {
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type StartCloudRecordRequestTenantContext struct {
	// example:
	//
	// 1
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s StartCloudRecordRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s StartCloudRecordRequestTenantContext) GoString() string {
	return s.String()
}

func (s *StartCloudRecordRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *StartCloudRecordRequestTenantContext) SetTenantId(v string) *StartCloudRecordRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *StartCloudRecordRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
