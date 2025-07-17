// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPrometheusGlobalViewRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *ListPrometheusGlobalViewRequest
	GetRegionId() *string
}

type ListPrometheusGlobalViewRequest struct {
	// The ID of the region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListPrometheusGlobalViewRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPrometheusGlobalViewRequest) GoString() string {
	return s.String()
}

func (s *ListPrometheusGlobalViewRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListPrometheusGlobalViewRequest) SetRegionId(v string) *ListPrometheusGlobalViewRequest {
	s.RegionId = &v
	return s
}

func (s *ListPrometheusGlobalViewRequest) Validate() error {
	return dara.Validate(s)
}
