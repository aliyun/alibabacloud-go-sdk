// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPhysicalNodeOperationLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnv(v string) *GetPhysicalNodeOperationLogRequest
	GetEnv() *string
	SetNodeId(v string) *GetPhysicalNodeOperationLogRequest
	GetNodeId() *string
	SetOpTenantId(v int64) *GetPhysicalNodeOperationLogRequest
	GetOpTenantId() *int64
}

type GetPhysicalNodeOperationLogRequest struct {
	// example:
	//
	// PROD
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// n_231131
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s GetPhysicalNodeOperationLogRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPhysicalNodeOperationLogRequest) GoString() string {
	return s.String()
}

func (s *GetPhysicalNodeOperationLogRequest) GetEnv() *string {
	return s.Env
}

func (s *GetPhysicalNodeOperationLogRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *GetPhysicalNodeOperationLogRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetPhysicalNodeOperationLogRequest) SetEnv(v string) *GetPhysicalNodeOperationLogRequest {
	s.Env = &v
	return s
}

func (s *GetPhysicalNodeOperationLogRequest) SetNodeId(v string) *GetPhysicalNodeOperationLogRequest {
	s.NodeId = &v
	return s
}

func (s *GetPhysicalNodeOperationLogRequest) SetOpTenantId(v int64) *GetPhysicalNodeOperationLogRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetPhysicalNodeOperationLogRequest) Validate() error {
	return dara.Validate(s)
}
