// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSourceRabbitMQMsgSyncParameters interface {
	dara.Model
	String() string
	GoString() string
	SetBodyDataType(v string) *SourceRabbitMQMsgSyncParameters
	GetBodyDataType() *string
	SetCloudMigrationTaskId(v string) *SourceRabbitMQMsgSyncParameters
	GetCloudMigrationTaskId() *string
	SetDstClusterHost(v string) *SourceRabbitMQMsgSyncParameters
	GetDstClusterHost() *string
	SetDstClusterType(v string) *SourceRabbitMQMsgSyncParameters
	GetDstClusterType() *string
	SetEndpoint(v string) *SourceRabbitMQMsgSyncParameters
	GetEndpoint() *string
	SetInstanceId(v string) *SourceRabbitMQMsgSyncParameters
	GetInstanceId() *string
	SetInstanceType(v string) *SourceRabbitMQMsgSyncParameters
	GetInstanceType() *string
	SetNetworkType(v string) *SourceRabbitMQMsgSyncParameters
	GetNetworkType() *string
	SetOrderConsume(v string) *SourceRabbitMQMsgSyncParameters
	GetOrderConsume() *string
	SetPassword(v string) *SourceRabbitMQMsgSyncParameters
	GetPassword() *string
	SetPrefetchCount(v string) *SourceRabbitMQMsgSyncParameters
	GetPrefetchCount() *string
	SetRegex(v string) *SourceRabbitMQMsgSyncParameters
	GetRegex() *string
	SetSecurityGroupId(v string) *SourceRabbitMQMsgSyncParameters
	GetSecurityGroupId() *string
	SetSrcClusterHost(v string) *SourceRabbitMQMsgSyncParameters
	GetSrcClusterHost() *string
	SetSrcClusterType(v string) *SourceRabbitMQMsgSyncParameters
	GetSrcClusterType() *string
	SetSyncMetaType(v string) *SourceRabbitMQMsgSyncParameters
	GetSyncMetaType() *string
	SetUsername(v string) *SourceRabbitMQMsgSyncParameters
	GetUsername() *string
	SetVSwitchIds(v string) *SourceRabbitMQMsgSyncParameters
	GetVSwitchIds() *string
	SetVirtualHostName(v string) *SourceRabbitMQMsgSyncParameters
	GetVirtualHostName() *string
	SetVpcId(v string) *SourceRabbitMQMsgSyncParameters
	GetVpcId() *string
}

type SourceRabbitMQMsgSyncParameters struct {
	BodyDataType         *string `json:"BodyDataType,omitempty" xml:"BodyDataType,omitempty"`
	CloudMigrationTaskId *string `json:"CloudMigrationTaskId,omitempty" xml:"CloudMigrationTaskId,omitempty"`
	DstClusterHost       *string `json:"DstClusterHost,omitempty" xml:"DstClusterHost,omitempty"`
	DstClusterType       *string `json:"DstClusterType,omitempty" xml:"DstClusterType,omitempty"`
	Endpoint             *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceType         *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	NetworkType          *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	OrderConsume         *string `json:"OrderConsume,omitempty" xml:"OrderConsume,omitempty"`
	Password             *string `json:"Password,omitempty" xml:"Password,omitempty"`
	PrefetchCount        *string `json:"PrefetchCount,omitempty" xml:"PrefetchCount,omitempty"`
	Regex                *string `json:"Regex,omitempty" xml:"Regex,omitempty"`
	SecurityGroupId      *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	SrcClusterHost       *string `json:"SrcClusterHost,omitempty" xml:"SrcClusterHost,omitempty"`
	SrcClusterType       *string `json:"SrcClusterType,omitempty" xml:"SrcClusterType,omitempty"`
	SyncMetaType         *string `json:"SyncMetaType,omitempty" xml:"SyncMetaType,omitempty"`
	Username             *string `json:"Username,omitempty" xml:"Username,omitempty"`
	VSwitchIds           *string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty"`
	VirtualHostName      *string `json:"VirtualHostName,omitempty" xml:"VirtualHostName,omitempty"`
	VpcId                *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s SourceRabbitMQMsgSyncParameters) String() string {
	return dara.Prettify(s)
}

func (s SourceRabbitMQMsgSyncParameters) GoString() string {
	return s.String()
}

func (s *SourceRabbitMQMsgSyncParameters) GetBodyDataType() *string {
	return s.BodyDataType
}

func (s *SourceRabbitMQMsgSyncParameters) GetCloudMigrationTaskId() *string {
	return s.CloudMigrationTaskId
}

func (s *SourceRabbitMQMsgSyncParameters) GetDstClusterHost() *string {
	return s.DstClusterHost
}

func (s *SourceRabbitMQMsgSyncParameters) GetDstClusterType() *string {
	return s.DstClusterType
}

func (s *SourceRabbitMQMsgSyncParameters) GetEndpoint() *string {
	return s.Endpoint
}

func (s *SourceRabbitMQMsgSyncParameters) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SourceRabbitMQMsgSyncParameters) GetInstanceType() *string {
	return s.InstanceType
}

func (s *SourceRabbitMQMsgSyncParameters) GetNetworkType() *string {
	return s.NetworkType
}

