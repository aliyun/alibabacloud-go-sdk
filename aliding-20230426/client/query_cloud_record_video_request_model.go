// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCloudRecordVideoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantContext(v *QueryCloudRecordVideoRequestTenantContext) *QueryCloudRecordVideoRequest
	GetTenantContext() *QueryCloudRecordVideoRequestTenantContext
	SetConferenceId(v string) *QueryCloudRecordVideoRequest
	GetConferenceId() *string
}

type QueryCloudRecordVideoRequest struct {
	TenantContext *QueryCloudRecordVideoRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 61289fxxx
	ConferenceId *string `json:"conferenceId,omitempty" xml:"conferenceId,omitempty"`
}

func (s QueryCloudRecordVideoRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryCloudRecordVideoRequest) GoString() string {
	return s.String()
}

func (s *QueryCloudRecordVideoRequest) GetTenantContext() *QueryCloudRecordVideoRequestTenantContext {
	return s.TenantContext
}

func (s *QueryCloudRecordVideoRequest) GetConferenceId() *string {
	return s.ConferenceId
}

func (s *QueryCloudRecordVideoRequest) SetTenantContext(v *QueryCloudRecordVideoRequestTenantContext) *QueryCloudRecordVideoRequest {
	s.TenantContext = v
	return s
}

func (s *QueryCloudRecordVideoRequest) SetConferenceId(v string) *QueryCloudRecordVideoRequest {
	s.ConferenceId = &v
	return s
}

func (s *QueryCloudRecordVideoRequest) Validate() error {
	return dara.Validate(s)
}

type QueryCloudRecordVideoRequestTenantContext struct {
	// example:
	//
	// 1
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s QueryCloudRecordVideoRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s QueryCloudRecordVideoRequestTenantContext) GoString() string {
	return s.String()
}

func (s *QueryCloudRecordVideoRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *QueryCloudRecordVideoRequestTenantContext) SetTenantId(v string) *QueryCloudRecordVideoRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *QueryCloudRecordVideoRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
