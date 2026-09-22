// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAuthorizedUsersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ListAuthorizedUsersRequest
	GetAppId() *string
	SetAppInstanceGroupId(v string) *ListAuthorizedUsersRequest
	GetAppInstanceGroupId() *string
	SetAppInstanceGroupSetId(v string) *ListAuthorizedUsersRequest
	GetAppInstanceGroupSetId() *string
	SetAppInstancePersistentId(v string) *ListAuthorizedUsersRequest
	GetAppInstancePersistentId() *string
	SetEndUserId(v string) *ListAuthorizedUsersRequest
	GetEndUserId() *string
	SetPageNumber(v int32) *ListAuthorizedUsersRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListAuthorizedUsersRequest
	GetPageSize() *int32
	SetProductType(v string) *ListAuthorizedUsersRequest
	GetProductType() *string
	SetUserIdFuzzy(v string) *ListAuthorizedUsersRequest
	GetUserIdFuzzy() *string
}

type ListAuthorizedUsersRequest struct {
	// The application ID. Specifies the application to filter users who are **authorized for that specific application*	- (authorized through the [AuthorizeUsersForApp](~~AuthorizeUsersForApp~~) operation). This parameter applies to delivery groups with the `App` authorization mode. Obtain the application ID from the Apps list returned by the [GetAppInstanceGroup](~~GetAppInstanceGroup~~) operation.
	//
	// If not specified, all authorized users under the delivery group are returned. This parameter is not supported when querying by delivery group set.
	//
	// example:
	//
	// ca-i87mycyn419nu****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The delivery group ID. Call the [ListAppInstanceGroup](~~ListAppInstanceGroup~~) operation to obtain this value. For cloud browser groups, specify the browser group ID returned by the [ListBrowserInstanceGroup](~~ListBrowserInstanceGroup~~) operation.
	//
	// **Exactly one of this parameter and AppInstanceGroupSetId must be specified.**
	//
	// example:
	//
	// aig-9ciijz60n4xsv****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// The delivery group set ID.
	//
	// **Exactly one of this parameter and AppInstanceGroupId must be specified.*	- When querying by set, do not specify AppId or AppInstancePersistentId. Otherwise, a parameter error is returned.
	//
	// example:
	//
	// set-3jm9d0abc00example
	AppInstanceGroupSetId *string `json:"AppInstanceGroupSetId,omitempty" xml:"AppInstanceGroupSetId,omitempty"`
	// The persistent session ID. Specifies the persistent session to filter users who are granted that session. This parameter applies to delivery groups with the `Session` authorization mode. Call the [ListPersistentAppInstances](~~ListPersistentAppInstances~~) operation to obtain this value.
	//
	// If specified, only users granted that session are returned. However, the response parameter AppInstancePersistentIds still lists all persistent sessions granted to each user. This parameter is not supported when querying by delivery group set.
	//
	// example:
	//
	// p-0cc7s3mw2fg4j****
	AppInstancePersistentId *string `json:"AppInstancePersistentId,omitempty" xml:"AppInstancePersistentId,omitempty"`
	// The username for **exact matching**. If not specified, no filtering by exact username is applied. Can be specified together with UserIdFuzzy, in which case both conditions must be met.
	//
	// example:
	//
	// alice
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The page number, starting from 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of records per page. Valid values: 1 to 100.
	//
	// When the authorization mode is `App` or `AppInstanceGroup`, pagination is based on authorization records. Multiple authorization records for the same user are merged into a single user entry. Therefore, the actual number of users returned on the current page may be less than this value.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The product type. The value must match the product type of the queried delivery group or delivery group set. If the value does not match, a resource-not-found error code is returned.
	//
	// Valid values:
	//
	// - CloudApp: Wuying Cloud Application.
	//
	// - CloudBrowser: Cloud Browser.
	//
	// - WuyingServer: Enterprise Edition Workstation.
	//
	// - WuyingWorkstation: Personal Edition Linggou Container Workstation.
	//
	// - WuyingWorkstationTeam: Linggou Team Edition Container Workstation.
	//
	// - WuyingWorkstationBusiness: Linggou Dedicated Edition Container Workstation.
	//
	// - AndroidCloud: Cloud Phone.
	//
	// - AIAgent: AgentBay (AI agent).
	//
	// This parameter is required.
	//
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The username keyword for **fuzzy matching**. A match occurs if the username contains this keyword. For example, if you specify `ali`, both `alice` and `ali.wang` are returned. If not specified, no keyword-based filtering is applied.
	//
	// example:
	//
	// ali
	UserIdFuzzy *string `json:"UserIdFuzzy,omitempty" xml:"UserIdFuzzy,omitempty"`
}

func (s ListAuthorizedUsersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorizedUsersRequest) GoString() string {
	return s.String()
}

func (s *ListAuthorizedUsersRequest) GetAppId() *string {
	return s.AppId
}

func (s *ListAuthorizedUsersRequest) GetAppInstanceGroupId() *string {
	return s.AppInstanceGroupId
}

func (s *ListAuthorizedUsersRequest) GetAppInstanceGroupSetId() *string {
	return s.AppInstanceGroupSetId
}

func (s *ListAuthorizedUsersRequest) GetAppInstancePersistentId() *string {
	return s.AppInstancePersistentId
}

func (s *ListAuthorizedUsersRequest) GetEndUserId() *string {
	return s.EndUserId
}

func (s *ListAuthorizedUsersRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAuthorizedUsersRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAuthorizedUsersRequest) GetProductType() *string {
	return s.ProductType
}

func (s *ListAuthorizedUsersRequest) GetUserIdFuzzy() *string {
	return s.UserIdFuzzy
}

func (s *ListAuthorizedUsersRequest) SetAppId(v string) *ListAuthorizedUsersRequest {
	s.AppId = &v
	return s
}

func (s *ListAuthorizedUsersRequest) SetAppInstanceGroupId(v string) *ListAuthorizedUsersRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *ListAuthorizedUsersRequest) SetAppInstanceGroupSetId(v string) *ListAuthorizedUsersRequest {
	s.AppInstanceGroupSetId = &v
	return s
}

func (s *ListAuthorizedUsersRequest) SetAppInstancePersistentId(v string) *ListAuthorizedUsersRequest {
	s.AppInstancePersistentId = &v
	return s
}

func (s *ListAuthorizedUsersRequest) SetEndUserId(v string) *ListAuthorizedUsersRequest {
	s.EndUserId = &v
	return s
}

func (s *ListAuthorizedUsersRequest) SetPageNumber(v int32) *ListAuthorizedUsersRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAuthorizedUsersRequest) SetPageSize(v int32) *ListAuthorizedUsersRequest {
	s.PageSize = &v
	return s
}

func (s *ListAuthorizedUsersRequest) SetProductType(v string) *ListAuthorizedUsersRequest {
	s.ProductType = &v
	return s
}

func (s *ListAuthorizedUsersRequest) SetUserIdFuzzy(v string) *ListAuthorizedUsersRequest {
	s.UserIdFuzzy = &v
	return s
}

func (s *ListAuthorizedUsersRequest) Validate() error {
	return dara.Validate(s)
}
