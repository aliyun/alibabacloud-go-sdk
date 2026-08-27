// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckHealthRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantId(v string) *CheckHealthRequest
	GetTenantId() *string
}

type CheckHealthRequest struct {
	// The tenant ID.
	//
	// example:
	//
	// 692318833855074
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s CheckHealthRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckHealthRequest) GoString() string {
	return s.String()
}

func (s *CheckHealthRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CheckHealthRequest) SetTenantId(v string) *CheckHealthRequest {
	s.TenantId = &v
	return s
}

func (s *CheckHealthRequest) Validate() error {
	return dara.Validate(s)
}
