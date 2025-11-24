// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWaypointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateWaypointResponseBody
	GetRequestId() *string
}

type CreateWaypointResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 71680038-8009-5073-B43E-C057E9******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateWaypointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateWaypointResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWaypointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateWaypointResponseBody) SetRequestId(v string) *CreateWaypointResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateWaypointResponseBody) Validate() error {
	return dara.Validate(s)
}
