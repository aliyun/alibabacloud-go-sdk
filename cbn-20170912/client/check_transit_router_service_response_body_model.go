// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckTransitRouterServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEnabled(v string) *CheckTransitRouterServiceResponseBody
	GetEnabled() *string
	SetRequestId(v string) *CheckTransitRouterServiceResponseBody
	GetRequestId() *string
}

type CheckTransitRouterServiceResponseBody struct {
	// Indicates whether the transit router feature is activated.
	//
	// 	- **true**: activated
	//
	// 	- If this value is not returned, the system prompts that the current account does not have the transit router feature activated.
	//
	// example:
	//
	// true
	Enabled *string `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 5D93C8B9-C354-5C3E-BEFB-BA8A2C314D68
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckTransitRouterServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckTransitRouterServiceResponseBody) GoString() string {
	return s.String()
}

func (s *CheckTransitRouterServiceResponseBody) GetEnabled() *string {
	return s.Enabled
}

func (s *CheckTransitRouterServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckTransitRouterServiceResponseBody) SetEnabled(v string) *CheckTransitRouterServiceResponseBody {
	s.Enabled = &v
	return s
}

func (s *CheckTransitRouterServiceResponseBody) SetRequestId(v string) *CheckTransitRouterServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckTransitRouterServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
