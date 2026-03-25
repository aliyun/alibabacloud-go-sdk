// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSecretReferencesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListSecretReferencesResponseBody
	GetCode() *string
	SetData(v *ListSecretReferencesResponseBodyData) *ListSecretReferencesResponseBody
	GetData() *ListSecretReferencesResponseBodyData
	SetMessage(v string) *ListSecretReferencesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListSecretReferencesResponseBody
	GetRequestId() *string
}

type ListSecretReferencesResponseBody struct {
	// The status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned data.
	Data *ListSecretReferencesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// CE857A85-251D-5018-8103-A38957D71***
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListSecretReferencesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSecretReferencesResponseBody) GoString() string {
	return s.String()
}

func (s *ListSecretReferencesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListSecretReferencesResponseBody) GetData() *ListSecretReferencesResponseBodyData {
	return s.Data
}

func (s *ListSecretReferencesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListSecretReferencesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSecretReferencesResponseBody) SetCode(v string) *ListSecretReferencesResponseBody {
	s.Code = &v
	return s
}

func (s *ListSecretReferencesResponseBody) SetData(v *ListSecretReferencesResponseBodyData) *ListSecretReferencesResponseBody {
	s.Data = v
	return s
}

func (s *ListSecretReferencesResponseBody) SetMessage(v string) *ListSecretReferencesResponseBody {
	s.Message = &v
	return s
}

func (s *ListSecretReferencesResponseBody) SetRequestId(v string) *ListSecretReferencesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSecretReferencesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListSecretReferencesResponseBodyData struct {
	// The list of reference details.
	Items []*ListSecretReferencesResponseBodyDataItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 25
	TotalSize *int32 `json:"totalSize,omitempty" xml:"totalSize,omitempty"`
}

func (s ListSecretReferencesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListSecretReferencesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListSecretReferencesResponseBodyData) GetItems() []*ListSecretReferencesResponseBodyDataItems {
	return s.Items
}

func (s *ListSecretReferencesResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListSecretReferencesResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSecretReferencesResponseBodyData) GetTotalSize() *int32 {
	return s.TotalSize
}

func (s *ListSecretReferencesResponseBodyData) SetItems(v []*ListSecretReferencesResponseBodyDataItems) *ListSecretReferencesResponseBodyData {
	s.Items = v
	return s
}

func (s *ListSecretReferencesResponseBodyData) SetPageNumber(v int32) *ListSecretReferencesResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListSecretReferencesResponseBodyData) SetPageSize(v int32) *ListSecretReferencesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListSecretReferencesResponseBodyData) SetTotalSize(v int32) *ListSecretReferencesResponseBodyData {
	s.TotalSize = &v
	return s
}

func (s *ListSecretReferencesResponseBodyData) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSecretReferencesResponseBodyDataItems struct {
	// The information about the plug-in that references the current key.
	PluginConfig *ListSecretReferencesResponseBodyDataItemsPluginConfig `json:"PluginConfig,omitempty" xml:"PluginConfig,omitempty" type:"Struct"`
	// The service information that references the current key.
	ServiceConfig *ListSecretReferencesResponseBodyDataItemsServiceConfig `json:"ServiceConfig,omitempty" xml:"ServiceConfig,omitempty" type:"Struct"`
	// The consumer information that references the current key.
	ConsumerConfig *ListSecretReferencesResponseBodyDataItemsConsumerConfig `json:"consumerConfig,omitempty" xml:"consumerConfig,omitempty" type:"Struct"`
	// The gateway instance ID.
	//
	// example:
	//
	// gw-xxxx
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// MCP service information that references the current key.
	McpServerConfig *ListSecretReferencesResponseBodyDataItemsMcpServerConfig `json:"mcpServerConfig,omitempty" xml:"mcpServerConfig,omitempty" type:"Struct"`
	// The type of resource.
	//
	// example:
	//
	// Plugin
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s ListSecretReferencesResponseBodyDataItems) String() string {
	return dara.Prettify(s)
}

func (s ListSecretReferencesResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *ListSecretReferencesResponseBodyDataItems) GetPluginConfig() *ListSecretReferencesResponseBodyDataItemsPluginConfig {
	return s.PluginConfig
}

func (s *ListSecretReferencesResponseBodyDataItems) GetServiceConfig() *ListSecretReferencesResponseBodyDataItemsServiceConfig {
	return s.ServiceConfig
}

func (s *ListSecretReferencesResponseBodyDataItems) GetConsumerConfig() *ListSecretReferencesResponseBodyDataItemsConsumerConfig {
	return s.ConsumerConfig
}

