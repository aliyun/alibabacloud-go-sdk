// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNetworkOptimizationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *GetNetworkOptimizationRequest
	GetConfigId() *int64
	SetSiteId(v int64) *GetNetworkOptimizationRequest
	GetSiteId() *int64
}

type GetNetworkOptimizationRequest struct {
	// This parameter is required.
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// This parameter is required.
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s GetNetworkOptimizationRequest) String() string {
	return dara.Prettify(s)
}

func (s GetNetworkOptimizationRequest) GoString() string {
	return s.String()
}

func (s *GetNetworkOptimizationRequest) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *GetNetworkOptimizationRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *GetNetworkOptimizationRequest) SetConfigId(v int64) *GetNetworkOptimizationRequest {
	s.ConfigId = &v
	return s
}

func (s *GetNetworkOptimizationRequest) SetSiteId(v int64) *GetNetworkOptimizationRequest {
	s.SiteId = &v
	return s
}

func (s *GetNetworkOptimizationRequest) Validate() error {
	return dara.Validate(s)
}
