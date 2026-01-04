// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRouteTargetGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateRouteTargetGroupResponseBody
	GetRequestId() *string
}

type UpdateRouteTargetGroupResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// 0ED8D006-F706-4D23-88ED-E11ED28DCAC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateRouteTargetGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateRouteTargetGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRouteTargetGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateRouteTargetGroupResponseBody) SetRequestId(v string) *UpdateRouteTargetGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRouteTargetGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
