// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDatabaseInstanceParametersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigParameters(v []*DescribeDatabaseInstanceParametersResponseBodyConfigParameters) *DescribeDatabaseInstanceParametersResponseBody
	GetConfigParameters() []*DescribeDatabaseInstanceParametersResponseBodyConfigParameters
	SetEngine(v string) *DescribeDatabaseInstanceParametersResponseBody
	GetEngine() *string
	SetEngineVersion(v string) *DescribeDatabaseInstanceParametersResponseBody
	GetEngineVersion() *string
	SetRequestId(v string) *DescribeDatabaseInstanceParametersResponseBody
	GetRequestId() *string
	SetRunningParameters(v []*DescribeDatabaseInstanceParametersResponseBodyRunningParameters) *DescribeDatabaseInstanceParametersResponseBody
	GetRunningParameters() []*DescribeDatabaseInstanceParametersResponseBodyRunningParameters
}

type DescribeDatabaseInstanceParametersResponseBody struct {
	// A list of configuration parameters being synchronized. After you modify and submit parameters, the instance starts to synchronize them. Once synchronization is complete, these parameters are removed from this list.
	ConfigParameters []*DescribeDatabaseInstanceParametersResponseBodyConfigParameters `json:"ConfigParameters,omitempty" xml:"ConfigParameters,omitempty" type:"Repeated"`
	// The database engine. Currently, only `MySQL` is supported.
	//
	// example:
	//
	// MySQL
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The engine version. Valid values:
	//
	// - `5.7`: MySQL 5.7.
	//
	// - `8.0`: MySQL 8.0.
	//
	// example:
	//
	// 5.5
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 30637AD6-D977-4833-A54C-CC89483E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// A list of running parameters.
	RunningParameters []*DescribeDatabaseInstanceParametersResponseBodyRunningParameters `json:"RunningParameters,omitempty" xml:"RunningParameters,omitempty" type:"Repeated"`
}

func (s DescribeDatabaseInstanceParametersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDatabaseInstanceParametersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDatabaseInstanceParametersResponseBody) GetConfigParameters() []*DescribeDatabaseInstanceParametersResponseBodyConfigParameters {
	return s.ConfigParameters
}

func (s *DescribeDatabaseInstanceParametersResponseBody) GetEngine() *string {
	return s.Engine
}

func (s *DescribeDatabaseInstanceParametersResponseBody) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *DescribeDatabaseInstanceParametersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDatabaseInstanceParametersResponseBody) GetRunningParameters() []*DescribeDatabaseInstanceParametersResponseBodyRunningParameters {
	return s.RunningParameters
}

func (s *DescribeDatabaseInstanceParametersResponseBody) SetConfigParameters(v []*DescribeDatabaseInstanceParametersResponseBodyConfigParameters) *DescribeDatabaseInstanceParametersResponseBody {
	s.ConfigParameters = v
	return s
}

func (s *DescribeDatabaseInstanceParametersResponseBody) SetEngine(v string) *DescribeDatabaseInstanceParametersResponseBody {
	s.Engine = &v
	return s
}

func (s *DescribeDatabaseInstanceParametersResponseBody) SetEngineVersion(v string) *DescribeDatabaseInstanceParametersResponseBody {
	s.EngineVersion = &v
	return s
}

func (s *DescribeDatabaseInstanceParametersResponseBody) SetRequestId(v string) *DescribeDatabaseInstanceParametersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDatabaseInstanceParametersResponseBody) SetRunningParameters(v []*DescribeDatabaseInstanceParametersResponseBodyRunningParameters) *DescribeDatabaseInstanceParametersResponseBody {
	s.RunningParameters = v
	return s
}