func (s *ListSecretReferencesResponseBodyDataItems) GetGatewayId() *string {
	return s.GatewayId
}

func (s *ListSecretReferencesResponseBodyDataItems) GetMcpServerConfig() *ListSecretReferencesResponseBodyDataItemsMcpServerConfig {
	return s.McpServerConfig
}

func (s *ListSecretReferencesResponseBodyDataItems) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListSecretReferencesResponseBodyDataItems) SetPluginConfig(v *ListSecretReferencesResponseBodyDataItemsPluginConfig) *ListSecretReferencesResponseBodyDataItems {
	s.PluginConfig = v
	return s
}

func (s *ListSecretReferencesResponseBodyDataItems) SetServiceConfig(v *ListSecretReferencesResponseBodyDataItemsServiceConfig) *ListSecretReferencesResponseBodyDataItems {
	s.ServiceConfig = v
	return s
}

func (s *ListSecretReferencesResponseBodyDataItems) SetConsumerConfig(v *ListSecretReferencesResponseBodyDataItemsConsumerConfig) *ListSecretReferencesResponseBodyDataItems {
	s.ConsumerConfig = v
	return s
}

func (s *ListSecretReferencesResponseBodyDataItems) SetGatewayId(v string) *ListSecretReferencesResponseBodyDataItems {
	s.GatewayId = &v
	return s
}

func (s *ListSecretReferencesResponseBodyDataItems) SetMcpServerConfig(v *ListSecretReferencesResponseBodyDataItemsMcpServerConfig) *ListSecretReferencesResponseBodyDataItems {
	s.McpServerConfig = v
	return s
}

func (s *ListSecretReferencesResponseBodyDataItems) SetResourceType(v string) *ListSecretReferencesResponseBodyDataItems {
	s.ResourceType = &v
	return s
}

