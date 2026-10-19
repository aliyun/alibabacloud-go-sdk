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
	// The compatible version.
	//
	// example:
	//
	// 5
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The list of parameter details.
	Parameters []*DescribeParameterGroupTemplateListResponseBodyParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	// The request ID.
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

type DescribeParameterGroupTemplateListResponseBodyParameters struct {
	// The value range of the parameter.
	//
	// example:
	//
	// [0|1]
	CheckingCode *string `json:"CheckingCode,omitempty" xml:"CheckingCode,omitempty"`
	// Indicates whether the parameter has taken effect. Valid values:
	//
	// 	- **0**: No.
	//
	// 	- **1**: Yes.
	//
	// example:
	//
	// 1
	Effective *int64 `json:"Effective,omitempty" xml:"Effective,omitempty"`
	// The divisor factor. For integer and byte type parameters, the parameter value must not only fall within the valid range but also be a multiple of Factor (when Factor is not 0).
	//
	// example:
	//
	// 1
	Factor *int64 `json:"Factor,omitempty" xml:"Factor,omitempty"`
	// The parameter description.
	//
	// example:
	//
	// Enable AOF persistence mode
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
	// - **0**: The parameter cannot be modified.
	//
	// - **1**: The parameter can be modified.
	//
	// example:
	//
	// 0
	Revisable *int64 `json:"Revisable,omitempty" xml:"Revisable,omitempty"`
	// Indicates whether modification for minor versions is supported. Valid values:
	//
	// 	- **true**: Supported.
	//
	// 	- **false**: Not supported.
	//
	// example:
	//
	// true
	SupportModifyForMinorVersion *bool `json:"SupportModifyForMinorVersion,omitempty" xml:"SupportModifyForMinorVersion,omitempty"`
	// The unit of the parameter value. Valid values:
	//
	// 	- **INT**: a common integer.
	//
	// 	- **STRING**: a fixed string.
	//
	// 	- **B**: bytes.
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
