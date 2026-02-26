// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeParametersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigParameters(v *DescribeParametersResponseBodyConfigParameters) *DescribeParametersResponseBody
	GetConfigParameters() *DescribeParametersResponseBodyConfigParameters
	SetEngine(v string) *DescribeParametersResponseBody
	GetEngine() *string
	SetEngineVersion(v string) *DescribeParametersResponseBody
	GetEngineVersion() *string
	SetParamGroupInfo(v *DescribeParametersResponseBodyParamGroupInfo) *DescribeParametersResponseBody
	GetParamGroupInfo() *DescribeParametersResponseBodyParamGroupInfo
	SetRequestId(v string) *DescribeParametersResponseBody
	GetRequestId() *string
	SetRunningParameters(v *DescribeParametersResponseBodyRunningParameters) *DescribeParametersResponseBody
	GetRunningParameters() *DescribeParametersResponseBodyRunningParameters
}

type DescribeParametersResponseBody struct {
	ConfigParameters *DescribeParametersResponseBodyConfigParameters `json:"ConfigParameters,omitempty" xml:"ConfigParameters,omitempty" type:"Struct"`
	// The type of the database engine.
	//
	// example:
	//
	// MySQL
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The version of the database engine.
	//
	// example:
	//
	// 5.5
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The information about the parameter template.
	ParamGroupInfo *DescribeParametersResponseBodyParamGroupInfo `json:"ParamGroupInfo,omitempty" xml:"ParamGroupInfo,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId         *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RunningParameters *DescribeParametersResponseBodyRunningParameters `json:"RunningParameters,omitempty" xml:"RunningParameters,omitempty" type:"Struct"`
}

func (s DescribeParametersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeParametersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeParametersResponseBody) GetConfigParameters() *DescribeParametersResponseBodyConfigParameters {
	return s.ConfigParameters
}

func (s *DescribeParametersResponseBody) GetEngine() *string {
	return s.Engine
}

func (s *DescribeParametersResponseBody) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *DescribeParametersResponseBody) GetParamGroupInfo() *DescribeParametersResponseBodyParamGroupInfo {
	return s.ParamGroupInfo
}

func (s *DescribeParametersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeParametersResponseBody) GetRunningParameters() *DescribeParametersResponseBodyRunningParameters {
	return s.RunningParameters
}

func (s *DescribeParametersResponseBody) SetConfigParameters(v *DescribeParametersResponseBodyConfigParameters) *DescribeParametersResponseBody {
	s.ConfigParameters = v
	return s
}

func (s *DescribeParametersResponseBody) SetEngine(v string) *DescribeParametersResponseBody {
	s.Engine = &v
	return s
}

func (s *DescribeParametersResponseBody) SetEngineVersion(v string) *DescribeParametersResponseBody {
	s.EngineVersion = &v
	return s
}

func (s *DescribeParametersResponseBody) SetParamGroupInfo(v *DescribeParametersResponseBodyParamGroupInfo) *DescribeParametersResponseBody {
	s.ParamGroupInfo = v
	return s
}

func (s *DescribeParametersResponseBody) SetRequestId(v string) *DescribeParametersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeParametersResponseBody) SetRunningParameters(v *DescribeParametersResponseBodyRunningParameters) *DescribeParametersResponseBody {
	s.RunningParameters = v
	return s
}

func (s *DescribeParametersResponseBody) Validate() error {
	if s.ConfigParameters != nil {
		if err := s.ConfigParameters.Validate(); err != nil {
			return err
		}
	}
	if s.ParamGroupInfo != nil {
		if err := s.ParamGroupInfo.Validate(); err != nil {
			return err
		}
	}
	if s.RunningParameters != nil {
		if err := s.RunningParameters.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeParametersResponseBodyConfigParameters struct {
	DBInstanceParameter []*DescribeParametersResponseBodyConfigParametersDBInstanceParameter `json:"DBInstanceParameter,omitempty" xml:"DBInstanceParameter,omitempty" type:"Repeated"`
}

func (s DescribeParametersResponseBodyConfigParameters) String() string {
	return dara.Prettify(s)
}

func (s DescribeParametersResponseBodyConfigParameters) GoString() string {
	return s.String()
}

func (s *DescribeParametersResponseBodyConfigParameters) GetDBInstanceParameter() []*DescribeParametersResponseBodyConfigParametersDBInstanceParameter {
	return s.DBInstanceParameter
}

func (s *DescribeParametersResponseBodyConfigParameters) SetDBInstanceParameter(v []*DescribeParametersResponseBodyConfigParametersDBInstanceParameter) *DescribeParametersResponseBodyConfigParameters {
	s.DBInstanceParameter = v
	return s
}

func (s *DescribeParametersResponseBodyConfigParameters) Validate() error {
	if s.DBInstanceParameter != nil {
		for _, item := range s.DBInstanceParameter {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeParametersResponseBodyConfigParametersDBInstanceParameter struct {
	ParameterDescription *string `json:"ParameterDescription,omitempty" xml:"ParameterDescription,omitempty"`
	ParameterName        *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
	ParameterValue       *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s DescribeParametersResponseBodyConfigParametersDBInstanceParameter) String() string {
	return dara.Prettify(s)
}

func (s DescribeParametersResponseBodyConfigParametersDBInstanceParameter) GoString() string {
	return s.String()
}

func (s *DescribeParametersResponseBodyConfigParametersDBInstanceParameter) GetParameterDescription() *string {
	return s.ParameterDescription
}

func (s *DescribeParametersResponseBodyConfigParametersDBInstanceParameter) GetParameterName() *string {
	return s.ParameterName
}

func (s *DescribeParametersResponseBodyConfigParametersDBInstanceParameter) GetParameterValue() *string {
	return s.ParameterValue
}

func (s *DescribeParametersResponseBodyConfigParametersDBInstanceParameter) SetParameterDescription(v string) *DescribeParametersResponseBodyConfigParametersDBInstanceParameter {
	s.ParameterDescription = &v
	return s
}

func (s *DescribeParametersResponseBodyConfigParametersDBInstanceParameter) SetParameterName(v string) *DescribeParametersResponseBodyConfigParametersDBInstanceParameter {
	s.ParameterName = &v
	return s
}

func (s *DescribeParametersResponseBodyConfigParametersDBInstanceParameter) SetParameterValue(v string) *DescribeParametersResponseBodyConfigParametersDBInstanceParameter {
	s.ParameterValue = &v
	return s
}

func (s *DescribeParametersResponseBodyConfigParametersDBInstanceParameter) Validate() error {
	return dara.Validate(s)
}

type DescribeParametersResponseBodyParamGroupInfo struct {
	// The ID of the parameter template.
	//
	// example:
	//
	// rpg-sys-01040401010200
	ParamGroupId *string `json:"ParamGroupId,omitempty" xml:"ParamGroupId,omitempty"`
	// The description of the parameter template.
	//
	// example:
	//
	// sync_binlog=1000, innodb_flush_log_at_trx_commit=2, async
	ParameterGroupDesc *string `json:"ParameterGroupDesc,omitempty" xml:"ParameterGroupDesc,omitempty"`
	// The name of the parameter template.
	//
	// example:
	//
	// mysql_innodb_8.0_basic_normal_high
	ParameterGroupName *string `json:"ParameterGroupName,omitempty" xml:"ParameterGroupName,omitempty"`
	// The type of the parameter template.
	//
	// example:
	//
	// 0
	ParameterGroupType *string `json:"ParameterGroupType,omitempty" xml:"ParameterGroupType,omitempty"`
}

func (s DescribeParametersResponseBodyParamGroupInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeParametersResponseBodyParamGroupInfo) GoString() string {
	return s.String()
}

func (s *DescribeParametersResponseBodyParamGroupInfo) GetParamGroupId() *string {
	return s.ParamGroupId
}

func (s *DescribeParametersResponseBodyParamGroupInfo) GetParameterGroupDesc() *string {
	return s.ParameterGroupDesc
}

func (s *DescribeParametersResponseBodyParamGroupInfo) GetParameterGroupName() *string {
	return s.ParameterGroupName
}

func (s *DescribeParametersResponseBodyParamGroupInfo) GetParameterGroupType() *string {
	return s.ParameterGroupType
}

func (s *DescribeParametersResponseBodyParamGroupInfo) SetParamGroupId(v string) *DescribeParametersResponseBodyParamGroupInfo {
	s.ParamGroupId = &v
	return s
}

func (s *DescribeParametersResponseBodyParamGroupInfo) SetParameterGroupDesc(v string) *DescribeParametersResponseBodyParamGroupInfo {
	s.ParameterGroupDesc = &v
	return s
}

func (s *DescribeParametersResponseBodyParamGroupInfo) SetParameterGroupName(v string) *DescribeParametersResponseBodyParamGroupInfo {
	s.ParameterGroupName = &v
	return s
}

func (s *DescribeParametersResponseBodyParamGroupInfo) SetParameterGroupType(v string) *DescribeParametersResponseBodyParamGroupInfo {
	s.ParameterGroupType = &v
	return s
}

func (s *DescribeParametersResponseBodyParamGroupInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeParametersResponseBodyRunningParameters struct {
	DBInstanceParameter []*DescribeParametersResponseBodyRunningParametersDBInstanceParameter `json:"DBInstanceParameter,omitempty" xml:"DBInstanceParameter,omitempty" type:"Repeated"`
}

func (s DescribeParametersResponseBodyRunningParameters) String() string {
	return dara.Prettify(s)
}

func (s DescribeParametersResponseBodyRunningParameters) GoString() string {
	return s.String()
}

func (s *DescribeParametersResponseBodyRunningParameters) GetDBInstanceParameter() []*DescribeParametersResponseBodyRunningParametersDBInstanceParameter {
	return s.DBInstanceParameter
}

func (s *DescribeParametersResponseBodyRunningParameters) SetDBInstanceParameter(v []*DescribeParametersResponseBodyRunningParametersDBInstanceParameter) *DescribeParametersResponseBodyRunningParameters {
	s.DBInstanceParameter = v
	return s
}

func (s *DescribeParametersResponseBodyRunningParameters) Validate() error {
	if s.DBInstanceParameter != nil {
		for _, item := range s.DBInstanceParameter {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeParametersResponseBodyRunningParametersDBInstanceParameter struct {
	ParameterDefaultValue *string `json:"ParameterDefaultValue,omitempty" xml:"ParameterDefaultValue,omitempty"`
	ParameterDescription  *string `json:"ParameterDescription,omitempty" xml:"ParameterDescription,omitempty"`
	ParameterName         *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
	ParameterValue        *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
	ParameterValueRange   *string `json:"ParameterValueRange,omitempty" xml:"ParameterValueRange,omitempty"`
}

func (s DescribeParametersResponseBodyRunningParametersDBInstanceParameter) String() string {
	return dara.Prettify(s)
}

func (s DescribeParametersResponseBodyRunningParametersDBInstanceParameter) GoString() string {
	return s.String()
}

func (s *DescribeParametersResponseBodyRunningParametersDBInstanceParameter) GetParameterDefaultValue() *string {
	return s.ParameterDefaultValue
}

func (s *DescribeParametersResponseBodyRunningParametersDBInstanceParameter) GetParameterDescription() *string {
	return s.ParameterDescription
}

func (s *DescribeParametersResponseBodyRunningParametersDBInstanceParameter) GetParameterName() *string {
	return s.ParameterName
}

func (s *DescribeParametersResponseBodyRunningParametersDBInstanceParameter) GetParameterValue() *string {
	return s.ParameterValue
}

func (s *DescribeParametersResponseBodyRunningParametersDBInstanceParameter) GetParameterValueRange() *string {
	return s.ParameterValueRange
}

func (s *DescribeParametersResponseBodyRunningParametersDBInstanceParameter) SetParameterDefaultValue(v string) *DescribeParametersResponseBodyRunningParametersDBInstanceParameter {
	s.ParameterDefaultValue = &v
	return s
}

func (s *DescribeParametersResponseBodyRunningParametersDBInstanceParameter) SetParameterDescription(v string) *DescribeParametersResponseBodyRunningParametersDBInstanceParameter {
	s.ParameterDescription = &v
	return s
}

func (s *DescribeParametersResponseBodyRunningParametersDBInstanceParameter) SetParameterName(v string) *DescribeParametersResponseBodyRunningParametersDBInstanceParameter {
	s.ParameterName = &v
	return s
}

func (s *DescribeParametersResponseBodyRunningParametersDBInstanceParameter) SetParameterValue(v string) *DescribeParametersResponseBodyRunningParametersDBInstanceParameter {
	s.ParameterValue = &v
	return s
}

func (s *DescribeParametersResponseBodyRunningParametersDBInstanceParameter) SetParameterValueRange(v string) *DescribeParametersResponseBodyRunningParametersDBInstanceParameter {
	s.ParameterValueRange = &v
	return s
}

func (s *DescribeParametersResponseBodyRunningParametersDBInstanceParameter) Validate() error {
	return dara.Validate(s)
}
