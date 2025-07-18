// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPrivateAccessTagsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *ListPrivateAccessTagsRequest
	GetApplicationId() *string
	SetCurrentPage(v int32) *ListPrivateAccessTagsRequest
	GetCurrentPage() *int32
	SetName(v string) *ListPrivateAccessTagsRequest
	GetName() *string
	SetPageSize(v int32) *ListPrivateAccessTagsRequest
	GetPageSize() *int32
	SetPolicyId(v string) *ListPrivateAccessTagsRequest
	GetPolicyId() *string
	SetSimpleMode(v bool) *ListPrivateAccessTagsRequest
	GetSimpleMode() *bool
	SetTagIds(v []*string) *ListPrivateAccessTagsRequest
	GetTagIds() []*string
}

type ListPrivateAccessTagsRequest struct {
	// The ID of the internal access application. You can obtain the application ID by calling the following operations:
	//
	// 	- [ListPrivateAccessApplications](~~ListPrivateAccessApplications~~): queries all internal access applications.
	//
	// 	- [CreatePrivateAccessApplication](~~CreatePrivateAccessApplication~~): creates an internal access application.
	//
	// example:
	//
	// pa-application-e12860ef6c48****
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The page number. Valid values: 1 to 10000.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The name of the internal access tag. The name must be 1 to 128 characters in length and can contain letters, digits, periods (.), underscores (_), and hyphens (-).
	//
	// example:
	//
	// tag_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The number of entries per page. Valid values: 1 to 1000.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the internal access policy. You can obtain the policy ID by calling the following operations:
	//
	// 	- [ListPrivateAccessPolices](~~ListPrivateAccessPolices~~): queries all internal access policies.
	//
	// 	- [CreatePrivateAccessPolicy](~~CreatePrivateAccessPolicy~~): creates an internal access policy.
	//
	// example:
	//
	// pa-policy-54a7838a48bf****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// Specifies whether to enable the simple query mode. A value of true specifies that policy IDs are not queried.
	//
	// example:
	//
	// true
	SimpleMode *bool `json:"SimpleMode,omitempty" xml:"SimpleMode,omitempty"`
	// The IDs of internal access tags. You can specify up to 100 tag IDs.
	TagIds []*string `json:"TagIds,omitempty" xml:"TagIds,omitempty" type:"Repeated"`
}

func (s ListPrivateAccessTagsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPrivateAccessTagsRequest) GoString() string {
	return s.String()
}

func (s *ListPrivateAccessTagsRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *ListPrivateAccessTagsRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListPrivateAccessTagsRequest) GetName() *string {
	return s.Name
}

func (s *ListPrivateAccessTagsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPrivateAccessTagsRequest) GetPolicyId() *string {
	return s.PolicyId
}

func (s *ListPrivateAccessTagsRequest) GetSimpleMode() *bool {
	return s.SimpleMode
}

func (s *ListPrivateAccessTagsRequest) GetTagIds() []*string {
	return s.TagIds
}

func (s *ListPrivateAccessTagsRequest) SetApplicationId(v string) *ListPrivateAccessTagsRequest {
	s.ApplicationId = &v
	return s
}

func (s *ListPrivateAccessTagsRequest) SetCurrentPage(v int32) *ListPrivateAccessTagsRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListPrivateAccessTagsRequest) SetName(v string) *ListPrivateAccessTagsRequest {
	s.Name = &v
	return s
}

func (s *ListPrivateAccessTagsRequest) SetPageSize(v int32) *ListPrivateAccessTagsRequest {
	s.PageSize = &v
	return s
}

func (s *ListPrivateAccessTagsRequest) SetPolicyId(v string) *ListPrivateAccessTagsRequest {
	s.PolicyId = &v
	return s
}

func (s *ListPrivateAccessTagsRequest) SetSimpleMode(v bool) *ListPrivateAccessTagsRequest {
	s.SimpleMode = &v
	return s
}

func (s *ListPrivateAccessTagsRequest) SetTagIds(v []*string) *ListPrivateAccessTagsRequest {
	s.TagIds = v
	return s
}

func (s *ListPrivateAccessTagsRequest) Validate() error {
	return dara.Validate(s)
}
