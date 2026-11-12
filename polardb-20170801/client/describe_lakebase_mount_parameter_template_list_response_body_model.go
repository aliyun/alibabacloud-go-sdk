// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLakebaseMountParameterTemplateListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetParameters(v []*DescribeLakebaseMountParameterTemplateListResponseBodyParameters) *DescribeLakebaseMountParameterTemplateListResponseBody
	GetParameters() []*DescribeLakebaseMountParameterTemplateListResponseBodyParameters
	SetRequestId(v string) *DescribeLakebaseMountParameterTemplateListResponseBody
	GetRequestId() *string
}

type DescribeLakebaseMountParameterTemplateListResponseBody struct {
	// The list of parameter templates.
	Parameters []*DescribeLakebaseMountParameterTemplateListResponseBodyParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// F45FFACC-xxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeLakebaseMountParameterTemplateListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLakebaseMountParameterTemplateListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLakebaseMountParameterTemplateListResponseBody) GetParameters() []*DescribeLakebaseMountParameterTemplateListResponseBodyParameters {
	return s.Parameters
}

func (s *DescribeLakebaseMountParameterTemplateListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLakebaseMountParameterTemplateListResponseBody) SetParameters(v []*DescribeLakebaseMountParameterTemplateListResponseBodyParameters) *DescribeLakebaseMountParameterTemplateListResponseBody {
	s.Parameters = v
	return s
}

func (s *DescribeLakebaseMountParameterTemplateListResponseBody) SetRequestId(v string) *DescribeLakebaseMountParameterTemplateListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLakebaseMountParameterTemplateListResponseBody) Validate() error {
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

type DescribeLakebaseMountParameterTemplateListResponseBodyParameters struct {
	// The parameter description.
	//
	// example:
	//
	// Auto-increment offset
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The default value.
	//
	// example:
	//
	// 1
	DefaultValue *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	// Indicates whether the parameter takes effect dynamically. Valid values: 1 and 0.
	//
	// example:
	//
	// 1
	IsDynamic *string `json:"IsDynamic,omitempty" xml:"IsDynamic,omitempty"`
	// Indicates whether the parameter can be modified by users. Valid values: 1 and 0.
	//
	// example:
	//
	// 1
	IsUserChangeable *string `json:"IsUserChangeable,omitempty" xml:"IsUserChangeable,omitempty"`
	// The value range or valid values. Example: [0-3600].
	//
	// example:
	//
	// [0-3600]
	Optional *string `json:"Optional,omitempty" xml:"Optional,omitempty"`
	// The parameter name.
	//
	// example:
	//
	// auto_increment_offset
	ParameterName *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
	// The unit.
	//
	// example:
	//
	// s
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
}

func (s DescribeLakebaseMountParameterTemplateListResponseBodyParameters) String() string {
	return dara.Prettify(s)
}

func (s DescribeLakebaseMountParameterTemplateListResponseBodyParameters) GoString() string {
	return s.String()
}

func (s *DescribeLakebaseMountParameterTemplateListResponseBodyParameters) GetComment() *string {
	return s.Comment
}

func (s *DescribeLakebaseMountParameterTemplateListResponseBodyParameters) GetDefaultValue() *string {
	return s.DefaultValue
}

func (s *DescribeLakebaseMountParameterTemplateListResponseBodyParameters) GetIsDynamic() *string {
	return s.IsDynamic
}

func (s *DescribeLakebaseMountParameterTemplateListResponseBodyParameters) GetIsUserChangeable() *string {
	return s.IsUserChangeable
}

func (s *DescribeLakebaseMountParameterTemplateListResponseBodyParameters) GetOptional() *string {
	return s.Optional
}

func (s *DescribeLakebaseMountParameterTemplateListResponseBodyParameters) GetParameterName() *string {
	return s.ParameterName
}

func (s *DescribeLakebaseMountParameterTemplateListResponseBodyParameters) GetUnit() *string {
	return s.Unit
}

func (s *DescribeLakebaseMountParameterTemplateListResponseBodyParameters) SetComment(v string) *DescribeLakebaseMountParameterTemplateListResponseBodyParameters {
	s.Comment = &v
	return s
}

func (s *DescribeLakebaseMountParameterTemplateListResponseBodyParameters) SetDefaultValue(v string) *DescribeLakebaseMountParameterTemplateListResponseBodyParameters {
	s.DefaultValue = &v
	return s
}

func (s *DescribeLakebaseMountParameterTemplateListResponseBodyParameters) SetIsDynamic(v string) *DescribeLakebaseMountParameterTemplateListResponseBodyParameters {
	s.IsDynamic = &v
	return s
}

func (s *DescribeLakebaseMountParameterTemplateListResponseBodyParameters) SetIsUserChangeable(v string) *DescribeLakebaseMountParameterTemplateListResponseBodyParameters {
	s.IsUserChangeable = &v
	return s
}

func (s *DescribeLakebaseMountParameterTemplateListResponseBodyParameters) SetOptional(v string) *DescribeLakebaseMountParameterTemplateListResponseBodyParameters {
	s.Optional = &v
	return s
}

func (s *DescribeLakebaseMountParameterTemplateListResponseBodyParameters) SetParameterName(v string) *DescribeLakebaseMountParameterTemplateListResponseBodyParameters {
	s.ParameterName = &v
	return s
}

func (s *DescribeLakebaseMountParameterTemplateListResponseBodyParameters) SetUnit(v string) *DescribeLakebaseMountParameterTemplateListResponseBodyParameters {
	s.Unit = &v
	return s
}

func (s *DescribeLakebaseMountParameterTemplateListResponseBodyParameters) Validate() error {
	return dara.Validate(s)
}
