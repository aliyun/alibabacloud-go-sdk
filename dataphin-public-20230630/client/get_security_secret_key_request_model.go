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
	SetOpUserId(v string) *GetSecuritySecretKeyRequest
	GetOpUserId() *string
}

type GetSecuritySecretKeyRequest struct {
	// The key name.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
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

func (s *GetSecuritySecretKeyRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *GetSecuritySecretKeyRequest) SetName(v string) *GetSecuritySecretKeyRequest {
	s.Name = &v
	return s
}

func (s *GetSecuritySecretKeyRequest) SetOpTenantId(v int64) *GetSecuritySecretKeyRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetSecuritySecretKeyRequest) SetOpUserId(v string) *GetSecuritySecretKeyRequest {
	s.OpUserId = &v
	return s
}

func (s *GetSecuritySecretKeyRequest) Validate() error {
	return dara.Validate(s)
}
