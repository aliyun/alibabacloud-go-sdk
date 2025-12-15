// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceGroupsWithAuthDetailsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAuthDetails(v []*ListResourceGroupsWithAuthDetailsResponseBodyAuthDetails) *ListResourceGroupsWithAuthDetailsResponseBody
	GetAuthDetails() []*ListResourceGroupsWithAuthDetailsResponseBodyAuthDetails
	SetPageNumber(v int32) *ListResourceGroupsWithAuthDetailsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListResourceGroupsWithAuthDetailsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListResourceGroupsWithAuthDetailsResponseBody
	GetRequestId() *string
	SetResourceGroups(v []*ListResourceGroupsWithAuthDetailsResponseBodyResourceGroups) *ListResourceGroupsWithAuthDetailsResponseBody
	GetResourceGroups() []*ListResourceGroupsWithAuthDetailsResponseBodyResourceGroups
	SetTotalCount(v int32) *ListResourceGroupsWithAuthDetailsResponseBody
	GetTotalCount() *int32
}

type ListResourceGroupsWithAuthDetailsResponseBody struct {
	AuthDetails []*ListResourceGroupsWithAuthDetailsResponseBodyAuthDetails `json:"AuthDetails,omitempty" xml:"AuthDetails,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The response parameters.
	//
	// example:
	//
	// 4141780B-4E3D-5D2A-A8F4-44D6D34F****
	RequestId      *string                                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceGroups []*ListResourceGroupsWithAuthDetailsResponseBodyResourceGroups `json:"ResourceGroups,omitempty" xml:"ResourceGroups,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListResourceGroupsWithAuthDetailsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListResourceGroupsWithAuthDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourceGroupsWithAuthDetailsResponseBody) GetAuthDetails() []*ListResourceGroupsWithAuthDetailsResponseBodyAuthDetails {
	return s.AuthDetails
}

func (s *ListResourceGroupsWithAuthDetailsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListResourceGroupsWithAuthDetailsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListResourceGroupsWithAuthDetailsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListResourceGroupsWithAuthDetailsResponseBody) GetResourceGroups() []*ListResourceGroupsWithAuthDetailsResponseBodyResourceGroups {
	return s.ResourceGroups
}

func (s *ListResourceGroupsWithAuthDetailsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListResourceGroupsWithAuthDetailsResponseBody) SetAuthDetails(v []*ListResourceGroupsWithAuthDetailsResponseBodyAuthDetails) *ListResourceGroupsWithAuthDetailsResponseBody {
	s.AuthDetails = v
	return s
}

func (s *ListResourceGroupsWithAuthDetailsResponseBody) SetPageNumber(v int32) *ListResourceGroupsWithAuthDetailsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListResourceGroupsWithAuthDetailsResponseBody) SetPageSize(v int32) *ListResourceGroupsWithAuthDetailsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListResourceGroupsWithAuthDetailsResponseBody) SetRequestId(v string) *ListResourceGroupsWithAuthDetailsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListResourceGroupsWithAuthDetailsResponseBody) SetResourceGroups(v []*ListResourceGroupsWithAuthDetailsResponseBodyResourceGroups) *ListResourceGroupsWithAuthDetailsResponseBody {
	s.ResourceGroups = v
	return s
}

