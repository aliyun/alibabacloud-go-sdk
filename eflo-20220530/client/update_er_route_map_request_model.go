// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateErRouteMapRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateErRouteMapRequest
	GetDescription() *string
	SetErId(v string) *UpdateErRouteMapRequest
	GetErId() *string
	SetErRouteMapId(v string) *UpdateErRouteMapRequest
	GetErRouteMapId() *string
	SetRegionId(v string) *UpdateErRouteMapRequest
	GetRegionId() *string
}

type UpdateErRouteMapRequest struct {
	// The description of the document.
	//
	// example:
	//
	// test-example
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Lingjun HUB ID
	//
	// This parameter is required.
	//
	// example:
	//
	// er-kkopgtne
	ErId *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	// routing policy ID
	//
	// This parameter is required.
	//
	// example:
	//
	// er-rmap-uwglhzom
	ErRouteMapId *string `json:"ErRouteMapId,omitempty" xml:"ErRouteMapId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateErRouteMapRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateErRouteMapRequest) GoString() string {
	return s.String()
}

func (s *UpdateErRouteMapRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateErRouteMapRequest) GetErId() *string {
	return s.ErId
}

func (s *UpdateErRouteMapRequest) GetErRouteMapId() *string {
	return s.ErRouteMapId
}

func (s *UpdateErRouteMapRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateErRouteMapRequest) SetDescription(v string) *UpdateErRouteMapRequest {
	s.Description = &v
	return s
}

func (s *UpdateErRouteMapRequest) SetErId(v string) *UpdateErRouteMapRequest {
	s.ErId = &v
	return s
}

func (s *UpdateErRouteMapRequest) SetErRouteMapId(v string) *UpdateErRouteMapRequest {
	s.ErRouteMapId = &v
	return s
}

func (s *UpdateErRouteMapRequest) SetRegionId(v string) *UpdateErRouteMapRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateErRouteMapRequest) Validate() error {
	return dara.Validate(s)
}
