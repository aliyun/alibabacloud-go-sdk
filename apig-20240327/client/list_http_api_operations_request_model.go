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
	// Filter the operation list based on a specific consumer authorization rule ID, and the interface list in the response only contains authorized operations.
	//
	// example:
	//
	// cas-xxx
	ConsumerAuthorizationRuleId *string `json:"consumerAuthorizationRuleId,omitempty" xml:"consumerAuthorizationRuleId,omitempty"`
	ForDeploy                   *bool   `json:"forDeploy,omitempty" xml:"forDeploy,omitempty"`
	GatewayId                   *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// List interfaces by Method.
	//
	// example:
	//
	// GET
	Method *string `json:"method,omitempty" xml:"method,omitempty"`
	// Search operations by exact name.
	//
	// example:
	//
	// getUserInfo
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Search operations by name prefix.
	//
	// example:
	//
	// GetUser
	NameLike *string `json:"nameLike,omitempty" xml:"nameLike,omitempty"`
	// Page number, starting from 1, default is 1 if not specified.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// Page size, valid range [1, 100], default is 10 if not specified.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Search operations by path prefix.
	//
	// example:
	//
	// /v1
	PathLike *string `json:"pathLike,omitempty" xml:"pathLike,omitempty"`
	// Each operation information in the response carries a list of authorization rules for the specified consumer under the specified environment ID. The withConsumerInEnvironmentId field needs to be additionally specified.
	//
	// example:
	//
	// env-xxx
	WithConsumerInEnvironmentId *string `json:"withConsumerInEnvironmentId,omitempty" xml:"withConsumerInEnvironmentId,omitempty"`
	// Each operation information in the response carries a list of authorization rules for the specified consumer under the specified environment ID. The withConsumerInEnvironmentId field needs to be additionally specified.
	//
	// example:
	//
	// cs-xxx
	WithConsumerInfoById *string `json:"withConsumerInfoById,omitempty" xml:"withConsumerInfoById,omitempty"`
	// Plugin ID, use this plugin ID to retrieve the plugin release information.
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
