// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConnectorProperties interface {
	dara.Model
	String() string
	GoString() string
	SetChannelType(v string) *ConnectorProperties
	GetChannelType() *string
	SetDryRunRequestBody(v string) *ConnectorProperties
	GetDryRunRequestBody() *string
	SetMaxRetries(v string) *ConnectorProperties
	GetMaxRetries() *string
	SetModelList(v string) *ConnectorProperties
	GetModelList() *string
	SetNetwork(v string) *ConnectorProperties
	GetNetwork() *string
	SetProtocol(v string) *ConnectorProperties
	GetProtocol() *string
	SetQpsLimit(v string) *ConnectorProperties
	GetQpsLimit() *string
	SetRegion(v string) *ConnectorProperties
	GetRegion() *string
	SetResponseBodyPath(v string) *ConnectorProperties
	GetResponseBodyPath() *string
	SetSecurityGroupId(v string) *ConnectorProperties
	GetSecurityGroupId() *string
	SetTimeoutMs(v string) *ConnectorProperties
	GetTimeoutMs() *string
	SetVSwitchId(v string) *ConnectorProperties
	GetVSwitchId() *string
	SetVpcId(v string) *ConnectorProperties
	GetVpcId() *string
}

type ConnectorProperties struct {
	// The channel type: custom or apig. This parameter is optional for model_service and defaults to custom.
	ChannelType *string `json:"channelType,omitempty" xml:"channelType,omitempty"`
	// The sample request body provided by the user for verifying endpoint connectivity. This parameter is required when dryRun is set to All and type is set to agent_app. The value is not persisted.
	DryRunRequestBody *string `json:"dryRunRequestBody,omitempty" xml:"dryRunRequestBody,omitempty"`
	// The number of retries after a failed invocation of the dial-test registration service. Valid values: 0 to 10. Default value: 3.
	//
	// example:
	//
	// 3
	MaxRetries *string `json:"maxRetries,omitempty" xml:"maxRetries,omitempty"`
	// The list of supported models in comma-separated format. This parameter is required for model_service.
	ModelList *string `json:"modelList,omitempty" xml:"modelList,omitempty"`
	// The network type: internet or aliyun-vpc.
	Network *string `json:"network,omitempty" xml:"network,omitempty"`
	// The protocol type: openai, openai-compatible, or anthropic. This parameter is required for model_service.
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// The QPS limit for the dial-test registration service. A value of 0 indicates no throttling. Otherwise, valid values: 0.1 to 1000. Default value: 20 for agent_app, 100 for model_service.
	//
	// example:
	//
	// 20
	QpsLimit *string `json:"qpsLimit,omitempty" xml:"qpsLimit,omitempty"`
	// The region. This parameter is required when the network type is aliyun-vpc.
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// The JSON Path extraction path for the response body. This parameter is optional for agent_app.
	ResponseBodyPath *string `json:"responseBodyPath,omitempty" xml:"responseBodyPath,omitempty"`
	// The security group ID. This parameter is optional for agent_app.
	SecurityGroupId *string `json:"securityGroupId,omitempty" xml:"securityGroupId,omitempty"`
	// The timeout for a single call, in milliseconds. Valid values: 1000 to 1800000. Default value: 300000.
	//
	// example:
	//
	// 30000
	TimeoutMs *string `json:"timeoutMs,omitempty" xml:"timeoutMs,omitempty"`
	// The vSwitch ID. This parameter is optional for agent_app.
	VSwitchId *string `json:"vSwitchId,omitempty" xml:"vSwitchId,omitempty"`
	// The VPC ID. This parameter is optional for agent_app.
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
}

func (s ConnectorProperties) String() string {
	return dara.Prettify(s)
}

func (s ConnectorProperties) GoString() string {
	return s.String()
}

func (s *ConnectorProperties) GetChannelType() *string {
	return s.ChannelType
}

func (s *ConnectorProperties) GetDryRunRequestBody() *string {
	return s.DryRunRequestBody
}

func (s *ConnectorProperties) GetMaxRetries() *string {
	return s.MaxRetries
}

func (s *ConnectorProperties) GetModelList() *string {
	return s.ModelList
}

func (s *ConnectorProperties) GetNetwork() *string {
	return s.Network
}

func (s *ConnectorProperties) GetProtocol() *string {
	return s.Protocol
}

func (s *ConnectorProperties) GetQpsLimit() *string {
	return s.QpsLimit
}

func (s *ConnectorProperties) GetRegion() *string {
	return s.Region
}

func (s *ConnectorProperties) GetResponseBodyPath() *string {
	return s.ResponseBodyPath
}

func (s *ConnectorProperties) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *ConnectorProperties) GetTimeoutMs() *string {
	return s.TimeoutMs
}

func (s *ConnectorProperties) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *ConnectorProperties) GetVpcId() *string {
	return s.VpcId
}

func (s *ConnectorProperties) SetChannelType(v string) *ConnectorProperties {
	s.ChannelType = &v
	return s
}

func (s *ConnectorProperties) SetDryRunRequestBody(v string) *ConnectorProperties {
	s.DryRunRequestBody = &v
	return s
}

func (s *ConnectorProperties) SetMaxRetries(v string) *ConnectorProperties {
	s.MaxRetries = &v
	return s
}

func (s *ConnectorProperties) SetModelList(v string) *ConnectorProperties {
	s.ModelList = &v
	return s
}

func (s *ConnectorProperties) SetNetwork(v string) *ConnectorProperties {
	s.Network = &v
	return s
}

func (s *ConnectorProperties) SetProtocol(v string) *ConnectorProperties {
	s.Protocol = &v
	return s
}

func (s *ConnectorProperties) SetQpsLimit(v string) *ConnectorProperties {
	s.QpsLimit = &v
	return s
}

func (s *ConnectorProperties) SetRegion(v string) *ConnectorProperties {
	s.Region = &v
	return s
}

func (s *ConnectorProperties) SetResponseBodyPath(v string) *ConnectorProperties {
	s.ResponseBodyPath = &v
	return s
}

func (s *ConnectorProperties) SetSecurityGroupId(v string) *ConnectorProperties {
	s.SecurityGroupId = &v
	return s
}

func (s *ConnectorProperties) SetTimeoutMs(v string) *ConnectorProperties {
	s.TimeoutMs = &v
	return s
}

func (s *ConnectorProperties) SetVSwitchId(v string) *ConnectorProperties {
	s.VSwitchId = &v
	return s
}

func (s *ConnectorProperties) SetVpcId(v string) *ConnectorProperties {
	s.VpcId = &v
	return s
}

func (s *ConnectorProperties) Validate() error {
	return dara.Validate(s)
}
