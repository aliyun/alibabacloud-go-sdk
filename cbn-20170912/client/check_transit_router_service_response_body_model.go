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
	// Indicates whether the transit router service is activated for the current Alibaba Cloud account.
	//
	// - **true**: The service is activated.
	//
	// - If this parameter is not returned, the transit router service is not activated for the current Alibaba Cloud account, and the system returns a corresponding message.
	//
	// example:
	//
	// true
	Enabled *string `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The request ID.
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
