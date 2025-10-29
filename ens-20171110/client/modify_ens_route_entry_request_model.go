// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyEnsRouteEntryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *ModifyEnsRouteEntryRequest
	GetDescription() *string
	SetRouteEntryId(v string) *ModifyEnsRouteEntryRequest
	GetRouteEntryId() *string
	SetRouteEntryName(v string) *ModifyEnsRouteEntryRequest
	GetRouteEntryName() *string
}

type ModifyEnsRouteEntryRequest struct {
	// The description of the route entry. The description must be 1 to 256 characters in length and cannot start with http:// or https://.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the custom route.
	//
	// This parameter is required.
	//
	// example:
	//
	// rte-5****
	RouteEntryId *string `json:"RouteEntryId,omitempty" xml:"RouteEntryId,omitempty"`
	// The name of the route.
	//
	// The name must be 1 to 128 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// test
	RouteEntryName *string `json:"RouteEntryName,omitempty" xml:"RouteEntryName,omitempty"`
}

func (s ModifyEnsRouteEntryRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyEnsRouteEntryRequest) GoString() string {
	return s.String()
}

func (s *ModifyEnsRouteEntryRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyEnsRouteEntryRequest) GetRouteEntryId() *string {
	return s.RouteEntryId
}

func (s *ModifyEnsRouteEntryRequest) GetRouteEntryName() *string {
	return s.RouteEntryName
}

func (s *ModifyEnsRouteEntryRequest) SetDescription(v string) *ModifyEnsRouteEntryRequest {
	s.Description = &v
	return s
}

func (s *ModifyEnsRouteEntryRequest) SetRouteEntryId(v string) *ModifyEnsRouteEntryRequest {
	s.RouteEntryId = &v
	return s
}

func (s *ModifyEnsRouteEntryRequest) SetRouteEntryName(v string) *ModifyEnsRouteEntryRequest {
	s.RouteEntryName = &v
	return s
}

func (s *ModifyEnsRouteEntryRequest) Validate() error {
	return dara.Validate(s)
}
