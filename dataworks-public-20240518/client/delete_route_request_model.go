// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRouteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *DeleteRouteRequest
	GetId() *int64
}

type DeleteRouteRequest struct {
	// The route ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1000
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteRouteRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRouteRequest) GoString() string {
	return s.String()
}

func (s *DeleteRouteRequest) GetId() *int64 {
	return s.Id
}

func (s *DeleteRouteRequest) SetId(v int64) *DeleteRouteRequest {
	s.Id = &v
	return s
}

func (s *DeleteRouteRequest) Validate() error {
	return dara.Validate(s)
}
