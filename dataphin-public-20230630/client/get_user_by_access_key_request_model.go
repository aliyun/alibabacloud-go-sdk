// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserByAccessKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *GetUserByAccessKeyRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *GetUserByAccessKeyRequest
	GetOpUserId() *string
}

type GetUserByAccessKeyRequest struct {
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The ID of the operator user.
	//
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
}

func (s GetUserByAccessKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUserByAccessKeyRequest) GoString() string {
	return s.String()
}

func (s *GetUserByAccessKeyRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetUserByAccessKeyRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *GetUserByAccessKeyRequest) SetOpTenantId(v int64) *GetUserByAccessKeyRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetUserByAccessKeyRequest) SetOpUserId(v string) *GetUserByAccessKeyRequest {
	s.OpUserId = &v
	return s
}

func (s *GetUserByAccessKeyRequest) Validate() error {
	return dara.Validate(s)
}
