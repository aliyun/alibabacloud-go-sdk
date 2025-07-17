// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPrometheusInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *GetPrometheusInstanceRequest
	GetClusterId() *string
	SetRegionId(v string) *GetPrometheusInstanceRequest
	GetRegionId() *string
}

type GetPrometheusInstanceRequest struct {
	// The ID of the Prometheus instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// c77f6f2397ea74672872acf5e31374a27
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetPrometheusInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPrometheusInstanceRequest) GoString() string {
	return s.String()
}

func (s *GetPrometheusInstanceRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetPrometheusInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetPrometheusInstanceRequest) SetClusterId(v string) *GetPrometheusInstanceRequest {
	s.ClusterId = &v
	return s
}

func (s *GetPrometheusInstanceRequest) SetRegionId(v string) *GetPrometheusInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *GetPrometheusInstanceRequest) Validate() error {
	return dara.Validate(s)
}
