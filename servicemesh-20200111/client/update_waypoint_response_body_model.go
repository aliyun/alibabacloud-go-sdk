// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWaypointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateWaypointResponseBody
	GetRequestId() *string
}

type UpdateWaypointResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 71680038-8009-5073-B43E-C057E9******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateWaypointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateWaypointResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateWaypointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateWaypointResponseBody) SetRequestId(v string) *UpdateWaypointResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateWaypointResponseBody) Validate() error {
	return dara.Validate(s)
}
