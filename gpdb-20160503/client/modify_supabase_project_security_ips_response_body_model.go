// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySupabaseProjectSecurityIpsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifySupabaseProjectSecurityIpsResponseBody
	GetRequestId() *string
}

type ModifySupabaseProjectSecurityIpsResponseBody struct {
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifySupabaseProjectSecurityIpsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifySupabaseProjectSecurityIpsResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySupabaseProjectSecurityIpsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifySupabaseProjectSecurityIpsResponseBody) SetRequestId(v string) *ModifySupabaseProjectSecurityIpsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySupabaseProjectSecurityIpsResponseBody) Validate() error {
	return dara.Validate(s)
}
