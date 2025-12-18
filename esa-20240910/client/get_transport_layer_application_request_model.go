// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTransportLayerApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v int64) *GetTransportLayerApplicationRequest
	GetApplicationId() *int64
	SetSiteId(v int64) *GetTransportLayerApplicationRequest
	GetSiteId() *int64
}

type GetTransportLayerApplicationRequest struct {
	// Number of forwarding rules contained in the transport layer acceleration application.
	//
	// This parameter is required.
	//
	// example:
	//
	// 170997271816****
	ApplicationId *int64 `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// Transport layer application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s GetTransportLayerApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTransportLayerApplicationRequest) GoString() string {
	return s.String()
}

func (s *GetTransportLayerApplicationRequest) GetApplicationId() *int64 {
	return s.ApplicationId
}

func (s *GetTransportLayerApplicationRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *GetTransportLayerApplicationRequest) SetApplicationId(v int64) *GetTransportLayerApplicationRequest {
	s.ApplicationId = &v
	return s
}

func (s *GetTransportLayerApplicationRequest) SetSiteId(v int64) *GetTransportLayerApplicationRequest {
	s.SiteId = &v
	return s
}

func (s *GetTransportLayerApplicationRequest) Validate() error {
	return dara.Validate(s)
}
