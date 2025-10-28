// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceMethodPageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *GetServiceMethodPageRequest
	GetAppId() *string
	SetGroup(v string) *GetServiceMethodPageRequest
	GetGroup() *string
	SetIp(v string) *GetServiceMethodPageRequest
	GetIp() *string
	SetMethodController(v string) *GetServiceMethodPageRequest
	GetMethodController() *string
	SetName(v string) *GetServiceMethodPageRequest
	GetName() *string
	SetNamespace(v string) *GetServiceMethodPageRequest
	GetNamespace() *string
	SetOrigin(v string) *GetServiceMethodPageRequest
	GetOrigin() *string
	SetPageNumber(v int32) *GetServiceMethodPageRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *GetServiceMethodPageRequest
	GetPageSize() *int32
	SetPath(v string) *GetServiceMethodPageRequest
	GetPath() *string
	SetRegion(v string) *GetServiceMethodPageRequest
	GetRegion() *string
	SetRegistryType(v string) *GetServiceMethodPageRequest
	GetRegistryType() *string
	SetServiceId(v string) *GetServiceMethodPageRequest
	GetServiceId() *string
	SetServiceName(v string) *GetServiceMethodPageRequest
	GetServiceName() *string
	SetServiceType(v string) *GetServiceMethodPageRequest
	GetServiceType() *string
	SetServiceVersion(v string) *GetServiceMethodPageRequest
	GetServiceVersion() *string
	SetSource(v string) *GetServiceMethodPageRequest
	GetSource() *string
}

type GetServiceMethodPageRequest struct {
	// The ID of the application.
	//
	// example:
	//
	// 310b18c3-1dbe-4807-****-18d7d637****
	AppId *string `json:"appId,omitempty" xml:"appId,omitempty"`
	// The group to which the service belongs.
	//
	// example:
	//
	// DUBBO
	Group *string `json:"group,omitempty" xml:"group,omitempty"`
	// The IP address of the service.
	//
	// example:
	//
	// 127.0.0.1
	Ip *string `json:"ip,omitempty" xml:"ip,omitempty"`
	// The controller method.
	//
	// example:
	//
	// com.aliware.edas.DemoController
	MethodController *string `json:"methodController,omitempty" xml:"methodController,omitempty"`
	// The name of the method.
	//
	// example:
	//
	// echo
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
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
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The request path.
	//
	// example:
	//
	// /echo/{str}
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// The type of the service registry.
	//
	// example:
	//
	// nacos
	RegistryType *string `json:"registryType,omitempty" xml:"registryType,omitempty"`
	// The ID of the service.
	//
	// example:
	//
	// edas.service.consumer
	ServiceId *string `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
	// The name of the service.
	//
	// example:
	//
	// edas.service.consumer
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

func (s GetServiceMethodPageRequest) String() string {
	return dara.Prettify(s)
}

func (s GetServiceMethodPageRequest) GoString() string {
	return s.String()
}

func (s *GetServiceMethodPageRequest) GetAppId() *string {
	return s.AppId
}

func (s *GetServiceMethodPageRequest) GetGroup() *string {
	return s.Group
}

func (s *GetServiceMethodPageRequest) GetIp() *string {
	return s.Ip
}

func (s *GetServiceMethodPageRequest) GetMethodController() *string {
	return s.MethodController
}

func (s *GetServiceMethodPageRequest) GetName() *string {
	return s.Name
}

func (s *GetServiceMethodPageRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *GetServiceMethodPageRequest) GetOrigin() *string {
	return s.Origin
}

func (s *GetServiceMethodPageRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *GetServiceMethodPageRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetServiceMethodPageRequest) GetPath() *string {
	return s.Path
}

func (s *GetServiceMethodPageRequest) GetRegion() *string {
	return s.Region
}

func (s *GetServiceMethodPageRequest) GetRegistryType() *string {
	return s.RegistryType
}

func (s *GetServiceMethodPageRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *GetServiceMethodPageRequest) GetServiceName() *string {
	return s.ServiceName
}

func (s *GetServiceMethodPageRequest) GetServiceType() *string {
	return s.ServiceType
}

func (s *GetServiceMethodPageRequest) GetServiceVersion() *string {
	return s.ServiceVersion
}

func (s *GetServiceMethodPageRequest) GetSource() *string {
	return s.Source
}

func (s *GetServiceMethodPageRequest) SetAppId(v string) *GetServiceMethodPageRequest {
	s.AppId = &v
	return s
}

func (s *GetServiceMethodPageRequest) SetGroup(v string) *GetServiceMethodPageRequest {
	s.Group = &v
	return s
}

func (s *GetServiceMethodPageRequest) SetIp(v string) *GetServiceMethodPageRequest {
	s.Ip = &v
	return s
}

func (s *GetServiceMethodPageRequest) SetMethodController(v string) *GetServiceMethodPageRequest {
	s.MethodController = &v
	return s
}

func (s *GetServiceMethodPageRequest) SetName(v string) *GetServiceMethodPageRequest {
	s.Name = &v
	return s
}

func (s *GetServiceMethodPageRequest) SetNamespace(v string) *GetServiceMethodPageRequest {
	s.Namespace = &v
	return s
}

func (s *GetServiceMethodPageRequest) SetOrigin(v string) *GetServiceMethodPageRequest {
	s.Origin = &v
	return s
}

func (s *GetServiceMethodPageRequest) SetPageNumber(v int32) *GetServiceMethodPageRequest {
	s.PageNumber = &v
	return s
}

func (s *GetServiceMethodPageRequest) SetPageSize(v int32) *GetServiceMethodPageRequest {
	s.PageSize = &v
	return s
}

func (s *GetServiceMethodPageRequest) SetPath(v string) *GetServiceMethodPageRequest {
	s.Path = &v
	return s
}

func (s *GetServiceMethodPageRequest) SetRegion(v string) *GetServiceMethodPageRequest {
	s.Region = &v
	return s
}

func (s *GetServiceMethodPageRequest) SetRegistryType(v string) *GetServiceMethodPageRequest {
	s.RegistryType = &v
	return s
}

func (s *GetServiceMethodPageRequest) SetServiceId(v string) *GetServiceMethodPageRequest {
	s.ServiceId = &v
	return s
}

func (s *GetServiceMethodPageRequest) SetServiceName(v string) *GetServiceMethodPageRequest {
	s.ServiceName = &v
	return s
}

func (s *GetServiceMethodPageRequest) SetServiceType(v string) *GetServiceMethodPageRequest {
	s.ServiceType = &v
	return s
}

func (s *GetServiceMethodPageRequest) SetServiceVersion(v string) *GetServiceMethodPageRequest {
	s.ServiceVersion = &v
	return s
}

func (s *GetServiceMethodPageRequest) SetSource(v string) *GetServiceMethodPageRequest {
	s.Source = &v
	return s
}

func (s *GetServiceMethodPageRequest) Validate() error {
	return dara.Validate(s)
}
