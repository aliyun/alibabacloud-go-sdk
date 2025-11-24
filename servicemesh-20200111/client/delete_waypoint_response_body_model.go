// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWaypointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteWaypointResponseBody
	GetRequestId() *string
}

type DeleteWaypointResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// BD65C0AD-D3C6-48D3-8D93-38D2015C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteWaypointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteWaypointResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWaypointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteWaypointResponseBody) SetRequestId(v string) *DeleteWaypointResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteWaypointResponseBody) Validate() error {
	return dara.Validate(s)
}
