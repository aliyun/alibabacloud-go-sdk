// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDynamicRouteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDynamicRouteId(v string) *GetDynamicRouteRequest
	GetDynamicRouteId() *string
}

type GetDynamicRouteRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// dr-16ff07c8207d****
	DynamicRouteId *string `json:"DynamicRouteId,omitempty" xml:"DynamicRouteId,omitempty"`
}

func (s GetDynamicRouteRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDynamicRouteRequest) GoString() string {
	return s.String()
}

func (s *GetDynamicRouteRequest) GetDynamicRouteId() *string {
	return s.DynamicRouteId
}

func (s *GetDynamicRouteRequest) SetDynamicRouteId(v string) *GetDynamicRouteRequest {
	s.DynamicRouteId = &v
	return s
}

func (s *GetDynamicRouteRequest) Validate() error {
	return dara.Validate(s)
}
