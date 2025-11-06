// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGatewayServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ListGatewayServiceRequest
	GetAcceptLanguage() *string
	SetDescSort(v bool) *ListGatewayServiceRequest
	GetDescSort() *bool
	SetFilterParams(v *ListGatewayServiceRequestFilterParams) *ListGatewayServiceRequest
	GetFilterParams() *ListGatewayServiceRequestFilterParams
	SetOrderItem(v string) *ListGatewayServiceRequest
	GetOrderItem() *string
	SetPageNumber(v int32) *ListGatewayServiceRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListGatewayServiceRequest
	GetPageSize() *int32
}

type ListGatewayServiceRequest struct {
	// The language of the response. Valid values:
	//
	// 	- zh: Chinese
	//
	// 	- en: English
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// Specifies whether to enable sorting.
	//
	// example:
	//
	// false
	DescSort *bool `json:"DescSort,omitempty" xml:"DescSort,omitempty"`
	// The parameters that are used to specify filter conditions. The values of the parameters are in the format of {"key1":"value1"}.
	FilterParams *ListGatewayServiceRequestFilterParams `json:"FilterParams,omitempty" xml:"FilterParams,omitempty" type:"Struct"`
	// The item based on which entries are sorted.
	//
	// example:
	//
	// GmtCreate
	OrderItem *string `json:"OrderItem,omitempty" xml:"OrderItem,omitempty"`
	// The number of the page to return.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListGatewayServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayServiceRequest) GoString() string {
	return s.String()
}

func (s *ListGatewayServiceRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ListGatewayServiceRequest) GetDescSort() *bool {
	return s.DescSort
}

func (s *ListGatewayServiceRequest) GetFilterParams() *ListGatewayServiceRequestFilterParams {
	return s.FilterParams
}

func (s *ListGatewayServiceRequest) GetOrderItem() *string {
	return s.OrderItem
}

func (s *ListGatewayServiceRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListGatewayServiceRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListGatewayServiceRequest) SetAcceptLanguage(v string) *ListGatewayServiceRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ListGatewayServiceRequest) SetDescSort(v bool) *ListGatewayServiceRequest {
	s.DescSort = &v
	return s
}

func (s *ListGatewayServiceRequest) SetFilterParams(v *ListGatewayServiceRequestFilterParams) *ListGatewayServiceRequest {
	s.FilterParams = v
	return s
}

func (s *ListGatewayServiceRequest) SetOrderItem(v string) *ListGatewayServiceRequest {
	s.OrderItem = &v
	return s
}

func (s *ListGatewayServiceRequest) SetPageNumber(v int32) *ListGatewayServiceRequest {
	s.PageNumber = &v
	return s
}

func (s *ListGatewayServiceRequest) SetPageSize(v int32) *ListGatewayServiceRequest {
	s.PageSize = &v
	return s
}

func (s *ListGatewayServiceRequest) Validate() error {
	if s.FilterParams != nil {
		if err := s.FilterParams.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListGatewayServiceRequestFilterParams struct {
	// The unique ID of the gateway.
	//
	// example:
	//
	// gw-5017305290e14centbrveca****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// The group.
	//
	// example:
	//
	// test
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The name of the service.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The namespace to which the service belongs.
	//
	// example:
	//
	// public
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The protocol of the service.
	//
	// 	- HTTP
	//
	// 	- HTTPS
	//
	// 	- HTTP2
	//
	// 	- GRPC
	//
	// 	- DUBBO
	//
	// example:
	//
	// HTTP
	ServiceProtocol *string `json:"ServiceProtocol,omitempty" xml:"ServiceProtocol,omitempty"`
	// The type of the source.
	//
	// example:
	//
	// MSE
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
}

func (s ListGatewayServiceRequestFilterParams) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayServiceRequestFilterParams) GoString() string {
	return s.String()
}

func (s *ListGatewayServiceRequestFilterParams) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *ListGatewayServiceRequestFilterParams) GetGroupName() *string {
	return s.GroupName
}

func (s *ListGatewayServiceRequestFilterParams) GetName() *string {
	return s.Name
}

func (s *ListGatewayServiceRequestFilterParams) GetNamespace() *string {
	return s.Namespace
}

func (s *ListGatewayServiceRequestFilterParams) GetServiceProtocol() *string {
	return s.ServiceProtocol
}

func (s *ListGatewayServiceRequestFilterParams) GetSourceType() *string {
	return s.SourceType
}

func (s *ListGatewayServiceRequestFilterParams) SetGatewayUniqueId(v string) *ListGatewayServiceRequestFilterParams {
	s.GatewayUniqueId = &v
	return s
}

func (s *ListGatewayServiceRequestFilterParams) SetGroupName(v string) *ListGatewayServiceRequestFilterParams {
	s.GroupName = &v
	return s
}

func (s *ListGatewayServiceRequestFilterParams) SetName(v string) *ListGatewayServiceRequestFilterParams {
	s.Name = &v
	return s
}

func (s *ListGatewayServiceRequestFilterParams) SetNamespace(v string) *ListGatewayServiceRequestFilterParams {
	s.Namespace = &v
	return s
}

func (s *ListGatewayServiceRequestFilterParams) SetServiceProtocol(v string) *ListGatewayServiceRequestFilterParams {
	s.ServiceProtocol = &v
	return s
}

func (s *ListGatewayServiceRequestFilterParams) SetSourceType(v string) *ListGatewayServiceRequestFilterParams {
	s.SourceType = &v
	return s
}

func (s *ListGatewayServiceRequestFilterParams) Validate() error {
	return dara.Validate(s)
}
