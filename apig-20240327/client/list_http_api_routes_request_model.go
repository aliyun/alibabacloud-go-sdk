// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHttpApiRoutesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackendServiceName(v string) *ListHttpApiRoutesRequest
	GetBackendServiceName() *string
	SetConsumerAuthorizationRuleId(v string) *ListHttpApiRoutesRequest
	GetConsumerAuthorizationRuleId() *string
	SetDeployStatuses(v string) *ListHttpApiRoutesRequest
	GetDeployStatuses() *string
	SetDomainId(v string) *ListHttpApiRoutesRequest
	GetDomainId() *string
	SetEnvironmentId(v string) *ListHttpApiRoutesRequest
	GetEnvironmentId() *string
	SetForDeploy(v bool) *ListHttpApiRoutesRequest
	GetForDeploy() *bool
	SetGatewayId(v string) *ListHttpApiRoutesRequest
	GetGatewayId() *string
	SetName(v string) *ListHttpApiRoutesRequest
	GetName() *string
	SetNameLike(v string) *ListHttpApiRoutesRequest
	GetNameLike() *string
	SetPageNumber(v int32) *ListHttpApiRoutesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListHttpApiRoutesRequest
	GetPageSize() *int32
	SetPathLike(v string) *ListHttpApiRoutesRequest
	GetPathLike() *string
	SetWithAuthPolicyInfo(v bool) *ListHttpApiRoutesRequest
	GetWithAuthPolicyInfo() *bool
	SetWithConsumerInfoById(v string) *ListHttpApiRoutesRequest
	GetWithConsumerInfoById() *string
	SetWithPluginAttachmentByPluginId(v string) *ListHttpApiRoutesRequest
	GetWithPluginAttachmentByPluginId() *string
}

type ListHttpApiRoutesRequest struct {
	// example:
	//
	// test-svc
	BackendServiceName *string `json:"backendServiceName,omitempty" xml:"backendServiceName,omitempty"`
	// The consumer authorization rule ID. If specified, the response includes only routes that are authorized by the specified rule.
	//
	// example:
	//
	// cas-xxx
	ConsumerAuthorizationRuleId *string `json:"consumerAuthorizationRuleId,omitempty" xml:"consumerAuthorizationRuleId,omitempty"`
	// The deployment status of the route.
	//
	// example:
	//
	// NotDeployed
	DeployStatuses *string `json:"deployStatuses,omitempty" xml:"deployStatuses,omitempty"`
	// The domain name ID used to filter routes.
	//
	// example:
	//
	// d-xxx
	DomainId *string `json:"domainId,omitempty" xml:"domainId,omitempty"`
	// The environment ID.
	//
	// example:
	//
	// env-cpqnr6tlhtgubc***
	EnvironmentId *string `json:"environmentId,omitempty" xml:"environmentId,omitempty"`
	// Specifies whether the query is for a deployment scenario.
	//
	// example:
	//
	// true
	ForDeploy *bool `json:"forDeploy,omitempty" xml:"forDeploy,omitempty"`
	// The cloud-native API gateway ID.
	//
	// example:
	//
	// gw-cpv4sqdl****
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// The route name.
	//
	// example:
	//
	// itemcenter-gateway
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The route name for fuzzy match.
	//
	// example:
	//
	// item
	NameLike *string `json:"nameLike,omitempty" xml:"nameLike,omitempty"`
	// The page number. Pages start from 1. Default value: 1.
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
	// The route path for fuzzy match.
	//
	// example:
	//
	// /v1
	PathLike *string `json:"pathLike,omitempty" xml:"pathLike,omitempty"`
	// Specifies whether to include consumer authorization information in the response.
	//
	// example:
	//
	// true
	WithAuthPolicyInfo *bool `json:"withAuthPolicyInfo,omitempty" xml:"withAuthPolicyInfo,omitempty"`
	// The consumer ID. If specified, the response includes the authorization rule list for the specified consumer in each route.
	//
	// example:
	//
	// cs-xxx
	WithConsumerInfoById *string `json:"withConsumerInfoById,omitempty" xml:"withConsumerInfoById,omitempty"`
	// The plug-in ID. If specified, the response includes the attachment information of the specified plug-in for each route.
	//
	// example:
	//
	// pl-xxx
	WithPluginAttachmentByPluginId *string `json:"withPluginAttachmentByPluginId,omitempty" xml:"withPluginAttachmentByPluginId,omitempty"`
}

