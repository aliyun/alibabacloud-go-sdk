// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySagPortRouteProtocolResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifySagPortRouteProtocolResponseBody
	GetRequestId() *string
}

type ModifySagPortRouteProtocolResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// CE6642D4-21EB-4168-9BF9-F217953F9892
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifySagPortRouteProtocolResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifySagPortRouteProtocolResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySagPortRouteProtocolResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifySagPortRouteProtocolResponseBody) SetRequestId(v string) *ModifySagPortRouteProtocolResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySagPortRouteProtocolResponseBody) Validate() error {
	return dara.Validate(s)
}
