// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantInstanceToTransitRouterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GrantInstanceToTransitRouterResponseBody
	GetRequestId() *string
}

type GrantInstanceToTransitRouterResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// C6E5992C-A57B-5A6C-9B26-568074DC68BA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GrantInstanceToTransitRouterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GrantInstanceToTransitRouterResponseBody) GoString() string {
	return s.String()
}

func (s *GrantInstanceToTransitRouterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GrantInstanceToTransitRouterResponseBody) SetRequestId(v string) *GrantInstanceToTransitRouterResponseBody {
	s.RequestId = &v
	return s
}

func (s *GrantInstanceToTransitRouterResponseBody) Validate() error {
	return dara.Validate(s)
}
