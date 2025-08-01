// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNacosInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *DeleteNacosInstanceRequest
	GetAcceptLanguage() *string
	SetClusterName(v string) *DeleteNacosInstanceRequest
	GetClusterName() *string
	SetEphemeral(v bool) *DeleteNacosInstanceRequest
	GetEphemeral() *bool
	SetGroupName(v string) *DeleteNacosInstanceRequest
	GetGroupName() *string
	SetInstanceId(v string) *DeleteNacosInstanceRequest
	GetInstanceId() *string
	SetIp(v string) *DeleteNacosInstanceRequest
	GetIp() *string
	SetNamespaceId(v string) *DeleteNacosInstanceRequest
	GetNamespaceId() *string
	SetPort(v int32) *DeleteNacosInstanceRequest
	GetPort() *int32
	SetServiceName(v string) *DeleteNacosInstanceRequest
	GetServiceName() *string
}

type DeleteNacosInstanceRequest struct {
	// The language of the response. Valid values:
	//
	// 	- zh: Chinese
	//
	// 	- en: English
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The alias of the cluster.
	//
	// example:
	//
	// DEFAULT
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// Specifies whether the node is an ephemeral node. Valid values:
	//
	// 	- `true`: yes
	//
	// 	- `false`: no
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	Ephemeral *bool `json:"Ephemeral,omitempty" xml:"Ephemeral,omitempty"`
	// The name of the group.
	//
	// This parameter is required.
	//
	// example:
	//
	// DEFAULT_GROUP
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The ID of the Nacos instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// mse-cn-st21v5****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The IP address of the Nacos instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10.237.1.32
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The ID of the namespace.
	//
	// example:
	//
	// 9e78a671-4b9b-4dd4-99c1-0b9da87****
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The port of the Nacos instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8080
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The name of the service.
	//
	// This parameter is required.
	//
	// example:
	//
	// hello_service
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s DeleteNacosInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteNacosInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeleteNacosInstanceRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *DeleteNacosInstanceRequest) GetClusterName() *string {
	return s.ClusterName
}

func (s *DeleteNacosInstanceRequest) GetEphemeral() *bool {
	return s.Ephemeral
}

func (s *DeleteNacosInstanceRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *DeleteNacosInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteNacosInstanceRequest) GetIp() *string {
	return s.Ip
}

func (s *DeleteNacosInstanceRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *DeleteNacosInstanceRequest) GetPort() *int32 {
	return s.Port
}

func (s *DeleteNacosInstanceRequest) GetServiceName() *string {
	return s.ServiceName
}

func (s *DeleteNacosInstanceRequest) SetAcceptLanguage(v string) *DeleteNacosInstanceRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *DeleteNacosInstanceRequest) SetClusterName(v string) *DeleteNacosInstanceRequest {
	s.ClusterName = &v
	return s
}

func (s *DeleteNacosInstanceRequest) SetEphemeral(v bool) *DeleteNacosInstanceRequest {
	s.Ephemeral = &v
	return s
}

func (s *DeleteNacosInstanceRequest) SetGroupName(v string) *DeleteNacosInstanceRequest {
	s.GroupName = &v
	return s
}

func (s *DeleteNacosInstanceRequest) SetInstanceId(v string) *DeleteNacosInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteNacosInstanceRequest) SetIp(v string) *DeleteNacosInstanceRequest {
	s.Ip = &v
	return s
}

func (s *DeleteNacosInstanceRequest) SetNamespaceId(v string) *DeleteNacosInstanceRequest {
	s.NamespaceId = &v
	return s
}

func (s *DeleteNacosInstanceRequest) SetPort(v int32) *DeleteNacosInstanceRequest {
	s.Port = &v
	return s
}

func (s *DeleteNacosInstanceRequest) SetServiceName(v string) *DeleteNacosInstanceRequest {
	s.ServiceName = &v
	return s
}

func (s *DeleteNacosInstanceRequest) Validate() error {
	return dara.Validate(s)
}
