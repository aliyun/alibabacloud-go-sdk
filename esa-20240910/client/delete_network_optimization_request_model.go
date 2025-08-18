// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNetworkOptimizationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *DeleteNetworkOptimizationRequest
	GetConfigId() *int64
	SetSiteId(v int64) *DeleteNetworkOptimizationRequest
	GetSiteId() *int64
}

type DeleteNetworkOptimizationRequest struct {
	// ConfigId of the configuration, which can be obtained by calling the ListNetworkOptimizations.
	//
	// This parameter is required.
	//
	// example:
	//
	// 352816**********
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// Site ID, which can be obtained by calling the [ListSites](~~ListSites~~) interface.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s DeleteNetworkOptimizationRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteNetworkOptimizationRequest) GoString() string {
	return s.String()
}

func (s *DeleteNetworkOptimizationRequest) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *DeleteNetworkOptimizationRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *DeleteNetworkOptimizationRequest) SetConfigId(v int64) *DeleteNetworkOptimizationRequest {
	s.ConfigId = &v
	return s
}

func (s *DeleteNetworkOptimizationRequest) SetSiteId(v int64) *DeleteNetworkOptimizationRequest {
	s.SiteId = &v
	return s
}

func (s *DeleteNetworkOptimizationRequest) Validate() error {
	return dara.Validate(s)
}
