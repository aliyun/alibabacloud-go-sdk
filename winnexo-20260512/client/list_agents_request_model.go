// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantId(v string) *ListAgentsRequest
	GetTenantId() *string
}

type ListAgentsRequest struct {
	// The tenant ID. This is a common parameter. If not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListAgentsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAgentsRequest) GoString() string {
	return s.String()
}

func (s *ListAgentsRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListAgentsRequest) SetTenantId(v string) *ListAgentsRequest {
	s.TenantId = &v
	return s
}

func (s *ListAgentsRequest) Validate() error {
	return dara.Validate(s)
}
