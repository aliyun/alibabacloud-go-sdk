// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *GetServiceDetailRequest
	GetAppId() *string
	SetGroup(v string) *GetServiceDetailRequest
	GetGroup() *string
	SetIp(v string) *GetServiceDetailRequest
	GetIp() *string
	SetNamespace(v string) *GetServiceDetailRequest
	GetNamespace() *string
	SetOrigin(v string) *GetServiceDetailRequest
	GetOrigin() *string
	SetRegion(v string) *GetServiceDetailRequest
	GetRegion() *string
	SetRegistryType(v string) *GetServiceDetailRequest
	GetRegistryType() *string
	SetServiceId(v string) *GetServiceDetailRequest
	GetServiceId() *string
	SetServiceName(v string) *GetServiceDetailRequest
	GetServiceName() *string
	SetServiceType(v string) *GetServiceDetailRequest
	GetServiceType() *string
	SetServiceVersion(v string) *GetServiceDetailRequest
	GetServiceVersion() *string
	SetSource(v string) *GetServiceDetailRequest
	GetSource() *string
}

type GetServiceDetailRequest struct {
	// The ID of the application.
	//
	// example:
	//
	// efbda488-7b33-432f-a40d-****0047****
	AppId *string `json:"appId,omitempty" xml:"appId,omitempty"`
	// The group to which the service belongs.
	//
	// example:
	//
	// DUBBO
	Group *string `json:"group,omitempty" xml:"group,omitempty"`
	// The IP address of the service provider. Fuzzy searches are supported.
	//
	// example:
	//
	// 10.20.x.xx
	Ip *string `json:"ip,omitempty" xml:"ip,omitempty"`
	// The ID of the namespace.
	//
	// example:
	//
	// cn-hangzhou:doc-test
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// The source of the data. Valid values:
	//
	// 	- agent: Use this value if you use the service query feature of the latest version to pass the query result.
	//
	// 	- registry: Use this value if you use the service query feature of the earlier version to pass the query result.
	//
	// example:
	//
	// agent
	Origin *string `json:"origin,omitempty" xml:"origin,omitempty"`
	// The region ID of the service.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// The type of the service registry. This parameter is deprecated. You can ignore it.
	//
	// example:
	//
	// nacos
	RegistryType *string `json:"registryType,omitempty" xml:"registryType,omitempty"`
	// The ID of the service. This parameter is deprecated. You can ignore it.
	//
	// example:
	//
	// com.alibabacloud.hipstershop.CartService
	ServiceId *string `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
	// The name of the service.
	//
	// example:
	//
	// com.alibabacloud.hipstershop.CartService
	ServiceName *string `json:"serviceName,omitempty" xml:"serviceName,omitempty"`
	// The type of the service. Valid values:
	//
	// 	- dubbo
	//
	// 	- springCloud
	//
	// 	- hsf
	//
	// 	- istio
	//
	// example:
	//
	// springCloud
	ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
	// The version of the service.
	//
	// example:
	//
	// 1.0.0
	ServiceVersion *string `json:"serviceVersion,omitempty" xml:"serviceVersion,omitempty"`
	// The source of the service. Set the value to edas.
	//
	// example:
	//
	// edas
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
}

func (s GetServiceDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetServiceDetailRequest) GoString() string {
	return s.String()
}

func (s *GetServiceDetailRequest) GetAppId() *string {
	return s.AppId
}

func (s *GetServiceDetailRequest) GetGroup() *string {
	return s.Group
}

func (s *GetServiceDetailRequest) GetIp() *string {
	return s.Ip
}

func (s *GetServiceDetailRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *GetServiceDetailRequest) GetOrigin() *string {
	return s.Origin
}

func (s *GetServiceDetailRequest) GetRegion() *string {
	return s.Region
}

func (s *GetServiceDetailRequest) GetRegistryType() *string {
	return s.RegistryType
}

func (s *GetServiceDetailRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *GetServiceDetailRequest) GetServiceName() *string {
	return s.ServiceName
}

func (s *GetServiceDetailRequest) GetServiceType() *string {
	return s.ServiceType
}

func (s *GetServiceDetailRequest) GetServiceVersion() *string {
	return s.ServiceVersion
}

func (s *GetServiceDetailRequest) GetSource() *string {
	return s.Source
}

func (s *GetServiceDetailRequest) SetAppId(v string) *GetServiceDetailRequest {
	s.AppId = &v
	return s
}

func (s *GetServiceDetailRequest) SetGroup(v string) *GetServiceDetailRequest {
	s.Group = &v
	return s
}

func (s *GetServiceDetailRequest) SetIp(v string) *GetServiceDetailRequest {
	s.Ip = &v
	return s
}

func (s *GetServiceDetailRequest) SetNamespace(v string) *GetServiceDetailRequest {
	s.Namespace = &v
	return s
}

func (s *GetServiceDetailRequest) SetOrigin(v string) *GetServiceDetailRequest {
	s.Origin = &v
	return s
}

func (s *GetServiceDetailRequest) SetRegion(v string) *GetServiceDetailRequest {
	s.Region = &v
	return s
}

func (s *GetServiceDetailRequest) SetRegistryType(v string) *GetServiceDetailRequest {
	s.RegistryType = &v
	return s
}

func (s *GetServiceDetailRequest) SetServiceId(v string) *GetServiceDetailRequest {
	s.ServiceId = &v
	return s
}

func (s *GetServiceDetailRequest) SetServiceName(v string) *GetServiceDetailRequest {
	s.ServiceName = &v
	return s
}

func (s *GetServiceDetailRequest) SetServiceType(v string) *GetServiceDetailRequest {
	s.ServiceType = &v
	return s
}

func (s *GetServiceDetailRequest) SetServiceVersion(v string) *GetServiceDetailRequest {
	s.ServiceVersion = &v
	return s
}

func (s *GetServiceDetailRequest) SetSource(v string) *GetServiceDetailRequest {
	s.Source = &v
	return s
}

func (s *GetServiceDetailRequest) Validate() error {
	return dara.Validate(s)
}
