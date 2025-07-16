// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyParameterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ModifyParameterRequest
	GetClientToken() *string
	SetDBInstanceId(v string) *ModifyParameterRequest
	GetDBInstanceId() *string
	SetParamLevel(v string) *ModifyParameterRequest
	GetParamLevel() *string
	SetParameterGroupId(v string) *ModifyParameterRequest
	GetParameterGroupId() *string
	SetParameters(v string) *ModifyParameterRequest
	GetParameters() *string
	SetRegionId(v string) *ModifyParameterRequest
	GetRegionId() *string
}

type ModifyParameterRequest struct {
	// example:
	//
	// FEA5DC20-6D8A-5979-97AA-FC57546ADC20
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pxc-hzjasdyuoo
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// example:
	//
	// compute
	ParamLevel       *string `json:"ParamLevel,omitempty" xml:"ParamLevel,omitempty"`
	ParameterGroupId *string `json:"ParameterGroupId,omitempty" xml:"ParameterGroupId,omitempty"`
	// example:
	//
	// {"CONN_POOL_BLOCK_TIMEOUT":6000}
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyParameterRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyParameterRequest) GoString() string {
	return s.String()
}

func (s *ModifyParameterRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyParameterRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyParameterRequest) GetParamLevel() *string {
	return s.ParamLevel
}

func (s *ModifyParameterRequest) GetParameterGroupId() *string {
	return s.ParameterGroupId
}

func (s *ModifyParameterRequest) GetParameters() *string {
	return s.Parameters
}

func (s *ModifyParameterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyParameterRequest) SetClientToken(v string) *ModifyParameterRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyParameterRequest) SetDBInstanceId(v string) *ModifyParameterRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyParameterRequest) SetParamLevel(v string) *ModifyParameterRequest {
	s.ParamLevel = &v
	return s
}

func (s *ModifyParameterRequest) SetParameterGroupId(v string) *ModifyParameterRequest {
	s.ParameterGroupId = &v
	return s
}

func (s *ModifyParameterRequest) SetParameters(v string) *ModifyParameterRequest {
	s.Parameters = &v
	return s
}

func (s *ModifyParameterRequest) SetRegionId(v string) *ModifyParameterRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyParameterRequest) Validate() error {
	return dara.Validate(s)
}
