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
	// The current page number, which is the same as the PageNumber request parameter.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of records per page, which is the same as the PageSize request parameter.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 5C1A4F2D-713A-5C98-8AF6-1B5D0868****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of records that match the query conditions. Use this value to determine whether to continue paging.
	//
	// - When the authorization mode is `App` or `AppInstanceGroup`, this is the number of authorization records. If the same user has multiple authorization records, the user is counted multiple times. Therefore, this value may be greater than the actual number of users.
	//
	// - When the authorization mode is `Session`, this is the deduplicated user count.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The list of authorized users on the current page. Multiple authorization records for the same user are merged into a single entry. An empty list is returned if no authorized users match the conditions.
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
	// The account type of the user. Valid values:
	//
	// - simple: Convenience account.
	//
	// - ad: Active Directory (AD) domain account, which originates from an enterprise AD domain.
	//
	// example:
	//
	// simple
	AccountType *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	// The application ID. Returned only when AppId is specified in the request. The value is the same as the request parameter. Not returned if AppId is not specified or when querying by delivery group set.
	//
	// example:
	//
	// ca-i87mycyn419nu****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The delivery group ID associated with the user\\"s authorization relationship. When querying by delivery group, this value is the same as the request parameter. When querying by delivery group set, this value is the primary delivery group ID of the set.
	//
	// example:
	//
	// aig-9ciijz60n4xsv****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// The delivery group set ID. Returned only when querying by delivery group set. The value is the same as the AppInstanceGroupSetId request parameter.
	//
	// example:
	//
	// set-3jm9d0abc00example
	AppInstanceGroupSetId *string `json:"AppInstanceGroupSetId,omitempty" xml:"AppInstanceGroupSetId,omitempty"`
	// The list of persistent session IDs granted to the user. Returned only when the delivery group authorization mode (AuthMode) is `Session`. This list is not affected by the AppInstancePersistentId request parameter and always includes all persistent sessions granted to the user.
	AppInstancePersistentIds []*string `json:"AppInstancePersistentIds,omitempty" xml:"AppInstancePersistentIds,omitempty" type:"Repeated"`
	// The authorization mode of the delivery group, which determines the scope of results returned by this operation. Valid values:
	//
	// - App: Application-level authorization. Applications within the delivery group are authorized to users without restricting which sessions the users can use.
	//
	// - Session: Session-level authorization. Persistent sessions within the delivery group are authorized to users without restricting which applications the users can use. In this case, AppInstancePersistentIds returns the persistent sessions granted to the user.
	//
	// - AppInstanceGroup: Delivery group-level authorization. The entire delivery group is authorized to users, allowing them to open any application using any session within the delivery group.
	//
	// When querying by delivery group set, the authorization mode of the primary delivery group in the set is returned.
	//
	// example:
	//
	// AppInstanceGroup
	AuthMode *string `json:"AuthMode,omitempty" xml:"AuthMode,omitempty"`
	// The email address of the user. Returned only when the account information of the user can be retrieved.
	//
	// example:
	//
	// alice@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The username. To remove authorization, pass this value to the UnAuthorizeUserIds parameter of the [AuthorizeInstanceGroup](~~AuthorizeInstanceGroup~~) or [AuthorizeUsersForApp](~~AuthorizeUsersForApp~~) operation.
	//
	// example:
	//
	// alice
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// Indicates whether the query is not restricted to a specific application. Valid values:
	//
	// - true: AppId is not specified in the request. All authorized users under the delivery group are returned.
	//
	// - false: AppId is specified in the request. Only users authorized for that specific application are returned.
	//
	// > This field is determined by whether the AppId request parameter is specified. It does not reflect the actual scope of applications authorized to the user and cannot be used to determine whether the user is authorized for all applications.
	//
	// example:
	//
	// true
	IsAuthAllApps *string `json:"IsAuthAllApps,omitempty" xml:"IsAuthAllApps,omitempty"`
	// The phone number of the user. Returned only when the account information of the user can be retrieved.
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
