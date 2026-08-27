// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAvailableConfigsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantId(v string) *ListAvailableConfigsRequest
	GetTenantId() *string
}

type ListAvailableConfigsRequest struct {
	// The tenant ID. This is a common parameter. Pass it explicitly through --tenant-id in winnexo-cli.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListAvailableConfigsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAvailableConfigsRequest) GoString() string {
	return s.String()
}

func (s *ListAvailableConfigsRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListAvailableConfigsRequest) SetTenantId(v string) *ListAvailableConfigsRequest {
	s.TenantId = &v
	return s
}

func (s *ListAvailableConfigsRequest) Validate() error {
	return dara.Validate(s)
}
