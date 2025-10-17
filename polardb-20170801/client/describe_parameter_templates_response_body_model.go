// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeParameterTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBType(v string) *DescribeParameterTemplatesResponseBody
	GetDBType() *string
	SetDBVersion(v string) *DescribeParameterTemplatesResponseBody
	GetDBVersion() *string
	SetEngine(v string) *DescribeParameterTemplatesResponseBody
	GetEngine() *string
	SetParameterCount(v string) *DescribeParameterTemplatesResponseBody
	GetParameterCount() *string
	SetParameters(v *DescribeParameterTemplatesResponseBodyParameters) *DescribeParameterTemplatesResponseBody
	GetParameters() *DescribeParameterTemplatesResponseBodyParameters
	SetRequestId(v string) *DescribeParameterTemplatesResponseBody
	GetRequestId() *string
}

type DescribeParameterTemplatesResponseBody struct {
	// The type of the database engine.
	//
	// example:
	//
	// MySQL
	DBType *string `json:"DBType,omitempty" xml:"DBType,omitempty"`
	// The version of the database engine.
	//
	// example:
	//
	// 5.7
	DBVersion *string `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	// The database engine of the cluster.
	//
	// example:
	//
	// POLARDB
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The number of parameters.
	//
	// example:
	//
	// 183
	ParameterCount *string `json:"ParameterCount,omitempty" xml:"ParameterCount,omitempty"`
	// The details of the parameters.
	Parameters *DescribeParameterTemplatesResponseBodyParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// D963934D-8605-4473-8EAC-54C719******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeParameterTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeParameterTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeParameterTemplatesResponseBody) GetDBType() *string {
	return s.DBType
}

func (s *DescribeParameterTemplatesResponseBody) GetDBVersion() *string {
	return s.DBVersion
}

func (s *DescribeParameterTemplatesResponseBody) GetEngine() *string {
	return s.Engine
}

func (s *DescribeParameterTemplatesResponseBody) GetParameterCount() *string {
	return s.ParameterCount
}

func (s *DescribeParameterTemplatesResponseBody) GetParameters() *DescribeParameterTemplatesResponseBodyParameters {
	return s.Parameters
}

func (s *DescribeParameterTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeParameterTemplatesResponseBody) SetDBType(v string) *DescribeParameterTemplatesResponseBody {
	s.DBType = &v
	return s
}

func (s *DescribeParameterTemplatesResponseBody) SetDBVersion(v string) *DescribeParameterTemplatesResponseBody {
	s.DBVersion = &v
	return s
}

func (s *DescribeParameterTemplatesResponseBody) SetEngine(v string) *DescribeParameterTemplatesResponseBody {
	s.Engine = &v
	return s
}

func (s *DescribeParameterTemplatesResponseBody) SetParameterCount(v string) *DescribeParameterTemplatesResponseBody {
	s.ParameterCount = &v
	return s
}

func (s *DescribeParameterTemplatesResponseBody) SetParameters(v *DescribeParameterTemplatesResponseBodyParameters) *DescribeParameterTemplatesResponseBody {
	s.Parameters = v
	return s
}

func (s *DescribeParameterTemplatesResponseBody) SetRequestId(v string) *DescribeParameterTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeParameterTemplatesResponseBody) Validate() error {
	if s.Parameters != nil {
		if err := s.Parameters.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeParameterTemplatesResponseBodyParameters struct {
	TemplateRecord []*DescribeParameterTemplatesResponseBodyParametersTemplateRecord `json:"TemplateRecord,omitempty" xml:"TemplateRecord,omitempty" type:"Repeated"`
}

func (s DescribeParameterTemplatesResponseBodyParameters) String() string {
	return dara.Prettify(s)
}

func (s DescribeParameterTemplatesResponseBodyParameters) GoString() string {
	return s.String()
}

func (s *DescribeParameterTemplatesResponseBodyParameters) GetTemplateRecord() []*DescribeParameterTemplatesResponseBodyParametersTemplateRecord {
	return s.TemplateRecord
}

func (s *DescribeParameterTemplatesResponseBodyParameters) SetTemplateRecord(v []*DescribeParameterTemplatesResponseBodyParametersTemplateRecord) *DescribeParameterTemplatesResponseBodyParameters {
	s.TemplateRecord = v
	return s
}

func (s *DescribeParameterTemplatesResponseBodyParameters) Validate() error {
	if s.TemplateRecord != nil {
		for _, item := range s.TemplateRecord {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeParameterTemplatesResponseBodyParametersTemplateRecord struct {
	// The valid values of the parameter.
	//
	// example:
	//
	// [ROW|STATEMENT|MIXED]
	CheckingCode *string `json:"CheckingCode,omitempty" xml:"CheckingCode,omitempty"`
	// Indicates whether the parameter setting can be modified. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	ForceModify *string `json:"ForceModify,omitempty" xml:"ForceModify,omitempty"`
	// Indicates whether a cluster restart is required to make the parameter modification take effect. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	ForceRestart *string `json:"ForceRestart,omitempty" xml:"ForceRestart,omitempty"`
	// Indicates whether the parameter is a global parameter. Valid values:
	//
	// 	- **0**: yes. The modified parameter value is synchronized to other nodes by default.
	//
	// 	- **1**: no. You can customize the nodes to which the modified parameter value can be synchronized.
	//
	// example:
	//
	// 1
	IsNodeAvailable *string `json:"IsNodeAvailable,omitempty" xml:"IsNodeAvailable,omitempty"`
	// The parameter dependencies.
	//
	// example:
	//
	// utf8
	ParamRelyRule *string `json:"ParamRelyRule,omitempty" xml:"ParamRelyRule,omitempty"`
	// The description of the parameter.
	//
	// example:
	//
	// What form of binary logging the master will use.
	ParameterDescription *string `json:"ParameterDescription,omitempty" xml:"ParameterDescription,omitempty"`
	// The name of the parameter.
	//
	// example:
	//
	// binlog_format
	ParameterName *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
	// The default value of the parameter.
	//
	// example:
	//
	// ROW
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s DescribeParameterTemplatesResponseBodyParametersTemplateRecord) String() string {
	return dara.Prettify(s)
}

func (s DescribeParameterTemplatesResponseBodyParametersTemplateRecord) GoString() string {
	return s.String()
}

func (s *DescribeParameterTemplatesResponseBodyParametersTemplateRecord) GetCheckingCode() *string {
	return s.CheckingCode
}

func (s *DescribeParameterTemplatesResponseBodyParametersTemplateRecord) GetForceModify() *string {
	return s.ForceModify
}

func (s *DescribeParameterTemplatesResponseBodyParametersTemplateRecord) GetForceRestart() *string {
	return s.ForceRestart
}

func (s *DescribeParameterTemplatesResponseBodyParametersTemplateRecord) GetIsNodeAvailable() *string {
	return s.IsNodeAvailable
}

func (s *DescribeParameterTemplatesResponseBodyParametersTemplateRecord) GetParamRelyRule() *string {
	return s.ParamRelyRule
}

func (s *DescribeParameterTemplatesResponseBodyParametersTemplateRecord) GetParameterDescription() *string {
	return s.ParameterDescription
}

func (s *DescribeParameterTemplatesResponseBodyParametersTemplateRecord) GetParameterName() *string {
	return s.ParameterName
}

func (s *DescribeParameterTemplatesResponseBodyParametersTemplateRecord) GetParameterValue() *string {
	return s.ParameterValue
}

func (s *DescribeParameterTemplatesResponseBodyParametersTemplateRecord) SetCheckingCode(v string) *DescribeParameterTemplatesResponseBodyParametersTemplateRecord {
	s.CheckingCode = &v
	return s
}

func (s *DescribeParameterTemplatesResponseBodyParametersTemplateRecord) SetForceModify(v string) *DescribeParameterTemplatesResponseBodyParametersTemplateRecord {
	s.ForceModify = &v
	return s
}

func (s *DescribeParameterTemplatesResponseBodyParametersTemplateRecord) SetForceRestart(v string) *DescribeParameterTemplatesResponseBodyParametersTemplateRecord {
	s.ForceRestart = &v
	return s
}

func (s *DescribeParameterTemplatesResponseBodyParametersTemplateRecord) SetIsNodeAvailable(v string) *DescribeParameterTemplatesResponseBodyParametersTemplateRecord {
	s.IsNodeAvailable = &v
	return s
}

func (s *DescribeParameterTemplatesResponseBodyParametersTemplateRecord) SetParamRelyRule(v string) *DescribeParameterTemplatesResponseBodyParametersTemplateRecord {
	s.ParamRelyRule = &v
	return s
}

func (s *DescribeParameterTemplatesResponseBodyParametersTemplateRecord) SetParameterDescription(v string) *DescribeParameterTemplatesResponseBodyParametersTemplateRecord {
	s.ParameterDescription = &v
	return s
}

func (s *DescribeParameterTemplatesResponseBodyParametersTemplateRecord) SetParameterName(v string) *DescribeParameterTemplatesResponseBodyParametersTemplateRecord {
	s.ParameterName = &v
	return s
}

func (s *DescribeParameterTemplatesResponseBodyParametersTemplateRecord) SetParameterValue(v string) *DescribeParameterTemplatesResponseBodyParametersTemplateRecord {
	s.ParameterValue = &v
	return s
}

func (s *DescribeParameterTemplatesResponseBodyParametersTemplateRecord) Validate() error {
	return dara.Validate(s)
}
