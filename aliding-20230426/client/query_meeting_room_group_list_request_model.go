// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMeetingRoomGroupListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRequest(v map[string]interface{}) *QueryMeetingRoomGroupListRequest
	GetRequest() map[string]interface{}
	SetTenantContext(v *QueryMeetingRoomGroupListRequestTenantContext) *QueryMeetingRoomGroupListRequest
	GetTenantContext() *QueryMeetingRoomGroupListRequestTenantContext
}

type QueryMeetingRoomGroupListRequest struct {
	Request       map[string]interface{}                         `json:"Request,omitempty" xml:"Request,omitempty"`
	TenantContext *QueryMeetingRoomGroupListRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
}

func (s QueryMeetingRoomGroupListRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryMeetingRoomGroupListRequest) GoString() string {
	return s.String()
}

func (s *QueryMeetingRoomGroupListRequest) GetRequest() map[string]interface{} {
	return s.Request
}

func (s *QueryMeetingRoomGroupListRequest) GetTenantContext() *QueryMeetingRoomGroupListRequestTenantContext {
	return s.TenantContext
}

func (s *QueryMeetingRoomGroupListRequest) SetRequest(v map[string]interface{}) *QueryMeetingRoomGroupListRequest {
	s.Request = v
	return s
}

func (s *QueryMeetingRoomGroupListRequest) SetTenantContext(v *QueryMeetingRoomGroupListRequestTenantContext) *QueryMeetingRoomGroupListRequest {
	s.TenantContext = v
	return s
}

func (s *QueryMeetingRoomGroupListRequest) Validate() error {
	return dara.Validate(s)
}

type QueryMeetingRoomGroupListRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s QueryMeetingRoomGroupListRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s QueryMeetingRoomGroupListRequestTenantContext) GoString() string {
	return s.String()
}

func (s *QueryMeetingRoomGroupListRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *QueryMeetingRoomGroupListRequestTenantContext) SetTenantId(v string) *QueryMeetingRoomGroupListRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *QueryMeetingRoomGroupListRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
