// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTransitRouterPolicyTableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPolicyTableId(v string) *CreateTransitRouterPolicyTableResponseBody
	GetPolicyTableId() *string
	SetRequestId(v string) *CreateTransitRouterPolicyTableResponseBody
	GetRequestId() *string
}

type CreateTransitRouterPolicyTableResponseBody struct {
	// The ID of the policy route table.
	//
	// example:
	//
	// ptbl-44m0p68spvlrqq****
	PolicyTableId *string `json:"PolicyTableId,omitempty" xml:"PolicyTableId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2BFA6822-240E-4E27-B4C8-AA400EF7474D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateTransitRouterPolicyTableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTransitRouterPolicyTableResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTransitRouterPolicyTableResponseBody) GetPolicyTableId() *string {
	return s.PolicyTableId
}

func (s *CreateTransitRouterPolicyTableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTransitRouterPolicyTableResponseBody) SetPolicyTableId(v string) *CreateTransitRouterPolicyTableResponseBody {
	s.PolicyTableId = &v
	return s
}

func (s *CreateTransitRouterPolicyTableResponseBody) SetRequestId(v string) *CreateTransitRouterPolicyTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTransitRouterPolicyTableResponseBody) Validate() error {
	return dara.Validate(s)
}
