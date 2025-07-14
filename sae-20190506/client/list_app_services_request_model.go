// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppServicesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ListAppServicesRequest
	GetAppId() *string
	SetNacosInstanceId(v string) *ListAppServicesRequest
	GetNacosInstanceId() *string
	SetNacosNamespaceId(v string) *ListAppServicesRequest
	GetNacosNamespaceId() *string
	SetNamespaceId(v string) *ListAppServicesRequest
	GetNamespaceId() *string
	SetPageNumber(v int32) *ListAppServicesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListAppServicesRequest
	GetPageSize() *int32
	SetRegistryType(v string) *ListAppServicesRequest
	GetRegistryType() *string
	SetServiceType(v string) *ListAppServicesRequest
	GetServiceType() *string
	SetVpcId(v string) *ListAppServicesRequest
	GetVpcId() *string
}

type ListAppServicesRequest struct {
	// The ID of the application. You must specify only one of the following parameters: vpcId, namespace ID, and application ID.
	//
	// example:
	//
	// 017f39b8-dfa4-4e16-a84b-1dcee4b1****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the MSE Nacos instance. This parameter is required when the registry type is set to MSE Nacos.
	//
	// example:
	//
	// mse-cn-sco3r0u****
	NacosInstanceId *string `json:"NacosInstanceId,omitempty" xml:"NacosInstanceId,omitempty"`
	// The ID of the MSE Nacos namespace. This parameter is required when the registry type is set to MSE Nacos.
	//
	// example:
	//
	// mse-test
	NacosNamespaceId *string `json:"NacosNamespaceId,omitempty" xml:"NacosNamespaceId,omitempty"`
	// The ID of the namespace. You must specify only one of the following parameters: VPC ID, namespace ID, and application ID.
	//
	// example:
	//
	// cn-beijing:test
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
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
	// The registry type. Valid values:
	//
	// 	- **0**: SAE Nacos
	//
	// 	- **1**: SAE built-in Nacos
	//
	// 	- **2*	- :MSE Nacos
	//
	// 	- **9**: SAE Kubernetes service
	//
	// example:
	//
	// 0
	RegistryType *string `json:"RegistryType,omitempty" xml:"RegistryType,omitempty"`
	// The service type. Valid values:
	//
	// 	- **dubbo**
	//
	// 	- **springCloud**
	//
	// 	- **hsf**
	//
	// 	- **k8sService**
	//
	// example:
	//
	// springCloud
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// The unique identifier of the VPC. You must specify only one of the following parameters: VPC ID, namespace ID, and application ID.
	//
	// example:
	//
	// vpc-2ze0i263cnn311nvj****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListAppServicesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAppServicesRequest) GoString() string {
	return s.String()
}

func (s *ListAppServicesRequest) GetAppId() *string {
	return s.AppId
}

func (s *ListAppServicesRequest) GetNacosInstanceId() *string {
	return s.NacosInstanceId
}

func (s *ListAppServicesRequest) GetNacosNamespaceId() *string {
	return s.NacosNamespaceId
}

func (s *ListAppServicesRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *ListAppServicesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAppServicesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAppServicesRequest) GetRegistryType() *string {
	return s.RegistryType
}

func (s *ListAppServicesRequest) GetServiceType() *string {
	return s.ServiceType
}

func (s *ListAppServicesRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *ListAppServicesRequest) SetAppId(v string) *ListAppServicesRequest {
	s.AppId = &v
	return s
}

func (s *ListAppServicesRequest) SetNacosInstanceId(v string) *ListAppServicesRequest {
	s.NacosInstanceId = &v
	return s
}

func (s *ListAppServicesRequest) SetNacosNamespaceId(v string) *ListAppServicesRequest {
	s.NacosNamespaceId = &v
	return s
}

func (s *ListAppServicesRequest) SetNamespaceId(v string) *ListAppServicesRequest {
	s.NamespaceId = &v
	return s
}

func (s *ListAppServicesRequest) SetPageNumber(v int32) *ListAppServicesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAppServicesRequest) SetPageSize(v int32) *ListAppServicesRequest {
	s.PageSize = &v
	return s
}

func (s *ListAppServicesRequest) SetRegistryType(v string) *ListAppServicesRequest {
	s.RegistryType = &v
	return s
}

func (s *ListAppServicesRequest) SetServiceType(v string) *ListAppServicesRequest {
	s.ServiceType = &v
	return s
}

func (s *ListAppServicesRequest) SetVpcId(v string) *ListAppServicesRequest {
	s.VpcId = &v
	return s
}

func (s *ListAppServicesRequest) Validate() error {
	return dara.Validate(s)
}
