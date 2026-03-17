// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySagRouteProtocolOspfResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifySagRouteProtocolOspfResponseBody
	GetRequestId() *string
}

type ModifySagRouteProtocolOspfResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 191DC00D-00C0-475C-99B8-ADBB82496405
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifySagRouteProtocolOspfResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifySagRouteProtocolOspfResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySagRouteProtocolOspfResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifySagRouteProtocolOspfResponseBody) SetRequestId(v string) *ModifySagRouteProtocolOspfResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySagRouteProtocolOspfResponseBody) Validate() error {
	return dara.Validate(s)
}
