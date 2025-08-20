// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceTypesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListResourceTypesResponseBody
	GetRequestId() *string
	SetResourceTypeSummaries(v []*ListResourceTypesResponseBodyResourceTypeSummaries) *ListResourceTypesResponseBody
	GetResourceTypeSummaries() []*ListResourceTypesResponseBodyResourceTypeSummaries
	SetResourceTypes(v []*string) *ListResourceTypesResponseBody
	GetResourceTypes() []*string
}

type ListResourceTypesResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// EA00860C-ECAF-5253-A1F9-8198695A7157
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The resource type summaries.
	ResourceTypeSummaries []*ListResourceTypesResponseBodyResourceTypeSummaries `json:"ResourceTypeSummaries,omitempty" xml:"ResourceTypeSummaries,omitempty" type:"Repeated"`
	// The array of resource types.
	ResourceTypes []*string `json:"ResourceTypes,omitempty" xml:"ResourceTypes,omitempty" type:"Repeated"`
}

func (s ListResourceTypesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListResourceTypesResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourceTypesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListResourceTypesResponseBody) GetResourceTypeSummaries() []*ListResourceTypesResponseBodyResourceTypeSummaries {
	return s.ResourceTypeSummaries
}

func (s *ListResourceTypesResponseBody) GetResourceTypes() []*string {
	return s.ResourceTypes
}

func (s *ListResourceTypesResponseBody) SetRequestId(v string) *ListResourceTypesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListResourceTypesResponseBody) SetResourceTypeSummaries(v []*ListResourceTypesResponseBodyResourceTypeSummaries) *ListResourceTypesResponseBody {
	s.ResourceTypeSummaries = v
	return s
}

func (s *ListResourceTypesResponseBody) SetResourceTypes(v []*string) *ListResourceTypesResponseBody {
	s.ResourceTypes = v
	return s
}

func (s *ListResourceTypesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListResourceTypesResponseBodyResourceTypeSummaries struct {
	// The creation time. The time is displayed in UTC. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ss format.
	//
	// example:
	//
	// 2023-02-24T08:25:21
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the default version.
	//
	// example:
	//
	// v1
	DefaultVersionId *string `json:"DefaultVersionId,omitempty" xml:"DefaultVersionId,omitempty"`
	// The description of the resource type.
	//
	// example:
	//
	// It is a demo.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The entity type. Valid values:
	//
	// 	- Resource: regular resources.
	//
	// 	- DataSource: DataSource resources.
	//
	// 	- Module: modules.
	//
	// example:
	//
	// Module
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// The ID of the latest version.
	//
	// example:
	//
	// v10
	LatestVersionId *string `json:"LatestVersionId,omitempty" xml:"LatestVersionId,omitempty"`
	// The provider of the resource type. Valid values:
	//
	// 	- ROS: The resource type is provided by ROS.
	//
	// 	- Self: The resource type is provided by you.
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
	// The number of versions.
	//
	// example:
	//
	// 10
	TotalVersionCount *int32 `json:"TotalVersionCount,omitempty" xml:"TotalVersionCount,omitempty"`
	// The update time. The time is displayed in UTC. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ss format.
	//
	// example:
	//
	// 2023-02-24T08:25:21
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListResourceTypesResponseBodyResourceTypeSummaries) String() string {
	return dara.Prettify(s)
}

func (s ListResourceTypesResponseBodyResourceTypeSummaries) GoString() string {
	return s.String()
}

func (s *ListResourceTypesResponseBodyResourceTypeSummaries) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListResourceTypesResponseBodyResourceTypeSummaries) GetDefaultVersionId() *string {
	return s.DefaultVersionId
}

func (s *ListResourceTypesResponseBodyResourceTypeSummaries) GetDescription() *string {
	return s.Description
}

func (s *ListResourceTypesResponseBodyResourceTypeSummaries) GetEntityType() *string {
	return s.EntityType
}

func (s *ListResourceTypesResponseBodyResourceTypeSummaries) GetLatestVersionId() *string {
	return s.LatestVersionId
}

func (s *ListResourceTypesResponseBodyResourceTypeSummaries) GetProvider() *string {
	return s.Provider
}

func (s *ListResourceTypesResponseBodyResourceTypeSummaries) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListResourceTypesResponseBodyResourceTypeSummaries) GetTotalVersionCount() *int32 {
	return s.TotalVersionCount
}

func (s *ListResourceTypesResponseBodyResourceTypeSummaries) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListResourceTypesResponseBodyResourceTypeSummaries) SetCreateTime(v string) *ListResourceTypesResponseBodyResourceTypeSummaries {
	s.CreateTime = &v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypeSummaries) SetDefaultVersionId(v string) *ListResourceTypesResponseBodyResourceTypeSummaries {
	s.DefaultVersionId = &v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypeSummaries) SetDescription(v string) *ListResourceTypesResponseBodyResourceTypeSummaries {
	s.Description = &v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypeSummaries) SetEntityType(v string) *ListResourceTypesResponseBodyResourceTypeSummaries {
	s.EntityType = &v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypeSummaries) SetLatestVersionId(v string) *ListResourceTypesResponseBodyResourceTypeSummaries {
	s.LatestVersionId = &v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypeSummaries) SetProvider(v string) *ListResourceTypesResponseBodyResourceTypeSummaries {
	s.Provider = &v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypeSummaries) SetResourceType(v string) *ListResourceTypesResponseBodyResourceTypeSummaries {
	s.ResourceType = &v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypeSummaries) SetTotalVersionCount(v int32) *ListResourceTypesResponseBodyResourceTypeSummaries {
	s.TotalVersionCount = &v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypeSummaries) SetUpdateTime(v string) *ListResourceTypesResponseBodyResourceTypeSummaries {
	s.UpdateTime = &v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypeSummaries) Validate() error {
	return dara.Validate(s)
}
