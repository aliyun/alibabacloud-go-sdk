// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLakebaseMountParameterGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetParamGroupCategory(v string) *DescribeLakebaseMountParameterGroupsRequest
	GetParamGroupCategory() *string
	SetPolarFsInstanceId(v string) *DescribeLakebaseMountParameterGroupsRequest
	GetPolarFsInstanceId() *string
	SetRegionId(v string) *DescribeLakebaseMountParameterGroupsRequest
	GetRegionId() *string
}

type DescribeLakebaseMountParameterGroupsRequest struct {
	// The instance category. Valid values: basic (Basic Edition) and high_performance (High-performance Edition).
	//
	// example:
	//
	// basic
	ParamGroupCategory *string `json:"ParamGroupCategory,omitempty" xml:"ParamGroupCategory,omitempty"`
	// The PolarFS instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pfs-xxx
	PolarFsInstanceId *string `json:"PolarFsInstanceId,omitempty" xml:"PolarFsInstanceId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeLakebaseMountParameterGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLakebaseMountParameterGroupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeLakebaseMountParameterGroupsRequest) GetParamGroupCategory() *string {
	return s.ParamGroupCategory
}

func (s *DescribeLakebaseMountParameterGroupsRequest) GetPolarFsInstanceId() *string {
	return s.PolarFsInstanceId
}

func (s *DescribeLakebaseMountParameterGroupsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLakebaseMountParameterGroupsRequest) SetParamGroupCategory(v string) *DescribeLakebaseMountParameterGroupsRequest {
	s.ParamGroupCategory = &v
	return s
}

func (s *DescribeLakebaseMountParameterGroupsRequest) SetPolarFsInstanceId(v string) *DescribeLakebaseMountParameterGroupsRequest {
	s.PolarFsInstanceId = &v
	return s
}

func (s *DescribeLakebaseMountParameterGroupsRequest) SetRegionId(v string) *DescribeLakebaseMountParameterGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLakebaseMountParameterGroupsRequest) Validate() error {
	return dara.Validate(s)
}
