// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetConferenceHostsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOperationType(v string) *SetConferenceHostsRequest
	GetOperationType() *string
	SetTenantContext(v *SetConferenceHostsRequestTenantContext) *SetConferenceHostsRequest
	GetTenantContext() *SetConferenceHostsRequestTenantContext
	SetUserIds(v []*string) *SetConferenceHostsRequest
	GetUserIds() []*string
	SetConferenceId(v string) *SetConferenceHostsRequest
	GetConferenceId() *string
}

type SetConferenceHostsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// add
	OperationType *string                                 `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	TenantContext *SetConferenceHostsRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// [ "012345"]
	UserIds []*string `json:"UserIds,omitempty" xml:"UserIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 607452e01401526ee39609e1
	ConferenceId *string `json:"conferenceId,omitempty" xml:"conferenceId,omitempty"`
}

func (s SetConferenceHostsRequest) String() string {
	return dara.Prettify(s)
}

func (s SetConferenceHostsRequest) GoString() string {
	return s.String()
}

func (s *SetConferenceHostsRequest) GetOperationType() *string {
	return s.OperationType
}

func (s *SetConferenceHostsRequest) GetTenantContext() *SetConferenceHostsRequestTenantContext {
	return s.TenantContext
}

func (s *SetConferenceHostsRequest) GetUserIds() []*string {
	return s.UserIds
}

func (s *SetConferenceHostsRequest) GetConferenceId() *string {
	return s.ConferenceId
}

func (s *SetConferenceHostsRequest) SetOperationType(v string) *SetConferenceHostsRequest {
	s.OperationType = &v
	return s
}

func (s *SetConferenceHostsRequest) SetTenantContext(v *SetConferenceHostsRequestTenantContext) *SetConferenceHostsRequest {
	s.TenantContext = v
	return s
}

func (s *SetConferenceHostsRequest) SetUserIds(v []*string) *SetConferenceHostsRequest {
	s.UserIds = v
	return s
}

func (s *SetConferenceHostsRequest) SetConferenceId(v string) *SetConferenceHostsRequest {
	s.ConferenceId = &v
	return s
}

func (s *SetConferenceHostsRequest) Validate() error {
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SetConferenceHostsRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s SetConferenceHostsRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s SetConferenceHostsRequestTenantContext) GoString() string {
	return s.String()
}

func (s *SetConferenceHostsRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *SetConferenceHostsRequestTenantContext) SetTenantId(v string) *SetConferenceHostsRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *SetConferenceHostsRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
