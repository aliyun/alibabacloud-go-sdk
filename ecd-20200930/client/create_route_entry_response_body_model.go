// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRouteEntryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateRouteEntryResponseBody
	GetRequestId() *string
	SetRouteEntryId(v string) *CreateRouteEntryResponseBody
	GetRouteEntryId() *string
}

type CreateRouteEntryResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RouteEntryId *string `json:"RouteEntryId,omitempty" xml:"RouteEntryId,omitempty"`
}

func (s CreateRouteEntryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRouteEntryResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRouteEntryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRouteEntryResponseBody) GetRouteEntryId() *string {
	return s.RouteEntryId
}

func (s *CreateRouteEntryResponseBody) SetRequestId(v string) *CreateRouteEntryResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRouteEntryResponseBody) SetRouteEntryId(v string) *CreateRouteEntryResponseBody {
	s.RouteEntryId = &v
	return s
}

func (s *CreateRouteEntryResponseBody) Validate() error {
	return dara.Validate(s)
}
