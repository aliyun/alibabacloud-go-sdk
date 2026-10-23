// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkflowInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetWorkflowInstanceRequest
	GetInstanceId() *string
	SetTenantId(v string) *GetWorkflowInstanceRequest
	GetTenantId() *string
}

type GetWorkflowInstanceRequest struct {
	// The stable logical instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// wfi_20261016_abc123
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The tenant ID. This is a common parameter. If you use winnexo-cli, pass this parameter explicitly by using --tenant-id.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GetWorkflowInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetWorkflowInstanceRequest) GoString() string {
	return s.String()
}

func (s *GetWorkflowInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetWorkflowInstanceRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetWorkflowInstanceRequest) SetInstanceId(v string) *GetWorkflowInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *GetWorkflowInstanceRequest) SetTenantId(v string) *GetWorkflowInstanceRequest {
	s.TenantId = &v
	return s
}

func (s *GetWorkflowInstanceRequest) Validate() error {
	return dara.Validate(s)
}
