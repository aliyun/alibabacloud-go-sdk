// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLakebaseMountParameterGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetParameterGroupId(v string) *DeleteLakebaseMountParameterGroupRequest
	GetParameterGroupId() *string
	SetPolarFsInstanceId(v string) *DeleteLakebaseMountParameterGroupRequest
	GetPolarFsInstanceId() *string
	SetRegionId(v string) *DeleteLakebaseMountParameterGroupRequest
	GetRegionId() *string
}

type DeleteLakebaseMountParameterGroupRequest struct {
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

func (s DeleteLakebaseMountParameterGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteLakebaseMountParameterGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteLakebaseMountParameterGroupRequest) GetParameterGroupId() *string {
	return s.ParameterGroupId
}

func (s *DeleteLakebaseMountParameterGroupRequest) GetPolarFsInstanceId() *string {
	return s.PolarFsInstanceId
}

func (s *DeleteLakebaseMountParameterGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteLakebaseMountParameterGroupRequest) SetParameterGroupId(v string) *DeleteLakebaseMountParameterGroupRequest {
	s.ParameterGroupId = &v
	return s
}

func (s *DeleteLakebaseMountParameterGroupRequest) SetPolarFsInstanceId(v string) *DeleteLakebaseMountParameterGroupRequest {
	s.PolarFsInstanceId = &v
	return s
}

func (s *DeleteLakebaseMountParameterGroupRequest) SetRegionId(v string) *DeleteLakebaseMountParameterGroupRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteLakebaseMountParameterGroupRequest) Validate() error {
	return dara.Validate(s)
}
