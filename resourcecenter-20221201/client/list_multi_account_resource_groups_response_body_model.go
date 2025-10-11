// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMultiAccountResourceGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *ListMultiAccountResourceGroupsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListMultiAccountResourceGroupsResponseBody
	GetRequestId() *string
	SetResourceGroups(v []*ListMultiAccountResourceGroupsResponseBodyResourceGroups) *ListMultiAccountResourceGroupsResponseBody
	GetResourceGroups() []*ListMultiAccountResourceGroupsResponseBodyResourceGroups
}

type ListMultiAccountResourceGroupsResponseBody struct {
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// AAAAAU5VsT9R1adMTuz9GzginZ3Y+7Y/5JATS+6q5GK9kT75
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 0FF0A66E-781F-51EE-9531-928F197558F2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the resource groups.
	ResourceGroups []*ListMultiAccountResourceGroupsResponseBodyResourceGroups `json:"ResourceGroups,omitempty" xml:"ResourceGroups,omitempty" type:"Repeated"`
}

func (s ListMultiAccountResourceGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMultiAccountResourceGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListMultiAccountResourceGroupsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListMultiAccountResourceGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMultiAccountResourceGroupsResponseBody) GetResourceGroups() []*ListMultiAccountResourceGroupsResponseBodyResourceGroups {
	return s.ResourceGroups
}

func (s *ListMultiAccountResourceGroupsResponseBody) SetNextToken(v string) *ListMultiAccountResourceGroupsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListMultiAccountResourceGroupsResponseBody) SetRequestId(v string) *ListMultiAccountResourceGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMultiAccountResourceGroupsResponseBody) SetResourceGroups(v []*ListMultiAccountResourceGroupsResponseBodyResourceGroups) *ListMultiAccountResourceGroupsResponseBody {
	s.ResourceGroups = v
	return s
}

func (s *ListMultiAccountResourceGroupsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListMultiAccountResourceGroupsResponseBodyResourceGroups struct {
	// The ID of the management account or member of the resource directory.
	//
	// example:
	//
	// 1394339739****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The time when the resource group was created.
	//
	// example:
	//
	// 2021-06-30T09:20:08Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The display name of the resource group.
	//
	// example:
	//
	// group1
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmzawhxxc****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The unique identifier of the resource group.
	//
	// example:
	//
	// my-project
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The status of the resource group. Valid values:
	//
	// 	- Creating: The resource group is being created.
	//
	// 	- OK: The resource group is created.
	//
	// 	- PendingDelete: The resource group is waiting to be deleted.
	//
	// example:
	//
	// OK
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListMultiAccountResourceGroupsResponseBodyResourceGroups) String() string {
	return dara.Prettify(s)
}

func (s ListMultiAccountResourceGroupsResponseBodyResourceGroups) GoString() string {
	return s.String()
}

func (s *ListMultiAccountResourceGroupsResponseBodyResourceGroups) GetAccountId() *string {
	return s.AccountId
}

func (s *ListMultiAccountResourceGroupsResponseBodyResourceGroups) GetCreateDate() *string {
	return s.CreateDate
}

func (s *ListMultiAccountResourceGroupsResponseBodyResourceGroups) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListMultiAccountResourceGroupsResponseBodyResourceGroups) GetId() *string {
	return s.Id
}

func (s *ListMultiAccountResourceGroupsResponseBodyResourceGroups) GetName() *string {
	return s.Name
}

func (s *ListMultiAccountResourceGroupsResponseBodyResourceGroups) GetStatus() *string {
	return s.Status
}

func (s *ListMultiAccountResourceGroupsResponseBodyResourceGroups) SetAccountId(v string) *ListMultiAccountResourceGroupsResponseBodyResourceGroups {
	s.AccountId = &v
	return s
}

func (s *ListMultiAccountResourceGroupsResponseBodyResourceGroups) SetCreateDate(v string) *ListMultiAccountResourceGroupsResponseBodyResourceGroups {
	s.CreateDate = &v
	return s
}

func (s *ListMultiAccountResourceGroupsResponseBodyResourceGroups) SetDisplayName(v string) *ListMultiAccountResourceGroupsResponseBodyResourceGroups {
	s.DisplayName = &v
	return s
}

func (s *ListMultiAccountResourceGroupsResponseBodyResourceGroups) SetId(v string) *ListMultiAccountResourceGroupsResponseBodyResourceGroups {
	s.Id = &v
	return s
}

func (s *ListMultiAccountResourceGroupsResponseBodyResourceGroups) SetName(v string) *ListMultiAccountResourceGroupsResponseBodyResourceGroups {
	s.Name = &v
	return s
}

func (s *ListMultiAccountResourceGroupsResponseBodyResourceGroups) SetStatus(v string) *ListMultiAccountResourceGroupsResponseBodyResourceGroups {
	s.Status = &v
	return s
}

func (s *ListMultiAccountResourceGroupsResponseBodyResourceGroups) Validate() error {
	return dara.Validate(s)
}
