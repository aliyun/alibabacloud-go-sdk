// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRouteTargetGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateRouteTargetGroupResponseBody
	GetRequestId() *string
	SetRouteTargetGroupId(v string) *CreateRouteTargetGroupResponseBody
	GetRouteTargetGroupId() *string
}

type CreateRouteTargetGroupResponseBody struct {
	// ID of the request.
	//
	// example:
	//
	// 8AA5CE21-2E6A-4530-BDF5-F055849476E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the route target group instance.
	//
	// example:
	//
	// rtg-xxxx
	RouteTargetGroupId *string `json:"RouteTargetGroupId,omitempty" xml:"RouteTargetGroupId,omitempty"`
}

func (s CreateRouteTargetGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRouteTargetGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRouteTargetGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRouteTargetGroupResponseBody) GetRouteTargetGroupId() *string {
	return s.RouteTargetGroupId
}

func (s *CreateRouteTargetGroupResponseBody) SetRequestId(v string) *CreateRouteTargetGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRouteTargetGroupResponseBody) SetRouteTargetGroupId(v string) *CreateRouteTargetGroupResponseBody {
	s.RouteTargetGroupId = &v
	return s
}

func (s *CreateRouteTargetGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
