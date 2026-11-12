// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLakebaseMountParameterGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetParamGroupCategory(v string) *CreateLakebaseMountParameterGroupRequest
	GetParamGroupCategory() *string
	SetParameterGroupDesc(v string) *CreateLakebaseMountParameterGroupRequest
	GetParameterGroupDesc() *string
	SetParameterGroupName(v string) *CreateLakebaseMountParameterGroupRequest
	GetParameterGroupName() *string
	SetParameters(v string) *CreateLakebaseMountParameterGroupRequest
	GetParameters() *string
	SetPolarFsInstanceId(v string) *CreateLakebaseMountParameterGroupRequest
	GetPolarFsInstanceId() *string
	SetRegionId(v string) *CreateLakebaseMountParameterGroupRequest
	GetRegionId() *string
}

type CreateLakebaseMountParameterGroupRequest struct {
	// The instance edition. Valid values:
	//
	// - basic: Basic Edition.
	//
	// - high_performance: High-performance Edition.
	//
	// This parameter is required.
	//
	// example:
	//
	// basic
	ParamGroupCategory *string `json:"ParamGroupCategory,omitempty" xml:"ParamGroupCategory,omitempty"`
	// The description of the parameter template.
	//
	// example:
	//
	// Test template
	ParameterGroupDesc *string `json:"ParameterGroupDesc,omitempty" xml:"ParameterGroupDesc,omitempty"`
	// The name of the parameter template.
	//
	// This parameter is required.
	//
	// example:
	//
	// pg-test
	ParameterGroupName *string `json:"ParameterGroupName,omitempty" xml:"ParameterGroupName,omitempty"`
	// The JSON string of parameters and their values.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"key1":"value1"}
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
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

func (s CreateLakebaseMountParameterGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLakebaseMountParameterGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateLakebaseMountParameterGroupRequest) GetParamGroupCategory() *string {
	return s.ParamGroupCategory
}

func (s *CreateLakebaseMountParameterGroupRequest) GetParameterGroupDesc() *string {
	return s.ParameterGroupDesc
}

func (s *CreateLakebaseMountParameterGroupRequest) GetParameterGroupName() *string {
	return s.ParameterGroupName
}

func (s *CreateLakebaseMountParameterGroupRequest) GetParameters() *string {
	return s.Parameters
}

func (s *CreateLakebaseMountParameterGroupRequest) GetPolarFsInstanceId() *string {
	return s.PolarFsInstanceId
}

func (s *CreateLakebaseMountParameterGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateLakebaseMountParameterGroupRequest) SetParamGroupCategory(v string) *CreateLakebaseMountParameterGroupRequest {
	s.ParamGroupCategory = &v
	return s
}

func (s *CreateLakebaseMountParameterGroupRequest) SetParameterGroupDesc(v string) *CreateLakebaseMountParameterGroupRequest {
	s.ParameterGroupDesc = &v
	return s
}

func (s *CreateLakebaseMountParameterGroupRequest) SetParameterGroupName(v string) *CreateLakebaseMountParameterGroupRequest {
	s.ParameterGroupName = &v
	return s
}

func (s *CreateLakebaseMountParameterGroupRequest) SetParameters(v string) *CreateLakebaseMountParameterGroupRequest {
	s.Parameters = &v
	return s
}

func (s *CreateLakebaseMountParameterGroupRequest) SetPolarFsInstanceId(v string) *CreateLakebaseMountParameterGroupRequest {
	s.PolarFsInstanceId = &v
	return s
}

func (s *CreateLakebaseMountParameterGroupRequest) SetRegionId(v string) *CreateLakebaseMountParameterGroupRequest {
	s.RegionId = &v
	return s
}

func (s *CreateLakebaseMountParameterGroupRequest) Validate() error {
	return dara.Validate(s)
}
