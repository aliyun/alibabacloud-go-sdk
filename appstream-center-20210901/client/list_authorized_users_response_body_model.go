// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAuthorizedUsersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListAuthorizedUsersResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListAuthorizedUsersResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListAuthorizedUsersResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListAuthorizedUsersResponseBody
	GetTotalCount() *int32
	SetUsers(v []*ListAuthorizedUsersResponseBodyUsers) *ListAuthorizedUsersResponseBody
	GetUsers() []*ListAuthorizedUsersResponseBodyUsers
}

type ListAuthorizedUsersResponseBody struct {
	// The current page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of records per page in this request.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID, which is used to locate this call.
	//
	// example:
	//
	// 5C1A4F2D-713A-5C98-8AF6-1B5D0868****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of authorization records that match the query conditions.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The list of authorized users on the current page. An empty list is returned if no authorization records are matched.
	Users []*ListAuthorizedUsersResponseBodyUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s ListAuthorizedUsersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorizedUsersResponseBody) GoString() string {
	return s.String()
}

func (s *ListAuthorizedUsersResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAuthorizedUsersResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAuthorizedUsersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAuthorizedUsersResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListAuthorizedUsersResponseBody) GetUsers() []*ListAuthorizedUsersResponseBodyUsers {
	return s.Users
}

func (s *ListAuthorizedUsersResponseBody) SetPageNumber(v int32) *ListAuthorizedUsersResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListAuthorizedUsersResponseBody) SetPageSize(v int32) *ListAuthorizedUsersResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListAuthorizedUsersResponseBody) SetRequestId(v string) *ListAuthorizedUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAuthorizedUsersResponseBody) SetTotalCount(v int32) *ListAuthorizedUsersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListAuthorizedUsersResponseBody) SetUsers(v []*ListAuthorizedUsersResponseBodyUsers) *ListAuthorizedUsersResponseBody {
	s.Users = v
	return s
}