func (s *SourceRabbitMQMsgSyncParameters) GetOrderConsume() *string {
	return s.OrderConsume
}

func (s *SourceRabbitMQMsgSyncParameters) GetPassword() *string {
	return s.Password
}

func (s *SourceRabbitMQMsgSyncParameters) GetPrefetchCount() *string {
	return s.PrefetchCount
}

func (s *SourceRabbitMQMsgSyncParameters) GetRegex() *string {
	return s.Regex
}

func (s *SourceRabbitMQMsgSyncParameters) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *SourceRabbitMQMsgSyncParameters) GetSrcClusterHost() *string {
	return s.SrcClusterHost
}

func (s *SourceRabbitMQMsgSyncParameters) GetSrcClusterType() *string {
	return s.SrcClusterType
}

func (s *SourceRabbitMQMsgSyncParameters) GetSyncMetaType() *string {
	return s.SyncMetaType
}

func (s *SourceRabbitMQMsgSyncParameters) GetUsername() *string {
	return s.Username
}

func (s *SourceRabbitMQMsgSyncParameters) GetVSwitchIds() *string {
	return s.VSwitchIds
}

func (s *SourceRabbitMQMsgSyncParameters) GetVirtualHostName() *string {
	return s.VirtualHostName
}

func (s *SourceRabbitMQMsgSyncParameters) GetVpcId() *string {
	return s.VpcId
}

func (s *SourceRabbitMQMsgSyncParameters) SetBodyDataType(v string) *SourceRabbitMQMsgSyncParameters {
	s.BodyDataType = &v
	return s
}

func (s *SourceRabbitMQMsgSyncParameters) SetCloudMigrationTaskId(v string) *SourceRabbitMQMsgSyncParameters {
	s.CloudMigrationTaskId = &v
	return s
}

func (s *SourceRabbitMQMsgSyncParameters) SetDstClusterHost(v string) *SourceRabbitMQMsgSyncParameters {
	s.DstClusterHost = &v
	return s
}

func (s *SourceRabbitMQMsgSyncParameters) SetDstClusterType(v string) *SourceRabbitMQMsgSyncParameters {
	s.DstClusterType = &v
	return s
}

func (s *SourceRabbitMQMsgSyncParameters) SetEndpoint(v string) *SourceRabbitMQMsgSyncParameters {
	s.Endpoint = &v
	return s
}

func (s *SourceRabbitMQMsgSyncParameters) SetInstanceId(v string) *SourceRabbitMQMsgSyncParameters {
	s.InstanceId = &v
	return s
}

func (s *SourceRabbitMQMsgSyncParameters) SetInstanceType(v string) *SourceRabbitMQMsgSyncParameters {
	s.InstanceType = &v
	return s
}

func (s *SourceRabbitMQMsgSyncParameters) SetNetworkType(v string) *SourceRabbitMQMsgSyncParameters {
	s.NetworkType = &v
	return s
}

func (s *SourceRabbitMQMsgSyncParameters) SetOrderConsume(v string) *SourceRabbitMQMsgSyncParameters {
	s.OrderConsume = &v
	return s
}

func (s *SourceRabbitMQMsgSyncParameters) SetPassword(v string) *SourceRabbitMQMsgSyncParameters {
	s.Password = &v
	return s
}

func (s *SourceRabbitMQMsgSyncParameters) SetPrefetchCount(v string) *SourceRabbitMQMsgSyncParameters {
	s.PrefetchCount = &v
	return s
}

func (s *SourceRabbitMQMsgSyncParameters) SetRegex(v string) *SourceRabbitMQMsgSyncParameters {
	s.Regex = &v
	return s
}

func (s *SourceRabbitMQMsgSyncParameters) SetSecurityGroupId(v string) *SourceRabbitMQMsgSyncParameters {
	s.SecurityGroupId = &v
	return s
}

func (s *SourceRabbitMQMsgSyncParameters) SetSrcClusterHost(v string) *SourceRabbitMQMsgSyncParameters {
	s.SrcClusterHost = &v
	return s
}

func (s *SourceRabbitMQMsgSyncParameters) SetSrcClusterType(v string) *SourceRabbitMQMsgSyncParameters {
	s.SrcClusterType = &v
	return s
}

func (s *SourceRabbitMQMsgSyncParameters) SetSyncMetaType(v string) *SourceRabbitMQMsgSyncParameters {
	s.SyncMetaType = &v
	return s
}

func (s *SourceRabbitMQMsgSyncParameters) SetUsername(v string) *SourceRabbitMQMsgSyncParameters {
	s.Username = &v
	return s
}

func (s *SourceRabbitMQMsgSyncParameters) SetVSwitchIds(v string) *SourceRabbitMQMsgSyncParameters {
	s.VSwitchIds = &v
	return s
}

func (s *SourceRabbitMQMsgSyncParameters) SetVirtualHostName(v string) *SourceRabbitMQMsgSyncParameters {
	s.VirtualHostName = &v
	return s
}

func (s *SourceRabbitMQMsgSyncParameters) SetVpcId(v string) *SourceRabbitMQMsgSyncParameters {
	s.VpcId = &v
	return s
}

func (s *SourceRabbitMQMsgSyncParameters) Validate() error {
	return dara.Validate(s)
}
