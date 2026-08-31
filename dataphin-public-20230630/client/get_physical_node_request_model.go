// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPhysicalNodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnv(v string) *GetPhysicalNodeRequest
	GetEnv() *string
	SetNodeId(v string) *GetPhysicalNodeRequest
	GetNodeId() *string
	SetOpTenantId(v int64) *GetPhysicalNodeRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *GetPhysicalNodeRequest
	GetOpUserId() *string
}

type GetPhysicalNodeRequest struct {
	// The environment identifier. Valid values:
	//
	// - DEV: development environment.
	//
	// - PROD (default): production environment.
	//
	// example:
	//
	// PROD
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// The node ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// n_232132
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
}

func (s GetPhysicalNodeRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPhysicalNodeRequest) GoString() string {
	return s.String()
}

func (s *GetPhysicalNodeRequest) GetEnv() *string {
	return s.Env
}

func (s *GetPhysicalNodeRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *GetPhysicalNodeRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetPhysicalNodeRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *GetPhysicalNodeRequest) SetEnv(v string) *GetPhysicalNodeRequest {
	s.Env = &v
	return s
}

func (s *GetPhysicalNodeRequest) SetNodeId(v string) *GetPhysicalNodeRequest {
	s.NodeId = &v
	return s
}

func (s *GetPhysicalNodeRequest) SetOpTenantId(v int64) *GetPhysicalNodeRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetPhysicalNodeRequest) SetOpUserId(v string) *GetPhysicalNodeRequest {
	s.OpUserId = &v
	return s
}

func (s *GetPhysicalNodeRequest) Validate() error {
	return dara.Validate(s)
}
