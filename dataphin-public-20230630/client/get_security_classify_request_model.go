// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSecurityClassifyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *GetSecurityClassifyRequest
	GetId() *int64
	SetOpTenantId(v int64) *GetSecurityClassifyRequest
	GetOpTenantId() *int64
}

type GetSecurityClassifyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s GetSecurityClassifyRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSecurityClassifyRequest) GoString() string {
	return s.String()
}

func (s *GetSecurityClassifyRequest) GetId() *int64 {
	return s.Id
}

func (s *GetSecurityClassifyRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetSecurityClassifyRequest) SetId(v int64) *GetSecurityClassifyRequest {
	s.Id = &v
	return s
}

func (s *GetSecurityClassifyRequest) SetOpTenantId(v int64) *GetSecurityClassifyRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetSecurityClassifyRequest) Validate() error {
	return dara.Validate(s)
}
