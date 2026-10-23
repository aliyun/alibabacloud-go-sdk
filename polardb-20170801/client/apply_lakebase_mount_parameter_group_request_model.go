// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyLakebaseMountParameterGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetParameterGroupId(v string) *ApplyLakebaseMountParameterGroupRequest
	GetParameterGroupId() *string
	SetPolarFsInstanceId(v string) *ApplyLakebaseMountParameterGroupRequest
	GetPolarFsInstanceId() *string
	SetRegionId(v string) *ApplyLakebaseMountParameterGroupRequest
	GetRegionId() *string
}

type ApplyLakebaseMountParameterGroupRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pagp-xxx
	ParameterGroupId *string `json:"ParameterGroupId,omitempty" xml:"ParameterGroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pfs-xxx
	PolarFsInstanceId *string `json:"PolarFsInstanceId,omitempty" xml:"PolarFsInstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ApplyLakebaseMountParameterGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ApplyLakebaseMountParameterGroupRequest) GoString() string {
	return s.String()
}

func (s *ApplyLakebaseMountParameterGroupRequest) GetParameterGroupId() *string {
	return s.ParameterGroupId
}

func (s *ApplyLakebaseMountParameterGroupRequest) GetPolarFsInstanceId() *string {
	return s.PolarFsInstanceId
}

func (s *ApplyLakebaseMountParameterGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ApplyLakebaseMountParameterGroupRequest) SetParameterGroupId(v string) *ApplyLakebaseMountParameterGroupRequest {
	s.ParameterGroupId = &v
	return s
}

func (s *ApplyLakebaseMountParameterGroupRequest) SetPolarFsInstanceId(v string) *ApplyLakebaseMountParameterGroupRequest {
	s.PolarFsInstanceId = &v
	return s
}

func (s *ApplyLakebaseMountParameterGroupRequest) SetRegionId(v string) *ApplyLakebaseMountParameterGroupRequest {
	s.RegionId = &v
	return s
}

func (s *ApplyLakebaseMountParameterGroupRequest) Validate() error {
	return dara.Validate(s)
}
