// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterParametersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeDBClusterParametersResponseBody
	GetDBClusterId() *string
	SetDBType(v string) *DescribeDBClusterParametersResponseBody
	GetDBType() *string
	SetDBVersion(v string) *DescribeDBClusterParametersResponseBody
	GetDBVersion() *string
	SetEngine(v string) *DescribeDBClusterParametersResponseBody
	GetEngine() *string
	SetParameterNumbers(v string) *DescribeDBClusterParametersResponseBody
	GetParameterNumbers() *string
	SetParameters(v *DescribeDBClusterParametersResponseBodyParameters) *DescribeDBClusterParametersResponseBody
	GetParameters() *DescribeDBClusterParametersResponseBodyParameters
	SetRequestId(v string) *DescribeDBClusterParametersResponseBody
	GetRequestId() *string
	SetRunningParameters(v *DescribeDBClusterParametersResponseBodyRunningParameters) *DescribeDBClusterParametersResponseBody
	GetRunningParameters() *DescribeDBClusterParametersResponseBodyRunningParameters
}

type DescribeDBClusterParametersResponseBody struct {
	// The ID of the cluster.
	//
	// example:
	//
	// pc-bp1s826a1up******
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The database engine that the clusters runs. Valid values:
	//
	// 	- **MySQL**
	//
	// 	- **PostgreSQL**
	//
	// 	- **Oracle**
	//
	// example:
	//
	// MySQL
	DBType *string `json:"DBType,omitempty" xml:"DBType,omitempty"`
	// The version of the database engine.
	//
	// - Valid values for the MySQL database engine:
	//
	//   - **5.6**
	//
	//   - **5.7**
	//
	//   - **8.0**
	//
	// - Valid value for the PostgreSQL database engine:
	//
	//   - **11**
	//
	//   - **14**
	//
	// - Valid value for the Oracle database engine:  **11**
	//
	// example:
	//
	// 5.6
	DBVersion *string `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	// The cluster engine.
	//
	// example:
	//
	// POLARDB
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The number of parameters.
	//
	// example:
	//
	// 1
	ParameterNumbers *string `json:"ParameterNumbers,omitempty" xml:"ParameterNumbers,omitempty"`
	// A comparison of parameters between the source RDS instance and the destination PolarDB cluster.
	Parameters *DescribeDBClusterParametersResponseBodyParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// EBEAA83D-1734-42E3-85E3-E25F6E******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The parameters of the PolarDB cluster.
	RunningParameters *DescribeDBClusterParametersResponseBodyRunningParameters `json:"RunningParameters,omitempty" xml:"RunningParameters,omitempty" type:"Struct"`
}

func (s DescribeDBClusterParametersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterParametersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterParametersResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeDBClusterParametersResponseBody) GetDBType() *string {
	return s.DBType
}

func (s *DescribeDBClusterParametersResponseBody) GetDBVersion() *string {
	return s.DBVersion
}

func (s *DescribeDBClusterParametersResponseBody) GetEngine() *string {
	return s.Engine
}

func (s *DescribeDBClusterParametersResponseBody) GetParameterNumbers() *string {
	return s.ParameterNumbers
}

func (s *DescribeDBClusterParametersResponseBody) GetParameters() *DescribeDBClusterParametersResponseBodyParameters {
	return s.Parameters
}

func (s *DescribeDBClusterParametersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBClusterParametersResponseBody) GetRunningParameters() *DescribeDBClusterParametersResponseBodyRunningParameters {
	return s.RunningParameters
}

func (s *DescribeDBClusterParametersResponseBody) SetDBClusterId(v string) *DescribeDBClusterParametersResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterParametersResponseBody) SetDBType(v string) *DescribeDBClusterParametersResponseBody {
	s.DBType = &v
	return s
}

func (s *DescribeDBClusterParametersResponseBody) SetDBVersion(v string) *DescribeDBClusterParametersResponseBody {
	s.DBVersion = &v
	return s
}

func (s *DescribeDBClusterParametersResponseBody) SetEngine(v string) *DescribeDBClusterParametersResponseBody {
	s.Engine = &v
	return s
}

func (s *DescribeDBClusterParametersResponseBody) SetParameterNumbers(v string) *DescribeDBClusterParametersResponseBody {
	s.ParameterNumbers = &v
	return s
}

func (s *DescribeDBClusterParametersResponseBody) SetParameters(v *DescribeDBClusterParametersResponseBodyParameters) *DescribeDBClusterParametersResponseBody {
	s.Parameters = v
	return s
}

func (s *DescribeDBClusterParametersResponseBody) SetRequestId(v string) *DescribeDBClusterParametersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBClusterParametersResponseBody) SetRunningParameters(v *DescribeDBClusterParametersResponseBodyRunningParameters) *DescribeDBClusterParametersResponseBody {
	s.RunningParameters = v
	return s
}

func (s *DescribeDBClusterParametersResponseBody) Validate() error {
	if s.Parameters != nil {
		if err := s.Parameters.Validate(); err != nil {
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

type DescribeDBClusterParametersResponseBodyParameters struct {
	Parameters []*DescribeDBClusterParametersResponseBodyParametersParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
}

func (s DescribeDBClusterParametersResponseBodyParameters) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterParametersResponseBodyParameters) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterParametersResponseBodyParameters) GetParameters() []*DescribeDBClusterParametersResponseBodyParametersParameters {
	return s.Parameters
}

func (s *DescribeDBClusterParametersResponseBodyParameters) SetParameters(v []*DescribeDBClusterParametersResponseBodyParametersParameters) *DescribeDBClusterParametersResponseBodyParameters {
	s.Parameters = v
	return s
}

func (s *DescribeDBClusterParametersResponseBodyParameters) Validate() error {
	if s.Parameters != nil {
		for _, item := range s.Parameters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBClusterParametersResponseBodyParametersParameters struct {
	// Indicates whether the source and current parameters have the same value.
	//
	// example:
	//
	// true
	IsEqual *string `json:"IsEqual,omitempty" xml:"IsEqual,omitempty"`
	// Indicate whether the parameter is a primary parameter of the destination cluster. Valid values:
	//
	// 	- **1**: The parameter is a primary parameter of the destination cluster.
	//
	// 	- **0**: The parameter is not a primary parameter of the destination cluster.
	//
	// example:
	//
	// 1
	IsInstancePolarDBKey *string `json:"IsInstancePolarDBKey,omitempty" xml:"IsInstancePolarDBKey,omitempty"`
	// Indicate whether the parameter is a primary parameter of the source instance. Valid values:
	//
	// 	- **1**: The parameter is a primary parameter of the source instance.
	//
	// 	- **0**: The parameter is not a primary parameter of the source instance.
	//
	// example:
	//
	// 0
	IsInstanceRdsKey *string `json:"IsInstanceRdsKey,omitempty" xml:"IsInstanceRdsKey,omitempty"`
	// Indicate whether the parameter is a primary parameter of the destination cluster. Valid values:
	//
	// 	- **1**: The parameter is a primary parameter of the destination cluster.
	//
	// 	- **0**: The parameter is not a primary parameter of the destination cluster.
	//
	// example:
	//
	// 0
	IsPolarDBKey *string `json:"IsPolarDBKey,omitempty" xml:"IsPolarDBKey,omitempty"`
	// Indicate whether the parameter is a primary parameter of the source instance. Valid values:
	//
	// 	- **1**: The parameter is a primary parameter of the source instance.
	//
	// 	- **0**: The parameter is not a primary parameter of the source instance.
	//
	// example:
	//
	// 1
	IsRdsKey *string `json:"IsRdsKey,omitempty" xml:"IsRdsKey,omitempty"`
	// The description of the parameter of the destination cluster.
	//
	// example:
	//
	// The server\\"s default character set.
	DistParameterDescription *string `json:"distParameterDescription,omitempty" xml:"distParameterDescription,omitempty"`
	// The name of the parameter of the destination cluster.
	//
	// example:
	//
	// character_set_server
	DistParameterName *string `json:"distParameterName,omitempty" xml:"distParameterName,omitempty"`
	// The valid values of the parameter of the destination cluster.
	//
	// example:
	//
	// - utf8
	//
	// - gbk
	DistParameterOptional *string `json:"distParameterOptional,omitempty" xml:"distParameterOptional,omitempty"`
	// The value of the parameter of the destination cluster.
	//
	// example:
	//
	// utf8
	DistParameterValue *string `json:"distParameterValue,omitempty" xml:"distParameterValue,omitempty"`
	// The description of the parameter of the source instance.
	//
	// example:
	//
	// The server\\"s default character set.
	RdsParameterDescription *string `json:"rdsParameterDescription,omitempty" xml:"rdsParameterDescription,omitempty"`
	// The name of the parameter of the source instance.
	//
	// example:
	//
	// character_set_server
	RdsParameterName *string `json:"rdsParameterName,omitempty" xml:"rdsParameterName,omitempty"`
	// The valid values of the parameter of the source instance.
	//
	// example:
	//
	// - utf8
	//
	// - gbk
	RdsParameterOptional *string `json:"rdsParameterOptional,omitempty" xml:"rdsParameterOptional,omitempty"`
	// The value of the parameter of the source instance.
	//
	// example:
	//
	// utf8
	RdsParameterValue *string `json:"rdsParameterValue,omitempty" xml:"rdsParameterValue,omitempty"`
}

func (s DescribeDBClusterParametersResponseBodyParametersParameters) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterParametersResponseBodyParametersParameters) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterParametersResponseBodyParametersParameters) GetIsEqual() *string {
	return s.IsEqual
}

func (s *DescribeDBClusterParametersResponseBodyParametersParameters) GetIsInstancePolarDBKey() *string {
	return s.IsInstancePolarDBKey
}

func (s *DescribeDBClusterParametersResponseBodyParametersParameters) GetIsInstanceRdsKey() *string {
	return s.IsInstanceRdsKey
}

func (s *DescribeDBClusterParametersResponseBodyParametersParameters) GetIsPolarDBKey() *string {
	return s.IsPolarDBKey
}

func (s *DescribeDBClusterParametersResponseBodyParametersParameters) GetIsRdsKey() *string {
	return s.IsRdsKey
}

func (s *DescribeDBClusterParametersResponseBodyParametersParameters) GetDistParameterDescription() *string {
	return s.DistParameterDescription
}

func (s *DescribeDBClusterParametersResponseBodyParametersParameters) GetDistParameterName() *string {
	return s.DistParameterName
}

func (s *DescribeDBClusterParametersResponseBodyParametersParameters) GetDistParameterOptional() *string {
	return s.DistParameterOptional
}

func (s *DescribeDBClusterParametersResponseBodyParametersParameters) GetDistParameterValue() *string {
	return s.DistParameterValue
}

func (s *DescribeDBClusterParametersResponseBodyParametersParameters) GetRdsParameterDescription() *string {
	return s.RdsParameterDescription
}

func (s *DescribeDBClusterParametersResponseBodyParametersParameters) GetRdsParameterName() *string {
	return s.RdsParameterName
}

func (s *DescribeDBClusterParametersResponseBodyParametersParameters) GetRdsParameterOptional() *string {
	return s.RdsParameterOptional
}

func (s *DescribeDBClusterParametersResponseBodyParametersParameters) GetRdsParameterValue() *string {
	return s.RdsParameterValue
}

func (s *DescribeDBClusterParametersResponseBodyParametersParameters) SetIsEqual(v string) *DescribeDBClusterParametersResponseBodyParametersParameters {
	s.IsEqual = &v
	return s
}

func (s *DescribeDBClusterParametersResponseBodyParametersParameters) SetIsInstancePolarDBKey(v string) *DescribeDBClusterParametersResponseBodyParametersParameters {
	s.IsInstancePolarDBKey = &v
	return s
}

func (s *DescribeDBClusterParametersResponseBodyParametersParameters) SetIsInstanceRdsKey(v string) *DescribeDBClusterParametersResponseBodyParametersParameters {
	s.IsInstanceRdsKey = &v
	return s
}

func (s *DescribeDBClusterParametersResponseBodyParametersParameters) SetIsPolarDBKey(v string) *DescribeDBClusterParametersResponseBodyParametersParameters {
	s.IsPolarDBKey = &v
	return s
}

func (s *DescribeDBClusterParametersResponseBodyParametersParameters) SetIsRdsKey(v string) *DescribeDBClusterParametersResponseBodyParametersParameters {
	s.IsRdsKey = &v
	return s
}

func (s *DescribeDBClusterParametersResponseBodyParametersParameters) SetDistParameterDescription(v string) *DescribeDBClusterParametersResponseBodyParametersParameters {
	s.DistParameterDescription = &v
	return s
}

func (s *DescribeDBClusterParametersResponseBodyParametersParameters) SetDistParameterName(v string) *DescribeDBClusterParametersResponseBodyParametersParameters {
	s.DistParameterName = &v
	return s
}

func (s *DescribeDBClusterParametersResponseBodyParametersParameters) SetDistParameterOptional(v string) *DescribeDBClusterParametersResponseBodyParametersParameters {
	s.DistParameterOptional = &v
	return s
}

func (s *DescribeDBClusterParametersResponseBodyParametersParameters) SetDistParameterValue(v string) *DescribeDBClusterParametersResponseBodyParametersParameters {
	s.DistParameterValue = &v
	return s
}

func (s *DescribeDBClusterParametersResponseBodyParametersParameters) SetRdsParameterDescription(v string) *DescribeDBClusterParametersResponseBodyParametersParameters {
	s.RdsParameterDescription = &v
	return s
}

func (s *DescribeDBClusterParametersResponseBodyParametersParameters) SetRdsParameterName(v string) *DescribeDBClusterParametersResponseBodyParametersParameters {
	s.RdsParameterName = &v
	return s
}

func (s *DescribeDBClusterParametersResponseBodyParametersParameters) SetRdsParameterOptional(v string) *DescribeDBClusterParametersResponseBodyParametersParameters {
	s.RdsParameterOptional = &v
	return s
}

func (s *DescribeDBClusterParametersResponseBodyParametersParameters) SetRdsParameterValue(v string) *DescribeDBClusterParametersResponseBodyParametersParameters {
	s.RdsParameterValue = &v
	return s
}

func (s *DescribeDBClusterParametersResponseBodyParametersParameters) Validate() error {
	return dara.Validate(s)
}

type DescribeDBClusterParametersResponseBodyRunningParameters struct {
	Parameter []*DescribeDBClusterParametersResponseBodyRunningParametersParameter `json:"Parameter,omitempty" xml:"Parameter,omitempty" type:"Repeated"`
}

func (s DescribeDBClusterParametersResponseBodyRunningParameters) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterParametersResponseBodyRunningParameters) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterParametersResponseBodyRunningParameters) GetParameter() []*DescribeDBClusterParametersResponseBodyRunningParametersParameter {
	return s.Parameter
}

func (s *DescribeDBClusterParametersResponseBodyRunningParameters) SetParameter(v []*DescribeDBClusterParametersResponseBodyRunningParametersParameter) *DescribeDBClusterParametersResponseBodyRunningParameters {
	s.Parameter = v
	return s
}

func (s *DescribeDBClusterParametersResponseBodyRunningParameters) Validate() error {
	if s.Parameter != nil {
		for _, item := range s.Parameter {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBClusterParametersResponseBodyRunningParametersParameter struct {
	// The valid values of the parameter.
	//
	// example:
	//
	// [utf8|latin1|gbk|utf8mb4]
	CheckingCode *string `json:"CheckingCode,omitempty" xml:"CheckingCode,omitempty"`
	// The data type of the parameter value. Valid values:
	//
	// 	- **INT**
	//
	// 	- **STRING**
	//
	// 	- **B**
	//
	// example:
	//
	// INT
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// The default value of the parameter.
	//
	// example:
	//
	// utf8
	DefaultParameterValue *string `json:"DefaultParameterValue,omitempty" xml:"DefaultParameterValue,omitempty"`
	// A divisor of the parameter. For a parameter of the integer or byte type, the valid values must be a multiple of Factor unless you set Factor to 0.
	//
	// example:
	//
	// 20
	Factor *string `json:"Factor,omitempty" xml:"Factor,omitempty"`
	// Indicates whether a cluster restart is required for the parameter modification to take effect. Valid values:
	//
	// 	- **false**
	//
	// 	- **true**
	//
	// example:
	//
	// true
	ForceRestart *bool `json:"ForceRestart,omitempty" xml:"ForceRestart,omitempty"`
	// Indicates whether the parameter can be modified. Valid values:
	//
	// 	- **false**
	//
	// 	- **true**
	//
	// example:
	//
	// true
	IsModifiable *bool `json:"IsModifiable,omitempty" xml:"IsModifiable,omitempty"`
	// Indicates whether the parameter is a global parameter. Valid values:
	//
	// 	- **0**: The parameter is a global parameter. The modified parameter value is synchronized to other nodes.
	//
	// 	- **1**: The parameter is not a global parameter. You can specify the nodes to which the modified parameter value can be synchronized.
	//
	// example:
	//
	// 0
	IsNodeAvailable *string `json:"IsNodeAvailable,omitempty" xml:"IsNodeAvailable,omitempty"`
	// The dependencies of the parameter.
	//
	// example:
	//
	// utf8
	ParamRelyRule *string `json:"ParamRelyRule,omitempty" xml:"ParamRelyRule,omitempty"`
	// The description of the parameter.
	//
	// example:
	//
	// The server\\"s default character set.
	ParameterDescription *string `json:"ParameterDescription,omitempty" xml:"ParameterDescription,omitempty"`
	// The name of the parameter.
	//
	// example:
	//
	// character_set_server
	ParameterName *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
	// The status of the parameter. Valid values:
	//
	// 	- **Normal**
	//
	// 	- **Modifying**
	//
	// example:
	//
	// Normal
	ParameterStatus *string `json:"ParameterStatus,omitempty" xml:"ParameterStatus,omitempty"`
	// The value of the parameter.
	//
	// example:
	//
	// utf8
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s DescribeDBClusterParametersResponseBodyRunningParametersParameter) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterParametersResponseBodyRunningParametersParameter) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterParametersResponseBodyRunningParametersParameter) GetCheckingCode() *string {
	return s.CheckingCode
}

func (s *DescribeDBClusterParametersResponseBodyRunningParametersParameter) GetDataType() *string {
	return s.DataType
}

func (s *DescribeDBClusterParametersResponseBodyRunningParametersParameter) GetDefaultParameterValue() *string {
	return s.DefaultParameterValue
}

func (s *DescribeDBClusterParametersResponseBodyRunningParametersParameter) GetFactor() *string {
	return s.Factor
}

func (s *DescribeDBClusterParametersResponseBodyRunningParametersParameter) GetForceRestart() *bool {
	return s.ForceRestart
}

func (s *DescribeDBClusterParametersResponseBodyRunningParametersParameter) GetIsModifiable() *bool {
	return s.IsModifiable
}

func (s *DescribeDBClusterParametersResponseBodyRunningParametersParameter) GetIsNodeAvailable() *string {
	return s.IsNodeAvailable
}

func (s *DescribeDBClusterParametersResponseBodyRunningParametersParameter) GetParamRelyRule() *string {
	return s.ParamRelyRule
}

func (s *DescribeDBClusterParametersResponseBodyRunningParametersParameter) GetParameterDescription() *string {
	return s.ParameterDescription
}

func (s *DescribeDBClusterParametersResponseBodyRunningParametersParameter) GetParameterName() *string {
	return s.ParameterName
}

func (s *DescribeDBClusterParametersResponseBodyRunningParametersParameter) GetParameterStatus() *string {
	return s.ParameterStatus
}

func (s *DescribeDBClusterParametersResponseBodyRunningParametersParameter) GetParameterValue() *string {
	return s.ParameterValue
}

func (s *DescribeDBClusterParametersResponseBodyRunningParametersParameter) SetCheckingCode(v string) *DescribeDBClusterParametersResponseBodyRunningParametersParameter {
	s.CheckingCode = &v
	return s
}

func (s *DescribeDBClusterParametersResponseBodyRunningParametersParameter) SetDataType(v string) *DescribeDBClusterParametersResponseBodyRunningParametersParameter {
	s.DataType = &v
	return s
}

func (s *DescribeDBClusterParametersResponseBodyRunningParametersParameter) SetDefaultParameterValue(v string) *DescribeDBClusterParametersResponseBodyRunningParametersParameter {
	s.DefaultParameterValue = &v
	return s
}

func (s *DescribeDBClusterParametersResponseBodyRunningParametersParameter) SetFactor(v string) *DescribeDBClusterParametersResponseBodyRunningParametersParameter {
	s.Factor = &v
	return s
}

func (s *DescribeDBClusterParametersResponseBodyRunningParametersParameter) SetForceRestart(v bool) *DescribeDBClusterParametersResponseBodyRunningParametersParameter {
	s.ForceRestart = &v
	return s
}

func (s *DescribeDBClusterParametersResponseBodyRunningParametersParameter) SetIsModifiable(v bool) *DescribeDBClusterParametersResponseBodyRunningParametersParameter {
	s.IsModifiable = &v
	return s
}

func (s *DescribeDBClusterParametersResponseBodyRunningParametersParameter) SetIsNodeAvailable(v string) *DescribeDBClusterParametersResponseBodyRunningParametersParameter {
	s.IsNodeAvailable = &v
	return s
}

func (s *DescribeDBClusterParametersResponseBodyRunningParametersParameter) SetParamRelyRule(v string) *DescribeDBClusterParametersResponseBodyRunningParametersParameter {
	s.ParamRelyRule = &v
	return s
}

func (s *DescribeDBClusterParametersResponseBodyRunningParametersParameter) SetParameterDescription(v string) *DescribeDBClusterParametersResponseBodyRunningParametersParameter {
	s.ParameterDescription = &v
	return s
}

func (s *DescribeDBClusterParametersResponseBodyRunningParametersParameter) SetParameterName(v string) *DescribeDBClusterParametersResponseBodyRunningParametersParameter {
	s.ParameterName = &v
	return s
}

func (s *DescribeDBClusterParametersResponseBodyRunningParametersParameter) SetParameterStatus(v string) *DescribeDBClusterParametersResponseBodyRunningParametersParameter {
	s.ParameterStatus = &v
	return s
}

func (s *DescribeDBClusterParametersResponseBodyRunningParametersParameter) SetParameterValue(v string) *DescribeDBClusterParametersResponseBodyRunningParametersParameter {
	s.ParameterValue = &v
	return s
}

func (s *DescribeDBClusterParametersResponseBodyRunningParametersParameter) Validate() error {
	return dara.Validate(s)
}
