// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReActivateAuditRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnableAudit(v bool) *ReActivateAuditRequest
	GetEnableAudit() *bool
	SetServiceMeshId(v string) *ReActivateAuditRequest
	GetServiceMeshId() *string
}

type ReActivateAuditRequest struct {
	// Specifies whether to recreate a project that is used to store audit logs. Valid values:
	//
	// 	- true: recreates a project.
	//
	// 	- false: does not recreate a project.
	//
	// example:
	//
	// true
	EnableAudit *bool `json:"EnableAudit,omitempty" xml:"EnableAudit,omitempty"`
	// The ID of the Service Mesh (ASM) instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// c5bf9eb05c4424b89985d6536a809****
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s ReActivateAuditRequest) String() string {
	return dara.Prettify(s)
}

func (s ReActivateAuditRequest) GoString() string {
	return s.String()
}

func (s *ReActivateAuditRequest) GetEnableAudit() *bool {
	return s.EnableAudit
}

func (s *ReActivateAuditRequest) GetServiceMeshId() *string {
	return s.ServiceMeshId
}

func (s *ReActivateAuditRequest) SetEnableAudit(v bool) *ReActivateAuditRequest {
	s.EnableAudit = &v
	return s
}

func (s *ReActivateAuditRequest) SetServiceMeshId(v string) *ReActivateAuditRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *ReActivateAuditRequest) Validate() error {
	return dara.Validate(s)
}
