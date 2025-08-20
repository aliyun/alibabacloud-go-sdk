// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceTypesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEntityType(v string) *ListResourceTypesRequest
	GetEntityType() *string
	SetProvider(v string) *ListResourceTypesRequest
	GetProvider() *string
	SetResourceType(v string) *ListResourceTypesRequest
	GetResourceType() *string
}

type ListResourceTypesRequest struct {
	// The entity type. Valid values:
	//
	// 	- All: all types of resources.
	//
	// 	- Resource (default): regular resources. For more information, see [Resources](https://help.aliyun.com/document_detail/28863.html).
	//
	// 	- DataSource: DataSource resources. For more information, see [DataSource resources](https://help.aliyun.com/document_detail/404753.html).
	//
	// 	- Module: modules.
	//
	// example:
	//
	// Resource
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// The provider of the resource type. Valid values:
	//
	// 	- ROS (default): The resource type is provided by Resource Orchestration Service (ROS).
	//
	// 	- Self: The resource type is provided by you.
	//
	// example:
	//
	// ROS
	Provider *string `json:"Provider,omitempty" xml:"Provider,omitempty"`
	// The resource type. The resource type can contain letters, digits, colons (:), and asterisks (\\*). You can use an asterisk (\\*) to perform a fuzzy match.
	//
	// example:
	//
	// MODULE::MyOrganization::MyService::MyUsecase
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ListResourceTypesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListResourceTypesRequest) GoString() string {
	return s.String()
}

func (s *ListResourceTypesRequest) GetEntityType() *string {
	return s.EntityType
}

func (s *ListResourceTypesRequest) GetProvider() *string {
	return s.Provider
}

func (s *ListResourceTypesRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListResourceTypesRequest) SetEntityType(v string) *ListResourceTypesRequest {
	s.EntityType = &v
	return s
}

func (s *ListResourceTypesRequest) SetProvider(v string) *ListResourceTypesRequest {
	s.Provider = &v
	return s
}

func (s *ListResourceTypesRequest) SetResourceType(v string) *ListResourceTypesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListResourceTypesRequest) Validate() error {
	return dara.Validate(s)
}
