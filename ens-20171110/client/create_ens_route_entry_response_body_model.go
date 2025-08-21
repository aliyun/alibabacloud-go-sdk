// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEnsRouteEntryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateEnsRouteEntryResponseBody
	GetRequestId() *string
	SetRouteEntryId(v string) *CreateEnsRouteEntryResponseBody
	GetRouteEntryId() *string
}

type CreateEnsRouteEntryResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the route.
	//
	// example:
	//
	// rte-5vb5q8sk0lyoscx8539ds
	RouteEntryId *string `json:"RouteEntryId,omitempty" xml:"RouteEntryId,omitempty"`
}

func (s CreateEnsRouteEntryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateEnsRouteEntryResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEnsRouteEntryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateEnsRouteEntryResponseBody) GetRouteEntryId() *string {
	return s.RouteEntryId
}

func (s *CreateEnsRouteEntryResponseBody) SetRequestId(v string) *CreateEnsRouteEntryResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateEnsRouteEntryResponseBody) SetRouteEntryId(v string) *CreateEnsRouteEntryResponseBody {
	s.RouteEntryId = &v
	return s
}

func (s *CreateEnsRouteEntryResponseBody) Validate() error {
	return dara.Validate(s)
}
