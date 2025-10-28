// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceConsumersPageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *GetServiceConsumersPageRequest
	GetAppId() *string
	SetGroup(v string) *GetServiceConsumersPageRequest
	GetGroup() *string
	SetIp(v string) *GetServiceConsumersPageRequest
	GetIp() *string
	SetNamespace(v string) *GetServiceConsumersPageRequest
	GetNamespace() *string
	SetOrigin(v string) *GetServiceConsumersPageRequest
	GetOrigin() *string
	SetPage(v int32) *GetServiceConsumersPageRequest
	GetPage() *int32
	SetRegion(v string) *GetServiceConsumersPageRequest
	GetRegion() *string
	SetRegistryType(v string) *GetServiceConsumersPageRequest
	GetRegistryType() *string
	SetServiceId(v string) *GetServiceConsumersPageRequest
	GetServiceId() *string
	SetServiceName(v string) *GetServiceConsumersPageRequest
	GetServiceName() *string
	SetServiceType(v string) *GetServiceConsumersPageRequest
	GetServiceType() *string
	SetServiceVersion(v string) *GetServiceConsumersPageRequest
	GetServiceVersion() *string
	SetSize(v int32) *GetServiceConsumersPageRequest
	GetSize() *int32
	SetSource(v string) *GetServiceConsumersPageRequest
	GetSource() *string
}

type GetServiceConsumersPageRequest struct {
	// The ID of the application.
	//
	// example:
	//
	// efbda488-7b33-432f-a40d-****0047****
	AppId *string `json:"appId,omitempty" xml:"appId,omitempty"`
	// The service group.
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
	// The number of the page to return.
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
	// The type of the service.
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
	// The number of entries returned per page.
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

func (s GetServiceConsumersPageRequest) String() string {
	return dara.Prettify(s)
}

func (s GetServiceConsumersPageRequest) GoString() string {
	return s.String()
}

func (s *GetServiceConsumersPageRequest) GetAppId() *string {
	return s.AppId
}

func (s *GetServiceConsumersPageRequest) GetGroup() *string {
	return s.Group
}

func (s *GetServiceConsumersPageRequest) GetIp() *string {
	return s.Ip
}

func (s *GetServiceConsumersPageRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *GetServiceConsumersPageRequest) GetOrigin() *string {
	return s.Origin
}

func (s *GetServiceConsumersPageRequest) GetPage() *int32 {
	return s.Page
}

func (s *GetServiceConsumersPageRequest) GetRegion() *string {
	return s.Region
}

func (s *GetServiceConsumersPageRequest) GetRegistryType() *string {
	return s.RegistryType
}

func (s *GetServiceConsumersPageRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *GetServiceConsumersPageRequest) GetServiceName() *string {
	return s.ServiceName
}

func (s *GetServiceConsumersPageRequest) GetServiceType() *string {
	return s.ServiceType
}

func (s *GetServiceConsumersPageRequest) GetServiceVersion() *string {
	return s.ServiceVersion
}

func (s *GetServiceConsumersPageRequest) GetSize() *int32 {
	return s.Size
}

func (s *GetServiceConsumersPageRequest) GetSource() *string {
	return s.Source
}

func (s *GetServiceConsumersPageRequest) SetAppId(v string) *GetServiceConsumersPageRequest {
	s.AppId = &v
	return s
}

func (s *GetServiceConsumersPageRequest) SetGroup(v string) *GetServiceConsumersPageRequest {
	s.Group = &v
	return s
}

func (s *GetServiceConsumersPageRequest) SetIp(v string) *GetServiceConsumersPageRequest {
	s.Ip = &v
	return s
}

func (s *GetServiceConsumersPageRequest) SetNamespace(v string) *GetServiceConsumersPageRequest {
	s.Namespace = &v
	return s
}

func (s *GetServiceConsumersPageRequest) SetOrigin(v string) *GetServiceConsumersPageRequest {
	s.Origin = &v
	return s
}

func (s *GetServiceConsumersPageRequest) SetPage(v int32) *GetServiceConsumersPageRequest {
	s.Page = &v
	return s
}

func (s *GetServiceConsumersPageRequest) SetRegion(v string) *GetServiceConsumersPageRequest {
	s.Region = &v
	return s
}

func (s *GetServiceConsumersPageRequest) SetRegistryType(v string) *GetServiceConsumersPageRequest {
	s.RegistryType = &v
	return s
}

func (s *GetServiceConsumersPageRequest) SetServiceId(v string) *GetServiceConsumersPageRequest {
	s.ServiceId = &v
	return s
}

func (s *GetServiceConsumersPageRequest) SetServiceName(v string) *GetServiceConsumersPageRequest {
	s.ServiceName = &v
	return s
}

func (s *GetServiceConsumersPageRequest) SetServiceType(v string) *GetServiceConsumersPageRequest {
	s.ServiceType = &v
	return s
}

func (s *GetServiceConsumersPageRequest) SetServiceVersion(v string) *GetServiceConsumersPageRequest {
	s.ServiceVersion = &v
	return s
}

func (s *GetServiceConsumersPageRequest) SetSize(v int32) *GetServiceConsumersPageRequest {
	s.Size = &v
	return s
}

func (s *GetServiceConsumersPageRequest) SetSource(v string) *GetServiceConsumersPageRequest {
	s.Source = &v
	return s
}

func (s *GetServiceConsumersPageRequest) Validate() error {
	return dara.Validate(s)
}
