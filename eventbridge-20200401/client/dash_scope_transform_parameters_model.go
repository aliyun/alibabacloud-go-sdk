// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDashScopeTransformParameters interface {
	dara.Model
	String() string
	GoString() string
	SetApiKey(v string) *DashScopeTransformParameters
	GetApiKey() *string
	SetMessages(v []*DashScopeTransformParametersMessages) *DashScopeTransformParameters
	GetMessages() []*DashScopeTransformParametersMessages
	SetModel(v string) *DashScopeTransformParameters
	GetModel() *string
	SetRequestPerMinute(v int64) *DashScopeTransformParameters
	GetRequestPerMinute() *int64
	SetStructuredOutputJsonSchema(v string) *DashScopeTransformParameters
	GetStructuredOutputJsonSchema() *string
	SetTokenPerMinute(v int64) *DashScopeTransformParameters
	GetTokenPerMinute() *int64
}

type DashScopeTransformParameters struct {
	ApiKey                     *string                                 `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	Messages                   []*DashScopeTransformParametersMessages `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Repeated"`
	Model                      *string                                 `json:"Model,omitempty" xml:"Model,omitempty"`
	RequestPerMinute           *int64                                  `json:"RequestPerMinute,omitempty" xml:"RequestPerMinute,omitempty"`
	StructuredOutputJsonSchema *string                                 `json:"StructuredOutputJsonSchema,omitempty" xml:"StructuredOutputJsonSchema,omitempty"`
	TokenPerMinute             *int64                                  `json:"TokenPerMinute,omitempty" xml:"TokenPerMinute,omitempty"`
}

func (s DashScopeTransformParameters) String() string {
	return dara.Prettify(s)
}

func (s DashScopeTransformParameters) GoString() string {
	return s.String()
}

func (s *DashScopeTransformParameters) GetApiKey() *string {
	return s.ApiKey
}

func (s *DashScopeTransformParameters) GetMessages() []*DashScopeTransformParametersMessages {
	return s.Messages
}

func (s *DashScopeTransformParameters) GetModel() *string {
	return s.Model
}

func (s *DashScopeTransformParameters) GetRequestPerMinute() *int64 {
	return s.RequestPerMinute
}

func (s *DashScopeTransformParameters) GetStructuredOutputJsonSchema() *string {
	return s.StructuredOutputJsonSchema
}

func (s *DashScopeTransformParameters) GetTokenPerMinute() *int64 {
	return s.TokenPerMinute
}

func (s *DashScopeTransformParameters) SetApiKey(v string) *DashScopeTransformParameters {
	s.ApiKey = &v
	return s
}

func (s *DashScopeTransformParameters) SetMessages(v []*DashScopeTransformParametersMessages) *DashScopeTransformParameters {
	s.Messages = v
	return s
}

func (s *DashScopeTransformParameters) SetModel(v string) *DashScopeTransformParameters {
	s.Model = &v
	return s
}

func (s *DashScopeTransformParameters) SetRequestPerMinute(v int64) *DashScopeTransformParameters {
	s.RequestPerMinute = &v
	return s
}

func (s *DashScopeTransformParameters) SetStructuredOutputJsonSchema(v string) *DashScopeTransformParameters {
	s.StructuredOutputJsonSchema = &v
	return s
}

func (s *DashScopeTransformParameters) SetTokenPerMinute(v int64) *DashScopeTransformParameters {
	s.TokenPerMinute = &v
	return s
}

func (s *DashScopeTransformParameters) Validate() error {
	if s.Messages != nil {
		for _, item := range s.Messages {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DashScopeTransformParametersMessages struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Role     *string `json:"Role,omitempty" xml:"Role,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DashScopeTransformParametersMessages) String() string {
	return dara.Prettify(s)
}

func (s DashScopeTransformParametersMessages) GoString() string {
	return s.String()
}

func (s *DashScopeTransformParametersMessages) GetForm() *string {
	return s.Form
}

func (s *DashScopeTransformParametersMessages) GetRole() *string {
	return s.Role
}

func (s *DashScopeTransformParametersMessages) GetTemplate() *string {
	return s.Template
}

func (s *DashScopeTransformParametersMessages) GetValue() *string {
	return s.Value
}

func (s *DashScopeTransformParametersMessages) SetForm(v string) *DashScopeTransformParametersMessages {
	s.Form = &v
	return s
}

func (s *DashScopeTransformParametersMessages) SetRole(v string) *DashScopeTransformParametersMessages {
	s.Role = &v
	return s
}

func (s *DashScopeTransformParametersMessages) SetTemplate(v string) *DashScopeTransformParametersMessages {
	s.Template = &v
	return s
}

func (s *DashScopeTransformParametersMessages) SetValue(v string) *DashScopeTransformParametersMessages {
	s.Value = &v
	return s
}

func (s *DashScopeTransformParametersMessages) Validate() error {
	return dara.Validate(s)
}
