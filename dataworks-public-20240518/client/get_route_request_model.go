// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRouteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *GetRouteRequest
	GetId() *int64
}

type GetRouteRequest struct {
	// The route ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1000
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetRouteRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRouteRequest) GoString() string {
	return s.String()
}

func (s *GetRouteRequest) GetId() *int64 {
	return s.Id
}

func (s *GetRouteRequest) SetId(v int64) *GetRouteRequest {
	s.Id = &v
	return s
}

func (s *GetRouteRequest) Validate() error {
	return dara.Validate(s)
}
