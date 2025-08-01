// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNacosInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *CreateNacosInstanceRequest
	GetAcceptLanguage() *string
	SetClusterName(v string) *CreateNacosInstanceRequest
	GetClusterName() *string
	SetEnabled(v bool) *CreateNacosInstanceRequest
	GetEnabled() *bool
	SetEphemeral(v bool) *CreateNacosInstanceRequest
	GetEphemeral() *bool
	SetGroupName(v string) *CreateNacosInstanceRequest
	GetGroupName() *string
	SetInstanceId(v string) *CreateNacosInstanceRequest
	GetInstanceId() *string
	SetIp(v string) *CreateNacosInstanceRequest
	GetIp() *string
	SetMetadata(v string) *CreateNacosInstanceRequest
	GetMetadata() *string
	SetNamespaceId(v string) *CreateNacosInstanceRequest
	GetNamespaceId() *string
	SetPort(v int32) *CreateNacosInstanceRequest
	GetPort() *int32
	SetServiceName(v string) *CreateNacosInstanceRequest
	GetServiceName() *string
	SetWeight(v string) *CreateNacosInstanceRequest
	GetWeight() *string
}

type CreateNacosInstanceRequest struct {
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
	// The name of the Nacos instance.
	//
	// example:
	//
	// DEFAULT
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// Specifies whether to enable the service for the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// Specifies whether to mark the instance as a temporary node.
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
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// mse-cn-123456
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The IP address of the Nacos instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1.2.xx.xx
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The node metadata of the instance.
	//
	// example:
	//
	// {\\"grayversion\\":\\"1.0\\",\\"preserved.register.source\\":\\"SPRING_CLOUD\\",\\"management.context-path\\":\\"\\"}
	Metadata *string `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// The ID of the namespace.
	//
	// example:
	//
	// 9e78a671-4b9b-4dd4-99c1-0b9da87****
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The port number of the Nacos instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12281
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The name of the service.
	//
	// This parameter is required.
	//
	// example:
	//
	// hello_service
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The weight. Valid values: 0 to 10000. The value must be an integer. A larger value indicates a higher frequency at which the instance is accessed.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Weight *string `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s CreateNacosInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateNacosInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateNacosInstanceRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *CreateNacosInstanceRequest) GetClusterName() *string {
	return s.ClusterName
}

func (s *CreateNacosInstanceRequest) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateNacosInstanceRequest) GetEphemeral() *bool {
	return s.Ephemeral
}

func (s *CreateNacosInstanceRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *CreateNacosInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateNacosInstanceRequest) GetIp() *string {
	return s.Ip
}

func (s *CreateNacosInstanceRequest) GetMetadata() *string {
	return s.Metadata
}

func (s *CreateNacosInstanceRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *CreateNacosInstanceRequest) GetPort() *int32 {
	return s.Port
}

func (s *CreateNacosInstanceRequest) GetServiceName() *string {
	return s.ServiceName
}

func (s *CreateNacosInstanceRequest) GetWeight() *string {
	return s.Weight
}

func (s *CreateNacosInstanceRequest) SetAcceptLanguage(v string) *CreateNacosInstanceRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *CreateNacosInstanceRequest) SetClusterName(v string) *CreateNacosInstanceRequest {
	s.ClusterName = &v
	return s
}

func (s *CreateNacosInstanceRequest) SetEnabled(v bool) *CreateNacosInstanceRequest {
	s.Enabled = &v
	return s
}

func (s *CreateNacosInstanceRequest) SetEphemeral(v bool) *CreateNacosInstanceRequest {
	s.Ephemeral = &v
	return s
}

func (s *CreateNacosInstanceRequest) SetGroupName(v string) *CreateNacosInstanceRequest {
	s.GroupName = &v
	return s
}

func (s *CreateNacosInstanceRequest) SetInstanceId(v string) *CreateNacosInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateNacosInstanceRequest) SetIp(v string) *CreateNacosInstanceRequest {
	s.Ip = &v
	return s
}

func (s *CreateNacosInstanceRequest) SetMetadata(v string) *CreateNacosInstanceRequest {
	s.Metadata = &v
	return s
}

func (s *CreateNacosInstanceRequest) SetNamespaceId(v string) *CreateNacosInstanceRequest {
	s.NamespaceId = &v
	return s
}

func (s *CreateNacosInstanceRequest) SetPort(v int32) *CreateNacosInstanceRequest {
	s.Port = &v
	return s
}

func (s *CreateNacosInstanceRequest) SetServiceName(v string) *CreateNacosInstanceRequest {
	s.ServiceName = &v
	return s
}

func (s *CreateNacosInstanceRequest) SetWeight(v string) *CreateNacosInstanceRequest {
	s.Weight = &v
	return s
}

func (s *CreateNacosInstanceRequest) Validate() error {
	return dara.Validate(s)
}
