// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGreyTagRouteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGreyTagRouteId(v int64) *DeleteGreyTagRouteRequest
	GetGreyTagRouteId() *int64
}

type DeleteGreyTagRouteRequest struct {
	// The rule ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	GreyTagRouteId *int64 `json:"GreyTagRouteId,omitempty" xml:"GreyTagRouteId,omitempty"`
}

func (s DeleteGreyTagRouteRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteGreyTagRouteRequest) GoString() string {
	return s.String()
}

func (s *DeleteGreyTagRouteRequest) GetGreyTagRouteId() *int64 {
	return s.GreyTagRouteId
}

func (s *DeleteGreyTagRouteRequest) SetGreyTagRouteId(v int64) *DeleteGreyTagRouteRequest {
	s.GreyTagRouteId = &v
	return s
}

func (s *DeleteGreyTagRouteRequest) Validate() error {
	return dara.Validate(s)
}
