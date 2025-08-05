// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTrFirewallV2ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteTrFirewallV2ResponseBody
	GetRequestId() *string
}

type DeleteTrFirewallV2ResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// A774E66E-B170-59FC-9AAF-3068B15E991F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteTrFirewallV2ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteTrFirewallV2ResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTrFirewallV2ResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteTrFirewallV2ResponseBody) SetRequestId(v string) *DeleteTrFirewallV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTrFirewallV2ResponseBody) Validate() error {
	return dara.Validate(s)
}
