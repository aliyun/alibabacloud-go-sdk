// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTransitRouterPolicyTableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteTransitRouterPolicyTableResponseBody
	GetRequestId() *string
}

type DeleteTransitRouterPolicyTableResponseBody struct {
	// example:
	//
	// C46FF5A8-C5F0-4024-8262-B16B639225A0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteTransitRouterPolicyTableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteTransitRouterPolicyTableResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTransitRouterPolicyTableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteTransitRouterPolicyTableResponseBody) SetRequestId(v string) *DeleteTransitRouterPolicyTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTransitRouterPolicyTableResponseBody) Validate() error {
	return dara.Validate(s)
}
