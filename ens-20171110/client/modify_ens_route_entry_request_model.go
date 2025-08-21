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
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// rte-5****
	RouteEntryId *string `json:"RouteEntryId,omitempty" xml:"RouteEntryId,omitempty"`
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
