// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServicesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGatewayId(v string) *ListServicesRequest
	GetGatewayId() *string
	SetName(v string) *ListServicesRequest
	GetName() *string
	SetPageNumber(v int32) *ListServicesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListServicesRequest
	GetPageSize() *int32
	SetResourceGroupId(v string) *ListServicesRequest
	GetResourceGroupId() *string
	SetSourceType(v string) *ListServicesRequest
	GetSourceType() *string
	SetSourceTypes(v string) *ListServicesRequest
	GetSourceTypes() *string
}

type ListServicesRequest struct {
	// The ID of the Cloud-native API Gateway instance.
	//
	// example:
	//
	// gw-cpv4sqdl*****
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// The service name.
	//
	// example:
	//
	// user-service
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The page number to return. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmxxe5rc6cvla
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// The service source. Valid values:
	//
	// 	- MSE_NACOS: a service in an MSE Nacos instance
	//
	// 	- K8S: a service in a Kubernetes (K8s) cluster in Container Service for Kubernetes (ACK)
	//
	// 	- FC3: a service in Function Compute
	//
	// 	- VIP: a fixed address
	//
	// 	- DNS: a domain name
	//
	// Enumerated values:
	//
	// 	- K8S
	//
	// 	- FC3
	//
	// 	- DNS
	//
	// 	- VIP
	//
	// 	- MSE_NACOS
	//
	// example:
	//
	// MSE_NACOS
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	// List of service source types
	//
	// example:
	//
	// ["K8S", "FC3"]
	SourceTypes *string `json:"sourceTypes,omitempty" xml:"sourceTypes,omitempty"`
}

func (s ListServicesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListServicesRequest) GoString() string {
	return s.String()
}

func (s *ListServicesRequest) GetGatewayId() *string {
	return s.GatewayId
}

func (s *ListServicesRequest) GetName() *string {
	return s.Name
}

func (s *ListServicesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListServicesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListServicesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListServicesRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *ListServicesRequest) GetSourceTypes() *string {
	return s.SourceTypes
}

func (s *ListServicesRequest) SetGatewayId(v string) *ListServicesRequest {
	s.GatewayId = &v
	return s
}

func (s *ListServicesRequest) SetName(v string) *ListServicesRequest {
	s.Name = &v
	return s
}

func (s *ListServicesRequest) SetPageNumber(v int32) *ListServicesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListServicesRequest) SetPageSize(v int32) *ListServicesRequest {
	s.PageSize = &v
	return s
}

func (s *ListServicesRequest) SetResourceGroupId(v string) *ListServicesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListServicesRequest) SetSourceType(v string) *ListServicesRequest {
	s.SourceType = &v
	return s
}

func (s *ListServicesRequest) SetSourceTypes(v string) *ListServicesRequest {
	s.SourceTypes = &v
	return s
}

func (s *ListServicesRequest) Validate() error {
	return dara.Validate(s)
}
