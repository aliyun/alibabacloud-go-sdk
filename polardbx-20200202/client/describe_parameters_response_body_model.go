// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeParametersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeParametersResponseBodyData) *DescribeParametersResponseBody
	GetData() *DescribeParametersResponseBodyData
	SetRequestId(v string) *DescribeParametersResponseBody
	GetRequestId() *string
}

type DescribeParametersResponseBody struct {
	// The parameter details.
	Data *DescribeParametersResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 6BA32080EEA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeParametersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeParametersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeParametersResponseBody) GetData() *DescribeParametersResponseBodyData {
	return s.Data
}

func (s *DescribeParametersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeParametersResponseBody) SetData(v *DescribeParametersResponseBodyData) *DescribeParametersResponseBody {
	s.Data = v
	return s
}

func (s *DescribeParametersResponseBody) SetRequestId(v string) *DescribeParametersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeParametersResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeParametersResponseBodyData struct {
	// The list of original parameters.
	ConfigParameters []*DescribeParametersResponseBodyDataConfigParameters `json:"ConfigParameters,omitempty" xml:"ConfigParameters,omitempty" type:"Repeated"`
	DBInstanceId     *string                                               `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The engine. Default value: polarx.
	//
	// example:
	//
	// polarx
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The DPI engine version. Default value: 2.0.
	//
	// example:
	//
	// 2.0
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The list of running parameter details.
	RunningParameters []*DescribeParametersResponseBodyDataRunningParameters `json:"RunningParameters,omitempty" xml:"RunningParameters,omitempty" type:"Repeated"`
}

func (s DescribeParametersResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeParametersResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeParametersResponseBodyData) GetConfigParameters() []*DescribeParametersResponseBodyDataConfigParameters {
	return s.ConfigParameters
}

func (s *DescribeParametersResponseBodyData) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeParametersResponseBodyData) GetEngine() *string {
	return s.Engine
}

func (s *DescribeParametersResponseBodyData) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *DescribeParametersResponseBodyData) GetRunningParameters() []*DescribeParametersResponseBodyDataRunningParameters {
	return s.RunningParameters
}

func (s *DescribeParametersResponseBodyData) SetConfigParameters(v []*DescribeParametersResponseBodyDataConfigParameters) *DescribeParametersResponseBodyData {
	s.ConfigParameters = v
	return s
}

func (s *DescribeParametersResponseBodyData) SetDBInstanceId(v string) *DescribeParametersResponseBodyData {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeParametersResponseBodyData) SetEngine(v string) *DescribeParametersResponseBodyData {
	s.Engine = &v
	return s
}

func (s *DescribeParametersResponseBodyData) SetEngineVersion(v string) *DescribeParametersResponseBodyData {
	s.EngineVersion = &v
	return s
}

func (s *DescribeParametersResponseBodyData) SetRunningParameters(v []*DescribeParametersResponseBodyDataRunningParameters) *DescribeParametersResponseBodyData {
	s.RunningParameters = v
	return s
}

func (s *DescribeParametersResponseBodyData) Validate() error {
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

type DescribeParametersResponseBodyDataConfigParameters struct {
	// The parameter description.
	//
	// example:
	//
	// 存储节点私有协议开关(-1关闭，0开启)
	ParameterDescription *string `json:"ParameterDescription,omitempty" xml:"ParameterDescription,omitempty"`
	// The parameter name.
	//
	// example:
	//
	// CONN_POOL_XPROTO_STORAGE_DB_PORT
	ParameterName *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
	// The parameter value.
	//
	// example:
	//
	// -1
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s DescribeParametersResponseBodyDataConfigParameters) String() string {
	return dara.Prettify(s)
}

func (s DescribeParametersResponseBodyDataConfigParameters) GoString() string {
	return s.String()
}

func (s *DescribeParametersResponseBodyDataConfigParameters) GetParameterDescription() *string {
	return s.ParameterDescription
}

func (s *DescribeParametersResponseBodyDataConfigParameters) GetParameterName() *string {
	return s.ParameterName
}

func (s *DescribeParametersResponseBodyDataConfigParameters) GetParameterValue() *string {
	return s.ParameterValue
}

func (s *DescribeParametersResponseBodyDataConfigParameters) SetParameterDescription(v string) *DescribeParametersResponseBodyDataConfigParameters {
	s.ParameterDescription = &v
	return s
}

func (s *DescribeParametersResponseBodyDataConfigParameters) SetParameterName(v string) *DescribeParametersResponseBodyDataConfigParameters {
	s.ParameterName = &v
	return s
}

func (s *DescribeParametersResponseBodyDataConfigParameters) SetParameterValue(v string) *DescribeParametersResponseBodyDataConfigParameters {
	s.ParameterValue = &v
	return s
}

func (s *DescribeParametersResponseBodyDataConfigParameters) Validate() error {
	return dara.Validate(s)
}

type DescribeParametersResponseBodyDataRunningParameters struct {
	// The parameter description.
	//
	// example:
	//
	// 存储节点私有协议开关(-1关闭，0开启)
	ParameterDescription *string `json:"ParameterDescription,omitempty" xml:"ParameterDescription,omitempty"`
	// The parameter name.
	//
	// example:
	//
	// CONN_POOL_XPROTO_STORAGE_DB_PORT
	ParameterName *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
	// The parameter value.
	//
	// example:
	//
	// -1
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s DescribeParametersResponseBodyDataRunningParameters) String() string {
	return dara.Prettify(s)
}

func (s DescribeParametersResponseBodyDataRunningParameters) GoString() string {
	return s.String()
}

func (s *DescribeParametersResponseBodyDataRunningParameters) GetParameterDescription() *string {
	return s.ParameterDescription
}

func (s *DescribeParametersResponseBodyDataRunningParameters) GetParameterName() *string {
	return s.ParameterName
}

func (s *DescribeParametersResponseBodyDataRunningParameters) GetParameterValue() *string {
	return s.ParameterValue
}

func (s *DescribeParametersResponseBodyDataRunningParameters) SetParameterDescription(v string) *DescribeParametersResponseBodyDataRunningParameters {
	s.ParameterDescription = &v
	return s
}

func (s *DescribeParametersResponseBodyDataRunningParameters) SetParameterName(v string) *DescribeParametersResponseBodyDataRunningParameters {
	s.ParameterName = &v
	return s
}

func (s *DescribeParametersResponseBodyDataRunningParameters) SetParameterValue(v string) *DescribeParametersResponseBodyDataRunningParameters {
	s.ParameterValue = &v
	return s
}

func (s *DescribeParametersResponseBodyDataRunningParameters) Validate() error {
	return dara.Validate(s)
}
