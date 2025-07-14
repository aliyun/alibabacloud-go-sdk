// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAppServiceDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeAppServiceDetailRequest
	GetAppId() *string
	SetNacosInstanceId(v string) *DescribeAppServiceDetailRequest
	GetNacosInstanceId() *string
	SetNacosNamespaceId(v string) *DescribeAppServiceDetailRequest
	GetNacosNamespaceId() *string
	SetServiceGroup(v string) *DescribeAppServiceDetailRequest
	GetServiceGroup() *string
	SetServiceName(v string) *DescribeAppServiceDetailRequest
	GetServiceName() *string
	SetServiceType(v string) *DescribeAppServiceDetailRequest
	GetServiceType() *string
	SetServiceVersion(v string) *DescribeAppServiceDetailRequest
	GetServiceVersion() *string
}

type DescribeAppServiceDetailRequest struct {
	// 6dcc8c9e-d3da-478a-a066-86dcf820\\*\\*\\*\\*
	//
	// This parameter is required.
	//
	// example:
	//
	// 6dcc8c9e-d3da-478a-a066-86dcf820****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the MSE Nacos instance.
	//
	// example:
	//
	// mse-cn-sco3r0u****
	NacosInstanceId *string `json:"NacosInstanceId,omitempty" xml:"NacosInstanceId,omitempty"`
	// The ID of the namespace for the MSE Nacos instance.
	//
	// example:
	//
	// public
	NacosNamespaceId *string `json:"NacosNamespaceId,omitempty" xml:"NacosNamespaceId,omitempty"`
	// springCloud
	//
	// example:
	//
	// springCloud
	ServiceGroup *string `json:"ServiceGroup,omitempty" xml:"ServiceGroup,omitempty"`
	// edas.service.provider
	//
	// example:
	//
	// edas.service.provider
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// springCloud
	//
	// example:
	//
	// springCloud
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// 1.0.0
	//
	// example:
	//
	// 1.0.0
	ServiceVersion *string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
}

func (s DescribeAppServiceDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppServiceDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeAppServiceDetailRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeAppServiceDetailRequest) GetNacosInstanceId() *string {
	return s.NacosInstanceId
}

func (s *DescribeAppServiceDetailRequest) GetNacosNamespaceId() *string {
	return s.NacosNamespaceId
}

func (s *DescribeAppServiceDetailRequest) GetServiceGroup() *string {
	return s.ServiceGroup
}

func (s *DescribeAppServiceDetailRequest) GetServiceName() *string {
	return s.ServiceName
}

func (s *DescribeAppServiceDetailRequest) GetServiceType() *string {
	return s.ServiceType
}

func (s *DescribeAppServiceDetailRequest) GetServiceVersion() *string {
	return s.ServiceVersion
}

func (s *DescribeAppServiceDetailRequest) SetAppId(v string) *DescribeAppServiceDetailRequest {
	s.AppId = &v
	return s
}

func (s *DescribeAppServiceDetailRequest) SetNacosInstanceId(v string) *DescribeAppServiceDetailRequest {
	s.NacosInstanceId = &v
	return s
}

func (s *DescribeAppServiceDetailRequest) SetNacosNamespaceId(v string) *DescribeAppServiceDetailRequest {
	s.NacosNamespaceId = &v
	return s
}

func (s *DescribeAppServiceDetailRequest) SetServiceGroup(v string) *DescribeAppServiceDetailRequest {
	s.ServiceGroup = &v
	return s
}

func (s *DescribeAppServiceDetailRequest) SetServiceName(v string) *DescribeAppServiceDetailRequest {
	s.ServiceName = &v
	return s
}

func (s *DescribeAppServiceDetailRequest) SetServiceType(v string) *DescribeAppServiceDetailRequest {
	s.ServiceType = &v
	return s
}

func (s *DescribeAppServiceDetailRequest) SetServiceVersion(v string) *DescribeAppServiceDetailRequest {
	s.ServiceVersion = &v
	return s
}

func (s *DescribeAppServiceDetailRequest) Validate() error {
	return dara.Validate(s)
}
