// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSearchParametersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeSearchParametersResponseBody
	GetDBClusterId() *string
	SetDBVersion(v string) *DescribeSearchParametersResponseBody
	GetDBVersion() *string
	SetEngine(v string) *DescribeSearchParametersResponseBody
	GetEngine() *string
	SetRequestId(v string) *DescribeSearchParametersResponseBody
	GetRequestId() *string
	SetRunningParameters(v *DescribeSearchParametersResponseBodyRunningParameters) *DescribeSearchParametersResponseBody
	GetRunningParameters() *DescribeSearchParametersResponseBodyRunningParameters
}

type DescribeSearchParametersResponseBody struct {
	// The ID of the PolarDB primary cluster.
	//
	// example:
	//
	// pc-2zex82f527he****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The Search engine version. This is not the MySQL version of the primary cluster.
	//
	// example:
	//
	// 3.0
	DBVersion *string `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	// The type of the Search engine.
	//
	// example:
	//
	// polardb_search
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The request ID.
	//
	// example:
	//
	// e3a24663-eee5-4009-b635-2e9e8d052e74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The visible parameters and their modification status.
	RunningParameters *DescribeSearchParametersResponseBodyRunningParameters `json:"RunningParameters,omitempty" xml:"RunningParameters,omitempty" type:"Struct"`
}

func (s DescribeSearchParametersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSearchParametersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSearchParametersResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeSearchParametersResponseBody) GetDBVersion() *string {
	return s.DBVersion
}

func (s *DescribeSearchParametersResponseBody) GetEngine() *string {
	return s.Engine
}

func (s *DescribeSearchParametersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSearchParametersResponseBody) GetRunningParameters() *DescribeSearchParametersResponseBodyRunningParameters {
	return s.RunningParameters
}

func (s *DescribeSearchParametersResponseBody) SetDBClusterId(v string) *DescribeSearchParametersResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeSearchParametersResponseBody) SetDBVersion(v string) *DescribeSearchParametersResponseBody {
	s.DBVersion = &v
	return s
}

func (s *DescribeSearchParametersResponseBody) SetEngine(v string) *DescribeSearchParametersResponseBody {
	s.Engine = &v
	return s
}

func (s *DescribeSearchParametersResponseBody) SetRequestId(v string) *DescribeSearchParametersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSearchParametersResponseBody) SetRunningParameters(v *DescribeSearchParametersResponseBodyRunningParameters) *DescribeSearchParametersResponseBody {
	s.RunningParameters = v
	return s
}

func (s *DescribeSearchParametersResponseBody) Validate() error {
	if s.RunningParameters != nil {
		if err := s.RunningParameters.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSearchParametersResponseBodyRunningParameters struct {
	// The list of parameters.
	Parameter []*DescribeSearchParametersResponseBodyRunningParametersParameter `json:"Parameter,omitempty" xml:"Parameter,omitempty" type:"Repeated"`
}

func (s DescribeSearchParametersResponseBodyRunningParameters) String() string {
	return dara.Prettify(s)
}

func (s DescribeSearchParametersResponseBodyRunningParameters) GoString() string {
	return s.String()
}

func (s *DescribeSearchParametersResponseBodyRunningParameters) GetParameter() []*DescribeSearchParametersResponseBodyRunningParametersParameter {
	return s.Parameter
}

func (s *DescribeSearchParametersResponseBodyRunningParameters) SetParameter(v []*DescribeSearchParametersResponseBodyRunningParametersParameter) *DescribeSearchParametersResponseBodyRunningParameters {
	s.Parameter = v
	return s
}

func (s *DescribeSearchParametersResponseBodyRunningParameters) Validate() error {
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

type DescribeSearchParametersResponseBodyRunningParametersParameter struct {
	// The value validation rule of the template. The rule can be a numeric range, an enumeration, or a regular expression, and must be interpreted in combination with the data type.
	//
	// example:
	//
	// [true\\|false]
	CheckingCode *string `json:"CheckingCode,omitempty" xml:"CheckingCode,omitempty"`
	// The data type of the parameter value. Valid values: INT, DOUBLE, B/BYTE, MEMORY, PERCENT, TIME, BOOLEAN, and STRING. The MEMORY type supports percentages or capacities with units.
	//
	// example:
	//
	// BOOLEAN
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// The default value of the current active template. This value is independent of the value modified by the user.
	//
	// example:
	//
	// false
	DefaultParameterValue *string `json:"DefaultParameterValue,omitempty" xml:"DefaultParameterValue,omitempty"`
	// The divisibility factor. If the value is not zero, the numeric value parsed based on the template type must be a multiple of this factor.
	//
	// example:
	//
	// 0
	Factor *string `json:"Factor,omitempty" xml:"Factor,omitempty"`
	// Specifies whether a restart is required for the parameter to take effect. The value is true for static parameters and false for dynamic parameters.
	//
	// example:
	//
	// false
	ForceRestart *bool `json:"ForceRestart,omitempty" xml:"ForceRestart,omitempty"`
	// Specifies whether the parameter is modifiable.
	//
	// example:
	//
	// false
	IsModifiable *bool `json:"IsModifiable,omitempty" xml:"IsModifiable,omitempty"`
	// The parameter description.
	//
	// example:
	//
	// Parameter description
	ParameterDescription *string `json:"ParameterDescription,omitempty" xml:"ParameterDescription,omitempty"`
	// The name of the parameter.
	//
	// example:
	//
	// prometheus.indices
	ParameterName *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
	// The modification status of the parameter. A value of Normal indicates that no modifications are pending. A value of Modifying indicates that the modification task is not completed.
	//
	// example:
	//
	// false
	ParameterStatus *string `json:"ParameterStatus,omitempty" xml:"ParameterStatus,omitempty"`
	// The configuration value on the control plane. If a modification is in progress, the value to be applied is displayed. This value is not read from the kernel in real time.
	//
	// example:
	//
	// false
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s DescribeSearchParametersResponseBodyRunningParametersParameter) String() string {
	return dara.Prettify(s)
}

func (s DescribeSearchParametersResponseBodyRunningParametersParameter) GoString() string {
	return s.String()
}

func (s *DescribeSearchParametersResponseBodyRunningParametersParameter) GetCheckingCode() *string {
	return s.CheckingCode
}

func (s *DescribeSearchParametersResponseBodyRunningParametersParameter) GetDataType() *string {
	return s.DataType
}

func (s *DescribeSearchParametersResponseBodyRunningParametersParameter) GetDefaultParameterValue() *string {
	return s.DefaultParameterValue
}

func (s *DescribeSearchParametersResponseBodyRunningParametersParameter) GetFactor() *string {
	return s.Factor
}

func (s *DescribeSearchParametersResponseBodyRunningParametersParameter) GetForceRestart() *bool {
	return s.ForceRestart
}

func (s *DescribeSearchParametersResponseBodyRunningParametersParameter) GetIsModifiable() *bool {
	return s.IsModifiable
}

func (s *DescribeSearchParametersResponseBodyRunningParametersParameter) GetParameterDescription() *string {
	return s.ParameterDescription
}

func (s *DescribeSearchParametersResponseBodyRunningParametersParameter) GetParameterName() *string {
	return s.ParameterName
}

func (s *DescribeSearchParametersResponseBodyRunningParametersParameter) GetParameterStatus() *string {
	return s.ParameterStatus
}

func (s *DescribeSearchParametersResponseBodyRunningParametersParameter) GetParameterValue() *string {
	return s.ParameterValue
}

func (s *DescribeSearchParametersResponseBodyRunningParametersParameter) SetCheckingCode(v string) *DescribeSearchParametersResponseBodyRunningParametersParameter {
	s.CheckingCode = &v
	return s
}

func (s *DescribeSearchParametersResponseBodyRunningParametersParameter) SetDataType(v string) *DescribeSearchParametersResponseBodyRunningParametersParameter {
	s.DataType = &v
	return s
}

func (s *DescribeSearchParametersResponseBodyRunningParametersParameter) SetDefaultParameterValue(v string) *DescribeSearchParametersResponseBodyRunningParametersParameter {
	s.DefaultParameterValue = &v
	return s
}

func (s *DescribeSearchParametersResponseBodyRunningParametersParameter) SetFactor(v string) *DescribeSearchParametersResponseBodyRunningParametersParameter {
	s.Factor = &v
	return s
}

func (s *DescribeSearchParametersResponseBodyRunningParametersParameter) SetForceRestart(v bool) *DescribeSearchParametersResponseBodyRunningParametersParameter {
	s.ForceRestart = &v
	return s
}

func (s *DescribeSearchParametersResponseBodyRunningParametersParameter) SetIsModifiable(v bool) *DescribeSearchParametersResponseBodyRunningParametersParameter {
	s.IsModifiable = &v
	return s
}

func (s *DescribeSearchParametersResponseBodyRunningParametersParameter) SetParameterDescription(v string) *DescribeSearchParametersResponseBodyRunningParametersParameter {
	s.ParameterDescription = &v
	return s
}

func (s *DescribeSearchParametersResponseBodyRunningParametersParameter) SetParameterName(v string) *DescribeSearchParametersResponseBodyRunningParametersParameter {
	s.ParameterName = &v
	return s
}

func (s *DescribeSearchParametersResponseBodyRunningParametersParameter) SetParameterStatus(v string) *DescribeSearchParametersResponseBodyRunningParametersParameter {
	s.ParameterStatus = &v
	return s
}

func (s *DescribeSearchParametersResponseBodyRunningParametersParameter) SetParameterValue(v string) *DescribeSearchParametersResponseBodyRunningParametersParameter {
	s.ParameterValue = &v
	return s
}

func (s *DescribeSearchParametersResponseBodyRunningParametersParameter) Validate() error {
	return dara.Validate(s)
}
