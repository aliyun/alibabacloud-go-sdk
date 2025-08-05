// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVpcFirewallDefaultIPSConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyVpcFirewallDefaultIPSConfigResponseBody
	GetRequestId() *string
}

type ModifyVpcFirewallDefaultIPSConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 3B51B5BF-3C26-5009-ADAB-190E58DE4D6E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyVpcFirewallDefaultIPSConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyVpcFirewallDefaultIPSConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyVpcFirewallDefaultIPSConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyVpcFirewallDefaultIPSConfigResponseBody) SetRequestId(v string) *ModifyVpcFirewallDefaultIPSConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyVpcFirewallDefaultIPSConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
