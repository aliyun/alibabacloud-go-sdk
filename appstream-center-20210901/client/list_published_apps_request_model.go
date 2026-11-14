// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPublishedAppsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ListPublishedAppsRequest
	GetAppId() *string
	SetAppInstanceGroupId(v string) *ListPublishedAppsRequest
	GetAppInstanceGroupId() *string
	SetAppInstanceGroupName(v string) *ListPublishedAppsRequest
	GetAppInstanceGroupName() *string
	SetAppName(v string) *ListPublishedAppsRequest
	GetAppName() *string
	SetExcludeUserId(v string) *ListPublishedAppsRequest
	GetExcludeUserId() *string
	SetPageNumber(v int32) *ListPublishedAppsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListPublishedAppsRequest
	GetPageSize() *int32
	SetProductType(v string) *ListPublishedAppsRequest
	GetProductType() *string
}

type ListPublishedAppsRequest struct {
	// The ID of the application used for filtering. Substring matching is supported. You can pass the complete ID or a continuous segment of it. If this parameter is not specified or an empty string is passed, filtering is not performed based on the application ID. If this parameter and `AppName` are specified at the same time, both conditions must be met by the same application.
	//
	// example:
	//
	// ca-i87mycyn419nu****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the delivery group used for filtering. Substring matching is supported. You can pass the complete ID or a continuous segment of it. If this parameter is not specified or an empty string is passed, filtering is not performed based on the delivery group ID. You can call the [ListAppInstanceGroup](~~ListAppInstanceGroup~~) operation to obtain the delivery group ID. This parameter can be used together with other filter conditions, and all conditions must be met.
	//
	// example:
	//
	// aig-9ciijz60n4xsv****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// The name of the delivery group used for filtering. Substring matching is supported. For example, if you pass `Office applications`, delivery groups whose names contain this text are matched. If this parameter is not specified or an empty string is passed, filtering is not performed based on the delivery group name. If this parameter and the delivery group ID are specified at the same time, both conditions must be met.
	//
	// example:
	//
	// Office applications
	AppInstanceGroupName *string `json:"AppInstanceGroupName,omitempty" xml:"AppInstanceGroupName,omitempty"`
	// The name of the application used for filtering. Substring matching is supported. For example, if you pass `Office applications`, applications whose names contain this text are matched. If this parameter is not specified or an empty string is passed, filtering is not performed based on the application name. If this parameter and `AppId` are specified at the same time, both conditions must be met by the same application.
	//
	// example:
	//
	// Office applications
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The username to be excluded. Exact matching is performed based on the full username, such as `alice`. If this parameter is specified, applications that have been authorized for this user on a per-application basis through [AuthorizeUsersForApp](~~AuthorizeUsersForApp~~) are not returned. This can be used to find applications that can still be authorized for the user. If this parameter is not specified or an empty string is passed, filtering is not performed based on user authorization.
	//
	// **Access permissions obtained through overall delivery group authorization or user groups are not evaluated by this condition.*	- Do not equate the returned results with a list of applications that the user has absolutely no access to.
	//
	// example:
	//
	// alice
	ExcludeUserId *string `json:"ExcludeUserId,omitempty" xml:"ExcludeUserId,omitempty"`
	// The page number. This parameter is required. Start the query from page `1` and use it in conjunction with `PageSize`. Keep other filter conditions unchanged when querying subsequent pages. If the value is invalid, the error code `InvalidParameter.PageNumber` is returned.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The maximum number of application records to return per page. This parameter is required. Valid values: `1` to `100`. If the value is out of range, the error code `InvalidParameter.PageSize` is returned.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The product type. This parameter is required and is not case-sensitive. Only applications in published delivery groups of this product type are returned. If an unrecognized value is passed, the error code `InvalidParameter.ProductType` is returned. Filtering and statistics related to per-application authorization (`ExcludeUserId` and `AuthorizedUserCount`) are primarily used in WUYING Cloud App scenarios.
	//
	// Valid values:
	//
	// - `CloudApp`: WUYING Cloud App.
	//
	// - `CloudBrowser`: Cloud browser.
	//
	// - `WuyingServer`: Enterprise Edition workstation.
	//
	// - `WuyingWorkstation`: Personal Edition Linggou container workstation.
	//
	// - `WuyingWorkstationTeam`: Team Edition Linggou container workstation.
	//
	// - `WuyingWorkstationBusiness`: Dedicated Edition Linggou container workstation.
	//
	// - `AndroidCloud`: Cloud phone.
	//
	// - `AIAgent`: AgentBay (AI agent).
	//
	// This parameter is required.
	//
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
}

func (s ListPublishedAppsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPublishedAppsRequest) GoString() string {
	return s.String()
}

func (s *ListPublishedAppsRequest) GetAppId() *string {
	return s.AppId
}

func (s *ListPublishedAppsRequest) GetAppInstanceGroupId() *string {
	return s.AppInstanceGroupId
}

func (s *ListPublishedAppsRequest) GetAppInstanceGroupName() *string {
	return s.AppInstanceGroupName
}

func (s *ListPublishedAppsRequest) GetAppName() *string {
	return s.AppName
}

func (s *ListPublishedAppsRequest) GetExcludeUserId() *string {
	return s.ExcludeUserId
}

func (s *ListPublishedAppsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListPublishedAppsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPublishedAppsRequest) GetProductType() *string {
	return s.ProductType
}

func (s *ListPublishedAppsRequest) SetAppId(v string) *ListPublishedAppsRequest {
	s.AppId = &v
	return s
}

func (s *ListPublishedAppsRequest) SetAppInstanceGroupId(v string) *ListPublishedAppsRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *ListPublishedAppsRequest) SetAppInstanceGroupName(v string) *ListPublishedAppsRequest {
	s.AppInstanceGroupName = &v
	return s
}

func (s *ListPublishedAppsRequest) SetAppName(v string) *ListPublishedAppsRequest {
	s.AppName = &v
	return s
}

func (s *ListPublishedAppsRequest) SetExcludeUserId(v string) *ListPublishedAppsRequest {
	s.ExcludeUserId = &v
	return s
}

func (s *ListPublishedAppsRequest) SetPageNumber(v int32) *ListPublishedAppsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListPublishedAppsRequest) SetPageSize(v int32) *ListPublishedAppsRequest {
	s.PageSize = &v
	return s
}

func (s *ListPublishedAppsRequest) SetProductType(v string) *ListPublishedAppsRequest {
	s.ProductType = &v
	return s
}

func (s *ListPublishedAppsRequest) Validate() error {
	return dara.Validate(s)
}