func (s *ListAuthorizedUsersResponseBody) Validate() error {
	if s.Users != nil {
		for _, item := range s.Users {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAuthorizedUsersResponseBodyUsers struct {
	// The user account type.
	//
	// - `simple`: convenience account.
	//
	// - `ad`: Active Directory (AD) domain account.
	//
	// example:
	//
	// simple
	AccountType *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	// The application ID specified in this query. This field is not returned if no application filter condition is specified.
	//
	// example:
	//
	// app-3jm9d0abc00example
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The delivery group ID to which the authorization relationship belongs. When querying cloud browsers, this is the browser group ID. When querying by set, this field is the primary delivery group ID of the set.
	//
	// example:
	//
	// big-3jm9d0abc00example
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// The delivery group set ID of this query. This field is returned when querying by set.
	//
	// example:
	//
	// set-3jm9d0abc00example
	AppInstanceGroupSetId *string `json:"AppInstanceGroupSetId,omitempty" xml:"AppInstanceGroupSetId,omitempty"`
	// The list of persistent session IDs authorized to the user. This field is returned when the authorization mode is `Session`.
	AppInstancePersistentIds []*string `json:"AppInstancePersistentIds,omitempty" xml:"AppInstancePersistentIds,omitempty" type:"Repeated"`
	// The authorization mode of the delivery group. Valid values:
	//
	// - `App`: Authorization by application.
	//
	// - `Session`: Authorization by persistent session.
	//
	// - `AppInstanceGroup`: Authorization by delivery group.
	//
	// example:
	//
	// AppInstanceGroup
	AuthMode *string `json:"AuthMode,omitempty" xml:"AuthMode,omitempty"`
	// The email address of the user. This field may not be returned if the email address is not available.
	//
	// example:
	//
	// alice@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The authorized username.
	//
	// example:
	//
	// alice
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// Indicates whether the query is not restricted to a specific application. Valid values:
	//
	// - `true`: No application filter condition is specified.
	//
	// - `false`: An application filter condition is specified.
	//
	// This field is determined by the query conditions and cannot be used alone to determine whether the user is authorized for all applications.
	//
	// example:
	//
	// true
	IsAuthAllApps *string `json:"IsAuthAllApps,omitempty" xml:"IsAuthAllApps,omitempty"`
	// The phone number of the user. This field may not be returned if the phone number is not available.
	//
	// example:
	//
	// 138****0000
	Phone *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
}

func (s ListAuthorizedUsersResponseBodyUsers) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorizedUsersResponseBodyUsers) GoString() string {
	return s.String()
}

func (s *ListAuthorizedUsersResponseBodyUsers) GetAccountType() *string {
	return s.AccountType
}

func (s *ListAuthorizedUsersResponseBodyUsers) GetAppId() *string {
	return s.AppId
}

func (s *ListAuthorizedUsersResponseBodyUsers) GetAppInstanceGroupId() *string {
	return s.AppInstanceGroupId
}

func (s *ListAuthorizedUsersResponseBodyUsers) GetAppInstanceGroupSetId() *string {
	return s.AppInstanceGroupSetId
}

func (s *ListAuthorizedUsersResponseBodyUsers) GetAppInstancePersistentIds() []*string {
	return s.AppInstancePersistentIds
}

func (s *ListAuthorizedUsersResponseBodyUsers) GetAuthMode() *string {
	return s.AuthMode
}

func (s *ListAuthorizedUsersResponseBodyUsers) GetEmail() *string {
	return s.Email
}

func (s *ListAuthorizedUsersResponseBodyUsers) GetEndUserId() *string {
	return s.EndUserId
}

func (s *ListAuthorizedUsersResponseBodyUsers) GetIsAuthAllApps() *string {
	return s.IsAuthAllApps
}

func (s *ListAuthorizedUsersResponseBodyUsers) GetPhone() *string {
	return s.Phone
}

func (s *ListAuthorizedUsersResponseBodyUsers) SetAccountType(v string) *ListAuthorizedUsersResponseBodyUsers {
	s.AccountType = &v
	return s
}

func (s *ListAuthorizedUsersResponseBodyUsers) SetAppId(v string) *ListAuthorizedUsersResponseBodyUsers {
	s.AppId = &v
	return s
}

func (s *ListAuthorizedUsersResponseBodyUsers) SetAppInstanceGroupId(v string) *ListAuthorizedUsersResponseBodyUsers {
	s.AppInstanceGroupId = &v
	return s
}

func (s *ListAuthorizedUsersResponseBodyUsers) SetAppInstanceGroupSetId(v string) *ListAuthorizedUsersResponseBodyUsers {
	s.AppInstanceGroupSetId = &v
	return s
}

func (s *ListAuthorizedUsersResponseBodyUsers) SetAppInstancePersistentIds(v []*string) *ListAuthorizedUsersResponseBodyUsers {
	s.AppInstancePersistentIds = v
	return s
}

func (s *ListAuthorizedUsersResponseBodyUsers) SetAuthMode(v string) *ListAuthorizedUsersResponseBodyUsers {
	s.AuthMode = &v
	return s
}

func (s *ListAuthorizedUsersResponseBodyUsers) SetEmail(v string) *ListAuthorizedUsersResponseBodyUsers {
	s.Email = &v
	return s
}

func (s *ListAuthorizedUsersResponseBodyUsers) SetEndUserId(v string) *ListAuthorizedUsersResponseBodyUsers {
	s.EndUserId = &v
	return s
}

func (s *ListAuthorizedUsersResponseBodyUsers) SetIsAuthAllApps(v string) *ListAuthorizedUsersResponseBodyUsers {
	s.IsAuthAllApps = &v
	return s
}

func (s *ListAuthorizedUsersResponseBodyUsers) SetPhone(v string) *ListAuthorizedUsersResponseBodyUsers {
	s.Phone = &v
	return s
}

func (s *ListAuthorizedUsersResponseBodyUsers) Validate() error {
	return dara.Validate(s)
}
