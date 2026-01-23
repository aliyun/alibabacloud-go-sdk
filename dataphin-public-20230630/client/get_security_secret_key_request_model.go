// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSecuritySecretKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *GetSecuritySecretKeyRequest
	GetName() *string
	SetOpTenantId(v int64) *GetSecuritySecretKeyRequest
	GetOpTenantId() *int64
}

type GetSecuritySecretKeyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s GetSecuritySecretKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSecuritySecretKeyRequest) GoString() string {
	return s.String()
}

func (s *GetSecuritySecretKeyRequest) GetName() *string {
	return s.Name
}

func (s *GetSecuritySecretKeyRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetSecuritySecretKeyRequest) SetName(v string) *GetSecuritySecretKeyRequest {
	s.Name = &v
	return s
}

func (s *GetSecuritySecretKeyRequest) SetOpTenantId(v int64) *GetSecuritySecretKeyRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetSecuritySecretKeyRequest) Validate() error {
	return dara.Validate(s)
}
