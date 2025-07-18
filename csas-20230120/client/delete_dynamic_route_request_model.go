// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDynamicRouteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDynamicRouteId(v string) *DeleteDynamicRouteRequest
	GetDynamicRouteId() *string
}

type DeleteDynamicRouteRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// dr-ca9fddfac7c6****
	DynamicRouteId *string `json:"DynamicRouteId,omitempty" xml:"DynamicRouteId,omitempty"`
}

func (s DeleteDynamicRouteRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDynamicRouteRequest) GoString() string {
	return s.String()
}

func (s *DeleteDynamicRouteRequest) GetDynamicRouteId() *string {
	return s.DynamicRouteId
}

func (s *DeleteDynamicRouteRequest) SetDynamicRouteId(v string) *DeleteDynamicRouteRequest {
	s.DynamicRouteId = &v
	return s
}

func (s *DeleteDynamicRouteRequest) Validate() error {
	return dara.Validate(s)
}
