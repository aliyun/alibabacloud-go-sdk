// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitTenantAliasRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessChannel(v string) *InitTenantAliasRequest
	GetBusinessChannel() *string
}

type InitTenantAliasRequest struct {
	// example:
	//
	// ENTERPRISE
	BusinessChannel *string `json:"BusinessChannel,omitempty" xml:"BusinessChannel,omitempty"`
}

func (s InitTenantAliasRequest) String() string {
	return dara.Prettify(s)
}

func (s InitTenantAliasRequest) GoString() string {
	return s.String()
}

func (s *InitTenantAliasRequest) GetBusinessChannel() *string {
	return s.BusinessChannel
}

func (s *InitTenantAliasRequest) SetBusinessChannel(v string) *InitTenantAliasRequest {
	s.BusinessChannel = &v
	return s
}

func (s *InitTenantAliasRequest) Validate() error {
	return dara.Validate(s)
}
