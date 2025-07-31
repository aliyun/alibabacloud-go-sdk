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
	SetRequestId(v string) *DescribeParametersResponseBody
	GetRequestId() *string
	SetRunningParameters(v *DescribeParametersResponseBodyRunningParameters) *DescribeParametersResponseBody
	GetRunningParameters() *DescribeParametersResponseBodyRunningParameters
}

type DescribeParametersResponseBody struct {
	// The parameter settings in the configuration template.
	ConfigParameters *DescribeParametersResponseBodyConfigParameters `json:"ConfigParameters,omitempty" xml:"ConfigParameters,omitempty" type:"Struct"`
	// The database engine of the instance. Default value: **mongodb**.
	//
	// example:
	//
	// mongodb
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The database engine version of the instance.
	//
	// example:
	//
	// 4.0
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 3ADD0C7D-2D2A-4F15-88FF-E7AC9B9FDCC8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The settings of the parameters that have taken effect.
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

func (s *DescribeParametersResponseBody) SetRequestId(v string) *DescribeParametersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeParametersResponseBody) SetRunningParameters(v *DescribeParametersResponseBodyRunningParameters) *DescribeParametersResponseBody {
	s.RunningParameters = v
	return s
}

func (s *DescribeParametersResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeParametersResponseBodyConfigParameters struct {
	Parameter []*DescribeParametersResponseBodyConfigParametersParameter `json:"Parameter,omitempty" xml:"Parameter,omitempty" type:"Repeated"`
}

func (s DescribeParametersResponseBodyConfigParameters) String() string {
	return dara.Prettify(s)
}

func (s DescribeParametersResponseBodyConfigParameters) GoString() string {
	return s.String()
}

func (s *DescribeParametersResponseBodyConfigParameters) GetParameter() []*DescribeParametersResponseBodyConfigParametersParameter {
	return s.Parameter
}

func (s *DescribeParametersResponseBodyConfigParameters) SetParameter(v []*DescribeParametersResponseBodyConfigParametersParameter) *DescribeParametersResponseBodyConfigParameters {
	s.Parameter = v
	return s
}

func (s *DescribeParametersResponseBodyConfigParameters) Validate() error {
	return dara.Validate(s)
}

type DescribeParametersResponseBodyConfigParametersParameter struct {
	// The valid values of the parameter.
	//
	// example:
	//
	// [0-65536]
	CheckingCode *string `json:"CheckingCode,omitempty" xml:"CheckingCode,omitempty"`
	// Indicates whether a restart is required for parameter modifications to take effect. Valid values:
	//
	// 	- **false**: A restart is not required. Modifications take effect immediately.
	//
	// 	- **true**: A restart is required for parameter modifications to take effect.
	//
	// example:
	//
	// true
	ForceRestart *bool `json:"ForceRestart,omitempty" xml:"ForceRestart,omitempty"`
	// Indicates whether the parameter value can be modified. Valid values:
	//
	// 	- **false**: The parameter value cannot be modified.
	//
	// 	- **true**: The parameter value can be modified.
	//
	// example:
	//
	// true
	ModifiableStatus *bool `json:"ModifiableStatus,omitempty" xml:"ModifiableStatus,omitempty"`
	// The description of the parameter.
	//
	// example:
	//
	// The threshold in milliseconds at which the database profiler considers a query slow, default is 100.
	ParameterDescription *string `json:"ParameterDescription,omitempty" xml:"ParameterDescription,omitempty"`
	// The name of the parameter.
	//
	// example:
	//
	// operationProfiling.slowOpThresholdMs
	ParameterName *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
	// The value of the parameter.
	//
	// example:
	//
	// 200
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s DescribeParametersResponseBodyConfigParametersParameter) String() string {
	return dara.Prettify(s)
}

func (s DescribeParametersResponseBodyConfigParametersParameter) GoString() string {
	return s.String()
}

func (s *DescribeParametersResponseBodyConfigParametersParameter) GetCheckingCode() *string {
	return s.CheckingCode
}

func (s *DescribeParametersResponseBodyConfigParametersParameter) GetForceRestart() *bool {
	return s.ForceRestart
}

func (s *DescribeParametersResponseBodyConfigParametersParameter) GetModifiableStatus() *bool {
	return s.ModifiableStatus
}

func (s *DescribeParametersResponseBodyConfigParametersParameter) GetParameterDescription() *string {
	return s.ParameterDescription
}

func (s *DescribeParametersResponseBodyConfigParametersParameter) GetParameterName() *string {
	return s.ParameterName
}

func (s *DescribeParametersResponseBodyConfigParametersParameter) GetParameterValue() *string {
	return s.ParameterValue
}

func (s *DescribeParametersResponseBodyConfigParametersParameter) SetCheckingCode(v string) *DescribeParametersResponseBodyConfigParametersParameter {
	s.CheckingCode = &v
	return s
}

func (s *DescribeParametersResponseBodyConfigParametersParameter) SetForceRestart(v bool) *DescribeParametersResponseBodyConfigParametersParameter {
	s.ForceRestart = &v
	return s
}

func (s *DescribeParametersResponseBodyConfigParametersParameter) SetModifiableStatus(v bool) *DescribeParametersResponseBodyConfigParametersParameter {
	s.ModifiableStatus = &v
	return s
}

func (s *DescribeParametersResponseBodyConfigParametersParameter) SetParameterDescription(v string) *DescribeParametersResponseBodyConfigParametersParameter {
	s.ParameterDescription = &v
	return s
}

func (s *DescribeParametersResponseBodyConfigParametersParameter) SetParameterName(v string) *DescribeParametersResponseBodyConfigParametersParameter {
	s.ParameterName = &v
	return s
}

func (s *DescribeParametersResponseBodyConfigParametersParameter) SetParameterValue(v string) *DescribeParametersResponseBodyConfigParametersParameter {
	s.ParameterValue = &v
	return s
}

func (s *DescribeParametersResponseBodyConfigParametersParameter) Validate() error {
	return dara.Validate(s)
}

type DescribeParametersResponseBodyRunningParameters struct {
	Parameter []*DescribeParametersResponseBodyRunningParametersParameter `json:"Parameter,omitempty" xml:"Parameter,omitempty" type:"Repeated"`
}

func (s DescribeParametersResponseBodyRunningParameters) String() string {
	return dara.Prettify(s)
}

func (s DescribeParametersResponseBodyRunningParameters) GoString() string {
	return s.String()
}

func (s *DescribeParametersResponseBodyRunningParameters) GetParameter() []*DescribeParametersResponseBodyRunningParametersParameter {
	return s.Parameter
}

func (s *DescribeParametersResponseBodyRunningParameters) SetParameter(v []*DescribeParametersResponseBodyRunningParametersParameter) *DescribeParametersResponseBodyRunningParameters {
	s.Parameter = v
	return s
}

func (s *DescribeParametersResponseBodyRunningParameters) Validate() error {
	return dara.Validate(s)
}

type DescribeParametersResponseBodyRunningParametersParameter struct {
	// 实例的角色类型，取值说明：
	//
	// - **db**：shard角色。
	//
	// - **cs**：config server角色。
	//
	// - **mongos**：mongos角色。
	//
	// example:
	//
	// db
	CharacterType *string `json:"CharacterType,omitempty" xml:"CharacterType,omitempty"`
	// The valid values of the parameter.
	//
	// example:
	//
	// [33554432-268435456]
	CheckingCode *string `json:"CheckingCode,omitempty" xml:"CheckingCode,omitempty"`
	// Indicates whether a restart is required for parameter modifications to take effect. Valid values:
	//
	// 	- **false**: A restart is not required. Modifications take effect immediately.
	//
	// 	- **true**: A restart is required for parameter modifications to take effect.
	//
	// example:
	//
	// false
	ForceRestart *string `json:"ForceRestart,omitempty" xml:"ForceRestart,omitempty"`
	// Indicates whether the parameter value can be modified. Valid values:
	//
	// 	- **false**: The parameter value cannot be modified.
	//
	// 	- **true**: The parameter value can be modified.
	//
	// example:
	//
	// true
	ModifiableStatus *string `json:"ModifiableStatus,omitempty" xml:"ModifiableStatus,omitempty"`
	// The description of the parameter.
	//
	// example:
	//
	// The maximum memory bytes that sort stage may use, default is 33554432(i.e. 32MB)
	ParameterDescription *string `json:"ParameterDescription,omitempty" xml:"ParameterDescription,omitempty"`
	// The name of the parameter.
	//
	// example:
	//
	// setParameter.internalQueryExecMaxBlockingSortBytes
	ParameterName *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
	// The value of the parameter.
	//
	// example:
	//
	// 33554432
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s DescribeParametersResponseBodyRunningParametersParameter) String() string {
	return dara.Prettify(s)
}

func (s DescribeParametersResponseBodyRunningParametersParameter) GoString() string {
	return s.String()
}

func (s *DescribeParametersResponseBodyRunningParametersParameter) GetCharacterType() *string {
	return s.CharacterType
}

func (s *DescribeParametersResponseBodyRunningParametersParameter) GetCheckingCode() *string {
	return s.CheckingCode
}

func (s *DescribeParametersResponseBodyRunningParametersParameter) GetForceRestart() *string {
	return s.ForceRestart
}

func (s *DescribeParametersResponseBodyRunningParametersParameter) GetModifiableStatus() *string {
	return s.ModifiableStatus
}

func (s *DescribeParametersResponseBodyRunningParametersParameter) GetParameterDescription() *string {
	return s.ParameterDescription
}

func (s *DescribeParametersResponseBodyRunningParametersParameter) GetParameterName() *string {
	return s.ParameterName
}

func (s *DescribeParametersResponseBodyRunningParametersParameter) GetParameterValue() *string {
	return s.ParameterValue
}

func (s *DescribeParametersResponseBodyRunningParametersParameter) SetCharacterType(v string) *DescribeParametersResponseBodyRunningParametersParameter {
	s.CharacterType = &v
	return s
}

func (s *DescribeParametersResponseBodyRunningParametersParameter) SetCheckingCode(v string) *DescribeParametersResponseBodyRunningParametersParameter {
	s.CheckingCode = &v
	return s
}

func (s *DescribeParametersResponseBodyRunningParametersParameter) SetForceRestart(v string) *DescribeParametersResponseBodyRunningParametersParameter {
	s.ForceRestart = &v
	return s
}

func (s *DescribeParametersResponseBodyRunningParametersParameter) SetModifiableStatus(v string) *DescribeParametersResponseBodyRunningParametersParameter {
	s.ModifiableStatus = &v
	return s
}

func (s *DescribeParametersResponseBodyRunningParametersParameter) SetParameterDescription(v string) *DescribeParametersResponseBodyRunningParametersParameter {
	s.ParameterDescription = &v
	return s
}

func (s *DescribeParametersResponseBodyRunningParametersParameter) SetParameterName(v string) *DescribeParametersResponseBodyRunningParametersParameter {
	s.ParameterName = &v
	return s
}

func (s *DescribeParametersResponseBodyRunningParametersParameter) SetParameterValue(v string) *DescribeParametersResponseBodyRunningParametersParameter {
	s.ParameterValue = &v
	return s
}

func (s *DescribeParametersResponseBodyRunningParametersParameter) Validate() error {
	return dara.Validate(s)
}
