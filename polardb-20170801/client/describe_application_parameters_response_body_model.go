// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApplicationParametersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetParameterTemplates(v *DescribeApplicationParametersResponseBodyParameterTemplates) *DescribeApplicationParametersResponseBody
	GetParameterTemplates() *DescribeApplicationParametersResponseBodyParameterTemplates
	SetParameters(v *DescribeApplicationParametersResponseBodyParameters) *DescribeApplicationParametersResponseBody
	GetParameters() *DescribeApplicationParametersResponseBodyParameters
	SetRequestId(v string) *DescribeApplicationParametersResponseBody
	GetRequestId() *string
}

type DescribeApplicationParametersResponseBody struct {
	ParameterTemplates *DescribeApplicationParametersResponseBodyParameterTemplates `json:"ParameterTemplates,omitempty" xml:"ParameterTemplates,omitempty" type:"Struct"`
	Parameters         *DescribeApplicationParametersResponseBodyParameters         `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Struct"`
	// example:
	//
	// 3E5CD764-FCCA-5C9C-838E-20E0DE84B2AF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeApplicationParametersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationParametersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApplicationParametersResponseBody) GetParameterTemplates() *DescribeApplicationParametersResponseBodyParameterTemplates {
	return s.ParameterTemplates
}

func (s *DescribeApplicationParametersResponseBody) GetParameters() *DescribeApplicationParametersResponseBodyParameters {
	return s.Parameters
}

func (s *DescribeApplicationParametersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeApplicationParametersResponseBody) SetParameterTemplates(v *DescribeApplicationParametersResponseBodyParameterTemplates) *DescribeApplicationParametersResponseBody {
	s.ParameterTemplates = v
	return s
}

func (s *DescribeApplicationParametersResponseBody) SetParameters(v *DescribeApplicationParametersResponseBodyParameters) *DescribeApplicationParametersResponseBody {
	s.Parameters = v
	return s
}

