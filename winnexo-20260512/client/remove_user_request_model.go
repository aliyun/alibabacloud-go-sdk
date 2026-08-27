// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantId(v string) *RemoveUserRequest
	GetTenantId() *string
	SetWnUserId(v string) *RemoveUserRequest
	GetWnUserId() *string
}

type RemoveUserRequest struct {
	// The tenant ID. This is a common parameter. If not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The ID of the target user (WINNEXO platform user ID).
	//
	// example:
	//
	// 1
	WnUserId *string `json:"wnUserId,omitempty" xml:"wnUserId,omitempty"`
}

func (s RemoveUserRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveUserRequest) GoString() string {
	return s.String()
}

func (s *RemoveUserRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *RemoveUserRequest) GetWnUserId() *string {
	return s.WnUserId
}

func (s *RemoveUserRequest) SetTenantId(v string) *RemoveUserRequest {
	s.TenantId = &v
	return s
}

func (s *RemoveUserRequest) SetWnUserId(v string) *RemoveUserRequest {
	s.WnUserId = &v
	return s
}

func (s *RemoveUserRequest) Validate() error {
	return dara.Validate(s)
}
