// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGraphsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantId(v string) *ListGraphsRequest
	GetTenantId() *string
}

type ListGraphsRequest struct {
	// 租户ID，公共参数；winnexo-cli 通过 --tenant-id 显式传入
	//
	// This parameter is required.
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListGraphsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListGraphsRequest) GoString() string {
	return s.String()
}

func (s *ListGraphsRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListGraphsRequest) SetTenantId(v string) *ListGraphsRequest {
	s.TenantId = &v
	return s
}

func (s *ListGraphsRequest) Validate() error {
	return dara.Validate(s)
}
