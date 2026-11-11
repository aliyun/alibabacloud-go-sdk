// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisassociateTransitRouterPolicyTableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisassociateTransitRouterPolicyTableResponseBody
	GetRequestId() *string
}

type DisassociateTransitRouterPolicyTableResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// F7DDDC17-FA06-4AC2-8F35-59D2470FCFC1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisassociateTransitRouterPolicyTableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisassociateTransitRouterPolicyTableResponseBody) GoString() string {
	return s.String()
}

func (s *DisassociateTransitRouterPolicyTableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisassociateTransitRouterPolicyTableResponseBody) SetRequestId(v string) *DisassociateTransitRouterPolicyTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisassociateTransitRouterPolicyTableResponseBody) Validate() error {
	return dara.Validate(s)
}
