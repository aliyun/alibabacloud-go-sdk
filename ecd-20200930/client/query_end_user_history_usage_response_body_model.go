// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryEndUserHistoryUsageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *QueryEndUserHistoryUsageResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *QueryEndUserHistoryUsageResponseBody
	GetTotalCount() *int64
	SetUserUsageInfoList(v []*QueryEndUserHistoryUsageResponseBodyUserUsageInfoList) *QueryEndUserHistoryUsageResponseBody
	GetUserUsageInfoList() []*QueryEndUserHistoryUsageResponseBodyUserUsageInfoList
}

type QueryEndUserHistoryUsageResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 269BDB16-2CD8-4865-84BD-11C40BC21DB0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of users that meet the query conditions.
	//
	// example:
	//
	// 50
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The list of user usage duration entries on the current page.
	UserUsageInfoList []*QueryEndUserHistoryUsageResponseBodyUserUsageInfoList `json:"UserUsageInfoList,omitempty" xml:"UserUsageInfoList,omitempty" type:"Repeated"`
}

func (s QueryEndUserHistoryUsageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryEndUserHistoryUsageResponseBody) GoString() string {
	return s.String()
}

func (s *QueryEndUserHistoryUsageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryEndUserHistoryUsageResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *QueryEndUserHistoryUsageResponseBody) GetUserUsageInfoList() []*QueryEndUserHistoryUsageResponseBodyUserUsageInfoList {
	return s.UserUsageInfoList
}

func (s *QueryEndUserHistoryUsageResponseBody) SetRequestId(v string) *QueryEndUserHistoryUsageResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryEndUserHistoryUsageResponseBody) SetTotalCount(v int64) *QueryEndUserHistoryUsageResponseBody {
	s.TotalCount = &v
	return s
}

func (s *QueryEndUserHistoryUsageResponseBody) SetUserUsageInfoList(v []*QueryEndUserHistoryUsageResponseBodyUserUsageInfoList) *QueryEndUserHistoryUsageResponseBody {
	s.UserUsageInfoList = v
	return s
}

