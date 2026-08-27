// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceExpireTimeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantId(v string) *GetInstanceExpireTimeRequest
	GetTenantId() *string
}

type GetInstanceExpireTimeRequest struct {
	// The effective tenant ID.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GetInstanceExpireTimeRequest) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceExpireTimeRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceExpireTimeRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetInstanceExpireTimeRequest) SetTenantId(v string) *GetInstanceExpireTimeRequest {
	s.TenantId = &v
	return s
}

func (s *GetInstanceExpireTimeRequest) Validate() error {
	return dara.Validate(s)
}
