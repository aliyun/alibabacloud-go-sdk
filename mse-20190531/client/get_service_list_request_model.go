// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *GetServiceListRequest
	GetAcceptLanguage() *string
	SetAppId(v string) *GetServiceListRequest
	GetAppId() *string
	SetIp(v string) *GetServiceListRequest
	GetIp() *string
	SetRegion(v string) *GetServiceListRequest
	GetRegion() *string
	SetServiceName(v string) *GetServiceListRequest
	GetServiceName() *string
	SetServiceType(v string) *GetServiceListRequest
	GetServiceType() *string
}

type GetServiceListRequest struct {
	// The language of the response.
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// xxx@xxx
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The IP address.
	//
	// example:
	//
	// 127.0.0.1
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-zhangjiakou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The name of the service.
	//
	// example:
	//
	// com.alibaba.xxx
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The type of the framework.
	//
	// example:
	//
	// dubbo
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
}

func (s GetServiceListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetServiceListRequest) GoString() string {
	return s.String()
}

func (s *GetServiceListRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *GetServiceListRequest) GetAppId() *string {
	return s.AppId
}

func (s *GetServiceListRequest) GetIp() *string {
	return s.Ip
}

func (s *GetServiceListRequest) GetRegion() *string {
	return s.Region
}

func (s *GetServiceListRequest) GetServiceName() *string {
	return s.ServiceName
}

func (s *GetServiceListRequest) GetServiceType() *string {
	return s.ServiceType
}

func (s *GetServiceListRequest) SetAcceptLanguage(v string) *GetServiceListRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *GetServiceListRequest) SetAppId(v string) *GetServiceListRequest {
	s.AppId = &v
	return s
}

func (s *GetServiceListRequest) SetIp(v string) *GetServiceListRequest {
	s.Ip = &v
	return s
}

func (s *GetServiceListRequest) SetRegion(v string) *GetServiceListRequest {
	s.Region = &v
	return s
}

func (s *GetServiceListRequest) SetServiceName(v string) *GetServiceListRequest {
	s.ServiceName = &v
	return s
}

func (s *GetServiceListRequest) SetServiceType(v string) *GetServiceListRequest {
	s.ServiceType = &v
	return s
}

func (s *GetServiceListRequest) Validate() error {
	return dara.Validate(s)
}