func (s *QueryEndUserHistoryUsageResponseBody) Validate() error {
	if s.UserUsageInfoList != nil {
		for _, item := range s.UserUsageInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryEndUserHistoryUsageResponseBodyUserUsageInfoList struct {
	// The remarks of the user. This parameter has a value only for convenience account users.
	//
	// example:
	//
	// Test user
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The list of usage duration details for each desktop.
	DesktopUsageList []*QueryEndUserHistoryUsageResponseBodyUserUsageInfoListDesktopUsageList `json:"DesktopUsageList,omitempty" xml:"DesktopUsageList,omitempty" type:"Repeated"`
	// The display name of the user. For convenience account users, this is the actual nickname. For AD users, this is the display name.
	//
	// example:
	//
	// John Smith
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The total usage duration, in seconds.
	//
	// example:
	//
	// 3600
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The end user ID.
	//
	// example:
	//
	// user01
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The username. For convenience account users, this is the nickname. For AD users, this is the UserPrincipalName.
	//
	// example:
	//
	// zhangsan
	EndUserName *string `json:"EndUserName,omitempty" xml:"EndUserName,omitempty"`
	// The list of organization paths. For convenience account users, this contains multiple organization paths. For AD users, this is the organizational unit (OU) path.
	OrgPathList []*string `json:"OrgPathList,omitempty" xml:"OrgPathList,omitempty" type:"Repeated"`
	// The list of user groups. This parameter has a value only for convenience account users.
	UserGroupList []*QueryEndUserHistoryUsageResponseBodyUserUsageInfoListUserGroupList `json:"UserGroupList,omitempty" xml:"UserGroupList,omitempty" type:"Repeated"`
}

func (s QueryEndUserHistoryUsageResponseBodyUserUsageInfoList) String() string {
	return dara.Prettify(s)
}

func (s QueryEndUserHistoryUsageResponseBodyUserUsageInfoList) GoString() string {
	return s.String()
}

func (s *QueryEndUserHistoryUsageResponseBodyUserUsageInfoList) GetDescription() *string {
	return s.Description
}

func (s *QueryEndUserHistoryUsageResponseBodyUserUsageInfoList) GetDesktopUsageList() []*QueryEndUserHistoryUsageResponseBodyUserUsageInfoListDesktopUsageList {
	return s.DesktopUsageList
}

func (s *QueryEndUserHistoryUsageResponseBodyUserUsageInfoList) GetDisplayName() *string {
	return s.DisplayName
}

func (s *QueryEndUserHistoryUsageResponseBodyUserUsageInfoList) GetDuration() *int64 {
	return s.Duration
}

func (s *QueryEndUserHistoryUsageResponseBodyUserUsageInfoList) GetEndUserId() *string {
	return s.EndUserId
}

func (s *QueryEndUserHistoryUsageResponseBodyUserUsageInfoList) GetEndUserName() *string {
	return s.EndUserName
}

func (s *QueryEndUserHistoryUsageResponseBodyUserUsageInfoList) GetOrgPathList() []*string {
	return s.OrgPathList
}

func (s *QueryEndUserHistoryUsageResponseBodyUserUsageInfoList) GetUserGroupList() []*QueryEndUserHistoryUsageResponseBodyUserUsageInfoListUserGroupList {
	return s.UserGroupList
}

func (s *QueryEndUserHistoryUsageResponseBodyUserUsageInfoList) SetDescription(v string) *QueryEndUserHistoryUsageResponseBodyUserUsageInfoList {
	s.Description = &v
	return s
}

func (s *QueryEndUserHistoryUsageResponseBodyUserUsageInfoList) SetDesktopUsageList(v []*QueryEndUserHistoryUsageResponseBodyUserUsageInfoListDesktopUsageList) *QueryEndUserHistoryUsageResponseBodyUserUsageInfoList {
	s.DesktopUsageList = v
	return s
}

func (s *QueryEndUserHistoryUsageResponseBodyUserUsageInfoList) SetDisplayName(v string) *QueryEndUserHistoryUsageResponseBodyUserUsageInfoList {
	s.DisplayName = &v
	return s
}

func (s *QueryEndUserHistoryUsageResponseBodyUserUsageInfoList) SetDuration(v int64) *QueryEndUserHistoryUsageResponseBodyUserUsageInfoList {
	s.Duration = &v
	return s
}

func (s *QueryEndUserHistoryUsageResponseBodyUserUsageInfoList) SetEndUserId(v string) *QueryEndUserHistoryUsageResponseBodyUserUsageInfoList {
	s.EndUserId = &v
	return s
}

func (s *QueryEndUserHistoryUsageResponseBodyUserUsageInfoList) SetEndUserName(v string) *QueryEndUserHistoryUsageResponseBodyUserUsageInfoList {
	s.EndUserName = &v
	return s
}

func (s *QueryEndUserHistoryUsageResponseBodyUserUsageInfoList) SetOrgPathList(v []*string) *QueryEndUserHistoryUsageResponseBodyUserUsageInfoList {
	s.OrgPathList = v
	return s
}

func (s *QueryEndUserHistoryUsageResponseBodyUserUsageInfoList) SetUserGroupList(v []*QueryEndUserHistoryUsageResponseBodyUserUsageInfoListUserGroupList) *QueryEndUserHistoryUsageResponseBodyUserUsageInfoList {
	s.UserGroupList = v
	return s
}

func (s *QueryEndUserHistoryUsageResponseBodyUserUsageInfoList) Validate() error {
	if s.DesktopUsageList != nil {
		for _, item := range s.DesktopUsageList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.UserGroupList != nil {
		for _, item := range s.UserGroupList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryEndUserHistoryUsageResponseBodyUserUsageInfoListDesktopUsageList struct {
	// The desktop ID.
	//
	// example:
	//
	// ecd-abc123
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	// The desktop name.
	//
	// example:
	//
	// DemoDesktop
	DesktopName *string `json:"DesktopName,omitempty" xml:"DesktopName,omitempty"`
	// The usage duration of the user on the desktop, in seconds.
	//
	// example:
	//
	// 1800
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
}

func (s QueryEndUserHistoryUsageResponseBodyUserUsageInfoListDesktopUsageList) String() string {
	return dara.Prettify(s)
}

func (s QueryEndUserHistoryUsageResponseBodyUserUsageInfoListDesktopUsageList) GoString() string {
	return s.String()
}

func (s *QueryEndUserHistoryUsageResponseBodyUserUsageInfoListDesktopUsageList) GetDesktopId() *string {
	return s.DesktopId
}

func (s *QueryEndUserHistoryUsageResponseBodyUserUsageInfoListDesktopUsageList) GetDesktopName() *string {
	return s.DesktopName
}

func (s *QueryEndUserHistoryUsageResponseBodyUserUsageInfoListDesktopUsageList) GetDuration() *int64 {
	return s.Duration
}

func (s *QueryEndUserHistoryUsageResponseBodyUserUsageInfoListDesktopUsageList) SetDesktopId(v string) *QueryEndUserHistoryUsageResponseBodyUserUsageInfoListDesktopUsageList {
	s.DesktopId = &v
	return s
}

func (s *QueryEndUserHistoryUsageResponseBodyUserUsageInfoListDesktopUsageList) SetDesktopName(v string) *QueryEndUserHistoryUsageResponseBodyUserUsageInfoListDesktopUsageList {
	s.DesktopName = &v
	return s
}

func (s *QueryEndUserHistoryUsageResponseBodyUserUsageInfoListDesktopUsageList) SetDuration(v int64) *QueryEndUserHistoryUsageResponseBodyUserUsageInfoListDesktopUsageList {
	s.Duration = &v
	return s
}

func (s *QueryEndUserHistoryUsageResponseBodyUserUsageInfoListDesktopUsageList) Validate() error {
	return dara.Validate(s)
}

type QueryEndUserHistoryUsageResponseBodyUserUsageInfoListUserGroupList struct {
	// The user group ID.
	//
	// example:
	//
	// ug-12345678
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	// The user group name.
	//
	// example:
	//
	// Default user group
	UserGroupName *string `json:"UserGroupName,omitempty" xml:"UserGroupName,omitempty"`
}

func (s QueryEndUserHistoryUsageResponseBodyUserUsageInfoListUserGroupList) String() string {
	return dara.Prettify(s)
}

func (s QueryEndUserHistoryUsageResponseBodyUserUsageInfoListUserGroupList) GoString() string {
	return s.String()
}

func (s *QueryEndUserHistoryUsageResponseBodyUserUsageInfoListUserGroupList) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *QueryEndUserHistoryUsageResponseBodyUserUsageInfoListUserGroupList) GetUserGroupName() *string {
	return s.UserGroupName
}

func (s *QueryEndUserHistoryUsageResponseBodyUserUsageInfoListUserGroupList) SetUserGroupId(v string) *QueryEndUserHistoryUsageResponseBodyUserUsageInfoListUserGroupList {
	s.UserGroupId = &v
	return s
}

func (s *QueryEndUserHistoryUsageResponseBodyUserUsageInfoListUserGroupList) SetUserGroupName(v string) *QueryEndUserHistoryUsageResponseBodyUserUsageInfoListUserGroupList {
	s.UserGroupName = &v
	return s
}

func (s *QueryEndUserHistoryUsageResponseBodyUserUsageInfoListUserGroupList) Validate() error {
	return dara.Validate(s)
}
