// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDynamicRouteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateDynamicRouteResponseBody
	GetRequestId() *string
}

type UpdateDynamicRouteResponseBody struct {
	// example:
	//
	// BE4FB974-11BC-5453-9BE1-1606A73EACA6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDynamicRouteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDynamicRouteResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDynamicRouteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDynamicRouteResponseBody) SetRequestId(v string) *UpdateDynamicRouteResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDynamicRouteResponseBody) Validate() error {
	return dara.Validate(s)
}
