// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSinkRabbitMQMsgSyncParameters interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *SinkRabbitMQMsgSyncParametersBody) *SinkRabbitMQMsgSyncParameters
	GetBody() *SinkRabbitMQMsgSyncParametersBody
	SetEndpoint(v string) *SinkRabbitMQMsgSyncParameters
	GetEndpoint() *string
	SetExchange(v *SinkRabbitMQMsgSyncParametersExchange) *SinkRabbitMQMsgSyncParameters
	GetExchange() *SinkRabbitMQMsgSyncParametersExchange
	SetInstanceId(v string) *SinkRabbitMQMsgSyncParameters
	GetInstanceId() *string
	SetInstanceType(v string) *SinkRabbitMQMsgSyncParameters
	GetInstanceType() *string
	SetMaxHops(v string) *SinkRabbitMQMsgSyncParameters
	GetMaxHops() *string
	SetMessageId(v *SinkRabbitMQMsgSyncParametersMessageId) *SinkRabbitMQMsgSyncParameters
	GetMessageId() *SinkRabbitMQMsgSyncParametersMessageId
	SetNetworkType(v string) *SinkRabbitMQMsgSyncParameters
	GetNetworkType() *string
	SetPassword(v string) *SinkRabbitMQMsgSyncParameters
	GetPassword() *string
	SetProperties(v *SinkRabbitMQMsgSyncParametersProperties) *SinkRabbitMQMsgSyncParameters
	GetProperties() *SinkRabbitMQMsgSyncParametersProperties
	SetRoutingKey(v *SinkRabbitMQMsgSyncParametersRoutingKey) *SinkRabbitMQMsgSyncParameters
	GetRoutingKey() *SinkRabbitMQMsgSyncParametersRoutingKey
	SetSecurityGroupId(v string) *SinkRabbitMQMsgSyncParameters
	GetSecurityGroupId() *string
	SetUsername(v string) *SinkRabbitMQMsgSyncParameters
	GetUsername() *string
	SetVSwitchIds(v string) *SinkRabbitMQMsgSyncParameters
	GetVSwitchIds() *string
	SetVirtualHostName(v string) *SinkRabbitMQMsgSyncParameters
	GetVirtualHostName() *string
	SetVpcId(v string) *SinkRabbitMQMsgSyncParameters
	GetVpcId() *string
}

