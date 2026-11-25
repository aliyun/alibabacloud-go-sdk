// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVpcFirewallCenConfigureResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteVpcFirewallCenConfigureResponseBody
	GetRequestId() *string
}

type DeleteVpcFirewallCenConfigureResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 850A84D6-0DE4-4797-A1E8-00090125k6j3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVpcFirewallCenConfigureResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteVpcFirewallCenConfigureResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVpcFirewallCenConfigureResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteVpcFirewallCenConfigureResponseBody) SetRequestId(v string) *DeleteVpcFirewallCenConfigureResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteVpcFirewallCenConfigureResponseBody) Validate() error {
	return dara.Validate(s)
}
