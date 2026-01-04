// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRouteTargetGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteRouteTargetGroupResponseBody
	GetRequestId() *string
}

type DeleteRouteTargetGroupResponseBody struct {
	// ID of the request
	//
	// example:
	//
	// 4EC47282-1B74-4534-BD0E-403F3EE64CAF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteRouteTargetGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteRouteTargetGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRouteTargetGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteRouteTargetGroupResponseBody) SetRequestId(v string) *DeleteRouteTargetGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRouteTargetGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