type SinkRabbitMQMsgSyncParameters struct {
	Body            *SinkRabbitMQMsgSyncParametersBody       `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
	Endpoint        *string                                  `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	Exchange        *SinkRabbitMQMsgSyncParametersExchange   `json:"Exchange,omitempty" xml:"Exchange,omitempty" type:"Struct"`
	InstanceId      *string                                  `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceType    *string                                  `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	MaxHops         *string                                  `json:"MaxHops,omitempty" xml:"MaxHops,omitempty"`
	MessageId       *SinkRabbitMQMsgSyncParametersMessageId  `json:"MessageId,omitempty" xml:"MessageId,omitempty" type:"Struct"`
	NetworkType     *string                                  `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	Password        *string                                  `json:"Password,omitempty" xml:"Password,omitempty"`
	Properties      *SinkRabbitMQMsgSyncParametersProperties `json:"Properties,omitempty" xml:"Properties,omitempty" type:"Struct"`
	RoutingKey      *SinkRabbitMQMsgSyncParametersRoutingKey `json:"RoutingKey,omitempty" xml:"RoutingKey,omitempty" type:"Struct"`
	SecurityGroupId *string                                  `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	Username        *string                                  `json:"Username,omitempty" xml:"Username,omitempty"`
	VSwitchIds      *string                                  `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty"`
	VirtualHostName *string                                  `json:"VirtualHostName,omitempty" xml:"VirtualHostName,omitempty"`
	VpcId           *string                                  `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s SinkRabbitMQMsgSyncParameters) String() string {
	return dara.Prettify(s)
}

func (s SinkRabbitMQMsgSyncParameters) GoString() string {
	return s.String()
}

func (s *SinkRabbitMQMsgSyncParameters) GetBody() *SinkRabbitMQMsgSyncParametersBody {
	return s.Body
}

func (s *SinkRabbitMQMsgSyncParameters) GetEndpoint() *string {
	return s.Endpoint
}

func (s *SinkRabbitMQMsgSyncParameters) GetExchange() *SinkRabbitMQMsgSyncParametersExchange {
	return s.Exchange
}

func (s *SinkRabbitMQMsgSyncParameters) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SinkRabbitMQMsgSyncParameters) GetInstanceType() *string {
	return s.InstanceType
}

func (s *SinkRabbitMQMsgSyncParameters) GetMaxHops() *string {
	return s.MaxHops
}

func (s *SinkRabbitMQMsgSyncParameters) GetMessageId() *SinkRabbitMQMsgSyncParametersMessageId {
	return s.MessageId
}

func (s *SinkRabbitMQMsgSyncParameters) GetNetworkType() *string {
	return s.NetworkType
}

func (s *SinkRabbitMQMsgSyncParameters) GetPassword() *string {
	return s.Password
}

func (s *SinkRabbitMQMsgSyncParameters) GetProperties() *SinkRabbitMQMsgSyncParametersProperties {
	return s.Properties
}

func (s *SinkRabbitMQMsgSyncParameters) GetRoutingKey() *SinkRabbitMQMsgSyncParametersRoutingKey {
	return s.RoutingKey
}

func (s *SinkRabbitMQMsgSyncParameters) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *SinkRabbitMQMsgSyncParameters) GetUsername() *string {
	return s.Username
}

func (s *SinkRabbitMQMsgSyncParameters) GetVSwitchIds() *string {
	return s.VSwitchIds
}

func (s *SinkRabbitMQMsgSyncParameters) GetVirtualHostName() *string {
	return s.VirtualHostName
}

func (s *SinkRabbitMQMsgSyncParameters) GetVpcId() *string {
	return s.VpcId
}

func (s *SinkRabbitMQMsgSyncParameters) SetBody(v *SinkRabbitMQMsgSyncParametersBody) *SinkRabbitMQMsgSyncParameters {
	s.Body = v
	return s
}

func (s *SinkRabbitMQMsgSyncParameters) SetEndpoint(v string) *SinkRabbitMQMsgSyncParameters {
	s.Endpoint = &v
	return s
}

func (s *SinkRabbitMQMsgSyncParameters) SetExchange(v *SinkRabbitMQMsgSyncParametersExchange) *SinkRabbitMQMsgSyncParameters {
	s.Exchange = v
	return s
}

func (s *SinkRabbitMQMsgSyncParameters) SetInstanceId(v string) *SinkRabbitMQMsgSyncParameters {
	s.InstanceId = &v
	return s
}

func (s *SinkRabbitMQMsgSyncParameters) SetInstanceType(v string) *SinkRabbitMQMsgSyncParameters {
	s.InstanceType = &v
	return s
}

func (s *SinkRabbitMQMsgSyncParameters) SetMaxHops(v string) *SinkRabbitMQMsgSyncParameters {
	s.MaxHops = &v
	return s
}

func (s *SinkRabbitMQMsgSyncParameters) SetMessageId(v *SinkRabbitMQMsgSyncParametersMessageId) *SinkRabbitMQMsgSyncParameters {
	s.MessageId = v
	return s
}

func (s *SinkRabbitMQMsgSyncParameters) SetNetworkType(v string) *SinkRabbitMQMsgSyncParameters {
	s.NetworkType = &v
	return s
}

func (s *SinkRabbitMQMsgSyncParameters) SetPassword(v string) *SinkRabbitMQMsgSyncParameters {
	s.Password = &v
	return s
}

func (s *SinkRabbitMQMsgSyncParameters) SetProperties(v *SinkRabbitMQMsgSyncParametersProperties) *SinkRabbitMQMsgSyncParameters {
	s.Properties = v
	return s
}

func (s *SinkRabbitMQMsgSyncParameters) SetRoutingKey(v *SinkRabbitMQMsgSyncParametersRoutingKey) *SinkRabbitMQMsgSyncParameters {
	s.RoutingKey = v
	return s
}

func (s *SinkRabbitMQMsgSyncParameters) SetSecurityGroupId(v string) *SinkRabbitMQMsgSyncParameters {
	s.SecurityGroupId = &v
	return s
}

func (s *SinkRabbitMQMsgSyncParameters) SetUsername(v string) *SinkRabbitMQMsgSyncParameters {
	s.Username = &v
	return s
}

func (s *SinkRabbitMQMsgSyncParameters) SetVSwitchIds(v string) *SinkRabbitMQMsgSyncParameters {
	s.VSwitchIds = &v
	return s
}

func (s *SinkRabbitMQMsgSyncParameters) SetVirtualHostName(v string) *SinkRabbitMQMsgSyncParameters {
	s.VirtualHostName = &v
	return s
}

func (s *SinkRabbitMQMsgSyncParameters) SetVpcId(v string) *SinkRabbitMQMsgSyncParameters {
	s.VpcId = &v
	return s
}

func (s *SinkRabbitMQMsgSyncParameters) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	if s.Exchange != nil {
		if err := s.Exchange.Validate(); err != nil {
			return err
		}
	}
	if s.MessageId != nil {
		if err := s.MessageId.Validate(); err != nil {
			return err
		}
	}
	if s.Properties != nil {
		if err := s.Properties.Validate(); err != nil {
			return err
		}
	}
	if s.RoutingKey != nil {
		if err := s.RoutingKey.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SinkRabbitMQMsgSyncParametersBody struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s SinkRabbitMQMsgSyncParametersBody) String() string {
	return dara.Prettify(s)
}

func (s SinkRabbitMQMsgSyncParametersBody) GoString() string {
	return s.String()
}

func (s *SinkRabbitMQMsgSyncParametersBody) GetForm() *string {
	return s.Form
}

func (s *SinkRabbitMQMsgSyncParametersBody) GetTemplate() *string {
	return s.Template
}

func (s *SinkRabbitMQMsgSyncParametersBody) GetValue() *string {
	return s.Value
}

func (s *SinkRabbitMQMsgSyncParametersBody) SetForm(v string) *SinkRabbitMQMsgSyncParametersBody {
	s.Form = &v
	return s
}

func (s *SinkRabbitMQMsgSyncParametersBody) SetTemplate(v string) *SinkRabbitMQMsgSyncParametersBody {
	s.Template = &v
	return s
}

func (s *SinkRabbitMQMsgSyncParametersBody) SetValue(v string) *SinkRabbitMQMsgSyncParametersBody {
	s.Value = &v
	return s
}

func (s *SinkRabbitMQMsgSyncParametersBody) Validate() error {
	return dara.Validate(s)
}

type SinkRabbitMQMsgSyncParametersExchange struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s SinkRabbitMQMsgSyncParametersExchange) String() string {
	return dara.Prettify(s)
}

func (s SinkRabbitMQMsgSyncParametersExchange) GoString() string {
	return s.String()
}

func (s *SinkRabbitMQMsgSyncParametersExchange) GetForm() *string {
	return s.Form
}

func (s *SinkRabbitMQMsgSyncParametersExchange) GetTemplate() *string {
	return s.Template
}

func (s *SinkRabbitMQMsgSyncParametersExchange) GetValue() *string {
	return s.Value
}

func (s *SinkRabbitMQMsgSyncParametersExchange) SetForm(v string) *SinkRabbitMQMsgSyncParametersExchange {
	s.Form = &v
	return s
}

func (s *SinkRabbitMQMsgSyncParametersExchange) SetTemplate(v string) *SinkRabbitMQMsgSyncParametersExchange {
	s.Template = &v
	return s
}

func (s *SinkRabbitMQMsgSyncParametersExchange) SetValue(v string) *SinkRabbitMQMsgSyncParametersExchange {
	s.Value = &v
	return s
}

func (s *SinkRabbitMQMsgSyncParametersExchange) Validate() error {
	return dara.Validate(s)
}

type SinkRabbitMQMsgSyncParametersMessageId struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s SinkRabbitMQMsgSyncParametersMessageId) String() string {
	return dara.Prettify(s)
}

func (s SinkRabbitMQMsgSyncParametersMessageId) GoString() string {
	return s.String()
}

func (s *SinkRabbitMQMsgSyncParametersMessageId) GetForm() *string {
	return s.Form
}

func (s *SinkRabbitMQMsgSyncParametersMessageId) GetTemplate() *string {
	return s.Template
}

func (s *SinkRabbitMQMsgSyncParametersMessageId) GetValue() *string {
	return s.Value
}

func (s *SinkRabbitMQMsgSyncParametersMessageId) SetForm(v string) *SinkRabbitMQMsgSyncParametersMessageId {
	s.Form = &v
	return s
}

func (s *SinkRabbitMQMsgSyncParametersMessageId) SetTemplate(v string) *SinkRabbitMQMsgSyncParametersMessageId {
	s.Template = &v
	return s
}

func (s *SinkRabbitMQMsgSyncParametersMessageId) SetValue(v string) *SinkRabbitMQMsgSyncParametersMessageId {
	s.Value = &v
	return s
}

func (s *SinkRabbitMQMsgSyncParametersMessageId) Validate() error {
	return dara.Validate(s)
}

type SinkRabbitMQMsgSyncParametersProperties struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s SinkRabbitMQMsgSyncParametersProperties) String() string {
	return dara.Prettify(s)
}

func (s SinkRabbitMQMsgSyncParametersProperties) GoString() string {
	return s.String()
}

func (s *SinkRabbitMQMsgSyncParametersProperties) GetForm() *string {
	return s.Form
}

func (s *SinkRabbitMQMsgSyncParametersProperties) GetTemplate() *string {
	return s.Template
}

func (s *SinkRabbitMQMsgSyncParametersProperties) GetValue() *string {
	return s.Value
}

func (s *SinkRabbitMQMsgSyncParametersProperties) SetForm(v string) *SinkRabbitMQMsgSyncParametersProperties {
	s.Form = &v
	return s
}

func (s *SinkRabbitMQMsgSyncParametersProperties) SetTemplate(v string) *SinkRabbitMQMsgSyncParametersProperties {
	s.Template = &v
	return s
}

func (s *SinkRabbitMQMsgSyncParametersProperties) SetValue(v string) *SinkRabbitMQMsgSyncParametersProperties {
	s.Value = &v
	return s
}

func (s *SinkRabbitMQMsgSyncParametersProperties) Validate() error {
	return dara.Validate(s)
}

type SinkRabbitMQMsgSyncParametersRoutingKey struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s SinkRabbitMQMsgSyncParametersRoutingKey) String() string {
	return dara.Prettify(s)
}

func (s SinkRabbitMQMsgSyncParametersRoutingKey) GoString() string {
	return s.String()
}

func (s *SinkRabbitMQMsgSyncParametersRoutingKey) GetForm() *string {
	return s.Form
}

func (s *SinkRabbitMQMsgSyncParametersRoutingKey) GetTemplate() *string {
	return s.Template
}

func (s *SinkRabbitMQMsgSyncParametersRoutingKey) GetValue() *string {
	return s.Value
}

func (s *SinkRabbitMQMsgSyncParametersRoutingKey) SetForm(v string) *SinkRabbitMQMsgSyncParametersRoutingKey {
	s.Form = &v
	return s
}

func (s *SinkRabbitMQMsgSyncParametersRoutingKey) SetTemplate(v string) *SinkRabbitMQMsgSyncParametersRoutingKey {
	s.Template = &v
	return s
}

func (s *SinkRabbitMQMsgSyncParametersRoutingKey) SetValue(v string) *SinkRabbitMQMsgSyncParametersRoutingKey {
	s.Value = &v
	return s
}

func (s *SinkRabbitMQMsgSyncParametersRoutingKey) Validate() error {
	return dara.Validate(s)
}
