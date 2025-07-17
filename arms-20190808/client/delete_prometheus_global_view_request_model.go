// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePrometheusGlobalViewRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGlobalViewClusterId(v string) *DeletePrometheusGlobalViewRequest
	GetGlobalViewClusterId() *string
	SetRegionId(v string) *DeletePrometheusGlobalViewRequest
	GetRegionId() *string
}

type DeletePrometheusGlobalViewRequest struct {
	// The ID of the global aggregation instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// global-v2-cn-1670100631025794-amaykca4
	GlobalViewClusterId *string `json:"GlobalViewClusterId,omitempty" xml:"GlobalViewClusterId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeletePrometheusGlobalViewRequest) String() string {
	return dara.Prettify(s)
}

func (s DeletePrometheusGlobalViewRequest) GoString() string {
	return s.String()
}

func (s *DeletePrometheusGlobalViewRequest) GetGlobalViewClusterId() *string {
	return s.GlobalViewClusterId
}

func (s *DeletePrometheusGlobalViewRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeletePrometheusGlobalViewRequest) SetGlobalViewClusterId(v string) *DeletePrometheusGlobalViewRequest {
	s.GlobalViewClusterId = &v
	return s
}

func (s *DeletePrometheusGlobalViewRequest) SetRegionId(v string) *DeletePrometheusGlobalViewRequest {
	s.RegionId = &v
	return s
}

func (s *DeletePrometheusGlobalViewRequest) Validate() error {
	return dara.Validate(s)
}
