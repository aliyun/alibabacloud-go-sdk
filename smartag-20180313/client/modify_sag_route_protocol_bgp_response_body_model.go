// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySagRouteProtocolBgpResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifySagRouteProtocolBgpResponseBody
	GetRequestId() *string
}

type ModifySagRouteProtocolBgpResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 48868976-45A2-4E87-B3AA-25089B8B7E49
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifySagRouteProtocolBgpResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifySagRouteProtocolBgpResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySagRouteProtocolBgpResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifySagRouteProtocolBgpResponseBody) SetRequestId(v string) *ModifySagRouteProtocolBgpResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySagRouteProtocolBgpResponseBody) Validate() error {
	return dara.Validate(s)
}
