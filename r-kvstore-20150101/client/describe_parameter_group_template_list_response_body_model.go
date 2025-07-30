// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeParameterGroupTemplateListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEngineVersion(v string) *DescribeParameterGroupTemplateListResponseBody
	GetEngineVersion() *string
	SetParameters(v []*DescribeParameterGroupTemplateListResponseBodyParameters) *DescribeParameterGroupTemplateListResponseBody
	GetParameters() []*DescribeParameterGroupTemplateListResponseBodyParameters
	SetRequestId(v string) *DescribeParameterGroupTemplateListResponseBody
	GetRequestId() *string
}

type DescribeParameterGroupTemplateListResponseBody struct {
	// The compatible engine version. Valid values:
	//
	// Redis Open Source Edition: 5.0, 6.0, and 7.0. Tair DRAM-based instances: 5.0 and 6.0. Tair persistent memory-optimized instances: 6.0. Tair ESSD-based instances: 4.0.
	//
	// example:
	//
	// 5
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The information about parameters.
	Parameters []*DescribeParameterGroupTemplateListResponseBodyParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// 5D622714-AEDD-4609-9167-F5DDD3D1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeParameterGroupTemplateListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeParameterGroupTemplateListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeParameterGroupTemplateListResponseBody) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *DescribeParameterGroupTemplateListResponseBody) GetParameters() []*DescribeParameterGroupTemplateListResponseBodyParameters {
	return s.Parameters
}

func (s *DescribeParameterGroupTemplateListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeParameterGroupTemplateListResponseBody) SetEngineVersion(v string) *DescribeParameterGroupTemplateListResponseBody {
	s.EngineVersion = &v
	return s
}

func (s *DescribeParameterGroupTemplateListResponseBody) SetParameters(v []*DescribeParameterGroupTemplateListResponseBodyParameters) *DescribeParameterGroupTemplateListResponseBody {
	s.Parameters = v
	return s
}

func (s *DescribeParameterGroupTemplateListResponseBody) SetRequestId(v string) *DescribeParameterGroupTemplateListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeParameterGroupTemplateListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeParameterGroupTemplateListResponseBodyParameters struct {
	// The regular expression used to validate input.
	//
	// example:
	//
	// "\\\\d+\\\\s+\\\\d+\\\\s+\\\\d+"
	CheckingCode *string `json:"CheckingCode,omitempty" xml:"CheckingCode,omitempty"`
	// Indicates whether the modification takes effect. Valid values: 0 and 1. A value of 0 indicates that the modification does not take effect. A value of 1 indicates that the modification takes effect.
	//
	// example:
	//
	// 1
	Effective *int64 `json:"Effective,omitempty" xml:"Effective,omitempty"`
	// A divisor of the parameter. For a parameter of the integer or byte type, the valid values must be a multiple of Factor unless you set Factor to 0.
	//
	// example:
	//
	// 1
	Factor *int64 `json:"Factor,omitempty" xml:"Factor,omitempty"`
	// The description of the parameter.
	//
	// example:
	//
	// Open AOF persistence mode
	ParameterDescription *string `json:"ParameterDescription,omitempty" xml:"ParameterDescription,omitempty"`
	// The parameter name.
	//
	// example:
	//
	// appendonly
	ParameterName *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
	// The default value of the parameter.
	//
	// example:
	//
	// 10
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
	// Indicates whether the parameter can be modified. Valid values:
	//
	// 	- **0: The parameter cannot be modified.**
	//
	// 	- **1**: The parameter can be modified.
	//
	// example:
	//
	// 0
	Revisable *int64 `json:"Revisable,omitempty" xml:"Revisable,omitempty"`
	// Indicates whether the minor version can be changed. Valid values: true and false.
	//
	// example:
	//
	// true
	SupportModifyForMinorVersion *bool `json:"SupportModifyForMinorVersion,omitempty" xml:"SupportModifyForMinorVersion,omitempty"`
	// The unit of the parameter value. Valid values: INT (ordinary integer), STRING (fixed string), and B (byte).
	//
	// example:
	//
	// STRING
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
}

func (s DescribeParameterGroupTemplateListResponseBodyParameters) String() string {
	return dara.Prettify(s)
}

func (s DescribeParameterGroupTemplateListResponseBodyParameters) GoString() string {
	return s.String()
}

func (s *DescribeParameterGroupTemplateListResponseBodyParameters) GetCheckingCode() *string {
	return s.CheckingCode
}

func (s *DescribeParameterGroupTemplateListResponseBodyParameters) GetEffective() *int64 {
	return s.Effective
}

func (s *DescribeParameterGroupTemplateListResponseBodyParameters) GetFactor() *int64 {
	return s.Factor
}

func (s *DescribeParameterGroupTemplateListResponseBodyParameters) GetParameterDescription() *string {
	return s.ParameterDescription
}

func (s *DescribeParameterGroupTemplateListResponseBodyParameters) GetParameterName() *string {
	return s.ParameterName
}

func (s *DescribeParameterGroupTemplateListResponseBodyParameters) GetParameterValue() *string {
	return s.ParameterValue
}

func (s *DescribeParameterGroupTemplateListResponseBodyParameters) GetRevisable() *int64 {
	return s.Revisable
}

func (s *DescribeParameterGroupTemplateListResponseBodyParameters) GetSupportModifyForMinorVersion() *bool {
	return s.SupportModifyForMinorVersion
}

func (s *DescribeParameterGroupTemplateListResponseBodyParameters) GetUnit() *string {
	return s.Unit
}

func (s *DescribeParameterGroupTemplateListResponseBodyParameters) SetCheckingCode(v string) *DescribeParameterGroupTemplateListResponseBodyParameters {
	s.CheckingCode = &v
	return s
}

func (s *DescribeParameterGroupTemplateListResponseBodyParameters) SetEffective(v int64) *DescribeParameterGroupTemplateListResponseBodyParameters {
	s.Effective = &v
	return s
}

func (s *DescribeParameterGroupTemplateListResponseBodyParameters) SetFactor(v int64) *DescribeParameterGroupTemplateListResponseBodyParameters {
	s.Factor = &v
	return s
}

func (s *DescribeParameterGroupTemplateListResponseBodyParameters) SetParameterDescription(v string) *DescribeParameterGroupTemplateListResponseBodyParameters {
	s.ParameterDescription = &v
	return s
}

func (s *DescribeParameterGroupTemplateListResponseBodyParameters) SetParameterName(v string) *DescribeParameterGroupTemplateListResponseBodyParameters {
	s.ParameterName = &v
	return s
}

func (s *DescribeParameterGroupTemplateListResponseBodyParameters) SetParameterValue(v string) *DescribeParameterGroupTemplateListResponseBodyParameters {
	s.ParameterValue = &v
	return s
}

func (s *DescribeParameterGroupTemplateListResponseBodyParameters) SetRevisable(v int64) *DescribeParameterGroupTemplateListResponseBodyParameters {
	s.Revisable = &v
	return s
}

func (s *DescribeParameterGroupTemplateListResponseBodyParameters) SetSupportModifyForMinorVersion(v bool) *DescribeParameterGroupTemplateListResponseBodyParameters {
	s.SupportModifyForMinorVersion = &v
	return s
}

func (s *DescribeParameterGroupTemplateListResponseBodyParameters) SetUnit(v string) *DescribeParameterGroupTemplateListResponseBodyParameters {
	s.Unit = &v
	return s
}

func (s *DescribeParameterGroupTemplateListResponseBodyParameters) Validate() error {
	return dara.Validate(s)
}
