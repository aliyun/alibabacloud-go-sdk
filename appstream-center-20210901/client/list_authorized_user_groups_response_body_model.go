// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAuthorizedUserGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListAuthorizedUserGroupsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListAuthorizedUserGroupsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListAuthorizedUserGroupsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListAuthorizedUserGroupsResponseBody
	GetTotalCount() *int32
	SetUserGroups(v []*ListAuthorizedUserGroupsResponseBodyUserGroups) *ListAuthorizedUserGroupsResponseBody
	GetUserGroups() []*ListAuthorizedUserGroupsResponseBodyUserGroups
}

type ListAuthorizedUserGroupsResponseBody struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 6
	TotalCount *int32                                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	UserGroups []*ListAuthorizedUserGroupsResponseBodyUserGroups `json:"UserGroups,omitempty" xml:"UserGroups,omitempty" type:"Repeated"`
}

func (s ListAuthorizedUserGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorizedUserGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAuthorizedUserGroupsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAuthorizedUserGroupsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAuthorizedUserGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAuthorizedUserGroupsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListAuthorizedUserGroupsResponseBody) GetUserGroups() []*ListAuthorizedUserGroupsResponseBodyUserGroups {
	return s.UserGroups
}

func (s *ListAuthorizedUserGroupsResponseBody) SetPageNumber(v int32) *ListAuthorizedUserGroupsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListAuthorizedUserGroupsResponseBody) SetPageSize(v int32) *ListAuthorizedUserGroupsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListAuthorizedUserGroupsResponseBody) SetRequestId(v string) *ListAuthorizedUserGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAuthorizedUserGroupsResponseBody) SetTotalCount(v int32) *ListAuthorizedUserGroupsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListAuthorizedUserGroupsResponseBody) SetUserGroups(v []*ListAuthorizedUserGroupsResponseBodyUserGroups) *ListAuthorizedUserGroupsResponseBody {
	s.UserGroups = v
	return s
}

func (s *ListAuthorizedUserGroupsResponseBody) Validate() error {
	if s.UserGroups != nil {
		for _, item := range s.UserGroups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAuthorizedUserGroupsResponseBodyUserGroups struct {
	// example:
	//
	// aig-9ciijz60n4xsv****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// example:
	//
	// App
	AuthMode *string `json:"AuthMode,omitempty" xml:"AuthMode,omitempty"`
	// example:
	//
	// ug-00001
	GroupId   *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
}

func (s ListAuthorizedUserGroupsResponseBodyUserGroups) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorizedUserGroupsResponseBodyUserGroups) GoString() string {
	return s.String()
}

func (s *ListAuthorizedUserGroupsResponseBodyUserGroups) GetAppInstanceGroupId() *string {
	return s.AppInstanceGroupId
}

func (s *ListAuthorizedUserGroupsResponseBodyUserGroups) GetAuthMode() *string {
	return s.AuthMode
}

func (s *ListAuthorizedUserGroupsResponseBodyUserGroups) GetGroupId() *string {
	return s.GroupId
}

func (s *ListAuthorizedUserGroupsResponseBodyUserGroups) GetGroupName() *string {
	return s.GroupName
}

func (s *ListAuthorizedUserGroupsResponseBodyUserGroups) SetAppInstanceGroupId(v string) *ListAuthorizedUserGroupsResponseBodyUserGroups {
	s.AppInstanceGroupId = &v
	return s
}

func (s *ListAuthorizedUserGroupsResponseBodyUserGroups) SetAuthMode(v string) *ListAuthorizedUserGroupsResponseBodyUserGroups {
	s.AuthMode = &v
	return s
}

func (s *ListAuthorizedUserGroupsResponseBodyUserGroups) SetGroupId(v string) *ListAuthorizedUserGroupsResponseBodyUserGroups {
	s.GroupId = &v
	return s
}

func (s *ListAuthorizedUserGroupsResponseBodyUserGroups) SetGroupName(v string) *ListAuthorizedUserGroupsResponseBodyUserGroups {
	s.GroupName = &v
	return s
}

func (s *ListAuthorizedUserGroupsResponseBodyUserGroups) Validate() error {
	return dara.Validate(s)
}
