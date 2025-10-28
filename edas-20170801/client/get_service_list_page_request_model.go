// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceListPageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNamespace(v string) *GetServiceListPageRequest
	GetNamespace() *string
	SetOrigin(v string) *GetServiceListPageRequest
	GetOrigin() *string
	SetPage(v int32) *GetServiceListPageRequest
	GetPage() *int32
	SetRegion(v string) *GetServiceListPageRequest
	GetRegion() *string
	SetSearchType(v string) *GetServiceListPageRequest
	GetSearchType() *string
	SetSearchValue(v string) *GetServiceListPageRequest
	GetSearchValue() *string
	SetServiceType(v string) *GetServiceListPageRequest
	GetServiceType() *string
	SetSide(v string) *GetServiceListPageRequest
	GetSide() *string
	SetSize(v int32) *GetServiceListPageRequest
	GetSize() *int32
}

type GetServiceListPageRequest struct {
	// The namespace.
	//
	// example:
	//
	// cn-hangzhou:doc-test
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// The source of the data. Valid values:
	//
	// 	- `agent`: Use this value if you use the service query feature of the latest version to pass the query result.
	//
	// 	- `registry`: Use this value if you use the service query feature of the earlier version to pass the query result.
	//
	// example:
	//
	// Agent
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
	// The type of the service. Valid values:
	//
	// 	- `app`: searches by application.
	//
	// 	- `service`: searches by service.
	//
	// 	- `providerIp`: searches by IP address.
	//
	// example:
	//
	// App
	SearchType *string `json:"searchType,omitempty" xml:"searchType,omitempty"`
	// The keyword used for the search.
	//
	// 	- Set this parameter to the ID of the application if you set the searchType parameter to app.``
	//
	// 	- Set this parameter to the name of the service if you set the serachType parameter to service.``
	//
	// 	- Set this parameter to the IP address of the application if you set the searchType parameter to providerIp.
	//
	// example:
	//
	// com.alibaba.edas.HelloService
	SearchValue *string `json:"searchValue,omitempty" xml:"searchValue,omitempty"`
	// The type of the service. Valid values:
	//
	// 	- `dubbo`
	//
	// 	- `springCloud`
	//
	// 	- `hsf`
	//
	// 	- `istio`
	//
	// example:
	//
	// SpringCloud
	ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
	// Specifies the provider side or the consumer side. Valid values:
	//
	// 	- provider
	//
	// 	- consumer
	//
	// example:
	//
	// provider
	Side *string `json:"side,omitempty" xml:"side,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 10
	Size *int32 `json:"size,omitempty" xml:"size,omitempty"`
}

func (s GetServiceListPageRequest) String() string {
	return dara.Prettify(s)
}

func (s GetServiceListPageRequest) GoString() string {
	return s.String()
}

func (s *GetServiceListPageRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *GetServiceListPageRequest) GetOrigin() *string {
	return s.Origin
}

func (s *GetServiceListPageRequest) GetPage() *int32 {
	return s.Page
}

func (s *GetServiceListPageRequest) GetRegion() *string {
	return s.Region
}

func (s *GetServiceListPageRequest) GetSearchType() *string {
	return s.SearchType
}

func (s *GetServiceListPageRequest) GetSearchValue() *string {
	return s.SearchValue
}

func (s *GetServiceListPageRequest) GetServiceType() *string {
	return s.ServiceType
}

func (s *GetServiceListPageRequest) GetSide() *string {
	return s.Side
}

func (s *GetServiceListPageRequest) GetSize() *int32 {
	return s.Size
}

func (s *GetServiceListPageRequest) SetNamespace(v string) *GetServiceListPageRequest {
	s.Namespace = &v
	return s
}

func (s *GetServiceListPageRequest) SetOrigin(v string) *GetServiceListPageRequest {
	s.Origin = &v
	return s
}

func (s *GetServiceListPageRequest) SetPage(v int32) *GetServiceListPageRequest {
	s.Page = &v
	return s
}

func (s *GetServiceListPageRequest) SetRegion(v string) *GetServiceListPageRequest {
	s.Region = &v
	return s
}

func (s *GetServiceListPageRequest) SetSearchType(v string) *GetServiceListPageRequest {
	s.SearchType = &v
	return s
}

func (s *GetServiceListPageRequest) SetSearchValue(v string) *GetServiceListPageRequest {
	s.SearchValue = &v
	return s
}

func (s *GetServiceListPageRequest) SetServiceType(v string) *GetServiceListPageRequest {
	s.ServiceType = &v
	return s
}

func (s *GetServiceListPageRequest) SetSide(v string) *GetServiceListPageRequest {
	s.Side = &v
	return s
}

func (s *GetServiceListPageRequest) SetSize(v int32) *GetServiceListPageRequest {
	s.Size = &v
	return s
}

func (s *GetServiceListPageRequest) Validate() error {
	return dara.Validate(s)
}
