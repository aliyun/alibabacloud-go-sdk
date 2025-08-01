// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNacosClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *UpdateNacosClusterRequest
	GetAcceptLanguage() *string
	SetCheckPort(v int32) *UpdateNacosClusterRequest
	GetCheckPort() *int32
	SetClusterName(v string) *UpdateNacosClusterRequest
	GetClusterName() *string
	SetGroupName(v string) *UpdateNacosClusterRequest
	GetGroupName() *string
	SetHealthChecker(v string) *UpdateNacosClusterRequest
	GetHealthChecker() *string
	SetInstanceId(v string) *UpdateNacosClusterRequest
	GetInstanceId() *string
	SetNamespaceId(v string) *UpdateNacosClusterRequest
	GetNamespaceId() *string
	SetServiceName(v string) *UpdateNacosClusterRequest
	GetServiceName() *string
	SetUseInstancePortForCheck(v bool) *UpdateNacosClusterRequest
	GetUseInstancePortForCheck() *bool
}

type UpdateNacosClusterRequest struct {
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
	// The port used for health checks.
	//
	// example:
	//
	// 80
	CheckPort *int32 `json:"CheckPort,omitempty" xml:"CheckPort,omitempty"`
	// The name of the Nacos cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// DEFAULT
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The name of the group.
	//
	// This parameter is required.
	//
	// example:
	//
	// DEFAULT_GROUP
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The type of the health check.
	//
	// example:
	//
	// {"type":"none"}
	HealthChecker *string `json:"HealthChecker,omitempty" xml:"HealthChecker,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// mse-cn-123456
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the namespace.
	//
	// example:
	//
	// 9e78a671-4b9b-4dd4-99c1-0b9da87d3dec
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The name of the service.
	//
	// This parameter is required.
	//
	// example:
	//
	// hello_service
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// Specifies whether to use the port of the instance for a health check.
	//
	// example:
	//
	// false
	UseInstancePortForCheck *bool `json:"UseInstancePortForCheck,omitempty" xml:"UseInstancePortForCheck,omitempty"`
}

func (s UpdateNacosClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateNacosClusterRequest) GoString() string {
	return s.String()
}

func (s *UpdateNacosClusterRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *UpdateNacosClusterRequest) GetCheckPort() *int32 {
	return s.CheckPort
}

func (s *UpdateNacosClusterRequest) GetClusterName() *string {
	return s.ClusterName
}

func (s *UpdateNacosClusterRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *UpdateNacosClusterRequest) GetHealthChecker() *string {
	return s.HealthChecker
}

func (s *UpdateNacosClusterRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateNacosClusterRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *UpdateNacosClusterRequest) GetServiceName() *string {
	return s.ServiceName
}

func (s *UpdateNacosClusterRequest) GetUseInstancePortForCheck() *bool {
	return s.UseInstancePortForCheck
}

func (s *UpdateNacosClusterRequest) SetAcceptLanguage(v string) *UpdateNacosClusterRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *UpdateNacosClusterRequest) SetCheckPort(v int32) *UpdateNacosClusterRequest {
	s.CheckPort = &v
	return s
}

func (s *UpdateNacosClusterRequest) SetClusterName(v string) *UpdateNacosClusterRequest {
	s.ClusterName = &v
	return s
}

func (s *UpdateNacosClusterRequest) SetGroupName(v string) *UpdateNacosClusterRequest {
	s.GroupName = &v
	return s
}

func (s *UpdateNacosClusterRequest) SetHealthChecker(v string) *UpdateNacosClusterRequest {
	s.HealthChecker = &v
	return s
}

func (s *UpdateNacosClusterRequest) SetInstanceId(v string) *UpdateNacosClusterRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateNacosClusterRequest) SetNamespaceId(v string) *UpdateNacosClusterRequest {
	s.NamespaceId = &v
	return s
}

func (s *UpdateNacosClusterRequest) SetServiceName(v string) *UpdateNacosClusterRequest {
	s.ServiceName = &v
	return s
}

func (s *UpdateNacosClusterRequest) SetUseInstancePortForCheck(v bool) *UpdateNacosClusterRequest {
	s.UseInstancePortForCheck = &v
	return s
}

func (s *UpdateNacosClusterRequest) Validate() error {
	return dara.Validate(s)
}
