// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceProvidersPageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *GetServiceProvidersPageRequest
	GetAppId() *string
	SetGroup(v string) *GetServiceProvidersPageRequest
	GetGroup() *string
	SetIp(v string) *GetServiceProvidersPageRequest
	GetIp() *string
	SetNamespace(v string) *GetServiceProvidersPageRequest
	GetNamespace() *string
	SetOrigin(v string) *GetServiceProvidersPageRequest
	GetOrigin() *string
	SetPage(v int32) *GetServiceProvidersPageRequest
	GetPage() *int32
	SetRegion(v string) *GetServiceProvidersPageRequest
	GetRegion() *string
	SetRegistryType(v string) *GetServiceProvidersPageRequest
	GetRegistryType() *string
	SetServiceId(v string) *GetServiceProvidersPageRequest
	GetServiceId() *string
	SetServiceName(v string) *GetServiceProvidersPageRequest
	GetServiceName() *string
	SetServiceType(v string) *GetServiceProvidersPageRequest
	GetServiceType() *string
	SetServiceVersion(v string) *GetServiceProvidersPageRequest
	GetServiceVersion() *string
	SetSize(v int32) *GetServiceProvidersPageRequest
	GetSize() *int32
	SetSource(v string) *GetServiceProvidersPageRequest
	GetSource() *string
}

type GetServiceProvidersPageRequest struct {
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
	// The source of data. Valid values:
	//
	// 	- agent: Use this value if you use the service query feature of the latest version to pass the query result.
	//
	// 	- registry: Use this value if you use the service query feature of the earlier version to pass the query result.
	//
	// example:
	//
	// agent
	Origin *string `json:"origin,omitempty" xml:"origin,omitempty"`
	// The number of the page to return. Pages start from Page 0.
	//
	// example:
	//
	// 0
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// The ID of the region.
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
	// 	- dubbo: Dubbo service
	//
	// 	- springCloud: Spring Cloud service
	//
	// 	- hsf: High-speed Service Framework (HSF) service
	//
	// example:
	//
	// dubbo
	ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
	// The version of the service.
	//
	// example:
	//
	// 1.0.0
	ServiceVersion *string `json:"serviceVersion,omitempty" xml:"serviceVersion,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 10
	Size *int32 `json:"size,omitempty" xml:"size,omitempty"`
	// The source of the service. Set the value to edas.
	//
	// example:
	//
	// edas
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
}

func (s GetServiceProvidersPageRequest) String() string {
	return dara.Prettify(s)
}

func (s GetServiceProvidersPageRequest) GoString() string {
	return s.String()
}

func (s *GetServiceProvidersPageRequest) GetAppId() *string {
	return s.AppId
}

func (s *GetServiceProvidersPageRequest) GetGroup() *string {
	return s.Group
}

func (s *GetServiceProvidersPageRequest) GetIp() *string {
	return s.Ip
}

func (s *GetServiceProvidersPageRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *GetServiceProvidersPageRequest) GetOrigin() *string {
	return s.Origin
}

func (s *GetServiceProvidersPageRequest) GetPage() *int32 {
	return s.Page
}

func (s *GetServiceProvidersPageRequest) GetRegion() *string {
	return s.Region
}

func (s *GetServiceProvidersPageRequest) GetRegistryType() *string {
	return s.RegistryType
}

func (s *GetServiceProvidersPageRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *GetServiceProvidersPageRequest) GetServiceName() *string {
	return s.ServiceName
}

func (s *GetServiceProvidersPageRequest) GetServiceType() *string {
	return s.ServiceType
}

func (s *GetServiceProvidersPageRequest) GetServiceVersion() *string {
	return s.ServiceVersion
}

func (s *GetServiceProvidersPageRequest) GetSize() *int32 {
	return s.Size
}

func (s *GetServiceProvidersPageRequest) GetSource() *string {
	return s.Source
}

func (s *GetServiceProvidersPageRequest) SetAppId(v string) *GetServiceProvidersPageRequest {
	s.AppId = &v
	return s
}

func (s *GetServiceProvidersPageRequest) SetGroup(v string) *GetServiceProvidersPageRequest {
	s.Group = &v
	return s
}

func (s *GetServiceProvidersPageRequest) SetIp(v string) *GetServiceProvidersPageRequest {
	s.Ip = &v
	return s
}

func (s *GetServiceProvidersPageRequest) SetNamespace(v string) *GetServiceProvidersPageRequest {
	s.Namespace = &v
	return s
}

func (s *GetServiceProvidersPageRequest) SetOrigin(v string) *GetServiceProvidersPageRequest {
	s.Origin = &v
	return s
}

func (s *GetServiceProvidersPageRequest) SetPage(v int32) *GetServiceProvidersPageRequest {
	s.Page = &v
	return s
}

func (s *GetServiceProvidersPageRequest) SetRegion(v string) *GetServiceProvidersPageRequest {
	s.Region = &v
	return s
}

func (s *GetServiceProvidersPageRequest) SetRegistryType(v string) *GetServiceProvidersPageRequest {
	s.RegistryType = &v
	return s
}

func (s *GetServiceProvidersPageRequest) SetServiceId(v string) *GetServiceProvidersPageRequest {
	s.ServiceId = &v
	return s
}

func (s *GetServiceProvidersPageRequest) SetServiceName(v string) *GetServiceProvidersPageRequest {
	s.ServiceName = &v
	return s
}

func (s *GetServiceProvidersPageRequest) SetServiceType(v string) *GetServiceProvidersPageRequest {
	s.ServiceType = &v
	return s
}

func (s *GetServiceProvidersPageRequest) SetServiceVersion(v string) *GetServiceProvidersPageRequest {
	s.ServiceVersion = &v
	return s
}

func (s *GetServiceProvidersPageRequest) SetSize(v int32) *GetServiceProvidersPageRequest {
	s.Size = &v
	return s
}

func (s *GetServiceProvidersPageRequest) SetSource(v string) *GetServiceProvidersPageRequest {
	s.Source = &v
	return s
}

func (s *GetServiceProvidersPageRequest) Validate() error {
	return dara.Validate(s)
}