func (s ListHttpApiRoutesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListHttpApiRoutesRequest) GoString() string {
	return s.String()
}

func (s *ListHttpApiRoutesRequest) GetBackendServiceName() *string {
	return s.BackendServiceName
}

func (s *ListHttpApiRoutesRequest) GetConsumerAuthorizationRuleId() *string {
	return s.ConsumerAuthorizationRuleId
}

func (s *ListHttpApiRoutesRequest) GetDeployStatuses() *string {
	return s.DeployStatuses
}

func (s *ListHttpApiRoutesRequest) GetDomainId() *string {
	return s.DomainId
}

func (s *ListHttpApiRoutesRequest) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *ListHttpApiRoutesRequest) GetForDeploy() *bool {
	return s.ForDeploy
}

func (s *ListHttpApiRoutesRequest) GetGatewayId() *string {
	return s.GatewayId
}

func (s *ListHttpApiRoutesRequest) GetName() *string {
	return s.Name
}

func (s *ListHttpApiRoutesRequest) GetNameLike() *string {
	return s.NameLike
}

func (s *ListHttpApiRoutesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListHttpApiRoutesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListHttpApiRoutesRequest) GetPathLike() *string {
	return s.PathLike
}

func (s *ListHttpApiRoutesRequest) GetWithAuthPolicyInfo() *bool {
	return s.WithAuthPolicyInfo
}

func (s *ListHttpApiRoutesRequest) GetWithConsumerInfoById() *string {
	return s.WithConsumerInfoById
}

func (s *ListHttpApiRoutesRequest) GetWithPluginAttachmentByPluginId() *string {
	return s.WithPluginAttachmentByPluginId
}

func (s *ListHttpApiRoutesRequest) SetBackendServiceName(v string) *ListHttpApiRoutesRequest {
	s.BackendServiceName = &v
	return s
}

func (s *ListHttpApiRoutesRequest) SetConsumerAuthorizationRuleId(v string) *ListHttpApiRoutesRequest {
	s.ConsumerAuthorizationRuleId = &v
	return s
}

func (s *ListHttpApiRoutesRequest) SetDeployStatuses(v string) *ListHttpApiRoutesRequest {
	s.DeployStatuses = &v
	return s
}

func (s *ListHttpApiRoutesRequest) SetDomainId(v string) *ListHttpApiRoutesRequest {
	s.DomainId = &v
	return s
}

func (s *ListHttpApiRoutesRequest) SetEnvironmentId(v string) *ListHttpApiRoutesRequest {
	s.EnvironmentId = &v
	return s
}

func (s *ListHttpApiRoutesRequest) SetForDeploy(v bool) *ListHttpApiRoutesRequest {
	s.ForDeploy = &v
	return s
}

func (s *ListHttpApiRoutesRequest) SetGatewayId(v string) *ListHttpApiRoutesRequest {
	s.GatewayId = &v
	return s
}

func (s *ListHttpApiRoutesRequest) SetName(v string) *ListHttpApiRoutesRequest {
	s.Name = &v
	return s
}

func (s *ListHttpApiRoutesRequest) SetNameLike(v string) *ListHttpApiRoutesRequest {
	s.NameLike = &v
	return s
}

func (s *ListHttpApiRoutesRequest) SetPageNumber(v int32) *ListHttpApiRoutesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListHttpApiRoutesRequest) SetPageSize(v int32) *ListHttpApiRoutesRequest {
	s.PageSize = &v
	return s
}

func (s *ListHttpApiRoutesRequest) SetPathLike(v string) *ListHttpApiRoutesRequest {
	s.PathLike = &v
	return s
}

func (s *ListHttpApiRoutesRequest) SetWithAuthPolicyInfo(v bool) *ListHttpApiRoutesRequest {
	s.WithAuthPolicyInfo = &v
	return s
}

func (s *ListHttpApiRoutesRequest) SetWithConsumerInfoById(v string) *ListHttpApiRoutesRequest {
	s.WithConsumerInfoById = &v
	return s
}

func (s *ListHttpApiRoutesRequest) SetWithPluginAttachmentByPluginId(v string) *ListHttpApiRoutesRequest {
	s.WithPluginAttachmentByPluginId = &v
	return s
}

func (s *ListHttpApiRoutesRequest) Validate() error {
	return dara.Validate(s)
}
