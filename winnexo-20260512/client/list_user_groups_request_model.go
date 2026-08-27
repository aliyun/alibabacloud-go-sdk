// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantId(v string) *ListUserGroupsRequest
	GetTenantId() *string
}

type ListUserGroupsRequest struct {
	// The tenant ID. This is a common parameter. In winnexo-cli, pass this parameter explicitly by using `--tenant-id`.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListUserGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUserGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListUserGroupsRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListUserGroupsRequest) SetTenantId(v string) *ListUserGroupsRequest {
	s.TenantId = &v
	return s
}

func (s *ListUserGroupsRequest) Validate() error {
	return dara.Validate(s)
}
