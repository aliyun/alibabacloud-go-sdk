// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReplaceTransitRouterRouteTableAssociationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ReplaceTransitRouterRouteTableAssociationResponseBody
	GetRequestId() *string
}

type ReplaceTransitRouterRouteTableAssociationResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 160BD7D3-3D1E-5702-9AF0-56E4B15FCB65
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReplaceTransitRouterRouteTableAssociationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReplaceTransitRouterRouteTableAssociationResponseBody) GoString() string {
	return s.String()
}

func (s *ReplaceTransitRouterRouteTableAssociationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReplaceTransitRouterRouteTableAssociationResponseBody) SetRequestId(v string) *ReplaceTransitRouterRouteTableAssociationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReplaceTransitRouterRouteTableAssociationResponseBody) Validate() error {
	return dara.Validate(s)
}
