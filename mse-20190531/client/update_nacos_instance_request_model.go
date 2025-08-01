// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNacosInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *UpdateNacosInstanceRequest
	GetAcceptLanguage() *string
	SetClusterName(v string) *UpdateNacosInstanceRequest
	GetClusterName() *string
	SetEnabled(v bool) *UpdateNacosInstanceRequest
	GetEnabled() *bool
	SetEphemeral(v bool) *UpdateNacosInstanceRequest
	GetEphemeral() *bool
	SetGroupName(v string) *UpdateNacosInstanceRequest
	GetGroupName() *string
	SetInstanceId(v string) *UpdateNacosInstanceRequest
	GetInstanceId() *string
	SetIp(v string) *UpdateNacosInstanceRequest
	GetIp() *string
	SetMetadata(v string) *UpdateNacosInstanceRequest
	GetMetadata() *string
	SetNamespaceId(v string) *UpdateNacosInstanceRequest
	GetNamespaceId() *string
	SetPort(v int32) *UpdateNacosInstanceRequest
	GetPort() *int32
	SetServiceName(v string) *UpdateNacosInstanceRequest
	GetServiceName() *string
	SetWeight(v string) *UpdateNacosInstanceRequest
	GetWeight() *string
}

type UpdateNacosInstanceRequest struct {
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
	// Specifies whether to disable this service. Valid values:
	//
	// 	- `true`: yes.
	//
	// 	- `false`: no.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// Specifies whether the node is a non-persistent node. Valid values:
	//
	// 	- `true`: yes.
	//
	// 	- `false`: no.
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
	// 1.2.X.X
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The metadata of the instance.
	//
	// example:
	//
	// [int]
	Metadata *string `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// The ID of the namespace.
	//
	// example:
	//
	// 9e78a671-4b9b-4dd4-99c1-0****
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

func (s UpdateNacosInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateNacosInstanceRequest) GoString() string {
	return s.String()
}

func (s *UpdateNacosInstanceRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *UpdateNacosInstanceRequest) GetClusterName() *string {
	return s.ClusterName
}

func (s *UpdateNacosInstanceRequest) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateNacosInstanceRequest) GetEphemeral() *bool {
	return s.Ephemeral
}

func (s *UpdateNacosInstanceRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *UpdateNacosInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateNacosInstanceRequest) GetIp() *string {
	return s.Ip
}

func (s *UpdateNacosInstanceRequest) GetMetadata() *string {
	return s.Metadata
}

func (s *UpdateNacosInstanceRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *UpdateNacosInstanceRequest) GetPort() *int32 {
	return s.Port
}

func (s *UpdateNacosInstanceRequest) GetServiceName() *string {
	return s.ServiceName
}

func (s *UpdateNacosInstanceRequest) GetWeight() *string {
	return s.Weight
}

func (s *UpdateNacosInstanceRequest) SetAcceptLanguage(v string) *UpdateNacosInstanceRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *UpdateNacosInstanceRequest) SetClusterName(v string) *UpdateNacosInstanceRequest {
	s.ClusterName = &v
	return s
}

func (s *UpdateNacosInstanceRequest) SetEnabled(v bool) *UpdateNacosInstanceRequest {
	s.Enabled = &v
	return s
}

func (s *UpdateNacosInstanceRequest) SetEphemeral(v bool) *UpdateNacosInstanceRequest {
	s.Ephemeral = &v
	return s
}

func (s *UpdateNacosInstanceRequest) SetGroupName(v string) *UpdateNacosInstanceRequest {
	s.GroupName = &v
	return s
}

func (s *UpdateNacosInstanceRequest) SetInstanceId(v string) *UpdateNacosInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateNacosInstanceRequest) SetIp(v string) *UpdateNacosInstanceRequest {
	s.Ip = &v
	return s
}

func (s *UpdateNacosInstanceRequest) SetMetadata(v string) *UpdateNacosInstanceRequest {
	s.Metadata = &v
	return s
}

func (s *UpdateNacosInstanceRequest) SetNamespaceId(v string) *UpdateNacosInstanceRequest {
	s.NamespaceId = &v
	return s
}

func (s *UpdateNacosInstanceRequest) SetPort(v int32) *UpdateNacosInstanceRequest {
	s.Port = &v
	return s
}

func (s *UpdateNacosInstanceRequest) SetServiceName(v string) *UpdateNacosInstanceRequest {
	s.ServiceName = &v
	return s
}

func (s *UpdateNacosInstanceRequest) SetWeight(v string) *UpdateNacosInstanceRequest {
	s.Weight = &v
	return s
}

func (s *UpdateNacosInstanceRequest) Validate() error {
	return dara.Validate(s)
}
