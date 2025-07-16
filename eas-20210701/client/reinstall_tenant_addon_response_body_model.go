// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReinstallTenantAddonResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *ReinstallTenantAddonResponseBody
	GetMessage() *string
	SetRequestId(v string) *ReinstallTenantAddonResponseBody
	GetRequestId() *string
}

type ReinstallTenantAddonResponseBody struct {
	// The returned message.
	//
	// example:
	//
	// Addon prometheus_discovery is successfully reinstalled
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReinstallTenantAddonResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReinstallTenantAddonResponseBody) GoString() string {
	return s.String()
}

func (s *ReinstallTenantAddonResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ReinstallTenantAddonResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReinstallTenantAddonResponseBody) SetMessage(v string) *ReinstallTenantAddonResponseBody {
	s.Message = &v
	return s
}

func (s *ReinstallTenantAddonResponseBody) SetRequestId(v string) *ReinstallTenantAddonResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReinstallTenantAddonResponseBody) Validate() error {
	return dara.Validate(s)
}
