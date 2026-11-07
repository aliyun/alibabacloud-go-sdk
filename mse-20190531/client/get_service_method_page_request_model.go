// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceMethodPageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *GetServiceMethodPageRequest
	GetAcceptLanguage() *string
	SetAppId(v string) *GetServiceMethodPageRequest
	GetAppId() *string
	SetAppName(v string) *GetServiceMethodPageRequest
	GetAppName() *string
	SetIp(v string) *GetServiceMethodPageRequest
	GetIp() *string
	SetMethodController(v string) *GetServiceMethodPageRequest
	GetMethodController() *string
	SetName(v string) *GetServiceMethodPageRequest
	GetName() *string
	SetNamespace(v string) *GetServiceMethodPageRequest
	GetNamespace() *string
	SetPageNumber(v int32) *GetServiceMethodPageRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *GetServiceMethodPageRequest
	GetPageSize() *int32
	SetPath(v string) *GetServiceMethodPageRequest
	GetPath() *string
	SetRegion(v string) *GetServiceMethodPageRequest
	GetRegion() *string
	SetServiceGroup(v string) *GetServiceMethodPageRequest
	GetServiceGroup() *string
	SetServiceName(v string) *GetServiceMethodPageRequest
	GetServiceName() *string
	SetServiceType(v string) *GetServiceMethodPageRequest
	GetServiceType() *string
	SetServiceVersion(v string) *GetServiceMethodPageRequest
	GetServiceVersion() *string
}

type GetServiceMethodPageRequest struct {
	// The language in which you want the results to be displayed. Valid values: zh and en. zh indicates Chinese, which is the default value. en indicates English.
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The application ID.
	//
	// example:
	//
	// hkhonxxxxx@f3f75ed8ffxxxxx
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The application name.
	//
	// example:
	//
	// example-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The IP address of the service provider.
	//
	// example:
	//
	// 223.5.5.5
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The full name of the Controller class to which the interface belongs.
	//
	// example:
	//
	// com.alibabacloud.mse.demo.a.AController
	MethodController *string `json:"MethodController,omitempty" xml:"MethodController,omitempty"`
	// The method name.
	//
	// example:
	//
	// aMethod
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The Microservices Engine (MSE) namespace to which the application belongs.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request path that corresponds to the interface.
	//
	// example:
	//
	// /a
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The region.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The group to which the service belongs.
	//
	// example:
	//
	// ALIYUN
	ServiceGroup *string `json:"ServiceGroup,omitempty" xml:"ServiceGroup,omitempty"`
	// The name of the service.
	//
	// example:
	//
	// sc-A
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The engine type.
	//
	// example:
	//
	// springCloud
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// The service version.
	//
	// example:
	//
	// 1.0.0
	ServiceVersion *string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
}

func (s GetServiceMethodPageRequest) String() string {
	return dara.Prettify(s)
}

func (s GetServiceMethodPageRequest) GoString() string {
	return s.String()
}

func (s *GetServiceMethodPageRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *GetServiceMethodPageRequest) GetAppId() *string {
	return s.AppId
}

func (s *GetServiceMethodPageRequest) GetAppName() *string {
	return s.AppName
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

func (s *GetServiceMethodPageRequest) GetServiceGroup() *string {
	return s.ServiceGroup
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

func (s *GetServiceMethodPageRequest) SetAcceptLanguage(v string) *GetServiceMethodPageRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *GetServiceMethodPageRequest) SetAppId(v string) *GetServiceMethodPageRequest {
	s.AppId = &v
	return s
}

func (s *GetServiceMethodPageRequest) SetAppName(v string) *GetServiceMethodPageRequest {
	s.AppName = &v
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

func (s *GetServiceMethodPageRequest) SetServiceGroup(v string) *GetServiceMethodPageRequest {
	s.ServiceGroup = &v
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

func (s *GetServiceMethodPageRequest) Validate() error {
	return dara.Validate(s)
}
