// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyTrFirewallV2ConfigurationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyTrFirewallV2ConfigurationResponseBody
	GetRequestId() *string
}

type ModifyTrFirewallV2ConfigurationResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// A74C8FDD-2BEF-52D5-8B01-EB6FD94606F9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyTrFirewallV2ConfigurationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyTrFirewallV2ConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyTrFirewallV2ConfigurationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyTrFirewallV2ConfigurationResponseBody) SetRequestId(v string) *ModifyTrFirewallV2ConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyTrFirewallV2ConfigurationResponseBody) Validate() error {
	return dara.Validate(s)
}