func (s *DescribeApplicationParametersResponseBody) SetRequestId(v string) *DescribeApplicationParametersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApplicationParametersResponseBody) Validate() error {
	if s.ParameterTemplates != nil {
		if err := s.ParameterTemplates.Validate(); err != nil {
			return err
		}
	}
	if s.Parameters != nil {
		if err := s.Parameters.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeApplicationParametersResponseBodyParameterTemplates struct {
	ComponentParameterTemplates []*DescribeApplicationParametersResponseBodyParameterTemplatesComponentParameterTemplates `json:"ComponentParameterTemplates,omitempty" xml:"ComponentParameterTemplates,omitempty" type:"Repeated"`
}

func (s DescribeApplicationParametersResponseBodyParameterTemplates) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationParametersResponseBodyParameterTemplates) GoString() string {
	return s.String()
}

func (s *DescribeApplicationParametersResponseBodyParameterTemplates) GetComponentParameterTemplates() []*DescribeApplicationParametersResponseBodyParameterTemplatesComponentParameterTemplates {
	return s.ComponentParameterTemplates
}

func (s *DescribeApplicationParametersResponseBodyParameterTemplates) SetComponentParameterTemplates(v []*DescribeApplicationParametersResponseBodyParameterTemplatesComponentParameterTemplates) *DescribeApplicationParametersResponseBodyParameterTemplates {
	s.ComponentParameterTemplates = v
	return s
}

func (s *DescribeApplicationParametersResponseBodyParameterTemplates) Validate() error {
	if s.ComponentParameterTemplates != nil {
		for _, item := range s.ComponentParameterTemplates {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeApplicationParametersResponseBodyParameterTemplatesComponentParameterTemplates struct {
	// example:
	//
	// pac-**************
	ComponentId *string `json:"ComponentId,omitempty" xml:"ComponentId,omitempty"`
	// example:
	//
	// supabase
	ComponentType *string                                                                                             `json:"ComponentType,omitempty" xml:"ComponentType,omitempty"`
	Parameters    []*DescribeApplicationParametersResponseBodyParameterTemplatesComponentParameterTemplatesParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
}

func (s DescribeApplicationParametersResponseBodyParameterTemplatesComponentParameterTemplates) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationParametersResponseBodyParameterTemplatesComponentParameterTemplates) GoString() string {
	return s.String()
}

func (s *DescribeApplicationParametersResponseBodyParameterTemplatesComponentParameterTemplates) GetComponentId() *string {
	return s.ComponentId
}

func (s *DescribeApplicationParametersResponseBodyParameterTemplatesComponentParameterTemplates) GetComponentType() *string {
	return s.ComponentType
}

func (s *DescribeApplicationParametersResponseBodyParameterTemplatesComponentParameterTemplates) GetParameters() []*DescribeApplicationParametersResponseBodyParameterTemplatesComponentParameterTemplatesParameters {
	return s.Parameters
}

func (s *DescribeApplicationParametersResponseBodyParameterTemplatesComponentParameterTemplates) SetComponentId(v string) *DescribeApplicationParametersResponseBodyParameterTemplatesComponentParameterTemplates {
	s.ComponentId = &v
	return s
}

func (s *DescribeApplicationParametersResponseBodyParameterTemplatesComponentParameterTemplates) SetComponentType(v string) *DescribeApplicationParametersResponseBodyParameterTemplatesComponentParameterTemplates {
	s.ComponentType = &v
	return s
}

func (s *DescribeApplicationParametersResponseBodyParameterTemplatesComponentParameterTemplates) SetParameters(v []*DescribeApplicationParametersResponseBodyParameterTemplatesComponentParameterTemplatesParameters) *DescribeApplicationParametersResponseBodyParameterTemplatesComponentParameterTemplates {
	s.Parameters = v
	return s
}

func (s *DescribeApplicationParametersResponseBodyParameterTemplatesComponentParameterTemplates) Validate() error {
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

type DescribeApplicationParametersResponseBodyParameterTemplatesComponentParameterTemplatesParameters struct {
	// example:
	//
	// default
	Default *string `json:"Default,omitempty" xml:"Default,omitempty"`
	// example:
	//
	// The name of the parameter
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// false
	NeedRestart *bool `json:"NeedRestart,omitempty" xml:"NeedRestart,omitempty"`
	// example:
	//
	// ^[a-zA-Z0-9]{1,20}$
	Pattern *string `json:"Pattern,omitempty" xml:"Pattern,omitempty"`
	// example:
	//
	// false
	ReadOnly *bool `json:"ReadOnly,omitempty" xml:"ReadOnly,omitempty"`
	// example:
	//
	// string
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeApplicationParametersResponseBodyParameterTemplatesComponentParameterTemplatesParameters) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationParametersResponseBodyParameterTemplatesComponentParameterTemplatesParameters) GoString() string {
	return s.String()
}

func (s *DescribeApplicationParametersResponseBodyParameterTemplatesComponentParameterTemplatesParameters) GetDefault() *string {
	return s.Default
}

func (s *DescribeApplicationParametersResponseBodyParameterTemplatesComponentParameterTemplatesParameters) GetDescription() *string {
	return s.Description
}

func (s *DescribeApplicationParametersResponseBodyParameterTemplatesComponentParameterTemplatesParameters) GetName() *string {
	return s.Name
}

func (s *DescribeApplicationParametersResponseBodyParameterTemplatesComponentParameterTemplatesParameters) GetNeedRestart() *bool {
	return s.NeedRestart
}

func (s *DescribeApplicationParametersResponseBodyParameterTemplatesComponentParameterTemplatesParameters) GetPattern() *string {
	return s.Pattern
}

func (s *DescribeApplicationParametersResponseBodyParameterTemplatesComponentParameterTemplatesParameters) GetReadOnly() *bool {
	return s.ReadOnly
}

func (s *DescribeApplicationParametersResponseBodyParameterTemplatesComponentParameterTemplatesParameters) GetType() *string {
	return s.Type
}

func (s *DescribeApplicationParametersResponseBodyParameterTemplatesComponentParameterTemplatesParameters) SetDefault(v string) *DescribeApplicationParametersResponseBodyParameterTemplatesComponentParameterTemplatesParameters {
	s.Default = &v
	return s
}

func (s *DescribeApplicationParametersResponseBodyParameterTemplatesComponentParameterTemplatesParameters) SetDescription(v string) *DescribeApplicationParametersResponseBodyParameterTemplatesComponentParameterTemplatesParameters {
	s.Description = &v
	return s
}

func (s *DescribeApplicationParametersResponseBodyParameterTemplatesComponentParameterTemplatesParameters) SetName(v string) *DescribeApplicationParametersResponseBodyParameterTemplatesComponentParameterTemplatesParameters {
	s.Name = &v
	return s
}

func (s *DescribeApplicationParametersResponseBodyParameterTemplatesComponentParameterTemplatesParameters) SetNeedRestart(v bool) *DescribeApplicationParametersResponseBodyParameterTemplatesComponentParameterTemplatesParameters {
	s.NeedRestart = &v
	return s
}

func (s *DescribeApplicationParametersResponseBodyParameterTemplatesComponentParameterTemplatesParameters) SetPattern(v string) *DescribeApplicationParametersResponseBodyParameterTemplatesComponentParameterTemplatesParameters {
	s.Pattern = &v
	return s
}

func (s *DescribeApplicationParametersResponseBodyParameterTemplatesComponentParameterTemplatesParameters) SetReadOnly(v bool) *DescribeApplicationParametersResponseBodyParameterTemplatesComponentParameterTemplatesParameters {
	s.ReadOnly = &v
	return s
}

func (s *DescribeApplicationParametersResponseBodyParameterTemplatesComponentParameterTemplatesParameters) SetType(v string) *DescribeApplicationParametersResponseBodyParameterTemplatesComponentParameterTemplatesParameters {
	s.Type = &v
	return s
}

func (s *DescribeApplicationParametersResponseBodyParameterTemplatesComponentParameterTemplatesParameters) Validate() error {
	return dara.Validate(s)
}

type DescribeApplicationParametersResponseBodyParameters struct {
	ComponentParameters []*DescribeApplicationParametersResponseBodyParametersComponentParameters `json:"ComponentParameters,omitempty" xml:"ComponentParameters,omitempty" type:"Repeated"`
}

func (s DescribeApplicationParametersResponseBodyParameters) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationParametersResponseBodyParameters) GoString() string {
	return s.String()
}

func (s *DescribeApplicationParametersResponseBodyParameters) GetComponentParameters() []*DescribeApplicationParametersResponseBodyParametersComponentParameters {
	return s.ComponentParameters
}

func (s *DescribeApplicationParametersResponseBodyParameters) SetComponentParameters(v []*DescribeApplicationParametersResponseBodyParametersComponentParameters) *DescribeApplicationParametersResponseBodyParameters {
	s.ComponentParameters = v
	return s
}

func (s *DescribeApplicationParametersResponseBodyParameters) Validate() error {
	if s.ComponentParameters != nil {
		for _, item := range s.ComponentParameters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeApplicationParametersResponseBodyParametersComponentParameters struct {
	// example:
	//
	// pac-**************
	ComponentId *string `json:"ComponentId,omitempty" xml:"ComponentId,omitempty"`
	// example:
	//
	// supabase
	ComponentType *string                                                                             `json:"ComponentType,omitempty" xml:"ComponentType,omitempty"`
	Parameters    []*DescribeApplicationParametersResponseBodyParametersComponentParametersParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
}

func (s DescribeApplicationParametersResponseBodyParametersComponentParameters) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationParametersResponseBodyParametersComponentParameters) GoString() string {
	return s.String()
}

func (s *DescribeApplicationParametersResponseBodyParametersComponentParameters) GetComponentId() *string {
	return s.ComponentId
}

func (s *DescribeApplicationParametersResponseBodyParametersComponentParameters) GetComponentType() *string {
	return s.ComponentType
}

func (s *DescribeApplicationParametersResponseBodyParametersComponentParameters) GetParameters() []*DescribeApplicationParametersResponseBodyParametersComponentParametersParameters {
	return s.Parameters
}

func (s *DescribeApplicationParametersResponseBodyParametersComponentParameters) SetComponentId(v string) *DescribeApplicationParametersResponseBodyParametersComponentParameters {
	s.ComponentId = &v
	return s
}

func (s *DescribeApplicationParametersResponseBodyParametersComponentParameters) SetComponentType(v string) *DescribeApplicationParametersResponseBodyParametersComponentParameters {
	s.ComponentType = &v
	return s
}

func (s *DescribeApplicationParametersResponseBodyParametersComponentParameters) SetParameters(v []*DescribeApplicationParametersResponseBodyParametersComponentParametersParameters) *DescribeApplicationParametersResponseBodyParametersComponentParameters {
	s.Parameters = v
	return s
}

func (s *DescribeApplicationParametersResponseBodyParametersComponentParameters) Validate() error {
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

type DescribeApplicationParametersResponseBodyParametersComponentParametersParameters struct {
	// example:
	//
	// default value
	Default *string `json:"Default,omitempty" xml:"Default,omitempty"`
	// example:
	//
	// The name of the parameter
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// false
	NeedRestart *bool `json:"NeedRestart,omitempty" xml:"NeedRestart,omitempty"`
	// example:
	//
	// ^[a-zA-Z0-9]{1,20}$
	Pattern *string `json:"Pattern,omitempty" xml:"Pattern,omitempty"`
	// example:
	//
	// false
	ReadOnly *bool `json:"ReadOnly,omitempty" xml:"ReadOnly,omitempty"`
	// example:
	//
	// Applied
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// string
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeApplicationParametersResponseBodyParametersComponentParametersParameters) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationParametersResponseBodyParametersComponentParametersParameters) GoString() string {
	return s.String()
}

func (s *DescribeApplicationParametersResponseBodyParametersComponentParametersParameters) GetDefault() *string {
	return s.Default
}

func (s *DescribeApplicationParametersResponseBodyParametersComponentParametersParameters) GetDescription() *string {
	return s.Description
}

func (s *DescribeApplicationParametersResponseBodyParametersComponentParametersParameters) GetName() *string {
	return s.Name
}

func (s *DescribeApplicationParametersResponseBodyParametersComponentParametersParameters) GetNeedRestart() *bool {
	return s.NeedRestart
}

func (s *DescribeApplicationParametersResponseBodyParametersComponentParametersParameters) GetPattern() *string {
	return s.Pattern
}

func (s *DescribeApplicationParametersResponseBodyParametersComponentParametersParameters) GetReadOnly() *bool {
	return s.ReadOnly
}

func (s *DescribeApplicationParametersResponseBodyParametersComponentParametersParameters) GetStatus() *string {
	return s.Status
}

func (s *DescribeApplicationParametersResponseBodyParametersComponentParametersParameters) GetType() *string {
	return s.Type
}

func (s *DescribeApplicationParametersResponseBodyParametersComponentParametersParameters) GetValue() *string {
	return s.Value
}

func (s *DescribeApplicationParametersResponseBodyParametersComponentParametersParameters) SetDefault(v string) *DescribeApplicationParametersResponseBodyParametersComponentParametersParameters {
	s.Default = &v
	return s
}

func (s *DescribeApplicationParametersResponseBodyParametersComponentParametersParameters) SetDescription(v string) *DescribeApplicationParametersResponseBodyParametersComponentParametersParameters {
	s.Description = &v
	return s
}

func (s *DescribeApplicationParametersResponseBodyParametersComponentParametersParameters) SetName(v string) *DescribeApplicationParametersResponseBodyParametersComponentParametersParameters {
	s.Name = &v
	return s
}

func (s *DescribeApplicationParametersResponseBodyParametersComponentParametersParameters) SetNeedRestart(v bool) *DescribeApplicationParametersResponseBodyParametersComponentParametersParameters {
	s.NeedRestart = &v
	return s
}

func (s *DescribeApplicationParametersResponseBodyParametersComponentParametersParameters) SetPattern(v string) *DescribeApplicationParametersResponseBodyParametersComponentParametersParameters {
	s.Pattern = &v
	return s
}

func (s *DescribeApplicationParametersResponseBodyParametersComponentParametersParameters) SetReadOnly(v bool) *DescribeApplicationParametersResponseBodyParametersComponentParametersParameters {
	s.ReadOnly = &v
	return s
}

func (s *DescribeApplicationParametersResponseBodyParametersComponentParametersParameters) SetStatus(v string) *DescribeApplicationParametersResponseBodyParametersComponentParametersParameters {
	s.Status = &v
	return s
}

func (s *DescribeApplicationParametersResponseBodyParametersComponentParametersParameters) SetType(v string) *DescribeApplicationParametersResponseBodyParametersComponentParametersParameters {
	s.Type = &v
	return s
}

func (s *DescribeApplicationParametersResponseBodyParametersComponentParametersParameters) SetValue(v string) *DescribeApplicationParametersResponseBodyParametersComponentParametersParameters {
	s.Value = &v
	return s
}

func (s *DescribeApplicationParametersResponseBodyParametersComponentParametersParameters) Validate() error {
	return dara.Validate(s)
}
