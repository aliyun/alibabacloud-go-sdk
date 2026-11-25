// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVpcFirewallConfigureResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteVpcFirewallConfigureResponseBody
	GetRequestId() *string
}

type DeleteVpcFirewallConfigureResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 850A84D6-0DE4-4797-A1E8-00090125k6j3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVpcFirewallConfigureResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteVpcFirewallConfigureResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVpcFirewallConfigureResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteVpcFirewallConfigureResponseBody) SetRequestId(v string) *DeleteVpcFirewallConfigureResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteVpcFirewallConfigureResponseBody) Validate() error {
	return dara.Validate(s)
}
