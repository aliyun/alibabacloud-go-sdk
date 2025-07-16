// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGatewayRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGatewayId(v string) *ListGatewayRequest
	GetGatewayId() *string
	SetGatewayName(v string) *ListGatewayRequest
	GetGatewayName() *string
	SetPageNumber(v int32) *ListGatewayRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListGatewayRequest
	GetPageSize() *int32
	SetResourceName(v string) *ListGatewayRequest
	GetResourceName() *string
}

type ListGatewayRequest struct {
	// The private gateway ID. To obtain the private gateway ID, see the private_gateway_id parameter in the response parameters of the ListResources operation.
	//
	// example:
	//
	// gw-1uhcqmsc7x22******
	GatewayId *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// The private gateway alias.
	//
	// example:
	//
	// mygateway1
	GatewayName *string `json:"GatewayName,omitempty" xml:"GatewayName,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 100.
	//
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the resource group. To obtain a resource group ID, see the ResourceId field in the response of the [ListResources](https://help.aliyun.com/document_detail/412133.html) operation.
	//
	// example:
	//
	// eas-r-4gt8twzwllfo******
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
}

func (s ListGatewayRequest) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayRequest) GoString() string {
	return s.String()
}

func (s *ListGatewayRequest) GetGatewayId() *string {
	return s.GatewayId
}

func (s *ListGatewayRequest) GetGatewayName() *string {
	return s.GatewayName
}

func (s *ListGatewayRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListGatewayRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListGatewayRequest) GetResourceName() *string {
	return s.ResourceName
}

func (s *ListGatewayRequest) SetGatewayId(v string) *ListGatewayRequest {
	s.GatewayId = &v
	return s
}

func (s *ListGatewayRequest) SetGatewayName(v string) *ListGatewayRequest {
	s.GatewayName = &v
	return s
}

func (s *ListGatewayRequest) SetPageNumber(v int32) *ListGatewayRequest {
	s.PageNumber = &v
	return s
}

func (s *ListGatewayRequest) SetPageSize(v int32) *ListGatewayRequest {
	s.PageSize = &v
	return s
}

func (s *ListGatewayRequest) SetResourceName(v string) *ListGatewayRequest {
	s.ResourceName = &v
	return s
}

func (s *ListGatewayRequest) Validate() error {
	return dara.Validate(s)
}
