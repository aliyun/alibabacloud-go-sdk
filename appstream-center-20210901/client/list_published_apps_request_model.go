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
	// The application ID used for filtering. Substring matching is supported. You can specify a complete ID or a consecutive segment of it. If this parameter is not specified or is set to an empty string, filtering by application ID is not applied. If both this parameter and `AppName` are specified, both conditions must be met by the same application.
	//
	// example:
	//
	// ca-i87mycyn419nu****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The delivery group ID used for filtering. Substring matching is supported. You can specify a complete ID or a consecutive segment of it. If this parameter is not specified or is set to an empty string, filtering by delivery group ID is not applied. You can call the [ListAppInstanceGroup](~~ListAppInstanceGroup~~) operation to obtain delivery group IDs. This parameter can be used together with other filter conditions, and all conditions must be met simultaneously.
	//
	// example:
	//
	// aig-9ciijz60n4xsv****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// The delivery group name used for filtering. Substring matching is supported. For example, if you specify `OfficeApps`, delivery groups whose names contain this text are matched. If this parameter is not specified or is set to an empty string, filtering by delivery group name is not applied. If both this parameter and the delivery group ID are specified, both conditions must be met.
	//
	// example:
	//
	// OfficeApps
	AppInstanceGroupName *string `json:"AppInstanceGroupName,omitempty" xml:"AppInstanceGroupName,omitempty"`
	// The application name used for filtering. Substring matching is supported. For example, if you specify `OfficeApps`, applications whose names contain this text are matched. If this parameter is not specified or is set to an empty string, filtering by application name is not applied. If both this parameter and `AppId` are specified, both conditions must be met by the same application.
	//
	// example:
	//
	// OfficeApps
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The username to exclude. Exact username matching is used. For example, `alice`. When specified, applications that have been authorized to this user through [AuthorizeUsersForApp](~~AuthorizeUsersForApp~~) by application are not returned. This helps you find applications that can still be authorized to the user. If this parameter is not specified or is set to an empty string, no exclusion based on user authorization is applied.
	//
	// **Access permissions granted through delivery-group-level authorization or user groups are not evaluated by this condition.*	- The returned results cannot be treated as a complete list of applications that the user has no access to.
	//
	// example:
	//
	// alice
	ExcludeUserId *string `json:"ExcludeUserId,omitempty" xml:"ExcludeUserId,omitempty"`
	// The page number. This parameter is required. Start from page `1` and use this parameter together with `PageSize`. Keep other filter conditions unchanged when querying subsequent pages. If an invalid value is specified, the error code `InvalidParameter.PageNumber` is returned.
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
	// The product type. This parameter is required. The value is case-insensitive. Only applications in published delivery groups of the specified product type are returned. If an unrecognized value is specified, the error code `InvalidParameter.ProductType` is returned. Filtering and statistics related to per-application authorization (`ExcludeUserId` and `AuthorizedUserCount`) are primarily used in WUYING Cloud Application common scenarios.
	//
	// Valid values:
	//
	// - `CloudApp`: WUYING Cloud Application.
	//
	// - `CloudBrowser`: Cloud Browser.
	//
	// - `WuyingServer`: Enterprise Workstation.
	//
	// - `WuyingWorkstation`: Personal Edition Lingjou Container Workstation.
	//
	// - `WuyingWorkstationTeam`: Team Edition Lingjou Container Workstation.
	//
	// - `WuyingWorkstationBusiness`: Dedicated Edition Lingjou Container Workstation.
	//
	// - `AndroidCloud`: Cloud Phone.
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
