// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopCloudRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantContext(v *StopCloudRecordRequestTenantContext) *StopCloudRecordRequest
	GetTenantContext() *StopCloudRecordRequestTenantContext
	SetConferenceId(v string) *StopCloudRecordRequest
	GetConferenceId() *string
}

type StopCloudRecordRequest struct {
	TenantContext *StopCloudRecordRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 61289fxxx
	ConferenceId *string `json:"conferenceId,omitempty" xml:"conferenceId,omitempty"`
}

func (s StopCloudRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s StopCloudRecordRequest) GoString() string {
	return s.String()
}

func (s *StopCloudRecordRequest) GetTenantContext() *StopCloudRecordRequestTenantContext {
	return s.TenantContext
}

func (s *StopCloudRecordRequest) GetConferenceId() *string {
	return s.ConferenceId
}

func (s *StopCloudRecordRequest) SetTenantContext(v *StopCloudRecordRequestTenantContext) *StopCloudRecordRequest {
	s.TenantContext = v
	return s
}

func (s *StopCloudRecordRequest) SetConferenceId(v string) *StopCloudRecordRequest {
	s.ConferenceId = &v
	return s
}

func (s *StopCloudRecordRequest) Validate() error {
	return dara.Validate(s)
}

type StopCloudRecordRequestTenantContext struct {
	// example:
	//
	// 1
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s StopCloudRecordRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s StopCloudRecordRequestTenantContext) GoString() string {
	return s.String()
}

func (s *StopCloudRecordRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *StopCloudRecordRequestTenantContext) SetTenantId(v string) *StopCloudRecordRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *StopCloudRecordRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