func (s *DescribeDatabaseInstanceParametersResponseBody) Validate() error {
	if s.ConfigParameters != nil {
		for _, item := range s.ConfigParameters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RunningParameters != nil {
		for _, item := range s.RunningParameters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDatabaseInstanceParametersResponseBodyConfigParameters struct {
	// The value range of the parameter.
	//
	// > The value range varies by parameter.
	//
	// example:
	//
	// [1-65535]
	CheckingCode *string `json:"CheckingCode,omitempty" xml:"CheckingCode,omitempty"`
	// Indicates whether you can modify the parameter value. Valid values:
	//
	// - `true`: You can modify the parameter value.
	//
	// - `false`: You cannot modify the parameter value.
	//
	// example:
	//
	// true
	ForceModify *string `json:"ForceModify,omitempty" xml:"ForceModify,omitempty"`
	// Indicates if the database must be restarted for a parameter modification to take effect. Valid values:
	//
	// - `true`: A restart is required. To apply changes for a parameter that requires a restart, you must set the `ForceRestart` parameter to `true`.
	//
	// - `false`: A restart is not required.
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	ForceRestart *string `json:"ForceRestart,omitempty" xml:"ForceRestart,omitempty"`
	// The description of the parameter.
	//
	// example:
	//
	// Auto-increment columns are incremented by this
	ParameterDescription *string `json:"ParameterDescription,omitempty" xml:"ParameterDescription,omitempty"`
	// The name of the parameter.
	//
	// example:
	//
	// auto_increment_increment
	ParameterName *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
	// The value of the parameter.
	//
	// > The value varies by parameter.
	//
	// example:
	//
	// 1
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s DescribeDatabaseInstanceParametersResponseBodyConfigParameters) String() string {
	return dara.Prettify(s)
}

func (s DescribeDatabaseInstanceParametersResponseBodyConfigParameters) GoString() string {
	return s.String()
}

func (s *DescribeDatabaseInstanceParametersResponseBodyConfigParameters) GetCheckingCode() *string {
	return s.CheckingCode
}

func (s *DescribeDatabaseInstanceParametersResponseBodyConfigParameters) GetForceModify() *string {
	return s.ForceModify
}

func (s *DescribeDatabaseInstanceParametersResponseBodyConfigParameters) GetForceRestart() *string {
	return s.ForceRestart
}

func (s *DescribeDatabaseInstanceParametersResponseBodyConfigParameters) GetParameterDescription() *string {
	return s.ParameterDescription
}

func (s *DescribeDatabaseInstanceParametersResponseBodyConfigParameters) GetParameterName() *string {
	return s.ParameterName
}

func (s *DescribeDatabaseInstanceParametersResponseBodyConfigParameters) GetParameterValue() *string {
	return s.ParameterValue
}

func (s *DescribeDatabaseInstanceParametersResponseBodyConfigParameters) SetCheckingCode(v string) *DescribeDatabaseInstanceParametersResponseBodyConfigParameters {
	s.CheckingCode = &v
	return s
}

func (s *DescribeDatabaseInstanceParametersResponseBodyConfigParameters) SetForceModify(v string) *DescribeDatabaseInstanceParametersResponseBodyConfigParameters {
	s.ForceModify = &v
	return s
}

func (s *DescribeDatabaseInstanceParametersResponseBodyConfigParameters) SetForceRestart(v string) *DescribeDatabaseInstanceParametersResponseBodyConfigParameters {
	s.ForceRestart = &v
	return s
}

func (s *DescribeDatabaseInstanceParametersResponseBodyConfigParameters) SetParameterDescription(v string) *DescribeDatabaseInstanceParametersResponseBodyConfigParameters {
	s.ParameterDescription = &v
	return s
}

func (s *DescribeDatabaseInstanceParametersResponseBodyConfigParameters) SetParameterName(v string) *DescribeDatabaseInstanceParametersResponseBodyConfigParameters {
	s.ParameterName = &v
	return s
}

func (s *DescribeDatabaseInstanceParametersResponseBodyConfigParameters) SetParameterValue(v string) *DescribeDatabaseInstanceParametersResponseBodyConfigParameters {
	s.ParameterValue = &v
	return s
}

func (s *DescribeDatabaseInstanceParametersResponseBodyConfigParameters) Validate() error {
	return dara.Validate(s)
}

type DescribeDatabaseInstanceParametersResponseBodyRunningParameters struct {
	// The value range of the parameter.
	//
	// > The value range varies by parameter.
	//
	// example:
	//
	// [ON|OFF]
	CheckingCode *string `json:"CheckingCode,omitempty" xml:"CheckingCode,omitempty"`
	// Indicates whether you can modify the parameter value. Valid values:
	//
	// - `true`: You can modify the parameter value.
	//
	// - `false`: You cannot modify the parameter value.
	//
	// example:
	//
	// true
	ForceModify *string `json:"ForceModify,omitempty" xml:"ForceModify,omitempty"`
	// Indicates if the database must be restarted for a parameter modification to take effect. Valid values:
	//
	// - `true`: A restart is required. To apply changes for a parameter that requires a restart, you must set the `ForceRestart` parameter to `true`.
	//
	// - `false`: A restart is not required.
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	ForceRestart *string `json:"ForceRestart,omitempty" xml:"ForceRestart,omitempty"`
	// The description of the running parameter.
	//
	// example:
	//
	// When this variable has a value of 1 (the default), the server automatically grants the EXECUTE and ALTER ROUTINE privileges to the creator of a stored routine, if the user cannot already execute and alter or drop the routine. (The ALTER ROUTINE privilege is required to drop the routine.) The server also automatically drops those privileges from the creator when the routine is dropped. If automatic_sp_privileges is 0, the server does not automatically add or drop these privileges.
	ParameterDescription *string `json:"ParameterDescription,omitempty" xml:"ParameterDescription,omitempty"`
	// The name of the parameter.
	//
	// example:
	//
	// autocommit
	ParameterName *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
	// The value of the running parameter.
	//
	// example:
	//
	// ON
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s DescribeDatabaseInstanceParametersResponseBodyRunningParameters) String() string {
	return dara.Prettify(s)
}

func (s DescribeDatabaseInstanceParametersResponseBodyRunningParameters) GoString() string {
	return s.String()
}

func (s *DescribeDatabaseInstanceParametersResponseBodyRunningParameters) GetCheckingCode() *string {
	return s.CheckingCode
}

func (s *DescribeDatabaseInstanceParametersResponseBodyRunningParameters) GetForceModify() *string {
	return s.ForceModify
}

func (s *DescribeDatabaseInstanceParametersResponseBodyRunningParameters) GetForceRestart() *string {
	return s.ForceRestart
}

func (s *DescribeDatabaseInstanceParametersResponseBodyRunningParameters) GetParameterDescription() *string {
	return s.ParameterDescription
}

func (s *DescribeDatabaseInstanceParametersResponseBodyRunningParameters) GetParameterName() *string {
	return s.ParameterName
}

func (s *DescribeDatabaseInstanceParametersResponseBodyRunningParameters) GetParameterValue() *string {
	return s.ParameterValue
}

func (s *DescribeDatabaseInstanceParametersResponseBodyRunningParameters) SetCheckingCode(v string) *DescribeDatabaseInstanceParametersResponseBodyRunningParameters {
	s.CheckingCode = &v
	return s
}

func (s *DescribeDatabaseInstanceParametersResponseBodyRunningParameters) SetForceModify(v string) *DescribeDatabaseInstanceParametersResponseBodyRunningParameters {
	s.ForceModify = &v
	return s
}

func (s *DescribeDatabaseInstanceParametersResponseBodyRunningParameters) SetForceRestart(v string) *DescribeDatabaseInstanceParametersResponseBodyRunningParameters {
	s.ForceRestart = &v
	return s
}

func (s *DescribeDatabaseInstanceParametersResponseBodyRunningParameters) SetParameterDescription(v string) *DescribeDatabaseInstanceParametersResponseBodyRunningParameters {
	s.ParameterDescription = &v
	return s
}

func (s *DescribeDatabaseInstanceParametersResponseBodyRunningParameters) SetParameterName(v string) *DescribeDatabaseInstanceParametersResponseBodyRunningParameters {
	s.ParameterName = &v
	return s
}

func (s *DescribeDatabaseInstanceParametersResponseBodyRunningParameters) SetParameterValue(v string) *DescribeDatabaseInstanceParametersResponseBodyRunningParameters {
	s.ParameterValue = &v
	return s
}

func (s *DescribeDatabaseInstanceParametersResponseBodyRunningParameters) Validate() error {
	return dara.Validate(s)
}
