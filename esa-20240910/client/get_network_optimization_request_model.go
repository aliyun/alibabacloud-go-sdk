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
	// Configuration ConfigId, which can be obtained by calling the [ListNetworkOptimizations](https://help.aliyun.com/document_detail/2869051.html) interface.
	//
	// This parameter is required.
	//
	// example:
	//
	// 35281609698****
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// Site ID, which can be obtained by calling the [ListSites](~~ListSites~~) interface.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12312312213212
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
