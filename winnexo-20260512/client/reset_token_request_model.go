// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantId(v string) *ResetTokenRequest
	GetTenantId() *string
	SetWnUserId(v string) *ResetTokenRequest
	GetWnUserId() *string
}

type ResetTokenRequest struct {
	// The tenant ID. This is a common parameter. If this parameter is not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The ID of the target user on the WINNEXO platform. If this parameter is left empty, the operation is performed on the caller. Administrators can specify the ID of another user to perform the operation on behalf of that user.
	//
	// example:
	//
	// 1
	WnUserId *string `json:"wnUserId,omitempty" xml:"wnUserId,omitempty"`
}

func (s ResetTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s ResetTokenRequest) GoString() string {
	return s.String()
}

func (s *ResetTokenRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ResetTokenRequest) GetWnUserId() *string {
	return s.WnUserId
}

func (s *ResetTokenRequest) SetTenantId(v string) *ResetTokenRequest {
	s.TenantId = &v
	return s
}

func (s *ResetTokenRequest) SetWnUserId(v string) *ResetTokenRequest {
	s.WnUserId = &v
	return s
}

func (s *ResetTokenRequest) Validate() error {
	return dara.Validate(s)
}
