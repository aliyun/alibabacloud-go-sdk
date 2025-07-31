// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPhysicalNodeContentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnv(v string) *GetPhysicalNodeContentRequest
	GetEnv() *string
	SetNodeId(v string) *GetPhysicalNodeContentRequest
	GetNodeId() *string
	SetOpTenantId(v int64) *GetPhysicalNodeContentRequest
	GetOpTenantId() *int64
}

type GetPhysicalNodeContentRequest struct {
	// example:
	//
	// PROD
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// n_232411
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s GetPhysicalNodeContentRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPhysicalNodeContentRequest) GoString() string {
	return s.String()
}

func (s *GetPhysicalNodeContentRequest) GetEnv() *string {
	return s.Env
}

func (s *GetPhysicalNodeContentRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *GetPhysicalNodeContentRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetPhysicalNodeContentRequest) SetEnv(v string) *GetPhysicalNodeContentRequest {
	s.Env = &v
	return s
}

func (s *GetPhysicalNodeContentRequest) SetNodeId(v string) *GetPhysicalNodeContentRequest {
	s.NodeId = &v
	return s
}

func (s *GetPhysicalNodeContentRequest) SetOpTenantId(v int64) *GetPhysicalNodeContentRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetPhysicalNodeContentRequest) Validate() error {
	return dara.Validate(s)
}
