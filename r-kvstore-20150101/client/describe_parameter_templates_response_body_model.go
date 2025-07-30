// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeParameterTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEngine(v string) *DescribeParameterTemplatesResponseBody
	GetEngine() *string
	SetEngineVersion(v string) *DescribeParameterTemplatesResponseBody
	GetEngineVersion() *string
	SetParameterCount(v string) *DescribeParameterTemplatesResponseBody
	GetParameterCount() *string
	SetParameters(v *DescribeParameterTemplatesResponseBodyParameters) *DescribeParameterTemplatesResponseBody
	GetParameters() *DescribeParameterTemplatesResponseBodyParameters
	SetRequestId(v string) *DescribeParameterTemplatesResponseBody
	GetRequestId() *string
}

type DescribeParameterTemplatesResponseBody struct {
	// The database engine that is run on the instance. The value **Redis*	- is returned for this parameter.
	//
	// example:
	//
	// redis
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The major version of the instance.
	//
	// example:
	//
	// 5.0
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The number of parameters that are supported by the instance.
	//
	// example:
	//
	// 24
	ParameterCount *string `json:"ParameterCount,omitempty" xml:"ParameterCount,omitempty"`
	// An array that consists of the details about the parameters returned.
	Parameters *DescribeParameterTemplatesResponseBodyParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 9DA28D8E-514D-4F12-ADED-70A9C818****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeParameterTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeParameterTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeParameterTemplatesResponseBody) GetEngine() *string {
	return s.Engine
}

func (s *DescribeParameterTemplatesResponseBody) GetEngineVersion() *string {
	return s.EngineVersion
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

func (s *DescribeParameterTemplatesResponseBody) SetEngine(v string) *DescribeParameterTemplatesResponseBody {
	s.Engine = &v
	return s
}

func (s *DescribeParameterTemplatesResponseBody) SetEngineVersion(v string) *DescribeParameterTemplatesResponseBody {
	s.EngineVersion = &v
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
	return dara.Validate(s)
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
	return dara.Validate(s)
}

type DescribeParameterTemplatesResponseBodyParametersTemplateRecord struct {
	// The valid values of the parameter.
	//
	// example:
	//
	// [yes|no]
	CheckingCode *string `json:"CheckingCode,omitempty" xml:"CheckingCode,omitempty"`
	// Indicates whether the parameter can be reconfigured. Valid values:
	//
	// 	- **true**: The parameter can be reconfigured.
	//
	// 	- **false**: The parameter cannot be reconfigured.
	//
	// example:
	//
	// true
	ForceModify *bool `json:"ForceModify,omitempty" xml:"ForceModify,omitempty"`
	// Indicates whether a restart of the instance is required after the parameter is reconfigured. Valid values:
	//
	// 	- **true**: After the parameter is reconfigured, you must restart the instance to make the new value of the parameter take effect.
	//
	// 	- **false**: After the parameter is reconfigured, the new value of the parameter immediately takes effect. You do not need to restart the instance.
	//
	// example:
	//
	// true
	ForceRestart *bool `json:"ForceRestart,omitempty" xml:"ForceRestart,omitempty"`
	// The description of the parameter.
	//
	// example:
	//
	// test description
	ParameterDescription *string `json:"ParameterDescription,omitempty" xml:"ParameterDescription,omitempty"`
	// The name of the parameter. For more information about the parameters and the parameter settings, see [Parameters](https://help.aliyun.com/document_detail/259681.html).
	//
	// example:
	//
	// appendonly
	ParameterName *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
	// The default value of the parameter.
	//
	// example:
	//
	// yes
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

func (s *DescribeParameterTemplatesResponseBodyParametersTemplateRecord) GetForceModify() *bool {
	return s.ForceModify
}

func (s *DescribeParameterTemplatesResponseBodyParametersTemplateRecord) GetForceRestart() *bool {
	return s.ForceRestart
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

func (s *DescribeParameterTemplatesResponseBodyParametersTemplateRecord) SetForceModify(v bool) *DescribeParameterTemplatesResponseBodyParametersTemplateRecord {
	s.ForceModify = &v
	return s
}

func (s *DescribeParameterTemplatesResponseBodyParametersTemplateRecord) SetForceRestart(v bool) *DescribeParameterTemplatesResponseBodyParametersTemplateRecord {
	s.ForceRestart = &v
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
