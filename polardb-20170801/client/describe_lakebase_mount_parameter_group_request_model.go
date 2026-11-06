// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLakebaseMountParameterGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetParameterGroupId(v string) *DescribeLakebaseMountParameterGroupRequest
	GetParameterGroupId() *string
	SetPolarFsInstanceId(v string) *DescribeLakebaseMountParameterGroupRequest
	GetPolarFsInstanceId() *string
	SetRegionId(v string) *DescribeLakebaseMountParameterGroupRequest
	GetRegionId() *string
}

type DescribeLakebaseMountParameterGroupRequest struct {
	// The parameter template ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pg-xxx
	ParameterGroupId *string `json:"ParameterGroupId,omitempty" xml:"ParameterGroupId,omitempty"`
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

func (s DescribeLakebaseMountParameterGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLakebaseMountParameterGroupRequest) GoString() string {
	return s.String()
}

func (s *DescribeLakebaseMountParameterGroupRequest) GetParameterGroupId() *string {
	return s.ParameterGroupId
}

func (s *DescribeLakebaseMountParameterGroupRequest) GetPolarFsInstanceId() *string {
	return s.PolarFsInstanceId
}

func (s *DescribeLakebaseMountParameterGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLakebaseMountParameterGroupRequest) SetParameterGroupId(v string) *DescribeLakebaseMountParameterGroupRequest {
	s.ParameterGroupId = &v
	return s
}

func (s *DescribeLakebaseMountParameterGroupRequest) SetPolarFsInstanceId(v string) *DescribeLakebaseMountParameterGroupRequest {
	s.PolarFsInstanceId = &v
	return s
}

func (s *DescribeLakebaseMountParameterGroupRequest) SetRegionId(v string) *DescribeLakebaseMountParameterGroupRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLakebaseMountParameterGroupRequest) Validate() error {
	return dara.Validate(s)
}
