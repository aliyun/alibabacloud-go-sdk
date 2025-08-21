// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEnsRouteEntryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRouteEntryId(v string) *DeleteEnsRouteEntryRequest
	GetRouteEntryId() *string
}

type DeleteEnsRouteEntryRequest struct {
	// The ID of the route that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// rte-8vbmb2890wiret5maqq25
	RouteEntryId *string `json:"RouteEntryId,omitempty" xml:"RouteEntryId,omitempty"`
}

func (s DeleteEnsRouteEntryRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteEnsRouteEntryRequest) GoString() string {
	return s.String()
}

func (s *DeleteEnsRouteEntryRequest) GetRouteEntryId() *string {
	return s.RouteEntryId
}

func (s *DeleteEnsRouteEntryRequest) SetRouteEntryId(v string) *DeleteEnsRouteEntryRequest {
	s.RouteEntryId = &v
	return s
}

func (s *DeleteEnsRouteEntryRequest) Validate() error {
	return dara.Validate(s)
}
