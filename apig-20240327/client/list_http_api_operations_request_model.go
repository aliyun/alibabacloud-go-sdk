// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHttpApiOperationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConsumerAuthorizationRuleId(v string) *ListHttpApiOperationsRequest
	GetConsumerAuthorizationRuleId() *string
	SetEnableAuth(v bool) *ListHttpApiOperationsRequest
	GetEnableAuth() *bool
	SetForDeploy(v bool) *ListHttpApiOperationsRequest
	GetForDeploy() *bool
	SetGatewayId(v string) *ListHttpApiOperationsRequest
	GetGatewayId() *string
	SetMethod(v string) *ListHttpApiOperationsRequest
	GetMethod() *string
	SetName(v string) *ListHttpApiOperationsRequest
	GetName() *string
	SetNameLike(v string) *ListHttpApiOperationsRequest
	GetNameLike() *string
	SetPageNumber(v int32) *ListHttpApiOperationsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListHttpApiOperationsRequest
	GetPageSize() *int32
	SetPathLike(v string) *ListHttpApiOperationsRequest
	GetPathLike() *string
	SetWithConsumerInEnvironmentId(v string) *ListHttpApiOperationsRequest
	GetWithConsumerInEnvironmentId() *string
	SetWithConsumerInfoById(v string) *ListHttpApiOperationsRequest
	GetWithConsumerInfoById() *string
	SetWithPluginAttachmentByPluginId(v string) *ListHttpApiOperationsRequest
	GetWithPluginAttachmentByPluginId() *string
}

type ListHttpApiOperationsRequest struct {
	// The consumer authorization rule ID used to filter the operation list. The response includes only operations that are authorized by the specified rule.
	//
	// example:
	//
	// cas-xxx
	ConsumerAuthorizationRuleId *string `json:"consumerAuthorizationRuleId,omitempty" xml:"consumerAuthorizationRuleId,omitempty"`
	// The authentication enablement filter.
	//
	// example:
	//
	// true
	EnableAuth *bool `json:"enableAuth,omitempty" xml:"enableAuth,omitempty"`
	// Specifies whether the request is for a deployment scenario.
	//
	// example:
	//
	// true
	ForDeploy *bool `json:"forDeploy,omitempty" xml:"forDeploy,omitempty"`
	// The gateway ID filter.
	//
	// example:
	//
	// gw-001
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// Lists operations by HTTP method.
	//
	// example:
	//
	// GET
	Method *string `json:"method,omitempty" xml:"method,omitempty"`
	// Searches for operations by exact name match.
	//
	// example:
	//
	// getUserInfo
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Searches for operations by name prefix.
	//
	// example:
	//
	// GetUser
	NameLike *string `json:"nameLike,omitempty" xml:"nameLike,omitempty"`
	// The page number, starting from 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The page size. Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Searches for operations by path prefix match.
	//
	// example:
	//
	// /v1
	PathLike *string `json:"pathLike,omitempty" xml:"pathLike,omitempty"`
	// The environment ID. When specified together with withConsumerInfoById, the response includes the authorization rule list of the specified consumer in the specified environment for each operation.
	//
	// example:
	//
	// env-xxx
	WithConsumerInEnvironmentId *string `json:"withConsumerInEnvironmentId,omitempty" xml:"withConsumerInEnvironmentId,omitempty"`
	// The consumer ID. When specified together with withConsumerInEnvironmentId, the response includes the authorization rule list of the specified consumer in the specified environment for each operation.
	//
	// example:
	//
	// cs-xxx
	WithConsumerInfoById *string `json:"withConsumerInfoById,omitempty" xml:"withConsumerInfoById,omitempty"`
	// The plugin ID used to retrieve plugin deployment information.
	//
	// example:
	//
	// pl-xxx
	WithPluginAttachmentByPluginId *string `json:"withPluginAttachmentByPluginId,omitempty" xml:"withPluginAttachmentByPluginId,omitempty"`
}

func (s ListHttpApiOperationsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListHttpApiOperationsRequest) GoString() string {
	return s.String()
}

func (s *ListHttpApiOperationsRequest) GetConsumerAuthorizationRuleId() *string {
	return s.ConsumerAuthorizationRuleId
}

func (s *ListHttpApiOperationsRequest) GetEnableAuth() *bool {
	return s.EnableAuth
}

func (s *ListHttpApiOperationsRequest) GetForDeploy() *bool {
	return s.ForDeploy
}

func (s *ListHttpApiOperationsRequest) GetGatewayId() *string {
	return s.GatewayId
}

func (s *ListHttpApiOperationsRequest) GetMethod() *string {
	return s.Method
}

func (s *ListHttpApiOperationsRequest) GetName() *string {
	return s.Name
}

func (s *ListHttpApiOperationsRequest) GetNameLike() *string {
	return s.NameLike
}

func (s *ListHttpApiOperationsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListHttpApiOperationsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListHttpApiOperationsRequest) GetPathLike() *string {
	return s.PathLike
}

func (s *ListHttpApiOperationsRequest) GetWithConsumerInEnvironmentId() *string {
	return s.WithConsumerInEnvironmentId
}

func (s *ListHttpApiOperationsRequest) GetWithConsumerInfoById() *string {
	return s.WithConsumerInfoById
}

func (s *ListHttpApiOperationsRequest) GetWithPluginAttachmentByPluginId() *string {
	return s.WithPluginAttachmentByPluginId
}

func (s *ListHttpApiOperationsRequest) SetConsumerAuthorizationRuleId(v string) *ListHttpApiOperationsRequest {
	s.ConsumerAuthorizationRuleId = &v
	return s
}

func (s *ListHttpApiOperationsRequest) SetEnableAuth(v bool) *ListHttpApiOperationsRequest {
	s.EnableAuth = &v
	return s
}

func (s *ListHttpApiOperationsRequest) SetForDeploy(v bool) *ListHttpApiOperationsRequest {
	s.ForDeploy = &v
	return s
}

func (s *ListHttpApiOperationsRequest) SetGatewayId(v string) *ListHttpApiOperationsRequest {
	s.GatewayId = &v
	return s
}

func (s *ListHttpApiOperationsRequest) SetMethod(v string) *ListHttpApiOperationsRequest {
	s.Method = &v
	return s
}

func (s *ListHttpApiOperationsRequest) SetName(v string) *ListHttpApiOperationsRequest {
	s.Name = &v
	return s
}

func (s *ListHttpApiOperationsRequest) SetNameLike(v string) *ListHttpApiOperationsRequest {
	s.NameLike = &v
	return s
}

func (s *ListHttpApiOperationsRequest) SetPageNumber(v int32) *ListHttpApiOperationsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListHttpApiOperationsRequest) SetPageSize(v int32) *ListHttpApiOperationsRequest {
	s.PageSize = &v
	return s
}

func (s *ListHttpApiOperationsRequest) SetPathLike(v string) *ListHttpApiOperationsRequest {
	s.PathLike = &v
	return s
}

func (s *ListHttpApiOperationsRequest) SetWithConsumerInEnvironmentId(v string) *ListHttpApiOperationsRequest {
	s.WithConsumerInEnvironmentId = &v
	return s
}

func (s *ListHttpApiOperationsRequest) SetWithConsumerInfoById(v string) *ListHttpApiOperationsRequest {
	s.WithConsumerInfoById = &v
	return s
}

func (s *ListHttpApiOperationsRequest) SetWithPluginAttachmentByPluginId(v string) *ListHttpApiOperationsRequest {
	s.WithPluginAttachmentByPluginId = &v
	return s
}

func (s *ListHttpApiOperationsRequest) Validate() error {
	return dara.Validate(s)
}
