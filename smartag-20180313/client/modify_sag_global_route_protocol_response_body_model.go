// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySagGlobalRouteProtocolResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifySagGlobalRouteProtocolResponseBody
	GetRequestId() *string
}

type ModifySagGlobalRouteProtocolResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// DB0A026C-A8E5-40AB-977E-3A87DD78F694
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifySagGlobalRouteProtocolResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifySagGlobalRouteProtocolResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySagGlobalRouteProtocolResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifySagGlobalRouteProtocolResponseBody) SetRequestId(v string) *ModifySagGlobalRouteProtocolResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySagGlobalRouteProtocolResponseBody) Validate() error {
	return dara.Validate(s)
}
