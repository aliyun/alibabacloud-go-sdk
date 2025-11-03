// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSinkApiDestinationParameters interface {
	dara.Model
	String() string
	GoString() string
	SetBodyParameters(v *SinkApiDestinationParametersBodyParameters) *SinkApiDestinationParameters
	GetBodyParameters() *SinkApiDestinationParametersBodyParameters
	SetHeaderParameters(v *SinkApiDestinationParametersHeaderParameters) *SinkApiDestinationParameters
	GetHeaderParameters() *SinkApiDestinationParametersHeaderParameters
	SetName(v string) *SinkApiDestinationParameters
	GetName() *string
	SetQueryStringParameters(v *SinkApiDestinationParametersQueryStringParameters) *SinkApiDestinationParameters
	GetQueryStringParameters() *SinkApiDestinationParametersQueryStringParameters
}

type SinkApiDestinationParameters struct {
	BodyParameters        *SinkApiDestinationParametersBodyParameters        `json:"BodyParameters,omitempty" xml:"BodyParameters,omitempty" type:"Struct"`
	HeaderParameters      *SinkApiDestinationParametersHeaderParameters      `json:"HeaderParameters,omitempty" xml:"HeaderParameters,omitempty" type:"Struct"`
	Name                  *string                                            `json:"Name,omitempty" xml:"Name,omitempty"`
	QueryStringParameters *SinkApiDestinationParametersQueryStringParameters `json:"QueryStringParameters,omitempty" xml:"QueryStringParameters,omitempty" type:"Struct"`
}

func (s SinkApiDestinationParameters) String() string {
	return dara.Prettify(s)
}

func (s SinkApiDestinationParameters) GoString() string {
	return s.String()
}

func (s *SinkApiDestinationParameters) GetBodyParameters() *SinkApiDestinationParametersBodyParameters {
	return s.BodyParameters
}

func (s *SinkApiDestinationParameters) GetHeaderParameters() *SinkApiDestinationParametersHeaderParameters {
	return s.HeaderParameters
}

func (s *SinkApiDestinationParameters) GetName() *string {
	return s.Name
}

func (s *SinkApiDestinationParameters) GetQueryStringParameters() *SinkApiDestinationParametersQueryStringParameters {
	return s.QueryStringParameters
}

func (s *SinkApiDestinationParameters) SetBodyParameters(v *SinkApiDestinationParametersBodyParameters) *SinkApiDestinationParameters {
	s.BodyParameters = v
	return s
}

func (s *SinkApiDestinationParameters) SetHeaderParameters(v *SinkApiDestinationParametersHeaderParameters) *SinkApiDestinationParameters {
	s.HeaderParameters = v
	return s
}

func (s *SinkApiDestinationParameters) SetName(v string) *SinkApiDestinationParameters {
	s.Name = &v
	return s
}

func (s *SinkApiDestinationParameters) SetQueryStringParameters(v *SinkApiDestinationParametersQueryStringParameters) *SinkApiDestinationParameters {
	s.QueryStringParameters = v
	return s
}

func (s *SinkApiDestinationParameters) Validate() error {
	if s.BodyParameters != nil {
		if err := s.BodyParameters.Validate(); err != nil {
			return err
		}
	}
	if s.HeaderParameters != nil {
		if err := s.HeaderParameters.Validate(); err != nil {
			return err
		}
	}
	if s.QueryStringParameters != nil {
		if err := s.QueryStringParameters.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SinkApiDestinationParametersBodyParameters struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s SinkApiDestinationParametersBodyParameters) String() string {
	return dara.Prettify(s)
}

func (s SinkApiDestinationParametersBodyParameters) GoString() string {
	return s.String()
}

func (s *SinkApiDestinationParametersBodyParameters) GetForm() *string {
	return s.Form
}

func (s *SinkApiDestinationParametersBodyParameters) GetTemplate() *string {
	return s.Template
}

func (s *SinkApiDestinationParametersBodyParameters) GetValue() *string {
	return s.Value
}

func (s *SinkApiDestinationParametersBodyParameters) SetForm(v string) *SinkApiDestinationParametersBodyParameters {
	s.Form = &v
	return s
}

func (s *SinkApiDestinationParametersBodyParameters) SetTemplate(v string) *SinkApiDestinationParametersBodyParameters {
	s.Template = &v
	return s
}

func (s *SinkApiDestinationParametersBodyParameters) SetValue(v string) *SinkApiDestinationParametersBodyParameters {
	s.Value = &v
	return s
}

func (s *SinkApiDestinationParametersBodyParameters) Validate() error {
	return dara.Validate(s)
}

type SinkApiDestinationParametersHeaderParameters struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s SinkApiDestinationParametersHeaderParameters) String() string {
	return dara.Prettify(s)
}

func (s SinkApiDestinationParametersHeaderParameters) GoString() string {
	return s.String()
}

func (s *SinkApiDestinationParametersHeaderParameters) GetForm() *string {
	return s.Form
}

func (s *SinkApiDestinationParametersHeaderParameters) GetTemplate() *string {
	return s.Template
}

func (s *SinkApiDestinationParametersHeaderParameters) GetValue() *string {
	return s.Value
}

func (s *SinkApiDestinationParametersHeaderParameters) SetForm(v string) *SinkApiDestinationParametersHeaderParameters {
	s.Form = &v
	return s
}

func (s *SinkApiDestinationParametersHeaderParameters) SetTemplate(v string) *SinkApiDestinationParametersHeaderParameters {
	s.Template = &v
	return s
}

func (s *SinkApiDestinationParametersHeaderParameters) SetValue(v string) *SinkApiDestinationParametersHeaderParameters {
	s.Value = &v
	return s
}

func (s *SinkApiDestinationParametersHeaderParameters) Validate() error {
	return dara.Validate(s)
}

type SinkApiDestinationParametersQueryStringParameters struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s SinkApiDestinationParametersQueryStringParameters) String() string {
	return dara.Prettify(s)
}

func (s SinkApiDestinationParametersQueryStringParameters) GoString() string {
	return s.String()
}

func (s *SinkApiDestinationParametersQueryStringParameters) GetForm() *string {
	return s.Form
}

func (s *SinkApiDestinationParametersQueryStringParameters) GetTemplate() *string {
	return s.Template
}

func (s *SinkApiDestinationParametersQueryStringParameters) GetValue() *string {
	return s.Value
}

func (s *SinkApiDestinationParametersQueryStringParameters) SetForm(v string) *SinkApiDestinationParametersQueryStringParameters {
	s.Form = &v
	return s
}

func (s *SinkApiDestinationParametersQueryStringParameters) SetTemplate(v string) *SinkApiDestinationParametersQueryStringParameters {
	s.Template = &v
	return s
}

func (s *SinkApiDestinationParametersQueryStringParameters) SetValue(v string) *SinkApiDestinationParametersQueryStringParameters {
	s.Value = &v
	return s
}

func (s *SinkApiDestinationParametersQueryStringParameters) Validate() error {
	return dara.Validate(s)
}
