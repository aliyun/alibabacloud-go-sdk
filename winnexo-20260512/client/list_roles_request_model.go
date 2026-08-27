// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRolesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantId(v string) *ListRolesRequest
	GetTenantId() *string
}

type ListRolesRequest struct {
	// The tenant ID. This is a common parameter. You can pass it explicitly by using --tenant-id in winnexo-cli.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListRolesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRolesRequest) GoString() string {
	return s.String()
}

func (s *ListRolesRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListRolesRequest) SetTenantId(v string) *ListRolesRequest {
	s.TenantId = &v
	return s
}

func (s *ListRolesRequest) Validate() error {
	return dara.Validate(s)
}
