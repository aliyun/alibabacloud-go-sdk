// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAuthorizedUserGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppInstanceGroupId(v string) *ListAuthorizedUserGroupsRequest
	GetAppInstanceGroupId() *string
	SetGroupId(v string) *ListAuthorizedUserGroupsRequest
	GetGroupId() *string
	SetGroupName(v string) *ListAuthorizedUserGroupsRequest
	GetGroupName() *string
	SetPageNumber(v int32) *ListAuthorizedUserGroupsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListAuthorizedUserGroupsRequest
	GetPageSize() *int32
	SetProductType(v string) *ListAuthorizedUserGroupsRequest
	GetProductType() *string
}

type ListAuthorizedUserGroupsRequest struct {
	// The ID of the delivery group.
	//
	// This parameter is required.
	//
	// example:
	//
	// aig-9ciijz60n4xsv****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// The ID of the user group (exact match).
	//
	// example:
	//
	// ug-00001
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the user group (fuzzy match).
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The page number. Pages start from page 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The maximum number of entries returned per page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The product type.
	//
	// Valid values:
	//
	// 	- CloudApp: App Streaming
	//
	// This parameter is required.
	//
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
}

func (s ListAuthorizedUserGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorizedUserGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListAuthorizedUserGroupsRequest) GetAppInstanceGroupId() *string {
	return s.AppInstanceGroupId
}

func (s *ListAuthorizedUserGroupsRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *ListAuthorizedUserGroupsRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *ListAuthorizedUserGroupsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAuthorizedUserGroupsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAuthorizedUserGroupsRequest) GetProductType() *string {
	return s.ProductType
}

func (s *ListAuthorizedUserGroupsRequest) SetAppInstanceGroupId(v string) *ListAuthorizedUserGroupsRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *ListAuthorizedUserGroupsRequest) SetGroupId(v string) *ListAuthorizedUserGroupsRequest {
	s.GroupId = &v
	return s
}

func (s *ListAuthorizedUserGroupsRequest) SetGroupName(v string) *ListAuthorizedUserGroupsRequest {
	s.GroupName = &v
	return s
}

func (s *ListAuthorizedUserGroupsRequest) SetPageNumber(v int32) *ListAuthorizedUserGroupsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAuthorizedUserGroupsRequest) SetPageSize(v int32) *ListAuthorizedUserGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *ListAuthorizedUserGroupsRequest) SetProductType(v string) *ListAuthorizedUserGroupsRequest {
	s.ProductType = &v
	return s
}

func (s *ListAuthorizedUserGroupsRequest) Validate() error {
	return dara.Validate(s)
}
