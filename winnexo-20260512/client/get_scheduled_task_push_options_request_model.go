// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScheduledTaskPushOptionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollaborationGroupId(v string) *GetScheduledTaskPushOptionsRequest
	GetCollaborationGroupId() *string
	SetDigitalEmployeeName(v string) *GetScheduledTaskPushOptionsRequest
	GetDigitalEmployeeName() *string
	SetTenantId(v string) *GetScheduledTaskPushOptionsRequest
	GetTenantId() *string
}

type GetScheduledTaskPushOptionsRequest struct {
	// The ID of the collaboration group (such as cg_101). If specified, a group workspace task is created (the caller must be a valid group member). If left empty, a personal task is created.
	//
	// example:
	//
	// cg_401
	CollaborationGroupId *string `json:"collaborationGroupId,omitempty" xml:"collaborationGroupId,omitempty"`
	// The name of the currently active digital employee. This value is empty if not configured.
	//
	// example:
	//
	// exampleDigitalEmployeeName
	DigitalEmployeeName *string `json:"digitalEmployeeName,omitempty" xml:"digitalEmployeeName,omitempty"`
	// The tenant ID. This is a common parameter. In winnexo-cli, pass it explicitly with --tenant-id.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GetScheduledTaskPushOptionsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetScheduledTaskPushOptionsRequest) GoString() string {
	return s.String()
}

func (s *GetScheduledTaskPushOptionsRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *GetScheduledTaskPushOptionsRequest) GetDigitalEmployeeName() *string {
	return s.DigitalEmployeeName
}

func (s *GetScheduledTaskPushOptionsRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetScheduledTaskPushOptionsRequest) SetCollaborationGroupId(v string) *GetScheduledTaskPushOptionsRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *GetScheduledTaskPushOptionsRequest) SetDigitalEmployeeName(v string) *GetScheduledTaskPushOptionsRequest {
	s.DigitalEmployeeName = &v
	return s
}

func (s *GetScheduledTaskPushOptionsRequest) SetTenantId(v string) *GetScheduledTaskPushOptionsRequest {
	s.TenantId = &v
	return s
}

func (s *GetScheduledTaskPushOptionsRequest) Validate() error {
	return dara.Validate(s)
}
