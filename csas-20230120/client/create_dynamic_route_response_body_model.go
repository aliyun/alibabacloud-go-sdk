// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDynamicRouteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDynamicRouteId(v string) *CreateDynamicRouteResponseBody
	GetDynamicRouteId() *string
	SetRequestId(v string) *CreateDynamicRouteResponseBody
	GetRequestId() *string
}

type CreateDynamicRouteResponseBody struct {
	// example:
	//
	// dr-ca9fddfac7c6****
	DynamicRouteId *string `json:"DynamicRouteId,omitempty" xml:"DynamicRouteId,omitempty"`
	// example:
	//
	// 102350E7-1A20-58F5-9D63-ABEA820AE6E1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDynamicRouteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDynamicRouteResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDynamicRouteResponseBody) GetDynamicRouteId() *string {
	return s.DynamicRouteId
}

func (s *CreateDynamicRouteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDynamicRouteResponseBody) SetDynamicRouteId(v string) *CreateDynamicRouteResponseBody {
	s.DynamicRouteId = &v
	return s
}

func (s *CreateDynamicRouteResponseBody) SetRequestId(v string) *CreateDynamicRouteResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDynamicRouteResponseBody) Validate() error {
	return dara.Validate(s)
}
