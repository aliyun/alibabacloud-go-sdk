// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloseVideoConferenceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantContext(v *CloseVideoConferenceRequestTenantContext) *CloseVideoConferenceRequest
	GetTenantContext() *CloseVideoConferenceRequestTenantContext
	SetConferenceId(v string) *CloseVideoConferenceRequest
	GetConferenceId() *string
}

type CloseVideoConferenceRequest struct {
	TenantContext *CloseVideoConferenceRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 607452e01401526ee39609e1
	ConferenceId *string `json:"conferenceId,omitempty" xml:"conferenceId,omitempty"`
}

func (s CloseVideoConferenceRequest) String() string {
	return dara.Prettify(s)
}

func (s CloseVideoConferenceRequest) GoString() string {
	return s.String()
}

func (s *CloseVideoConferenceRequest) GetTenantContext() *CloseVideoConferenceRequestTenantContext {
	return s.TenantContext
}

func (s *CloseVideoConferenceRequest) GetConferenceId() *string {
	return s.ConferenceId
}

func (s *CloseVideoConferenceRequest) SetTenantContext(v *CloseVideoConferenceRequestTenantContext) *CloseVideoConferenceRequest {
	s.TenantContext = v
	return s
}

func (s *CloseVideoConferenceRequest) SetConferenceId(v string) *CloseVideoConferenceRequest {
	s.ConferenceId = &v
	return s
}

func (s *CloseVideoConferenceRequest) Validate() error {
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CloseVideoConferenceRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s CloseVideoConferenceRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s CloseVideoConferenceRequestTenantContext) GoString() string {
	return s.String()
}

func (s *CloseVideoConferenceRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *CloseVideoConferenceRequestTenantContext) SetTenantId(v string) *CloseVideoConferenceRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *CloseVideoConferenceRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