func (s *ListResourceGroupsWithAuthDetailsResponseBody) SetTotalCount(v int32) *ListResourceGroupsWithAuthDetailsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListResourceGroupsWithAuthDetailsResponseBody) Validate() error {
	if s.AuthDetails != nil {
		for _, item := range s.AuthDetails {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ResourceGroups != nil {
		for _, item := range s.ResourceGroups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListResourceGroupsWithAuthDetailsResponseBodyAuthDetails struct {
	AccountScopeAuth     *bool                                                                           `json:"AccountScopeAuth,omitempty" xml:"AccountScopeAuth,omitempty"`
	AuthOfResourceGroups []*ListResourceGroupsWithAuthDetailsResponseBodyAuthDetailsAuthOfResourceGroups `json:"AuthOfResourceGroups,omitempty" xml:"AuthOfResourceGroups,omitempty" type:"Repeated"`
	// example:
	//
	// instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// ecs
	Service *string `json:"Service,omitempty" xml:"Service,omitempty"`
}

func (s ListResourceGroupsWithAuthDetailsResponseBodyAuthDetails) String() string {
	return dara.Prettify(s)
}

func (s ListResourceGroupsWithAuthDetailsResponseBodyAuthDetails) GoString() string {
	return s.String()
}

func (s *ListResourceGroupsWithAuthDetailsResponseBodyAuthDetails) GetAccountScopeAuth() *bool {
	return s.AccountScopeAuth
}

func (s *ListResourceGroupsWithAuthDetailsResponseBodyAuthDetails) GetAuthOfResourceGroups() []*ListResourceGroupsWithAuthDetailsResponseBodyAuthDetailsAuthOfResourceGroups {
	return s.AuthOfResourceGroups
}

func (s *ListResourceGroupsWithAuthDetailsResponseBodyAuthDetails) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListResourceGroupsWithAuthDetailsResponseBodyAuthDetails) GetService() *string {
	return s.Service
}

func (s *ListResourceGroupsWithAuthDetailsResponseBodyAuthDetails) SetAccountScopeAuth(v bool) *ListResourceGroupsWithAuthDetailsResponseBodyAuthDetails {
	s.AccountScopeAuth = &v
	return s
}

func (s *ListResourceGroupsWithAuthDetailsResponseBodyAuthDetails) SetAuthOfResourceGroups(v []*ListResourceGroupsWithAuthDetailsResponseBodyAuthDetailsAuthOfResourceGroups) *ListResourceGroupsWithAuthDetailsResponseBodyAuthDetails {
	s.AuthOfResourceGroups = v
	return s
}

func (s *ListResourceGroupsWithAuthDetailsResponseBodyAuthDetails) SetResourceType(v string) *ListResourceGroupsWithAuthDetailsResponseBodyAuthDetails {
	s.ResourceType = &v
	return s
}

func (s *ListResourceGroupsWithAuthDetailsResponseBodyAuthDetails) SetService(v string) *ListResourceGroupsWithAuthDetailsResponseBodyAuthDetails {
	s.Service = &v
	return s
}

func (s *ListResourceGroupsWithAuthDetailsResponseBodyAuthDetails) Validate() error {
	if s.AuthOfResourceGroups != nil {
		for _, item := range s.AuthOfResourceGroups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListResourceGroupsWithAuthDetailsResponseBodyAuthDetailsAuthOfResourceGroups struct {
	// example:
	//
	// true
	HasPermission *bool `json:"HasPermission,omitempty" xml:"HasPermission,omitempty"`
	// example:
	//
	// rg-aekzscexx6w3u2y
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ListResourceGroupsWithAuthDetailsResponseBodyAuthDetailsAuthOfResourceGroups) String() string {
	return dara.Prettify(s)
}

func (s ListResourceGroupsWithAuthDetailsResponseBodyAuthDetailsAuthOfResourceGroups) GoString() string {
	return s.String()
}

func (s *ListResourceGroupsWithAuthDetailsResponseBodyAuthDetailsAuthOfResourceGroups) GetHasPermission() *bool {
	return s.HasPermission
}

func (s *ListResourceGroupsWithAuthDetailsResponseBodyAuthDetailsAuthOfResourceGroups) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListResourceGroupsWithAuthDetailsResponseBodyAuthDetailsAuthOfResourceGroups) SetHasPermission(v bool) *ListResourceGroupsWithAuthDetailsResponseBodyAuthDetailsAuthOfResourceGroups {
	s.HasPermission = &v
	return s
}

func (s *ListResourceGroupsWithAuthDetailsResponseBodyAuthDetailsAuthOfResourceGroups) SetResourceGroupId(v string) *ListResourceGroupsWithAuthDetailsResponseBodyAuthDetailsAuthOfResourceGroups {
	s.ResourceGroupId = &v
	return s
}

func (s *ListResourceGroupsWithAuthDetailsResponseBodyAuthDetailsAuthOfResourceGroups) Validate() error {
	return dara.Validate(s)
}

type ListResourceGroupsWithAuthDetailsResponseBodyResourceGroups struct {
	// example:
	//
	// 123456789****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// example:
	//
	// 2015-01-23T12:33:18Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// example:
	//
	// my-project
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// rg-9gLOoK****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// my-project
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// OK
	Status *string                                                            `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags   []*ListResourceGroupsWithAuthDetailsResponseBodyResourceGroupsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListResourceGroupsWithAuthDetailsResponseBodyResourceGroups) String() string {
	return dara.Prettify(s)
}

func (s ListResourceGroupsWithAuthDetailsResponseBodyResourceGroups) GoString() string {
	return s.String()
}

func (s *ListResourceGroupsWithAuthDetailsResponseBodyResourceGroups) GetAccountId() *string {
	return s.AccountId
}

func (s *ListResourceGroupsWithAuthDetailsResponseBodyResourceGroups) GetCreateDate() *string {
	return s.CreateDate
}

func (s *ListResourceGroupsWithAuthDetailsResponseBodyResourceGroups) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListResourceGroupsWithAuthDetailsResponseBodyResourceGroups) GetId() *string {
	return s.Id
}

func (s *ListResourceGroupsWithAuthDetailsResponseBodyResourceGroups) GetName() *string {
	return s.Name
}

func (s *ListResourceGroupsWithAuthDetailsResponseBodyResourceGroups) GetStatus() *string {
	return s.Status
}

func (s *ListResourceGroupsWithAuthDetailsResponseBodyResourceGroups) GetTags() []*ListResourceGroupsWithAuthDetailsResponseBodyResourceGroupsTags {
	return s.Tags
}

func (s *ListResourceGroupsWithAuthDetailsResponseBodyResourceGroups) SetAccountId(v string) *ListResourceGroupsWithAuthDetailsResponseBodyResourceGroups {
	s.AccountId = &v
	return s
}

func (s *ListResourceGroupsWithAuthDetailsResponseBodyResourceGroups) SetCreateDate(v string) *ListResourceGroupsWithAuthDetailsResponseBodyResourceGroups {
	s.CreateDate = &v
	return s
}

func (s *ListResourceGroupsWithAuthDetailsResponseBodyResourceGroups) SetDisplayName(v string) *ListResourceGroupsWithAuthDetailsResponseBodyResourceGroups {
	s.DisplayName = &v
	return s
}

func (s *ListResourceGroupsWithAuthDetailsResponseBodyResourceGroups) SetId(v string) *ListResourceGroupsWithAuthDetailsResponseBodyResourceGroups {
	s.Id = &v
	return s
}

func (s *ListResourceGroupsWithAuthDetailsResponseBodyResourceGroups) SetName(v string) *ListResourceGroupsWithAuthDetailsResponseBodyResourceGroups {
	s.Name = &v
	return s
}

func (s *ListResourceGroupsWithAuthDetailsResponseBodyResourceGroups) SetStatus(v string) *ListResourceGroupsWithAuthDetailsResponseBodyResourceGroups {
	s.Status = &v
	return s
}

func (s *ListResourceGroupsWithAuthDetailsResponseBodyResourceGroups) SetTags(v []*ListResourceGroupsWithAuthDetailsResponseBodyResourceGroupsTags) *ListResourceGroupsWithAuthDetailsResponseBodyResourceGroups {
	s.Tags = v
	return s
}

func (s *ListResourceGroupsWithAuthDetailsResponseBodyResourceGroups) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListResourceGroupsWithAuthDetailsResponseBodyResourceGroupsTags struct {
	// example:
	//
	// k1
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// example:
	//
	// v1
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListResourceGroupsWithAuthDetailsResponseBodyResourceGroupsTags) String() string {
	return dara.Prettify(s)
}

func (s ListResourceGroupsWithAuthDetailsResponseBodyResourceGroupsTags) GoString() string {
	return s.String()
}

func (s *ListResourceGroupsWithAuthDetailsResponseBodyResourceGroupsTags) GetTagKey() *string {
	return s.TagKey
}

func (s *ListResourceGroupsWithAuthDetailsResponseBodyResourceGroupsTags) GetTagValue() *string {
	return s.TagValue
}

func (s *ListResourceGroupsWithAuthDetailsResponseBodyResourceGroupsTags) SetTagKey(v string) *ListResourceGroupsWithAuthDetailsResponseBodyResourceGroupsTags {
	s.TagKey = &v
	return s
}

func (s *ListResourceGroupsWithAuthDetailsResponseBodyResourceGroupsTags) SetTagValue(v string) *ListResourceGroupsWithAuthDetailsResponseBodyResourceGroupsTags {
	s.TagValue = &v
	return s
}

func (s *ListResourceGroupsWithAuthDetailsResponseBodyResourceGroupsTags) Validate() error {
	return dara.Validate(s)
}
