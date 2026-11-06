// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLakebaseMountParameterGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeleteParameters(v string) *ModifyLakebaseMountParameterGroupRequest
	GetDeleteParameters() *string
	SetParameterGroupId(v string) *ModifyLakebaseMountParameterGroupRequest
	GetParameterGroupId() *string
	SetParameterGroupName(v string) *ModifyLakebaseMountParameterGroupRequest
	GetParameterGroupName() *string
	SetParameters(v string) *ModifyLakebaseMountParameterGroupRequest
	GetParameters() *string
	SetPolarFsInstanceId(v string) *ModifyLakebaseMountParameterGroupRequest
	GetPolarFsInstanceId() *string
	SetRegionId(v string) *ModifyLakebaseMountParameterGroupRequest
	GetRegionId() *string
}

type ModifyLakebaseMountParameterGroupRequest struct {
	// The names of the parameters to delete, separated by commas.
	//
	// example:
	//
	// param1,param2
	DeleteParameters *string `json:"DeleteParameters,omitempty" xml:"DeleteParameters,omitempty"`
	// The parameter template ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pg-xxx
	ParameterGroupId *string `json:"ParameterGroupId,omitempty" xml:"ParameterGroupId,omitempty"`
	// The new parameter template name.
	//
	// example:
	//
	// pg-new
	ParameterGroupName *string `json:"ParameterGroupName,omitempty" xml:"ParameterGroupName,omitempty"`
	// The JSON string of added or modified parameter key-value pairs.
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

func (s ModifyLakebaseMountParameterGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyLakebaseMountParameterGroupRequest) GoString() string {
	return s.String()
}

func (s *ModifyLakebaseMountParameterGroupRequest) GetDeleteParameters() *string {
	return s.DeleteParameters
}

func (s *ModifyLakebaseMountParameterGroupRequest) GetParameterGroupId() *string {
	return s.ParameterGroupId
}

func (s *ModifyLakebaseMountParameterGroupRequest) GetParameterGroupName() *string {
	return s.ParameterGroupName
}

func (s *ModifyLakebaseMountParameterGroupRequest) GetParameters() *string {
	return s.Parameters
}

func (s *ModifyLakebaseMountParameterGroupRequest) GetPolarFsInstanceId() *string {
	return s.PolarFsInstanceId
}

func (s *ModifyLakebaseMountParameterGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyLakebaseMountParameterGroupRequest) SetDeleteParameters(v string) *ModifyLakebaseMountParameterGroupRequest {
	s.DeleteParameters = &v
	return s
}

func (s *ModifyLakebaseMountParameterGroupRequest) SetParameterGroupId(v string) *ModifyLakebaseMountParameterGroupRequest {
	s.ParameterGroupId = &v
	return s
}

func (s *ModifyLakebaseMountParameterGroupRequest) SetParameterGroupName(v string) *ModifyLakebaseMountParameterGroupRequest {
	s.ParameterGroupName = &v
	return s
}

func (s *ModifyLakebaseMountParameterGroupRequest) SetParameters(v string) *ModifyLakebaseMountParameterGroupRequest {
	s.Parameters = &v
	return s
}

func (s *ModifyLakebaseMountParameterGroupRequest) SetPolarFsInstanceId(v string) *ModifyLakebaseMountParameterGroupRequest {
	s.PolarFsInstanceId = &v
	return s
}

func (s *ModifyLakebaseMountParameterGroupRequest) SetRegionId(v string) *ModifyLakebaseMountParameterGroupRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyLakebaseMountParameterGroupRequest) Validate() error {
	return dara.Validate(s)
}
