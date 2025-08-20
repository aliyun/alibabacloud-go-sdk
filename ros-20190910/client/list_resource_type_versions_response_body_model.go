// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceTypeVersionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListResourceTypeVersionsResponseBody
	GetRequestId() *string
	SetResourceTypeVersions(v []*ListResourceTypeVersionsResponseBodyResourceTypeVersions) *ListResourceTypeVersionsResponseBody
	GetResourceTypeVersions() []*ListResourceTypeVersionsResponseBodyResourceTypeVersions
}

type ListResourceTypeVersionsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// B288A0BE-D927-4888-B0F7-B35EF84B6E6F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The versions of the resource types.
	ResourceTypeVersions []*ListResourceTypeVersionsResponseBodyResourceTypeVersions `json:"ResourceTypeVersions,omitempty" xml:"ResourceTypeVersions,omitempty" type:"Repeated"`
}

func (s ListResourceTypeVersionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListResourceTypeVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourceTypeVersionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListResourceTypeVersionsResponseBody) GetResourceTypeVersions() []*ListResourceTypeVersionsResponseBodyResourceTypeVersions {
	return s.ResourceTypeVersions
}

func (s *ListResourceTypeVersionsResponseBody) SetRequestId(v string) *ListResourceTypeVersionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListResourceTypeVersionsResponseBody) SetResourceTypeVersions(v []*ListResourceTypeVersionsResponseBodyResourceTypeVersions) *ListResourceTypeVersionsResponseBody {
	s.ResourceTypeVersions = v
	return s
}

func (s *ListResourceTypeVersionsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListResourceTypeVersionsResponseBodyResourceTypeVersions struct {
	// The time when the version was created. The time is displayed in UTC. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ss format.
	//
	// example:
	//
	// 2023-02-24T08:25:21
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the version.
	//
	// example:
	//
	// It is a demo.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The entity type. Only Module may be returned.
	//
	// example:
	//
	// Module
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// Indicates whether the version is the default version. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	IsDefaultVersion *bool `json:"IsDefaultVersion,omitempty" xml:"IsDefaultVersion,omitempty"`
	// The provider of the resource type. Valid values:
	//
	// 	- ROS: ROS
	//
	// 	- Self: yourself
	//
	// example:
	//
	// ROS
	Provider *string `json:"Provider,omitempty" xml:"Provider,omitempty"`
	// The resource type.
	//
	// example:
	//
	// MODULE::MyOrganization::MyService::MyUsecase
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The time when the version was updated. The time is displayed in UTC. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ss format.
	//
	// example:
	//
	// 2023-02-24T08:25:21
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The version ID.
	//
	// example:
	//
	// v1
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s ListResourceTypeVersionsResponseBodyResourceTypeVersions) String() string {
	return dara.Prettify(s)
}

func (s ListResourceTypeVersionsResponseBodyResourceTypeVersions) GoString() string {
	return s.String()
}

func (s *ListResourceTypeVersionsResponseBodyResourceTypeVersions) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListResourceTypeVersionsResponseBodyResourceTypeVersions) GetDescription() *string {
	return s.Description
}

func (s *ListResourceTypeVersionsResponseBodyResourceTypeVersions) GetEntityType() *string {
	return s.EntityType
}

func (s *ListResourceTypeVersionsResponseBodyResourceTypeVersions) GetIsDefaultVersion() *bool {
	return s.IsDefaultVersion
}

func (s *ListResourceTypeVersionsResponseBodyResourceTypeVersions) GetProvider() *string {
	return s.Provider
}

func (s *ListResourceTypeVersionsResponseBodyResourceTypeVersions) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListResourceTypeVersionsResponseBodyResourceTypeVersions) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListResourceTypeVersionsResponseBodyResourceTypeVersions) GetVersionId() *string {
	return s.VersionId
}

func (s *ListResourceTypeVersionsResponseBodyResourceTypeVersions) SetCreateTime(v string) *ListResourceTypeVersionsResponseBodyResourceTypeVersions {
	s.CreateTime = &v
	return s
}

func (s *ListResourceTypeVersionsResponseBodyResourceTypeVersions) SetDescription(v string) *ListResourceTypeVersionsResponseBodyResourceTypeVersions {
	s.Description = &v
	return s
}

func (s *ListResourceTypeVersionsResponseBodyResourceTypeVersions) SetEntityType(v string) *ListResourceTypeVersionsResponseBodyResourceTypeVersions {
	s.EntityType = &v
	return s
}

func (s *ListResourceTypeVersionsResponseBodyResourceTypeVersions) SetIsDefaultVersion(v bool) *ListResourceTypeVersionsResponseBodyResourceTypeVersions {
	s.IsDefaultVersion = &v
	return s
}

func (s *ListResourceTypeVersionsResponseBodyResourceTypeVersions) SetProvider(v string) *ListResourceTypeVersionsResponseBodyResourceTypeVersions {
	s.Provider = &v
	return s
}

func (s *ListResourceTypeVersionsResponseBodyResourceTypeVersions) SetResourceType(v string) *ListResourceTypeVersionsResponseBodyResourceTypeVersions {
	s.ResourceType = &v
	return s
}

func (s *ListResourceTypeVersionsResponseBodyResourceTypeVersions) SetUpdateTime(v string) *ListResourceTypeVersionsResponseBodyResourceTypeVersions {
	s.UpdateTime = &v
	return s
}

func (s *ListResourceTypeVersionsResponseBodyResourceTypeVersions) SetVersionId(v string) *ListResourceTypeVersionsResponseBodyResourceTypeVersions {
	s.VersionId = &v
	return s
}

func (s *ListResourceTypeVersionsResponseBodyResourceTypeVersions) Validate() error {
	return dara.Validate(s)
}