func (s *ListSecretReferencesResponseBodyDataItems) Validate() error {
	if s.PluginConfig != nil {
		if err := s.PluginConfig.Validate(); err != nil {
			return err
		}
	}
	if s.ServiceConfig != nil {
		if err := s.ServiceConfig.Validate(); err != nil {
			return err
		}
	}
	if s.ConsumerConfig != nil {
		if err := s.ConsumerConfig.Validate(); err != nil {
			return err
		}
	}
	if s.McpServerConfig != nil {
		if err := s.McpServerConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListSecretReferencesResponseBodyDataItemsPluginConfig struct {
	// The plug-in name.
	//
	// example:
	//
	// oauth
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The plug-in type ID.
	//
	// example:
	//
	// pls-xxxxxxxx
	PluginClassId *string `json:"pluginClassId,omitempty" xml:"pluginClassId,omitempty"`
	// The plug-in ID.
	//
	// example:
	//
	// pl-d4ijk56m1hkhxxxxxxxx
	PluginId *string `json:"pluginId,omitempty" xml:"pluginId,omitempty"`
}

func (s ListSecretReferencesResponseBodyDataItemsPluginConfig) String() string {
	return dara.Prettify(s)
}

func (s ListSecretReferencesResponseBodyDataItemsPluginConfig) GoString() string {
	return s.String()
}

func (s *ListSecretReferencesResponseBodyDataItemsPluginConfig) GetName() *string {
	return s.Name
}

func (s *ListSecretReferencesResponseBodyDataItemsPluginConfig) GetPluginClassId() *string {
	return s.PluginClassId
}

func (s *ListSecretReferencesResponseBodyDataItemsPluginConfig) GetPluginId() *string {
	return s.PluginId
}

func (s *ListSecretReferencesResponseBodyDataItemsPluginConfig) SetName(v string) *ListSecretReferencesResponseBodyDataItemsPluginConfig {
	s.Name = &v
	return s
}

func (s *ListSecretReferencesResponseBodyDataItemsPluginConfig) SetPluginClassId(v string) *ListSecretReferencesResponseBodyDataItemsPluginConfig {
	s.PluginClassId = &v
	return s
}

func (s *ListSecretReferencesResponseBodyDataItemsPluginConfig) SetPluginId(v string) *ListSecretReferencesResponseBodyDataItemsPluginConfig {
	s.PluginId = &v
	return s
}

func (s *ListSecretReferencesResponseBodyDataItemsPluginConfig) Validate() error {
	return dara.Validate(s)
}

type ListSecretReferencesResponseBodyDataItemsServiceConfig struct {
	// The service name.
	//
	// example:
	//
	// myService
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The service ID.
	//
	// example:
	//
	// svc-cvgbtcmm1hkmxxxxxxxx
	ServiceId *string `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
}

func (s ListSecretReferencesResponseBodyDataItemsServiceConfig) String() string {
	return dara.Prettify(s)
}

func (s ListSecretReferencesResponseBodyDataItemsServiceConfig) GoString() string {
	return s.String()
}

func (s *ListSecretReferencesResponseBodyDataItemsServiceConfig) GetName() *string {
	return s.Name
}

func (s *ListSecretReferencesResponseBodyDataItemsServiceConfig) GetServiceId() *string {
	return s.ServiceId
}

func (s *ListSecretReferencesResponseBodyDataItemsServiceConfig) SetName(v string) *ListSecretReferencesResponseBodyDataItemsServiceConfig {
	s.Name = &v
	return s
}

func (s *ListSecretReferencesResponseBodyDataItemsServiceConfig) SetServiceId(v string) *ListSecretReferencesResponseBodyDataItemsServiceConfig {
	s.ServiceId = &v
	return s
}

func (s *ListSecretReferencesResponseBodyDataItemsServiceConfig) Validate() error {
	return dara.Validate(s)
}

type ListSecretReferencesResponseBodyDataItemsConsumerConfig struct {
	// The consumer ID.
	//
	// example:
	//
	// cs-d0iltnem1hkhxxxxxxxx
	ConsumerId *string `json:"consumerId,omitempty" xml:"consumerId,omitempty"`
	// The consumer name.
	//
	// example:
	//
	// myconsumer
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListSecretReferencesResponseBodyDataItemsConsumerConfig) String() string {
	return dara.Prettify(s)
}

func (s ListSecretReferencesResponseBodyDataItemsConsumerConfig) GoString() string {
	return s.String()
}

func (s *ListSecretReferencesResponseBodyDataItemsConsumerConfig) GetConsumerId() *string {
	return s.ConsumerId
}

func (s *ListSecretReferencesResponseBodyDataItemsConsumerConfig) GetName() *string {
	return s.Name
}

func (s *ListSecretReferencesResponseBodyDataItemsConsumerConfig) SetConsumerId(v string) *ListSecretReferencesResponseBodyDataItemsConsumerConfig {
	s.ConsumerId = &v
	return s
}

func (s *ListSecretReferencesResponseBodyDataItemsConsumerConfig) SetName(v string) *ListSecretReferencesResponseBodyDataItemsConsumerConfig {
	s.Name = &v
	return s
}

func (s *ListSecretReferencesResponseBodyDataItemsConsumerConfig) Validate() error {
	return dara.Validate(s)
}

type ListSecretReferencesResponseBodyDataItemsMcpServerConfig struct {
	// The HTTP API ID.
	//
	// example:
	//
	// api-d2vv43em201hxxxxxxxx
	HttpApiId *string `json:"httpApiId,omitempty" xml:"httpApiId,omitempty"`
	// The route name.
	//
	// example:
	//
	// mcp
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The route ID.
	//
	// example:
	//
	// hr-cv0i5oum1hkhxxxxxxxx
	RouteId *string `json:"routeId,omitempty" xml:"routeId,omitempty"`
}

func (s ListSecretReferencesResponseBodyDataItemsMcpServerConfig) String() string {
	return dara.Prettify(s)
}

func (s ListSecretReferencesResponseBodyDataItemsMcpServerConfig) GoString() string {
	return s.String()
}

func (s *ListSecretReferencesResponseBodyDataItemsMcpServerConfig) GetHttpApiId() *string {
	return s.HttpApiId
}

func (s *ListSecretReferencesResponseBodyDataItemsMcpServerConfig) GetName() *string {
	return s.Name
}

func (s *ListSecretReferencesResponseBodyDataItemsMcpServerConfig) GetRouteId() *string {
	return s.RouteId
}

func (s *ListSecretReferencesResponseBodyDataItemsMcpServerConfig) SetHttpApiId(v string) *ListSecretReferencesResponseBodyDataItemsMcpServerConfig {
	s.HttpApiId = &v
	return s
}

func (s *ListSecretReferencesResponseBodyDataItemsMcpServerConfig) SetName(v string) *ListSecretReferencesResponseBodyDataItemsMcpServerConfig {
	s.Name = &v
	return s
}

func (s *ListSecretReferencesResponseBodyDataItemsMcpServerConfig) SetRouteId(v string) *ListSecretReferencesResponseBodyDataItemsMcpServerConfig {
	s.RouteId = &v
	return s
}

func (s *ListSecretReferencesResponseBodyDataItemsMcpServerConfig) Validate() error {
	return dara.Validate(s)
}
