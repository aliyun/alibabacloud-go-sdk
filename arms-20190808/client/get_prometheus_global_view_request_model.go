// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPrometheusGlobalViewRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGlobalViewClusterId(v string) *GetPrometheusGlobalViewRequest
	GetGlobalViewClusterId() *string
	SetRegionId(v string) *GetPrometheusGlobalViewRequest
	GetRegionId() *string
}

type GetPrometheusGlobalViewRequest struct {
	// The ID of the global aggregation instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// global-v2-cn-1478326682034601-vss8pd0i
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

func (s GetPrometheusGlobalViewRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPrometheusGlobalViewRequest) GoString() string {
	return s.String()
}

func (s *GetPrometheusGlobalViewRequest) GetGlobalViewClusterId() *string {
	return s.GlobalViewClusterId
}

func (s *GetPrometheusGlobalViewRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetPrometheusGlobalViewRequest) SetGlobalViewClusterId(v string) *GetPrometheusGlobalViewRequest {
	s.GlobalViewClusterId = &v
	return s
}

func (s *GetPrometheusGlobalViewRequest) SetRegionId(v string) *GetPrometheusGlobalViewRequest {
	s.RegionId = &v
	return s
}

func (s *GetPrometheusGlobalViewRequest) Validate() error {
	return dara.Validate(s)
}
