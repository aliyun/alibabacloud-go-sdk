// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSupplementDagrunRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnv(v string) *GetSupplementDagrunRequest
	GetEnv() *string
	SetOpTenantId(v int64) *GetSupplementDagrunRequest
	GetOpTenantId() *int64
	SetSupplementId(v string) *GetSupplementDagrunRequest
	GetSupplementId() *string
}

type GetSupplementDagrunRequest struct {
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
	// This parameter is required.
	//
	// example:
	//
	// f_8241792_20201202_2099680
	SupplementId *string `json:"SupplementId,omitempty" xml:"SupplementId,omitempty"`
}

func (s GetSupplementDagrunRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSupplementDagrunRequest) GoString() string {
	return s.String()
}

func (s *GetSupplementDagrunRequest) GetEnv() *string {
	return s.Env
}

func (s *GetSupplementDagrunRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetSupplementDagrunRequest) GetSupplementId() *string {
	return s.SupplementId
}

func (s *GetSupplementDagrunRequest) SetEnv(v string) *GetSupplementDagrunRequest {
	s.Env = &v
	return s
}

func (s *GetSupplementDagrunRequest) SetOpTenantId(v int64) *GetSupplementDagrunRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetSupplementDagrunRequest) SetSupplementId(v string) *GetSupplementDagrunRequest {
	s.SupplementId = &v
	return s
}

func (s *GetSupplementDagrunRequest) Validate() error {
	return dara.Validate(s)
}
