// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSinkMQTTParameters interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *SinkMQTTParameters
	GetInstanceId() *string
	SetMqtt5UserProperty(v *SinkMQTTParametersMqtt5UserProperty) *SinkMQTTParameters
	GetMqtt5UserProperty() *SinkMQTTParametersMqtt5UserProperty
	SetParentTopic(v string) *SinkMQTTParameters
	GetParentTopic() *string
	SetPayload(v *SinkMQTTParametersPayload) *SinkMQTTParameters
	GetPayload() *SinkMQTTParametersPayload
	SetSubTopic(v *SinkMQTTParametersSubTopic) *SinkMQTTParameters
	GetSubTopic() *SinkMQTTParametersSubTopic
}

type SinkMQTTParameters struct {
	// The ID of the destination MQTT instance.
	//
	// example:
	//
	// mqtt-cn-a1b2c3d4e5f
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The user properties.
	Mqtt5UserProperty *SinkMQTTParametersMqtt5UserProperty `json:"Mqtt5UserProperty,omitempty" xml:"Mqtt5UserProperty,omitempty" type:"Struct"`
	// The parent topic of the destination MQTT instance.
	//
	// example:
	//
	// example-topic
	ParentTopic *string `json:"ParentTopic,omitempty" xml:"ParentTopic,omitempty"`
	// The message payload.
	Payload *SinkMQTTParametersPayload `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	// The subtopic name.
	SubTopic *SinkMQTTParametersSubTopic `json:"SubTopic,omitempty" xml:"SubTopic,omitempty" type:"Struct"`
}

func (s SinkMQTTParameters) String() string {
	return dara.Prettify(s)
}

func (s SinkMQTTParameters) GoString() string {
	return s.String()
}

func (s *SinkMQTTParameters) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SinkMQTTParameters) GetMqtt5UserProperty() *SinkMQTTParametersMqtt5UserProperty {
	return s.Mqtt5UserProperty
}

func (s *SinkMQTTParameters) GetParentTopic() *string {
	return s.ParentTopic
}

func (s *SinkMQTTParameters) GetPayload() *SinkMQTTParametersPayload {
	return s.Payload
}

func (s *SinkMQTTParameters) GetSubTopic() *SinkMQTTParametersSubTopic {
	return s.SubTopic
}

func (s *SinkMQTTParameters) SetInstanceId(v string) *SinkMQTTParameters {
	s.InstanceId = &v
	return s
}

func (s *SinkMQTTParameters) SetMqtt5UserProperty(v *SinkMQTTParametersMqtt5UserProperty) *SinkMQTTParameters {
	s.Mqtt5UserProperty = v
	return s
}

func (s *SinkMQTTParameters) SetParentTopic(v string) *SinkMQTTParameters {
	s.ParentTopic = &v
	return s
}

func (s *SinkMQTTParameters) SetPayload(v *SinkMQTTParametersPayload) *SinkMQTTParameters {
	s.Payload = v
	return s
}

func (s *SinkMQTTParameters) SetSubTopic(v *SinkMQTTParametersSubTopic) *SinkMQTTParameters {
	s.SubTopic = v
	return s
}

func (s *SinkMQTTParameters) Validate() error {
	if s.Mqtt5UserProperty != nil {
		if err := s.Mqtt5UserProperty.Validate(); err != nil {
			return err
		}
	}
	if s.Payload != nil {
		if err := s.Payload.Validate(); err != nil {
			return err
		}
	}
	if s.SubTopic != nil {
		if err := s.SubTopic.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SinkMQTTParametersMqtt5UserProperty struct {
	// The method used to generate the parameter value. Valid values: CONSTANT (constant), JSONPATH (extracted from event content), and TEMPLATE (template).
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template string used to format the parameter value. This parameter is used only when Form is set to TEMPLATE.
	//
	// example:
	//
	// The value of ${key} is ${value}!
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The source content of the parameter value specified by Form. When Form is set to CONSTANT, this value is the content.
	//
	// example:
	//
	// DEMO CONTENT
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s SinkMQTTParametersMqtt5UserProperty) String() string {
	return dara.Prettify(s)
}

func (s SinkMQTTParametersMqtt5UserProperty) GoString() string {
	return s.String()
}

func (s *SinkMQTTParametersMqtt5UserProperty) GetForm() *string {
	return s.Form
}

func (s *SinkMQTTParametersMqtt5UserProperty) GetTemplate() *string {
	return s.Template
}

func (s *SinkMQTTParametersMqtt5UserProperty) GetValue() *string {
	return s.Value
}

func (s *SinkMQTTParametersMqtt5UserProperty) SetForm(v string) *SinkMQTTParametersMqtt5UserProperty {
	s.Form = &v
	return s
}

func (s *SinkMQTTParametersMqtt5UserProperty) SetTemplate(v string) *SinkMQTTParametersMqtt5UserProperty {
	s.Template = &v
	return s
}

func (s *SinkMQTTParametersMqtt5UserProperty) SetValue(v string) *SinkMQTTParametersMqtt5UserProperty {
	s.Value = &v
	return s
}

func (s *SinkMQTTParametersMqtt5UserProperty) Validate() error {
	return dara.Validate(s)
}

type SinkMQTTParametersPayload struct {
	// The method used to generate the message content.
	//
	// example:
	//
	// ORIGINAL
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template string used to format the parameter value. This parameter is used only when Form is set to TEMPLATE.
	//
	// example:
	//
	// The value of ${key} is ${value}!
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The source content of the parameter value specified by Form. When Form is set to CONSTANT, this value is the content.
	//
	// example:
	//
	// DEMO CONTENT
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s SinkMQTTParametersPayload) String() string {
	return dara.Prettify(s)
}

func (s SinkMQTTParametersPayload) GoString() string {
	return s.String()
}

func (s *SinkMQTTParametersPayload) GetForm() *string {
	return s.Form
}

func (s *SinkMQTTParametersPayload) GetTemplate() *string {
	return s.Template
}

func (s *SinkMQTTParametersPayload) GetValue() *string {
	return s.Value
}

func (s *SinkMQTTParametersPayload) SetForm(v string) *SinkMQTTParametersPayload {
	s.Form = &v
	return s
}

func (s *SinkMQTTParametersPayload) SetTemplate(v string) *SinkMQTTParametersPayload {
	s.Template = &v
	return s
}

func (s *SinkMQTTParametersPayload) SetValue(v string) *SinkMQTTParametersPayload {
	s.Value = &v
	return s
}

func (s *SinkMQTTParametersPayload) Validate() error {
	return dara.Validate(s)
}

type SinkMQTTParametersSubTopic struct {
	// The method used to generate the parameter value. Valid values: CONSTANT (constant), JSONPATH (extracted from event content), and TEMPLATE (template).
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template string used to format the parameter value. This parameter is used only when Form is set to TEMPLATE.
	//
	// example:
	//
	// The value of ${key} is ${value}!
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The source content of the parameter value specified by Form. When Form is set to CONSTANT, this value is the content.
	//
	// example:
	//
	// SUB-TOPIC
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s SinkMQTTParametersSubTopic) String() string {
	return dara.Prettify(s)
}

func (s SinkMQTTParametersSubTopic) GoString() string {
	return s.String()
}

func (s *SinkMQTTParametersSubTopic) GetForm() *string {
	return s.Form
}

func (s *SinkMQTTParametersSubTopic) GetTemplate() *string {
	return s.Template
}

func (s *SinkMQTTParametersSubTopic) GetValue() *string {
	return s.Value
}

func (s *SinkMQTTParametersSubTopic) SetForm(v string) *SinkMQTTParametersSubTopic {
	s.Form = &v
	return s
}

func (s *SinkMQTTParametersSubTopic) SetTemplate(v string) *SinkMQTTParametersSubTopic {
	s.Template = &v
	return s
}

func (s *SinkMQTTParametersSubTopic) SetValue(v string) *SinkMQTTParametersSubTopic {
	s.Value = &v
	return s
}

func (s *SinkMQTTParametersSubTopic) Validate() error {
	return dara.Validate(s)
}
