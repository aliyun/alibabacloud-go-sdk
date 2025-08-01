// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceListPageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *GetServiceListPageRequest
	GetAcceptLanguage() *string
	SetAppId(v string) *GetServiceListPageRequest
	GetAppId() *string
	SetAppName(v string) *GetServiceListPageRequest
	GetAppName() *string
	SetIp(v string) *GetServiceListPageRequest
	GetIp() *string
	SetNamespace(v string) *GetServiceListPageRequest
	GetNamespace() *string
	SetPageNumber(v int32) *GetServiceListPageRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *GetServiceListPageRequest
	GetPageSize() *int32
	SetRegion(v string) *GetServiceListPageRequest
	GetRegion() *string
	SetServiceName(v string) *GetServiceListPageRequest
	GetServiceName() *string
	SetServiceType(v string) *GetServiceListPageRequest
	GetServiceType() *string
}

type GetServiceListPageRequest struct {
	// The language of the response. Valid values: zh and en. Default value: zh. The value zh indicates Chinese, and the value en indicates English.
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The application ID.
	//
	// example:
	//
	// dez4xxxxx@f3f75ed8ffxxxxx
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The application name.
	//
	// example:
	//
	// example-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The IP address from which the query is initiated.
	//
	// example:
	//
	// 223.5.5.5
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The name of the MSE namespace.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The page number.
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
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The service name.
	//
	// example:
	//
	// sc-A
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The service type.
	//
	// example:
	//
	// springCloud
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
}

func (s GetServiceListPageRequest) String() string {
	return dara.Prettify(s)
}

func (s GetServiceListPageRequest) GoString() string {
	return s.String()
}

func (s *GetServiceListPageRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *GetServiceListPageRequest) GetAppId() *string {
	return s.AppId
}

func (s *GetServiceListPageRequest) GetAppName() *string {
	return s.AppName
}

func (s *GetServiceListPageRequest) GetIp() *string {
	return s.Ip
}

func (s *GetServiceListPageRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *GetServiceListPageRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *GetServiceListPageRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetServiceListPageRequest) GetRegion() *string {
	return s.Region
}

func (s *GetServiceListPageRequest) GetServiceName() *string {
	return s.ServiceName
}

func (s *GetServiceListPageRequest) GetServiceType() *string {
	return s.ServiceType
}

func (s *GetServiceListPageRequest) SetAcceptLanguage(v string) *GetServiceListPageRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *GetServiceListPageRequest) SetAppId(v string) *GetServiceListPageRequest {
	s.AppId = &v
	return s
}

func (s *GetServiceListPageRequest) SetAppName(v string) *GetServiceListPageRequest {
	s.AppName = &v
	return s
}

func (s *GetServiceListPageRequest) SetIp(v string) *GetServiceListPageRequest {
	s.Ip = &v
	return s
}

func (s *GetServiceListPageRequest) SetNamespace(v string) *GetServiceListPageRequest {
	s.Namespace = &v
	return s
}

func (s *GetServiceListPageRequest) SetPageNumber(v int32) *GetServiceListPageRequest {
	s.PageNumber = &v
	return s
}

func (s *GetServiceListPageRequest) SetPageSize(v int32) *GetServiceListPageRequest {
	s.PageSize = &v
	return s
}

func (s *GetServiceListPageRequest) SetRegion(v string) *GetServiceListPageRequest {
	s.Region = &v
	return s
}

func (s *GetServiceListPageRequest) SetServiceName(v string) *GetServiceListPageRequest {
	s.ServiceName = &v
	return s
}

func (s *GetServiceListPageRequest) SetServiceType(v string) *GetServiceListPageRequest {
	s.ServiceType = &v
	return s
}

func (s *GetServiceListPageRequest) Validate() error {
	return dara.Validate(s)
}
