// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRouteMapResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteRouteMapResponseBody
	GetRequestId() *string
}

type DeleteRouteMapResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// F7DDDC17-FA06-4AC2-8F35-59D2470FCFC1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteRouteMapResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteRouteMapResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRouteMapResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteRouteMapResponseBody) SetRequestId(v string) *DeleteRouteMapResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRouteMapResponseBody) Validate() error {
	return dara.Validate(s)
}
