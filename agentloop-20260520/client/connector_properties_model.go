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
	SetModelList(v string) *ConnectorProperties
	GetModelList() *string
	SetNetwork(v string) *ConnectorProperties
	GetNetwork() *string
	SetProtocol(v string) *ConnectorProperties
	GetProtocol() *string
	SetRegion(v string) *ConnectorProperties
	GetRegion() *string
	SetResponseBodyPath(v string) *ConnectorProperties
	GetResponseBodyPath() *string
	SetSecurityGroupId(v string) *ConnectorProperties
	GetSecurityGroupId() *string
	SetVSwitchId(v string) *ConnectorProperties
	GetVSwitchId() *string
	SetVpcId(v string) *ConnectorProperties
	GetVpcId() *string
}

type ConnectorProperties struct {
	ChannelType       *string `json:"channelType,omitempty" xml:"channelType,omitempty"`
	DryRunRequestBody *string `json:"dryRunRequestBody,omitempty" xml:"dryRunRequestBody,omitempty"`
	ModelList         *string `json:"modelList,omitempty" xml:"modelList,omitempty"`
	Network           *string `json:"network,omitempty" xml:"network,omitempty"`
	Protocol          *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	Region            *string `json:"region,omitempty" xml:"region,omitempty"`
	ResponseBodyPath  *string `json:"responseBodyPath,omitempty" xml:"responseBodyPath,omitempty"`
	SecurityGroupId   *string `json:"securityGroupId,omitempty" xml:"securityGroupId,omitempty"`
	VSwitchId         *string `json:"vSwitchId,omitempty" xml:"vSwitchId,omitempty"`
	VpcId             *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
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

func (s *ConnectorProperties) GetModelList() *string {
	return s.ModelList
}

func (s *ConnectorProperties) GetNetwork() *string {
	return s.Network
}

func (s *ConnectorProperties) GetProtocol() *string {
	return s.Protocol
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
