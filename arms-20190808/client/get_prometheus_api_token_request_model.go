// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPrometheusApiTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *GetPrometheusApiTokenRequest
	GetRegionId() *string
}

type GetPrometheusApiTokenRequest struct {
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetPrometheusApiTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPrometheusApiTokenRequest) GoString() string {
	return s.String()
}

func (s *GetPrometheusApiTokenRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetPrometheusApiTokenRequest) SetRegionId(v string) *GetPrometheusApiTokenRequest {
	s.RegionId = &v
	return s
}

func (s *GetPrometheusApiTokenRequest) Validate() error {
	return dara.Validate(s)
}
