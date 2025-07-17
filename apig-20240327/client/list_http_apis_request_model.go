// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHttpApisRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGatewayId(v string) *ListHttpApisRequest
	GetGatewayId() *string
	SetGatewayType(v string) *ListHttpApisRequest
	GetGatewayType() *string
	SetKeyword(v string) *ListHttpApisRequest
	GetKeyword() *string
	SetName(v string) *ListHttpApisRequest
	GetName() *string
	SetPageNumber(v int32) *ListHttpApisRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListHttpApisRequest
	GetPageSize() *int32
	SetResourceGroupId(v string) *ListHttpApisRequest
	GetResourceGroupId() *string
	SetTypes(v string) *ListHttpApisRequest
	GetTypes() *string
	SetWithAPIsPublishedToEnvironment(v bool) *ListHttpApisRequest
	GetWithAPIsPublishedToEnvironment() *bool
	SetWithAuthPolicyInEnvironmentId(v string) *ListHttpApisRequest
	GetWithAuthPolicyInEnvironmentId() *string
	SetWithAuthPolicyList(v bool) *ListHttpApisRequest
	GetWithAuthPolicyList() *bool
	SetWithConsumerInfoById(v string) *ListHttpApisRequest
	GetWithConsumerInfoById() *string
	SetWithEnvironmentInfo(v bool) *ListHttpApisRequest
	GetWithEnvironmentInfo() *bool
	SetWithEnvironmentInfoById(v string) *ListHttpApisRequest
	GetWithEnvironmentInfoById() *string
	SetWithIngressInfo(v bool) *ListHttpApisRequest
	GetWithIngressInfo() *bool
	SetWithPluginAttachmentByPluginId(v string) *ListHttpApisRequest
	GetWithPluginAttachmentByPluginId() *string
	SetWithPolicyConfigs(v bool) *ListHttpApisRequest
	GetWithPolicyConfigs() *bool
}

type ListHttpApisRequest struct {
	// The ID of the Cloud-native API Gateway instance.
	//
	// example:
	//
	// gw-cq2avtllh****
	GatewayId   *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	GatewayType *string `json:"gatewayType,omitempty" xml:"gatewayType,omitempty"`
	// The search keyword. You can fuzzy-search by API name or exact-search by API ID.
	//
	// example:
	//
	// test-
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// The API name that is used for the search. In this case, exact search is performed.
	//
	// example:
	//
	// login
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The page number of the page to return. Pages start from page 1. Default value: 1.
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
	// rg-ahr5uil8raz0rq3b
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// The API type. You can specify multiple types and separate them with a comma (,).
	//
	// 	- Http
	//
	// 	- Rest
	//
	// 	- WebSocket
	//
	// 	- HttpIngress
	//
	// example:
	//
	// Http,Rest
	Types                          *string `json:"types,omitempty" xml:"types,omitempty"`
	WithAPIsPublishedToEnvironment *bool   `json:"withAPIsPublishedToEnvironment,omitempty" xml:"withAPIsPublishedToEnvironment,omitempty"`
	// The consumer authentication policy in the specified environment in each returned API.
	//
	// example:
	//
	// env-xxx
	WithAuthPolicyInEnvironmentId *string `json:"withAuthPolicyInEnvironmentId,omitempty" xml:"withAuthPolicyInEnvironmentId,omitempty"`
	// Specifies whether authentication is enabled.
	//
	// example:
	//
	// true
	WithAuthPolicyList *bool `json:"withAuthPolicyList,omitempty" xml:"withAuthPolicyList,omitempty"`
	// The authorization rules of the specified consumer in each returned API.
	//
	// example:
	//
	// cs-xxx
	WithConsumerInfoById *string `json:"withConsumerInfoById,omitempty" xml:"withConsumerInfoById,omitempty"`
	// The environment information.
	//
	// example:
	//
	// true
	WithEnvironmentInfo *bool `json:"withEnvironmentInfo,omitempty" xml:"withEnvironmentInfo,omitempty"`
	// The environment ID.
	//
	// example:
	//
	// env-ctovu5mm1hksb4q8ln40
	WithEnvironmentInfoById *string `json:"withEnvironmentInfoById,omitempty" xml:"withEnvironmentInfoById,omitempty"`
	// The Ingress information.
	//
	// example:
	//
	// false
	WithIngressInfo *bool `json:"withIngressInfo,omitempty" xml:"withIngressInfo,omitempty"`
	// The plug-in ID. You can use the returned value of this parameter to query the plug-in.
	//
	// example:
	//
	// pl-ct9qn3um1hktue8dqol0
	WithPluginAttachmentByPluginId *string `json:"withPluginAttachmentByPluginId,omitempty" xml:"withPluginAttachmentByPluginId,omitempty"`
	WithPolicyConfigs              *bool   `json:"withPolicyConfigs,omitempty" xml:"withPolicyConfigs,omitempty"`
}

