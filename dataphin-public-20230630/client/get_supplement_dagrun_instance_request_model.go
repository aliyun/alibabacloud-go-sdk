// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSupplementDagrunInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDagrunId(v string) *GetSupplementDagrunInstanceRequest
	GetDagrunId() *string
	SetEnv(v string) *GetSupplementDagrunInstanceRequest
	GetEnv() *string
	SetOpTenantId(v int64) *GetSupplementDagrunInstanceRequest
	GetOpTenantId() *int64
}

type GetSupplementDagrunInstanceRequest struct {
	// Dagrun ID
	//
	// This parameter is required.
	//
	// example:
	//
	// dr_2242792_14542
	DagrunId *string `json:"DagrunId,omitempty" xml:"DagrunId,omitempty"`
	// example:
	//
	// PROD
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s GetSupplementDagrunInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSupplementDagrunInstanceRequest) GoString() string {
	return s.String()
}

func (s *GetSupplementDagrunInstanceRequest) GetDagrunId() *string {
	return s.DagrunId
}

func (s *GetSupplementDagrunInstanceRequest) GetEnv() *string {
	return s.Env
}

func (s *GetSupplementDagrunInstanceRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetSupplementDagrunInstanceRequest) SetDagrunId(v string) *GetSupplementDagrunInstanceRequest {
	s.DagrunId = &v
	return s
}

func (s *GetSupplementDagrunInstanceRequest) SetEnv(v string) *GetSupplementDagrunInstanceRequest {
	s.Env = &v
	return s
}

func (s *GetSupplementDagrunInstanceRequest) SetOpTenantId(v int64) *GetSupplementDagrunInstanceRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetSupplementDagrunInstanceRequest) Validate() error {
	return dara.Validate(s)
}
