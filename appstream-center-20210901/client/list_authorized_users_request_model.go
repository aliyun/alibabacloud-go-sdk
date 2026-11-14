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
	// The application ID. This parameter is used to filter users who are **specifically authorized for this application*	- (authorized through the [AuthorizeUsersForApp](~~AuthorizeUsersForApp~~) operation). This parameter applies to app instance groups whose authorization mode is `App`. You can obtain the application ID from the Apps list returned by the [GetAppInstanceGroup](~~GetAppInstanceGroup~~) operation.
	//
	// If this parameter is not specified, all authorized users under the app instance group are returned. This parameter is not supported when querying by app instance group set.
	//
	// example:
	//
	// ca-i87mycyn419nu****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The app instance group ID. You can call the [ListAppInstanceGroup](~~ListAppInstanceGroup~~) operation to obtain the ID. For cloud browser groups, specify the browser group ID returned by the [ListBrowserInstanceGroup](~~ListBrowserInstanceGroup~~) operation.
	//
	// **You must specify either this parameter or AppInstanceGroupSetId, but not both.**
	//
	// example:
	//
	// aig-9ciijz60n4xsv****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// The app instance group set ID.
	//
	// **You must specify either this parameter or AppInstanceGroupId, but not both.*	- When querying by set, do not specify AppId or AppInstancePersistentId. Otherwise, a parameter error is returned.
	//
	// example:
	//
	// set-3jm9d0abc00example
	AppInstanceGroupSetId *string `json:"AppInstanceGroupSetId,omitempty" xml:"AppInstanceGroupSetId,omitempty"`
	// The persistent session ID. This parameter is used to filter users who are granted this persistent session. This parameter applies to app instance groups whose authorization mode is `Session`. You can call the [ListPersistentAppInstances](~~ListPersistentAppInstances~~) operation to obtain the ID.
	//
	// If this parameter is specified, only the users granted this session are returned. However, the AppInstancePersistentIds response parameter still lists all persistent sessions granted to the user. This parameter is not supported when querying by app instance group set.
	//
	// example:
	//
	// p-0cc7s3mw2fg4j****
	AppInstancePersistentId *string `json:"AppInstancePersistentId,omitempty" xml:"AppInstancePersistentId,omitempty"`
	// The username. Performs **exact match*	- by username. If this parameter is not specified, filtering by full username is not applied. This parameter can be specified together with UserIdFuzzy, and both conditions must be met.
	//
	// example:
	//
	// alice
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The page number. Pages start from 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of records returned per page. Valid values: 1 to 100.
	//
	// When the authorization mode is `App` or `AppInstanceGroup`, pagination is based on authorization records. Multiple authorization records for the same user are merged into one user in the response. Therefore, the actual number of users returned on the current page may be less than this value.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The product type. The value must be consistent with the product type of the queried app instance group or app instance group set. If they are inconsistent, a resource not found error code is returned.
	//
	// Valid values:
	//
	// - CloudApp: Wuying Cloud App
	//
	// - CloudBrowser: Cloud Browser
	//
	// - WuyingServer: Enterprise Edition workstation
	//
	// - WuyingWorkstation: Personal Edition Linggou container workstation
	//
	// - WuyingWorkstationTeam: Team Edition Linggou container workstation
	//
	// - WuyingWorkstationBusiness: Dedicated Edition Linggou container workstation
	//
	// - AndroidCloud: Cloud Phone
	//
	// - AIAgent: AgentBay (AI agent)
	//
	// This parameter is required.
	//
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The username keyword. Performs **fuzzy match*	- by username. A hit occurs if the username contains this keyword. For example, if you specify `ali`, both `alice` and `ali.wang` are matched. If this parameter is not specified, filtering by keyword is not applied.
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