func (s ListHttpApisRequest) String() string {
	return dara.Prettify(s)
}

func (s ListHttpApisRequest) GoString() string {
	return s.String()
}

func (s *ListHttpApisRequest) GetGatewayId() *string {
	return s.GatewayId
}

func (s *ListHttpApisRequest) GetGatewayType() *string {
	return s.GatewayType
}

func (s *ListHttpApisRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListHttpApisRequest) GetName() *string {
	return s.Name
}

func (s *ListHttpApisRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListHttpApisRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListHttpApisRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListHttpApisRequest) GetTypes() *string {
	return s.Types
}

func (s *ListHttpApisRequest) GetWithAPIsPublishedToEnvironment() *bool {
	return s.WithAPIsPublishedToEnvironment
}

func (s *ListHttpApisRequest) GetWithAuthPolicyInEnvironmentId() *string {
	return s.WithAuthPolicyInEnvironmentId
}

func (s *ListHttpApisRequest) GetWithAuthPolicyList() *bool {
	return s.WithAuthPolicyList
}

func (s *ListHttpApisRequest) GetWithConsumerInfoById() *string {
	return s.WithConsumerInfoById
}

func (s *ListHttpApisRequest) GetWithEnvironmentInfo() *bool {
	return s.WithEnvironmentInfo
}

func (s *ListHttpApisRequest) GetWithEnvironmentInfoById() *string {
	return s.WithEnvironmentInfoById
}

func (s *ListHttpApisRequest) GetWithIngressInfo() *bool {
	return s.WithIngressInfo
}

func (s *ListHttpApisRequest) GetWithPluginAttachmentByPluginId() *string {
	return s.WithPluginAttachmentByPluginId
}

func (s *ListHttpApisRequest) GetWithPolicyConfigs() *bool {
	return s.WithPolicyConfigs
}

func (s *ListHttpApisRequest) SetGatewayId(v string) *ListHttpApisRequest {
	s.GatewayId = &v
	return s
}

func (s *ListHttpApisRequest) SetGatewayType(v string) *ListHttpApisRequest {
	s.GatewayType = &v
	return s
}

func (s *ListHttpApisRequest) SetKeyword(v string) *ListHttpApisRequest {
	s.Keyword = &v
	return s
}

func (s *ListHttpApisRequest) SetName(v string) *ListHttpApisRequest {
	s.Name = &v
	return s
}

func (s *ListHttpApisRequest) SetPageNumber(v int32) *ListHttpApisRequest {
	s.PageNumber = &v
	return s
}

func (s *ListHttpApisRequest) SetPageSize(v int32) *ListHttpApisRequest {
	s.PageSize = &v
	return s
}

func (s *ListHttpApisRequest) SetResourceGroupId(v string) *ListHttpApisRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListHttpApisRequest) SetTypes(v string) *ListHttpApisRequest {
	s.Types = &v
	return s
}

func (s *ListHttpApisRequest) SetWithAPIsPublishedToEnvironment(v bool) *ListHttpApisRequest {
	s.WithAPIsPublishedToEnvironment = &v
	return s
}

func (s *ListHttpApisRequest) SetWithAuthPolicyInEnvironmentId(v string) *ListHttpApisRequest {
	s.WithAuthPolicyInEnvironmentId = &v
	return s
}

func (s *ListHttpApisRequest) SetWithAuthPolicyList(v bool) *ListHttpApisRequest {
	s.WithAuthPolicyList = &v
	return s
}

func (s *ListHttpApisRequest) SetWithConsumerInfoById(v string) *ListHttpApisRequest {
	s.WithConsumerInfoById = &v
	return s
}

func (s *ListHttpApisRequest) SetWithEnvironmentInfo(v bool) *ListHttpApisRequest {
	s.WithEnvironmentInfo = &v
	return s
}

func (s *ListHttpApisRequest) SetWithEnvironmentInfoById(v string) *ListHttpApisRequest {
	s.WithEnvironmentInfoById = &v
	return s
}

func (s *ListHttpApisRequest) SetWithIngressInfo(v bool) *ListHttpApisRequest {
	s.WithIngressInfo = &v
	return s
}

func (s *ListHttpApisRequest) SetWithPluginAttachmentByPluginId(v string) *ListHttpApisRequest {
	s.WithPluginAttachmentByPluginId = &v
	return s
}

func (s *ListHttpApisRequest) SetWithPolicyConfigs(v bool) *ListHttpApisRequest {
	s.WithPolicyConfigs = &v
	return s
}

func (s *ListHttpApisRequest) Validate() error {
	return dara.Validate(s)
}
