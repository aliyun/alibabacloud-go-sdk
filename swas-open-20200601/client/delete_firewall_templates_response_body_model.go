// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFirewallTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteFirewallTemplatesResponseBody
	GetRequestId() *string
}

type DeleteFirewallTemplatesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 20758A-585D-4A41-A9B2-28DA8F4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteFirewallTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteFirewallTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFirewallTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteFirewallTemplatesResponseBody) SetRequestId(v string) *DeleteFirewallTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFirewallTemplatesResponseBody) Validate() error {
	return dara.Validate(s)
}
